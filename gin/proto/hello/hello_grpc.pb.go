// 协议类型

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: hello.proto

package hello

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

const (
	HelloDB_SayHello_FullMethodName   = "/HelloDB/SayHello"
	HelloDB_SayGoodbye_FullMethodName = "/HelloDB/SayGoodbye"
)

// HelloDBClient is the client API for HelloDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloDBClient interface {
	SayHello(ctx context.Context, in *HelloDBRequest, opts ...grpc.CallOption) (*HelloDBResponse, error)
	SayGoodbye(ctx context.Context, in *GoodByeDBRequest, opts ...grpc.CallOption) (*GoodByeDBResponse, error)
}

type helloDBClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloDBClient(cc grpc.ClientConnInterface) HelloDBClient {
	return &helloDBClient{cc}
}

func (c *helloDBClient) SayHello(ctx context.Context, in *HelloDBRequest, opts ...grpc.CallOption) (*HelloDBResponse, error) {
	out := new(HelloDBResponse)
	err := c.cc.Invoke(ctx, HelloDB_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDBClient) SayGoodbye(ctx context.Context, in *GoodByeDBRequest, opts ...grpc.CallOption) (*GoodByeDBResponse, error) {
	out := new(GoodByeDBResponse)
	err := c.cc.Invoke(ctx, HelloDB_SayGoodbye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloDBServer is the server API for HelloDB service.
// All implementations must embed UnimplementedHelloDBServer
// for forward compatibility
type HelloDBServer interface {
	SayHello(context.Context, *HelloDBRequest) (*HelloDBResponse, error)
	SayGoodbye(context.Context, *GoodByeDBRequest) (*GoodByeDBResponse, error)
	mustEmbedUnimplementedHelloDBServer()
}

// UnimplementedHelloDBServer must be embedded to have forward compatible implementations.
type UnimplementedHelloDBServer struct {
}

func (UnimplementedHelloDBServer) SayHello(context.Context, *HelloDBRequest) (*HelloDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloDBServer) SayGoodbye(context.Context, *GoodByeDBRequest) (*GoodByeDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}
func (UnimplementedHelloDBServer) mustEmbedUnimplementedHelloDBServer() {}

// UnsafeHelloDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloDBServer will
// result in compilation errors.
type UnsafeHelloDBServer interface {
	mustEmbedUnimplementedHelloDBServer()
}

func RegisterHelloDBServer(s grpc.ServiceRegistrar, srv HelloDBServer) {
	s.RegisterService(&HelloDB_ServiceDesc, srv)
}

func _HelloDB_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDBServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloDB_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDBServer).SayHello(ctx, req.(*HelloDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDB_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodByeDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDBServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloDB_SayGoodbye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDBServer).SayGoodbye(ctx, req.(*GoodByeDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloDB_ServiceDesc is the grpc.ServiceDesc for HelloDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HelloDB",
	HandlerType: (*HelloDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloDB_SayHello_Handler,
		},
		{
			MethodName: "SayGoodbye",
			Handler:    _HelloDB_SayGoodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

const (
	HelloHttp_SayHello_FullMethodName   = "/HelloHttp/SayHello"
	HelloHttp_SayGoodbye_FullMethodName = "/HelloHttp/SayGoodbye"
)

// HelloHttpClient is the client API for HelloHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloHttpClient interface {
	SayHello(ctx context.Context, in *HelloHttpRequest, opts ...grpc.CallOption) (*HelloHttpResponse, error)
	SayGoodbye(ctx context.Context, in *GoodByeHttpRequest, opts ...grpc.CallOption) (*GoodByeHttpResponse, error)
}

type helloHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloHttpClient(cc grpc.ClientConnInterface) HelloHttpClient {
	return &helloHttpClient{cc}
}

func (c *helloHttpClient) SayHello(ctx context.Context, in *HelloHttpRequest, opts ...grpc.CallOption) (*HelloHttpResponse, error) {
	out := new(HelloHttpResponse)
	err := c.cc.Invoke(ctx, HelloHttp_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloHttpClient) SayGoodbye(ctx context.Context, in *GoodByeHttpRequest, opts ...grpc.CallOption) (*GoodByeHttpResponse, error) {
	out := new(GoodByeHttpResponse)
	err := c.cc.Invoke(ctx, HelloHttp_SayGoodbye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloHttpServer is the server API for HelloHttp service.
// All implementations must embed UnimplementedHelloHttpServer
// for forward compatibility
type HelloHttpServer interface {
	SayHello(context.Context, *HelloHttpRequest) (*HelloHttpResponse, error)
	SayGoodbye(context.Context, *GoodByeHttpRequest) (*GoodByeHttpResponse, error)
	mustEmbedUnimplementedHelloHttpServer()
}

// UnimplementedHelloHttpServer must be embedded to have forward compatible implementations.
type UnimplementedHelloHttpServer struct {
}

func (UnimplementedHelloHttpServer) SayHello(context.Context, *HelloHttpRequest) (*HelloHttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloHttpServer) SayGoodbye(context.Context, *GoodByeHttpRequest) (*GoodByeHttpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}
func (UnimplementedHelloHttpServer) mustEmbedUnimplementedHelloHttpServer() {}

// UnsafeHelloHttpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloHttpServer will
// result in compilation errors.
type UnsafeHelloHttpServer interface {
	mustEmbedUnimplementedHelloHttpServer()
}

func RegisterHelloHttpServer(s grpc.ServiceRegistrar, srv HelloHttpServer) {
	s.RegisterService(&HelloHttp_ServiceDesc, srv)
}

func _HelloHttp_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloHttpServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloHttp_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloHttpServer).SayHello(ctx, req.(*HelloHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloHttp_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodByeHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloHttpServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloHttp_SayGoodbye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloHttpServer).SayGoodbye(ctx, req.(*GoodByeHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloHttp_ServiceDesc is the grpc.ServiceDesc for HelloHttp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloHttp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HelloHttp",
	HandlerType: (*HelloHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloHttp_SayHello_Handler,
		},
		{
			MethodName: "SayGoodbye",
			Handler:    _HelloHttp_SayGoodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
