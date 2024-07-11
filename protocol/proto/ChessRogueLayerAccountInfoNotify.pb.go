// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueLayerAccountInfoNotify.proto

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

type ChessRogueLayerAccountInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishInfo          *ChessRogueFinishInfo `protobuf:"bytes,7,opt,name=finish_info,json=finishInfo,proto3" json:"finish_info,omitempty"`
	AreaDifficultyLevel uint32                `protobuf:"varint,9,opt,name=area_difficulty_level,json=areaDifficultyLevel,proto3" json:"area_difficulty_level,omitempty"`
	ENPGPALPCDL         []uint32              `protobuf:"varint,11,rep,packed,name=ENPGPALPCDL,proto3" json:"ENPGPALPCDL,omitempty"`
	LevelInfo           *ChessRogueLevelInfo  `protobuf:"bytes,12,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	LayerId             uint32                `protobuf:"varint,1,opt,name=layer_id,json=layerId,proto3" json:"layer_id,omitempty"`
}

func (x *ChessRogueLayerAccountInfoNotify) Reset() {
	*x = ChessRogueLayerAccountInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueLayerAccountInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueLayerAccountInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueLayerAccountInfoNotify) ProtoMessage() {}

func (x *ChessRogueLayerAccountInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueLayerAccountInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueLayerAccountInfoNotify.ProtoReflect.Descriptor instead.
func (*ChessRogueLayerAccountInfoNotify) Descriptor() ([]byte, []int) {
	return file_ChessRogueLayerAccountInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueLayerAccountInfoNotify) GetFinishInfo() *ChessRogueFinishInfo {
	if x != nil {
		return x.FinishInfo
	}
	return nil
}

func (x *ChessRogueLayerAccountInfoNotify) GetAreaDifficultyLevel() uint32 {
	if x != nil {
		return x.AreaDifficultyLevel
	}
	return 0
}

func (x *ChessRogueLayerAccountInfoNotify) GetENPGPALPCDL() []uint32 {
	if x != nil {
		return x.ENPGPALPCDL
	}
	return nil
}

func (x *ChessRogueLayerAccountInfoNotify) GetLevelInfo() *ChessRogueLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

func (x *ChessRogueLayerAccountInfoNotify) GetLayerId() uint32 {
	if x != nil {
		return x.LayerId
	}
	return 0
}

var File_ChessRogueLayerAccountInfoNotify_proto protoreflect.FileDescriptor

var file_ChessRogueLayerAccountInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x80, 0x02, 0x0a, 0x20, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x15,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x72, 0x65,
	0x61, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4e, 0x50, 0x47, 0x50, 0x41, 0x4c, 0x50, 0x43, 0x44, 0x4c, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4e, 0x50, 0x47, 0x50, 0x41, 0x4c, 0x50, 0x43,
	0x44, 0x4c, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueLayerAccountInfoNotify_proto_rawDescOnce sync.Once
	file_ChessRogueLayerAccountInfoNotify_proto_rawDescData = file_ChessRogueLayerAccountInfoNotify_proto_rawDesc
)

func file_ChessRogueLayerAccountInfoNotify_proto_rawDescGZIP() []byte {
	file_ChessRogueLayerAccountInfoNotify_proto_rawDescOnce.Do(func() {
		file_ChessRogueLayerAccountInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueLayerAccountInfoNotify_proto_rawDescData)
	})
	return file_ChessRogueLayerAccountInfoNotify_proto_rawDescData
}

var file_ChessRogueLayerAccountInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueLayerAccountInfoNotify_proto_goTypes = []interface{}{
	(*ChessRogueLayerAccountInfoNotify)(nil), // 0: ChessRogueLayerAccountInfoNotify
	(*ChessRogueFinishInfo)(nil),             // 1: ChessRogueFinishInfo
	(*ChessRogueLevelInfo)(nil),              // 2: ChessRogueLevelInfo
}
var file_ChessRogueLayerAccountInfoNotify_proto_depIdxs = []int32{
	1, // 0: ChessRogueLayerAccountInfoNotify.finish_info:type_name -> ChessRogueFinishInfo
	2, // 1: ChessRogueLayerAccountInfoNotify.level_info:type_name -> ChessRogueLevelInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChessRogueLayerAccountInfoNotify_proto_init() }
func file_ChessRogueLayerAccountInfoNotify_proto_init() {
	if File_ChessRogueLayerAccountInfoNotify_proto != nil {
		return
	}
	file_ChessRogueLevelInfo_proto_init()
	file_ChessRogueFinishInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueLayerAccountInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueLayerAccountInfoNotify); i {
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
			RawDescriptor: file_ChessRogueLayerAccountInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueLayerAccountInfoNotify_proto_goTypes,
		DependencyIndexes: file_ChessRogueLayerAccountInfoNotify_proto_depIdxs,
		MessageInfos:      file_ChessRogueLayerAccountInfoNotify_proto_msgTypes,
	}.Build()
	File_ChessRogueLayerAccountInfoNotify_proto = out.File
	file_ChessRogueLayerAccountInfoNotify_proto_rawDesc = nil
	file_ChessRogueLayerAccountInfoNotify_proto_goTypes = nil
	file_ChessRogueLayerAccountInfoNotify_proto_depIdxs = nil
}
