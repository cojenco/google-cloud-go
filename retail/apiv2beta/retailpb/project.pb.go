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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/retail/v2beta/project.proto

package retailpb

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

// The enrollment status enum for alert policy.
type AlertConfig_AlertPolicy_EnrollStatus int32

const (
	// Default value. Used for customers who have not responded to the
	// alert policy.
	AlertConfig_AlertPolicy_ENROLL_STATUS_UNSPECIFIED AlertConfig_AlertPolicy_EnrollStatus = 0
	// Customer is enrolled in this policy.
	AlertConfig_AlertPolicy_ENROLLED AlertConfig_AlertPolicy_EnrollStatus = 1
	// Customer declined this policy.
	AlertConfig_AlertPolicy_DECLINED AlertConfig_AlertPolicy_EnrollStatus = 2
)

// Enum value maps for AlertConfig_AlertPolicy_EnrollStatus.
var (
	AlertConfig_AlertPolicy_EnrollStatus_name = map[int32]string{
		0: "ENROLL_STATUS_UNSPECIFIED",
		1: "ENROLLED",
		2: "DECLINED",
	}
	AlertConfig_AlertPolicy_EnrollStatus_value = map[string]int32{
		"ENROLL_STATUS_UNSPECIFIED": 0,
		"ENROLLED":                  1,
		"DECLINED":                  2,
	}
)

func (x AlertConfig_AlertPolicy_EnrollStatus) Enum() *AlertConfig_AlertPolicy_EnrollStatus {
	p := new(AlertConfig_AlertPolicy_EnrollStatus)
	*p = x
	return p
}

func (x AlertConfig_AlertPolicy_EnrollStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertConfig_AlertPolicy_EnrollStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2beta_project_proto_enumTypes[0].Descriptor()
}

func (AlertConfig_AlertPolicy_EnrollStatus) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2beta_project_proto_enumTypes[0]
}

