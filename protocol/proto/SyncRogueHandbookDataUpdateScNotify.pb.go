// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncRogueHandbookDataUpdateScNotify.proto

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

type SyncRogueHandbookDataUpdateScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OOIAEJKDKCF []*EDPHLINACNJ `protobuf:"bytes,14,rep,name=OOIAEJKDKCF,proto3" json:"OOIAEJKDKCF,omitempty"`
	FMFBOBJDIEC []*IFJPLLMDHPL `protobuf:"bytes,8,rep,name=FMFBOBJDIEC,proto3" json:"FMFBOBJDIEC,omitempty"`
	KGJFGMLGOCP []*EGNCJJJOLJA `protobuf:"bytes,4,rep,name=KGJFGMLGOCP,proto3" json:"KGJFGMLGOCP,omitempty"`
}

func (x *SyncRogueHandbookDataUpdateScNotify) Reset() {
	*x = SyncRogueHandbookDataUpdateScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueHandbookDataUpdateScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueHandbookDataUpdateScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueHandbookDataUpdateScNotify) ProtoMessage() {}

func (x *SyncRogueHandbookDataUpdateScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueHandbookDataUpdateScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueHandbookDataUpdateScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueHandbookDataUpdateScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueHandbookDataUpdateScNotify) GetOOIAEJKDKCF() []*EDPHLINACNJ {
	if x != nil {
		return x.OOIAEJKDKCF
	}
	return nil
}

func (x *SyncRogueHandbookDataUpdateScNotify) GetFMFBOBJDIEC() []*IFJPLLMDHPL {
	if x != nil {
		return x.FMFBOBJDIEC
	}
	return nil
}

func (x *SyncRogueHandbookDataUpdateScNotify) GetKGJFGMLGOCP() []*EGNCJJJOLJA {
	if x != nil {
		return x.KGJFGMLGOCP
	}
	return nil
}

var File_SyncRogueHandbookDataUpdateScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueHandbookDataUpdateScNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x46, 0x4a,
	0x50, 0x4c, 0x4c, 0x4d, 0x44, 0x48, 0x50, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x45, 0x47, 0x4e, 0x43, 0x4a, 0x4a, 0x4a, 0x4f, 0x4c, 0x4a, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x45, 0x44, 0x50, 0x48, 0x4c, 0x49, 0x4e, 0x41, 0x43, 0x4e, 0x4a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x23, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2e, 0x0a, 0x0b,
	0x4f, 0x4f, 0x49, 0x41, 0x45, 0x4a, 0x4b, 0x44, 0x4b, 0x43, 0x46, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x44, 0x50, 0x48, 0x4c, 0x49, 0x4e, 0x41, 0x43, 0x4e, 0x4a, 0x52,
	0x0b, 0x4f, 0x4f, 0x49, 0x41, 0x45, 0x4a, 0x4b, 0x44, 0x4b, 0x43, 0x46, 0x12, 0x2e, 0x0a, 0x0b,
	0x46, 0x4d, 0x46, 0x42, 0x4f, 0x42, 0x4a, 0x44, 0x49, 0x45, 0x43, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x46, 0x4a, 0x50, 0x4c, 0x4c, 0x4d, 0x44, 0x48, 0x50, 0x4c, 0x52,
	0x0b, 0x46, 0x4d, 0x46, 0x42, 0x4f, 0x42, 0x4a, 0x44, 0x49, 0x45, 0x43, 0x12, 0x2e, 0x0a, 0x0b,
	0x4b, 0x47, 0x4a, 0x46, 0x47, 0x4d, 0x4c, 0x47, 0x4f, 0x43, 0x50, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x47, 0x4e, 0x43, 0x4a, 0x4a, 0x4a, 0x4f, 0x4c, 0x4a, 0x41, 0x52,
	0x0b, 0x4b, 0x47, 0x4a, 0x46, 0x47, 0x4d, 0x4c, 0x47, 0x4f, 0x43, 0x50, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescData = file_SyncRogueHandbookDataUpdateScNotify_proto_rawDesc
)

func file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescData)
	})
	return file_SyncRogueHandbookDataUpdateScNotify_proto_rawDescData
}

var file_SyncRogueHandbookDataUpdateScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueHandbookDataUpdateScNotify_proto_goTypes = []interface{}{
	(*SyncRogueHandbookDataUpdateScNotify)(nil), // 0: SyncRogueHandbookDataUpdateScNotify
	(*EDPHLINACNJ)(nil),                         // 1: EDPHLINACNJ
	(*IFJPLLMDHPL)(nil),                         // 2: IFJPLLMDHPL
	(*EGNCJJJOLJA)(nil),                         // 3: EGNCJJJOLJA
}
var file_SyncRogueHandbookDataUpdateScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueHandbookDataUpdateScNotify.OOIAEJKDKCF:type_name -> EDPHLINACNJ
	2, // 1: SyncRogueHandbookDataUpdateScNotify.FMFBOBJDIEC:type_name -> IFJPLLMDHPL
	3, // 2: SyncRogueHandbookDataUpdateScNotify.KGJFGMLGOCP:type_name -> EGNCJJJOLJA
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SyncRogueHandbookDataUpdateScNotify_proto_init() }
func file_SyncRogueHandbookDataUpdateScNotify_proto_init() {
	if File_SyncRogueHandbookDataUpdateScNotify_proto != nil {
		return
	}
	file_IFJPLLMDHPL_proto_init()
	file_EGNCJJJOLJA_proto_init()
	file_EDPHLINACNJ_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueHandbookDataUpdateScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueHandbookDataUpdateScNotify); i {
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
			RawDescriptor: file_SyncRogueHandbookDataUpdateScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueHandbookDataUpdateScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueHandbookDataUpdateScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueHandbookDataUpdateScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueHandbookDataUpdateScNotify_proto = out.File
	file_SyncRogueHandbookDataUpdateScNotify_proto_rawDesc = nil
	file_SyncRogueHandbookDataUpdateScNotify_proto_goTypes = nil
	file_SyncRogueHandbookDataUpdateScNotify_proto_depIdxs = nil
}
