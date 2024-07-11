// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueLevelInfo.proto

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

type ChessRogueLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FGCMINNFNHN        int32                     `protobuf:"varint,14,opt,name=FGCMINNFNHN,proto3" json:"FGCMINNFNHN,omitempty"`
	LevelStatus        ChessRogueLevelStatusType `protobuf:"varint,5,opt,name=level_status,json=levelStatus,proto3,enum=ChessRogueLevelStatusType" json:"level_status,omitempty"`
	ActionPoint        int32                     `protobuf:"varint,8,opt,name=action_point,json=actionPoint,proto3" json:"action_point,omitempty"`
	LMJDBNHFJOO        uint32                    `protobuf:"varint,13,opt,name=LMJDBNHFJOO,proto3" json:"LMJDBNHFJOO,omitempty"`
	AreaInfo           *ChessRogueAreaInfo       `protobuf:"bytes,12,opt,name=area_info,json=areaInfo,proto3" json:"area_info,omitempty"`
	Id                 uint32                    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	AreaIdList         []uint32                  `protobuf:"varint,1,rep,packed,name=area_id_list,json=areaIdList,proto3" json:"area_id_list,omitempty"`
	LayerId            uint32                    `protobuf:"varint,15,opt,name=layer_id,json=layerId,proto3" json:"layer_id,omitempty"`
	ExploredAreaIdList []uint32                  `protobuf:"varint,3,rep,packed,name=explored_area_id_list,json=exploredAreaIdList,proto3" json:"explored_area_id_list,omitempty"`
}

func (x *ChessRogueLevelInfo) Reset() {
	*x = ChessRogueLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueLevelInfo) ProtoMessage() {}

func (x *ChessRogueLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueLevelInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueLevelInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueLevelInfo) GetFGCMINNFNHN() int32 {
	if x != nil {
		return x.FGCMINNFNHN
	}
	return 0
}

func (x *ChessRogueLevelInfo) GetLevelStatus() ChessRogueLevelStatusType {
	if x != nil {
		return x.LevelStatus
	}
	return ChessRogueLevelStatusType_CHESS_ROGUE_LEVEL_IDLE
}

func (x *ChessRogueLevelInfo) GetActionPoint() int32 {
	if x != nil {
		return x.ActionPoint
	}
	return 0
}

func (x *ChessRogueLevelInfo) GetLMJDBNHFJOO() uint32 {
	if x != nil {
		return x.LMJDBNHFJOO
	}
	return 0
}

func (x *ChessRogueLevelInfo) GetAreaInfo() *ChessRogueAreaInfo {
	if x != nil {
		return x.AreaInfo
	}
	return nil
}

func (x *ChessRogueLevelInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChessRogueLevelInfo) GetAreaIdList() []uint32 {
	if x != nil {
		return x.AreaIdList
	}
	return nil
}

func (x *ChessRogueLevelInfo) GetLayerId() uint32 {
	if x != nil {
		return x.LayerId
	}
	return 0
}

func (x *ChessRogueLevelInfo) GetExploredAreaIdList() []uint32 {
	if x != nil {
		return x.ExploredAreaIdList
	}
	return nil
}

var File_ChessRogueLevelInfo_proto protoreflect.FileDescriptor

var file_ChessRogueLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x46, 0x47, 0x43, 0x4d, 0x49, 0x4e, 0x4e, 0x46, 0x4e, 0x48, 0x4e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x46, 0x47, 0x43, 0x4d, 0x49, 0x4e, 0x4e, 0x46, 0x4e, 0x48, 0x4e,
	0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4d, 0x4a, 0x44, 0x42, 0x4e, 0x48, 0x46, 0x4a, 0x4f,
	0x4f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4d, 0x4a, 0x44, 0x42, 0x4e, 0x48,
	0x46, 0x4a, 0x4f, 0x4f, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x12, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueLevelInfo_proto_rawDescOnce sync.Once
	file_ChessRogueLevelInfo_proto_rawDescData = file_ChessRogueLevelInfo_proto_rawDesc
)

func file_ChessRogueLevelInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueLevelInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueLevelInfo_proto_rawDescData)
	})
	return file_ChessRogueLevelInfo_proto_rawDescData
}

var file_ChessRogueLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueLevelInfo_proto_goTypes = []interface{}{
	(*ChessRogueLevelInfo)(nil),    // 0: ChessRogueLevelInfo
	(ChessRogueLevelStatusType)(0), // 1: ChessRogueLevelStatusType
	(*ChessRogueAreaInfo)(nil),     // 2: ChessRogueAreaInfo
}
var file_ChessRogueLevelInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueLevelInfo.level_status:type_name -> ChessRogueLevelStatusType
	2, // 1: ChessRogueLevelInfo.area_info:type_name -> ChessRogueAreaInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChessRogueLevelInfo_proto_init() }
func file_ChessRogueLevelInfo_proto_init() {
	if File_ChessRogueLevelInfo_proto != nil {
		return
	}
	file_ChessRogueAreaInfo_proto_init()
	file_ChessRogueLevelStatusType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueLevelInfo); i {
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
			RawDescriptor: file_ChessRogueLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueLevelInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueLevelInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueLevelInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueLevelInfo_proto = out.File
	file_ChessRogueLevelInfo_proto_rawDesc = nil
	file_ChessRogueLevelInfo_proto_goTypes = nil
	file_ChessRogueLevelInfo_proto_depIdxs = nil
}
