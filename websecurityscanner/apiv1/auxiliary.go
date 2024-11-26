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

package websecurityscanner

import (
	websecurityscannerpb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"google.golang.org/api/iterator"
)

// CrawledUrlIterator manages a stream of *websecurityscannerpb.CrawledUrl.
type CrawledUrlIterator struct {
	items    []*websecurityscannerpb.CrawledUrl
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.CrawledUrl, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *CrawledUrlIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CrawledUrlIterator) Next() (*websecurityscannerpb.CrawledUrl, error) {
	var item *websecurityscannerpb.CrawledUrl
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CrawledUrlIterator) bufLen() int {
	return len(it.items)
}

func (it *CrawledUrlIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// FindingIterator manages a stream of *websecurityscannerpb.Finding.
type FindingIterator struct {
	items    []*websecurityscannerpb.Finding
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.Finding, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *FindingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FindingIterator) Next() (*websecurityscannerpb.Finding, error) {
	var item *websecurityscannerpb.Finding
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FindingIterator) bufLen() int {
	return len(it.items)
}

func (it *FindingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ScanConfigIterator manages a stream of *websecurityscannerpb.ScanConfig.
type ScanConfigIterator struct {
	items    []*websecurityscannerpb.ScanConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.ScanConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ScanConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ScanConfigIterator) Next() (*websecurityscannerpb.ScanConfig, error) {
	var item *websecurityscannerpb.ScanConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ScanConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *ScanConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ScanRunIterator manages a stream of *websecurityscannerpb.ScanRun.
type ScanRunIterator struct {
	items    []*websecurityscannerpb.ScanRun
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
	InternalFetch func(pageSize int, pageToken string) (results []*websecurityscannerpb.ScanRun, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ScanRunIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ScanRunIterator) Next() (*websecurityscannerpb.ScanRun, error) {
	var item *websecurityscannerpb.ScanRun
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ScanRunIterator) bufLen() int {
	return len(it.items)
}

func (it *ScanRunIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
