// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kvserver_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/keys"
	"github.com/cockroachdb/cockroach/pkg/kv"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/storagebase"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/testutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func relocateAndCheck(
	t *testing.T,
	tc *testcluster.TestCluster,
	startKey roachpb.RKey,
	targets []roachpb.ReplicationTarget,
) (retries int) {
	testutils.SucceedsSoon(t, func() error {
		err := tc.Servers[0].DB().
			AdminRelocateRange(context.Background(), startKey.AsRawKey(), targets)
		if err != nil {
			retries++
		}
		return err
	})
	desc, err := tc.Servers[0].LookupRange(startKey.AsRawKey())
	require.NoError(t, err)
	requireDescMembers(t, desc, targets)
	requireLeaseAt(t, tc, desc, targets[0])
	return retries
}

func requireDescMembers(
	t *testing.T, desc roachpb.RangeDescriptor, targets []roachpb.ReplicationTarget,
) {
	t.Helper()
	targets = append([]roachpb.ReplicationTarget(nil), targets...)
	sort.Slice(targets, func(i, j int) bool { return targets[i].StoreID < targets[j].StoreID })

	have := make([]roachpb.ReplicationTarget, 0, len(targets))
	for _, rDesc := range desc.Replicas().All() {
		have = append(have, roachpb.ReplicationTarget{
			NodeID:  rDesc.NodeID,
			StoreID: rDesc.StoreID,
		})
	}
	sort.Slice(have, func(i, j int) bool { return have[i].StoreID < have[j].StoreID })
	require.Equal(t, targets, have)
}

func requireLeaseAt(
	t *testing.T,
	tc *testcluster.TestCluster,
	desc roachpb.RangeDescriptor,
	target roachpb.ReplicationTarget,
) {
	t.Helper()
	// NB: under stressrace the lease will sometimes be inactive by the time
	// it's returned here, so don't use FindRangeLeaseHolder which fails when
	// that happens.
	testutils.SucceedsSoon(t, func() error {
		lease, _, err := tc.FindRangeLease(desc, &target)
		if err != nil {
			return err
		}
		if target != (roachpb.ReplicationTarget{
			NodeID:  lease.Replica.NodeID,
			StoreID: lease.Replica.StoreID,
		}) {
			return errors.Errorf("lease %v is not held by %+v", lease, target)
		}
		return nil
	})
}

func TestAdminRelocateRange(t *testing.T) {
	defer leaktest.AfterTest(t)()

	ctx := context.Background()

	type intercept struct {
		ops         []roachpb.ReplicationChange
		leaseTarget *roachpb.ReplicationTarget
		err         error
	}
	var intercepted []intercept

	requireNumAtomic := func(expAtomic int, expSingle int, f func() (retries int)) {
		t.Helper()
		intercepted = nil
		retries := f()
		var actAtomic, actSingle int
		for _, ic := range intercepted {
			if ic.err != nil {
				continue
			}
			if len(ic.ops) == 2 && ic.ops[0].ChangeType == roachpb.ADD_REPLICA && ic.ops[1].ChangeType == roachpb.REMOVE_REPLICA {
				actAtomic++
			} else {
				actSingle++
			}
		}
		actAtomic -= retries
		require.Equal(t, expAtomic, actAtomic, "wrong number of atomic changes: %+v", intercepted)
		require.Equal(t, expSingle, actSingle, "wrong number of single changes: %+v", intercepted)
	}

	knobs := base.TestingKnobs{
		Store: &kvserver.StoreTestingKnobs{
			BeforeRelocateOne: func(ops []roachpb.ReplicationChange, leaseTarget *roachpb.ReplicationTarget, err error) {
				intercepted = append(intercepted, intercept{
					ops:         ops,
					leaseTarget: leaseTarget,
					err:         err,
				})
			},
		},
	}
	args := base.TestClusterArgs{
		ServerArgs:      base.TestServerArgs{Knobs: knobs},
		ReplicationMode: base.ReplicationManual,
	}
	tc := testcluster.StartTestCluster(t, 6, args)
	defer tc.Stopper().Stop(ctx)

	// s1 (LH) ---> s2 (LH) s1 s3
	// Pure upreplication.
	k := keys.MustAddr(tc.ScratchRange(t))
	{
		targets := tc.Targets(1, 0, 2)
		// Expect two single additions, and that's it.
		requireNumAtomic(0, 2, func() int {
			return relocateAndCheck(t, tc, k, targets)
		})
	}

	// s1 (LH) s2 s3 ---> s4 (LH) s5 s6.
	// This is trickier because the leaseholder gets removed, and so do all
	// other replicas (i.e. a simple lease transfer at the beginning won't solve
	// the problem).
	{
		targets := tc.Targets(3, 4, 5)
		// Should carry out three swaps. Note that the leaseholder gets removed
		// in the process (i.e. internally the lease must've been moved around
		// to achieve that).
		requireNumAtomic(3, 0, func() int {
			return relocateAndCheck(t, tc, k, targets)
		})
	}

	// s4 (LH) s5 s6 ---> s5 (LH)
	// Pure downreplication.
	{
		requireNumAtomic(0, 2, func() int {
			return relocateAndCheck(t, tc, k, tc.Targets(4))
		})
	}

	// s5 (LH) ---> s3 (LH)
	// Lateral movement while at replication factor one. In this case atomic
	// replication changes cannot be used; we add-then-remove instead.
	{
		requireNumAtomic(0, 2, func() int {
			return relocateAndCheck(t, tc, k, tc.Targets(2))
		})
	}

	// s3 (LH) ---> s2 (LH) s4 s1 --> s4 (LH) s2 s6 s1 --> s3 (LH) s5
	// A grab bag.
	{
		// s3 -(add)-> s3 s2 -(swap)-> s4 s2 -(add)-> s4 s2 s1 (=s2 s4 s1)
		requireNumAtomic(1, 2, func() int {
			return relocateAndCheck(t, tc, k, tc.Targets(1, 3, 0))
		})
		// s2 s4 s1 -(add)-> s2 s4 s1 s6 (=s4 s2 s6 s1)
		requireNumAtomic(0, 1, func() int {
			return relocateAndCheck(t, tc, k, tc.Targets(3, 1, 5, 0))
		})
		// s4 s2 s6 s1 -(swap)-> s3 s2 s6 s1 -(swap)-> s3 s5 s6 s1 -(del)-> s3 s5 s6 -(del)-> s3 s5
		requireNumAtomic(2, 2, func() int {
			return relocateAndCheck(t, tc, k, tc.Targets(2, 4))
		})
	}
}

