// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: cmd/cmd_routine.proto

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

type PlayerRoutineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType uint32 `protobuf:"varint,8,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
	FinishedNum uint32 `protobuf:"varint,15,opt,name=finished_num,json=finishedNum,proto3" json:"finished_num,omitempty"`
}

func (x *PlayerRoutineInfo) Reset() {
	*x = PlayerRoutineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRoutineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRoutineInfo) ProtoMessage() {}

func (x *PlayerRoutineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRoutineInfo.ProtoReflect.Descriptor instead.
func (*PlayerRoutineInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRoutineInfo) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

func (x *PlayerRoutineInfo) GetFinishedNum() uint32 {
	if x != nil {
		return x.FinishedNum
	}
	return 0
}

type PlayerRoutineDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineInfoList []*PlayerRoutineInfo `protobuf:"bytes,11,rep,name=routine_info_list,json=routineInfoList,proto3" json:"routine_info_list,omitempty"`
}

func (x *PlayerRoutineDataNotify) Reset() {
	*x = PlayerRoutineDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRoutineDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRoutineDataNotify) ProtoMessage() {}

func (x *PlayerRoutineDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRoutineDataNotify.ProtoReflect.Descriptor instead.
func (*PlayerRoutineDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerRoutineDataNotify) GetRoutineInfoList() []*PlayerRoutineInfo {
	if x != nil {
		return x.RoutineInfoList
	}
	return nil
}

type WorldRoutineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress       uint32 `protobuf:"varint,4,opt,name=progress,proto3" json:"progress,omitempty"`
	IsFinished     bool   `protobuf:"varint,14,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	FinishProgress uint32 `protobuf:"varint,3,opt,name=finish_progress,json=finishProgress,proto3" json:"finish_progress,omitempty"`
	RoutineId      uint32 `protobuf:"varint,11,opt,name=routine_id,json=routineId,proto3" json:"routine_id,omitempty"`
}

func (x *WorldRoutineInfo) Reset() {
	*x = WorldRoutineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineInfo) ProtoMessage() {}

func (x *WorldRoutineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineInfo.ProtoReflect.Descriptor instead.
func (*WorldRoutineInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{2}
}

func (x *WorldRoutineInfo) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *WorldRoutineInfo) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *WorldRoutineInfo) GetFinishProgress() uint32 {
	if x != nil {
		return x.FinishProgress
	}
	return 0
}

func (x *WorldRoutineInfo) GetRoutineId() uint32 {
	if x != nil {
		return x.RoutineId
	}
	return 0
}

type WorldRoutineTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType          uint32              `protobuf:"varint,13,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
	NextRefreshTime      uint32              `protobuf:"varint,12,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	WorldRoutineInfoList []*WorldRoutineInfo `protobuf:"bytes,3,rep,name=world_routine_info_list,json=worldRoutineInfoList,proto3" json:"world_routine_info_list,omitempty"`
}

func (x *WorldRoutineTypeInfo) Reset() {
	*x = WorldRoutineTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeInfo) ProtoMessage() {}

func (x *WorldRoutineTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeInfo.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{3}
}

func (x *WorldRoutineTypeInfo) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetWorldRoutineInfoList() []*WorldRoutineInfo {
	if x != nil {
		return x.WorldRoutineInfoList
	}
	return nil
}

type WorldAllRoutineTypeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldRoutineTypeList []*WorldRoutineTypeInfo `protobuf:"bytes,12,rep,name=world_routine_type_list,json=worldRoutineTypeList,proto3" json:"world_routine_type_list,omitempty"`
}

func (x *WorldAllRoutineTypeNotify) Reset() {
	*x = WorldAllRoutineTypeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldAllRoutineTypeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldAllRoutineTypeNotify) ProtoMessage() {}

func (x *WorldAllRoutineTypeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldAllRoutineTypeNotify.ProtoReflect.Descriptor instead.
func (*WorldAllRoutineTypeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{4}
}

func (x *WorldAllRoutineTypeNotify) GetWorldRoutineTypeList() []*WorldRoutineTypeInfo {
	if x != nil {
		return x.WorldRoutineTypeList
	}
	return nil
}

type WorldRoutineTypeRefreshNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldRoutineType *WorldRoutineTypeInfo `protobuf:"bytes,7,opt,name=world_routine_type,json=worldRoutineType,proto3" json:"world_routine_type,omitempty"`
}

func (x *WorldRoutineTypeRefreshNotify) Reset() {
	*x = WorldRoutineTypeRefreshNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeRefreshNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeRefreshNotify) ProtoMessage() {}

func (x *WorldRoutineTypeRefreshNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeRefreshNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeRefreshNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{5}
}

