// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PropExtraInfo.proto

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

type PropExtraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueTournWorkbenchInfo *RogueTournWorkbenchInfo `protobuf:"bytes,2,opt,name=rogue_tourn_workbench_info,json=rogueTournWorkbenchInfo,proto3" json:"rogue_tourn_workbench_info,omitempty"`
	// Types that are assignable to InfoOneofCase:
	//
	//	*PropExtraInfo_RogueInfo
	//	*PropExtraInfo_AeonInfo
	//	*PropExtraInfo_ChessRogueInfo
	//	*PropExtraInfo_RogueTournDoorInfo
	//	*PropExtraInfo_RogueGambleMachineInfo
	//	*PropExtraInfo_RogueCurseChestInfo
	InfoOneofCase isPropExtraInfo_InfoOneofCase `protobuf_oneof:"InfoOneofCase"`
}

func (x *PropExtraInfo) Reset() {
	*x = PropExtraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PropExtraInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropExtraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropExtraInfo) ProtoMessage() {}

func (x *PropExtraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PropExtraInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropExtraInfo.ProtoReflect.Descriptor instead.
func (*PropExtraInfo) Descriptor() ([]byte, []int) {
	return file_PropExtraInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PropExtraInfo) GetRogueTournWorkbenchInfo() *RogueTournWorkbenchInfo {
	if x != nil {
		return x.RogueTournWorkbenchInfo
	}
	return nil
}

func (m *PropExtraInfo) GetInfoOneofCase() isPropExtraInfo_InfoOneofCase {
	if m != nil {
		return m.InfoOneofCase
	}
	return nil
}

func (x *PropExtraInfo) GetRogueInfo() *PropRogueInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_RogueInfo); ok {
		return x.RogueInfo
	}
	return nil
}

func (x *PropExtraInfo) GetAeonInfo() *PropAeonInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_AeonInfo); ok {
		return x.AeonInfo
	}
	return nil
}

func (x *PropExtraInfo) GetChessRogueInfo() *PropChessRogueInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_ChessRogueInfo); ok {
		return x.ChessRogueInfo
	}
	return nil
}

func (x *PropExtraInfo) GetRogueTournDoorInfo() *RogueTournDoorInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_RogueTournDoorInfo); ok {
		return x.RogueTournDoorInfo
	}
	return nil
}

func (x *PropExtraInfo) GetRogueGambleMachineInfo() *RogueGambleMachineInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_RogueGambleMachineInfo); ok {
		return x.RogueGambleMachineInfo
	}
	return nil
}

func (x *PropExtraInfo) GetRogueCurseChestInfo() *RogueCurseChestInfo {
	if x, ok := x.GetInfoOneofCase().(*PropExtraInfo_RogueCurseChestInfo); ok {
		return x.RogueCurseChestInfo
	}
	return nil
}

type isPropExtraInfo_InfoOneofCase interface {
	isPropExtraInfo_InfoOneofCase()
}

type PropExtraInfo_RogueInfo struct {
	RogueInfo *PropRogueInfo `protobuf:"bytes,10,opt,name=rogue_info,json=rogueInfo,proto3,oneof"`
}

type PropExtraInfo_AeonInfo struct {
	AeonInfo *PropAeonInfo `protobuf:"bytes,7,opt,name=aeon_info,json=aeonInfo,proto3,oneof"`
}

type PropExtraInfo_ChessRogueInfo struct {
	ChessRogueInfo *PropChessRogueInfo `protobuf:"bytes,8,opt,name=chess_rogue_info,json=chessRogueInfo,proto3,oneof"`
}

type PropExtraInfo_RogueTournDoorInfo struct {
	RogueTournDoorInfo *RogueTournDoorInfo `protobuf:"bytes,12,opt,name=rogue_tourn_door_info,json=rogueTournDoorInfo,proto3,oneof"`
}

type PropExtraInfo_RogueGambleMachineInfo struct {
	RogueGambleMachineInfo *RogueGambleMachineInfo `protobuf:"bytes,13,opt,name=rogue_gamble_machine_info,json=rogueGambleMachineInfo,proto3,oneof"`
}

type PropExtraInfo_RogueCurseChestInfo struct {
	RogueCurseChestInfo *RogueCurseChestInfo `protobuf:"bytes,14,opt,name=rogue_curse_chest_info,json=rogueCurseChestInfo,proto3,oneof"`
}

func (*PropExtraInfo_RogueInfo) isPropExtraInfo_InfoOneofCase() {}

func (*PropExtraInfo_AeonInfo) isPropExtraInfo_InfoOneofCase() {}

