// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: v1/interface.proto

package v1

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

// InterfaceClient is the client API for Interface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterfaceClient interface {
	Data(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type interfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterfaceClient(cc grpc.ClientConnInterface) InterfaceClient {
	return &interfaceClient{cc}
}

func (c *interfaceClient) Data(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/api.interface.v1.Interface/Data", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfaceServer is the server API for Interface service.
// All implementations must embed UnimplementedInterfaceServer
// for forward compatibility
type InterfaceServer interface {
	Data(context.Context, *DataRequest) (*DataResponse, error)
	mustEmbedUnimplementedInterfaceServer()
}

// UnimplementedInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedInterfaceServer struct {
}

func (UnimplementedInterfaceServer) Data(context.Context, *DataRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Data not implemented")
}
func (UnimplementedInterfaceServer) mustEmbedUnimplementedInterfaceServer() {}

// UnsafeInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterfaceServer will
// result in compilation errors.
type UnsafeInterfaceServer interface {
	mustEmbedUnimplementedInterfaceServer()
}

func RegisterInterfaceServer(s grpc.ServiceRegistrar, srv InterfaceServer) {
	s.RegisterService(&Interface_ServiceDesc, srv)
}

func _Interface_Data_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).Data(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.interface.v1.Interface/Data",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).Data(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Interface_ServiceDesc is the grpc.ServiceDesc for Interface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.interface.v1.Interface",
	HandlerType: (*InterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Data",
			Handler:    _Interface_Data_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/interface.proto",
}
