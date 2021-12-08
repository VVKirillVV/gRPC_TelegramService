// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package telegram

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

// TelegramClient is the client API for Telegram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type telegramClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramClient(cc grpc.ClientConnInterface) TelegramClient {
	return &telegramClient{cc}
}

func (c *telegramClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/api.Telegram/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramServer is the server API for Telegram service.
// All implementations must embed UnimplementedTelegramServer
// for forward compatibility
type TelegramServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	mustEmbedUnimplementedTelegramServer()
}

// UnimplementedTelegramServer must be embedded to have forward compatible implementations.
type UnimplementedTelegramServer struct {
}

func (UnimplementedTelegramServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTelegramServer) mustEmbedUnimplementedTelegramServer() {}

// UnsafeTelegramServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramServer will
// result in compilation errors.
type UnsafeTelegramServer interface {
	mustEmbedUnimplementedTelegramServer()
}

func RegisterTelegramServer(s grpc.ServiceRegistrar, srv TelegramServer) {
	s.RegisterService(&Telegram_ServiceDesc, srv)
}

func _Telegram_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Telegram/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Telegram_ServiceDesc is the grpc.ServiceDesc for Telegram service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Telegram_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Telegram",
	HandlerType: (*TelegramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Telegram_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adder.proto",
}
