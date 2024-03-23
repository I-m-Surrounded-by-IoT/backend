// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: waterquality/waterquality.proto

package waterquality

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

type Quality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp   int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Temperature float32 `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Ph          float32 `protobuf:"fixed32,3,opt,name=ph,proto3" json:"ph,omitempty"`
}

func (x *Quality) Reset() {
	*x = Quality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quality) ProtoMessage() {}

func (x *Quality) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quality.ProtoReflect.Descriptor instead.
func (*Quality) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{0}
}

func (x *Quality) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Quality) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Quality) GetPh() float32 {
	if x != nil {
		return x.Ph
	}
	return 0
}

type PredictReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qualities []*Quality `protobuf:"bytes,1,rep,name=qualities,proto3" json:"qualities,omitempty"`
	LookBack  int64      `protobuf:"varint,2,opt,name=look_back,json=lookBack,proto3" json:"look_back,omitempty"`
	Horizon   int64      `protobuf:"varint,3,opt,name=horizon,proto3" json:"horizon,omitempty"`
}

func (x *PredictReq) Reset() {
	*x = PredictReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictReq) ProtoMessage() {}

func (x *PredictReq) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictReq.ProtoReflect.Descriptor instead.
func (*PredictReq) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{1}
}

func (x *PredictReq) GetQualities() []*Quality {
	if x != nil {
		return x.Qualities
	}
	return nil
}

func (x *PredictReq) GetLookBack() int64 {
	if x != nil {
		return x.LookBack
	}
	return 0
}

func (x *PredictReq) GetHorizon() int64 {
	if x != nil {
		return x.Horizon
	}
	return 0
}

type PredictResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qualities []*Quality `protobuf:"bytes,1,rep,name=qualities,proto3" json:"qualities,omitempty"`
}

func (x *PredictResp) Reset() {
	*x = PredictResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResp) ProtoMessage() {}

func (x *PredictResp) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResp.ProtoReflect.Descriptor instead.
func (*PredictResp) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{2}
}

func (x *PredictResp) GetQualities() []*Quality {
	if x != nil {
		return x.Qualities
	}
	return nil
}

type GuessLevelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int64 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GuessLevelResp) Reset() {
	*x = GuessLevelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessLevelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessLevelResp) ProtoMessage() {}

func (x *GuessLevelResp) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessLevelResp.ProtoReflect.Descriptor instead.
func (*GuessLevelResp) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{3}
}

func (x *GuessLevelResp) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

type PredictAndGuessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qualities []*Quality `protobuf:"bytes,1,rep,name=qualities,proto3" json:"qualities,omitempty"`
	LookBack  int64      `protobuf:"varint,2,opt,name=look_back,json=lookBack,proto3" json:"look_back,omitempty"`
	Horizon   int64      `protobuf:"varint,3,opt,name=horizon,proto3" json:"horizon,omitempty"`
}

func (x *PredictAndGuessReq) Reset() {
	*x = PredictAndGuessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictAndGuessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictAndGuessReq) ProtoMessage() {}

func (x *PredictAndGuessReq) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictAndGuessReq.ProtoReflect.Descriptor instead.
func (*PredictAndGuessReq) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{4}
}

func (x *PredictAndGuessReq) GetQualities() []*Quality {
	if x != nil {
		return x.Qualities
	}
	return nil
}

func (x *PredictAndGuessReq) GetLookBack() int64 {
	if x != nil {
		return x.LookBack
	}
	return 0
}

func (x *PredictAndGuessReq) GetHorizon() int64 {
	if x != nil {
		return x.Horizon
	}
	return 0
}

type PredictAndGuessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qualities []*Quality `protobuf:"bytes,1,rep,name=qualities,proto3" json:"qualities,omitempty"`
	Level     []int64    `protobuf:"varint,2,rep,packed,name=level,proto3" json:"level,omitempty"`
}

func (x *PredictAndGuessResp) Reset() {
	*x = PredictAndGuessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waterquality_waterquality_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictAndGuessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictAndGuessResp) ProtoMessage() {}

func (x *PredictAndGuessResp) ProtoReflect() protoreflect.Message {
	mi := &file_waterquality_waterquality_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictAndGuessResp.ProtoReflect.Descriptor instead.
func (*PredictAndGuessResp) Descriptor() ([]byte, []int) {
	return file_waterquality_waterquality_proto_rawDescGZIP(), []int{5}
}

func (x *PredictAndGuessResp) GetQualities() []*Quality {
	if x != nil {
		return x.Qualities
	}
	return nil
}

func (x *PredictAndGuessResp) GetLevel() []int64 {
	if x != nil {
		return x.Level
	}
	return nil
}

var File_waterquality_waterquality_proto protoreflect.FileDescriptor

var file_waterquality_waterquality_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0x59, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x70, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x70, 0x68, 0x22, 0x7c,
	0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x09,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x42, 0x61,
	0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x0b,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x84, 0x01, 0x0a,
	0x12, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74,
	0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x41, 0x6e,
	0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x8e, 0x02, 0x0a, 0x13, 0x57, 0x61,
	0x74, 0x65, 0x72, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0a, 0x47,
	0x75, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x47,
	0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53, 0x75, 0x72,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x3b, 0x77, 0x61, 0x74, 0x65, 0x72, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waterquality_waterquality_proto_rawDescOnce sync.Once
	file_waterquality_waterquality_proto_rawDescData = file_waterquality_waterquality_proto_rawDesc
)

func file_waterquality_waterquality_proto_rawDescGZIP() []byte {
	file_waterquality_waterquality_proto_rawDescOnce.Do(func() {
		file_waterquality_waterquality_proto_rawDescData = protoimpl.X.CompressGZIP(file_waterquality_waterquality_proto_rawDescData)
	})
	return file_waterquality_waterquality_proto_rawDescData
}

var file_waterquality_waterquality_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_waterquality_waterquality_proto_goTypes = []interface{}{
	(*Quality)(nil),             // 0: api.waterquality.Quality
	(*PredictReq)(nil),          // 1: api.waterquality.PredictReq
	(*PredictResp)(nil),         // 2: api.waterquality.PredictResp
	(*GuessLevelResp)(nil),      // 3: api.waterquality.GuessLevelResp
	(*PredictAndGuessReq)(nil),  // 4: api.waterquality.PredictAndGuessReq
	(*PredictAndGuessResp)(nil), // 5: api.waterquality.PredictAndGuessResp
}
var file_waterquality_waterquality_proto_depIdxs = []int32{
	0, // 0: api.waterquality.PredictReq.qualities:type_name -> api.waterquality.Quality
	0, // 1: api.waterquality.PredictResp.qualities:type_name -> api.waterquality.Quality
	0, // 2: api.waterquality.PredictAndGuessReq.qualities:type_name -> api.waterquality.Quality
	0, // 3: api.waterquality.PredictAndGuessResp.qualities:type_name -> api.waterquality.Quality
	1, // 4: api.waterquality.WaterQualityService.Predict:input_type -> api.waterquality.PredictReq
	0, // 5: api.waterquality.WaterQualityService.GuessLevel:input_type -> api.waterquality.Quality
	4, // 6: api.waterquality.WaterQualityService.PredictAndGuess:input_type -> api.waterquality.PredictAndGuessReq
	2, // 7: api.waterquality.WaterQualityService.Predict:output_type -> api.waterquality.PredictResp
	3, // 8: api.waterquality.WaterQualityService.GuessLevel:output_type -> api.waterquality.GuessLevelResp
	5, // 9: api.waterquality.WaterQualityService.PredictAndGuess:output_type -> api.waterquality.PredictAndGuessResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_waterquality_waterquality_proto_init() }
func file_waterquality_waterquality_proto_init() {
	if File_waterquality_waterquality_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waterquality_waterquality_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quality); i {
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
		file_waterquality_waterquality_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictReq); i {
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
		file_waterquality_waterquality_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResp); i {
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
		file_waterquality_waterquality_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessLevelResp); i {
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
		file_waterquality_waterquality_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictAndGuessReq); i {
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
		file_waterquality_waterquality_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictAndGuessResp); i {
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
			RawDescriptor: file_waterquality_waterquality_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_waterquality_waterquality_proto_goTypes,
		DependencyIndexes: file_waterquality_waterquality_proto_depIdxs,
		MessageInfos:      file_waterquality_waterquality_proto_msgTypes,
	}.Build()
	File_waterquality_waterquality_proto = out.File
	file_waterquality_waterquality_proto_rawDesc = nil
	file_waterquality_waterquality_proto_goTypes = nil
	file_waterquality_waterquality_proto_depIdxs = nil
}
