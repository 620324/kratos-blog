// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v4.25.1
// source: api/v1/comment/comment.proto

package comment

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

const OperationCommentAddComment = "/api.comment.Comment/AddComment"
const OperationCommentAddReward = "/api.comment.Comment/AddReward"
const OperationCommentExtractParentComments = "/api.comment.Comment/ExtractParentComments"

type CommentHTTPServer interface {
	AddComment(context.Context, *CreateCommentRequest) (*CreateCommentReply, error)
	AddReward(context.Context, *CreateRewardRequest) (*CreateRewardReply, error)
	ExtractParentComments(context.Context, *ExtractParentCommentsRequest) (*ExtractParentCommentsReply, error)
}

func RegisterCommentHTTPServer(s *http.Server, srv CommentHTTPServer) {
	r := s.Route("/")
	r.POST("/api/addComment", _Comment_AddComment0_HTTP_Handler(srv))
	r.POST("/api/addReward/{reward_id}", _Comment_AddReward0_HTTP_Handler(srv))
	r.GET("/api/getComment/{id}", _Comment_ExtractParentComments0_HTTP_Handler(srv))
}

func _Comment_AddComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentAddComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComment(ctx, req.(*CreateCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCommentReply)
		return ctx.Result(200, reply)
	}
}

func _Comment_AddReward0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRewardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentAddReward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddReward(ctx, req.(*CreateRewardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRewardReply)
		return ctx.Result(200, reply)
	}
}

func _Comment_ExtractParentComments0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExtractParentCommentsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentExtractParentComments)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExtractParentComments(ctx, req.(*ExtractParentCommentsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExtractParentCommentsReply)
		return ctx.Result(200, reply)
	}
}

type CommentHTTPClient interface {
	AddComment(ctx context.Context, req *CreateCommentRequest, opts ...http.CallOption) (rsp *CreateCommentReply, err error)
	AddReward(ctx context.Context, req *CreateRewardRequest, opts ...http.CallOption) (rsp *CreateRewardReply, err error)
	ExtractParentComments(ctx context.Context, req *ExtractParentCommentsRequest, opts ...http.CallOption) (rsp *ExtractParentCommentsReply, err error)
}

type CommentHTTPClientImpl struct {
	cc *http.Client
}

func NewCommentHTTPClient(client *http.Client) CommentHTTPClient {
	return &CommentHTTPClientImpl{client}
}

func (c *CommentHTTPClientImpl) AddComment(ctx context.Context, in *CreateCommentRequest, opts ...http.CallOption) (*CreateCommentReply, error) {
	var out CreateCommentReply
	pattern := "/api/addComment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentAddComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) AddReward(ctx context.Context, in *CreateRewardRequest, opts ...http.CallOption) (*CreateRewardReply, error) {
	var out CreateRewardReply
	pattern := "/api/addReward/{reward_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentAddReward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) ExtractParentComments(ctx context.Context, in *ExtractParentCommentsRequest, opts ...http.CallOption) (*ExtractParentCommentsReply, error) {
	var out ExtractParentCommentsReply
	pattern := "/api/getComment/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentExtractParentComments))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
