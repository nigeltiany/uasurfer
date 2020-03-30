// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: types/platform/platform.proto

package platform

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Platform int32

const (
	Platform_Unknown      Platform = 0
	Platform_Windows      Platform = 1
	Platform_Mac          Platform = 2
	Platform_Linux        Platform = 3
	Platform_iPad         Platform = 4
	Platform_iPhone       Platform = 5
	Platform_iPod         Platform = 6
	Platform_Blackberry   Platform = 7
	Platform_WindowsPhone Platform = 8
	Platform_Playstation  Platform = 9
	Platform_Xbox         Platform = 10
	Platform_Nintendo     Platform = 11
	Platform_Bot          Platform = 12
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0:  "Unknown",
		1:  "Windows",
		2:  "Mac",
		3:  "Linux",
		4:  "iPad",
		5:  "iPhone",
		6:  "iPod",
		7:  "Blackberry",
		8:  "WindowsPhone",
		9:  "Playstation",
		10: "Xbox",
		11: "Nintendo",
		12: "Bot",
	}
	Platform_value = map[string]int32{
		"Unknown":      0,
		"Windows":      1,
		"Mac":          2,
		"Linux":        3,
		"iPad":         4,
		"iPhone":       5,
		"iPod":         6,
		"Blackberry":   7,
		"WindowsPhone": 8,
		"Playstation":  9,
		"Xbox":         10,
		"Nintendo":     11,
		"Bot":          12,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_types_platform_platform_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_types_platform_platform_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_types_platform_platform_proto_rawDescGZIP(), []int{0}
}

var File_types_platform_platform_proto protoreflect.FileDescriptor

var file_types_platform_platform_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x75, 0x61, 0x73, 0x75, 0x72, 0x66, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2a, 0xac, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x61, 0x63, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x75,
	0x78, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x50, 0x61, 0x64, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x69, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x50, 0x6f,
	0x64, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x65, 0x72, 0x72,
	0x79, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x58, 0x62, 0x6f, 0x78, 0x10, 0x0a,
	0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x6f, 0x10, 0x0b, 0x12, 0x07,
	0x0a, 0x03, 0x42, 0x6f, 0x74, 0x10, 0x0c, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67, 0x65, 0x6c, 0x74, 0x69, 0x61, 0x6e, 0x79,
	0x2f, 0x75, 0x61, 0x73, 0x75, 0x72, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_platform_platform_proto_rawDescOnce sync.Once
	file_types_platform_platform_proto_rawDescData = file_types_platform_platform_proto_rawDesc
)

func file_types_platform_platform_proto_rawDescGZIP() []byte {
	file_types_platform_platform_proto_rawDescOnce.Do(func() {
		file_types_platform_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_platform_platform_proto_rawDescData)
	})
	return file_types_platform_platform_proto_rawDescData
}

var file_types_platform_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_platform_platform_proto_goTypes = []interface{}{
	(Platform)(0), // 0: uasurfer.types.platform.Platform
}
var file_types_platform_platform_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_platform_platform_proto_init() }
func file_types_platform_platform_proto_init() {
	if File_types_platform_platform_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_platform_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_platform_platform_proto_goTypes,
		DependencyIndexes: file_types_platform_platform_proto_depIdxs,
		EnumInfos:         file_types_platform_platform_proto_enumTypes,
	}.Build()
	File_types_platform_platform_proto = out.File
	file_types_platform_platform_proto_rawDesc = nil
	file_types_platform_platform_proto_goTypes = nil
	file_types_platform_platform_proto_depIdxs = nil
}
