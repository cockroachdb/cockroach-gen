// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/index_usage_stats.proto

package roachpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// IndexUsageStatistics represents the index usage statistics per index.
// N.B. When fields are added to this struct, make sure to update
// (*IndexUsageStatistics).Add in roachpb/index_usage_stats.go.
type IndexUsageStatistics struct {
	// TotalReadCount is the number of times this index has been read from.
	TotalReadCount uint64 `protobuf:"varint,1,opt,name=total_read_count,json=totalReadCount" json:"total_read_count"`
	// LastRead is the timestamp that this index was last being read from.
	LastRead time.Time `protobuf:"bytes,2,opt,name=last_read,json=lastRead,stdtime" json:"last_read"`
	// TotalRowsRead is the number rows that has read from this index.
	// TODO(azhng): Currently this field is unused.
	TotalRowsRead uint64 `protobuf:"varint,3,opt,name=total_rows_read,json=totalRowsRead" json:"total_rows_read"`
	// TotalWriteCount is the number of times this index has been written to.
	// TODO(azhng): Currently this field is unused.
	TotalWriteCount uint64 `protobuf:"varint,4,opt,name=total_write_count,json=totalWriteCount" json:"total_write_count"`
	// LastWrite is the timestamp that this index was last being written to.
	// TODO(azhng): Currently this field is unused.
	LastWrite time.Time `protobuf:"bytes,5,opt,name=last_write,json=lastWrite,stdtime" json:"last_write"`
	// TotalRowsWritten is the number rows that have been written to this index.
	// TODO(azhng): Currently this field is unused.
	TotalRowsWritten uint64 `protobuf:"varint,6,opt,name=total_rows_written,json=totalRowsWritten" json:"total_rows_written"`
}

func (m *IndexUsageStatistics) Reset()         { *m = IndexUsageStatistics{} }
func (m *IndexUsageStatistics) String() string { return proto.CompactTextString(m) }
func (*IndexUsageStatistics) ProtoMessage()    {}
func (*IndexUsageStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523d71560a74d6c, []int{0}
}
func (m *IndexUsageStatistics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexUsageStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IndexUsageStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexUsageStatistics.Merge(m, src)
}
func (m *IndexUsageStatistics) XXX_Size() int {
	return m.Size()
}
func (m *IndexUsageStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexUsageStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_IndexUsageStatistics proto.InternalMessageInfo

// IndexUsageKey uniquely identifies an index. It's a tuple of TableID and a
// IndexID.
type IndexUsageKey struct {
	// TableID is the ID of the table this index is created on. This is same as
	// descpb.TableID and is unique within the cluster.
	TableID TableID `protobuf:"varint,1,opt,name=table_id,json=tableId,casttype=TableID" json:"table_id"`
	// IndexID is the ID of the index within the scope of the given table.
	IndexID IndexID `protobuf:"varint,2,opt,name=index_id,json=indexId,casttype=IndexID" json:"index_id"`
}

func (m *IndexUsageKey) Reset()         { *m = IndexUsageKey{} }
func (m *IndexUsageKey) String() string { return proto.CompactTextString(m) }
func (*IndexUsageKey) ProtoMessage()    {}
func (*IndexUsageKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523d71560a74d6c, []int{1}
}
func (m *IndexUsageKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexUsageKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IndexUsageKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexUsageKey.Merge(m, src)
}
func (m *IndexUsageKey) XXX_Size() int {
	return m.Size()
}
func (m *IndexUsageKey) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexUsageKey.DiscardUnknown(m)
}

var xxx_messageInfo_IndexUsageKey proto.InternalMessageInfo

// CollectedIndexUsageStatistics wraps collected index key and its usage
// statistics.
type CollectedIndexUsageStatistics struct {
	Key   IndexUsageKey        `protobuf:"bytes,1,opt,name=key" json:"key"`
	Stats IndexUsageStatistics `protobuf:"bytes,2,opt,name=stats" json:"stats"`
}

func (m *CollectedIndexUsageStatistics) Reset()         { *m = CollectedIndexUsageStatistics{} }
func (m *CollectedIndexUsageStatistics) String() string { return proto.CompactTextString(m) }
func (*CollectedIndexUsageStatistics) ProtoMessage()    {}
func (*CollectedIndexUsageStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7523d71560a74d6c, []int{2}
}
func (m *CollectedIndexUsageStatistics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectedIndexUsageStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CollectedIndexUsageStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectedIndexUsageStatistics.Merge(m, src)
}
func (m *CollectedIndexUsageStatistics) XXX_Size() int {
	return m.Size()
}
func (m *CollectedIndexUsageStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectedIndexUsageStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_CollectedIndexUsageStatistics proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IndexUsageStatistics)(nil), "cockroach.sql.IndexUsageStatistics")
	proto.RegisterType((*IndexUsageKey)(nil), "cockroach.sql.IndexUsageKey")
	proto.RegisterType((*CollectedIndexUsageStatistics)(nil), "cockroach.sql.CollectedIndexUsageStatistics")
}