func (x *WorldRoutineTypeRefreshNotify) GetWorldRoutineType() *WorldRoutineTypeInfo {
	if x != nil {
		return x.WorldRoutineType
	}
	return nil
}

type WorldRoutineChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineInfo *WorldRoutineInfo `protobuf:"bytes,3,opt,name=routine_info,json=routineInfo,proto3" json:"routine_info,omitempty"`
	RoutineType uint32            `protobuf:"varint,11,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
}

func (x *WorldRoutineChangeNotify) Reset() {
	*x = WorldRoutineChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineChangeNotify) ProtoMessage() {}

func (x *WorldRoutineChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineChangeNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineChangeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{6}
}

func (x *WorldRoutineChangeNotify) GetRoutineInfo() *WorldRoutineInfo {
	if x != nil {
		return x.RoutineInfo
	}
	return nil
}

func (x *WorldRoutineChangeNotify) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

type WorldRoutineTypeCloseNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType uint32 `protobuf:"varint,7,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
}

func (x *WorldRoutineTypeCloseNotify) Reset() {
	*x = WorldRoutineTypeCloseNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_cmd_routine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeCloseNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeCloseNotify) ProtoMessage() {}

func (x *WorldRoutineTypeCloseNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_routine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeCloseNotify.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeCloseNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_routine_proto_rawDescGZIP(), []int{7}
}

func (x *WorldRoutineTypeCloseNotify) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

var File_cmd_cmd_routine_proto protoreflect.FileDescriptor

var file_cmd_cmd_routine_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0x5f, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x44, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x17, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x19, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x52, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x1d, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x49, 0x0a, 0x12,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x79, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x40, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_routine_proto_rawDescOnce sync.Once
	file_cmd_cmd_routine_proto_rawDescData = file_cmd_cmd_routine_proto_rawDesc
)

func file_cmd_cmd_routine_proto_rawDescGZIP() []byte {
	file_cmd_cmd_routine_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_routine_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_routine_proto_rawDescData)
	})
	return file_cmd_cmd_routine_proto_rawDescData
}

var file_cmd_cmd_routine_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_cmd_routine_proto_goTypes = []interface{}{
	(*PlayerRoutineInfo)(nil),             // 0: proto.PlayerRoutineInfo
	(*PlayerRoutineDataNotify)(nil),       // 1: proto.PlayerRoutineDataNotify
	(*WorldRoutineInfo)(nil),              // 2: proto.WorldRoutineInfo
	(*WorldRoutineTypeInfo)(nil),          // 3: proto.WorldRoutineTypeInfo
	(*WorldAllRoutineTypeNotify)(nil),     // 4: proto.WorldAllRoutineTypeNotify
	(*WorldRoutineTypeRefreshNotify)(nil), // 5: proto.WorldRoutineTypeRefreshNotify
	(*WorldRoutineChangeNotify)(nil),      // 6: proto.WorldRoutineChangeNotify
	(*WorldRoutineTypeCloseNotify)(nil),   // 7: proto.WorldRoutineTypeCloseNotify
}
var file_cmd_cmd_routine_proto_depIdxs = []int32{
	0, // 0: proto.PlayerRoutineDataNotify.routine_info_list:type_name -> proto.PlayerRoutineInfo
	2, // 1: proto.WorldRoutineTypeInfo.world_routine_info_list:type_name -> proto.WorldRoutineInfo
	3, // 2: proto.WorldAllRoutineTypeNotify.world_routine_type_list:type_name -> proto.WorldRoutineTypeInfo
	3, // 3: proto.WorldRoutineTypeRefreshNotify.world_routine_type:type_name -> proto.WorldRoutineTypeInfo
	2, // 4: proto.WorldRoutineChangeNotify.routine_info:type_name -> proto.WorldRoutineInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cmd_cmd_routine_proto_init() }
func file_cmd_cmd_routine_proto_init() {
	if File_cmd_cmd_routine_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_cmd_routine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRoutineInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRoutineDataNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeInfo); i {
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
		file_cmd_cmd_routine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldAllRoutineTypeNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeRefreshNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineChangeNotify); i {
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
		file_cmd_cmd_routine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeCloseNotify); i {
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
			RawDescriptor: file_cmd_cmd_routine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_routine_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_routine_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_routine_proto_msgTypes,
	}.Build()
	File_cmd_cmd_routine_proto = out.File
	file_cmd_cmd_routine_proto_rawDesc = nil
	file_cmd_cmd_routine_proto_goTypes = nil
	file_cmd_cmd_routine_proto_depIdxs = nil
}