// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/serverpb/index_recommendations.proto

package serverpb

import (
	fmt "fmt"
	github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
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

type IndexRecommendation_RecommendationType int32

const (
	IndexRecommendation_DROP_UNUSED IndexRecommendation_RecommendationType = 0
)

var IndexRecommendation_RecommendationType_name = map[int32]string{
	0: "DROP_UNUSED",
}

var IndexRecommendation_RecommendationType_value = map[string]int32{
	"DROP_UNUSED": 0,
}

func (x IndexRecommendation_RecommendationType) Enum() *IndexRecommendation_RecommendationType {
	p := new(IndexRecommendation_RecommendationType)
	*p = x
	return p
}

func (x IndexRecommendation_RecommendationType) String() string {
	return proto.EnumName(IndexRecommendation_RecommendationType_name, int32(x))
}

func (x *IndexRecommendation_RecommendationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IndexRecommendation_RecommendationType_value, data, "IndexRecommendation_RecommendationType")
	if err != nil {
		return err
	}
	*x = IndexRecommendation_RecommendationType(value)
	return nil
}

func (IndexRecommendation_RecommendationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0593af027d009b10, []int{0, 0}
}

type IndexRecommendation struct {
	// TableID is the ID of the table this index is created on. This is same as
	// descpb.TableID and is unique within the cluster.
	TableID github_com_cockroachdb_cockroach_pkg_roachpb.TableID `protobuf:"varint,1,opt,name=table_id,json=tableId,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.TableID" json:"table_id"`
	// IndexID is the ID of the index within the scope of the given table.
	IndexID github_com_cockroachdb_cockroach_pkg_roachpb.IndexID `protobuf:"varint,2,opt,name=index_id,json=indexId,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.IndexID" json:"index_id"`
	// Type of recommendation for the index.
	Type IndexRecommendation_RecommendationType `protobuf:"varint,3,opt,name=type,enum=cockroach.sql.IndexRecommendation_RecommendationType" json:"type"`
	// Reason for our recommendation type.
	Reason string `protobuf:"bytes,4,opt,name=reason" json:"reason"`
}

func (m *IndexRecommendation) Reset()         { *m = IndexRecommendation{} }
func (m *IndexRecommendation) String() string { return proto.CompactTextString(m) }
func (*IndexRecommendation) ProtoMessage()    {}
func (*IndexRecommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0593af027d009b10, []int{0}
}
func (m *IndexRecommendation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexRecommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IndexRecommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRecommendation.Merge(m, src)
}
func (m *IndexRecommendation) XXX_Size() int {
	return m.Size()
}
func (m *IndexRecommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRecommendation.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRecommendation proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.sql.IndexRecommendation_RecommendationType", IndexRecommendation_RecommendationType_name, IndexRecommendation_RecommendationType_value)
	proto.RegisterType((*IndexRecommendation)(nil), "cockroach.sql.IndexRecommendation")
}

func init() {
	proto.RegisterFile("server/serverpb/index_recommendations.proto", fileDescriptor_0593af027d009b10)
}

