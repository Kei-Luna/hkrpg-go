// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RaidStatus.proto

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

type RaidStatus int32

const (
	RaidStatus_RAID_STATUS_NONE   RaidStatus = 0
	RaidStatus_RAID_STATUS_DOING  RaidStatus = 1
	RaidStatus_RAID_STATUS_FINISH RaidStatus = 2
	RaidStatus_RAID_STATUS_FAILED RaidStatus = 3
)

// Enum value maps for RaidStatus.
var (
	RaidStatus_name = map[int32]string{
		0: "RAID_STATUS_NONE",
		1: "RAID_STATUS_DOING",
		2: "RAID_STATUS_FINISH",
		3: "RAID_STATUS_FAILED",
	}
	RaidStatus_value = map[string]int32{
		"RAID_STATUS_NONE":   0,
		"RAID_STATUS_DOING":  1,
		"RAID_STATUS_FINISH": 2,
		"RAID_STATUS_FAILED": 3,
	}
)

func (x RaidStatus) Enum() *RaidStatus {
	p := new(RaidStatus)
	*p = x
	return p
}

func (x RaidStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaidStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RaidStatus_proto_enumTypes[0].Descriptor()
}

func (RaidStatus) Type() protoreflect.EnumType {
	return &file_RaidStatus_proto_enumTypes[0]
}

func (x RaidStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RaidStatus.Descriptor instead.
func (RaidStatus) EnumDescriptor() ([]byte, []int) {
	return file_RaidStatus_proto_rawDescGZIP(), []int{0}
}

var File_RaidStatus_proto protoreflect.FileDescriptor

var file_RaidStatus_proto_rawDesc = []byte{
	0x0a, 0x10, 0x52, 0x61, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x69, 0x0a, 0x0a, 0x52, 0x61, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x10, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RaidStatus_proto_rawDescOnce sync.Once
	file_RaidStatus_proto_rawDescData = file_RaidStatus_proto_rawDesc
)

func file_RaidStatus_proto_rawDescGZIP() []byte {
	file_RaidStatus_proto_rawDescOnce.Do(func() {
		file_RaidStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_RaidStatus_proto_rawDescData)
	})
	return file_RaidStatus_proto_rawDescData
}

var file_RaidStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RaidStatus_proto_goTypes = []interface{}{
	(RaidStatus)(0), // 0: RaidStatus
}
var file_RaidStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RaidStatus_proto_init() }
func file_RaidStatus_proto_init() {
	if File_RaidStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RaidStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RaidStatus_proto_goTypes,
		DependencyIndexes: file_RaidStatus_proto_depIdxs,
		EnumInfos:         file_RaidStatus_proto_enumTypes,
	}.Build()
	File_RaidStatus_proto = out.File
	file_RaidStatus_proto_rawDesc = nil
	file_RaidStatus_proto_goTypes = nil
	file_RaidStatus_proto_depIdxs = nil
}