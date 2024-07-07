// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MechanismBarInfo.proto

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

type MechanismBarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OAFEFEBJABO uint32 `protobuf:"varint,1,opt,name=OAFEFEBJABO,proto3" json:"OAFEFEBJABO,omitempty"`
	Value       uint32 `protobuf:"varint,11,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MechanismBarInfo) Reset() {
	*x = MechanismBarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MechanismBarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MechanismBarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanismBarInfo) ProtoMessage() {}

func (x *MechanismBarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MechanismBarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanismBarInfo.ProtoReflect.Descriptor instead.
func (*MechanismBarInfo) Descriptor() ([]byte, []int) {
	return file_MechanismBarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MechanismBarInfo) GetOAFEFEBJABO() uint32 {
	if x != nil {
		return x.OAFEFEBJABO
	}
	return 0
}

func (x *MechanismBarInfo) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_MechanismBarInfo_proto protoreflect.FileDescriptor

var file_MechanismBarInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x42, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x4d, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x73, 0x6d, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x4f, 0x41, 0x46, 0x45, 0x46, 0x45, 0x42, 0x4a, 0x41, 0x42, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4f, 0x41, 0x46, 0x45, 0x46, 0x45, 0x42, 0x4a, 0x41, 0x42, 0x4f, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MechanismBarInfo_proto_rawDescOnce sync.Once
	file_MechanismBarInfo_proto_rawDescData = file_MechanismBarInfo_proto_rawDesc
)

func file_MechanismBarInfo_proto_rawDescGZIP() []byte {
	file_MechanismBarInfo_proto_rawDescOnce.Do(func() {
		file_MechanismBarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MechanismBarInfo_proto_rawDescData)
	})
	return file_MechanismBarInfo_proto_rawDescData
}

var file_MechanismBarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MechanismBarInfo_proto_goTypes = []interface{}{
	(*MechanismBarInfo)(nil), // 0: MechanismBarInfo
}
var file_MechanismBarInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MechanismBarInfo_proto_init() }
func file_MechanismBarInfo_proto_init() {
	if File_MechanismBarInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MechanismBarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MechanismBarInfo); i {
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
			RawDescriptor: file_MechanismBarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MechanismBarInfo_proto_goTypes,
		DependencyIndexes: file_MechanismBarInfo_proto_depIdxs,
		MessageInfos:      file_MechanismBarInfo_proto_msgTypes,
	}.Build()
	File_MechanismBarInfo_proto = out.File
	file_MechanismBarInfo_proto_rawDesc = nil
	file_MechanismBarInfo_proto_goTypes = nil
	file_MechanismBarInfo_proto_depIdxs = nil
}