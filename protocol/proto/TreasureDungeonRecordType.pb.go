// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonRecordType.proto

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

type TreasureDungeonRecordType int32

const (
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_NONE                        TreasureDungeonRecordType = 0
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_HP                      TreasureDungeonRecordType = 1
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SUB_HP                      TreasureDungeonRecordType = 2
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SUB_HP_NO_EXPLORE           TreasureDungeonRecordType = 3
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_ATTACK                  TreasureDungeonRecordType = 5
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_DEFENCE                 TreasureDungeonRecordType = 6
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_EXPLORE                 TreasureDungeonRecordType = 9
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SUB_EXPLORE                 TreasureDungeonRecordType = 10
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_EXPLORE_OVERFLOW        TreasureDungeonRecordType = 11
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SUMMON                      TreasureDungeonRecordType = 15
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_KILL                        TreasureDungeonRecordType = 16
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_TRIAL_AVATAR            TreasureDungeonRecordType = 20
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ADD_BUFF                    TreasureDungeonRecordType = 24
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_UNLOCK_DOOR                 TreasureDungeonRecordType = 25
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ENEMY_ENHANCE               TreasureDungeonRecordType = 27
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ENEMY_WEAKEN                TreasureDungeonRecordType = 28
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ENEMY_AURA_REMOVE           TreasureDungeonRecordType = 29
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_RUN         TreasureDungeonRecordType = 30
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_KILL        TreasureDungeonRecordType = 31
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_SUCCESS TreasureDungeonRecordType = 33
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_FAIL    TreasureDungeonRecordType = 34
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_EXPLORE     TreasureDungeonRecordType = 35
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_BATTLE_BUFF_OPEN_GRID       TreasureDungeonRecordType = 36
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_ITEM        TreasureDungeonRecordType = 37
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_AVATAR_DEAD                 TreasureDungeonRecordType = 40
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_TRIAL_AVATAR_DEAD           TreasureDungeonRecordType = 41
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_ALL_AVATAR_DEAD             TreasureDungeonRecordType = 42
	TreasureDungeonRecordType_TREASURE_DUNGEON_RECORD_OPEN_ITEM_CHEST             TreasureDungeonRecordType = 43
)

