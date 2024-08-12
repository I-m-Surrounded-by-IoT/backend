// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: log/log.proto

package log

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Log_CreateDeviceLog_FullMethodName = "/api.log.Log/CreateDeviceLog"
	Log_ListDeviceLog_FullMethodName   = "/api.log.Log/ListDeviceLog"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	CreateDeviceLog(ctx context.Context, in *DeviceLog, opts ...grpc.CallOption) (*Empty, error)
	ListDeviceLog(ctx context.Context, in *ListDeviceLogReq, opts ...grpc.CallOption) (*ListDeviceLogResp, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) CreateDeviceLog(ctx context.Context, in *DeviceLog, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Log_CreateDeviceLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ListDeviceLog(ctx context.Context, in *ListDeviceLogReq, opts ...grpc.CallOption) (*ListDeviceLogResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDeviceLogResp)
	err := c.cc.Invoke(ctx, Log_ListDeviceLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility.
type LogServer interface {
	CreateDeviceLog(context.Context, *DeviceLog) (*Empty, error)
	ListDeviceLog(context.Context, *ListDeviceLogReq) (*ListDeviceLogResp, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServer struct{}

func (UnimplementedLogServer) CreateDeviceLog(context.Context, *DeviceLog) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceLog not implemented")
}
func (UnimplementedLogServer) ListDeviceLog(context.Context, *ListDeviceLogReq) (*ListDeviceLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeviceLog not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}
func (UnimplementedLogServer) testEmbeddedByValue()             {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	// If the following call pancis, it indicates UnimplementedLogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_CreateDeviceLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).CreateDeviceLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_CreateDeviceLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).CreateDeviceLog(ctx, req.(*DeviceLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ListDeviceLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).ListDeviceLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_ListDeviceLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).ListDeviceLog(ctx, req.(*ListDeviceLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.log.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceLog",
			Handler:    _Log_CreateDeviceLog_Handler,
		},
		{
			MethodName: "ListDeviceLog",
			Handler:    _Log_ListDeviceLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log/log.proto",
}
