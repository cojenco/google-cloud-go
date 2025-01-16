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

package css

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	csspb "cloud.google.com/go/shopping/css/apiv1/csspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newAccountLabelsClientHook clientHook

// AccountLabelsCallOptions contains the retry settings for each method of AccountLabelsClient.
type AccountLabelsCallOptions struct {
	ListAccountLabels  []gax.CallOption
	CreateAccountLabel []gax.CallOption
	UpdateAccountLabel []gax.CallOption
	DeleteAccountLabel []gax.CallOption
}

func defaultAccountLabelsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("css.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("css.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("css.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://css.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAccountLabelsCallOptions() *AccountLabelsCallOptions {
	return &AccountLabelsCallOptions{
		ListAccountLabels: []gax.CallOption{
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
		CreateAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultAccountLabelsRESTCallOptions() *AccountLabelsCallOptions {
	return &AccountLabelsCallOptions{
		ListAccountLabels: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		CreateAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteAccountLabel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalAccountLabelsClient is an interface that defines the methods available from CSS API.
type internalAccountLabelsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAccountLabels(context.Context, *csspb.ListAccountLabelsRequest, ...gax.CallOption) *AccountLabelIterator
	CreateAccountLabel(context.Context, *csspb.CreateAccountLabelRequest, ...gax.CallOption) (*csspb.AccountLabel, error)
	UpdateAccountLabel(context.Context, *csspb.UpdateAccountLabelRequest, ...gax.CallOption) (*csspb.AccountLabel, error)
	DeleteAccountLabel(context.Context, *csspb.DeleteAccountLabelRequest, ...gax.CallOption) error
}

// AccountLabelsClient is a client for interacting with CSS API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages Merchant Center and CSS accounts labels.
type AccountLabelsClient struct {
	// The internal transport-dependent client.
	internalClient internalAccountLabelsClient

	// The call options for this service.
	CallOptions *AccountLabelsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AccountLabelsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AccountLabelsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AccountLabelsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAccountLabels lists the labels owned by an account.
func (c *AccountLabelsClient) ListAccountLabels(ctx context.Context, req *csspb.ListAccountLabelsRequest, opts ...gax.CallOption) *AccountLabelIterator {
	return c.internalClient.ListAccountLabels(ctx, req, opts...)
}

// CreateAccountLabel creates a new label, not assigned to any account.
func (c *AccountLabelsClient) CreateAccountLabel(ctx context.Context, req *csspb.CreateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	return c.internalClient.CreateAccountLabel(ctx, req, opts...)
}

// UpdateAccountLabel updates a label.
func (c *AccountLabelsClient) UpdateAccountLabel(ctx context.Context, req *csspb.UpdateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	return c.internalClient.UpdateAccountLabel(ctx, req, opts...)
}

// DeleteAccountLabel deletes a label and removes it from all accounts to which it was assigned.
func (c *AccountLabelsClient) DeleteAccountLabel(ctx context.Context, req *csspb.DeleteAccountLabelRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAccountLabel(ctx, req, opts...)
}

// accountLabelsGRPCClient is a client for interacting with CSS API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type accountLabelsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AccountLabelsClient
	CallOptions **AccountLabelsCallOptions

	// The gRPC API client.
	accountLabelsClient csspb.AccountLabelsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAccountLabelsClient creates a new account labels service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages Merchant Center and CSS accounts labels.
func NewAccountLabelsClient(ctx context.Context, opts ...option.ClientOption) (*AccountLabelsClient, error) {
	clientOpts := defaultAccountLabelsGRPCClientOptions()
	if newAccountLabelsClientHook != nil {
		hookOpts, err := newAccountLabelsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AccountLabelsClient{CallOptions: defaultAccountLabelsCallOptions()}

	c := &accountLabelsGRPCClient{
		connPool:            connPool,
		accountLabelsClient: csspb.NewAccountLabelsServiceClient(connPool),
		CallOptions:         &client.CallOptions,
		logger:              internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *accountLabelsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *accountLabelsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *accountLabelsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type accountLabelsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing AccountLabelsClient
	CallOptions **AccountLabelsCallOptions

	logger *slog.Logger
}

// NewAccountLabelsRESTClient creates a new account labels service rest client.
//
// Manages Merchant Center and CSS accounts labels.
func NewAccountLabelsRESTClient(ctx context.Context, opts ...option.ClientOption) (*AccountLabelsClient, error) {
	clientOpts := append(defaultAccountLabelsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultAccountLabelsRESTCallOptions()
	c := &accountLabelsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &AccountLabelsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultAccountLabelsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://css.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://css.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://css.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://css.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *accountLabelsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *accountLabelsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *accountLabelsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *accountLabelsGRPCClient) ListAccountLabels(ctx context.Context, req *csspb.ListAccountLabelsRequest, opts ...gax.CallOption) *AccountLabelIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAccountLabels[0:len((*c.CallOptions).ListAccountLabels):len((*c.CallOptions).ListAccountLabels)], opts...)
	it := &AccountLabelIterator{}
	req = proto.Clone(req).(*csspb.ListAccountLabelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*csspb.AccountLabel, string, error) {
		resp := &csspb.ListAccountLabelsResponse{}
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
			resp, err = executeRPC(ctx, c.accountLabelsClient.ListAccountLabels, req, settings.GRPC, c.logger, "ListAccountLabels")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAccountLabels(), resp.GetNextPageToken(), nil
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

func (c *accountLabelsGRPCClient) CreateAccountLabel(ctx context.Context, req *csspb.CreateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateAccountLabel[0:len((*c.CallOptions).CreateAccountLabel):len((*c.CallOptions).CreateAccountLabel)], opts...)
	var resp *csspb.AccountLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.accountLabelsClient.CreateAccountLabel, req, settings.GRPC, c.logger, "CreateAccountLabel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *accountLabelsGRPCClient) UpdateAccountLabel(ctx context.Context, req *csspb.UpdateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "account_label.name", url.QueryEscape(req.GetAccountLabel().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAccountLabel[0:len((*c.CallOptions).UpdateAccountLabel):len((*c.CallOptions).UpdateAccountLabel)], opts...)
	var resp *csspb.AccountLabel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.accountLabelsClient.UpdateAccountLabel, req, settings.GRPC, c.logger, "UpdateAccountLabel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *accountLabelsGRPCClient) DeleteAccountLabel(ctx context.Context, req *csspb.DeleteAccountLabelRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteAccountLabel[0:len((*c.CallOptions).DeleteAccountLabel):len((*c.CallOptions).DeleteAccountLabel)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.accountLabelsClient.DeleteAccountLabel, req, settings.GRPC, c.logger, "DeleteAccountLabel")
		return err
	}, opts...)
	return err
}

// ListAccountLabels lists the labels owned by an account.
func (c *accountLabelsRESTClient) ListAccountLabels(ctx context.Context, req *csspb.ListAccountLabelsRequest, opts ...gax.CallOption) *AccountLabelIterator {
	it := &AccountLabelIterator{}
	req = proto.Clone(req).(*csspb.ListAccountLabelsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*csspb.AccountLabel, string, error) {
		resp := &csspb.ListAccountLabelsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/labels", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListAccountLabels")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetAccountLabels(), resp.GetNextPageToken(), nil
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

// CreateAccountLabel creates a new label, not assigned to any account.
func (c *accountLabelsRESTClient) CreateAccountLabel(ctx context.Context, req *csspb.CreateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetAccountLabel()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/labels", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateAccountLabel[0:len((*c.CallOptions).CreateAccountLabel):len((*c.CallOptions).CreateAccountLabel)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &csspb.AccountLabel{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateAccountLabel")
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

// UpdateAccountLabel updates a label.
func (c *accountLabelsRESTClient) UpdateAccountLabel(ctx context.Context, req *csspb.UpdateAccountLabelRequest, opts ...gax.CallOption) (*csspb.AccountLabel, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetAccountLabel()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetAccountLabel().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "account_label.name", url.QueryEscape(req.GetAccountLabel().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAccountLabel[0:len((*c.CallOptions).UpdateAccountLabel):len((*c.CallOptions).UpdateAccountLabel)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &csspb.AccountLabel{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateAccountLabel")
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

// DeleteAccountLabel deletes a label and removes it from all accounts to which it was assigned.
func (c *accountLabelsRESTClient) DeleteAccountLabel(ctx context.Context, req *csspb.DeleteAccountLabelRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteAccountLabel")
		return err
	}, opts...)
}
