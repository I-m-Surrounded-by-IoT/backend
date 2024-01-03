// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: collector/collector.proto

package collector

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

const OperationCollectorReportImmediately = "/api.collector.Collector/ReportImmediately"
const OperationCollectorSetReportInterval = "/api.collector.Collector/SetReportInterval"

type CollectorHTTPServer interface {
	ReportImmediately(context.Context, *ReportImmediatelyReq) (*ReportImmediatelyResp, error)
	SetReportInterval(context.Context, *SetReportIntervalReq) (*SetReportIntervalResp, error)
}

func RegisterCollectorHTTPServer(s *http.Server, srv CollectorHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/collector/report_immediately", _Collector_ReportImmediately0_HTTP_Handler(srv))
	r.POST("/v1/collector/set_report_interval", _Collector_SetReportInterval0_HTTP_Handler(srv))
}

func _Collector_ReportImmediately0_HTTP_Handler(srv CollectorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReportImmediatelyReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectorReportImmediately)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReportImmediately(ctx, req.(*ReportImmediatelyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReportImmediatelyResp)
		return ctx.Result(200, reply)
	}
}

func _Collector_SetReportInterval0_HTTP_Handler(srv CollectorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetReportIntervalReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectorSetReportInterval)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetReportInterval(ctx, req.(*SetReportIntervalReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetReportIntervalResp)
		return ctx.Result(200, reply)
	}
}

type CollectorHTTPClient interface {
	ReportImmediately(ctx context.Context, req *ReportImmediatelyReq, opts ...http.CallOption) (rsp *ReportImmediatelyResp, err error)
	SetReportInterval(ctx context.Context, req *SetReportIntervalReq, opts ...http.CallOption) (rsp *SetReportIntervalResp, err error)
}

type CollectorHTTPClientImpl struct {
	cc *http.Client
}

func NewCollectorHTTPClient(client *http.Client) CollectorHTTPClient {
	return &CollectorHTTPClientImpl{client}
}

func (c *CollectorHTTPClientImpl) ReportImmediately(ctx context.Context, in *ReportImmediatelyReq, opts ...http.CallOption) (*ReportImmediatelyResp, error) {
	var out ReportImmediatelyResp
	pattern := "/v1/collector/report_immediately"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectorReportImmediately))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectorHTTPClientImpl) SetReportInterval(ctx context.Context, in *SetReportIntervalReq, opts ...http.CallOption) (*SetReportIntervalResp, error) {
	var out SetReportIntervalResp
	pattern := "/v1/collector/set_report_interval"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectorSetReportInterval))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
