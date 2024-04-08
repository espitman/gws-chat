// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/message-service/service.proto

package message_service

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	V1CreateRoom(ctx context.Context, in *V1CreateRoomRequest, opts ...grpc.CallOption) (*V1CreateRoomResponse, error)
	V1AddMemberToRoom(ctx context.Context, in *V1AddMemberToRoomRequest, opts ...grpc.CallOption) (*V1AddMemberToRoomResponse, error)
	V1GetRoom(ctx context.Context, in *V1GetRoomRequest, opts ...grpc.CallOption) (*V1GetRoomResponse, error)
	V1AddMessage(ctx context.Context, in *V1AddMessageRequest, opts ...grpc.CallOption) (*V1AddMessageResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) V1CreateRoom(ctx context.Context, in *V1CreateRoomRequest, opts ...grpc.CallOption) (*V1CreateRoomResponse, error) {
	out := new(V1CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageService/V1CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) V1AddMemberToRoom(ctx context.Context, in *V1AddMemberToRoomRequest, opts ...grpc.CallOption) (*V1AddMemberToRoomResponse, error) {
	out := new(V1AddMemberToRoomResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageService/V1AddMemberToRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) V1GetRoom(ctx context.Context, in *V1GetRoomRequest, opts ...grpc.CallOption) (*V1GetRoomResponse, error) {
	out := new(V1GetRoomResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageService/V1GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) V1AddMessage(ctx context.Context, in *V1AddMessageRequest, opts ...grpc.CallOption) (*V1AddMessageResponse, error) {
	out := new(V1AddMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.MessageService/V1AddMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	V1CreateRoom(context.Context, *V1CreateRoomRequest) (*V1CreateRoomResponse, error)
	V1AddMemberToRoom(context.Context, *V1AddMemberToRoomRequest) (*V1AddMemberToRoomResponse, error)
	V1GetRoom(context.Context, *V1GetRoomRequest) (*V1GetRoomResponse, error)
	V1AddMessage(context.Context, *V1AddMessageRequest) (*V1AddMessageResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) V1CreateRoom(context.Context, *V1CreateRoomRequest) (*V1CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V1CreateRoom not implemented")
}
func (UnimplementedMessageServiceServer) V1AddMemberToRoom(context.Context, *V1AddMemberToRoomRequest) (*V1AddMemberToRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V1AddMemberToRoom not implemented")
}
func (UnimplementedMessageServiceServer) V1GetRoom(context.Context, *V1GetRoomRequest) (*V1GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V1GetRoom not implemented")
}
func (UnimplementedMessageServiceServer) V1AddMessage(context.Context, *V1AddMessageRequest) (*V1AddMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V1AddMessage not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_V1CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V1CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).V1CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageService/V1CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).V1CreateRoom(ctx, req.(*V1CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_V1AddMemberToRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V1AddMemberToRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).V1AddMemberToRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageService/V1AddMemberToRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).V1AddMemberToRoom(ctx, req.(*V1AddMemberToRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_V1GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V1GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).V1GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageService/V1GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).V1GetRoom(ctx, req.(*V1GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_V1AddMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(V1AddMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).V1AddMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageService/V1AddMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).V1AddMessage(ctx, req.(*V1AddMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "V1CreateRoom",
			Handler:    _MessageService_V1CreateRoom_Handler,
		},
		{
			MethodName: "V1AddMemberToRoom",
			Handler:    _MessageService_V1AddMemberToRoom_Handler,
		},
		{
			MethodName: "V1GetRoom",
			Handler:    _MessageService_V1GetRoom_Handler,
		},
		{
			MethodName: "V1AddMessage",
			Handler:    _MessageService_V1AddMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message-service/service.proto",
}
