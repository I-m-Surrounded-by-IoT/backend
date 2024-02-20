// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.3
// source: user/user.proto

package user

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

const OperationUserCreateUser = "/api.user.User/CreateUser"
const OperationUserGetUser = "/api.user.User/GetUser"
const OperationUserSetUserName = "/api.user.User/SetUserName"
const OperationUserSetUserPassword = "/api.user.User/SetUserPassword"
const OperationUserSetUserRole = "/api.user.User/SetUserRole"
const OperationUserSetUserStatus = "/api.user.User/SetUserStatus"
const OperationUserValidateUserPassword = "/api.user.User/ValidateUserPassword"

type UserHTTPServer interface {
	CreateUser(context.Context, *CreateUserReq) (*UserInfo, error)
	GetUser(context.Context, *GetUserReq) (*UserInfo, error)
	SetUserName(context.Context, *SetUserNameReq) (*Empty, error)
	SetUserPassword(context.Context, *SetUserPasswordReq) (*Empty, error)
	SetUserRole(context.Context, *SetUserRoleReq) (*Empty, error)
	SetUserStatus(context.Context, *SetUserStatusReq) (*Empty, error)
	ValidateUserPassword(context.Context, *ValidateUserPasswordReq) (*ValidateUserPasswordResp, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/user", _User_CreateUser0_HTTP_Handler(srv))
	r.GET("/user/{id}", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/user/validate", _User_ValidateUserPassword0_HTTP_Handler(srv))
	r.PUT("/user/password", _User_SetUserPassword0_HTTP_Handler(srv))
	r.PUT("/user/role", _User_SetUserRole0_HTTP_Handler(srv))
	r.PUT("/user/status", _User_SetUserStatus0_HTTP_Handler(srv))
	r.PUT("/user/name", _User_SetUserName0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfo)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfo)
		return ctx.Result(200, reply)
	}
}

func _User_ValidateUserPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ValidateUserPasswordReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserValidateUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ValidateUserPassword(ctx, req.(*ValidateUserPasswordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValidateUserPasswordResp)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserPasswordReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserPassword(ctx, req.(*SetUserPasswordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserRole0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserRoleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetUserRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserRole(ctx, req.(*SetUserRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserStatus0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserStatusReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetUserStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserStatus(ctx, req.(*SetUserStatusReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserName0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserNameReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetUserName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserName(ctx, req.(*SetUserNameReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserReq, opts ...http.CallOption) (rsp *UserInfo, err error)
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *UserInfo, err error)
	SetUserName(ctx context.Context, req *SetUserNameReq, opts ...http.CallOption) (rsp *Empty, err error)
	SetUserPassword(ctx context.Context, req *SetUserPasswordReq, opts ...http.CallOption) (rsp *Empty, err error)
	SetUserRole(ctx context.Context, req *SetUserRoleReq, opts ...http.CallOption) (rsp *Empty, err error)
	SetUserStatus(ctx context.Context, req *SetUserStatusReq, opts ...http.CallOption) (rsp *Empty, err error)
	ValidateUserPassword(ctx context.Context, req *ValidateUserPasswordReq, opts ...http.CallOption) (rsp *ValidateUserPasswordResp, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserReq, opts ...http.CallOption) (*UserInfo, error) {
	var out UserInfo
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*UserInfo, error) {
	var out UserInfo
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserName(ctx context.Context, in *SetUserNameReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/user/name"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSetUserName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserPassword(ctx context.Context, in *SetUserPasswordReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/user/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSetUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserRole(ctx context.Context, in *SetUserRoleReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/user/role"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSetUserRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserStatus(ctx context.Context, in *SetUserStatusReq, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/user/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSetUserStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ValidateUserPassword(ctx context.Context, in *ValidateUserPasswordReq, opts ...http.CallOption) (*ValidateUserPasswordResp, error) {
	var out ValidateUserPasswordResp
	pattern := "/user/validate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserValidateUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
