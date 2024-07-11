// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartRogueScRsp.proto

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

type StartRogueScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scene      *SceneInfo          `protobuf:"bytes,2,opt,name=scene,proto3" json:"scene,omitempty"`
	RotateInfo *RogueMapRotateInfo `protobuf:"bytes,5,opt,name=rotate_info,json=rotateInfo,proto3" json:"rotate_info,omitempty"`
	RogueInfo  *RogueInfo          `protobuf:"bytes,14,opt,name=rogue_info,json=rogueInfo,proto3" json:"rogue_info,omitempty"`
	Lineup     *LineupInfo         `protobuf:"bytes,15,opt,name=lineup,proto3" json:"lineup,omitempty"`
	Retcode    uint32              `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *StartRogueScRsp) Reset() {
	*x = StartRogueScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartRogueScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRogueScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRogueScRsp) ProtoMessage() {}

func (x *StartRogueScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartRogueScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRogueScRsp.ProtoReflect.Descriptor instead.
func (*StartRogueScRsp) Descriptor() ([]byte, []int) {
	return file_StartRogueScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartRogueScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *StartRogueScRsp) GetRotateInfo() *RogueMapRotateInfo {
	if x != nil {
		return x.RotateInfo
	}
	return nil
}

func (x *StartRogueScRsp) GetRogueInfo() *RogueInfo {
	if x != nil {
		return x.RogueInfo
	}
	return nil
}

func (x *StartRogueScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *StartRogueScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_StartRogueScRsp_proto protoreflect.FileDescriptor

var file_StartRogueScRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x29, 0x0a, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x06,
	0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartRogueScRsp_proto_rawDescOnce sync.Once
	file_StartRogueScRsp_proto_rawDescData = file_StartRogueScRsp_proto_rawDesc
)

func file_StartRogueScRsp_proto_rawDescGZIP() []byte {
	file_StartRogueScRsp_proto_rawDescOnce.Do(func() {
		file_StartRogueScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartRogueScRsp_proto_rawDescData)
	})
	return file_StartRogueScRsp_proto_rawDescData
}

var file_StartRogueScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartRogueScRsp_proto_goTypes = []interface{}{
	(*StartRogueScRsp)(nil),    // 0: StartRogueScRsp
	(*SceneInfo)(nil),          // 1: SceneInfo
	(*RogueMapRotateInfo)(nil), // 2: RogueMapRotateInfo
	(*RogueInfo)(nil),          // 3: RogueInfo
	(*LineupInfo)(nil),         // 4: LineupInfo
}
var file_StartRogueScRsp_proto_depIdxs = []int32{
	1, // 0: StartRogueScRsp.scene:type_name -> SceneInfo
	2, // 1: StartRogueScRsp.rotate_info:type_name -> RogueMapRotateInfo
	3, // 2: StartRogueScRsp.rogue_info:type_name -> RogueInfo
	4, // 3: StartRogueScRsp.lineup:type_name -> LineupInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_StartRogueScRsp_proto_init() }
func file_StartRogueScRsp_proto_init() {
	if File_StartRogueScRsp_proto != nil {
		return
	}
	file_SceneInfo_proto_init()
	file_RogueInfo_proto_init()
	file_RogueMapRotateInfo_proto_init()
	file_LineupInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartRogueScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRogueScRsp); i {
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
			RawDescriptor: file_StartRogueScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartRogueScRsp_proto_goTypes,
		DependencyIndexes: file_StartRogueScRsp_proto_depIdxs,
		MessageInfos:      file_StartRogueScRsp_proto_msgTypes,
	}.Build()
	File_StartRogueScRsp_proto = out.File
	file_StartRogueScRsp_proto_rawDesc = nil
	file_StartRogueScRsp_proto_goTypes = nil
	file_StartRogueScRsp_proto_depIdxs = nil
}
