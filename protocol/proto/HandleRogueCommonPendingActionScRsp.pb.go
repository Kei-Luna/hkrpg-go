// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleRogueCommonPendingActionScRsp.proto

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

type HandleRogueCommonPendingActionScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueLocation uint32 `protobuf:"varint,12,opt,name=queue_location,json=queueLocation,proto3" json:"queue_location,omitempty"`
	QueuePosition uint32 `protobuf:"varint,1,opt,name=queue_position,json=queuePosition,proto3" json:"queue_position,omitempty"`
	Retcode       uint32 `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	// Types that are assignable to Action:
	//
	//	*HandleRogueCommonPendingActionScRsp_BuffSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_LIIFCFMCIBK
	//	*HandleRogueCommonPendingActionScRsp_LFJPKPLFCBN
	//	*HandleRogueCommonPendingActionScRsp_BuffRerollCallback
	//	*HandleRogueCommonPendingActionScRsp_BHCKAPOFMOO
	//	*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_FGCJLJFNGIJ
	//	*HandleRogueCommonPendingActionScRsp_GPDFGBCBAKJ
	//	*HandleRogueCommonPendingActionScRsp_LHJFFJDCFOB
	//	*HandleRogueCommonPendingActionScRsp_CJCDNDNIPOB
	//	*HandleRogueCommonPendingActionScRsp_DFEIHCPAHDH
	//	*HandleRogueCommonPendingActionScRsp_FDADIMKCLOB
	//	*HandleRogueCommonPendingActionScRsp_BonusSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback
	//	*HandleRogueCommonPendingActionScRsp_DAKOGKIJCGC
	//	*HandleRogueCommonPendingActionScRsp_DEHAKFMCMNM
	//	*HandleRogueCommonPendingActionScRsp_NAGEBCDKKAF
	Action isHandleRogueCommonPendingActionScRsp_Action `protobuf_oneof:"action"`
}

func (x *HandleRogueCommonPendingActionScRsp) Reset() {
	*x = HandleRogueCommonPendingActionScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleRogueCommonPendingActionScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleRogueCommonPendingActionScRsp) ProtoMessage() {}

func (x *HandleRogueCommonPendingActionScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleRogueCommonPendingActionScRsp.ProtoReflect.Descriptor instead.
func (*HandleRogueCommonPendingActionScRsp) Descriptor() ([]byte, []int) {
	return file_HandleRogueCommonPendingActionScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *HandleRogueCommonPendingActionScRsp) GetQueueLocation() uint32 {
	if x != nil {
		return x.QueueLocation
	}
	return 0
}

func (x *HandleRogueCommonPendingActionScRsp) GetQueuePosition() uint32 {
	if x != nil {
		return x.QueuePosition
	}
	return 0
}

func (x *HandleRogueCommonPendingActionScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (m *HandleRogueCommonPendingActionScRsp) GetAction() isHandleRogueCommonPendingActionScRsp_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBuffSelectCallback() *RogueBuffSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BuffSelectCallback); ok {
		return x.BuffSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetLIIFCFMCIBK() *GNJEHFCCIAB {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_LIIFCFMCIBK); ok {
		return x.LIIFCFMCIBK
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetLFJPKPLFCBN() *IHBKINEFFAO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_LFJPKPLFCBN); ok {
		return x.LFJPKPLFCBN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBuffRerollCallback() *RogueBuffRerollCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BuffRerollCallback); ok {
		return x.BuffRerollCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBHCKAPOFMOO() *FFCCLEECGPP {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BHCKAPOFMOO); ok {
		return x.BHCKAPOFMOO
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetMiracleSelectCallback() *RogueMiracleSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback); ok {
		return x.MiracleSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetFGCJLJFNGIJ() *ABFFCPMIJIL {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_FGCJLJFNGIJ); ok {
		return x.FGCJLJFNGIJ
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetGPDFGBCBAKJ() *MGMOLIHMEOJ {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_GPDFGBCBAKJ); ok {
		return x.GPDFGBCBAKJ
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetLHJFFJDCFOB() *DOKPJBJCMAO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_LHJFFJDCFOB); ok {
		return x.LHJFFJDCFOB
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetCJCDNDNIPOB() *FAODICLJJKL {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_CJCDNDNIPOB); ok {
		return x.CJCDNDNIPOB
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDFEIHCPAHDH() *DBCDLKBGGDP {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DFEIHCPAHDH); ok {
		return x.DFEIHCPAHDH
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetFDADIMKCLOB() *OEKIOPKIGEM {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_FDADIMKCLOB); ok {
		return x.FDADIMKCLOB
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBonusSelectCallback() *RogueBonusSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BonusSelectCallback); ok {
		return x.BonusSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetRogueTournFormulaCallback() *RogueTournFormulaCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback); ok {
		return x.RogueTournFormulaCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDAKOGKIJCGC() *HJAINGCDKMH {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DAKOGKIJCGC); ok {
		return x.DAKOGKIJCGC
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDEHAKFMCMNM() *DCINOGOGBHF {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DEHAKFMCMNM); ok {
		return x.DEHAKFMCMNM
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetNAGEBCDKKAF() *BPEGLDKMPGA {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_NAGEBCDKKAF); ok {
		return x.NAGEBCDKKAF
	}
	return nil
}

type isHandleRogueCommonPendingActionScRsp_Action interface {
	isHandleRogueCommonPendingActionScRsp_Action()
}

type HandleRogueCommonPendingActionScRsp_BuffSelectCallback struct {
	BuffSelectCallback *RogueBuffSelectCallback `protobuf:"bytes,688,opt,name=buff_select_callback,json=buffSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_LIIFCFMCIBK struct {
	LIIFCFMCIBK *GNJEHFCCIAB `protobuf:"bytes,1933,opt,name=LIIFCFMCIBK,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_LFJPKPLFCBN struct {
	LFJPKPLFCBN *IHBKINEFFAO `protobuf:"bytes,1646,opt,name=LFJPKPLFCBN,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_BuffRerollCallback struct {
	BuffRerollCallback *RogueBuffRerollCallback `protobuf:"bytes,523,opt,name=buff_reroll_callback,json=buffRerollCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_BHCKAPOFMOO struct {
	BHCKAPOFMOO *FFCCLEECGPP `protobuf:"bytes,741,opt,name=BHCKAPOFMOO,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_MiracleSelectCallback struct {
	MiracleSelectCallback *RogueMiracleSelectCallback `protobuf:"bytes,214,opt,name=miracle_select_callback,json=miracleSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_FGCJLJFNGIJ struct {
	FGCJLJFNGIJ *ABFFCPMIJIL `protobuf:"bytes,1353,opt,name=FGCJLJFNGIJ,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_GPDFGBCBAKJ struct {
	GPDFGBCBAKJ *MGMOLIHMEOJ `protobuf:"bytes,183,opt,name=GPDFGBCBAKJ,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_LHJFFJDCFOB struct {
	LHJFFJDCFOB *DOKPJBJCMAO `protobuf:"bytes,1459,opt,name=LHJFFJDCFOB,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_CJCDNDNIPOB struct {
	CJCDNDNIPOB *FAODICLJJKL `protobuf:"bytes,544,opt,name=CJCDNDNIPOB,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DFEIHCPAHDH struct {
	DFEIHCPAHDH *DBCDLKBGGDP `protobuf:"bytes,496,opt,name=DFEIHCPAHDH,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_FDADIMKCLOB struct {
	FDADIMKCLOB *OEKIOPKIGEM `protobuf:"bytes,1994,opt,name=FDADIMKCLOB,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_BonusSelectCallback struct {
	BonusSelectCallback *RogueBonusSelectCallback `protobuf:"bytes,59,opt,name=bonus_select_callback,json=bonusSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback struct {
	RogueTournFormulaCallback *RogueTournFormulaCallback `protobuf:"bytes,991,opt,name=rogue_tourn_formula_callback,json=rogueTournFormulaCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DAKOGKIJCGC struct {
	DAKOGKIJCGC *HJAINGCDKMH `protobuf:"bytes,471,opt,name=DAKOGKIJCGC,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DEHAKFMCMNM struct {
	DEHAKFMCMNM *DCINOGOGBHF `protobuf:"bytes,235,opt,name=DEHAKFMCMNM,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_NAGEBCDKKAF struct {
	NAGEBCDKKAF *BPEGLDKMPGA `protobuf:"bytes,167,opt,name=NAGEBCDKKAF,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionScRsp_BuffSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_LIIFCFMCIBK) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_LFJPKPLFCBN) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_BuffRerollCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_BHCKAPOFMOO) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_FGCJLJFNGIJ) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_GPDFGBCBAKJ) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_LHJFFJDCFOB) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_CJCDNDNIPOB) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DFEIHCPAHDH) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_FDADIMKCLOB) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_BonusSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DAKOGKIJCGC) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DEHAKFMCMNM) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_NAGEBCDKKAF) isHandleRogueCommonPendingActionScRsp_Action() {
}

