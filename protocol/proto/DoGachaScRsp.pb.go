// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DoGachaScRsp.proto

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

type DoGachaScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PNIDDEENFFH   uint32       `protobuf:"varint,1,opt,name=PNIDDEENFFH,proto3" json:"PNIDDEENFFH,omitempty"`
	Retcode       uint32       `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MIBGEEFGAOO   uint32       `protobuf:"varint,8,opt,name=MIBGEEFGAOO,proto3" json:"MIBGEEFGAOO,omitempty"`
	DDMCEOBBMHB   uint32       `protobuf:"varint,5,opt,name=DDMCEOBBMHB,proto3" json:"DDMCEOBBMHB,omitempty"`
	AGHOIMAJFEI   uint32       `protobuf:"varint,7,opt,name=AGHOIMAJFEI,proto3" json:"AGHOIMAJFEI,omitempty"`
	GachaId       uint32       `protobuf:"varint,12,opt,name=gacha_id,json=gachaId,proto3" json:"gacha_id,omitempty"`
	CeilingNum    uint32       `protobuf:"varint,15,opt,name=ceiling_num,json=ceilingNum,proto3" json:"ceiling_num,omitempty"`
	GachaNum      uint32       `protobuf:"varint,4,opt,name=gacha_num,json=gachaNum,proto3" json:"gacha_num,omitempty"`
	GachaItemList []*GachaItem `protobuf:"bytes,9,rep,name=gacha_item_list,json=gachaItemList,proto3" json:"gacha_item_list,omitempty"`
}

func (x *DoGachaScRsp) Reset() {
	*x = DoGachaScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DoGachaScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoGachaScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoGachaScRsp) ProtoMessage() {}

func (x *DoGachaScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_DoGachaScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoGachaScRsp.ProtoReflect.Descriptor instead.
func (*DoGachaScRsp) Descriptor() ([]byte, []int) {
	return file_DoGachaScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *DoGachaScRsp) GetPNIDDEENFFH() uint32 {
	if x != nil {
		return x.PNIDDEENFFH
	}
	return 0
}

func (x *DoGachaScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *DoGachaScRsp) GetMIBGEEFGAOO() uint32 {
	if x != nil {
		return x.MIBGEEFGAOO
	}
	return 0
}

func (x *DoGachaScRsp) GetDDMCEOBBMHB() uint32 {
	if x != nil {
		return x.DDMCEOBBMHB
	}
	return 0
}

func (x *DoGachaScRsp) GetAGHOIMAJFEI() uint32 {
	if x != nil {
		return x.AGHOIMAJFEI
	}
	return 0
}

func (x *DoGachaScRsp) GetGachaId() uint32 {
	if x != nil {
		return x.GachaId
	}
	return 0
}

func (x *DoGachaScRsp) GetCeilingNum() uint32 {
	if x != nil {
		return x.CeilingNum
	}
	return 0
}

func (x *DoGachaScRsp) GetGachaNum() uint32 {
	if x != nil {
		return x.GachaNum
	}
	return 0
}

func (x *DoGachaScRsp) GetGachaItemList() []*GachaItem {
	if x != nil {
		return x.GachaItemList
	}
	return nil
}

var File_DoGachaScRsp_proto protoreflect.FileDescriptor

var file_DoGachaScRsp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x44, 0x6f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x0c, 0x44, 0x6f, 0x47, 0x61, 0x63, 0x68,
	0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4e, 0x49, 0x44, 0x44, 0x45,
	0x45, 0x4e, 0x46, 0x46, 0x48, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4e, 0x49,
	0x44, 0x44, 0x45, 0x45, 0x4e, 0x46, 0x46, 0x48, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x49, 0x42, 0x47, 0x45, 0x45, 0x46, 0x47, 0x41, 0x4f,
	0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x49, 0x42, 0x47, 0x45, 0x45, 0x46,
	0x47, 0x41, 0x4f, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x44, 0x4d, 0x43, 0x45, 0x4f, 0x42, 0x42,
	0x4d, 0x48, 0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x44, 0x4d, 0x43, 0x45,
	0x4f, 0x42, 0x42, 0x4d, 0x48, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x47, 0x48, 0x4f, 0x49, 0x4d,
	0x41, 0x4a, 0x46, 0x45, 0x49, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x47, 0x48,
	0x4f, 0x49, 0x4d, 0x41, 0x4a, 0x46, 0x45, 0x49, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x65, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x63, 0x68, 0x61, 0x4e, 0x75,
	0x6d, 0x12, 0x32, 0x0a, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x47, 0x61, 0x63,
	0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x67, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DoGachaScRsp_proto_rawDescOnce sync.Once
	file_DoGachaScRsp_proto_rawDescData = file_DoGachaScRsp_proto_rawDesc
)

func file_DoGachaScRsp_proto_rawDescGZIP() []byte {
	file_DoGachaScRsp_proto_rawDescOnce.Do(func() {
		file_DoGachaScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_DoGachaScRsp_proto_rawDescData)
	})
	return file_DoGachaScRsp_proto_rawDescData
}

var file_DoGachaScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DoGachaScRsp_proto_goTypes = []interface{}{
	(*DoGachaScRsp)(nil), // 0: DoGachaScRsp
	(*GachaItem)(nil),    // 1: GachaItem
}
var file_DoGachaScRsp_proto_depIdxs = []int32{
	1, // 0: DoGachaScRsp.gacha_item_list:type_name -> GachaItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DoGachaScRsp_proto_init() }
func file_DoGachaScRsp_proto_init() {
	if File_DoGachaScRsp_proto != nil {
		return
	}
	file_GachaItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DoGachaScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoGachaScRsp); i {
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
			RawDescriptor: file_DoGachaScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DoGachaScRsp_proto_goTypes,
		DependencyIndexes: file_DoGachaScRsp_proto_depIdxs,
		MessageInfos:      file_DoGachaScRsp_proto_msgTypes,
	}.Build()
	File_DoGachaScRsp_proto = out.File
	file_DoGachaScRsp_proto_rawDesc = nil
	file_DoGachaScRsp_proto_goTypes = nil
	file_DoGachaScRsp_proto_depIdxs = nil
}
