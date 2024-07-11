// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AvatarBattleInfo.proto

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

type AvatarBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType      AvatarType              `protobuf:"varint,1,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	Id              uint32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AvatarLevel     uint32                  `protobuf:"varint,3,opt,name=avatar_level,json=avatarLevel,proto3" json:"avatar_level,omitempty"`
	AvatarRank      uint32                  `protobuf:"varint,4,opt,name=avatar_rank,json=avatarRank,proto3" json:"avatar_rank,omitempty"`
	AvatarPromotion uint32                  `protobuf:"varint,5,opt,name=avatar_promotion,json=avatarPromotion,proto3" json:"avatar_promotion,omitempty"`
	AvatarStatus    *AvatarProperty         `protobuf:"bytes,6,opt,name=avatar_status,json=avatarStatus,proto3" json:"avatar_status,omitempty"`
	HOMLHNPHDOE     []*AvatarSkillTree      `protobuf:"bytes,7,rep,name=HOMLHNPHDOE,proto3" json:"HOMLHNPHDOE,omitempty"`
	JBJPCALLCHI     []*EquipmentProperty    `protobuf:"bytes,8,rep,name=JBJPCALLCHI,proto3" json:"JBJPCALLCHI,omitempty"`
	IFBNCFBOAIC     uint32                  `protobuf:"varint,9,opt,name=IFBNCFBOAIC,proto3" json:"IFBNCFBOAIC,omitempty"`
	CNBMGEBJCLL     float64                 `protobuf:"fixed64,10,opt,name=CNBMGEBJCLL,proto3" json:"CNBMGEBJCLL,omitempty"`
	HJPLFLJIFBJ     float64                 `protobuf:"fixed64,11,opt,name=HJPLFLJIFBJ,proto3" json:"HJPLFLJIFBJ,omitempty"`
	JOJKCHDPJIE     float64                 `protobuf:"fixed64,12,opt,name=JOJKCHDPJIE,proto3" json:"JOJKCHDPJIE,omitempty"`
	MJGNEKMLIIO     float64                 `protobuf:"fixed64,13,opt,name=MJGNEKMLIIO,proto3" json:"MJGNEKMLIIO,omitempty"`
	NOOHHDMAMDJ     float64                 `protobuf:"fixed64,14,opt,name=NOOHHDMAMDJ,proto3" json:"NOOHHDMAMDJ,omitempty"`
	StageId         uint32                  `protobuf:"varint,15,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	MAOIEEMDKAO     uint32                  `protobuf:"varint,16,opt,name=MAOIEEMDKAO,proto3" json:"MAOIEEMDKAO,omitempty"`
	HNNEOEMANLL     float64                 `protobuf:"fixed64,17,opt,name=HNNEOEMANLL,proto3" json:"HNNEOEMANLL,omitempty"`
	MPECNBHNKDE     []*AttackDamageProperty `protobuf:"bytes,18,rep,name=MPECNBHNKDE,proto3" json:"MPECNBHNKDE,omitempty"`
	LEBCLCHBOAN     []*AttackDamageProperty `protobuf:"bytes,19,rep,name=LEBCLCHBOAN,proto3" json:"LEBCLCHBOAN,omitempty"`
	IDGHCCBOBDP     []*AttackDamageProperty `protobuf:"bytes,20,rep,name=IDGHCCBOBDP,proto3" json:"IDGHCCBOBDP,omitempty"`
	LKCPDPMCLPD     []*SkillUseProperty     `protobuf:"bytes,21,rep,name=LKCPDPMCLPD,proto3" json:"LKCPDPMCLPD,omitempty"`
	ODOJOOCKHAE     float64                 `protobuf:"fixed64,22,opt,name=ODOJOOCKHAE,proto3" json:"ODOJOOCKHAE,omitempty"`
	OLGNOIDFOPD     uint32                  `protobuf:"varint,23,opt,name=OLGNOIDFOPD,proto3" json:"OLGNOIDFOPD,omitempty"`
	EGLFLNGFPIG     []*SpAddSource          `protobuf:"bytes,24,rep,name=EGLFLNGFPIG,proto3" json:"EGLFLNGFPIG,omitempty"`
	NJOHGLBIAHO     uint32                  `protobuf:"varint,25,opt,name=NJOHGLBIAHO,proto3" json:"NJOHGLBIAHO,omitempty"`
	MCPJBFLNIKN     uint32                  `protobuf:"varint,26,opt,name=MCPJBFLNIKN,proto3" json:"MCPJBFLNIKN,omitempty"`
	OLPHMFBLHMH     uint32                  `protobuf:"varint,27,opt,name=OLPHMFBLHMH,proto3" json:"OLPHMFBLHMH,omitempty"`
	BIFCFJEEMIB     uint32                  `protobuf:"varint,28,opt,name=BIFCFJEEMIB,proto3" json:"BIFCFJEEMIB,omitempty"`
	IEJINOGMEDK     uint32                  `protobuf:"varint,29,opt,name=IEJINOGMEDK,proto3" json:"IEJINOGMEDK,omitempty"`
	BHHHPOGMGLC     float64                 `protobuf:"fixed64,30,opt,name=BHHHPOGMGLC,proto3" json:"BHHHPOGMGLC,omitempty"`
	KCNGNMIHBCE     float64                 `protobuf:"fixed64,31,opt,name=KCNGNMIHBCE,proto3" json:"KCNGNMIHBCE,omitempty"`
	DEFDDFIDLFM     float64                 `protobuf:"fixed64,32,opt,name=DEFDDFIDLFM,proto3" json:"DEFDDFIDLFM,omitempty"`
	FCDFNJALENJ     *AvatarProperty         `protobuf:"bytes,33,opt,name=FCDFNJALENJ,proto3" json:"FCDFNJALENJ,omitempty"`
	BKMIEEOKPBN     []*BattleRelic          `protobuf:"bytes,34,rep,name=BKMIEEOKPBN,proto3" json:"BKMIEEOKPBN,omitempty"`
	AssistUid       uint32                  `protobuf:"varint,35,opt,name=assist_uid,json=assistUid,proto3" json:"assist_uid,omitempty"`
	LJMPDNAHOPD     []*AttackDamageProperty `protobuf:"bytes,36,rep,name=LJMPDNAHOPD,proto3" json:"LJMPDNAHOPD,omitempty"`
	AINLHBOJBOK     float64                 `protobuf:"fixed64,37,opt,name=AINLHBOJBOK,proto3" json:"AINLHBOJBOK,omitempty"`
	OOKKLHPJGMD     float64                 `protobuf:"fixed64,38,opt,name=OOKKLHPJGMD,proto3" json:"OOKKLHPJGMD,omitempty"`
	EKAEMACBAGB     float64                 `protobuf:"fixed64,39,opt,name=EKAEMACBAGB,proto3" json:"EKAEMACBAGB,omitempty"`
	JBDIJFDBGHF     float64                 `protobuf:"fixed64,40,opt,name=JBDIJFDBGHF,proto3" json:"JBDIJFDBGHF,omitempty"`
	DONKCNNCCPC     []*AbilityUseStt        `protobuf:"bytes,41,rep,name=DONKCNNCCPC,proto3" json:"DONKCNNCCPC,omitempty"`
	IINDPNEGNCC     uint32                  `protobuf:"varint,42,opt,name=IINDPNEGNCC,proto3" json:"IINDPNEGNCC,omitempty"`
	NMIKBCDPCAN     uint32                  `protobuf:"varint,43,opt,name=NMIKBCDPCAN,proto3" json:"NMIKBCDPCAN,omitempty"`
}

