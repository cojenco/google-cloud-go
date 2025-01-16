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
// source: google/cloud/run/v2/status.proto

package runpb

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

// Effective settings for the current revision
type RevisionScalingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current number of min instances provisioned for this revision.
	DesiredMinInstanceCount int32 `protobuf:"varint,1,opt,name=desired_min_instance_count,json=desiredMinInstanceCount,proto3" json:"desired_min_instance_count,omitempty"`
}

func (x *RevisionScalingStatus) Reset() {
	*x = RevisionScalingStatus{}
	mi := &file_google_cloud_run_v2_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevisionScalingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionScalingStatus) ProtoMessage() {}

func (x *RevisionScalingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v2_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionScalingStatus.ProtoReflect.Descriptor instead.
func (*RevisionScalingStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v2_status_proto_rawDescGZIP(), []int{0}
}

func (x *RevisionScalingStatus) GetDesiredMinInstanceCount() int32 {
	if x != nil {
		return x.DesiredMinInstanceCount
	}
	return 0
}

var File_google_cloud_run_v2_status_proto protoreflect.FileDescriptor

var file_google_cloud_run_v2_status_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x22, 0x54, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x4d, 0x69, 0x6e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x53, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x75, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x75, 0x6e, 0x70, 0x62, 0x3b, 0x72, 0x75, 0x6e,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v2_status_proto_rawDescOnce sync.Once
	file_google_cloud_run_v2_status_proto_rawDescData = file_google_cloud_run_v2_status_proto_rawDesc
)

func file_google_cloud_run_v2_status_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v2_status_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v2_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v2_status_proto_rawDescData)
	})
	return file_google_cloud_run_v2_status_proto_rawDescData
}

var file_google_cloud_run_v2_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_run_v2_status_proto_goTypes = []any{
	(*RevisionScalingStatus)(nil), // 0: google.cloud.run.v2.RevisionScalingStatus
}
var file_google_cloud_run_v2_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v2_status_proto_init() }
func file_google_cloud_run_v2_status_proto_init() {
	if File_google_cloud_run_v2_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v2_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v2_status_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v2_status_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v2_status_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v2_status_proto = out.File
	file_google_cloud_run_v2_status_proto_rawDesc = nil
	file_google_cloud_run_v2_status_proto_goTypes = nil
	file_google_cloud_run_v2_status_proto_depIdxs = nil
}
