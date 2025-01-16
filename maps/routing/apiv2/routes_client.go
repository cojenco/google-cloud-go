// Copyright 2025 Google LLC
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

package routing

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	routingpb "cloud.google.com/go/maps/routing/apiv2/routingpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newRoutesClientHook clientHook

// RoutesCallOptions contains the retry settings for each method of RoutesClient.
type RoutesCallOptions struct {
	ComputeRoutes      []gax.CallOption
	ComputeRouteMatrix []gax.CallOption
}

func defaultRoutesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("routes.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("routes.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("routes.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://routes.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultRoutesCallOptions() *RoutesCallOptions {
	return &RoutesCallOptions{
		ComputeRoutes:      []gax.CallOption{},
		ComputeRouteMatrix: []gax.CallOption{},
	}
}

func defaultRoutesRESTCallOptions() *RoutesCallOptions {
	return &RoutesCallOptions{
		ComputeRoutes:      []gax.CallOption{},
		ComputeRouteMatrix: []gax.CallOption{},
	}
}

// internalRoutesClient is an interface that defines the methods available from Routes API.
type internalRoutesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ComputeRoutes(context.Context, *routingpb.ComputeRoutesRequest, ...gax.CallOption) (*routingpb.ComputeRoutesResponse, error)
	ComputeRouteMatrix(context.Context, *routingpb.ComputeRouteMatrixRequest, ...gax.CallOption) (routingpb.Routes_ComputeRouteMatrixClient, error)
}

