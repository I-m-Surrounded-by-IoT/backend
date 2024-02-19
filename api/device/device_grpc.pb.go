// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: device/device.proto

package device

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
	Device_GetDevice_FullMethodName         = "/api.device.Device/GetDevice"
	Device_GetDeviceByMac_FullMethodName    = "/api.device.Device/GetDeviceByMac"
	Device_CreateDevice_FullMethodName      = "/api.device.Device/CreateDevice"
	Device_GetOrCreateDevice_FullMethodName = "/api.device.Device/GetOrCreateDevice"
)

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	GetDevice(ctx context.Context, in *GetDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error)
	GetDeviceByMac(ctx context.Context, in *GetDeviceByMacReq, opts ...grpc.CallOption) (*DeviceRecord, error)
	CreateDevice(ctx context.Context, in *CreateDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error)
	GetOrCreateDevice(ctx context.Context, in *GetOrCreateDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) GetDevice(ctx context.Context, in *GetDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error) {
	out := new(DeviceRecord)
	err := c.cc.Invoke(ctx, Device_GetDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDeviceByMac(ctx context.Context, in *GetDeviceByMacReq, opts ...grpc.CallOption) (*DeviceRecord, error) {
	out := new(DeviceRecord)
	err := c.cc.Invoke(ctx, Device_GetDeviceByMac_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) CreateDevice(ctx context.Context, in *CreateDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error) {
	out := new(DeviceRecord)
	err := c.cc.Invoke(ctx, Device_CreateDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetOrCreateDevice(ctx context.Context, in *GetOrCreateDeviceReq, opts ...grpc.CallOption) (*DeviceRecord, error) {
	out := new(DeviceRecord)
	err := c.cc.Invoke(ctx, Device_GetOrCreateDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations must embed UnimplementedDeviceServer
// for forward compatibility
type DeviceServer interface {
	GetDevice(context.Context, *GetDeviceReq) (*DeviceRecord, error)
	GetDeviceByMac(context.Context, *GetDeviceByMacReq) (*DeviceRecord, error)
	CreateDevice(context.Context, *CreateDeviceReq) (*DeviceRecord, error)
	GetOrCreateDevice(context.Context, *GetOrCreateDeviceReq) (*DeviceRecord, error)
	mustEmbedUnimplementedDeviceServer()
}

// UnimplementedDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (UnimplementedDeviceServer) GetDevice(context.Context, *GetDeviceReq) (*DeviceRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedDeviceServer) GetDeviceByMac(context.Context, *GetDeviceByMacReq) (*DeviceRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByMac not implemented")
}
func (UnimplementedDeviceServer) CreateDevice(context.Context, *CreateDeviceReq) (*DeviceRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDeviceServer) GetOrCreateDevice(context.Context, *GetOrCreateDeviceReq) (*DeviceRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateDevice not implemented")
}
func (UnimplementedDeviceServer) mustEmbedUnimplementedDeviceServer() {}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_GetDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDevice(ctx, req.(*GetDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDeviceByMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceByMacReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceByMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_GetDeviceByMac_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceByMac(ctx, req.(*GetDeviceByMacReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).CreateDevice(ctx, req.(*CreateDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetOrCreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetOrCreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_GetOrCreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetOrCreateDevice(ctx, req.(*GetOrCreateDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.device.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDevice",
			Handler:    _Device_GetDevice_Handler,
		},
		{
			MethodName: "GetDeviceByMac",
			Handler:    _Device_GetDeviceByMac_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _Device_CreateDevice_Handler,
		},
		{
			MethodName: "GetOrCreateDevice",
			Handler:    _Device_GetOrCreateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device/device.proto",
}
