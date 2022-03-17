// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexecjoin

import "github.com/cockroachdb/cockroach/pkg/col/coldata"

const _ = "template_collectProbeOuter"

const _ = "template_collectProbeNoOuter"

// This code snippet collects the "matches" for LEFT ANTI and EXCEPT ALL joins.
// "Matches" are in quotes because we're actually interested in non-matches
// from the left side.
const _ = "template_collectLeftAnti"

// collectRightSemiAnti processes all matches for right semi/anti joins. Note
// that during the probing phase we do not emit any output for these joins and
// are simply tracking whether build rows had a match. The output will be
// populated when in hjEmittingRight state.
func collectRightSemiAnti(hj *hashJoiner, batchSize int) {
	// Early bounds checks.
	// Capture the slice in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	_ = HeadIDs[batchSize-1]
	for i := 0; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]
		for currentID != 0 {
			hj.probeState.buildRowMatched[currentID-1] = true
			currentID = hj.ht.Same[currentID]
		}
	}
}

const _ = "template_distinctCollectProbeOuter"

const _ = "template_distinctCollectProbeNoOuter"

// collect prepares the buildIdx and probeIdx arrays where the buildIdx and
// probeIdx at each index are joined to make an output row. The total number of
// resulting rows is returned.
func (hj *hashJoiner) collect(batch coldata.Batch, batchSize int, sel []int) int {
	nResults := 0

	if !hj.spec.JoinType.ShouldIncludeLeftColsInOutput() {
		collectRightSemiAnti(hj, batchSize)
		return 0
	}

	if hj.spec.JoinType.IsLeftOuterOrFullOuter() {
		if sel != nil {
			nResults = collectProbeOuter_true(hj, batchSize, nResults, batch, sel)
		} else {
			nResults = collectProbeOuter_false(hj, batchSize, nResults, batch, sel)
		}
	} else {
		if sel != nil {
			if hj.spec.JoinType.IsLeftAntiOrExceptAll() {
				nResults = collectLeftAnti_true(hj, batchSize, nResults, batch, sel)
			} else {
				nResults = collectProbeNoOuter_true(hj, batchSize, nResults, batch, sel)
			}
		} else {
			if hj.spec.JoinType.IsLeftAntiOrExceptAll() {
				nResults = collectLeftAnti_false(hj, batchSize, nResults, batch, sel)
			} else {
				nResults = collectProbeNoOuter_false(hj, batchSize, nResults, batch, sel)
			}
		}
	}

	return nResults
}

// distinctCollect prepares the batch with the joined output columns where the build
// row index for each probe row is given in the ToCheckID slice. This function
// requires assumes a N-1 hash join.
func (hj *hashJoiner) distinctCollect(batch coldata.Batch, batchSize int, sel []int) int {
	nResults := 0

	if !hj.spec.JoinType.ShouldIncludeLeftColsInOutput() {
		collectRightSemiAnti(hj, batchSize)
		return 0
	}

	if hj.spec.JoinType.IsLeftOuterOrFullOuter() {
		nResults = batchSize

		if sel != nil {
			distinctCollectProbeOuter_true(hj, batchSize, sel)
		} else {
			distinctCollectProbeOuter_false(hj, batchSize, sel)
		}
	} else {
		if sel != nil {
			if hj.spec.JoinType.IsLeftAntiOrExceptAll() {
				// For LEFT ANTI and EXCEPT ALL joins we don't care whether the build
				// (right) side was distinct, so we only have single variation of COLLECT
				// method.
				nResults = collectLeftAnti_true(hj, batchSize, nResults, batch, sel)
			} else {
				nResults = distinctCollectProbeNoOuter_true(hj, batchSize, nResults, sel)
			}
		} else {
			if hj.spec.JoinType.IsLeftAntiOrExceptAll() {
				// For LEFT ANTI and EXCEPT ALL joins we don't care whether the build
				// (right) side was distinct, so we only have single variation of COLLECT
				// method.
				nResults = collectLeftAnti_false(hj, batchSize, nResults, batch, sel)
			} else {
				nResults = distinctCollectProbeNoOuter_false(hj, batchSize, nResults, sel)
			}
		}
	}

	return nResults
}

// execgen:inline
const _ = "template_getIdx"

