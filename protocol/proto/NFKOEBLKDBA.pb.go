// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NFKOEBLKDBA.proto

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

type NFKOEBLKDBA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDFOBIKKAJG uint32 `protobuf:"varint,13,opt,name=IDFOBIKKAJG,proto3" json:"IDFOBIKKAJG,omitempty"`
	CEBCFEOMABN uint32 `protobuf:"varint,6,opt,name=CEBCFEOMABN,proto3" json:"CEBCFEOMABN,omitempty"`
	FJAHEJCCDHI uint32 `protobuf:"varint,1,opt,name=FJAHEJCCDHI,proto3" json:"FJAHEJCCDHI,omitempty"`
	KOIMAACPIJO uint32 `protobuf:"varint,4,opt,name=KOIMAACPIJO,proto3" json:"KOIMAACPIJO,omitempty"`
	JNJJMGJPMAM uint32 `protobuf:"varint,12,opt,name=JNJJMGJPMAM,proto3" json:"JNJJMGJPMAM,omitempty"`
	AKJAKAAFFFA uint32 `protobuf:"varint,2,opt,name=AKJAKAAFFFA,proto3" json:"AKJAKAAFFFA,omitempty"`
	NIFNEJPFBKI bool   `protobuf:"varint,9,opt,name=NIFNEJPFBKI,proto3" json:"NIFNEJPFBKI,omitempty"`
	INPDFOHLLGO uint32 `protobuf:"varint,15,opt,name=INPDFOHLLGO,proto3" json:"INPDFOHLLGO,omitempty"`
	BFAELPFLGKD uint32 `protobuf:"varint,11,opt,name=BFAELPFLGKD,proto3" json:"BFAELPFLGKD,omitempty"`
}

func (x *NFKOEBLKDBA) Reset() {
	*x = NFKOEBLKDBA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NFKOEBLKDBA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFKOEBLKDBA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFKOEBLKDBA) ProtoMessage() {}

