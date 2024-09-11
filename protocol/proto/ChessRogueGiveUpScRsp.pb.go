// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueGiveUpScRsp.proto

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

type ChessRogueGiveUpScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishInfo    *ChessRogueFinishInfo `protobuf:"bytes,10,opt,name=finish_info,json=finishInfo,proto3" json:"finish_info,omitempty"`
	RogueAeonInfo *ChessRogueAeonInfo   `protobuf:"bytes,1,opt,name=rogue_aeon_info,json=rogueAeonInfo,proto3" json:"rogue_aeon_info,omitempty"`
	Retcode       uint32                `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	QueryInfo     *ChessRogueQueryInfo  `protobuf:"bytes,11,opt,name=query_info,json=queryInfo,proto3" json:"query_info,omitempty"`
	RogueGetInfo  *ChessRogueGetInfo    `protobuf:"bytes,13,opt,name=rogue_get_info,json=rogueGetInfo,proto3" json:"rogue_get_info,omitempty"`
	StageInfo     *ChessRogueInfo       `protobuf:"bytes,6,opt,name=stage_info,json=stageInfo,proto3" json:"stage_info,omitempty"`
}

func (x *ChessRogueGiveUpScRsp) Reset() {
	*x = ChessRogueGiveUpScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueGiveUpScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueGiveUpScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueGiveUpScRsp) ProtoMessage() {}

func (x *ChessRogueGiveUpScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueGiveUpScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueGiveUpScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueGiveUpScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueGiveUpScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueGiveUpScRsp) GetFinishInfo() *ChessRogueFinishInfo {
	if x != nil {
		return x.FinishInfo
	}
	return nil
}

func (x *ChessRogueGiveUpScRsp) GetRogueAeonInfo() *ChessRogueAeonInfo {
	if x != nil {
		return x.RogueAeonInfo
	}
	return nil
}

func (x *ChessRogueGiveUpScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChessRogueGiveUpScRsp) GetQueryInfo() *ChessRogueQueryInfo {
	if x != nil {
		return x.QueryInfo
	}
	return nil
}

func (x *ChessRogueGiveUpScRsp) GetRogueGetInfo() *ChessRogueGetInfo {
	if x != nil {
		return x.RogueGetInfo
	}
	return nil
}

func (x *ChessRogueGiveUpScRsp) GetStageInfo() *ChessRogueInfo {
	if x != nil {
		return x.StageInfo
	}
	return nil
}

var File_ChessRogueGiveUpScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueGiveUpScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x69, 0x76, 0x65,
	0x55, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02,
	0x0a, 0x15, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x69, 0x76, 0x65,
	0x55, 0x70, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3b, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0e, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueGiveUpScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueGiveUpScRsp_proto_rawDescData = file_ChessRogueGiveUpScRsp_proto_rawDesc
)

func file_ChessRogueGiveUpScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueGiveUpScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueGiveUpScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueGiveUpScRsp_proto_rawDescData)
	})
	return file_ChessRogueGiveUpScRsp_proto_rawDescData
}

var file_ChessRogueGiveUpScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueGiveUpScRsp_proto_goTypes = []interface{}{
	(*ChessRogueGiveUpScRsp)(nil), // 0: ChessRogueGiveUpScRsp
	(*ChessRogueFinishInfo)(nil),  // 1: ChessRogueFinishInfo
	(*ChessRogueAeonInfo)(nil),    // 2: ChessRogueAeonInfo
	(*ChessRogueQueryInfo)(nil),   // 3: ChessRogueQueryInfo
	(*ChessRogueGetInfo)(nil),     // 4: ChessRogueGetInfo
	(*ChessRogueInfo)(nil),        // 5: ChessRogueInfo
}
var file_ChessRogueGiveUpScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueGiveUpScRsp.finish_info:type_name -> ChessRogueFinishInfo
	2, // 1: ChessRogueGiveUpScRsp.rogue_aeon_info:type_name -> ChessRogueAeonInfo
	3, // 2: ChessRogueGiveUpScRsp.query_info:type_name -> ChessRogueQueryInfo
	4, // 3: ChessRogueGiveUpScRsp.rogue_get_info:type_name -> ChessRogueGetInfo
	5, // 4: ChessRogueGiveUpScRsp.stage_info:type_name -> ChessRogueInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ChessRogueGiveUpScRsp_proto_init() }
func file_ChessRogueGiveUpScRsp_proto_init() {
	if File_ChessRogueGiveUpScRsp_proto != nil {
		return
	}
	file_ChessRogueAeonInfo_proto_init()
	file_ChessRogueInfo_proto_init()
	file_ChessRogueFinishInfo_proto_init()
	file_ChessRogueQueryInfo_proto_init()
	file_ChessRogueGetInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueGiveUpScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueGiveUpScRsp); i {
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
			RawDescriptor: file_ChessRogueGiveUpScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueGiveUpScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueGiveUpScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueGiveUpScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueGiveUpScRsp_proto = out.File
	file_ChessRogueGiveUpScRsp_proto_rawDesc = nil
	file_ChessRogueGiveUpScRsp_proto_goTypes = nil
	file_ChessRogueGiveUpScRsp_proto_depIdxs = nil
}
