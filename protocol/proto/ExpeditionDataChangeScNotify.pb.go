// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ExpeditionDataChangeScNotify.proto

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

type ExpeditionDataChangeScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INMIJJGHKEO uint32                `protobuf:"varint,11,opt,name=INMIJJGHKEO,proto3" json:"INMIJJGHKEO,omitempty"`
	KEKNKGCDNND []uint32              `protobuf:"varint,5,rep,packed,name=KEKNKGCDNND,proto3" json:"KEKNKGCDNND,omitempty"`
	OPMKFNGGGAN []uint32              `protobuf:"varint,4,rep,packed,name=OPMKFNGGGAN,proto3" json:"OPMKFNGGGAN,omitempty"`
	NPKCMMLBJKD []*ActivityExpedition `protobuf:"bytes,1,rep,name=NPKCMMLBJKD,proto3" json:"NPKCMMLBJKD,omitempty"`
	ILEKFPLFMDC []*KKMCMAFLCLF        `protobuf:"bytes,9,rep,name=ILEKFPLFMDC,proto3" json:"ILEKFPLFMDC,omitempty"`
}

func (x *ExpeditionDataChangeScNotify) Reset() {
	*x = ExpeditionDataChangeScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExpeditionDataChangeScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpeditionDataChangeScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpeditionDataChangeScNotify) ProtoMessage() {}

func (x *ExpeditionDataChangeScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ExpeditionDataChangeScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpeditionDataChangeScNotify.ProtoReflect.Descriptor instead.
func (*ExpeditionDataChangeScNotify) Descriptor() ([]byte, []int) {
	return file_ExpeditionDataChangeScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ExpeditionDataChangeScNotify) GetINMIJJGHKEO() uint32 {
	if x != nil {
		return x.INMIJJGHKEO
	}
	return 0
}

func (x *ExpeditionDataChangeScNotify) GetKEKNKGCDNND() []uint32 {
	if x != nil {
		return x.KEKNKGCDNND
	}
	return nil
}

func (x *ExpeditionDataChangeScNotify) GetOPMKFNGGGAN() []uint32 {
	if x != nil {
		return x.OPMKFNGGGAN
	}
	return nil
}

func (x *ExpeditionDataChangeScNotify) GetNPKCMMLBJKD() []*ActivityExpedition {
	if x != nil {
		return x.NPKCMMLBJKD
	}
	return nil
}

func (x *ExpeditionDataChangeScNotify) GetILEKFPLFMDC() []*KKMCMAFLCLF {
	if x != nil {
		return x.ILEKFPLFMDC
	}
	return nil
}

var File_ExpeditionDataChangeScNotify_proto protoreflect.FileDescriptor

var file_ExpeditionDataChangeScNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4b, 0x4d, 0x43, 0x4d, 0x41, 0x46, 0x4c, 0x43, 0x4c,
	0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x1c, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x4d, 0x49, 0x4a, 0x4a, 0x47, 0x48, 0x4b, 0x45,
	0x4f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x4d, 0x49, 0x4a, 0x4a, 0x47,
	0x48, 0x4b, 0x45, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x45, 0x4b, 0x4e, 0x4b, 0x47, 0x43, 0x44,
	0x4e, 0x4e, 0x44, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x45, 0x4b, 0x4e, 0x4b,
	0x47, 0x43, 0x44, 0x4e, 0x4e, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x50, 0x4d, 0x4b, 0x46, 0x4e,
	0x47, 0x47, 0x47, 0x41, 0x4e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x50, 0x4d,
	0x4b, 0x46, 0x4e, 0x47, 0x47, 0x47, 0x41, 0x4e, 0x12, 0x35, 0x0a, 0x0b, 0x4e, 0x50, 0x4b, 0x43,
	0x4d, 0x4d, 0x4c, 0x42, 0x4a, 0x4b, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x4e, 0x50, 0x4b, 0x43, 0x4d, 0x4d, 0x4c, 0x42, 0x4a, 0x4b, 0x44, 0x12,
	0x2e, 0x0a, 0x0b, 0x49, 0x4c, 0x45, 0x4b, 0x46, 0x50, 0x4c, 0x46, 0x4d, 0x44, 0x43, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4b, 0x4d, 0x43, 0x4d, 0x41, 0x46, 0x4c, 0x43,
	0x4c, 0x46, 0x52, 0x0b, 0x49, 0x4c, 0x45, 0x4b, 0x46, 0x50, 0x4c, 0x46, 0x4d, 0x44, 0x43, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ExpeditionDataChangeScNotify_proto_rawDescOnce sync.Once
	file_ExpeditionDataChangeScNotify_proto_rawDescData = file_ExpeditionDataChangeScNotify_proto_rawDesc
)

func file_ExpeditionDataChangeScNotify_proto_rawDescGZIP() []byte {
	file_ExpeditionDataChangeScNotify_proto_rawDescOnce.Do(func() {
		file_ExpeditionDataChangeScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExpeditionDataChangeScNotify_proto_rawDescData)
	})
	return file_ExpeditionDataChangeScNotify_proto_rawDescData
}

var file_ExpeditionDataChangeScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExpeditionDataChangeScNotify_proto_goTypes = []interface{}{
	(*ExpeditionDataChangeScNotify)(nil), // 0: ExpeditionDataChangeScNotify
	(*ActivityExpedition)(nil),           // 1: ActivityExpedition
	(*KKMCMAFLCLF)(nil),                  // 2: KKMCMAFLCLF
}
var file_ExpeditionDataChangeScNotify_proto_depIdxs = []int32{
	1, // 0: ExpeditionDataChangeScNotify.NPKCMMLBJKD:type_name -> ActivityExpedition
	2, // 1: ExpeditionDataChangeScNotify.ILEKFPLFMDC:type_name -> KKMCMAFLCLF
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ExpeditionDataChangeScNotify_proto_init() }
func file_ExpeditionDataChangeScNotify_proto_init() {
	if File_ExpeditionDataChangeScNotify_proto != nil {
		return
	}
	file_KKMCMAFLCLF_proto_init()
	file_ActivityExpedition_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExpeditionDataChangeScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpeditionDataChangeScNotify); i {
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
			RawDescriptor: file_ExpeditionDataChangeScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExpeditionDataChangeScNotify_proto_goTypes,
		DependencyIndexes: file_ExpeditionDataChangeScNotify_proto_depIdxs,
		MessageInfos:      file_ExpeditionDataChangeScNotify_proto_msgTypes,
	}.Build()
	File_ExpeditionDataChangeScNotify_proto = out.File
	file_ExpeditionDataChangeScNotify_proto_rawDesc = nil
	file_ExpeditionDataChangeScNotify_proto_goTypes = nil
	file_ExpeditionDataChangeScNotify_proto_depIdxs = nil
}