func (*PropExtraInfo_ChessRogueInfo) isPropExtraInfo_InfoOneofCase() {}

func (*PropExtraInfo_RogueTournDoorInfo) isPropExtraInfo_InfoOneofCase() {}

func (*PropExtraInfo_RogueGambleMachineInfo) isPropExtraInfo_InfoOneofCase() {}

func (*PropExtraInfo_RogueCurseChestInfo) isPropExtraInfo_InfoOneofCase() {}

var File_PropExtraInfo_proto protoreflect.FileDescriptor

var file_PropExtraInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x62,
	0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x57,
	0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x50, 0x72, 0x6f, 0x70, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x73,
	0x65, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x44, 0x6f, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x1a,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x17, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x09, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x41, 0x65,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x61, 0x65, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x15, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75,
	0x72, 0x6e, 0x5f, 0x64, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x44,
	0x6f, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a,
	0x19, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x67, 0x61, 0x6d, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x62, 0x6c, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x16, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x47, 0x61, 0x6d, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x16, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x73, 0x65,
	0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x13, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x43, 0x75, 0x72, 0x73, 0x65, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x0f, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x61, 0x73,
	0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropExtraInfo_proto_rawDescOnce sync.Once
	file_PropExtraInfo_proto_rawDescData = file_PropExtraInfo_proto_rawDesc
)

func file_PropExtraInfo_proto_rawDescGZIP() []byte {
	file_PropExtraInfo_proto_rawDescOnce.Do(func() {
		file_PropExtraInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropExtraInfo_proto_rawDescData)
	})
	return file_PropExtraInfo_proto_rawDescData
}

var file_PropExtraInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PropExtraInfo_proto_goTypes = []interface{}{
	(*PropExtraInfo)(nil),           // 0: PropExtraInfo
	(*RogueTournWorkbenchInfo)(nil), // 1: RogueTournWorkbenchInfo
	(*PropRogueInfo)(nil),           // 2: PropRogueInfo
	(*PropAeonInfo)(nil),            // 3: PropAeonInfo
	(*PropChessRogueInfo)(nil),      // 4: PropChessRogueInfo
	(*RogueTournDoorInfo)(nil),      // 5: RogueTournDoorInfo
	(*RogueGambleMachineInfo)(nil),  // 6: RogueGambleMachineInfo
	(*RogueCurseChestInfo)(nil),     // 7: RogueCurseChestInfo
}
var file_PropExtraInfo_proto_depIdxs = []int32{
	1, // 0: PropExtraInfo.rogue_tourn_workbench_info:type_name -> RogueTournWorkbenchInfo
	2, // 1: PropExtraInfo.rogue_info:type_name -> PropRogueInfo
	3, // 2: PropExtraInfo.aeon_info:type_name -> PropAeonInfo
	4, // 3: PropExtraInfo.chess_rogue_info:type_name -> PropChessRogueInfo
	5, // 4: PropExtraInfo.rogue_tourn_door_info:type_name -> RogueTournDoorInfo
	6, // 5: PropExtraInfo.rogue_gamble_machine_info:type_name -> RogueGambleMachineInfo
	7, // 6: PropExtraInfo.rogue_curse_chest_info:type_name -> RogueCurseChestInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_PropExtraInfo_proto_init() }
func file_PropExtraInfo_proto_init() {
	if File_PropExtraInfo_proto != nil {
		return
	}
	file_RogueGambleMachineInfo_proto_init()
	file_RogueTournWorkbenchInfo_proto_init()
	file_PropRogueInfo_proto_init()
	file_PropChessRogueInfo_proto_init()
	file_PropAeonInfo_proto_init()
	file_RogueCurseChestInfo_proto_init()
	file_RogueTournDoorInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PropExtraInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropExtraInfo); i {
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
	file_PropExtraInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PropExtraInfo_RogueInfo)(nil),
		(*PropExtraInfo_AeonInfo)(nil),
		(*PropExtraInfo_ChessRogueInfo)(nil),
		(*PropExtraInfo_RogueTournDoorInfo)(nil),
		(*PropExtraInfo_RogueGambleMachineInfo)(nil),
		(*PropExtraInfo_RogueCurseChestInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PropExtraInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropExtraInfo_proto_goTypes,
		DependencyIndexes: file_PropExtraInfo_proto_depIdxs,
		MessageInfos:      file_PropExtraInfo_proto_msgTypes,
	}.Build()
	File_PropExtraInfo_proto = out.File
	file_PropExtraInfo_proto_rawDesc = nil
	file_PropExtraInfo_proto_goTypes = nil
	file_PropExtraInfo_proto_depIdxs = nil
}
