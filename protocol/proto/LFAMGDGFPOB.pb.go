// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LFAMGDGFPOB.proto

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

type LFAMGDGFPOB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BJGHOFCJFOD *KENHBBHJBDK `protobuf:"bytes,803,opt,name=BJGHOFCJFOD,proto3" json:"BJGHOFCJFOD,omitempty"`
	DCOEAGKMCDK *JGKKIACMHFE `protobuf:"bytes,1332,opt,name=DCOEAGKMCDK,proto3" json:"DCOEAGKMCDK,omitempty"`
}

func (x *LFAMGDGFPOB) Reset() {
	*x = LFAMGDGFPOB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LFAMGDGFPOB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LFAMGDGFPOB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LFAMGDGFPOB) ProtoMessage() {}

func (x *LFAMGDGFPOB) ProtoReflect() protoreflect.Message {
	mi := &file_LFAMGDGFPOB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LFAMGDGFPOB.ProtoReflect.Descriptor instead.
func (*LFAMGDGFPOB) Descriptor() ([]byte, []int) {
	return file_LFAMGDGFPOB_proto_rawDescGZIP(), []int{0}
}

func (x *LFAMGDGFPOB) GetBJGHOFCJFOD() *KENHBBHJBDK {
	if x != nil {
		return x.BJGHOFCJFOD
	}
	return nil
}

func (x *LFAMGDGFPOB) GetDCOEAGKMCDK() *JGKKIACMHFE {
	if x != nil {
		return x.DCOEAGKMCDK
	}
	return nil
}

var File_LFAMGDGFPOB_proto protoreflect.FileDescriptor

var file_LFAMGDGFPOB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x46, 0x41, 0x4d, 0x47, 0x44, 0x47, 0x46, 0x50, 0x4f, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x45, 0x4e, 0x48, 0x42, 0x42, 0x48, 0x4a, 0x42, 0x44, 0x4b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x47, 0x4b, 0x4b, 0x49, 0x41, 0x43, 0x4d,
	0x48, 0x46, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0b, 0x4c, 0x46, 0x41,
	0x4d, 0x47, 0x44, 0x47, 0x46, 0x50, 0x4f, 0x42, 0x12, 0x2f, 0x0a, 0x0b, 0x42, 0x4a, 0x47, 0x48,
	0x4f, 0x46, 0x43, 0x4a, 0x46, 0x4f, 0x44, 0x18, 0xa3, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4b, 0x45, 0x4e, 0x48, 0x42, 0x42, 0x48, 0x4a, 0x42, 0x44, 0x4b, 0x52, 0x0b, 0x42, 0x4a,
	0x47, 0x48, 0x4f, 0x46, 0x43, 0x4a, 0x46, 0x4f, 0x44, 0x12, 0x2f, 0x0a, 0x0b, 0x44, 0x43, 0x4f,
	0x45, 0x41, 0x47, 0x4b, 0x4d, 0x43, 0x44, 0x4b, 0x18, 0xb4, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4a, 0x47, 0x4b, 0x4b, 0x49, 0x41, 0x43, 0x4d, 0x48, 0x46, 0x45, 0x52, 0x0b, 0x44,
	0x43, 0x4f, 0x45, 0x41, 0x47, 0x4b, 0x4d, 0x43, 0x44, 0x4b, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LFAMGDGFPOB_proto_rawDescOnce sync.Once
	file_LFAMGDGFPOB_proto_rawDescData = file_LFAMGDGFPOB_proto_rawDesc
)

func file_LFAMGDGFPOB_proto_rawDescGZIP() []byte {
	file_LFAMGDGFPOB_proto_rawDescOnce.Do(func() {
		file_LFAMGDGFPOB_proto_rawDescData = protoimpl.X.CompressGZIP(file_LFAMGDGFPOB_proto_rawDescData)
	})
	return file_LFAMGDGFPOB_proto_rawDescData
}

var file_LFAMGDGFPOB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LFAMGDGFPOB_proto_goTypes = []interface{}{
	(*LFAMGDGFPOB)(nil), // 0: LFAMGDGFPOB
	(*KENHBBHJBDK)(nil), // 1: KENHBBHJBDK
	(*JGKKIACMHFE)(nil), // 2: JGKKIACMHFE
}
var file_LFAMGDGFPOB_proto_depIdxs = []int32{
	1, // 0: LFAMGDGFPOB.BJGHOFCJFOD:type_name -> KENHBBHJBDK
	2, // 1: LFAMGDGFPOB.DCOEAGKMCDK:type_name -> JGKKIACMHFE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LFAMGDGFPOB_proto_init() }
func file_LFAMGDGFPOB_proto_init() {
	if File_LFAMGDGFPOB_proto != nil {
		return
	}
	file_KENHBBHJBDK_proto_init()
	file_JGKKIACMHFE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LFAMGDGFPOB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LFAMGDGFPOB); i {
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
			RawDescriptor: file_LFAMGDGFPOB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LFAMGDGFPOB_proto_goTypes,
		DependencyIndexes: file_LFAMGDGFPOB_proto_depIdxs,
		MessageInfos:      file_LFAMGDGFPOB_proto_msgTypes,
	}.Build()
	File_LFAMGDGFPOB_proto = out.File
	file_LFAMGDGFPOB_proto_rawDesc = nil
	file_LFAMGDGFPOB_proto_goTypes = nil
	file_LFAMGDGFPOB_proto_depIdxs = nil
}
