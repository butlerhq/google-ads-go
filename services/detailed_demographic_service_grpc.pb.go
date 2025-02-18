// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/butlerhq/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DetailedDemographicServiceClient is the client API for DetailedDemographicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetailedDemographicServiceClient interface {
	// Returns the requested detailed demographic in full detail.
	GetDetailedDemographic(ctx context.Context, in *GetDetailedDemographicRequest, opts ...grpc.CallOption) (*resources.DetailedDemographic, error)
}

type detailedDemographicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailedDemographicServiceClient(cc grpc.ClientConnInterface) DetailedDemographicServiceClient {
	return &detailedDemographicServiceClient{cc}
}

func (c *detailedDemographicServiceClient) GetDetailedDemographic(ctx context.Context, in *GetDetailedDemographicRequest, opts ...grpc.CallOption) (*resources.DetailedDemographic, error) {
	out := new(resources.DetailedDemographic)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.DetailedDemographicService/GetDetailedDemographic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailedDemographicServiceServer is the server API for DetailedDemographicService service.
// All implementations must embed UnimplementedDetailedDemographicServiceServer
// for forward compatibility
type DetailedDemographicServiceServer interface {
	// Returns the requested detailed demographic in full detail.
	GetDetailedDemographic(context.Context, *GetDetailedDemographicRequest) (*resources.DetailedDemographic, error)
	mustEmbedUnimplementedDetailedDemographicServiceServer()
}

// UnimplementedDetailedDemographicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDetailedDemographicServiceServer struct {
}

func (UnimplementedDetailedDemographicServiceServer) GetDetailedDemographic(context.Context, *GetDetailedDemographicRequest) (*resources.DetailedDemographic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailedDemographic not implemented")
}
func (UnimplementedDetailedDemographicServiceServer) mustEmbedUnimplementedDetailedDemographicServiceServer() {
}

// UnsafeDetailedDemographicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetailedDemographicServiceServer will
// result in compilation errors.
type UnsafeDetailedDemographicServiceServer interface {
	mustEmbedUnimplementedDetailedDemographicServiceServer()
}

func RegisterDetailedDemographicServiceServer(s grpc.ServiceRegistrar, srv DetailedDemographicServiceServer) {
	s.RegisterService(&DetailedDemographicService_ServiceDesc, srv)
}

func _DetailedDemographicService_GetDetailedDemographic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailedDemographicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailedDemographicServiceServer).GetDetailedDemographic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.DetailedDemographicService/GetDetailedDemographic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailedDemographicServiceServer).GetDetailedDemographic(ctx, req.(*GetDetailedDemographicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DetailedDemographicService_ServiceDesc is the grpc.ServiceDesc for DetailedDemographicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DetailedDemographicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.DetailedDemographicService",
	HandlerType: (*DetailedDemographicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailedDemographic",
			Handler:    _DetailedDemographicService_GetDetailedDemographic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/detailed_demographic_service.proto",
}
