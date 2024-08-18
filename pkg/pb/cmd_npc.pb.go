// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_npc.proto

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

type NpcTalkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId    uint32 `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	NpcEntityId uint32 `protobuf:"varint,9,opt,name=npc_entity_id,json=npcEntityId,proto3" json:"npc_entity_id,omitempty"`
	TalkId      uint32 `protobuf:"varint,7,opt,name=talk_id,json=talkId,proto3" json:"talk_id,omitempty"`
}

func (x *NpcTalkReq) Reset() {
	*x = NpcTalkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NpcTalkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NpcTalkReq) ProtoMessage() {}

func (x *NpcTalkReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NpcTalkReq.ProtoReflect.Descriptor instead.
func (*NpcTalkReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{0}
}

func (x *NpcTalkReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *NpcTalkReq) GetNpcEntityId() uint32 {
	if x != nil {
		return x.NpcEntityId
	}
	return 0
}

func (x *NpcTalkReq) GetTalkId() uint32 {
	if x != nil {
		return x.TalkId
	}
	return 0
}

type NpcTalkRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurTalkId   uint32 `protobuf:"varint,9,opt,name=cur_talk_id,json=curTalkId,proto3" json:"cur_talk_id,omitempty"`
	NpcEntityId uint32 `protobuf:"varint,6,opt,name=npc_entity_id,json=npcEntityId,proto3" json:"npc_entity_id,omitempty"`
	Retcode     int32  `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EntityId    uint32 `protobuf:"varint,13,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *NpcTalkRsp) Reset() {
	*x = NpcTalkRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NpcTalkRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NpcTalkRsp) ProtoMessage() {}

func (x *NpcTalkRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NpcTalkRsp.ProtoReflect.Descriptor instead.
func (*NpcTalkRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{1}
}

func (x *NpcTalkRsp) GetCurTalkId() uint32 {
	if x != nil {
		return x.CurTalkId
	}
	return 0
}

func (x *NpcTalkRsp) GetNpcEntityId() uint32 {
	if x != nil {
		return x.NpcEntityId
	}
	return 0
}

func (x *NpcTalkRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *NpcTalkRsp) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

type GetSceneNpcPositionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NpcIdList []uint32 `protobuf:"varint,6,rep,packed,name=npc_id_list,json=npcIdList,proto3" json:"npc_id_list,omitempty"`
	SceneId   uint32   `protobuf:"varint,8,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *GetSceneNpcPositionReq) Reset() {
	*x = GetSceneNpcPositionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneNpcPositionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneNpcPositionReq) ProtoMessage() {}

