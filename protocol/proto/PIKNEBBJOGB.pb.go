// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PIKNEBBJOGB.proto

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

type PIKNEBBJOGB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BCFPEDCABCM *AEAMBLOFMPI `protobuf:"bytes,1,opt,name=BCFPEDCABCM,proto3" json:"BCFPEDCABCM,omitempty"`
	LLHAMMHBNFO *IOGBIDFKHPM `protobuf:"bytes,13,opt,name=LLHAMMHBNFO,proto3" json:"LLHAMMHBNFO,omitempty"`
	HMBOHBOFLCD *IOGBIDFKHPM `protobuf:"bytes,5,opt,name=HMBOHBOFLCD,proto3" json:"HMBOHBOFLCD,omitempty"`
	GKEIKEBJODO *BBCAPKMGAEA `protobuf:"bytes,7,opt,name=GKEIKEBJODO,proto3" json:"GKEIKEBJODO,omitempty"`
	NNOLDICJKJE *CAIAJMHEBPE `protobuf:"bytes,11,opt,name=NNOLDICJKJE,proto3" json:"NNOLDICJKJE,omitempty"`
	NLNLEFFJPKB *MIHPPHICNIH `protobuf:"bytes,12,opt,name=NLNLEFFJPKB,proto3" json:"NLNLEFFJPKB,omitempty"`
	LFJAPKMPEIB *LPDOHKMGBFM `protobuf:"bytes,14,opt,name=LFJAPKMPEIB,proto3" json:"LFJAPKMPEIB,omitempty"`
	ONDNKBOENMO uint32       `protobuf:"varint,10,opt,name=ONDNKBOENMO,proto3" json:"ONDNKBOENMO,omitempty"`
	MPMANDBOFBC uint32       `protobuf:"varint,1519,opt,name=MPMANDBOFBC,proto3" json:"MPMANDBOFBC,omitempty"`
	DMEKIKJAKCG *CBGNJFGBGEE `protobuf:"bytes,871,opt,name=DMEKIKJAKCG,proto3" json:"DMEKIKJAKCG,omitempty"`
	NMDBKCPCBII *ONJHNHIKEOC `protobuf:"bytes,1686,opt,name=NMDBKCPCBII,proto3" json:"NMDBKCPCBII,omitempty"`
	EDOFKFHEFNG uint32       `protobuf:"varint,1038,opt,name=EDOFKFHEFNG,proto3" json:"EDOFKFHEFNG,omitempty"`
}

func (x *PIKNEBBJOGB) Reset() {
	*x = PIKNEBBJOGB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PIKNEBBJOGB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIKNEBBJOGB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIKNEBBJOGB) ProtoMessage() {}