func (x *AvatarBattleInfo) Reset() {
	*x = AvatarBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarBattleInfo) ProtoMessage() {}

func (x *AvatarBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarBattleInfo.ProtoReflect.Descriptor instead.
func (*AvatarBattleInfo) Descriptor() ([]byte, []int) {
	return file_AvatarBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarBattleInfo) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *AvatarBattleInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarLevel() uint32 {
	if x != nil {
		return x.AvatarLevel
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarRank() uint32 {
	if x != nil {
		return x.AvatarRank
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarPromotion() uint32 {
	if x != nil {
		return x.AvatarPromotion
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarStatus() *AvatarProperty {
	if x != nil {
		return x.AvatarStatus
	}
	return nil
}

func (x *AvatarBattleInfo) GetHOMLHNPHDOE() []*AvatarSkillTree {
	if x != nil {
		return x.HOMLHNPHDOE
	}
	return nil
}

func (x *AvatarBattleInfo) GetJBJPCALLCHI() []*EquipmentProperty {
	if x != nil {
		return x.JBJPCALLCHI
	}
	return nil
}

func (x *AvatarBattleInfo) GetIFBNCFBOAIC() uint32 {
	if x != nil {
		return x.IFBNCFBOAIC
	}
	return 0
}

func (x *AvatarBattleInfo) GetCNBMGEBJCLL() float64 {
	if x != nil {
		return x.CNBMGEBJCLL
	}
	return 0
}

func (x *AvatarBattleInfo) GetHJPLFLJIFBJ() float64 {
	if x != nil {
		return x.HJPLFLJIFBJ
	}
	return 0
}

func (x *AvatarBattleInfo) GetJOJKCHDPJIE() float64 {
	if x != nil {
		return x.JOJKCHDPJIE
	}
	return 0
}

func (x *AvatarBattleInfo) GetMJGNEKMLIIO() float64 {
	if x != nil {
		return x.MJGNEKMLIIO
	}
	return 0
}

func (x *AvatarBattleInfo) GetNOOHHDMAMDJ() float64 {
	if x != nil {
		return x.NOOHHDMAMDJ
	}
	return 0
}

func (x *AvatarBattleInfo) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *AvatarBattleInfo) GetMAOIEEMDKAO() uint32 {
	if x != nil {
		return x.MAOIEEMDKAO
	}
	return 0
}

func (x *AvatarBattleInfo) GetHNNEOEMANLL() float64 {
	if x != nil {
		return x.HNNEOEMANLL
	}
	return 0
}

func (x *AvatarBattleInfo) GetMPECNBHNKDE() []*AttackDamageProperty {
	if x != nil {
		return x.MPECNBHNKDE
	}
	return nil
}

func (x *AvatarBattleInfo) GetLEBCLCHBOAN() []*AttackDamageProperty {
	if x != nil {
		return x.LEBCLCHBOAN
	}
	return nil
}

func (x *AvatarBattleInfo) GetIDGHCCBOBDP() []*AttackDamageProperty {
	if x != nil {
		return x.IDGHCCBOBDP
	}
	return nil
}

func (x *AvatarBattleInfo) GetLKCPDPMCLPD() []*SkillUseProperty {
	if x != nil {
		return x.LKCPDPMCLPD
	}
	return nil
}

func (x *AvatarBattleInfo) GetODOJOOCKHAE() float64 {
	if x != nil {
		return x.ODOJOOCKHAE
	}
	return 0
}

func (x *AvatarBattleInfo) GetOLGNOIDFOPD() uint32 {
	if x != nil {
		return x.OLGNOIDFOPD
	}
	return 0
}

func (x *AvatarBattleInfo) GetEGLFLNGFPIG() []*SpAddSource {
	if x != nil {
		return x.EGLFLNGFPIG
	}
	return nil
}

func (x *AvatarBattleInfo) GetNJOHGLBIAHO() uint32 {
	if x != nil {
		return x.NJOHGLBIAHO
	}
	return 0
}

func (x *AvatarBattleInfo) GetMCPJBFLNIKN() uint32 {
	if x != nil {
		return x.MCPJBFLNIKN
	}
	return 0
}

func (x *AvatarBattleInfo) GetOLPHMFBLHMH() uint32 {
	if x != nil {
		return x.OLPHMFBLHMH
	}
	return 0
}

func (x *AvatarBattleInfo) GetBIFCFJEEMIB() uint32 {
	if x != nil {
		return x.BIFCFJEEMIB
	}
	return 0
}

func (x *AvatarBattleInfo) GetIEJINOGMEDK() uint32 {
	if x != nil {
		return x.IEJINOGMEDK
	}
	return 0
}

func (x *AvatarBattleInfo) GetBHHHPOGMGLC() float64 {
	if x != nil {
		return x.BHHHPOGMGLC
	}
	return 0
}

func (x *AvatarBattleInfo) GetKCNGNMIHBCE() float64 {
	if x != nil {
		return x.KCNGNMIHBCE
	}
	return 0
}

func (x *AvatarBattleInfo) GetDEFDDFIDLFM() float64 {
	if x != nil {
		return x.DEFDDFIDLFM
	}
	return 0
}

func (x *AvatarBattleInfo) GetFCDFNJALENJ() *AvatarProperty {
	if x != nil {
		return x.FCDFNJALENJ
	}
	return nil
}

func (x *AvatarBattleInfo) GetBKMIEEOKPBN() []*BattleRelic {
	if x != nil {
		return x.BKMIEEOKPBN
	}
	return nil
}

func (x *AvatarBattleInfo) GetAssistUid() uint32 {
	if x != nil {
		return x.AssistUid
	}
	return 0
}

func (x *AvatarBattleInfo) GetLJMPDNAHOPD() []*AttackDamageProperty {
	if x != nil {
		return x.LJMPDNAHOPD
	}
	return nil
}

func (x *AvatarBattleInfo) GetAINLHBOJBOK() float64 {
	if x != nil {
		return x.AINLHBOJBOK
	}
	return 0
}

func (x *AvatarBattleInfo) GetOOKKLHPJGMD() float64 {
	if x != nil {
		return x.OOKKLHPJGMD
	}
	return 0
}

func (x *AvatarBattleInfo) GetEKAEMACBAGB() float64 {
	if x != nil {
		return x.EKAEMACBAGB
	}
	return 0
}

func (x *AvatarBattleInfo) GetJBDIJFDBGHF() float64 {
	if x != nil {
		return x.JBDIJFDBGHF
	}
	return 0
}

func (x *AvatarBattleInfo) GetDONKCNNCCPC() []*AbilityUseStt {
	if x != nil {
		return x.DONKCNNCCPC
	}
	return nil
}

func (x *AvatarBattleInfo) GetIINDPNEGNCC() uint32 {
	if x != nil {
		return x.IINDPNEGNCC
	}
	return 0
}

func (x *AvatarBattleInfo) GetNMIKBCDPCAN() uint32 {
	if x != nil {
		return x.NMIKBCDPCAN
	}
	return 0
}

var File_AvatarBattleInfo_proto protoreflect.FileDescriptor

var file_AvatarBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x53, 0x70, 0x41, 0x64, 0x64, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x73, 0x65, 0x53, 0x74, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x54, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa7, 0x0d, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0b, 0x48, 0x4f, 0x4d, 0x4c, 0x48,
	0x4e, 0x50, 0x48, 0x44, 0x4f, 0x45, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0b,
	0x48, 0x4f, 0x4d, 0x4c, 0x48, 0x4e, 0x50, 0x48, 0x44, 0x4f, 0x45, 0x12, 0x34, 0x0a, 0x0b, 0x4a,
	0x42, 0x4a, 0x50, 0x43, 0x41, 0x4c, 0x4c, 0x43, 0x48, 0x49, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x0b, 0x4a, 0x42, 0x4a, 0x50, 0x43, 0x41, 0x4c, 0x4c, 0x43, 0x48,
	0x49, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x46, 0x42, 0x4e, 0x43, 0x46, 0x42, 0x4f, 0x41, 0x49, 0x43,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x46, 0x42, 0x4e, 0x43, 0x46, 0x42, 0x4f,
	0x41, 0x49, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4e, 0x42, 0x4d, 0x47, 0x45, 0x42, 0x4a, 0x43,
	0x4c, 0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x43, 0x4e, 0x42, 0x4d, 0x47, 0x45,
	0x42, 0x4a, 0x43, 0x4c, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4a, 0x50, 0x4c, 0x46, 0x4c, 0x4a,
	0x49, 0x46, 0x42, 0x4a, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x48, 0x4a, 0x50, 0x4c,
	0x46, 0x4c, 0x4a, 0x49, 0x46, 0x42, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4f, 0x4a, 0x4b, 0x43,
	0x48, 0x44, 0x50, 0x4a, 0x49, 0x45, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x4a, 0x4f,
	0x4a, 0x4b, 0x43, 0x48, 0x44, 0x50, 0x4a, 0x49, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4a, 0x47,
	0x4e, 0x45, 0x4b, 0x4d, 0x4c, 0x49, 0x49, 0x4f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x4d, 0x4a, 0x47, 0x4e, 0x45, 0x4b, 0x4d, 0x4c, 0x49, 0x49, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4e,
	0x4f, 0x4f, 0x48, 0x48, 0x44, 0x4d, 0x41, 0x4d, 0x44, 0x4a, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x4e, 0x4f, 0x4f, 0x48, 0x48, 0x44, 0x4d, 0x41, 0x4d, 0x44, 0x4a, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x41, 0x4f, 0x49,
	0x45, 0x45, 0x4d, 0x44, 0x4b, 0x41, 0x4f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d,
	0x41, 0x4f, 0x49, 0x45, 0x45, 0x4d, 0x44, 0x4b, 0x41, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4e,
	0x4e, 0x45, 0x4f, 0x45, 0x4d, 0x41, 0x4e, 0x4c, 0x4c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x48, 0x4e, 0x4e, 0x45, 0x4f, 0x45, 0x4d, 0x41, 0x4e, 0x4c, 0x4c, 0x12, 0x37, 0x0a, 0x0b,
	0x4d, 0x50, 0x45, 0x43, 0x4e, 0x42, 0x48, 0x4e, 0x4b, 0x44, 0x45, 0x18, 0x12, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0b, 0x4d, 0x50, 0x45, 0x43, 0x4e, 0x42,
	0x48, 0x4e, 0x4b, 0x44, 0x45, 0x12, 0x37, 0x0a, 0x0b, 0x4c, 0x45, 0x42, 0x43, 0x4c, 0x43, 0x48,
	0x42, 0x4f, 0x41, 0x4e, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x0b, 0x4c, 0x45, 0x42, 0x43, 0x4c, 0x43, 0x48, 0x42, 0x4f, 0x41, 0x4e, 0x12, 0x37,
	0x0a, 0x0b, 0x49, 0x44, 0x47, 0x48, 0x43, 0x43, 0x42, 0x4f, 0x42, 0x44, 0x50, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0b, 0x49, 0x44, 0x47, 0x48,
	0x43, 0x43, 0x42, 0x4f, 0x42, 0x44, 0x50, 0x12, 0x33, 0x0a, 0x0b, 0x4c, 0x4b, 0x43, 0x50, 0x44,
	0x50, 0x4d, 0x43, 0x4c, 0x50, 0x44, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x0b, 0x4c, 0x4b, 0x43, 0x50, 0x44, 0x50, 0x4d, 0x43, 0x4c, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x4f, 0x44, 0x4f, 0x4a, 0x4f, 0x4f, 0x43, 0x4b, 0x48, 0x41, 0x45, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x4f, 0x44, 0x4f, 0x4a, 0x4f, 0x4f, 0x43, 0x4b, 0x48, 0x41, 0x45, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x4c, 0x47, 0x4e, 0x4f, 0x49, 0x44, 0x46, 0x4f, 0x50, 0x44, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x47, 0x4e, 0x4f, 0x49, 0x44, 0x46, 0x4f, 0x50, 0x44,
	0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x47, 0x4c, 0x46, 0x4c, 0x4e, 0x47, 0x46, 0x50, 0x49, 0x47, 0x18,
	0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x70, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0b, 0x45, 0x47, 0x4c, 0x46, 0x4c, 0x4e, 0x47, 0x46, 0x50, 0x49, 0x47,
	0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4a, 0x4f, 0x48, 0x47, 0x4c, 0x42, 0x49, 0x41, 0x48, 0x4f, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4a, 0x4f, 0x48, 0x47, 0x4c, 0x42, 0x49, 0x41,
	0x48, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x43, 0x50, 0x4a, 0x42, 0x46, 0x4c, 0x4e, 0x49, 0x4b,
	0x4e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x43, 0x50, 0x4a, 0x42, 0x46, 0x4c,
	0x4e, 0x49, 0x4b, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4c, 0x50, 0x48, 0x4d, 0x46, 0x42, 0x4c,
	0x48, 0x4d, 0x48, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x50, 0x48, 0x4d,
	0x46, 0x42, 0x4c, 0x48, 0x4d, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x49, 0x46, 0x43, 0x46, 0x4a,
	0x45, 0x45, 0x4d, 0x49, 0x42, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x49, 0x46,
	0x43, 0x46, 0x4a, 0x45, 0x45, 0x4d, 0x49, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x45, 0x4a, 0x49,
	0x4e, 0x4f, 0x47, 0x4d, 0x45, 0x44, 0x4b, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49,
	0x45, 0x4a, 0x49, 0x4e, 0x4f, 0x47, 0x4d, 0x45, 0x44, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x48,
	0x48, 0x48, 0x50, 0x4f, 0x47, 0x4d, 0x47, 0x4c, 0x43, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x42, 0x48, 0x48, 0x48, 0x50, 0x4f, 0x47, 0x4d, 0x47, 0x4c, 0x43, 0x12, 0x20, 0x0a, 0x0b,
	0x4b, 0x43, 0x4e, 0x47, 0x4e, 0x4d, 0x49, 0x48, 0x42, 0x43, 0x45, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x4b, 0x43, 0x4e, 0x47, 0x4e, 0x4d, 0x49, 0x48, 0x42, 0x43, 0x45, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x45, 0x46, 0x44, 0x44, 0x46, 0x49, 0x44, 0x4c, 0x46, 0x4d, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x44, 0x45, 0x46, 0x44, 0x44, 0x46, 0x49, 0x44, 0x4c, 0x46, 0x4d,
	0x12, 0x31, 0x0a, 0x0b, 0x46, 0x43, 0x44, 0x46, 0x4e, 0x4a, 0x41, 0x4c, 0x45, 0x4e, 0x4a, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0b, 0x46, 0x43, 0x44, 0x46, 0x4e, 0x4a, 0x41, 0x4c,
	0x45, 0x4e, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x4b, 0x4d, 0x49, 0x45, 0x45, 0x4f, 0x4b, 0x50,
	0x42, 0x4e, 0x18, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x0b, 0x42, 0x4b, 0x4d, 0x49, 0x45, 0x45, 0x4f, 0x4b,
	0x50, 0x42, 0x4e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x4c, 0x4a, 0x4d, 0x50, 0x44, 0x4e, 0x41, 0x48, 0x4f, 0x50,
	0x44, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0b,
	0x4c, 0x4a, 0x4d, 0x50, 0x44, 0x4e, 0x41, 0x48, 0x4f, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x49, 0x4e, 0x4c, 0x48, 0x42, 0x4f, 0x4a, 0x42, 0x4f, 0x4b, 0x18, 0x25, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x41, 0x49, 0x4e, 0x4c, 0x48, 0x42, 0x4f, 0x4a, 0x42, 0x4f, 0x4b, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x4f, 0x4b, 0x4b, 0x4c, 0x48, 0x50, 0x4a, 0x47, 0x4d, 0x44, 0x18, 0x26, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x4f, 0x4f, 0x4b, 0x4b, 0x4c, 0x48, 0x50, 0x4a, 0x47, 0x4d, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x45, 0x4b, 0x41, 0x45, 0x4d, 0x41, 0x43, 0x42, 0x41, 0x47, 0x42, 0x18, 0x27,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x45, 0x4b, 0x41, 0x45, 0x4d, 0x41, 0x43, 0x42, 0x41, 0x47,
	0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x42, 0x44, 0x49, 0x4a, 0x46, 0x44, 0x42, 0x47, 0x48, 0x46,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x4a, 0x42, 0x44, 0x49, 0x4a, 0x46, 0x44, 0x42,
	0x47, 0x48, 0x46, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x4f, 0x4e, 0x4b, 0x43, 0x4e, 0x4e, 0x43, 0x43,
	0x50, 0x43, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x55, 0x73, 0x65, 0x53, 0x74, 0x74, 0x52, 0x0b, 0x44, 0x4f, 0x4e, 0x4b, 0x43, 0x4e,
	0x4e, 0x43, 0x43, 0x50, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x49, 0x4e, 0x44, 0x50, 0x4e, 0x45,
	0x47, 0x4e, 0x43, 0x43, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x49, 0x4e, 0x44,
	0x50, 0x4e, 0x45, 0x47, 0x4e, 0x43, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4d, 0x49, 0x4b, 0x42,
	0x43, 0x44, 0x50, 0x43, 0x41, 0x4e, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4d,
	0x49, 0x4b, 0x42, 0x43, 0x44, 0x50, 0x43, 0x41, 0x4e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarBattleInfo_proto_rawDescOnce sync.Once
	file_AvatarBattleInfo_proto_rawDescData = file_AvatarBattleInfo_proto_rawDesc
)

func file_AvatarBattleInfo_proto_rawDescGZIP() []byte {
	file_AvatarBattleInfo_proto_rawDescOnce.Do(func() {
		file_AvatarBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarBattleInfo_proto_rawDescData)
	})
	return file_AvatarBattleInfo_proto_rawDescData
}

var file_AvatarBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarBattleInfo_proto_goTypes = []interface{}{
	(*AvatarBattleInfo)(nil),     // 0: AvatarBattleInfo
	(AvatarType)(0),              // 1: AvatarType
	(*AvatarProperty)(nil),       // 2: AvatarProperty
	(*AvatarSkillTree)(nil),      // 3: AvatarSkillTree
	(*EquipmentProperty)(nil),    // 4: EquipmentProperty
	(*AttackDamageProperty)(nil), // 5: AttackDamageProperty
	(*SkillUseProperty)(nil),     // 6: SkillUseProperty
	(*SpAddSource)(nil),          // 7: SpAddSource
	(*BattleRelic)(nil),          // 8: BattleRelic
	(*AbilityUseStt)(nil),        // 9: AbilityUseStt
}
var file_AvatarBattleInfo_proto_depIdxs = []int32{
	1,  // 0: AvatarBattleInfo.avatar_type:type_name -> AvatarType
	2,  // 1: AvatarBattleInfo.avatar_status:type_name -> AvatarProperty
	3,  // 2: AvatarBattleInfo.HOMLHNPHDOE:type_name -> AvatarSkillTree
	4,  // 3: AvatarBattleInfo.JBJPCALLCHI:type_name -> EquipmentProperty
	5,  // 4: AvatarBattleInfo.MPECNBHNKDE:type_name -> AttackDamageProperty
	5,  // 5: AvatarBattleInfo.LEBCLCHBOAN:type_name -> AttackDamageProperty
	5,  // 6: AvatarBattleInfo.IDGHCCBOBDP:type_name -> AttackDamageProperty
	6,  // 7: AvatarBattleInfo.LKCPDPMCLPD:type_name -> SkillUseProperty
	7,  // 8: AvatarBattleInfo.EGLFLNGFPIG:type_name -> SpAddSource
	2,  // 9: AvatarBattleInfo.FCDFNJALENJ:type_name -> AvatarProperty
	8,  // 10: AvatarBattleInfo.BKMIEEOKPBN:type_name -> BattleRelic
	5,  // 11: AvatarBattleInfo.LJMPDNAHOPD:type_name -> AttackDamageProperty
	9,  // 12: AvatarBattleInfo.DONKCNNCCPC:type_name -> AbilityUseStt
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_AvatarBattleInfo_proto_init() }
func file_AvatarBattleInfo_proto_init() {
	if File_AvatarBattleInfo_proto != nil {
		return
	}
	file_SpAddSource_proto_init()
	file_BattleRelic_proto_init()
	file_AbilityUseStt_proto_init()
	file_AvatarSkillTree_proto_init()
	file_EquipmentProperty_proto_init()
	file_AvatarProperty_proto_init()
	file_AttackDamageProperty_proto_init()
	file_SkillUseProperty_proto_init()
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarBattleInfo); i {
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
			RawDescriptor: file_AvatarBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarBattleInfo_proto_goTypes,
		DependencyIndexes: file_AvatarBattleInfo_proto_depIdxs,
		MessageInfos:      file_AvatarBattleInfo_proto_msgTypes,
	}.Build()
	File_AvatarBattleInfo_proto = out.File
	file_AvatarBattleInfo_proto_rawDesc = nil
	file_AvatarBattleInfo_proto_goTypes = nil
	file_AvatarBattleInfo_proto_depIdxs = nil
}
