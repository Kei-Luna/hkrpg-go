// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MatchBoxingClubOpponentScRsp.proto

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

type MatchBoxingClubOpponentScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JLPEAOCIODN *IFBDAHALCLH `protobuf:"bytes,9,opt,name=JLPEAOCIODN,proto3" json:"JLPEAOCIODN,omitempty"`
	Retcode     uint32       `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MatchBoxingClubOpponentScRsp) Reset() {
	*x = MatchBoxingClubOpponentScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MatchBoxingClubOpponentScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchBoxingClubOpponentScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchBoxingClubOpponentScRsp) ProtoMessage() {}

func (x *MatchBoxingClubOpponentScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MatchBoxingClubOpponentScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchBoxingClubOpponentScRsp.ProtoReflect.Descriptor instead.
func (*MatchBoxingClubOpponentScRsp) Descriptor() ([]byte, []int) {
	return file_MatchBoxingClubOpponentScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MatchBoxingClubOpponentScRsp) GetJLPEAOCIODN() *IFBDAHALCLH {
	if x != nil {
		return x.JLPEAOCIODN
	}
	return nil
}

func (x *MatchBoxingClubOpponentScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_MatchBoxingClubOpponentScRsp_proto protoreflect.FileDescriptor

var file_MatchBoxingClubOpponentScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x75,
	0x62, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x46, 0x42, 0x44, 0x41, 0x48, 0x41, 0x4c, 0x43, 0x4c,
	0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x1c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x42, 0x6f, 0x78, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x75, 0x62, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x4c, 0x50, 0x45, 0x41,
	0x4f, 0x43, 0x49, 0x4f, 0x44, 0x4e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49,
	0x46, 0x42, 0x44, 0x41, 0x48, 0x41, 0x4c, 0x43, 0x4c, 0x48, 0x52, 0x0b, 0x4a, 0x4c, 0x50, 0x45,
	0x41, 0x4f, 0x43, 0x49, 0x4f, 0x44, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_MatchBoxingClubOpponentScRsp_proto_rawDescOnce sync.Once
	file_MatchBoxingClubOpponentScRsp_proto_rawDescData = file_MatchBoxingClubOpponentScRsp_proto_rawDesc
)

func file_MatchBoxingClubOpponentScRsp_proto_rawDescGZIP() []byte {
	file_MatchBoxingClubOpponentScRsp_proto_rawDescOnce.Do(func() {
		file_MatchBoxingClubOpponentScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MatchBoxingClubOpponentScRsp_proto_rawDescData)
	})
	return file_MatchBoxingClubOpponentScRsp_proto_rawDescData
}

var file_MatchBoxingClubOpponentScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MatchBoxingClubOpponentScRsp_proto_goTypes = []interface{}{
	(*MatchBoxingClubOpponentScRsp)(nil), // 0: MatchBoxingClubOpponentScRsp
	(*IFBDAHALCLH)(nil),                  // 1: IFBDAHALCLH
}
var file_MatchBoxingClubOpponentScRsp_proto_depIdxs = []int32{
	1, // 0: MatchBoxingClubOpponentScRsp.JLPEAOCIODN:type_name -> IFBDAHALCLH
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MatchBoxingClubOpponentScRsp_proto_init() }
func file_MatchBoxingClubOpponentScRsp_proto_init() {
	if File_MatchBoxingClubOpponentScRsp_proto != nil {
		return
	}
	file_IFBDAHALCLH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MatchBoxingClubOpponentScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchBoxingClubOpponentScRsp); i {
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
			RawDescriptor: file_MatchBoxingClubOpponentScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MatchBoxingClubOpponentScRsp_proto_goTypes,
		DependencyIndexes: file_MatchBoxingClubOpponentScRsp_proto_depIdxs,
		MessageInfos:      file_MatchBoxingClubOpponentScRsp_proto_msgTypes,
	}.Build()
	File_MatchBoxingClubOpponentScRsp_proto = out.File
	file_MatchBoxingClubOpponentScRsp_proto_rawDesc = nil
	file_MatchBoxingClubOpponentScRsp_proto_goTypes = nil
	file_MatchBoxingClubOpponentScRsp_proto_depIdxs = nil
}