func init() { proto.RegisterFile("roachpb/index_usage_stats.proto", fileDescriptor_7523d71560a74d6c) }

var fileDescriptor_7523d71560a74d6c = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xf7, 0xb5, 0x29, 0x09, 0x57, 0x05, 0xca, 0xa9, 0x43, 0x14, 0xc1, 0xb9, 0x0a, 0x4b, 0x91,
	0xd0, 0x19, 0x45, 0x30, 0x23, 0x1c, 0x16, 0x8b, 0xcd, 0x14, 0x55, 0x62, 0xb1, 0x2e, 0xf6, 0x61,
	0xac, 0xba, 0xb9, 0xe0, 0x7b, 0x51, 0xe8, 0xc0, 0x77, 0xe8, 0xc4, 0x67, 0xca, 0xd8, 0xb1, 0x53,
	0x01, 0xe7, 0x5b, 0x30, 0xa1, 0x7b, 0xe7, 0x12, 0x47, 0x82, 0xa1, 0xdb, 0xb3, 0xdf, 0xef, 0xdf,
	0xbd, 0x7b, 0x47, 0xfd, 0x4a, 0xcb, 0xf4, 0xf3, 0x7c, 0x1a, 0x14, 0xb3, 0x4c, 0x7d, 0x4d, 0x16,
	0x46, 0xe6, 0x2a, 0x31, 0x20, 0xc1, 0x88, 0x79, 0xa5, 0x41, 0xb3, 0x7e, 0xaa, 0xd3, 0x33, 0x04,
	0x09, 0xf3, 0xa5, 0x1c, 0x1e, 0xe6, 0x3a, 0xd7, 0xd8, 0x09, 0x6c, 0xe5, 0x40, 0x43, 0x3f, 0xd7,
	0x3a, 0x2f, 0x55, 0x80, 0x5f, 0xd3, 0xc5, 0xa7, 0x00, 0x8a, 0x73, 0x65, 0x40, 0x9e, 0xcf, 0x1d,
	0x60, 0x54, 0xef, 0xd0, 0xc3, 0xc8, 0x3a, 0x7c, 0xb0, 0x06, 0xef, 0x41, 0x42, 0x61, 0xa0, 0x48,
	0x0d, 0x13, 0xf4, 0x00, 0x34, 0xc8, 0x32, 0xa9, 0x94, 0xcc, 0x92, 0x54, 0x2f, 0x66, 0x30, 0x20,
	0x47, 0xe4, 0xb8, 0x13, 0x76, 0x56, 0x37, 0xbe, 0x17, 0x3f, 0xc0, 0x6e, 0xac, 0x64, 0x36, 0xb1,
	0x3d, 0xf6, 0x86, 0xde, 0x2f, 0xa5, 0x01, 0x84, 0x0f, 0x76, 0x8e, 0xc8, 0xf1, 0xfe, 0x78, 0x28,
	0x9c, 0xbb, 0xb8, 0x75, 0x17, 0x27, 0xb7, 0xee, 0x61, 0xcf, 0x8a, 0x5c, 0xfe, 0xf0, 0x49, 0xdc,
	0xb3, 0x34, 0xab, 0xc3, 0x9e, 0xd3, 0x87, 0x8d, 0xa5, 0x5e, 0x1a, 0x27, 0xb4, 0xdb, 0x72, 0xec,
	0x3b, 0x47, 0xbd, 0x34, 0x88, 0x7e, 0x41, 0x1f, 0x39, 0xf4, 0xb2, 0x2a, 0x40, 0x35, 0x09, 0x3b,
	0x2d, 0xbc, 0x13, 0x3b, 0xb5, 0x5d, 0x17, 0x71, 0x42, 0x29, 0x46, 0x44, 0xc2, 0x60, 0xef, 0x0e,
	0x19, 0xf1, 0x68, 0xa8, 0xc4, 0xc6, 0x94, 0xb5, 0x42, 0x5a, 0x29, 0x50, 0xb3, 0xc1, 0xbd, 0x96,
	0xef, 0xc1, 0xdf, 0x9c, 0xa7, 0xae, 0x3b, 0xfa, 0x46, 0xfb, 0x9b, 0x19, 0xbf, 0x53, 0x17, 0xec,
	0x15, 0xed, 0x81, 0x9c, 0x96, 0x2a, 0x29, 0x32, 0x1c, 0x6a, 0x3f, 0x1c, 0x5a, 0x6a, 0x7d, 0xe3,
	0x77, 0x4f, 0xec, 0xff, 0xe8, 0xed, 0xef, 0x4d, 0x19, 0x77, 0x11, 0x1b, 0x65, 0x96, 0xe6, 0xb6,
	0xa1, 0x70, 0x23, 0x6e, 0xd1, 0x50, 0xdf, 0xd1, 0x9a, 0x32, 0xee, 0x22, 0x36, 0xca, 0x46, 0xdf,
	0x09, 0x7d, 0x32, 0xd1, 0x65, 0xa9, 0x52, 0x50, 0xd9, 0x3f, 0x2f, 0xfb, 0x25, 0xdd, 0x3d, 0x53,
	0x17, 0x18, 0x65, 0x7f, 0xfc, 0x58, 0x6c, 0x6d, 0x96, 0xd8, 0x8a, 0xde, 0x9c, 0xd1, 0xc2, 0xd9,
	0x6b, 0xba, 0x87, 0x0b, 0xd9, 0x5c, 0xf7, 0xd3, 0xff, 0xf2, 0x36, 0x4e, 0x0d, 0xdd, 0xf1, 0xc2,
	0x67, 0xab, 0x5f, 0xdc, 0x5b, 0xd5, 0x9c, 0x5c, 0xd5, 0x9c, 0x5c, 0xd7, 0x9c, 0xfc, 0xac, 0x39,
	0xb9, 0x5c, 0x73, 0xef, 0x6a, 0xcd, 0xbd, 0xeb, 0x35, 0xf7, 0x3e, 0x76, 0x9b, 0x57, 0xf0, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x14, 0x7a, 0xa2, 0x07, 0x0f, 0x03, 0x00, 0x00,
}

