// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/kritzware/google-ads-go/resources"
	servicespb "github.com/kritzware/google-ads-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newExtensionFeedItemClientHook clientHook

// ExtensionFeedItemCallOptions contains the retry settings for each method of ExtensionFeedItemClient.
type ExtensionFeedItemCallOptions struct {
	GetExtensionFeedItem []gax.CallOption
	MutateExtensionFeedItems []gax.CallOption
}

func defaultExtensionFeedItemGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultExtensionFeedItemCallOptions() *ExtensionFeedItemCallOptions {
	return &ExtensionFeedItemCallOptions{
		GetExtensionFeedItem: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MutateExtensionFeedItems: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalExtensionFeedItemClient is an interface that defines the methods availaible from Google Ads API.
type internalExtensionFeedItemClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetExtensionFeedItem(context.Context, *servicespb.GetExtensionFeedItemRequest, ...gax.CallOption) (*resourcespb.ExtensionFeedItem, error)
	MutateExtensionFeedItems(context.Context, *servicespb.MutateExtensionFeedItemsRequest, ...gax.CallOption) (*servicespb.MutateExtensionFeedItemsResponse, error)
}

// ExtensionFeedItemClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage extension feed items.
type ExtensionFeedItemClient struct {
	// The internal transport-dependent client.
	internalClient internalExtensionFeedItemClient

	// The call options for this service.
	CallOptions *ExtensionFeedItemCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ExtensionFeedItemClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ExtensionFeedItemClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ExtensionFeedItemClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetExtensionFeedItem returns the requested extension feed item in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExtensionFeedItemClient) GetExtensionFeedItem(ctx context.Context, req *servicespb.GetExtensionFeedItemRequest, opts ...gax.CallOption) (*resourcespb.ExtensionFeedItem, error) {
	return c.internalClient.GetExtensionFeedItem(ctx, req, opts...)
}

// MutateExtensionFeedItems creates, updates, or removes extension feed items. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// CountryCodeError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// ExtensionFeedItemError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// ImageError (at )
// InternalError (at )
// LanguageCodeError (at )
// MutateError (at )
// NewResourceCreationError (at )
// OperationAccessDeniedError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *ExtensionFeedItemClient) MutateExtensionFeedItems(ctx context.Context, req *servicespb.MutateExtensionFeedItemsRequest, opts ...gax.CallOption) (*servicespb.MutateExtensionFeedItemsResponse, error) {
	return c.internalClient.MutateExtensionFeedItems(ctx, req, opts...)
}

// extensionFeedItemGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type extensionFeedItemGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ExtensionFeedItemClient
	CallOptions **ExtensionFeedItemCallOptions

	// The gRPC API client.
	extensionFeedItemClient servicespb.ExtensionFeedItemServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewExtensionFeedItemClient creates a new extension feed item service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage extension feed items.
func NewExtensionFeedItemClient(ctx context.Context, opts ...option.ClientOption) (*ExtensionFeedItemClient, error) {
	clientOpts := defaultExtensionFeedItemGRPCClientOptions()
	if newExtensionFeedItemClientHook != nil {
		hookOpts, err := newExtensionFeedItemClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ExtensionFeedItemClient{CallOptions: defaultExtensionFeedItemCallOptions()}

	c := &extensionFeedItemGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		extensionFeedItemClient: servicespb.NewExtensionFeedItemServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *extensionFeedItemGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *extensionFeedItemGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *extensionFeedItemGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *extensionFeedItemGRPCClient) GetExtensionFeedItem(ctx context.Context, req *servicespb.GetExtensionFeedItemRequest, opts ...gax.CallOption) (*resourcespb.ExtensionFeedItem, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetExtensionFeedItem[0:len((*c.CallOptions).GetExtensionFeedItem):len((*c.CallOptions).GetExtensionFeedItem)], opts...)
	var resp *resourcespb.ExtensionFeedItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.extensionFeedItemClient.GetExtensionFeedItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *extensionFeedItemGRPCClient) MutateExtensionFeedItems(ctx context.Context, req *servicespb.MutateExtensionFeedItemsRequest, opts ...gax.CallOption) (*servicespb.MutateExtensionFeedItemsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateExtensionFeedItems[0:len((*c.CallOptions).MutateExtensionFeedItems):len((*c.CallOptions).MutateExtensionFeedItems)], opts...)
	var resp *servicespb.MutateExtensionFeedItemsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.extensionFeedItemClient.MutateExtensionFeedItems(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}