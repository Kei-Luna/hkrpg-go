// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GachaInfo.proto

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

type GachaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndTime            int64         `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	PrizeItemList      []uint32      `protobuf:"varint,11,rep,packed,name=prize_item_list,json=prizeItemList,proto3" json:"prize_item_list,omitempty"`
	DropHistoryWebview string        `protobuf:"bytes,4,opt,name=drop_history_webview,json=dropHistoryWebview,proto3" json:"drop_history_webview,omitempty"`
	GachaInfoList      *GachaCeiling `protobuf:"bytes,5,opt,name=gacha_info_list,json=gachaInfoList,proto3" json:"gacha_info_list,omitempty"`
	BeginTime          int64         `protobuf:"varint,1,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	GachaId            uint32        `protobuf:"varint,10,opt,name=gacha_id,json=gachaId,proto3" json:"gacha_id,omitempty"`
	AGHOIMAJFEI        uint32        `protobuf:"varint,14,opt,name=AGHOIMAJFEI,proto3" json:"AGHOIMAJFEI,omitempty"`
	DetailWebview      string        `protobuf:"bytes,6,opt,name=detail_webview,json=detailWebview,proto3" json:"detail_webview,omitempty"`
	DDMCEOBBMHB        uint32        `protobuf:"varint,9,opt,name=DDMCEOBBMHB,proto3" json:"DDMCEOBBMHB,omitempty"`
	ItemDetailList     []uint32      `protobuf:"varint,13,rep,packed,name=item_detail_list,json=itemDetailList,proto3" json:"item_detail_list,omitempty"`
}

func (x *GachaInfo) Reset() {
	*x = GachaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaInfo) ProtoMessage() {}

func (x *GachaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GachaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaInfo.ProtoReflect.Descriptor instead.
func (*GachaInfo) Descriptor() ([]byte, []int) {
	return file_GachaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GachaInfo) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GachaInfo) GetPrizeItemList() []uint32 {
	if x != nil {
		return x.PrizeItemList
	}
	return nil
}

func (x *GachaInfo) GetDropHistoryWebview() string {
	if x != nil {
		return x.DropHistoryWebview
	}
	return ""
}

func (x *GachaInfo) GetGachaInfoList() *GachaCeiling {
	if x != nil {
		return x.GachaInfoList
	}
	return nil
}

func (x *GachaInfo) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *GachaInfo) GetGachaId() uint32 {
	if x != nil {
		return x.GachaId
	}
	return 0
}

func (x *GachaInfo) GetAGHOIMAJFEI() uint32 {
	if x != nil {
		return x.AGHOIMAJFEI
	}
	return 0
}

func (x *GachaInfo) GetDetailWebview() string {
	if x != nil {
		return x.DetailWebview
	}
	return ""
}

func (x *GachaInfo) GetDDMCEOBBMHB() uint32 {
	if x != nil {
		return x.DDMCEOBBMHB
	}
	return 0
}

func (x *GachaInfo) GetItemDetailList() []uint32 {
	if x != nil {
		return x.ItemDetailList
	}
	return nil
}

var File_GachaInfo_proto protoreflect.FileDescriptor

var file_GachaInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x47, 0x61, 0x63, 0x68, 0x61, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x09, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x72, 0x6f, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x57, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x12, 0x35, 0x0a, 0x0f, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x47, 0x61, 0x63, 0x68, 0x61, 0x43, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x52, 0x0d, 0x67, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x67, 0x61, 0x63, 0x68, 0x61, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x47, 0x48,
	0x4f, 0x49, 0x4d, 0x41, 0x4a, 0x46, 0x45, 0x49, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x41, 0x47, 0x48, 0x4f, 0x49, 0x4d, 0x41, 0x4a, 0x46, 0x45, 0x49, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x77, 0x65, 0x62, 0x76, 0x69, 0x65, 0x77, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x65, 0x62, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x44, 0x4d, 0x43, 0x45, 0x4f, 0x42, 0x42, 0x4d, 0x48,
	0x42, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x44, 0x4d, 0x43, 0x45, 0x4f, 0x42,
	0x42, 0x4d, 0x48, 0x42, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e,
	0x69, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GachaInfo_proto_rawDescOnce sync.Once
	file_GachaInfo_proto_rawDescData = file_GachaInfo_proto_rawDesc
)

func file_GachaInfo_proto_rawDescGZIP() []byte {
	file_GachaInfo_proto_rawDescOnce.Do(func() {
		file_GachaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaInfo_proto_rawDescData)
	})
	return file_GachaInfo_proto_rawDescData
}

var file_GachaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaInfo_proto_goTypes = []interface{}{
	(*GachaInfo)(nil),    // 0: GachaInfo
	(*GachaCeiling)(nil), // 1: GachaCeiling
}
var file_GachaInfo_proto_depIdxs = []int32{
	1, // 0: GachaInfo.gacha_info_list:type_name -> GachaCeiling
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GachaInfo_proto_init() }
func file_GachaInfo_proto_init() {
	if File_GachaInfo_proto != nil {
		return
	}
	file_GachaCeiling_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GachaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaInfo); i {
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
			RawDescriptor: file_GachaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaInfo_proto_goTypes,
		DependencyIndexes: file_GachaInfo_proto_depIdxs,
		MessageInfos:      file_GachaInfo_proto_msgTypes,
	}.Build()
	File_GachaInfo_proto = out.File
	file_GachaInfo_proto_rawDesc = nil
	file_GachaInfo_proto_goTypes = nil
	file_GachaInfo_proto_depIdxs = nil
}
