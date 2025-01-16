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

package meet

import (
	meetpb "cloud.google.com/go/apps/meet/apiv2beta/meetpb"
	"google.golang.org/api/iterator"
)

// ConferenceRecordIterator manages a stream of *meetpb.ConferenceRecord.
type ConferenceRecordIterator struct {
	items    []*meetpb.ConferenceRecord
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.ConferenceRecord, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ConferenceRecordIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ConferenceRecordIterator) Next() (*meetpb.ConferenceRecord, error) {
	var item *meetpb.ConferenceRecord
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ConferenceRecordIterator) bufLen() int {
	return len(it.items)
}

func (it *ConferenceRecordIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ParticipantIterator manages a stream of *meetpb.Participant.
type ParticipantIterator struct {
	items    []*meetpb.Participant
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.Participant, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ParticipantIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ParticipantIterator) Next() (*meetpb.Participant, error) {
	var item *meetpb.Participant
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ParticipantIterator) bufLen() int {
	return len(it.items)
}

func (it *ParticipantIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ParticipantSessionIterator manages a stream of *meetpb.ParticipantSession.
type ParticipantSessionIterator struct {
	items    []*meetpb.ParticipantSession
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.ParticipantSession, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *ParticipantSessionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ParticipantSessionIterator) Next() (*meetpb.ParticipantSession, error) {
	var item *meetpb.ParticipantSession
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ParticipantSessionIterator) bufLen() int {
	return len(it.items)
}

func (it *ParticipantSessionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RecordingIterator manages a stream of *meetpb.Recording.
type RecordingIterator struct {
	items    []*meetpb.Recording
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.Recording, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *RecordingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RecordingIterator) Next() (*meetpb.Recording, error) {
	var item *meetpb.Recording
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RecordingIterator) bufLen() int {
	return len(it.items)
}

func (it *RecordingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TranscriptEntryIterator manages a stream of *meetpb.TranscriptEntry.
type TranscriptEntryIterator struct {
	items    []*meetpb.TranscriptEntry
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.TranscriptEntry, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *TranscriptEntryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TranscriptEntryIterator) Next() (*meetpb.TranscriptEntry, error) {
	var item *meetpb.TranscriptEntry
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TranscriptEntryIterator) bufLen() int {
	return len(it.items)
}

func (it *TranscriptEntryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TranscriptIterator manages a stream of *meetpb.Transcript.
type TranscriptIterator struct {
	items    []*meetpb.Transcript
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
	InternalFetch func(pageSize int, pageToken string) (results []*meetpb.Transcript, nextPageToken string, err error)
}

// PageInfo supports pagination. See the [google.golang.org/api/iterator] package for details.
func (it *TranscriptIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TranscriptIterator) Next() (*meetpb.Transcript, error) {
	var item *meetpb.Transcript
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TranscriptIterator) bufLen() int {
	return len(it.items)
}

func (it *TranscriptIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
