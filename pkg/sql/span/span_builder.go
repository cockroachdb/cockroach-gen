// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package span

import (
	"sort"
	"sync"

	"github.com/cockroachdb/cockroach/pkg/keys"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/inverted"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/constraint"
	"github.com/cockroachdb/cockroach/pkg/sql/rowenc"
	"github.com/cockroachdb/cockroach/pkg/sql/rowenc/keyside"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/encoding"
	"github.com/cockroachdb/errors"
)

// Builder is a single struct for generating key spans from Constraints, Datums,
// encDatums, and InvertedSpans.
type Builder struct {
	evalCtx       *tree.EvalContext
	codec         keys.SQLCodec
	table         catalog.TableDescriptor
	index         catalog.Index
	indexColTypes []*types.T
	indexColDirs  []descpb.IndexDescriptor_Direction

	// KeyPrefix is the prefix of keys generated by the builder.
	KeyPrefix []byte
	alloc     tree.DatumAlloc
}

var builderPool = sync.Pool{
	New: func() interface{} { return &Builder{} },
}

// MakeBuilder creates a Builder for a table and index. The returned object must
// be Release()d when no longer needed.
func MakeBuilder(
	evalCtx *tree.EvalContext,
	codec keys.SQLCodec,
	table catalog.TableDescriptor,
	index catalog.Index,
) *Builder {
	s := builderPool.Get().(*Builder)
	*s = Builder{
		evalCtx:       evalCtx,
		codec:         codec,
		table:         table,
		index:         index,
		indexColTypes: s.indexColTypes,
		KeyPrefix:     rowenc.MakeIndexKeyPrefix(codec, table.GetID(), index.GetID()),
	}

	s.indexColDirs = s.table.IndexFullColumnDirections(index)
	columns := s.table.IndexFullColumns(index)
	if cap(s.indexColTypes) < len(columns) {
		s.indexColTypes = make([]*types.T, len(columns))
	} else {
		s.indexColTypes = s.indexColTypes[:len(columns)]
	}
	for i, col := range columns {
		// TODO (rohany): do I need to look at table columns with mutations here as well?
		if col != nil && col.Public() {
			s.indexColTypes[i] = col.GetType()
		}
	}

	return s
}

// SpanFromEncDatums encodes a span with prefixLen constraint columns from the
// index prefixed with the index key prefix that includes the table and index
// ID. SpanFromEncDatums assumes that the EncDatums in values are in the order
// of the index columns. It also returns whether or not the input values contain
// a null value or not, which can be used as input for
// CanSplitSpanIntoFamilySpans.
func (s *Builder) SpanFromEncDatums(
	values rowenc.EncDatumRow, prefixLen int,
) (_ roachpb.Span, containsNull bool, _ error) {
	return rowenc.MakeSpanFromEncDatums(
		values[:prefixLen], s.indexColTypes[:prefixLen], s.indexColDirs[:prefixLen], s.index, &s.alloc, s.KeyPrefix)
}

