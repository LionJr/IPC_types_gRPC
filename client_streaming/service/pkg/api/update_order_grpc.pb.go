// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/update_order.proto

package client_stream_api

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UpdateOrderService_UpdateOrders_FullMethodName = "/client_stream_api.UpdateOrderService/updateOrders"
)

// UpdateOrderServiceClient is the client API for UpdateOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateOrderServiceClient interface {
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Order, wrappers.StringValue], error)
}

type updateOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateOrderServiceClient(cc grpc.ClientConnInterface) UpdateOrderServiceClient {
	return &updateOrderServiceClient{cc}
}

func (c *updateOrderServiceClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Order, wrappers.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UpdateOrderService_ServiceDesc.Streams[0], UpdateOrderService_UpdateOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Order, wrappers.StringValue]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UpdateOrderService_UpdateOrdersClient = grpc.ClientStreamingClient[Order, wrappers.StringValue]

// UpdateOrderServiceServer is the server API for UpdateOrderService service.
// All implementations must embed UnimplementedUpdateOrderServiceServer
// for forward compatibility.
type UpdateOrderServiceServer interface {
	UpdateOrders(grpc.ClientStreamingServer[Order, wrappers.StringValue]) error
	mustEmbedUnimplementedUpdateOrderServiceServer()
}

// UnimplementedUpdateOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUpdateOrderServiceServer struct{}

func (UnimplementedUpdateOrderServiceServer) UpdateOrders(grpc.ClientStreamingServer[Order, wrappers.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedUpdateOrderServiceServer) mustEmbedUnimplementedUpdateOrderServiceServer() {}
func (UnimplementedUpdateOrderServiceServer) testEmbeddedByValue()                            {}

// UnsafeUpdateOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateOrderServiceServer will
// result in compilation errors.
type UnsafeUpdateOrderServiceServer interface {
	mustEmbedUnimplementedUpdateOrderServiceServer()
}

func RegisterUpdateOrderServiceServer(s grpc.ServiceRegistrar, srv UpdateOrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedUpdateOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UpdateOrderService_ServiceDesc, srv)
}

func _UpdateOrderService_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpdateOrderServiceServer).UpdateOrders(&grpc.GenericServerStream[Order, wrappers.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UpdateOrderService_UpdateOrdersServer = grpc.ClientStreamingServer[Order, wrappers.StringValue]

// UpdateOrderService_ServiceDesc is the grpc.ServiceDesc for UpdateOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client_stream_api.UpdateOrderService",
	HandlerType: (*UpdateOrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "updateOrders",
			Handler:       _UpdateOrderService_UpdateOrders_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/update_order.proto",
}
