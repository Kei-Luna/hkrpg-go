// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPlayerBoardDataScRsp.proto

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

type GetPlayerBoardDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockedHeadIconList []*HeadIconData   `protobuf:"bytes,12,rep,name=unlocked_head_icon_list,json=unlockedHeadIconList,proto3" json:"unlocked_head_icon_list,omitempty"`
	AssistAvatarIdList   []uint32          `protobuf:"varint,1,rep,packed,name=assist_avatar_id_list,json=assistAvatarIdList,proto3" json:"assist_avatar_id_list,omitempty"`
	Signature            string            `protobuf:"bytes,13,opt,name=signature,proto3" json:"signature,omitempty"`
	CurrentHeadIconId    uint32            `protobuf:"varint,14,opt,name=current_head_icon_id,json=currentHeadIconId,proto3" json:"current_head_icon_id,omitempty"`
	DisplayAvatarVec     *DisplayAvatarVec `protobuf:"bytes,4,opt,name=display_avatar_vec,json=displayAvatarVec,proto3" json:"display_avatar_vec,omitempty"`
	Retcode              uint32            `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetPlayerBoardDataScRsp) Reset() {
	*x = GetPlayerBoardDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPlayerBoardDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerBoardDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerBoardDataScRsp) ProtoMessage() {}

func (x *GetPlayerBoardDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPlayerBoardDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerBoardDataScRsp.ProtoReflect.Descriptor instead.
func (*GetPlayerBoardDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetPlayerBoardDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerBoardDataScRsp) GetUnlockedHeadIconList() []*HeadIconData {
	if x != nil {
		return x.UnlockedHeadIconList
	}
	return nil
}

func (x *GetPlayerBoardDataScRsp) GetAssistAvatarIdList() []uint32 {
	if x != nil {
		return x.AssistAvatarIdList
	}
	return nil
}

func (x *GetPlayerBoardDataScRsp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *GetPlayerBoardDataScRsp) GetCurrentHeadIconId() uint32 {
	if x != nil {
		return x.CurrentHeadIconId
	}
	return 0
}

func (x *GetPlayerBoardDataScRsp) GetDisplayAvatarVec() *DisplayAvatarVec {
	if x != nil {
		return x.DisplayAvatarVec
	}
	return nil
}

func (x *GetPlayerBoardDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetPlayerBoardDataScRsp_proto protoreflect.FileDescriptor

var file_GetPlayerBoardDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x56, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x17, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x49,
	0x63, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x15, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2f,
	0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x3f, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x76, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x56, 0x65, 0x63, 0x52, 0x10,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x56, 0x65, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GetPlayerBoardDataScRsp_proto_rawDescOnce sync.Once
	file_GetPlayerBoardDataScRsp_proto_rawDescData = file_GetPlayerBoardDataScRsp_proto_rawDesc
)

func file_GetPlayerBoardDataScRsp_proto_rawDescGZIP() []byte {
	file_GetPlayerBoardDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetPlayerBoardDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPlayerBoardDataScRsp_proto_rawDescData)
	})
	return file_GetPlayerBoardDataScRsp_proto_rawDescData
}

var file_GetPlayerBoardDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPlayerBoardDataScRsp_proto_goTypes = []interface{}{
	(*GetPlayerBoardDataScRsp)(nil), // 0: GetPlayerBoardDataScRsp
	(*HeadIconData)(nil),            // 1: HeadIconData
	(*DisplayAvatarVec)(nil),        // 2: DisplayAvatarVec
}
var file_GetPlayerBoardDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetPlayerBoardDataScRsp.unlocked_head_icon_list:type_name -> HeadIconData
	2, // 1: GetPlayerBoardDataScRsp.display_avatar_vec:type_name -> DisplayAvatarVec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetPlayerBoardDataScRsp_proto_init() }
func file_GetPlayerBoardDataScRsp_proto_init() {
	if File_GetPlayerBoardDataScRsp_proto != nil {
		return
	}
	file_HeadIconData_proto_init()
	file_DisplayAvatarVec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPlayerBoardDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerBoardDataScRsp); i {
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
			RawDescriptor: file_GetPlayerBoardDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPlayerBoardDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetPlayerBoardDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetPlayerBoardDataScRsp_proto_msgTypes,
	}.Build()
	File_GetPlayerBoardDataScRsp_proto = out.File
	file_GetPlayerBoardDataScRsp_proto_rawDesc = nil
	file_GetPlayerBoardDataScRsp_proto_goTypes = nil
	file_GetPlayerBoardDataScRsp_proto_depIdxs = nil
}