// Enum value maps for TreasureDungeonRecordType.
var (
	TreasureDungeonRecordType_name = map[int32]string{
		0:  "TREASURE_DUNGEON_RECORD_NONE",
		1:  "TREASURE_DUNGEON_RECORD_ADD_HP",
		2:  "TREASURE_DUNGEON_RECORD_SUB_HP",
		3:  "TREASURE_DUNGEON_RECORD_SUB_HP_NO_EXPLORE",
		5:  "TREASURE_DUNGEON_RECORD_ADD_ATTACK",
		6:  "TREASURE_DUNGEON_RECORD_ADD_DEFENCE",
		9:  "TREASURE_DUNGEON_RECORD_ADD_EXPLORE",
		10: "TREASURE_DUNGEON_RECORD_SUB_EXPLORE",
		11: "TREASURE_DUNGEON_RECORD_ADD_EXPLORE_OVERFLOW",
		15: "TREASURE_DUNGEON_RECORD_SUMMON",
		16: "TREASURE_DUNGEON_RECORD_KILL",
		20: "TREASURE_DUNGEON_RECORD_ADD_TRIAL_AVATAR",
		24: "TREASURE_DUNGEON_RECORD_ADD_BUFF",
		25: "TREASURE_DUNGEON_RECORD_UNLOCK_DOOR",
		27: "TREASURE_DUNGEON_RECORD_ENEMY_ENHANCE",
		28: "TREASURE_DUNGEON_RECORD_ENEMY_WEAKEN",
		29: "TREASURE_DUNGEON_RECORD_ENEMY_AURA_REMOVE",
		30: "TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_RUN",
		31: "TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_KILL",
		33: "TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_SUCCESS",
		34: "TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_FAIL",
		35: "TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_EXPLORE",
		36: "TREASURE_DUNGEON_RECORD_BATTLE_BUFF_OPEN_GRID",
		37: "TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_ITEM",
		40: "TREASURE_DUNGEON_RECORD_AVATAR_DEAD",
		41: "TREASURE_DUNGEON_RECORD_TRIAL_AVATAR_DEAD",
		42: "TREASURE_DUNGEON_RECORD_ALL_AVATAR_DEAD",
		43: "TREASURE_DUNGEON_RECORD_OPEN_ITEM_CHEST",
	}
	TreasureDungeonRecordType_value = map[string]int32{
		"TREASURE_DUNGEON_RECORD_NONE":                        0,
		"TREASURE_DUNGEON_RECORD_ADD_HP":                      1,
		"TREASURE_DUNGEON_RECORD_SUB_HP":                      2,
		"TREASURE_DUNGEON_RECORD_SUB_HP_NO_EXPLORE":           3,
		"TREASURE_DUNGEON_RECORD_ADD_ATTACK":                  5,
		"TREASURE_DUNGEON_RECORD_ADD_DEFENCE":                 6,
		"TREASURE_DUNGEON_RECORD_ADD_EXPLORE":                 9,
		"TREASURE_DUNGEON_RECORD_SUB_EXPLORE":                 10,
		"TREASURE_DUNGEON_RECORD_ADD_EXPLORE_OVERFLOW":        11,
		"TREASURE_DUNGEON_RECORD_SUMMON":                      15,
		"TREASURE_DUNGEON_RECORD_KILL":                        16,
		"TREASURE_DUNGEON_RECORD_ADD_TRIAL_AVATAR":            20,
		"TREASURE_DUNGEON_RECORD_ADD_BUFF":                    24,
		"TREASURE_DUNGEON_RECORD_UNLOCK_DOOR":                 25,
		"TREASURE_DUNGEON_RECORD_ENEMY_ENHANCE":               27,
		"TREASURE_DUNGEON_RECORD_ENEMY_WEAKEN":                28,
		"TREASURE_DUNGEON_RECORD_ENEMY_AURA_REMOVE":           29,
		"TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_RUN":         30,
		"TREASURE_DUNGEON_RECORD_SPECIAL_MONSTER_KILL":        31,
		"TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_SUCCESS": 33,
		"TREASURE_DUNGEON_RECORD_BATTLE_BUFF_TRIGGER_FAIL":    34,
		"TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_EXPLORE":     35,
		"TREASURE_DUNGEON_RECORD_BATTLE_BUFF_OPEN_GRID":       36,
		"TREASURE_DUNGEON_RECORD_BATTLE_BUFF_ADD_ITEM":        37,
		"TREASURE_DUNGEON_RECORD_AVATAR_DEAD":                 40,
		"TREASURE_DUNGEON_RECORD_TRIAL_AVATAR_DEAD":           41,
		"TREASURE_DUNGEON_RECORD_ALL_AVATAR_DEAD":             42,
		"TREASURE_DUNGEON_RECORD_OPEN_ITEM_CHEST":             43,
	}
)

func (x TreasureDungeonRecordType) Enum() *TreasureDungeonRecordType {
	p := new(TreasureDungeonRecordType)
	*p = x
	return p
}

func (x TreasureDungeonRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TreasureDungeonRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_TreasureDungeonRecordType_proto_enumTypes[0].Descriptor()
}

func (TreasureDungeonRecordType) Type() protoreflect.EnumType {
	return &file_TreasureDungeonRecordType_proto_enumTypes[0]
}