var File_HandleRogueCommonPendingActionScRsp_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionScRsp_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x47, 0x4d,
	0x4f, 0x4c, 0x49, 0x48, 0x4d, 0x45, 0x4f, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42,
	0x50, 0x45, 0x47, 0x4c, 0x44, 0x4b, 0x4d, 0x50, 0x47, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x46, 0x46, 0x43, 0x43, 0x4c, 0x45, 0x45, 0x43, 0x47, 0x50, 0x50, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x48, 0x42, 0x4b, 0x49, 0x4e, 0x45, 0x46, 0x46,
	0x41, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x43, 0x49, 0x4e, 0x4f, 0x47,
	0x4f, 0x47, 0x42, 0x48, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4e, 0x4a,
	0x45, 0x48, 0x46, 0x43, 0x43, 0x49, 0x41, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x44, 0x4f, 0x4b, 0x50, 0x4a, 0x42, 0x4a, 0x43, 0x4d, 0x41, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72,
	0x6f, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4f, 0x45, 0x4b, 0x49, 0x4f, 0x50, 0x4b, 0x49, 0x47, 0x45, 0x4d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x46, 0x46, 0x43, 0x50, 0x4d, 0x49, 0x4a, 0x49,
	0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x42, 0x43, 0x44, 0x4c, 0x4b, 0x42,
	0x47, 0x47, 0x44, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4a, 0x41, 0x49,
	0x4e, 0x47, 0x43, 0x44, 0x4b, 0x4d, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46,
	0x41, 0x4f, 0x44, 0x49, 0x43, 0x4c, 0x4a, 0x4a, 0x4b, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa2, 0x09, 0x0a, 0x23, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x4d, 0x0a, 0x14, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0xb0, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x12, 0x62, 0x75, 0x66,
	0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x31, 0x0a, 0x0b, 0x4c, 0x49, 0x49, 0x46, 0x43, 0x46, 0x4d, 0x43, 0x49, 0x42, 0x4b, 0x18, 0x8d,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4e, 0x4a, 0x45, 0x48, 0x46, 0x43, 0x43,
	0x49, 0x41, 0x42, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x49, 0x49, 0x46, 0x43, 0x46, 0x4d, 0x43, 0x49,
	0x42, 0x4b, 0x12, 0x31, 0x0a, 0x0b, 0x4c, 0x46, 0x4a, 0x50, 0x4b, 0x50, 0x4c, 0x46, 0x43, 0x42,
	0x4e, 0x18, 0xee, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x48, 0x42, 0x4b, 0x49,
	0x4e, 0x45, 0x46, 0x46, 0x41, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x46, 0x4a, 0x50, 0x4b, 0x50,
	0x4c, 0x46, 0x43, 0x42, 0x4e, 0x12, 0x4d, 0x0a, 0x14, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72, 0x65,
	0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x8b, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00,
	0x52, 0x12, 0x62, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x42, 0x48, 0x43, 0x4b, 0x41, 0x50, 0x4f, 0x46,
	0x4d, 0x4f, 0x4f, 0x18, 0xe5, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x46, 0x43,
	0x43, 0x4c, 0x45, 0x45, 0x43, 0x47, 0x50, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x42, 0x48, 0x43, 0x4b,
	0x41, 0x50, 0x4f, 0x46, 0x4d, 0x4f, 0x4f, 0x12, 0x56, 0x0a, 0x17, 0x6d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0xd6, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x15, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x31, 0x0a, 0x0b, 0x46, 0x47, 0x43, 0x4a, 0x4c, 0x4a, 0x46, 0x4e, 0x47, 0x49, 0x4a, 0x18, 0xc9,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x42, 0x46, 0x46, 0x43, 0x50, 0x4d, 0x49,
	0x4a, 0x49, 0x4c, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x47, 0x43, 0x4a, 0x4c, 0x4a, 0x46, 0x4e, 0x47,
	0x49, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x50, 0x44, 0x46, 0x47, 0x42, 0x43, 0x42, 0x41, 0x4b,
	0x4a, 0x18, 0xb7, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x47, 0x4d, 0x4f, 0x4c,
	0x49, 0x48, 0x4d, 0x45, 0x4f, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x50, 0x44, 0x46, 0x47, 0x42,
	0x43, 0x42, 0x41, 0x4b, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x4c, 0x48, 0x4a, 0x46, 0x46, 0x4a, 0x44,
	0x43, 0x46, 0x4f, 0x42, 0x18, 0xb3, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4f,
	0x4b, 0x50, 0x4a, 0x42, 0x4a, 0x43, 0x4d, 0x41, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x48, 0x4a,
	0x46, 0x46, 0x4a, 0x44, 0x43, 0x46, 0x4f, 0x42, 0x12, 0x31, 0x0a, 0x0b, 0x43, 0x4a, 0x43, 0x44,
	0x4e, 0x44, 0x4e, 0x49, 0x50, 0x4f, 0x42, 0x18, 0xa0, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x46, 0x41, 0x4f, 0x44, 0x49, 0x43, 0x4c, 0x4a, 0x4a, 0x4b, 0x4c, 0x48, 0x00, 0x52, 0x0b,
	0x43, 0x4a, 0x43, 0x44, 0x4e, 0x44, 0x4e, 0x49, 0x50, 0x4f, 0x42, 0x12, 0x31, 0x0a, 0x0b, 0x44,
	0x46, 0x45, 0x49, 0x48, 0x43, 0x50, 0x41, 0x48, 0x44, 0x48, 0x18, 0xf0, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x42, 0x43, 0x44, 0x4c, 0x4b, 0x42, 0x47, 0x47, 0x44, 0x50, 0x48,
	0x00, 0x52, 0x0b, 0x44, 0x46, 0x45, 0x49, 0x48, 0x43, 0x50, 0x41, 0x48, 0x44, 0x48, 0x12, 0x31,
	0x0a, 0x0b, 0x46, 0x44, 0x41, 0x44, 0x49, 0x4d, 0x4b, 0x43, 0x4c, 0x4f, 0x42, 0x18, 0xca, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x45, 0x4b, 0x49, 0x4f, 0x50, 0x4b, 0x49, 0x47,
	0x45, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x44, 0x41, 0x44, 0x49, 0x4d, 0x4b, 0x43, 0x4c, 0x4f,
	0x42, 0x12, 0x4f, 0x0a, 0x15, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x13, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x5e, 0x0a, 0x1c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0xdf, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x19, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x41, 0x4b, 0x4f, 0x47, 0x4b, 0x49, 0x4a, 0x43, 0x47,
	0x43, 0x18, 0xd7, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4a, 0x41, 0x49, 0x4e,
	0x47, 0x43, 0x44, 0x4b, 0x4d, 0x48, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x41, 0x4b, 0x4f, 0x47, 0x4b,
	0x49, 0x4a, 0x43, 0x47, 0x43, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x45, 0x48, 0x41, 0x4b, 0x46, 0x4d,
	0x43, 0x4d, 0x4e, 0x4d, 0x18, 0xeb, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x43,
	0x49, 0x4e, 0x4f, 0x47, 0x4f, 0x47, 0x42, 0x48, 0x46, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x45, 0x48,
	0x41, 0x4b, 0x46, 0x4d, 0x43, 0x4d, 0x4e, 0x4d, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x41, 0x47, 0x45,
	0x42, 0x43, 0x44, 0x4b, 0x4b, 0x41, 0x46, 0x18, 0xa7, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x42, 0x50, 0x45, 0x47, 0x4c, 0x44, 0x4b, 0x4d, 0x50, 0x47, 0x41, 0x48, 0x00, 0x52, 0x0b,
	0x4e, 0x41, 0x47, 0x45, 0x42, 0x43, 0x44, 0x4b, 0x4b, 0x41, 0x46, 0x42, 0x08, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescOnce sync.Once
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescData = file_HandleRogueCommonPendingActionScRsp_proto_rawDesc
)

