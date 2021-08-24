// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_problem_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OcpProblemClient is the client API for OcpProblem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpProblemClient interface {
	CreateProblemV1(ctx context.Context, in *ProblemV1, opts ...grpc.CallOption) (*ResultSaveV1, error)
	DescribeProblemV1(ctx context.Context, in *ProblemQueryV1, opts ...grpc.CallOption) (*ProblemV1, error)
	ListProblemsV1(ctx context.Context, in *ProblemListQueryV1, opts ...grpc.CallOption) (*ProblemListV1, error)
	RemoveProblemV1(ctx context.Context, in *ProblemQueryV1, opts ...grpc.CallOption) (*ProblemResultV1, error)
}

type ocpProblemClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpProblemClient(cc grpc.ClientConnInterface) OcpProblemClient {
	return &ocpProblemClient{cc}
}

func (c *ocpProblemClient) CreateProblemV1(ctx context.Context, in *ProblemV1, opts ...grpc.CallOption) (*ResultSaveV1, error) {
	out := new(ResultSaveV1)
	err := c.cc.Invoke(ctx, "/ocp.problem.api.OcpProblem/CreateProblemV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpProblemClient) DescribeProblemV1(ctx context.Context, in *ProblemQueryV1, opts ...grpc.CallOption) (*ProblemV1, error) {
	out := new(ProblemV1)
	err := c.cc.Invoke(ctx, "/ocp.problem.api.OcpProblem/DescribeProblemV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpProblemClient) ListProblemsV1(ctx context.Context, in *ProblemListQueryV1, opts ...grpc.CallOption) (*ProblemListV1, error) {
	out := new(ProblemListV1)
	err := c.cc.Invoke(ctx, "/ocp.problem.api.OcpProblem/ListProblemsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpProblemClient) RemoveProblemV1(ctx context.Context, in *ProblemQueryV1, opts ...grpc.CallOption) (*ProblemResultV1, error) {
	out := new(ProblemResultV1)
	err := c.cc.Invoke(ctx, "/ocp.problem.api.OcpProblem/RemoveProblemV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpProblemServer is the server API for OcpProblem service.
// All implementations must embed UnimplementedOcpProblemServer
// for forward compatibility
type OcpProblemServer interface {
	CreateProblemV1(context.Context, *ProblemV1) (*ResultSaveV1, error)
	DescribeProblemV1(context.Context, *ProblemQueryV1) (*ProblemV1, error)
	ListProblemsV1(context.Context, *ProblemListQueryV1) (*ProblemListV1, error)
	RemoveProblemV1(context.Context, *ProblemQueryV1) (*ProblemResultV1, error)
	mustEmbedUnimplementedOcpProblemServer()
}

// UnimplementedOcpProblemServer must be embedded to have forward compatible implementations.
type UnimplementedOcpProblemServer struct {
}

func (UnimplementedOcpProblemServer) CreateProblemV1(context.Context, *ProblemV1) (*ResultSaveV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProblemV1 not implemented")
}
func (UnimplementedOcpProblemServer) DescribeProblemV1(context.Context, *ProblemQueryV1) (*ProblemV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProblemV1 not implemented")
}
func (UnimplementedOcpProblemServer) ListProblemsV1(context.Context, *ProblemListQueryV1) (*ProblemListV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProblemsV1 not implemented")
}
func (UnimplementedOcpProblemServer) RemoveProblemV1(context.Context, *ProblemQueryV1) (*ProblemResultV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProblemV1 not implemented")
}
func (UnimplementedOcpProblemServer) mustEmbedUnimplementedOcpProblemServer() {}

// UnsafeOcpProblemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpProblemServer will
// result in compilation errors.
type UnsafeOcpProblemServer interface {
	mustEmbedUnimplementedOcpProblemServer()
}

func RegisterOcpProblemServer(s grpc.ServiceRegistrar, srv OcpProblemServer) {
	s.RegisterService(&OcpProblem_ServiceDesc, srv)
}

func _OcpProblem_CreateProblemV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpProblemServer).CreateProblemV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.problem.api.OcpProblem/CreateProblemV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpProblemServer).CreateProblemV1(ctx, req.(*ProblemV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpProblem_DescribeProblemV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemQueryV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpProblemServer).DescribeProblemV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.problem.api.OcpProblem/DescribeProblemV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpProblemServer).DescribeProblemV1(ctx, req.(*ProblemQueryV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpProblem_ListProblemsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemListQueryV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpProblemServer).ListProblemsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.problem.api.OcpProblem/ListProblemsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpProblemServer).ListProblemsV1(ctx, req.(*ProblemListQueryV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpProblem_RemoveProblemV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemQueryV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpProblemServer).RemoveProblemV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.problem.api.OcpProblem/RemoveProblemV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpProblemServer).RemoveProblemV1(ctx, req.(*ProblemQueryV1))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpProblem_ServiceDesc is the grpc.ServiceDesc for OcpProblem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpProblem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.problem.api.OcpProblem",
	HandlerType: (*OcpProblemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProblemV1",
			Handler:    _OcpProblem_CreateProblemV1_Handler,
		},
		{
			MethodName: "DescribeProblemV1",
			Handler:    _OcpProblem_DescribeProblemV1_Handler,
		},
		{
			MethodName: "ListProblemsV1",
			Handler:    _OcpProblem_ListProblemsV1_Handler,
		},
		{
			MethodName: "RemoveProblemV1",
			Handler:    _OcpProblem_RemoveProblemV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-problem-api/ocp-problem-api.proto",
}