func collectProbeOuter_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	startIdx := hj.probeState.prevBatchResumeIdx
	_ = HeadIDs[startIdx]
	_ = HeadIDs[batchSize-1]
	_ = sel[batchSize-1]
	maxResults := len(hj.probeState.buildIdx)
	buildIdx := hj.probeState.buildIdx
	probeIdx := hj.probeState.probeIdx
	_ = buildIdx[nResults]
	_ = probeIdx[nResults]
	_ = buildIdx[maxResults-1]
	_ = probeIdx[maxResults-1]
	for i := startIdx; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]

		for ; nResults < maxResults; nResults++ {
			rowUnmatched := currentID == 0
			// For some reason, BCE doesn't occur for probeRowUnmatched slice.
			// TODO(yuzefovich): figure it out.
			hj.probeState.probeRowUnmatched[nResults] = rowUnmatched
			if rowUnmatched {
				// The row is unmatched, and we set the corresponding buildIdx
				// to zero so that (as long as the build hash table has at least
				// one row) we can copy the values vector without paying
				// attention to probeRowUnmatched.
				//gcassert:bce
				buildIdx[nResults] = 0
			} else {
				//gcassert:bce
				buildIdx[nResults] = int(currentID - 1)
			}
			var pIdx int
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				pIdx = __retval_0
			}
			//gcassert:bce
			probeIdx[nResults] = pIdx
			currentID = hj.ht.Same[currentID]
			//gcassert:bce
			HeadIDs[i] = currentID

			if currentID == 0 {
				nResults++
				break
			}
		}

		if nResults == maxResults {
			// We have collected the maximum number of results that fit into the
			// current output batch.
			if currentID != 0 {
				// We haven't finished probing the ith tuple of the current
				// probing batch, so we'll need to resume from the same state.
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
			} else {
				// We're done probing the ith tuple.
				if i+1 < batchSize {
					// But we're not done probing the batch yet.
					hj.probeState.prevBatch = batch
					hj.probeState.prevBatchResumeIdx = i + 1
				}
			}
			return nResults
		}
	}
	return nResults
}

func collectProbeOuter_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	startIdx := hj.probeState.prevBatchResumeIdx
	_ = HeadIDs[startIdx]
	_ = HeadIDs[batchSize-1]
	maxResults := len(hj.probeState.buildIdx)
	buildIdx := hj.probeState.buildIdx
	probeIdx := hj.probeState.probeIdx
	_ = buildIdx[nResults]
	_ = probeIdx[nResults]
	_ = buildIdx[maxResults-1]
	_ = probeIdx[maxResults-1]
	for i := startIdx; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]

		for ; nResults < maxResults; nResults++ {
			rowUnmatched := currentID == 0
			// For some reason, BCE doesn't occur for probeRowUnmatched slice.
			// TODO(yuzefovich): figure it out.
			hj.probeState.probeRowUnmatched[nResults] = rowUnmatched
			if rowUnmatched {
				// The row is unmatched, and we set the corresponding buildIdx
				// to zero so that (as long as the build hash table has at least
				// one row) we can copy the values vector without paying
				// attention to probeRowUnmatched.
				//gcassert:bce
				buildIdx[nResults] = 0
			} else {
				//gcassert:bce
				buildIdx[nResults] = int(currentID - 1)
			}
			var pIdx int
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				pIdx = __retval_0
			}
			//gcassert:bce
			probeIdx[nResults] = pIdx
			currentID = hj.ht.Same[currentID]
			//gcassert:bce
			HeadIDs[i] = currentID

			if currentID == 0 {
				nResults++
				break
			}
		}

		if nResults == maxResults {
			// We have collected the maximum number of results that fit into the
			// current output batch.
			if currentID != 0 {
				// We haven't finished probing the ith tuple of the current
				// probing batch, so we'll need to resume from the same state.
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
			} else {
				// We're done probing the ith tuple.
				if i+1 < batchSize {
					// But we're not done probing the batch yet.
					hj.probeState.prevBatch = batch
					hj.probeState.prevBatchResumeIdx = i + 1
				}
			}
			return nResults
		}
	}
	return nResults
}

