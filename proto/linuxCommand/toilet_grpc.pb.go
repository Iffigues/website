// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/toilet.proto

package linuxCommand

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

// ToiletServiceClient is the client API for ToiletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToiletServiceClient interface {
	GetToilet(ctx context.Context, in *Toilet, opts ...grpc.CallOption) (*Responses, error)
	GetToiletFFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
	GetToiletEFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
	GetToiletFFFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
}

type toiletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToiletServiceClient(cc grpc.ClientConnInterface) ToiletServiceClient {
	return &toiletServiceClient{cc}
}

func (c *toiletServiceClient) GetToilet(ctx context.Context, in *Toilet, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.ToiletService/GetToilet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toiletServiceClient) GetToiletFFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.ToiletService/GetToiletFFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toiletServiceClient) GetToiletEFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.ToiletService/GetToiletEFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toiletServiceClient) GetToiletFFFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.ToiletService/GetToiletFFFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToiletServiceServer is the server API for ToiletService service.
// All implementations must embed UnimplementedToiletServiceServer
// for forward compatibility
type ToiletServiceServer interface {
	GetToilet(context.Context, *Toilet) (*Responses, error)
	GetToiletFFile(context.Context, *Empty) (*Responses, error)
	GetToiletEFile(context.Context, *Empty) (*Responses, error)
	GetToiletFFFile(context.Context, *Empty) (*Responses, error)
	mustEmbedUnimplementedToiletServiceServer()
}

// UnimplementedToiletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToiletServiceServer struct {
}

func (UnimplementedToiletServiceServer) GetToilet(context.Context, *Toilet) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToilet not implemented")
}
func (UnimplementedToiletServiceServer) GetToiletFFile(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToiletFFile not implemented")
}
func (UnimplementedToiletServiceServer) GetToiletEFile(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToiletEFile not implemented")
}
func (UnimplementedToiletServiceServer) GetToiletFFFile(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToiletFFFile not implemented")
}
func (UnimplementedToiletServiceServer) mustEmbedUnimplementedToiletServiceServer() {}

// UnsafeToiletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToiletServiceServer will
// result in compilation errors.
type UnsafeToiletServiceServer interface {
	mustEmbedUnimplementedToiletServiceServer()
}

func RegisterToiletServiceServer(s grpc.ServiceRegistrar, srv ToiletServiceServer) {
	s.RegisterService(&ToiletService_ServiceDesc, srv)
}

func _ToiletService_GetToilet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Toilet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToiletServiceServer).GetToilet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.ToiletService/GetToilet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToiletServiceServer).GetToilet(ctx, req.(*Toilet))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToiletService_GetToiletFFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToiletServiceServer).GetToiletFFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.ToiletService/GetToiletFFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToiletServiceServer).GetToiletFFile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToiletService_GetToiletEFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToiletServiceServer).GetToiletEFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.ToiletService/GetToiletEFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToiletServiceServer).GetToiletEFile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToiletService_GetToiletFFFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToiletServiceServer).GetToiletFFFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.ToiletService/GetToiletFFFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToiletServiceServer).GetToiletFFFile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ToiletService_ServiceDesc is the grpc.ServiceDesc for ToiletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToiletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linuxCommand.ToiletService",
	HandlerType: (*ToiletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToilet",
			Handler:    _ToiletService_GetToilet_Handler,
		},
		{
			MethodName: "GetToiletFFile",
			Handler:    _ToiletService_GetToiletFFile_Handler,
		},
		{
			MethodName: "GetToiletEFile",
			Handler:    _ToiletService_GetToiletEFile_Handler,
		},
		{
			MethodName: "GetToiletFFFile",
			Handler:    _ToiletService_GetToiletFFFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/toilet.proto",
}
