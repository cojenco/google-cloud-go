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

package migrationcenter

import (
	"iter"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	migrationcenterpb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	"github.com/googleapis/gax-go/v2/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AssetIterator) All() iter.Seq2[*migrationcenterpb.Asset, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ErrorFrameIterator) All() iter.Seq2[*migrationcenterpb.ErrorFrame, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *GroupIterator) All() iter.Seq2[*migrationcenterpb.Group, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ImportDataFileIterator) All() iter.Seq2[*migrationcenterpb.ImportDataFile, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ImportJobIterator) All() iter.Seq2[*migrationcenterpb.ImportJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LocationIterator) All() iter.Seq2[*locationpb.Location, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*longrunningpb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PreferenceSetIterator) All() iter.Seq2[*migrationcenterpb.PreferenceSet, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReportConfigIterator) All() iter.Seq2[*migrationcenterpb.ReportConfig, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReportIterator) All() iter.Seq2[*migrationcenterpb.Report, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SourceIterator) All() iter.Seq2[*migrationcenterpb.Source, error] {
	return iterator.RangeAdapter(it.Next)
}
