// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ClockParkHandleWaitOperationScRsp.proto

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

type ClockParkHandleWaitOperationScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32              `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CALNDLDIAOP ClockParkPlayStatus `protobuf:"varint,4,opt,name=CALNDLDIAOP,proto3,enum=ClockParkPlayStatus" json:"CALNDLDIAOP,omitempty"`
	BFAJJLLGCEE uint32              `protobuf:"varint,11,opt,name=BFAJJLLGCEE,proto3" json:"BFAJJLLGCEE,omitempty"`
	EFHFCMKAKAA uint32              `protobuf:"varint,5,opt,name=EFHFCMKAKAA,proto3" json:"EFHFCMKAKAA,omitempty"`
	BattleInfo  *SceneBattleInfo    `protobuf:"bytes,1,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
	BDDHIMGOGLC uint32              `protobuf:"varint,14,opt,name=BDDHIMGOGLC,proto3" json:"BDDHIMGOGLC,omitempty"`
}

func (x *ClockParkHandleWaitOperationScRsp) Reset() {
	*x = ClockParkHandleWaitOperationScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClockParkHandleWaitOperationScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockParkHandleWaitOperationScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockParkHandleWaitOperationScRsp) ProtoMessage() {}

func (x *ClockParkHandleWaitOperationScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ClockParkHandleWaitOperationScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockParkHandleWaitOperationScRsp.ProtoReflect.Descriptor instead.
func (*ClockParkHandleWaitOperationScRsp) Descriptor() ([]byte, []int) {
	return file_ClockParkHandleWaitOperationScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ClockParkHandleWaitOperationScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ClockParkHandleWaitOperationScRsp) GetCALNDLDIAOP() ClockParkPlayStatus {
	if x != nil {
		return x.CALNDLDIAOP
	}
	return ClockParkPlayStatus_CLOCK_PARK_PLAY_NONE
}

func (x *ClockParkHandleWaitOperationScRsp) GetBFAJJLLGCEE() uint32 {
	if x != nil {
		return x.BFAJJLLGCEE
	}
	return 0
}

func (x *ClockParkHandleWaitOperationScRsp) GetEFHFCMKAKAA() uint32 {
	if x != nil {
		return x.EFHFCMKAKAA
	}
	return 0
}

func (x *ClockParkHandleWaitOperationScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

func (x *ClockParkHandleWaitOperationScRsp) GetBDDHIMGOGLC() uint32 {
	if x != nil {
		return x.BDDHIMGOGLC
	}
	return 0
}

var File_ClockParkHandleWaitOperationScRsp_proto protoreflect.FileDescriptor

var file_ClockParkHandleWaitOperationScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x57, 0x61, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x21,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x57,
	0x61, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x43,
	0x41, 0x4c, 0x4e, 0x44, 0x4c, 0x44, 0x49, 0x41, 0x4f, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x50, 0x6c, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x43, 0x41, 0x4c, 0x4e, 0x44, 0x4c, 0x44, 0x49,
	0x41, 0x4f, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46, 0x41, 0x4a, 0x4a, 0x4c, 0x4c, 0x47, 0x43,
	0x45, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x46, 0x41, 0x4a, 0x4a, 0x4c,
	0x4c, 0x47, 0x43, 0x45, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x48, 0x46, 0x43, 0x4d, 0x4b,
	0x41, 0x4b, 0x41, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x46, 0x48, 0x46,
	0x43, 0x4d, 0x4b, 0x41, 0x4b, 0x41, 0x41, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44,
	0x44, 0x48, 0x49, 0x4d, 0x47, 0x4f, 0x47, 0x4c, 0x43, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x44, 0x44, 0x48, 0x49, 0x4d, 0x47, 0x4f, 0x47, 0x4c, 0x43, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClockParkHandleWaitOperationScRsp_proto_rawDescOnce sync.Once
	file_ClockParkHandleWaitOperationScRsp_proto_rawDescData = file_ClockParkHandleWaitOperationScRsp_proto_rawDesc
)

func file_ClockParkHandleWaitOperationScRsp_proto_rawDescGZIP() []byte {
	file_ClockParkHandleWaitOperationScRsp_proto_rawDescOnce.Do(func() {
		file_ClockParkHandleWaitOperationScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClockParkHandleWaitOperationScRsp_proto_rawDescData)
	})
	return file_ClockParkHandleWaitOperationScRsp_proto_rawDescData
}

var file_ClockParkHandleWaitOperationScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClockParkHandleWaitOperationScRsp_proto_goTypes = []interface{}{
	(*ClockParkHandleWaitOperationScRsp)(nil), // 0: ClockParkHandleWaitOperationScRsp
	(ClockParkPlayStatus)(0),                  // 1: ClockParkPlayStatus
	(*SceneBattleInfo)(nil),                   // 2: SceneBattleInfo
}
var file_ClockParkHandleWaitOperationScRsp_proto_depIdxs = []int32{
	1, // 0: ClockParkHandleWaitOperationScRsp.CALNDLDIAOP:type_name -> ClockParkPlayStatus
	2, // 1: ClockParkHandleWaitOperationScRsp.battle_info:type_name -> SceneBattleInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ClockParkHandleWaitOperationScRsp_proto_init() }
func file_ClockParkHandleWaitOperationScRsp_proto_init() {
	if File_ClockParkHandleWaitOperationScRsp_proto != nil {
		return
	}
	file_SceneBattleInfo_proto_init()
	file_ClockParkPlayStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClockParkHandleWaitOperationScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockParkHandleWaitOperationScRsp); i {
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
			RawDescriptor: file_ClockParkHandleWaitOperationScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClockParkHandleWaitOperationScRsp_proto_goTypes,
		DependencyIndexes: file_ClockParkHandleWaitOperationScRsp_proto_depIdxs,
		MessageInfos:      file_ClockParkHandleWaitOperationScRsp_proto_msgTypes,
	}.Build()
	File_ClockParkHandleWaitOperationScRsp_proto = out.File
	file_ClockParkHandleWaitOperationScRsp_proto_rawDesc = nil
	file_ClockParkHandleWaitOperationScRsp_proto_goTypes = nil
	file_ClockParkHandleWaitOperationScRsp_proto_depIdxs = nil
}