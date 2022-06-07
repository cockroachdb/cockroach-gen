// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package storage

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/cli/exit"
	"github.com/cockroachdb/cockroach/pkg/clusterversion"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/settings"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/storage/enginepb"
	"github.com/cockroachdb/cockroach/pkg/storage/fs"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/cockroach/pkg/util/encoding"
	"github.com/cockroachdb/cockroach/pkg/util/envutil"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/humanizeutil"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/protoutil"
	"github.com/cockroachdb/cockroach/pkg/util/syncutil"
	"github.com/cockroachdb/cockroach/pkg/util/tracing"
	"github.com/cockroachdb/cockroach/pkg/util/uuid"
	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/oserror"
	"github.com/cockroachdb/logtags"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
	"github.com/cockroachdb/pebble/sstable"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/cockroachdb/redact"
	humanize "github.com/dustin/go-humanize"
)

const maxSyncDurationFatalOnExceededDefault = true

// Default for MaxSyncDuration below.
var maxSyncDurationDefault = envutil.EnvOrDefaultDuration("COCKROACH_ENGINE_MAX_SYNC_DURATION_DEFAULT", 20*time.Second)

// MaxSyncDuration is the threshold above which an observed engine sync duration
// triggers either a warning or a fatal error.
var MaxSyncDuration = settings.RegisterDurationSetting(
	settings.TenantWritable,
	"storage.max_sync_duration",
	"maximum duration for disk operations; any operations that take longer"+
		" than this setting trigger a warning log entry or process crash",
	maxSyncDurationDefault,
)

// MaxSyncDurationFatalOnExceeded governs whether disk stalls longer than
// MaxSyncDuration fatal the Cockroach process. Defaults to true.
var MaxSyncDurationFatalOnExceeded = settings.RegisterBoolSetting(
	settings.TenantWritable,
	"storage.max_sync_duration.fatal.enabled",
	"if true, fatal the process when a disk operation exceeds storage.max_sync_duration",
	maxSyncDurationFatalOnExceededDefault,
)

// EngineKeyCompare compares cockroach keys, including the version (which
// could be MVCC timestamps).
func EngineKeyCompare(a, b []byte) int {
	// NB: For performance, this routine manually splits the key into the
	// user-key and version components rather than using DecodeEngineKey. In
	// most situations, use DecodeEngineKey or GetKeyPartFromEngineKey or
	// SplitMVCCKey instead of doing this.
	aEnd := len(a) - 1
	bEnd := len(b) - 1
	if aEnd < 0 || bEnd < 0 {
		// This should never happen unless there is some sort of corruption of
		// the keys.
		return bytes.Compare(a, b)
	}

	// Compute the index of the separator between the key and the version. If the
	// separator is found to be at -1 for both keys, then we are comparing bare
	// suffixes without a user key part. Pebble requires bare suffixes to be
	// comparable with the same ordering as if they had a common user key.
	aSep := aEnd - int(a[aEnd])
	bSep := bEnd - int(b[bEnd])
	if aSep == -1 && bSep == -1 {
		aSep, bSep = 0, 0 // comparing bare suffixes
	}
	if aSep < 0 || bSep < 0 {
		// This should never happen unless there is some sort of corruption of
		// the keys.
		return bytes.Compare(a, b)
	}

	// Compare the "user key" part of the key.
	if c := bytes.Compare(a[:aSep], b[:bSep]); c != 0 {
		return c
	}

	// Compare the version part of the key. Note that when the version is a
	// timestamp, the timestamp encoding causes byte comparison to be equivalent
	// to timestamp comparison.
	aVer := a[aSep:aEnd]
	bVer := b[bSep:bEnd]
	if len(aVer) == 0 {
		if len(bVer) == 0 {
			return 0
		}
		return -1
	} else if len(bVer) == 0 {
		return 1
	}
	aVer = normalizeEngineKeyVersionForCompare(aVer)
	bVer = normalizeEngineKeyVersionForCompare(bVer)
	return bytes.Compare(bVer, aVer)
}

// EngineKeyEqual checks for equality of cockroach keys, including the version
// (which could be MVCC timestamps).
func EngineKeyEqual(a, b []byte) bool {
	// NB: For performance, this routine manually splits the key into the
	// user-key and version components rather than using DecodeEngineKey. In
	// most situations, use DecodeEngineKey or GetKeyPartFromEngineKey or
	// SplitMVCCKey instead of doing this.
	aEnd := len(a) - 1
	bEnd := len(b) - 1
	if aEnd < 0 || bEnd < 0 {
		// This should never happen unless there is some sort of corruption of
		// the keys.
		return bytes.Equal(a, b)
	}

	// Last byte is the version length + 1 when there is a version,
	// else it is 0.
	aVerLen := int(a[aEnd])
	bVerLen := int(b[bEnd])

	// Fast-path. If the key version is empty or contains only a walltime
	// component then normalizeEngineKeyVersionForCompare is a no-op, so we don't
	// need to split the "user key" from the version suffix before comparing to
	// compute equality. Instead, we can check for byte equality immediately.
	const withWall = mvccEncodedTimeSentinelLen + mvccEncodedTimeWallLen
	if aVerLen <= withWall && bVerLen <= withWall {
		return bytes.Equal(a, b)
	}

	// Compute the index of the separator between the key and the version. If the
	// separator is found to be at -1 for both keys, then we are comparing bare
	// suffixes without a user key part. Pebble requires bare suffixes to be
	// comparable with the same ordering as if they had a common user key.
	aSep := aEnd - aVerLen
	bSep := bEnd - bVerLen
	if aSep == -1 && bSep == -1 {
		aSep, bSep = 0, 0 // comparing bare suffixes
	}
	if aSep < 0 || bSep < 0 {
		// This should never happen unless there is some sort of corruption of
		// the keys.
		return bytes.Equal(a, b)
	}

	// Compare the "user key" part of the key.
	if !bytes.Equal(a[:aSep], b[:bSep]) {
		return false
	}

	// Compare the version part of the key.
	aVer := a[aSep:aEnd]
	bVer := b[bSep:bEnd]
	aVer = normalizeEngineKeyVersionForCompare(aVer)
	bVer = normalizeEngineKeyVersionForCompare(bVer)
	return bytes.Equal(aVer, bVer)
}

var zeroLogical [mvccEncodedTimeLogicalLen]byte

//gcassert:inline
func normalizeEngineKeyVersionForCompare(a []byte) []byte {
	// In general, the version could also be a non-timestamp version, but we know
	// that engineKeyVersionLockTableLen+mvccEncodedTimeSentinelLen is a different
	// constant than the above, so there is no danger here of stripping parts from
	// a non-timestamp version.
	const withWall = mvccEncodedTimeSentinelLen + mvccEncodedTimeWallLen
	const withLogical = withWall + mvccEncodedTimeLogicalLen
	const withSynthetic = withLogical + mvccEncodedTimeSyntheticLen
	if len(a) == withSynthetic {
		// Strip the synthetic bit component from the timestamp version. The
		// presence of the synthetic bit does not affect key ordering or equality.
		a = a[:withLogical]
	}
	if len(a) == withLogical {
		// If the timestamp version contains a logical timestamp component that is
		// zero, strip the component. encodeMVCCTimestampToBuf will typically omit
		// the entire logical component in these cases as an optimization, but it
		// does not guarantee to never include a zero logical component.
		// Additionally, we can fall into this case after stripping off other
		// components of the key version earlier on in this function.
		if bytes.Equal(a[withWall:], zeroLogical[:]) {
			a = a[:withWall]
		}
	}
	return a
}

// EngineComparer is a pebble.Comparer object that implements MVCC-specific
// comparator settings for use with Pebble.
var EngineComparer = &pebble.Comparer{
	Compare: EngineKeyCompare,

	Equal: EngineKeyEqual,

	AbbreviatedKey: func(k []byte) uint64 {
		key, ok := GetKeyPartFromEngineKey(k)
		if !ok {
			return 0
		}
		return pebble.DefaultComparer.AbbreviatedKey(key)
	},

	FormatKey: func(k []byte) fmt.Formatter {
		decoded, ok := DecodeEngineKey(k)
		if !ok {
			return mvccKeyFormatter{err: errors.Errorf("invalid encoded engine key: %x", k)}
		}
		if decoded.IsMVCCKey() {
			mvccKey, err := decoded.ToMVCCKey()
			if err != nil {
				return mvccKeyFormatter{err: err}
			}
			return mvccKeyFormatter{key: mvccKey}
		}
		return EngineKeyFormatter{key: decoded}
	},

	Separator: func(dst, a, b []byte) []byte {
		aKey, ok := GetKeyPartFromEngineKey(a)
		if !ok {
			return append(dst, a...)
		}
		bKey, ok := GetKeyPartFromEngineKey(b)
		if !ok {
			return append(dst, a...)
		}
		// If the keys are the same just return a.
		if bytes.Equal(aKey, bKey) {
			return append(dst, a...)
		}
		n := len(dst)
		// Engine key comparison uses bytes.Compare on the roachpb.Key, which is the same semantics as
		// pebble.DefaultComparer, so reuse the latter's Separator implementation.
		dst = pebble.DefaultComparer.Separator(dst, aKey, bKey)
		// Did it pick a separator different than aKey -- if it did not we can't do better than a.
		buf := dst[n:]
		if bytes.Equal(aKey, buf) {
			return append(dst[:n], a...)
		}
		// The separator is > aKey, so we only need to add the sentinel.
		return append(dst, 0)
	},

	Successor: func(dst, a []byte) []byte {
		aKey, ok := GetKeyPartFromEngineKey(a)
		if !ok {
			return append(dst, a...)
		}
		n := len(dst)
		// Engine key comparison uses bytes.Compare on the roachpb.Key, which is the same semantics as
		// pebble.DefaultComparer, so reuse the latter's Successor implementation.
		dst = pebble.DefaultComparer.Successor(dst, aKey)
		// Did it pick a successor different than aKey -- if it did not we can't do better than a.
		buf := dst[n:]
		if bytes.Equal(aKey, buf) {
			return append(dst[:n], a...)
		}
		// The successor is > aKey, so we only need to add the sentinel.
		return append(dst, 0)
	},

	Split: func(k []byte) int {
		key, ok := GetKeyPartFromEngineKey(k)
		if !ok {
			return len(k)
		}
		// Pebble requires that keys generated via a split be comparable with
		// normal encoded engine keys. Encoded engine keys have a suffix
		// indicating the number of bytes of version data. Engine keys without a
		// version have a suffix of 0. We're careful in EncodeKey to make sure
		// that the user-key always has a trailing 0. If there is no version this
		// falls out naturally. If there is a version we prepend a 0 to the
		// encoded version data.
		return len(key) + 1
	},

	Name: "cockroach_comparator",
}

// MVCCMerger is a pebble.Merger object that implements the merge operator used
// by Cockroach.
var MVCCMerger = &pebble.Merger{
	Name: "cockroach_merge_operator",
	Merge: func(_, value []byte) (pebble.ValueMerger, error) {
		res := &MVCCValueMerger{}
		err := res.MergeNewer(value)
		if err != nil {
			return nil, err
		}
		return res, nil
	},
}

// TODO(sumeer): consider removing pebbleTimeBoundPropCollector for the 22.2
// release, since 22.1 nodes will be using BlockPropertyCollectors and there
// shouldn't be significant sstables that do not get rewritten for 6 months
// (we would have no time bound optimization for such sstables). Since it is
// possible for users to upgrade to 22.1 and then to 22.2 in quick succession
// this duration argument is not really valid. To be safe we could wait until
// 23.1.