func (m *IndexUsageStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexUsageStatistics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexUsageStatistics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.TotalRowsWritten))
	i--
	dAtA[i] = 0x30
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastWrite, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastWrite):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.TotalWriteCount))
	i--
	dAtA[i] = 0x20
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.TotalRowsRead))
	i--
	dAtA[i] = 0x18
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastRead, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastRead):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.TotalReadCount))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *IndexUsageKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexUsageKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexUsageKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.IndexID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintIndexUsageStats(dAtA, i, uint64(m.TableID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *CollectedIndexUsageStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectedIndexUsageStatistics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectedIndexUsageStatistics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIndexUsageStats(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIndexUsageStats(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIndexUsageStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexUsageStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexUsageStatistics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovIndexUsageStats(uint64(m.TotalReadCount))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastRead)
	n += 1 + l + sovIndexUsageStats(uint64(l))
	n += 1 + sovIndexUsageStats(uint64(m.TotalRowsRead))
	n += 1 + sovIndexUsageStats(uint64(m.TotalWriteCount))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastWrite)
	n += 1 + l + sovIndexUsageStats(uint64(l))
	n += 1 + sovIndexUsageStats(uint64(m.TotalRowsWritten))
	return n
}

func (m *IndexUsageKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovIndexUsageStats(uint64(m.TableID))
	n += 1 + sovIndexUsageStats(uint64(m.IndexID))
	return n
}

func (m *CollectedIndexUsageStatistics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Key.Size()
	n += 1 + l + sovIndexUsageStats(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovIndexUsageStats(uint64(l))
	return n
}

func sovIndexUsageStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexUsageStats(x uint64) (n int) {
	return sovIndexUsageStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexUsageStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexUsageStats
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
			return fmt.Errorf("proto: IndexUsageStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexUsageStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReadCount", wireType)
			}
			m.TotalReadCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRead", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
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
				return ErrInvalidLengthIndexUsageStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexUsageStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastRead, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRowsRead", wireType)
			}
			m.TotalRowsRead = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalRowsRead |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWriteCount", wireType)
			}
			m.TotalWriteCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalWriteCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastWrite", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
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
				return ErrInvalidLengthIndexUsageStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexUsageStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastWrite, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRowsWritten", wireType)
			}
			m.TotalRowsWritten = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalRowsWritten |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndexUsageStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexUsageStats
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
func (m *IndexUsageKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexUsageStats
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
			return fmt.Errorf("proto: IndexUsageKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexUsageKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableID", wireType)
			}
			m.TableID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableID |= TableID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexID", wireType)
			}
			m.IndexID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexID |= IndexID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIndexUsageStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexUsageStats
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
func (m *CollectedIndexUsageStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexUsageStats
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
			return fmt.Errorf("proto: CollectedIndexUsageStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectedIndexUsageStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
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
				return ErrInvalidLengthIndexUsageStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexUsageStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexUsageStats
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
				return ErrInvalidLengthIndexUsageStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexUsageStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexUsageStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexUsageStats
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
func skipIndexUsageStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexUsageStats
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
					return 0, ErrIntOverflowIndexUsageStats
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
					return 0, ErrIntOverflowIndexUsageStats
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
				return 0, ErrInvalidLengthIndexUsageStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndexUsageStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndexUsageStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndexUsageStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexUsageStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndexUsageStats = fmt.Errorf("proto: unexpected end of group")
)
