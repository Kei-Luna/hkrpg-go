// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HEEENFFOBLE.proto

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

type HEEENFFOBLE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NMBPIEBNJPG bool           `protobuf:"varint,9,opt,name=NMBPIEBNJPG,proto3" json:"NMBPIEBNJPG,omitempty"`
	OJNOMLGBILP bool           `protobuf:"varint,7,opt,name=OJNOMLGBILP,proto3" json:"OJNOMLGBILP,omitempty"`
	HPHCAEDDOCM uint32         `protobuf:"varint,15,opt,name=HPHCAEDDOCM,proto3" json:"HPHCAEDDOCM,omitempty"`
	ALNBFFGIKNM bool           `protobuf:"varint,14,opt,name=ALNBFFGIKNM,proto3" json:"ALNBFFGIKNM,omitempty"`
	HHPDJFILHLP uint32         `protobuf:"varint,1,opt,name=HHPDJFILHLP,proto3" json:"HHPDJFILHLP,omitempty"`
	DBENCJELCIF bool           `protobuf:"varint,10,opt,name=DBENCJELCIF,proto3" json:"DBENCJELCIF,omitempty"`
	CIDDCGHFFOC uint32         `protobuf:"varint,6,opt,name=CIDDCGHFFOC,proto3" json:"CIDDCGHFFOC,omitempty"`
	BuffList    []*OLMNOHDHFHN `protobuf:"bytes,1456,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
}

func (x *HEEENFFOBLE) Reset() {
	*x = HEEENFFOBLE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HEEENFFOBLE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HEEENFFOBLE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HEEENFFOBLE) ProtoMessage() {}

func (x *HEEENFFOBLE) ProtoReflect() protoreflect.Message {
	mi := &file_HEEENFFOBLE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HEEENFFOBLE.ProtoReflect.Descriptor instead.
func (*HEEENFFOBLE) Descriptor() ([]byte, []int) {
	return file_HEEENFFOBLE_proto_rawDescGZIP(), []int{0}
}

func (x *HEEENFFOBLE) GetNMBPIEBNJPG() bool {
	if x != nil {
		return x.NMBPIEBNJPG
	}
	return false
}

func (x *HEEENFFOBLE) GetOJNOMLGBILP() bool {
	if x != nil {
		return x.OJNOMLGBILP
	}
	return false
}

func (x *HEEENFFOBLE) GetHPHCAEDDOCM() uint32 {
	if x != nil {
		return x.HPHCAEDDOCM
	}
	return 0
}

func (x *HEEENFFOBLE) GetALNBFFGIKNM() bool {
	if x != nil {
		return x.ALNBFFGIKNM
	}
	return false
}

func (x *HEEENFFOBLE) GetHHPDJFILHLP() uint32 {
	if x != nil {
		return x.HHPDJFILHLP
	}
	return 0
}

func (x *HEEENFFOBLE) GetDBENCJELCIF() bool {
	if x != nil {
		return x.DBENCJELCIF
	}
	return false
}

func (x *HEEENFFOBLE) GetCIDDCGHFFOC() uint32 {
	if x != nil {
		return x.CIDDCGHFFOC
	}
	return 0
}

func (x *HEEENFFOBLE) GetBuffList() []*OLMNOHDHFHN {
	if x != nil {
		return x.BuffList
	}
	return nil
}

var File_HEEENFFOBLE_proto protoreflect.FileDescriptor

var file_HEEENFFOBLE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x45, 0x45, 0x45, 0x4e, 0x46, 0x46, 0x4f, 0x42, 0x4c, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4c, 0x4d, 0x4e, 0x4f, 0x48, 0x44, 0x48, 0x46, 0x48, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x48, 0x45, 0x45, 0x45, 0x4e,
	0x46, 0x46, 0x4f, 0x42, 0x4c, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4d, 0x42, 0x50, 0x49, 0x45,
	0x42, 0x4e, 0x4a, 0x50, 0x47, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x4d, 0x42,
	0x50, 0x49, 0x45, 0x42, 0x4e, 0x4a, 0x50, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4a, 0x4e, 0x4f,
	0x4d, 0x4c, 0x47, 0x42, 0x49, 0x4c, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f,
	0x4a, 0x4e, 0x4f, 0x4d, 0x4c, 0x47, 0x42, 0x49, 0x4c, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x50,
	0x48, 0x43, 0x41, 0x45, 0x44, 0x44, 0x4f, 0x43, 0x4d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x48, 0x50, 0x48, 0x43, 0x41, 0x45, 0x44, 0x44, 0x4f, 0x43, 0x4d, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x4c, 0x4e, 0x42, 0x46, 0x46, 0x47, 0x49, 0x4b, 0x4e, 0x4d, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x41, 0x4c, 0x4e, 0x42, 0x46, 0x46, 0x47, 0x49, 0x4b, 0x4e, 0x4d, 0x12, 0x20,
	0x0a, 0x0b, 0x48, 0x48, 0x50, 0x44, 0x4a, 0x46, 0x49, 0x4c, 0x48, 0x4c, 0x50, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x48, 0x50, 0x44, 0x4a, 0x46, 0x49, 0x4c, 0x48, 0x4c, 0x50,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x42, 0x45, 0x4e, 0x43, 0x4a, 0x45, 0x4c, 0x43, 0x49, 0x46, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x42, 0x45, 0x4e, 0x43, 0x4a, 0x45, 0x4c, 0x43,
	0x49, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x44, 0x44, 0x43, 0x47, 0x48, 0x46, 0x46, 0x4f,
	0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x49, 0x44, 0x44, 0x43, 0x47, 0x48,
	0x46, 0x46, 0x4f, 0x43, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0xb0, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4c, 0x4d, 0x4e, 0x4f,
	0x48, 0x44, 0x48, 0x46, 0x48, 0x4e, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_HEEENFFOBLE_proto_rawDescOnce sync.Once
	file_HEEENFFOBLE_proto_rawDescData = file_HEEENFFOBLE_proto_rawDesc
)

func file_HEEENFFOBLE_proto_rawDescGZIP() []byte {
	file_HEEENFFOBLE_proto_rawDescOnce.Do(func() {
		file_HEEENFFOBLE_proto_rawDescData = protoimpl.X.CompressGZIP(file_HEEENFFOBLE_proto_rawDescData)
	})
	return file_HEEENFFOBLE_proto_rawDescData
}

var file_HEEENFFOBLE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HEEENFFOBLE_proto_goTypes = []interface{}{
	(*HEEENFFOBLE)(nil), // 0: HEEENFFOBLE
	(*OLMNOHDHFHN)(nil), // 1: OLMNOHDHFHN
}
var file_HEEENFFOBLE_proto_depIdxs = []int32{
	1, // 0: HEEENFFOBLE.buff_list:type_name -> OLMNOHDHFHN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HEEENFFOBLE_proto_init() }
func file_HEEENFFOBLE_proto_init() {
	if File_HEEENFFOBLE_proto != nil {
		return
	}
	file_OLMNOHDHFHN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HEEENFFOBLE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HEEENFFOBLE); i {
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
			RawDescriptor: file_HEEENFFOBLE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HEEENFFOBLE_proto_goTypes,
		DependencyIndexes: file_HEEENFFOBLE_proto_depIdxs,
		MessageInfos:      file_HEEENFFOBLE_proto_msgTypes,
	}.Build()
	File_HEEENFFOBLE_proto = out.File
	file_HEEENFFOBLE_proto_rawDesc = nil
	file_HEEENFFOBLE_proto_goTypes = nil
	file_HEEENFFOBLE_proto_depIdxs = nil
}