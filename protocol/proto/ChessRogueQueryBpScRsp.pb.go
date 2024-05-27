// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueQueryBpScRsp.proto

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

type ChessRogueQueryBpScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *KDOLHDNKOOC `protobuf:"bytes,12,opt,name=info,proto3" json:"info,omitempty"`
	Retcode uint32       `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *ChessRogueQueryBpScRsp) Reset() {
	*x = ChessRogueQueryBpScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueQueryBpScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueQueryBpScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueQueryBpScRsp) ProtoMessage() {}

func (x *ChessRogueQueryBpScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueQueryBpScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueQueryBpScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueQueryBpScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueQueryBpScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueQueryBpScRsp) GetInfo() *KDOLHDNKOOC {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChessRogueQueryBpScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_ChessRogueQueryBpScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueQueryBpScRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x4b, 0x44, 0x4f, 0x4c, 0x48, 0x44, 0x4e, 0x4b, 0x4f, 0x4f, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x44, 0x4f, 0x4c,
	0x48, 0x44, 0x4e, 0x4b, 0x4f, 0x4f, 0x43, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueQueryBpScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueQueryBpScRsp_proto_rawDescData = file_ChessRogueQueryBpScRsp_proto_rawDesc
)

func file_ChessRogueQueryBpScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueQueryBpScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueQueryBpScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueQueryBpScRsp_proto_rawDescData)
	})
	return file_ChessRogueQueryBpScRsp_proto_rawDescData
}

var file_ChessRogueQueryBpScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueQueryBpScRsp_proto_goTypes = []interface{}{
	(*ChessRogueQueryBpScRsp)(nil), // 0: ChessRogueQueryBpScRsp
	(*KDOLHDNKOOC)(nil),            // 1: KDOLHDNKOOC
}
var file_ChessRogueQueryBpScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueQueryBpScRsp.info:type_name -> KDOLHDNKOOC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessRogueQueryBpScRsp_proto_init() }
func file_ChessRogueQueryBpScRsp_proto_init() {
	if File_ChessRogueQueryBpScRsp_proto != nil {
		return
	}
	file_KDOLHDNKOOC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueQueryBpScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueQueryBpScRsp); i {
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
			RawDescriptor: file_ChessRogueQueryBpScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueQueryBpScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueQueryBpScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueQueryBpScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueQueryBpScRsp_proto = out.File
	file_ChessRogueQueryBpScRsp_proto_rawDesc = nil
	file_ChessRogueQueryBpScRsp_proto_goTypes = nil
	file_ChessRogueQueryBpScRsp_proto_depIdxs = nil
}
