// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NGPBEMIAJAL.proto

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

type NGPBEMIAJAL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MLIGMJCFNOI *CDAHKLGEFEG `protobuf:"bytes,1301,opt,name=MLIGMJCFNOI,proto3" json:"MLIGMJCFNOI,omitempty"`
	HDGKDOEODGI *EMMLCGHKHKC `protobuf:"bytes,1602,opt,name=HDGKDOEODGI,proto3" json:"HDGKDOEODGI,omitempty"`
	FFCMIAJDCCI *ADLMGDKBOEO `protobuf:"bytes,3,opt,name=FFCMIAJDCCI,proto3" json:"FFCMIAJDCCI,omitempty"`
	KLDDPLKKGHP *ALOEFGOHKLE `protobuf:"bytes,10,opt,name=KLDDPLKKGHP,proto3" json:"KLDDPLKKGHP,omitempty"`
}

func (x *NGPBEMIAJAL) Reset() {
	*x = NGPBEMIAJAL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NGPBEMIAJAL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NGPBEMIAJAL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NGPBEMIAJAL) ProtoMessage() {}

func (x *NGPBEMIAJAL) ProtoReflect() protoreflect.Message {
	mi := &file_NGPBEMIAJAL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NGPBEMIAJAL.ProtoReflect.Descriptor instead.
func (*NGPBEMIAJAL) Descriptor() ([]byte, []int) {
	return file_NGPBEMIAJAL_proto_rawDescGZIP(), []int{0}
}

func (x *NGPBEMIAJAL) GetMLIGMJCFNOI() *CDAHKLGEFEG {
	if x != nil {
		return x.MLIGMJCFNOI
	}
	return nil
}

func (x *NGPBEMIAJAL) GetHDGKDOEODGI() *EMMLCGHKHKC {
	if x != nil {
		return x.HDGKDOEODGI
	}
	return nil
}

func (x *NGPBEMIAJAL) GetFFCMIAJDCCI() *ADLMGDKBOEO {
	if x != nil {
		return x.FFCMIAJDCCI
	}
	return nil
}

func (x *NGPBEMIAJAL) GetKLDDPLKKGHP() *ALOEFGOHKLE {
	if x != nil {
		return x.KLDDPLKKGHP
	}
	return nil
}

var File_NGPBEMIAJAL_proto protoreflect.FileDescriptor

var file_NGPBEMIAJAL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x47, 0x50, 0x42, 0x45, 0x4d, 0x49, 0x41, 0x4a, 0x41, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4c, 0x4f, 0x45, 0x46, 0x47, 0x4f, 0x48, 0x4b, 0x4c, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x44, 0x4c, 0x4d, 0x47, 0x44, 0x4b, 0x42,
	0x4f, 0x45, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4d, 0x4d, 0x4c, 0x43,
	0x47, 0x48, 0x4b, 0x48, 0x4b, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x44,
	0x41, 0x48, 0x4b, 0x4c, 0x47, 0x45, 0x46, 0x45, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcf, 0x01, 0x0a, 0x0b, 0x4e, 0x47, 0x50, 0x42, 0x45, 0x4d, 0x49, 0x41, 0x4a, 0x41, 0x4c, 0x12,
	0x2f, 0x0a, 0x0b, 0x4d, 0x4c, 0x49, 0x47, 0x4d, 0x4a, 0x43, 0x46, 0x4e, 0x4f, 0x49, 0x18, 0x95,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x44, 0x41, 0x48, 0x4b, 0x4c, 0x47, 0x45,
	0x46, 0x45, 0x47, 0x52, 0x0b, 0x4d, 0x4c, 0x49, 0x47, 0x4d, 0x4a, 0x43, 0x46, 0x4e, 0x4f, 0x49,
	0x12, 0x2f, 0x0a, 0x0b, 0x48, 0x44, 0x47, 0x4b, 0x44, 0x4f, 0x45, 0x4f, 0x44, 0x47, 0x49, 0x18,
	0xc2, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4d, 0x4d, 0x4c, 0x43, 0x47, 0x48,
	0x4b, 0x48, 0x4b, 0x43, 0x52, 0x0b, 0x48, 0x44, 0x47, 0x4b, 0x44, 0x4f, 0x45, 0x4f, 0x44, 0x47,
	0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x46, 0x43, 0x4d, 0x49, 0x41, 0x4a, 0x44, 0x43, 0x43, 0x49,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x44, 0x4c, 0x4d, 0x47, 0x44, 0x4b,
	0x42, 0x4f, 0x45, 0x4f, 0x52, 0x0b, 0x46, 0x46, 0x43, 0x4d, 0x49, 0x41, 0x4a, 0x44, 0x43, 0x43,
	0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x4c, 0x44, 0x44, 0x50, 0x4c, 0x4b, 0x4b, 0x47, 0x48, 0x50,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4c, 0x4f, 0x45, 0x46, 0x47, 0x4f,
	0x48, 0x4b, 0x4c, 0x45, 0x52, 0x0b, 0x4b, 0x4c, 0x44, 0x44, 0x50, 0x4c, 0x4b, 0x4b, 0x47, 0x48,
	0x50, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NGPBEMIAJAL_proto_rawDescOnce sync.Once
	file_NGPBEMIAJAL_proto_rawDescData = file_NGPBEMIAJAL_proto_rawDesc
)

