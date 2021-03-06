// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_obligation_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ObligationRpcClient is the client API for ObligationRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObligationRpcClient interface {
	CreateObligation(ctx context.Context, in *CreateObligationRequest, opts ...grpc.CallOption) (*CreateObligationResponce, error)
	DescribeObligation(ctx context.Context, in *DescribeObligationRequest, opts ...grpc.CallOption) (*DescribeObligationResponse, error)
	ListObligations(ctx context.Context, in *ListObligationsRequest, opts ...grpc.CallOption) (*ListObligationsResponse, error)
	RemoveObligation(ctx context.Context, in *RemoveObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MultiCreateObligation(ctx context.Context, in *MultiCreateObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateObligation(ctx context.Context, in *UpdateObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type obligationRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewObligationRpcClient(cc grpc.ClientConnInterface) ObligationRpcClient {
	return &obligationRpcClient{cc}
}

func (c *obligationRpcClient) CreateObligation(ctx context.Context, in *CreateObligationRequest, opts ...grpc.CallOption) (*CreateObligationResponce, error) {
	out := new(CreateObligationResponce)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/CreateObligation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obligationRpcClient) DescribeObligation(ctx context.Context, in *DescribeObligationRequest, opts ...grpc.CallOption) (*DescribeObligationResponse, error) {
	out := new(DescribeObligationResponse)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/DescribeObligation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obligationRpcClient) ListObligations(ctx context.Context, in *ListObligationsRequest, opts ...grpc.CallOption) (*ListObligationsResponse, error) {
	out := new(ListObligationsResponse)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/ListObligations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obligationRpcClient) RemoveObligation(ctx context.Context, in *RemoveObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/RemoveObligation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obligationRpcClient) MultiCreateObligation(ctx context.Context, in *MultiCreateObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/MultiCreateObligation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obligationRpcClient) UpdateObligation(ctx context.Context, in *UpdateObligationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.obligation.api.ObligationRpc/UpdateObligation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObligationRpcServer is the server API for ObligationRpc service.
// All implementations must embed UnimplementedObligationRpcServer
// for forward compatibility
type ObligationRpcServer interface {
	CreateObligation(context.Context, *CreateObligationRequest) (*CreateObligationResponce, error)
	DescribeObligation(context.Context, *DescribeObligationRequest) (*DescribeObligationResponse, error)
	ListObligations(context.Context, *ListObligationsRequest) (*ListObligationsResponse, error)
	RemoveObligation(context.Context, *RemoveObligationRequest) (*emptypb.Empty, error)
	MultiCreateObligation(context.Context, *MultiCreateObligationRequest) (*emptypb.Empty, error)
	UpdateObligation(context.Context, *UpdateObligationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedObligationRpcServer()
}

// UnimplementedObligationRpcServer must be embedded to have forward compatible implementations.
type UnimplementedObligationRpcServer struct {
}

func (UnimplementedObligationRpcServer) CreateObligation(context.Context, *CreateObligationRequest) (*CreateObligationResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObligation not implemented")
}
func (UnimplementedObligationRpcServer) DescribeObligation(context.Context, *DescribeObligationRequest) (*DescribeObligationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeObligation not implemented")
}
func (UnimplementedObligationRpcServer) ListObligations(context.Context, *ListObligationsRequest) (*ListObligationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObligations not implemented")
}
func (UnimplementedObligationRpcServer) RemoveObligation(context.Context, *RemoveObligationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveObligation not implemented")
}
func (UnimplementedObligationRpcServer) MultiCreateObligation(context.Context, *MultiCreateObligationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateObligation not implemented")
}
func (UnimplementedObligationRpcServer) UpdateObligation(context.Context, *UpdateObligationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObligation not implemented")
}
func (UnimplementedObligationRpcServer) mustEmbedUnimplementedObligationRpcServer() {}

// UnsafeObligationRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObligationRpcServer will
// result in compilation errors.
type UnsafeObligationRpcServer interface {
	mustEmbedUnimplementedObligationRpcServer()
}

func RegisterObligationRpcServer(s grpc.ServiceRegistrar, srv ObligationRpcServer) {
	s.RegisterService(&ObligationRpc_ServiceDesc, srv)
}

func _ObligationRpc_CreateObligation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObligationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).CreateObligation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/CreateObligation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).CreateObligation(ctx, req.(*CreateObligationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObligationRpc_DescribeObligation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeObligationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).DescribeObligation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/DescribeObligation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).DescribeObligation(ctx, req.(*DescribeObligationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObligationRpc_ListObligations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObligationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).ListObligations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/ListObligations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).ListObligations(ctx, req.(*ListObligationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObligationRpc_RemoveObligation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveObligationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).RemoveObligation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/RemoveObligation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).RemoveObligation(ctx, req.(*RemoveObligationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObligationRpc_MultiCreateObligation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateObligationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).MultiCreateObligation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/MultiCreateObligation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).MultiCreateObligation(ctx, req.(*MultiCreateObligationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObligationRpc_UpdateObligation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObligationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObligationRpcServer).UpdateObligation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.obligation.api.ObligationRpc/UpdateObligation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObligationRpcServer).UpdateObligation(ctx, req.(*UpdateObligationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObligationRpc_ServiceDesc is the grpc.ServiceDesc for ObligationRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObligationRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.obligation.api.ObligationRpc",
	HandlerType: (*ObligationRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObligation",
			Handler:    _ObligationRpc_CreateObligation_Handler,
		},
		{
			MethodName: "DescribeObligation",
			Handler:    _ObligationRpc_DescribeObligation_Handler,
		},
		{
			MethodName: "ListObligations",
			Handler:    _ObligationRpc_ListObligations_Handler,
		},
		{
			MethodName: "RemoveObligation",
			Handler:    _ObligationRpc_RemoveObligation_Handler,
		},
		{
			MethodName: "MultiCreateObligation",
			Handler:    _ObligationRpc_MultiCreateObligation_Handler,
		},
		{
			MethodName: "UpdateObligation",
			Handler:    _ObligationRpc_UpdateObligation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ova-obligation-api.proto",
}
