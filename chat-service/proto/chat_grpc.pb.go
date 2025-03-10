// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: chat.proto

package __

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
	ChatService_SaveMessage_FullMethodName         = "/chat.ChatService/SaveMessage"
	ChatService_GetRoomParticipants_FullMethodName = "/chat.ChatService/GetRoomParticipants"
	ChatService_GetRoomMessages_FullMethodName     = "/chat.ChatService/GetRoomMessages"
	ChatService_CreateRoom_FullMethodName          = "/chat.ChatService/CreateRoom"
	ChatService_AddRoomParticipant_FullMethodName  = "/chat.ChatService/AddRoomParticipant"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition
type ChatServiceClient interface {
	SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*SaveMessageResponse, error)
	GetRoomParticipants(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*RoomParticipantsResponse, error)
	GetRoomMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*PaginatedMessagesResponse, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	AddRoomParticipant(ctx context.Context, in *AddRoomParticipantRequest, opts ...grpc.CallOption) (*RoomParticipantsResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*SaveMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SaveMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRoomParticipants(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*RoomParticipantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomParticipantsResponse)
	err := c.cc.Invoke(ctx, ChatService_GetRoomParticipants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRoomMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*PaginatedMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaginatedMessagesResponse)
	err := c.cc.Invoke(ctx, ChatService_GetRoomMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddRoomParticipant(ctx context.Context, in *AddRoomParticipantRequest, opts ...grpc.CallOption) (*RoomParticipantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomParticipantsResponse)
	err := c.cc.Invoke(ctx, ChatService_AddRoomParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
//
// Service definition
type ChatServiceServer interface {
	SaveMessage(context.Context, *SaveMessageRequest) (*SaveMessageResponse, error)
	GetRoomParticipants(context.Context, *GetRoomRequest) (*RoomParticipantsResponse, error)
	GetRoomMessages(context.Context, *GetMessagesRequest) (*PaginatedMessagesResponse, error)
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	AddRoomParticipant(context.Context, *AddRoomParticipantRequest) (*RoomParticipantsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) SaveMessage(context.Context, *SaveMessageRequest) (*SaveMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (UnimplementedChatServiceServer) GetRoomParticipants(context.Context, *GetRoomRequest) (*RoomParticipantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomParticipants not implemented")
}
func (UnimplementedChatServiceServer) GetRoomMessages(context.Context, *GetMessagesRequest) (*PaginatedMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomMessages not implemented")
}
func (UnimplementedChatServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedChatServiceServer) AddRoomParticipant(context.Context, *AddRoomParticipantRequest) (*RoomParticipantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoomParticipant not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SaveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SaveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SaveMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SaveMessage(ctx, req.(*SaveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRoomParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetRoomParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetRoomParticipants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetRoomParticipants(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRoomMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetRoomMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetRoomMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetRoomMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddRoomParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoomParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddRoomParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddRoomParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddRoomParticipant(ctx, req.(*AddRoomParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMessage",
			Handler:    _ChatService_SaveMessage_Handler,
		},
		{
			MethodName: "GetRoomParticipants",
			Handler:    _ChatService_GetRoomParticipants_Handler,
		},
		{
			MethodName: "GetRoomMessages",
			Handler:    _ChatService_GetRoomMessages_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _ChatService_CreateRoom_Handler,
		},
		{
			MethodName: "AddRoomParticipant",
			Handler:    _ChatService_AddRoomParticipant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
