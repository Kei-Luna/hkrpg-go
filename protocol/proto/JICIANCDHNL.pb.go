// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JICIANCDHNL.proto

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

type JICIANCDHNL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LFPMPLLHGEP []*EANIBKNOEKJ `protobuf:"bytes,14,rep,name=LFPMPLLHGEP,proto3" json:"LFPMPLLHGEP,omitempty"`
	ILJGCGMDKFC uint64         `protobuf:"varint,5,opt,name=ILJGCGMDKFC,proto3" json:"ILJGCGMDKFC,omitempty"`
	JIHCKGMGCNI uint32         `protobuf:"varint,6,opt,name=JIHCKGMGCNI,proto3" json:"JIHCKGMGCNI,omitempty"`
	IDOJEMJKOCG uint32         `protobuf:"varint,8,opt,name=IDOJEMJKOCG,proto3" json:"IDOJEMJKOCG,omitempty"`
	BHNPEGHHEIJ string         `protobuf:"bytes,39,opt,name=BHNPEGHHEIJ,proto3" json:"BHNPEGHHEIJ,omitempty"`
	KIOENEKCBMN uint32         `protobuf:"varint,13,opt,name=KIOENEKCBMN,proto3" json:"KIOENEKCBMN,omitempty"`
	NNJONDGNAOC []*NILMPNJDCHP `protobuf:"bytes,4,rep,name=NNJONDGNAOC,proto3" json:"NNJONDGNAOC,omitempty"`
	MIHGBIBKFPD []*EANIBKNOEKJ `protobuf:"bytes,3,rep,name=MIHGBIBKFPD,proto3" json:"MIHGBIBKFPD,omitempty"`
	PoolId      uint32         `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	AEFGEMHJMAK uint64         `protobuf:"varint,15,opt,name=AEFGEMHJMAK,proto3" json:"AEFGEMHJMAK,omitempty"`
	IAEDCABHNJL uint64         `protobuf:"varint,7,opt,name=IAEDCABHNJL,proto3" json:"IAEDCABHNJL,omitempty"`
	OCMBKJNAGIF []*EANIBKNOEKJ `protobuf:"bytes,10,rep,name=OCMBKJNAGIF,proto3" json:"OCMBKJNAGIF,omitempty"`
	PMPIOLCDFEM uint32         `protobuf:"varint,9,opt,name=PMPIOLCDFEM,proto3" json:"PMPIOLCDFEM,omitempty"`
	LOKBDIGDCCN uint32         `protobuf:"varint,2,opt,name=LOKBDIGDCCN,proto3" json:"LOKBDIGDCCN,omitempty"`
	IKFMFKBMHHH string         `protobuf:"bytes,1923,opt,name=IKFMFKBMHHH,proto3" json:"IKFMFKBMHHH,omitempty"`
	LALCLOONFLE bool           `protobuf:"varint,11,opt,name=LALCLOONFLE,proto3" json:"LALCLOONFLE,omitempty"`
	NIMPMPMLJDI string         `protobuf:"bytes,948,opt,name=NIMPMPMLJDI,proto3" json:"NIMPMPMLJDI,omitempty"`
	IGDJAOELMKM []*EANIBKNOEKJ `protobuf:"bytes,12,rep,name=IGDJAOELMKM,proto3" json:"IGDJAOELMKM,omitempty"`
}

func (x *JICIANCDHNL) Reset() {
	*x = JICIANCDHNL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JICIANCDHNL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JICIANCDHNL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JICIANCDHNL) ProtoMessage() {}

func (x *JICIANCDHNL) ProtoReflect() protoreflect.Message {
	mi := &file_JICIANCDHNL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JICIANCDHNL.ProtoReflect.Descriptor instead.
func (*JICIANCDHNL) Descriptor() ([]byte, []int) {
	return file_JICIANCDHNL_proto_rawDescGZIP(), []int{0}
}

func (x *JICIANCDHNL) GetLFPMPLLHGEP() []*EANIBKNOEKJ {
	if x != nil {
		return x.LFPMPLLHGEP
	}
	return nil
}

func (x *JICIANCDHNL) GetILJGCGMDKFC() uint64 {
	if x != nil {
		return x.ILJGCGMDKFC
	}
	return 0
}

func (x *JICIANCDHNL) GetJIHCKGMGCNI() uint32 {
	if x != nil {
		return x.JIHCKGMGCNI
	}
	return 0
}

func (x *JICIANCDHNL) GetIDOJEMJKOCG() uint32 {
	if x != nil {
		return x.IDOJEMJKOCG
	}
	return 0
}

func (x *JICIANCDHNL) GetBHNPEGHHEIJ() string {
	if x != nil {
		return x.BHNPEGHHEIJ
	}
	return ""
}

func (x *JICIANCDHNL) GetKIOENEKCBMN() uint32 {
	if x != nil {
		return x.KIOENEKCBMN
	}
	return 0
}

func (x *JICIANCDHNL) GetNNJONDGNAOC() []*NILMPNJDCHP {
	if x != nil {
		return x.NNJONDGNAOC
	}
	return nil
}

func (x *JICIANCDHNL) GetMIHGBIBKFPD() []*EANIBKNOEKJ {
	if x != nil {
		return x.MIHGBIBKFPD
	}
	return nil
}

func (x *JICIANCDHNL) GetPoolId() uint32 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *JICIANCDHNL) GetAEFGEMHJMAK() uint64 {
	if x != nil {
		return x.AEFGEMHJMAK
	}
	return 0
}

func (x *JICIANCDHNL) GetIAEDCABHNJL() uint64 {
	if x != nil {
		return x.IAEDCABHNJL
	}
	return 0
}

func (x *JICIANCDHNL) GetOCMBKJNAGIF() []*EANIBKNOEKJ {
	if x != nil {
		return x.OCMBKJNAGIF
	}
	return nil
}

func (x *JICIANCDHNL) GetPMPIOLCDFEM() uint32 {
	if x != nil {
		return x.PMPIOLCDFEM
	}
	return 0
}

func (x *JICIANCDHNL) GetLOKBDIGDCCN() uint32 {
	if x != nil {
		return x.LOKBDIGDCCN
	}
	return 0
}

func (x *JICIANCDHNL) GetIKFMFKBMHHH() string {
	if x != nil {
		return x.IKFMFKBMHHH
	}
	return ""
}

func (x *JICIANCDHNL) GetLALCLOONFLE() bool {
	if x != nil {
		return x.LALCLOONFLE
	}
	return false
}

func (x *JICIANCDHNL) GetNIMPMPMLJDI() string {
	if x != nil {
		return x.NIMPMPMLJDI
	}
	return ""
}

func (x *JICIANCDHNL) GetIGDJAOELMKM() []*EANIBKNOEKJ {
	if x != nil {
		return x.IGDJAOELMKM
	}
	return nil
}

var File_JICIANCDHNL_proto protoreflect.FileDescriptor

var file_JICIANCDHNL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x49, 0x43, 0x49, 0x41, 0x4e, 0x43, 0x44, 0x48, 0x4e, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x41, 0x4e, 0x49, 0x42, 0x4b, 0x4e, 0x4f, 0x45, 0x4b, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x49, 0x4c, 0x4d, 0x50, 0x4e, 0x4a, 0x44,
	0x43, 0x48, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x05, 0x0a, 0x0b, 0x4a, 0x49,
	0x43, 0x49, 0x41, 0x4e, 0x43, 0x44, 0x48, 0x4e, 0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x46, 0x50,
	0x4d, 0x50, 0x4c, 0x4c, 0x48, 0x47, 0x45, 0x50, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x45, 0x41, 0x4e, 0x49, 0x42, 0x4b, 0x4e, 0x4f, 0x45, 0x4b, 0x4a, 0x52, 0x0b, 0x4c, 0x46,
	0x50, 0x4d, 0x50, 0x4c, 0x4c, 0x48, 0x47, 0x45, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4c, 0x4a,
	0x47, 0x43, 0x47, 0x4d, 0x44, 0x4b, 0x46, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x49, 0x4c, 0x4a, 0x47, 0x43, 0x47, 0x4d, 0x44, 0x4b, 0x46, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x49, 0x48, 0x43, 0x4b, 0x47, 0x4d, 0x47, 0x43, 0x4e, 0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x49, 0x48, 0x43, 0x4b, 0x47, 0x4d, 0x47, 0x43, 0x4e, 0x49, 0x12, 0x20, 0x0a,
	0x0b, 0x49, 0x44, 0x4f, 0x4a, 0x45, 0x4d, 0x4a, 0x4b, 0x4f, 0x43, 0x47, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x49, 0x44, 0x4f, 0x4a, 0x45, 0x4d, 0x4a, 0x4b, 0x4f, 0x43, 0x47, 0x12,
	0x20, 0x0a, 0x0b, 0x42, 0x48, 0x4e, 0x50, 0x45, 0x47, 0x48, 0x48, 0x45, 0x49, 0x4a, 0x18, 0x27,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x48, 0x4e, 0x50, 0x45, 0x47, 0x48, 0x48, 0x45, 0x49,
	0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x4f, 0x45, 0x4e, 0x45, 0x4b, 0x43, 0x42, 0x4d, 0x4e,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x49, 0x4f, 0x45, 0x4e, 0x45, 0x4b, 0x43,
	0x42, 0x4d, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x4e, 0x4a, 0x4f, 0x4e, 0x44, 0x47, 0x4e, 0x41,
	0x4f, 0x43, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x49, 0x4c, 0x4d, 0x50,
	0x4e, 0x4a, 0x44, 0x43, 0x48, 0x50, 0x52, 0x0b, 0x4e, 0x4e, 0x4a, 0x4f, 0x4e, 0x44, 0x47, 0x4e,
	0x41, 0x4f, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x49, 0x48, 0x47, 0x42, 0x49, 0x42, 0x4b, 0x46,
	0x50, 0x44, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x41, 0x4e, 0x49, 0x42,
	0x4b, 0x4e, 0x4f, 0x45, 0x4b, 0x4a, 0x52, 0x0b, 0x4d, 0x49, 0x48, 0x47, 0x42, 0x49, 0x42, 0x4b,
	0x46, 0x50, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x45, 0x46, 0x47, 0x45, 0x4d, 0x48, 0x4a, 0x4d, 0x41, 0x4b, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x41, 0x45, 0x46, 0x47, 0x45, 0x4d, 0x48, 0x4a, 0x4d, 0x41, 0x4b, 0x12, 0x20,
	0x0a, 0x0b, 0x49, 0x41, 0x45, 0x44, 0x43, 0x41, 0x42, 0x48, 0x4e, 0x4a, 0x4c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x49, 0x41, 0x45, 0x44, 0x43, 0x41, 0x42, 0x48, 0x4e, 0x4a, 0x4c,
	0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x43, 0x4d, 0x42, 0x4b, 0x4a, 0x4e, 0x41, 0x47, 0x49, 0x46, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x41, 0x4e, 0x49, 0x42, 0x4b, 0x4e, 0x4f,
	0x45, 0x4b, 0x4a, 0x52, 0x0b, 0x4f, 0x43, 0x4d, 0x42, 0x4b, 0x4a, 0x4e, 0x41, 0x47, 0x49, 0x46,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4d, 0x50, 0x49, 0x4f, 0x4c, 0x43, 0x44, 0x46, 0x45, 0x4d, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4d, 0x50, 0x49, 0x4f, 0x4c, 0x43, 0x44, 0x46,
	0x45, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x4b, 0x42, 0x44, 0x49, 0x47, 0x44, 0x43, 0x43,
	0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4f, 0x4b, 0x42, 0x44, 0x49, 0x47,
	0x44, 0x43, 0x43, 0x4e, 0x12, 0x21, 0x0a, 0x0b, 0x49, 0x4b, 0x46, 0x4d, 0x46, 0x4b, 0x42, 0x4d,
	0x48, 0x48, 0x48, 0x18, 0x83, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x4b, 0x46, 0x4d,
	0x46, 0x4b, 0x42, 0x4d, 0x48, 0x48, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x41, 0x4c, 0x43, 0x4c,
	0x4f, 0x4f, 0x4e, 0x46, 0x4c, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x41,
	0x4c, 0x43, 0x4c, 0x4f, 0x4f, 0x4e, 0x46, 0x4c, 0x45, 0x12, 0x21, 0x0a, 0x0b, 0x4e, 0x49, 0x4d,
	0x50, 0x4d, 0x50, 0x4d, 0x4c, 0x4a, 0x44, 0x49, 0x18, 0xb4, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4e, 0x49, 0x4d, 0x50, 0x4d, 0x50, 0x4d, 0x4c, 0x4a, 0x44, 0x49, 0x12, 0x2e, 0x0a, 0x0b,
	0x49, 0x47, 0x44, 0x4a, 0x41, 0x4f, 0x45, 0x4c, 0x4d, 0x4b, 0x4d, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x41, 0x4e, 0x49, 0x42, 0x4b, 0x4e, 0x4f, 0x45, 0x4b, 0x4a, 0x52,
	0x0b, 0x49, 0x47, 0x44, 0x4a, 0x41, 0x4f, 0x45, 0x4c, 0x4d, 0x4b, 0x4d, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JICIANCDHNL_proto_rawDescOnce sync.Once
	file_JICIANCDHNL_proto_rawDescData = file_JICIANCDHNL_proto_rawDesc
)

func file_JICIANCDHNL_proto_rawDescGZIP() []byte {
	file_JICIANCDHNL_proto_rawDescOnce.Do(func() {
		file_JICIANCDHNL_proto_rawDescData = protoimpl.X.CompressGZIP(file_JICIANCDHNL_proto_rawDescData)
	})
	return file_JICIANCDHNL_proto_rawDescData
}

var file_JICIANCDHNL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JICIANCDHNL_proto_goTypes = []interface{}{
	(*JICIANCDHNL)(nil), // 0: JICIANCDHNL
	(*EANIBKNOEKJ)(nil), // 1: EANIBKNOEKJ
	(*NILMPNJDCHP)(nil), // 2: NILMPNJDCHP
}
var file_JICIANCDHNL_proto_depIdxs = []int32{
	1, // 0: JICIANCDHNL.LFPMPLLHGEP:type_name -> EANIBKNOEKJ
	2, // 1: JICIANCDHNL.NNJONDGNAOC:type_name -> NILMPNJDCHP
	1, // 2: JICIANCDHNL.MIHGBIBKFPD:type_name -> EANIBKNOEKJ
	1, // 3: JICIANCDHNL.OCMBKJNAGIF:type_name -> EANIBKNOEKJ
	1, // 4: JICIANCDHNL.IGDJAOELMKM:type_name -> EANIBKNOEKJ
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_JICIANCDHNL_proto_init() }
func file_JICIANCDHNL_proto_init() {
	if File_JICIANCDHNL_proto != nil {
		return
	}
	file_EANIBKNOEKJ_proto_init()
	file_NILMPNJDCHP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JICIANCDHNL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JICIANCDHNL); i {
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
			RawDescriptor: file_JICIANCDHNL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JICIANCDHNL_proto_goTypes,
		DependencyIndexes: file_JICIANCDHNL_proto_depIdxs,
		MessageInfos:      file_JICIANCDHNL_proto_msgTypes,
	}.Build()
	File_JICIANCDHNL_proto = out.File
	file_JICIANCDHNL_proto_rawDesc = nil
	file_JICIANCDHNL_proto_goTypes = nil
	file_JICIANCDHNL_proto_depIdxs = nil
}
