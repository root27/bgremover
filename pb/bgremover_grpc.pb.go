// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/bgremover.proto

package pb

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
	Remove_RemoveBG_FullMethodName = "/pb.Remove/RemoveBG"
)

// RemoveClient is the client API for Remove service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoveClient interface {
	RemoveBG(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
}

type removeClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoveClient(cc grpc.ClientConnInterface) RemoveClient {
	return &removeClient{cc}
}

func (c *removeClient) RemoveBG(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, Remove_RemoveBG_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoveServer is the server API for Remove service.
// All implementations must embed UnimplementedRemoveServer
// for forward compatibility.
type RemoveServer interface {
	RemoveBG(context.Context, *ImageRequest) (*ImageResponse, error)
	mustEmbedUnimplementedRemoveServer()
}

// UnimplementedRemoveServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRemoveServer struct{}

func (UnimplementedRemoveServer) RemoveBG(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBG not implemented")
}
func (UnimplementedRemoveServer) mustEmbedUnimplementedRemoveServer() {}
func (UnimplementedRemoveServer) testEmbeddedByValue()                {}

// UnsafeRemoveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoveServer will
// result in compilation errors.
type UnsafeRemoveServer interface {
	mustEmbedUnimplementedRemoveServer()
}

func RegisterRemoveServer(s grpc.ServiceRegistrar, srv RemoveServer) {
	// If the following call pancis, it indicates UnimplementedRemoveServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Remove_ServiceDesc, srv)
}

func _Remove_RemoveBG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoveServer).RemoveBG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Remove_RemoveBG_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoveServer).RemoveBG(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Remove_ServiceDesc is the grpc.ServiceDesc for Remove service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Remove_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Remove",
	HandlerType: (*RemoveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveBG",
			Handler:    _Remove_RemoveBG_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bgremover.proto",
}