func (x *NFKOEBLKDBA) ProtoReflect() protoreflect.Message {
	mi := &file_NFKOEBLKDBA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFKOEBLKDBA.ProtoReflect.Descriptor instead.
func (*NFKOEBLKDBA) Descriptor() ([]byte, []int) {
	return file_NFKOEBLKDBA_proto_rawDescGZIP(), []int{0}
}

func (x *NFKOEBLKDBA) GetIDFOBIKKAJG() uint32 {
	if x != nil {
		return x.IDFOBIKKAJG
	}
	return 0
}

func (x *NFKOEBLKDBA) GetCEBCFEOMABN() uint32 {
	if x != nil {
		return x.CEBCFEOMABN
	}
	return 0
}

func (x *NFKOEBLKDBA) GetFJAHEJCCDHI() uint32 {
	if x != nil {
		return x.FJAHEJCCDHI
	}
	return 0
}

func (x *NFKOEBLKDBA) GetKOIMAACPIJO() uint32 {
	if x != nil {
		return x.KOIMAACPIJO
	}
	return 0
}

func (x *NFKOEBLKDBA) GetJNJJMGJPMAM() uint32 {
	if x != nil {
		return x.JNJJMGJPMAM
	}
	return 0
}

func (x *NFKOEBLKDBA) GetAKJAKAAFFFA() uint32 {
	if x != nil {
		return x.AKJAKAAFFFA
	}
	return 0
}

func (x *NFKOEBLKDBA) GetNIFNEJPFBKI() bool {
	if x != nil {
		return x.NIFNEJPFBKI
	}
	return false
}

func (x *NFKOEBLKDBA) GetINPDFOHLLGO() uint32 {
	if x != nil {
		return x.INPDFOHLLGO
	}
	return 0
}

func (x *NFKOEBLKDBA) GetBFAELPFLGKD() uint32 {
	if x != nil {
		return x.BFAELPFLGKD
	}
	return 0
}

var File_NFKOEBLKDBA_proto protoreflect.FileDescriptor

var file_NFKOEBLKDBA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x46, 0x4b, 0x4f, 0x45, 0x42, 0x4c, 0x4b, 0x44, 0x42, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x0b, 0x4e, 0x46, 0x4b, 0x4f, 0x45, 0x42, 0x4c, 0x4b,
	0x44, 0x42, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x44, 0x46, 0x4f, 0x42, 0x49, 0x4b, 0x4b, 0x41,
	0x4a, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x44, 0x46, 0x4f, 0x42, 0x49,
	0x4b, 0x4b, 0x41, 0x4a, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x45, 0x42, 0x43, 0x46, 0x45, 0x4f,
	0x4d, 0x41, 0x42, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x45, 0x42, 0x43,
	0x46, 0x45, 0x4f, 0x4d, 0x41, 0x42, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4a, 0x41, 0x48, 0x45,
	0x4a, 0x43, 0x43, 0x44, 0x48, 0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4a,
	0x41, 0x48, 0x45, 0x4a, 0x43, 0x43, 0x44, 0x48, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4f, 0x49,
	0x4d, 0x41, 0x41, 0x43, 0x50, 0x49, 0x4a, 0x4f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4b, 0x4f, 0x49, 0x4d, 0x41, 0x41, 0x43, 0x50, 0x49, 0x4a, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x4e, 0x4a, 0x4a, 0x4d, 0x47, 0x4a, 0x50, 0x4d, 0x41, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x4e, 0x4a, 0x4a, 0x4d, 0x47, 0x4a, 0x50, 0x4d, 0x41, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x4b, 0x4a, 0x41, 0x4b, 0x41, 0x41, 0x46, 0x46, 0x46, 0x41, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4b, 0x4a, 0x41, 0x4b, 0x41, 0x41, 0x46, 0x46, 0x46, 0x41, 0x12,
	0x20, 0x0a, 0x0b, 0x4e, 0x49, 0x46, 0x4e, 0x45, 0x4a, 0x50, 0x46, 0x42, 0x4b, 0x49, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x49, 0x46, 0x4e, 0x45, 0x4a, 0x50, 0x46, 0x42, 0x4b,
	0x49, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x50, 0x44, 0x46, 0x4f, 0x48, 0x4c, 0x4c, 0x47, 0x4f,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x50, 0x44, 0x46, 0x4f, 0x48, 0x4c,
	0x4c, 0x47, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46, 0x41, 0x45, 0x4c, 0x50, 0x46, 0x4c, 0x47,
	0x4b, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x46, 0x41, 0x45, 0x4c, 0x50,
	0x46, 0x4c, 0x47, 0x4b, 0x44, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NFKOEBLKDBA_proto_rawDescOnce sync.Once
	file_NFKOEBLKDBA_proto_rawDescData = file_NFKOEBLKDBA_proto_rawDesc
)

func file_NFKOEBLKDBA_proto_rawDescGZIP() []byte {
	file_NFKOEBLKDBA_proto_rawDescOnce.Do(func() {
		file_NFKOEBLKDBA_proto_rawDescData = protoimpl.X.CompressGZIP(file_NFKOEBLKDBA_proto_rawDescData)
	})
	return file_NFKOEBLKDBA_proto_rawDescData
}

var file_NFKOEBLKDBA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NFKOEBLKDBA_proto_goTypes = []interface{}{
	(*NFKOEBLKDBA)(nil), // 0: NFKOEBLKDBA
}
var file_NFKOEBLKDBA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NFKOEBLKDBA_proto_init() }
func file_NFKOEBLKDBA_proto_init() {
	if File_NFKOEBLKDBA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NFKOEBLKDBA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFKOEBLKDBA); i {
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
			RawDescriptor: file_NFKOEBLKDBA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NFKOEBLKDBA_proto_goTypes,
		DependencyIndexes: file_NFKOEBLKDBA_proto_depIdxs,
		MessageInfos:      file_NFKOEBLKDBA_proto_msgTypes,
	}.Build()
	File_NFKOEBLKDBA_proto = out.File
	file_NFKOEBLKDBA_proto_rawDesc = nil
	file_NFKOEBLKDBA_proto_goTypes = nil
	file_NFKOEBLKDBA_proto_depIdxs = nil
}