// pebbleTimeBoundPropCollector implements a property collector for MVCC
// Timestamps. Its behavior matches TimeBoundTblPropCollector in
// table_props.cc.
//
// The handling of timestamps in intents is mildly complicated. Consider:
//
//   a@<meta>   -> <MVCCMetadata: Timestamp=t2>
//   a@t2       -> <value>
//   a@t1       -> <value>
//
// The metadata record (a.k.a. the intent) for a key always sorts first. The
// timestamp field always points to the next record. In this case, the meta
// record contains t2 and the next record is t2. Because of this duplication of
// the timestamp both in the intent and in the timestamped record that
// immediately follows it, we only need to unmarshal the MVCCMetadata if it is
// the last key in the sstable.
// The metadata unmarshaling logic discussed in the previous paragraph is
// flawed (see the comment
// https://github.com/cockroachdb/cockroach/issues/43799#issuecomment-576830234
// and its preceding comment), and we've worked around it as indicated in that
// issue. Due to that workaround, we need not unmarshal the metadata record at
// all.
type pebbleTimeBoundPropCollector struct {
	min, max  []byte
	lastValue []byte
}

func (t *pebbleTimeBoundPropCollector) Add(key pebble.InternalKey, value []byte) error {
	engineKey, ok := DecodeEngineKey(key.UserKey)
	if !ok {
		return errors.Errorf("failed to split engine key")
	}
	if engineKey.IsMVCCKey() && len(engineKey.Version) > 0 {
		t.lastValue = t.lastValue[:0]
		t.updateBounds(engineKey.Version)
	} else {
		t.lastValue = append(t.lastValue[:0], value...)
	}
	return nil
}

func (t *pebbleTimeBoundPropCollector) Finish(userProps map[string]string) error {
	if len(t.lastValue) > 0 {
		// The last record in the sstable was an intent. Unmarshal the metadata and
		// update the bounds with the timestamp it contains.
		meta := &enginepb.MVCCMetadata{}
		if err := protoutil.Unmarshal(t.lastValue, meta); err != nil {
			// We're unable to parse the MVCCMetadata. Fail open by not setting the
			// min/max timestamp properties. This mimics the behavior of
			// TimeBoundTblPropCollector.
			// TODO(petermattis): Return the error here and in C++, see #43422.
			return nil //nolint:returnerrcheck
		}
		if meta.Txn != nil {
			ts := encodeMVCCTimestamp(meta.Timestamp.ToTimestamp())
			t.updateBounds(ts)
		}
	}

	userProps["crdb.ts.min"] = string(t.min)
	userProps["crdb.ts.max"] = string(t.max)
	return nil
}

func (t *pebbleTimeBoundPropCollector) updateBounds(ts []byte) {
	if len(t.min) == 0 || bytes.Compare(ts, t.min) < 0 {
		t.min = append(t.min[:0], ts...)
	}
	if len(t.max) == 0 || bytes.Compare(ts, t.max) > 0 {
		t.max = append(t.max[:0], ts...)
	}
}

func (t *pebbleTimeBoundPropCollector) Name() string {
	// This constant needs to match the one used by the RocksDB version of this
	// table property collector. DO NOT CHANGE.
	return "TimeBoundTblPropCollectorFactory"
}

func (t *pebbleTimeBoundPropCollector) UpdateKeySuffixes(
	oldProps map[string]string, oldSuffix []byte, newSuffix []byte,
) error {
	t.updateBounds(newSuffix)
	return nil
}

var _ sstable.SuffixReplaceableTableCollector = (*pebbleTimeBoundPropCollector)(nil)

// pebbleDeleteRangeCollector is the equivalent table collector as the RocksDB
// DeleteRangeTblPropCollector. Pebble does not require it because Pebble will
// prioritize its own compactions of range tombstones.
type pebbleDeleteRangeCollector struct{}

func (pebbleDeleteRangeCollector) Add(_ pebble.InternalKey, _ []byte) error {
	return nil
}

func (pebbleDeleteRangeCollector) Finish(_ map[string]string) error {
	return nil
}

func (pebbleDeleteRangeCollector) Name() string {
	// This constant needs to match the one used by the RocksDB version of this
	// table property collector. DO NOT CHANGE.
	return "DeleteRangeTblPropCollectorFactory"
}

func (t *pebbleDeleteRangeCollector) UpdateKeySuffixes(
	_ map[string]string, _ []byte, _ []byte,
) error {
	return nil
}

var _ sstable.SuffixReplaceableTableCollector = (*pebbleTimeBoundPropCollector)(nil)

// PebbleTablePropertyCollectors is the list of Pebble TablePropertyCollectors.
var PebbleTablePropertyCollectors = []func() pebble.TablePropertyCollector{
	func() pebble.TablePropertyCollector { return &pebbleTimeBoundPropCollector{} },
	func() pebble.TablePropertyCollector { return &pebbleDeleteRangeCollector{} },
}

// pebbleDataBlockMVCCTimeIntervalCollector provides an implementation of
// pebble.DataBlockIntervalCollector that is used to construct a
// pebble.BlockPropertyCollector. This provides per-block filtering, which
// also gets aggregated to the sstable-level and filters out sstables. It must
// only be used for MVCCKeyIterKind iterators, since it will ignore
// blocks/sstables that contain intents (and any other key that is not a real
// MVCC key).
type pebbleDataBlockMVCCTimeIntervalCollector struct {
	// min, max are the encoded timestamps.
	min, max []byte
}

var _ sstable.DataBlockIntervalCollector = &pebbleDataBlockMVCCTimeIntervalCollector{}
var _ sstable.SuffixReplaceableBlockCollector = (*pebbleDataBlockMVCCTimeIntervalCollector)(nil)

func (tc *pebbleDataBlockMVCCTimeIntervalCollector) Add(key pebble.InternalKey, _ []byte) error {
	return tc.add(key.UserKey)
}

// add collects the given slice in the collector. The slice may be an entire
// encoded MVCC key, or the bare suffix of an encoded key.
func (tc *pebbleDataBlockMVCCTimeIntervalCollector) add(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	// Last byte is the version length + 1 when there is a version,
	// else it is 0.
	versionLen := int(b[len(b)-1])
	if versionLen == 0 {
		// This is not an MVCC key that we can collect.
		return nil
	}
	// prefixPartEnd points to the sentinel byte, unless this is a bare suffix, in
	// which case the index is -1.
	prefixPartEnd := len(b) - 1 - versionLen
	// Sanity check: the index should be >= -1. Additionally, if the index is >=
	// 0, it should point to the sentinel byte, as this is a full EngineKey.
	if prefixPartEnd < -1 || (prefixPartEnd >= 0 && b[prefixPartEnd] != sentinel) {
		return errors.Errorf("invalid key %s", roachpb.Key(b).String())
	}
	// We don't need the last byte (the version length).
	versionLen--
	// Only collect if this looks like an MVCC timestamp.
	if versionLen == engineKeyVersionWallTimeLen ||
		versionLen == engineKeyVersionWallAndLogicalTimeLen ||
		versionLen == engineKeyVersionWallLogicalAndSyntheticTimeLen {
		// INVARIANT: -1 <= prefixPartEnd < len(b) - 1.
		// Version consists of the bytes after the sentinel and before the length.
		b = b[prefixPartEnd+1 : len(b)-1]
		// Lexicographic comparison on the encoded timestamps is equivalent to the
		// comparison on decoded timestamps, so delay decoding.
		if len(tc.min) == 0 || bytes.Compare(b, tc.min) < 0 {
			tc.min = append(tc.min[:0], b...)
		}
		if len(tc.max) == 0 || bytes.Compare(b, tc.max) > 0 {
			tc.max = append(tc.max[:0], b...)
		}
	}
	return nil
}

func decodeWallTime(ts []byte) uint64 {
	return binary.BigEndian.Uint64(ts[0:engineKeyVersionWallTimeLen])
}

func (tc *pebbleDataBlockMVCCTimeIntervalCollector) FinishDataBlock() (
	lower uint64,
	upper uint64,
	err error,
) {
	if len(tc.min) == 0 {
		// No calls to Add that contained a timestamped key.
		return 0, 0, nil
	}
	// Construct a [lower, upper) walltime that will contain all the
	// hlc.Timestamps in this block.
	lower = decodeWallTime(tc.min)
	// Remember that we have to reset tc.min and tc.max to get ready for the
	// next data block, as specified in the DataBlockIntervalCollector interface
	// help and help too.
	tc.min = tc.min[:0]
	// The actual value encoded into walltime is an int64, so +1 will not
	// overflow.
	upper = decodeWallTime(tc.max) + 1
	tc.max = tc.max[:0]
	if lower >= upper {
		return 0, 0,
			errors.Errorf("corrupt timestamps lower %d >= upper %d", lower, upper)
	}
	return lower, upper, nil
}

func (tc *pebbleDataBlockMVCCTimeIntervalCollector) UpdateKeySuffixes(
	_ []byte, _, newSuffix []byte,
) error {
	return tc.add(newSuffix)
}

const mvccWallTimeIntervalCollector = "MVCCTimeInterval"

// PebbleBlockPropertyCollectors is the list of functions to construct
// BlockPropertyCollectors.
var PebbleBlockPropertyCollectors = []func() pebble.BlockPropertyCollector{
	func() pebble.BlockPropertyCollector {
		return sstable.NewBlockIntervalCollector(
			mvccWallTimeIntervalCollector,
			&pebbleDataBlockMVCCTimeIntervalCollector{}, /* points */
			nil, /* ranges */
		)
	},
}

// DefaultPebbleOptions returns the default pebble options.
func DefaultPebbleOptions() *pebble.Options {
	// In RocksDB, the concurrency setting corresponds to both flushes and
	// compactions. In Pebble, there is always a slot for a flush, and
	// compactions are counted separately.
	maxConcurrentCompactions := rocksdbConcurrency - 1
	if maxConcurrentCompactions < 1 {
		maxConcurrentCompactions = 1
	}

	opts := &pebble.Options{
		Comparer:                    EngineComparer,
		FS:                          vfs.Default,
		L0CompactionThreshold:       2,
		L0StopWritesThreshold:       1000,
		LBaseMaxBytes:               64 << 20, // 64 MB
		Levels:                      make([]pebble.LevelOptions, 7),
		MaxConcurrentCompactions:    maxConcurrentCompactions,
		MemTableSize:                64 << 20, // 64 MB
		MemTableStopWritesThreshold: 4,
		Merger:                      MVCCMerger,
		TablePropertyCollectors:     PebbleTablePropertyCollectors,
		BlockPropertyCollectors:     PebbleBlockPropertyCollectors,
	}
	// Used for experimental MVCC range tombstones.
	opts.Experimental.RangeKeys = new(pebble.RangeKeysArena)
	// Automatically flush 10s after the first range tombstone is added to a
	// memtable. This ensures that we can reclaim space even when there's no
	// activity on the database generating flushes.
	opts.Experimental.DeleteRangeFlushDelay = 10 * time.Second
	// Enable deletion pacing. This helps prevent disk slowness events on some
	// SSDs, that kick off an expensive GC if a lot of files are deleted at
	// once.
	opts.Experimental.MinDeletionRate = 128 << 20 // 128 MB
	// Validate min/max keys in each SSTable when performing a compaction. This
	// serves as a simple protection against corruption or programmer-error in
	// Pebble.
	opts.Experimental.KeyValidationFunc = func(userKey []byte) error {
		engineKey, ok := DecodeEngineKey(userKey)
		if !ok {
			return errors.Newf("key %s could not be decoded as an EngineKey", string(userKey))
		}
		if err := engineKey.Validate(); err != nil {
			return err
		}
		return nil
	}

	for i := 0; i < len(opts.Levels); i++ {
		l := &opts.Levels[i]
		l.BlockSize = 32 << 10       // 32 KB
		l.IndexBlockSize = 256 << 10 // 256 KB
		l.FilterPolicy = bloom.FilterPolicy(10)
		l.FilterType = pebble.TableFilter
		if i > 0 {
			l.TargetFileSize = opts.Levels[i-1].TargetFileSize * 2
		}
		l.EnsureDefaults()
	}

	return opts
}

