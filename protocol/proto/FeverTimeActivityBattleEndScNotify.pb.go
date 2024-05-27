// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FeverTimeActivityBattleEndScNotify.proto

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

type FeverTimeActivityBattleEndScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FPNHMCMHFHO uint32              `protobuf:"varint,8,opt,name=FPNHMCMHFHO,proto3" json:"FPNHMCMHFHO,omitempty"`
	CAAPHPHGJGG uint32              `protobuf:"varint,14,opt,name=CAAPHPHGJGG,proto3" json:"CAAPHPHGJGG,omitempty"`
	Id          uint32              `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	GFAOGLKEHHJ FeverTimeBattleRank `protobuf:"varint,4,opt,name=GFAOGLKEHHJ,proto3,enum=FeverTimeBattleRank" json:"GFAOGLKEHHJ,omitempty"`
}

func (x *FeverTimeActivityBattleEndScNotify) Reset() {
	*x = FeverTimeActivityBattleEndScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FeverTimeActivityBattleEndScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeverTimeActivityBattleEndScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeverTimeActivityBattleEndScNotify) ProtoMessage() {}

func (x *FeverTimeActivityBattleEndScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FeverTimeActivityBattleEndScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeverTimeActivityBattleEndScNotify.ProtoReflect.Descriptor instead.
func (*FeverTimeActivityBattleEndScNotify) Descriptor() ([]byte, []int) {
	return file_FeverTimeActivityBattleEndScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FeverTimeActivityBattleEndScNotify) GetFPNHMCMHFHO() uint32 {
	if x != nil {
		return x.FPNHMCMHFHO
	}
	return 0
}

func (x *FeverTimeActivityBattleEndScNotify) GetCAAPHPHGJGG() uint32 {
	if x != nil {
		return x.CAAPHPHGJGG
	}
	return 0
}

func (x *FeverTimeActivityBattleEndScNotify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeverTimeActivityBattleEndScNotify) GetGFAOGLKEHHJ() FeverTimeBattleRank {
	if x != nil {
		return x.GFAOGLKEHHJ
	}
	return FeverTimeBattleRank_FEVER_TIME_BATTLE_RANK_C
}

var File_FeverTimeActivityBattleEndScNotify_proto protoreflect.FileDescriptor

var file_FeverTimeActivityBattleEndScNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x46, 0x65, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x46, 0x65, 0x76, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x22, 0x46, 0x65, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x50, 0x4e, 0x48, 0x4d, 0x43, 0x4d, 0x48, 0x46, 0x48, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x46, 0x50, 0x4e, 0x48, 0x4d, 0x43, 0x4d, 0x48, 0x46, 0x48, 0x4f, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x41, 0x41, 0x50, 0x48, 0x50, 0x48, 0x47, 0x4a, 0x47, 0x47, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x41, 0x41, 0x50, 0x48, 0x50, 0x48, 0x47, 0x4a, 0x47, 0x47,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x36, 0x0a, 0x0b, 0x47, 0x46, 0x41, 0x4f, 0x47, 0x4c, 0x4b, 0x45, 0x48, 0x48, 0x4a, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x46, 0x65, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x0b, 0x47, 0x46, 0x41,
	0x4f, 0x47, 0x4c, 0x4b, 0x45, 0x48, 0x48, 0x4a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FeverTimeActivityBattleEndScNotify_proto_rawDescOnce sync.Once
	file_FeverTimeActivityBattleEndScNotify_proto_rawDescData = file_FeverTimeActivityBattleEndScNotify_proto_rawDesc
)

func file_FeverTimeActivityBattleEndScNotify_proto_rawDescGZIP() []byte {
	file_FeverTimeActivityBattleEndScNotify_proto_rawDescOnce.Do(func() {
		file_FeverTimeActivityBattleEndScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FeverTimeActivityBattleEndScNotify_proto_rawDescData)
	})
	return file_FeverTimeActivityBattleEndScNotify_proto_rawDescData
}

var file_FeverTimeActivityBattleEndScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FeverTimeActivityBattleEndScNotify_proto_goTypes = []interface{}{
	(*FeverTimeActivityBattleEndScNotify)(nil), // 0: FeverTimeActivityBattleEndScNotify
	(FeverTimeBattleRank)(0),                   // 1: FeverTimeBattleRank
}
var file_FeverTimeActivityBattleEndScNotify_proto_depIdxs = []int32{
	1, // 0: FeverTimeActivityBattleEndScNotify.GFAOGLKEHHJ:type_name -> FeverTimeBattleRank
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FeverTimeActivityBattleEndScNotify_proto_init() }
func file_FeverTimeActivityBattleEndScNotify_proto_init() {
	if File_FeverTimeActivityBattleEndScNotify_proto != nil {
		return
	}
	file_FeverTimeBattleRank_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FeverTimeActivityBattleEndScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeverTimeActivityBattleEndScNotify); i {
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
			RawDescriptor: file_FeverTimeActivityBattleEndScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FeverTimeActivityBattleEndScNotify_proto_goTypes,
		DependencyIndexes: file_FeverTimeActivityBattleEndScNotify_proto_depIdxs,
		MessageInfos:      file_FeverTimeActivityBattleEndScNotify_proto_msgTypes,
	}.Build()
	File_FeverTimeActivityBattleEndScNotify_proto = out.File
	file_FeverTimeActivityBattleEndScNotify_proto_rawDesc = nil
	file_FeverTimeActivityBattleEndScNotify_proto_goTypes = nil
	file_FeverTimeActivityBattleEndScNotify_proto_depIdxs = nil
}
