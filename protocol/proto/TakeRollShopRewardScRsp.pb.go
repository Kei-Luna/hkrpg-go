// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeRollShopRewardScRsp.proto

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

type TakeRollShopRewardScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode              uint32    `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GroupTypeFieldNumber uint32    `protobuf:"varint,10,opt,name=GroupTypeFieldNumber,proto3" json:"GroupTypeFieldNumber,omitempty"`
	Reward               *ItemList `protobuf:"bytes,15,opt,name=reward,proto3" json:"reward,omitempty"`
	RollshopId           uint32    `protobuf:"varint,3,opt,name=Rollshop_id,json=RollshopId,proto3" json:"Rollshop_id,omitempty"`
}

func (x *TakeRollShopRewardScRsp) Reset() {
	*x = TakeRollShopRewardScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeRollShopRewardScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeRollShopRewardScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeRollShopRewardScRsp) ProtoMessage() {}

func (x *TakeRollShopRewardScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeRollShopRewardScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeRollShopRewardScRsp.ProtoReflect.Descriptor instead.
func (*TakeRollShopRewardScRsp) Descriptor() ([]byte, []int) {
	return file_TakeRollShopRewardScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeRollShopRewardScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TakeRollShopRewardScRsp) GetGroupTypeFieldNumber() uint32 {
	if x != nil {
		return x.GroupTypeFieldNumber
	}
	return 0
}

func (x *TakeRollShopRewardScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *TakeRollShopRewardScRsp) GetRollshopId() uint32 {
	if x != nil {
		return x.RollshopId
	}
	return 0
}

var File_TakeRollShopRewardScRsp_proto protoreflect.FileDescriptor

var file_TakeRollShopRewardScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xab, 0x01, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeRollShopRewardScRsp_proto_rawDescOnce sync.Once
	file_TakeRollShopRewardScRsp_proto_rawDescData = file_TakeRollShopRewardScRsp_proto_rawDesc
)

func file_TakeRollShopRewardScRsp_proto_rawDescGZIP() []byte {
	file_TakeRollShopRewardScRsp_proto_rawDescOnce.Do(func() {
		file_TakeRollShopRewardScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeRollShopRewardScRsp_proto_rawDescData)
	})
	return file_TakeRollShopRewardScRsp_proto_rawDescData
}

var file_TakeRollShopRewardScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeRollShopRewardScRsp_proto_goTypes = []interface{}{
	(*TakeRollShopRewardScRsp)(nil), // 0: TakeRollShopRewardScRsp
	(*ItemList)(nil),                // 1: ItemList
}
var file_TakeRollShopRewardScRsp_proto_depIdxs = []int32{
	1, // 0: TakeRollShopRewardScRsp.reward:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TakeRollShopRewardScRsp_proto_init() }
func file_TakeRollShopRewardScRsp_proto_init() {
	if File_TakeRollShopRewardScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeRollShopRewardScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeRollShopRewardScRsp); i {
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
			RawDescriptor: file_TakeRollShopRewardScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeRollShopRewardScRsp_proto_goTypes,
		DependencyIndexes: file_TakeRollShopRewardScRsp_proto_depIdxs,
		MessageInfos:      file_TakeRollShopRewardScRsp_proto_msgTypes,
	}.Build()
	File_TakeRollShopRewardScRsp_proto = out.File
	file_TakeRollShopRewardScRsp_proto_rawDesc = nil
	file_TakeRollShopRewardScRsp_proto_goTypes = nil
	file_TakeRollShopRewardScRsp_proto_depIdxs = nil
}