// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package foobar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Hello int32

const (
	Hello_HelloBar1  Hello = 0
	Hello_HelloBar2  Hello = 2
	Hello_HelloBar3  Hello = 3
	Hello_HelloBar4  Hello = 4
	Hello_HelloBar5  Hello = 5
	Hello_HelloBar6  Hello = 6
	Hello_HelloBar7  Hello = 7
	Hello_HelloBar8  Hello = 8
	Hello_HelloBar9  Hello = 9
	Hello_HelloBar10 Hello = 10
)

var Hello_name = map[int32]string{
	0:  "HelloBar1",
	2:  "HelloBar2",
	3:  "HelloBar3",
	4:  "HelloBar4",
	5:  "HelloBar5",
	6:  "HelloBar6",
	7:  "HelloBar7",
	8:  "HelloBar8",
	9:  "HelloBar9",
	10: "HelloBar10",
}

var Hello_value = map[string]int32{
	"HelloBar1":  0,
	"HelloBar2":  2,
	"HelloBar3":  3,
	"HelloBar4":  4,
	"HelloBar5":  5,
	"HelloBar6":  6,
	"HelloBar7":  7,
	"HelloBar8":  8,
	"HelloBar9":  9,
	"HelloBar10": 10,
}

func (x Hello) String() string {
	return proto.EnumName(Hello_name, int32(x))
}

func (Hello) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

type Bar int32

const (
	Bar_BarBar1  Bar = 0
	Bar_BarBar2  Bar = 2
	Bar_BarBar3  Bar = 3
	Bar_BarBar4  Bar = 4
	Bar_BarBar5  Bar = 5
	Bar_BarBar6  Bar = 6
	Bar_BarBar7  Bar = 7
	Bar_BarBar8  Bar = 8
	Bar_BarBar9  Bar = 9
	Bar_BarBar10 Bar = 10
)

var Bar_name = map[int32]string{
	0:  "BarBar1",
	2:  "BarBar2",
	3:  "BarBar3",
	4:  "BarBar4",
	5:  "BarBar5",
	6:  "BarBar6",
	7:  "BarBar7",
	8:  "BarBar8",
	9:  "BarBar9",
	10: "BarBar10",
}

var Bar_value = map[string]int32{
	"BarBar1":  0,
	"BarBar2":  2,
	"BarBar3":  3,
	"BarBar4":  4,
	"BarBar5":  5,
	"BarBar6":  6,
	"BarBar7":  7,
	"BarBar8":  8,
	"BarBar9":  9,
	"BarBar10": 10,
}

func (x Bar) String() string {
	return proto.EnumName(Bar_name, int32(x))
}

func (Bar) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

var E_StringVal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66196,
	Name:          "foobar.string_val",
	Tag:           "bytes,66196,opt,name=string_val",
	Filename:      "hello.proto",
}

