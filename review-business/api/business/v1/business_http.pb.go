// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.3
// source: api/business/v1/business.proto

package v1

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

const OperationBusinessCreateAppeal = "/api.business.v1.Business/CreateAppeal"
const OperationBusinessCreateReply = "/api.business.v1.Business/CreateReply"

type BusinessHTTPServer interface {
	CreateAppeal(context.Context, *CreateAppealRequest) (*CreateAppealReply, error)
	CreateReply(context.Context, *CreateReplyRequest) (*CreateReplyReply, error)
}

func RegisterBusinessHTTPServer(s *http.Server, srv BusinessHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/reply", _Business_CreateReply0_HTTP_Handler(srv))
	r.POST("/v1/appeal", _Business_CreateAppeal0_HTTP_Handler(srv))
}

func _Business_CreateReply0_HTTP_Handler(srv BusinessHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReplyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBusinessCreateReply)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateReply(ctx, req.(*CreateReplyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateReplyReply)
		return ctx.Result(200, reply)
	}
}

func _Business_CreateAppeal0_HTTP_Handler(srv BusinessHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAppealRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBusinessCreateAppeal)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAppeal(ctx, req.(*CreateAppealRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAppealReply)
		return ctx.Result(200, reply)
	}
}

type BusinessHTTPClient interface {
	CreateAppeal(ctx context.Context, req *CreateAppealRequest, opts ...http.CallOption) (rsp *CreateAppealReply, err error)
	CreateReply(ctx context.Context, req *CreateReplyRequest, opts ...http.CallOption) (rsp *CreateReplyReply, err error)
}

type BusinessHTTPClientImpl struct {
	cc *http.Client
}

func NewBusinessHTTPClient(client *http.Client) BusinessHTTPClient {
	return &BusinessHTTPClientImpl{client}
}

func (c *BusinessHTTPClientImpl) CreateAppeal(ctx context.Context, in *CreateAppealRequest, opts ...http.CallOption) (*CreateAppealReply, error) {
	var out CreateAppealReply
	pattern := "/v1/appeal"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBusinessCreateAppeal))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BusinessHTTPClientImpl) CreateReply(ctx context.Context, in *CreateReplyRequest, opts ...http.CallOption) (*CreateReplyReply, error) {
	var out CreateReplyReply
	pattern := "/v1/reply"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBusinessCreateReply))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
