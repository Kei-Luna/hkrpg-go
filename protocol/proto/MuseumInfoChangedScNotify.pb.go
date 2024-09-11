// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MuseumInfoChangedScNotify.proto

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

type MuseumInfoChangedScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       uint32         `protobuf:"varint,7,opt,name=level,proto3" json:"level,omitempty"`
	PIBEFEFGMNN *LPNAAPGCINF   `protobuf:"bytes,6,opt,name=PIBEFEFGMNN,proto3" json:"PIBEFEFGMNN,omitempty"`
	EGFMHNCAKML []*HEKMJICGNOK `protobuf:"bytes,2,rep,name=EGFMHNCAKML,proto3" json:"EGFMHNCAKML,omitempty"`
	NMFNDFLLOIM uint32         `protobuf:"varint,4,opt,name=NMFNDFLLOIM,proto3" json:"NMFNDFLLOIM,omitempty"`
	LADPOGLHNJJ []uint32       `protobuf:"varint,8,rep,packed,name=LADPOGLHNJJ,proto3" json:"LADPOGLHNJJ,omitempty"`
	BFKBLKNBHFB uint32         `protobuf:"varint,5,opt,name=BFKBLKNBHFB,proto3" json:"BFKBLKNBHFB,omitempty"`
	NBGPOPOKELO uint32         `protobuf:"varint,9,opt,name=NBGPOPOKELO,proto3" json:"NBGPOPOKELO,omitempty"`
	Exp         uint32         `protobuf:"varint,15,opt,name=exp,proto3" json:"exp,omitempty"`
	JPIKKFGJMCE uint32         `protobuf:"varint,11,opt,name=JPIKKFGJMCE,proto3" json:"JPIKKFGJMCE,omitempty"`
	HHDFFEDNLGM *HMJGECKCKMM   `protobuf:"bytes,12,opt,name=HHDFFEDNLGM,proto3" json:"HHDFFEDNLGM,omitempty"`
	FLNOGELFODP uint32         `protobuf:"varint,14,opt,name=FLNOGELFODP,proto3" json:"FLNOGELFODP,omitempty"`
	CPKMIAOLGJK []uint32       `protobuf:"varint,3,rep,packed,name=CPKMIAOLGJK,proto3" json:"CPKMIAOLGJK,omitempty"`
	JOFNJMPGNHB []*KOPBMGBKABE `protobuf:"bytes,10,rep,name=JOFNJMPGNHB,proto3" json:"JOFNJMPGNHB,omitempty"`
}

func (x *MuseumInfoChangedScNotify) Reset() {
	*x = MuseumInfoChangedScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MuseumInfoChangedScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuseumInfoChangedScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuseumInfoChangedScNotify) ProtoMessage() {}