func (x *PIKNEBBJOGB) ProtoReflect() protoreflect.Message {
	mi := &file_PIKNEBBJOGB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIKNEBBJOGB.ProtoReflect.Descriptor instead.
func (*PIKNEBBJOGB) Descriptor() ([]byte, []int) {
	return file_PIKNEBBJOGB_proto_rawDescGZIP(), []int{0}
}

func (x *PIKNEBBJOGB) GetBCFPEDCABCM() *AEAMBLOFMPI {
	if x != nil {
		return x.BCFPEDCABCM
	}
	return nil
}

func (x *PIKNEBBJOGB) GetLLHAMMHBNFO() *IOGBIDFKHPM {
	if x != nil {
		return x.LLHAMMHBNFO
	}
	return nil
}

func (x *PIKNEBBJOGB) GetHMBOHBOFLCD() *IOGBIDFKHPM {
	if x != nil {
		return x.HMBOHBOFLCD
	}
	return nil
}

func (x *PIKNEBBJOGB) GetGKEIKEBJODO() *BBCAPKMGAEA {
	if x != nil {
		return x.GKEIKEBJODO
	}
	return nil
}

func (x *PIKNEBBJOGB) GetNNOLDICJKJE() *CAIAJMHEBPE {
	if x != nil {
		return x.NNOLDICJKJE
	}
	return nil
}

func (x *PIKNEBBJOGB) GetNLNLEFFJPKB() *MIHPPHICNIH {
	if x != nil {
		return x.NLNLEFFJPKB
	}
	return nil
}

func (x *PIKNEBBJOGB) GetLFJAPKMPEIB() *LPDOHKMGBFM {
	if x != nil {
		return x.LFJAPKMPEIB
	}
	return nil
}

func (x *PIKNEBBJOGB) GetONDNKBOENMO() uint32 {
	if x != nil {
		return x.ONDNKBOENMO
	}
	return 0
}

func (x *PIKNEBBJOGB) GetMPMANDBOFBC() uint32 {
	if x != nil {
		return x.MPMANDBOFBC
	}
	return 0
}

func (x *PIKNEBBJOGB) GetDMEKIKJAKCG() *CBGNJFGBGEE {
	if x != nil {
		return x.DMEKIKJAKCG
	}
	return nil
}

func (x *PIKNEBBJOGB) GetNMDBKCPCBII() *ONJHNHIKEOC {
	if x != nil {
		return x.NMDBKCPCBII
	}
	return nil
}

func (x *PIKNEBBJOGB) GetEDOFKFHEFNG() uint32 {
	if x != nil {
		return x.EDOFKFHEFNG
	}
	return 0
}

var File_PIKNEBBJOGB_proto protoreflect.FileDescriptor

var file_PIKNEBBJOGB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x49, 0x4b, 0x4e, 0x45, 0x42, 0x42, 0x4a, 0x4f, 0x47, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x50, 0x44, 0x4f, 0x48, 0x4b, 0x4d, 0x47, 0x42, 0x46, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x42, 0x43, 0x41, 0x50, 0x4b, 0x4d, 0x47,
	0x41, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x49, 0x48, 0x50, 0x50,
	0x48, 0x49, 0x43, 0x4e, 0x49, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4f,
	0x47, 0x42, 0x49, 0x44, 0x46, 0x4b, 0x48, 0x50, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x43, 0x42, 0x47, 0x4e, 0x4a, 0x46, 0x47, 0x42, 0x47, 0x45, 0x45, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4e, 0x4a, 0x48, 0x4e, 0x48, 0x49, 0x4b, 0x45, 0x4f, 0x43, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x41, 0x49, 0x41, 0x4a, 0x4d, 0x48, 0x45, 0x42,
	0x50, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x45, 0x41, 0x4d, 0x42, 0x4c,
	0x4f, 0x46, 0x4d, 0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x04, 0x0a, 0x0b,
	0x50, 0x49, 0x4b, 0x4e, 0x45, 0x42, 0x42, 0x4a, 0x4f, 0x47, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x42,
	0x43, 0x46, 0x50, 0x45, 0x44, 0x43, 0x41, 0x42, 0x43, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x41, 0x45, 0x41, 0x4d, 0x42, 0x4c, 0x4f, 0x46, 0x4d, 0x50, 0x49, 0x52, 0x0b,
	0x42, 0x43, 0x46, 0x50, 0x45, 0x44, 0x43, 0x41, 0x42, 0x43, 0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x4c,
	0x4c, 0x48, 0x41, 0x4d, 0x4d, 0x48, 0x42, 0x4e, 0x46, 0x4f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x4f, 0x47, 0x42, 0x49, 0x44, 0x46, 0x4b, 0x48, 0x50, 0x4d, 0x52, 0x0b,
	0x4c, 0x4c, 0x48, 0x41, 0x4d, 0x4d, 0x48, 0x42, 0x4e, 0x46, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x48,
	0x4d, 0x42, 0x4f, 0x48, 0x42, 0x4f, 0x46, 0x4c, 0x43, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x4f, 0x47, 0x42, 0x49, 0x44, 0x46, 0x4b, 0x48, 0x50, 0x4d, 0x52, 0x0b,
	0x48, 0x4d, 0x42, 0x4f, 0x48, 0x42, 0x4f, 0x46, 0x4c, 0x43, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x4b, 0x45, 0x49, 0x4b, 0x45, 0x42, 0x4a, 0x4f, 0x44, 0x4f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x42, 0x42, 0x43, 0x41, 0x50, 0x4b, 0x4d, 0x47, 0x41, 0x45, 0x41, 0x52, 0x0b,
	0x47, 0x4b, 0x45, 0x49, 0x4b, 0x45, 0x42, 0x4a, 0x4f, 0x44, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x4e,
	0x4e, 0x4f, 0x4c, 0x44, 0x49, 0x43, 0x4a, 0x4b, 0x4a, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x43, 0x41, 0x49, 0x41, 0x4a, 0x4d, 0x48, 0x45, 0x42, 0x50, 0x45, 0x52, 0x0b,
	0x4e, 0x4e, 0x4f, 0x4c, 0x44, 0x49, 0x43, 0x4a, 0x4b, 0x4a, 0x45, 0x12, 0x2e, 0x0a, 0x0b, 0x4e,
	0x4c, 0x4e, 0x4c, 0x45, 0x46, 0x46, 0x4a, 0x50, 0x4b, 0x42, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4d, 0x49, 0x48, 0x50, 0x50, 0x48, 0x49, 0x43, 0x4e, 0x49, 0x48, 0x52, 0x0b,
	0x4e, 0x4c, 0x4e, 0x4c, 0x45, 0x46, 0x46, 0x4a, 0x50, 0x4b, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x4c,
	0x46, 0x4a, 0x41, 0x50, 0x4b, 0x4d, 0x50, 0x45, 0x49, 0x42, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x44, 0x4f, 0x48, 0x4b, 0x4d, 0x47, 0x42, 0x46, 0x4d, 0x52, 0x0b,
	0x4c, 0x46, 0x4a, 0x41, 0x50, 0x4b, 0x4d, 0x50, 0x45, 0x49, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x4e, 0x44, 0x4e, 0x4b, 0x42, 0x4f, 0x45, 0x4e, 0x4d, 0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4f, 0x4e, 0x44, 0x4e, 0x4b, 0x42, 0x4f, 0x45, 0x4e, 0x4d, 0x4f, 0x12, 0x21, 0x0a,
	0x0b, 0x4d, 0x50, 0x4d, 0x41, 0x4e, 0x44, 0x42, 0x4f, 0x46, 0x42, 0x43, 0x18, 0xef, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x50, 0x4d, 0x41, 0x4e, 0x44, 0x42, 0x4f, 0x46, 0x42, 0x43,
	0x12, 0x2f, 0x0a, 0x0b, 0x44, 0x4d, 0x45, 0x4b, 0x49, 0x4b, 0x4a, 0x41, 0x4b, 0x43, 0x47, 0x18,
	0xe7, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x42, 0x47, 0x4e, 0x4a, 0x46, 0x47,
	0x42, 0x47, 0x45, 0x45, 0x52, 0x0b, 0x44, 0x4d, 0x45, 0x4b, 0x49, 0x4b, 0x4a, 0x41, 0x4b, 0x43,
	0x47, 0x12, 0x2f, 0x0a, 0x0b, 0x4e, 0x4d, 0x44, 0x42, 0x4b, 0x43, 0x50, 0x43, 0x42, 0x49, 0x49,
	0x18, 0x96, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4e, 0x4a, 0x48, 0x4e, 0x48,
	0x49, 0x4b, 0x45, 0x4f, 0x43, 0x52, 0x0b, 0x4e, 0x4d, 0x44, 0x42, 0x4b, 0x43, 0x50, 0x43, 0x42,
	0x49, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x45, 0x44, 0x4f, 0x46, 0x4b, 0x46, 0x48, 0x45, 0x46, 0x4e,
	0x47, 0x18, 0x8e, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x44, 0x4f, 0x46, 0x4b, 0x46,
	0x48, 0x45, 0x46, 0x4e, 0x47, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PIKNEBBJOGB_proto_rawDescOnce sync.Once
	file_PIKNEBBJOGB_proto_rawDescData = file_PIKNEBBJOGB_proto_rawDesc
)