// SpanFromEncDatumsWithRange encodes a range span. The inequality is assumed to
// be the end of the span and the start/end keys are generated by putting them
// in the values row at the prefixLen - 1 position. Only one of start or end
// need be non-nil, omitted one causing an open ended range span to be
// generated. Since the exec code knows nothing about index column sorting
// direction we assume ascending if they are descending we deal with that here.
func (s *Builder) SpanFromEncDatumsWithRange(
	values rowenc.EncDatumRow,
	prefixLen int,
	startDatum tree.Datum,
	startInclusive bool,
	endDatum tree.Datum,
	endInclusive bool,
) (_ roachpb.Span, containsNull bool, err error) {

	if s.indexColDirs[prefixLen-1] == descpb.IndexDescriptor_DESC {
		startDatum, endDatum = endDatum, startDatum
		startInclusive, endInclusive = endInclusive, startInclusive
	}

	makeKeyFromRow := func(r rowenc.EncDatumRow, l int) (k roachpb.Key, cn bool, e error) {
		k, _, cn, e = rowenc.MakeKeyFromEncDatums(
			r[:l], s.indexColTypes[:l], s.indexColDirs[:l], s.index, &s.alloc, s.KeyPrefix,
		)
		return
	}

	var startKey, endKey roachpb.Key
	var startContainsNull, endContainsNull bool
	if startDatum != nil {
		values[prefixLen-1] = rowenc.EncDatum{Datum: startDatum}
		startKey, startContainsNull, err = makeKeyFromRow(values, prefixLen)
		if !startInclusive {
			startKey = startKey.Next()
		}
	} else {
		startKey, startContainsNull, err = makeKeyFromRow(values, prefixLen-1)
		// If we have a ascending index make sure not to include NULLs.
		if s.indexColDirs[prefixLen-1] == descpb.IndexDescriptor_ASC {
			startKey = encoding.EncodeNullAscending(startKey)
		}
		startKey = startKey.Next()
	}

	if err != nil {
		return roachpb.Span{}, false, err
	}

	if endDatum != nil {
		values[prefixLen-1] = rowenc.EncDatum{Datum: endDatum}
		endKey, endContainsNull, err = makeKeyFromRow(values, prefixLen)
		if endInclusive {
			endKey = endKey.PrefixEnd()
		}
	} else {
		endKey, endContainsNull, err = makeKeyFromRow(values, prefixLen-1)
		// If we have a descending index make sure not to include NULLs.
		if s.indexColDirs[prefixLen-1] == descpb.IndexDescriptor_DESC {
			endKey = encoding.EncodeNullDescending(endKey)
		} else {
			endKey = endKey.PrefixEnd()
		}
	}

	if err != nil {
		return roachpb.Span{}, false, err
	}

	return roachpb.Span{Key: startKey, EndKey: endKey}, startContainsNull || endContainsNull, nil
}

// SpanFromDatumRow generates an index span with prefixLen constraint columns from the index.
// SpanFromDatumRow assumes that values is a valid table row for the Builder's table.
// It also returns whether or not the input values contain a null value or not, which can be
// used as input for CanSplitSpanIntoFamilySpans.
func (s *Builder) SpanFromDatumRow(
	values tree.Datums, prefixLen int, colMap catalog.TableColMap,
) (_ roachpb.Span, containsNull bool, _ error) {
	return rowenc.EncodePartialIndexSpan(s.index, prefixLen, colMap, values, s.KeyPrefix)
}

// SpanToPointSpan converts a span into a span that represents a point lookup on a
// specific family. It is up to the caller to ensure that this is a safe operation,
// by calling CanSplitSpanIntoFamilySpans before using it.
func (s *Builder) SpanToPointSpan(span roachpb.Span, family descpb.FamilyID) roachpb.Span {
	key := keys.MakeFamilyKey(span.Key, uint32(family))
	return roachpb.Span{Key: key, EndKey: roachpb.Key(key).PrefixEnd()}
}

// Functions for optimizer related span generation are below.

// SpansFromConstraint generates spans from an optimizer constraint.
// A span.Splitter can be used to generate more specific family spans.
//
// TODO (rohany): In future work, there should be a single API to generate spans
//  from constraints, datums and encdatums.
func (s *Builder) SpansFromConstraint(
	c *constraint.Constraint, splitter Splitter,
) (roachpb.Spans, error) {
	var spans roachpb.Spans
	var err error
	if c == nil || c.IsUnconstrained() {
		// Encode a full span.
		spans, err = s.appendSpansFromConstraintSpan(spans, &constraint.UnconstrainedSpan, splitter)
		if err != nil {
			return nil, err
		}
		return spans, nil
	}

	spans = make(roachpb.Spans, 0, c.Spans.Count())
	for i := 0; i < c.Spans.Count(); i++ {
		spans, err = s.appendSpansFromConstraintSpan(spans, c.Spans.Get(i), splitter)
		if err != nil {
			return nil, err
		}
	}
	return spans, nil
}

