// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.3
// source: notify/notify.proto

package notify

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNotifyNotifyDeviceOffline = "/api.notify.Notify/NotifyDeviceOffline"
const OperationNotifyNotifyDeviceOnline = "/api.notify.Notify/NotifyDeviceOnline"

type NotifyHTTPServer interface {
	NotifyDeviceOffline(context.Context, *NotifyDeviceOfflineReq) (*Empty, error)
	NotifyDeviceOnline(context.Context, *NotifyDeviceOnlineReq) (*Empty, error)
}

func RegisterNotifyHTTPServer(s *http.Server, srv NotifyHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/notify/device/online", _Notify_NotifyDeviceOnline0_HTTP_Handler(srv))
	r.POST("/v1/notify/device/offline", _Notify_NotifyDeviceOffline0_HTTP_Handler(srv))
}

func _Notify_NotifyDeviceOnline0_HTTP_Handler(srv NotifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotifyDeviceOnlineReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotifyNotifyDeviceOnline)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NotifyDeviceOnline(ctx, req.(*NotifyDeviceOnlineReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Notify_NotifyDeviceOffline0_HTTP_Handler(srv NotifyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NotifyDeviceOfflineReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotifyNotifyDeviceOffline)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NotifyDeviceOffline(ctx, req.(*NotifyDeviceOfflineReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type NotifyHTTPClient interface {
	NotifyDeviceOffline(ctx context.Context, req *NotifyDeviceOfflineReq, opts ...http.CallOption) (rsp *Empty, err error)
	NotifyDeviceOnline(ctx context.Context, req *NotifyDeviceOnlineReq, opts ...http.CallOption) (rsp *Empty, err error)
}

type NotifyHTTPClientImpl struct {
	cc *http.Client
}

func NewNotifyHTTPClient(client *http.Client) NotifyHTTPClient {
	return &NotifyHTTPClientImpl{client}
}

func (c *NotifyHTTPClientImpl) NotifyDeviceOffline(ctx context.Context, in *NotifyDeviceOfflineReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/v1/notify/device/offline"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotifyNotifyDeviceOffline))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NotifyHTTPClientImpl) NotifyDeviceOnline(ctx context.Context, in *NotifyDeviceOnlineReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/v1/notify/device/online"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotifyNotifyDeviceOnline))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
