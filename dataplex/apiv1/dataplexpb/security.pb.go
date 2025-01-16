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
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/cloud/dataplex/v1/security.proto

package dataplexpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ResourceAccessSpec holds the access control configuration to be enforced
// on the resources, for example, Cloud Storage bucket, BigQuery dataset,
// BigQuery table.
type ResourceAccessSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The format of strings follows the pattern followed by IAM in the
	// bindings. user:{email}, serviceAccount:{email} group:{email}.
	// The set of principals to be granted reader role on the resource.
	Readers []string `protobuf:"bytes,1,rep,name=readers,proto3" json:"readers,omitempty"`
	// Optional. The set of principals to be granted writer role on the resource.
	Writers []string `protobuf:"bytes,2,rep,name=writers,proto3" json:"writers,omitempty"`
	// Optional. The set of principals to be granted owner role on the resource.
	Owners []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
}

func (x *ResourceAccessSpec) Reset() {
	*x = ResourceAccessSpec{}
	mi := &file_google_cloud_dataplex_v1_security_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceAccessSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAccessSpec) ProtoMessage() {}

func (x *ResourceAccessSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_security_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAccessSpec.ProtoReflect.Descriptor instead.
func (*ResourceAccessSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_security_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceAccessSpec) GetReaders() []string {
	if x != nil {
		return x.Readers
	}
	return nil
}

func (x *ResourceAccessSpec) GetWriters() []string {
	if x != nil {
		return x.Writers
	}
	return nil
}

func (x *ResourceAccessSpec) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

// DataAccessSpec holds the access control configuration to be enforced on data
// stored within resources (eg: rows, columns in BigQuery Tables). When
// associated with data, the data is only accessible to
// principals explicitly granted access through the DataAccessSpec. Principals
// with access to the containing resource are not implicitly granted access.
type DataAccessSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The format of strings follows the pattern followed by IAM in the
	// bindings. user:{email}, serviceAccount:{email} group:{email}.
	// The set of principals to be granted reader role on data
	// stored within resources.
	Readers []string `protobuf:"bytes,1,rep,name=readers,proto3" json:"readers,omitempty"`
}

func (x *DataAccessSpec) Reset() {
	*x = DataAccessSpec{}
	mi := &file_google_cloud_dataplex_v1_security_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataAccessSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAccessSpec) ProtoMessage() {}

func (x *DataAccessSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataplex_v1_security_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAccessSpec.ProtoReflect.Descriptor instead.
func (*DataAccessSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataplex_v1_security_proto_rawDescGZIP(), []int{1}
}

func (x *DataAccessSpec) GetReaders() []string {
	if x != nil {
		return x.Readers
	}
	return nil
}

var File_google_cloud_dataplex_v1_security_proto protoreflect.FileDescriptor

var file_google_cloud_dataplex_v1_security_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x42, 0x69, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x65, 0x78, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataplex_v1_security_proto_rawDescOnce sync.Once
	file_google_cloud_dataplex_v1_security_proto_rawDescData = file_google_cloud_dataplex_v1_security_proto_rawDesc
)

func file_google_cloud_dataplex_v1_security_proto_rawDescGZIP() []byte {
	file_google_cloud_dataplex_v1_security_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataplex_v1_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataplex_v1_security_proto_rawDescData)
	})
	return file_google_cloud_dataplex_v1_security_proto_rawDescData
}

var file_google_cloud_dataplex_v1_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dataplex_v1_security_proto_goTypes = []any{
	(*ResourceAccessSpec)(nil), // 0: google.cloud.dataplex.v1.ResourceAccessSpec
	(*DataAccessSpec)(nil),     // 1: google.cloud.dataplex.v1.DataAccessSpec
}
var file_google_cloud_dataplex_v1_security_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_dataplex_v1_security_proto_init() }
func file_google_cloud_dataplex_v1_security_proto_init() {
	if File_google_cloud_dataplex_v1_security_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dataplex_v1_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataplex_v1_security_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataplex_v1_security_proto_depIdxs,
		MessageInfos:      file_google_cloud_dataplex_v1_security_proto_msgTypes,
	}.Build()
	File_google_cloud_dataplex_v1_security_proto = out.File
	file_google_cloud_dataplex_v1_security_proto_rawDesc = nil
	file_google_cloud_dataplex_v1_security_proto_goTypes = nil
	file_google_cloud_dataplex_v1_security_proto_depIdxs = nil
}
