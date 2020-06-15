// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"unsafe"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
)

type sumInt16Agg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt16Agg{}

const sizeOfSumInt16Agg = int64(unsafe.Sizeof(sumInt16Agg{}))

func (a *sumInt16Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumInt16Agg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumInt16Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt16Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumInt16Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int16(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt16Agg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumInt16Agg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumInt16AggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt16Agg
}

var _ aggregateFuncAlloc = &sumInt16AggAlloc{}

func (a *sumInt16AggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt16Agg * a.allocSize)
		a.aggFuncs = make([]sumInt16Agg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumInt32Agg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt32Agg{}

const sizeOfSumInt32Agg = int64(unsafe.Sizeof(sumInt32Agg{}))

func (a *sumInt32Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumInt32Agg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumInt32Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt32Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumInt32Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int32(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt32Agg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumInt32Agg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumInt32AggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt32Agg
}

var _ aggregateFuncAlloc = &sumInt32AggAlloc{}

func (a *sumInt32AggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt32Agg * a.allocSize)
		a.aggFuncs = make([]sumInt32Agg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumInt64Agg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt64Agg{}

const sizeOfSumInt64Agg = int64(unsafe.Sizeof(sumInt64Agg{}))

func (a *sumInt64Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumInt64Agg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumInt64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumInt64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumInt64Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Int64(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt64Agg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumInt64Agg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumInt64AggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt64Agg
}

var _ aggregateFuncAlloc = &sumInt64AggAlloc{}

func (a *sumInt64AggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt64Agg * a.allocSize)
		a.aggFuncs = make([]sumInt64Agg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumDecimalAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumDecimalAgg{}

const sizeOfSumDecimalAgg = int64(unsafe.Sizeof(sumDecimalAgg{}))

func (a *sumDecimalAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumDecimalAgg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumDecimalAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumDecimalAgg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumDecimalAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Decimal(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumDecimalAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumDecimalAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumDecimalAggAlloc struct {
	aggAllocBase
	aggFuncs []sumDecimalAgg
}

var _ aggregateFuncAlloc = &sumDecimalAggAlloc{}

func (a *sumDecimalAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumDecimalAgg * a.allocSize)
		a.aggFuncs = make([]sumDecimalAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumFloat64Agg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg float64
		// vec points to the output vector we are updating.
		vec []float64
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumFloat64Agg{}

const sizeOfSumFloat64Agg = int64(unsafe.Sizeof(sumFloat64Agg{}))

func (a *sumFloat64Agg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Float64()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumFloat64Agg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumFloat64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumFloat64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumFloat64Agg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Float64(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroFloat64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroFloat64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroFloat64Value

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroFloat64Value

				}
				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumFloat64Agg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumFloat64Agg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumFloat64AggAlloc struct {
	aggAllocBase
	aggFuncs []sumFloat64Agg
}

var _ aggregateFuncAlloc = &sumFloat64AggAlloc{}

func (a *sumFloat64AggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumFloat64Agg * a.allocSize)
		a.aggFuncs = make([]sumFloat64Agg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntervalAgg struct {
	groups  []bool
	scratch struct {
		curIdx int
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg duration.Duration
		// vec points to the output vector we are updating.
		vec []duration.Duration
		// nulls points to the output null vector that we are updating.
		nulls *coldata.Nulls
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumIntervalAgg{}

const sizeOfSumIntervalAgg = int64(unsafe.Sizeof(sumIntervalAgg{}))

func (a *sumIntervalAgg) Init(groups []bool, v coldata.Vec) {
	a.groups = groups
	a.scratch.vec = v.Interval()
	a.scratch.nulls = v.Nulls()
	a.Reset()
}

func (a *sumIntervalAgg) Reset() {
	a.scratch.curIdx = -1
	a.scratch.foundNonNullForCurrentGroup = false
	a.scratch.nulls.UnsetNulls()
}

func (a *sumIntervalAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *sumIntervalAgg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		a.scratch.nulls.UnsetNullsAfter(idx + 1)
	}
}

func (a *sumIntervalAgg) Compute(b coldata.Batch, inputIdxs []uint32) {
	inputLen := b.Length()
	vec, sel := b.ColVec(int(inputIdxs[0])), b.Selection()
	col, nulls := vec.Interval(), vec.Nulls()
	if nulls.MaybeHasNulls() {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroIntervalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroIntervalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}
				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		if sel != nil {
			sel = sel[:inputLen]
			for _, i := range sel {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroIntervalValue

				}
				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			col = col[:inputLen]
			for i := range col {

				if a.groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null. If
					// a.scratch.curIdx is negative, it means that this is the first group.
					if a.scratch.curIdx >= 0 {
						if !a.scratch.foundNonNullForCurrentGroup {
							a.scratch.nulls.SetNull(a.scratch.curIdx)
						} else {
							a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
						}
					}
					a.scratch.curIdx++
					a.scratch.curAgg = zeroIntervalValue

				}
				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntervalAgg) Flush() {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.scratch.nulls.SetNull(a.scratch.curIdx)
	} else {
		a.scratch.vec[a.scratch.curIdx] = a.scratch.curAgg
	}
	a.scratch.curIdx++
}

func (a *sumIntervalAgg) HandleEmptyInputScalar() {
	a.scratch.nulls.SetNull(0)
}

type sumIntervalAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntervalAgg
}

var _ aggregateFuncAlloc = &sumIntervalAggAlloc{}

func (a *sumIntervalAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumIntervalAgg * a.allocSize)
		a.aggFuncs = make([]sumIntervalAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
