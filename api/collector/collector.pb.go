// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: collector/collector.proto

package collector

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

type ReportImmediatelyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReportImmediatelyReq) Reset() {
	*x = ReportImmediatelyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportImmediatelyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportImmediatelyReq) ProtoMessage() {}

func (x *ReportImmediatelyReq) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportImmediatelyReq.ProtoReflect.Descriptor instead.
func (*ReportImmediatelyReq) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{0}
}

func (x *ReportImmediatelyReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReportImmediatelyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportImmediatelyResp) Reset() {
	*x = ReportImmediatelyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportImmediatelyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportImmediatelyResp) ProtoMessage() {}

func (x *ReportImmediatelyResp) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportImmediatelyResp.ProtoReflect.Descriptor instead.
func (*ReportImmediatelyResp) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{1}
}

type SetReportIntervalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Interval uint64 `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *SetReportIntervalReq) Reset() {
	*x = SetReportIntervalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReportIntervalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReportIntervalReq) ProtoMessage() {}

func (x *SetReportIntervalReq) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReportIntervalReq.ProtoReflect.Descriptor instead.
func (*SetReportIntervalReq) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{2}
}

func (x *SetReportIntervalReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetReportIntervalReq) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type SetReportIntervalResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetReportIntervalResp) Reset() {
	*x = SetReportIntervalResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReportIntervalResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReportIntervalResp) ProtoMessage() {}

func (x *SetReportIntervalResp) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReportIntervalResp.ProtoReflect.Descriptor instead.
func (*SetReportIntervalResp) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{3}
}

type GetDeviceStreamLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" binding:"required"`                                     // @gotags: form:"id" binding:"required"
	LevelFilter string `protobuf:"bytes,2,opt,name=level_filter,json=levelFilter,proto3" json:"levelFilter" form:"levelFilter"` // @gotags: form:"levelFilter" json:"levelFilter"
}

func (x *GetDeviceStreamLogReq) Reset() {
	*x = GetDeviceStreamLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStreamLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStreamLogReq) ProtoMessage() {}

func (x *GetDeviceStreamLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStreamLogReq.ProtoReflect.Descriptor instead.
func (*GetDeviceStreamLogReq) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceStreamLogReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDeviceStreamLogReq) GetLevelFilter() string {
	if x != nil {
		return x.LevelFilter
	}
	return ""
}

type GetDeviceStreamLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level     uint32 `protobuf:"varint,1,opt,name=level,proto3" json:"level"`         // @gotags: json:"level"
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`      // @gotags: json:"message"
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp"` // @gotags: json:"timestamp"
}

func (x *GetDeviceStreamLogResp) Reset() {
	*x = GetDeviceStreamLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStreamLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStreamLogResp) ProtoMessage() {}

func (x *GetDeviceStreamLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStreamLogResp.ProtoReflect.Descriptor instead.
func (*GetDeviceStreamLogResp) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeviceStreamLogResp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetDeviceStreamLogResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetDeviceStreamLogResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_collector_collector_proto protoreflect.FileDescriptor

var file_collector_collector_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x42, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x17, 0x0a,
	0x15, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x66, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xb5, 0x03, 0x0a, 0x09, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x65, 0x6c, 0x79, 0x12, 0x89, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22,
	0x1e, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x90, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x6f, 0x67,
	0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53, 0x75, 0x72, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d,
	0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collector_collector_proto_rawDescOnce sync.Once
	file_collector_collector_proto_rawDescData = file_collector_collector_proto_rawDesc
)

func file_collector_collector_proto_rawDescGZIP() []byte {
	file_collector_collector_proto_rawDescOnce.Do(func() {
		file_collector_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_collector_collector_proto_rawDescData)
	})
	return file_collector_collector_proto_rawDescData
}

var file_collector_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_collector_collector_proto_goTypes = []interface{}{
	(*ReportImmediatelyReq)(nil),   // 0: api.collector.ReportImmediatelyReq
	(*ReportImmediatelyResp)(nil),  // 1: api.collector.ReportImmediatelyResp
	(*SetReportIntervalReq)(nil),   // 2: api.collector.SetReportIntervalReq
	(*SetReportIntervalResp)(nil),  // 3: api.collector.SetReportIntervalResp
	(*GetDeviceStreamLogReq)(nil),  // 4: api.collector.GetDeviceStreamLogReq
	(*GetDeviceStreamLogResp)(nil), // 5: api.collector.GetDeviceStreamLogResp
}
var file_collector_collector_proto_depIdxs = []int32{
	0, // 0: api.collector.Collector.ReportImmediately:input_type -> api.collector.ReportImmediatelyReq
	2, // 1: api.collector.Collector.SetReportInterval:input_type -> api.collector.SetReportIntervalReq
	4, // 2: api.collector.Collector.GetDeviceStreamLog:input_type -> api.collector.GetDeviceStreamLogReq
	1, // 3: api.collector.Collector.ReportImmediately:output_type -> api.collector.ReportImmediatelyResp
	3, // 4: api.collector.Collector.SetReportInterval:output_type -> api.collector.SetReportIntervalResp
	5, // 5: api.collector.Collector.GetDeviceStreamLog:output_type -> api.collector.GetDeviceStreamLogResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_collector_collector_proto_init() }
func file_collector_collector_proto_init() {
	if File_collector_collector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collector_collector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportImmediatelyReq); i {
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
		file_collector_collector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportImmediatelyResp); i {
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
		file_collector_collector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReportIntervalReq); i {
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
		file_collector_collector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetReportIntervalResp); i {
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
		file_collector_collector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStreamLogReq); i {
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
		file_collector_collector_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStreamLogResp); i {
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
			RawDescriptor: file_collector_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collector_collector_proto_goTypes,
		DependencyIndexes: file_collector_collector_proto_depIdxs,
		MessageInfos:      file_collector_collector_proto_msgTypes,
	}.Build()
	File_collector_collector_proto = out.File
	file_collector_collector_proto_rawDesc = nil
	file_collector_collector_proto_goTypes = nil
	file_collector_collector_proto_depIdxs = nil
}
