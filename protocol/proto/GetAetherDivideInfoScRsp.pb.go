// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAetherDivideInfoScRsp.proto

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

type GetAetherDivideInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IJEEHMIGNFI     uint32                    `protobuf:"varint,9,opt,name=IJEEHMIGNFI,proto3" json:"IJEEHMIGNFI,omitempty"`
	OMKNCCGDKNP     uint32                    `protobuf:"varint,2,opt,name=OMKNCCGDKNP,proto3" json:"OMKNCCGDKNP,omitempty"`
	AvatarList      []*AetherDivideSpiritInfo `protobuf:"bytes,8,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	LCNFCGHKACO     uint32                    `protobuf:"varint,1,opt,name=LCNFCGHKACO,proto3" json:"LCNFCGHKACO,omitempty"`
	LineupList      []*AetherDivideLineupInfo `protobuf:"bytes,7,rep,name=lineup_list,json=lineupList,proto3" json:"lineup_list,omitempty"`
	Retcode         uint32                    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NPKMPDHJHCL     uint32                    `protobuf:"varint,14,opt,name=NPKMPDHJHCL,proto3" json:"NPKMPDHJHCL,omitempty"`
	AetherSkillList []*AetherSkillInfo        `protobuf:"bytes,13,rep,name=aether_skill_list,json=aetherSkillList,proto3" json:"aether_skill_list,omitempty"`
	NEOAIFGNKNI     uint32                    `protobuf:"varint,6,opt,name=NEOAIFGNKNI,proto3" json:"NEOAIFGNKNI,omitempty"`
}

func (x *GetAetherDivideInfoScRsp) Reset() {
	*x = GetAetherDivideInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAetherDivideInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAetherDivideInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAetherDivideInfoScRsp) ProtoMessage() {}

func (x *GetAetherDivideInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAetherDivideInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAetherDivideInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetAetherDivideInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetAetherDivideInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAetherDivideInfoScRsp) GetIJEEHMIGNFI() uint32 {
	if x != nil {
		return x.IJEEHMIGNFI
	}
	return 0
}

func (x *GetAetherDivideInfoScRsp) GetOMKNCCGDKNP() uint32 {
	if x != nil {
		return x.OMKNCCGDKNP
	}
	return 0
}

func (x *GetAetherDivideInfoScRsp) GetAvatarList() []*AetherDivideSpiritInfo {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *GetAetherDivideInfoScRsp) GetLCNFCGHKACO() uint32 {
	if x != nil {
		return x.LCNFCGHKACO
	}
	return 0
}

func (x *GetAetherDivideInfoScRsp) GetLineupList() []*AetherDivideLineupInfo {
	if x != nil {
		return x.LineupList
	}
	return nil
}

func (x *GetAetherDivideInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAetherDivideInfoScRsp) GetNPKMPDHJHCL() uint32 {
	if x != nil {
		return x.NPKMPDHJHCL
	}
	return 0
}

func (x *GetAetherDivideInfoScRsp) GetAetherSkillList() []*AetherSkillInfo {
	if x != nil {
		return x.AetherSkillList
	}
	return nil
}

func (x *GetAetherDivideInfoScRsp) GetNEOAIFGNKNI() uint32 {
	if x != nil {
		return x.NEOAIFGNKNI
	}
	return 0
}

var File_GetAetherDivideInfoScRsp_proto protoreflect.FileDescriptor

var file_GetAetherDivideInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4a, 0x45, 0x45, 0x48, 0x4d, 0x49, 0x47, 0x4e, 0x46, 0x49, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4a, 0x45, 0x45, 0x48, 0x4d, 0x49, 0x47, 0x4e,
	0x46, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4d, 0x4b, 0x4e, 0x43, 0x43, 0x47, 0x44, 0x4b, 0x4e,
	0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4d, 0x4b, 0x4e, 0x43, 0x43, 0x47,
	0x44, 0x4b, 0x4e, 0x50, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x4c, 0x43, 0x4e, 0x46, 0x43, 0x47, 0x48, 0x4b, 0x41, 0x43, 0x4f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x43, 0x4e, 0x46, 0x43, 0x47, 0x48, 0x4b, 0x41, 0x43, 0x4f,
	0x12, 0x38, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x50, 0x4b, 0x4d, 0x50, 0x44, 0x48, 0x4a,
	0x48, 0x43, 0x4c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x50, 0x4b, 0x4d, 0x50,
	0x44, 0x48, 0x4a, 0x48, 0x43, 0x4c, 0x12, 0x3c, 0x0a, 0x11, 0x61, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x61, 0x65, 0x74, 0x68, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x45, 0x4f, 0x41, 0x49, 0x46, 0x47, 0x4e,
	0x4b, 0x4e, 0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x45, 0x4f, 0x41, 0x49,
	0x46, 0x47, 0x4e, 0x4b, 0x4e, 0x49, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAetherDivideInfoScRsp_proto_rawDescOnce sync.Once
	file_GetAetherDivideInfoScRsp_proto_rawDescData = file_GetAetherDivideInfoScRsp_proto_rawDesc
)

func file_GetAetherDivideInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetAetherDivideInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetAetherDivideInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAetherDivideInfoScRsp_proto_rawDescData)
	})
	return file_GetAetherDivideInfoScRsp_proto_rawDescData
}

var file_GetAetherDivideInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAetherDivideInfoScRsp_proto_goTypes = []interface{}{
	(*GetAetherDivideInfoScRsp)(nil), // 0: GetAetherDivideInfoScRsp
	(*AetherDivideSpiritInfo)(nil),   // 1: AetherDivideSpiritInfo
	(*AetherDivideLineupInfo)(nil),   // 2: AetherDivideLineupInfo
	(*AetherSkillInfo)(nil),          // 3: AetherSkillInfo
}
var file_GetAetherDivideInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetAetherDivideInfoScRsp.avatar_list:type_name -> AetherDivideSpiritInfo
	2, // 1: GetAetherDivideInfoScRsp.lineup_list:type_name -> AetherDivideLineupInfo
	3, // 2: GetAetherDivideInfoScRsp.aether_skill_list:type_name -> AetherSkillInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetAetherDivideInfoScRsp_proto_init() }
func file_GetAetherDivideInfoScRsp_proto_init() {
	if File_GetAetherDivideInfoScRsp_proto != nil {
		return
	}
	file_AetherDivideLineupInfo_proto_init()
	file_AetherDivideSpiritInfo_proto_init()
	file_AetherSkillInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetAetherDivideInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAetherDivideInfoScRsp); i {
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
			RawDescriptor: file_GetAetherDivideInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAetherDivideInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetAetherDivideInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetAetherDivideInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetAetherDivideInfoScRsp_proto = out.File
	file_GetAetherDivideInfoScRsp_proto_rawDesc = nil
	file_GetAetherDivideInfoScRsp_proto_goTypes = nil
	file_GetAetherDivideInfoScRsp_proto_depIdxs = nil
}
