// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournGetMiscRealTimeDataScRsp.proto

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

type RogueTournGetMiscRealTimeDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DHPKPLOJIOD *EJFLFKFJBJC `protobuf:"bytes,7,opt,name=DHPKPLOJIOD,proto3" json:"DHPKPLOJIOD,omitempty"`
	ECCJMFFLMMG *IBPNIMLNPLI `protobuf:"bytes,13,opt,name=ECCJMFFLMMG,proto3" json:"ECCJMFFLMMG,omitempty"`
	GCNCKFCFOID *FGOIANHPJHH `protobuf:"bytes,12,opt,name=GCNCKFCFOID,proto3" json:"GCNCKFCFOID,omitempty"`
	Retcode     uint32       `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *RogueTournGetMiscRealTimeDataScRsp) Reset() {
	*x = RogueTournGetMiscRealTimeDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournGetMiscRealTimeDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournGetMiscRealTimeDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournGetMiscRealTimeDataScRsp) ProtoMessage() {}

func (x *RogueTournGetMiscRealTimeDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournGetMiscRealTimeDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournGetMiscRealTimeDataScRsp.ProtoReflect.Descriptor instead.
func (*RogueTournGetMiscRealTimeDataScRsp) Descriptor() ([]byte, []int) {
	return file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournGetMiscRealTimeDataScRsp) GetDHPKPLOJIOD() *EJFLFKFJBJC {
	if x != nil {
		return x.DHPKPLOJIOD
	}
	return nil
}

func (x *RogueTournGetMiscRealTimeDataScRsp) GetECCJMFFLMMG() *IBPNIMLNPLI {
	if x != nil {
		return x.ECCJMFFLMMG
	}
	return nil
}

func (x *RogueTournGetMiscRealTimeDataScRsp) GetGCNCKFCFOID() *FGOIANHPJHH {
	if x != nil {
		return x.GCNCKFCFOID
	}
	return nil
}

func (x *RogueTournGetMiscRealTimeDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_RogueTournGetMiscRealTimeDataScRsp_proto protoreflect.FileDescriptor

var file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x47, 0x65, 0x74, 0x4d,
	0x69, 0x73, 0x63, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4a, 0x46, 0x4c,
	0x46, 0x4b, 0x46, 0x4a, 0x42, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49,
	0x42, 0x50, 0x4e, 0x49, 0x4d, 0x4c, 0x4e, 0x50, 0x4c, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x46, 0x47, 0x4f, 0x49, 0x41, 0x4e, 0x48, 0x50, 0x4a, 0x48, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x63, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x48,
	0x50, 0x4b, 0x50, 0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x45, 0x4a, 0x46, 0x4c, 0x46, 0x4b, 0x46, 0x4a, 0x42, 0x4a, 0x43, 0x52, 0x0b, 0x44,
	0x48, 0x50, 0x4b, 0x50, 0x4c, 0x4f, 0x4a, 0x49, 0x4f, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x43,
	0x43, 0x4a, 0x4d, 0x46, 0x46, 0x4c, 0x4d, 0x4d, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x42, 0x50, 0x4e, 0x49, 0x4d, 0x4c, 0x4e, 0x50, 0x4c, 0x49, 0x52, 0x0b, 0x45,
	0x43, 0x43, 0x4a, 0x4d, 0x46, 0x46, 0x4c, 0x4d, 0x4d, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x43,
	0x4e, 0x43, 0x4b, 0x46, 0x43, 0x46, 0x4f, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x46, 0x47, 0x4f, 0x49, 0x41, 0x4e, 0x48, 0x50, 0x4a, 0x48, 0x48, 0x52, 0x0b, 0x47,
	0x43, 0x4e, 0x43, 0x4b, 0x46, 0x43, 0x46, 0x4f, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescOnce sync.Once
	file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescData = file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDesc
)

func file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescGZIP() []byte {
	file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescOnce.Do(func() {
		file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescData)
	})
	return file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDescData
}

var file_RogueTournGetMiscRealTimeDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournGetMiscRealTimeDataScRsp_proto_goTypes = []interface{}{
	(*RogueTournGetMiscRealTimeDataScRsp)(nil), // 0: RogueTournGetMiscRealTimeDataScRsp
	(*EJFLFKFJBJC)(nil),                        // 1: EJFLFKFJBJC
	(*IBPNIMLNPLI)(nil),                        // 2: IBPNIMLNPLI
	(*FGOIANHPJHH)(nil),                        // 3: FGOIANHPJHH
}
var file_RogueTournGetMiscRealTimeDataScRsp_proto_depIdxs = []int32{
	1, // 0: RogueTournGetMiscRealTimeDataScRsp.DHPKPLOJIOD:type_name -> EJFLFKFJBJC
	2, // 1: RogueTournGetMiscRealTimeDataScRsp.ECCJMFFLMMG:type_name -> IBPNIMLNPLI
	3, // 2: RogueTournGetMiscRealTimeDataScRsp.GCNCKFCFOID:type_name -> FGOIANHPJHH
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RogueTournGetMiscRealTimeDataScRsp_proto_init() }
func file_RogueTournGetMiscRealTimeDataScRsp_proto_init() {
	if File_RogueTournGetMiscRealTimeDataScRsp_proto != nil {
		return
	}
	file_EJFLFKFJBJC_proto_init()
	file_IBPNIMLNPLI_proto_init()
	file_FGOIANHPJHH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournGetMiscRealTimeDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournGetMiscRealTimeDataScRsp); i {
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
			RawDescriptor: file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournGetMiscRealTimeDataScRsp_proto_goTypes,
		DependencyIndexes: file_RogueTournGetMiscRealTimeDataScRsp_proto_depIdxs,
		MessageInfos:      file_RogueTournGetMiscRealTimeDataScRsp_proto_msgTypes,
	}.Build()
	File_RogueTournGetMiscRealTimeDataScRsp_proto = out.File
	file_RogueTournGetMiscRealTimeDataScRsp_proto_rawDesc = nil
	file_RogueTournGetMiscRealTimeDataScRsp_proto_goTypes = nil
	file_RogueTournGetMiscRealTimeDataScRsp_proto_depIdxs = nil
}