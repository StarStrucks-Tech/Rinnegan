// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: onboarding/onboarding_service.proto

package onboarding

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OnboardingService_UpdateCurrentStage_FullMethodName = "/onboarding.OnboardingService/UpdateCurrentStage"
)

// OnboardingServiceClient is the client API for OnboardingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the service
type OnboardingServiceClient interface {
	UpdateCurrentStage(ctx context.Context, in *UpdateCurrentStageRequest, opts ...grpc.CallOption) (*UpdateCurrentStageResponse, error)
}

type onboardingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOnboardingServiceClient(cc grpc.ClientConnInterface) OnboardingServiceClient {
	return &onboardingServiceClient{cc}
}

func (c *onboardingServiceClient) UpdateCurrentStage(ctx context.Context, in *UpdateCurrentStageRequest, opts ...grpc.CallOption) (*UpdateCurrentStageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentStageResponse)
	err := c.cc.Invoke(ctx, OnboardingService_UpdateCurrentStage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnboardingServiceServer is the server API for OnboardingService service.
// All implementations must embed UnimplementedOnboardingServiceServer
// for forward compatibility
//
// Define the service
type OnboardingServiceServer interface {
	UpdateCurrentStage(context.Context, *UpdateCurrentStageRequest) (*UpdateCurrentStageResponse, error)
	mustEmbedUnimplementedOnboardingServiceServer()
}

// UnimplementedOnboardingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOnboardingServiceServer struct {
}

func (UnimplementedOnboardingServiceServer) UpdateCurrentStage(context.Context, *UpdateCurrentStageRequest) (*UpdateCurrentStageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentStage not implemented")
}
func (UnimplementedOnboardingServiceServer) mustEmbedUnimplementedOnboardingServiceServer() {}

// UnsafeOnboardingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnboardingServiceServer will
// result in compilation errors.
type UnsafeOnboardingServiceServer interface {
	mustEmbedUnimplementedOnboardingServiceServer()
}

func RegisterOnboardingServiceServer(s grpc.ServiceRegistrar, srv OnboardingServiceServer) {
	s.RegisterService(&OnboardingService_ServiceDesc, srv)
}

func _OnboardingService_UpdateCurrentStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnboardingServiceServer).UpdateCurrentStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnboardingService_UpdateCurrentStage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnboardingServiceServer).UpdateCurrentStage(ctx, req.(*UpdateCurrentStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OnboardingService_ServiceDesc is the grpc.ServiceDesc for OnboardingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnboardingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "onboarding.OnboardingService",
	HandlerType: (*OnboardingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateCurrentStage",
			Handler:    _OnboardingService_UpdateCurrentStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding/onboarding_service.proto",
}
