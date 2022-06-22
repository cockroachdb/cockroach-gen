// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/readsummary/rspb/summary.proto

package rspb

import (
	fmt "fmt"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
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

// ReadSummary contains a summary of all read requests served on a range, along
// with the timestamps that those reads were served at. The structure is a
// summary in that it may not represent these reads with perfect precision.
// Instead, it is allowed to lose resolution in exchange for reduced space, as
// long as the resulting timestamp for each key does not regress. During such
// compression, the timestamp of a given key is only allowed to advance as
// precision drops. This parallels a similar ratcheting policy in the timestamp
// cache (tscache.Cache).
//
// For example, a high-resolution version of the summary may look like:
//
//                         #
//    ^       ##     ##    #
// ts |    #  ##     ##    #    #######
//    |    #  ##     ####  #    #######   ##
//       ###  ##     #######    #######   ##  ###
//       ----------- ----------------------------
//          local               global
//
// While a low-resolution (compressed) version of the summary may look like:
//
//                   ############################
//    ^  ########### ############################
// ts |  ########### ############################
//    |  ########### ############################
//       ########### ############################
//       ----------- ----------------------------
//          local               global
//
type ReadSummary struct {
	Local  Segment `protobuf:"bytes,1,opt,name=local,proto3" json:"local"`
	Global Segment `protobuf:"bytes,2,opt,name=global,proto3" json:"global"`
}

func (m *ReadSummary) Reset()         { *m = ReadSummary{} }
func (m *ReadSummary) String() string { return proto.CompactTextString(m) }
func (*ReadSummary) ProtoMessage()    {}
func (*ReadSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b9794767da51e44, []int{0}
}
func (m *ReadSummary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReadSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSummary.Merge(m, src)
}
func (m *ReadSummary) XXX_Size() int {
	return m.Size()
}
func (m *ReadSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSummary proto.InternalMessageInfo

// Segment is a segement of a Range's complete ReadSummary. A ReadSummary has a
// segment for each of the Range's replicated, addressable contiguous keyspaces
// (i.e. range-local and global).
type Segment struct {
	LowWater hlc.Timestamp `protobuf:"bytes,1,opt,name=low_water,json=lowWater,proto3" json:"low_water"`
}

func (m *Segment) Reset()         { *m = Segment{} }
func (m *Segment) String() string { return proto.CompactTextString(m) }
func (*Segment) ProtoMessage()    {}
func (*Segment) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b9794767da51e44, []int{1}
}
func (m *Segment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Segment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Segment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segment.Merge(m, src)
}
func (m *Segment) XXX_Size() int {
	return m.Size()
}
func (m *Segment) XXX_DiscardUnknown() {
	xxx_messageInfo_Segment.DiscardUnknown(m)
}

var xxx_messageInfo_Segment proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReadSummary)(nil), "cockroach.kv.kvserver.readsummary.ReadSummary")
	proto.RegisterType((*Segment)(nil), "cockroach.kv.kvserver.readsummary.Segment")
}

func init() {
	proto.RegisterFile("kv/kvserver/readsummary/rspb/summary.proto", fileDescriptor_7b9794767da51e44)
}

var fileDescriptor_7b9794767da51e44 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x2e, 0xd3, 0xcf,
	0x2e, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x4a, 0x4d, 0x4c, 0x29, 0x2e, 0xcd, 0xcd,
	0x4d, 0x2c, 0xaa, 0xd4, 0x2f, 0x2a, 0x2e, 0x48, 0xd2, 0x87, 0x72, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x14, 0x93, 0xf3, 0x93, 0xb3, 0x8b, 0xf2, 0x13, 0x93, 0x33, 0xf4, 0xb2, 0xcb, 0xf4,
	0x60, 0xba, 0xf4, 0x90, 0x74, 0x49, 0x49, 0x94, 0x96, 0x64, 0xe6, 0xe8, 0x67, 0xe4, 0x24, 0xeb,
	0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x34, 0x4b, 0x89, 0xa4, 0xe7, 0xa7,
	0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x69, 0x29, 0x23, 0x17, 0x77, 0x50, 0x6a, 0x62,
	0x4a, 0x30, 0x44, 0xbf, 0x90, 0x1b, 0x17, 0x6b, 0x4e, 0x7e, 0x72, 0x62, 0x8e, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xb7, 0x91, 0x96, 0x1e, 0x41, 0x2b, 0xf5, 0x82, 0x53, 0xd3, 0x73, 0x53, 0xf3, 0x4a,
	0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0x68, 0x17, 0xf2, 0xe0, 0x62, 0x4b, 0xcf, 0xc9,
	0x4f, 0x4a, 0xcc, 0x91, 0x60, 0x22, 0xd3, 0x20, 0xa8, 0x7e, 0x2b, 0x96, 0x17, 0x0b, 0xe4, 0x19,
	0x95, 0x02, 0xb9, 0xd8, 0xa1, 0xd2, 0x42, 0x0e, 0x5c, 0x9c, 0x39, 0xf9, 0xe5, 0xf1, 0xe5, 0x89,
	0x25, 0xa9, 0x45, 0x50, 0x67, 0xca, 0x22, 0x99, 0x0e, 0x0a, 0x00, 0xbd, 0x8c, 0x9c, 0x64, 0xbd,
	0x10, 0x58, 0x00, 0x40, 0x0d, 0xe4, 0xc8, 0xc9, 0x2f, 0x0f, 0x07, 0x69, 0x82, 0x18, 0xe9, 0xa4,
	0x76, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xc5, 0x02, 0x8a, 0x83, 0x24, 0x36, 0x70, 0x48, 0x19, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x74, 0x7b, 0xe2, 0x8b, 0xaa, 0x01, 0x00, 0x00,
}

func (this *ReadSummary) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReadSummary)
	if !ok {
		that2, ok := that.(ReadSummary)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Local.Equal(&that1.Local) {
		return false
	}
	if !this.Global.Equal(&that1.Global) {
		return false
	}
	return true
}
func (this *Segment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Segment)
	if !ok {
		that2, ok := that.(Segment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.LowWater.Equal(&that1.LowWater) {
		return false
	}
	return true
}
func (m *ReadSummary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadSummary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadSummary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Global.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSummary(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Local.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSummary(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Segment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Segment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Segment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LowWater.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSummary(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSummary(dAtA []byte, offset int, v uint64) int {
	offset -= sovSummary(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReadSummary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Local.Size()
	n += 1 + l + sovSummary(uint64(l))
	l = m.Global.Size()
	n += 1 + l + sovSummary(uint64(l))
	return n
}

func (m *Segment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LowWater.Size()
	n += 1 + l + sovSummary(uint64(l))
	return n
}

func sovSummary(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSummary(x uint64) (n int) {
	return sovSummary(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReadSummary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSummary
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
			return fmt.Errorf("proto: ReadSummary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadSummary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Local", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSummary
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
				return ErrInvalidLengthSummary
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSummary
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Local.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Global", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSummary
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
				return ErrInvalidLengthSummary
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSummary
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Global.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSummary(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSummary
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
func (m *Segment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSummary
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
			return fmt.Errorf("proto: Segment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Segment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowWater", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSummary
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
				return ErrInvalidLengthSummary
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSummary
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LowWater.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSummary(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSummary
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
func skipSummary(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSummary
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
					return 0, ErrIntOverflowSummary
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
					return 0, ErrIntOverflowSummary
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
				return 0, ErrInvalidLengthSummary
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSummary
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSummary
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSummary        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSummary          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSummary = fmt.Errorf("proto: unexpected end of group")
)