// UnconstrainedSpans returns the full span corresponding to the Builder's
// table and index.
func (s *Builder) UnconstrainedSpans() (roachpb.Spans, error) {
	return s.SpansFromConstraint(nil, NoopSplitter())
}

// appendSpansFromConstraintSpan converts a constraint.Span to one or more
// roachpb.Spans and appends them to the provided spans. It appends multiple
// spans in the case that multiple, non-adjacent column families should be
// scanned.
func (s *Builder) appendSpansFromConstraintSpan(
	appendTo roachpb.Spans, cs *constraint.Span, splitter Splitter,
) (roachpb.Spans, error) {
	var span roachpb.Span
	var err error
	var containsNull bool
	// Encode each logical part of the start key.
	span.Key, containsNull, err = s.encodeConstraintKey(cs.StartKey())
	if err != nil {
		return nil, err
	}
	if cs.StartBoundary() == constraint.IncludeBoundary {
		if cs.StartKey().IsEmpty() {
			span.Key = append(span.Key, s.KeyPrefix...)
		}
	} else {
		// We need to exclude the value this logical part refers to.
		span.Key = span.Key.PrefixEnd()
	}
	// Encode each logical part of the end key.
	span.EndKey, _, err = s.encodeConstraintKey(cs.EndKey())
	if err != nil {
		return nil, err
	}
	if cs.EndKey().IsEmpty() {
		span.EndKey = append(span.EndKey, s.KeyPrefix...)
	}

	// Optimization: for single row lookups on a table with one or more column
	// families, only scan the relevant column families, and use GetRequests
	// instead of ScanRequests when doing the column family fetches.
	if splitter.CanSplitSpanIntoFamilySpans(cs.StartKey().Length(), containsNull) && span.Key.Equal(span.EndKey) {
		return rowenc.SplitRowKeyIntoFamilySpans(appendTo, span.Key, splitter.neededFamilies), nil
	}

	// We need to advance the end key if it is inclusive.
	if cs.EndBoundary() == constraint.IncludeBoundary {
		span.EndKey = span.EndKey.PrefixEnd()
	}

	return append(appendTo, span), nil
}

// encodeConstraintKey encodes each logical part of a constraint.Key into a
// roachpb.Key.
func (s *Builder) encodeConstraintKey(
	ck constraint.Key,
) (key roachpb.Key, containsNull bool, err error) {
	if ck.IsEmpty() {
		return key, containsNull, nil
	}
	key = append(key, s.KeyPrefix...)
	for i := 0; i < ck.Length(); i++ {
		val := ck.Value(i)
		if val == tree.DNull {
			containsNull = true
		}

		// For extra columns (like implicit columns), the direction
		// is ascending.
		dir := encoding.Ascending
		if i < s.index.NumKeyColumns() {
			dir, err = s.index.GetKeyColumnDirection(i).ToEncodingDirection()
			if err != nil {
				return nil, false, err
			}
		}

		key, err = keyside.Encode(key, val, dir)
		if err != nil {
			return nil, false, err
		}
	}
	return key, containsNull, nil
}

// InvertedSpans represent inverted index spans that can be encoded into
// key spans.
type InvertedSpans interface {
	// Len returns the number of spans represented.
	Len() int

	// Start returns the start bytes of the ith span.
	Start(i int) []byte

	// End returns the end bytes of the ith span.
	End(i int) []byte
}

var _ InvertedSpans = inverted.Spans{}
var _ InvertedSpans = inverted.SpanExpressionProtoSpans{}

