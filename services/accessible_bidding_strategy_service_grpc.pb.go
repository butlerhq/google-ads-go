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

// AccessibleBiddingStrategyServiceClient is the client API for AccessibleBiddingStrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessibleBiddingStrategyServiceClient interface {
	// Returns the requested accessible bidding strategy in full detail.
	GetAccessibleBiddingStrategy(ctx context.Context, in *GetAccessibleBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.AccessibleBiddingStrategy, error)
}

type accessibleBiddingStrategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessibleBiddingStrategyServiceClient(cc grpc.ClientConnInterface) AccessibleBiddingStrategyServiceClient {
	return &accessibleBiddingStrategyServiceClient{cc}
}

func (c *accessibleBiddingStrategyServiceClient) GetAccessibleBiddingStrategy(ctx context.Context, in *GetAccessibleBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.AccessibleBiddingStrategy, error) {
	out := new(resources.AccessibleBiddingStrategy)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AccessibleBiddingStrategyService/GetAccessibleBiddingStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessibleBiddingStrategyServiceServer is the server API for AccessibleBiddingStrategyService service.
// All implementations must embed UnimplementedAccessibleBiddingStrategyServiceServer
// for forward compatibility
type AccessibleBiddingStrategyServiceServer interface {
	// Returns the requested accessible bidding strategy in full detail.
	GetAccessibleBiddingStrategy(context.Context, *GetAccessibleBiddingStrategyRequest) (*resources.AccessibleBiddingStrategy, error)
	mustEmbedUnimplementedAccessibleBiddingStrategyServiceServer()
}

// UnimplementedAccessibleBiddingStrategyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessibleBiddingStrategyServiceServer struct {
}

func (UnimplementedAccessibleBiddingStrategyServiceServer) GetAccessibleBiddingStrategy(context.Context, *GetAccessibleBiddingStrategyRequest) (*resources.AccessibleBiddingStrategy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessibleBiddingStrategy not implemented")
}
func (UnimplementedAccessibleBiddingStrategyServiceServer) mustEmbedUnimplementedAccessibleBiddingStrategyServiceServer() {
}

// UnsafeAccessibleBiddingStrategyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessibleBiddingStrategyServiceServer will
// result in compilation errors.
type UnsafeAccessibleBiddingStrategyServiceServer interface {
	mustEmbedUnimplementedAccessibleBiddingStrategyServiceServer()
}

func RegisterAccessibleBiddingStrategyServiceServer(s grpc.ServiceRegistrar, srv AccessibleBiddingStrategyServiceServer) {
	s.RegisterService(&AccessibleBiddingStrategyService_ServiceDesc, srv)
}

func _AccessibleBiddingStrategyService_GetAccessibleBiddingStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessibleBiddingStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessibleBiddingStrategyServiceServer).GetAccessibleBiddingStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AccessibleBiddingStrategyService/GetAccessibleBiddingStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessibleBiddingStrategyServiceServer).GetAccessibleBiddingStrategy(ctx, req.(*GetAccessibleBiddingStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessibleBiddingStrategyService_ServiceDesc is the grpc.ServiceDesc for AccessibleBiddingStrategyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessibleBiddingStrategyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AccessibleBiddingStrategyService",
	HandlerType: (*AccessibleBiddingStrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccessibleBiddingStrategy",
			Handler:    _AccessibleBiddingStrategyService_GetAccessibleBiddingStrategy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/accessible_bidding_strategy_service.proto",
}