func (x *GetSceneNpcPositionReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneNpcPositionReq.ProtoReflect.Descriptor instead.
func (*GetSceneNpcPositionReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{2}
}

func (x *GetSceneNpcPositionReq) GetNpcIdList() []uint32 {
	if x != nil {
		return x.NpcIdList
	}
	return nil
}

func (x *GetSceneNpcPositionReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

type GetSceneNpcPositionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32              `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NpcInfoList []*NpcPositionInfo `protobuf:"bytes,14,rep,name=npc_info_list,json=npcInfoList,proto3" json:"npc_info_list,omitempty"`
	SceneId     uint32             `protobuf:"varint,4,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *GetSceneNpcPositionRsp) Reset() {
	*x = GetSceneNpcPositionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneNpcPositionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneNpcPositionRsp) ProtoMessage() {}

func (x *GetSceneNpcPositionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneNpcPositionRsp.ProtoReflect.Descriptor instead.
func (*GetSceneNpcPositionRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{3}
}

func (x *GetSceneNpcPositionRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetSceneNpcPositionRsp) GetNpcInfoList() []*NpcPositionInfo {
	if x != nil {
		return x.NpcInfoList
	}
	return nil
}

func (x *GetSceneNpcPositionRsp) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

type MetNpcIdListNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NpcFirstMetIdList []uint32 `protobuf:"varint,9,rep,packed,name=npc_first_met_id_list,json=npcFirstMetIdList,proto3" json:"npc_first_met_id_list,omitempty"`
}

func (x *MetNpcIdListNotify) Reset() {
	*x = MetNpcIdListNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetNpcIdListNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetNpcIdListNotify) ProtoMessage() {}

func (x *MetNpcIdListNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetNpcIdListNotify.ProtoReflect.Descriptor instead.
func (*MetNpcIdListNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{4}
}

func (x *MetNpcIdListNotify) GetNpcFirstMetIdList() []uint32 {
	if x != nil {
		return x.NpcFirstMetIdList
	}
	return nil
}

type MeetNpcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NpcId uint32 `protobuf:"varint,4,opt,name=npc_id,json=npcId,proto3" json:"npc_id,omitempty"`
}

func (x *MeetNpcReq) Reset() {
	*x = MeetNpcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeetNpcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeetNpcReq) ProtoMessage() {}

func (x *MeetNpcReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeetNpcReq.ProtoReflect.Descriptor instead.
func (*MeetNpcReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{5}
}

func (x *MeetNpcReq) GetNpcId() uint32 {
	if x != nil {
		return x.NpcId
	}
	return 0
}

type MeetNpcRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32  `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NpcFirstMetId uint32 `protobuf:"varint,8,opt,name=npc_first_met_id,json=npcFirstMetId,proto3" json:"npc_first_met_id,omitempty"`
}

func (x *MeetNpcRsp) Reset() {
	*x = MeetNpcRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeetNpcRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeetNpcRsp) ProtoMessage() {}

func (x *MeetNpcRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeetNpcRsp.ProtoReflect.Descriptor instead.
func (*MeetNpcRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{6}
}

func (x *MeetNpcRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MeetNpcRsp) GetNpcFirstMetId() uint32 {
	if x != nil {
		return x.NpcFirstMetId
	}
	return 0
}

type FinishedTalkIdListNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedTalkIdList []uint32 `protobuf:"varint,1,rep,packed,name=finished_talk_id_list,json=finishedTalkIdList,proto3" json:"finished_talk_id_list,omitempty"`
}

func (x *FinishedTalkIdListNotify) Reset() {
	*x = FinishedTalkIdListNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_npc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishedTalkIdListNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishedTalkIdListNotify) ProtoMessage() {}

func (x *FinishedTalkIdListNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_npc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishedTalkIdListNotify.ProtoReflect.Descriptor instead.
func (*FinishedTalkIdListNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_npc_proto_rawDescGZIP(), []int{7}
}

func (x *FinishedTalkIdListNotify) GetFinishedTalkIdList() []uint32 {
	if x != nil {
		return x.FinishedTalkIdList
	}
	return nil
}

var File_cmd_cmd_npc_proto protoreflect.FileDescriptor

var file_cmd_cmd_npc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x6e, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0a, 0x4e, 0x70, 0x63, 0x54,
	0x61, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x70, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x70, 0x63, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x6c, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x6c, 0x6b, 0x49, 0x64,
	0x22, 0x87, 0x01, 0x0a, 0x0a, 0x4e, 0x70, 0x63, 0x54, 0x61, 0x6c, 0x6b, 0x52, 0x73, 0x70, 0x12,
	0x1e, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x5f, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x75, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x6e, 0x70, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x70, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x70, 0x63, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x22,
	0x89, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x6e, 0x70, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6e, 0x70, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x4d,
	0x65, 0x74, 0x4e, 0x70, 0x63, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x30, 0x0a, 0x15, 0x6e, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x11, 0x6e, 0x70, 0x63, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x0a, 0x4d, 0x65, 0x65, 0x74, 0x4e, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6e, 0x70, 0x63, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x0a, 0x4d, 0x65, 0x65, 0x74,
	0x4e, 0x70, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x27, 0x0a, 0x10, 0x6e, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x70, 0x63, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x18, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x61, 0x6c, 0x6b, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x61,
	0x6c, 0x6b, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_npc_proto_rawDescOnce sync.Once
	file_cmd_cmd_npc_proto_rawDescData = file_cmd_cmd_npc_proto_rawDesc
)

func file_cmd_cmd_npc_proto_rawDescGZIP() []byte {
	file_cmd_cmd_npc_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_npc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_npc_proto_rawDescData)
	})
	return file_cmd_cmd_npc_proto_rawDescData
}

var file_cmd_cmd_npc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_cmd_npc_proto_goTypes = []interface{}{
	(*NpcTalkReq)(nil),               // 0: proto.NpcTalkReq
	(*NpcTalkRsp)(nil),               // 1: proto.NpcTalkRsp
	(*GetSceneNpcPositionReq)(nil),   // 2: proto.GetSceneNpcPositionReq
	(*GetSceneNpcPositionRsp)(nil),   // 3: proto.GetSceneNpcPositionRsp
	(*MetNpcIdListNotify)(nil),       // 4: proto.MetNpcIdListNotify
	(*MeetNpcReq)(nil),               // 5: proto.MeetNpcReq
	(*MeetNpcRsp)(nil),               // 6: proto.MeetNpcRsp
	(*FinishedTalkIdListNotify)(nil), // 7: proto.FinishedTalkIdListNotify
	(*NpcPositionInfo)(nil),          // 8: proto.NpcPositionInfo
}
var file_cmd_cmd_npc_proto_depIdxs = []int32{
	8, // 0: proto.GetSceneNpcPositionRsp.npc_info_list:type_name -> proto.NpcPositionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_cmd_npc_proto_init() }
func file_cmd_cmd_npc_proto_init() {
	if File_cmd_cmd_npc_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_npc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NpcTalkReq); i {
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
		file_cmd_cmd_npc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NpcTalkRsp); i {
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
		file_cmd_cmd_npc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneNpcPositionReq); i {
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
		file_cmd_cmd_npc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneNpcPositionRsp); i {
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
		file_cmd_cmd_npc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetNpcIdListNotify); i {
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
		file_cmd_cmd_npc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeetNpcReq); i {
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
		file_cmd_cmd_npc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeetNpcRsp); i {
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
		file_cmd_cmd_npc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishedTalkIdListNotify); i {
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
			RawDescriptor: file_cmd_cmd_npc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_npc_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_npc_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_npc_proto_msgTypes,
	}.Build()
	File_cmd_cmd_npc_proto = out.File
	file_cmd_cmd_npc_proto_rawDesc = nil
	file_cmd_cmd_npc_proto_goTypes = nil
	file_cmd_cmd_npc_proto_depIdxs = nil
}