func (x *MuseumInfoChangedScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MuseumInfoChangedScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuseumInfoChangedScNotify.ProtoReflect.Descriptor instead.
func (*MuseumInfoChangedScNotify) Descriptor() ([]byte, []int) {
	return file_MuseumInfoChangedScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MuseumInfoChangedScNotify) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetPIBEFEFGMNN() *LPNAAPGCINF {
	if x != nil {
		return x.PIBEFEFGMNN
	}
	return nil
}

func (x *MuseumInfoChangedScNotify) GetEGFMHNCAKML() []*HEKMJICGNOK {
	if x != nil {
		return x.EGFMHNCAKML
	}
	return nil
}

func (x *MuseumInfoChangedScNotify) GetNMFNDFLLOIM() uint32 {
	if x != nil {
		return x.NMFNDFLLOIM
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetLADPOGLHNJJ() []uint32 {
	if x != nil {
		return x.LADPOGLHNJJ
	}
	return nil
}

func (x *MuseumInfoChangedScNotify) GetBFKBLKNBHFB() uint32 {
	if x != nil {
		return x.BFKBLKNBHFB
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetNBGPOPOKELO() uint32 {
	if x != nil {
		return x.NBGPOPOKELO
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetJPIKKFGJMCE() uint32 {
	if x != nil {
		return x.JPIKKFGJMCE
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetHHDFFEDNLGM() *HMJGECKCKMM {
	if x != nil {
		return x.HHDFFEDNLGM
	}
	return nil
}

func (x *MuseumInfoChangedScNotify) GetFLNOGELFODP() uint32 {
	if x != nil {
		return x.FLNOGELFODP
	}
	return 0
}

func (x *MuseumInfoChangedScNotify) GetCPKMIAOLGJK() []uint32 {
	if x != nil {
		return x.CPKMIAOLGJK
	}
	return nil
}

func (x *MuseumInfoChangedScNotify) GetJOFNJMPGNHB() []*KOPBMGBKABE {
	if x != nil {
		return x.JOFNJMPGNHB
	}
	return nil
}

var File_MuseumInfoChangedScNotify_proto protoreflect.FileDescriptor

var file_MuseumInfoChangedScNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x48, 0x4d, 0x4a, 0x47, 0x45, 0x43, 0x4b, 0x43, 0x4b, 0x4d, 0x4d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x50, 0x4e, 0x41, 0x41, 0x50, 0x47, 0x43, 0x49, 0x4e,
	0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x45, 0x4b, 0x4d, 0x4a, 0x49, 0x43,
	0x47, 0x4e, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4f, 0x50, 0x42,
	0x4d, 0x47, 0x42, 0x4b, 0x41, 0x42, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x03,
	0x0a, 0x19, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x49, 0x42, 0x45, 0x46, 0x45, 0x46, 0x47, 0x4d, 0x4e, 0x4e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x4e, 0x41, 0x41, 0x50, 0x47,
	0x43, 0x49, 0x4e, 0x46, 0x52, 0x0b, 0x50, 0x49, 0x42, 0x45, 0x46, 0x45, 0x46, 0x47, 0x4d, 0x4e,
	0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x47, 0x46, 0x4d, 0x48, 0x4e, 0x43, 0x41, 0x4b, 0x4d, 0x4c,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x45, 0x4b, 0x4d, 0x4a, 0x49, 0x43,
	0x47, 0x4e, 0x4f, 0x4b, 0x52, 0x0b, 0x45, 0x47, 0x46, 0x4d, 0x48, 0x4e, 0x43, 0x41, 0x4b, 0x4d,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4d, 0x46, 0x4e, 0x44, 0x46, 0x4c, 0x4c, 0x4f, 0x49, 0x4d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4d, 0x46, 0x4e, 0x44, 0x46, 0x4c, 0x4c,
	0x4f, 0x49, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x41, 0x44, 0x50, 0x4f, 0x47, 0x4c, 0x48, 0x4e,
	0x4a, 0x4a, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x41, 0x44, 0x50, 0x4f, 0x47,
	0x4c, 0x48, 0x4e, 0x4a, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46, 0x4b, 0x42, 0x4c, 0x4b, 0x4e,
	0x42, 0x48, 0x46, 0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x46, 0x4b, 0x42,
	0x4c, 0x4b, 0x4e, 0x42, 0x48, 0x46, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x42, 0x47, 0x50, 0x4f,
	0x50, 0x4f, 0x4b, 0x45, 0x4c, 0x4f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42,
	0x47, 0x50, 0x4f, 0x50, 0x4f, 0x4b, 0x45, 0x4c, 0x4f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x50, 0x49, 0x4b, 0x4b, 0x46, 0x47, 0x4a, 0x4d, 0x43, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x50, 0x49, 0x4b, 0x4b, 0x46, 0x47, 0x4a, 0x4d, 0x43, 0x45, 0x12, 0x2e, 0x0a,
	0x0b, 0x48, 0x48, 0x44, 0x46, 0x46, 0x45, 0x44, 0x4e, 0x4c, 0x47, 0x4d, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4d, 0x4a, 0x47, 0x45, 0x43, 0x4b, 0x43, 0x4b, 0x4d, 0x4d,
	0x52, 0x0b, 0x48, 0x48, 0x44, 0x46, 0x46, 0x45, 0x44, 0x4e, 0x4c, 0x47, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x46, 0x4c, 0x4e, 0x4f, 0x47, 0x45, 0x4c, 0x46, 0x4f, 0x44, 0x50, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4c, 0x4e, 0x4f, 0x47, 0x45, 0x4c, 0x46, 0x4f, 0x44, 0x50, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x50, 0x4b, 0x4d, 0x49, 0x41, 0x4f, 0x4c, 0x47, 0x4a, 0x4b, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x4b, 0x4d, 0x49, 0x41, 0x4f, 0x4c, 0x47, 0x4a,
	0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x4f, 0x46, 0x4e, 0x4a, 0x4d, 0x50, 0x47, 0x4e, 0x48, 0x42,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4f, 0x50, 0x42, 0x4d, 0x47, 0x42,
	0x4b, 0x41, 0x42, 0x45, 0x52, 0x0b, 0x4a, 0x4f, 0x46, 0x4e, 0x4a, 0x4d, 0x50, 0x47, 0x4e, 0x48,
	0x42, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MuseumInfoChangedScNotify_proto_rawDescOnce sync.Once
	file_MuseumInfoChangedScNotify_proto_rawDescData = file_MuseumInfoChangedScNotify_proto_rawDesc
)

func file_MuseumInfoChangedScNotify_proto_rawDescGZIP() []byte {
	file_MuseumInfoChangedScNotify_proto_rawDescOnce.Do(func() {
		file_MuseumInfoChangedScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MuseumInfoChangedScNotify_proto_rawDescData)
	})
	return file_MuseumInfoChangedScNotify_proto_rawDescData
}

var file_MuseumInfoChangedScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MuseumInfoChangedScNotify_proto_goTypes = []interface{}{
	(*MuseumInfoChangedScNotify)(nil), // 0: MuseumInfoChangedScNotify
	(*LPNAAPGCINF)(nil),               // 1: LPNAAPGCINF
	(*HEKMJICGNOK)(nil),               // 2: HEKMJICGNOK
	(*HMJGECKCKMM)(nil),               // 3: HMJGECKCKMM
	(*KOPBMGBKABE)(nil),               // 4: KOPBMGBKABE
}
var file_MuseumInfoChangedScNotify_proto_depIdxs = []int32{
	1, // 0: MuseumInfoChangedScNotify.PIBEFEFGMNN:type_name -> LPNAAPGCINF
	2, // 1: MuseumInfoChangedScNotify.EGFMHNCAKML:type_name -> HEKMJICGNOK
	3, // 2: MuseumInfoChangedScNotify.HHDFFEDNLGM:type_name -> HMJGECKCKMM
	4, // 3: MuseumInfoChangedScNotify.JOFNJMPGNHB:type_name -> KOPBMGBKABE
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_MuseumInfoChangedScNotify_proto_init() }
func file_MuseumInfoChangedScNotify_proto_init() {
	if File_MuseumInfoChangedScNotify_proto != nil {
		return
	}
	file_HMJGECKCKMM_proto_init()
	file_LPNAAPGCINF_proto_init()
	file_HEKMJICGNOK_proto_init()
	file_KOPBMGBKABE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MuseumInfoChangedScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuseumInfoChangedScNotify); i {
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
			RawDescriptor: file_MuseumInfoChangedScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MuseumInfoChangedScNotify_proto_goTypes,
		DependencyIndexes: file_MuseumInfoChangedScNotify_proto_depIdxs,
		MessageInfos:      file_MuseumInfoChangedScNotify_proto_msgTypes,
	}.Build()
	File_MuseumInfoChangedScNotify_proto = out.File
	file_MuseumInfoChangedScNotify_proto_rawDesc = nil
	file_MuseumInfoChangedScNotify_proto_goTypes = nil
	file_MuseumInfoChangedScNotify_proto_depIdxs = nil
}
