// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishAeonDialogueGroupScRsp.proto

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

type FinishAeonDialogueGroupScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueAeonInfo *BOILBLGPBKN `protobuf:"bytes,8,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	Reward        *ItemList    `protobuf:"bytes,9,opt,name=reward,proto3" json:"reward,omitempty"`
	Retcode       uint32       `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *FinishAeonDialogueGroupScRsp) Reset() {
	*x = FinishAeonDialogueGroupScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishAeonDialogueGroupScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishAeonDialogueGroupScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishAeonDialogueGroupScRsp) ProtoMessage() {}

func (x *FinishAeonDialogueGroupScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_FinishAeonDialogueGroupScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishAeonDialogueGroupScRsp.ProtoReflect.Descriptor instead.
func (*FinishAeonDialogueGroupScRsp) Descriptor() ([]byte, []int) {
	return file_FinishAeonDialogueGroupScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FinishAeonDialogueGroupScRsp) GetRogueAeonInfo() *BOILBLGPBKN {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *FinishAeonDialogueGroupScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *FinishAeonDialogueGroupScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_FinishAeonDialogueGroupScRsp_proto protoreflect.FileDescriptor

var file_FinishAeonDialogueGroupScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x41, 0x65, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x4f, 0x49, 0x4c, 0x42, 0x4c, 0x47, 0x50, 0x42, 0x4b,
	0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x41, 0x65, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x4f, 0x49, 0x4c, 0x42, 0x4c, 0x47, 0x50, 0x42, 0x4b, 0x4e, 0x52,
	0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21,
	0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FinishAeonDialogueGroupScRsp_proto_rawDescOnce sync.Once
	file_FinishAeonDialogueGroupScRsp_proto_rawDescData = file_FinishAeonDialogueGroupScRsp_proto_rawDesc
)

func file_FinishAeonDialogueGroupScRsp_proto_rawDescGZIP() []byte {
	file_FinishAeonDialogueGroupScRsp_proto_rawDescOnce.Do(func() {
		file_FinishAeonDialogueGroupScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishAeonDialogueGroupScRsp_proto_rawDescData)
	})
	return file_FinishAeonDialogueGroupScRsp_proto_rawDescData
}

var file_FinishAeonDialogueGroupScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishAeonDialogueGroupScRsp_proto_goTypes = []interface{}{
	(*FinishAeonDialogueGroupScRsp)(nil), // 0: FinishAeonDialogueGroupScRsp
	(*BOILBLGPBKN)(nil),                  // 1: BOILBLGPBKN
	(*ItemList)(nil),                     // 2: ItemList
}
var file_FinishAeonDialogueGroupScRsp_proto_depIdxs = []int32{
	1, // 0: FinishAeonDialogueGroupScRsp.rogue_aeon_info:type_name -> BOILBLGPBKN
	2, // 1: FinishAeonDialogueGroupScRsp.reward:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FinishAeonDialogueGroupScRsp_proto_init() }
func file_FinishAeonDialogueGroupScRsp_proto_init() {
	if File_FinishAeonDialogueGroupScRsp_proto != nil {
		return
	}
	file_BOILBLGPBKN_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FinishAeonDialogueGroupScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishAeonDialogueGroupScRsp); i {
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
			RawDescriptor: file_FinishAeonDialogueGroupScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishAeonDialogueGroupScRsp_proto_goTypes,
		DependencyIndexes: file_FinishAeonDialogueGroupScRsp_proto_depIdxs,
		MessageInfos:      file_FinishAeonDialogueGroupScRsp_proto_msgTypes,
	}.Build()
	File_FinishAeonDialogueGroupScRsp_proto = out.File
	file_FinishAeonDialogueGroupScRsp_proto_rawDesc = nil
	file_FinishAeonDialogueGroupScRsp_proto_goTypes = nil
	file_FinishAeonDialogueGroupScRsp_proto_depIdxs = nil
}
