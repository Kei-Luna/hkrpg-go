// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueCellSpecialType.proto

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

type ChessRogueCellSpecialType int32

const (
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_NONE      ChessRogueCellSpecialType = 0
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_LOCKED    ChessRogueCellSpecialType = 1
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_REPLICATE ChessRogueCellSpecialType = 2
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_PROTECTED ChessRogueCellSpecialType = 3
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_SEED      ChessRogueCellSpecialType = 4
	ChessRogueCellSpecialType_CHESS_ROGUE_CELL_SPECIAL_TYPE_STAMP     ChessRogueCellSpecialType = 5
)

// Enum value maps for ChessRogueCellSpecialType.
var (
	ChessRogueCellSpecialType_name = map[int32]string{
		0: "CHESS_ROGUE_CELL_SPECIAL_TYPE_NONE",
		1: "CHESS_ROGUE_CELL_SPECIAL_TYPE_LOCKED",
		2: "CHESS_ROGUE_CELL_SPECIAL_TYPE_REPLICATE",
		3: "CHESS_ROGUE_CELL_SPECIAL_TYPE_PROTECTED",
		4: "CHESS_ROGUE_CELL_SPECIAL_TYPE_SEED",
		5: "CHESS_ROGUE_CELL_SPECIAL_TYPE_STAMP",
	}
	ChessRogueCellSpecialType_value = map[string]int32{
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_NONE":      0,
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_LOCKED":    1,
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_REPLICATE": 2,
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_PROTECTED": 3,
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_SEED":      4,
		"CHESS_ROGUE_CELL_SPECIAL_TYPE_STAMP":     5,
	}
)

func (x ChessRogueCellSpecialType) Enum() *ChessRogueCellSpecialType {
	p := new(ChessRogueCellSpecialType)
	*p = x
	return p
}

func (x ChessRogueCellSpecialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChessRogueCellSpecialType) Descriptor() protoreflect.EnumDescriptor {
	return file_ChessRogueCellSpecialType_proto_enumTypes[0].Descriptor()
}

func (ChessRogueCellSpecialType) Type() protoreflect.EnumType {
	return &file_ChessRogueCellSpecialType_proto_enumTypes[0]
}

func (x ChessRogueCellSpecialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChessRogueCellSpecialType.Descriptor instead.
func (ChessRogueCellSpecialType) EnumDescriptor() ([]byte, []int) {
	return file_ChessRogueCellSpecialType_proto_rawDescGZIP(), []int{0}
}

var File_ChessRogueCellSpecialType_proto protoreflect.FileDescriptor

var file_ChessRogueCellSpecialType_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x98, 0x02, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x65, 0x6c, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x22, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43,
	0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x45, 0x53, 0x53,
	0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x2b,
	0x0a, 0x27, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x45,
	0x4c, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x43,
	0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47,
	0x55, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueCellSpecialType_proto_rawDescOnce sync.Once
	file_ChessRogueCellSpecialType_proto_rawDescData = file_ChessRogueCellSpecialType_proto_rawDesc
)

func file_ChessRogueCellSpecialType_proto_rawDescGZIP() []byte {
	file_ChessRogueCellSpecialType_proto_rawDescOnce.Do(func() {
		file_ChessRogueCellSpecialType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueCellSpecialType_proto_rawDescData)
	})
	return file_ChessRogueCellSpecialType_proto_rawDescData
}

var file_ChessRogueCellSpecialType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChessRogueCellSpecialType_proto_goTypes = []interface{}{
	(ChessRogueCellSpecialType)(0), // 0: ChessRogueCellSpecialType
}
var file_ChessRogueCellSpecialType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueCellSpecialType_proto_init() }
func file_ChessRogueCellSpecialType_proto_init() {
	if File_ChessRogueCellSpecialType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChessRogueCellSpecialType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueCellSpecialType_proto_goTypes,
		DependencyIndexes: file_ChessRogueCellSpecialType_proto_depIdxs,
		EnumInfos:         file_ChessRogueCellSpecialType_proto_enumTypes,
	}.Build()
	File_ChessRogueCellSpecialType_proto = out.File
	file_ChessRogueCellSpecialType_proto_rawDesc = nil
	file_ChessRogueCellSpecialType_proto_goTypes = nil
	file_ChessRogueCellSpecialType_proto_depIdxs = nil
}