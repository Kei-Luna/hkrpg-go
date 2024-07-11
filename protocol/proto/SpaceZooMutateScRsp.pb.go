// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SpaceZooMutateScRsp.proto

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

type SpaceZooMutateScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IHGHAGMAFJN bool           `protobuf:"varint,13,opt,name=IHGHAGMAFJN,proto3" json:"IHGHAGMAFJN,omitempty"`
	MGEBCJBDGLL []*PNGOOBNCNKJ `protobuf:"bytes,15,rep,name=MGEBCJBDGLL,proto3" json:"MGEBCJBDGLL,omitempty"`
	Retcode     uint32         `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DNFMENIODLA *PPCOPNIACJE   `protobuf:"bytes,9,opt,name=DNFMENIODLA,proto3" json:"DNFMENIODLA,omitempty"`
}

func (x *SpaceZooMutateScRsp) Reset() {
	*x = SpaceZooMutateScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpaceZooMutateScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceZooMutateScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceZooMutateScRsp) ProtoMessage() {}

func (x *SpaceZooMutateScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SpaceZooMutateScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceZooMutateScRsp.ProtoReflect.Descriptor instead.
func (*SpaceZooMutateScRsp) Descriptor() ([]byte, []int) {
	return file_SpaceZooMutateScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceZooMutateScRsp) GetIHGHAGMAFJN() bool {
	if x != nil {
		return x.IHGHAGMAFJN
	}
	return false
}

func (x *SpaceZooMutateScRsp) GetMGEBCJBDGLL() []*PNGOOBNCNKJ {
	if x != nil {
		return x.MGEBCJBDGLL
	}
	return nil
}

func (x *SpaceZooMutateScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SpaceZooMutateScRsp) GetDNFMENIODLA() *PPCOPNIACJE {
	if x != nil {
		return x.DNFMENIODLA
	}
	return nil
}

var File_SpaceZooMutateScRsp_proto protoreflect.FileDescriptor

var file_SpaceZooMutateScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4e, 0x47,
	0x4f, 0x4f, 0x42, 0x4e, 0x43, 0x4e, 0x4b, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x50, 0x50, 0x43, 0x4f, 0x50, 0x4e, 0x49, 0x41, 0x43, 0x4a, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x48, 0x47,
	0x48, 0x41, 0x47, 0x4d, 0x41, 0x46, 0x4a, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x49, 0x48, 0x47, 0x48, 0x41, 0x47, 0x4d, 0x41, 0x46, 0x4a, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4d,
	0x47, 0x45, 0x42, 0x43, 0x4a, 0x42, 0x44, 0x47, 0x4c, 0x4c, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x50, 0x4e, 0x47, 0x4f, 0x4f, 0x42, 0x4e, 0x43, 0x4e, 0x4b, 0x4a, 0x52, 0x0b,
	0x4d, 0x47, 0x45, 0x42, 0x43, 0x4a, 0x42, 0x44, 0x47, 0x4c, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x4e, 0x46, 0x4d, 0x45, 0x4e, 0x49,
	0x4f, 0x44, 0x4c, 0x41, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x50, 0x43,
	0x4f, 0x50, 0x4e, 0x49, 0x41, 0x43, 0x4a, 0x45, 0x52, 0x0b, 0x44, 0x4e, 0x46, 0x4d, 0x45, 0x4e,
	0x49, 0x4f, 0x44, 0x4c, 0x41, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpaceZooMutateScRsp_proto_rawDescOnce sync.Once
	file_SpaceZooMutateScRsp_proto_rawDescData = file_SpaceZooMutateScRsp_proto_rawDesc
)

func file_SpaceZooMutateScRsp_proto_rawDescGZIP() []byte {
	file_SpaceZooMutateScRsp_proto_rawDescOnce.Do(func() {
		file_SpaceZooMutateScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpaceZooMutateScRsp_proto_rawDescData)
	})
	return file_SpaceZooMutateScRsp_proto_rawDescData
}

var file_SpaceZooMutateScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpaceZooMutateScRsp_proto_goTypes = []interface{}{
	(*SpaceZooMutateScRsp)(nil), // 0: SpaceZooMutateScRsp
	(*PNGOOBNCNKJ)(nil),         // 1: PNGOOBNCNKJ
	(*PPCOPNIACJE)(nil),         // 2: PPCOPNIACJE
}
var file_SpaceZooMutateScRsp_proto_depIdxs = []int32{
	1, // 0: SpaceZooMutateScRsp.MGEBCJBDGLL:type_name -> PNGOOBNCNKJ
	2, // 1: SpaceZooMutateScRsp.DNFMENIODLA:type_name -> PPCOPNIACJE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SpaceZooMutateScRsp_proto_init() }
func file_SpaceZooMutateScRsp_proto_init() {
	if File_SpaceZooMutateScRsp_proto != nil {
		return
	}
	file_PNGOOBNCNKJ_proto_init()
	file_PPCOPNIACJE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SpaceZooMutateScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceZooMutateScRsp); i {
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
			RawDescriptor: file_SpaceZooMutateScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpaceZooMutateScRsp_proto_goTypes,
		DependencyIndexes: file_SpaceZooMutateScRsp_proto_depIdxs,
		MessageInfos:      file_SpaceZooMutateScRsp_proto_msgTypes,
	}.Build()
	File_SpaceZooMutateScRsp_proto = out.File
	file_SpaceZooMutateScRsp_proto_rawDesc = nil
	file_SpaceZooMutateScRsp_proto_goTypes = nil
	file_SpaceZooMutateScRsp_proto_depIdxs = nil
}