// This code snippet collects the "matches" for LEFT ANTI and EXCEPT ALL joins.
// "Matches" are in quotes because we're actually interested in non-matches
// from the left side.
func collectLeftAnti_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slice in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	_ = HeadIDs[batchSize-1]
	_ = sel[batchSize-1]
	for i := 0; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]
		if currentID == 0 {
			// currentID of 0 indicates that ith probing row didn't have a match, so
			// we include it into the output.
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				// currentID of 0 indicates that ith probing row didn't have a match, so
				// we include it into the output.
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

func collectProbeNoOuter_true(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	startIdx := hj.probeState.prevBatchResumeIdx
	_ = HeadIDs[startIdx]
	_ = HeadIDs[batchSize-1]
	_ = sel[batchSize-1]
	maxResults := len(hj.probeState.buildIdx)
	probeIdx := hj.probeState.probeIdx
	_ = probeIdx[nResults]
	_ = probeIdx[maxResults-1]
	for i := startIdx; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]
		for ; currentID != 0 && nResults < maxResults; nResults++ {
			// For some reason, BCE doesn't occur for buildIdx slice.
			// TODO(yuzefovich): figure it out.
			hj.probeState.buildIdx[nResults] = int(currentID - 1)
			var pIdx int
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				pIdx = __retval_0
			}
			//gcassert:bce
			probeIdx[nResults] = pIdx
			currentID = hj.ht.Same[currentID]
			//gcassert:bce
			HeadIDs[i] = currentID
		}

		if nResults == maxResults {
			// We have collected the maximum number of results that fit into the
			// current output batch.
			if currentID != 0 {
				// We haven't finished probing the ith tuple of the current
				// probing batch, so we'll need to resume from the same state.
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
			} else {
				// We're done probing the ith tuple.
				if i+1 < batchSize {
					// But we're not done probing the batch yet.
					hj.probeState.prevBatch = batch
					hj.probeState.prevBatchResumeIdx = i + 1
				}
			}
			return nResults
		}
	}
	return nResults
}

// This code snippet collects the "matches" for LEFT ANTI and EXCEPT ALL joins.
// "Matches" are in quotes because we're actually interested in non-matches
// from the left side.
func collectLeftAnti_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slice in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	_ = HeadIDs[batchSize-1]
	for i := 0; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]
		if currentID == 0 {
			// currentID of 0 indicates that ith probing row didn't have a match, so
			// we include it into the output.
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				// currentID of 0 indicates that ith probing row didn't have a match, so
				// we include it into the output.
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

func collectProbeNoOuter_false(
	hj *hashJoiner, batchSize int, nResults int, batch coldata.Batch, sel []int) int {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	HeadIDs := hj.ht.ProbeScratch.HeadID
	startIdx := hj.probeState.prevBatchResumeIdx
	_ = HeadIDs[startIdx]
	_ = HeadIDs[batchSize-1]
	maxResults := len(hj.probeState.buildIdx)
	probeIdx := hj.probeState.probeIdx
	_ = probeIdx[nResults]
	_ = probeIdx[maxResults-1]
	for i := startIdx; i < batchSize; i++ {
		//gcassert:bce
		currentID := HeadIDs[i]
		for ; currentID != 0 && nResults < maxResults; nResults++ {
			// For some reason, BCE doesn't occur for buildIdx slice.
			// TODO(yuzefovich): figure it out.
			hj.probeState.buildIdx[nResults] = int(currentID - 1)
			var pIdx int
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				pIdx = __retval_0
			}
			//gcassert:bce
			probeIdx[nResults] = pIdx
			currentID = hj.ht.Same[currentID]
			//gcassert:bce
			HeadIDs[i] = currentID
		}

		if nResults == maxResults {
			// We have collected the maximum number of results that fit into the
			// current output batch.
			if currentID != 0 {
				// We haven't finished probing the ith tuple of the current
				// probing batch, so we'll need to resume from the same state.
				hj.probeState.prevBatch = batch
				hj.probeState.prevBatchResumeIdx = i
			} else {
				// We're done probing the ith tuple.
				if i+1 < batchSize {
					// But we're not done probing the batch yet.
					hj.probeState.prevBatch = batch
					hj.probeState.prevBatchResumeIdx = i + 1
				}
			}
			return nResults
		}
	}
	return nResults
}

