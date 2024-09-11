// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueStartScRsp.proto

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

type ChessRogueStartScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info             *ChessRogueGameInfo              `protobuf:"bytes,8,opt,name=info,proto3" json:"info,omitempty"`
	StageInfo        *ChessRogueInfo                  `protobuf:"bytes,12,opt,name=stage_info,json=stageInfo,proto3" json:"stage_info,omitempty"`
	Retcode          uint32                           `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Scene            *SceneInfo                       `protobuf:"bytes,7,opt,name=scene,proto3" json:"scene,omitempty"`
	Lineup           *LineupInfo                      `protobuf:"bytes,9,opt,name=lineup,proto3" json:"lineup,omitempty"`
	BoardEventInfo   *ChessRogueLayerInitialEventInfo `protobuf:"bytes,14,opt,name=board_event_info,json=boardEventInfo,proto3" json:"board_event_info,omitempty"`
	RogueCurrentInfo *ChessRogueCurrentInfo           `protobuf:"bytes,4,opt,name=rogue_current_info,json=rogueCurrentInfo,proto3" json:"rogue_current_info,omitempty"`
}

func (x *ChessRogueStartScRsp) Reset() {
	*x = ChessRogueStartScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueStartScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueStartScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueStartScRsp) ProtoMessage() {}

func (x *ChessRogueStartScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueStartScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueStartScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueStartScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueStartScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueStartScRsp) GetInfo() *ChessRogueGameInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChessRogueStartScRsp) GetStageInfo() *ChessRogueInfo {
	if x != nil {
		return x.StageInfo
	}
	return nil
}

func (x *ChessRogueStartScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChessRogueStartScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *ChessRogueStartScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *ChessRogueStartScRsp) GetBoardEventInfo() *ChessRogueLayerInitialEventInfo {
	if x != nil {
		return x.BoardEventInfo
	}
	return nil
}

func (x *ChessRogueStartScRsp) GetRogueCurrentInfo() *ChessRogueCurrentInfo {
	if x != nil {
		return x.RogueCurrentInfo
	}
	return nil
}

var File_ChessRogueStartScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueStartScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x10, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x44, 0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueStartScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueStartScRsp_proto_rawDescData = file_ChessRogueStartScRsp_proto_rawDesc
)

func file_ChessRogueStartScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueStartScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueStartScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueStartScRsp_proto_rawDescData)
	})
	return file_ChessRogueStartScRsp_proto_rawDescData
}

var file_ChessRogueStartScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueStartScRsp_proto_goTypes = []interface{}{
	(*ChessRogueStartScRsp)(nil),            // 0: ChessRogueStartScRsp
	(*ChessRogueGameInfo)(nil),              // 1: ChessRogueGameInfo
	(*ChessRogueInfo)(nil),                  // 2: ChessRogueInfo
	(*SceneInfo)(nil),                       // 3: SceneInfo
	(*LineupInfo)(nil),                      // 4: LineupInfo
	(*ChessRogueLayerInitialEventInfo)(nil), // 5: ChessRogueLayerInitialEventInfo
	(*ChessRogueCurrentInfo)(nil),           // 6: ChessRogueCurrentInfo
}
var file_ChessRogueStartScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueStartScRsp.info:type_name -> ChessRogueGameInfo
	2, // 1: ChessRogueStartScRsp.stage_info:type_name -> ChessRogueInfo
	3, // 2: ChessRogueStartScRsp.scene:type_name -> SceneInfo
	4, // 3: ChessRogueStartScRsp.lineup:type_name -> LineupInfo
	5, // 4: ChessRogueStartScRsp.board_event_info:type_name -> ChessRogueLayerInitialEventInfo
	6, // 5: ChessRogueStartScRsp.rogue_current_info:type_name -> ChessRogueCurrentInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ChessRogueStartScRsp_proto_init() }
func file_ChessRogueStartScRsp_proto_init() {
	if File_ChessRogueStartScRsp_proto != nil {
		return
	}
	file_ChessRogueInfo_proto_init()
	file_LineupInfo_proto_init()
	file_ChessRogueLayerInitialEventInfo_proto_init()
	file_ChessRogueGameInfo_proto_init()
	file_ChessRogueCurrentInfo_proto_init()
	file_SceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueStartScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueStartScRsp); i {
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
			RawDescriptor: file_ChessRogueStartScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueStartScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueStartScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueStartScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueStartScRsp_proto = out.File
	file_ChessRogueStartScRsp_proto_rawDesc = nil
	file_ChessRogueStartScRsp_proto_goTypes = nil
	file_ChessRogueStartScRsp_proto_depIdxs = nil
}