func file_PIKNEBBJOGB_proto_rawDescGZIP() []byte {
	file_PIKNEBBJOGB_proto_rawDescOnce.Do(func() {
		file_PIKNEBBJOGB_proto_rawDescData = protoimpl.X.CompressGZIP(file_PIKNEBBJOGB_proto_rawDescData)
	})
	return file_PIKNEBBJOGB_proto_rawDescData
}

var file_PIKNEBBJOGB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PIKNEBBJOGB_proto_goTypes = []interface{}{
	(*PIKNEBBJOGB)(nil), // 0: PIKNEBBJOGB
	(*AEAMBLOFMPI)(nil), // 1: AEAMBLOFMPI
	(*IOGBIDFKHPM)(nil), // 2: IOGBIDFKHPM
	(*BBCAPKMGAEA)(nil), // 3: BBCAPKMGAEA
	(*CAIAJMHEBPE)(nil), // 4: CAIAJMHEBPE
	(*MIHPPHICNIH)(nil), // 5: MIHPPHICNIH
	(*LPDOHKMGBFM)(nil), // 6: LPDOHKMGBFM
	(*CBGNJFGBGEE)(nil), // 7: CBGNJFGBGEE
	(*ONJHNHIKEOC)(nil), // 8: ONJHNHIKEOC
}
var file_PIKNEBBJOGB_proto_depIdxs = []int32{
	1, // 0: PIKNEBBJOGB.BCFPEDCABCM:type_name -> AEAMBLOFMPI
	2, // 1: PIKNEBBJOGB.LLHAMMHBNFO:type_name -> IOGBIDFKHPM
	2, // 2: PIKNEBBJOGB.HMBOHBOFLCD:type_name -> IOGBIDFKHPM
	3, // 3: PIKNEBBJOGB.GKEIKEBJODO:type_name -> BBCAPKMGAEA
	4, // 4: PIKNEBBJOGB.NNOLDICJKJE:type_name -> CAIAJMHEBPE
	5, // 5: PIKNEBBJOGB.NLNLEFFJPKB:type_name -> MIHPPHICNIH
	6, // 6: PIKNEBBJOGB.LFJAPKMPEIB:type_name -> LPDOHKMGBFM
	7, // 7: PIKNEBBJOGB.DMEKIKJAKCG:type_name -> CBGNJFGBGEE
	8, // 8: PIKNEBBJOGB.NMDBKCPCBII:type_name -> ONJHNHIKEOC
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_PIKNEBBJOGB_proto_init() }
func file_PIKNEBBJOGB_proto_init() {
	if File_PIKNEBBJOGB_proto != nil {
		return
	}
	file_LPDOHKMGBFM_proto_init()
	file_BBCAPKMGAEA_proto_init()
	file_MIHPPHICNIH_proto_init()
	file_IOGBIDFKHPM_proto_init()
	file_CBGNJFGBGEE_proto_init()
	file_ONJHNHIKEOC_proto_init()
	file_CAIAJMHEBPE_proto_init()
	file_AEAMBLOFMPI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PIKNEBBJOGB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIKNEBBJOGB); i {
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
			RawDescriptor: file_PIKNEBBJOGB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PIKNEBBJOGB_proto_goTypes,
		DependencyIndexes: file_PIKNEBBJOGB_proto_depIdxs,
		MessageInfos:      file_PIKNEBBJOGB_proto_msgTypes,
	}.Build()
	File_PIKNEBBJOGB_proto = out.File
	file_PIKNEBBJOGB_proto_rawDesc = nil
	file_PIKNEBBJOGB_proto_goTypes = nil
	file_PIKNEBBJOGB_proto_depIdxs = nil
}