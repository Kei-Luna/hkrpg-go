// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MessageGroupStatus.proto

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

type MessageGroupStatus int32

const (
	MessageGroupStatus_MESSAGE_GROUP_NONE   MessageGroupStatus = 0
	MessageGroupStatus_MESSAGE_GROUP_DOING  MessageGroupStatus = 1
	MessageGroupStatus_MESSAGE_GROUP_FINISH MessageGroupStatus = 2
	MessageGroupStatus_MESSAGE_GROUP_FROZEN MessageGroupStatus = 3
)

// Enum value maps for MessageGroupStatus.
var (
	MessageGroupStatus_name = map[int32]string{
		0: "MESSAGE_GROUP_NONE",
		1: "MESSAGE_GROUP_DOING",
		2: "MESSAGE_GROUP_FINISH",
		3: "MESSAGE_GROUP_FROZEN",
	}
	MessageGroupStatus_value = map[string]int32{
		"MESSAGE_GROUP_NONE":   0,
		"MESSAGE_GROUP_DOING":  1,
		"MESSAGE_GROUP_FINISH": 2,
		"MESSAGE_GROUP_FROZEN": 3,
	}
)

func (x MessageGroupStatus) Enum() *MessageGroupStatus {
	p := new(MessageGroupStatus)
	*p = x
	return p
}

func (x MessageGroupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageGroupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_MessageGroupStatus_proto_enumTypes[0].Descriptor()
}

func (MessageGroupStatus) Type() protoreflect.EnumType {
	return &file_MessageGroupStatus_proto_enumTypes[0]
}

func (x MessageGroupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageGroupStatus.Descriptor instead.
func (MessageGroupStatus) EnumDescriptor() ([]byte, []int) {
	return file_MessageGroupStatus_proto_rawDescGZIP(), []int{0}
}

var File_MessageGroupStatus_proto protoreflect.FileDescriptor

var file_MessageGroupStatus_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x79, 0x0a, 0x12, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x46, 0x52, 0x4f,
	0x5a, 0x45, 0x4e, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MessageGroupStatus_proto_rawDescOnce sync.Once
	file_MessageGroupStatus_proto_rawDescData = file_MessageGroupStatus_proto_rawDesc
)

func file_MessageGroupStatus_proto_rawDescGZIP() []byte {
	file_MessageGroupStatus_proto_rawDescOnce.Do(func() {
		file_MessageGroupStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_MessageGroupStatus_proto_rawDescData)
	})
	return file_MessageGroupStatus_proto_rawDescData
}

var file_MessageGroupStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MessageGroupStatus_proto_goTypes = []interface{}{
	(MessageGroupStatus)(0), // 0: MessageGroupStatus
}
var file_MessageGroupStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MessageGroupStatus_proto_init() }
func file_MessageGroupStatus_proto_init() {
	if File_MessageGroupStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MessageGroupStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MessageGroupStatus_proto_goTypes,
		DependencyIndexes: file_MessageGroupStatus_proto_depIdxs,
		EnumInfos:         file_MessageGroupStatus_proto_enumTypes,
	}.Build()
	File_MessageGroupStatus_proto = out.File
	file_MessageGroupStatus_proto_rawDesc = nil
	file_MessageGroupStatus_proto_goTypes = nil
	file_MessageGroupStatus_proto_depIdxs = nil
}
