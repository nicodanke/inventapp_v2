// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: service_inventapp-v1.proto

package pb

import (
	context "context"
	account "github.com/nicodanke/inventapp_v2/pb/requests/v1/account"
	login "github.com/nicodanke/inventapp_v2/pb/requests/v1/login"
	role "github.com/nicodanke/inventapp_v2/pb/requests/v1/role"
	user "github.com/nicodanke/inventapp_v2/pb/requests/v1/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	InventAppV1_Login_FullMethodName         = "/pb.InventAppV1/Login"
	InventAppV1_CreateAccount_FullMethodName = "/pb.InventAppV1/CreateAccount"
	InventAppV1_UpdateAccount_FullMethodName = "/pb.InventAppV1/UpdateAccount"
	InventAppV1_CreateUser_FullMethodName    = "/pb.InventAppV1/CreateUser"
	InventAppV1_GetRoles_FullMethodName      = "/pb.InventAppV1/GetRoles"
	InventAppV1_CreateRole_FullMethodName    = "/pb.InventAppV1/CreateRole"
)

// InventAppV1Client is the client API for InventAppV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventAppV1Client interface {
	// LOGIN
	Login(ctx context.Context, in *login.LoginRequest, opts ...grpc.CallOption) (*login.LoginResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *account.CreateAccountRequest, opts ...grpc.CallOption) (*account.CreateAccountResponse, error)
	UpdateAccount(ctx context.Context, in *account.UpdateAccountRequest, opts ...grpc.CallOption) (*account.UpdateAccountResponse, error)
	// USER
	CreateUser(ctx context.Context, in *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error)
	// ROLE
	GetRoles(ctx context.Context, in *role.GetRolesRequest, opts ...grpc.CallOption) (*role.GetRolesResponse, error)
	CreateRole(ctx context.Context, in *role.CreateRoleRequest, opts ...grpc.CallOption) (*role.CreateRoleResponse, error)
}

type inventAppV1Client struct {
	cc grpc.ClientConnInterface
}

func NewInventAppV1Client(cc grpc.ClientConnInterface) InventAppV1Client {
	return &inventAppV1Client{cc}
}

func (c *inventAppV1Client) Login(ctx context.Context, in *login.LoginRequest, opts ...grpc.CallOption) (*login.LoginResponse, error) {
	out := new(login.LoginResponse)
	err := c.cc.Invoke(ctx, InventAppV1_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventAppV1Client) CreateAccount(ctx context.Context, in *account.CreateAccountRequest, opts ...grpc.CallOption) (*account.CreateAccountResponse, error) {
	out := new(account.CreateAccountResponse)
	err := c.cc.Invoke(ctx, InventAppV1_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventAppV1Client) UpdateAccount(ctx context.Context, in *account.UpdateAccountRequest, opts ...grpc.CallOption) (*account.UpdateAccountResponse, error) {
	out := new(account.UpdateAccountResponse)
	err := c.cc.Invoke(ctx, InventAppV1_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventAppV1Client) CreateUser(ctx context.Context, in *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error) {
	out := new(user.CreateUserResponse)
	err := c.cc.Invoke(ctx, InventAppV1_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventAppV1Client) GetRoles(ctx context.Context, in *role.GetRolesRequest, opts ...grpc.CallOption) (*role.GetRolesResponse, error) {
	out := new(role.GetRolesResponse)
	err := c.cc.Invoke(ctx, InventAppV1_GetRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventAppV1Client) CreateRole(ctx context.Context, in *role.CreateRoleRequest, opts ...grpc.CallOption) (*role.CreateRoleResponse, error) {
	out := new(role.CreateRoleResponse)
	err := c.cc.Invoke(ctx, InventAppV1_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventAppV1Server is the server API for InventAppV1 service.
// All implementations must embed UnimplementedInventAppV1Server
// for forward compatibility
type InventAppV1Server interface {
	// LOGIN
	Login(context.Context, *login.LoginRequest) (*login.LoginResponse, error)
	// ACCOUNT
	CreateAccount(context.Context, *account.CreateAccountRequest) (*account.CreateAccountResponse, error)
	UpdateAccount(context.Context, *account.UpdateAccountRequest) (*account.UpdateAccountResponse, error)
	// USER
	CreateUser(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error)
	// ROLE
	GetRoles(context.Context, *role.GetRolesRequest) (*role.GetRolesResponse, error)
	CreateRole(context.Context, *role.CreateRoleRequest) (*role.CreateRoleResponse, error)
	mustEmbedUnimplementedInventAppV1Server()
}

// UnimplementedInventAppV1Server must be embedded to have forward compatible implementations.
type UnimplementedInventAppV1Server struct {
}

func (UnimplementedInventAppV1Server) Login(context.Context, *login.LoginRequest) (*login.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedInventAppV1Server) CreateAccount(context.Context, *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedInventAppV1Server) UpdateAccount(context.Context, *account.UpdateAccountRequest) (*account.UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedInventAppV1Server) CreateUser(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedInventAppV1Server) GetRoles(context.Context, *role.GetRolesRequest) (*role.GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedInventAppV1Server) CreateRole(context.Context, *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedInventAppV1Server) mustEmbedUnimplementedInventAppV1Server() {}

// UnsafeInventAppV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventAppV1Server will
// result in compilation errors.
type UnsafeInventAppV1Server interface {
	mustEmbedUnimplementedInventAppV1Server()
}

func RegisterInventAppV1Server(s grpc.ServiceRegistrar, srv InventAppV1Server) {
	s.RegisterService(&InventAppV1_ServiceDesc, srv)
}

func _InventAppV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(login.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).Login(ctx, req.(*login.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventAppV1_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).CreateAccount(ctx, req.(*account.CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventAppV1_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).UpdateAccount(ctx, req.(*account.UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventAppV1_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).CreateUser(ctx, req.(*user.CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventAppV1_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_GetRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).GetRoles(ctx, req.(*role.GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventAppV1_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(role.CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventAppV1Server).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventAppV1_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventAppV1Server).CreateRole(ctx, req.(*role.CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventAppV1_ServiceDesc is the grpc.ServiceDesc for InventAppV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventAppV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.InventAppV1",
	HandlerType: (*InventAppV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _InventAppV1_Login_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _InventAppV1_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _InventAppV1_UpdateAccount_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _InventAppV1_CreateUser_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _InventAppV1_GetRoles_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _InventAppV1_CreateRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_inventapp-v1.proto",
}
