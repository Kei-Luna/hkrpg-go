// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonActivityData.proto

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

type TreasureDungeonActivityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JMBBHKAABHL uint32 `protobuf:"varint,8,opt,name=JMBBHKAABHL,proto3" json:"JMBBHKAABHL,omitempty"`
	LHAOGKKNNCC bool   `protobuf:"varint,1,opt,name=LHAOGKKNNCC,proto3" json:"LHAOGKKNNCC,omitempty"`
	BPAGNPADCDE uint32 `protobuf:"varint,11,opt,name=BPAGNPADCDE,proto3" json:"BPAGNPADCDE,omitempty"`
	IPLEKIJDMLH uint32 `protobuf:"varint,3,opt,name=IPLEKIJDMLH,proto3" json:"IPLEKIJDMLH,omitempty"`
	IsFinished  bool   `protobuf:"varint,6,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	PCOBHIGDHEK uint32 `protobuf:"varint,2,opt,name=PCOBHIGDHEK,proto3" json:"PCOBHIGDHEK,omitempty"`
	AOKPAOJINNC uint32 `protobuf:"varint,4,opt,name=AOKPAOJINNC,proto3" json:"AOKPAOJINNC,omitempty"`
}

func (x *TreasureDungeonActivityData) Reset() {
	*x = TreasureDungeonActivityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureDungeonActivityData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureDungeonActivityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureDungeonActivityData) ProtoMessage() {}

func (x *TreasureDungeonActivityData) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureDungeonActivityData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureDungeonActivityData.ProtoReflect.Descriptor instead.
func (*TreasureDungeonActivityData) Descriptor() ([]byte, []int) {
	return file_TreasureDungeonActivityData_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureDungeonActivityData) GetJMBBHKAABHL() uint32 {
	if x != nil {
		return x.JMBBHKAABHL
	}
	return 0
}

func (x *TreasureDungeonActivityData) GetLHAOGKKNNCC() bool {
	if x != nil {
		return x.LHAOGKKNNCC
	}
	return false
}

func (x *TreasureDungeonActivityData) GetBPAGNPADCDE() uint32 {
	if x != nil {
		return x.BPAGNPADCDE
	}
	return 0
}

func (x *TreasureDungeonActivityData) GetIPLEKIJDMLH() uint32 {
	if x != nil {
		return x.IPLEKIJDMLH
	}
	return 0
}

func (x *TreasureDungeonActivityData) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *TreasureDungeonActivityData) GetPCOBHIGDHEK() uint32 {
	if x != nil {
		return x.PCOBHIGDHEK
	}
	return 0
}

func (x *TreasureDungeonActivityData) GetAOKPAOJINNC() uint32 {
	if x != nil {
		return x.AOKPAOJINNC
	}
	return 0
}

var File_TreasureDungeonActivityData_proto protoreflect.FileDescriptor

var file_TreasureDungeonActivityData_proto_rawDesc = []byte{
	0x0a, 0x21, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x1b, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4d, 0x42, 0x42, 0x48, 0x4b, 0x41, 0x41, 0x42,
	0x48, 0x4c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4d, 0x42, 0x42, 0x48, 0x4b,
	0x41, 0x41, 0x42, 0x48, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x48, 0x41, 0x4f, 0x47, 0x4b, 0x4b,
	0x4e, 0x4e, 0x43, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x48, 0x41, 0x4f,
	0x47, 0x4b, 0x4b, 0x4e, 0x4e, 0x43, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x50, 0x41, 0x47, 0x4e,
	0x50, 0x41, 0x44, 0x43, 0x44, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x50,
	0x41, 0x47, 0x4e, 0x50, 0x41, 0x44, 0x43, 0x44, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x50, 0x4c,
	0x45, 0x4b, 0x49, 0x4a, 0x44, 0x4d, 0x4c, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x50, 0x4c, 0x45, 0x4b, 0x49, 0x4a, 0x44, 0x4d, 0x4c, 0x48, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x43, 0x4f, 0x42, 0x48, 0x49, 0x47, 0x44, 0x48, 0x45, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x50, 0x43, 0x4f, 0x42, 0x48, 0x49, 0x47, 0x44, 0x48, 0x45, 0x4b, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x4f, 0x4b, 0x50, 0x41, 0x4f, 0x4a, 0x49, 0x4e, 0x4e, 0x43, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4f, 0x4b, 0x50, 0x41, 0x4f, 0x4a, 0x49, 0x4e, 0x4e, 0x43,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonActivityData_proto_rawDescOnce sync.Once
	file_TreasureDungeonActivityData_proto_rawDescData = file_TreasureDungeonActivityData_proto_rawDesc
)

func file_TreasureDungeonActivityData_proto_rawDescGZIP() []byte {
	file_TreasureDungeonActivityData_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonActivityData_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonActivityData_proto_rawDescData)
	})
	return file_TreasureDungeonActivityData_proto_rawDescData
}

var file_TreasureDungeonActivityData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureDungeonActivityData_proto_goTypes = []interface{}{
	(*TreasureDungeonActivityData)(nil), // 0: TreasureDungeonActivityData
}
var file_TreasureDungeonActivityData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TreasureDungeonActivityData_proto_init() }
func file_TreasureDungeonActivityData_proto_init() {
	if File_TreasureDungeonActivityData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TreasureDungeonActivityData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureDungeonActivityData); i {
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
			RawDescriptor: file_TreasureDungeonActivityData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonActivityData_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonActivityData_proto_depIdxs,
		MessageInfos:      file_TreasureDungeonActivityData_proto_msgTypes,
	}.Build()
	File_TreasureDungeonActivityData_proto = out.File
	file_TreasureDungeonActivityData_proto_rawDesc = nil
	file_TreasureDungeonActivityData_proto_goTypes = nil
	file_TreasureDungeonActivityData_proto_depIdxs = nil
}
