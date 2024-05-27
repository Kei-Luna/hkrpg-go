// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TelevisionActivityBattleEndScNotify.proto

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

type TelevisionActivityBattleEndScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OPPIGGHCCHO uint32                  `protobuf:"varint,8,opt,name=OPPIGGHCCHO,proto3" json:"OPPIGGHCCHO,omitempty"`
	PCLBPMHICGL *TelevisionActivityData `protobuf:"bytes,12,opt,name=PCLBPMHICGL,proto3" json:"PCLBPMHICGL,omitempty"`
	CAAPHPHGJGG uint32                  `protobuf:"varint,15,opt,name=CAAPHPHGJGG,proto3" json:"CAAPHPHGJGG,omitempty"`
	TotalScore  uint32                  `protobuf:"varint,10,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"`
	NEPEKOENFCM uint32                  `protobuf:"varint,3,opt,name=NEPEKOENFCM,proto3" json:"NEPEKOENFCM,omitempty"`
}

func (x *TelevisionActivityBattleEndScNotify) Reset() {
	*x = TelevisionActivityBattleEndScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TelevisionActivityBattleEndScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelevisionActivityBattleEndScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelevisionActivityBattleEndScNotify) ProtoMessage() {}

func (x *TelevisionActivityBattleEndScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TelevisionActivityBattleEndScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelevisionActivityBattleEndScNotify.ProtoReflect.Descriptor instead.
func (*TelevisionActivityBattleEndScNotify) Descriptor() ([]byte, []int) {
	return file_TelevisionActivityBattleEndScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TelevisionActivityBattleEndScNotify) GetOPPIGGHCCHO() uint32 {
	if x != nil {
		return x.OPPIGGHCCHO
	}
	return 0
}

func (x *TelevisionActivityBattleEndScNotify) GetPCLBPMHICGL() *TelevisionActivityData {
	if x != nil {
		return x.PCLBPMHICGL
	}
	return nil
}

func (x *TelevisionActivityBattleEndScNotify) GetCAAPHPHGJGG() uint32 {
	if x != nil {
		return x.CAAPHPHGJGG
	}
	return 0
}

func (x *TelevisionActivityBattleEndScNotify) GetTotalScore() uint32 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *TelevisionActivityBattleEndScNotify) GetNEPEKOENFCM() uint32 {
	if x != nil {
		return x.NEPEKOENFCM
	}
	return 0
}

var File_TelevisionActivityBattleEndScNotify_proto protoreflect.FileDescriptor

var file_TelevisionActivityBattleEndScNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x54, 0x65, 0x6c, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x54, 0x65, 0x6c,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x23, 0x54, 0x65,
	0x6c, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x50, 0x50, 0x49, 0x47, 0x47, 0x48, 0x43, 0x43, 0x48, 0x4f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x50, 0x50, 0x49, 0x47, 0x47, 0x48, 0x43,
	0x43, 0x48, 0x4f, 0x12, 0x39, 0x0a, 0x0b, 0x50, 0x43, 0x4c, 0x42, 0x50, 0x4d, 0x48, 0x49, 0x43,
	0x47, 0x4c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0b, 0x50, 0x43, 0x4c, 0x42, 0x50, 0x4d, 0x48, 0x49, 0x43, 0x47, 0x4c, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x41, 0x41, 0x50, 0x48, 0x50, 0x48, 0x47, 0x4a, 0x47, 0x47, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x41, 0x41, 0x50, 0x48, 0x50, 0x48, 0x47, 0x4a, 0x47, 0x47,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x45, 0x50, 0x45, 0x4b, 0x4f, 0x45, 0x4e, 0x46, 0x43, 0x4d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x45, 0x50, 0x45, 0x4b, 0x4f, 0x45, 0x4e,
	0x46, 0x43, 0x4d, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TelevisionActivityBattleEndScNotify_proto_rawDescOnce sync.Once
	file_TelevisionActivityBattleEndScNotify_proto_rawDescData = file_TelevisionActivityBattleEndScNotify_proto_rawDesc
)

func file_TelevisionActivityBattleEndScNotify_proto_rawDescGZIP() []byte {
	file_TelevisionActivityBattleEndScNotify_proto_rawDescOnce.Do(func() {
		file_TelevisionActivityBattleEndScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TelevisionActivityBattleEndScNotify_proto_rawDescData)
	})
	return file_TelevisionActivityBattleEndScNotify_proto_rawDescData
}

var file_TelevisionActivityBattleEndScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TelevisionActivityBattleEndScNotify_proto_goTypes = []interface{}{
	(*TelevisionActivityBattleEndScNotify)(nil), // 0: TelevisionActivityBattleEndScNotify
	(*TelevisionActivityData)(nil),              // 1: TelevisionActivityData
}
var file_TelevisionActivityBattleEndScNotify_proto_depIdxs = []int32{
	1, // 0: TelevisionActivityBattleEndScNotify.PCLBPMHICGL:type_name -> TelevisionActivityData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TelevisionActivityBattleEndScNotify_proto_init() }
func file_TelevisionActivityBattleEndScNotify_proto_init() {
	if File_TelevisionActivityBattleEndScNotify_proto != nil {
		return
	}
	file_TelevisionActivityData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TelevisionActivityBattleEndScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelevisionActivityBattleEndScNotify); i {
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
			RawDescriptor: file_TelevisionActivityBattleEndScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TelevisionActivityBattleEndScNotify_proto_goTypes,
		DependencyIndexes: file_TelevisionActivityBattleEndScNotify_proto_depIdxs,
		MessageInfos:      file_TelevisionActivityBattleEndScNotify_proto_msgTypes,
	}.Build()
	File_TelevisionActivityBattleEndScNotify_proto = out.File
	file_TelevisionActivityBattleEndScNotify_proto_rawDesc = nil
	file_TelevisionActivityBattleEndScNotify_proto_goTypes = nil
	file_TelevisionActivityBattleEndScNotify_proto_depIdxs = nil
}
