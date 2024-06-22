// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: log/log.proto

package log

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DeviceLogOrder int32

const (
	DeviceLogOrder_TIMESTAMP  DeviceLogOrder = 0
	DeviceLogOrder_CREATED_AT DeviceLogOrder = 1
)

// Enum value maps for DeviceLogOrder.
var (
	DeviceLogOrder_name = map[int32]string{
		0: "TIMESTAMP",
		1: "CREATED_AT",
	}
	DeviceLogOrder_value = map[string]int32{
		"TIMESTAMP":  0,
		"CREATED_AT": 1,
	}
)

func (x DeviceLogOrder) Enum() *DeviceLogOrder {
	p := new(DeviceLogOrder)
	*p = x
	return p
}

func (x DeviceLogOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceLogOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_log_log_proto_enumTypes[0].Descriptor()
}

func (DeviceLogOrder) Type() protoreflect.EnumType {
	return &file_log_log_proto_enumTypes[0]
}

func (x DeviceLogOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceLogOrder.Descriptor instead.
func (DeviceLogOrder) EnumDescriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{0}
}

type Sort int32

const (
	Sort_DESC Sort = 0
	Sort_ASC  Sort = 1
)

// Enum value maps for Sort.
var (
	Sort_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	Sort_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x Sort) Enum() *Sort {
	p := new(Sort)
	*p = x
	return p
}

func (x Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_log_log_proto_enumTypes[1].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_log_log_proto_enumTypes[1]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{0}
}

type DeviceLogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId  uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"deviceId"` // @gotags: json:"deviceId"
	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Level     uint32 `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DeviceLogData) Reset() {
	*x = DeviceLogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceLogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLogData) ProtoMessage() {}

func (x *DeviceLogData) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLogData.ProtoReflect.Descriptor instead.
func (*DeviceLogData) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceLogData) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *DeviceLogData) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *DeviceLogData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeviceLogData) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DeviceLogData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type DeviceLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *DeviceLogData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeviceLog) Reset() {
	*x = DeviceLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceLog) ProtoMessage() {}

func (x *DeviceLog) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceLog.ProtoReflect.Descriptor instead.
func (*DeviceLog) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceLog) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceLog) GetData() *DeviceLogData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListDeviceLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId    uint64         `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"deviceId" form:"deviceId"`         // @gotags: form:"deviceId" json:"deviceId"
	Page        int64          `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`                                 // @gotags: form:"page"
	Size        int64          `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" form:"size"`                                 // @gotags: form:"size"
	Before      int64          `protobuf:"varint,4,opt,name=before,proto3" json:"before,omitempty" form:"before"`                             // @gotags: form:"before"
	After       int64          `protobuf:"varint,5,opt,name=after,proto3" json:"after,omitempty" form:"after"`                               // @gotags: form:"after"
	LevelFilter string         `protobuf:"bytes,6,opt,name=level_filter,json=levelFilter,proto3" json:"levelFilter" form:"levelFilter"` // @gotags: form:"levelFilter" json:"levelFilter"
	Order       DeviceLogOrder `protobuf:"varint,7,opt,name=order,proto3,enum=api.log.DeviceLogOrder" json:"order,omitempty" form:"order"`   // @gotags: form:"order"
	Sort        Sort           `protobuf:"varint,8,opt,name=sort,proto3,enum=api.log.Sort" json:"sort,omitempty" form:"sort"`               // @gotags: form:"sort"
}

func (x *ListDeviceLogReq) Reset() {
	*x = ListDeviceLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceLogReq) ProtoMessage() {}

func (x *ListDeviceLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceLogReq.ProtoReflect.Descriptor instead.
func (*ListDeviceLogReq) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{3}
}

func (x *ListDeviceLogReq) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *ListDeviceLogReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListDeviceLogReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListDeviceLogReq) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *ListDeviceLogReq) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ListDeviceLogReq) GetLevelFilter() string {
	if x != nil {
		return x.LevelFilter
	}
	return ""
}

