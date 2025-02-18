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

// IncomeRangeViewServiceClient is the client API for IncomeRangeViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncomeRangeViewServiceClient interface {
	// Returns the requested income range view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetIncomeRangeView(ctx context.Context, in *GetIncomeRangeViewRequest, opts ...grpc.CallOption) (*resources.IncomeRangeView, error)
}

type incomeRangeViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncomeRangeViewServiceClient(cc grpc.ClientConnInterface) IncomeRangeViewServiceClient {
	return &incomeRangeViewServiceClient{cc}
}

func (c *incomeRangeViewServiceClient) GetIncomeRangeView(ctx context.Context, in *GetIncomeRangeViewRequest, opts ...grpc.CallOption) (*resources.IncomeRangeView, error) {
	out := new(resources.IncomeRangeView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.IncomeRangeViewService/GetIncomeRangeView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncomeRangeViewServiceServer is the server API for IncomeRangeViewService service.
// All implementations must embed UnimplementedIncomeRangeViewServiceServer
// for forward compatibility
type IncomeRangeViewServiceServer interface {
	// Returns the requested income range view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetIncomeRangeView(context.Context, *GetIncomeRangeViewRequest) (*resources.IncomeRangeView, error)
	mustEmbedUnimplementedIncomeRangeViewServiceServer()
}

// UnimplementedIncomeRangeViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIncomeRangeViewServiceServer struct {
}

func (UnimplementedIncomeRangeViewServiceServer) GetIncomeRangeView(context.Context, *GetIncomeRangeViewRequest) (*resources.IncomeRangeView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomeRangeView not implemented")
}
func (UnimplementedIncomeRangeViewServiceServer) mustEmbedUnimplementedIncomeRangeViewServiceServer() {
}

// UnsafeIncomeRangeViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncomeRangeViewServiceServer will
// result in compilation errors.
type UnsafeIncomeRangeViewServiceServer interface {
	mustEmbedUnimplementedIncomeRangeViewServiceServer()
}

func RegisterIncomeRangeViewServiceServer(s grpc.ServiceRegistrar, srv IncomeRangeViewServiceServer) {
	s.RegisterService(&IncomeRangeViewService_ServiceDesc, srv)
}

func _IncomeRangeViewService_GetIncomeRangeView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncomeRangeViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomeRangeViewServiceServer).GetIncomeRangeView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.IncomeRangeViewService/GetIncomeRangeView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomeRangeViewServiceServer).GetIncomeRangeView(ctx, req.(*GetIncomeRangeViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IncomeRangeViewService_ServiceDesc is the grpc.ServiceDesc for IncomeRangeViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncomeRangeViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.IncomeRangeViewService",
	HandlerType: (*IncomeRangeViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIncomeRangeView",
			Handler:    _IncomeRangeViewService_GetIncomeRangeView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/income_range_view_service.proto",
}
