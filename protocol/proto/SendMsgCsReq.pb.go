// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SendMsgCsReq.proto

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

type SendMsgCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetList  []uint32 `protobuf:"varint,4,rep,packed,name=target_list,json=targetList,proto3" json:"target_list,omitempty"`
	ChatType    ChatType `protobuf:"varint,8,opt,name=chat_type,json=chatType,proto3,enum=ChatType" json:"chat_type,omitempty"`
	MessageType MsgType  `protobuf:"varint,9,opt,name=message_type,json=messageType,proto3,enum=MsgType" json:"message_type,omitempty"`
	ExtraId     uint32   `protobuf:"varint,7,opt,name=extra_id,json=extraId,proto3" json:"extra_id,omitempty"`
	MessageText string   `protobuf:"bytes,2,opt,name=message_text,json=messageText,proto3" json:"message_text,omitempty"`
}

func (x *SendMsgCsReq) Reset() {
	*x = SendMsgCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SendMsgCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgCsReq) ProtoMessage() {}

func (x *SendMsgCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SendMsgCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgCsReq.ProtoReflect.Descriptor instead.
func (*SendMsgCsReq) Descriptor() ([]byte, []int) {
	return file_SendMsgCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SendMsgCsReq) GetTargetList() []uint32 {
	if x != nil {
		return x.TargetList
	}
	return nil
}

func (x *SendMsgCsReq) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_TYPE_NONE
}

func (x *SendMsgCsReq) GetMessageType() MsgType {
	if x != nil {
		return x.MessageType
	}
	return MsgType_MSG_TYPE_NONE
}

func (x *SendMsgCsReq) GetExtraId() uint32 {
	if x != nil {
		return x.ExtraId
	}
	return 0
}

func (x *SendMsgCsReq) GetMessageText() string {
	if x != nil {
		return x.MessageText
	}
	return ""
}

var File_SendMsgCsReq_proto protoreflect.FileDescriptor

var file_SendMsgCsReq_proto_rawDesc = []byte{
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x78, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SendMsgCsReq_proto_rawDescOnce sync.Once
	file_SendMsgCsReq_proto_rawDescData = file_SendMsgCsReq_proto_rawDesc
)

func file_SendMsgCsReq_proto_rawDescGZIP() []byte {
	file_SendMsgCsReq_proto_rawDescOnce.Do(func() {
		file_SendMsgCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SendMsgCsReq_proto_rawDescData)
	})
	return file_SendMsgCsReq_proto_rawDescData
}

var file_SendMsgCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SendMsgCsReq_proto_goTypes = []interface{}{
	(*SendMsgCsReq)(nil), // 0: SendMsgCsReq
	(ChatType)(0),        // 1: ChatType
	(MsgType)(0),         // 2: MsgType
}
var file_SendMsgCsReq_proto_depIdxs = []int32{
	1, // 0: SendMsgCsReq.chat_type:type_name -> ChatType
	2, // 1: SendMsgCsReq.message_type:type_name -> MsgType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SendMsgCsReq_proto_init() }
func file_SendMsgCsReq_proto_init() {
	if File_SendMsgCsReq_proto != nil {
		return
	}
	file_ChatType_proto_init()
	file_MsgType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SendMsgCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgCsReq); i {
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
			RawDescriptor: file_SendMsgCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SendMsgCsReq_proto_goTypes,
		DependencyIndexes: file_SendMsgCsReq_proto_depIdxs,
		MessageInfos:      file_SendMsgCsReq_proto_msgTypes,
	}.Build()
	File_SendMsgCsReq_proto = out.File
	file_SendMsgCsReq_proto_rawDesc = nil
	file_SendMsgCsReq_proto_goTypes = nil
	file_SendMsgCsReq_proto_depIdxs = nil
}
