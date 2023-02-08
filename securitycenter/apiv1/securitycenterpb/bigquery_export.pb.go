// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: google/cloud/securitycenter/v1/bigquery_export.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configures how to deliver Findings to BigQuery Instance.
type BigQueryExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative resource name of this export. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name.
	// Example format:
	// "organizations/{organization_id}/bigQueryExports/{export_id}" Example
	// format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format:
	// "projects/{project_id}/bigQueryExports/{export_id}"
	// This field is provided in responses, and is ignored when provided in create
	// requests.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the export (max of 1024 characters).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Expression that defines the filter to apply across create/update events
	// of findings. The expression is a list of zero or more restrictions combined
	// via logical operators `AND` and `OR`. Parentheses are supported, and `OR`
	// has higher precedence than `AND`.
	//
	// Restrictions have the form `<field> <operator> <value>` and may have a
	// `-` character in front of them to indicate negation. The fields map to
	// those defined in the corresponding resource.
	//
	// The supported operators are:
	//
	// * `=` for all value types.
	// * `>`, `<`, `>=`, `<=` for integer values.
	// * `:`, meaning substring matching, for strings.
	//
	// The supported value types are:
	//
	// * string literals in quotes.
	// * integer literals without quotes.
	// * boolean literals `true` and `false` without quotes.
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// The dataset to write findings' updates to. Its format is
	// "projects/[project_id]/datasets/[bigquery_dataset_id]".
	// BigQuery Dataset unique ID  must contain only letters (a-z, A-Z), numbers
	// (0-9), or underscores (_).
	Dataset string `protobuf:"bytes,4,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Output only. The time at which the big query export was created.
	// This field is set by the server and will be ignored if provided on export
	// on creation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The most recent time at which the big export was updated.
	// This field is set by the server and will be ignored if provided on export
	// creation or update.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Email address of the user who last edited the big query export.
	// This field is set by the server and will be ignored if provided on export
	// creation or update.
	MostRecentEditor string `protobuf:"bytes,7,opt,name=most_recent_editor,json=mostRecentEditor,proto3" json:"most_recent_editor,omitempty"`
	// Output only. The service account that needs permission to create table, upload data to
	// the big query dataset.
	Principal string `protobuf:"bytes,8,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *BigQueryExport) Reset() {
	*x = BigQueryExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_bigquery_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryExport) ProtoMessage() {}

func (x *BigQueryExport) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_bigquery_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryExport.ProtoReflect.Descriptor instead.
func (*BigQueryExport) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescGZIP(), []int{0}
}

func (x *BigQueryExport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BigQueryExport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BigQueryExport) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *BigQueryExport) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *BigQueryExport) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BigQueryExport) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BigQueryExport) GetMostRecentEditor() string {
	if x != nil {
		return x.MostRecentEditor
	}
	return ""
}

func (x *BigQueryExport) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

var File_google_cloud_securitycenter_v1_bigquery_export_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x96, 0x04, 0x0a, 0x0e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x31, 0x0a, 0x12, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x10, 0x6d, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x45,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x3a, 0xc1, 0x01, 0xea, 0x41, 0xbd, 0x01, 0x0a,
	0x2c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x62, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x7d, 0x12, 0x29, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x62, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0x12,
	0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x62, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0x42, 0xed, 0x01, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescData = file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_bigquery_export_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_bigquery_export_proto_goTypes = []interface{}{
	(*BigQueryExport)(nil),        // 0: google.cloud.securitycenter.v1.BigQueryExport
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_v1_bigquery_export_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.BigQueryExport.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: google.cloud.securitycenter.v1.BigQueryExport.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_bigquery_export_proto_init() }
func file_google_cloud_securitycenter_v1_bigquery_export_proto_init() {
	if File_google_cloud_securitycenter_v1_bigquery_export_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_bigquery_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryExport); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_bigquery_export_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_bigquery_export_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_bigquery_export_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_bigquery_export_proto = out.File
	file_google_cloud_securitycenter_v1_bigquery_export_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_bigquery_export_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_bigquery_export_proto_depIdxs = nil
}
