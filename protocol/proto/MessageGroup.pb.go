// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MessageGroup.proto

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

type MessageGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshTime        int64              `protobuf:"varint,4,opt,name=refresh_time,json=refreshTime,proto3" json:"refresh_time,omitempty"`
	MessageSectionId   uint32             `protobuf:"varint,13,opt,name=message_section_id,json=messageSectionId,proto3" json:"message_section_id,omitempty"`
	MessageSectionList []*MessageSection  `protobuf:"bytes,12,rep,name=message_section_list,json=messageSectionList,proto3" json:"message_section_list,omitempty"`
	Id                 uint32             `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	Status             MessageGroupStatus `protobuf:"varint,11,opt,name=status,proto3,enum=MessageGroupStatus" json:"status,omitempty"`
}

func (x *MessageGroup) Reset() {
	*x = MessageGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroup) ProtoMessage() {}

func (x *MessageGroup) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroup.ProtoReflect.Descriptor instead.
func (*MessageGroup) Descriptor() ([]byte, []int) {
	return file_MessageGroup_proto_rawDescGZIP(), []int{0}
}

func (x *MessageGroup) GetRefreshTime() int64 {
	if x != nil {
		return x.RefreshTime
	}
	return 0
}

func (x *MessageGroup) GetMessageSectionId() uint32 {
	if x != nil {
		return x.MessageSectionId
	}
	return 0
}

func (x *MessageGroup) GetMessageSectionList() []*MessageSection {
	if x != nil {
		return x.MessageSectionList
	}
	return nil
}

func (x *MessageGroup) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageGroup) GetStatus() MessageGroupStatus {
	if x != nil {
		return x.Status
	}
	return MessageGroupStatus_MESSAGE_GROUP_NONE
}

var File_MessageGroup_proto protoreflect.FileDescriptor

var file_MessageGroup_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MessageGroup_proto_rawDescOnce sync.Once
	file_MessageGroup_proto_rawDescData = file_MessageGroup_proto_rawDesc
)

func file_MessageGroup_proto_rawDescGZIP() []byte {
	file_MessageGroup_proto_rawDescOnce.Do(func() {
		file_MessageGroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_MessageGroup_proto_rawDescData)
	})
	return file_MessageGroup_proto_rawDescData
}

var file_MessageGroup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MessageGroup_proto_goTypes = []interface{}{
	(*MessageGroup)(nil),    // 0: MessageGroup
	(*MessageSection)(nil),  // 1: MessageSection
	(MessageGroupStatus)(0), // 2: MessageGroupStatus
}
var file_MessageGroup_proto_depIdxs = []int32{
	1, // 0: MessageGroup.message_section_list:type_name -> MessageSection
	2, // 1: MessageGroup.status:type_name -> MessageGroupStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MessageGroup_proto_init() }
func file_MessageGroup_proto_init() {
	if File_MessageGroup_proto != nil {
		return
	}
	file_MessageGroupStatus_proto_init()
	file_MessageSection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MessageGroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MessageGroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MessageGroup_proto_goTypes,
		DependencyIndexes: file_MessageGroup_proto_depIdxs,
		MessageInfos:      file_MessageGroup_proto_msgTypes,
	}.Build()
	File_MessageGroup_proto = out.File
	file_MessageGroup_proto_rawDesc = nil
	file_MessageGroup_proto_goTypes = nil
	file_MessageGroup_proto_depIdxs = nil
}
