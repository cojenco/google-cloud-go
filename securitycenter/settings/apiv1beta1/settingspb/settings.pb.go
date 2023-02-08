// Copyright 2020 Google LLC
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
// source: google/cloud/securitycenter/settings/v1beta1/settings.proto

package settingspb

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

// Defines the onboarding states for SCC
//
// Potentially is just an indicator that a user has reviewed some subset of
// our configuration surface, even if it's still currently set to its
// API-default state.
type Settings_OnboardingState int32

const (
	// No onboarding state has been set. Should not be seen in practice, but
	// should be functionally equivalent to DISABLED.
	Settings_ONBOARDING_STATE_UNSPECIFIED Settings_OnboardingState = 0
	// SCC is fully on boarded
	Settings_ENABLED Settings_OnboardingState = 1
	// SCC has been disabled after being on boarded
	Settings_DISABLED Settings_OnboardingState = 2
	// SCC's onboarding tier has been explicitly set
	Settings_BILLING_SELECTED Settings_OnboardingState = 3
	// SCC's CTD FindingsProviders have been chosen
	Settings_PROVIDERS_SELECTED Settings_OnboardingState = 4
	// SCC's Service-Resource mappings have been set
	Settings_RESOURCES_SELECTED Settings_OnboardingState = 5
	// SCC's core Service Account was created
	Settings_ORG_SERVICE_ACCOUNT_CREATED Settings_OnboardingState = 6
)

// Enum value maps for Settings_OnboardingState.
var (
	Settings_OnboardingState_name = map[int32]string{
		0: "ONBOARDING_STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
		3: "BILLING_SELECTED",
		4: "PROVIDERS_SELECTED",
		5: "RESOURCES_SELECTED",
		6: "ORG_SERVICE_ACCOUNT_CREATED",
	}
	Settings_OnboardingState_value = map[string]int32{
		"ONBOARDING_STATE_UNSPECIFIED": 0,
		"ENABLED":                      1,
		"DISABLED":                     2,
		"BILLING_SELECTED":             3,
		"PROVIDERS_SELECTED":           4,
		"RESOURCES_SELECTED":           5,
		"ORG_SERVICE_ACCOUNT_CREATED":  6,
	}
)

func (x Settings_OnboardingState) Enum() *Settings_OnboardingState {
	p := new(Settings_OnboardingState)
	*p = x
	return p
}

func (x Settings_OnboardingState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Settings_OnboardingState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_settings_v1beta1_settings_proto_enumTypes[0].Descriptor()
}

func (Settings_OnboardingState) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_settings_v1beta1_settings_proto_enumTypes[0]
}

func (x Settings_OnboardingState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Settings_OnboardingState.Descriptor instead.
func (Settings_OnboardingState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescGZIP(), []int{0, 0}
}

// Common configuration settings for all of Security Center.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative resource name of the settings resource.
	// Formats:
	//  * `organizations/{organization}/settings`
	//  * `folders/{folder}/settings`
	//  * `projects/{project}/settings`
	//  * `projects/{project}/locations/{location}/clusters/{cluster}/settings`
	//  * `projects/{project}/regions/{region}/clusters/{cluster}/settings`
	//  * `projects/{project}/zones/{zone}/clusters/{cluster}/settings`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Billing settings
	BillingSettings *BillingSettings `protobuf:"bytes,2,opt,name=billing_settings,json=billingSettings,proto3" json:"billing_settings,omitempty"`
	// An enum representing the current on boarding state of SCC.
	State Settings_OnboardingState `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.securitycenter.settings.v1beta1.Settings_OnboardingState" json:"state,omitempty"`
	// Output only. The organization-level service account to be used for security center
	// components. The component must have permission to "act as" the service
	// account.
	OrgServiceAccount string `protobuf:"bytes,5,opt,name=org_service_account,json=orgServiceAccount,proto3" json:"org_service_account,omitempty"`
	// Sink settings.
	SinkSettings *SinkSettings `protobuf:"bytes,6,opt,name=sink_settings,json=sinkSettings,proto3" json:"sink_settings,omitempty"`
	// The settings for detectors and/or scanners.
	ComponentSettings map[string]*ComponentSettings `protobuf:"bytes,7,rep,name=component_settings,json=componentSettings,proto3" json:"component_settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Detector group settings for all Security Center components.
	// The key is the name of the detector group and the value is the settings for
	// that group.
	DetectorGroupSettings map[string]*Settings_DetectorGroupSettings `protobuf:"bytes,8,rep,name=detector_group_settings,json=detectorGroupSettings,proto3" json:"detector_group_settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A fingerprint used for optimistic concurrency. If none is provided
	// on updates then the existing metadata will be blindly overwritten.
	Etag string `protobuf:"bytes,9,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. The time these settings were last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Settings) GetBillingSettings() *BillingSettings {
	if x != nil {
		return x.BillingSettings
	}
	return nil
}

func (x *Settings) GetState() Settings_OnboardingState {
	if x != nil {
		return x.State
	}
	return Settings_ONBOARDING_STATE_UNSPECIFIED
}

func (x *Settings) GetOrgServiceAccount() string {
	if x != nil {
		return x.OrgServiceAccount
	}
	return ""
}

func (x *Settings) GetSinkSettings() *SinkSettings {
	if x != nil {
		return x.SinkSettings
	}
	return nil
}

func (x *Settings) GetComponentSettings() map[string]*ComponentSettings {
	if x != nil {
		return x.ComponentSettings
	}
	return nil
}

func (x *Settings) GetDetectorGroupSettings() map[string]*Settings_DetectorGroupSettings {
	if x != nil {
		return x.DetectorGroupSettings
	}
	return nil
}

func (x *Settings) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Settings) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// The DetectorGroupSettings define the configuration for a detector group.
type Settings_DetectorGroupSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state determines if the group is enabled or not.
	State ComponentEnablementState `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState" json:"state,omitempty"`
}

