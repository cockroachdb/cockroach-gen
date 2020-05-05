// Code generated by execgen; DO NOT EDIT.
// Copyright 2020 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"bytes"
	"fmt"
	"math"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
)

// Remove unused warning.
var _ = execgen.UNSAFEGET

// populateSels populates intermediate selection vectors (stored in
// op.scratch.sels) for each hash code present in b. hashBuffer must contain
// the hash codes for all of the tuples in b.
func (op *hashAggregator) populateSels(b coldata.Batch, hashBuffer []uint64) {
	// Note: we don't need to reset any of the slices in op.scratch.sels since
	// they all are of zero length here (see the comment for op.scratch.sels
	// for context).
	op.scratch.hashCodeForSelsSlot = op.scratch.hashCodeForSelsSlot[:0]
	if batchSelection := b.Selection(); batchSelection != nil {
		for selIdx, hashCode := range hashBuffer {
			selsSlot := -1
			for slot, hash := range op.scratch.hashCodeForSelsSlot {
				if hash == hashCode {
					// We have already seen a tuple with the same hashCode
					// previously, so we will append into the same sels slot.
					selsSlot = slot
					break
				}
			}
			if selsSlot < 0 {
				// This is the first tuple in hashBuffer with this hashCode, so we
				// will add this tuple to the next available sels slot.
				selsSlot = len(op.scratch.hashCodeForSelsSlot)
				op.scratch.hashCodeForSelsSlot = append(op.scratch.hashCodeForSelsSlot, hashCode)
			}
			op.scratch.sels[selsSlot] = append(op.scratch.sels[selsSlot], batchSelection[selIdx])
		}
	} else {
		for selIdx, hashCode := range hashBuffer {
			selsSlot := -1
			for slot, hash := range op.scratch.hashCodeForSelsSlot {
				if hash == hashCode {
					// We have already seen a tuple with the same hashCode
					// previously, so we will append into the same sels slot.
					selsSlot = slot
					break
				}
			}
			if selsSlot < 0 {
				// This is the first tuple in hashBuffer with this hashCode, so we
				// will add this tuple to the next available sels slot.
				selsSlot = len(op.scratch.hashCodeForSelsSlot)
				op.scratch.hashCodeForSelsSlot = append(op.scratch.hashCodeForSelsSlot, hashCode)
			}
			op.scratch.sels[selsSlot] = append(op.scratch.sels[selsSlot], selIdx)
		}
	}
}

