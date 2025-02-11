// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: api/operator/v1/operator.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Operator_ResolveAppeal_FullMethodName = "/api.operator.v1.Operator/ResolveAppeal"
)

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	ResolveAppeal(ctx context.Context, in *ResolveAppealRequest, opts ...grpc.CallOption) (*ResolveAppealReply, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) ResolveAppeal(ctx context.Context, in *ResolveAppealRequest, opts ...grpc.CallOption) (*ResolveAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveAppealReply)
	err := c.cc.Invoke(ctx, Operator_ResolveAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations must embed UnimplementedOperatorServer
// for forward compatibility.
type OperatorServer interface {
	ResolveAppeal(context.Context, *ResolveAppealRequest) (*ResolveAppealReply, error)
	mustEmbedUnimplementedOperatorServer()
}

// UnimplementedOperatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOperatorServer struct{}

func (UnimplementedOperatorServer) ResolveAppeal(context.Context, *ResolveAppealRequest) (*ResolveAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAppeal not implemented")
}
func (UnimplementedOperatorServer) mustEmbedUnimplementedOperatorServer() {}
func (UnimplementedOperatorServer) testEmbeddedByValue()                  {}

// UnsafeOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServer will
// result in compilation errors.
type UnsafeOperatorServer interface {
	mustEmbedUnimplementedOperatorServer()
}

func RegisterOperatorServer(s grpc.ServiceRegistrar, srv OperatorServer) {
	// If the following call pancis, it indicates UnimplementedOperatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Operator_ServiceDesc, srv)
}

func _Operator_ResolveAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).ResolveAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_ResolveAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).ResolveAppeal(ctx, req.(*ResolveAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operator_ServiceDesc is the grpc.ServiceDesc for Operator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operator.v1.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveAppeal",
			Handler:    _Operator_ResolveAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/operator/v1/operator.proto",
}
