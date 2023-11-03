// Copyright 2023 Google LLC
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

package gkemulticloud_test

import (
	"context"

	gkemulticloud "cloud.google.com/go/gkemulticloud/apiv1"
	gkemulticloudpb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/iterator"
)

func ExampleNewAttachedClustersClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAttachedClustersClient_CreateAttachedCluster() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.CreateAttachedClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#CreateAttachedClusterRequest.
	}
	op, err := c.CreateAttachedCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_DeleteAttachedCluster() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.DeleteAttachedClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#DeleteAttachedClusterRequest.
	}
	op, err := c.DeleteAttachedCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAttachedClustersClient_GenerateAttachedClusterInstallManifest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.GenerateAttachedClusterInstallManifestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#GenerateAttachedClusterInstallManifestRequest.
	}
	resp, err := c.GenerateAttachedClusterInstallManifest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_GetAttachedCluster() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.GetAttachedClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#GetAttachedClusterRequest.
	}
	resp, err := c.GetAttachedCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_GetAttachedServerConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.GetAttachedServerConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#GetAttachedServerConfigRequest.
	}
	resp, err := c.GetAttachedServerConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_ImportAttachedCluster() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.ImportAttachedClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#ImportAttachedClusterRequest.
	}
	op, err := c.ImportAttachedCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_ListAttachedClusters() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.ListAttachedClustersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#ListAttachedClustersRequest.
	}
	it := c.ListAttachedClusters(ctx, req)
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

func ExampleAttachedClustersClient_UpdateAttachedCluster() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &gkemulticloudpb.UpdateAttachedClusterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb#UpdateAttachedClusterRequest.
	}
	op, err := c.UpdateAttachedCluster(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAttachedClustersClient_DeleteOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.DeleteOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#DeleteOperationRequest.
	}
	err = c.DeleteOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAttachedClustersClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAttachedClustersClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := gkemulticloud.NewAttachedClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
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
