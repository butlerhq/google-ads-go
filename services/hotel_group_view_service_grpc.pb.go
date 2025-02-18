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

// HotelGroupViewServiceClient is the client API for HotelGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelGroupViewServiceClient interface {
	// Returns the requested Hotel Group View in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error)
}

type hotelGroupViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelGroupViewServiceClient(cc grpc.ClientConnInterface) HotelGroupViewServiceClient {
	return &hotelGroupViewServiceClient{cc}
}

func (c *hotelGroupViewServiceClient) GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error) {
	out := new(resources.HotelGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.HotelGroupViewService/GetHotelGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelGroupViewServiceServer is the server API for HotelGroupViewService service.
// All implementations must embed UnimplementedHotelGroupViewServiceServer
// for forward compatibility
type HotelGroupViewServiceServer interface {
	// Returns the requested Hotel Group View in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error)
	mustEmbedUnimplementedHotelGroupViewServiceServer()
}

// UnimplementedHotelGroupViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelGroupViewServiceServer struct {
}

func (UnimplementedHotelGroupViewServiceServer) GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotelGroupView not implemented")
}
func (UnimplementedHotelGroupViewServiceServer) mustEmbedUnimplementedHotelGroupViewServiceServer() {}

// UnsafeHotelGroupViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelGroupViewServiceServer will
// result in compilation errors.
type UnsafeHotelGroupViewServiceServer interface {
	mustEmbedUnimplementedHotelGroupViewServiceServer()
}

func RegisterHotelGroupViewServiceServer(s grpc.ServiceRegistrar, srv HotelGroupViewServiceServer) {
	s.RegisterService(&HotelGroupViewService_ServiceDesc, srv)
}

func _HotelGroupViewService_GetHotelGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.HotelGroupViewService/GetHotelGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, req.(*GetHotelGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelGroupViewService_ServiceDesc is the grpc.ServiceDesc for HotelGroupViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelGroupViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.HotelGroupViewService",
	HandlerType: (*HotelGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelGroupView",
			Handler:    _HotelGroupViewService_GetHotelGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/hotel_group_view_service.proto",
}
