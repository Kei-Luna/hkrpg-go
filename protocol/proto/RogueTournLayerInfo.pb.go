// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournLayerInfo.proto

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

type RogueTournLayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        RogueTournLevelStatus `protobuf:"varint,8,opt,name=status,proto3,enum=RogueTournLevelStatus" json:"status,omitempty"`
	LayerInfoList []*RogueTournLayer    `protobuf:"bytes,2,rep,name=layer_info_list,json=layerInfoList,proto3" json:"layer_info_list,omitempty"`
	Reason        GIOGIPJLONO           `protobuf:"varint,5,opt,name=reason,proto3,enum=GIOGIPJLONO" json:"reason,omitempty"`
	CurLayerIndex uint32                `protobuf:"varint,6,opt,name=cur_layer_index,json=curLayerIndex,proto3" json:"cur_layer_index,omitempty"`
}

func (x *RogueTournLayerInfo) Reset() {
	*x = RogueTournLayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournLayerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournLayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournLayerInfo) ProtoMessage() {}

func (x *RogueTournLayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournLayerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournLayerInfo.ProtoReflect.Descriptor instead.
func (*RogueTournLayerInfo) Descriptor() ([]byte, []int) {
	return file_RogueTournLayerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournLayerInfo) GetStatus() RogueTournLevelStatus {
	if x != nil {
		return x.Status
	}
	return RogueTournLevelStatus_ROGUE_TOURN_LEVEL_STATUS_NONE
}

func (x *RogueTournLayerInfo) GetLayerInfoList() []*RogueTournLayer {
	if x != nil {
		return x.LayerInfoList
	}
	return nil
}

func (x *RogueTournLayerInfo) GetReason() GIOGIPJLONO {
	if x != nil {
		return x.Reason
	}
	return GIOGIPJLONO_ROGUE_TOURN_SETTLE_REASON_NONE
}

func (x *RogueTournLayerInfo) GetCurLayerIndex() uint32 {
	if x != nil {
		return x.CurLayerIndex
	}
	return 0
}

var File_RogueTournLayerInfo_proto protoreflect.FileDescriptor

var file_RogueTournLayerInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x49, 0x4f,
	0x47, 0x49, 0x50, 0x4a, 0x4c, 0x4f, 0x4e, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0d, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x47, 0x49, 0x4f, 0x47, 0x49, 0x50, 0x4a, 0x4c, 0x4f,
	0x4e, 0x4f, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75,
	0x72, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournLayerInfo_proto_rawDescOnce sync.Once
	file_RogueTournLayerInfo_proto_rawDescData = file_RogueTournLayerInfo_proto_rawDesc
)

func file_RogueTournLayerInfo_proto_rawDescGZIP() []byte {
	file_RogueTournLayerInfo_proto_rawDescOnce.Do(func() {
		file_RogueTournLayerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournLayerInfo_proto_rawDescData)
	})
	return file_RogueTournLayerInfo_proto_rawDescData
}

var file_RogueTournLayerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournLayerInfo_proto_goTypes = []interface{}{
	(*RogueTournLayerInfo)(nil), // 0: RogueTournLayerInfo
	(RogueTournLevelStatus)(0),  // 1: RogueTournLevelStatus
	(*RogueTournLayer)(nil),     // 2: RogueTournLayer
	(GIOGIPJLONO)(0),            // 3: GIOGIPJLONO
}
var file_RogueTournLayerInfo_proto_depIdxs = []int32{
	1, // 0: RogueTournLayerInfo.status:type_name -> RogueTournLevelStatus
	2, // 1: RogueTournLayerInfo.layer_info_list:type_name -> RogueTournLayer
	3, // 2: RogueTournLayerInfo.reason:type_name -> GIOGIPJLONO
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RogueTournLayerInfo_proto_init() }
func file_RogueTournLayerInfo_proto_init() {
	if File_RogueTournLayerInfo_proto != nil {
		return
	}
	file_GIOGIPJLONO_proto_init()
	file_RogueTournLayer_proto_init()
	file_RogueTournLevelStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournLayerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournLayerInfo); i {
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
			RawDescriptor: file_RogueTournLayerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournLayerInfo_proto_goTypes,
		DependencyIndexes: file_RogueTournLayerInfo_proto_depIdxs,
		MessageInfos:      file_RogueTournLayerInfo_proto_msgTypes,
	}.Build()
	File_RogueTournLayerInfo_proto = out.File
	file_RogueTournLayerInfo_proto_rawDesc = nil
	file_RogueTournLayerInfo_proto_goTypes = nil
	file_RogueTournLayerInfo_proto_depIdxs = nil
}
