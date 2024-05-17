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
// source: google/cloud/cloudcontrolspartner/v1beta/partner_permissions.proto

package cloudcontrolspartnerpb

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

type PartnerPermissions_Permission int32

const (
	// Unspecified partner permission
	PartnerPermissions_PERMISSION_UNSPECIFIED PartnerPermissions_Permission = 0
	// Permission for Access Transparency and emergency logs
	PartnerPermissions_ACCESS_TRANSPARENCY_AND_EMERGENCY_ACCESS_LOGS PartnerPermissions_Permission = 1
	// Permission for Assured Workloads monitoring violations
	PartnerPermissions_ASSURED_WORKLOADS_MONITORING PartnerPermissions_Permission = 2
	// Permission for Access Approval requests
	PartnerPermissions_ACCESS_APPROVAL_REQUESTS PartnerPermissions_Permission = 3
	// Permission for External Key Manager connection status
	PartnerPermissions_ASSURED_WORKLOADS_EKM_CONNECTION_STATUS PartnerPermissions_Permission = 4
)

// Enum value maps for PartnerPermissions_Permission.
var (
	PartnerPermissions_Permission_name = map[int32]string{
		0: "PERMISSION_UNSPECIFIED",
		1: "ACCESS_TRANSPARENCY_AND_EMERGENCY_ACCESS_LOGS",
		2: "ASSURED_WORKLOADS_MONITORING",
		3: "ACCESS_APPROVAL_REQUESTS",
		4: "ASSURED_WORKLOADS_EKM_CONNECTION_STATUS",
	}
	PartnerPermissions_Permission_value = map[string]int32{
		"PERMISSION_UNSPECIFIED":                        0,
		"ACCESS_TRANSPARENCY_AND_EMERGENCY_ACCESS_LOGS": 1,
		"ASSURED_WORKLOADS_MONITORING":                  2,
		"ACCESS_APPROVAL_REQUESTS":                      3,
		"ASSURED_WORKLOADS_EKM_CONNECTION_STATUS":       4,
	}
)

func (x PartnerPermissions_Permission) Enum() *PartnerPermissions_Permission {
	p := new(PartnerPermissions_Permission)
	*p = x
	return p
}

func (x PartnerPermissions_Permission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartnerPermissions_Permission) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_enumTypes[0].Descriptor()
}

func (PartnerPermissions_Permission) Type() protoreflect.EnumType {
	return &file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_enumTypes[0]
}

func (x PartnerPermissions_Permission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartnerPermissions_Permission.Descriptor instead.
func (PartnerPermissions_Permission) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescGZIP(), []int{0, 0}
}

// The permissions granted to the partner for a workload
type PartnerPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Format:
	// organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/partnerPermissions
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The partner permissions granted for the workload
	PartnerPermissions []PartnerPermissions_Permission `protobuf:"varint,2,rep,packed,name=partner_permissions,json=partnerPermissions,proto3,enum=google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions_Permission" json:"partner_permissions,omitempty"`
}

func (x *PartnerPermissions) Reset() {
	*x = PartnerPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartnerPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerPermissions) ProtoMessage() {}

func (x *PartnerPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerPermissions.ProtoReflect.Descriptor instead.
func (*PartnerPermissions) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *PartnerPermissions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartnerPermissions) GetPartnerPermissions() []PartnerPermissions_Permission {
	if x != nil {
		return x.PartnerPermissions
	}
	return nil
}

// Request for getting the partner permissions granted for a workload
type GetPartnerPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the resource to get in the format:
	// organizations/{organization}/locations/{location}/customers/{customer}/workloads/{workload}/partnerPermissions
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPartnerPermissionsRequest) Reset() {
	*x = GetPartnerPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerPermissionsRequest) ProtoMessage() {}

func (x *GetPartnerPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartnerPermissionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x12, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x78, 0x0a, 0x13, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x31, 0x0a, 0x2d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50,
	0x41, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x45, 0x4d, 0x45, 0x52, 0x47,
	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x53,
	0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x53, 0x53, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x4c, 0x4f, 0x41, 0x44, 0x53, 0x5f, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53,
	0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x41, 0x53, 0x53, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x4c, 0x4f, 0x41, 0x44, 0x53, 0x5f, 0x45, 0x4b, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x04, 0x3a,
	0xac, 0x01, 0xea, 0x41, 0xa8, 0x01, 0x0a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x72,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x38, 0x0a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0xaf, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x42, 0x17, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62,
	0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescOnce sync.Once
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescData = file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDesc
)

func file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescGZIP() []byte {
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescOnce.Do(func() {
		file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescData)
	})
	return file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDescData
}

var file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_goTypes = []interface{}{
	(PartnerPermissions_Permission)(0),   // 0: google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions.Permission
	(*PartnerPermissions)(nil),           // 1: google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions
	(*GetPartnerPermissionsRequest)(nil), // 2: google.cloud.cloudcontrolspartner.v1beta.GetPartnerPermissionsRequest
}
var file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_depIdxs = []int32{
	0, // 0: google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions.partner_permissions:type_name -> google.cloud.cloudcontrolspartner.v1beta.PartnerPermissions.Permission
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartnerPermissions); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerPermissionsRequest); i {
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
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_depIdxs,
		EnumInfos:         file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_enumTypes,
		MessageInfos:      file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_msgTypes,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1beta_partner_permissions_proto_depIdxs = nil
}
