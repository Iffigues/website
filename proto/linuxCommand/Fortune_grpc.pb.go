// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/Fortune.proto

package linuxCommand

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FortuneServiceClient is the client API for FortuneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FortuneServiceClient interface {
	GetFortune(ctx context.Context, in *Fortune, opts ...grpc.CallOption) (*Responses, error)
	GetFortuneFile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Responses, error)
}

type fortuneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFortuneServiceClient(cc grpc.ClientConnInterface) FortuneServiceClient {
	return &fortuneServiceClient{cc}
}

func (c *fortuneServiceClient) GetFortune(ctx context.Context, in *Fortune, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.FortuneService/GetFortune", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fortuneServiceClient) GetFortuneFile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.FortuneService/GetFortuneFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FortuneServiceServer is the server API for FortuneService service.
// All implementations must embed UnimplementedFortuneServiceServer
// for forward compatibility
type FortuneServiceServer interface {
	GetFortune(context.Context, *Fortune) (*Responses, error)
	GetFortuneFile(context.Context, *empty.Empty) (*Responses, error)
	mustEmbedUnimplementedFortuneServiceServer()
}

// UnimplementedFortuneServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFortuneServiceServer struct {
}

func (UnimplementedFortuneServiceServer) GetFortune(context.Context, *Fortune) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFortune not implemented")
}
func (UnimplementedFortuneServiceServer) GetFortuneFile(context.Context, *empty.Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFortuneFile not implemented")
}
func (UnimplementedFortuneServiceServer) mustEmbedUnimplementedFortuneServiceServer() {}

// UnsafeFortuneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FortuneServiceServer will
// result in compilation errors.
type UnsafeFortuneServiceServer interface {
	mustEmbedUnimplementedFortuneServiceServer()
}

func RegisterFortuneServiceServer(s grpc.ServiceRegistrar, srv FortuneServiceServer) {
	s.RegisterService(&FortuneService_ServiceDesc, srv)
}

func _FortuneService_GetFortune_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fortune)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FortuneServiceServer).GetFortune(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.FortuneService/GetFortune",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FortuneServiceServer).GetFortune(ctx, req.(*Fortune))
	}
	return interceptor(ctx, in, info, handler)
}

func _FortuneService_GetFortuneFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FortuneServiceServer).GetFortuneFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.FortuneService/GetFortuneFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FortuneServiceServer).GetFortuneFile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FortuneService_ServiceDesc is the grpc.ServiceDesc for FortuneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FortuneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linuxCommand.FortuneService",
	HandlerType: (*FortuneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFortune",
			Handler:    _FortuneService_GetFortune_Handler,
		},
		{
			MethodName: "GetFortuneFile",
			Handler:    _FortuneService_GetFortuneFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fortune.proto",
}
