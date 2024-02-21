// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: collection/collection.proto

package collection

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
	Collection_CreateCollectionRecord_FullMethodName = "/api.collection.Collection/CreateCollectionRecord"
	Collection_ListCollectionRecord_FullMethodName   = "/api.collection.Collection/ListCollectionRecord"
)

// CollectionClient is the client API for Collection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionClient interface {
	CreateCollectionRecord(ctx context.Context, in *CollectionRecord, opts ...grpc.CallOption) (*Empty, error)
	ListCollectionRecord(ctx context.Context, in *ListCollectionRecordReq, opts ...grpc.CallOption) (*ListCollectionRecordResp, error)
}

type collectionClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionClient(cc grpc.ClientConnInterface) CollectionClient {
	return &collectionClient{cc}
}

func (c *collectionClient) CreateCollectionRecord(ctx context.Context, in *CollectionRecord, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Collection_CreateCollectionRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionClient) ListCollectionRecord(ctx context.Context, in *ListCollectionRecordReq, opts ...grpc.CallOption) (*ListCollectionRecordResp, error) {
	out := new(ListCollectionRecordResp)
	err := c.cc.Invoke(ctx, Collection_ListCollectionRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServer is the server API for Collection service.
// All implementations must embed UnimplementedCollectionServer
// for forward compatibility
type CollectionServer interface {
	CreateCollectionRecord(context.Context, *CollectionRecord) (*Empty, error)
	ListCollectionRecord(context.Context, *ListCollectionRecordReq) (*ListCollectionRecordResp, error)
	mustEmbedUnimplementedCollectionServer()
}

// UnimplementedCollectionServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionServer struct {
}

func (UnimplementedCollectionServer) CreateCollectionRecord(context.Context, *CollectionRecord) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollectionRecord not implemented")
}
func (UnimplementedCollectionServer) ListCollectionRecord(context.Context, *ListCollectionRecordReq) (*ListCollectionRecordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionRecord not implemented")
}
func (UnimplementedCollectionServer) mustEmbedUnimplementedCollectionServer() {}

// UnsafeCollectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServer will
// result in compilation errors.
type UnsafeCollectionServer interface {
	mustEmbedUnimplementedCollectionServer()
}

func RegisterCollectionServer(s grpc.ServiceRegistrar, srv CollectionServer) {
	s.RegisterService(&Collection_ServiceDesc, srv)
}

func _Collection_CreateCollectionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).CreateCollectionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_CreateCollectionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).CreateCollectionRecord(ctx, req.(*CollectionRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Collection_ListCollectionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServer).ListCollectionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Collection_ListCollectionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServer).ListCollectionRecord(ctx, req.(*ListCollectionRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Collection_ServiceDesc is the grpc.ServiceDesc for Collection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.collection.Collection",
	HandlerType: (*CollectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollectionRecord",
			Handler:    _Collection_CreateCollectionRecord_Handler,
		},
		{
			MethodName: "ListCollectionRecord",
			Handler:    _Collection_ListCollectionRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection/collection.proto",
}