var fileDescriptor_0593af027d009b10 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x05, 0x49, 0xfa, 0x99, 0x79, 0x29, 0xa9, 0x15, 0xf1, 0x45, 0xa9,
	0xc9, 0xf9, 0xb9, 0xb9, 0xa9, 0x79, 0x29, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xbc, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19, 0x7a, 0xc5,
	0x85, 0x39, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x19, 0x7d, 0x10, 0x0b, 0xa2, 0x48, 0xe9,
	0x33, 0x13, 0x97, 0xb0, 0x27, 0xc8, 0x90, 0x20, 0x14, 0x33, 0x84, 0x92, 0xb8, 0x38, 0x4a, 0x12,
	0x93, 0x72, 0x52, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x9d, 0xdc, 0x4f, 0xdc,
	0x93, 0x67, 0x78, 0x74, 0x4f, 0x9e, 0x3d, 0x04, 0x24, 0xee, 0xe9, 0xf2, 0xeb, 0x9e, 0xbc, 0x49,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xdc, 0xb2, 0x94, 0x24, 0x04,
	0x5b, 0xbf, 0x20, 0x3b, 0x5d, 0x1f, 0xcc, 0x2a, 0x48, 0xd2, 0x83, 0xea, 0x0b, 0x62, 0x07, 0x1b,
	0xec, 0x99, 0x02, 0xb2, 0x03, 0xe2, 0xfe, 0xcc, 0x14, 0x09, 0x26, 0x54, 0x3b, 0xc0, 0x4e, 0x22,
	0xc3, 0x0e, 0xa8, 0xbe, 0x20, 0x76, 0xb0, 0xc1, 0x9e, 0x29, 0x42, 0xfe, 0x5c, 0x2c, 0x25, 0x95,
	0x05, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0xa6, 0x7a, 0x28, 0x61, 0xa2, 0x87, 0xc5,
	0xe7, 0x7a, 0xa8, 0xdc, 0x90, 0xca, 0x82, 0x54, 0x27, 0x16, 0x90, 0xb3, 0x82, 0xc0, 0x06, 0x09,
	0xc9, 0x70, 0xb1, 0x15, 0xa5, 0x26, 0x16, 0xe7, 0xe7, 0x49, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x42,
	0xe5, 0xa0, 0x62, 0x4a, 0xaa, 0x5c, 0x42, 0x98, 0xfa, 0x85, 0xf8, 0xb9, 0xb8, 0x5d, 0x82, 0xfc,
	0x03, 0xe2, 0x43, 0xfd, 0x42, 0x83, 0x5d, 0x5d, 0x04, 0x18, 0x9c, 0xb4, 0x4e, 0x3c, 0x94, 0x63,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x1b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x38,
	0x60, 0x51, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xcc, 0x14, 0x7c, 0xf4, 0x01, 0x00, 0x00,
}

func (m *IndexRecommendation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexRecommendation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexRecommendation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Reason)
	copy(dAtA[i:], m.Reason)
	i = encodeVarintIndexRecommendations(dAtA, i, uint64(len(m.Reason)))
	i--
	dAtA[i] = 0x22
	i = encodeVarintIndexRecommendations(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x18
	i = encodeVarintIndexRecommendations(dAtA, i, uint64(m.IndexID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintIndexRecommendations(dAtA, i, uint64(m.TableID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintIndexRecommendations(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexRecommendations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexRecommendation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovIndexRecommendations(uint64(m.TableID))
	n += 1 + sovIndexRecommendations(uint64(m.IndexID))
	n += 1 + sovIndexRecommendations(uint64(m.Type))
	l = len(m.Reason)
	n += 1 + l + sovIndexRecommendations(uint64(l))
	return n
}

func sovIndexRecommendations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexRecommendations(x uint64) (n int) {
	return sovIndexRecommendations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexRecommendation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexRecommendations
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
			return fmt.Errorf("proto: IndexRecommendation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexRecommendation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableID", wireType)
			}
			m.TableID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexRecommendations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableID |= github_com_cockroachdb_cockroach_pkg_roachpb.TableID(b&0x7F) << shift
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
					return ErrIntOverflowIndexRecommendations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexID |= github_com_cockroachdb_cockroach_pkg_roachpb.IndexID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexRecommendations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= IndexRecommendation_RecommendationType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexRecommendations
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
				return ErrInvalidLengthIndexRecommendations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexRecommendations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexRecommendations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexRecommendations
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
func skipIndexRecommendations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexRecommendations
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
					return 0, ErrIntOverflowIndexRecommendations
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
					return 0, ErrIntOverflowIndexRecommendations
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
				return 0, ErrInvalidLengthIndexRecommendations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndexRecommendations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndexRecommendations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndexRecommendations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexRecommendations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndexRecommendations = fmt.Errorf("proto: unexpected end of group")
)
