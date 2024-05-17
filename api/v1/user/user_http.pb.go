// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: user.proto

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

const OperationUserAdminLogin = "/api.user.User/AdminLogin"
const OperationUserCreateUser = "/api.user.User/CreateUser"
const OperationUserGetUser = "/api.user.User/GetUser"
const OperationUserLogOut = "/api.user.User/LogOut"
const OperationUserLoginUser = "/api.user.User/LoginUser"
const OperationUserSendEmail = "/api.user.User/SendEmail"
const OperationUserSetBlack = "/api.user.User/SetBlack"
const OperationUserUpdatePassword = "/api.user.User/UpdatePassword"

type UserHTTPServer interface {
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	LogOut(context.Context, *LogOutRequest) (*LogOutReply, error)
	LoginUser(context.Context, *LoginRequest) (*LoginReply, error)
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailReply, error)
	SetBlack(context.Context, *SetBlackRequest) (*SetBlackReply, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/register/{code}", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/api/login", _User_LoginUser0_HTTP_Handler(srv))
	r.GET("/api/sendEmail/{email}", _User_SendEmail0_HTTP_Handler(srv))
	r.POST("/api/updatePassword/{code}", _User_UpdatePassword0_HTTP_Handler(srv))
	r.GET("/api/setBlack/{name}", _User_SetBlack0_HTTP_Handler(srv))
	r.GET("/api/getUserMessage/{name}", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/api/admin", _User_AdminLogin0_HTTP_Handler(srv))
	r.GET("/api/logOut/{name}", _User_LogOut0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginUser(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_SendEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendEmailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSendEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendEmail(ctx, req.(*SendEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendEmailReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdatePassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdatePassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePassword(ctx, req.(*UpdatePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePasswordReply)
		return ctx.Result(200, reply)
	}
}

func _User_SetBlack0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetBlackRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetBlack)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetBlack(ctx, req.(*SetBlackRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetBlackReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_AdminLogin0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAdminLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminLogin(ctx, req.(*AdminLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminLoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_LogOut0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogOutRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogOut)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LogOut(ctx, req.(*LogOutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogOutReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AdminLogin(ctx context.Context, req *AdminLoginRequest, opts ...http.CallOption) (rsp *AdminLoginReply, err error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	LogOut(ctx context.Context, req *LogOutRequest, opts ...http.CallOption) (rsp *LogOutReply, err error)
	LoginUser(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	SendEmail(ctx context.Context, req *SendEmailRequest, opts ...http.CallOption) (rsp *SendEmailReply, err error)
	SetBlack(ctx context.Context, req *SetBlackRequest, opts ...http.CallOption) (rsp *SetBlackReply, err error)
	UpdatePassword(ctx context.Context, req *UpdatePasswordRequest, opts ...http.CallOption) (rsp *UpdatePasswordReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...http.CallOption) (*AdminLoginReply, error) {
	var out AdminLoginReply
	pattern := "/api/admin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAdminLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/api/register/{code}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api/getUserMessage/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) LogOut(ctx context.Context, in *LogOutRequest, opts ...http.CallOption) (*LogOutReply, error) {
	var out LogOutReply
	pattern := "/api/logOut/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserLogOut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) LoginUser(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...http.CallOption) (*SendEmailReply, error) {
	var out SendEmailReply
	pattern := "/api/sendEmail/{email}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserSendEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) SetBlack(ctx context.Context, in *SetBlackRequest, opts ...http.CallOption) (*SetBlackReply, error) {
	var out SetBlackReply
	pattern := "/api/setBlack/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserSetBlack))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHTTPClientImpl) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...http.CallOption) (*UpdatePasswordReply, error) {
	var out UpdatePasswordReply
	pattern := "/api/updatePassword/{code}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdatePassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
