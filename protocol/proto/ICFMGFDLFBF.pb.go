// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ICFMGFDLFBF.proto

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

type ICFMGFDLFBF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ANBBCPICEHA []*LogisticsScore `protobuf:"bytes,15,rep,name=ANBBCPICEHA,proto3" json:"ANBBCPICEHA,omitempty"`
	OMHMNLEPINA []*OHJHHHGPCGG    `protobuf:"bytes,13,rep,name=OMHMNLEPINA,proto3" json:"OMHMNLEPINA,omitempty"`
	KICNDJJGLKJ []*KCJBOHGLBKG    `protobuf:"bytes,8,rep,name=KICNDJJGLKJ,proto3" json:"KICNDJJGLKJ,omitempty"`
	NBKEEHKBJAD []uint32          `protobuf:"varint,9,rep,packed,name=NBKEEHKBJAD,proto3" json:"NBKEEHKBJAD,omitempty"`
}

func (x *ICFMGFDLFBF) Reset() {
	*x = ICFMGFDLFBF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ICFMGFDLFBF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ICFMGFDLFBF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICFMGFDLFBF) ProtoMessage() {}

func (x *ICFMGFDLFBF) ProtoReflect() protoreflect.Message {
	mi := &file_ICFMGFDLFBF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICFMGFDLFBF.ProtoReflect.Descriptor instead.
func (*ICFMGFDLFBF) Descriptor() ([]byte, []int) {
	return file_ICFMGFDLFBF_proto_rawDescGZIP(), []int{0}
}

func (x *ICFMGFDLFBF) GetANBBCPICEHA() []*LogisticsScore {
	if x != nil {
		return x.ANBBCPICEHA
	}
	return nil
}

func (x *ICFMGFDLFBF) GetOMHMNLEPINA() []*OHJHHHGPCGG {
	if x != nil {
		return x.OMHMNLEPINA
	}
	return nil
}

func (x *ICFMGFDLFBF) GetKICNDJJGLKJ() []*KCJBOHGLBKG {
	if x != nil {
		return x.KICNDJJGLKJ
	}
	return nil
}

func (x *ICFMGFDLFBF) GetNBKEEHKBJAD() []uint32 {
	if x != nil {
		return x.NBKEEHKBJAD
	}
	return nil
}

var File_ICFMGFDLFBF_proto protoreflect.FileDescriptor

var file_ICFMGFDLFBF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x43, 0x46, 0x4d, 0x47, 0x46, 0x44, 0x4c, 0x46, 0x42, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x43, 0x4a, 0x42, 0x4f, 0x48, 0x47, 0x4c, 0x42, 0x4b, 0x47,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x48, 0x4a, 0x48, 0x48, 0x48, 0x47, 0x50,
	0x43, 0x47, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x01, 0x0a, 0x0b, 0x49, 0x43, 0x46, 0x4d, 0x47, 0x46, 0x44, 0x4c, 0x46, 0x42, 0x46, 0x12,
	0x31, 0x0a, 0x0b, 0x41, 0x4e, 0x42, 0x42, 0x43, 0x50, 0x49, 0x43, 0x45, 0x48, 0x41, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0b, 0x41, 0x4e, 0x42, 0x42, 0x43, 0x50, 0x49, 0x43, 0x45,
	0x48, 0x41, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4d, 0x48, 0x4d, 0x4e, 0x4c, 0x45, 0x50, 0x49, 0x4e,
	0x41, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x48, 0x4a, 0x48, 0x48, 0x48,
	0x47, 0x50, 0x43, 0x47, 0x47, 0x52, 0x0b, 0x4f, 0x4d, 0x48, 0x4d, 0x4e, 0x4c, 0x45, 0x50, 0x49,
	0x4e, 0x41, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x49, 0x43, 0x4e, 0x44, 0x4a, 0x4a, 0x47, 0x4c, 0x4b,
	0x4a, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x43, 0x4a, 0x42, 0x4f, 0x48,
	0x47, 0x4c, 0x42, 0x4b, 0x47, 0x52, 0x0b, 0x4b, 0x49, 0x43, 0x4e, 0x44, 0x4a, 0x4a, 0x47, 0x4c,
	0x4b, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x42, 0x4b, 0x45, 0x45, 0x48, 0x4b, 0x42, 0x4a, 0x41,
	0x44, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42, 0x4b, 0x45, 0x45, 0x48, 0x4b,
	0x42, 0x4a, 0x41, 0x44, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ICFMGFDLFBF_proto_rawDescOnce sync.Once
	file_ICFMGFDLFBF_proto_rawDescData = file_ICFMGFDLFBF_proto_rawDesc
)

func file_ICFMGFDLFBF_proto_rawDescGZIP() []byte {
	file_ICFMGFDLFBF_proto_rawDescOnce.Do(func() {
		file_ICFMGFDLFBF_proto_rawDescData = protoimpl.X.CompressGZIP(file_ICFMGFDLFBF_proto_rawDescData)
	})
	return file_ICFMGFDLFBF_proto_rawDescData
}

var file_ICFMGFDLFBF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ICFMGFDLFBF_proto_goTypes = []interface{}{
	(*ICFMGFDLFBF)(nil),    // 0: ICFMGFDLFBF
	(*LogisticsScore)(nil), // 1: LogisticsScore
	(*OHJHHHGPCGG)(nil),    // 2: OHJHHHGPCGG
	(*KCJBOHGLBKG)(nil),    // 3: KCJBOHGLBKG
}
var file_ICFMGFDLFBF_proto_depIdxs = []int32{
	1, // 0: ICFMGFDLFBF.ANBBCPICEHA:type_name -> LogisticsScore
	2, // 1: ICFMGFDLFBF.OMHMNLEPINA:type_name -> OHJHHHGPCGG
	3, // 2: ICFMGFDLFBF.KICNDJJGLKJ:type_name -> KCJBOHGLBKG
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ICFMGFDLFBF_proto_init() }
func file_ICFMGFDLFBF_proto_init() {
	if File_ICFMGFDLFBF_proto != nil {
		return
	}
	file_KCJBOHGLBKG_proto_init()
	file_OHJHHHGPCGG_proto_init()
	file_LogisticsScore_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ICFMGFDLFBF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ICFMGFDLFBF); i {
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
			RawDescriptor: file_ICFMGFDLFBF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ICFMGFDLFBF_proto_goTypes,
		DependencyIndexes: file_ICFMGFDLFBF_proto_depIdxs,
		MessageInfos:      file_ICFMGFDLFBF_proto_msgTypes,
	}.Build()
	File_ICFMGFDLFBF_proto = out.File
	file_ICFMGFDLFBF_proto_rawDesc = nil
	file_ICFMGFDLFBF_proto_goTypes = nil
	file_ICFMGFDLFBF_proto_depIdxs = nil
}