// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetSpringRecoverDataScRsp.proto

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

type GetSpringRecoverDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JLMLFKBMJNL *SpringRecoverConfig `protobuf:"bytes,13,opt,name=JLMLFKBMJNL,proto3" json:"JLMLFKBMJNL,omitempty"`
	KMNCAPJPGDP *HealPoolInfo        `protobuf:"bytes,12,opt,name=KMNCAPJPGDP,proto3" json:"KMNCAPJPGDP,omitempty"`
	Retcode     uint32               `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetSpringRecoverDataScRsp) Reset() {
	*x = GetSpringRecoverDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSpringRecoverDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpringRecoverDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpringRecoverDataScRsp) ProtoMessage() {}

func (x *GetSpringRecoverDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetSpringRecoverDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpringRecoverDataScRsp.ProtoReflect.Descriptor instead.
func (*GetSpringRecoverDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetSpringRecoverDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetSpringRecoverDataScRsp) GetJLMLFKBMJNL() *SpringRecoverConfig {
	if x != nil {
		return x.JLMLFKBMJNL
	}
	return nil
}

func (x *GetSpringRecoverDataScRsp) GetKMNCAPJPGDP() *HealPoolInfo {
	if x != nil {
		return x.KMNCAPJPGDP
	}
	return nil
}

func (x *GetSpringRecoverDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetSpringRecoverDataScRsp_proto protoreflect.FileDescriptor

var file_GetSpringRecoverDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9e, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x36,
	0x0a, 0x0b, 0x4a, 0x4c, 0x4d, 0x4c, 0x46, 0x4b, 0x42, 0x4d, 0x4a, 0x4e, 0x4c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x4a, 0x4c, 0x4d, 0x4c, 0x46,
	0x4b, 0x42, 0x4d, 0x4a, 0x4e, 0x4c, 0x12, 0x2f, 0x0a, 0x0b, 0x4b, 0x4d, 0x4e, 0x43, 0x41, 0x50,
	0x4a, 0x50, 0x47, 0x44, 0x50, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x4b, 0x4d, 0x4e, 0x43,
	0x41, 0x50, 0x4a, 0x50, 0x47, 0x44, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetSpringRecoverDataScRsp_proto_rawDescOnce sync.Once
	file_GetSpringRecoverDataScRsp_proto_rawDescData = file_GetSpringRecoverDataScRsp_proto_rawDesc
)

func file_GetSpringRecoverDataScRsp_proto_rawDescGZIP() []byte {
	file_GetSpringRecoverDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetSpringRecoverDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSpringRecoverDataScRsp_proto_rawDescData)
	})
	return file_GetSpringRecoverDataScRsp_proto_rawDescData
}

var file_GetSpringRecoverDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSpringRecoverDataScRsp_proto_goTypes = []interface{}{
	(*GetSpringRecoverDataScRsp)(nil), // 0: GetSpringRecoverDataScRsp
	(*SpringRecoverConfig)(nil),       // 1: SpringRecoverConfig
	(*HealPoolInfo)(nil),              // 2: HealPoolInfo
}
var file_GetSpringRecoverDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetSpringRecoverDataScRsp.JLMLFKBMJNL:type_name -> SpringRecoverConfig
	2, // 1: GetSpringRecoverDataScRsp.KMNCAPJPGDP:type_name -> HealPoolInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetSpringRecoverDataScRsp_proto_init() }
func file_GetSpringRecoverDataScRsp_proto_init() {
	if File_GetSpringRecoverDataScRsp_proto != nil {
		return
	}
	file_SpringRecoverConfig_proto_init()
	file_HealPoolInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetSpringRecoverDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpringRecoverDataScRsp); i {
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
			RawDescriptor: file_GetSpringRecoverDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSpringRecoverDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetSpringRecoverDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetSpringRecoverDataScRsp_proto_msgTypes,
	}.Build()
	File_GetSpringRecoverDataScRsp_proto = out.File
	file_GetSpringRecoverDataScRsp_proto_rawDesc = nil
	file_GetSpringRecoverDataScRsp_proto_goTypes = nil
	file_GetSpringRecoverDataScRsp_proto_depIdxs = nil
}