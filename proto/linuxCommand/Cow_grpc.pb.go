// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/Cow.proto

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

// CowServiceClient is the client API for CowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CowServiceClient interface {
	GetCow(ctx context.Context, in *Cow, opts ...grpc.CallOption) (*Responses, error)
	GetCowFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
}

type cowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCowServiceClient(cc grpc.ClientConnInterface) CowServiceClient {
	return &cowServiceClient{cc}
}

func (c *cowServiceClient) GetCow(ctx context.Context, in *Cow, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.CowService/GetCow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cowServiceClient) GetCowFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.CowService/GetCowFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CowServiceServer is the server API for CowService service.
// All implementations must embed UnimplementedCowServiceServer
// for forward compatibility
type CowServiceServer interface {
	GetCow(context.Context, *Cow) (*Responses, error)
	GetCowFile(context.Context, *Empty) (*Responses, error)
	mustEmbedUnimplementedCowServiceServer()
}

// UnimplementedCowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCowServiceServer struct {
}

func (UnimplementedCowServiceServer) GetCow(context.Context, *Cow) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCow not implemented")
}
func (UnimplementedCowServiceServer) GetCowFile(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCowFile not implemented")
}
func (UnimplementedCowServiceServer) mustEmbedUnimplementedCowServiceServer() {}

// UnsafeCowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CowServiceServer will
// result in compilation errors.
type UnsafeCowServiceServer interface {
	mustEmbedUnimplementedCowServiceServer()
}

func RegisterCowServiceServer(s grpc.ServiceRegistrar, srv CowServiceServer) {
	s.RegisterService(&CowService_ServiceDesc, srv)
}

func _CowService_GetCow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CowServiceServer).GetCow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.CowService/GetCow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CowServiceServer).GetCow(ctx, req.(*Cow))
	}
	return interceptor(ctx, in, info, handler)
}

func _CowService_GetCowFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CowServiceServer).GetCowFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.CowService/GetCowFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CowServiceServer).GetCowFile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CowService_ServiceDesc is the grpc.ServiceDesc for CowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linuxCommand.CowService",
	HandlerType: (*CowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCow",
			Handler:    _CowService_GetCow_Handler,
		},
		{
			MethodName: "GetCowFile",
			Handler:    _CowService_GetCowFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Cow.proto",
}