// RoutesClient is a client for interacting with Routes API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Routes API.
type RoutesClient struct {
	// The internal transport-dependent client.
	internalClient internalRoutesClient

	// The call options for this service.
	CallOptions *RoutesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RoutesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RoutesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RoutesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ComputeRoutes returns the primary route along with optional alternate routes, given a set
// of terminal and intermediate waypoints.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using URL parameter
// $fields or fields, or by using an HTTP/gRPC header X-Goog-FieldMask
// (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters)). The value
// is a comma separated list of field paths. See detailed documentation about
// how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//	Field mask of all available fields (for manual inspection):
//	X-Goog-FieldMask: *
//
//	Field mask of Route-level duration, distance, and polyline (an example
//	production setup):
//	X-Goog-FieldMask: routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline
//
// Google discourage the use of the wildcard (*) response field mask, or
// specifying the field mask at the top level (routes), because:
//
//	Selecting only the fields that you need helps our server save computation
//	cycles, allowing us to return the result to you with a lower latency.
//
//	Selecting only the fields that you need
//	in your production job ensures stable latency performance. We might add
//	more response fields in the future, and those new fields might require
//	extra computation time. If you select all fields, or if you select all
//	fields at the top level, then you might experience performance degradation
//	because any new field we add will be automatically included in the
//	response.
//
//	Selecting only the fields that you need results in a smaller response
//	size, and thus higher network throughput.
func (c *RoutesClient) ComputeRoutes(ctx context.Context, req *routingpb.ComputeRoutesRequest, opts ...gax.CallOption) (*routingpb.ComputeRoutesResponse, error) {
	return c.internalClient.ComputeRoutes(ctx, req, opts...)
}

// ComputeRouteMatrix takes in a list of origins and destinations and returns a stream containing
// route information for each combination of origin and destination.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using the URL
// parameter $fields or fields, or by using the HTTP/gRPC header
// X-Goog-FieldMask (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters)).
// The value is a comma separated list of field paths. See this detailed
// documentation about how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//	Field mask of all available fields (for manual inspection):
//	X-Goog-FieldMask: *
//
//	Field mask of route durations, distances, element status, condition, and
//	element indices (an example production setup):
//	X-Goog-FieldMask: originIndex,destinationIndex,status,condition,distanceMeters,duration
//
// It is critical that you include status in your field mask as otherwise
// all messages will appear to be OK. Google discourages the use of the
// wildcard (*) response field mask, because:
//
//	Selecting only the fields that you need helps our server save computation
//	cycles, allowing us to return the result to you with a lower latency.
//
//	Selecting only the fields that you need in your production job ensures
//	stable latency performance. We might add more response fields in the
//	future, and those new fields might require extra computation time. If you
//	select all fields, or if you select all fields at the top level, then you
//	might experience performance degradation because any new field we add will
//	be automatically included in the response.
//
//	Selecting only the fields that you need results in a smaller response
//	size, and thus higher network throughput.
func (c *RoutesClient) ComputeRouteMatrix(ctx context.Context, req *routingpb.ComputeRouteMatrixRequest, opts ...gax.CallOption) (routingpb.Routes_ComputeRouteMatrixClient, error) {
	return c.internalClient.ComputeRouteMatrix(ctx, req, opts...)
}

// routesGRPCClient is a client for interacting with Routes API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type routesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RoutesClient
	CallOptions **RoutesCallOptions

	// The gRPC API client.
	routesClient routingpb.RoutesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewRoutesClient creates a new routes client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Routes API.
func NewRoutesClient(ctx context.Context, opts ...option.ClientOption) (*RoutesClient, error) {
	clientOpts := defaultRoutesGRPCClientOptions()
	if newRoutesClientHook != nil {
		hookOpts, err := newRoutesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RoutesClient{CallOptions: defaultRoutesCallOptions()}

	c := &routesGRPCClient{
		connPool:     connPool,
		routesClient: routingpb.NewRoutesClient(connPool),
		CallOptions:  &client.CallOptions,
		logger:       internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *routesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *routesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *routesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type routesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing RoutesClient
	CallOptions **RoutesCallOptions

	logger *slog.Logger
}

// NewRoutesRESTClient creates a new routes rest client.
//
// The Routes API.
func NewRoutesRESTClient(ctx context.Context, opts ...option.ClientOption) (*RoutesClient, error) {
	clientOpts := append(defaultRoutesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRoutesRESTCallOptions()
	c := &routesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &RoutesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRoutesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://routes.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://routes.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://routes.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://routes.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *routesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *routesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *routesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *routesGRPCClient) ComputeRoutes(ctx context.Context, req *routingpb.ComputeRoutesRequest, opts ...gax.CallOption) (*routingpb.ComputeRoutesResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ComputeRoutes[0:len((*c.CallOptions).ComputeRoutes):len((*c.CallOptions).ComputeRoutes)], opts...)
	var resp *routingpb.ComputeRoutesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.routesClient.ComputeRoutes, req, settings.GRPC, c.logger, "ComputeRoutes")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routesGRPCClient) ComputeRouteMatrix(ctx context.Context, req *routingpb.ComputeRouteMatrixRequest, opts ...gax.CallOption) (routingpb.Routes_ComputeRouteMatrixClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ComputeRouteMatrix[0:len((*c.CallOptions).ComputeRouteMatrix):len((*c.CallOptions).ComputeRouteMatrix)], opts...)
	var resp routingpb.Routes_ComputeRouteMatrixClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		c.logger.DebugContext(ctx, "api streaming client request", "serviceName", serviceName, "rpcName", "ComputeRouteMatrix")
		resp, err = c.routesClient.ComputeRouteMatrix(ctx, req, settings.GRPC...)
		c.logger.DebugContext(ctx, "api streaming client response", "serviceName", serviceName, "rpcName", "ComputeRouteMatrix")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ComputeRoutes returns the primary route along with optional alternate routes, given a set
// of terminal and intermediate waypoints.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using URL parameter
// $fields or fields, or by using an HTTP/gRPC header X-Goog-FieldMask
// (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters)). The value
// is a comma separated list of field paths. See detailed documentation about
// how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//	Field mask of all available fields (for manual inspection):
//	X-Goog-FieldMask: *
//
//	Field mask of Route-level duration, distance, and polyline (an example
//	production setup):
//	X-Goog-FieldMask: routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline
//
// Google discourage the use of the wildcard (*) response field mask, or
// specifying the field mask at the top level (routes), because:
//
//	Selecting only the fields that you need helps our server save computation
//	cycles, allowing us to return the result to you with a lower latency.
//
//	Selecting only the fields that you need
//	in your production job ensures stable latency performance. We might add
//	more response fields in the future, and those new fields might require
//	extra computation time. If you select all fields, or if you select all
//	fields at the top level, then you might experience performance degradation
//	because any new field we add will be automatically included in the
//	response.
//
//	Selecting only the fields that you need results in a smaller response
//	size, and thus higher network throughput.
func (c *routesRESTClient) ComputeRoutes(ctx context.Context, req *routingpb.ComputeRoutesRequest, opts ...gax.CallOption) (*routingpb.ComputeRoutesResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/directions/v2:computeRoutes")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ComputeRoutes[0:len((*c.CallOptions).ComputeRoutes):len((*c.CallOptions).ComputeRoutes)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &routingpb.ComputeRoutesResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ComputeRoutes")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ComputeRouteMatrix takes in a list of origins and destinations and returns a stream containing
// route information for each combination of origin and destination.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using the URL
// parameter $fields or fields, or by using the HTTP/gRPC header
// X-Goog-FieldMask (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters)).
// The value is a comma separated list of field paths. See this detailed
// documentation about how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//	Field mask of all available fields (for manual inspection):
//	X-Goog-FieldMask: *
//
//	Field mask of route durations, distances, element status, condition, and
//	element indices (an example production setup):
//	X-Goog-FieldMask: originIndex,destinationIndex,status,condition,distanceMeters,duration
//
// It is critical that you include status in your field mask as otherwise
// all messages will appear to be OK. Google discourages the use of the
// wildcard (*) response field mask, because:
//
//	Selecting only the fields that you need helps our server save computation
//	cycles, allowing us to return the result to you with a lower latency.
//
//	Selecting only the fields that you need in your production job ensures
//	stable latency performance. We might add more response fields in the
//	future, and those new fields might require extra computation time. If you
//	select all fields, or if you select all fields at the top level, then you
//	might experience performance degradation because any new field we add will
//	be automatically included in the response.
//
//	Selecting only the fields that you need results in a smaller response
//	size, and thus higher network throughput.
func (c *routesRESTClient) ComputeRouteMatrix(ctx context.Context, req *routingpb.ComputeRouteMatrixRequest, opts ...gax.CallOption) (routingpb.Routes_ComputeRouteMatrixClient, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/distanceMatrix/v2:computeRouteMatrix")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	var streamClient *computeRouteMatrixRESTClient
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := executeStreamingHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ComputeRouteMatrix")
		if err != nil {
			return err
		}

		streamClient = &computeRouteMatrixRESTClient{
			ctx:    ctx,
			md:     metadata.MD(httpRsp.Header),
			stream: gax.NewProtoJSONStreamReader(httpRsp.Body, (&routingpb.RouteMatrixElement{}).ProtoReflect().Type()),
		}
		return nil
	}, opts...)

	return streamClient, e
}

// computeRouteMatrixRESTClient is the stream client used to consume the server stream created by
// the REST implementation of ComputeRouteMatrix.
type computeRouteMatrixRESTClient struct {
	ctx    context.Context
	md     metadata.MD
	stream *gax.ProtoJSONStream
}

func (c *computeRouteMatrixRESTClient) Recv() (*routingpb.RouteMatrixElement, error) {
	if err := c.ctx.Err(); err != nil {
		defer c.stream.Close()
		return nil, err
	}
	msg, err := c.stream.Recv()
	if err != nil {
		defer c.stream.Close()
		return nil, err
	}
	res := msg.(*routingpb.RouteMatrixElement)
	return res, nil
}

func (c *computeRouteMatrixRESTClient) Header() (metadata.MD, error) {
	return c.md, nil
}

func (c *computeRouteMatrixRESTClient) Trailer() metadata.MD {
	return c.md
}

func (c *computeRouteMatrixRESTClient) CloseSend() error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented for a server-stream")
}

func (c *computeRouteMatrixRESTClient) Context() context.Context {
	return c.ctx
}

func (c *computeRouteMatrixRESTClient) SendMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented for a server-stream")
}

func (c *computeRouteMatrixRESTClient) RecvMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented, use Recv")
}