func file_HandleRogueCommonPendingActionScRsp_proto_rawDescGZIP() []byte {
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescOnce.Do(func() {
		file_HandleRogueCommonPendingActionScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleRogueCommonPendingActionScRsp_proto_rawDescData)
	})
	return file_HandleRogueCommonPendingActionScRsp_proto_rawDescData
}

var file_HandleRogueCommonPendingActionScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleRogueCommonPendingActionScRsp_proto_goTypes = []interface{}{
	(*HandleRogueCommonPendingActionScRsp)(nil), // 0: HandleRogueCommonPendingActionScRsp
	(*RogueBuffSelectCallback)(nil),             // 1: RogueBuffSelectCallback
	(*GNJEHFCCIAB)(nil),                         // 2: GNJEHFCCIAB
	(*IHBKINEFFAO)(nil),                         // 3: IHBKINEFFAO
	(*RogueBuffRerollCallback)(nil),             // 4: RogueBuffRerollCallback
	(*FFCCLEECGPP)(nil),                         // 5: FFCCLEECGPP
	(*RogueMiracleSelectCallback)(nil),          // 6: RogueMiracleSelectCallback
	(*ABFFCPMIJIL)(nil),                         // 7: ABFFCPMIJIL
	(*MGMOLIHMEOJ)(nil),                         // 8: MGMOLIHMEOJ
	(*DOKPJBJCMAO)(nil),                         // 9: DOKPJBJCMAO
	(*FAODICLJJKL)(nil),                         // 10: FAODICLJJKL
	(*DBCDLKBGGDP)(nil),                         // 11: DBCDLKBGGDP
	(*OEKIOPKIGEM)(nil),                         // 12: OEKIOPKIGEM
	(*RogueBonusSelectCallback)(nil),            // 13: RogueBonusSelectCallback
	(*RogueTournFormulaCallback)(nil),           // 14: RogueTournFormulaCallback
	(*HJAINGCDKMH)(nil),                         // 15: HJAINGCDKMH
	(*DCINOGOGBHF)(nil),                         // 16: DCINOGOGBHF
	(*BPEGLDKMPGA)(nil),                         // 17: BPEGLDKMPGA
}
var file_HandleRogueCommonPendingActionScRsp_proto_depIdxs = []int32{
	1,  // 0: HandleRogueCommonPendingActionScRsp.buff_select_callback:type_name -> RogueBuffSelectCallback
	2,  // 1: HandleRogueCommonPendingActionScRsp.LIIFCFMCIBK:type_name -> GNJEHFCCIAB
	3,  // 2: HandleRogueCommonPendingActionScRsp.LFJPKPLFCBN:type_name -> IHBKINEFFAO
	4,  // 3: HandleRogueCommonPendingActionScRsp.buff_reroll_callback:type_name -> RogueBuffRerollCallback
	5,  // 4: HandleRogueCommonPendingActionScRsp.BHCKAPOFMOO:type_name -> FFCCLEECGPP
	6,  // 5: HandleRogueCommonPendingActionScRsp.miracle_select_callback:type_name -> RogueMiracleSelectCallback
	7,  // 6: HandleRogueCommonPendingActionScRsp.FGCJLJFNGIJ:type_name -> ABFFCPMIJIL
	8,  // 7: HandleRogueCommonPendingActionScRsp.GPDFGBCBAKJ:type_name -> MGMOLIHMEOJ
	9,  // 8: HandleRogueCommonPendingActionScRsp.LHJFFJDCFOB:type_name -> DOKPJBJCMAO
	10, // 9: HandleRogueCommonPendingActionScRsp.CJCDNDNIPOB:type_name -> FAODICLJJKL
	11, // 10: HandleRogueCommonPendingActionScRsp.DFEIHCPAHDH:type_name -> DBCDLKBGGDP
	12, // 11: HandleRogueCommonPendingActionScRsp.FDADIMKCLOB:type_name -> OEKIOPKIGEM
	13, // 12: HandleRogueCommonPendingActionScRsp.bonus_select_callback:type_name -> RogueBonusSelectCallback
	14, // 13: HandleRogueCommonPendingActionScRsp.rogue_tourn_formula_callback:type_name -> RogueTournFormulaCallback
	15, // 14: HandleRogueCommonPendingActionScRsp.DAKOGKIJCGC:type_name -> HJAINGCDKMH
	16, // 15: HandleRogueCommonPendingActionScRsp.DEHAKFMCMNM:type_name -> DCINOGOGBHF
	17, // 16: HandleRogueCommonPendingActionScRsp.NAGEBCDKKAF:type_name -> BPEGLDKMPGA
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionScRsp_proto_init() }
func file_HandleRogueCommonPendingActionScRsp_proto_init() {
	if File_HandleRogueCommonPendingActionScRsp_proto != nil {
		return
	}
	file_MGMOLIHMEOJ_proto_init()
	file_RogueBuffSelectCallback_proto_init()
	file_BPEGLDKMPGA_proto_init()
	file_RogueMiracleSelectCallback_proto_init()
	file_RogueBonusSelectCallback_proto_init()
	file_FFCCLEECGPP_proto_init()
	file_IHBKINEFFAO_proto_init()
	file_DCINOGOGBHF_proto_init()
	file_GNJEHFCCIAB_proto_init()
	file_DOKPJBJCMAO_proto_init()
	file_RogueTournFormulaCallback_proto_init()
	file_RogueBuffRerollCallback_proto_init()
	file_OEKIOPKIGEM_proto_init()
	file_ABFFCPMIJIL_proto_init()
	file_DBCDLKBGGDP_proto_init()
	file_HJAINGCDKMH_proto_init()
	file_FAODICLJJKL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleRogueCommonPendingActionScRsp); i {
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
	file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HandleRogueCommonPendingActionScRsp_BuffSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_LIIFCFMCIBK)(nil),
		(*HandleRogueCommonPendingActionScRsp_LFJPKPLFCBN)(nil),
		(*HandleRogueCommonPendingActionScRsp_BuffRerollCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_BHCKAPOFMOO)(nil),
		(*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_FGCJLJFNGIJ)(nil),
		(*HandleRogueCommonPendingActionScRsp_GPDFGBCBAKJ)(nil),
		(*HandleRogueCommonPendingActionScRsp_LHJFFJDCFOB)(nil),
		(*HandleRogueCommonPendingActionScRsp_CJCDNDNIPOB)(nil),
		(*HandleRogueCommonPendingActionScRsp_DFEIHCPAHDH)(nil),
		(*HandleRogueCommonPendingActionScRsp_FDADIMKCLOB)(nil),
		(*HandleRogueCommonPendingActionScRsp_BonusSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_DAKOGKIJCGC)(nil),
		(*HandleRogueCommonPendingActionScRsp_DEHAKFMCMNM)(nil),
		(*HandleRogueCommonPendingActionScRsp_NAGEBCDKKAF)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HandleRogueCommonPendingActionScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleRogueCommonPendingActionScRsp_proto_goTypes,
		DependencyIndexes: file_HandleRogueCommonPendingActionScRsp_proto_depIdxs,
		MessageInfos:      file_HandleRogueCommonPendingActionScRsp_proto_msgTypes,
	}.Build()
	File_HandleRogueCommonPendingActionScRsp_proto = out.File
	file_HandleRogueCommonPendingActionScRsp_proto_rawDesc = nil
	file_HandleRogueCommonPendingActionScRsp_proto_goTypes = nil
	file_HandleRogueCommonPendingActionScRsp_proto_depIdxs = nil
}
