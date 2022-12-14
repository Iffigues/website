// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/figlet.proto

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

// FigletServiceClient is the client API for FigletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FigletServiceClient interface {
	GetFiglet(ctx context.Context, in *Figlet, opts ...grpc.CallOption) (*Responses, error)
	GetFigletFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
}

type figletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFigletServiceClient(cc grpc.ClientConnInterface) FigletServiceClient {
	return &figletServiceClient{cc}
}

func (c *figletServiceClient) GetFiglet(ctx context.Context, in *Figlet, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.FigletService/GetFiglet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *figletServiceClient) GetFigletFile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.FigletService/GetFigletFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FigletServiceServer is the server API for FigletService service.
// All implementations must embed UnimplementedFigletServiceServer
// for forward compatibility
type FigletServiceServer interface {
	GetFiglet(context.Context, *Figlet) (*Responses, error)
	GetFigletFile(context.Context, *Empty) (*Responses, error)
	mustEmbedUnimplementedFigletServiceServer()
}

// UnimplementedFigletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFigletServiceServer struct {
}

func (UnimplementedFigletServiceServer) GetFiglet(context.Context, *Figlet) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiglet not implemented")
}
func (UnimplementedFigletServiceServer) GetFigletFile(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFigletFile not implemented")
}
func (UnimplementedFigletServiceServer) mustEmbedUnimplementedFigletServiceServer() {}

// UnsafeFigletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FigletServiceServer will
// result in compilation errors.
type UnsafeFigletServiceServer interface {
	mustEmbedUnimplementedFigletServiceServer()
}

func RegisterFigletServiceServer(s grpc.ServiceRegistrar, srv FigletServiceServer) {
	s.RegisterService(&FigletService_ServiceDesc, srv)
}

func _FigletService_GetFiglet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Figlet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FigletServiceServer).GetFiglet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.FigletService/GetFiglet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FigletServiceServer).GetFiglet(ctx, req.(*Figlet))
	}
	return interceptor(ctx, in, info, handler)
}

func _FigletService_GetFigletFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FigletServiceServer).GetFigletFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.FigletService/GetFigletFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FigletServiceServer).GetFigletFile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FigletService_ServiceDesc is the grpc.ServiceDesc for FigletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FigletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linuxCommand.FigletService",
	HandlerType: (*FigletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiglet",
			Handler:    _FigletService_GetFiglet_Handler,
		},
		{
			MethodName: "GetFigletFile",
			Handler:    _FigletService_GetFigletFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/figlet.proto",
}
