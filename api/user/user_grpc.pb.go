// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: user/user.proto

package user

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
	User_CreateUser_FullMethodName           = "/api.user.User/CreateUser"
	User_GetUser_FullMethodName              = "/api.user.User/GetUser"
	User_ValidateUserPassword_FullMethodName = "/api.user.User/ValidateUserPassword"
	User_SetUserPassword_FullMethodName      = "/api.user.User/SetUserPassword"
	User_SetUserRole_FullMethodName          = "/api.user.User/SetUserRole"
	User_SetUserStatus_FullMethodName        = "/api.user.User/SetUserStatus"
	User_SetUserName_FullMethodName          = "/api.user.User/SetUserName"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	ValidateUserPassword(ctx context.Context, in *ValidateUserPasswordReq, opts ...grpc.CallOption) (*ValidateUserPasswordResp, error)
	SetUserPassword(ctx context.Context, in *SetUserPasswordReq, opts ...grpc.CallOption) (*Empty, error)
	SetUserRole(ctx context.Context, in *SetUserRoleReq, opts ...grpc.CallOption) (*Empty, error)
	SetUserStatus(ctx context.Context, in *SetUserStatusReq, opts ...grpc.CallOption) (*Empty, error)
	SetUserName(ctx context.Context, in *SetUserNameReq, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, User_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateUserPassword(ctx context.Context, in *ValidateUserPasswordReq, opts ...grpc.CallOption) (*ValidateUserPasswordResp, error) {
	out := new(ValidateUserPasswordResp)
	err := c.cc.Invoke(ctx, User_ValidateUserPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserPassword(ctx context.Context, in *SetUserPasswordReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_SetUserPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserRole(ctx context.Context, in *SetUserRoleReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_SetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserStatus(ctx context.Context, in *SetUserStatusReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_SetUserStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserName(ctx context.Context, in *SetUserNameReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, User_SetUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUser(context.Context, *CreateUserReq) (*GetUserResp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	ValidateUserPassword(context.Context, *ValidateUserPasswordReq) (*ValidateUserPasswordResp, error)
	SetUserPassword(context.Context, *SetUserPasswordReq) (*Empty, error)
	SetUserRole(context.Context, *SetUserRoleReq) (*Empty, error)
	SetUserStatus(context.Context, *SetUserStatusReq) (*Empty, error)
	SetUserName(context.Context, *SetUserNameReq) (*Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) ValidateUserPassword(context.Context, *ValidateUserPasswordReq) (*ValidateUserPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserPassword not implemented")
}
func (UnimplementedUserServer) SetUserPassword(context.Context, *SetUserPasswordReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPassword not implemented")
}
func (UnimplementedUserServer) SetUserRole(context.Context, *SetUserRoleReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserRole not implemented")
}
func (UnimplementedUserServer) SetUserStatus(context.Context, *SetUserStatusReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserStatus not implemented")
}
func (UnimplementedUserServer) SetUserName(context.Context, *SetUserNameReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserName not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ValidateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ValidateUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidateUserPassword(ctx, req.(*ValidateUserPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserPassword(ctx, req.(*SetUserPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserRole(ctx, req.(*SetUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserStatus(ctx, req.(*SetUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserName(ctx, req.(*SetUserNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "ValidateUserPassword",
			Handler:    _User_ValidateUserPassword_Handler,
		},
		{
			MethodName: "SetUserPassword",
			Handler:    _User_SetUserPassword_Handler,
		},
		{
			MethodName: "SetUserRole",
			Handler:    _User_SetUserRole_Handler,
		},
		{
			MethodName: "SetUserStatus",
			Handler:    _User_SetUserStatus_Handler,
		},
		{
			MethodName: "SetUserName",
			Handler:    _User_SetUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
