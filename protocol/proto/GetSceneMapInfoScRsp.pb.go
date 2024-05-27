// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetSceneMapInfoScRsp.proto

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

type GetSceneMapInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId            uint32           `protobuf:"varint,6,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	UnlockTeleportList []uint32         `protobuf:"varint,3,rep,packed,name=unlock_teleport_list,json=unlockTeleportList,proto3" json:"unlock_teleport_list,omitempty"`
	MazePropList       []*MazePropState `protobuf:"bytes,8,rep,name=maze_prop_list,json=mazePropList,proto3" json:"maze_prop_list,omitempty"`
	CurMapEntryId      uint32           `protobuf:"varint,10,opt,name=cur_map_entry_id,json=curMapEntryId,proto3" json:"cur_map_entry_id,omitempty"`
	ChestList          []*ChestInfo     `protobuf:"bytes,14,rep,name=chest_list,json=chestList,proto3" json:"chest_list,omitempty"`
	LightenSectionList []uint32         `protobuf:"varint,9,rep,packed,name=lighten_section_list,json=lightenSectionList,proto3" json:"lighten_section_list,omitempty"`
	MazeGroupList      []*MazeGroup     `protobuf:"bytes,13,rep,name=maze_group_list,json=mazeGroupList,proto3" json:"maze_group_list,omitempty"`
	SceneMapInfo       []*SceneMapInfo  `protobuf:"bytes,2,rep,name=scene_map_info,json=sceneMapInfo,proto3" json:"scene_map_info,omitempty"`
	DHDOICCKIFL        bool             `protobuf:"varint,5,opt,name=DHDOICCKIFL,proto3" json:"DHDOICCKIFL,omitempty"`
	JFAMOJFHHDL        uint32           `protobuf:"varint,4,opt,name=JFAMOJFHHDL,proto3" json:"JFAMOJFHHDL,omitempty"`
	Retcode            uint32           `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetSceneMapInfoScRsp) Reset() {
	*x = GetSceneMapInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneMapInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneMapInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneMapInfoScRsp) ProtoMessage() {}

func (x *GetSceneMapInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneMapInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneMapInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetSceneMapInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetSceneMapInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneMapInfoScRsp) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *GetSceneMapInfoScRsp) GetUnlockTeleportList() []uint32 {
	if x != nil {
		return x.UnlockTeleportList
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetMazePropList() []*MazePropState {
	if x != nil {
		return x.MazePropList
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetCurMapEntryId() uint32 {
	if x != nil {
		return x.CurMapEntryId
	}
	return 0
}

func (x *GetSceneMapInfoScRsp) GetChestList() []*ChestInfo {
	if x != nil {
		return x.ChestList
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetLightenSectionList() []uint32 {
	if x != nil {
		return x.LightenSectionList
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetMazeGroupList() []*MazeGroup {
	if x != nil {
		return x.MazeGroupList
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetSceneMapInfo() []*SceneMapInfo {
	if x != nil {
		return x.SceneMapInfo
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetDHDOICCKIFL() bool {
	if x != nil {
		return x.DHDOICCKIFL
	}
	return false
}

func (x *GetSceneMapInfoScRsp) GetJFAMOJFHHDL() uint32 {
	if x != nil {
		return x.JFAMOJFHHDL
	}
	return 0
}

func (x *GetSceneMapInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetSceneMapInfoScRsp_proto protoreflect.FileDescriptor

var file_GetSceneMapInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61,
	0x7a, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x43,
	0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x4d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x03, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x0e, 0x6d, 0x61, 0x7a, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x6d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63,
	0x75, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x6e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0f, 0x6d, 0x61, 0x7a,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d,
	0x6d, 0x61, 0x7a, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x0e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x48, 0x44, 0x4f, 0x49, 0x43, 0x43, 0x4b, 0x49, 0x46,
	0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x48, 0x44, 0x4f, 0x49, 0x43, 0x43,
	0x4b, 0x49, 0x46, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x46, 0x41, 0x4d, 0x4f, 0x4a, 0x46, 0x48,
	0x48, 0x44, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x46, 0x41, 0x4d, 0x4f,
	0x4a, 0x46, 0x48, 0x48, 0x44, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSceneMapInfoScRsp_proto_rawDescOnce sync.Once
	file_GetSceneMapInfoScRsp_proto_rawDescData = file_GetSceneMapInfoScRsp_proto_rawDesc
)

func file_GetSceneMapInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetSceneMapInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetSceneMapInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneMapInfoScRsp_proto_rawDescData)
	})
	return file_GetSceneMapInfoScRsp_proto_rawDescData
}

var file_GetSceneMapInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneMapInfoScRsp_proto_goTypes = []interface{}{
	(*GetSceneMapInfoScRsp)(nil), // 0: GetSceneMapInfoScRsp
	(*MazePropState)(nil),        // 1: MazePropState
	(*ChestInfo)(nil),            // 2: ChestInfo
	(*MazeGroup)(nil),            // 3: MazeGroup
	(*SceneMapInfo)(nil),         // 4: SceneMapInfo
}
var file_GetSceneMapInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetSceneMapInfoScRsp.maze_prop_list:type_name -> MazePropState
	2, // 1: GetSceneMapInfoScRsp.chest_list:type_name -> ChestInfo
	3, // 2: GetSceneMapInfoScRsp.maze_group_list:type_name -> MazeGroup
	4, // 3: GetSceneMapInfoScRsp.scene_map_info:type_name -> SceneMapInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GetSceneMapInfoScRsp_proto_init() }
func file_GetSceneMapInfoScRsp_proto_init() {
	if File_GetSceneMapInfoScRsp_proto != nil {
		return
	}
	file_MazeGroup_proto_init()
	file_ChestInfo_proto_init()
	file_SceneMapInfo_proto_init()
	file_MazePropState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetSceneMapInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneMapInfoScRsp); i {
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
			RawDescriptor: file_GetSceneMapInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneMapInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetSceneMapInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetSceneMapInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetSceneMapInfoScRsp_proto = out.File
	file_GetSceneMapInfoScRsp_proto_rawDesc = nil
	file_GetSceneMapInfoScRsp_proto_goTypes = nil
	file_GetSceneMapInfoScRsp_proto_depIdxs = nil
}
