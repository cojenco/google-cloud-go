// Copyright 2022 Google LLC
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

package retail_test

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2"
	"google.golang.org/api/iterator"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewCatalogClient() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCatalogClient_ListCatalogs() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ListCatalogsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#ListCatalogsRequest.
	}
	it := c.ListCatalogs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCatalogClient_UpdateCatalog() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateCatalogRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#UpdateCatalogRequest.
	}
	resp, err := c.UpdateCatalog(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_SetDefaultBranch() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.SetDefaultBranchRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#SetDefaultBranchRequest.
	}
	err = c.SetDefaultBranch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCatalogClient_GetDefaultBranch() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetDefaultBranchRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#GetDefaultBranchRequest.
	}
	resp, err := c.GetDefaultBranch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_GetCompletionConfig() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetCompletionConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#GetCompletionConfigRequest.
	}
	resp, err := c.GetCompletionConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_UpdateCompletionConfig() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateCompletionConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#UpdateCompletionConfigRequest.
	}
	resp, err := c.UpdateCompletionConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_GetAttributesConfig() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetAttributesConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#GetAttributesConfigRequest.
	}
	resp, err := c.GetAttributesConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_UpdateAttributesConfig() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateAttributesConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#UpdateAttributesConfigRequest.
	}
	resp, err := c.UpdateAttributesConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_AddCatalogAttribute() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.AddCatalogAttributeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#AddCatalogAttributeRequest.
	}
	resp, err := c.AddCatalogAttribute(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_RemoveCatalogAttribute() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.RemoveCatalogAttributeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#RemoveCatalogAttributeRequest.
	}
	resp, err := c.RemoveCatalogAttribute(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_ReplaceCatalogAttribute() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ReplaceCatalogAttributeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#ReplaceCatalogAttributeRequest.
	}
	resp, err := c.ReplaceCatalogAttribute(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_GetOperation() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_ListOperations() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