// wrapFilesystemMiddleware wraps the Option's vfs.FS with disk-health checking
// and ENOSPC detection. It mutates the provided options to set the FS and
// returns a Closer that should be invoked when the filesystem will no longer be
// used.
func wrapFilesystemMiddleware(opts *pebble.Options) io.Closer {
	// Set disk-health check interval to min(5s, maxSyncDurationDefault). This
	// is mostly to ease testing; the default of 5s is too infrequent to test
	// conveniently. See the disk-stalled roachtest for an example of how this
	// is used.
	diskHealthCheckInterval := 5 * time.Second
	if diskHealthCheckInterval.Seconds() > maxSyncDurationDefault.Seconds() {
		diskHealthCheckInterval = maxSyncDurationDefault
	}
	// Instantiate a file system with disk health checking enabled. This FS
	// wraps the filesystem with a layer that times all write-oriented
	// operations.
	var closer io.Closer
	opts.FS, closer = vfs.WithDiskHealthChecks(opts.FS, diskHealthCheckInterval,
		func(name string, duration time.Duration) {
			opts.EventListener.DiskSlow(pebble.DiskSlowInfo{
				Path:     name,
				Duration: duration,
			})
		})
	// If we encounter ENOSPC, exit with an informative exit code.
	opts.FS = vfs.OnDiskFull(opts.FS, func() {
		exit.WithCode(exit.DiskFull())
	})
	return closer
}

type pebbleLogger struct {
	ctx   context.Context
	depth int
}

func (l pebbleLogger) Infof(format string, args ...interface{}) {
	log.Storage.InfofDepth(l.ctx, l.depth, format, args...)
}

func (l pebbleLogger) Fatalf(format string, args ...interface{}) {
	log.Storage.FatalfDepth(l.ctx, l.depth, format, args...)
}

// PebbleConfig holds all configuration parameters and knobs used in setting up
// a new Pebble instance.
type PebbleConfig struct {
	// StorageConfig contains storage configs for all storage engines.
	// A non-nil cluster.Settings must be provided in the StorageConfig for a
	// Pebble instance that will be used to write intents.
	base.StorageConfig
	// Pebble specific options.
	Opts *pebble.Options
	// Temporary option while there exist file descriptor leaks. See the
	// DisableFilesystemMiddlewareTODO ConfigOption that sets this, and #81389.
	DisableFilesystemMiddlewareTODO bool
}

// EncryptionStatsHandler provides encryption related stats.
type EncryptionStatsHandler interface {
	// Returns a serialized enginepbccl.EncryptionStatus.
	GetEncryptionStatus() ([]byte, error)
	// Returns a serialized enginepbccl.DataKeysRegistry, scrubbed of key contents.
	GetDataKeysRegistry() ([]byte, error)
	// Returns the ID of the active data key, or "plain" if none.
	GetActiveDataKeyID() (string, error)
	// Returns the enum value of the encryption type.
	GetActiveStoreKeyType() int32
	// Returns the KeyID embedded in the serialized EncryptionSettings.
	GetKeyIDFromSettings(settings []byte) (string, error)
}

// Pebble is a wrapper around a Pebble database instance.
type Pebble struct {
	db *pebble.DB

	closed      bool
	readOnly    bool
	path        string
	auxDir      string
	ballastPath string
	ballastSize int64
	maxSize     int64
	attrs       roachpb.Attributes
	properties  roachpb.StoreProperties
	// settings must be non-nil if this Pebble instance will be used to write
	// intents.
	settings     *cluster.Settings
	encryption   *EncryptionEnv
	fileRegistry *PebbleFileRegistry

	// Stats updated by pebble.EventListener invocations, and returned in
	// GetMetrics. Updated and retrieved atomically.
	writeStallCount int64
	diskSlowCount   int64
	diskStallCount  int64

	// Relevant options copied over from pebble.Options.
	fs            vfs.FS
	unencryptedFS vfs.FS
	logger        pebble.Logger
	eventListener *pebble.EventListener
	mu            struct {
		// This mutex is the lowest in any lock ordering.
		syncutil.Mutex
		flushCompletedCallback func()
	}

	// supportsRangeKeys is 1 if the database supports range keys. It must
	// be accessed atomically.
	//
	// TODO(erikgrinaker): Remove this after 22.2 when all databases support it.
	supportsRangeKeys int32

	// closer is populated when the database is opened. The closer is associated
	// with the filesyetem
	closer io.Closer

	wrappedIntentWriter intentDemuxWriter

	storeIDPebbleLog *base.StoreIDContainer
}

// EncryptionEnv describes the encryption-at-rest environment, providing
// access to a filesystem with on-the-fly encryption.
type EncryptionEnv struct {
	// Closer closes the encryption-at-rest environment. Once the
	// environment is closed, the environment's VFS may no longer be
	// used.
	Closer io.Closer
	// FS provides the encrypted virtual filesystem. New files are
	// transparently encrypted.
	FS vfs.FS
	// StatsHandler exposes encryption-at-rest state for observability.
	StatsHandler EncryptionStatsHandler
}

var _ Engine = &Pebble{}

// NewEncryptedEnvFunc creates an encrypted environment and returns the vfs.FS to use for reading
// and writing data. This should be initialized by calling engineccl.Init() before calling
// NewPebble(). The optionBytes is a binary serialized baseccl.EncryptionOptions, so that non-CCL
// code does not depend on CCL code.
var NewEncryptedEnvFunc func(fs vfs.FS, fr *PebbleFileRegistry, dbDir string, readOnly bool, optionBytes []byte) (*EncryptionEnv, error)

// StoreIDSetter is used to set the store id in the log.
type StoreIDSetter interface {
	// SetStoreID can be used to atomically set the store
	// id as a tag in the pebble logs. Once set, the store id will be visible
	// in pebble logs in cockroach.
	SetStoreID(ctx context.Context, storeID int32)
}

// SetStoreID adds the store id to pebble logs.
func (p *Pebble) SetStoreID(ctx context.Context, storeID int32) {
	if p == nil {
		return
	}
	if p.storeIDPebbleLog == nil {
		return
	}
	p.storeIDPebbleLog.Set(ctx, storeID)
}

// ResolveEncryptedEnvOptions fills in cfg.Opts.FS with an encrypted vfs if this
// store has encryption-at-rest enabled. Also returns the associated file
// registry and EncryptionStatsHandler.
func ResolveEncryptedEnvOptions(cfg *PebbleConfig) (*PebbleFileRegistry, *EncryptionEnv, error) {
	fileRegistry := &PebbleFileRegistry{FS: cfg.Opts.FS, DBDir: cfg.Dir, ReadOnly: cfg.Opts.ReadOnly}
	if cfg.UseFileRegistry {
		if err := fileRegistry.Load(); err != nil {
			return nil, nil, err
		}
	} else {
		if err := fileRegistry.CheckNoRegistryFile(); err != nil {
			return nil, nil, fmt.Errorf("encryption was used on this store before, but no encryption flags " +
				"specified. You need a CCL build and must fully specify the --enterprise-encryption flag")
		}
		fileRegistry = nil
	}

	var env *EncryptionEnv
	if cfg.IsEncrypted() {
		// Encryption is enabled.
		if !cfg.UseFileRegistry {
			return nil, nil, fmt.Errorf("file registry is needed to support encryption")
		}
		if NewEncryptedEnvFunc == nil {
			return nil, nil, fmt.Errorf("encryption is enabled but no function to create the encrypted env")
		}
		var err error
		env, err = NewEncryptedEnvFunc(
			cfg.Opts.FS,
			fileRegistry,
			cfg.Dir,
			cfg.Opts.ReadOnly,
			cfg.EncryptionOptions,
		)
		if err != nil {
			return nil, nil, err
		}
		// TODO(jackson): Should this just return an EncryptionEnv,
		// rather than mutating cfg.Opts?
		cfg.Opts.FS = env.FS
	}
	return fileRegistry, env, nil
}

// NewPebble creates a new Pebble instance, at the specified path.
func NewPebble(ctx context.Context, cfg PebbleConfig) (p *Pebble, err error) {
	// pebble.Open also calls EnsureDefaults, but only after doing a clone. Call
	// EnsureDefaults beforehand so we have a matching cfg here for when we save
	// cfg.FS and cfg.ReadOnly later on.
	if cfg.Opts == nil {
		cfg.Opts = DefaultPebbleOptions()
	}

	// Initialize the FS, wrapping it with disk health-checking and
	// ENOSPC-detection.
	var filesystemCloser io.Closer
	if !cfg.DisableFilesystemMiddlewareTODO {
		filesystemCloser = wrapFilesystemMiddleware(cfg.Opts)
		defer func() {
			if err != nil {
				filesystemCloser.Close()
			}
		}()
	}

	cfg.Opts.EnsureDefaults()
	cfg.Opts.ErrorIfNotExists = cfg.MustExist
	if settings := cfg.Settings; settings != nil {
		cfg.Opts.WALMinSyncInterval = func() time.Duration {
			return minWALSyncInterval.Get(&settings.SV)
		}
	}

	auxDir := cfg.Opts.FS.PathJoin(cfg.Dir, base.AuxiliaryDir)
	if err := cfg.Opts.FS.MkdirAll(auxDir, 0755); err != nil {
		return nil, err
	}
	ballastPath := base.EmergencyBallastFile(cfg.Opts.FS.PathJoin, cfg.Dir)

	// For some purposes, we want to always use an unencrypted
	// filesystem. The call below to ResolveEncryptedEnvOptions will
	// replace cfg.Opts.FS with a VFS wrapped with encryption-at-rest if
	// necessary. Before we do that, save a handle on the unencrypted
	// FS for those that need it. Some call sites need the unencrypted
	// FS for the purpose of atomic renames.
	unencryptedFS := cfg.Opts.FS
	fileRegistry, env, err := ResolveEncryptedEnvOptions(&cfg)
	if err != nil {
		return nil, err
	}

	// The context dance here is done so that we have a clean context without
	// timeouts that has a copy of the log tags.
	logCtx := logtags.WithTags(context.Background(), logtags.FromContext(ctx))
	logCtx = logtags.AddTag(logCtx, "pebble", nil)
	// The store id, could not necessarily be determined when this function
	// is called. Therefore, we use a container for the store id.
	storeIDContainer := &base.StoreIDContainer{}
	logCtx = logtags.AddTag(logCtx, "s", storeIDContainer)

	if cfg.Opts.Logger == nil {
		cfg.Opts.Logger = pebbleLogger{
			ctx:   logCtx,
			depth: 1,
		}
	}

	// Establish the emergency ballast if we can. If there's not sufficient
	// disk space, the ballast will be reestablished from Capacity when the
	// store's capacity is queried periodically.
	if !cfg.Opts.ReadOnly {
		du, err := unencryptedFS.GetDiskUsage(cfg.Dir)
		// If the FS is an in-memory FS, GetDiskUsage returns
		// vfs.ErrUnsupported and we skip ballast creation.
		if err != nil && !errors.Is(err, vfs.ErrUnsupported) {
			return nil, errors.Wrap(err, "retrieving disk usage")
		} else if err == nil {
			resized, err := maybeEstablishBallast(unencryptedFS, ballastPath, cfg.BallastSize, du)
			if err != nil {
				return nil, errors.Wrap(err, "resizing ballast")
			}
			if resized {
				cfg.Opts.Logger.Infof("resized ballast %s to size %s",
					ballastPath, humanizeutil.IBytes(cfg.BallastSize))
			}
		}
	}

	storeProps := computeStoreProperties(ctx, cfg.Dir, cfg.Opts.ReadOnly, env != nil /* encryptionEnabled */)

	p = &Pebble{
		readOnly:         cfg.Opts.ReadOnly,
		path:             cfg.Dir,
		auxDir:           auxDir,
		ballastPath:      ballastPath,
		ballastSize:      cfg.BallastSize,
		maxSize:          cfg.MaxSize,
		attrs:            cfg.Attrs,
		properties:       storeProps,
		settings:         cfg.Settings,
		encryption:       env,
		fileRegistry:     fileRegistry,
		fs:               cfg.Opts.FS,
		unencryptedFS:    unencryptedFS,
		logger:           cfg.Opts.Logger,
		storeIDPebbleLog: storeIDContainer,
		closer:           filesystemCloser,
	}
	cfg.Opts.EventListener = pebble.TeeEventListener(
		pebble.MakeLoggingEventListener(pebbleLogger{
			ctx:   logCtx,
			depth: 2, // skip over the EventListener stack frame
		}),
		p.makeMetricEtcEventListener(ctx),
	)
	p.eventListener = &cfg.Opts.EventListener
	p.wrappedIntentWriter = wrapIntentWriter(p)

	// Read the current store cluster version.
	storeClusterVersion, err := getMinVersion(unencryptedFS, cfg.Dir)
	if err != nil {
		return nil, err
	}

	db, err := pebble.Open(cfg.StorageConfig.Dir, cfg.Opts)
	if err != nil {
		return nil, err
	}
	p.db = db

	if storeClusterVersion != (roachpb.Version{}) {
		// The storage engine performs its own internal migrations
		// through the setting of the store cluster version. When
		// storage's min version is set, SetMinVersion writes to disk to
		// commit to the new store cluster version. Then it idempotently
		// applies any internal storage engine migrations necessitated
		// or enabled by the new store cluster version. If we crash
		// after committing the new store cluster version but before
		// applying the internal migrations, we're left in an in-between
		// state.
		//
		// To account for this, after the engine is open,
		// unconditionally set the min cluster version again. If any
		// storage engine state has not been updated, the call to
		// SetMinVersion will update it.  If all storage engine state is
		// already updated, SetMinVersion is a noop.
		if err := p.SetMinVersion(storeClusterVersion); err != nil {
			p.Close()
			return nil, err
		}
	}
	if p.db.FormatMajorVersion() >= pebble.FormatRangeKeys {
		atomic.StoreInt32(&p.supportsRangeKeys, 1)
	}

	return p, nil
}

