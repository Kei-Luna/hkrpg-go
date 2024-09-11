// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMonopolyInfoScRsp.proto

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

type GetMonopolyInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OMDFOLOFPDB   *LJMJPODBCAE `protobuf:"bytes,4,opt,name=OMDFOLOFPDB,proto3" json:"OMDFOLOFPDB,omitempty"`
	RoomMap       *OOPNGMHCNNN `protobuf:"bytes,14,opt,name=room_map,json=roomMap,proto3" json:"room_map,omitempty"`
	AJOKJDEGEPC   []uint32     `protobuf:"varint,15,rep,packed,name=AJOKJDEGEPC,proto3" json:"AJOKJDEGEPC,omitempty"`
	CAIMMNMCPCJ   *MGGHEHLPFMH `protobuf:"bytes,13,opt,name=CAIMMNMCPCJ,proto3" json:"CAIMMNMCPCJ,omitempty"`
	RogueBuffInfo *CKFFBACIGPG `protobuf:"bytes,9,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	IKAFFGFKBFN   *DEFMCECJJBI `protobuf:"bytes,11,opt,name=IKAFFGFKBFN,proto3" json:"IKAFFGFKBFN,omitempty"`
	JKKGEBEMGOF   *GIODDOIHLCN `protobuf:"bytes,12,opt,name=JKKGEBEMGOF,proto3" json:"JKKGEBEMGOF,omitempty"`
	DCPFFMNFMKJ   *JOEPAJDGPHK `protobuf:"bytes,5,opt,name=DCPFFMNFMKJ,proto3" json:"DCPFFMNFMKJ,omitempty"`
	CAJDMFNLNKD   *EONNIEFDOCI `protobuf:"bytes,6,opt,name=CAJDMFNLNKD,proto3" json:"CAJDMFNLNKD,omitempty"`
	Retcode       uint32       `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Stt           *KFBEAPLBEFB `protobuf:"bytes,1,opt,name=stt,proto3" json:"stt,omitempty"`
	JADMBMOHANO   *JOGIAGPFDML `protobuf:"bytes,3,opt,name=JADMBMOHANO,proto3" json:"JADMBMOHANO,omitempty"`
	JLBDHCCPJDP   *KHDOBJBGPNH `protobuf:"bytes,7,opt,name=JLBDHCCPJDP,proto3" json:"JLBDHCCPJDP,omitempty"`
}

func (x *GetMonopolyInfoScRsp) Reset() {
	*x = GetMonopolyInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonopolyInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonopolyInfoScRsp) ProtoMessage() {}

