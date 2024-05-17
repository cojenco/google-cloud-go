// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/firestore/v1/aggregation_result.proto

package firestorepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The result of a single bucket from a Firestore aggregation query.
//
// The keys of `aggregate_fields` are the same for all results in an aggregation
// query, unlike document queries which can have different fields present for
// each result.
type AggregationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result of the aggregation functions, ex: `COUNT(*) AS total_docs`.
	//
	// The key is the
	// [alias][google.firestore.v1.StructuredAggregationQuery.Aggregation.alias]
	// assigned to the aggregation function on input and the size of this map
	// equals the number of aggregation functions in the query.
	AggregateFields map[string]*Value `protobuf:"bytes,2,rep,name=aggregate_fields,json=aggregateFields,proto3" json:"aggregate_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AggregationResult) Reset() {
	*x = AggregationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1_aggregation_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationResult) ProtoMessage() {}

func (x *AggregationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1_aggregation_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationResult.ProtoReflect.Descriptor instead.
func (*AggregationResult) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1_aggregation_result_proto_rawDescGZIP(), []int{0}
}

func (x *AggregationResult) GetAggregateFields() map[string]*Value {
	if x != nil {
		return x.AggregateFields
	}
	return nil
}

var File_google_firestore_v1_aggregation_result_proto protoreflect.FileDescriptor

var file_google_firestore_v1_aggregation_result_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x66, 0x0a,
	0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x5e, 0x0a, 0x14, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xce, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x3b, 0x66, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa,
	0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_v1_aggregation_result_proto_rawDescOnce sync.Once
	file_google_firestore_v1_aggregation_result_proto_rawDescData = file_google_firestore_v1_aggregation_result_proto_rawDesc
)

func file_google_firestore_v1_aggregation_result_proto_rawDescGZIP() []byte {
	file_google_firestore_v1_aggregation_result_proto_rawDescOnce.Do(func() {
		file_google_firestore_v1_aggregation_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_v1_aggregation_result_proto_rawDescData)
	})
	return file_google_firestore_v1_aggregation_result_proto_rawDescData
}

var file_google_firestore_v1_aggregation_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_firestore_v1_aggregation_result_proto_goTypes = []interface{}{
	(*AggregationResult)(nil), // 0: google.firestore.v1.AggregationResult
	nil,                       // 1: google.firestore.v1.AggregationResult.AggregateFieldsEntry
	(*Value)(nil),             // 2: google.firestore.v1.Value
}
var file_google_firestore_v1_aggregation_result_proto_depIdxs = []int32{
	1, // 0: google.firestore.v1.AggregationResult.aggregate_fields:type_name -> google.firestore.v1.AggregationResult.AggregateFieldsEntry
	2, // 1: google.firestore.v1.AggregationResult.AggregateFieldsEntry.value:type_name -> google.firestore.v1.Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_firestore_v1_aggregation_result_proto_init() }
func file_google_firestore_v1_aggregation_result_proto_init() {
	if File_google_firestore_v1_aggregation_result_proto != nil {
		return
	}
	file_google_firestore_v1_document_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_v1_aggregation_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_v1_aggregation_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_v1_aggregation_result_proto_goTypes,
		DependencyIndexes: file_google_firestore_v1_aggregation_result_proto_depIdxs,
		MessageInfos:      file_google_firestore_v1_aggregation_result_proto_msgTypes,
	}.Build()
	File_google_firestore_v1_aggregation_result_proto = out.File
	file_google_firestore_v1_aggregation_result_proto_rawDesc = nil
	file_google_firestore_v1_aggregation_result_proto_goTypes = nil
	file_google_firestore_v1_aggregation_result_proto_depIdxs = nil
}