func (p *Pebble) makeMetricEtcEventListener(ctx context.Context) pebble.EventListener {
	return pebble.EventListener{
		WriteStallBegin: func(info pebble.WriteStallBeginInfo) {
			atomic.AddInt64(&p.writeStallCount, 1)
		},
		DiskSlow: func(info pebble.DiskSlowInfo) {
			maxSyncDuration := maxSyncDurationDefault
			fatalOnExceeded := maxSyncDurationFatalOnExceededDefault
			if p.settings != nil {
				maxSyncDuration = MaxSyncDuration.Get(&p.settings.SV)
				fatalOnExceeded = MaxSyncDurationFatalOnExceeded.Get(&p.settings.SV)
			}
			if info.Duration.Seconds() >= maxSyncDuration.Seconds() {
				atomic.AddInt64(&p.diskStallCount, 1)
				// Note that the below log messages go to the main cockroach log, not
				// the pebble-specific log.
				if fatalOnExceeded {
					log.Fatalf(ctx, "disk stall detected: pebble unable to write to %s in %.2f seconds",
						info.Path, redact.Safe(info.Duration.Seconds()))
				} else {
					log.Errorf(ctx, "disk stall detected: pebble unable to write to %s in %.2f seconds",
						info.Path, redact.Safe(info.Duration.Seconds()))
				}
				return
			}
			atomic.AddInt64(&p.diskSlowCount, 1)
		},
		FlushEnd: func(info pebble.FlushInfo) {
			if info.Err != nil {
				return
			}
			p.mu.Lock()
			cb := p.mu.flushCompletedCallback
			p.mu.Unlock()
			if cb != nil {
				cb()
			}
		},
	}
}

func (p *Pebble) String() string {
	dir := p.path
	if dir == "" {
		dir = "<in-mem>"
	}
	attrs := p.attrs.String()
	if attrs == "" {
		attrs = "<no-attributes>"
	}
	return fmt.Sprintf("%s=%s", attrs, dir)
}

// Close implements the Engine interface.
func (p *Pebble) Close() {
	if p.closed {
		p.logger.Infof("closing unopened pebble instance")
		return
	}
	p.closed = true
	_ = p.db.Close()
	if p.fileRegistry != nil {
		_ = p.fileRegistry.Close()
	}
	if p.encryption != nil {
		_ = p.encryption.Closer.Close()
	}
	if p.closer != nil {
		_ = p.closer.Close()
	}
}

// Closed implements the Engine interface.
func (p *Pebble) Closed() bool {
	return p.closed
}

// ExportMVCCToSst is part of the engine.Reader interface.
func (p *Pebble) ExportMVCCToSst(
	ctx context.Context, exportOptions ExportOptions, dest io.Writer,
) (roachpb.BulkOpSummary, roachpb.Key, hlc.Timestamp, error) {
	r := wrapReader(p)
	// Doing defer r.Free() does not inline.
	summary, k, err := pebbleExportToSst(ctx, p.settings, r, exportOptions, dest)
	r.Free()
	return summary, k.Key, k.Timestamp, err
}

// MVCCGet implements the Engine interface.
func (p *Pebble) MVCCGet(key MVCCKey) ([]byte, error) {
	return mvccGetHelper(key, p)
}

func mvccGetHelper(key MVCCKey, reader wrappableReader) ([]byte, error) {
	if len(key.Key) == 0 {
		return nil, emptyKeyError()
	}
	r := wrapReader(reader)
	// Doing defer r.Free() does not inline.
	v, err := r.MVCCGet(key)
	r.Free()
	return v, err
}

func (p *Pebble) rawMVCCGet(key []byte) ([]byte, error) {
	ret, closer, err := p.db.Get(key)
	if closer != nil {
		retCopy := make([]byte, len(ret))
		copy(retCopy, ret)
		ret = retCopy
		closer.Close()
	}
	if errors.Is(err, pebble.ErrNotFound) || len(ret) == 0 {
		return nil, nil
	}
	return ret, err
}

// MVCCGetProto implements the Engine interface.
func (p *Pebble) MVCCGetProto(
	key MVCCKey, msg protoutil.Message,
) (ok bool, keyBytes, valBytes int64, err error) {
	return pebbleGetProto(p, key, msg)
}

// MVCCIterate implements the Engine interface.
func (p *Pebble) MVCCIterate(
	start, end roachpb.Key, iterKind MVCCIterKind, f func(MVCCKeyValue) error,
) error {
	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		err := iterateOnReader(r, start, end, iterKind, f)
		r.Free()
		return err
	}
	return iterateOnReader(p, start, end, iterKind, f)
}

// NewMVCCIterator implements the Engine interface.
func (p *Pebble) NewMVCCIterator(iterKind MVCCIterKind, opts IterOptions) MVCCIterator {
	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		iter := r.NewMVCCIterator(iterKind, opts)
		r.Free()
		if util.RaceEnabled {
			iter = wrapInUnsafeIter(iter)
		}
		return iter
	}

	iter := newPebbleIterator(p.db, nil, opts, StandardDurability, p.SupportsRangeKeys())
	if iter == nil {
		panic("couldn't create a new iterator")
	}
	if util.RaceEnabled {
		return wrapInUnsafeIter(iter)
	}
	return iter
}

// NewEngineIterator implements the Engine interface.
func (p *Pebble) NewEngineIterator(opts IterOptions) EngineIterator {
	if opts.KeyTypes != IterKeyTypePointsOnly {
		panic("EngineIterator does not support range keys")
	}
	iter := newPebbleIterator(p.db, nil, opts, StandardDurability, p.SupportsRangeKeys())
	if iter == nil {
		panic("couldn't create a new iterator")
	}
	return iter
}

// ConsistentIterators implements the Engine interface.
func (p *Pebble) ConsistentIterators() bool {
	return false
}

// SupportsRangeKeys implements the Engine interface.
func (p *Pebble) SupportsRangeKeys() bool {
	return atomic.LoadInt32(&p.supportsRangeKeys) == 1
}

// PinEngineStateForIterators implements the Engine interface.
func (p *Pebble) PinEngineStateForIterators() error {
	return errors.AssertionFailedf(
		"PinEngineStateForIterators must not be called when ConsistentIterators returns false")
}

// ApplyBatchRepr implements the Engine interface.
func (p *Pebble) ApplyBatchRepr(repr []byte, sync bool) error {
	// batch.SetRepr takes ownership of the underlying slice, so make a copy.
	reprCopy := make([]byte, len(repr))
	copy(reprCopy, repr)

	batch := p.db.NewBatch()
	if err := batch.SetRepr(reprCopy); err != nil {
		return err
	}

	opts := pebble.NoSync
	if sync {
		opts = pebble.Sync
	}
	return batch.Commit(opts)
}

// ClearMVCC implements the Engine interface.
func (p *Pebble) ClearMVCC(key MVCCKey) error {
	if key.Timestamp.IsEmpty() {
		panic("ClearMVCC timestamp is empty")
	}
	return p.clear(key)
}

// ClearUnversioned implements the Engine interface.
func (p *Pebble) ClearUnversioned(key roachpb.Key) error {
	return p.clear(MVCCKey{Key: key})
}

// ClearIntent implements the Engine interface.
func (p *Pebble) ClearIntent(key roachpb.Key, txnDidNotUpdateMeta bool, txnUUID uuid.UUID) error {
	_, err := p.wrappedIntentWriter.ClearIntent(key, txnDidNotUpdateMeta, txnUUID, nil)
	return err
}

// ClearEngineKey implements the Engine interface.
func (p *Pebble) ClearEngineKey(key EngineKey) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.Delete(key.Encode(), pebble.Sync)
}

func (p *Pebble) clear(key MVCCKey) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.Delete(EncodeMVCCKey(key), pebble.Sync)
}

// SingleClearEngineKey implements the Engine interface.
func (p *Pebble) SingleClearEngineKey(key EngineKey) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.SingleDelete(key.Encode(), pebble.Sync)
}

// ClearRawRange implements the Engine interface.
func (p *Pebble) ClearRawRange(start, end roachpb.Key) error {
	startKey, endKey := EncodeMVCCKey(MVCCKey{Key: start}), EncodeMVCCKey(MVCCKey{Key: end})
	if err := p.db.DeleteRange(startKey, endKey, pebble.Sync); err != nil {
		return err
	}
	return p.ExperimentalClearAllMVCCRangeKeys(start, end)
}

// ClearMVCCRange implements the Engine interface.
func (p *Pebble) ClearMVCCRange(start, end roachpb.Key) error {
	_, err := p.wrappedIntentWriter.ClearMVCCRange(start, end, nil)
	return err
}

// ClearMVCCVersions implements the Engine interface.
func (p *Pebble) ClearMVCCVersions(start, end MVCCKey) error {
	return p.db.DeleteRange(EncodeMVCCKey(start), EncodeMVCCKey(end), pebble.Sync)
}

