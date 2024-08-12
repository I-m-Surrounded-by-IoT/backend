// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v5.27.3
// source: captcha/captcha.proto

package captcha

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

const OperationCaptchaSendEmailCaptcha = "/api.captcha.Captcha/SendEmailCaptcha"
const OperationCaptchaVerifyEmailCaptcha = "/api.captcha.Captcha/VerifyEmailCaptcha"

type CaptchaHTTPServer interface {
	SendEmailCaptcha(context.Context, *SendEmailCaptchaReq) (*Empty, error)
	VerifyEmailCaptcha(context.Context, *VerifyEmailCaptchaReq) (*Empty, error)
}

func RegisterCaptchaHTTPServer(s *http.Server, srv CaptchaHTTPServer) {
	r := s.Route("/")
	r.POST("/captcha/sendEmailCaptcha", _Captcha_SendEmailCaptcha0_HTTP_Handler(srv))
	r.POST("/captcha/verifyEmailCaptcha", _Captcha_VerifyEmailCaptcha0_HTTP_Handler(srv))
}

func _Captcha_SendEmailCaptcha0_HTTP_Handler(srv CaptchaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendEmailCaptchaReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaptchaSendEmailCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendEmailCaptcha(ctx, req.(*SendEmailCaptchaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _Captcha_VerifyEmailCaptcha0_HTTP_Handler(srv CaptchaHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyEmailCaptchaReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaptchaVerifyEmailCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyEmailCaptcha(ctx, req.(*VerifyEmailCaptchaReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type CaptchaHTTPClient interface {
	SendEmailCaptcha(ctx context.Context, req *SendEmailCaptchaReq, opts ...http.CallOption) (rsp *Empty, err error)
	VerifyEmailCaptcha(ctx context.Context, req *VerifyEmailCaptchaReq, opts ...http.CallOption) (rsp *Empty, err error)
}

type CaptchaHTTPClientImpl struct {
	cc *http.Client
}

func NewCaptchaHTTPClient(client *http.Client) CaptchaHTTPClient {
	return &CaptchaHTTPClientImpl{client}
}

func (c *CaptchaHTTPClientImpl) SendEmailCaptcha(ctx context.Context, in *SendEmailCaptchaReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/captcha/sendEmailCaptcha"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCaptchaSendEmailCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CaptchaHTTPClientImpl) VerifyEmailCaptcha(ctx context.Context, in *VerifyEmailCaptchaReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/captcha/verifyEmailCaptcha"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCaptchaVerifyEmailCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