// match takes a selection vector and compares it against the values of the key
// of its aggregation function. It returns a selection vector representing the
// unmatched tuples and a boolean to indicate whether or not there are any
// matching tuples. It directly writes the result of matched tuples into the
// selection vector of 'b' and sets the length of the batch to the number of
// matching tuples. match also takes a diff boolean slice for internal use.
// This slice need to be allocated to be at at least as big as sel and set to
// all false. diff will be reset to all false when match returns. This is to
// avoid additional slice allocation.
// - firstDefiniteMatch indicates whether we know that tuple with index sel[0]
//   matches the key of the aggregation function and whether we can short
//   circuit probing that tuple.
// NOTE: the return vector will reuse the memory allocated for the selection
//       vector.
func (v hashAggFuncs) match(
	sel []int,
	b coldata.Batch,
	keyCols []uint32,
	keyTypes []*types.T,
	keyCanonicalTypeFamilies []types.Family,
	keyMapping coldata.Batch,
	diff []bool,
	firstDefiniteMatch bool,
) (bool, []int) {
	// We want to directly write to the selection vector to avoid extra
	// allocation.
	b.SetSelection(true)
	matched := b.Selection()[:0]

	aggKeyIdx := v.keyIdx

	if firstDefiniteMatch {
		matched = append(matched, sel[0])
		sel = sel[1:]
		diff = diff[:len(diff)-1]
	}

	if len(sel) > 0 {
		for keyIdx, colIdx := range keyCols {
			lhs := keyMapping.ColVec(keyIdx)
			lhsHasNull := lhs.MaybeHasNulls()

			rhs := b.ColVec(int(colIdx))
			rhsHasNull := rhs.MaybeHasNulls()

			switch keyCanonicalTypeFamilies[keyIdx] {
			case types.BoolFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Bool()
					rhsCol := rhs.Bool()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if !lhsVal && rhsVal {
										cmpResult = -1
									} else if lhsVal && !rhsVal {
										cmpResult = 1
									} else {
										cmpResult = 0
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if !lhsVal && rhsVal {
										cmpResult = -1
									} else if lhsVal && !rhsVal {
										cmpResult = 1
									} else {
										cmpResult = 0
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if !lhsVal && rhsVal {
										cmpResult = -1
									} else if lhsVal && !rhsVal {
										cmpResult = 1
									} else {
										cmpResult = 0
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if !lhsVal && rhsVal {
										cmpResult = -1
									} else if lhsVal && !rhsVal {
										cmpResult = 1
									} else {
										cmpResult = 0
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.BytesFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Bytes()
					rhsCol := rhs.Bytes()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol.Get(aggKeyIdx)

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol.Get(rowIdx)

								var cmp bool

								{
									var cmpResult int
									cmpResult = bytes.Compare(lhsVal, rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol.Get(aggKeyIdx)

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol.Get(rowIdx)

								var cmp bool

								{
									var cmpResult int
									cmpResult = bytes.Compare(lhsVal, rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol.Get(aggKeyIdx)

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol.Get(rowIdx)

								var cmp bool

								{
									var cmpResult int
									cmpResult = bytes.Compare(lhsVal, rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol.Get(aggKeyIdx)

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol.Get(rowIdx)

								var cmp bool

								{
									var cmpResult int
									cmpResult = bytes.Compare(lhsVal, rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.DecimalFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Decimal()
					rhsCol := rhs.Decimal()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = tree.CompareDecimals(&lhsVal, &rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.IntFamily:
				switch keyTypes[keyIdx].Width() {
				case 16:
					lhsCol := lhs.Int16()
					rhsCol := rhs.Int16()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				case 32:
					lhsCol := lhs.Int32()
					rhsCol := rhs.Int32()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				case -1:
				default:
					lhsCol := lhs.Int64()
					rhsCol := rhs.Int64()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := int64(lhsVal), int64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else {
											cmpResult = 0
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.FloatFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Float64()
					rhsCol := rhs.Float64()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := float64(lhsVal), float64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else if a == b {
											cmpResult = 0
										} else if math.IsNaN(a) {
											if math.IsNaN(b) {
												cmpResult = 0
											} else {
												cmpResult = -1
											}
										} else {
											cmpResult = 1
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := float64(lhsVal), float64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else if a == b {
											cmpResult = 0
										} else if math.IsNaN(a) {
											if math.IsNaN(b) {
												cmpResult = 0
											} else {
												cmpResult = -1
											}
										} else {
											cmpResult = 1
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := float64(lhsVal), float64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else if a == b {
											cmpResult = 0
										} else if math.IsNaN(a) {
											if math.IsNaN(b) {
												cmpResult = 0
											} else {
												cmpResult = -1
											}
										} else {
											cmpResult = 1
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									{
										a, b := float64(lhsVal), float64(rhsVal)
										if a < b {
											cmpResult = -1
										} else if a > b {
											cmpResult = 1
										} else if a == b {
											cmpResult = 0
										} else if math.IsNaN(a) {
											if math.IsNaN(b) {
												cmpResult = 0
											} else {
												cmpResult = -1
											}
										} else {
											cmpResult = 1
										}
									}

									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.TimestampTZFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Timestamp()
					rhsCol := rhs.Timestamp()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if lhsVal.Before(rhsVal) {
										cmpResult = -1
									} else if rhsVal.Before(lhsVal) {
										cmpResult = 1
									} else {
										cmpResult = 0
									}
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if lhsVal.Before(rhsVal) {
										cmpResult = -1
									} else if rhsVal.Before(lhsVal) {
										cmpResult = 1
									} else {
										cmpResult = 0
									}
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if lhsVal.Before(rhsVal) {
										cmpResult = -1
									} else if rhsVal.Before(lhsVal) {
										cmpResult = 1
									} else {
										cmpResult = 0
									}
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int

									if lhsVal.Before(rhsVal) {
										cmpResult = -1
									} else if rhsVal.Before(lhsVal) {
										cmpResult = 1
									} else {
										cmpResult = 0
									}
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			case types.IntervalFamily:
				switch keyTypes[keyIdx].Width() {
				case -1:
				default:
					lhsCol := lhs.Interval()
					rhsCol := rhs.Interval()
					if lhsHasNull {
						lhsNull := lhs.Nulls().NullAt(v.keyIdx)
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								rhsNull := rhs.Nulls().NullAt(rowIdx)
								if lhsNull && rhsNull {
									// Both values are NULLs, and we do not consider them different.
									continue
								} else if lhsNull || rhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = lhsVal.Compare(rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {
								if lhsNull {
									diff[selIdx] = true
									continue
								}

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = lhsVal.Compare(rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					} else {
						if rhsHasNull {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = lhsVal.Compare(rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						} else {
							lhsVal := lhsCol[aggKeyIdx]

							for selIdx, rowIdx := range sel {

								rhsVal := rhsCol[rowIdx]

								var cmp bool

								{
									var cmpResult int
									cmpResult = lhsVal.Compare(rhsVal)
									cmp = cmpResult != 0
								}

								diff[selIdx] = diff[selIdx] || cmp
							}

						}
					}
				}
			default:
				colexecerror.InternalError(fmt.Sprintf("unhandled type %s", keyTypes[keyIdx]))
			}
		}
	}

	remaining := sel[:0]
	for selIdx, tupleIdx := range sel {
		if diff[selIdx] {
			remaining = append(remaining, tupleIdx)
		} else {
			matched = append(matched, tupleIdx)
		}
	}

	anyMatched := false
	if len(matched) > 0 {
		b.SetLength(len(matched))
		anyMatched = true
	}

	// Reset diff slice back to all false.
	for n := 0; n < len(diff); n += copy(diff, zeroBoolColumn) {
	}

	return anyMatched, remaining
}
