// Copyright 2024 Google LLC
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

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newSnoozeClientHook clientHook

// SnoozeCallOptions contains the retry settings for each method of SnoozeClient.
type SnoozeCallOptions struct {
	CreateSnooze []gax.CallOption
	ListSnoozes  []gax.CallOption
	GetSnooze    []gax.CallOption
	UpdateSnooze []gax.CallOption
}

func defaultSnoozeGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSnoozeCallOptions() *SnoozeCallOptions {
	return &SnoozeCallOptions{
		CreateSnooze: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		ListSnoozes: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSnooze: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateSnooze: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
	}
}

// internalSnoozeClient is an interface that defines the methods available from Cloud Monitoring API.
type internalSnoozeClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateSnooze(context.Context, *monitoringpb.CreateSnoozeRequest, ...gax.CallOption) (*monitoringpb.Snooze, error)
	ListSnoozes(context.Context, *monitoringpb.ListSnoozesRequest, ...gax.CallOption) *SnoozeIterator
	GetSnooze(context.Context, *monitoringpb.GetSnoozeRequest, ...gax.CallOption) (*monitoringpb.Snooze, error)
	UpdateSnooze(context.Context, *monitoringpb.UpdateSnoozeRequest, ...gax.CallOption) (*monitoringpb.Snooze, error)
}

// SnoozeClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The SnoozeService API is used to temporarily prevent an alert policy from
// generating alerts. A Snooze is a description of the criteria under which one
// or more alert policies should not fire alerts for the specified duration.
type SnoozeClient struct {
	// The internal transport-dependent client.
	internalClient internalSnoozeClient

	// The call options for this service.
	CallOptions *SnoozeCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SnoozeClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SnoozeClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SnoozeClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateSnooze creates a Snooze that will prevent alerts, which match the provided
// criteria, from being opened. The Snooze applies for a specific time
// interval.
func (c *SnoozeClient) CreateSnooze(ctx context.Context, req *monitoringpb.CreateSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	return c.internalClient.CreateSnooze(ctx, req, opts...)
}

// ListSnoozes lists the Snoozes associated with a project. Can optionally pass in
// filter, which specifies predicates to match Snoozes.
func (c *SnoozeClient) ListSnoozes(ctx context.Context, req *monitoringpb.ListSnoozesRequest, opts ...gax.CallOption) *SnoozeIterator {
	return c.internalClient.ListSnoozes(ctx, req, opts...)
}

// GetSnooze retrieves a Snooze by name.
func (c *SnoozeClient) GetSnooze(ctx context.Context, req *monitoringpb.GetSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	return c.internalClient.GetSnooze(ctx, req, opts...)
}

// UpdateSnooze updates a Snooze, identified by its name, with the parameters in the
// given Snooze object.
func (c *SnoozeClient) UpdateSnooze(ctx context.Context, req *monitoringpb.UpdateSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	return c.internalClient.UpdateSnooze(ctx, req, opts...)
}

// snoozeGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type snoozeGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SnoozeClient
	CallOptions **SnoozeCallOptions

	// The gRPC API client.
	snoozeClient monitoringpb.SnoozeServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewSnoozeClient creates a new snooze service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The SnoozeService API is used to temporarily prevent an alert policy from
// generating alerts. A Snooze is a description of the criteria under which one
// or more alert policies should not fire alerts for the specified duration.
func NewSnoozeClient(ctx context.Context, opts ...option.ClientOption) (*SnoozeClient, error) {
	clientOpts := defaultSnoozeGRPCClientOptions()
	if newSnoozeClientHook != nil {
		hookOpts, err := newSnoozeClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SnoozeClient{CallOptions: defaultSnoozeCallOptions()}

	c := &snoozeGRPCClient{
		connPool:     connPool,
		snoozeClient: monitoringpb.NewSnoozeServiceClient(connPool),
		CallOptions:  &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *snoozeGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *snoozeGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *snoozeGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *snoozeGRPCClient) CreateSnooze(ctx context.Context, req *monitoringpb.CreateSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateSnooze[0:len((*c.CallOptions).CreateSnooze):len((*c.CallOptions).CreateSnooze)], opts...)
	var resp *monitoringpb.Snooze
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snoozeClient.CreateSnooze(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snoozeGRPCClient) ListSnoozes(ctx context.Context, req *monitoringpb.ListSnoozesRequest, opts ...gax.CallOption) *SnoozeIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSnoozes[0:len((*c.CallOptions).ListSnoozes):len((*c.CallOptions).ListSnoozes)], opts...)
	it := &SnoozeIterator{}
	req = proto.Clone(req).(*monitoringpb.ListSnoozesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.Snooze, string, error) {
		resp := &monitoringpb.ListSnoozesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.snoozeClient.ListSnoozes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSnoozes(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *snoozeGRPCClient) GetSnooze(ctx context.Context, req *monitoringpb.GetSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSnooze[0:len((*c.CallOptions).GetSnooze):len((*c.CallOptions).GetSnooze)], opts...)
	var resp *monitoringpb.Snooze
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snoozeClient.GetSnooze(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *snoozeGRPCClient) UpdateSnooze(ctx context.Context, req *monitoringpb.UpdateSnoozeRequest, opts ...gax.CallOption) (*monitoringpb.Snooze, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "snooze.name", url.QueryEscape(req.GetSnooze().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateSnooze[0:len((*c.CallOptions).UpdateSnooze):len((*c.CallOptions).UpdateSnooze)], opts...)
	var resp *monitoringpb.Snooze
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.snoozeClient.UpdateSnooze(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