func file_NGPBEMIAJAL_proto_rawDescGZIP() []byte {
	file_NGPBEMIAJAL_proto_rawDescOnce.Do(func() {
		file_NGPBEMIAJAL_proto_rawDescData = protoimpl.X.CompressGZIP(file_NGPBEMIAJAL_proto_rawDescData)
	})
	return file_NGPBEMIAJAL_proto_rawDescData
}

var file_NGPBEMIAJAL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NGPBEMIAJAL_proto_goTypes = []interface{}{
	(*NGPBEMIAJAL)(nil), // 0: NGPBEMIAJAL
	(*CDAHKLGEFEG)(nil), // 1: CDAHKLGEFEG
	(*EMMLCGHKHKC)(nil), // 2: EMMLCGHKHKC
	(*ADLMGDKBOEO)(nil), // 3: ADLMGDKBOEO
	(*ALOEFGOHKLE)(nil), // 4: ALOEFGOHKLE
}
var file_NGPBEMIAJAL_proto_depIdxs = []int32{
	1, // 0: NGPBEMIAJAL.MLIGMJCFNOI:type_name -> CDAHKLGEFEG
	2, // 1: NGPBEMIAJAL.HDGKDOEODGI:type_name -> EMMLCGHKHKC
	3, // 2: NGPBEMIAJAL.FFCMIAJDCCI:type_name -> ADLMGDKBOEO
	4, // 3: NGPBEMIAJAL.KLDDPLKKGHP:type_name -> ALOEFGOHKLE
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_NGPBEMIAJAL_proto_init() }
func file_NGPBEMIAJAL_proto_init() {
	if File_NGPBEMIAJAL_proto != nil {
		return
	}
	file_ALOEFGOHKLE_proto_init()
	file_ADLMGDKBOEO_proto_init()
	file_EMMLCGHKHKC_proto_init()
	file_CDAHKLGEFEG_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NGPBEMIAJAL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NGPBEMIAJAL); i {
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
			RawDescriptor: file_NGPBEMIAJAL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NGPBEMIAJAL_proto_goTypes,
		DependencyIndexes: file_NGPBEMIAJAL_proto_depIdxs,
		MessageInfos:      file_NGPBEMIAJAL_proto_msgTypes,
	}.Build()
	File_NGPBEMIAJAL_proto = out.File
	file_NGPBEMIAJAL_proto_rawDesc = nil
	file_NGPBEMIAJAL_proto_goTypes = nil
	file_NGPBEMIAJAL_proto_depIdxs = nil
}
