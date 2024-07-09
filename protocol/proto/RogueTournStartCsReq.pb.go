// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournStartCsReq.proto

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

type RogueTournStartCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId           uint32   `protobuf:"varint,5,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	BaseAvatarIdList []uint32 `protobuf:"varint,12,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	DifficultyIdList []uint32 `protobuf:"varint,11,rep,packed,name=difficulty_id_list,json=difficultyIdList,proto3" json:"difficulty_id_list,omitempty"`
	Week             uint32   `protobuf:"varint,7,opt,name=week,proto3" json:"week,omitempty"`
}

func (x *RogueTournStartCsReq) Reset() {
	*x = RogueTournStartCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournStartCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournStartCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournStartCsReq) ProtoMessage() {}

func (x *RogueTournStartCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournStartCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournStartCsReq.ProtoReflect.Descriptor instead.
func (*RogueTournStartCsReq) Descriptor() ([]byte, []int) {
	return file_RogueTournStartCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournStartCsReq) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *RogueTournStartCsReq) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *RogueTournStartCsReq) GetDifficultyIdList() []uint32 {
	if x != nil {
		return x.DifficultyIdList
	}
	return nil
}

func (x *RogueTournStartCsReq) GetWeek() uint32 {
	if x != nil {
		return x.Week
	}
	return 0
}

var File_RogueTournStartCsReq_proto protoreflect.FileDescriptor

var file_RogueTournStartCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77,
	0x65, 0x65, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RogueTournStartCsReq_proto_rawDescOnce sync.Once
	file_RogueTournStartCsReq_proto_rawDescData = file_RogueTournStartCsReq_proto_rawDesc
)

func file_RogueTournStartCsReq_proto_rawDescGZIP() []byte {
	file_RogueTournStartCsReq_proto_rawDescOnce.Do(func() {
		file_RogueTournStartCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournStartCsReq_proto_rawDescData)
	})
	return file_RogueTournStartCsReq_proto_rawDescData
}

var file_RogueTournStartCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournStartCsReq_proto_goTypes = []interface{}{
	(*RogueTournStartCsReq)(nil), // 0: RogueTournStartCsReq
}
var file_RogueTournStartCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueTournStartCsReq_proto_init() }
func file_RogueTournStartCsReq_proto_init() {
	if File_RogueTournStartCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueTournStartCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournStartCsReq); i {
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
			RawDescriptor: file_RogueTournStartCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournStartCsReq_proto_goTypes,
		DependencyIndexes: file_RogueTournStartCsReq_proto_depIdxs,
		MessageInfos:      file_RogueTournStartCsReq_proto_msgTypes,
	}.Build()
	File_RogueTournStartCsReq_proto = out.File
	file_RogueTournStartCsReq_proto_rawDesc = nil
	file_RogueTournStartCsReq_proto_goTypes = nil
	file_RogueTournStartCsReq_proto_depIdxs = nil
}
