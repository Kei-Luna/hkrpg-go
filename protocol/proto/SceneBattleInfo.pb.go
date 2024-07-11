// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneBattleInfo.proto

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

type SceneBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterWaveList  []*SceneMonsterWave          `protobuf:"bytes,7,rep,name=monster_wave_list,json=monsterWaveList,proto3" json:"monster_wave_list,omitempty"`
	BattleTargetInfo map[uint32]*BattleTargetList `protobuf:"bytes,583,rep,name=battle_target_info,json=battleTargetInfo,proto3" json:"battle_target_info,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BattleId         uint32                       `protobuf:"varint,15,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
	WorldLevel       uint32                       `protobuf:"varint,1,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	BuffList         []*BattleBuff                `protobuf:"bytes,6,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	BGHGLGGPBGA      uint32                       `protobuf:"varint,10,opt,name=BGHGLGGPBGA,proto3" json:"BGHGLGGPBGA,omitempty"`
	NHOKHJGKKLO      *LNIFNCALIOM                 `protobuf:"bytes,1604,opt,name=NHOKHJGKKLO,proto3" json:"NHOKHJGKKLO,omitempty"`
	OJLBBHGGFJH      *EvolveBuildBattleInfo       `protobuf:"bytes,209,opt,name=OJLBBHGGFJH,proto3" json:"OJLBBHGGFJH,omitempty"`
	RoundsLimit      uint32                       `protobuf:"varint,8,opt,name=rounds_limit,json=roundsLimit,proto3" json:"rounds_limit,omitempty"`
	LogicRandomSeed  uint32                       `protobuf:"varint,3,opt,name=logic_random_seed,json=logicRandomSeed,proto3" json:"logic_random_seed,omitempty"`
	AvatarBattleList []*BattleAvatar              `protobuf:"bytes,4,rep,name=avatar_battle_list,json=avatarBattleList,proto3" json:"avatar_battle_list,omitempty"`
	FCNEFNEFGDA      bool                         `protobuf:"varint,14,opt,name=FCNEFNEFGDA,proto3" json:"FCNEFNEFGDA,omitempty"`
	BattleEvent      []*BattleEventBattleInfo     `protobuf:"bytes,413,rep,name=battle_event,json=battleEvent,proto3" json:"battle_event,omitempty"`
	StageId          uint32                       `protobuf:"varint,11,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *SceneBattleInfo) Reset() {
	*x = SceneBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneBattleInfo) ProtoMessage() {}

func (x *SceneBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneBattleInfo.ProtoReflect.Descriptor instead.
func (*SceneBattleInfo) Descriptor() ([]byte, []int) {
	return file_SceneBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneBattleInfo) GetMonsterWaveList() []*SceneMonsterWave {
	if x != nil {
		return x.MonsterWaveList
	}
	return nil
}

func (x *SceneBattleInfo) GetBattleTargetInfo() map[uint32]*BattleTargetList {
	if x != nil {
		return x.BattleTargetInfo
	}
	return nil
}

func (x *SceneBattleInfo) GetBattleId() uint32 {
	if x != nil {
		return x.BattleId
	}
	return 0
}

func (x *SceneBattleInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *SceneBattleInfo) GetBuffList() []*BattleBuff {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *SceneBattleInfo) GetBGHGLGGPBGA() uint32 {
	if x != nil {
		return x.BGHGLGGPBGA
	}
	return 0
}

func (x *SceneBattleInfo) GetNHOKHJGKKLO() *LNIFNCALIOM {
	if x != nil {
		return x.NHOKHJGKKLO
	}
	return nil
}

func (x *SceneBattleInfo) GetOJLBBHGGFJH() *EvolveBuildBattleInfo {
	if x != nil {
		return x.OJLBBHGGFJH
	}
	return nil
}

func (x *SceneBattleInfo) GetRoundsLimit() uint32 {
	if x != nil {
		return x.RoundsLimit
	}
	return 0
}

func (x *SceneBattleInfo) GetLogicRandomSeed() uint32 {
	if x != nil {
		return x.LogicRandomSeed
	}
	return 0
}

func (x *SceneBattleInfo) GetAvatarBattleList() []*BattleAvatar {
	if x != nil {
		return x.AvatarBattleList
	}
	return nil
}

func (x *SceneBattleInfo) GetFCNEFNEFGDA() bool {
	if x != nil {
		return x.FCNEFNEFGDA
	}
	return false
}

func (x *SceneBattleInfo) GetBattleEvent() []*BattleEventBattleInfo {
	if x != nil {
		return x.BattleEvent
	}
	return nil
}

func (x *SceneBattleInfo) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

var File_SceneBattleInfo_proto protoreflect.FileDescriptor

var file_SceneBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x45,
	0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4c, 0x4e, 0x49, 0x46, 0x4e, 0x43, 0x41, 0x4c, 0x49, 0x4f, 0x4d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x05, 0x0a, 0x0f, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d,
	0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x76, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x52, 0x0f, 0x6d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55, 0x0a,
	0x12, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0xc7, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x28, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x42, 0x47, 0x48, 0x47, 0x4c, 0x47, 0x47, 0x50, 0x42, 0x47, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x42, 0x47, 0x48, 0x47, 0x4c, 0x47, 0x47, 0x50, 0x42, 0x47, 0x41, 0x12, 0x2f,
	0x0a, 0x0b, 0x4e, 0x48, 0x4f, 0x4b, 0x48, 0x4a, 0x47, 0x4b, 0x4b, 0x4c, 0x4f, 0x18, 0xc4, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4e, 0x49, 0x46, 0x4e, 0x43, 0x41, 0x4c, 0x49,
	0x4f, 0x4d, 0x52, 0x0b, 0x4e, 0x48, 0x4f, 0x4b, 0x48, 0x4a, 0x47, 0x4b, 0x4b, 0x4c, 0x4f, 0x12,
	0x39, 0x0a, 0x0b, 0x4f, 0x4a, 0x4c, 0x42, 0x42, 0x48, 0x47, 0x47, 0x46, 0x4a, 0x48, 0x18, 0xd1,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x4f,
	0x4a, 0x4c, 0x42, 0x42, 0x48, 0x47, 0x47, 0x46, 0x4a, 0x48, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x73, 0x65,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x12, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x43, 0x4e, 0x45, 0x46, 0x4e,
	0x45, 0x46, 0x47, 0x44, 0x41, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x43, 0x4e,
	0x45, 0x46, 0x4e, 0x45, 0x46, 0x47, 0x44, 0x41, 0x12, 0x3a, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x9d, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x1a,
	0x56, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneBattleInfo_proto_rawDescOnce sync.Once
	file_SceneBattleInfo_proto_rawDescData = file_SceneBattleInfo_proto_rawDesc
)

func file_SceneBattleInfo_proto_rawDescGZIP() []byte {
	file_SceneBattleInfo_proto_rawDescOnce.Do(func() {
		file_SceneBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneBattleInfo_proto_rawDescData)
	})
	return file_SceneBattleInfo_proto_rawDescData
}

var file_SceneBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneBattleInfo_proto_goTypes = []interface{}{
	(*SceneBattleInfo)(nil),       // 0: SceneBattleInfo
	nil,                           // 1: SceneBattleInfo.BattleTargetInfoEntry
	(*SceneMonsterWave)(nil),      // 2: SceneMonsterWave
	(*BattleBuff)(nil),            // 3: BattleBuff
	(*LNIFNCALIOM)(nil),           // 4: LNIFNCALIOM
	(*EvolveBuildBattleInfo)(nil), // 5: EvolveBuildBattleInfo
	(*BattleAvatar)(nil),          // 6: BattleAvatar
	(*BattleEventBattleInfo)(nil), // 7: BattleEventBattleInfo
	(*BattleTargetList)(nil),      // 8: BattleTargetList
}
var file_SceneBattleInfo_proto_depIdxs = []int32{
	2, // 0: SceneBattleInfo.monster_wave_list:type_name -> SceneMonsterWave
	1, // 1: SceneBattleInfo.battle_target_info:type_name -> SceneBattleInfo.BattleTargetInfoEntry
	3, // 2: SceneBattleInfo.buff_list:type_name -> BattleBuff
	4, // 3: SceneBattleInfo.NHOKHJGKKLO:type_name -> LNIFNCALIOM
	5, // 4: SceneBattleInfo.OJLBBHGGFJH:type_name -> EvolveBuildBattleInfo
	6, // 5: SceneBattleInfo.avatar_battle_list:type_name -> BattleAvatar
	7, // 6: SceneBattleInfo.battle_event:type_name -> BattleEventBattleInfo
	8, // 7: SceneBattleInfo.BattleTargetInfoEntry.value:type_name -> BattleTargetList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_SceneBattleInfo_proto_init() }
func file_SceneBattleInfo_proto_init() {
	if File_SceneBattleInfo_proto != nil {
		return
	}
	file_BattleAvatar_proto_init()
	file_BattleBuff_proto_init()
	file_EvolveBuildBattleInfo_proto_init()
	file_BattleEventBattleInfo_proto_init()
	file_SceneMonsterWave_proto_init()
	file_LNIFNCALIOM_proto_init()
	file_BattleTargetList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneBattleInfo); i {
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
			RawDescriptor: file_SceneBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneBattleInfo_proto_goTypes,
		DependencyIndexes: file_SceneBattleInfo_proto_depIdxs,
		MessageInfos:      file_SceneBattleInfo_proto_msgTypes,
	}.Build()
	File_SceneBattleInfo_proto = out.File
	file_SceneBattleInfo_proto_rawDesc = nil
	file_SceneBattleInfo_proto_goTypes = nil
	file_SceneBattleInfo_proto_depIdxs = nil
}
