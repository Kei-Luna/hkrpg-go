// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SettleTrackPhotoStageCsReq.proto

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

type SettleTrackPhotoStageCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostTime    uint32         `protobuf:"varint,6,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	NFBNCOGLIBK []*LPPPJAPJLAE `protobuf:"bytes,7,rep,name=NFBNCOGLIBK,proto3" json:"NFBNCOGLIBK,omitempty"`
	StageId     uint32         `protobuf:"varint,3,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *SettleTrackPhotoStageCsReq) Reset() {
	*x = SettleTrackPhotoStageCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SettleTrackPhotoStageCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleTrackPhotoStageCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleTrackPhotoStageCsReq) ProtoMessage() {}

func (x *SettleTrackPhotoStageCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SettleTrackPhotoStageCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleTrackPhotoStageCsReq.ProtoReflect.Descriptor instead.
func (*SettleTrackPhotoStageCsReq) Descriptor() ([]byte, []int) {
	return file_SettleTrackPhotoStageCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SettleTrackPhotoStageCsReq) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *SettleTrackPhotoStageCsReq) GetNFBNCOGLIBK() []*LPPPJAPJLAE {
	if x != nil {
		return x.NFBNCOGLIBK
	}
	return nil
}

func (x *SettleTrackPhotoStageCsReq) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

var File_SettleTrackPhotoStageCsReq_proto protoreflect.FileDescriptor

var file_SettleTrackPhotoStageCsReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x50, 0x50, 0x50, 0x4a, 0x41, 0x50, 0x4a, 0x4c, 0x41, 0x45, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x46, 0x42, 0x4e, 0x43, 0x4f, 0x47, 0x4c, 0x49, 0x42, 0x4b,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x50, 0x50, 0x4a, 0x41, 0x50,
	0x4a, 0x4c, 0x41, 0x45, 0x52, 0x0b, 0x4e, 0x46, 0x42, 0x4e, 0x43, 0x4f, 0x47, 0x4c, 0x49, 0x42,
	0x4b, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SettleTrackPhotoStageCsReq_proto_rawDescOnce sync.Once
	file_SettleTrackPhotoStageCsReq_proto_rawDescData = file_SettleTrackPhotoStageCsReq_proto_rawDesc
)

func file_SettleTrackPhotoStageCsReq_proto_rawDescGZIP() []byte {
	file_SettleTrackPhotoStageCsReq_proto_rawDescOnce.Do(func() {
		file_SettleTrackPhotoStageCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SettleTrackPhotoStageCsReq_proto_rawDescData)
	})
	return file_SettleTrackPhotoStageCsReq_proto_rawDescData
}

var file_SettleTrackPhotoStageCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SettleTrackPhotoStageCsReq_proto_goTypes = []interface{}{
	(*SettleTrackPhotoStageCsReq)(nil), // 0: SettleTrackPhotoStageCsReq
	(*LPPPJAPJLAE)(nil),                // 1: LPPPJAPJLAE
}
var file_SettleTrackPhotoStageCsReq_proto_depIdxs = []int32{
	1, // 0: SettleTrackPhotoStageCsReq.NFBNCOGLIBK:type_name -> LPPPJAPJLAE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SettleTrackPhotoStageCsReq_proto_init() }
func file_SettleTrackPhotoStageCsReq_proto_init() {
	if File_SettleTrackPhotoStageCsReq_proto != nil {
		return
	}
	file_LPPPJAPJLAE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SettleTrackPhotoStageCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleTrackPhotoStageCsReq); i {
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
			RawDescriptor: file_SettleTrackPhotoStageCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SettleTrackPhotoStageCsReq_proto_goTypes,
		DependencyIndexes: file_SettleTrackPhotoStageCsReq_proto_depIdxs,
		MessageInfos:      file_SettleTrackPhotoStageCsReq_proto_msgTypes,
	}.Build()
	File_SettleTrackPhotoStageCsReq_proto = out.File
	file_SettleTrackPhotoStageCsReq_proto_rawDesc = nil
	file_SettleTrackPhotoStageCsReq_proto_goTypes = nil
	file_SettleTrackPhotoStageCsReq_proto_depIdxs = nil
}