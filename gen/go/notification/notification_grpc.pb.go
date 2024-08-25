// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: notification/notification.proto

package notificationpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NotificationService_CreateNotification_FullMethodName = "/NotificationService/CreateNotification"
	NotificationService_GetNotification_FullMethodName    = "/NotificationService/GetNotification"
	NotificationService_AddNotification_FullMethodName    = "/NotificationService/AddNotification"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	CreateNotification(ctx context.Context, in *CreateNotificationReq, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationRes, error)
	AddNotification(ctx context.Context, in *AddNotificationReq, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationReq, opts ...grpc.CallOption) (*EmptyMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, NotificationService_CreateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationRes)
	err := c.cc.Invoke(ctx, NotificationService_GetNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) AddNotification(ctx context.Context, in *AddNotificationReq, opts ...grpc.CallOption) (*EmptyMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, NotificationService_AddNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	CreateNotification(context.Context, *CreateNotificationReq) (*EmptyMessage, error)
	GetNotification(context.Context, *GetNotificationReq) (*GetNotificationRes, error)
	AddNotification(context.Context, *AddNotificationReq) (*EmptyMessage, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) CreateNotification(context.Context, *CreateNotificationReq) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotification(context.Context, *GetNotificationReq) (*GetNotificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedNotificationServiceServer) AddNotification(context.Context, *AddNotificationReq) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNotification not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotification(ctx, req.(*CreateNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotification(ctx, req.(*GetNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_AddNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).AddNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_AddNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).AddNotification(ctx, req.(*AddNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationService_CreateNotification_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _NotificationService_GetNotification_Handler,
		},
		{
			MethodName: "AddNotification",
			Handler:    _NotificationService_AddNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/notification.proto",
}
