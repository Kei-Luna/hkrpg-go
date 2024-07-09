// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonDialogueInfo.proto

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

type RogueCommonDialogueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogueBasicInfo *RogueCommonDialogueBasicInfo `protobuf:"bytes,12,opt,name=dialogue_basic_info,json=dialogueBasicInfo,proto3" json:"dialogue_basic_info,omitempty"`
	GGGJACPNJNP       *EBMPINJCHFB                  `protobuf:"bytes,2,opt,name=GGGJACPNJNP,proto3" json:"GGGJACPNJNP,omitempty"`
	IPKNIJNEFIJ       *NDCKDMOMLFK                  `protobuf:"bytes,13,opt,name=IPKNIJNEFIJ,proto3" json:"IPKNIJNEFIJ,omitempty"`
	LCKCBBKMDNI       *ABCMHGNGFGL                  `protobuf:"bytes,3,opt,name=LCKCBBKMDNI,proto3" json:"LCKCBBKMDNI,omitempty"`
}

func (x *RogueCommonDialogueInfo) Reset() {
	*x = RogueCommonDialogueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCommonDialogueInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCommonDialogueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCommonDialogueInfo) ProtoMessage() {}

func (x *RogueCommonDialogueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCommonDialogueInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCommonDialogueInfo.ProtoReflect.Descriptor instead.
func (*RogueCommonDialogueInfo) Descriptor() ([]byte, []int) {
	return file_RogueCommonDialogueInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueCommonDialogueInfo) GetDialogueBasicInfo() *RogueCommonDialogueBasicInfo {
	if x != nil {
		return x.DialogueBasicInfo
	}
	return nil
}

func (x *RogueCommonDialogueInfo) GetGGGJACPNJNP() *EBMPINJCHFB {
	if x != nil {
		return x.GGGJACPNJNP
	}
	return nil
}

func (x *RogueCommonDialogueInfo) GetIPKNIJNEFIJ() *NDCKDMOMLFK {
	if x != nil {
		return x.IPKNIJNEFIJ
	}
	return nil
}

func (x *RogueCommonDialogueInfo) GetLCKCBBKMDNI() *ABCMHGNGFGL {
	if x != nil {
		return x.LCKCBBKMDNI
	}
	return nil
}

var File_RogueCommonDialogueInfo_proto protoreflect.FileDescriptor

var file_RogueCommonDialogueInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x42, 0x4d, 0x50, 0x49, 0x4e, 0x4a, 0x43, 0x48, 0x46, 0x42,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x43, 0x4d, 0x48, 0x47, 0x4e, 0x47,
	0x46, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x44, 0x43, 0x4b, 0x44,
	0x4d, 0x4f, 0x4d, 0x4c, 0x46, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a,
	0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x13, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x47, 0x47, 0x4a, 0x41,
	0x43, 0x50, 0x4e, 0x4a, 0x4e, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x42, 0x4d, 0x50, 0x49, 0x4e, 0x4a, 0x43, 0x48, 0x46, 0x42, 0x52, 0x0b, 0x47, 0x47, 0x47, 0x4a,
	0x41, 0x43, 0x50, 0x4e, 0x4a, 0x4e, 0x50, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x50, 0x4b, 0x4e, 0x49,
	0x4a, 0x4e, 0x45, 0x46, 0x49, 0x4a, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e,
	0x44, 0x43, 0x4b, 0x44, 0x4d, 0x4f, 0x4d, 0x4c, 0x46, 0x4b, 0x52, 0x0b, 0x49, 0x50, 0x4b, 0x4e,
	0x49, 0x4a, 0x4e, 0x45, 0x46, 0x49, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x43, 0x4b, 0x43, 0x42,
	0x42, 0x4b, 0x4d, 0x44, 0x4e, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41,
	0x42, 0x43, 0x4d, 0x48, 0x47, 0x4e, 0x47, 0x46, 0x47, 0x4c, 0x52, 0x0b, 0x4c, 0x43, 0x4b, 0x43,
	0x42, 0x42, 0x4b, 0x4d, 0x44, 0x4e, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonDialogueInfo_proto_rawDescOnce sync.Once
	file_RogueCommonDialogueInfo_proto_rawDescData = file_RogueCommonDialogueInfo_proto_rawDesc
)

func file_RogueCommonDialogueInfo_proto_rawDescGZIP() []byte {
	file_RogueCommonDialogueInfo_proto_rawDescOnce.Do(func() {
		file_RogueCommonDialogueInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonDialogueInfo_proto_rawDescData)
	})
	return file_RogueCommonDialogueInfo_proto_rawDescData
}

var file_RogueCommonDialogueInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCommonDialogueInfo_proto_goTypes = []interface{}{
	(*RogueCommonDialogueInfo)(nil),      // 0: RogueCommonDialogueInfo
	(*RogueCommonDialogueBasicInfo)(nil), // 1: RogueCommonDialogueBasicInfo
	(*EBMPINJCHFB)(nil),                  // 2: EBMPINJCHFB
	(*NDCKDMOMLFK)(nil),                  // 3: NDCKDMOMLFK
	(*ABCMHGNGFGL)(nil),                  // 4: ABCMHGNGFGL
}
var file_RogueCommonDialogueInfo_proto_depIdxs = []int32{
	1, // 0: RogueCommonDialogueInfo.dialogue_basic_info:type_name -> RogueCommonDialogueBasicInfo
	2, // 1: RogueCommonDialogueInfo.GGGJACPNJNP:type_name -> EBMPINJCHFB
	3, // 2: RogueCommonDialogueInfo.IPKNIJNEFIJ:type_name -> NDCKDMOMLFK
	4, // 3: RogueCommonDialogueInfo.LCKCBBKMDNI:type_name -> ABCMHGNGFGL
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueCommonDialogueInfo_proto_init() }
func file_RogueCommonDialogueInfo_proto_init() {
	if File_RogueCommonDialogueInfo_proto != nil {
		return
	}
	file_RogueCommonDialogueBasicInfo_proto_init()
	file_EBMPINJCHFB_proto_init()
	file_ABCMHGNGFGL_proto_init()
	file_NDCKDMOMLFK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCommonDialogueInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCommonDialogueInfo); i {
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
			RawDescriptor: file_RogueCommonDialogueInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonDialogueInfo_proto_goTypes,
		DependencyIndexes: file_RogueCommonDialogueInfo_proto_depIdxs,
		MessageInfos:      file_RogueCommonDialogueInfo_proto_msgTypes,
	}.Build()
	File_RogueCommonDialogueInfo_proto = out.File
	file_RogueCommonDialogueInfo_proto_rawDesc = nil
	file_RogueCommonDialogueInfo_proto_goTypes = nil
	file_RogueCommonDialogueInfo_proto_depIdxs = nil
}
