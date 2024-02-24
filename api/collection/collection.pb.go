// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: collection/collection.proto

package collection

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

type CollectionRecordOrder int32

const (
	CollectionRecordOrder_TIMESTAMP  CollectionRecordOrder = 0
	CollectionRecordOrder_CREATED_AT CollectionRecordOrder = 1
)

// Enum value maps for CollectionRecordOrder.
var (
	CollectionRecordOrder_name = map[int32]string{
		0: "TIMESTAMP",
		1: "CREATED_AT",
	}
	CollectionRecordOrder_value = map[string]int32{
		"TIMESTAMP":  0,
		"CREATED_AT": 1,
	}
)

func (x CollectionRecordOrder) Enum() *CollectionRecordOrder {
	p := new(CollectionRecordOrder)
	*p = x
	return p
}

func (x CollectionRecordOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CollectionRecordOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_collection_collection_proto_enumTypes[0].Descriptor()
}

func (CollectionRecordOrder) Type() protoreflect.EnumType {
	return &file_collection_collection_proto_enumTypes[0]
}

func (x CollectionRecordOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CollectionRecordOrder.Descriptor instead.
func (CollectionRecordOrder) EnumDescriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{0}
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
	return file_collection_collection_proto_enumTypes[1].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_collection_collection_proto_enumTypes[1]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_collection_collection_proto_rawDescGZIP(), []int{1}
}

type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[0]
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
	return file_collection_collection_proto_rawDescGZIP(), []int{0}
}

func (x *GeoPoint) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeoPoint) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type CollectionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId    uint64    `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"deviceId"`    // @gotags: json:"deviceId"
	CreatedAt   int64     `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"createdAt"` // @gotags: json:"createdAt"
	Timestamp   int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GeoPoint    *GeoPoint `protobuf:"bytes,5,opt,name=geo_point,json=geoPoint,proto3" json:"geoPoint"` // @gotags: json:"geoPoint"
	Temperature float32   `protobuf:"fixed32,6,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *CollectionRecord) Reset() {
	*x = CollectionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRecord) ProtoMessage() {}

func (x *CollectionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[1]
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
	return file_collection_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionRecord) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CollectionRecord) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *CollectionRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
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
		mi := &file_collection_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[2]
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
	return file_collection_collection_proto_rawDescGZIP(), []int{2}
}

type ListCollectionRecordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId uint64                `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"deviceId" form:"deviceId"`                     // @gotags: form:"deviceId" json:"deviceId"
	Page     uint64                `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`                                             // @gotags: form:"page"
	Size     uint64                `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" form:"size"`                                             // @gotags: form:"size"
	Before   int64                 `protobuf:"varint,4,opt,name=before,proto3" json:"before,omitempty" form:"before"`                                         // @gotags: form:"before"
	After    int64                 `protobuf:"varint,5,opt,name=after,proto3" json:"after,omitempty" form:"after"`                                           // @gotags: form:"after"
	Order    CollectionRecordOrder `protobuf:"varint,6,opt,name=order,proto3,enum=api.collection.CollectionRecordOrder" json:"order,omitempty" form:"order"` // @gotags: form:"order"
	Sort     Sort                  `protobuf:"varint,7,opt,name=sort,proto3,enum=api.collection.Sort" json:"sort,omitempty" form:"sort"`                    // @gotags: form:"sort"
}

func (x *ListCollectionRecordReq) Reset() {
	*x = ListCollectionRecordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRecordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRecordReq) ProtoMessage() {}

func (x *ListCollectionRecordReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[3]
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
	return file_collection_collection_proto_rawDescGZIP(), []int{3}
}

func (x *ListCollectionRecordReq) GetDeviceId() uint64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *ListCollectionRecordReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCollectionRecordReq) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCollectionRecordReq) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *ListCollectionRecordReq) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ListCollectionRecordReq) GetOrder() CollectionRecordOrder {
	if x != nil {
		return x.Order
	}
	return CollectionRecordOrder_TIMESTAMP
}

func (x *ListCollectionRecordReq) GetSort() Sort {
	if x != nil {
		return x.Sort
	}
	return Sort_DESC
}

type ListCollectionRecordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*CollectionRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Total   int64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListCollectionRecordResp) Reset() {
	*x = ListCollectionRecordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRecordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRecordResp) ProtoMessage() {}

func (x *ListCollectionRecordResp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_collection_proto_msgTypes[4]
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
	return file_collection_collection_proto_rawDescGZIP(), []int{4}
}

func (x *ListCollectionRecordResp) GetRecords() []*CollectionRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *ListCollectionRecordResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_collection_collection_proto protoreflect.FileDescriptor

var file_collection_collection_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x47,
	0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0xd5, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x09, 0x67, 0x65,
	0x6f, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xf3, 0x01, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x22, 0x6c, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x2a, 0x36, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x01, 0x2a, 0x19, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53,
	0x43, 0x10, 0x01, 0x32, 0x86, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x70, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a,
	0x22, 0x12, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x44, 0x5a, 0x42,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53,
	0x75, 0x72, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_collection_proto_rawDescOnce sync.Once
	file_collection_collection_proto_rawDescData = file_collection_collection_proto_rawDesc
)

func file_collection_collection_proto_rawDescGZIP() []byte {
	file_collection_collection_proto_rawDescOnce.Do(func() {
		file_collection_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_collection_proto_rawDescData)
	})
	return file_collection_collection_proto_rawDescData
}

var file_collection_collection_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_collection_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_collection_collection_proto_goTypes = []interface{}{
	(CollectionRecordOrder)(0),       // 0: api.collection.CollectionRecordOrder
	(Sort)(0),                        // 1: api.collection.Sort
	(*GeoPoint)(nil),                 // 2: api.collection.GeoPoint
	(*CollectionRecord)(nil),         // 3: api.collection.CollectionRecord
	(*Empty)(nil),                    // 4: api.collection.Empty
	(*ListCollectionRecordReq)(nil),  // 5: api.collection.ListCollectionRecordReq
	(*ListCollectionRecordResp)(nil), // 6: api.collection.ListCollectionRecordResp
}
var file_collection_collection_proto_depIdxs = []int32{
	2, // 0: api.collection.CollectionRecord.geo_point:type_name -> api.collection.GeoPoint
	0, // 1: api.collection.ListCollectionRecordReq.order:type_name -> api.collection.CollectionRecordOrder
	1, // 2: api.collection.ListCollectionRecordReq.sort:type_name -> api.collection.Sort
	3, // 3: api.collection.ListCollectionRecordResp.records:type_name -> api.collection.CollectionRecord
	3, // 4: api.collection.Collection.CreateCollectionRecord:input_type -> api.collection.CollectionRecord
	5, // 5: api.collection.Collection.ListCollectionRecord:input_type -> api.collection.ListCollectionRecordReq
	4, // 6: api.collection.Collection.CreateCollectionRecord:output_type -> api.collection.Empty
	6, // 7: api.collection.Collection.ListCollectionRecord:output_type -> api.collection.ListCollectionRecordResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_collection_collection_proto_init() }
func file_collection_collection_proto_init() {
	if File_collection_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_collection_collection_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_collection_proto_goTypes,
		DependencyIndexes: file_collection_collection_proto_depIdxs,
		EnumInfos:         file_collection_collection_proto_enumTypes,
		MessageInfos:      file_collection_collection_proto_msgTypes,
	}.Build()
	File_collection_collection_proto = out.File
	file_collection_collection_proto_rawDesc = nil
	file_collection_collection_proto_goTypes = nil
	file_collection_collection_proto_depIdxs = nil
}
