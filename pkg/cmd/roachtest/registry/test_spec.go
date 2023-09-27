// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package registry

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/cluster"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/spec"
	"github.com/cockroachdb/cockroach/pkg/cmd/roachtest/test"
)

// LibGEOS is a list of native libraries for libgeos.
var LibGEOS = []string{"libgeos", "libgeos_c"}

// TestSpec is a spec for a roachtest.
type TestSpec struct {
	Skip string // if non-empty, test will be skipped
	// When Skip is set, this can contain more text to be printed in the logs
	// after the "--- SKIP" line.
	SkipDetails string

	Name string
	// Owner is the name of the team responsible for signing off on failures of
	// this test that happen in the release process. This must be one of a limited
	// set of values (the keys in the roachtestTeams map).
	Owner Owner
	// The maximum duration the test is allowed to run before it is considered
	// failed. If not specified, the default timeout is 10m before the test's
	// associated cluster expires. The timeout is always truncated to 10m before
	// the test's cluster expires.
	Timeout time.Duration
	// Denotes whether the test is a roachperf benchmark. If true, the test is expected, but not required, to produce
	// artifacts in Test.PerfArtifactsDir(), which get exported to the roachperf dashboard
	// (see getPerfArtifacts() in test_runner.go).
	// N.B. performance tests may choose to _assert_ on latency, throughput, or some other perf. metric, _without_
	// exporting artifacts to the dashboard. E.g., see 'registerKVContention' and 'verifyTxnPerSecond'.
	// N.B. performance tests may have different requirements than correctness tests, e.g., machine type/architecture.
	// Thus, they must be opted into explicitly via this field.
	Benchmark bool
	// Tags is a set of tags associated with the test that allow grouping
	// tests. If no tags are specified, the set ["default"] is automatically
	// given.
	Tags map[string]struct{}
	// Cluster provides the specification for the cluster to use for the test.
	Cluster spec.ClusterSpec
	// NativeLibs specifies the native libraries required to be present on
	// the cluster during execution.
	NativeLibs []string

	// UseIOBarrier controls the local-ssd-no-ext4-barrier flag passed to
	// roachprod when creating a cluster. If set, the flag is not passed, and so
	// you get durable writes. If not set (the default!), the filesystem is
	// mounted without the barrier.
	//
	// The default (false) is chosen because it the no-barrier option is needed
	// explicitly by some tests (particularly benchmarks, ironically, since they'd
	// rather measure other things than I/O) and the vast majority of other tests
	// don't care - there's no durability across machine crashes that roachtests
	// care about.
	UseIOBarrier bool

	// NonReleaseBlocker indicates that a test failure should not be tagged as a
	// release blocker. Use this for tests that are not yet stable but should
	// still be run regularly.
	NonReleaseBlocker bool

	// RequiresLicense indicates that the test requires an
	// enterprise license to run correctly. Use this to ensure
	// tests will fail-early if COCKROACH_DEV_LICENSE is not set
	// in the environment.
	RequiresLicense bool

	// EncryptionSupport encodes to what extent tests supports
	// encryption-at-rest. See the EncryptionSupport type for details.
	// Encryption support is opt-in -- i.e., if the TestSpec does not
	// pass a value to this field, it will be assumed that the test
	// cannot be run with encryption enabled.
	EncryptionSupport EncryptionSupport

	// Run is the test function.
	Run func(ctx context.Context, t test.Test, c cluster.Cluster)
}

// MatchType is the type of match a file has to a TestFilter.
type MatchType int

const (
	// Matched means that the file passes the filter and the tags.
	Matched MatchType = iota
	// FailedFilter means that the file fails the filter.
	FailedFilter
	// FailedTags means that the file passed the filter but failed the tags
	// match.
	FailedTags
)

// Match returns Matched if the filter matches the test. If the filter does
// not match the test because the tag filter does not match, the test is
// marked as FailedTags.
func (t *TestSpec) Match(filter *TestFilter) MatchType {
	if !filter.Name.MatchString(t.Name) {
		return FailedFilter
	}

	if len(filter.Tags) == 0 {
		return Matched
	}

	for tag := range filter.Tags {
		// If the tag is a single CSV e.g. "foo,bar,baz", we match all the tags
		if matchesAll(t.Tags, strings.Split(tag, ",")) {
			return Matched
		}
	}

	return FailedTags
}

func matchesAll(testTags map[string]struct{}, filterTags []string) bool {
	for _, tag := range filterTags {
		negate := false
		if tag[0] == '!' {
			negate = true
			tag = tag[1:]
		}
		_, tagExists := testTags[tag]

		if negate == tagExists {
			return false
		}
	}
	return true
}

// Tags returns a set of strings.
func Tags(values ...string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, s := range values {
		set[s] = struct{}{}
	}
	return set
}
