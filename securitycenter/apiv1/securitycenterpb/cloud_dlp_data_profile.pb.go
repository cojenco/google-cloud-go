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
// source: google/cloud/securitycenter/v1/cloud_dlp_data_profile.proto

package securitycenterpb

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

// Parents for configurations that produce data profile findings.
type CloudDlpDataProfile_ParentType int32

const (
	// Unspecified parent type.
	CloudDlpDataProfile_PARENT_TYPE_UNSPECIFIED CloudDlpDataProfile_ParentType = 0
	// Organization-level configurations.
	CloudDlpDataProfile_ORGANIZATION CloudDlpDataProfile_ParentType = 1
	// Project-level configurations.
	CloudDlpDataProfile_PROJECT CloudDlpDataProfile_ParentType = 2
)

// Enum value maps for CloudDlpDataProfile_ParentType.
var (
	CloudDlpDataProfile_ParentType_name = map[int32]string{
		0: "PARENT_TYPE_UNSPECIFIED",
		1: "ORGANIZATION",
		2: "PROJECT",
	}
	CloudDlpDataProfile_ParentType_value = map[string]int32{
		"PARENT_TYPE_UNSPECIFIED": 0,
		"ORGANIZATION":            1,
		"PROJECT":                 2,
	}
)

func (x CloudDlpDataProfile_ParentType) Enum() *CloudDlpDataProfile_ParentType {
	p := new(CloudDlpDataProfile_ParentType)
	*p = x
	return p
}

func (x CloudDlpDataProfile_ParentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudDlpDataProfile_ParentType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_enumTypes[0].Descriptor()
}

func (CloudDlpDataProfile_ParentType) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_enumTypes[0]
}

func (x CloudDlpDataProfile_ParentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudDlpDataProfile_ParentType.Descriptor instead.
func (CloudDlpDataProfile_ParentType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescGZIP(), []int{0, 0}
}

// The [data profile](https://cloud.google.com/dlp/docs/data-profiles)
// associated with the finding.
type CloudDlpDataProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the data profile, for example,
	// `projects/123/locations/europe/tableProfiles/8383929`.
	DataProfile string `protobuf:"bytes,1,opt,name=data_profile,json=dataProfile,proto3" json:"data_profile,omitempty"`
	// The resource hierarchy level at which the data profile was generated.
	ParentType CloudDlpDataProfile_ParentType `protobuf:"varint,2,opt,name=parent_type,json=parentType,proto3,enum=google.cloud.securitycenter.v1.CloudDlpDataProfile_ParentType" json:"parent_type,omitempty"`
}

func (x *CloudDlpDataProfile) Reset() {
	*x = CloudDlpDataProfile{}
	mi := &file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudDlpDataProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudDlpDataProfile) ProtoMessage() {}

func (x *CloudDlpDataProfile) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudDlpDataProfile.ProtoReflect.Descriptor instead.
func (*CloudDlpDataProfile) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescGZIP(), []int{0}
}

func (x *CloudDlpDataProfile) GetDataProfile() string {
	if x != nil {
		return x.DataProfile
	}
	return ""
}

func (x *CloudDlpDataProfile) GetParentType() CloudDlpDataProfile_ParentType {
	if x != nil {
		return x.ParentType
	}
	return CloudDlpDataProfile_PARENT_TYPE_UNSPECIFIED
}

var File_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x64, 0x6c, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x13, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x44, 0x6c, 0x70, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x4b, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x64, 0x6c, 0x70,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x5f, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6c, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x48,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x42, 0x94, 0x03, 0xea, 0x41, 0x9e, 0x01, 0x0a,
	0x23, 0x64, 0x6c, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x7d, 0x12, 0x45, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x7d, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x18, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6c, 0x70, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescData = file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_goTypes = []any{
	(CloudDlpDataProfile_ParentType)(0), // 0: google.cloud.securitycenter.v1.CloudDlpDataProfile.ParentType
	(*CloudDlpDataProfile)(nil),         // 1: google.cloud.securitycenter.v1.CloudDlpDataProfile
}
var file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_depIdxs = []int32{
	0, // 0: google.cloud.securitycenter.v1.CloudDlpDataProfile.parent_type:type_name -> google.cloud.securitycenter.v1.CloudDlpDataProfile.ParentType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_init() }
func file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_init() {
	if File_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto = out.File
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_cloud_dlp_data_profile_proto_depIdxs = nil
}