func init() {
	proto.RegisterEnum("foobar.Hello", Hello_name, Hello_value)
	proto.RegisterEnum("foobar.Bar", Bar_name, Bar_value)
	proto.RegisterExtension(E_StringVal)
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x86, 0xb5, 0x5f, 0x9b, 0x36, 0x53, 0xa4, 0xc3, 0x58, 0xa4, 0x06, 0xd1, 0x29, 0xa2, 0x48,
	0x17, 0xa9, 0xc9, 0xd4, 0xdf, 0x65, 0x40, 0x70, 0x27, 0xb8, 0xe8, 0x56, 0x26, 0x69, 0x52, 0x0b,
	0xb1, 0x53, 0x4e, 0xa6, 0xe6, 0x0a, 0xbc, 0x03, 0xaf, 0x60, 0x56, 0xce, 0xbd, 0x79, 0x11, 0x32,
	0x49, 0x27, 0x7f, 0xb8, 0x6b, 0xf3, 0x3e, 0xe7, 0xf0, 0x9e, 0x87, 0x41, 0xf3, 0xef, 0x69, 0x9e,
	0x0b, 0xff, 0x02, 0x42, 0x0a, 0xe2, 0x64, 0x42, 0xc4, 0x1c, 0x3c, 0x7a, 0x14, 0xe2, 0x98, 0xa7,
	0xdb, 0xea, 0x6b, 0x7c, 0xcd, 0xb6, 0x87, 0xb4, 0x48, 0xe0, 0x74, 0x91, 0x02, 0x6a, 0x72, 0xf3,
	0x67, 0x84, 0x26, 0x9f, 0xcd, 0x24, 0xf1, 0x90, 0x5b, 0xfd, 0x88, 0x38, 0x04, 0xf8, 0x9e, 0x37,
	0x57, 0x9a, 0x4e, 0x63, 0x0e, 0x41, 0xc8, 0x18, 0x79, 0xd2, 0x66, 0x21, 0x1e, 0x79, 0x0f, 0x94,
	0xa6, 0x6e, 0xcc, 0x83, 0x90, 0x05, 0x21, 0x03, 0xf2, 0xb8, 0x4d, 0x19, 0xbe, 0xf3, 0x90, 0xd2,
	0xd4, 0x89, 0x83, 0x90, 0xf1, 0x5e, 0xb4, 0xc3, 0xe3, 0x5b, 0xc4, 0x87, 0x53, 0x6f, 0xf0, 0xa4,
	0x17, 0x75, 0xaa, 0xbc, 0xc5, 0x8e, 0xad, 0x12, 0x84, 0x6c, 0xd7, 0x1b, 0x7b, 0x87, 0xa7, 0x76,
	0x2c, 0x61, 0x12, 0xba, 0x2d, 0xdf, 0xe3, 0x99, 0x6d, 0x59, 0xa6, 0x12, 0x4a, 0x00, 0xf2, 0xb4,
	0x4d, 0x3f, 0x60, 0xd7, 0x5b, 0x28, 0x4d, 0xe7, 0x31, 0x4f, 0x24, 0x64, 0xbc, 0x38, 0x64, 0x40,
	0x9e, 0x21, 0xd4, 0xdc, 0xff, 0x1a, 0x23, 0x0b, 0x94, 0x89, 0x84, 0x32, 0x05, 0x09, 0x9b, 0xbf,
	0x23, 0x74, 0x17, 0x71, 0x20, 0x6b, 0x34, 0x8d, 0x38, 0xdc, 0x34, 0x2d, 0x95, 0xa6, 0xb8, 0xfe,
	0x1b, 0x71, 0xb0, 0xbe, 0x9e, 0x5b, 0xc4, 0xd8, 0x7a, 0xa4, 0x34, 0x25, 0x1d, 0xc4, 0x6a, 0xa3,
	0x16, 0x32, 0xd2, 0x1e, 0x2a, 0x4d, 0x17, 0x2d, 0x54, 0xdb, 0x6b, 0x08, 0xe3, 0x6e, 0x40, 0xf0,
	0xfe, 0x0e, 0xa3, 0xf0, 0xbf, 0x44, 0xd3, 0xd6, 0x98, 0x1c, 0xb6, 0xad, 0x95, 0x36, 0x4b, 0x8c,
	0xd0, 0xe1, 0x92, 0xca, 0x6c, 0x73, 0x8f, 0xf1, 0x3a, 0xbc, 0xc7, 0x0a, 0x7e, 0x61, 0x21, 0xa3,
	0x77, 0xa5, 0x34, 0x5d, 0x76, 0xd7, 0x34, 0x9e, 0x5f, 0xa2, 0xd9, 0x4d, 0x9f, 0xb1, 0x3c, 0xe4,
	0x1a, 0xdd, 0x1f, 0x23, 0x84, 0x0a, 0x09, 0xa7, 0xf3, 0xf1, 0xdb, 0x4f, 0x9e, 0x93, 0xb5, 0x5f,
	0x3f, 0x65, 0xdf, 0x3e, 0x65, 0xff, 0xd3, 0xf9, 0xfa, 0x63, 0xcf, 0xf3, 0x6b, 0xfa, 0xe5, 0x22,
	0x4f, 0xe2, 0x5c, 0xac, 0x7e, 0xff, 0x1a, 0xd3, 0xfb, 0xaf, 0xdc, 0xaf, 0x6e, 0x3d, 0xb6, 0xe7,
	0x79, 0xec, 0x54, 0x34, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x28, 0x83, 0x8d, 0x18, 0x1d, 0x03,
	0x00, 0x00,
}
