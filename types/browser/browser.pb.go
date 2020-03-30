// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: types/browser/browser.proto

package browser

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

type BrowserType int32

const (
	BrowserType_unknown           BrowserType = 0
	BrowserType_Chrome            BrowserType = 1
	BrowserType_Internet_Explorer BrowserType = 2
	BrowserType_Safari            BrowserType = 3
	BrowserType_Firefox           BrowserType = 4
	BrowserType_Android           BrowserType = 5
	BrowserType_Opera             BrowserType = 6
	BrowserType_Blackberry        BrowserType = 7
	BrowserType_UC_Browser        BrowserType = 8
	BrowserType_Silk              BrowserType = 9
	BrowserType_Nokia             BrowserType = 10
	BrowserType_NetFront          BrowserType = 11
	BrowserType_QQ                BrowserType = 12
	BrowserType_Maxthon           BrowserType = 13
	BrowserType_Sogou_Explorer    BrowserType = 14
	BrowserType_Spotify           BrowserType = 15
	BrowserType_Nintendo          BrowserType = 16
	BrowserType_Samsung           BrowserType = 17
	BrowserType_Yandex            BrowserType = 18
	BrowserType_Coc_Coc           BrowserType = 19
	// ----- Bot List ----- //
	BrowserType_Bot           BrowserType = 20
	BrowserType_AppleBot      BrowserType = 21
	BrowserType_BaiduBot      BrowserType = 22
	BrowserType_BingBot       BrowserType = 23
	BrowserType_DuckDuckGoBot BrowserType = 24
	BrowserType_FacebookBot   BrowserType = 25
	BrowserType_GoogleBot     BrowserType = 26
	BrowserType_LinkedInBot   BrowserType = 27
	BrowserType_MsnBot        BrowserType = 28
	BrowserType_PingdomBot    BrowserType = 29
	BrowserType_TwitterBot    BrowserType = 30
	BrowserType_YandexBot     BrowserType = 31
	BrowserType_CocCocBot     BrowserType = 32
	BrowserType_YahooBot      BrowserType = 33
)

// Enum value maps for BrowserType.
var (
	BrowserType_name = map[int32]string{
		0:  "unknown",
		1:  "Chrome",
		2:  "Internet_Explorer",
		3:  "Safari",
		4:  "Firefox",
		5:  "Android",
		6:  "Opera",
		7:  "Blackberry",
		8:  "UC_Browser",
		9:  "Silk",
		10: "Nokia",
		11: "NetFront",
		12: "QQ",
		13: "Maxthon",
		14: "Sogou_Explorer",
		15: "Spotify",
		16: "Nintendo",
		17: "Samsung",
		18: "Yandex",
		19: "Coc_Coc",
		20: "Bot",
		21: "AppleBot",
		22: "BaiduBot",
		23: "BingBot",
		24: "DuckDuckGoBot",
		25: "FacebookBot",
		26: "GoogleBot",
		27: "LinkedInBot",
		28: "MsnBot",
		29: "PingdomBot",
		30: "TwitterBot",
		31: "YandexBot",
		32: "CocCocBot",
		33: "YahooBot",
	}
	BrowserType_value = map[string]int32{
		"unknown":           0,
		"Chrome":            1,
		"Internet_Explorer": 2,
		"Safari":            3,
		"Firefox":           4,
		"Android":           5,
		"Opera":             6,
		"Blackberry":        7,
		"UC_Browser":        8,
		"Silk":              9,
		"Nokia":             10,
		"NetFront":          11,
		"QQ":                12,
		"Maxthon":           13,
		"Sogou_Explorer":    14,
		"Spotify":           15,
		"Nintendo":          16,
		"Samsung":           17,
		"Yandex":            18,
		"Coc_Coc":           19,
		"Bot":               20,
		"AppleBot":          21,
		"BaiduBot":          22,
		"BingBot":           23,
		"DuckDuckGoBot":     24,
		"FacebookBot":       25,
		"GoogleBot":         26,
		"LinkedInBot":       27,
		"MsnBot":            28,
		"PingdomBot":        29,
		"TwitterBot":        30,
		"YandexBot":         31,
		"CocCocBot":         32,
		"YahooBot":          33,
	}
)

func (x BrowserType) Enum() *BrowserType {
	p := new(BrowserType)
	*p = x
	return p
}

