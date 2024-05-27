// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SelectChatBubbleScRsp.proto

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

type SelectChatBubbleScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurChatBubble uint32 `protobuf:"varint,10,opt,name=cur_chat_bubble,json=curChatBubble,proto3" json:"cur_chat_bubble,omitempty"`
	Retcode       uint32 `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GCKNOBHOOOA   uint32 `protobuf:"varint,15,opt,name=GCKNOBHOOOA,proto3" json:"GCKNOBHOOOA,omitempty"`
}

func (x *SelectChatBubbleScRsp) Reset() {
	*x = SelectChatBubbleScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SelectChatBubbleScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectChatBubbleScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectChatBubbleScRsp) ProtoMessage() {}

func (x *SelectChatBubbleScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SelectChatBubbleScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectChatBubbleScRsp.ProtoReflect.Descriptor instead.
func (*SelectChatBubbleScRsp) Descriptor() ([]byte, []int) {
	return file_SelectChatBubbleScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SelectChatBubbleScRsp) GetCurChatBubble() uint32 {
	if x != nil {
		return x.CurChatBubble
	}
	return 0
}

func (x *SelectChatBubbleScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SelectChatBubbleScRsp) GetGCKNOBHOOOA() uint32 {
	if x != nil {
		return x.GCKNOBHOOOA
	}
	return 0
}

var File_SelectChatBubbleScRsp_proto protoreflect.FileDescriptor

var file_SelectChatBubbleScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x42, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a,
	0x15, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x74, 0x42, 0x75, 0x62, 0x62, 0x6c,
	0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x43, 0x68, 0x61, 0x74, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x43, 0x4b, 0x4e,
	0x4f, 0x42, 0x48, 0x4f, 0x4f, 0x4f, 0x41, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47,
	0x43, 0x4b, 0x4e, 0x4f, 0x42, 0x48, 0x4f, 0x4f, 0x4f, 0x41, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SelectChatBubbleScRsp_proto_rawDescOnce sync.Once
	file_SelectChatBubbleScRsp_proto_rawDescData = file_SelectChatBubbleScRsp_proto_rawDesc
)

func file_SelectChatBubbleScRsp_proto_rawDescGZIP() []byte {
	file_SelectChatBubbleScRsp_proto_rawDescOnce.Do(func() {
		file_SelectChatBubbleScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SelectChatBubbleScRsp_proto_rawDescData)
	})
	return file_SelectChatBubbleScRsp_proto_rawDescData
}

var file_SelectChatBubbleScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SelectChatBubbleScRsp_proto_goTypes = []interface{}{
	(*SelectChatBubbleScRsp)(nil), // 0: SelectChatBubbleScRsp
}
var file_SelectChatBubbleScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SelectChatBubbleScRsp_proto_init() }
func file_SelectChatBubbleScRsp_proto_init() {
	if File_SelectChatBubbleScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SelectChatBubbleScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectChatBubbleScRsp); i {
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
			RawDescriptor: file_SelectChatBubbleScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SelectChatBubbleScRsp_proto_goTypes,
		DependencyIndexes: file_SelectChatBubbleScRsp_proto_depIdxs,
		MessageInfos:      file_SelectChatBubbleScRsp_proto_msgTypes,
	}.Build()
	File_SelectChatBubbleScRsp_proto = out.File
	file_SelectChatBubbleScRsp_proto_rawDesc = nil
	file_SelectChatBubbleScRsp_proto_goTypes = nil
	file_SelectChatBubbleScRsp_proto_depIdxs = nil
}
