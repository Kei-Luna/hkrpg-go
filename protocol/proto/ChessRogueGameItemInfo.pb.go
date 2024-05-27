// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueGameItemInfo.proto

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

type ChessRogueGameItemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemMap map[uint32]uint32 `protobuf:"bytes,4,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ChessRogueGameItemInfo) Reset() {
	*x = ChessRogueGameItemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueGameItemInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueGameItemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueGameItemInfo) ProtoMessage() {}

func (x *ChessRogueGameItemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueGameItemInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueGameItemInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueGameItemInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueGameItemInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueGameItemInfo) GetItemMap() map[uint32]uint32 {
	if x != nil {
		return x.ItemMap
	}
	return nil
}

var File_ChessRogueGameItemInfo_proto protoreflect.FileDescriptor

var file_ChessRogueGameItemInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95,
	0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x4d, 0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x49, 0x74,
	0x65, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueGameItemInfo_proto_rawDescOnce sync.Once
	file_ChessRogueGameItemInfo_proto_rawDescData = file_ChessRogueGameItemInfo_proto_rawDesc
)

func file_ChessRogueGameItemInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueGameItemInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueGameItemInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueGameItemInfo_proto_rawDescData)
	})
	return file_ChessRogueGameItemInfo_proto_rawDescData
}

var file_ChessRogueGameItemInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ChessRogueGameItemInfo_proto_goTypes = []interface{}{
	(*ChessRogueGameItemInfo)(nil), // 0: ChessRogueGameItemInfo
	nil,                            // 1: ChessRogueGameItemInfo.ItemMapEntry
}
var file_ChessRogueGameItemInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueGameItemInfo.item_map:type_name -> ChessRogueGameItemInfo.ItemMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessRogueGameItemInfo_proto_init() }
func file_ChessRogueGameItemInfo_proto_init() {
	if File_ChessRogueGameItemInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueGameItemInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueGameItemInfo); i {
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
			RawDescriptor: file_ChessRogueGameItemInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueGameItemInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueGameItemInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueGameItemInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueGameItemInfo_proto = out.File
	file_ChessRogueGameItemInfo_proto_rawDesc = nil
	file_ChessRogueGameItemInfo_proto_goTypes = nil
	file_ChessRogueGameItemInfo_proto_depIdxs = nil
}