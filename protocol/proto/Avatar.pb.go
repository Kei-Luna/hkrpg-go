// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Avatar.proto

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

type Avatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAvatarId                uint32             `protobuf:"varint,15,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	Level                       uint32             `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	EquipmentUniqueId           uint32             `protobuf:"varint,3,opt,name=equipment_unique_id,json=equipmentUniqueId,proto3" json:"equipment_unique_id,omitempty"`
	EquipRelicList              []*EquipRelic      `protobuf:"bytes,2,rep,name=equip_relic_list,json=equipRelicList,proto3" json:"equip_relic_list,omitempty"`
	SkilltreeList               []*AvatarSkillTree `protobuf:"bytes,14,rep,name=skilltree_list,json=skilltreeList,proto3" json:"skilltree_list,omitempty"`
	FirstMetTimeStamp           uint64             `protobuf:"varint,4,opt,name=first_met_time_stamp,json=firstMetTimeStamp,proto3" json:"first_met_time_stamp,omitempty"`
	Rank                        uint32             `protobuf:"varint,13,opt,name=rank,proto3" json:"rank,omitempty"`
	DressedSkinId               uint32             `protobuf:"varint,8,opt,name=dressed_skin_id,json=dressedSkinId,proto3" json:"dressed_skin_id,omitempty"`
	IsMarked                    bool               `protobuf:"varint,6,opt,name=is_marked,json=isMarked,proto3" json:"is_marked,omitempty"`
	Exp                         uint32             `protobuf:"varint,10,opt,name=exp,proto3" json:"exp,omitempty"`
	Promotion                   uint32             `protobuf:"varint,9,opt,name=promotion,proto3" json:"promotion,omitempty"`
	HasTakenPromotionRewardList []uint32           `protobuf:"varint,11,rep,packed,name=has_taken_promotion_reward_list,json=hasTakenPromotionRewardList,proto3" json:"has_taken_promotion_reward_list,omitempty"`
}

func (x *Avatar) Reset() {
	*x = Avatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Avatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Avatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Avatar) ProtoMessage() {}

func (x *Avatar) ProtoReflect() protoreflect.Message {
	mi := &file_Avatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Avatar.ProtoReflect.Descriptor instead.
func (*Avatar) Descriptor() ([]byte, []int) {
	return file_Avatar_proto_rawDescGZIP(), []int{0}
}

func (x *Avatar) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *Avatar) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Avatar) GetEquipmentUniqueId() uint32 {
	if x != nil {
		return x.EquipmentUniqueId
	}
	return 0
}

func (x *Avatar) GetEquipRelicList() []*EquipRelic {
	if x != nil {
		return x.EquipRelicList
	}
	return nil
}

func (x *Avatar) GetSkilltreeList() []*AvatarSkillTree {
	if x != nil {
		return x.SkilltreeList
	}
	return nil
}

func (x *Avatar) GetFirstMetTimeStamp() uint64 {
	if x != nil {
		return x.FirstMetTimeStamp
	}
	return 0
}

func (x *Avatar) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Avatar) GetDressedSkinId() uint32 {
	if x != nil {
		return x.DressedSkinId
	}
	return 0
}

func (x *Avatar) GetIsMarked() bool {
	if x != nil {
		return x.IsMarked
	}
	return false
}

func (x *Avatar) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Avatar) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

func (x *Avatar) GetHasTakenPromotionRewardList() []uint32 {
	if x != nil {
		return x.HasTakenPromotionRewardList
	}
	return nil
}

var File_Avatar_proto protoreflect.FileDescriptor

var file_Avatar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e,
	0x0a, 0x13, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x0e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x74, 0x72, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x14, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x73,
	0x6b, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x1f, 0x68, 0x61, 0x73, 0x5f,
	0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x1b, 0x68, 0x61, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Avatar_proto_rawDescOnce sync.Once
	file_Avatar_proto_rawDescData = file_Avatar_proto_rawDesc
)

func file_Avatar_proto_rawDescGZIP() []byte {
	file_Avatar_proto_rawDescOnce.Do(func() {
		file_Avatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_Avatar_proto_rawDescData)
	})
	return file_Avatar_proto_rawDescData
}

var file_Avatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Avatar_proto_goTypes = []interface{}{
	(*Avatar)(nil),          // 0: Avatar
	(*EquipRelic)(nil),      // 1: EquipRelic
	(*AvatarSkillTree)(nil), // 2: AvatarSkillTree
}
var file_Avatar_proto_depIdxs = []int32{
	1, // 0: Avatar.equip_relic_list:type_name -> EquipRelic
	2, // 1: Avatar.skilltree_list:type_name -> AvatarSkillTree
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Avatar_proto_init() }
func file_Avatar_proto_init() {
	if File_Avatar_proto != nil {
		return
	}
	file_AvatarSkillTree_proto_init()
	file_EquipRelic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Avatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Avatar); i {
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
			RawDescriptor: file_Avatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Avatar_proto_goTypes,
		DependencyIndexes: file_Avatar_proto_depIdxs,
		MessageInfos:      file_Avatar_proto_msgTypes,
	}.Build()
	File_Avatar_proto = out.File
	file_Avatar_proto_rawDesc = nil
	file_Avatar_proto_goTypes = nil
	file_Avatar_proto_depIdxs = nil
}
