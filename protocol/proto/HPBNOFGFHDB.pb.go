// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HPBNOFGFHDB.proto

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

type HPBNOFGFHDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EEDIFOLPKKO uint32 `protobuf:"varint,1,opt,name=EEDIFOLPKKO,proto3" json:"EEDIFOLPKKO,omitempty"`
	JIICGEAGECB uint32 `protobuf:"varint,2,opt,name=JIICGEAGECB,proto3" json:"JIICGEAGECB,omitempty"`
	JOPKDGFBHHA uint32 `protobuf:"varint,3,opt,name=JOPKDGFBHHA,proto3" json:"JOPKDGFBHHA,omitempty"`
	DIOCLGIOEGJ uint32 `protobuf:"varint,4,opt,name=DIOCLGIOEGJ,proto3" json:"DIOCLGIOEGJ,omitempty"`
	AFCLFJOFBGG uint32 `protobuf:"varint,5,opt,name=AFCLFJOFBGG,proto3" json:"AFCLFJOFBGG,omitempty"`
	HBKPAEOGJEL uint32 `protobuf:"varint,6,opt,name=HBKPAEOGJEL,proto3" json:"HBKPAEOGJEL,omitempty"`
	FEKGNAMFGKP uint32 `protobuf:"varint,7,opt,name=FEKGNAMFGKP,proto3" json:"FEKGNAMFGKP,omitempty"`
	HIBONPGJBDH uint32 `protobuf:"varint,8,opt,name=HIBONPGJBDH,proto3" json:"HIBONPGJBDH,omitempty"`
	EHLELHGJOBH uint32 `protobuf:"varint,9,opt,name=EHLELHGJOBH,proto3" json:"EHLELHGJOBH,omitempty"`
	GNKKGBMLIAO uint32 `protobuf:"varint,10,opt,name=GNKKGBMLIAO,proto3" json:"GNKKGBMLIAO,omitempty"`
}

func (x *HPBNOFGFHDB) Reset() {
	*x = HPBNOFGFHDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HPBNOFGFHDB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HPBNOFGFHDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HPBNOFGFHDB) ProtoMessage() {}