func (x TreasureDungeonRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TreasureDungeonRecordType.Descriptor instead.
func (TreasureDungeonRecordType) EnumDescriptor() ([]byte, []int) {
	return file_TreasureDungeonRecordType_proto_rawDescGZIP(), []int{0}
}

var File_TreasureDungeonRecordType_proto protoreflect.FileDescriptor

var file_TreasureDungeonRecordType_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xee, 0x09, 0x0a, 0x19, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x1c, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47,
	0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55,
	0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x48, 0x50, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x53, 0x55, 0x42, 0x5f, 0x48, 0x50, 0x10, 0x02, 0x12, 0x2d, 0x0a, 0x29, 0x54, 0x52, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x48, 0x50, 0x5f, 0x4e, 0x4f, 0x5f, 0x45,
	0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x52, 0x45, 0x41,
	0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x05,
	0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e,
	0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x44, 0x45, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x06, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45,
	0x10, 0x09, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44,
	0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x53, 0x55,
	0x42, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45, 0x10, 0x0a, 0x12, 0x30, 0x0a, 0x2c, 0x54,
	0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f,
	0x52, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x0b, 0x12, 0x22, 0x0a,
	0x1e, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x4f, 0x4e, 0x10,
	0x0f, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55,
	0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4b, 0x49, 0x4c,
	0x4c, 0x10, 0x10, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f,
	0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x10,
	0x14, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55,
	0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x42, 0x55, 0x46, 0x46, 0x10, 0x18, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x45, 0x41, 0x53,
	0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x44, 0x4f, 0x4f, 0x52, 0x10, 0x19,
	0x12, 0x29, 0x0a, 0x25, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e,
	0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x45, 0x4d,
	0x59, 0x5f, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x1b, 0x12, 0x28, 0x0a, 0x24, 0x54,
	0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x45, 0x4e, 0x45, 0x4d, 0x59, 0x5f, 0x57, 0x45, 0x41,
	0x4b, 0x45, 0x4e, 0x10, 0x1c, 0x12, 0x2d, 0x0a, 0x29, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x45, 0x4e, 0x45, 0x4d, 0x59, 0x5f, 0x41, 0x55, 0x52, 0x41, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x10, 0x1d, 0x12, 0x2f, 0x0a, 0x2b, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45,
	0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f,
	0x52, 0x55, 0x4e, 0x10, 0x1e, 0x12, 0x30, 0x0a, 0x2c, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x1f, 0x12, 0x37, 0x0a, 0x33, 0x54, 0x52, 0x45, 0x41, 0x53,
	0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x54,
	0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x21,
	0x12, 0x34, 0x0a, 0x30, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e,
	0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x42, 0x41, 0x54, 0x54,
	0x4c, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x22, 0x12, 0x33, 0x0a, 0x2f, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55,
	0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52,
	0x44, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x41, 0x44,
	0x44, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x4f, 0x52, 0x45, 0x10, 0x23, 0x12, 0x31, 0x0a, 0x2d, 0x54,
	0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x55,
	0x46, 0x46, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x47, 0x52, 0x49, 0x44, 0x10, 0x24, 0x12, 0x30,
	0x0a, 0x2c, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45,
	0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x25,
	0x12, 0x27, 0x0a, 0x23, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e,
	0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x56, 0x41, 0x54,
	0x41, 0x52, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x28, 0x12, 0x2d, 0x0a, 0x29, 0x54, 0x52, 0x45,
	0x41, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41,
	0x52, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x29, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x52, 0x45, 0x41,
	0x53, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x44,
	0x45, 0x41, 0x44, 0x10, 0x2a, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x52, 0x45, 0x41, 0x53, 0x55, 0x52,
	0x45, 0x5f, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x54,
	0x10, 0x2b, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonRecordType_proto_rawDescOnce sync.Once
	file_TreasureDungeonRecordType_proto_rawDescData = file_TreasureDungeonRecordType_proto_rawDesc
)

func file_TreasureDungeonRecordType_proto_rawDescGZIP() []byte {
	file_TreasureDungeonRecordType_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonRecordType_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonRecordType_proto_rawDescData)
	})
	return file_TreasureDungeonRecordType_proto_rawDescData
}

var file_TreasureDungeonRecordType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TreasureDungeonRecordType_proto_goTypes = []interface{}{
	(TreasureDungeonRecordType)(0), // 0: TreasureDungeonRecordType
}
var file_TreasureDungeonRecordType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TreasureDungeonRecordType_proto_init() }
func file_TreasureDungeonRecordType_proto_init() {
	if File_TreasureDungeonRecordType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TreasureDungeonRecordType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonRecordType_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonRecordType_proto_depIdxs,
		EnumInfos:         file_TreasureDungeonRecordType_proto_enumTypes,
	}.Build()
	File_TreasureDungeonRecordType_proto = out.File
	file_TreasureDungeonRecordType_proto_rawDesc = nil
	file_TreasureDungeonRecordType_proto_goTypes = nil
	file_TreasureDungeonRecordType_proto_depIdxs = nil
}