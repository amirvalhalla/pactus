// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: wallet.proto

package pactus

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

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	// CreateWallet creates a new wallet with the specified parameters.
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	// LoadWallet loads an existing wallet with the given name.
	LoadWallet(ctx context.Context, in *LoadWalletRequest, opts ...grpc.CallOption) (*LoadWalletResponse, error)
	// UnloadWallet unloads a currently loaded wallet with the specified name.
	UnloadWallet(ctx context.Context, in *UnloadWalletRequest, opts ...grpc.CallOption) (*UnloadWalletResponse, error)
	// LockWallet locks a currently loaded wallet with the provided password and timeout.
	LockWallet(ctx context.Context, in *LockWalletRequest, opts ...grpc.CallOption) (*LockWalletResponse, error)
	// UnlockWallet unlocks a locked wallet with the provided password and timeout.
	UnlockWallet(ctx context.Context, in *UnlockWalletRequest, opts ...grpc.CallOption) (*UnlockWalletResponse, error)
	// SignRawTransaction Signs a raw transaction for a specified wallet.
	SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) LoadWallet(ctx context.Context, in *LoadWalletRequest, opts ...grpc.CallOption) (*LoadWalletResponse, error) {
	out := new(LoadWalletResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/LoadWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UnloadWallet(ctx context.Context, in *UnloadWalletRequest, opts ...grpc.CallOption) (*UnloadWalletResponse, error) {
	out := new(UnloadWalletResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/UnloadWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) LockWallet(ctx context.Context, in *LockWalletRequest, opts ...grpc.CallOption) (*LockWalletResponse, error) {
	out := new(LockWalletResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/LockWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UnlockWallet(ctx context.Context, in *UnlockWalletRequest, opts ...grpc.CallOption) (*UnlockWalletResponse, error) {
	out := new(UnlockWalletResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/UnlockWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error) {
	out := new(SignRawTransactionResponse)
	err := c.cc.Invoke(ctx, "/pactus.Wallet/SignRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations should embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	// CreateWallet creates a new wallet with the specified parameters.
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	// LoadWallet loads an existing wallet with the given name.
	LoadWallet(context.Context, *LoadWalletRequest) (*LoadWalletResponse, error)
	// UnloadWallet unloads a currently loaded wallet with the specified name.
	UnloadWallet(context.Context, *UnloadWalletRequest) (*UnloadWalletResponse, error)
	// LockWallet locks a currently loaded wallet with the provided password and timeout.
	LockWallet(context.Context, *LockWalletRequest) (*LockWalletResponse, error)
	// UnlockWallet unlocks a locked wallet with the provided password and timeout.
	UnlockWallet(context.Context, *UnlockWalletRequest) (*UnlockWalletResponse, error)
	// SignRawTransaction Signs a raw transaction for a specified wallet.
	SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error)
}

// UnimplementedWalletServer should be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServer) LoadWallet(context.Context, *LoadWalletRequest) (*LoadWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadWallet not implemented")
}
func (UnimplementedWalletServer) UnloadWallet(context.Context, *UnloadWalletRequest) (*UnloadWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnloadWallet not implemented")
}
func (UnimplementedWalletServer) LockWallet(context.Context, *LockWalletRequest) (*LockWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockWallet not implemented")
}
func (UnimplementedWalletServer) UnlockWallet(context.Context, *UnlockWalletRequest) (*UnlockWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockWallet not implemented")
}
func (UnimplementedWalletServer) SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignRawTransaction not implemented")
}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_LoadWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).LoadWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/LoadWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).LoadWallet(ctx, req.(*LoadWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UnloadWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnloadWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UnloadWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/UnloadWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UnloadWallet(ctx, req.(*UnloadWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_LockWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).LockWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/LockWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).LockWallet(ctx, req.(*LockWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UnlockWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UnlockWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/UnlockWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UnlockWallet(ctx, req.(*UnlockWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SignRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SignRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Wallet/SignRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SignRawTransaction(ctx, req.(*SignRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pactus.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _Wallet_CreateWallet_Handler,
		},
		{
			MethodName: "LoadWallet",
			Handler:    _Wallet_LoadWallet_Handler,
		},
		{
			MethodName: "UnloadWallet",
			Handler:    _Wallet_UnloadWallet_Handler,
		},
		{
			MethodName: "LockWallet",
			Handler:    _Wallet_LockWallet_Handler,
		},
		{
			MethodName: "UnlockWallet",
			Handler:    _Wallet_UnlockWallet_Handler,
		},
		{
			MethodName: "SignRawTransaction",
			Handler:    _Wallet_SignRawTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
