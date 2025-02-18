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

// PaidOrganicSearchTermViewServiceClient is the client API for PaidOrganicSearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaidOrganicSearchTermViewServiceClient interface {
	// Returns the requested paid organic search term view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error)
}

type paidOrganicSearchTermViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaidOrganicSearchTermViewServiceClient(cc grpc.ClientConnInterface) PaidOrganicSearchTermViewServiceClient {
	return &paidOrganicSearchTermViewServiceClient{cc}
}

func (c *paidOrganicSearchTermViewServiceClient) GetPaidOrganicSearchTermView(ctx context.Context, in *GetPaidOrganicSearchTermViewRequest, opts ...grpc.CallOption) (*resources.PaidOrganicSearchTermView, error) {
	out := new(resources.PaidOrganicSearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaidOrganicSearchTermViewServiceServer is the server API for PaidOrganicSearchTermViewService service.
// All implementations must embed UnimplementedPaidOrganicSearchTermViewServiceServer
// for forward compatibility
type PaidOrganicSearchTermViewServiceServer interface {
	// Returns the requested paid organic search term view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetPaidOrganicSearchTermView(context.Context, *GetPaidOrganicSearchTermViewRequest) (*resources.PaidOrganicSearchTermView, error)
	mustEmbedUnimplementedPaidOrganicSearchTermViewServiceServer()
}

// UnimplementedPaidOrganicSearchTermViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaidOrganicSearchTermViewServiceServer struct {
}

func (UnimplementedPaidOrganicSearchTermViewServiceServer) GetPaidOrganicSearchTermView(context.Context, *GetPaidOrganicSearchTermViewRequest) (*resources.PaidOrganicSearchTermView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaidOrganicSearchTermView not implemented")
}
func (UnimplementedPaidOrganicSearchTermViewServiceServer) mustEmbedUnimplementedPaidOrganicSearchTermViewServiceServer() {
}

// UnsafePaidOrganicSearchTermViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaidOrganicSearchTermViewServiceServer will
// result in compilation errors.
type UnsafePaidOrganicSearchTermViewServiceServer interface {
	mustEmbedUnimplementedPaidOrganicSearchTermViewServiceServer()
}

func RegisterPaidOrganicSearchTermViewServiceServer(s grpc.ServiceRegistrar, srv PaidOrganicSearchTermViewServiceServer) {
	s.RegisterService(&PaidOrganicSearchTermViewService_ServiceDesc, srv)
}

func _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaidOrganicSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.PaidOrganicSearchTermViewService/GetPaidOrganicSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaidOrganicSearchTermViewServiceServer).GetPaidOrganicSearchTermView(ctx, req.(*GetPaidOrganicSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaidOrganicSearchTermViewService_ServiceDesc is the grpc.ServiceDesc for PaidOrganicSearchTermViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaidOrganicSearchTermViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.PaidOrganicSearchTermViewService",
	HandlerType: (*PaidOrganicSearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaidOrganicSearchTermView",
			Handler:    _PaidOrganicSearchTermViewService_GetPaidOrganicSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/paid_organic_search_term_view_service.proto",
}
