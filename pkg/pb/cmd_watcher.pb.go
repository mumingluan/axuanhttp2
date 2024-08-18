// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_watcher.proto

package pb

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

type PushTipsState int32

const (
	PushTipsState_PUSH_TIPS_STATE_NONE   PushTipsState = 0
	PushTipsState_PUSH_TIPS_STATE_START  PushTipsState = 1
	PushTipsState_PUSH_TIPS_STATE_READ   PushTipsState = 2
	PushTipsState_PUSH_TIPS_STATE_FINISH PushTipsState = 3
)

// Enum value maps for PushTipsState.
var (
	PushTipsState_name = map[int32]string{
		0: "PUSH_TIPS_STATE_NONE",
		1: "PUSH_TIPS_STATE_START",
		2: "PUSH_TIPS_STATE_READ",
		3: "PUSH_TIPS_STATE_FINISH",
	}
	PushTipsState_value = map[string]int32{
		"PUSH_TIPS_STATE_NONE":   0,
		"PUSH_TIPS_STATE_START":  1,
		"PUSH_TIPS_STATE_READ":   2,
		"PUSH_TIPS_STATE_FINISH": 3,
	}
)

func (x PushTipsState) Enum() *PushTipsState {
	p := new(PushTipsState)
	*p = x
	return p
}

func (x PushTipsState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushTipsState) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_watcher_proto_enumTypes[0].Descriptor()
}

func (PushTipsState) Type() protoreflect.EnumType {
	return &file_cmd_cmd_watcher_proto_enumTypes[0]
}

func (x PushTipsState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushTipsState.Descriptor instead.
func (PushTipsState) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{0}
}

type WatcherAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatcherList []uint32 `protobuf:"varint,4,rep,packed,name=watcher_list,json=watcherList,proto3" json:"watcher_list,omitempty"`
}

func (x *WatcherAllDataNotify) Reset() {
	*x = WatcherAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherAllDataNotify) ProtoMessage() {}

func (x *WatcherAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherAllDataNotify.ProtoReflect.Descriptor instead.
func (*WatcherAllDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *WatcherAllDataNotify) GetWatcherList() []uint32 {
	if x != nil {
		return x.WatcherList
	}
	return nil
}

type WatcherChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemovedWatcherList []uint32 `protobuf:"varint,2,rep,packed,name=removed_watcher_list,json=removedWatcherList,proto3" json:"removed_watcher_list,omitempty"`
	NewWatcherList     []uint32 `protobuf:"varint,15,rep,packed,name=new_watcher_list,json=newWatcherList,proto3" json:"new_watcher_list,omitempty"`
}

func (x *WatcherChangeNotify) Reset() {
	*x = WatcherChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherChangeNotify) ProtoMessage() {}

func (x *WatcherChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherChangeNotify.ProtoReflect.Descriptor instead.
func (*WatcherChangeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *WatcherChangeNotify) GetRemovedWatcherList() []uint32 {
	if x != nil {
		return x.RemovedWatcherList
	}
	return nil
}

func (x *WatcherChangeNotify) GetNewWatcherList() []uint32 {
	if x != nil {
		return x.NewWatcherList
	}
	return nil
}

type WatcherEventNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddProgress uint32 `protobuf:"varint,6,opt,name=add_progress,json=addProgress,proto3" json:"add_progress,omitempty"`
	WatcherId   uint32 `protobuf:"varint,9,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
}

func (x *WatcherEventNotify) Reset() {
	*x = WatcherEventNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherEventNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherEventNotify) ProtoMessage() {}

func (x *WatcherEventNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherEventNotify.ProtoReflect.Descriptor instead.
func (*WatcherEventNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{2}
}

func (x *WatcherEventNotify) GetAddProgress() uint32 {
	if x != nil {
		return x.AddProgress
	}
	return 0
}

func (x *WatcherEventNotify) GetWatcherId() uint32 {
	if x != nil {
		return x.WatcherId
	}
	return 0
}

type WatcherEventTypeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParamList          []uint32 `protobuf:"varint,14,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
	AddProgress        uint32   `protobuf:"varint,15,opt,name=add_progress,json=addProgress,proto3" json:"add_progress,omitempty"`
	WatcherTriggerType uint32   `protobuf:"varint,11,opt,name=watcher_trigger_type,json=watcherTriggerType,proto3" json:"watcher_trigger_type,omitempty"`
}

func (x *WatcherEventTypeNotify) Reset() {
	*x = WatcherEventTypeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherEventTypeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherEventTypeNotify) ProtoMessage() {}

