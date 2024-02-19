// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: collection-database/collection-database.proto

package collection_database

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

type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPoint.ProtoReflect.Descriptor instead.
func (*GeoPoint) Descriptor() ([]byte, []int) {
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{0}
}

func (x *GeoPoint) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeoPoint) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type CollectionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId    uint64    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Timestamp   int64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GeoPoint    *GeoPoint `protobuf:"bytes,3,opt,name=geo_point,json=geoPoint,proto3" json:"geo_point,omitempty"`
	Temperature float32   `protobuf:"fixed32,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *CollectionRecord) Reset() {
	*x = CollectionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRecord) ProtoMessage() {}

func (x *CollectionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionRecord.ProtoReflect.Descriptor instead.
func (*CollectionRecord) Descriptor() ([]byte, []int) {
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionRecord) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *CollectionRecord) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CollectionRecord) GetGeoPoint() *GeoPoint {
	if x != nil {
		return x.GeoPoint
	}
	return nil
}

func (x *CollectionRecord) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[2]
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
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{2}
}

type ListCollectionRecordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page           uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize       uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	DeviceId       uint64 `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	StartTimestamp int64  `protobuf:"varint,4,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   int64  `protobuf:"varint,5,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
}

func (x *ListCollectionRecordReq) Reset() {
	*x = ListCollectionRecordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRecordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRecordReq) ProtoMessage() {}

func (x *ListCollectionRecordReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionRecordReq.ProtoReflect.Descriptor instead.
func (*ListCollectionRecordReq) Descriptor() ([]byte, []int) {
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{3}
}

func (x *ListCollectionRecordReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCollectionRecordReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCollectionRecordReq) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *ListCollectionRecordReq) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *ListCollectionRecordReq) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

type ListCollectionRecordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionInfos []*ListCollectionRecordResp_CollectionRecord `protobuf:"bytes,1,rep,name=collection_infos,json=collectionInfos,proto3" json:"collection_infos,omitempty"`
}

func (x *ListCollectionRecordResp) Reset() {
	*x = ListCollectionRecordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRecordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRecordResp) ProtoMessage() {}

func (x *ListCollectionRecordResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionRecordResp.ProtoReflect.Descriptor instead.
func (*ListCollectionRecordResp) Descriptor() ([]byte, []int) {
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{4}
}

func (x *ListCollectionRecordResp) GetCollectionInfos() []*ListCollectionRecordResp_CollectionRecord {
	if x != nil {
		return x.CollectionInfos
	}
	return nil
}

type ListCollectionRecordResp_CollectionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt int64     `protobuf:"varint,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeviceId  uint64    `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GeoPoint  *GeoPoint `protobuf:"bytes,4,opt,name=geo_point,json=geoPoint,proto3" json:"geo_point,omitempty"`
}

func (x *ListCollectionRecordResp_CollectionRecord) Reset() {
	*x = ListCollectionRecordResp_CollectionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_database_collection_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRecordResp_CollectionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRecordResp_CollectionRecord) ProtoMessage() {}

func (x *ListCollectionRecordResp_CollectionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_collection_database_collection_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionRecordResp_CollectionRecord.ProtoReflect.Descriptor instead.
func (*ListCollectionRecordResp_CollectionRecord) Descriptor() ([]byte, []int) {
	return file_collection_database_collection_database_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ListCollectionRecordResp_CollectionRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListCollectionRecordResp_CollectionRecord) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *ListCollectionRecordResp_CollectionRecord) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ListCollectionRecordResp_CollectionRecord) GetGeoPoint() *GeoPoint {
	if x != nil {
		return x.GeoPoint
	}
	return nil
}

var File_collection_database_collection_database_proto protoreflect.FileDescriptor

var file_collection_database_collection_database_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0xaf, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a, 0x09, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x65,
	0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb8, 0x02, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x6d, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0xac, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a, 0x09, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x32, 0xa4, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x7b, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x90, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x56, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53, 0x75,
	0x72, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3b,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_database_collection_database_proto_rawDescOnce sync.Once
	file_collection_database_collection_database_proto_rawDescData = file_collection_database_collection_database_proto_rawDesc
)

func file_collection_database_collection_database_proto_rawDescGZIP() []byte {
	file_collection_database_collection_database_proto_rawDescOnce.Do(func() {
		file_collection_database_collection_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_database_collection_database_proto_rawDescData)
	})
	return file_collection_database_collection_database_proto_rawDescData
}

var file_collection_database_collection_database_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_collection_database_collection_database_proto_goTypes = []interface{}{
	(*GeoPoint)(nil),                                  // 0: api.collection_database.GeoPoint
	(*CollectionRecord)(nil),                          // 1: api.collection_database.CollectionRecord
	(*Empty)(nil),                                     // 2: api.collection_database.Empty
	(*ListCollectionRecordReq)(nil),                   // 3: api.collection_database.ListCollectionRecordReq
	(*ListCollectionRecordResp)(nil),                  // 4: api.collection_database.ListCollectionRecordResp
	(*ListCollectionRecordResp_CollectionRecord)(nil), // 5: api.collection_database.ListCollectionRecordResp.CollectionRecord
}
var file_collection_database_collection_database_proto_depIdxs = []int32{
	0, // 0: api.collection_database.CollectionRecord.geo_point:type_name -> api.collection_database.GeoPoint
	5, // 1: api.collection_database.ListCollectionRecordResp.collection_infos:type_name -> api.collection_database.ListCollectionRecordResp.CollectionRecord
	0, // 2: api.collection_database.ListCollectionRecordResp.CollectionRecord.geo_point:type_name -> api.collection_database.GeoPoint
	1, // 3: api.collection_database.CollectionDatabase.CreateCollectionRecord:input_type -> api.collection_database.CollectionRecord
	3, // 4: api.collection_database.CollectionDatabase.ListCollectionRecord:input_type -> api.collection_database.ListCollectionRecordReq
	2, // 5: api.collection_database.CollectionDatabase.CreateCollectionRecord:output_type -> api.collection_database.Empty
	4, // 6: api.collection_database.CollectionDatabase.ListCollectionRecord:output_type -> api.collection_database.ListCollectionRecordResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_collection_database_collection_database_proto_init() }
func file_collection_database_collection_database_proto_init() {
	if File_collection_database_collection_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_database_collection_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoPoint); i {
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
		file_collection_database_collection_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionRecord); i {
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
		file_collection_database_collection_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_database_collection_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionRecordReq); i {
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
		file_collection_database_collection_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionRecordResp); i {
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
		file_collection_database_collection_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionRecordResp_CollectionRecord); i {
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
			RawDescriptor: file_collection_database_collection_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_database_collection_database_proto_goTypes,
		DependencyIndexes: file_collection_database_collection_database_proto_depIdxs,
		MessageInfos:      file_collection_database_collection_database_proto_msgTypes,
	}.Build()
	File_collection_database_collection_database_proto = out.File
	file_collection_database_collection_database_proto_rawDesc = nil
	file_collection_database_collection_database_proto_goTypes = nil
	file_collection_database_collection_database_proto_depIdxs = nil
}
