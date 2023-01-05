// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/haru.proto

package proto

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

// Version1Client is the client API for Version1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Version1Client interface {
	CreateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	GetProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UpdateProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	CreateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error)
	UpdateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error)
	GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error)
	GetGameMulti(ctx context.Context, in *GameMultiRequest, opts ...grpc.CallOption) (*GameMultiReply, error)
	GetFilterdGames(ctx context.Context, in *FilterdGamesRequest, opts ...grpc.CallOption) (*FilterdGamesReply, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
	GetMyJoins(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
	GetGameJoins(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
	UpdateJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
	GetChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error)
	AddChatMessage(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*ChatReply, error)
	GetDataPlace(ctx context.Context, in *DataPlaceRequest, opts ...grpc.CallOption) (*DataPlaceReply, error)
}

type version1Client struct {
	cc grpc.ClientConnInterface
}

func NewVersion1Client(cc grpc.ClientConnInterface) Version1Client {
	return &version1Client{cc}
}

func (c *version1Client) CreateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/haru.version1/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) UpdateProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/haru.version1/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) CreateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error) {
	out := new(GameReply)
	err := c.cc.Invoke(ctx, "/haru.version1/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) UpdateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error) {
	out := new(GameReply)
	err := c.cc.Invoke(ctx, "/haru.version1/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error) {
	out := new(GameReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetGameMulti(ctx context.Context, in *GameMultiRequest, opts ...grpc.CallOption) (*GameMultiReply, error) {
	out := new(GameMultiReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetGameMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetFilterdGames(ctx context.Context, in *FilterdGamesRequest, opts ...grpc.CallOption) (*FilterdGamesReply, error) {
	out := new(FilterdGamesReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetFilterdGames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/haru.version1/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetMyJoins(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetMyJoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetGameJoins(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetGameJoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) UpdateJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/haru.version1/UpdateJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetChat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatReply, error) {
	out := new(ChatReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) AddChatMessage(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*ChatReply, error) {
	out := new(ChatReply)
	err := c.cc.Invoke(ctx, "/haru.version1/AddChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *version1Client) GetDataPlace(ctx context.Context, in *DataPlaceRequest, opts ...grpc.CallOption) (*DataPlaceReply, error) {
	out := new(DataPlaceReply)
	err := c.cc.Invoke(ctx, "/haru.version1/GetDataPlace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Version1Server is the server API for Version1 service.
// All implementations must embed UnimplementedVersion1Server
// for forward compatibility
type Version1Server interface {
	CreateAccount(context.Context, *AccountRequest) (*AccountReply, error)
	GetProfile(context.Context, *ProfileRequest) (*ProfileReply, error)
	UpdateProfile(context.Context, *ProfileRequest) (*ProfileReply, error)
	CreateGame(context.Context, *GameRequest) (*GameReply, error)
	UpdateGame(context.Context, *GameRequest) (*GameReply, error)
	GetGame(context.Context, *GameRequest) (*GameReply, error)
	GetGameMulti(context.Context, *GameMultiRequest) (*GameMultiReply, error)
	GetFilterdGames(context.Context, *FilterdGamesRequest) (*FilterdGamesReply, error)
	Join(context.Context, *JoinRequest) (*JoinReply, error)
	GetMyJoins(context.Context, *JoinRequest) (*JoinReply, error)
	GetGameJoins(context.Context, *JoinRequest) (*JoinReply, error)
	UpdateJoin(context.Context, *JoinRequest) (*JoinReply, error)
	GetChat(context.Context, *ChatRequest) (*ChatReply, error)
	AddChatMessage(context.Context, *ChatMessageRequest) (*ChatReply, error)
	GetDataPlace(context.Context, *DataPlaceRequest) (*DataPlaceReply, error)
	mustEmbedUnimplementedVersion1Server()
}

// UnimplementedVersion1Server must be embedded to have forward compatible implementations.
type UnimplementedVersion1Server struct {
}

func (UnimplementedVersion1Server) CreateAccount(context.Context, *AccountRequest) (*AccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedVersion1Server) GetProfile(context.Context, *ProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedVersion1Server) UpdateProfile(context.Context, *ProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedVersion1Server) CreateGame(context.Context, *GameRequest) (*GameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedVersion1Server) UpdateGame(context.Context, *GameRequest) (*GameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}
func (UnimplementedVersion1Server) GetGame(context.Context, *GameRequest) (*GameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedVersion1Server) GetGameMulti(context.Context, *GameMultiRequest) (*GameMultiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameMulti not implemented")
}
func (UnimplementedVersion1Server) GetFilterdGames(context.Context, *FilterdGamesRequest) (*FilterdGamesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterdGames not implemented")
}
func (UnimplementedVersion1Server) Join(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedVersion1Server) GetMyJoins(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyJoins not implemented")
}
func (UnimplementedVersion1Server) GetGameJoins(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameJoins not implemented")
}
func (UnimplementedVersion1Server) UpdateJoin(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJoin not implemented")
}
func (UnimplementedVersion1Server) GetChat(context.Context, *ChatRequest) (*ChatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedVersion1Server) AddChatMessage(context.Context, *ChatMessageRequest) (*ChatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChatMessage not implemented")
}
func (UnimplementedVersion1Server) GetDataPlace(context.Context, *DataPlaceRequest) (*DataPlaceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataPlace not implemented")
}
func (UnimplementedVersion1Server) mustEmbedUnimplementedVersion1Server() {}

// UnsafeVersion1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Version1Server will
// result in compilation errors.
type UnsafeVersion1Server interface {
	mustEmbedUnimplementedVersion1Server()
}

func RegisterVersion1Server(s grpc.ServiceRegistrar, srv Version1Server) {
	s.RegisterService(&Version1_ServiceDesc, srv)
}

func _Version1_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).CreateAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetProfile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).UpdateProfile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).CreateGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).UpdateGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetGameMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetGameMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetGameMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetGameMulti(ctx, req.(*GameMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetFilterdGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterdGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetFilterdGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetFilterdGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetFilterdGames(ctx, req.(*FilterdGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetMyJoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetMyJoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetMyJoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetMyJoins(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetGameJoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetGameJoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetGameJoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetGameJoins(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_UpdateJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).UpdateJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/UpdateJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).UpdateJoin(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetChat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_AddChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).AddChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/AddChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).AddChatMessage(ctx, req.(*ChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Version1_GetDataPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Version1Server).GetDataPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/haru.version1/GetDataPlace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Version1Server).GetDataPlace(ctx, req.(*DataPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Version1_ServiceDesc is the grpc.ServiceDesc for Version1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Version1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "haru.version1",
	HandlerType: (*Version1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Version1_CreateAccount_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Version1_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Version1_UpdateProfile_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Version1_CreateGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _Version1_UpdateGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _Version1_GetGame_Handler,
		},
		{
			MethodName: "GetGameMulti",
			Handler:    _Version1_GetGameMulti_Handler,
		},
		{
			MethodName: "GetFilterdGames",
			Handler:    _Version1_GetFilterdGames_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Version1_Join_Handler,
		},
		{
			MethodName: "GetMyJoins",
			Handler:    _Version1_GetMyJoins_Handler,
		},
		{
			MethodName: "GetGameJoins",
			Handler:    _Version1_GetGameJoins_Handler,
		},
		{
			MethodName: "UpdateJoin",
			Handler:    _Version1_UpdateJoin_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _Version1_GetChat_Handler,
		},
		{
			MethodName: "AddChatMessage",
			Handler:    _Version1_AddChatMessage_Handler,
		},
		{
			MethodName: "GetDataPlace",
			Handler:    _Version1_GetDataPlace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/haru.proto",
}
