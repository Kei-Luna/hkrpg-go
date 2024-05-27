// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordMonsterInfoNotifyReason.proto

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

type PunkLordMonsterInfoNotifyReason int32

const (
	PunkLordMonsterInfoNotifyReason_PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_NONE       PunkLordMonsterInfoNotifyReason = 0
	PunkLordMonsterInfoNotifyReason_PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_ENTER_RAID PunkLordMonsterInfoNotifyReason = 1
	PunkLordMonsterInfoNotifyReason_PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_BATTLE_END PunkLordMonsterInfoNotifyReason = 2
	PunkLordMonsterInfoNotifyReason_PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_LEAVE_RAID PunkLordMonsterInfoNotifyReason = 3
)

// Enum value maps for PunkLordMonsterInfoNotifyReason.
var (
	PunkLordMonsterInfoNotifyReason_name = map[int32]string{
		0: "PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_NONE",
		1: "PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_ENTER_RAID",
		2: "PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_BATTLE_END",
		3: "PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_LEAVE_RAID",
	}
	PunkLordMonsterInfoNotifyReason_value = map[string]int32{
		"PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_NONE":       0,
		"PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_ENTER_RAID": 1,
		"PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_BATTLE_END": 2,
		"PUNK_LORD_MONSTER_INFO_NOTIFY_REASON_LEAVE_RAID": 3,
	}
)

func (x PunkLordMonsterInfoNotifyReason) Enum() *PunkLordMonsterInfoNotifyReason {
	p := new(PunkLordMonsterInfoNotifyReason)
	*p = x
	return p
}

func (x PunkLordMonsterInfoNotifyReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PunkLordMonsterInfoNotifyReason) Descriptor() protoreflect.EnumDescriptor {
	return file_PunkLordMonsterInfoNotifyReason_proto_enumTypes[0].Descriptor()
}

func (PunkLordMonsterInfoNotifyReason) Type() protoreflect.EnumType {
	return &file_PunkLordMonsterInfoNotifyReason_proto_enumTypes[0]
}

func (x PunkLordMonsterInfoNotifyReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PunkLordMonsterInfoNotifyReason.Descriptor instead.
func (PunkLordMonsterInfoNotifyReason) EnumDescriptor() ([]byte, []int) {
	return file_PunkLordMonsterInfoNotifyReason_proto_rawDescGZIP(), []int{0}
}

var File_PunkLordMonsterInfoNotifyReason_proto protoreflect.FileDescriptor

var file_PunkLordMonsterInfoNotifyReason_proto_rawDesc = []byte{
	0x0a, 0x25, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xef, 0x01, 0x0a, 0x1f, 0x50, 0x75, 0x6e, 0x6b,
	0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x29, 0x50,
	0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x33, 0x0a, 0x2f, 0x50, 0x55,
	0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x52, 0x41, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x33, 0x0a, 0x2f, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x4f, 0x4e,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x45,
	0x4e, 0x44, 0x10, 0x02, 0x12, 0x33, 0x0a, 0x2f, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52,
	0x44, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e,
	0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x41,
	0x56, 0x45, 0x5f, 0x52, 0x41, 0x49, 0x44, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PunkLordMonsterInfoNotifyReason_proto_rawDescOnce sync.Once
	file_PunkLordMonsterInfoNotifyReason_proto_rawDescData = file_PunkLordMonsterInfoNotifyReason_proto_rawDesc
)

func file_PunkLordMonsterInfoNotifyReason_proto_rawDescGZIP() []byte {
	file_PunkLordMonsterInfoNotifyReason_proto_rawDescOnce.Do(func() {
		file_PunkLordMonsterInfoNotifyReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordMonsterInfoNotifyReason_proto_rawDescData)
	})
	return file_PunkLordMonsterInfoNotifyReason_proto_rawDescData
}

var file_PunkLordMonsterInfoNotifyReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PunkLordMonsterInfoNotifyReason_proto_goTypes = []interface{}{
	(PunkLordMonsterInfoNotifyReason)(0), // 0: PunkLordMonsterInfoNotifyReason
}
var file_PunkLordMonsterInfoNotifyReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PunkLordMonsterInfoNotifyReason_proto_init() }
func file_PunkLordMonsterInfoNotifyReason_proto_init() {
	if File_PunkLordMonsterInfoNotifyReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PunkLordMonsterInfoNotifyReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordMonsterInfoNotifyReason_proto_goTypes,
		DependencyIndexes: file_PunkLordMonsterInfoNotifyReason_proto_depIdxs,
		EnumInfos:         file_PunkLordMonsterInfoNotifyReason_proto_enumTypes,
	}.Build()
	File_PunkLordMonsterInfoNotifyReason_proto = out.File
	file_PunkLordMonsterInfoNotifyReason_proto_rawDesc = nil
	file_PunkLordMonsterInfoNotifyReason_proto_goTypes = nil
	file_PunkLordMonsterInfoNotifyReason_proto_depIdxs = nil
}
