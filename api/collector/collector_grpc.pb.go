// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: collector/collector.proto

package collector

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Collector_ReportImmediately_FullMethodName = "/api.collector.Collector/ReportImmediately"
	Collector_SetReportInterval_FullMethodName = "/api.collector.Collector/SetReportInterval"
)

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorClient interface {
	ReportImmediately(ctx context.Context, in *ReportImmediatelyReq, opts ...grpc.CallOption) (*Empty, error)
	SetReportInterval(ctx context.Context, in *SetReportIntervalReq, opts ...grpc.CallOption) (*SetReportIntervalResp, error)
}

type collectorClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorClient(cc grpc.ClientConnInterface) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) ReportImmediately(ctx context.Context, in *ReportImmediatelyReq, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Collector_ReportImmediately_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorClient) SetReportInterval(ctx context.Context, in *SetReportIntervalReq, opts ...grpc.CallOption) (*SetReportIntervalResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetReportIntervalResp)
	err := c.cc.Invoke(ctx, Collector_SetReportInterval_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServer is the server API for Collector service.
// All implementations must embed UnimplementedCollectorServer
// for forward compatibility
type CollectorServer interface {
	ReportImmediately(context.Context, *ReportImmediatelyReq) (*Empty, error)
	SetReportInterval(context.Context, *SetReportIntervalReq) (*SetReportIntervalResp, error)
	mustEmbedUnimplementedCollectorServer()
}

// UnimplementedCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorServer struct {
}

func (UnimplementedCollectorServer) ReportImmediately(context.Context, *ReportImmediatelyReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportImmediately not implemented")
}
func (UnimplementedCollectorServer) SetReportInterval(context.Context, *SetReportIntervalReq) (*SetReportIntervalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReportInterval not implemented")
}
func (UnimplementedCollectorServer) mustEmbedUnimplementedCollectorServer() {}

// UnsafeCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorServer will
// result in compilation errors.
type UnsafeCollectorServer interface {
	mustEmbedUnimplementedCollectorServer()
}

func RegisterCollectorServer(s grpc.ServiceRegistrar, srv CollectorServer) {
	s.RegisterService(&Collector_ServiceDesc, srv)
}

func _Collector_ReportImmediately_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportImmediatelyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).ReportImmediately(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_ReportImmediately_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).ReportImmediately(ctx, req.(*ReportImmediatelyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collector_SetReportInterval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReportIntervalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).SetReportInterval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collector_SetReportInterval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).SetReportInterval(ctx, req.(*SetReportIntervalReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Collector_ServiceDesc is the grpc.ServiceDesc for Collector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.collector.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportImmediately",
			Handler:    _Collector_ReportImmediately_Handler,
		},
		{
			MethodName: "SetReportInterval",
			Handler:    _Collector_SetReportInterval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector/collector.proto",
}
