// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueAeonInfo.proto

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

type ChessRogueAeonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HNCGPMCBNLH int32                    `protobuf:"varint,11,opt,name=HNCGPMCBNLH,proto3" json:"HNCGPMCBNLH,omitempty"`
	IDMNMAPAHOM *DANEEHMKDKN             `protobuf:"bytes,4,opt,name=IDMNMAPAHOM,proto3" json:"IDMNMAPAHOM,omitempty"`
	AeonInfo    *ChessRogueQueryAeonInfo `protobuf:"bytes,13,opt,name=aeon_info,json=aeonInfo,proto3" json:"aeon_info,omitempty"`
	AeonId      uint32                   `protobuf:"varint,10,opt,name=aeon_id,json=aeonId,proto3" json:"aeon_id,omitempty"`
	AeonIdList  []uint32                 `protobuf:"varint,1,rep,packed,name=aeon_id_list,json=aeonIdList,proto3" json:"aeon_id_list,omitempty"`
}

func (x *ChessRogueAeonInfo) Reset() {
	*x = ChessRogueAeonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueAeonInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueAeonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueAeonInfo) ProtoMessage() {}

func (x *ChessRogueAeonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueAeonInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueAeonInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueAeonInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueAeonInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueAeonInfo) GetHNCGPMCBNLH() int32 {
	if x != nil {
		return x.HNCGPMCBNLH
	}
	return 0
}

func (x *ChessRogueAeonInfo) GetIDMNMAPAHOM() *DANEEHMKDKN {
	if x != nil {
		return x.IDMNMAPAHOM
	}
	return nil
}

func (x *ChessRogueAeonInfo) GetAeonInfo() *ChessRogueQueryAeonInfo {
	if x != nil {
		return x.AeonInfo
	}
	return nil
}

func (x *ChessRogueAeonInfo) GetAeonId() uint32 {
	if x != nil {
		return x.AeonId
	}
	return 0
}

func (x *ChessRogueAeonInfo) GetAeonIdList() []uint32 {
	if x != nil {
		return x.AeonIdList
	}
	return nil
}

var File_ChessRogueAeonInfo_proto protoreflect.FileDescriptor

var file_ChessRogueAeonInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x41, 0x4e, 0x45,
	0x45, 0x48, 0x4d, 0x4b, 0x44, 0x4b, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x65,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a,
	0x12, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4e, 0x43, 0x47, 0x50, 0x4d, 0x43, 0x42, 0x4e,
	0x4c, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x48, 0x4e, 0x43, 0x47, 0x50, 0x4d,
	0x43, 0x42, 0x4e, 0x4c, 0x48, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x44, 0x4d, 0x4e, 0x4d, 0x41, 0x50,
	0x41, 0x48, 0x4f, 0x4d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x41, 0x4e,
	0x45, 0x45, 0x48, 0x4d, 0x4b, 0x44, 0x4b, 0x4e, 0x52, 0x0b, 0x49, 0x44, 0x4d, 0x4e, 0x4d, 0x41,
	0x50, 0x41, 0x48, 0x4f, 0x4d, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x61, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61,
	0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x65, 0x6f,
	0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueAeonInfo_proto_rawDescOnce sync.Once
	file_ChessRogueAeonInfo_proto_rawDescData = file_ChessRogueAeonInfo_proto_rawDesc
)

func file_ChessRogueAeonInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueAeonInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueAeonInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueAeonInfo_proto_rawDescData)
	})
	return file_ChessRogueAeonInfo_proto_rawDescData
}

var file_ChessRogueAeonInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueAeonInfo_proto_goTypes = []interface{}{
	(*ChessRogueAeonInfo)(nil),      // 0: ChessRogueAeonInfo
	(*DANEEHMKDKN)(nil),             // 1: DANEEHMKDKN
	(*ChessRogueQueryAeonInfo)(nil), // 2: ChessRogueQueryAeonInfo
}
var file_ChessRogueAeonInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueAeonInfo.IDMNMAPAHOM:type_name -> DANEEHMKDKN
	2, // 1: ChessRogueAeonInfo.aeon_info:type_name -> ChessRogueQueryAeonInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChessRogueAeonInfo_proto_init() }
func file_ChessRogueAeonInfo_proto_init() {
	if File_ChessRogueAeonInfo_proto != nil {
		return
	}
	file_DANEEHMKDKN_proto_init()
	file_ChessRogueQueryAeonInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueAeonInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueAeonInfo); i {
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
			RawDescriptor: file_ChessRogueAeonInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueAeonInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueAeonInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueAeonInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueAeonInfo_proto = out.File
	file_ChessRogueAeonInfo_proto_rawDesc = nil
	file_ChessRogueAeonInfo_proto_goTypes = nil
	file_ChessRogueAeonInfo_proto_depIdxs = nil
}
