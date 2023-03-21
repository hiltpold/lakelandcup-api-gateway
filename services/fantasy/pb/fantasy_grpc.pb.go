// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: services/fantasy/pb/fantasy.proto

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
	GetLeagues(ctx context.Context, in *GetLeaguesRequest, opts ...grpc.CallOption) (*GetLeaguesResponse, error)
	GetLeague(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueResponse, error)
	UpdateLeague(ctx context.Context, in *LeagueUpdateRequest, opts ...grpc.CallOption) (*LeagueResponse, error)
	GetLeagueFranchises(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueFranchisesResponse, error)
	CreateFranchise(ctx context.Context, in *FranchiseRequest, opts ...grpc.CallOption) (*FranchiseResponse, error)
	GetFranchise(ctx context.Context, in *GetFranchiseRequest, opts ...grpc.CallOption) (*GetFranchiseResponse, error)
	CreateProspect(ctx context.Context, in *CreateProspectRequest, opts ...grpc.CallOption) (*CreateProspectResponse, error)
	CreateProspectsBulk(ctx context.Context, in *CreateProspectsBulkRequest, opts ...grpc.CallOption) (*CreateProspectsBulkResponse, error)
	TextSearchProspects(ctx context.Context, in *TextSearchRequest, opts ...grpc.CallOption) (*TextSearchProspectsResponse, error)
	GetPicks(ctx context.Context, in *GetFranchiseRequest, opts ...grpc.CallOption) (*GetPicksResponse, error)
	Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	CreateOrUpdatePicks(ctx context.Context, in *CreateOrUpdatePicksRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
	//
	DraftProspect(ctx context.Context, in *DraftProspectRequest, opts ...grpc.CallOption) (*DraftProspectResponse, error)
	UndraftProspect(ctx context.Context, in *DraftProspectRequest, opts ...grpc.CallOption) (*DraftProspectResponse, error)
	GetLeagueFranchisePairs(ctx context.Context, in *GetLeagueFranchisePairsRequest, opts ...grpc.CallOption) (*GetLeagueFranchisePairsResponse, error)
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

func (c *fantasyServiceClient) GetLeagues(ctx context.Context, in *GetLeaguesRequest, opts ...grpc.CallOption) (*GetLeaguesResponse, error) {
	out := new(GetLeaguesResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetLeagues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetLeague(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueResponse, error) {
	out := new(GetLeagueResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetLeague", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) UpdateLeague(ctx context.Context, in *LeagueUpdateRequest, opts ...grpc.CallOption) (*LeagueResponse, error) {
	out := new(LeagueResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/UpdateLeague", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetLeagueFranchises(ctx context.Context, in *GetLeagueRequest, opts ...grpc.CallOption) (*GetLeagueFranchisesResponse, error) {
	out := new(GetLeagueFranchisesResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetLeagueFranchises", in, out, opts...)
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

func (c *fantasyServiceClient) GetFranchise(ctx context.Context, in *GetFranchiseRequest, opts ...grpc.CallOption) (*GetFranchiseResponse, error) {
	out := new(GetFranchiseResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetFranchise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) CreateProspect(ctx context.Context, in *CreateProspectRequest, opts ...grpc.CallOption) (*CreateProspectResponse, error) {
	out := new(CreateProspectResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/CreateProspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) CreateProspectsBulk(ctx context.Context, in *CreateProspectsBulkRequest, opts ...grpc.CallOption) (*CreateProspectsBulkResponse, error) {
	out := new(CreateProspectsBulkResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/CreateProspectsBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) TextSearchProspects(ctx context.Context, in *TextSearchRequest, opts ...grpc.CallOption) (*TextSearchProspectsResponse, error) {
	out := new(TextSearchProspectsResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/TextSearchProspects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetPicks(ctx context.Context, in *GetFranchiseRequest, opts ...grpc.CallOption) (*GetPicksResponse, error) {
	out := new(GetPicksResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetPicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/Trade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) CreateOrUpdatePicks(ctx context.Context, in *CreateOrUpdatePicksRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/CreateOrUpdatePicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) DraftProspect(ctx context.Context, in *DraftProspectRequest, opts ...grpc.CallOption) (*DraftProspectResponse, error) {
	out := new(DraftProspectResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/DraftProspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) UndraftProspect(ctx context.Context, in *DraftProspectRequest, opts ...grpc.CallOption) (*DraftProspectResponse, error) {
	out := new(DraftProspectResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/UndraftProspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fantasyServiceClient) GetLeagueFranchisePairs(ctx context.Context, in *GetLeagueFranchisePairsRequest, opts ...grpc.CallOption) (*GetLeagueFranchisePairsResponse, error) {
	out := new(GetLeagueFranchisePairsResponse)
	err := c.cc.Invoke(ctx, "/fantasy.FantasyService/GetLeagueFranchisePairs", in, out, opts...)
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
	GetLeagues(context.Context, *GetLeaguesRequest) (*GetLeaguesResponse, error)
	GetLeague(context.Context, *GetLeagueRequest) (*GetLeagueResponse, error)
	UpdateLeague(context.Context, *LeagueUpdateRequest) (*LeagueResponse, error)
	GetLeagueFranchises(context.Context, *GetLeagueRequest) (*GetLeagueFranchisesResponse, error)
	CreateFranchise(context.Context, *FranchiseRequest) (*FranchiseResponse, error)
	GetFranchise(context.Context, *GetFranchiseRequest) (*GetFranchiseResponse, error)
	CreateProspect(context.Context, *CreateProspectRequest) (*CreateProspectResponse, error)
	CreateProspectsBulk(context.Context, *CreateProspectsBulkRequest) (*CreateProspectsBulkResponse, error)
	TextSearchProspects(context.Context, *TextSearchRequest) (*TextSearchProspectsResponse, error)
	GetPicks(context.Context, *GetFranchiseRequest) (*GetPicksResponse, error)
	Trade(context.Context, *TradeRequest) (*DefaultResponse, error)
	CreateOrUpdatePicks(context.Context, *CreateOrUpdatePicksRequest) (*DefaultResponse, error)
	//
	DraftProspect(context.Context, *DraftProspectRequest) (*DraftProspectResponse, error)
	UndraftProspect(context.Context, *DraftProspectRequest) (*DraftProspectResponse, error)
	GetLeagueFranchisePairs(context.Context, *GetLeagueFranchisePairsRequest) (*GetLeagueFranchisePairsResponse, error)
	mustEmbedUnimplementedFantasyServiceServer()
}

// UnimplementedFantasyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFantasyServiceServer struct {
}

func (UnimplementedFantasyServiceServer) CreateLeague(context.Context, *LeagueRequest) (*LeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLeague not implemented")
}
func (UnimplementedFantasyServiceServer) GetLeagues(context.Context, *GetLeaguesRequest) (*GetLeaguesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagues not implemented")
}
func (UnimplementedFantasyServiceServer) GetLeague(context.Context, *GetLeagueRequest) (*GetLeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeague not implemented")
}
func (UnimplementedFantasyServiceServer) UpdateLeague(context.Context, *LeagueUpdateRequest) (*LeagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeague not implemented")
}
func (UnimplementedFantasyServiceServer) GetLeagueFranchises(context.Context, *GetLeagueRequest) (*GetLeagueFranchisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagueFranchises not implemented")
}
func (UnimplementedFantasyServiceServer) CreateFranchise(context.Context, *FranchiseRequest) (*FranchiseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFranchise not implemented")
}
func (UnimplementedFantasyServiceServer) GetFranchise(context.Context, *GetFranchiseRequest) (*GetFranchiseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFranchise not implemented")
}
func (UnimplementedFantasyServiceServer) CreateProspect(context.Context, *CreateProspectRequest) (*CreateProspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProspect not implemented")
}
func (UnimplementedFantasyServiceServer) CreateProspectsBulk(context.Context, *CreateProspectsBulkRequest) (*CreateProspectsBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProspectsBulk not implemented")
}
func (UnimplementedFantasyServiceServer) TextSearchProspects(context.Context, *TextSearchRequest) (*TextSearchProspectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextSearchProspects not implemented")
}
func (UnimplementedFantasyServiceServer) GetPicks(context.Context, *GetFranchiseRequest) (*GetPicksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPicks not implemented")
}
func (UnimplementedFantasyServiceServer) Trade(context.Context, *TradeRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trade not implemented")
}
func (UnimplementedFantasyServiceServer) CreateOrUpdatePicks(context.Context, *CreateOrUpdatePicksRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdatePicks not implemented")
}
func (UnimplementedFantasyServiceServer) DraftProspect(context.Context, *DraftProspectRequest) (*DraftProspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DraftProspect not implemented")
}
func (UnimplementedFantasyServiceServer) UndraftProspect(context.Context, *DraftProspectRequest) (*DraftProspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndraftProspect not implemented")
}
func (UnimplementedFantasyServiceServer) GetLeagueFranchisePairs(context.Context, *GetLeagueFranchisePairsRequest) (*GetLeagueFranchisePairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagueFranchisePairs not implemented")
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

func _FantasyService_GetLeagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaguesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetLeagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetLeagues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetLeagues(ctx, req.(*GetLeaguesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetLeague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetLeague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetLeague",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetLeague(ctx, req.(*GetLeagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_UpdateLeague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeagueUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).UpdateLeague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/UpdateLeague",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).UpdateLeague(ctx, req.(*LeagueUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetLeagueFranchises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetLeagueFranchises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetLeagueFranchises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetLeagueFranchises(ctx, req.(*GetLeagueRequest))
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

func _FantasyService_GetFranchise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFranchiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetFranchise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetFranchise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetFranchise(ctx, req.(*GetFranchiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_CreateProspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).CreateProspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/CreateProspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).CreateProspect(ctx, req.(*CreateProspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_CreateProspectsBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProspectsBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).CreateProspectsBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/CreateProspectsBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).CreateProspectsBulk(ctx, req.(*CreateProspectsBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_TextSearchProspects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).TextSearchProspects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/TextSearchProspects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).TextSearchProspects(ctx, req.(*TextSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetPicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFranchiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetPicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetPicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetPicks(ctx, req.(*GetFranchiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_Trade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).Trade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/Trade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).Trade(ctx, req.(*TradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_CreateOrUpdatePicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdatePicksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).CreateOrUpdatePicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/CreateOrUpdatePicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).CreateOrUpdatePicks(ctx, req.(*CreateOrUpdatePicksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_DraftProspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftProspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).DraftProspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/DraftProspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).DraftProspect(ctx, req.(*DraftProspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_UndraftProspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftProspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).UndraftProspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/UndraftProspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).UndraftProspect(ctx, req.(*DraftProspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FantasyService_GetLeagueFranchisePairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeagueFranchisePairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FantasyServiceServer).GetLeagueFranchisePairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fantasy.FantasyService/GetLeagueFranchisePairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FantasyServiceServer).GetLeagueFranchisePairs(ctx, req.(*GetLeagueFranchisePairsRequest))
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
			MethodName: "GetLeagues",
			Handler:    _FantasyService_GetLeagues_Handler,
		},
		{
			MethodName: "GetLeague",
			Handler:    _FantasyService_GetLeague_Handler,
		},
		{
			MethodName: "UpdateLeague",
			Handler:    _FantasyService_UpdateLeague_Handler,
		},
		{
			MethodName: "GetLeagueFranchises",
			Handler:    _FantasyService_GetLeagueFranchises_Handler,
		},
		{
			MethodName: "CreateFranchise",
			Handler:    _FantasyService_CreateFranchise_Handler,
		},
		{
			MethodName: "GetFranchise",
			Handler:    _FantasyService_GetFranchise_Handler,
		},
		{
			MethodName: "CreateProspect",
			Handler:    _FantasyService_CreateProspect_Handler,
		},
		{
			MethodName: "CreateProspectsBulk",
			Handler:    _FantasyService_CreateProspectsBulk_Handler,
		},
		{
			MethodName: "TextSearchProspects",
			Handler:    _FantasyService_TextSearchProspects_Handler,
		},
		{
			MethodName: "GetPicks",
			Handler:    _FantasyService_GetPicks_Handler,
		},
		{
			MethodName: "Trade",
			Handler:    _FantasyService_Trade_Handler,
		},
		{
			MethodName: "CreateOrUpdatePicks",
			Handler:    _FantasyService_CreateOrUpdatePicks_Handler,
		},
		{
			MethodName: "DraftProspect",
			Handler:    _FantasyService_DraftProspect_Handler,
		},
		{
			MethodName: "UndraftProspect",
			Handler:    _FantasyService_UndraftProspect_Handler,
		},
		{
			MethodName: "GetLeagueFranchisePairs",
			Handler:    _FantasyService_GetLeagueFranchisePairs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/fantasy/pb/fantasy.proto",
}