// ClearMVCCIteratorRange implements the Engine interface.
func (p *Pebble) ClearMVCCIteratorRange(start, end roachpb.Key) error {
	// Write all the tombstones in one batch.
	batch := p.NewUnindexedBatch(false /* writeOnly */)
	defer batch.Close()

	if err := batch.ClearMVCCIteratorRange(start, end); err != nil {
		return err
	}
	return batch.Commit(true)
}

// ExperimentalClearMVCCRangeKey implements the Engine interface.
func (p *Pebble) ExperimentalClearMVCCRangeKey(rangeKey MVCCRangeKey) error {
	if !p.SupportsRangeKeys() {
		// These databases cannot contain range keys, so clearing is a noop.
		return nil
	}
	if err := rangeKey.Validate(); err != nil {
		return err
	}
	return p.db.Experimental().RangeKeyUnset(
		EncodeMVCCKeyPrefix(rangeKey.StartKey),
		EncodeMVCCKeyPrefix(rangeKey.EndKey),
		EncodeMVCCTimestampSuffix(rangeKey.Timestamp),
		pebble.Sync)
}

// ExperimentalClearAllMVCCRangeKeys implements the Engine interface.
func (p *Pebble) ExperimentalClearAllMVCCRangeKeys(start, end roachpb.Key) error {
	if !p.SupportsRangeKeys() {
		return nil // noop
	}
	rangeKey := MVCCRangeKey{StartKey: start, EndKey: end, Timestamp: hlc.MinTimestamp}
	if err := rangeKey.Validate(); err != nil {
		return err
	}
	// Look for any range keys in the span before dropping a range tombstone, and
	// use the smallest possible span that covers them, to avoid dropping range
	// tombstones across unnecessary spans. We don't worry about races here,
	// because this is a non-MVCC operation where the caller must guarantee
	// appropriate isolation.
	clearFrom, clearTo, err := pebbleFindRangeKeySpan(p.db,
		EncodeMVCCKeyPrefix(start), EncodeMVCCKeyPrefix(end))
	if err != nil {
		return err
	} else if clearFrom == nil || clearTo == nil {
		return nil
	}
	return p.db.Experimental().RangeKeyDelete(clearFrom, clearTo, pebble.Sync)
}

// ExperimentalPutMVCCRangeKey implements the Engine interface.
func (p *Pebble) ExperimentalPutMVCCRangeKey(rangeKey MVCCRangeKey, value MVCCValue) error {
	if !p.SupportsRangeKeys() {
		return errors.Errorf("range keys not supported by Pebble database version %s",
			p.db.FormatMajorVersion())
	}
	if err := rangeKey.Validate(); err != nil {
		return err
	}
	// NB: all MVCC APIs currently assume all range keys are range tombstones.
	if !value.IsTombstone() {
		return errors.New("range keys can only be MVCC range tombstones")
	}
	valueBytes, err := EncodeMVCCValue(value)
	if err != nil {
		return errors.Wrapf(err, "failed to encode MVCC value for range key %s", rangeKey)
	}
	return p.db.Experimental().RangeKeySet(
		EncodeMVCCKeyPrefix(rangeKey.StartKey),
		EncodeMVCCKeyPrefix(rangeKey.EndKey),
		EncodeMVCCTimestampSuffix(rangeKey.Timestamp),
		valueBytes,
		pebble.Sync)
}

// Merge implements the Engine interface.
func (p *Pebble) Merge(key MVCCKey, value []byte) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.Merge(EncodeMVCCKey(key), value, pebble.Sync)
}

// PutMVCC implements the Engine interface.
func (p *Pebble) PutMVCC(key MVCCKey, value MVCCValue) error {
	if key.Timestamp.IsEmpty() {
		panic("PutMVCC timestamp is empty")
	}
	encValue, err := EncodeMVCCValue(value)
	if err != nil {
		return err
	}
	return p.put(key, encValue)
}

// PutRawMVCC implements the Engine interface.
func (p *Pebble) PutRawMVCC(key MVCCKey, value []byte) error {
	if key.Timestamp.IsEmpty() {
		panic("PutRawMVCC timestamp is empty")
	}
	return p.put(key, value)
}

// PutUnversioned implements the Engine interface.
func (p *Pebble) PutUnversioned(key roachpb.Key, value []byte) error {
	return p.put(MVCCKey{Key: key}, value)
}

// PutIntent implements the Engine interface.
func (p *Pebble) PutIntent(
	ctx context.Context, key roachpb.Key, value []byte, txnUUID uuid.UUID,
) error {
	_, err := p.wrappedIntentWriter.PutIntent(ctx, key, value, txnUUID, nil)
	return err
}

// PutEngineKey implements the Engine interface.
func (p *Pebble) PutEngineKey(key EngineKey, value []byte) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.Set(key.Encode(), value, pebble.Sync)
}

func (p *Pebble) put(key MVCCKey, value []byte) error {
	if len(key.Key) == 0 {
		return emptyKeyError()
	}
	return p.db.Set(EncodeMVCCKey(key), value, pebble.Sync)
}

// LogData implements the Engine interface.
func (p *Pebble) LogData(data []byte) error {
	return p.db.LogData(data, pebble.Sync)
}

// LogLogicalOp implements the Engine interface.
func (p *Pebble) LogLogicalOp(op MVCCLogicalOpType, details MVCCLogicalOpDetails) {
	// No-op. Logical logging disabled.
}

// localTimestampsEnabled controls whether local timestamps are written in MVCC
// values. A true setting is also gated on clusterversion.LocalTimestamps. After
// all nodes in a cluster are at or beyond clusterversion.LocalTimestamps,
// different nodes will see the version state transition at different times.
// Nodes that have not yet seen the transition may remove the local timestamp
// from an intent that has one during intent resolution. This will not cause
// problems.
//
// TODO(nvanbenschoten): remove this cluster setting and its associated plumbing
// when removing the cluster version, once we're confident in the efficacy and
// stability of local timestamps.
var localTimestampsEnabled = settings.RegisterBoolSetting(
	settings.SystemOnly,
	"storage.transaction.local_timestamps.enabled",
	"if enabled, MVCC keys will be written with local timestamps",
	true,
)

func shouldWriteLocalTimestamps(ctx context.Context, settings *cluster.Settings) bool {
	if !localTimestampsEnabled.Get(&settings.SV) {
		// Not enabled.
		return false
	}
	ver := settings.Version.ActiveVersionOrEmpty(ctx)
	if ver == (clusterversion.ClusterVersion{}) {
		// Some tests fail to configure settings. In these cases, assume that it
		// is safe to write local timestamps.
		return true
	}
	return ver.IsActive(clusterversion.LocalTimestamps)
}

// ShouldWriteLocalTimestamps implements the Writer interface.
func (p *Pebble) ShouldWriteLocalTimestamps(ctx context.Context) bool {
	// This is not fast. Pebble should not be used by writers that want
	// performance. They should use pebbleBatch.
	return shouldWriteLocalTimestamps(ctx, p.settings)
}

// Attrs implements the Engine interface.
func (p *Pebble) Attrs() roachpb.Attributes {
	return p.attrs
}

// Properties implements the Engine interface.
func (p *Pebble) Properties() roachpb.StoreProperties {
	return p.properties
}

// Capacity implements the Engine interface.
func (p *Pebble) Capacity() (roachpb.StoreCapacity, error) {
	dir := p.path
	if dir != "" {
		var err error
		// Eval directory if it is a symbolic links.
		if dir, err = filepath.EvalSymlinks(dir); err != nil {
			return roachpb.StoreCapacity{}, err
		}
	}
	du, err := p.unencryptedFS.GetDiskUsage(dir)
	if errors.Is(err, vfs.ErrUnsupported) {
		// This is an in-memory instance. Pretend we're empty since we
		// don't know better and only use this for testing. Using any
		// part of the actual file system here can throw off allocator
		// rebalancing in a hard-to-trace manner. See #7050.
		return roachpb.StoreCapacity{
			Capacity:  p.maxSize,
			Available: p.maxSize,
		}, nil
	} else if err != nil {
		return roachpb.StoreCapacity{}, err
	}

	if du.TotalBytes > math.MaxInt64 {
		return roachpb.StoreCapacity{}, fmt.Errorf("unsupported disk size %s, max supported size is %s",
			humanize.IBytes(du.TotalBytes), humanizeutil.IBytes(math.MaxInt64))
	}
	if du.AvailBytes > math.MaxInt64 {
		return roachpb.StoreCapacity{}, fmt.Errorf("unsupported disk size %s, max supported size is %s",
			humanize.IBytes(du.AvailBytes), humanizeutil.IBytes(math.MaxInt64))
	}
	fsuTotal := int64(du.TotalBytes)
	fsuAvail := int64(du.AvailBytes)

	// If the emergency ballast isn't appropriately sized, try to resize it.
	// This is a no-op if the ballast is already sized or if there's not
	// enough available capacity to resize it. Capacity is called periodically
	// by the kvserver, and that drives the automatic resizing of the ballast.
	if !p.readOnly {
		resized, err := maybeEstablishBallast(p.unencryptedFS, p.ballastPath, p.ballastSize, du)
		if err != nil {
			return roachpb.StoreCapacity{}, errors.Wrap(err, "resizing ballast")
		}
		if resized {
			p.logger.Infof("resized ballast %s to size %s",
				p.ballastPath, humanizeutil.IBytes(p.ballastSize))
			du, err = p.unencryptedFS.GetDiskUsage(dir)
			if err != nil {
				return roachpb.StoreCapacity{}, err
			}
		}
	}

	// Pebble has detailed accounting of its own disk space usage, and it's
	// incrementally updated which helps avoid O(# files) work here.
	m := p.db.Metrics()
	totalUsedBytes := int64(m.DiskSpaceUsage())

	// We don't have incremental accounting of the disk space usage of files
	// in the auxiliary directory. Walk the auxiliary directory and all its
	// subdirectories, adding to the total used bytes.
	if errOuter := filepath.Walk(p.auxDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// This can happen if CockroachDB removes files out from under us -
			// just keep going to get the best estimate we can.
			if oserror.IsNotExist(err) {
				return nil
			}
			// Special-case: if the store-dir is configured using the root of some fs,
			// e.g. "/mnt/db", we might have special fs-created files like lost+found
			// that we can't read, so just ignore them rather than crashing.
			if oserror.IsPermission(err) && filepath.Base(path) == "lost+found" {
				return nil
			}
			return err
		}
		if path == p.ballastPath {
			// Skip the ballast. Counting it as used is likely to confuse
			// users, and it's more akin to space that is just unavailable
			// like disk space often restricted to a root user.
			return nil
		}
		if info.Mode().IsRegular() {
			totalUsedBytes += info.Size()
		}
		return nil
	}); errOuter != nil {
		return roachpb.StoreCapacity{}, errOuter
	}

	// If no size limitation have been placed on the store size or if the
	// limitation is greater than what's available, just return the actual
	// totals.
	if p.maxSize == 0 || p.maxSize >= fsuTotal || p.path == "" {
		return roachpb.StoreCapacity{
			Capacity:  fsuTotal,
			Available: fsuAvail,
			Used:      totalUsedBytes,
		}, nil
	}

	available := p.maxSize - totalUsedBytes
	if available > fsuAvail {
		available = fsuAvail
	}
	if available < 0 {
		available = 0
	}

	return roachpb.StoreCapacity{
		Capacity:  p.maxSize,
		Available: available,
		Used:      totalUsedBytes,
	}, nil
}

// Flush implements the Engine interface.
func (p *Pebble) Flush() error {
	return p.db.Flush()
}

// GetMetrics implements the Engine interface.
func (p *Pebble) GetMetrics() Metrics {
	m := p.db.Metrics()
	return Metrics{
		Metrics:         m,
		WriteStallCount: atomic.LoadInt64(&p.writeStallCount),
		DiskSlowCount:   atomic.LoadInt64(&p.diskSlowCount),
		DiskStallCount:  atomic.LoadInt64(&p.diskStallCount),
	}
}

