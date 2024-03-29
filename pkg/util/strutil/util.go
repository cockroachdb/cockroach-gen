// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package strutil

import "strconv"

// AppendInt appends the decimal form of x to b and returns the result.
// If the decimal form is shorter than width, the result is padded with leading 0's.
// If the decimal is longer than width, returns formatted decimal without
// any truncation.
func AppendInt(b []byte, x int, width int) []byte {
	if x < 0 {
		width--
		x = -x
		b = append(b, '-')
	}

	var scratch [16]byte
	xb := strconv.AppendInt(scratch[:0], int64(x), 10)

	// Add 0-padding.
	for w := len(xb); w < width; w++ {
		b = append(b, '0')
	}
	return append(b, xb...)
}
