// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TextJoinSaveScRsp.proto

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

type TextJoinSaveScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OMIFPKKDLNK uint32 `protobuf:"varint,14,opt,name=OMIFPKKDLNK,proto3" json:"OMIFPKKDLNK,omitempty"`
	OKJFFCBNFAO string `protobuf:"bytes,5,opt,name=OKJFFCBNFAO,proto3" json:"OKJFFCBNFAO,omitempty"`
	OFLNDMFEJBP uint32 `protobuf:"varint,15,opt,name=OFLNDMFEJBP,proto3" json:"OFLNDMFEJBP,omitempty"`
	Retcode     uint32 `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TextJoinSaveScRsp) Reset() {
	*x = TextJoinSaveScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TextJoinSaveScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextJoinSaveScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextJoinSaveScRsp) ProtoMessage() {}

func (x *TextJoinSaveScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TextJoinSaveScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextJoinSaveScRsp.ProtoReflect.Descriptor instead.
func (*TextJoinSaveScRsp) Descriptor() ([]byte, []int) {
	return file_TextJoinSaveScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TextJoinSaveScRsp) GetOMIFPKKDLNK() uint32 {
	if x != nil {
		return x.OMIFPKKDLNK
	}
	return 0
}

func (x *TextJoinSaveScRsp) GetOKJFFCBNFAO() string {
	if x != nil {
		return x.OKJFFCBNFAO
	}
	return ""
}

func (x *TextJoinSaveScRsp) GetOFLNDMFEJBP() uint32 {
	if x != nil {
		return x.OFLNDMFEJBP
	}
	return 0
}

func (x *TextJoinSaveScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TextJoinSaveScRsp_proto protoreflect.FileDescriptor

var file_TextJoinSaveScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x54, 0x65, 0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x54, 0x65,
	0x78, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4d, 0x49, 0x46, 0x50, 0x4b, 0x4b, 0x44, 0x4c, 0x4e, 0x4b, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4d, 0x49, 0x46, 0x50, 0x4b, 0x4b, 0x44, 0x4c, 0x4e,
	0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x4a, 0x46, 0x46, 0x43, 0x42, 0x4e, 0x46, 0x41, 0x4f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x4b, 0x4a, 0x46, 0x46, 0x43, 0x42, 0x4e,
	0x46, 0x41, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x46, 0x4c, 0x4e, 0x44, 0x4d, 0x46, 0x45, 0x4a,
	0x42, 0x50, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x46, 0x4c, 0x4e, 0x44, 0x4d,
	0x46, 0x45, 0x4a, 0x42, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TextJoinSaveScRsp_proto_rawDescOnce sync.Once
	file_TextJoinSaveScRsp_proto_rawDescData = file_TextJoinSaveScRsp_proto_rawDesc
)

func file_TextJoinSaveScRsp_proto_rawDescGZIP() []byte {
	file_TextJoinSaveScRsp_proto_rawDescOnce.Do(func() {
		file_TextJoinSaveScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TextJoinSaveScRsp_proto_rawDescData)
	})
	return file_TextJoinSaveScRsp_proto_rawDescData
}

var file_TextJoinSaveScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TextJoinSaveScRsp_proto_goTypes = []interface{}{
	(*TextJoinSaveScRsp)(nil), // 0: TextJoinSaveScRsp
}
var file_TextJoinSaveScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TextJoinSaveScRsp_proto_init() }
func file_TextJoinSaveScRsp_proto_init() {
	if File_TextJoinSaveScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TextJoinSaveScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextJoinSaveScRsp); i {
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
			RawDescriptor: file_TextJoinSaveScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TextJoinSaveScRsp_proto_goTypes,
		DependencyIndexes: file_TextJoinSaveScRsp_proto_depIdxs,
		MessageInfos:      file_TextJoinSaveScRsp_proto_msgTypes,
	}.Build()
	File_TextJoinSaveScRsp_proto = out.File
	file_TextJoinSaveScRsp_proto_rawDesc = nil
	file_TextJoinSaveScRsp_proto_goTypes = nil
	file_TextJoinSaveScRsp_proto_depIdxs = nil
}
