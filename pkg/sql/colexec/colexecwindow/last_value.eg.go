// Code generated by execgen; DO NOT EDIT.
// Copyright 2021 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecwindow

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/colexecutils"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecop"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

// NewLastValueOperator creates a new Operator that computes window
// function lastValue. outputColIdx specifies in which coldata.Vec the operator
// should put its output (if there is no such column, a new column is appended).
func NewLastValueOperator(
	args *WindowArgs,
	frame *execinfrapb.WindowerSpec_Frame,
	ordering *execinfrapb.Ordering,
	argIdxs []int,
) (colexecop.Operator, error) {
	framer := newWindowFramer(args.EvalCtx, frame, ordering, args.InputTypes, args.PeersColIdx)
	colsToStore := framer.getColsToStore([]int{argIdxs[0]})

	// Allow the direct-access buffer 10% of the available memory. The rest will
	// be given to the bufferedWindowOp queue. While it is somewhat more important
	// for the direct-access buffer tuples to be kept in-memory, it only has to
	// store a single column. TODO(drewk): play around with benchmarks to find a
	// good empirically-supported fraction to use.
	bufferMemLimit := int64(float64(args.MemoryLimit) * 0.10)
	buffer := colexecutils.NewSpillingBuffer(
		args.BufferAllocator, bufferMemLimit, args.QueueCfg,
		args.FdSemaphore, args.InputTypes, args.DiskAcc, colsToStore...)
	base := lastValueBase{
		partitionSeekerBase: partitionSeekerBase{
			buffer:          buffer,
			partitionColIdx: args.PartitionColIdx,
		},
		framer:       framer,
		outputColIdx: args.OutputColIdx,
		bufferArgIdx: 0, // The arg column is the first column in the buffer.
	}
	argType := args.InputTypes[argIdxs[0]]
	switch typeconv.TypeFamilyToCanonicalTypeFamily(argType.Family()) {
	case types.BoolFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueBoolWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.BytesFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueBytesWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.DecimalFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueDecimalWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.IntFamily:
		switch argType.Width() {
		case 16:
			windower := &lastValueInt16Window{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		case 32:
			windower := &lastValueInt32Window{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		case -1:
		default:
			windower := &lastValueInt64Window{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.FloatFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueFloat64Window{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.TimestampTZFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueTimestampWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.IntervalFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueIntervalWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case types.JsonFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueJSONWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch argType.Width() {
		case -1:
		default:
			windower := &lastValueDatumWindow{lastValueBase: base}
			return newBufferedWindowOperator(args, windower, argType), nil
		}
	}
	return nil, errors.Errorf("unsupported lastValue window operator type %s", argType.Name())
}

type lastValueBase struct {
	partitionSeekerBase
	colexecop.CloserHelper
	framer windowFramer

	outputColIdx int
	bufferArgIdx int
}

type lastValueBoolWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueBoolWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueBoolWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Bool()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Bool()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueBytesWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueBytesWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueBytesWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Bytes()
	outputNulls := outputVec.Nulls()

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Bytes()
		val := col.Get(idx)
		outputCol.Set(i, val)
	}
}

type lastValueDecimalWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueDecimalWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueDecimalWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Decimal()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Decimal()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueInt16Window struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueInt16Window{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueInt16Window) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Int16()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Int16()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueInt32Window struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueInt32Window{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueInt32Window) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Int32()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Int32()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueInt64Window struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueInt64Window{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueInt64Window) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Int64()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Int64()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueFloat64Window struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueFloat64Window{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueFloat64Window) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Float64()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Float64()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueTimestampWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueTimestampWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueTimestampWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Timestamp()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Timestamp()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueIntervalWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueIntervalWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueIntervalWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Interval()
	outputNulls := outputVec.Nulls()
	_, _ = outputCol.Get(startIdx), outputCol.Get(endIdx-1)

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Interval()
		val := col.Get(idx)
		//gcassert:bce
		outputCol.Set(i, val)
	}
}

type lastValueJSONWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueJSONWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueJSONWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.JSON()
	outputNulls := outputVec.Nulls()

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.JSON()
		val := col.Get(idx)
		outputCol.Set(i, val)
	}
}

type lastValueDatumWindow struct {
	lastValueBase
}

var _ bufferedWindower = &lastValueDatumWindow{}

// processBatch implements the bufferedWindower interface.
func (w *lastValueDatumWindow) processBatch(batch coldata.Batch, startIdx, endIdx int) {
	if startIdx >= endIdx {
		// No processing needs to be done for this portion of the current partition.
		return
	}
	outputVec := batch.ColVec(w.outputColIdx)
	outputCol := outputVec.Datum()
	outputNulls := outputVec.Nulls()

	for i := startIdx; i < endIdx; i++ {
		w.framer.next(w.Ctx)
		requestedIdx := w.framer.frameLastIdx()
		if requestedIdx == -1 {
			// The requested row does not exist.
			outputNulls.SetNull(i)
			continue
		}

		vec, idx, _ := w.buffer.GetVecWithTuple(w.Ctx, w.bufferArgIdx, requestedIdx)
		if vec.Nulls().MaybeHasNulls() && vec.Nulls().NullAt(idx) {
			outputNulls.SetNull(i)
			continue
		}
		col := vec.Datum()
		val := col.Get(idx)
		outputCol.Set(i, val)
	}
}

// transitionToProcessing implements the bufferedWindower interface.
func (b *lastValueBase) transitionToProcessing() {
	b.framer.startPartition(b.Ctx, b.partitionSize, b.buffer)
}

// startNewPartition implements the bufferedWindower interface.
func (b *lastValueBase) startNewPartition() {
	b.partitionSize = 0
	b.buffer.Reset(b.Ctx)
}

// Init implements the bufferedWindower interface.
func (b *lastValueBase) Init(ctx context.Context) {
	if !b.InitHelper.Init(ctx) {
		return
	}
}

// Close implements the bufferedWindower interface.
func (b *lastValueBase) Close() {
	if !b.CloserHelper.Close() {
		return
	}
	b.buffer.Close(b.EnsureCtx())
}
