// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/telemetry.proto

package eventpb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SampledQuery is the SQL query event logged to the telemetry channel. It
// contains common SQL event/execution details.
type SampledQuery struct {
	CommonEventDetails    `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonSQLEventDetails `protobuf:"bytes,2,opt,name=sql,proto3,embedded=sql" json:""`
	CommonSQLExecDetails  `protobuf:"bytes,3,opt,name=exec,proto3,embedded=exec" json:""`
	// skipped_queries indicate how many SQL statements were not
	// considered for sampling prior to this one. If the field is
	// omitted, or its value is zero, this indicates that no statement
	// was omitted since the last event.
	SkippedQueries uint64 `protobuf:"varint,4,opt,name=skipped_queries,json=skippedQueries,proto3" json:",omitempty"`
	// Cost of the query as estimated by the optimizer.
	CostEstimate float64 `protobuf:"fixed64,5,opt,name=cost_estimate,json=costEstimate,proto3" json:",omitempty"`
	// The distribution of the DistSQL query plan (local, full, or partial).
	Distribution string `protobuf:"bytes,6,opt,name=distribution,proto3" json:",omitempty" redact:"nonsensitive"`
	// SessionID is the ID of the session that initiated the query.
	SessionID string `protobuf:"bytes,8,opt,name=session_id,json=sessionId,proto3" json:",omitempty" redact:"nonsensitive"`
	// Name of the database that initiated the query.
	Database string `protobuf:"bytes,9,opt,name=database,proto3" json:",omitempty" redact:"nonsensitive"`
	// Statement ID of the query.
	StatementID string `protobuf:"bytes,10,opt,name=statement_id,json=statementId,proto3" json:",omitempty" redact:"nonsensitive"`
	// Transaction ID of the query.
	TransactionID string `protobuf:"bytes,11,opt,name=transaction_id,json=transactionId,proto3" json:",omitempty" redact:"nonsensitive"`
	// Database ID of the query.
	DatabaseID uint32 `protobuf:"varint,12,opt,name=database_id,json=databaseId,proto3" json:",omitempty"`
	// Statement fingerprint ID of the query.
	StatementFingerprintID uint64 `protobuf:"varint,13,opt,name=statement_fingerprint_id,json=statementFingerprintId,proto3" json:",omitempty"`
}

func (m *SampledQuery) Reset()         { *m = SampledQuery{} }
func (m *SampledQuery) String() string { return proto.CompactTextString(m) }
func (*SampledQuery) ProtoMessage()    {}
func (*SampledQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d317b4ad74be4f7, []int{0}
}
func (m *SampledQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SampledQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SampledQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampledQuery.Merge(m, src)
}
func (m *SampledQuery) XXX_Size() int {
	return m.Size()
}
func (m *SampledQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SampledQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SampledQuery proto.InternalMessageInfo

// CapturedIndexUsageStats
type CapturedIndexUsageStats struct {
	CommonEventDetails `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	// TotalReadCount is the number of times this index has been read from.
	TotalReadCount uint64 `protobuf:"varint,2,opt,name=total_read_count,json=totalReadCount,proto3" json:"total_read_count,omitempty"`
	// LastRead is the timestamp that this index was last being read from.
	LastRead string `protobuf:"bytes,3,opt,name=last_read,json=lastRead,proto3" json:",omitempty" redact:"nonsensitive"`
	// TableID is the ID of the table this index is created on. This is same as
	// descpb.TableID and is unique within the cluster.
	TableID uint32 `protobuf:"varint,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	// IndexID is the ID of the index within the scope of the given table.
	IndexID      uint32 `protobuf:"varint,5,opt,name=index_id,json=indexId,proto3" json:"index_id,omitempty"`
	DatabaseName string `protobuf:"bytes,6,opt,name=database_name,json=databaseName,proto3" json:",omitempty" redact:"nonsensitive"`
	TableName    string `protobuf:"bytes,7,opt,name=table_name,json=tableName,proto3" json:",omitempty" redact:"nonsensitive"`
	IndexName    string `protobuf:"bytes,8,opt,name=index_name,json=indexName,proto3" json:",omitempty" redact:"nonsensitive"`
	IndexType    string `protobuf:"bytes,9,opt,name=index_type,json=indexType,proto3" json:",omitempty" redact:"nonsensitive"`
	IsUnique     bool   `protobuf:"varint,10,opt,name=is_unique,json=isUnique,proto3" json:",omitempty"`
	IsInverted   bool   `protobuf:"varint,11,opt,name=is_inverted,json=isInverted,proto3" json:",omitempty"`
}

func (m *CapturedIndexUsageStats) Reset()         { *m = CapturedIndexUsageStats{} }
func (m *CapturedIndexUsageStats) String() string { return proto.CompactTextString(m) }
func (*CapturedIndexUsageStats) ProtoMessage()    {}
func (*CapturedIndexUsageStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d317b4ad74be4f7, []int{1}
}
func (m *CapturedIndexUsageStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CapturedIndexUsageStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CapturedIndexUsageStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapturedIndexUsageStats.Merge(m, src)
}
func (m *CapturedIndexUsageStats) XXX_Size() int {
	return m.Size()
}
func (m *CapturedIndexUsageStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CapturedIndexUsageStats.DiscardUnknown(m)
}

var xxx_messageInfo_CapturedIndexUsageStats proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SampledQuery)(nil), "cockroach.util.log.eventpb.SampledQuery")
	proto.RegisterType((*CapturedIndexUsageStats)(nil), "cockroach.util.log.eventpb.CapturedIndexUsageStats")
}

func init() { proto.RegisterFile("util/log/eventpb/telemetry.proto", fileDescriptor_3d317b4ad74be4f7) }

var fileDescriptor_3d317b4ad74be4f7 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcd, 0x6e, 0xfb, 0x44,
	0x10, 0xc0, 0x63, 0x9a, 0x36, 0xc9, 0xe6, 0x03, 0xb4, 0x82, 0x62, 0x55, 0x60, 0x47, 0x39, 0xd0,
	0x20, 0x50, 0x02, 0x14, 0x01, 0x42, 0xe2, 0x92, 0xa4, 0x95, 0x2c, 0x2a, 0xa4, 0x38, 0xe9, 0x01,
	0x2e, 0xd6, 0xc6, 0x3b, 0x84, 0x55, 0x6d, 0xaf, 0xe3, 0x5d, 0x57, 0xcd, 0x5b, 0xf0, 0x2e, 0x3c,
	0x00, 0xd7, 0x1e, 0x7b, 0xec, 0xc9, 0x02, 0xf7, 0xd6, 0x23, 0x4f, 0x80, 0xbc, 0x76, 0x92, 0x96,
	0x14, 0x48, 0x0e, 0xff, 0x53, 0xe2, 0x99, 0xdf, 0xfc, 0x34, 0xfb, 0x31, 0x36, 0x6a, 0xc7, 0x92,
	0x79, 0x7d, 0x8f, 0xcf, 0xfb, 0x70, 0x03, 0x81, 0x0c, 0x67, 0x7d, 0x09, 0x1e, 0xf8, 0x20, 0xa3,
	0x65, 0x2f, 0x8c, 0xb8, 0xe4, 0xf8, 0xc4, 0xe5, 0xee, 0x75, 0xc4, 0x89, 0xfb, 0x4b, 0x2f, 0x63,
	0x7b, 0x1e, 0x9f, 0xf7, 0x0a, 0xf6, 0xe4, 0xdd, 0x39, 0x9f, 0x73, 0x85, 0xf5, 0xb3, 0x7f, 0x79,
	0xc5, 0xc9, 0x87, 0x5b, 0x4e, 0xf5, 0x2b, 0x8a, 0xf4, 0xe9, 0x56, 0x5a, 0x2c, 0x3c, 0x87, 0xc4,
	0x94, 0x49, 0xe7, 0x39, 0xd8, 0xf9, 0xbd, 0x82, 0x1a, 0x13, 0xe2, 0x87, 0x1e, 0xd0, 0x71, 0x0c,
	0xd1, 0x12, 0x4f, 0xd1, 0x91, 0xcb, 0x7d, 0x9f, 0x07, 0xba, 0xd6, 0xd6, 0xba, 0xf5, 0x2f, 0x7a,
	0xbd, 0x7f, 0xef, 0xad, 0x37, 0x54, 0xe4, 0x79, 0xf6, 0x34, 0x02, 0x49, 0x98, 0x27, 0x06, 0x8d,
	0xbb, 0xc4, 0x2c, 0xdd, 0x27, 0xa6, 0xf6, 0x94, 0x98, 0x25, 0xbb, 0x70, 0xe1, 0x31, 0x3a, 0x10,
	0x0b, 0x4f, 0x7f, 0x4b, 0x29, 0x3f, 0xff, 0x7f, 0xe5, 0x64, 0x7c, 0xf9, 0x1f, 0xd6, 0xcc, 0x85,
	0x6d, 0x54, 0x86, 0x5b, 0x70, 0xf5, 0x03, 0xe5, 0xfc, 0x6c, 0x37, 0xe7, 0x2d, 0xb8, 0xaf, 0x2b,
	0x95, 0x0b, 0x7f, 0x8d, 0xde, 0x16, 0xd7, 0x2c, 0x0c, 0x81, 0x3a, 0x8b, 0x18, 0x22, 0x06, 0x42,
	0x2f, 0xb7, 0xb5, 0x6e, 0x79, 0xd0, 0x7a, 0x4a, 0x4c, 0xf4, 0x29, 0xf7, 0x99, 0x04, 0x3f, 0x94,
	0x4b, 0xbb, 0x55, 0x60, 0xe3, 0x9c, 0xc2, 0x67, 0xa8, 0xe9, 0x72, 0x21, 0x1d, 0x10, 0x92, 0xf9,
	0x44, 0x82, 0x7e, 0xd8, 0xd6, 0xba, 0xda, 0x56, 0x59, 0x23, 0x83, 0xce, 0x0b, 0x06, 0x7f, 0x8f,
	0x1a, 0x94, 0x09, 0x19, 0xb1, 0x59, 0x2c, 0x19, 0x0f, 0xf4, 0xa3, 0xb6, 0xd6, 0xad, 0x0d, 0x4e,
	0x5f, 0xd6, 0xfc, 0x95, 0x98, 0xef, 0x45, 0x40, 0x89, 0x2b, 0xbf, 0xed, 0x04, 0x3c, 0x10, 0x10,
	0x08, 0x26, 0xd9, 0x0d, 0x74, 0xec, 0x17, 0xc5, 0x78, 0x82, 0x90, 0x00, 0x21, 0x18, 0x0f, 0x1c,
	0x46, 0xf5, 0xaa, 0x52, 0x7d, 0x99, 0x26, 0x66, 0x6d, 0x92, 0x47, 0xad, 0xd1, 0xae, 0xde, 0x5a,
	0xe1, 0xb1, 0x28, 0x1e, 0xa2, 0x2a, 0x25, 0x92, 0xcc, 0x88, 0x00, 0xbd, 0xb6, 0x5f, 0x77, 0xeb,
	0x42, 0xfc, 0x23, 0x6a, 0x08, 0x49, 0x24, 0xf8, 0x10, 0xc8, 0xac, 0x37, 0xa4, 0x44, 0x5f, 0xa5,
	0x89, 0x59, 0x9f, 0xac, 0xe2, 0xbb, 0x77, 0x57, 0x5f, 0xbb, 0x2c, 0x8a, 0x1d, 0xd4, 0x92, 0x11,
	0x09, 0x04, 0x71, 0x65, 0xb1, 0xf0, 0xba, 0x92, 0x7f, 0x93, 0x26, 0x66, 0x73, 0xba, 0xc9, 0xec,
	0xae, 0x6f, 0x3e, 0xf3, 0x59, 0x14, 0x7f, 0x87, 0xea, 0xab, 0x75, 0x64, 0xf6, 0x46, 0x5b, 0xeb,
	0x36, 0x07, 0x1f, 0xa4, 0x89, 0x89, 0x46, 0x45, 0xf8, 0x9f, 0x6a, 0x1b, 0xad, 0x0a, 0x2c, 0x8a,
	0x67, 0x48, 0xdf, 0x2c, 0xfd, 0x67, 0x16, 0xcc, 0x21, 0x0a, 0x23, 0x96, 0x6f, 0x43, 0x53, 0x5d,
	0xac, 0x6e, 0x9a, 0x98, 0xc7, 0xeb, 0x6d, 0xb8, 0xd8, 0x20, 0x5b, 0xde, 0x63, 0xf1, 0x1a, 0x45,
	0x3b, 0xbf, 0x1d, 0xa2, 0xf7, 0x87, 0x24, 0x94, 0x71, 0x04, 0xd4, 0x0a, 0x28, 0xdc, 0x5e, 0x09,
	0x32, 0x87, 0x4c, 0x29, 0xde, 0xd0, 0x30, 0x77, 0xd1, 0x3b, 0x92, 0x4b, 0xe2, 0x39, 0x11, 0x10,
	0xea, 0xb8, 0x3c, 0x0e, 0xa4, 0x9a, 0xec, 0xb2, 0xdd, 0x52, 0x71, 0x1b, 0x08, 0x1d, 0x66, 0x51,
	0x3c, 0x42, 0x35, 0x8f, 0x08, 0xa9, 0x40, 0x35, 0xa8, 0xfb, 0x5c, 0xa0, 0xac, 0x32, 0x53, 0xe1,
	0x8f, 0x50, 0x55, 0x92, 0x99, 0xa7, 0x4e, 0xa0, 0xac, 0x4e, 0xa0, 0x9e, 0x26, 0x66, 0x65, 0x9a,
	0xc5, 0xac, 0x91, 0x5d, 0x51, 0x49, 0x4b, 0x71, 0x2c, 0xdb, 0x80, 0x8c, 0x3b, 0xdc, 0x70, 0x6a,
	0x53, 0x32, 0x4e, 0x25, 0x2d, 0x8a, 0x2f, 0x51, 0x73, 0x7d, 0xa8, 0x01, 0xf1, 0x61, 0xff, 0xc1,
	0x2b, 0xaa, 0x7f, 0x20, 0x3e, 0xe0, 0x0b, 0x84, 0xf2, 0xee, 0x94, 0xaa, 0xb2, 0x9f, 0xaa, 0xa6,
	0x4a, 0x57, 0x9e, 0xbc, 0x7b, 0xe5, 0xa9, 0xee, 0xe9, 0x51, 0xa5, 0x2f, 0x3d, 0x72, 0x19, 0xee,
	0x3d, 0xb5, 0xb9, 0x67, 0xba, 0x0c, 0x01, 0x7f, 0x82, 0x6a, 0x4c, 0x38, 0x71, 0xc0, 0x16, 0x31,
	0xa8, 0x99, 0xad, 0x6e, 0xbd, 0xce, 0xaa, 0x4c, 0x5c, 0xa9, 0x3c, 0xee, 0xa3, 0x3a, 0x13, 0x0e,
	0x0b, 0x6e, 0x20, 0x92, 0x90, 0x4f, 0xe1, 0x36, 0x8e, 0x98, 0xb0, 0x0a, 0x62, 0xf0, 0xf1, 0xdd,
	0x9f, 0x46, 0xe9, 0x2e, 0x35, 0xb4, 0xfb, 0xd4, 0xd0, 0x1e, 0x52, 0x43, 0xfb, 0x23, 0x35, 0xb4,
	0x5f, 0x1f, 0x8d, 0xd2, 0xfd, 0xa3, 0x51, 0x7a, 0x78, 0x34, 0x4a, 0x3f, 0x55, 0x8a, 0x7b, 0x39,
	0x3b, 0x52, 0x5f, 0xaa, 0xb3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x49, 0x95, 0x80, 0x31, 0x47,
	0x07, 0x00, 0x00,
}

func (m *SampledQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SampledQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SampledQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StatementFingerprintID != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.StatementFingerprintID))
		i--
		dAtA[i] = 0x68
	}
	if m.DatabaseID != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.DatabaseID))
		i--
		dAtA[i] = 0x60
	}
	if len(m.TransactionID) > 0 {
		i -= len(m.TransactionID)
		copy(dAtA[i:], m.TransactionID)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.TransactionID)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.StatementID) > 0 {
		i -= len(m.StatementID)
		copy(dAtA[i:], m.StatementID)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.StatementID)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Database) > 0 {
		i -= len(m.Database)
		copy(dAtA[i:], m.Database)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.Database)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SessionID) > 0 {
		i -= len(m.SessionID)
		copy(dAtA[i:], m.SessionID)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.SessionID)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Distribution) > 0 {
		i -= len(m.Distribution)
		copy(dAtA[i:], m.Distribution)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.Distribution)))
		i--
		dAtA[i] = 0x32
	}
	if m.CostEstimate != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CostEstimate))))
		i--
		dAtA[i] = 0x29
	}
	if m.SkippedQueries != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.SkippedQueries))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.CommonSQLExecDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTelemetry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.CommonSQLEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTelemetry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTelemetry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CapturedIndexUsageStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CapturedIndexUsageStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CapturedIndexUsageStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsInverted {
		i--
		if m.IsInverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.IsUnique {
		i--
		if m.IsUnique {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if len(m.IndexType) > 0 {
		i -= len(m.IndexType)
		copy(dAtA[i:], m.IndexType)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.IndexType)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IndexName) > 0 {
		i -= len(m.IndexName)
		copy(dAtA[i:], m.IndexName)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.IndexName)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TableName) > 0 {
		i -= len(m.TableName)
		copy(dAtA[i:], m.TableName)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.TableName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.DatabaseName) > 0 {
		i -= len(m.DatabaseName)
		copy(dAtA[i:], m.DatabaseName)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.DatabaseName)))
		i--
		dAtA[i] = 0x32
	}
	if m.IndexID != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.IndexID))
		i--
		dAtA[i] = 0x28
	}
	if m.TableID != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.TableID))
		i--
		dAtA[i] = 0x20
	}
	if len(m.LastRead) > 0 {
		i -= len(m.LastRead)
		copy(dAtA[i:], m.LastRead)
		i = encodeVarintTelemetry(dAtA, i, uint64(len(m.LastRead)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TotalReadCount != 0 {
		i = encodeVarintTelemetry(dAtA, i, uint64(m.TotalReadCount))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.CommonEventDetails.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTelemetry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTelemetry(dAtA []byte, offset int, v uint64) int {
	offset -= sovTelemetry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SampledQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovTelemetry(uint64(l))
	l = m.CommonSQLEventDetails.Size()
	n += 1 + l + sovTelemetry(uint64(l))
	l = m.CommonSQLExecDetails.Size()
	n += 1 + l + sovTelemetry(uint64(l))
	if m.SkippedQueries != 0 {
		n += 1 + sovTelemetry(uint64(m.SkippedQueries))
	}
	if m.CostEstimate != 0 {
		n += 9
	}
	l = len(m.Distribution)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.Database)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.StatementID)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.TransactionID)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	if m.DatabaseID != 0 {
		n += 1 + sovTelemetry(uint64(m.DatabaseID))
	}
	if m.StatementFingerprintID != 0 {
		n += 1 + sovTelemetry(uint64(m.StatementFingerprintID))
	}
	return n
}

func (m *CapturedIndexUsageStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovTelemetry(uint64(l))
	if m.TotalReadCount != 0 {
		n += 1 + sovTelemetry(uint64(m.TotalReadCount))
	}
	l = len(m.LastRead)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	if m.TableID != 0 {
		n += 1 + sovTelemetry(uint64(m.TableID))
	}
	if m.IndexID != 0 {
		n += 1 + sovTelemetry(uint64(m.IndexID))
	}
	l = len(m.DatabaseName)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.TableName)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.IndexName)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	l = len(m.IndexType)
	if l > 0 {
		n += 1 + l + sovTelemetry(uint64(l))
	}
	if m.IsUnique {
		n += 2
	}
	if m.IsInverted {
		n += 2
	}
	return n
}

func sovTelemetry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTelemetry(x uint64) (n int) {
	return sovTelemetry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SampledQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelemetry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SampledQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SampledQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLExecDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLExecDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedQueries", wireType)
			}
			m.SkippedQueries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedQueries |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CostEstimate", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CostEstimate = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distribution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distribution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatementID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatementID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatabaseID", wireType)
			}
			m.DatabaseID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DatabaseID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatementFingerprintID", wireType)
			}
			m.StatementFingerprintID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatementFingerprintID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTelemetry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTelemetry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CapturedIndexUsageStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTelemetry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CapturedIndexUsageStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CapturedIndexUsageStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReadCount", wireType)
			}
			m.TotalReadCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalReadCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRead", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastRead = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableID", wireType)
			}
			m.TableID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexID", wireType)
			}
			m.IndexID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatabaseName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatabaseName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TableName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndexName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTelemetry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTelemetry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndexType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUnique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsUnique = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsInverted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTelemetry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTelemetry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTelemetry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTelemetry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTelemetry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTelemetry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTelemetry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTelemetry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTelemetry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTelemetry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTelemetry = fmt.Errorf("proto: unexpected end of group")
)
