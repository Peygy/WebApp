// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc3
// source: protos/graph_token/graph_token.proto

package graph_token

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CreateTokensPairService_CreateTokensPair_FullMethodName = "/graph_token.CreateTokensPairService/CreateTokensPair"
)

// CreateTokensPairServiceClient is the client API for CreateTokensPairService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateTokensPairServiceClient interface {
	CreateTokensPair(ctx context.Context, in *CreateTokensPairRequest, opts ...grpc.CallOption) (*CreateTokensPairResponce, error)
}

type createTokensPairServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateTokensPairServiceClient(cc grpc.ClientConnInterface) CreateTokensPairServiceClient {
	return &createTokensPairServiceClient{cc}
}

func (c *createTokensPairServiceClient) CreateTokensPair(ctx context.Context, in *CreateTokensPairRequest, opts ...grpc.CallOption) (*CreateTokensPairResponce, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTokensPairResponce)
	err := c.cc.Invoke(ctx, CreateTokensPairService_CreateTokensPair_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateTokensPairServiceServer is the server API for CreateTokensPairService service.
// All implementations must embed UnimplementedCreateTokensPairServiceServer
// for forward compatibility.
type CreateTokensPairServiceServer interface {
	CreateTokensPair(context.Context, *CreateTokensPairRequest) (*CreateTokensPairResponce, error)
	mustEmbedUnimplementedCreateTokensPairServiceServer()
}

// UnimplementedCreateTokensPairServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreateTokensPairServiceServer struct{}

func (UnimplementedCreateTokensPairServiceServer) CreateTokensPair(context.Context, *CreateTokensPairRequest) (*CreateTokensPairResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTokensPair not implemented")
}
func (UnimplementedCreateTokensPairServiceServer) mustEmbedUnimplementedCreateTokensPairServiceServer() {
}
func (UnimplementedCreateTokensPairServiceServer) testEmbeddedByValue() {}

// UnsafeCreateTokensPairServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateTokensPairServiceServer will
// result in compilation errors.
type UnsafeCreateTokensPairServiceServer interface {
	mustEmbedUnimplementedCreateTokensPairServiceServer()
}

func RegisterCreateTokensPairServiceServer(s grpc.ServiceRegistrar, srv CreateTokensPairServiceServer) {
	// If the following call pancis, it indicates UnimplementedCreateTokensPairServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreateTokensPairService_ServiceDesc, srv)
}

func _CreateTokensPairService_CreateTokensPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokensPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateTokensPairServiceServer).CreateTokensPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateTokensPairService_CreateTokensPair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateTokensPairServiceServer).CreateTokensPair(ctx, req.(*CreateTokensPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateTokensPairService_ServiceDesc is the grpc.ServiceDesc for CreateTokensPairService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateTokensPairService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "graph_token.CreateTokensPairService",
	HandlerType: (*CreateTokensPairServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTokensPair",
			Handler:    _CreateTokensPairService_CreateTokensPair_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/graph_token/graph_token.proto",
}