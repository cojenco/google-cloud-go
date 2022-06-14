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

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewSecuritySettingsClient() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleSecuritySettingsClient_CreateSecuritySettings() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.CreateSecuritySettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#CreateSecuritySettingsRequest.
	}
	resp, err := c.CreateSecuritySettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecuritySettingsClient_GetSecuritySettings() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.GetSecuritySettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#GetSecuritySettingsRequest.
	}
	resp, err := c.GetSecuritySettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecuritySettingsClient_UpdateSecuritySettings() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.UpdateSecuritySettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#UpdateSecuritySettingsRequest.
	}
	resp, err := c.UpdateSecuritySettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecuritySettingsClient_ListSecuritySettings() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListSecuritySettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#ListSecuritySettingsRequest.
	}
	it := c.ListSecuritySettings(ctx, req)
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

func ExampleSecuritySettingsClient_DeleteSecuritySettings() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.DeleteSecuritySettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#DeleteSecuritySettingsRequest.
	}
	err = c.DeleteSecuritySettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecuritySettingsClient_GetLocation() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecuritySettingsClient_ListLocations() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
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

func ExampleSecuritySettingsClient_CancelOperation() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSecuritySettingsClient_GetOperation() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
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

func ExampleSecuritySettingsClient_ListOperations() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
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
