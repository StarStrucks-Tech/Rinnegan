// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: payments/payments_service.proto

package payments

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// PaymentsServiceClient is the client API for PaymentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the service
type PaymentsServiceClient interface {
}

type paymentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsServiceClient(cc grpc.ClientConnInterface) PaymentsServiceClient {
	return &paymentsServiceClient{cc}
}

// PaymentsServiceServer is the server API for PaymentsService service.
// All implementations must embed UnimplementedPaymentsServiceServer
// for forward compatibility
//
// Define the service
type PaymentsServiceServer interface {
	mustEmbedUnimplementedPaymentsServiceServer()
}

// UnimplementedPaymentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsServiceServer struct {
}

func (UnimplementedPaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {}

// UnsafePaymentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsServiceServer will
// result in compilation errors.
type UnsafePaymentsServiceServer interface {
	mustEmbedUnimplementedPaymentsServiceServer()
}

func RegisterPaymentsServiceServer(s grpc.ServiceRegistrar, srv PaymentsServiceServer) {
	s.RegisterService(&PaymentsService_ServiceDesc, srv)
}

// PaymentsService_ServiceDesc is the grpc.ServiceDesc for PaymentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payments.PaymentsService",
	HandlerType: (*PaymentsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "payments/payments_service.proto",
}