// GetEncryptionRegistries implements the Engine interface.
func (p *Pebble) GetEncryptionRegistries() (*EncryptionRegistries, error) {
	rv := &EncryptionRegistries{}
	var err error
	if p.encryption != nil {
		rv.KeyRegistry, err = p.encryption.StatsHandler.GetDataKeysRegistry()
		if err != nil {
			return nil, err
		}
	}
	if p.fileRegistry != nil {
		rv.FileRegistry, err = protoutil.Marshal(p.fileRegistry.getRegistryCopy())
		if err != nil {
			return nil, err
		}
	}
	return rv, nil
}

// GetEnvStats implements the Engine interface.
func (p *Pebble) GetEnvStats() (*EnvStats, error) {
	// TODO(sumeer): make the stats complete. There are no bytes stats. The TotalFiles is missing
	// files that are not in the registry (from before encryption was enabled).
	stats := &EnvStats{}
	if p.encryption == nil {
		return stats, nil
	}
	stats.EncryptionType = p.encryption.StatsHandler.GetActiveStoreKeyType()
	var err error
	stats.EncryptionStatus, err = p.encryption.StatsHandler.GetEncryptionStatus()
	if err != nil {
		return nil, err
	}
	fr := p.fileRegistry.getRegistryCopy()
	activeKeyID, err := p.encryption.StatsHandler.GetActiveDataKeyID()
	if err != nil {
		return nil, err
	}

	m := p.db.Metrics()
	stats.TotalFiles = 3 /* CURRENT, MANIFEST, OPTIONS */
	stats.TotalFiles += uint64(m.WAL.Files + m.Table.ZombieCount + m.WAL.ObsoleteFiles + m.Table.ObsoleteCount)
	stats.TotalBytes = m.WAL.Size + m.Table.ZombieSize + m.Table.ObsoleteSize
	for _, l := range m.Levels {
		stats.TotalFiles += uint64(l.NumFiles)
		stats.TotalBytes += uint64(l.Size)
	}

	sstSizes := make(map[pebble.FileNum]uint64)
	sstInfos, err := p.db.SSTables()
	if err != nil {
		return nil, err
	}
	for _, ssts := range sstInfos {
		for _, sst := range ssts {
			sstSizes[sst.FileNum] = sst.Size
		}
	}

	for filePath, entry := range fr.Files {
		keyID, err := p.encryption.StatsHandler.GetKeyIDFromSettings(entry.EncryptionSettings)
		if err != nil {
			return nil, err
		}
		if len(keyID) == 0 {
			keyID = "plain"
		}
		if keyID != activeKeyID {
			continue
		}
		stats.ActiveKeyFiles++

		filename := p.fs.PathBase(filePath)
		numStr := strings.TrimSuffix(filename, ".sst")
		if len(numStr) == len(filename) {
			continue // not a sstable
		}
		u, err := strconv.ParseUint(numStr, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing filename %q", errors.Safe(filename))
		}
		stats.ActiveKeyBytes += sstSizes[pebble.FileNum(u)]
	}
	return stats, nil
}

// GetAuxiliaryDir implements the Engine interface.
func (p *Pebble) GetAuxiliaryDir() string {
	return p.auxDir
}

// NewBatch implements the Engine interface.
func (p *Pebble) NewBatch() Batch {
	return newPebbleBatch(p.db, p.db.NewIndexedBatch(), false /* writeOnly */, p.settings)
}

// NewReadOnly implements the Engine interface.
func (p *Pebble) NewReadOnly(durability DurabilityRequirement) ReadWriter {
	return newPebbleReadOnly(p, durability)
}

// NewUnindexedBatch implements the Engine interface.
func (p *Pebble) NewUnindexedBatch(writeOnly bool) Batch {
	return newPebbleBatch(p.db, p.db.NewBatch(), writeOnly, p.settings)
}

// NewSnapshot implements the Engine interface.
func (p *Pebble) NewSnapshot() Reader {
	return &pebbleSnapshot{
		snapshot: p.db.NewSnapshot(),
		parent:   p,
	}
}

// Type implements the Engine interface.
func (p *Pebble) Type() enginepb.EngineType {
	return enginepb.EngineTypePebble
}

// IngestExternalFiles implements the Engine interface.
func (p *Pebble) IngestExternalFiles(ctx context.Context, paths []string) error {
	return p.db.Ingest(paths)
}

// PreIngestDelay implements the Engine interface.
func (p *Pebble) PreIngestDelay(ctx context.Context) {
	preIngestDelay(ctx, p, p.settings)
}

// ApproximateDiskBytes implements the Engine interface.
func (p *Pebble) ApproximateDiskBytes(from, to roachpb.Key) (uint64, error) {
	count, err := p.db.EstimateDiskUsage(from, to)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Compact implements the Engine interface.
func (p *Pebble) Compact() error {
	return p.db.Compact(nil, EncodeMVCCKey(MVCCKeyMax), true /* parallel */)
}

// CompactRange implements the Engine interface.
func (p *Pebble) CompactRange(start, end roachpb.Key) error {
	bufStart := EncodeMVCCKey(MVCCKey{start, hlc.Timestamp{}})
	bufEnd := EncodeMVCCKey(MVCCKey{end, hlc.Timestamp{}})
	return p.db.Compact(bufStart, bufEnd, true /* parallel */)
}

// InMem returns true if the receiver is an in-memory engine and false
// otherwise.
func (p *Pebble) InMem() bool {
	return p.path == ""
}

// RegisterFlushCompletedCallback implements the Engine interface.
func (p *Pebble) RegisterFlushCompletedCallback(cb func()) {
	p.mu.Lock()
	p.mu.flushCompletedCallback = cb
	p.mu.Unlock()
}

// ReadFile implements the Engine interface.
func (p *Pebble) ReadFile(filename string) ([]byte, error) {
	file, err := p.fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

// WriteFile writes data to a file in this RocksDB's env.
func (p *Pebble) WriteFile(filename string, data []byte) error {
	file, err := p.fs.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(data))
	return err
}

// Remove implements the FS interface.
func (p *Pebble) Remove(filename string) error {
	return p.fs.Remove(filename)
}

// RemoveAll implements the Engine interface.
func (p *Pebble) RemoveAll(dir string) error {
	return p.fs.RemoveAll(dir)
}

// Link implements the FS interface.
func (p *Pebble) Link(oldname, newname string) error {
	return p.fs.Link(oldname, newname)
}

var _ fs.FS = &Pebble{}

// Create implements the FS interface.
func (p *Pebble) Create(name string) (fs.File, error) {
	return p.fs.Create(name)
}

// CreateWithSync implements the FS interface.
func (p *Pebble) CreateWithSync(name string, bytesPerSync int) (fs.File, error) {
	f, err := p.fs.Create(name)
	if err != nil {
		return nil, err
	}
	return vfs.NewSyncingFile(f, vfs.SyncingFileOptions{BytesPerSync: bytesPerSync}), nil
}

// Open implements the FS interface.
func (p *Pebble) Open(name string) (fs.File, error) {
	return p.fs.Open(name)
}

// OpenDir implements the FS interface.
func (p *Pebble) OpenDir(name string) (fs.File, error) {
	return p.fs.OpenDir(name)
}

// Rename implements the FS interface.
func (p *Pebble) Rename(oldname, newname string) error {
	return p.fs.Rename(oldname, newname)
}

// MkdirAll implements the FS interface.
func (p *Pebble) MkdirAll(name string) error {
	return p.fs.MkdirAll(name, 0755)
}

// List implements the FS interface.
func (p *Pebble) List(name string) ([]string, error) {
	dirents, err := p.fs.List(name)
	sort.Strings(dirents)
	return dirents, err
}

// Stat implements the FS interface.
func (p *Pebble) Stat(name string) (os.FileInfo, error) {
	return p.fs.Stat(name)
}

// CreateCheckpoint implements the Engine interface.
func (p *Pebble) CreateCheckpoint(dir string) error {
	return p.db.Checkpoint(dir)
}

// SetMinVersion implements the Engine interface.
func (p *Pebble) SetMinVersion(version roachpb.Version) error {
	// NB: SetMinVersion must be idempotent. It may called multiple
	// times with the same version.

	// Writing the min version file commits this storage engine to the
	// provided cluster version.
	if err := writeMinVersionFile(p.unencryptedFS, p.path, version); err != nil {
		return err
	}

	// Pebble has a concept of format major versions, similar to cluster
	// versions. Backwards incompatible changes to Pebble's on-disk
	// format are gated behind new format major versions. Bumping the
	// storage engine's format major version is tied to a CockroachDB
	// cluster version.
	//
	// Format major versions and cluster versions both only ratchet
	// upwards. Here we map the persisted cluster version to the
	// corresponding format major version, ratcheting Pebble's format
	// major version if necessary.
	//
	// Note that when introducing a new Pebble format version that relies on _all_
	// engines in a cluster being at the same, newer format major version, two
	// cluster versions should be used. The first is used to enable the feature in
	// Pebble, and should control the version ratchet below. The second is used as
	// a feature flag. The use of two cluster versions relies on a guarantee
	// provided by the migration framework (see pkg/migration) that if a node is
	// at a version X+1, it is guaranteed that all nodes have already ratcheted
	// their store version to the version X that enabled the feature at the Pebble
	// level.
	formatVers := pebble.FormatMostCompatible
	// Cases are ordered from newer to older versions.
	switch {
	case !version.Less(clusterversion.ByKey(clusterversion.EnsurePebbleFormatVersionRangeKeys)):
		if formatVers < pebble.FormatRangeKeys {
			formatVers = pebble.FormatRangeKeys
		}
	case !version.Less(clusterversion.ByKey(clusterversion.PebbleFormatSplitUserKeysMarked)):
		if formatVers < pebble.FormatSplitUserKeysMarked {
			formatVers = pebble.FormatSplitUserKeysMarked
		}
	case !version.Less(clusterversion.ByKey(clusterversion.PebbleFormatBlockPropertyCollector)):
		if formatVers < pebble.FormatBlockPropertyCollector {
			formatVers = pebble.FormatBlockPropertyCollector
		}
	case !version.Less(clusterversion.ByKey(clusterversion.TODOPreV21_2)):
		if formatVers < pebble.FormatSetWithDelete {
			formatVers = pebble.FormatSetWithDelete
		}
	}
	if p.db.FormatMajorVersion() < formatVers {
		if err := p.db.RatchetFormatMajorVersion(formatVers); err != nil {
			return errors.Wrap(err, "ratcheting format major version")
		}
		if formatVers >= pebble.FormatRangeKeys {
			atomic.StoreInt32(&p.supportsRangeKeys, 1)
		}
	}
	return nil
}

// MinVersionIsAtLeastTargetVersion implements the Engine interface.
func (p *Pebble) MinVersionIsAtLeastTargetVersion(target roachpb.Version) (bool, error) {
	return MinVersionIsAtLeastTargetVersion(p.unencryptedFS, p.path, target)
}

type pebbleReadOnly struct {
	parent *Pebble
	// The iterator reuse optimization in pebbleReadOnly is for servicing a
	// BatchRequest, such that the iterators get reused across different
	// requests in the batch.
	// Reuse iterators for {normal,prefix} x {MVCCKey,EngineKey} iteration. We
	// need separate iterators for EngineKey and MVCCKey iteration since
	// iterators that make separated locks/intents look as interleaved need to
	// use both simultaneously.
	// When the first iterator is initialized, or when
	// PinEngineStateForIterators is called (whichever happens first), the
	// underlying *pebble.Iterator is stashed in iter, so that subsequent
	// iterator initialization can use Iterator.Clone to use the same underlying
	// engine state. This relies on the fact that all pebbleIterators created
	// here are marked as reusable, which causes pebbleIterator.Close to not
	// close iter. iter will be closed when pebbleReadOnly.Close is called.
	prefixIter       pebbleIterator
	normalIter       pebbleIterator
	prefixEngineIter pebbleIterator
	normalEngineIter pebbleIterator

	iter       cloneableIter
	iterUnused bool
	durability DurabilityRequirement
	closed     bool
}

var _ ReadWriter = &pebbleReadOnly{}

var pebbleReadOnlyPool = sync.Pool{
	New: func() interface{} {
		return &pebbleReadOnly{
			// Defensively set reusable=true. One has to be careful about this since
			// an accidental false value would cause these iterators, that are value
			// members of pebbleReadOnly, to be put in the pebbleIterPool.
			prefixIter:       pebbleIterator{reusable: true},
			normalIter:       pebbleIterator{reusable: true},
			prefixEngineIter: pebbleIterator{reusable: true},
			normalEngineIter: pebbleIterator{reusable: true},
		}
	},
}

// Instantiates a new pebbleReadOnly.
func newPebbleReadOnly(parent *Pebble, durability DurabilityRequirement) *pebbleReadOnly {
	p := pebbleReadOnlyPool.Get().(*pebbleReadOnly)
	// When p is a reused pebbleReadOnly from the pool, the iter fields preserve
	// the original reusable=true that was set above in pebbleReadOnlyPool.New(),
	// and some buffers that are safe to reuse. Everything else has been reset by
	// pebbleIterator.destroy().
	*p = pebbleReadOnly{
		parent:           parent,
		prefixIter:       p.prefixIter,
		normalIter:       p.normalIter,
		prefixEngineIter: p.prefixEngineIter,
		normalEngineIter: p.normalEngineIter,
		durability:       durability,
	}
	return p
}

func (p *pebbleReadOnly) Close() {
	if p.closed {
		panic("closing an already-closed pebbleReadOnly")
	}
	p.closed = true
	if p.iterUnused {
		err := p.iter.Close()
		if err != nil {
			panic(err)
		}
	}

	// Setting iter to nil is sufficient since it will be closed by one of the
	// subsequent destroy calls.
	p.iter = nil
	p.prefixIter.destroy()
	p.normalIter.destroy()
	p.prefixEngineIter.destroy()
	p.normalEngineIter.destroy()
	p.durability = StandardDurability

	pebbleReadOnlyPool.Put(p)
}

func (p *pebbleReadOnly) Closed() bool {
	return p.closed
}

// ExportMVCCToSst is part of the engine.Reader interface.
func (p *pebbleReadOnly) ExportMVCCToSst(
	ctx context.Context, exportOptions ExportOptions, dest io.Writer,
) (roachpb.BulkOpSummary, roachpb.Key, hlc.Timestamp, error) {
	r := wrapReader(p)
	// Doing defer r.Free() does not inline.
	summary, k, err := pebbleExportToSst(ctx, p.parent.settings, r, exportOptions, dest)
	r.Free()
	return summary, k.Key, k.Timestamp, err
}

func (p *pebbleReadOnly) MVCCGet(key MVCCKey) ([]byte, error) {
	return mvccGetHelper(key, p)
}

func (p *pebbleReadOnly) rawMVCCGet(key []byte) ([]byte, error) {
	if p.closed {
		panic("using a closed pebbleReadOnly")
	}
	// Cannot delegate to p.parent.rawMVCCGet since we need to use p.durability.
	onlyReadGuaranteedDurable := false
	if p.durability == GuaranteedDurability {
		onlyReadGuaranteedDurable = true
	}
	options := pebble.IterOptions{
		LowerBound:                key,
		UpperBound:                encoding.BytesNext(key),
		OnlyReadGuaranteedDurable: onlyReadGuaranteedDurable,
	}
	iter := p.parent.db.NewIter(&options)
	defer func() {
		// Already handled error.
		_ = iter.Close()
	}()
	valid := iter.SeekGE(key)
	if !valid {
		return nil, iter.Error()
	}
	return iter.Value(), nil
}

func (p *pebbleReadOnly) MVCCGetProto(
	key MVCCKey, msg protoutil.Message,
) (ok bool, keyBytes, valBytes int64, err error) {
	return pebbleGetProto(p, key, msg)
}

func (p *pebbleReadOnly) MVCCIterate(
	start, end roachpb.Key, iterKind MVCCIterKind, f func(MVCCKeyValue) error,
) error {
	if p.closed {
		panic("using a closed pebbleReadOnly")
	}
	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		err := iterateOnReader(r, start, end, iterKind, f)
		r.Free()
		return err
	}
	return iterateOnReader(p, start, end, iterKind, f)
}

