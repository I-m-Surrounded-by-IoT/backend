// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v5.27.3
// source: collection/collection.proto

package collection

import (
	context "context"
	waterquality "github.com/I-m-Surrounded-by-IoT/backend/api/waterquality"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCollectionCreateCollectionRecord = "/api.collection.Collection/CreateCollectionRecord"
const OperationCollectionGetDeviceLastReport = "/api.collection.Collection/GetDeviceLastReport"
const OperationCollectionGetLatestIdWithinRange = "/api.collection.Collection/GetLatestIdWithinRange"
const OperationCollectionGetLatestRecordsWithinRange = "/api.collection.Collection/GetLatestRecordsWithinRange"
const OperationCollectionGetPredictQuality = "/api.collection.Collection/GetPredictQuality"
const OperationCollectionListCollectionRecord = "/api.collection.Collection/ListCollectionRecord"

type CollectionHTTPServer interface {
	CreateCollectionRecord(context.Context, *CreateCollectionRecordReq) (*Empty, error)
	GetDeviceLastReport(context.Context, *GetDeviceLastReportReq) (*CollectionRecord, error)
	GetLatestIdWithinRange(context.Context, *GetLatestWithinRangeReq) (*GetLatestIdWithinRangeResp, error)
	GetLatestRecordsWithinRange(context.Context, *GetLatestWithinRangeReq) (*GetLatestRecordsWithinRangeResp, error)
	GetPredictQuality(context.Context, *GetPredictQualityReq) (*waterquality.PredictAndGuessResp, error)
	ListCollectionRecord(context.Context, *ListCollectionRecordReq) (*ListCollectionRecordResp, error)
}

func RegisterCollectionHTTPServer(s *http.Server, srv CollectionHTTPServer) {
	r := s.Route("/")
	r.POST("/collection/record", _Collection_CreateCollectionRecord0_HTTP_Handler(srv))
	r.GET("/collection/record", _Collection_ListCollectionRecord0_HTTP_Handler(srv))
	r.GET("/collection/predict_quality", _Collection_GetPredictQuality0_HTTP_Handler(srv))
	r.GET("/device/{id}/last_report", _Collection_GetDeviceLastReport0_HTTP_Handler(srv))
	r.POST("/collection/get_latest_within_range", _Collection_GetLatestRecordsWithinRange0_HTTP_Handler(srv))
	r.POST("/collection/get_latest_id_within_range", _Collection_GetLatestIdWithinRange0_HTTP_Handler(srv))
}

func _Collection_CreateCollectionRecord0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCollectionRecordReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionCreateCollectionRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCollectionRecord(ctx, req.(*CreateCollectionRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Collection_ListCollectionRecord0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCollectionRecordReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionListCollectionRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCollectionRecord(ctx, req.(*ListCollectionRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCollectionRecordResp)
		return ctx.Result(200, reply)
	}
}

func _Collection_GetPredictQuality0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPredictQualityReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionGetPredictQuality)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPredictQuality(ctx, req.(*GetPredictQualityReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*waterquality.PredictAndGuessResp)
		return ctx.Result(200, reply)
	}
}

func _Collection_GetDeviceLastReport0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeviceLastReportReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionGetDeviceLastReport)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDeviceLastReport(ctx, req.(*GetDeviceLastReportReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollectionRecord)
		return ctx.Result(200, reply)
	}
}

func _Collection_GetLatestRecordsWithinRange0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLatestWithinRangeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionGetLatestRecordsWithinRange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLatestRecordsWithinRange(ctx, req.(*GetLatestWithinRangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLatestRecordsWithinRangeResp)
		return ctx.Result(200, reply)
	}
}

func _Collection_GetLatestIdWithinRange0_HTTP_Handler(srv CollectionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLatestWithinRangeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionGetLatestIdWithinRange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLatestIdWithinRange(ctx, req.(*GetLatestWithinRangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLatestIdWithinRangeResp)
		return ctx.Result(200, reply)
	}
}

type CollectionHTTPClient interface {
	CreateCollectionRecord(ctx context.Context, req *CreateCollectionRecordReq, opts ...http.CallOption) (rsp *Empty, err error)
	GetDeviceLastReport(ctx context.Context, req *GetDeviceLastReportReq, opts ...http.CallOption) (rsp *CollectionRecord, err error)
	GetLatestIdWithinRange(ctx context.Context, req *GetLatestWithinRangeReq, opts ...http.CallOption) (rsp *GetLatestIdWithinRangeResp, err error)
	GetLatestRecordsWithinRange(ctx context.Context, req *GetLatestWithinRangeReq, opts ...http.CallOption) (rsp *GetLatestRecordsWithinRangeResp, err error)
	GetPredictQuality(ctx context.Context, req *GetPredictQualityReq, opts ...http.CallOption) (rsp *waterquality.PredictAndGuessResp, err error)
	ListCollectionRecord(ctx context.Context, req *ListCollectionRecordReq, opts ...http.CallOption) (rsp *ListCollectionRecordResp, err error)
}

type CollectionHTTPClientImpl struct {
	cc *http.Client
}

func NewCollectionHTTPClient(client *http.Client) CollectionHTTPClient {
	return &CollectionHTTPClientImpl{client}
}

func (c *CollectionHTTPClientImpl) CreateCollectionRecord(ctx context.Context, in *CreateCollectionRecordReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/collection/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionCreateCollectionRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectionHTTPClientImpl) GetDeviceLastReport(ctx context.Context, in *GetDeviceLastReportReq, opts ...http.CallOption) (*CollectionRecord, error) {
	var out CollectionRecord
	pattern := "/device/{id}/last_report"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCollectionGetDeviceLastReport))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectionHTTPClientImpl) GetLatestIdWithinRange(ctx context.Context, in *GetLatestWithinRangeReq, opts ...http.CallOption) (*GetLatestIdWithinRangeResp, error) {
	var out GetLatestIdWithinRangeResp
	pattern := "/collection/get_latest_id_within_range"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionGetLatestIdWithinRange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectionHTTPClientImpl) GetLatestRecordsWithinRange(ctx context.Context, in *GetLatestWithinRangeReq, opts ...http.CallOption) (*GetLatestRecordsWithinRangeResp, error) {
	var out GetLatestRecordsWithinRangeResp
	pattern := "/collection/get_latest_within_range"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionGetLatestRecordsWithinRange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectionHTTPClientImpl) GetPredictQuality(ctx context.Context, in *GetPredictQualityReq, opts ...http.CallOption) (*waterquality.PredictAndGuessResp, error) {
	var out waterquality.PredictAndGuessResp
	pattern := "/collection/predict_quality"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCollectionGetPredictQuality))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CollectionHTTPClientImpl) ListCollectionRecord(ctx context.Context, in *ListCollectionRecordReq, opts ...http.CallOption) (*ListCollectionRecordResp, error) {
	var out ListCollectionRecordResp
	pattern := "/collection/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCollectionListCollectionRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
