// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ccl/storageccl/engineccl/enginepbccl/stats.proto

package enginepbccl

import (
	fmt "fmt"
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

// EncryptionStatus contains encryption-related information.
type EncryptionStatus struct {
	// Information about the active store key, if any.
	ActiveStoreKey *KeyInfo `protobuf:"bytes,1,opt,name=active_store_key,json=activeStoreKey,proto3" json:"active_store_key,omitempty"`
	// Information about the active data key, if any.
	ActiveDataKey *KeyInfo `protobuf:"bytes,2,opt,name=active_data_key,json=activeDataKey,proto3" json:"active_data_key,omitempty"`
}

func (m *EncryptionStatus) Reset()         { *m = EncryptionStatus{} }
func (m *EncryptionStatus) String() string { return proto.CompactTextString(m) }
func (*EncryptionStatus) ProtoMessage()    {}
func (*EncryptionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_372d12b2ba58675a, []int{0}
}
func (m *EncryptionStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EncryptionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionStatus.Merge(m, src)
}
func (m *EncryptionStatus) XXX_Size() int {
	return m.Size()
}
func (m *EncryptionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EncryptionStatus)(nil), "cockroach.ccl.storageccl.engineccl.enginepbccl.EncryptionStatus")
}

func init() {
	proto.RegisterFile("ccl/storageccl/engineccl/enginepbccl/stats.proto", fileDescriptor_372d12b2ba58675a)
}

var fileDescriptor_372d12b2ba58675a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0x4e, 0xce, 0xd1,
	0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x05, 0x31, 0x53, 0xf3, 0xd2, 0x33, 0xf3, 0x90, 0x58,
	0x05, 0x49, 0x10, 0x05, 0x89, 0x25, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7a, 0xc9,
	0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19, 0x7a, 0xc9, 0xc9, 0x39, 0x7a, 0x08, 0xbd, 0x7a,
	0x70, 0xbd, 0x7a, 0x48, 0x7a, 0xa5, 0xcc, 0x89, 0xb2, 0x21, 0x3b, 0xb5, 0x32, 0xbe, 0x28, 0x35,
	0x3d, 0xb3, 0xb8, 0xa4, 0xa8, 0x12, 0x62, 0x91, 0xd2, 0x35, 0x46, 0x2e, 0x01, 0xd7, 0xbc, 0xe4,
	0xa2, 0xca, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xa1, 0x44, 0x2e,
	0x81, 0xc4, 0xe4, 0x92, 0xcc, 0xb2, 0xd4, 0x78, 0x90, 0x91, 0xa9, 0xf1, 0xd9, 0xa9, 0x95, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xe6, 0x24, 0x3a, 0x4c, 0xcf, 0x3b, 0xb5, 0xd2, 0x33, 0x2f,
	0x2d, 0x3f, 0x88, 0x0f, 0x62, 0x60, 0x30, 0xc8, 0x3c, 0xef, 0xd4, 0x4a, 0xa1, 0x78, 0x2e, 0x7e,
	0xa8, 0x15, 0x29, 0x89, 0x25, 0x89, 0x60, 0x1b, 0x98, 0x28, 0xb3, 0x81, 0x17, 0x62, 0x9e, 0x4b,
	0x62, 0x49, 0xa2, 0x77, 0x6a, 0xa5, 0x93, 0xee, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x37, 0x92, 0x29, 0x49, 0x6c,
	0xe0, 0xe0, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x58, 0xbd, 0xb4, 0x9c, 0xab, 0x01, 0x00,
	0x00,
}

func (m *EncryptionStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptionStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptionStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActiveDataKey != nil {
		{
			size, err := m.ActiveDataKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStats(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ActiveStoreKey != nil {
		{
			size, err := m.ActiveStoreKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStats(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EncryptionStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ActiveStoreKey != nil {
		l = m.ActiveStoreKey.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	if m.ActiveDataKey != nil {
		l = m.ActiveDataKey.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EncryptionStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: EncryptionStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptionStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveStoreKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActiveStoreKey == nil {
				m.ActiveStoreKey = &KeyInfo{}
			}
			if err := m.ActiveStoreKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDataKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActiveDataKey == nil {
				m.ActiveDataKey = &KeyInfo{}
			}
			if err := m.ActiveDataKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)