// NewMVCCIterator implements the Engine interface.
func (p *pebbleReadOnly) NewMVCCIterator(iterKind MVCCIterKind, opts IterOptions) MVCCIterator {
	if p.closed {
		panic("using a closed pebbleReadOnly")
	}

	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		iter := r.NewMVCCIterator(iterKind, opts)
		r.Free()
		if util.RaceEnabled {
			iter = wrapInUnsafeIter(iter)
		}
		return iter
	}

	iter := &p.normalIter
	if opts.Prefix {
		iter = &p.prefixIter
	}
	if iter.inuse {
		return newPebbleIterator(p.parent.db, p.iter, opts, p.durability, p.SupportsRangeKeys())
	}

	if iter.iter != nil {
		iter.setOptions(opts, p.durability)
	} else {
		iter.init(p.parent.db, p.iter, p.iterUnused, opts, p.durability, p.SupportsRangeKeys())
		if p.iter == nil {
			// For future cloning.
			p.iter = iter.iter
		}
		p.iterUnused = false
		iter.reusable = true
	}

	iter.inuse = true
	var rv MVCCIterator = iter
	if util.RaceEnabled {
		rv = wrapInUnsafeIter(rv)
	}
	return rv
}

// NewEngineIterator implements the Engine interface.
func (p *pebbleReadOnly) NewEngineIterator(opts IterOptions) EngineIterator {
	if p.closed {
		panic("using a closed pebbleReadOnly")
	}
	if opts.KeyTypes != IterKeyTypePointsOnly {
		panic("EngineIterator does not support range keys")
	}

	iter := &p.normalEngineIter
	if opts.Prefix {
		iter = &p.prefixEngineIter
	}
	if iter.inuse {
		return newPebbleIterator(p.parent.db, p.iter, opts, p.durability, p.SupportsRangeKeys())
	}

	if iter.iter != nil {
		iter.setOptions(opts, p.durability)
	} else {
		iter.init(p.parent.db, p.iter, p.iterUnused, opts, p.durability, p.SupportsRangeKeys())
		if p.iter == nil {
			// For future cloning.
			p.iter = iter.iter
		}
		p.iterUnused = false
		iter.reusable = true
	}

	iter.inuse = true
	return iter
}

// ConsistentIterators implements the Engine interface.
func (p *pebbleReadOnly) ConsistentIterators() bool {
	return true
}

// SupportsRangeKeys implements the Engine interface.
func (p *pebbleReadOnly) SupportsRangeKeys() bool {
	return p.parent.SupportsRangeKeys()
}

// PinEngineStateForIterators implements the Engine interface.
func (p *pebbleReadOnly) PinEngineStateForIterators() error {
	if p.iter == nil {
		o := (*pebble.IterOptions)(nil)
		if p.durability == GuaranteedDurability {
			o = &pebble.IterOptions{OnlyReadGuaranteedDurable: true}
		}
		p.iter = p.parent.db.NewIter(o)
		// Since the iterator is being created just to pin the state of the engine
		// for future iterators, we'll avoid cloning it the next time we want an
		// iterator and instead just re-use what we created here.
		p.iterUnused = true
	}
	return nil
}

// Writer methods are not implemented for pebbleReadOnly. Ideally, the code
// could be refactored so that a Reader could be supplied to evaluateBatch

// Writer is the write interface to an engine's data.
func (p *pebbleReadOnly) ApplyBatchRepr(repr []byte, sync bool) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearMVCC(key MVCCKey) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearUnversioned(key roachpb.Key) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearIntent(
	key roachpb.Key, txnDidNotUpdateMeta bool, txnUUID uuid.UUID,
) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearEngineKey(key EngineKey) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) SingleClearEngineKey(key EngineKey) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearRawRange(start, end roachpb.Key) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearMVCCRange(start, end roachpb.Key) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearMVCCVersions(start, end MVCCKey) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ClearMVCCIteratorRange(start, end roachpb.Key) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ExperimentalPutMVCCRangeKey(MVCCRangeKey, MVCCValue) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ExperimentalClearMVCCRangeKey(MVCCRangeKey) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) ExperimentalClearAllMVCCRangeKeys(roachpb.Key, roachpb.Key) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) Merge(key MVCCKey, value []byte) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) PutMVCC(key MVCCKey, value MVCCValue) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) PutRawMVCC(key MVCCKey, value []byte) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) PutUnversioned(key roachpb.Key, value []byte) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) PutIntent(
	ctx context.Context, key roachpb.Key, value []byte, txnUUID uuid.UUID,
) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) PutEngineKey(key EngineKey, value []byte) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) LogData(data []byte) error {
	panic("not implemented")
}

func (p *pebbleReadOnly) LogLogicalOp(op MVCCLogicalOpType, details MVCCLogicalOpDetails) {
	panic("not implemented")
}

func (p *pebbleReadOnly) ShouldWriteLocalTimestamps(ctx context.Context) bool {
	panic("not implemented")
}

// pebbleSnapshot represents a snapshot created using Pebble.NewSnapshot().
type pebbleSnapshot struct {
	snapshot *pebble.Snapshot
	parent   *Pebble
	closed   bool
}

var _ Reader = &pebbleSnapshot{}

// Close implements the Reader interface.
func (p *pebbleSnapshot) Close() {
	_ = p.snapshot.Close()
	p.closed = true
}

// Closed implements the Reader interface.
func (p *pebbleSnapshot) Closed() bool {
	return p.closed
}

// ExportMVCCToSst is part of the engine.Reader interface.
func (p *pebbleSnapshot) ExportMVCCToSst(
	ctx context.Context, exportOptions ExportOptions, dest io.Writer,
) (roachpb.BulkOpSummary, roachpb.Key, hlc.Timestamp, error) {
	r := wrapReader(p)
	// Doing defer r.Free() does not inline.
	summary, k, err := pebbleExportToSst(ctx, p.parent.settings, r, exportOptions, dest)
	r.Free()
	return summary, k.Key, k.Timestamp, err
}

// Get implements the Reader interface.
func (p *pebbleSnapshot) MVCCGet(key MVCCKey) ([]byte, error) {
	if len(key.Key) == 0 {
		return nil, emptyKeyError()
	}
	r := wrapReader(p)
	// Doing defer r.Free() does not inline.
	v, err := r.MVCCGet(key)
	r.Free()
	return v, err
}

func (p *pebbleSnapshot) rawMVCCGet(key []byte) ([]byte, error) {
	ret, closer, err := p.snapshot.Get(key)
	if closer != nil {
		retCopy := make([]byte, len(ret))
		copy(retCopy, ret)
		ret = retCopy
		closer.Close()
	}
	if errors.Is(err, pebble.ErrNotFound) || len(ret) == 0 {
		return nil, nil
	}
	return ret, err
}

// MVCCGetProto implements the Reader interface.
func (p *pebbleSnapshot) MVCCGetProto(
	key MVCCKey, msg protoutil.Message,
) (ok bool, keyBytes, valBytes int64, err error) {
	return pebbleGetProto(p, key, msg)
}

