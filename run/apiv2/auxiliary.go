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

package run

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	runpb "cloud.google.com/go/run/apiv2/runpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// CancelExecutionOperation manages a long-running operation from CancelExecution.
type CancelExecutionOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CancelExecutionOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CancelExecutionOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CancelExecutionOperation) Metadata() (*runpb.Execution, error) {
	var meta runpb.Execution
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CancelExecutionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CancelExecutionOperation) Name() string {
	return op.lro.Name()
}

// CreateJobOperation manages a long-running operation from CreateJob.
type CreateJobOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateJobOperation) Metadata() (*runpb.Job, error) {
	var meta runpb.Job
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateJobOperation) Name() string {
	return op.lro.Name()
}

// CreateServiceOperation manages a long-running operation from CreateService.
type CreateServiceOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateServiceOperation) Name() string {
	return op.lro.Name()
}

// DeleteExecutionOperation manages a long-running operation from DeleteExecution.
type DeleteExecutionOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteExecutionOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteExecutionOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteExecutionOperation) Metadata() (*runpb.Execution, error) {
	var meta runpb.Execution
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteExecutionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteExecutionOperation) Name() string {
	return op.lro.Name()
}

// DeleteJobOperation manages a long-running operation from DeleteJob.
type DeleteJobOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteJobOperation) Metadata() (*runpb.Job, error) {
	var meta runpb.Job
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteJobOperation) Name() string {
	return op.lro.Name()
}

// DeleteRevisionOperation manages a long-running operation from DeleteRevision.
type DeleteRevisionOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteRevisionOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Revision, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Revision
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteRevisionOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Revision, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Revision
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteRevisionOperation) Metadata() (*runpb.Revision, error) {
	var meta runpb.Revision
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteRevisionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteRevisionOperation) Name() string {
	return op.lro.Name()
}

// DeleteServiceOperation manages a long-running operation from DeleteService.
type DeleteServiceOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteServiceOperation) Name() string {
	return op.lro.Name()
}

// RunJobOperation manages a long-running operation from RunJob.
type RunJobOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Execution, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Execution
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunJobOperation) Metadata() (*runpb.Execution, error) {
	var meta runpb.Execution
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunJobOperation) Name() string {
	return op.lro.Name()
}

// UpdateJobOperation manages a long-running operation from UpdateJob.
type UpdateJobOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Job, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Job
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateJobOperation) Metadata() (*runpb.Job, error) {
	var meta runpb.Job
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateJobOperation) Name() string {
	return op.lro.Name()
}

// UpdateServiceOperation manages a long-running operation from UpdateService.
type UpdateServiceOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateServiceOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateServiceOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*runpb.Service, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp runpb.Service
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateServiceOperation) Metadata() (*runpb.Service, error) {
	var meta runpb.Service
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateServiceOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateServiceOperation) Name() string {
	return op.lro.Name()
}

// ExecutionIterator manages a stream of *runpb.Execution.
type ExecutionIterator struct {
	items    []*runpb.Execution
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Execution, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ExecutionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ExecutionIterator) Next() (*runpb.Execution, error) {
	var item *runpb.Execution
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ExecutionIterator) bufLen() int {
	return len(it.items)
}

func (it *ExecutionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// JobIterator manages a stream of *runpb.Job.
type JobIterator struct {
	items    []*runpb.Job
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*runpb.Job, error) {
	var item *runpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RevisionIterator manages a stream of *runpb.Revision.
type RevisionIterator struct {
	items    []*runpb.Revision
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Revision, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *RevisionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RevisionIterator) Next() (*runpb.Revision, error) {
	var item *runpb.Revision
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RevisionIterator) bufLen() int {
	return len(it.items)
}

func (it *RevisionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ServiceIterator manages a stream of *runpb.Service.
type ServiceIterator struct {
	items    []*runpb.Service
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Service, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ServiceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceIterator) Next() (*runpb.Service, error) {
	var item *runpb.Service
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TaskIterator manages a stream of *runpb.Task.
type TaskIterator struct {
	items    []*runpb.Task
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*runpb.Task, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *TaskIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TaskIterator) Next() (*runpb.Task, error) {
	var item *runpb.Task
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TaskIterator) bufLen() int {
	return len(it.items)
}

func (it *TaskIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