// Regression test for https://github.com/cockroachdb/cockroach/issues/64325
// which makes sure an in-flight read operation during replica removal won't
// return empty results.
func TestReplicaRemovalDuringRequestEvaluation(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	type magicKey struct{}

	// delayReadC is used to synchronize the in-flight read request with the main
	// test goroutine. It is read from twice:
	//
	// 1. The first read allows the test to block until the request eval filter
	//    is called, i.e. when the read request is ready.
	// 2. The second read allows the test to close the channel to unblock
	//    the eval filter, causing the read request to be evaluated.
	delayReadC := make(chan struct{})
	evalFilter := func(args storagebase.FilterArgs) *roachpb.Error {
		if args.Ctx.Value(magicKey{}) != nil {
			<-delayReadC
			<-delayReadC
		}
		return nil
	}

	ctx := context.Background()
	manual := hlc.NewHybridManualClock()
	args := base.TestClusterArgs{
		ReplicationMode: base.ReplicationManual,
		ServerArgs: base.TestServerArgs{
			Knobs: base.TestingKnobs{
				Store: &kvserver.StoreTestingKnobs{
					EvalKnobs: storagebase.BatchEvalTestingKnobs{
						TestingEvalFilter: evalFilter,
					},
				},
				Server: &server.TestingKnobs{
					ClockSource: manual.UnixNano,
				},
			},
		},
	}
	tc := testcluster.StartTestCluster(t, 2, args)
	defer tc.Stopper().Stop(ctx)

	// Create range and upreplicate.
	key := tc.ScratchRange(t)
	tc.AddReplicasOrFatal(t, key, tc.Target(1))

	// Perform write.
	pArgs := putArgs(key, []byte("foo"))
	_, pErr := kv.SendWrapped(ctx, tc.Servers[0].DistSender(), pArgs)
	require.Nil(t, pErr)

	// Perform read on write and wait for read to block.
	type reply struct {
		resp roachpb.Response
		err  *roachpb.Error
	}
	readResC := make(chan reply)
	err := tc.Stopper().RunAsyncTask(ctx, "get", func(ctx context.Context) {
		readCtx := context.WithValue(ctx, magicKey{}, struct{}{})
		gArgs := getArgs(key)
		resp, pErr := kv.SendWrapped(readCtx, tc.Servers[0].DistSender(), gArgs)
		readResC <- reply{resp, pErr}
	})
	require.NoError(t, err)
	delayReadC <- struct{}{}

	// Transfer leaseholder to other store.
	rangeDesc, err := tc.LookupRange(key)
	require.NoError(t, err)
	store, err := tc.Server(0).GetStores().(*kvserver.Stores).GetStore(tc.Server(0).GetFirstStoreID())
	require.NoError(t, err)
	repl, err := store.GetReplica(rangeDesc.RangeID)
	require.NoError(t, err)
	err = tc.MoveRangeLeaseNonCooperatively(rangeDesc, tc.Target(1), manual)
	require.NoError(t, err)

	// Remove first store from raft group.
	tc.RemoveReplicasOrFatal(t, key, tc.Target(0))

	// This is a bit iffy. We want to make sure that, in the buggy case, we
	// will typically fail (i.e. the read returns empty because the replica was
	// removed). However, in the non-buggy case the in-flight read request will
	// be holding readOnlyCmdMu until evaluated, blocking the replica removal,
	// so waiting for replica removal would deadlock. We therefore take the
	// easy way out by starting an async replica GC and sleeping for a bit.
	err = tc.Stopper().RunAsyncTask(ctx, "replicaGC", func(ctx context.Context) {
		assert.NoError(t, store.ManualReplicaGC(repl))
	})
	require.NoError(t, err)
	time.Sleep(500 * time.Millisecond)

	// Allow read to resume. Should return "foo".
	close(delayReadC)
	r := <-readResC
	require.Nil(t, r.err)
	require.NotNil(t, r.resp)
	require.NotNil(t, r.resp.(*roachpb.GetResponse).Value)
	val, err := r.resp.(*roachpb.GetResponse).Value.GetBytes()
	require.NoError(t, err)
	require.Equal(t, []byte("foo"), val)
}
