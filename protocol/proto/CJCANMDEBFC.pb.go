// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CJCANMDEBFC.proto

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

type CJCANMDEBFC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FDACLBNKPMJ []uint32            `protobuf:"varint,10,rep,packed,name=FDACLBNKPMJ,proto3" json:"FDACLBNKPMJ,omitempty"`
	BAHIPGNCFJO RogueDialogueResult `protobuf:"varint,2,opt,name=BAHIPGNCFJO,proto3,enum=RogueDialogueResult" json:"BAHIPGNCFJO,omitempty"`
	// Types that are assignable to ONNMJDFLDIJ:
	//
	//	*CJCANMDEBFC_LBHLMLAHHME
	ONNMJDFLDIJ isCJCANMDEBFC_ONNMJDFLDIJ `protobuf_oneof:"ONNMJDFLDIJ"`
}

func (x *CJCANMDEBFC) Reset() {
	*x = CJCANMDEBFC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CJCANMDEBFC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CJCANMDEBFC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CJCANMDEBFC) ProtoMessage() {}

func (x *CJCANMDEBFC) ProtoReflect() protoreflect.Message {
	mi := &file_CJCANMDEBFC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CJCANMDEBFC.ProtoReflect.Descriptor instead.
func (*CJCANMDEBFC) Descriptor() ([]byte, []int) {
	return file_CJCANMDEBFC_proto_rawDescGZIP(), []int{0}
}

func (x *CJCANMDEBFC) GetFDACLBNKPMJ() []uint32 {
	if x != nil {
		return x.FDACLBNKPMJ
	}
	return nil
}

func (x *CJCANMDEBFC) GetBAHIPGNCFJO() RogueDialogueResult {
	if x != nil {
		return x.BAHIPGNCFJO
	}
	return RogueDialogueResult_ROGUE_DIALOGUE_RESULT_SUCC
}

func (m *CJCANMDEBFC) GetONNMJDFLDIJ() isCJCANMDEBFC_ONNMJDFLDIJ {
	if m != nil {
		return m.ONNMJDFLDIJ
	}
	return nil
}

func (x *CJCANMDEBFC) GetLBHLMLAHHME() *ItemList {
	if x, ok := x.GetONNMJDFLDIJ().(*CJCANMDEBFC_LBHLMLAHHME); ok {
		return x.LBHLMLAHHME
	}
	return nil
}

type isCJCANMDEBFC_ONNMJDFLDIJ interface {
	isCJCANMDEBFC_ONNMJDFLDIJ()
}

type CJCANMDEBFC_LBHLMLAHHME struct {
	LBHLMLAHHME *ItemList `protobuf:"bytes,13,opt,name=LBHLMLAHHME,proto3,oneof"`
}

func (*CJCANMDEBFC_LBHLMLAHHME) isCJCANMDEBFC_ONNMJDFLDIJ() {}

var File_CJCANMDEBFC_proto protoreflect.FileDescriptor

var file_CJCANMDEBFC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4a, 0x43, 0x41, 0x4e, 0x4d, 0x44, 0x45, 0x42, 0x46, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x01, 0x0a, 0x0b, 0x43, 0x4a, 0x43, 0x41, 0x4e, 0x4d, 0x44, 0x45, 0x42, 0x46, 0x43, 0x12, 0x20,
	0x0a, 0x0b, 0x46, 0x44, 0x41, 0x43, 0x4c, 0x42, 0x4e, 0x4b, 0x50, 0x4d, 0x4a, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x44, 0x41, 0x43, 0x4c, 0x42, 0x4e, 0x4b, 0x50, 0x4d, 0x4a,
	0x12, 0x36, 0x0a, 0x0b, 0x42, 0x41, 0x48, 0x49, 0x50, 0x47, 0x4e, 0x43, 0x46, 0x4a, 0x4f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x42, 0x41, 0x48,
	0x49, 0x50, 0x47, 0x4e, 0x43, 0x46, 0x4a, 0x4f, 0x12, 0x2d, 0x0a, 0x0b, 0x4c, 0x42, 0x48, 0x4c,
	0x4d, 0x4c, 0x41, 0x48, 0x48, 0x4d, 0x45, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x42, 0x48, 0x4c,
	0x4d, 0x4c, 0x41, 0x48, 0x48, 0x4d, 0x45, 0x42, 0x0d, 0x0a, 0x0b, 0x4f, 0x4e, 0x4e, 0x4d, 0x4a,
	0x44, 0x46, 0x4c, 0x44, 0x49, 0x4a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CJCANMDEBFC_proto_rawDescOnce sync.Once
	file_CJCANMDEBFC_proto_rawDescData = file_CJCANMDEBFC_proto_rawDesc
)

func file_CJCANMDEBFC_proto_rawDescGZIP() []byte {
	file_CJCANMDEBFC_proto_rawDescOnce.Do(func() {
		file_CJCANMDEBFC_proto_rawDescData = protoimpl.X.CompressGZIP(file_CJCANMDEBFC_proto_rawDescData)
	})
	return file_CJCANMDEBFC_proto_rawDescData
}

var file_CJCANMDEBFC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CJCANMDEBFC_proto_goTypes = []interface{}{
	(*CJCANMDEBFC)(nil),      // 0: CJCANMDEBFC
	(RogueDialogueResult)(0), // 1: RogueDialogueResult
	(*ItemList)(nil),         // 2: ItemList
}
var file_CJCANMDEBFC_proto_depIdxs = []int32{
	1, // 0: CJCANMDEBFC.BAHIPGNCFJO:type_name -> RogueDialogueResult
	2, // 1: CJCANMDEBFC.LBHLMLAHHME:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CJCANMDEBFC_proto_init() }
func file_CJCANMDEBFC_proto_init() {
	if File_CJCANMDEBFC_proto != nil {
		return
	}
	file_RogueDialogueResult_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CJCANMDEBFC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CJCANMDEBFC); i {
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
	file_CJCANMDEBFC_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CJCANMDEBFC_LBHLMLAHHME)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CJCANMDEBFC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CJCANMDEBFC_proto_goTypes,
		DependencyIndexes: file_CJCANMDEBFC_proto_depIdxs,
		MessageInfos:      file_CJCANMDEBFC_proto_msgTypes,
	}.Build()
	File_CJCANMDEBFC_proto = out.File
	file_CJCANMDEBFC_proto_rawDesc = nil
	file_CJCANMDEBFC_proto_goTypes = nil
	file_CJCANMDEBFC_proto_depIdxs = nil
}