func (x *WatcherEventTypeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherEventTypeNotify.ProtoReflect.Descriptor instead.
func (*WatcherEventTypeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{3}
}

func (x *WatcherEventTypeNotify) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

func (x *WatcherEventTypeNotify) GetAddProgress() uint32 {
	if x != nil {
		return x.AddProgress
	}
	return 0
}

func (x *WatcherEventTypeNotify) GetWatcherTriggerType() uint32 {
	if x != nil {
		return x.WatcherTriggerType
	}
	return 0
}

type WatcherEventStageNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddProgress uint32 `protobuf:"varint,4,opt,name=add_progress,json=addProgress,proto3" json:"add_progress,omitempty"`
	Stage       uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	WatcherId   uint32 `protobuf:"varint,12,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
}

func (x *WatcherEventStageNotify) Reset() {
	*x = WatcherEventStageNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatcherEventStageNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherEventStageNotify) ProtoMessage() {}

func (x *WatcherEventStageNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherEventStageNotify.ProtoReflect.Descriptor instead.
func (*WatcherEventStageNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{4}
}

func (x *WatcherEventStageNotify) GetAddProgress() uint32 {
	if x != nil {
		return x.AddProgress
	}
	return 0
}

func (x *WatcherEventStageNotify) GetStage() uint32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *WatcherEventStageNotify) GetWatcherId() uint32 {
	if x != nil {
		return x.WatcherId
	}
	return 0
}

type PushTipsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsId uint32 `protobuf:"varint,13,opt,name=push_tips_id,json=pushTipsId,proto3" json:"push_tips_id,omitempty"`
	State      uint32 `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *PushTipsData) Reset() {
	*x = PushTipsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsData) ProtoMessage() {}

func (x *PushTipsData) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsData.ProtoReflect.Descriptor instead.
func (*PushTipsData) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{5}
}

func (x *PushTipsData) GetPushTipsId() uint32 {
	if x != nil {
		return x.PushTipsId
	}
	return 0
}

func (x *PushTipsData) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type PushTipsAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsList []*PushTipsData `protobuf:"bytes,4,rep,name=push_tips_list,json=pushTipsList,proto3" json:"push_tips_list,omitempty"`
}

func (x *PushTipsAllDataNotify) Reset() {
	*x = PushTipsAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsAllDataNotify) ProtoMessage() {}

func (x *PushTipsAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsAllDataNotify.ProtoReflect.Descriptor instead.
func (*PushTipsAllDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{6}
}

func (x *PushTipsAllDataNotify) GetPushTipsList() []*PushTipsData {
	if x != nil {
		return x.PushTipsList
	}
	return nil
}

type PushTipsChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsList []*PushTipsData `protobuf:"bytes,9,rep,name=push_tips_list,json=pushTipsList,proto3" json:"push_tips_list,omitempty"`
}

func (x *PushTipsChangeNotify) Reset() {
	*x = PushTipsChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsChangeNotify) ProtoMessage() {}

func (x *PushTipsChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsChangeNotify.ProtoReflect.Descriptor instead.
func (*PushTipsChangeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{7}
}

func (x *PushTipsChangeNotify) GetPushTipsList() []*PushTipsData {
	if x != nil {
		return x.PushTipsList
	}
	return nil
}

type PushTipsReadFinishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsId uint32 `protobuf:"varint,11,opt,name=push_tips_id,json=pushTipsId,proto3" json:"push_tips_id,omitempty"`
}

func (x *PushTipsReadFinishReq) Reset() {
	*x = PushTipsReadFinishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsReadFinishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsReadFinishReq) ProtoMessage() {}