func (x *ListDeviceLogReq) GetOrder() DeviceLogOrder {
	if x != nil {
		return x.Order
	}
	return DeviceLogOrder_TIMESTAMP
}

func (x *ListDeviceLogReq) GetSort() Sort {
	if x != nil {
		return x.Sort
	}
	return Sort_DESC
}

type ListDeviceLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs  []*DeviceLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	Total int64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListDeviceLogResp) Reset() {
	*x = ListDeviceLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceLogResp) ProtoMessage() {}

func (x *ListDeviceLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceLogResp.ProtoReflect.Descriptor instead.
func (*ListDeviceLogResp) Descriptor() ([]byte, []int) {
	return file_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *ListDeviceLogResp) GetLogs() []*DeviceLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *ListDeviceLogResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_log_log_proto protoreflect.FileDescriptor

var file_log_log_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x90, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x47, 0x0a, 0x09, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xfa, 0x01, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x51, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x2f, 0x0a, 0x0e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x01, 0x2a, 0x19, 0x0a, 0x04,
	0x53, 0x6f, 0x72, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x32, 0xb1, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x4d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01,
	0x2a, 0x22, 0x0b, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53, 0x75,
	0x72, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x3b,
	0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_log_proto_rawDescOnce sync.Once
	file_log_log_proto_rawDescData = file_log_log_proto_rawDesc
)

func file_log_log_proto_rawDescGZIP() []byte {
	file_log_log_proto_rawDescOnce.Do(func() {
		file_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_log_proto_rawDescData)
	})
	return file_log_log_proto_rawDescData
}

var file_log_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_log_log_proto_goTypes = []any{
	(DeviceLogOrder)(0),       // 0: api.log.DeviceLogOrder
	(Sort)(0),                 // 1: api.log.Sort
	(*Empty)(nil),             // 2: api.log.Empty
	(*DeviceLogData)(nil),     // 3: api.log.DeviceLogData
	(*DeviceLog)(nil),         // 4: api.log.DeviceLog
	(*ListDeviceLogReq)(nil),  // 5: api.log.ListDeviceLogReq
	(*ListDeviceLogResp)(nil), // 6: api.log.ListDeviceLogResp
}
var file_log_log_proto_depIdxs = []int32{
	3, // 0: api.log.DeviceLog.data:type_name -> api.log.DeviceLogData
	0, // 1: api.log.ListDeviceLogReq.order:type_name -> api.log.DeviceLogOrder
	1, // 2: api.log.ListDeviceLogReq.sort:type_name -> api.log.Sort
	4, // 3: api.log.ListDeviceLogResp.logs:type_name -> api.log.DeviceLog
	4, // 4: api.log.Log.CreateDeviceLog:input_type -> api.log.DeviceLog
	5, // 5: api.log.Log.ListDeviceLog:input_type -> api.log.ListDeviceLogReq
	2, // 6: api.log.Log.CreateDeviceLog:output_type -> api.log.Empty
	6, // 7: api.log.Log.ListDeviceLog:output_type -> api.log.ListDeviceLogResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_log_log_proto_init() }
func file_log_log_proto_init() {
	if File_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_log_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceLogData); i {
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
		file_log_log_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceLog); i {
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
		file_log_log_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListDeviceLogReq); i {
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
		file_log_log_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListDeviceLogResp); i {
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
			RawDescriptor: file_log_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_log_proto_goTypes,
		DependencyIndexes: file_log_log_proto_depIdxs,
		EnumInfos:         file_log_log_proto_enumTypes,
		MessageInfos:      file_log_log_proto_msgTypes,
	}.Build()
	File_log_log_proto = out.File
	file_log_log_proto_rawDesc = nil
	file_log_log_proto_goTypes = nil
	file_log_log_proto_depIdxs = nil
}
