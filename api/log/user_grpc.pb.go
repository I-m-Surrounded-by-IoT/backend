// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: log/user.proto

package log

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Log_AddDeviceLog_FullMethodName  = "/api.log.Log/AddDeviceLog"
	Log_GetDeviceLogs_FullMethodName = "/api.log.Log/GetDeviceLogs"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	AddDeviceLog(ctx context.Context, in *DeviceLog, opts ...grpc.CallOption) (*Empty, error)
	GetDeviceLogs(ctx context.Context, in *GetDeviceLogsReq, opts ...grpc.CallOption) (*DeviceLog, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) AddDeviceLog(ctx context.Context, in *DeviceLog, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Log_AddDeviceLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) GetDeviceLogs(ctx context.Context, in *GetDeviceLogsReq, opts ...grpc.CallOption) (*DeviceLog, error) {
	out := new(DeviceLog)
	err := c.cc.Invoke(ctx, Log_GetDeviceLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility
type LogServer interface {
	AddDeviceLog(context.Context, *DeviceLog) (*Empty, error)
	GetDeviceLogs(context.Context, *GetDeviceLogsReq) (*DeviceLog, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (UnimplementedLogServer) AddDeviceLog(context.Context, *DeviceLog) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeviceLog not implemented")
}
func (UnimplementedLogServer) GetDeviceLogs(context.Context, *GetDeviceLogsReq) (*DeviceLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceLogs not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_AddDeviceLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).AddDeviceLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_AddDeviceLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).AddDeviceLog(ctx, req.(*DeviceLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_GetDeviceLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceLogsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetDeviceLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_GetDeviceLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetDeviceLogs(ctx, req.(*GetDeviceLogsReq))
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
			MethodName: "AddDeviceLog",
			Handler:    _Log_AddDeviceLog_Handler,
		},
		{
			MethodName: "GetDeviceLogs",
			Handler:    _Log_GetDeviceLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log/user.proto",
}