func (x BrowserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrowserType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_browser_browser_proto_enumTypes[0].Descriptor()
}

func (BrowserType) Type() protoreflect.EnumType {
	return &file_types_browser_browser_proto_enumTypes[0]
}

func (x BrowserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrowserType.Descriptor instead.
func (BrowserType) EnumDescriptor() ([]byte, []int) {
	return file_types_browser_browser_proto_rawDescGZIP(), []int{0}
}

var File_types_browser_browser_proto protoreflect.FileDescriptor

var file_types_browser_browser_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2f,
	0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75,
	0x61, 0x73, 0x75, 0x72, 0x66, 0x65, 0x72, 0x2e, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x2a,
	0xef, 0x03, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x5f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x61, 0x66, 0x61, 0x72, 0x69, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x69, 0x72, 0x65, 0x66, 0x6f, 0x78, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x10, 0x06,
	0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x65, 0x72, 0x72, 0x79, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x43, 0x5f, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x10, 0x08,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x69, 0x6c, 0x6b, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f,
	0x6b, 0x69, 0x61, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x10, 0x0b, 0x12, 0x06, 0x0a, 0x02, 0x51, 0x51, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x4d,
	0x61, 0x78, 0x74, 0x68, 0x6f, 0x6e, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x6f, 0x67, 0x6f,
	0x75, 0x5f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x6f, 0x10, 0x10, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x61, 0x6d, 0x73, 0x75,
	0x6e, 0x67, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x12,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x63, 0x5f, 0x43, 0x6f, 0x63, 0x10, 0x13, 0x12, 0x07, 0x0a,
	0x03, 0x42, 0x6f, 0x74, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x42,
	0x6f, 0x74, 0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x61, 0x69, 0x64, 0x75, 0x42, 0x6f, 0x74,
	0x10, 0x16, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x74, 0x10, 0x17, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x75, 0x63, 0x6b, 0x44, 0x75, 0x63, 0x6b, 0x47, 0x6f, 0x42, 0x6f, 0x74,
	0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x6f,
	0x74, 0x10, 0x19, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42, 0x6f, 0x74,
	0x10, 0x1a, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x6e, 0x42, 0x6f,
	0x74, 0x10, 0x1b, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x73, 0x6e, 0x42, 0x6f, 0x74, 0x10, 0x1c, 0x12,
	0x0e, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x64, 0x6f, 0x6d, 0x42, 0x6f, 0x74, 0x10, 0x1d, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x74, 0x10, 0x1e, 0x12,
	0x0d, 0x0a, 0x09, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x6f, 0x74, 0x10, 0x1f, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x6f, 0x63, 0x43, 0x6f, 0x63, 0x42, 0x6f, 0x74, 0x10, 0x20, 0x12, 0x0c, 0x0a,
	0x08, 0x59, 0x61, 0x68, 0x6f, 0x6f, 0x42, 0x6f, 0x74, 0x10, 0x21, 0x22, 0x04, 0x08, 0x22, 0x10,
	0x3c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x69, 0x67, 0x65, 0x6c, 0x74, 0x69, 0x61, 0x6e, 0x79, 0x2f, 0x75, 0x61, 0x73, 0x75, 0x72,
	0x66, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_browser_browser_proto_rawDescOnce sync.Once
	file_types_browser_browser_proto_rawDescData = file_types_browser_browser_proto_rawDesc
)

func file_types_browser_browser_proto_rawDescGZIP() []byte {
	file_types_browser_browser_proto_rawDescOnce.Do(func() {
		file_types_browser_browser_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_browser_browser_proto_rawDescData)
	})
	return file_types_browser_browser_proto_rawDescData
}

var file_types_browser_browser_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_browser_browser_proto_goTypes = []interface{}{
	(BrowserType)(0), // 0: uasurfer.browser.BrowserType
}
var file_types_browser_browser_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_browser_browser_proto_init() }
func file_types_browser_browser_proto_init() {
	if File_types_browser_browser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_browser_browser_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_browser_browser_proto_goTypes,
		DependencyIndexes: file_types_browser_browser_proto_depIdxs,
		EnumInfos:         file_types_browser_browser_proto_enumTypes,
	}.Build()
	File_types_browser_browser_proto = out.File
	file_types_browser_browser_proto_rawDesc = nil
	file_types_browser_browser_proto_goTypes = nil
	file_types_browser_browser_proto_depIdxs = nil
}