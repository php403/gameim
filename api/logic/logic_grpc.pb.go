// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/logic/logic.proto

package logic

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

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicClient interface {
	OnAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error)
	OnConnect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	OnMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	OnClose(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type logicClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicClient(cc grpc.ClientConnInterface) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) OnAuth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/logic.Logic/OnAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) OnConnect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logic.Logic/OnConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) OnMessage(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logic.Logic/OnMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) OnClose(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logic.Logic/OnClose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
// All implementations must embed UnimplementedLogicServer
// for forward compatibility
type LogicServer interface {
	OnAuth(context.Context, *AuthReq) (*AuthReply, error)
	OnConnect(context.Context, *ConnectReq) (*emptypb.Empty, error)
	OnMessage(context.Context, *MessageReq) (*emptypb.Empty, error)
	OnClose(context.Context, *CloseReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedLogicServer()
}

// UnimplementedLogicServer must be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (UnimplementedLogicServer) OnAuth(context.Context, *AuthReq) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnAuth not implemented")
}
func (UnimplementedLogicServer) OnConnect(context.Context, *ConnectReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnConnect not implemented")
}
func (UnimplementedLogicServer) OnMessage(context.Context, *MessageReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessage not implemented")
}
func (UnimplementedLogicServer) OnClose(context.Context, *CloseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnClose not implemented")
}
func (UnimplementedLogicServer) mustEmbedUnimplementedLogicServer() {}

// UnsafeLogicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicServer will
// result in compilation errors.
type UnsafeLogicServer interface {
	mustEmbedUnimplementedLogicServer()
}

func RegisterLogicServer(s grpc.ServiceRegistrar, srv LogicServer) {
	s.RegisterService(&Logic_ServiceDesc, srv)
}

func _Logic_OnAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).OnAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/OnAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).OnAuth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_OnConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).OnConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/OnConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).OnConnect(ctx, req.(*ConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_OnMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).OnMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/OnMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).OnMessage(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_OnClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).OnClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/OnClose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).OnClose(ctx, req.(*CloseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Logic_ServiceDesc is the grpc.ServiceDesc for Logic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logic.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnAuth",
			Handler:    _Logic_OnAuth_Handler,
		},
		{
			MethodName: "OnConnect",
			Handler:    _Logic_OnConnect_Handler,
		},
		{
			MethodName: "OnMessage",
			Handler:    _Logic_OnMessage_Handler,
		},
		{
			MethodName: "OnClose",
			Handler:    _Logic_OnClose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/logic/logic.proto",
}