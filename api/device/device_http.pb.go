// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.3
// source: device/device.proto

package device

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

const OperationDeviceCreateDevice = "/api.device.Device/CreateDevice"
const OperationDeviceGetDevice = "/api.device.Device/GetDevice"
const OperationDeviceGetDeviceByMac = "/api.device.Device/GetDeviceByMac"
const OperationDeviceGetOrCreateDevice = "/api.device.Device/GetOrCreateDevice"

type DeviceHTTPServer interface {
	CreateDevice(context.Context, *CreateDeviceReq) (*DeviceRecord, error)
	GetDevice(context.Context, *GetDeviceReq) (*DeviceRecord, error)
	GetDeviceByMac(context.Context, *GetDeviceByMacReq) (*DeviceRecord, error)
	GetOrCreateDevice(context.Context, *GetOrCreateDeviceReq) (*DeviceRecord, error)
}

func RegisterDeviceHTTPServer(s *http.Server, srv DeviceHTTPServer) {
	r := s.Route("/")
	r.GET("/device/{id}", _Device_GetDevice0_HTTP_Handler(srv))
	r.GET("/device/mac/{mac}", _Device_GetDeviceByMac0_HTTP_Handler(srv))
	r.POST("/device", _Device_CreateDevice0_HTTP_Handler(srv))
	r.POST("/device/get_or_create", _Device_GetOrCreateDevice0_HTTP_Handler(srv))
}

func _Device_GetDevice0_HTTP_Handler(srv DeviceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeviceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeviceGetDevice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDevice(ctx, req.(*GetDeviceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeviceRecord)
		return ctx.Result(200, reply)
	}
}

func _Device_GetDeviceByMac0_HTTP_Handler(srv DeviceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeviceByMacReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeviceGetDeviceByMac)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDeviceByMac(ctx, req.(*GetDeviceByMacReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeviceRecord)
		return ctx.Result(200, reply)
	}
}

func _Device_CreateDevice0_HTTP_Handler(srv DeviceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDeviceReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeviceCreateDevice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDevice(ctx, req.(*CreateDeviceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeviceRecord)
		return ctx.Result(200, reply)
	}
}

func _Device_GetOrCreateDevice0_HTTP_Handler(srv DeviceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrCreateDeviceReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDeviceGetOrCreateDevice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrCreateDevice(ctx, req.(*GetOrCreateDeviceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeviceRecord)
		return ctx.Result(200, reply)
	}
}

type DeviceHTTPClient interface {
	CreateDevice(ctx context.Context, req *CreateDeviceReq, opts ...http.CallOption) (rsp *DeviceRecord, err error)
	GetDevice(ctx context.Context, req *GetDeviceReq, opts ...http.CallOption) (rsp *DeviceRecord, err error)
	GetDeviceByMac(ctx context.Context, req *GetDeviceByMacReq, opts ...http.CallOption) (rsp *DeviceRecord, err error)
	GetOrCreateDevice(ctx context.Context, req *GetOrCreateDeviceReq, opts ...http.CallOption) (rsp *DeviceRecord, err error)
}

type DeviceHTTPClientImpl struct {
	cc *http.Client
}

func NewDeviceHTTPClient(client *http.Client) DeviceHTTPClient {
	return &DeviceHTTPClientImpl{client}
}

func (c *DeviceHTTPClientImpl) CreateDevice(ctx context.Context, in *CreateDeviceReq, opts ...http.CallOption) (*DeviceRecord, error) {
	var out DeviceRecord
	pattern := "/device"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDeviceCreateDevice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeviceHTTPClientImpl) GetDevice(ctx context.Context, in *GetDeviceReq, opts ...http.CallOption) (*DeviceRecord, error) {
	var out DeviceRecord
	pattern := "/device/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeviceGetDevice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeviceHTTPClientImpl) GetDeviceByMac(ctx context.Context, in *GetDeviceByMacReq, opts ...http.CallOption) (*DeviceRecord, error) {
	var out DeviceRecord
	pattern := "/device/mac/{mac}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDeviceGetDeviceByMac))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DeviceHTTPClientImpl) GetOrCreateDevice(ctx context.Context, in *GetOrCreateDeviceReq, opts ...http.CallOption) (*DeviceRecord, error) {
	var out DeviceRecord
	pattern := "/device/get_or_create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDeviceGetOrCreateDevice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
