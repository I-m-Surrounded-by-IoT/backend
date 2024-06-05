// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
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

var File_collector_collector_proto protoreflect.FileDescriptor

var file_collector_collector_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x17, 0x0a,
	0x15, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x32, 0x91, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x78, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22,
	0x1d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x12, 0x89,
	0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x2d, 0x6d, 0x2d, 0x53, 0x75, 0x72,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x2d, 0x62, 0x79, 0x2d, 0x49, 0x6f, 0x54, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x3b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_collector_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_collector_collector_proto_goTypes = []interface{}{
	(*ReportImmediatelyReq)(nil),  // 0: api.collector.ReportImmediatelyReq
	(*Empty)(nil),                 // 1: api.collector.Empty
	(*SetReportIntervalReq)(nil),  // 2: api.collector.SetReportIntervalReq
	(*SetReportIntervalResp)(nil), // 3: api.collector.SetReportIntervalResp
}
var file_collector_collector_proto_depIdxs = []int32{
	0, // 0: api.collector.Collector.ReportImmediately:input_type -> api.collector.ReportImmediatelyReq
	2, // 1: api.collector.Collector.SetReportInterval:input_type -> api.collector.SetReportIntervalReq
	1, // 2: api.collector.Collector.ReportImmediately:output_type -> api.collector.Empty
	3, // 3: api.collector.Collector.SetReportInterval:output_type -> api.collector.SetReportIntervalResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_collector_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
