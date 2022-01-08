// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package backupccl

import (
	"context"
	fmt "fmt"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

func BenchmarkCoverageChecks(b *testing.B) {
	r, _ := randutil.NewPseudoRand()

	for _, numBackups := range []int{1, 7, 24, 24 * 4} {
		b.Run(fmt.Sprintf("numBackups=%d", numBackups), func(b *testing.B) {
			for _, numSpans := range []int{10, 20, 100} {
				b.Run(fmt.Sprintf("numSpans=%d", numSpans), func(b *testing.B) {
					for _, baseFiles := range []int{0, 10, 100, 1000, 10000} {
						b.Run(fmt.Sprintf("numFiles=%d", baseFiles), func(b *testing.B) {
							b.StopTimer()
							ctx := context.Background()
							backups := MockBackupChain(numBackups, numSpans, baseFiles, r)
							b.ResetTimer()

							b.Run("checkCoverage", func(b *testing.B) {
								b.ResetTimer()
								for i := 0; i < b.N; i++ {
									if err := checkCoverage(ctx, backups[numBackups-1].Spans, backups); err != nil {
										b.Fatal(err)
									}
								}
							})
							b.Run("makeImportSpans", func(b *testing.B) {
								b.ResetTimer()
								for i := 0; i < b.N; i++ {
									_, ts, err := makeImportSpans(backups[numBackups-1].Spans, backups, nil, nil, errOnMissingRange)
									if err != nil {
										b.Fatal(err)
									}
									if got, expected := ts, backups[len(backups)-1].EndTime; !got.Equal(expected) {
										b.Fatal(expected, got)
									}
								}
							})
						})
					}
				})
			}
		})
	}
}

func BenchmarkRestoreEntryCover(b *testing.B) {
	r, _ := randutil.NewPseudoRand()

	for _, numBackups := range []int{1, 2, 24, 24 * 4} {
		b.Run(fmt.Sprintf("numBackups=%d", numBackups), func(b *testing.B) {
			for _, baseFiles := range []int{0, 100, 10000} {
				b.Run(fmt.Sprintf("numFiles=%d", baseFiles), func(b *testing.B) {
					for _, numSpans := range []int{10, 100} {
						b.Run(fmt.Sprintf("numSpans=%d", numSpans), func(b *testing.B) {
							b.StopTimer()
							ctx := context.Background()
							backups := MockBackupChain(numBackups, numSpans, baseFiles, r)
							b.ResetTimer()
							b.Run("simple", func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									if err := checkCoverage(ctx, backups[numBackups-1].Spans, backups); err != nil {
										b.Fatal(err)
									}
									cov := makeSimpleImportSpans(backups[numBackups-1].Spans, backups, nil, nil)
									b.ReportMetric(float64(len(cov)), "coverSize")
								}
							})
							b.Run("coveringMerge", func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									cov, _, err := makeImportSpans(backups[numBackups-1].Spans, backups, nil, nil, errOnMissingRange)
									if err != nil {
										b.Fatal(err)
									}
									b.ReportMetric(float64(len(cov)), "coverSize")
								}
							})
						})
					}
				})
			}
		})
	}
}
