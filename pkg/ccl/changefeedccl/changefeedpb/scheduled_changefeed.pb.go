// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ccl/changefeedccl/changefeedpb/scheduled_changefeed.proto

package changefeedpb

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

// ScheduledExportExecutionArgs is the arguments to the scheduled backup executor.
type ScheduledChangefeedExecutionArgs struct {
	ChangefeedStatement string `protobuf:"bytes,1,opt,name=changefeed_statement,json=changefeedStatement,proto3" json:"changefeed_statement,omitempty"`
}

func (m *ScheduledChangefeedExecutionArgs) Reset()         { *m = ScheduledChangefeedExecutionArgs{} }
func (m *ScheduledChangefeedExecutionArgs) String() string { return proto.CompactTextString(m) }
func (*ScheduledChangefeedExecutionArgs) ProtoMessage()    {}
func (*ScheduledChangefeedExecutionArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8434c882fdd9312a, []int{0}
}
func (m *ScheduledChangefeedExecutionArgs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledChangefeedExecutionArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduledChangefeedExecutionArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledChangefeedExecutionArgs.Merge(m, src)
}
func (m *ScheduledChangefeedExecutionArgs) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledChangefeedExecutionArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledChangefeedExecutionArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledChangefeedExecutionArgs proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduledChangefeedExecutionArgs)(nil), "cockroach.ccl.changefeedccl.ScheduledChangefeedExecutionArgs")
}

func init() {
	proto.RegisterFile("ccl/changefeedccl/changefeedpb/scheduled_changefeed.proto", fileDescriptor_8434c882fdd9312a)
}

var fileDescriptor_8434c882fdd9312a = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0x4e, 0xce, 0xd1,
	0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x4d, 0x4b, 0x4d, 0x4d, 0x41, 0xe5, 0x15, 0x24, 0xe9, 0x17,
	0x27, 0x67, 0xa4, 0xa6, 0x94, 0xe6, 0xa4, 0xa6, 0xc4, 0x23, 0x84, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0xa4, 0x93, 0xf3, 0x93, 0xb3, 0x8b, 0xf2, 0x13, 0x93, 0x33, 0xf4, 0x92, 0x93, 0x73,
	0xf4, 0x50, 0x0c, 0x51, 0x0a, 0xe5, 0x52, 0x08, 0x86, 0x69, 0x75, 0x86, 0xcb, 0xb8, 0x56, 0xa4,
	0x26, 0x97, 0x96, 0x64, 0xe6, 0xe7, 0x39, 0x16, 0xa5, 0x17, 0x0b, 0x19, 0x72, 0x89, 0x20, 0x34,
	0xc5, 0x17, 0x97, 0x24, 0x96, 0xa4, 0xe6, 0xa6, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x09, 0x23, 0xe4, 0x82, 0x61, 0x52, 0x4e, 0x7a, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x8d, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x3c, 0xc8, 0xce, 0x4f,
	0x62, 0x03, 0x3b, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x35, 0xcb, 0x3a, 0x6c, 0xe7, 0x00,
	0x00, 0x00,
}

func (m *ScheduledChangefeedExecutionArgs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledChangefeedExecutionArgs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledChangefeedExecutionArgs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChangefeedStatement) > 0 {
		i -= len(m.ChangefeedStatement)
		copy(dAtA[i:], m.ChangefeedStatement)
		i = encodeVarintScheduledChangefeed(dAtA, i, uint64(len(m.ChangefeedStatement)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduledChangefeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduledChangefeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduledChangefeedExecutionArgs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChangefeedStatement)
	if l > 0 {
		n += 1 + l + sovScheduledChangefeed(uint64(l))
	}
	return n
}

func sovScheduledChangefeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduledChangefeed(x uint64) (n int) {
	return sovScheduledChangefeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduledChangefeedExecutionArgs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduledChangefeed
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
			return fmt.Errorf("proto: ScheduledChangefeedExecutionArgs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledChangefeedExecutionArgs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangefeedStatement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledChangefeed
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
				return ErrInvalidLengthScheduledChangefeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledChangefeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChangefeedStatement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduledChangefeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScheduledChangefeed
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
func skipScheduledChangefeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduledChangefeed
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
					return 0, ErrIntOverflowScheduledChangefeed
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
					return 0, ErrIntOverflowScheduledChangefeed
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
				return 0, ErrInvalidLengthScheduledChangefeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduledChangefeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduledChangefeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduledChangefeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduledChangefeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduledChangefeed = fmt.Errorf("proto: unexpected end of group")
)