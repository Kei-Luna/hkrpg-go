// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeCityShopRewardCsReq.proto

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

type TakeCityShopRewardCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level  uint32 `protobuf:"varint,9,opt,name=level,proto3" json:"level,omitempty"`
	ShopId uint32 `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *TakeCityShopRewardCsReq) Reset() {
	*x = TakeCityShopRewardCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeCityShopRewardCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeCityShopRewardCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeCityShopRewardCsReq) ProtoMessage() {}

func (x *TakeCityShopRewardCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_TakeCityShopRewardCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeCityShopRewardCsReq.ProtoReflect.Descriptor instead.
func (*TakeCityShopRewardCsReq) Descriptor() ([]byte, []int) {
	return file_TakeCityShopRewardCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *TakeCityShopRewardCsReq) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *TakeCityShopRewardCsReq) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

var File_TakeCityShopRewardCsReq_proto protoreflect.FileDescriptor

var file_TakeCityShopRewardCsReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x43, 0x69, 0x74, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x48, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x43, 0x69, 0x74, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeCityShopRewardCsReq_proto_rawDescOnce sync.Once
	file_TakeCityShopRewardCsReq_proto_rawDescData = file_TakeCityShopRewardCsReq_proto_rawDesc
)

func file_TakeCityShopRewardCsReq_proto_rawDescGZIP() []byte {
	file_TakeCityShopRewardCsReq_proto_rawDescOnce.Do(func() {
		file_TakeCityShopRewardCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeCityShopRewardCsReq_proto_rawDescData)
	})
	return file_TakeCityShopRewardCsReq_proto_rawDescData
}

var file_TakeCityShopRewardCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeCityShopRewardCsReq_proto_goTypes = []interface{}{
	(*TakeCityShopRewardCsReq)(nil), // 0: TakeCityShopRewardCsReq
}
var file_TakeCityShopRewardCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeCityShopRewardCsReq_proto_init() }
func file_TakeCityShopRewardCsReq_proto_init() {
	if File_TakeCityShopRewardCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeCityShopRewardCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeCityShopRewardCsReq); i {
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
			RawDescriptor: file_TakeCityShopRewardCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeCityShopRewardCsReq_proto_goTypes,
		DependencyIndexes: file_TakeCityShopRewardCsReq_proto_depIdxs,
		MessageInfos:      file_TakeCityShopRewardCsReq_proto_msgTypes,
	}.Build()
	File_TakeCityShopRewardCsReq_proto = out.File
	file_TakeCityShopRewardCsReq_proto_rawDesc = nil
	file_TakeCityShopRewardCsReq_proto_goTypes = nil
	file_TakeCityShopRewardCsReq_proto_depIdxs = nil
}
