// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecagg

import (
	"unsafe"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
)

// Remove unused warning.
var _ = colexecerror.InternalError

func newBoolAndHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolAndHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolAndHashAgg struct {
	unorderedAggregateFuncBase
	curAgg bool
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &boolAndHashAgg{}

func (a *boolAndHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, startIdx, endIdx int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		{
			sel = sel[startIdx:endIdx]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for _, i := range sel {

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg && col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsageAfterAllocation(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolAndHashAgg) Flush(outputIdx int) {
	col := a.vec.Bool()
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		col[outputIdx] = a.curAgg
	}
}

func (a *boolAndHashAgg) Reset() {
	a.curAgg = true
	a.foundNonNullForCurrentGroup = false
}

type boolAndHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolAndHashAgg
}

var _ aggregateFuncAlloc = &boolAndHashAggAlloc{}

const sizeOfBoolAndHashAgg = int64(unsafe.Sizeof(boolAndHashAgg{}))
const boolAndHashAggSliceOverhead = int64(unsafe.Sizeof([]boolAndHashAgg{}))

func (a *boolAndHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolAndHashAggSliceOverhead + sizeOfBoolAndHashAgg*a.allocSize)
		a.aggFuncs = make([]boolAndHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

func newBoolOrHashAggAlloc(
	allocator *colmem.Allocator, allocSize int64,
) aggregateFuncAlloc {
	return &boolOrHashAggAlloc{aggAllocBase: aggAllocBase{
		allocator: allocator,
		allocSize: allocSize,
	}}
}

type boolOrHashAgg struct {
	unorderedAggregateFuncBase
	curAgg bool
	// foundNonNullForCurrentGroup tracks if we have seen any non-null values
	// for the group that is currently being aggregated.
	foundNonNullForCurrentGroup bool
}

var _ AggregateFunc = &boolOrHashAgg{}

func (a *boolOrHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, startIdx, endIdx int, sel []int,
) {
	var oldCurAggSize uintptr
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Bool(), vec.Nulls()
	a.allocator.PerformOperation([]coldata.Vec{a.vec}, func() {
		{
			sel = sel[startIdx:endIdx]
			if nulls.MaybeHasNulls() {
				for _, i := range sel {

					var isNull bool
					isNull = nulls.NullAt(i)
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			} else {
				for _, i := range sel {

					var isNull bool
					isNull = false
					if !isNull {
						a.curAgg = a.curAgg || col[i]
						a.foundNonNullForCurrentGroup = true
					}

				}
			}
		}
	},
	)
	var newCurAggSize uintptr
	if newCurAggSize != oldCurAggSize {
		a.allocator.AdjustMemoryUsageAfterAllocation(int64(newCurAggSize - oldCurAggSize))
	}
}

func (a *boolOrHashAgg) Flush(outputIdx int) {
	col := a.vec.Bool()
	if !a.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		col[outputIdx] = a.curAgg
	}
}

func (a *boolOrHashAgg) Reset() {
	a.curAgg = false
	a.foundNonNullForCurrentGroup = false
}

type boolOrHashAggAlloc struct {
	aggAllocBase
	aggFuncs []boolOrHashAgg
}

var _ aggregateFuncAlloc = &boolOrHashAggAlloc{}

const sizeOfBoolOrHashAgg = int64(unsafe.Sizeof(boolOrHashAgg{}))
const boolOrHashAggSliceOverhead = int64(unsafe.Sizeof([]boolOrHashAgg{}))

func (a *boolOrHashAggAlloc) newAggFunc() AggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(boolOrHashAggSliceOverhead + sizeOfBoolOrHashAgg*a.allocSize)
		a.aggFuncs = make([]boolOrHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	f.allocator = a.allocator
	f.Reset()
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
