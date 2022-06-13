// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/concurrency/poison/policy.proto

package poison

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// Policy determines how a request will react to encountering a poisoned
// latch. A poisoned latch is a latch for which the holder is unable to make
// progress. That is, waiters of this latch should not expect to be able to
// acquire this latch "for some time"; in practice this is the case of an
// unavailable Replica.
//
// The name is inspired by Rust's mutexes, which undergo poisoning[^1] when a
// thread panics while holding the mutex.
//
// [^1]: https://doc.rust-lang.org/std/sync/struct.Mutex.html#poisoning
type Policy int32

const (
	// Policy_Wait instructs a request to return an error upon encountering
	// a poisoned latch.
	Policy_Wait Policy = 0
	// Policy_Error instructs a request to return an error upon encountering
	// a poisoned latch.
	Policy_Error Policy = 1
)

var Policy_name = map[int32]string{
	0: "Wait",
	1: "Error",
}

var Policy_value = map[string]int32{
	"Wait":  0,
	"Error": 1,
}

func (x Policy) String() string {
	return proto.EnumName(Policy_name, int32(x))
}

func (Policy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43c896794a9a5973, []int{0}
}

func init() {
	proto.RegisterEnum("cockroach.kv.kvserver.concurrency.poison.Policy", Policy_name, Policy_value)
}

func init() {
	proto.RegisterFile("kv/kvserver/concurrency/poison/policy.proto", fileDescriptor_43c896794a9a5973)
}

var fileDescriptor_43c896794a9a5973 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcc, 0xb1, 0x0a, 0x82, 0x40,
	0x18, 0xc0, 0xf1, 0x3b, 0x28, 0xa9, 0x9b, 0x44, 0x9a, 0x82, 0xbe, 0x59, 0x0a, 0xee, 0x86, 0xde,
	0x20, 0x68, 0x6f, 0x0b, 0xda, 0xec, 0x43, 0x4c, 0x2e, 0xfc, 0xe4, 0xd3, 0x0e, 0x7c, 0x8b, 0x1e,
	0xcb, 0xd1, 0xd1, 0xb1, 0xce, 0x17, 0x89, 0x94, 0xa0, 0xed, 0xbf, 0xfc, 0xfe, 0x6a, 0x67, 0x9d,
	0xb1, 0xae, 0x4a, 0xd9, 0xa5, 0x6c, 0x90, 0x0a, 0x7c, 0x30, 0xa7, 0x05, 0x36, 0xa6, 0xa4, 0xbc,
	0xa2, 0xc2, 0x94, 0x74, 0xcf, 0xb1, 0xd1, 0x25, 0x53, 0x4d, 0x51, 0x8c, 0x84, 0x96, 0x29, 0xc1,
	0x9b, 0xb6, 0x4e, 0xff, 0x98, 0xfe, 0x63, 0x7a, 0x62, 0xeb, 0x55, 0x46, 0x19, 0x8d, 0xc8, 0x7c,
	0x6b, 0xf2, 0xdb, 0x8d, 0x0a, 0x4e, 0xe3, 0x2f, 0x5a, 0xa8, 0xd9, 0x39, 0xc9, 0xeb, 0x50, 0x44,
	0x4b, 0x35, 0x3f, 0x32, 0x13, 0x87, 0xf2, 0x10, 0xb7, 0x6f, 0x10, 0xad, 0x07, 0xd9, 0x79, 0x90,
	0xbd, 0x07, 0xf9, 0xf2, 0x20, 0x9f, 0x03, 0x88, 0x6e, 0x00, 0xd1, 0x0f, 0x20, 0x2e, 0xc1, 0xb4,
	0xbf, 0x06, 0xe3, 0x6f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x16, 0x10, 0x70, 0xbe, 0x00,
	0x00, 0x00,
}