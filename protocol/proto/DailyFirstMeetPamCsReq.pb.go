// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DailyFirstMeetPamCsReq.proto

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

type DailyFirstMeetPamCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DailyFirstMeetPamCsReq) Reset() {
	*x = DailyFirstMeetPamCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyFirstMeetPamCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyFirstMeetPamCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyFirstMeetPamCsReq) ProtoMessage() {}

func (x *DailyFirstMeetPamCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_DailyFirstMeetPamCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyFirstMeetPamCsReq.ProtoReflect.Descriptor instead.
func (*DailyFirstMeetPamCsReq) Descriptor() ([]byte, []int) {
	return file_DailyFirstMeetPamCsReq_proto_rawDescGZIP(), []int{0}
}

var File_DailyFirstMeetPamCsReq_proto protoreflect.FileDescriptor

var file_DailyFirstMeetPamCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74,
	0x50, 0x61, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18,
	0x0a, 0x16, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74,
	0x50, 0x61, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DailyFirstMeetPamCsReq_proto_rawDescOnce sync.Once
	file_DailyFirstMeetPamCsReq_proto_rawDescData = file_DailyFirstMeetPamCsReq_proto_rawDesc
)

func file_DailyFirstMeetPamCsReq_proto_rawDescGZIP() []byte {
	file_DailyFirstMeetPamCsReq_proto_rawDescOnce.Do(func() {
		file_DailyFirstMeetPamCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyFirstMeetPamCsReq_proto_rawDescData)
	})
	return file_DailyFirstMeetPamCsReq_proto_rawDescData
}

var file_DailyFirstMeetPamCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyFirstMeetPamCsReq_proto_goTypes = []interface{}{
	(*DailyFirstMeetPamCsReq)(nil), // 0: DailyFirstMeetPamCsReq
}
var file_DailyFirstMeetPamCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DailyFirstMeetPamCsReq_proto_init() }
func file_DailyFirstMeetPamCsReq_proto_init() {
	if File_DailyFirstMeetPamCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DailyFirstMeetPamCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyFirstMeetPamCsReq); i {
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
			RawDescriptor: file_DailyFirstMeetPamCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyFirstMeetPamCsReq_proto_goTypes,
		DependencyIndexes: file_DailyFirstMeetPamCsReq_proto_depIdxs,
		MessageInfos:      file_DailyFirstMeetPamCsReq_proto_msgTypes,
	}.Build()
	File_DailyFirstMeetPamCsReq_proto = out.File
	file_DailyFirstMeetPamCsReq_proto_rawDesc = nil
	file_DailyFirstMeetPamCsReq_proto_goTypes = nil
	file_DailyFirstMeetPamCsReq_proto_depIdxs = nil
}