func distinctCollectProbeOuter_true(hj *hashJoiner, batchSize int, sel []int) {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	toCheckIDs := hj.ht.ProbeScratch.ToCheckID
	probeRowUnmatched := hj.probeState.probeRowUnmatched
	buildIdx := hj.probeState.buildIdx
	probeIdx := hj.probeState.probeIdx
	_ = toCheckIDs[batchSize-1]
	_ = probeRowUnmatched[batchSize-1]
	_ = buildIdx[batchSize-1]
	_ = probeIdx[batchSize-1]
	_ = sel[batchSize-1]
	for i := 0; i < batchSize; i++ {
		// Index of keys and outputs in the hash table is calculated as ID - 1.
		//gcassert:bce
		id := toCheckIDs[i]
		rowUnmatched := id == 0
		//gcassert:bce
		probeRowUnmatched[i] = rowUnmatched
		if rowUnmatched {
			// The row is unmatched, and we set the corresponding buildIdx
			// to zero so that (as long as the build hash table has at least
			// one row) we can copy the values vector without paying
			// attention to probeRowUnmatched.
			//gcassert:bce
			buildIdx[i] = 0
		} else {
			//gcassert:bce
			buildIdx[i] = int(id - 1)
		}
		var pIdx int
		{
			var __retval_0 int
			{
				{
					__retval_0 = sel[i]
				}
			}
			pIdx = __retval_0
		}
		//gcassert:bce
		probeIdx[i] = pIdx
	}
}

func distinctCollectProbeOuter_false(hj *hashJoiner, batchSize int, sel []int) {
	// Early bounds checks.
	// Capture the slices in order for BCE to occur.
	toCheckIDs := hj.ht.ProbeScratch.ToCheckID
	probeRowUnmatched := hj.probeState.probeRowUnmatched
	buildIdx := hj.probeState.buildIdx
	probeIdx := hj.probeState.probeIdx
	_ = toCheckIDs[batchSize-1]
	_ = probeRowUnmatched[batchSize-1]
	_ = buildIdx[batchSize-1]
	_ = probeIdx[batchSize-1]
	for i := 0; i < batchSize; i++ {
		// Index of keys and outputs in the hash table is calculated as ID - 1.
		//gcassert:bce
		id := toCheckIDs[i]
		rowUnmatched := id == 0
		//gcassert:bce
		probeRowUnmatched[i] = rowUnmatched
		if rowUnmatched {
			// The row is unmatched, and we set the corresponding buildIdx
			// to zero so that (as long as the build hash table has at least
			// one row) we can copy the values vector without paying
			// attention to probeRowUnmatched.
			//gcassert:bce
			buildIdx[i] = 0
		} else {
			//gcassert:bce
			buildIdx[i] = int(id - 1)
		}
		var pIdx int
		{
			var __retval_0 int
			{
				{
					__retval_0 = i
				}
			}
			pIdx = __retval_0
		}
		//gcassert:bce
		probeIdx[i] = pIdx
	}
}

func distinctCollectProbeNoOuter_true(
	hj *hashJoiner, batchSize int, nResults int, sel []int) int {
	// Early bounds checks.
	// Capture the slice in order for BCE to occur.
	toCheckIDs := hj.ht.ProbeScratch.ToCheckID
	_ = toCheckIDs[batchSize-1]
	_ = sel[batchSize-1]
	for i := 0; i < batchSize; i++ {
		//gcassert:bce
		id := toCheckIDs[i]
		if id != 0 {
			// Index of keys and outputs in the hash table is calculated as ID - 1.
			hj.probeState.buildIdx[nResults] = int(id - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = sel[i]
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

func distinctCollectProbeNoOuter_false(
	hj *hashJoiner, batchSize int, nResults int, sel []int) int {
	// Early bounds checks.
	// Capture the slice in order for BCE to occur.
	toCheckIDs := hj.ht.ProbeScratch.ToCheckID
	_ = toCheckIDs[batchSize-1]
	for i := 0; i < batchSize; i++ {
		//gcassert:bce
		id := toCheckIDs[i]
		if id != 0 {
			// Index of keys and outputs in the hash table is calculated as ID - 1.
			hj.probeState.buildIdx[nResults] = int(id - 1)
			{
				var __retval_0 int
				{
					{
						__retval_0 = i
					}
				}
				hj.probeState.probeIdx[nResults] = __retval_0
			}
			nResults++
		}
	}
	return nResults
}

// execgen:inline
const _ = "inlined_getIdx_true"

// execgen:inline
const _ = "inlined_getIdx_false"
