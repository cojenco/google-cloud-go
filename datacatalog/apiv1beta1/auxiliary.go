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

package datacatalog

import (
	datacatalogpb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	"google.golang.org/api/iterator"
)

// EntryGroupIterator manages a stream of *datacatalogpb.EntryGroup.
type EntryGroupIterator struct {
	items    []*datacatalogpb.EntryGroup
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.EntryGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *EntryGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EntryGroupIterator) Next() (*datacatalogpb.EntryGroup, error) {
	var item *datacatalogpb.EntryGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EntryGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *EntryGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// EntryIterator manages a stream of *datacatalogpb.Entry.
type EntryIterator struct {
	items    []*datacatalogpb.Entry
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.Entry, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *EntryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *EntryIterator) Next() (*datacatalogpb.Entry, error) {
	var item *datacatalogpb.Entry
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *EntryIterator) bufLen() int {
	return len(it.items)
}

func (it *EntryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PolicyTagIterator manages a stream of *datacatalogpb.PolicyTag.
type PolicyTagIterator struct {
	items    []*datacatalogpb.PolicyTag
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.PolicyTag, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *PolicyTagIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PolicyTagIterator) Next() (*datacatalogpb.PolicyTag, error) {
	var item *datacatalogpb.PolicyTag
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PolicyTagIterator) bufLen() int {
	return len(it.items)
}

func (it *PolicyTagIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SearchCatalogResultIterator manages a stream of *datacatalogpb.SearchCatalogResult.
type SearchCatalogResultIterator struct {
	items    []*datacatalogpb.SearchCatalogResult
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.SearchCatalogResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *SearchCatalogResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SearchCatalogResultIterator) Next() (*datacatalogpb.SearchCatalogResult, error) {
	var item *datacatalogpb.SearchCatalogResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SearchCatalogResultIterator) bufLen() int {
	return len(it.items)
}

func (it *SearchCatalogResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TagIterator manages a stream of *datacatalogpb.Tag.
type TagIterator struct {
	items    []*datacatalogpb.Tag
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.Tag, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *TagIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TagIterator) Next() (*datacatalogpb.Tag, error) {
	var item *datacatalogpb.Tag
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TagIterator) bufLen() int {
	return len(it.items)
}

func (it *TagIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TaxonomyIterator manages a stream of *datacatalogpb.Taxonomy.
type TaxonomyIterator struct {
	items    []*datacatalogpb.Taxonomy
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
	InternalFetch func(pageSize int, pageToken string) (results []*datacatalogpb.Taxonomy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *TaxonomyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TaxonomyIterator) Next() (*datacatalogpb.Taxonomy, error) {
	var item *datacatalogpb.Taxonomy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TaxonomyIterator) bufLen() int {
	return len(it.items)
}

func (it *TaxonomyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
