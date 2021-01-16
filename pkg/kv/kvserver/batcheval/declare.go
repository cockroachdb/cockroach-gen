// Copyright 2014 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package batcheval

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/keys"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/spanset"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/storage/enginepb"
)

// DefaultDeclareKeys is the default implementation of Command.DeclareKeys.
func DefaultDeclareKeys(
	_ *roachpb.RangeDescriptor,
	header roachpb.Header,
	req roachpb.Request,
	latchSpans, _ *spanset.SpanSet,
) {
	access := spanset.SpanReadWrite
	if roachpb.IsReadOnly(req) && !roachpb.IsLocking(req) {
		access = spanset.SpanReadOnly
	}
	latchSpans.AddMVCC(access, req.Header().Span(), header.Timestamp)
}

// DefaultDeclareIsolatedKeys is similar to DefaultDeclareKeys, but it declares
// both lock spans in addition to latch spans. When used, commands will wait on
// locks and wait-queues owned by other transactions before evaluating. This
// ensures that the commands are fully isolated from conflicting transactions
// when it evaluated.
func DefaultDeclareIsolatedKeys(
	_ *roachpb.RangeDescriptor,
	header roachpb.Header,
	req roachpb.Request,
	latchSpans, lockSpans *spanset.SpanSet,
) {
	access := spanset.SpanReadWrite
	timestamp := header.Timestamp
	if roachpb.IsReadOnly(req) && !roachpb.IsLocking(req) {
		access = spanset.SpanReadOnly
		if header.Txn != nil {
			// For transactional reads, acquire read latches all the way up to
			// the transaction's MaxTimestamp, because reads may observe writes
			// all the way up to this timestamp.
			//
			// It is critical that reads declare latches up through their
			// uncertainty interval so that they are properly synchronized with
			// earlier writes that may have a happened-before relationship with
			// the read. These writes could not have completed and returned to
			// the client until they were durable in the Range's Raft log.
			// However, they may not have been applied to the replica's state
			// machine by the time the write was acknowledged, because Raft
			// entry application occurs asynchronously with respect to the
			// writer (see AckCommittedEntriesBeforeApplication). Latching is
			// the only mechanism that ensures that any observers of the write
			// wait for the write apply before reading.
			timestamp.Forward(header.Txn.MaxTimestamp)
		}
	}
	latchSpans.AddMVCC(access, req.Header().Span(), timestamp)
	lockSpans.AddNonMVCC(access, req.Header().Span())
}

// DeclareKeysForBatch adds all keys that the batch with the provided header
// touches to the given SpanSet. This does not include keys touched during the
// processing of the batch's individual commands.
func DeclareKeysForBatch(
	desc *roachpb.RangeDescriptor, header roachpb.Header, latchSpans *spanset.SpanSet,
) {
	if header.Txn != nil {
		header.Txn.AssertInitialized(context.TODO())
		latchSpans.AddNonMVCC(spanset.SpanReadOnly, roachpb.Span{
			Key: keys.AbortSpanKey(header.RangeID, header.Txn.ID),
		})
	}
}

// CommandArgs contains all the arguments to a command.
// TODO(bdarnell): consider merging with kvserverbase.FilterArgs (which
// would probably require removing the EvalCtx field due to import order
// constraints).
type CommandArgs struct {
	EvalCtx EvalContext
	Header  roachpb.Header
	Args    roachpb.Request
	// *Stats should be mutated to reflect any writes made by the command.
	Stats *enginepb.MVCCStats
}
