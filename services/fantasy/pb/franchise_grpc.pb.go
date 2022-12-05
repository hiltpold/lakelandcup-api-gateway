// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: services/fantasy/pb/franchise.proto

package pb

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

// FantasyServiceClient is the client API for FantasyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FantasyServiceClient interface {
	CreateLeague(ctx context.Context, in *LeagueRequest, opts ...grpc.CallOption) (*LeagueResponse, error)
	CreateFranchise(ctx context.Context, in *FranchiseRequest, opts ...grpc.CallOption) (*FranchiseResponse, error)
	GetLeagueById(ctx context.Context, in *GetLeagueByIdRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	GetFranchiseById(ctx context.Context, in *GetFranchiseByIdRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type fantasyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFantasyServiceClient(cc grpc.ClientConnInterface) FantasyServiceClient {
	return &fantasyServiceClient{cc}
}

func (c *fantasyServiceClient) CreateLeague(ctx context.Context, in *LeagueRequest, opts ...grpc.CallOption) (*LeagueResponse, error) {
	out := new(LeagueResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/CreateLeague", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) CreateFranchise(ctx context.Context, in *FranchiseRequest, opts ...grpc.CallOption) (*FranchiseResponse, error) {
	out := new(FranchiseResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/CreateFranchise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetLeagueById(ctx context.Context, in *GetLeagueByIdRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetLeagueById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetFranchiseById(ctx context.Context, in *GetFranchiseByIdRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetFranchiseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FantasyServiceServer is the server API for FantasyService service.
// All implementations must embed UnimplementedFantasyServiceServer
// for forward compatibility
type FantasyServiceServer interface {
	CreateLeague(context.Context, *LeagueRequest) (*LeagueResponse, error)
	CreateFranchise(context.Context, *FranchiseRequest) (*FranchiseResponse, error)
	GetLeagueById(context.Context, *GetLeagueByIdRequest) (*QueryResponse, error)
	GetFranchiseById(context.Context, *GetFranchiseByIdRequest) (*QueryResponse, error)
	mustEmbedUnimplementedFantasyServiceServer()
}

// UnimplementedFantasyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFantasyServiceServer struct {
}

func (UnimplementedFantasyServiceServer) CreateLeague(context.Context, *LeagueRequest) (*LeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeague not implemented")
}
func (UnimplementedFantasyServiceServer) CreateFranchise(context.Context, *FranchiseRequest) (*FranchiseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFranchise not implemented")
}
func (UnimplementedFantasyServiceServer) GetLeagueById(context.Context, *GetLeagueByIdRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagueById not implemented")
}
func (UnimplementedFantasyServiceServer) GetFranchiseById(context.Context, *GetFranchiseByIdRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFranchiseById not implemented")
}
func (UnimplementedFantasyServiceServer) mustEmbedUnimplementedFantasyServiceServer() {}

// UnsafeFantasyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FantasyServiceServer will
// result in compilation errors.
type UnsafeFantasyServiceServer interface {
	mustEmbedUnimplementedFantasyServiceServer()
}

func RegisterFantasyServiceServer(s grpc.ServiceRegistrar, srv FantasyServiceServer) {
	s.RegisterService(&FantasyService_ServiceDesc, srv)
}

func _FantasyService_CreateLeague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).CreateLeague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/CreateLeague",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).CreateLeague(ctx, req.(*LeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_CreateFranchise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FranchiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).CreateFranchise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/CreateFranchise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).CreateFranchise(ctx, req.(*FranchiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetLeagueById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeagueByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetLeagueById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetLeagueById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetLeagueById(ctx, req.(*GetLeagueByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetFranchiseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFranchiseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetFranchiseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetFranchiseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetFranchiseById(ctx, req.(*GetFranchiseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FantasyService_ServiceDesc is the grpc.ServiceDesc for FantasyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FantasyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fantasy.FantasyService",
	HandlerType: (*FantasyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLeague",
			Handler:    _FantasyService_CreateLeague_Handler,
		},
		{
			MethodName: "CreateFranchise",
			Handler:    _FantasyService_CreateFranchise_Handler,
		},
		{
			MethodName: "GetLeagueById",
			Handler:    _FantasyService_GetLeagueById_Handler,
		},
		{
			MethodName: "GetFranchiseById",
			Handler:    _FantasyService_GetFranchiseById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/fantasy/pb/franchise.proto",
}