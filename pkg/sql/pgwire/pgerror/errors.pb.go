// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/pgwire/pgerror/errors.proto

package pgerror

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

// Error contains all Postgres wire protocol error fields.
// See https://www.postgresql.org/docs/current/static/protocol-error-fields.html
// for a list of all Postgres error fields, most of which are optional and can
// be used to provide auxiliary error information.
type Error struct {
	// standard pg error fields. This can be passed
	// over the pg wire protocol.
	Code           string        `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message        string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail         string        `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Hint           string        `protobuf:"bytes,4,opt,name=hint,proto3" json:"hint,omitempty"`
	Severity       string        `protobuf:"bytes,8,opt,name=severity,proto3" json:"severity,omitempty"`
	ConstraintName string        `protobuf:"bytes,9,opt,name=constraint_name,json=constraintName,proto3" json:"constraint_name,omitempty"`
	Source         *Error_Source `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fd880909da5328, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

type Error_Source struct {
	File     string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Line     int32  `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
}

func (m *Error_Source) Reset()         { *m = Error_Source{} }
func (m *Error_Source) String() string { return proto.CompactTextString(m) }
func (*Error_Source) ProtoMessage()    {}
func (*Error_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_79fd880909da5328, []int{0, 0}
}
func (m *Error_Source) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Error_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error_Source.Merge(m, src)
}
func (m *Error_Source) XXX_Size() int {
	return m.Size()
}
func (m *Error_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Error_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Error_Source proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Error)(nil), "cockroach.pgerror.Error")
	proto.RegisterType((*Error_Source)(nil), "cockroach.pgerror.Error.Source")
}

func init() { proto.RegisterFile("sql/pgwire/pgerror/errors.proto", fileDescriptor_79fd880909da5328) }

var fileDescriptor_79fd880909da5328 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xfe, 0x69, 0x92, 0xfa, 0x97, 0xa0, 0x78, 0x40, 0x56, 0x07, 0xb7, 0x62, 0xa1,
	0x2c, 0xa9, 0x04, 0x03, 0x3b, 0x12, 0x4b, 0x85, 0x18, 0xca, 0xc6, 0x82, 0x8c, 0xeb, 0xa6, 0x16,
	0x89, 0x1d, 0x6c, 0x17, 0xc4, 0x5b, 0xf0, 0x3a, 0xbc, 0x41, 0xc7, 0x8e, 0x1d, 0x21, 0x79, 0x11,
	0x64, 0x27, 0x84, 0x81, 0xc5, 0x3a, 0xe7, 0xdc, 0xab, 0xab, 0xf3, 0xc9, 0x70, 0x6c, 0x9e, 0xf3,
	0x59, 0x99, 0xbd, 0x0a, 0xcd, 0x67, 0x65, 0xc6, 0xb5, 0x56, 0x7a, 0xe6, 0x5f, 0x93, 0x96, 0x5a,
	0x59, 0x85, 0x8e, 0x98, 0x62, 0x4f, 0x5a, 0x51, 0xb6, 0x4e, 0xdb, 0xf9, 0xc9, 0x47, 0x0f, 0xf6,
	0xaf, 0x9d, 0x42, 0x08, 0x86, 0x4c, 0x2d, 0x39, 0x06, 0x13, 0x30, 0x1d, 0x2c, 0xbc, 0x46, 0x18,
	0xc6, 0x05, 0x37, 0x86, 0x66, 0x1c, 0xf7, 0x7c, 0xfc, 0x63, 0xd1, 0x31, 0x8c, 0x96, 0xdc, 0x52,
	0x91, 0xe3, 0x7f, 0x7e, 0xd0, 0x3a, 0x77, 0x65, 0x2d, 0xa4, 0xc5, 0x61, 0x73, 0xc5, 0x69, 0x34,
	0x82, 0x89, 0xe1, 0x2f, 0x5c, 0x0b, 0xfb, 0x86, 0x13, 0x9f, 0x77, 0x1e, 0x9d, 0xc2, 0x43, 0xa6,
	0xa4, 0xb1, 0x9a, 0x0a, 0x69, 0x1f, 0x24, 0x2d, 0x38, 0x1e, 0xf8, 0x95, 0x83, 0xdf, 0xf8, 0x96,
	0x16, 0x1c, 0x5d, 0xc2, 0xc8, 0xa8, 0x8d, 0x66, 0x1c, 0xf7, 0x27, 0x60, 0xfa, 0xff, 0x7c, 0x9c,
	0xfe, 0x81, 0x49, 0x3d, 0x48, 0x7a, 0xe7, 0xd7, 0x16, 0xed, 0xfa, 0xe8, 0x06, 0x46, 0x4d, 0xe2,
	0xba, 0xad, 0x44, 0xde, 0x11, 0x3a, 0xed, 0xb2, 0x5c, 0xc8, 0x06, 0xaf, 0xbf, 0xf0, 0xda, 0xf5,
	0x5d, 0x6d, 0x24, 0xb3, 0x42, 0xc9, 0x96, 0xae, 0xf3, 0xf3, 0x30, 0x89, 0x86, 0xf1, 0x3c, 0x4c,
	0xe2, 0x61, 0x72, 0x75, 0xb6, 0xfd, 0x22, 0xc1, 0xb6, 0x22, 0x60, 0x57, 0x11, 0xb0, 0xaf, 0x08,
	0xf8, 0xac, 0x08, 0x78, 0xaf, 0x49, 0xb0, 0xab, 0x49, 0xb0, 0xaf, 0x49, 0x70, 0x1f, 0xb7, 0xcd,
	0x1e, 0x23, 0xff, 0x01, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x5e, 0x1b, 0x8b, 0xa3,
	0x01, 0x00, 0x00,
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConstraintName) > 0 {
		i -= len(m.ConstraintName)
		copy(dAtA[i:], m.ConstraintName)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.ConstraintName)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Severity) > 0 {
		i -= len(m.Severity)
		copy(dAtA[i:], m.Severity)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Severity)))
		i--
		dAtA[i] = 0x42
	}
	if m.Source != nil {
		{
			size, err := m.Source.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintErrors(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Hint) > 0 {
		i -= len(m.Hint)
		copy(dAtA[i:], m.Hint)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Hint)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Detail) > 0 {
		i -= len(m.Detail)
		copy(dAtA[i:], m.Detail)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Detail)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Error_Source) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error_Source) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error_Source) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Function) > 0 {
		i -= len(m.Function)
		copy(dAtA[i:], m.Function)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Function)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Line != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Line))
		i--
		dAtA[i] = 0x10
	}
	if len(m.File) > 0 {
		i -= len(m.File)
		copy(dAtA[i:], m.File)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.File)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Detail)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Hint)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Source != nil {
		l = m.Source.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.Severity)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	l = len(m.ConstraintName)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func (m *Error_Source) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.Line != 0 {
		n += 1 + sovErrors(uint64(m.Line))
	}
	l = len(m.Function)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Source == nil {
				m.Source = &Error_Source{}
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Severity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstraintName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConstraintName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
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
func (m *Error_Source) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Source: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Source: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Function", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Function = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
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
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
