// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetChapterCsReq.proto

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

type GetChapterCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChapterCsReq) Reset() {
	*x = GetChapterCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChapterCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChapterCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChapterCsReq) ProtoMessage() {}

func (x *GetChapterCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetChapterCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChapterCsReq.ProtoReflect.Descriptor instead.
func (*GetChapterCsReq) Descriptor() ([]byte, []int) {
	return file_GetChapterCsReq_proto_rawDescGZIP(), []int{0}
}

var File_GetChapterCsReq_proto protoreflect.FileDescriptor

var file_GetChapterCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetChapterCsReq_proto_rawDescOnce sync.Once
	file_GetChapterCsReq_proto_rawDescData = file_GetChapterCsReq_proto_rawDesc
)

func file_GetChapterCsReq_proto_rawDescGZIP() []byte {
	file_GetChapterCsReq_proto_rawDescOnce.Do(func() {
		file_GetChapterCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChapterCsReq_proto_rawDescData)
	})
	return file_GetChapterCsReq_proto_rawDescData
}

var file_GetChapterCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChapterCsReq_proto_goTypes = []interface{}{
	(*GetChapterCsReq)(nil), // 0: GetChapterCsReq
}
var file_GetChapterCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetChapterCsReq_proto_init() }
func file_GetChapterCsReq_proto_init() {
	if File_GetChapterCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetChapterCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChapterCsReq); i {
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
			RawDescriptor: file_GetChapterCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChapterCsReq_proto_goTypes,
		DependencyIndexes: file_GetChapterCsReq_proto_depIdxs,
		MessageInfos:      file_GetChapterCsReq_proto_msgTypes,
	}.Build()
	File_GetChapterCsReq_proto = out.File
	file_GetChapterCsReq_proto_rawDesc = nil
	file_GetChapterCsReq_proto_goTypes = nil
	file_GetChapterCsReq_proto_depIdxs = nil
}