func (x *PushTipsReadFinishReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsReadFinishReq.ProtoReflect.Descriptor instead.
func (*PushTipsReadFinishReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{8}
}

func (x *PushTipsReadFinishReq) GetPushTipsId() uint32 {
	if x != nil {
		return x.PushTipsId
	}
	return 0
}

type PushTipsReadFinishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsId uint32 `protobuf:"varint,3,opt,name=push_tips_id,json=pushTipsId,proto3" json:"push_tips_id,omitempty"`
	Retcode    int32  `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PushTipsReadFinishRsp) Reset() {
	*x = PushTipsReadFinishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushTipsReadFinishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushTipsReadFinishRsp) ProtoMessage() {}

func (x *PushTipsReadFinishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushTipsReadFinishRsp.ProtoReflect.Descriptor instead.
func (*PushTipsReadFinishRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{9}
}

func (x *PushTipsReadFinishRsp) GetPushTipsId() uint32 {
	if x != nil {
		return x.PushTipsId
	}
	return 0
}

func (x *PushTipsReadFinishRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type GetPushTipsRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushTipsIdList []uint32 `protobuf:"varint,4,rep,packed,name=push_tips_id_list,json=pushTipsIdList,proto3" json:"push_tips_id_list,omitempty"`
}

func (x *GetPushTipsRewardReq) Reset() {
	*x = GetPushTipsRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushTipsRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushTipsRewardReq) ProtoMessage() {}

func (x *GetPushTipsRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushTipsRewardReq.ProtoReflect.Descriptor instead.
func (*GetPushTipsRewardReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{10}
}

func (x *GetPushTipsRewardReq) GetPushTipsIdList() []uint32 {
	if x != nil {
		return x.PushTipsIdList
	}
	return nil
}

type GetPushTipsRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode        int32    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PushTipsIdList []uint32 `protobuf:"varint,9,rep,packed,name=push_tips_id_list,json=pushTipsIdList,proto3" json:"push_tips_id_list,omitempty"`
}

func (x *GetPushTipsRewardRsp) Reset() {
	*x = GetPushTipsRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_watcher_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPushTipsRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPushTipsRewardRsp) ProtoMessage() {}

func (x *GetPushTipsRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_watcher_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPushTipsRewardRsp.ProtoReflect.Descriptor instead.
func (*GetPushTipsRewardRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_watcher_proto_rawDescGZIP(), []int{11}
}

func (x *GetPushTipsRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPushTipsRewardRsp) GetPushTipsIdList() []uint32 {
	if x != nil {
		return x.PushTipsIdList
	}
	return nil
}

var File_cmd_cmd_watcher_proto protoreflect.FileDescriptor

var file_cmd_cmd_watcher_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x14, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x71, 0x0a, 0x13, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x65,
	0x77, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x12,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x71, 0x0a, 0x17, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69,
	0x70, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x75,
	0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x52,
	0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x51, 0x0a, 0x14, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0e, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x54,
	0x69, 0x70, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x20,
	0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x49, 0x64,
	0x22, 0x53, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x61, 0x64,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x75, 0x73,
	0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a,
	0x11, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69,
	0x70, 0x73, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x11, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70, 0x73, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x7a, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x54, 0x69, 0x70,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x54,
	0x49, 0x50, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x54, 0x49, 0x50, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x55, 0x53, 0x48, 0x5f, 0x54, 0x49, 0x50, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x54, 0x49,
	0x50, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10,
	0x03, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b,
	0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_watcher_proto_rawDescOnce sync.Once
	file_cmd_cmd_watcher_proto_rawDescData = file_cmd_cmd_watcher_proto_rawDesc
)

func file_cmd_cmd_watcher_proto_rawDescGZIP() []byte {
	file_cmd_cmd_watcher_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_watcher_proto_rawDescData)
	})
	return file_cmd_cmd_watcher_proto_rawDescData
}

var file_cmd_cmd_watcher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_cmd_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cmd_cmd_watcher_proto_goTypes = []interface{}{
	(PushTipsState)(0),              // 0: proto.PushTipsState
	(*WatcherAllDataNotify)(nil),    // 1: proto.WatcherAllDataNotify
	(*WatcherChangeNotify)(nil),     // 2: proto.WatcherChangeNotify
	(*WatcherEventNotify)(nil),      // 3: proto.WatcherEventNotify
	(*WatcherEventTypeNotify)(nil),  // 4: proto.WatcherEventTypeNotify
	(*WatcherEventStageNotify)(nil), // 5: proto.WatcherEventStageNotify
	(*PushTipsData)(nil),            // 6: proto.PushTipsData
	(*PushTipsAllDataNotify)(nil),   // 7: proto.PushTipsAllDataNotify
	(*PushTipsChangeNotify)(nil),    // 8: proto.PushTipsChangeNotify
	(*PushTipsReadFinishReq)(nil),   // 9: proto.PushTipsReadFinishReq
	(*PushTipsReadFinishRsp)(nil),   // 10: proto.PushTipsReadFinishRsp
	(*GetPushTipsRewardReq)(nil),    // 11: proto.GetPushTipsRewardReq
	(*GetPushTipsRewardRsp)(nil),    // 12: proto.GetPushTipsRewardRsp
}
var file_cmd_cmd_watcher_proto_depIdxs = []int32{
	6, // 0: proto.PushTipsAllDataNotify.push_tips_list:type_name -> proto.PushTipsData
	6, // 1: proto.PushTipsChangeNotify.push_tips_list:type_name -> proto.PushTipsData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_cmd_watcher_proto_init() }
func file_cmd_cmd_watcher_proto_init() {
	if File_cmd_cmd_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherAllDataNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherChangeNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherEventNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherEventTypeNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatcherEventStageNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsData); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsAllDataNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsChangeNotify); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsReadFinishReq); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushTipsReadFinishRsp); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushTipsRewardReq); i {
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
		file_cmd_cmd_watcher_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPushTipsRewardRsp); i {
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
			RawDescriptor: file_cmd_cmd_watcher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_watcher_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_watcher_proto_depIdxs,
		EnumInfos:         file_cmd_cmd_watcher_proto_enumTypes,
		MessageInfos:      file_cmd_cmd_watcher_proto_msgTypes,
	}.Build()
	File_cmd_cmd_watcher_proto = out.File
	file_cmd_cmd_watcher_proto_rawDesc = nil
	file_cmd_cmd_watcher_proto_goTypes = nil
	file_cmd_cmd_watcher_proto_depIdxs = nil
}