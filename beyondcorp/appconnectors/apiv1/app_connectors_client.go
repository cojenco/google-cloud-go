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

package appconnectors

import (
	"context"
	"fmt"
	"math"
	"net/url"

	appconnectorspb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListAppConnectors  []gax.CallOption
	GetAppConnector    []gax.CallOption
	CreateAppConnector []gax.CallOption
	UpdateAppConnector []gax.CallOption
	DeleteAppConnector []gax.CallOption
	ReportStatus       []gax.CallOption
	GetLocation        []gax.CallOption
	ListLocations      []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
	CancelOperation    []gax.CallOption
	DeleteOperation    []gax.CallOption
	GetOperation       []gax.CallOption
	ListOperations     []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("beyondcorp.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("beyondcorp.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://beyondcorp.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListAppConnectors:  []gax.CallOption{},
		GetAppConnector:    []gax.CallOption{},
		CreateAppConnector: []gax.CallOption{},
		UpdateAppConnector: []gax.CallOption{},
		DeleteAppConnector: []gax.CallOption{},
		ReportStatus:       []gax.CallOption{},
		GetLocation:        []gax.CallOption{},
		ListLocations:      []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
		CancelOperation:    []gax.CallOption{},
		DeleteOperation:    []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
		ListOperations:     []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from BeyondCorp API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAppConnectors(context.Context, *appconnectorspb.ListAppConnectorsRequest, ...gax.CallOption) *AppConnectorIterator
	GetAppConnector(context.Context, *appconnectorspb.GetAppConnectorRequest, ...gax.CallOption) (*appconnectorspb.AppConnector, error)
	CreateAppConnector(context.Context, *appconnectorspb.CreateAppConnectorRequest, ...gax.CallOption) (*CreateAppConnectorOperation, error)
	CreateAppConnectorOperation(name string) *CreateAppConnectorOperation
	UpdateAppConnector(context.Context, *appconnectorspb.UpdateAppConnectorRequest, ...gax.CallOption) (*UpdateAppConnectorOperation, error)
	UpdateAppConnectorOperation(name string) *UpdateAppConnectorOperation
	DeleteAppConnector(context.Context, *appconnectorspb.DeleteAppConnectorRequest, ...gax.CallOption) (*DeleteAppConnectorOperation, error)
	DeleteAppConnectorOperation(name string) *DeleteAppConnectorOperation
	ReportStatus(context.Context, *appconnectorspb.ReportStatusRequest, ...gax.CallOption) (*ReportStatusOperation, error)
	ReportStatusOperation(name string) *ReportStatusOperation
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// Client is a client for interacting with BeyondCorp API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API Overview:
//
// The beyondcorp.googleapis.com service implements the Google Cloud
// BeyondCorp API.
//
// Data Model:
//
// The AppConnectorsService exposes the following resource:
//
//	AppConnectors, named as follows:
//	projects/{project_id}/locations/{location_id}/appConnectors/{app_connector_id}.
//
// The AppConnectorsService provides methods to manage
// (create/read/update/delete) BeyondCorp AppConnectors.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAppConnectors lists AppConnectors in a given project and location.
func (c *Client) ListAppConnectors(ctx context.Context, req *appconnectorspb.ListAppConnectorsRequest, opts ...gax.CallOption) *AppConnectorIterator {
	return c.internalClient.ListAppConnectors(ctx, req, opts...)
}

// GetAppConnector gets details of a single AppConnector.
func (c *Client) GetAppConnector(ctx context.Context, req *appconnectorspb.GetAppConnectorRequest, opts ...gax.CallOption) (*appconnectorspb.AppConnector, error) {
	return c.internalClient.GetAppConnector(ctx, req, opts...)
}

// CreateAppConnector creates a new AppConnector in a given project and location.
func (c *Client) CreateAppConnector(ctx context.Context, req *appconnectorspb.CreateAppConnectorRequest, opts ...gax.CallOption) (*CreateAppConnectorOperation, error) {
	return c.internalClient.CreateAppConnector(ctx, req, opts...)
}

// CreateAppConnectorOperation returns a new CreateAppConnectorOperation from a given name.
// The name must be that of a previously created CreateAppConnectorOperation, possibly from a different process.
func (c *Client) CreateAppConnectorOperation(name string) *CreateAppConnectorOperation {
	return c.internalClient.CreateAppConnectorOperation(name)
}

// UpdateAppConnector updates the parameters of a single AppConnector.
func (c *Client) UpdateAppConnector(ctx context.Context, req *appconnectorspb.UpdateAppConnectorRequest, opts ...gax.CallOption) (*UpdateAppConnectorOperation, error) {
	return c.internalClient.UpdateAppConnector(ctx, req, opts...)
}

// UpdateAppConnectorOperation returns a new UpdateAppConnectorOperation from a given name.
// The name must be that of a previously created UpdateAppConnectorOperation, possibly from a different process.
func (c *Client) UpdateAppConnectorOperation(name string) *UpdateAppConnectorOperation {
	return c.internalClient.UpdateAppConnectorOperation(name)
}

// DeleteAppConnector deletes a single AppConnector.
func (c *Client) DeleteAppConnector(ctx context.Context, req *appconnectorspb.DeleteAppConnectorRequest, opts ...gax.CallOption) (*DeleteAppConnectorOperation, error) {
	return c.internalClient.DeleteAppConnector(ctx, req, opts...)
}

// DeleteAppConnectorOperation returns a new DeleteAppConnectorOperation from a given name.
// The name must be that of a previously created DeleteAppConnectorOperation, possibly from a different process.
func (c *Client) DeleteAppConnectorOperation(name string) *DeleteAppConnectorOperation {
	return c.internalClient.DeleteAppConnectorOperation(name)
}

// ReportStatus report status for a given connector.
func (c *Client) ReportStatus(ctx context.Context, req *appconnectorspb.ReportStatusRequest, opts ...gax.CallOption) (*ReportStatusOperation, error) {
	return c.internalClient.ReportStatus(ctx, req, opts...)
}

// ReportStatusOperation returns a new ReportStatusOperation from a given name.
// The name must be that of a previously created ReportStatusOperation, possibly from a different process.
func (c *Client) ReportStatusOperation(name string) *ReportStatusOperation {
	return c.internalClient.ReportStatusOperation(name)
}

// GetLocation gets information about a location.
func (c *Client) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *Client) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *Client) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *Client) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *Client) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *Client) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// gRPCClient is a client for interacting with BeyondCorp API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client appconnectorspb.AppConnectorsServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new app connectors service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API Overview:
//
// The beyondcorp.googleapis.com service implements the Google Cloud
// BeyondCorp API.
//
// Data Model:
//
// The AppConnectorsService exposes the following resource:
//
//	AppConnectors, named as follows:
//	projects/{project_id}/locations/{location_id}/appConnectors/{app_connector_id}.
//
// The AppConnectorsService provides methods to manage
// (create/read/update/delete) BeyondCorp AppConnectors.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		client:           appconnectorspb.NewAppConnectorsServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:  iampb.NewIAMPolicyClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListAppConnectors(ctx context.Context, req *appconnectorspb.ListAppConnectorsRequest, opts ...gax.CallOption) *AppConnectorIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAppConnectors[0:len((*c.CallOptions).ListAppConnectors):len((*c.CallOptions).ListAppConnectors)], opts...)
	it := &AppConnectorIterator{}
	req = proto.Clone(req).(*appconnectorspb.ListAppConnectorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appconnectorspb.AppConnector, string, error) {
		resp := &appconnectorspb.ListAppConnectorsResponse{}
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
			resp, err = c.client.ListAppConnectors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAppConnectors(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetAppConnector(ctx context.Context, req *appconnectorspb.GetAppConnectorRequest, opts ...gax.CallOption) (*appconnectorspb.AppConnector, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAppConnector[0:len((*c.CallOptions).GetAppConnector):len((*c.CallOptions).GetAppConnector)], opts...)
	var resp *appconnectorspb.AppConnector
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAppConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateAppConnector(ctx context.Context, req *appconnectorspb.CreateAppConnectorRequest, opts ...gax.CallOption) (*CreateAppConnectorOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateAppConnector[0:len((*c.CallOptions).CreateAppConnector):len((*c.CallOptions).CreateAppConnector)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAppConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UpdateAppConnector(ctx context.Context, req *appconnectorspb.UpdateAppConnectorRequest, opts ...gax.CallOption) (*UpdateAppConnectorOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "app_connector.name", url.QueryEscape(req.GetAppConnector().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAppConnector[0:len((*c.CallOptions).UpdateAppConnector):len((*c.CallOptions).UpdateAppConnector)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateAppConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteAppConnector(ctx context.Context, req *appconnectorspb.DeleteAppConnectorRequest, opts ...gax.CallOption) (*DeleteAppConnectorOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteAppConnector[0:len((*c.CallOptions).DeleteAppConnector):len((*c.CallOptions).DeleteAppConnector)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteAppConnector(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ReportStatus(ctx context.Context, req *appconnectorspb.ReportStatusRequest, opts ...gax.CallOption) (*ReportStatusOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "app_connector", url.QueryEscape(req.GetAppConnector()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ReportStatus[0:len((*c.CallOptions).ReportStatus):len((*c.CallOptions).ReportStatus)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportStatus(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ReportStatusOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
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

func (c *gRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
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

func (c *gRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
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

func (c *gRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
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

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
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

func (c *gRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
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

// CreateAppConnectorOperation returns a new CreateAppConnectorOperation from a given name.
// The name must be that of a previously created CreateAppConnectorOperation, possibly from a different process.
func (c *gRPCClient) CreateAppConnectorOperation(name string) *CreateAppConnectorOperation {
	return &CreateAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteAppConnectorOperation returns a new DeleteAppConnectorOperation from a given name.
// The name must be that of a previously created DeleteAppConnectorOperation, possibly from a different process.
func (c *gRPCClient) DeleteAppConnectorOperation(name string) *DeleteAppConnectorOperation {
	return &DeleteAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// ReportStatusOperation returns a new ReportStatusOperation from a given name.
// The name must be that of a previously created ReportStatusOperation, possibly from a different process.
func (c *gRPCClient) ReportStatusOperation(name string) *ReportStatusOperation {
	return &ReportStatusOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateAppConnectorOperation returns a new UpdateAppConnectorOperation from a given name.
// The name must be that of a previously created UpdateAppConnectorOperation, possibly from a different process.
func (c *gRPCClient) UpdateAppConnectorOperation(name string) *UpdateAppConnectorOperation {
	return &UpdateAppConnectorOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
