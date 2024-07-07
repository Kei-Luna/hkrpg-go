// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AAFHGGHCNLO.proto

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

type AAFHGGHCNLO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GCNCKFCFOID       *FGOIANHPJHH       `protobuf:"bytes,15,opt,name=GCNCKFCFOID,proto3" json:"GCNCKFCFOID,omitempty"`
	ECCJMFFLMMG       *IBPNIMLNPLI       `protobuf:"bytes,2,opt,name=ECCJMFFLMMG,proto3" json:"ECCJMFFLMMG,omitempty"`
	DHPKPLOJIOD       *EJFLFKFJBJC       `protobuf:"bytes,10,opt,name=DHPKPLOJIOD,proto3" json:"DHPKPLOJIOD,omitempty"`
	RogueLineupInfo   *LineupInfo        `protobuf:"bytes,4,opt,name=rogue_lineup_info,json=rogueLineupInfo,proto3" json:"rogue_lineup_info,omitempty"`
	RogueTournCurInfo *RogueTournCurInfo `protobuf:"bytes,5,opt,name=rogue_tourn_cur_info,json=rogueTournCurInfo,proto3" json:"rogue_tourn_cur_info,omitempty"`
}

func (x *AAFHGGHCNLO) Reset() {
	*x = AAFHGGHCNLO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AAFHGGHCNLO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AAFHGGHCNLO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AAFHGGHCNLO) ProtoMessage() {}

func (x *AAFHGGHCNLO) ProtoReflect() protoreflect.Message {
	mi := &file_AAFHGGHCNLO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AAFHGGHCNLO.ProtoReflect.Descriptor instead.
func (*AAFHGGHCNLO) Descriptor() ([]byte, []int) {
	return file_AAFHGGHCNLO_proto_rawDescGZIP(), []int{0}
}

func (x *AAFHGGHCNLO) GetGCNCKFCFOID() *FGOIANHPJHH {
	if x != nil {
		return x.GCNCKFCFOID
	}
	return nil
}

func (x *AAFHGGHCNLO) GetECCJMFFLMMG() *IBPNIMLNPLI {
	if x != nil {
		return x.ECCJMFFLMMG
	}
	return nil
}

func (x *AAFHGGHCNLO) GetDHPKPLOJIOD() *EJFLFKFJBJC {
	if x != nil {
		return x.DHPKPLOJIOD
	}
	return nil
}

func (x *AAFHGGHCNLO) GetRogueLineupInfo() *LineupInfo {
	if x != nil {
		return x.RogueLineupInfo
	}
	return nil
}

func (x *AAFHGGHCNLO) GetRogueTournCurInfo() *RogueTournCurInfo {
	if x != nil {
		return x.RogueTournCurInfo
	}
	return nil
}

var File_AAFHGGHCNLO_proto protoreflect.FileDescriptor

var file_AAFHGGHCNLO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x41, 0x46, 0x48, 0x47, 0x47, 0x48, 0x43, 0x4e, 0x4c, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4a, 0x46, 0x4c, 0x46, 0x4b, 0x46, 0x4a, 0x42, 0x4a, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x42, 0x50, 0x4e, 0x49, 0x4d, 0x4c, 0x4e,
	0x50, 0x4c, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x47, 0x4f, 0x49, 0x41,
	0x4e, 0x48, 0x50, 0x4a, 0x48, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0b, 0x41, 0x41, 0x46, 0x48,
	0x47, 0x47, 0x48, 0x43, 0x4e, 0x4c, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x43, 0x4e, 0x43, 0x4b,
	0x46, 0x43, 0x46, 0x4f, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46,
	0x47, 0x4f, 0x49, 0x41, 0x4e, 0x48, 0x50, 0x4a, 0x48, 0x48, 0x52, 0x0b, 0x47, 0x43, 0x4e, 0x43,
	0x4b, 0x46, 0x43, 0x46, 0x4f, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x43, 0x43, 0x4a, 0x4d,
	0x46, 0x46, 0x4c, 0x4d, 0x4d, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49,
	0x42, 0x50, 0x4e, 0x49, 0x4d, 0x4c, 0x4e, 0x50, 0x4c, 0x49, 0x52, 0x0b, 0x45, 0x43, 0x43, 0x4a,
	0x4d, 0x46, 0x46, 0x4c, 0x4d, 0x4d, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x48, 0x50, 0x4b, 0x50,
	0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x4a, 0x46, 0x4c, 0x46, 0x4b, 0x46, 0x4a, 0x42, 0x4a, 0x43, 0x52, 0x0b, 0x44, 0x48, 0x50, 0x4b,
	0x50, 0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x44, 0x12, 0x37, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x43, 0x0a, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f,
	0x63, 0x75, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AAFHGGHCNLO_proto_rawDescOnce sync.Once
	file_AAFHGGHCNLO_proto_rawDescData = file_AAFHGGHCNLO_proto_rawDesc
)

func file_AAFHGGHCNLO_proto_rawDescGZIP() []byte {
	file_AAFHGGHCNLO_proto_rawDescOnce.Do(func() {
		file_AAFHGGHCNLO_proto_rawDescData = protoimpl.X.CompressGZIP(file_AAFHGGHCNLO_proto_rawDescData)
	})
	return file_AAFHGGHCNLO_proto_rawDescData
}

var file_AAFHGGHCNLO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AAFHGGHCNLO_proto_goTypes = []interface{}{
	(*AAFHGGHCNLO)(nil),       // 0: AAFHGGHCNLO
	(*FGOIANHPJHH)(nil),       // 1: FGOIANHPJHH
	(*IBPNIMLNPLI)(nil),       // 2: IBPNIMLNPLI
	(*EJFLFKFJBJC)(nil),       // 3: EJFLFKFJBJC
	(*LineupInfo)(nil),        // 4: LineupInfo
	(*RogueTournCurInfo)(nil), // 5: RogueTournCurInfo
}
var file_AAFHGGHCNLO_proto_depIdxs = []int32{
	1, // 0: AAFHGGHCNLO.GCNCKFCFOID:type_name -> FGOIANHPJHH
	2, // 1: AAFHGGHCNLO.ECCJMFFLMMG:type_name -> IBPNIMLNPLI
	3, // 2: AAFHGGHCNLO.DHPKPLOJIOD:type_name -> EJFLFKFJBJC
	4, // 3: AAFHGGHCNLO.rogue_lineup_info:type_name -> LineupInfo
	5, // 4: AAFHGGHCNLO.rogue_tourn_cur_info:type_name -> RogueTournCurInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_AAFHGGHCNLO_proto_init() }
func file_AAFHGGHCNLO_proto_init() {
	if File_AAFHGGHCNLO_proto != nil {
		return
	}
	file_EJFLFKFJBJC_proto_init()
	file_IBPNIMLNPLI_proto_init()
	file_FGOIANHPJHH_proto_init()
	file_RogueTournCurInfo_proto_init()
	file_LineupInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AAFHGGHCNLO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AAFHGGHCNLO); i {
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
			RawDescriptor: file_AAFHGGHCNLO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AAFHGGHCNLO_proto_goTypes,
		DependencyIndexes: file_AAFHGGHCNLO_proto_depIdxs,
		MessageInfos:      file_AAFHGGHCNLO_proto_msgTypes,
	}.Build()
	File_AAFHGGHCNLO_proto = out.File
	file_AAFHGGHCNLO_proto_rawDesc = nil
	file_AAFHGGHCNLO_proto_goTypes = nil
	file_AAFHGGHCNLO_proto_depIdxs = nil
}