func (x *HPBNOFGFHDB) ProtoReflect() protoreflect.Message {
	mi := &file_HPBNOFGFHDB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HPBNOFGFHDB.ProtoReflect.Descriptor instead.
func (*HPBNOFGFHDB) Descriptor() ([]byte, []int) {
	return file_HPBNOFGFHDB_proto_rawDescGZIP(), []int{0}
}

func (x *HPBNOFGFHDB) GetEEDIFOLPKKO() uint32 {
	if x != nil {
		return x.EEDIFOLPKKO
	}
	return 0
}

func (x *HPBNOFGFHDB) GetJIICGEAGECB() uint32 {
	if x != nil {
		return x.JIICGEAGECB
	}
	return 0
}

func (x *HPBNOFGFHDB) GetJOPKDGFBHHA() uint32 {
	if x != nil {
		return x.JOPKDGFBHHA
	}
	return 0
}

func (x *HPBNOFGFHDB) GetDIOCLGIOEGJ() uint32 {
	if x != nil {
		return x.DIOCLGIOEGJ
	}
	return 0
}

func (x *HPBNOFGFHDB) GetAFCLFJOFBGG() uint32 {
	if x != nil {
		return x.AFCLFJOFBGG
	}
	return 0
}

func (x *HPBNOFGFHDB) GetHBKPAEOGJEL() uint32 {
	if x != nil {
		return x.HBKPAEOGJEL
	}
	return 0
}

func (x *HPBNOFGFHDB) GetFEKGNAMFGKP() uint32 {
	if x != nil {
		return x.FEKGNAMFGKP
	}
	return 0
}

func (x *HPBNOFGFHDB) GetHIBONPGJBDH() uint32 {
	if x != nil {
		return x.HIBONPGJBDH
	}
	return 0
}

func (x *HPBNOFGFHDB) GetEHLELHGJOBH() uint32 {
	if x != nil {
		return x.EHLELHGJOBH
	}
	return 0
}

func (x *HPBNOFGFHDB) GetGNKKGBMLIAO() uint32 {
	if x != nil {
		return x.GNKKGBMLIAO
	}
	return 0
}

var File_HPBNOFGFHDB_proto protoreflect.FileDescriptor

var file_HPBNOFGFHDB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x50, 0x42, 0x4e, 0x4f, 0x46, 0x47, 0x46, 0x48, 0x44, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x0b, 0x48, 0x50, 0x42, 0x4e, 0x4f, 0x46, 0x47, 0x46,
	0x48, 0x44, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x45, 0x44, 0x49, 0x46, 0x4f, 0x4c, 0x50, 0x4b,
	0x4b, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x45, 0x44, 0x49, 0x46, 0x4f,
	0x4c, 0x50, 0x4b, 0x4b, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x49, 0x43, 0x47, 0x45, 0x41,
	0x47, 0x45, 0x43, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x49, 0x49, 0x43,
	0x47, 0x45, 0x41, 0x47, 0x45, 0x43, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4f, 0x50, 0x4b, 0x44,
	0x47, 0x46, 0x42, 0x48, 0x48, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4f,
	0x50, 0x4b, 0x44, 0x47, 0x46, 0x42, 0x48, 0x48, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x49, 0x4f,
	0x43, 0x4c, 0x47, 0x49, 0x4f, 0x45, 0x47, 0x4a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x49, 0x4f, 0x43, 0x4c, 0x47, 0x49, 0x4f, 0x45, 0x47, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x46, 0x43, 0x4c, 0x46, 0x4a, 0x4f, 0x46, 0x42, 0x47, 0x47, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x46, 0x43, 0x4c, 0x46, 0x4a, 0x4f, 0x46, 0x42, 0x47, 0x47, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x42, 0x4b, 0x50, 0x41, 0x45, 0x4f, 0x47, 0x4a, 0x45, 0x4c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x48, 0x42, 0x4b, 0x50, 0x41, 0x45, 0x4f, 0x47, 0x4a, 0x45, 0x4c, 0x12,
	0x20, 0x0a, 0x0b, 0x46, 0x45, 0x4b, 0x47, 0x4e, 0x41, 0x4d, 0x46, 0x47, 0x4b, 0x50, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x45, 0x4b, 0x47, 0x4e, 0x41, 0x4d, 0x46, 0x47, 0x4b,
	0x50, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x49, 0x42, 0x4f, 0x4e, 0x50, 0x47, 0x4a, 0x42, 0x44, 0x48,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x49, 0x42, 0x4f, 0x4e, 0x50, 0x47, 0x4a,
	0x42, 0x44, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x48, 0x4c, 0x45, 0x4c, 0x48, 0x47, 0x4a, 0x4f,
	0x42, 0x48, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x48, 0x4c, 0x45, 0x4c, 0x48,
	0x47, 0x4a, 0x4f, 0x42, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4e, 0x4b, 0x4b, 0x47, 0x42, 0x4d,
	0x4c, 0x49, 0x41, 0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4e, 0x4b, 0x4b,
	0x47, 0x42, 0x4d, 0x4c, 0x49, 0x41, 0x4f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HPBNOFGFHDB_proto_rawDescOnce sync.Once
	file_HPBNOFGFHDB_proto_rawDescData = file_HPBNOFGFHDB_proto_rawDesc
)

func file_HPBNOFGFHDB_proto_rawDescGZIP() []byte {
	file_HPBNOFGFHDB_proto_rawDescOnce.Do(func() {
		file_HPBNOFGFHDB_proto_rawDescData = protoimpl.X.CompressGZIP(file_HPBNOFGFHDB_proto_rawDescData)
	})
	return file_HPBNOFGFHDB_proto_rawDescData
}

var file_HPBNOFGFHDB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HPBNOFGFHDB_proto_goTypes = []interface{}{
	(*HPBNOFGFHDB)(nil), // 0: HPBNOFGFHDB
}
var file_HPBNOFGFHDB_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HPBNOFGFHDB_proto_init() }
func file_HPBNOFGFHDB_proto_init() {
	if File_HPBNOFGFHDB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HPBNOFGFHDB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HPBNOFGFHDB); i {
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
			RawDescriptor: file_HPBNOFGFHDB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HPBNOFGFHDB_proto_goTypes,
		DependencyIndexes: file_HPBNOFGFHDB_proto_depIdxs,
		MessageInfos:      file_HPBNOFGFHDB_proto_msgTypes,
	}.Build()
	File_HPBNOFGFHDB_proto = out.File
	file_HPBNOFGFHDB_proto_rawDesc = nil
	file_HPBNOFGFHDB_proto_goTypes = nil
	file_HPBNOFGFHDB_proto_depIdxs = nil
}