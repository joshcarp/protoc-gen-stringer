// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: tests/simple/simple.proto

package simple

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/joshcarp/protoc-gen-stringer/stringer"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Example int32

const (
	Example_exampleBar1  Example = 0
	Example_exampleBar2  Example = 2
	Example_exampleBar3  Example = 3
	Example_exampleBar4  Example = 4
	Example_exampleBar5  Example = 5
	Example_exampleBar6  Example = 6
	Example_exampleBar7  Example = 7
	Example_exampleBar8  Example = 8
	Example_exampleBar9  Example = 9
	Example_exampleBar10 Example = 10
)

// Enum value maps for Example.
var (
	Example_name = map[int32]string{
		0:  "exampleBar1",
		2:  "exampleBar2",
		3:  "exampleBar3",
		4:  "exampleBar4",
		5:  "exampleBar5",
		6:  "exampleBar6",
		7:  "exampleBar7",
		8:  "exampleBar8",
		9:  "exampleBar9",
		10: "exampleBar10",
	}
	Example_value = map[string]int32{
		"exampleBar1":  0,
		"exampleBar2":  2,
		"exampleBar3":  3,
		"exampleBar4":  4,
		"exampleBar5":  5,
		"exampleBar6":  6,
		"exampleBar7":  7,
		"exampleBar8":  8,
		"exampleBar9":  9,
		"exampleBar10": 10,
	}
)

func (x Example) Enum() *Example {
	p := new(Example)
	*p = x
	return p
}

func (x Example) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Example) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_simple_simple_proto_enumTypes[0].Descriptor()
}

func (Example) Type() protoreflect.EnumType {
	return &file_tests_simple_simple_proto_enumTypes[0]
}

func (x Example) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Example.Descriptor instead.
func (Example) EnumDescriptor() ([]byte, []int) {
	return file_tests_simple_simple_proto_rawDescGZIP(), []int{0}
}

var File_tests_simple_simple_proto protoreflect.FileDescriptor

var file_tests_simple_simple_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xdf, 0x02, 0x0a, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x42, 0x61, 0x72, 0x31, 0x10, 0x00, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x31, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x32, 0x10, 0x02, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x32, 0x12, 0x20, 0x0a, 0x0b, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x33, 0x10, 0x03, 0x1a, 0x0f, 0xa2, 0xa9,
	0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x33, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x34, 0x10, 0x04, 0x1a, 0x0f,
	0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x34, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x35, 0x10, 0x05,
	0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72,
	0x35, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x36,
	0x10, 0x06, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42,
	0x61, 0x72, 0x36, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61,
	0x72, 0x37, 0x10, 0x07, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x42, 0x61, 0x72, 0x37, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x42, 0x61, 0x72, 0x38, 0x10, 0x08, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x38, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x42, 0x61, 0x72, 0x39, 0x10, 0x09, 0x1a, 0x0f, 0xa2, 0xa9, 0x20, 0x0b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x39, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x31, 0x30, 0x10, 0x0a, 0x1a, 0x10, 0xa2, 0xa9, 0x20,
	0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x61, 0x72, 0x31, 0x30, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_simple_simple_proto_rawDescOnce sync.Once
	file_tests_simple_simple_proto_rawDescData = file_tests_simple_simple_proto_rawDesc
)

func file_tests_simple_simple_proto_rawDescGZIP() []byte {
	file_tests_simple_simple_proto_rawDescOnce.Do(func() {
		file_tests_simple_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_simple_simple_proto_rawDescData)
	})
	return file_tests_simple_simple_proto_rawDescData
}

var file_tests_simple_simple_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_simple_simple_proto_goTypes = []interface{}{
	(Example)(0), // 0: example.example
}
var file_tests_simple_simple_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_simple_simple_proto_init() }
func file_tests_simple_simple_proto_init() {
	if File_tests_simple_simple_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_simple_simple_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_simple_simple_proto_goTypes,
		DependencyIndexes: file_tests_simple_simple_proto_depIdxs,
		EnumInfos:         file_tests_simple_simple_proto_enumTypes,
	}.Build()
	File_tests_simple_simple_proto = out.File
	file_tests_simple_simple_proto_rawDesc = nil
	file_tests_simple_simple_proto_goTypes = nil
	file_tests_simple_simple_proto_depIdxs = nil
}