func (x AlertConfig_AlertPolicy_EnrollStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertConfig_AlertPolicy_EnrollStatus.Descriptor instead.
func (AlertConfig_AlertPolicy_EnrollStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_project_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Project level alert config.
type AlertConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The name of the AlertConfig singleton resource.
	// Format: projects/*/alertConfig
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Alert policies for a customer.
	// They must be unique by [AlertPolicy.alert_group]
	AlertPolicies []*AlertConfig_AlertPolicy `protobuf:"bytes,2,rep,name=alert_policies,json=alertPolicies,proto3" json:"alert_policies,omitempty"`
}

func (x *AlertConfig) Reset() {
	*x = AlertConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertConfig) ProtoMessage() {}

func (x *AlertConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertConfig.ProtoReflect.Descriptor instead.
func (*AlertConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_project_proto_rawDescGZIP(), []int{0}
}

func (x *AlertConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertConfig) GetAlertPolicies() []*AlertConfig_AlertPolicy {
	if x != nil {
		return x.AlertPolicies
	}
	return nil
}

// Alert policy for a customer.
type AlertConfig_AlertPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The feature that provides alerting capability.
	// Supported value:
	// - `search-data-quality` for retail search customers.
	// - `conv-data-quality` for retail conversation customers.
	AlertGroup string `protobuf:"bytes,1,opt,name=alert_group,json=alertGroup,proto3" json:"alert_group,omitempty"`
	// The enrollment status of a customer.
	EnrollStatus AlertConfig_AlertPolicy_EnrollStatus `protobuf:"varint,2,opt,name=enroll_status,json=enrollStatus,proto3,enum=google.cloud.retail.v2beta.AlertConfig_AlertPolicy_EnrollStatus" json:"enroll_status,omitempty"`
	// Recipients for the alert policy.
	// One alert policy should not exceed 20 recipients.
	Recipients []*AlertConfig_AlertPolicy_Recipient `protobuf:"bytes,3,rep,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *AlertConfig_AlertPolicy) Reset() {
	*x = AlertConfig_AlertPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertConfig_AlertPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertConfig_AlertPolicy) ProtoMessage() {}

func (x *AlertConfig_AlertPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertConfig_AlertPolicy.ProtoReflect.Descriptor instead.
func (*AlertConfig_AlertPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_project_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AlertConfig_AlertPolicy) GetAlertGroup() string {
	if x != nil {
		return x.AlertGroup
	}
	return ""
}

func (x *AlertConfig_AlertPolicy) GetEnrollStatus() AlertConfig_AlertPolicy_EnrollStatus {
	if x != nil {
		return x.EnrollStatus
	}
	return AlertConfig_AlertPolicy_ENROLL_STATUS_UNSPECIFIED
}

func (x *AlertConfig_AlertPolicy) GetRecipients() []*AlertConfig_AlertPolicy_Recipient {
	if x != nil {
		return x.Recipients
	}
	return nil
}

// Recipient contact information.
type AlertConfig_AlertPolicy_Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Email address of the recipient.
	EmailAddress string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
}

func (x *AlertConfig_AlertPolicy_Recipient) Reset() {
	*x = AlertConfig_AlertPolicy_Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertConfig_AlertPolicy_Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertConfig_AlertPolicy_Recipient) ProtoMessage() {}

func (x *AlertConfig_AlertPolicy_Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertConfig_AlertPolicy_Recipient.ProtoReflect.Descriptor instead.
func (*AlertConfig_AlertPolicy_Recipient) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_project_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AlertConfig_AlertPolicy_Recipient) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

var File_google_cloud_retail_v2beta_project_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2beta_project_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc1, 0x04, 0x0a, 0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5a,
	0x0a, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x1a, 0xf1, 0x02, 0x0a, 0x0b, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x65, 0x0a, 0x0d, 0x65,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x5d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x0c, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x46,
	0xea, 0x41, 0x43, 0x0a, 0x21, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xcb, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70,
	0x62, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x56, 0x32, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32,
	0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2beta_project_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2beta_project_proto_rawDescData = file_google_cloud_retail_v2beta_project_proto_rawDesc
)

func file_google_cloud_retail_v2beta_project_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2beta_project_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2beta_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2beta_project_proto_rawDescData)
	})
	return file_google_cloud_retail_v2beta_project_proto_rawDescData
}

var file_google_cloud_retail_v2beta_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_retail_v2beta_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_retail_v2beta_project_proto_goTypes = []any{
	(AlertConfig_AlertPolicy_EnrollStatus)(0), // 0: google.cloud.retail.v2beta.AlertConfig.AlertPolicy.EnrollStatus
	(*AlertConfig)(nil),                       // 1: google.cloud.retail.v2beta.AlertConfig
	(*AlertConfig_AlertPolicy)(nil),           // 2: google.cloud.retail.v2beta.AlertConfig.AlertPolicy
	(*AlertConfig_AlertPolicy_Recipient)(nil), // 3: google.cloud.retail.v2beta.AlertConfig.AlertPolicy.Recipient
}
var file_google_cloud_retail_v2beta_project_proto_depIdxs = []int32{
	2, // 0: google.cloud.retail.v2beta.AlertConfig.alert_policies:type_name -> google.cloud.retail.v2beta.AlertConfig.AlertPolicy
	0, // 1: google.cloud.retail.v2beta.AlertConfig.AlertPolicy.enroll_status:type_name -> google.cloud.retail.v2beta.AlertConfig.AlertPolicy.EnrollStatus
	3, // 2: google.cloud.retail.v2beta.AlertConfig.AlertPolicy.recipients:type_name -> google.cloud.retail.v2beta.AlertConfig.AlertPolicy.Recipient
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2beta_project_proto_init() }
func file_google_cloud_retail_v2beta_project_proto_init() {
	if File_google_cloud_retail_v2beta_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2beta_project_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AlertConfig); i {
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
		file_google_cloud_retail_v2beta_project_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AlertConfig_AlertPolicy); i {
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
		file_google_cloud_retail_v2beta_project_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AlertConfig_AlertPolicy_Recipient); i {
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
			RawDescriptor: file_google_cloud_retail_v2beta_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2beta_project_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2beta_project_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2beta_project_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2beta_project_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2beta_project_proto = out.File
	file_google_cloud_retail_v2beta_project_proto_rawDesc = nil
	file_google_cloud_retail_v2beta_project_proto_goTypes = nil
	file_google_cloud_retail_v2beta_project_proto_depIdxs = nil
}
