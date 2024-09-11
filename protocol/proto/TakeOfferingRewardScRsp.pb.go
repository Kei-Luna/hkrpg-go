// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeOfferingRewardScRsp.proto

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

type TakeOfferingRewardScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward      *ItemList    `protobuf:"bytes,2,opt,name=reward,proto3" json:"reward,omitempty"`
	FAPLCOFOCMD *EPMNGCPCBKL `protobuf:"bytes,3,opt,name=FAPLCOFOCMD,proto3" json:"FAPLCOFOCMD,omitempty"`
	Retcode     uint32       `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TakeOfferingRewardScRsp) Reset() {
	*x = TakeOfferingRewardScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeOfferingRewardScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeOfferingRewardScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeOfferingRewardScRsp) ProtoMessage() {}

func (x *TakeOfferingRewardScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeOfferingRewardScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeOfferingRewardScRsp.ProtoReflect.Descriptor instead.
func (*TakeOfferingRewardScRsp) Descriptor() ([]byte, []int) {
	return file_TakeOfferingRewardScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeOfferingRewardScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *TakeOfferingRewardScRsp) GetFAPLCOFOCMD() *EPMNGCPCBKL {
	if x != nil {
		return x.FAPLCOFOCMD
	}
	return nil
}

func (x *TakeOfferingRewardScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TakeOfferingRewardScRsp_proto protoreflect.FileDescriptor

var file_TakeOfferingRewardScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x45, 0x50, 0x4d, 0x4e, 0x47, 0x43, 0x50, 0x43, 0x42, 0x4b, 0x4c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x21,
	0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x41, 0x50, 0x4c, 0x43, 0x4f, 0x46, 0x4f, 0x43, 0x4d, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x50, 0x4d, 0x4e, 0x47, 0x43, 0x50,
	0x43, 0x42, 0x4b, 0x4c, 0x52, 0x0b, 0x46, 0x41, 0x50, 0x4c, 0x43, 0x4f, 0x46, 0x4f, 0x43, 0x4d,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TakeOfferingRewardScRsp_proto_rawDescOnce sync.Once
	file_TakeOfferingRewardScRsp_proto_rawDescData = file_TakeOfferingRewardScRsp_proto_rawDesc
)

func file_TakeOfferingRewardScRsp_proto_rawDescGZIP() []byte {
	file_TakeOfferingRewardScRsp_proto_rawDescOnce.Do(func() {
		file_TakeOfferingRewardScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeOfferingRewardScRsp_proto_rawDescData)
	})
	return file_TakeOfferingRewardScRsp_proto_rawDescData
}

var file_TakeOfferingRewardScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeOfferingRewardScRsp_proto_goTypes = []interface{}{
	(*TakeOfferingRewardScRsp)(nil), // 0: TakeOfferingRewardScRsp
	(*ItemList)(nil),                // 1: ItemList
	(*EPMNGCPCBKL)(nil),             // 2: EPMNGCPCBKL
}
var file_TakeOfferingRewardScRsp_proto_depIdxs = []int32{
	1, // 0: TakeOfferingRewardScRsp.reward:type_name -> ItemList
	2, // 1: TakeOfferingRewardScRsp.FAPLCOFOCMD:type_name -> EPMNGCPCBKL
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TakeOfferingRewardScRsp_proto_init() }
func file_TakeOfferingRewardScRsp_proto_init() {
	if File_TakeOfferingRewardScRsp_proto != nil {
		return
	}
	file_EPMNGCPCBKL_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeOfferingRewardScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeOfferingRewardScRsp); i {
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
			RawDescriptor: file_TakeOfferingRewardScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeOfferingRewardScRsp_proto_goTypes,
		DependencyIndexes: file_TakeOfferingRewardScRsp_proto_depIdxs,
		MessageInfos:      file_TakeOfferingRewardScRsp_proto_msgTypes,
	}.Build()
	File_TakeOfferingRewardScRsp_proto = out.File
	file_TakeOfferingRewardScRsp_proto_rawDesc = nil
	file_TakeOfferingRewardScRsp_proto_goTypes = nil
	file_TakeOfferingRewardScRsp_proto_depIdxs = nil
}
