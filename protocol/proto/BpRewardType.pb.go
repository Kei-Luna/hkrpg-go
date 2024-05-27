// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BpRewardType.proto

package proto

import (
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

type BpRewardType int32

const (
	BpRewardType_BP_REWARAD_TYPE_NONE             BpRewardType = 0
	BpRewardType_BP_REWARAD_TYPE_FREE             BpRewardType = 1
	BpRewardType_BP_REWARAD_TYPE_PREMIUM_1        BpRewardType = 2
	BpRewardType_BP_REWARAD_TYPE_PREMIUM_2        BpRewardType = 3
	BpRewardType_BP_REWARAD_TYPE_PREMIUM_OPTIONAL BpRewardType = 4
)

// Enum value maps for BpRewardType.
var (
	BpRewardType_name = map[int32]string{
		0: "BP_REWARAD_TYPE_NONE",
		1: "BP_REWARAD_TYPE_FREE",
		2: "BP_REWARAD_TYPE_PREMIUM_1",
		3: "BP_REWARAD_TYPE_PREMIUM_2",
		4: "BP_REWARAD_TYPE_PREMIUM_OPTIONAL",
	}
	BpRewardType_value = map[string]int32{
		"BP_REWARAD_TYPE_NONE":             0,
		"BP_REWARAD_TYPE_FREE":             1,
		"BP_REWARAD_TYPE_PREMIUM_1":        2,
		"BP_REWARAD_TYPE_PREMIUM_2":        3,
		"BP_REWARAD_TYPE_PREMIUM_OPTIONAL": 4,
	}
)

func (x BpRewardType) Enum() *BpRewardType {
	p := new(BpRewardType)
	*p = x
	return p
}

func (x BpRewardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BpRewardType) Descriptor() protoreflect.EnumDescriptor {
	return file_BpRewardType_proto_enumTypes[0].Descriptor()
}

func (BpRewardType) Type() protoreflect.EnumType {
	return &file_BpRewardType_proto_enumTypes[0]
}

func (x BpRewardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BpRewardType.Descriptor instead.
func (BpRewardType) EnumDescriptor() ([]byte, []int) {
	return file_BpRewardType_proto_rawDescGZIP(), []int{0}
}

var File_BpRewardType_proto protoreflect.FileDescriptor

var file_BpRewardType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x42, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa6, 0x01, 0x0a, 0x0c, 0x42, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x50, 0x5f, 0x52, 0x45, 0x57, 0x41,
	0x52, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x42, 0x50, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x41, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x50, 0x5f,
	0x52, 0x45, 0x57, 0x41, 0x52, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45,
	0x4d, 0x49, 0x55, 0x4d, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x50, 0x5f, 0x52,
	0x45, 0x57, 0x41, 0x52, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x4d,
	0x49, 0x55, 0x4d, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x42, 0x50, 0x5f, 0x52, 0x45,
	0x57, 0x41, 0x52, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x4d, 0x49,
	0x55, 0x4d, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_BpRewardType_proto_rawDescOnce sync.Once
	file_BpRewardType_proto_rawDescData = file_BpRewardType_proto_rawDesc
)

func file_BpRewardType_proto_rawDescGZIP() []byte {
	file_BpRewardType_proto_rawDescOnce.Do(func() {
		file_BpRewardType_proto_rawDescData = protoimpl.X.CompressGZIP(file_BpRewardType_proto_rawDescData)
	})
	return file_BpRewardType_proto_rawDescData
}

var file_BpRewardType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BpRewardType_proto_goTypes = []interface{}{
	(BpRewardType)(0), // 0: BpRewardType
}
var file_BpRewardType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BpRewardType_proto_init() }
func file_BpRewardType_proto_init() {
	if File_BpRewardType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BpRewardType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BpRewardType_proto_goTypes,
		DependencyIndexes: file_BpRewardType_proto_depIdxs,
		EnumInfos:         file_BpRewardType_proto_enumTypes,
	}.Build()
	File_BpRewardType_proto = out.File
	file_BpRewardType_proto_rawDesc = nil
	file_BpRewardType_proto_goTypes = nil
	file_BpRewardType_proto_depIdxs = nil
}
