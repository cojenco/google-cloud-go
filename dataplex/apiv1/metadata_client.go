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

package dataplex

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	dataplexpb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newMetadataClientHook clientHook

// MetadataCallOptions contains the retry settings for each method of MetadataClient.
type MetadataCallOptions struct {
	CreateEntity    []gax.CallOption
	UpdateEntity    []gax.CallOption
	DeleteEntity    []gax.CallOption
	GetEntity       []gax.CallOption
	ListEntities    []gax.CallOption
	CreatePartition []gax.CallOption
	DeletePartition []gax.CallOption
	GetPartition    []gax.CallOption
	ListPartitions  []gax.CallOption
	GetLocation     []gax.CallOption
	ListLocations   []gax.CallOption
	CancelOperation []gax.CallOption
	DeleteOperation []gax.CallOption
	GetOperation    []gax.CallOption
	ListOperations  []gax.CallOption
}

func defaultMetadataGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataplex.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataplex.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataplex.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetadataCallOptions() *MetadataCallOptions {
	return &MetadataCallOptions{
		CreateEntity: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateEntity: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteEntity: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetEntity: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListEntities: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreatePartition: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeletePartition: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetPartition: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListPartitions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetLocation:     []gax.CallOption{},
		ListLocations:   []gax.CallOption{},
		CancelOperation: []gax.CallOption{},
		DeleteOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalMetadataClient is an interface that defines the methods available from Cloud Dataplex API.
type internalMetadataClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateEntity(context.Context, *dataplexpb.CreateEntityRequest, ...gax.CallOption) (*dataplexpb.Entity, error)
	UpdateEntity(context.Context, *dataplexpb.UpdateEntityRequest, ...gax.CallOption) (*dataplexpb.Entity, error)
	DeleteEntity(context.Context, *dataplexpb.DeleteEntityRequest, ...gax.CallOption) error
	GetEntity(context.Context, *dataplexpb.GetEntityRequest, ...gax.CallOption) (*dataplexpb.Entity, error)
	ListEntities(context.Context, *dataplexpb.ListEntitiesRequest, ...gax.CallOption) *EntityIterator
	CreatePartition(context.Context, *dataplexpb.CreatePartitionRequest, ...gax.CallOption) (*dataplexpb.Partition, error)
	DeletePartition(context.Context, *dataplexpb.DeletePartitionRequest, ...gax.CallOption) error
	GetPartition(context.Context, *dataplexpb.GetPartitionRequest, ...gax.CallOption) (*dataplexpb.Partition, error)
	ListPartitions(context.Context, *dataplexpb.ListPartitionsRequest, ...gax.CallOption) *PartitionIterator
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// MetadataClient is a client for interacting with Cloud Dataplex API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Metadata service manages metadata resources such as tables, filesets and
// partitions.
type MetadataClient struct {
	// The internal transport-dependent client.
	internalClient internalMetadataClient

	// The call options for this service.
	CallOptions *MetadataCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetadataClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetadataClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MetadataClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateEntity create a metadata entity.
func (c *MetadataClient) CreateEntity(ctx context.Context, req *dataplexpb.CreateEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	return c.internalClient.CreateEntity(ctx, req, opts...)
}

// UpdateEntity update a metadata entity. Only supports full resource update.
func (c *MetadataClient) UpdateEntity(ctx context.Context, req *dataplexpb.UpdateEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	return c.internalClient.UpdateEntity(ctx, req, opts...)
}

// DeleteEntity delete a metadata entity.
func (c *MetadataClient) DeleteEntity(ctx context.Context, req *dataplexpb.DeleteEntityRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteEntity(ctx, req, opts...)
}

// GetEntity get a metadata entity.
func (c *MetadataClient) GetEntity(ctx context.Context, req *dataplexpb.GetEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	return c.internalClient.GetEntity(ctx, req, opts...)
}

// ListEntities list metadata entities in a zone.
func (c *MetadataClient) ListEntities(ctx context.Context, req *dataplexpb.ListEntitiesRequest, opts ...gax.CallOption) *EntityIterator {
	return c.internalClient.ListEntities(ctx, req, opts...)
}

// CreatePartition create a metadata partition.
func (c *MetadataClient) CreatePartition(ctx context.Context, req *dataplexpb.CreatePartitionRequest, opts ...gax.CallOption) (*dataplexpb.Partition, error) {
	return c.internalClient.CreatePartition(ctx, req, opts...)
}

// DeletePartition delete a metadata partition.
func (c *MetadataClient) DeletePartition(ctx context.Context, req *dataplexpb.DeletePartitionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePartition(ctx, req, opts...)
}

// GetPartition get a metadata partition of an entity.
func (c *MetadataClient) GetPartition(ctx context.Context, req *dataplexpb.GetPartitionRequest, opts ...gax.CallOption) (*dataplexpb.Partition, error) {
	return c.internalClient.GetPartition(ctx, req, opts...)
}

// ListPartitions list metadata partitions of an entity.
func (c *MetadataClient) ListPartitions(ctx context.Context, req *dataplexpb.ListPartitionsRequest, opts ...gax.CallOption) *PartitionIterator {
	return c.internalClient.ListPartitions(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *MetadataClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *MetadataClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *MetadataClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *MetadataClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *MetadataClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *MetadataClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// metadataGRPCClient is a client for interacting with Cloud Dataplex API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metadataGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MetadataClient
	CallOptions **MetadataCallOptions

	// The gRPC API client.
	metadataClient dataplexpb.MetadataServiceClient

	operationsClient longrunningpb.OperationsClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewMetadataClient creates a new metadata service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Metadata service manages metadata resources such as tables, filesets and
// partitions.
func NewMetadataClient(ctx context.Context, opts ...option.ClientOption) (*MetadataClient, error) {
	clientOpts := defaultMetadataGRPCClientOptions()
	if newMetadataClientHook != nil {
		hookOpts, err := newMetadataClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MetadataClient{CallOptions: defaultMetadataCallOptions()}

	c := &metadataGRPCClient{
		connPool:         connPool,
		metadataClient:   dataplexpb.NewMetadataServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *metadataGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metadataGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metadataGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *metadataGRPCClient) CreateEntity(ctx context.Context, req *dataplexpb.CreateEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateEntity[0:len((*c.CallOptions).CreateEntity):len((*c.CallOptions).CreateEntity)], opts...)
	var resp *dataplexpb.Entity
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metadataClient.CreateEntity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) UpdateEntity(ctx context.Context, req *dataplexpb.UpdateEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "entity.name", url.QueryEscape(req.GetEntity().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateEntity[0:len((*c.CallOptions).UpdateEntity):len((*c.CallOptions).UpdateEntity)], opts...)
	var resp *dataplexpb.Entity
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metadataClient.UpdateEntity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) DeleteEntity(ctx context.Context, req *dataplexpb.DeleteEntityRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteEntity[0:len((*c.CallOptions).DeleteEntity):len((*c.CallOptions).DeleteEntity)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metadataClient.DeleteEntity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metadataGRPCClient) GetEntity(ctx context.Context, req *dataplexpb.GetEntityRequest, opts ...gax.CallOption) (*dataplexpb.Entity, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetEntity[0:len((*c.CallOptions).GetEntity):len((*c.CallOptions).GetEntity)], opts...)
	var resp *dataplexpb.Entity
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metadataClient.GetEntity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) ListEntities(ctx context.Context, req *dataplexpb.ListEntitiesRequest, opts ...gax.CallOption) *EntityIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListEntities[0:len((*c.CallOptions).ListEntities):len((*c.CallOptions).ListEntities)], opts...)
	it := &EntityIterator{}
	req = proto.Clone(req).(*dataplexpb.ListEntitiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataplexpb.Entity, string, error) {
		resp := &dataplexpb.ListEntitiesResponse{}
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
			resp, err = c.metadataClient.ListEntities(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetEntities(), resp.GetNextPageToken(), nil
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

func (c *metadataGRPCClient) CreatePartition(ctx context.Context, req *dataplexpb.CreatePartitionRequest, opts ...gax.CallOption) (*dataplexpb.Partition, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreatePartition[0:len((*c.CallOptions).CreatePartition):len((*c.CallOptions).CreatePartition)], opts...)
	var resp *dataplexpb.Partition
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metadataClient.CreatePartition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) DeletePartition(ctx context.Context, req *dataplexpb.DeletePartitionRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeletePartition[0:len((*c.CallOptions).DeletePartition):len((*c.CallOptions).DeletePartition)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metadataClient.DeletePartition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metadataGRPCClient) GetPartition(ctx context.Context, req *dataplexpb.GetPartitionRequest, opts ...gax.CallOption) (*dataplexpb.Partition, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPartition[0:len((*c.CallOptions).GetPartition):len((*c.CallOptions).GetPartition)], opts...)
	var resp *dataplexpb.Partition
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metadataClient.GetPartition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) ListPartitions(ctx context.Context, req *dataplexpb.ListPartitionsRequest, opts ...gax.CallOption) *PartitionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListPartitions[0:len((*c.CallOptions).ListPartitions):len((*c.CallOptions).ListPartitions)], opts...)
	it := &PartitionIterator{}
	req = proto.Clone(req).(*dataplexpb.ListPartitionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataplexpb.Partition, string, error) {
		resp := &dataplexpb.ListPartitionsResponse{}
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
			resp, err = c.metadataClient.ListPartitions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPartitions(), resp.GetNextPageToken(), nil
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

func (c *metadataGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *metadataGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metadataGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metadataGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metadataGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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
