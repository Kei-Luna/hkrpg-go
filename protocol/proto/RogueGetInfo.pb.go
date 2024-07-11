// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueGetInfo.proto

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

type RogueGetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueSeasonInfo      *RogueSeasonInfo         `protobuf:"bytes,7,opt,name=rogue_season_info,json=rogueSeasonInfo,proto3" json:"rogue_season_info,omitempty"`
	RogueScoreRewardInfo *RogueScoreRewardInfo    `protobuf:"bytes,9,opt,name=rogue_score_reward_info,json=rogueScoreRewardInfo,proto3" json:"rogue_score_reward_info,omitempty"`
	RogueVirtualItemInfo *RogueGetVirtualItemInfo `protobuf:"bytes,10,opt,name=rogue_virtual_item_info,json=rogueVirtualItemInfo,proto3" json:"rogue_virtual_item_info,omitempty"`
	RogueAreaInfo        *RogueAreaInfo           `protobuf:"bytes,5,opt,name=rogue_area_info,json=rogueAreaInfo,proto3" json:"rogue_area_info,omitempty"`
	RogueAeonInfo        *RogueAeonInfo           `protobuf:"bytes,11,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
}

func (x *RogueGetInfo) Reset() {
	*x = RogueGetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueGetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueGetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueGetInfo) ProtoMessage() {}

func (x *RogueGetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueGetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueGetInfo.ProtoReflect.Descriptor instead.
func (*RogueGetInfo) Descriptor() ([]byte, []int) {
	return file_RogueGetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueGetInfo) GetRogueSeasonInfo() *RogueSeasonInfo {
	if x != nil {
		return x.RogueSeasonInfo
	}
	return nil
}

func (x *RogueGetInfo) GetRogueScoreRewardInfo() *RogueScoreRewardInfo {
	if x != nil {
		return x.RogueScoreRewardInfo
	}
	return nil
}

func (x *RogueGetInfo) GetRogueVirtualItemInfo() *RogueGetVirtualItemInfo {
	if x != nil {
		return x.RogueVirtualItemInfo
	}
	return nil
}

func (x *RogueGetInfo) GetRogueAreaInfo() *RogueAreaInfo {
	if x != nil {
		return x.RogueAreaInfo
	}
	return nil
}

func (x *RogueGetInfo) GetRogueAeonInfo() *RogueAeonInfo {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

var File_RogueGetInfo_proto protoreflect.FileDescriptor

var file_RogueGetInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdb, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x4c, 0x0a, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4f, 0x0a,
	0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36,
	0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72,
	0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f,
	0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueGetInfo_proto_rawDescOnce sync.Once
	file_RogueGetInfo_proto_rawDescData = file_RogueGetInfo_proto_rawDesc
)

func file_RogueGetInfo_proto_rawDescGZIP() []byte {
	file_RogueGetInfo_proto_rawDescOnce.Do(func() {
		file_RogueGetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueGetInfo_proto_rawDescData)
	})
	return file_RogueGetInfo_proto_rawDescData
}

var file_RogueGetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueGetInfo_proto_goTypes = []interface{}{
	(*RogueGetInfo)(nil),            // 0: RogueGetInfo
	(*RogueSeasonInfo)(nil),         // 1: RogueSeasonInfo
	(*RogueScoreRewardInfo)(nil),    // 2: RogueScoreRewardInfo
	(*RogueGetVirtualItemInfo)(nil), // 3: RogueGetVirtualItemInfo
	(*RogueAreaInfo)(nil),           // 4: RogueAreaInfo
	(*RogueAeonInfo)(nil),           // 5: RogueAeonInfo
}
var file_RogueGetInfo_proto_depIdxs = []int32{
	1, // 0: RogueGetInfo.rogue_season_info:type_name -> RogueSeasonInfo
	2, // 1: RogueGetInfo.rogue_score_reward_info:type_name -> RogueScoreRewardInfo
	3, // 2: RogueGetInfo.rogue_virtual_item_info:type_name -> RogueGetVirtualItemInfo
	4, // 3: RogueGetInfo.rogue_area_info:type_name -> RogueAreaInfo
	5, // 4: RogueGetInfo.rogue_aeon_info:type_name -> RogueAeonInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_RogueGetInfo_proto_init() }
func file_RogueGetInfo_proto_init() {
	if File_RogueGetInfo_proto != nil {
		return
	}
	file_RogueAreaInfo_proto_init()
	file_RogueGetVirtualItemInfo_proto_init()
	file_RogueScoreRewardInfo_proto_init()
	file_RogueSeasonInfo_proto_init()
	file_RogueAeonInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueGetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueGetInfo); i {
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
			RawDescriptor: file_RogueGetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueGetInfo_proto_goTypes,
		DependencyIndexes: file_RogueGetInfo_proto_depIdxs,
		MessageInfos:      file_RogueGetInfo_proto_msgTypes,
	}.Build()
	File_RogueGetInfo_proto = out.File
	file_RogueGetInfo_proto_rawDesc = nil
	file_RogueGetInfo_proto_goTypes = nil
	file_RogueGetInfo_proto_depIdxs = nil
}