// SpansFromInvertedSpans constructs spans to scan an inverted index.
//
// If the index is a single-column inverted index, c should be nil.
//
// If the index is a multi-column inverted index, c should constrain the
// non-inverted prefix columns of the index. Each span in c must have a single
// key. The resulting roachpb.Spans are created by performing a cross product of
// keys in c and the invertedSpan keys.
//
// scratch can be an optional roachpb.Spans slice that will be reused to
// populate the result.
func (s *Builder) SpansFromInvertedSpans(
	invertedSpans InvertedSpans, c *constraint.Constraint, scratch roachpb.Spans,
) (roachpb.Spans, error) {
	if invertedSpans == nil {
		return nil, errors.AssertionFailedf("invertedSpans cannot be nil")
	}

	var scratchRows []rowenc.EncDatumRow
	if c != nil {
		// For each span in c, create a scratchRow that starts with the span's
		// keys. The last slot in each scratchRow is reserved for encoding the
		// inverted span key.
		scratchRows = make([]rowenc.EncDatumRow, c.Spans.Count())
		for i, n := 0, c.Spans.Count(); i < n; i++ {
			span := c.Spans.Get(i)

			// The spans must have the same start and end key.
			if !span.HasSingleKey(s.evalCtx) {
				return nil, errors.AssertionFailedf("constraint span %s does not have a single key", span)
			}

			keyLength := span.StartKey().Length()
			scratchRows[i] = make(rowenc.EncDatumRow, keyLength+1)
			for j := 0; j < keyLength; j++ {
				val := span.StartKey().Value(j)
				scratchRows[i][j] = rowenc.DatumToEncDatum(val.ResolvedType(), val)
			}
		}
	} else {
		// If c is nil, then the spans must constrain a single-column inverted
		// index. In this case, only 1 scratchRow of length 1 is needed to
		// encode the inverted spans.
		scratchRows = make([]rowenc.EncDatumRow, 1)
		scratchRows[0] = make(rowenc.EncDatumRow, 1)
	}

	scratch = scratch[:0]
	for i := range scratchRows {
		for j, n := 0, invertedSpans.Len(); j < n; j++ {
			var indexSpan roachpb.Span
			var err error
			if indexSpan.Key, err = s.generateInvertedSpanKey(invertedSpans.Start(j), scratchRows[i]); err != nil {
				return nil, err
			}
			if indexSpan.EndKey, err = s.generateInvertedSpanKey(invertedSpans.End(j), scratchRows[i]); err != nil {
				return nil, err
			}
			scratch = append(scratch, indexSpan)
		}
	}
	sort.Sort(scratch)
	return scratch, nil
}

// generateInvertedSpanKey returns a key that encodes enc and scratchRow. The
// last slot in scratchRow is overwritten in order to encode enc. If the length
// of scratchRow is greater than one, the EncDatums that precede the last slot
// are encoded as prefix keys of enc.
func (s *Builder) generateInvertedSpanKey(
	enc []byte, scratchRow rowenc.EncDatumRow,
) (roachpb.Key, error) {
	keyLen := len(scratchRow) - 1
	scratchRow = scratchRow[:keyLen]
	if len(enc) > 0 {
		// Pretend that the encoded inverted val is an EncDatum. This isn't always
		// true, since JSON inverted columns use a custom encoding. But since we
		// are providing an already encoded Datum, the following will eventually
		// fall through to EncDatum.Encode() which will reuse the encoded bytes.
		encDatum := rowenc.EncDatumFromEncoded(descpb.DatumEncoding_ASCENDING_KEY, enc)
		scratchRow = append(scratchRow, encDatum)
		keyLen++
	}
	// Else, this is the case of scanning all the inverted keys under the
	// prefix of scratchRow (including the case where there is no prefix when
	// the inverted column is the first column). Note, the inverted span in
	// that case will be [nil, RKeyMax), and the caller calls this method with
	// both nil and RKeyMax. The first call will fall through here, and
	// generate a span, of which we will only use Span.Key. Span.EndKey is
	// generated by the caller in the second call, with RKeyMax.

	span, _, err := s.SpanFromEncDatums(scratchRow, keyLen)
	return span.Key, err
}

// Release implements the execinfra.Releasable interface.
func (s *Builder) Release() {
	*s = Builder{
		// Note that the types are small objects, so we don't bother deeply
		// resetting the indexColTypes slice.
		indexColTypes: s.indexColTypes[:0],
	}
	builderPool.Put(s)
}
