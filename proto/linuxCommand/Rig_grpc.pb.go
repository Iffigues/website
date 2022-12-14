// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/Rig.proto

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

// RigServiceClient is the client API for RigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RigServiceClient interface {
	GetRig(ctx context.Context, in *Rig, opts ...grpc.CallOption) (*Responses, error)
}

type rigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRigServiceClient(cc grpc.ClientConnInterface) RigServiceClient {
	return &rigServiceClient{cc}
}

func (c *rigServiceClient) GetRig(ctx context.Context, in *Rig, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, "/linuxCommand.RigService/GetRig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RigServiceServer is the server API for RigService service.
// All implementations must embed UnimplementedRigServiceServer
// for forward compatibility
type RigServiceServer interface {
	GetRig(context.Context, *Rig) (*Responses, error)
	mustEmbedUnimplementedRigServiceServer()
}

// UnimplementedRigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRigServiceServer struct {
}

func (UnimplementedRigServiceServer) GetRig(context.Context, *Rig) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRig not implemented")
}
func (UnimplementedRigServiceServer) mustEmbedUnimplementedRigServiceServer() {}

// UnsafeRigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RigServiceServer will
// result in compilation errors.
type UnsafeRigServiceServer interface {
	mustEmbedUnimplementedRigServiceServer()
}

func RegisterRigServiceServer(s grpc.ServiceRegistrar, srv RigServiceServer) {
	s.RegisterService(&RigService_ServiceDesc, srv)
}

func _RigService_GetRig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RigServiceServer).GetRig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxCommand.RigService/GetRig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RigServiceServer).GetRig(ctx, req.(*Rig))
	}
	return interceptor(ctx, in, info, handler)
}

// RigService_ServiceDesc is the grpc.ServiceDesc for RigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linuxCommand.RigService",
	HandlerType: (*RigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRig",
			Handler:    _RigService_GetRig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Rig.proto",
}