func (x *Settings_DetectorGroupSettings) Reset() {
	*x = Settings_DetectorGroupSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_DetectorGroupSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_DetectorGroupSettings) ProtoMessage() {}

func (x *Settings_DetectorGroupSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_DetectorGroupSettings.ProtoReflect.Descriptor instead.
func (*Settings_DetectorGroupSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Settings_DetectorGroupSettings) GetState() ComponentEnablementState {
	if x != nil {
		return x.State
	}
	return ComponentEnablementState_COMPONENT_ENABLEMENT_STATE_UNSPECIFIED
}

var File_google_cloud_securitycenter_settings_v1beta1_settings_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x5c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x33, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x11, 0x6f, 0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x7c, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x17, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x75, 0x0a, 0x15, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5c,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x85, 0x01, 0x0a,
	0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x96, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x62, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x01,
	0x0a, 0x0f, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52,
	0x53, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x06, 0x3a, 0xce, 0x02, 0xea, 0x41, 0xca, 0x02, 0x0a, 0x26, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x19, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x2f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0xa5, 0x02, 0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x62, 0x3b,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x2c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x2c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x30, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescData = file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDesc
)

func file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDescData
}

var file_google_cloud_securitycenter_settings_v1beta1_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_securitycenter_settings_v1beta1_settings_proto_goTypes = []interface{}{
	(Settings_OnboardingState)(0),          // 0: google.cloud.securitycenter.settings.v1beta1.Settings.OnboardingState
	(*Settings)(nil),                       // 1: google.cloud.securitycenter.settings.v1beta1.Settings
	(*Settings_DetectorGroupSettings)(nil), // 2: google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettings
	nil,                                    // 3: google.cloud.securitycenter.settings.v1beta1.Settings.ComponentSettingsEntry
	nil,                                    // 4: google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettingsEntry
	(*BillingSettings)(nil),                // 5: google.cloud.securitycenter.settings.v1beta1.BillingSettings
	(*SinkSettings)(nil),                   // 6: google.cloud.securitycenter.settings.v1beta1.SinkSettings
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
	(ComponentEnablementState)(0),          // 8: google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState
	(*ComponentSettings)(nil),              // 9: google.cloud.securitycenter.settings.v1beta1.ComponentSettings
}
var file_google_cloud_securitycenter_settings_v1beta1_settings_proto_depIdxs = []int32{
	5, // 0: google.cloud.securitycenter.settings.v1beta1.Settings.billing_settings:type_name -> google.cloud.securitycenter.settings.v1beta1.BillingSettings
	0, // 1: google.cloud.securitycenter.settings.v1beta1.Settings.state:type_name -> google.cloud.securitycenter.settings.v1beta1.Settings.OnboardingState
	6, // 2: google.cloud.securitycenter.settings.v1beta1.Settings.sink_settings:type_name -> google.cloud.securitycenter.settings.v1beta1.SinkSettings
	3, // 3: google.cloud.securitycenter.settings.v1beta1.Settings.component_settings:type_name -> google.cloud.securitycenter.settings.v1beta1.Settings.ComponentSettingsEntry
	4, // 4: google.cloud.securitycenter.settings.v1beta1.Settings.detector_group_settings:type_name -> google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettingsEntry
	7, // 5: google.cloud.securitycenter.settings.v1beta1.Settings.update_time:type_name -> google.protobuf.Timestamp
	8, // 6: google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettings.state:type_name -> google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState
	9, // 7: google.cloud.securitycenter.settings.v1beta1.Settings.ComponentSettingsEntry.value:type_name -> google.cloud.securitycenter.settings.v1beta1.ComponentSettings
	2, // 8: google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettingsEntry.value:type_name -> google.cloud.securitycenter.settings.v1beta1.Settings.DetectorGroupSettings
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_settings_v1beta1_settings_proto_init() }
func file_google_cloud_securitycenter_settings_v1beta1_settings_proto_init() {
	if File_google_cloud_securitycenter_settings_v1beta1_settings_proto != nil {
		return
	}
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_init()
	file_google_cloud_securitycenter_settings_v1beta1_component_settings_proto_init()
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_DetectorGroupSettings); i {
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
			RawDescriptor: file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_settings_v1beta1_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_settings_v1beta1_settings_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_settings_v1beta1_settings_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_settings_v1beta1_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_settings_v1beta1_settings_proto = out.File
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_rawDesc = nil
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_goTypes = nil
	file_google_cloud_securitycenter_settings_v1beta1_settings_proto_depIdxs = nil
}