func (x *GetMonopolyInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonopolyInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetMonopolyInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetMonopolyInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMonopolyInfoScRsp) GetOMDFOLOFPDB() *LJMJPODBCAE {
	if x != nil {
		return x.OMDFOLOFPDB
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRoomMap() *OOPNGMHCNNN {
	if x != nil {
		return x.RoomMap
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetAJOKJDEGEPC() []uint32 {
	if x != nil {
		return x.AJOKJDEGEPC
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetCAIMMNMCPCJ() *MGGHEHLPFMH {
	if x != nil {
		return x.CAIMMNMCPCJ
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRogueBuffInfo() *CKFFBACIGPG {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetIKAFFGFKBFN() *DEFMCECJJBI {
	if x != nil {
		return x.IKAFFGFKBFN
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetJKKGEBEMGOF() *GIODDOIHLCN {
	if x != nil {
		return x.JKKGEBEMGOF
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetDCPFFMNFMKJ() *JOEPAJDGPHK {
	if x != nil {
		return x.DCPFFMNFMKJ
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetCAJDMFNLNKD() *EONNIEFDOCI {
	if x != nil {
		return x.CAJDMFNLNKD
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMonopolyInfoScRsp) GetStt() *KFBEAPLBEFB {
	if x != nil {
		return x.Stt
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetJADMBMOHANO() *JOGIAGPFDML {
	if x != nil {
		return x.JADMBMOHANO
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetJLBDHCCPJDP() *KHDOBJBGPNH {
	if x != nil {
		return x.JLBDHCCPJDP
	}
	return nil
}

var File_GetMonopolyInfoScRsp_proto protoreflect.FileDescriptor

var file_GetMonopolyInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4b,
	0x46, 0x46, 0x42, 0x41, 0x43, 0x49, 0x47, 0x50, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4a, 0x4f, 0x47, 0x49, 0x41, 0x47, 0x50, 0x46, 0x44, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4a, 0x4d, 0x4a, 0x50, 0x4f, 0x44, 0x42, 0x43, 0x41, 0x45, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x49, 0x4f, 0x44, 0x44, 0x4f, 0x49, 0x48, 0x4c,
	0x43, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4f, 0x4e, 0x4e, 0x49, 0x45,
	0x46, 0x44, 0x4f, 0x43, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x46, 0x42,
	0x45, 0x41, 0x50, 0x4c, 0x42, 0x45, 0x46, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4d, 0x47, 0x47, 0x48, 0x45, 0x48, 0x4c, 0x50, 0x46, 0x4d, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4f, 0x4f, 0x50, 0x4e, 0x47, 0x4d, 0x48, 0x43, 0x4e, 0x4e, 0x4e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x48, 0x44, 0x4f, 0x42, 0x4a, 0x42, 0x47, 0x50, 0x4e,
	0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x45, 0x46, 0x4d, 0x43, 0x45, 0x43,
	0x4a, 0x4a, 0x42, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4f, 0x45, 0x50,
	0x41, 0x4a, 0x44, 0x47, 0x50, 0x48, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x04,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4d, 0x44, 0x46, 0x4f, 0x4c,
	0x4f, 0x46, 0x50, 0x44, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4a,
	0x4d, 0x4a, 0x50, 0x4f, 0x44, 0x42, 0x43, 0x41, 0x45, 0x52, 0x0b, 0x4f, 0x4d, 0x44, 0x46, 0x4f,
	0x4c, 0x4f, 0x46, 0x50, 0x44, 0x42, 0x12, 0x27, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4f, 0x50, 0x4e, 0x47,
	0x4d, 0x48, 0x43, 0x4e, 0x4e, 0x4e, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x4a, 0x4f, 0x4b, 0x4a, 0x44, 0x45, 0x47, 0x45, 0x50, 0x43, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4a, 0x4f, 0x4b, 0x4a, 0x44, 0x45, 0x47, 0x45, 0x50,
	0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x41, 0x49, 0x4d, 0x4d, 0x4e, 0x4d, 0x43, 0x50, 0x43, 0x4a,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x47, 0x47, 0x48, 0x45, 0x48, 0x4c,
	0x50, 0x46, 0x4d, 0x48, 0x52, 0x0b, 0x43, 0x41, 0x49, 0x4d, 0x4d, 0x4e, 0x4d, 0x43, 0x50, 0x43,
	0x4a, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4b, 0x46,
	0x46, 0x42, 0x41, 0x43, 0x49, 0x47, 0x50, 0x47, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x4b, 0x41, 0x46, 0x46,
	0x47, 0x46, 0x4b, 0x42, 0x46, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44,
	0x45, 0x46, 0x4d, 0x43, 0x45, 0x43, 0x4a, 0x4a, 0x42, 0x49, 0x52, 0x0b, 0x49, 0x4b, 0x41, 0x46,
	0x46, 0x47, 0x46, 0x4b, 0x42, 0x46, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x4b, 0x4b, 0x47, 0x45,
	0x42, 0x45, 0x4d, 0x47, 0x4f, 0x46, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47,
	0x49, 0x4f, 0x44, 0x44, 0x4f, 0x49, 0x48, 0x4c, 0x43, 0x4e, 0x52, 0x0b, 0x4a, 0x4b, 0x4b, 0x47,
	0x45, 0x42, 0x45, 0x4d, 0x47, 0x4f, 0x46, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x43, 0x50, 0x46, 0x46,
	0x4d, 0x4e, 0x46, 0x4d, 0x4b, 0x4a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a,
	0x4f, 0x45, 0x50, 0x41, 0x4a, 0x44, 0x47, 0x50, 0x48, 0x4b, 0x52, 0x0b, 0x44, 0x43, 0x50, 0x46,
	0x46, 0x4d, 0x4e, 0x46, 0x4d, 0x4b, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x41, 0x4a, 0x44, 0x4d,
	0x46, 0x4e, 0x4c, 0x4e, 0x4b, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x4f, 0x4e, 0x4e, 0x49, 0x45, 0x46, 0x44, 0x4f, 0x43, 0x49, 0x52, 0x0b, 0x43, 0x41, 0x4a, 0x44,
	0x4d, 0x46, 0x4e, 0x4c, 0x4e, 0x4b, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x74, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4b, 0x46, 0x42, 0x45, 0x41, 0x50, 0x4c, 0x42, 0x45, 0x46, 0x42, 0x52, 0x03, 0x73, 0x74,
	0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x41, 0x44, 0x4d, 0x42, 0x4d, 0x4f, 0x48, 0x41, 0x4e, 0x4f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4f, 0x47, 0x49, 0x41, 0x47, 0x50,
	0x46, 0x44, 0x4d, 0x4c, 0x52, 0x0b, 0x4a, 0x41, 0x44, 0x4d, 0x42, 0x4d, 0x4f, 0x48, 0x41, 0x4e,
	0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x4c, 0x42, 0x44, 0x48, 0x43, 0x43, 0x50, 0x4a, 0x44, 0x50,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x48, 0x44, 0x4f, 0x42, 0x4a, 0x42,
	0x47, 0x50, 0x4e, 0x48, 0x52, 0x0b, 0x4a, 0x4c, 0x42, 0x44, 0x48, 0x43, 0x43, 0x50, 0x4a, 0x44,
	0x50, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMonopolyInfoScRsp_proto_rawDescOnce sync.Once
	file_GetMonopolyInfoScRsp_proto_rawDescData = file_GetMonopolyInfoScRsp_proto_rawDesc
)

func file_GetMonopolyInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetMonopolyInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetMonopolyInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMonopolyInfoScRsp_proto_rawDescData)
	})
	return file_GetMonopolyInfoScRsp_proto_rawDescData
}

var file_GetMonopolyInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMonopolyInfoScRsp_proto_goTypes = []interface{}{
	(*GetMonopolyInfoScRsp)(nil), // 0: GetMonopolyInfoScRsp
	(*LJMJPODBCAE)(nil),          // 1: LJMJPODBCAE
	(*OOPNGMHCNNN)(nil),          // 2: OOPNGMHCNNN
	(*MGGHEHLPFMH)(nil),          // 3: MGGHEHLPFMH
	(*CKFFBACIGPG)(nil),          // 4: CKFFBACIGPG
	(*DEFMCECJJBI)(nil),          // 5: DEFMCECJJBI
	(*GIODDOIHLCN)(nil),          // 6: GIODDOIHLCN
	(*JOEPAJDGPHK)(nil),          // 7: JOEPAJDGPHK
	(*EONNIEFDOCI)(nil),          // 8: EONNIEFDOCI
	(*KFBEAPLBEFB)(nil),          // 9: KFBEAPLBEFB
	(*JOGIAGPFDML)(nil),          // 10: JOGIAGPFDML
	(*KHDOBJBGPNH)(nil),          // 11: KHDOBJBGPNH
}
var file_GetMonopolyInfoScRsp_proto_depIdxs = []int32{
	1,  // 0: GetMonopolyInfoScRsp.OMDFOLOFPDB:type_name -> LJMJPODBCAE
	2,  // 1: GetMonopolyInfoScRsp.room_map:type_name -> OOPNGMHCNNN
	3,  // 2: GetMonopolyInfoScRsp.CAIMMNMCPCJ:type_name -> MGGHEHLPFMH
	4,  // 3: GetMonopolyInfoScRsp.rogue_buff_info:type_name -> CKFFBACIGPG
	5,  // 4: GetMonopolyInfoScRsp.IKAFFGFKBFN:type_name -> DEFMCECJJBI
	6,  // 5: GetMonopolyInfoScRsp.JKKGEBEMGOF:type_name -> GIODDOIHLCN
	7,  // 6: GetMonopolyInfoScRsp.DCPFFMNFMKJ:type_name -> JOEPAJDGPHK
	8,  // 7: GetMonopolyInfoScRsp.CAJDMFNLNKD:type_name -> EONNIEFDOCI
	9,  // 8: GetMonopolyInfoScRsp.stt:type_name -> KFBEAPLBEFB
	10, // 9: GetMonopolyInfoScRsp.JADMBMOHANO:type_name -> JOGIAGPFDML
	11, // 10: GetMonopolyInfoScRsp.JLBDHCCPJDP:type_name -> KHDOBJBGPNH
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_GetMonopolyInfoScRsp_proto_init() }
func file_GetMonopolyInfoScRsp_proto_init() {
	if File_GetMonopolyInfoScRsp_proto != nil {
		return
	}
	file_CKFFBACIGPG_proto_init()
	file_JOGIAGPFDML_proto_init()
	file_LJMJPODBCAE_proto_init()
	file_GIODDOIHLCN_proto_init()
	file_EONNIEFDOCI_proto_init()
	file_KFBEAPLBEFB_proto_init()
	file_MGGHEHLPFMH_proto_init()
	file_OOPNGMHCNNN_proto_init()
	file_KHDOBJBGPNH_proto_init()
	file_DEFMCECJJBI_proto_init()
	file_JOEPAJDGPHK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMonopolyInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonopolyInfoScRsp); i {
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
			RawDescriptor: file_GetMonopolyInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMonopolyInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetMonopolyInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetMonopolyInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetMonopolyInfoScRsp_proto = out.File
	file_GetMonopolyInfoScRsp_proto_rawDesc = nil
	file_GetMonopolyInfoScRsp_proto_goTypes = nil
	file_GetMonopolyInfoScRsp_proto_depIdxs = nil
}
