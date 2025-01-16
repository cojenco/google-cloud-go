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

//go:build go1.23

package deploy_test

import (
	"context"

	deploy "cloud.google.com/go/deploy/apiv1"
	deploypb "cloud.google.com/go/deploy/apiv1/deploypb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleCloudDeployClient_ListAutomationRuns_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListAutomationRunsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListAutomationRunsRequest.
	}
	for resp, err := range c.ListAutomationRuns(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListAutomations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListAutomationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListAutomationsRequest.
	}
	for resp, err := range c.ListAutomations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListCustomTargetTypes_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListCustomTargetTypesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListCustomTargetTypesRequest.
	}
	for resp, err := range c.ListCustomTargetTypes(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListDeliveryPipelines_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListDeliveryPipelinesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListDeliveryPipelinesRequest.
	}
	for resp, err := range c.ListDeliveryPipelines(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListDeployPolicies_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListDeployPoliciesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListDeployPoliciesRequest.
	}
	for resp, err := range c.ListDeployPolicies(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListJobRuns_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListJobRunsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListJobRunsRequest.
	}
	for resp, err := range c.ListJobRuns(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListReleases_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListReleasesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListReleasesRequest.
	}
	for resp, err := range c.ListReleases(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListRollouts_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListRolloutsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListRolloutsRequest.
	}
	for resp, err := range c.ListRollouts(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListTargets_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &deploypb.ListTargetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/deploy/apiv1/deploypb#ListTargetsRequest.
	}
	for resp, err := range c.ListTargets(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListLocations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	for resp, err := range c.ListLocations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudDeployClient_ListOperations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	for resp, err := range c.ListOperations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