// MVCCIterate implements the Reader interface.
func (p *pebbleSnapshot) MVCCIterate(
	start, end roachpb.Key, iterKind MVCCIterKind, f func(MVCCKeyValue) error,
) error {
	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		err := iterateOnReader(r, start, end, iterKind, f)
		r.Free()
		return err
	}
	return iterateOnReader(p, start, end, iterKind, f)
}

// NewMVCCIterator implements the Reader interface.
func (p *pebbleSnapshot) NewMVCCIterator(iterKind MVCCIterKind, opts IterOptions) MVCCIterator {
	if iterKind == MVCCKeyAndIntentsIterKind {
		r := wrapReader(p)
		// Doing defer r.Free() does not inline.
		iter := r.NewMVCCIterator(iterKind, opts)
		r.Free()
		if util.RaceEnabled {
			iter = wrapInUnsafeIter(iter)
		}
		return iter
	}

	iter := MVCCIterator(newPebbleIterator(
		p.snapshot, nil, opts, StandardDurability, p.SupportsRangeKeys()))
	if util.RaceEnabled {
		iter = wrapInUnsafeIter(iter)
	}
	return iter
}

// NewEngineIterator implements the Reader interface.
func (p pebbleSnapshot) NewEngineIterator(opts IterOptions) EngineIterator {
	if opts.KeyTypes != IterKeyTypePointsOnly {
		panic("EngineIterator does not support range keys")
	}
	return newPebbleIterator(p.snapshot, nil, opts, StandardDurability, p.SupportsRangeKeys())
}

// ConsistentIterators implements the Reader interface.
func (p pebbleSnapshot) ConsistentIterators() bool {
	return true
}

// SupportsRangeKeys implements the Reader interface.
func (p *pebbleSnapshot) SupportsRangeKeys() bool {
	return p.parent.SupportsRangeKeys()
}

// PinEngineStateForIterators implements the Reader interface.
func (p *pebbleSnapshot) PinEngineStateForIterators() error {
	// Snapshot already pins state, so nothing to do.
	return nil
}

// pebbleGetProto uses Reader.MVCCGet, so it not as efficient as a function
// that can unmarshal without copying bytes. But we don't care about
// efficiency, since this is used to implement Reader.MVCCGetProto, which is
// deprecated and only used in tests.
func pebbleGetProto(
	reader Reader, key MVCCKey, msg protoutil.Message,
) (ok bool, keyBytes, valBytes int64, err error) {
	val, err := reader.MVCCGet(key)
	if err != nil || val == nil {
		return false, 0, 0, err
	}
	keyBytes = int64(key.Len())
	valBytes = int64(len(val))
	if msg != nil {
		err = protoutil.Unmarshal(val, msg)
	}
	return true, keyBytes, valBytes, err
}

// ExceedMaxSizeError is the error returned when an export request
// fails due the export size exceeding the budget. This can be caused
// by large KVs that have many revisions.
type ExceedMaxSizeError struct {
	reached int64
	maxSize uint64
}

var _ error = &ExceedMaxSizeError{}

func (e *ExceedMaxSizeError) Error() string {
	return fmt.Sprintf("export size (%d bytes) exceeds max size (%d bytes)", e.reached, e.maxSize)
}

func pebbleExportToSst(
	ctx context.Context, cs *cluster.Settings, reader Reader, options ExportOptions, dest io.Writer,
) (roachpb.BulkOpSummary, MVCCKey, error) {
	var span *tracing.Span
	ctx, span = tracing.ChildSpan(ctx, "pebbleExportToSst")
	defer span.Finish()
	sstWriter := MakeBackupSSTWriter(ctx, cs, dest)
	defer sstWriter.Close()

	var rows RowCounter
	iter := NewMVCCIncrementalIterator(
		reader,
		MVCCIncrementalIterOptions{
			EndKey:                              options.EndKey,
			EnableTimeBoundIteratorOptimization: options.UseTBI,
			StartTime:                           options.StartTS,
			EndTime:                             options.EndTS,
			IntentPolicy:                        MVCCIncrementalIterIntentPolicyAggregate,
		})
	defer iter.Close()
	var curKey roachpb.Key // only used if exportAllRevisions
	var resumeKey roachpb.Key
	var resumeTS hlc.Timestamp
	paginated := options.TargetSize > 0
	trackKeyBoundary := paginated || options.ResourceLimiter != nil
	firstIteration := true
	for iter.SeekGE(options.StartKey); ; {
		ok, err := iter.Valid()
		if err != nil {
			return roachpb.BulkOpSummary{}, MVCCKey{}, err
		}
		if !ok {
			break
		}
		unsafeKey := iter.UnsafeKey()
		if unsafeKey.Key.Compare(options.EndKey) >= 0 {
			break
		}

		if iter.NumCollectedIntents() > 0 {
			break
		}

		isNewKey := !options.ExportAllRevisions || !unsafeKey.Key.Equal(curKey)
		if trackKeyBoundary && options.ExportAllRevisions && isNewKey {
			curKey = append(curKey[:0], unsafeKey.Key...)
		}

		if options.ResourceLimiter != nil {
			// Don't check resources on first iteration to ensure we can make some progress regardless
			// of starvation. Otherwise operations could spin indefinitely.
			if firstIteration {
				firstIteration = false
			} else {
				// In happy day case we want to only stop at key boundaries as it allows callers to use
				// produced sst's directly. But if we can't find key boundary within reasonable number of
				// iterations we would split mid key.
				// To achieve that we use soft and hard thresholds in limiter. Once soft limit is reached
				// we would start searching for key boundary and return as soon as it is reached. If we
				// can't find it before hard limit is reached and caller requested mid key stop we would
				// immediately return.
				limit := options.ResourceLimiter.IsExhausted()
				// We can stop at key once any threshold is reached or force stop at hard limit if midkey
				// split is allowed.
				if limit >= ResourceLimitReachedSoft && isNewKey || limit == ResourceLimitReachedHard && options.StopMidKey {
					// Reached iteration limit, stop with resume span
					resumeKey = append(make(roachpb.Key, 0, len(unsafeKey.Key)), unsafeKey.Key...)
					if !isNewKey {
						resumeTS = unsafeKey.Timestamp
					}
					break
				}
			}
		}

		unsafeValue := iter.UnsafeValue()
		skip := false
		if unsafeKey.IsValue() {
			mvccValue, ok, err := tryDecodeSimpleMVCCValue(unsafeValue)
			if !ok && err == nil {
				mvccValue, err = decodeExtendedMVCCValue(unsafeValue)
			}
			if err != nil {
				return roachpb.BulkOpSummary{}, MVCCKey{}, errors.Wrapf(err, "decoding mvcc value %s", unsafeKey)
			}

			// Export only the inner roachpb.Value, not the MVCCValue header.
			unsafeValue = mvccValue.Value.RawBytes

			// Skip tombstone records when start time is zero (non-incremental)
			// and we are not exporting all versions.
			skipTombstones := !options.ExportAllRevisions && options.StartTS.IsEmpty()
			skip = skipTombstones && mvccValue.IsTombstone()
		}

		if !skip {
			if err := rows.Count(unsafeKey.Key); err != nil {
				return roachpb.BulkOpSummary{}, MVCCKey{}, errors.Wrapf(err, "decoding %s", unsafeKey)
			}
			curSize := rows.BulkOpSummary.DataSize
			reachedTargetSize := curSize > 0 && uint64(curSize) >= options.TargetSize
			newSize := curSize + int64(len(unsafeKey.Key)+len(unsafeValue))
			reachedMaxSize := options.MaxSize > 0 && newSize > int64(options.MaxSize)
			// When paginating we stop writing in two cases:
			// - target size is reached and we wrote all versions of a key
			// - maximum size reached and we are allowed to stop mid key
			if paginated && (isNewKey && reachedTargetSize || options.StopMidKey && reachedMaxSize) {
				// Allocate the right size for resumeKey rather than using curKey.
				resumeKey = append(make(roachpb.Key, 0, len(unsafeKey.Key)), unsafeKey.Key...)
				if options.StopMidKey && !isNewKey {
					resumeTS = unsafeKey.Timestamp
				}
				break
			}
			if reachedMaxSize {
				return roachpb.BulkOpSummary{}, MVCCKey{}, &ExceedMaxSizeError{reached: newSize, maxSize: options.MaxSize}
			}
			if unsafeKey.Timestamp.IsEmpty() {
				// This should never be an intent since the incremental iterator returns
				// an error when encountering intents.
				if err := sstWriter.PutUnversioned(unsafeKey.Key, unsafeValue); err != nil {
					return roachpb.BulkOpSummary{}, MVCCKey{}, errors.Wrapf(err, "adding key %s", unsafeKey)
				}
			} else {
				if err := sstWriter.PutRawMVCC(unsafeKey, unsafeValue); err != nil {
					return roachpb.BulkOpSummary{}, MVCCKey{}, errors.Wrapf(err, "adding key %s", unsafeKey)
				}
			}
			rows.BulkOpSummary.DataSize = newSize
		}

		if options.ExportAllRevisions {
			iter.Next()
		} else {
			iter.NextKey()
		}
	}

	// First check if we encountered an intent while iterating the data.
	// If we do it means this export can't complete and is aborted. We need to loop over remaining data
	// to collect all matching intents before returning them in an error to the caller.
	if iter.NumCollectedIntents() > 0 {
		for uint64(iter.NumCollectedIntents()) < options.MaxIntents {
			iter.NextKey()
			// If we encounter other errors during intent collection, we return our original write intent failure.
			// We would find this new error again upon retry.
			ok, _ := iter.Valid()
			if !ok {
				break
			}
		}
		err := iter.TryGetIntentError()
		return roachpb.BulkOpSummary{}, MVCCKey{}, err
	}

	if rows.BulkOpSummary.DataSize == 0 {
		// If no records were added to the sstable, skip completing it and return a
		// nil slice – the export code will discard it anyway (based on 0 DataSize).
		return roachpb.BulkOpSummary{}, MVCCKey{}, nil
	}

	if err := sstWriter.Finish(); err != nil {
		return roachpb.BulkOpSummary{}, MVCCKey{}, err
	}

	return rows.BulkOpSummary, MVCCKey{Key: resumeKey, Timestamp: resumeTS}, nil
}

// pebbleFindRangeKeySpan returns the minimum span within the given bounds that
// covers all contained range keys. If there are no range keys within the
// bounds, this returns nil keys.
func pebbleFindRangeKeySpan(r pebble.Reader, lower, upper []byte) ([]byte, []byte, error) {
	iter := r.NewIter(&pebble.IterOptions{
		KeyTypes:   pebble.IterKeyTypeRangesOnly,
		LowerBound: lower,
		UpperBound: upper,
	})
	defer func() {
		// We handle errors during iteration.
		_ = iter.Close()
	}()

	// Look for a range key. If none are found, return nil bounds.
	if !iter.SeekGE(lower) {
		return nil, nil, iter.Error()
	}
	rangeStart, _ := iter.RangeBounds()
	start := append([]byte{}, rangeStart...)

	// Find the end of the span.
	if !iter.SeekLT(upper) {
		if err := iter.Error(); err != nil {
			return nil, nil, err
		}
		return nil, nil, errors.AssertionFailedf("unexpected missing range key in %s-%s", lower, upper)
	}
	_, rangeEnd := iter.RangeBounds()
	end := append([]byte{}, rangeEnd...)

	if bytes.Compare(start, end) >= 0 {
		return nil, nil, errors.AssertionFailedf("range key end %s at or before start %s", end, start)
	}
	return start, end, nil
}
