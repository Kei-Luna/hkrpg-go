// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartPartialChallengeScRsp.proto

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

type StartPartialChallengeScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lineup       *LineupInfo   `protobuf:"bytes,9,opt,name=lineup,proto3" json:"lineup,omitempty"`
	CurChallenge *CurChallenge `protobuf:"bytes,12,opt,name=cur_challenge,json=curChallenge,proto3" json:"cur_challenge,omitempty"`
	Scene        *SceneInfo    `protobuf:"bytes,7,opt,name=scene,proto3" json:"scene,omitempty"`
	Retcode      uint32        `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *StartPartialChallengeScRsp) Reset() {
	*x = StartPartialChallengeScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartPartialChallengeScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPartialChallengeScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPartialChallengeScRsp) ProtoMessage() {}

func (x *StartPartialChallengeScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartPartialChallengeScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPartialChallengeScRsp.ProtoReflect.Descriptor instead.
func (*StartPartialChallengeScRsp) Descriptor() ([]byte, []int) {
	return file_StartPartialChallengeScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartPartialChallengeScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *StartPartialChallengeScRsp) GetCurChallenge() *CurChallenge {
	if x != nil {
		return x.CurChallenge
	}
	return nil
}

func (x *StartPartialChallengeScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *StartPartialChallengeScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_StartPartialChallengeScRsp_proto protoreflect.FileDescriptor

var file_StartPartialChallengeScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x1a, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x32, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartPartialChallengeScRsp_proto_rawDescOnce sync.Once
	file_StartPartialChallengeScRsp_proto_rawDescData = file_StartPartialChallengeScRsp_proto_rawDesc
)

func file_StartPartialChallengeScRsp_proto_rawDescGZIP() []byte {
	file_StartPartialChallengeScRsp_proto_rawDescOnce.Do(func() {
		file_StartPartialChallengeScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartPartialChallengeScRsp_proto_rawDescData)
	})
	return file_StartPartialChallengeScRsp_proto_rawDescData
}

var file_StartPartialChallengeScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartPartialChallengeScRsp_proto_goTypes = []interface{}{
	(*StartPartialChallengeScRsp)(nil), // 0: StartPartialChallengeScRsp
	(*LineupInfo)(nil),                 // 1: LineupInfo
	(*CurChallenge)(nil),               // 2: CurChallenge
	(*SceneInfo)(nil),                  // 3: SceneInfo
}
var file_StartPartialChallengeScRsp_proto_depIdxs = []int32{
	1, // 0: StartPartialChallengeScRsp.lineup:type_name -> LineupInfo
	2, // 1: StartPartialChallengeScRsp.cur_challenge:type_name -> CurChallenge
	3, // 2: StartPartialChallengeScRsp.scene:type_name -> SceneInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_StartPartialChallengeScRsp_proto_init() }
func file_StartPartialChallengeScRsp_proto_init() {
	if File_StartPartialChallengeScRsp_proto != nil {
		return
	}
	file_SceneInfo_proto_init()
	file_LineupInfo_proto_init()
	file_CurChallenge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartPartialChallengeScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPartialChallengeScRsp); i {
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
			RawDescriptor: file_StartPartialChallengeScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartPartialChallengeScRsp_proto_goTypes,
		DependencyIndexes: file_StartPartialChallengeScRsp_proto_depIdxs,
		MessageInfos:      file_StartPartialChallengeScRsp_proto_msgTypes,
	}.Build()
	File_StartPartialChallengeScRsp_proto = out.File
	file_StartPartialChallengeScRsp_proto_rawDesc = nil
	file_StartPartialChallengeScRsp_proto_goTypes = nil
	file_StartPartialChallengeScRsp_proto_depIdxs = nil
}
