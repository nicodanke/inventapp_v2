// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: service_inventapp.proto

package pb

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
	InventApp_CreateAccount_FullMethodName = "/pb.InventApp/CreateAccount"
)

// InventAppClient is the client API for InventApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventAppClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type inventAppClient struct {
	cc grpc.ClientConnInterface
}

func NewInventAppClient(cc grpc.ClientConnInterface) InventAppClient {
	return &inventAppClient{cc}
}

func (c *inventAppClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, InventApp_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventAppServer is the server API for InventApp service.
// All implementations must embed UnimplementedInventAppServer
// for forward compatibility
type InventAppServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	mustEmbedUnimplementedInventAppServer()
}

// UnimplementedInventAppServer must be embedded to have forward compatible implementations.
type UnimplementedInventAppServer struct {
}

func (UnimplementedInventAppServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedInventAppServer) mustEmbedUnimplementedInventAppServer() {}

// UnsafeInventAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventAppServer will
// result in compilation errors.
type UnsafeInventAppServer interface {
	mustEmbedUnimplementedInventAppServer()
}

func RegisterInventAppServer(s grpc.ServiceRegistrar, srv InventAppServer) {
	s.RegisterService(&InventApp_ServiceDesc, srv)
}

func _InventApp_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventApp_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventApp_ServiceDesc is the grpc.ServiceDesc for InventApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.InventApp",
	HandlerType: (*InventAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _InventApp_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_inventapp.proto",
}
