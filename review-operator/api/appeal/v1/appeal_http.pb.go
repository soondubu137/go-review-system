// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.3
// source: api/appeal/v1/appeal.proto

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

const OperationAppealCreateAppeal = "/api.appeal.v1.Appeal/CreateAppeal"
const OperationAppealResolveAppeal = "/api.appeal.v1.Appeal/ResolveAppeal"

type AppealHTTPServer interface {
	CreateAppeal(context.Context, *CreateAppealRequest) (*CreateAppealReply, error)
	ResolveAppeal(context.Context, *ResolveAppealRequest) (*ResolveAppealReply, error)
}

func RegisterAppealHTTPServer(s *http.Server, srv AppealHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/appeal", _Appeal_CreateAppeal0_HTTP_Handler(srv))
	r.POST("/v1/appeal/resolve", _Appeal_ResolveAppeal0_HTTP_Handler(srv))
}

func _Appeal_CreateAppeal0_HTTP_Handler(srv AppealHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAppealRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppealCreateAppeal)
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

func _Appeal_ResolveAppeal0_HTTP_Handler(srv AppealHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResolveAppealRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppealResolveAppeal)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResolveAppeal(ctx, req.(*ResolveAppealRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResolveAppealReply)
		return ctx.Result(200, reply)
	}
}

type AppealHTTPClient interface {
	CreateAppeal(ctx context.Context, req *CreateAppealRequest, opts ...http.CallOption) (rsp *CreateAppealReply, err error)
	ResolveAppeal(ctx context.Context, req *ResolveAppealRequest, opts ...http.CallOption) (rsp *ResolveAppealReply, err error)
}

type AppealHTTPClientImpl struct {
	cc *http.Client
}

func NewAppealHTTPClient(client *http.Client) AppealHTTPClient {
	return &AppealHTTPClientImpl{client}
}

func (c *AppealHTTPClientImpl) CreateAppeal(ctx context.Context, in *CreateAppealRequest, opts ...http.CallOption) (*CreateAppealReply, error) {
	var out CreateAppealReply
	pattern := "/v1/appeal"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppealCreateAppeal))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AppealHTTPClientImpl) ResolveAppeal(ctx context.Context, in *ResolveAppealRequest, opts ...http.CallOption) (*ResolveAppealReply, error) {
	var out ResolveAppealReply
	pattern := "/v1/appeal/resolve"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppealResolveAppeal))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}