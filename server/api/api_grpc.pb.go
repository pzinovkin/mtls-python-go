// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// DiceServiceClient is the client API for DiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiceServiceClient interface {
	RollDie(ctx context.Context, in *RollDieRequest, opts ...grpc.CallOption) (*RollDieResponse, error)
}

type diceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiceServiceClient(cc grpc.ClientConnInterface) DiceServiceClient {
	return &diceServiceClient{cc}
}

func (c *diceServiceClient) RollDie(ctx context.Context, in *RollDieRequest, opts ...grpc.CallOption) (*RollDieResponse, error) {
	out := new(RollDieResponse)
	err := c.cc.Invoke(ctx, "/api.DiceService/RollDie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiceServiceServer is the server API for DiceService service.
// All implementations must embed UnimplementedDiceServiceServer
// for forward compatibility
type DiceServiceServer interface {
	RollDie(context.Context, *RollDieRequest) (*RollDieResponse, error)
	mustEmbedUnimplementedDiceServiceServer()
}

// UnimplementedDiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiceServiceServer struct {
}

func (UnimplementedDiceServiceServer) RollDie(context.Context, *RollDieRequest) (*RollDieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDie not implemented")
}
func (UnimplementedDiceServiceServer) mustEmbedUnimplementedDiceServiceServer() {}

// UnsafeDiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiceServiceServer will
// result in compilation errors.
type UnsafeDiceServiceServer interface {
	mustEmbedUnimplementedDiceServiceServer()
}

func RegisterDiceServiceServer(s grpc.ServiceRegistrar, srv DiceServiceServer) {
	s.RegisterService(&DiceService_ServiceDesc, srv)
}

func _DiceService_RollDie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollDieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiceServiceServer).RollDie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DiceService/RollDie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiceServiceServer).RollDie(ctx, req.(*RollDieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiceService_ServiceDesc is the grpc.ServiceDesc for DiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.DiceService",
	HandlerType: (*DiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RollDie",
			Handler:    _DiceService_RollDie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}