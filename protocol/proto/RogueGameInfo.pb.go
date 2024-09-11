// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueGameInfo.proto

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

type RogueGameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueBuffInfo       *ChessRogueBuffInfo       `protobuf:"bytes,12,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	GameMiracleInfo     *ChessRogueMiracleInfo    `protobuf:"bytes,9,opt,name=game_miracle_info,json=gameMiracleInfo,proto3" json:"game_miracle_info,omitempty"`
	GameItemValue       *RogueGameItemValue       `protobuf:"bytes,2,opt,name=game_item_value,json=gameItemValue,proto3" json:"game_item_value,omitempty"`
	RogueAeonInfo       *ChessRogueGameAeonInfo   `protobuf:"bytes,1,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	RogueDifficultyInfo *RogueDifficultyLevelInfo `protobuf:"bytes,14,opt,name=rogue_difficulty_info,json=rogueDifficultyInfo,proto3" json:"rogue_difficulty_info,omitempty"`
	DPHEHCODKDE         *RogueTournFormulaInfo    `protobuf:"bytes,3,opt,name=DPHEHCODKDE,proto3" json:"DPHEHCODKDE,omitempty"`
	IPPILPLLBGK         *KeywordUnlockValue       `protobuf:"bytes,7,opt,name=IPPILPLLBGK,proto3" json:"IPPILPLLBGK,omitempty"`
	RogueLineupInfo     *RogueTournLineupInfo     `protobuf:"bytes,10,opt,name=rogue_lineup_info,json=rogueLineupInfo,proto3" json:"rogue_lineup_info,omitempty"`
}

func (x *RogueGameInfo) Reset() {
	*x = RogueGameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueGameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueGameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueGameInfo) ProtoMessage() {}

func (x *RogueGameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueGameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueGameInfo.ProtoReflect.Descriptor instead.
func (*RogueGameInfo) Descriptor() ([]byte, []int) {
	return file_RogueGameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueGameInfo) GetRogueBuffInfo() *ChessRogueBuffInfo {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *RogueGameInfo) GetGameMiracleInfo() *ChessRogueMiracleInfo {
	if x != nil {
		return x.GameMiracleInfo
	}
	return nil
}

func (x *RogueGameInfo) GetGameItemValue() *RogueGameItemValue {
	if x != nil {
		return x.GameItemValue
	}
	return nil
}

func (x *RogueGameInfo) GetRogueAeonInfo() *ChessRogueGameAeonInfo {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *RogueGameInfo) GetRogueDifficultyInfo() *RogueDifficultyLevelInfo {
	if x != nil {
		return x.RogueDifficultyInfo
	}
	return nil
}

func (x *RogueGameInfo) GetDPHEHCODKDE() *RogueTournFormulaInfo {
	if x != nil {
		return x.DPHEHCODKDE
	}
	return nil
}

func (x *RogueGameInfo) GetIPPILPLLBGK() *KeywordUnlockValue {
	if x != nil {
		return x.IPPILPLLBGK
	}
	return nil
}

func (x *RogueGameInfo) GetRogueLineupInfo() *RogueTournLineupInfo {
	if x != nil {
		return x.RogueLineupInfo
	}
	return nil
}

var File_RogueGameInfo_proto protoreflect.FileDescriptor

var file_RogueGameInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x04, 0x0a, 0x0d, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0f, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d,
	0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x15, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x13, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0b, 0x44, 0x50, 0x48, 0x45, 0x48,
	0x43, 0x4f, 0x44, 0x4b, 0x44, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x44, 0x50, 0x48, 0x45, 0x48, 0x43, 0x4f, 0x44, 0x4b, 0x44,
	0x45, 0x12, 0x35, 0x0a, 0x0b, 0x49, 0x50, 0x50, 0x49, 0x4c, 0x50, 0x4c, 0x4c, 0x42, 0x47, 0x4b,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x49, 0x50, 0x50,
	0x49, 0x4c, 0x50, 0x4c, 0x4c, 0x42, 0x47, 0x4b, 0x12, 0x41, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RogueGameInfo_proto_rawDescOnce sync.Once
	file_RogueGameInfo_proto_rawDescData = file_RogueGameInfo_proto_rawDesc
)

func file_RogueGameInfo_proto_rawDescGZIP() []byte {
	file_RogueGameInfo_proto_rawDescOnce.Do(func() {
		file_RogueGameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueGameInfo_proto_rawDescData)
	})
	return file_RogueGameInfo_proto_rawDescData
}

var file_RogueGameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueGameInfo_proto_goTypes = []interface{}{
	(*RogueGameInfo)(nil),            // 0: RogueGameInfo
	(*ChessRogueBuffInfo)(nil),       // 1: ChessRogueBuffInfo
	(*ChessRogueMiracleInfo)(nil),    // 2: ChessRogueMiracleInfo
	(*RogueGameItemValue)(nil),       // 3: RogueGameItemValue
	(*ChessRogueGameAeonInfo)(nil),   // 4: ChessRogueGameAeonInfo
	(*RogueDifficultyLevelInfo)(nil), // 5: RogueDifficultyLevelInfo
	(*RogueTournFormulaInfo)(nil),    // 6: RogueTournFormulaInfo
	(*KeywordUnlockValue)(nil),       // 7: KeywordUnlockValue
	(*RogueTournLineupInfo)(nil),     // 8: RogueTournLineupInfo
}
var file_RogueGameInfo_proto_depIdxs = []int32{
	1, // 0: RogueGameInfo.rogue_buff_info:type_name -> ChessRogueBuffInfo
	2, // 1: RogueGameInfo.game_miracle_info:type_name -> ChessRogueMiracleInfo
	3, // 2: RogueGameInfo.game_item_value:type_name -> RogueGameItemValue
	4, // 3: RogueGameInfo.rogue_aeon_info:type_name -> ChessRogueGameAeonInfo
	5, // 4: RogueGameInfo.rogue_difficulty_info:type_name -> RogueDifficultyLevelInfo
	6, // 5: RogueGameInfo.DPHEHCODKDE:type_name -> RogueTournFormulaInfo
	7, // 6: RogueGameInfo.IPPILPLLBGK:type_name -> KeywordUnlockValue
	8, // 7: RogueGameInfo.rogue_lineup_info:type_name -> RogueTournLineupInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RogueGameInfo_proto_init() }
func file_RogueGameInfo_proto_init() {
	if File_RogueGameInfo_proto != nil {
		return
	}
	file_ChessRogueGameAeonInfo_proto_init()
	file_ChessRogueBuffInfo_proto_init()
	file_RogueDifficultyLevelInfo_proto_init()
	file_RogueGameItemValue_proto_init()
	file_ChessRogueMiracleInfo_proto_init()
	file_KeywordUnlockValue_proto_init()
	file_RogueTournLineupInfo_proto_init()
	file_RogueTournFormulaInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueGameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueGameInfo); i {
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
			RawDescriptor: file_RogueGameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueGameInfo_proto_goTypes,
		DependencyIndexes: file_RogueGameInfo_proto_depIdxs,
		MessageInfos:      file_RogueGameInfo_proto_msgTypes,
	}.Build()
	File_RogueGameInfo_proto = out.File
	file_RogueGameInfo_proto_rawDesc = nil
	file_RogueGameInfo_proto_goTypes = nil
	file_RogueGameInfo_proto_depIdxs = nil
}
