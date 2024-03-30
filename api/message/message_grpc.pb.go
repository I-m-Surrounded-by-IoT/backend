// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: message/message.proto

package message

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
	Message_SendMessage_FullMethodName    = "/api.message.Message/SendMessage"
	Message_MarkAllRead_FullMethodName    = "/api.message.Message/MarkAllRead"
	Message_GetUnreadNum_FullMethodName   = "/api.message.Message/GetUnreadNum"
	Message_GetMessage_FullMethodName     = "/api.message.Message/GetMessage"
	Message_GetMessageList_FullMethodName = "/api.message.Message/GetMessageList"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	SendMessage(ctx context.Context, in *MessagePayload, opts ...grpc.CallOption) (*Empty, error)
	MarkAllRead(ctx context.Context, in *MarkAllReadReq, opts ...grpc.CallOption) (*Empty, error)
	GetUnreadNum(ctx context.Context, in *GetUnreadNumReq, opts ...grpc.CallOption) (*GetUnreadNumResp, error)
	GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (*MessageRecord, error)
	GetMessageList(ctx context.Context, in *GetMessageListReq, opts ...grpc.CallOption) (*GetMessageListResp, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SendMessage(ctx context.Context, in *MessagePayload, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Message_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) MarkAllRead(ctx context.Context, in *MarkAllReadReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Message_MarkAllRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetUnreadNum(ctx context.Context, in *GetUnreadNumReq, opts ...grpc.CallOption) (*GetUnreadNumResp, error) {
	out := new(GetUnreadNumResp)
	err := c.cc.Invoke(ctx, Message_GetUnreadNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetMessage(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (*MessageRecord, error) {
	out := new(MessageRecord)
	err := c.cc.Invoke(ctx, Message_GetMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetMessageList(ctx context.Context, in *GetMessageListReq, opts ...grpc.CallOption) (*GetMessageListResp, error) {
	out := new(GetMessageListResp)
	err := c.cc.Invoke(ctx, Message_GetMessageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	SendMessage(context.Context, *MessagePayload) (*Empty, error)
	MarkAllRead(context.Context, *MarkAllReadReq) (*Empty, error)
	GetUnreadNum(context.Context, *GetUnreadNumReq) (*GetUnreadNumResp, error)
	GetMessage(context.Context, *GetMessageReq) (*MessageRecord, error)
	GetMessageList(context.Context, *GetMessageListReq) (*GetMessageListResp, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) SendMessage(context.Context, *MessagePayload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessageServer) MarkAllRead(context.Context, *MarkAllReadReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAllRead not implemented")
}
func (UnimplementedMessageServer) GetUnreadNum(context.Context, *GetUnreadNumReq) (*GetUnreadNumResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnreadNum not implemented")
}
func (UnimplementedMessageServer) GetMessage(context.Context, *GetMessageReq) (*MessageRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessageServer) GetMessageList(context.Context, *GetMessageListReq) (*GetMessageListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageList not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessage(ctx, req.(*MessagePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_MarkAllRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAllReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MarkAllRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MarkAllRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MarkAllRead(ctx, req.(*MarkAllReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetUnreadNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnreadNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetUnreadNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetUnreadNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetUnreadNum(ctx, req.(*GetUnreadNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetMessage(ctx, req.(*GetMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetMessageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetMessageList(ctx, req.(*GetMessageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Message_SendMessage_Handler,
		},
		{
			MethodName: "MarkAllRead",
			Handler:    _Message_MarkAllRead_Handler,
		},
		{
			MethodName: "GetUnreadNum",
			Handler:    _Message_GetUnreadNum_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Message_GetMessage_Handler,
		},
		{
			MethodName: "GetMessageList",
			Handler:    _Message_GetMessageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/message.proto",
}
