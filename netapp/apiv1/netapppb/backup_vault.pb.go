// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/netapp/v1/backup_vault.proto

package netapppb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Backup Vault States
type BackupVault_State int32

const (
	// State not set.
	BackupVault_STATE_UNSPECIFIED BackupVault_State = 0
	// BackupVault is being created.
	BackupVault_CREATING BackupVault_State = 1
	// BackupVault is available for use.
	BackupVault_READY BackupVault_State = 2
	// BackupVault is being deleted.
	BackupVault_DELETING BackupVault_State = 3
	// BackupVault is not valid and cannot be used.
	BackupVault_ERROR BackupVault_State = 4
	// BackupVault is being updated.
	BackupVault_UPDATING BackupVault_State = 5
)

// Enum value maps for BackupVault_State.
var (
	BackupVault_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
		4: "ERROR",
		5: "UPDATING",
	}
	BackupVault_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
		"ERROR":             4,
		"UPDATING":          5,
	}
)

func (x BackupVault_State) Enum() *BackupVault_State {
	p := new(BackupVault_State)
	*p = x
	return p
}

func (x BackupVault_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupVault_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_netapp_v1_backup_vault_proto_enumTypes[0].Descriptor()
}

func (BackupVault_State) Type() protoreflect.EnumType {
	return &file_google_cloud_netapp_v1_backup_vault_proto_enumTypes[0]
}

func (x BackupVault_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupVault_State.Descriptor instead.
func (BackupVault_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{0, 0}
}

// A NetApp BackupVault.
type BackupVault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the backup vault.
	// Format:
	// `projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The backup vault state.
	State BackupVault_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.netapp.v1.BackupVault_State" json:"state,omitempty"`
	// Output only. Create time of the backup vault.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Description of the backup vault.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels to represent user provided metadata.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BackupVault) Reset() {
	*x = BackupVault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupVault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupVault) ProtoMessage() {}

func (x *BackupVault) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupVault.ProtoReflect.Descriptor instead.
func (*BackupVault) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{0}
}

func (x *BackupVault) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BackupVault) GetState() BackupVault_State {
	if x != nil {
		return x.State
	}
	return BackupVault_STATE_UNSPECIFIED
}

func (x *BackupVault) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BackupVault) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BackupVault) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// GetBackupVaultRequest gets the state of a backupVault.
type GetBackupVaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The backupVault resource name, in the format
	// `projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBackupVaultRequest) Reset() {
	*x = GetBackupVaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackupVaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackupVaultRequest) ProtoMessage() {}

func (x *GetBackupVaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackupVaultRequest.ProtoReflect.Descriptor instead.
func (*GetBackupVaultRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{1}
}

func (x *GetBackupVaultRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListBackupVaultsRequest lists backupVaults.
type ListBackupVaultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The location for which to retrieve backupVault information,
	// in the format
	// `projects/{project_id}/locations/{location}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value to use if there are additional
	// results to retrieve for this list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Sort results. Supported values are "name", "name desc" or "" (unsorted).
	OrderBy string `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// List filter.
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListBackupVaultsRequest) Reset() {
	*x = ListBackupVaultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupVaultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupVaultsRequest) ProtoMessage() {}

func (x *ListBackupVaultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupVaultsRequest.ProtoReflect.Descriptor instead.
func (*ListBackupVaultsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{2}
}

func (x *ListBackupVaultsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListBackupVaultsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBackupVaultsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListBackupVaultsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListBackupVaultsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ListBackupVaultsResponse is the result of ListBackupVaultsRequest.
type ListBackupVaultsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of backupVaults in the project for the specified location.
	BackupVaults []*BackupVault `protobuf:"bytes,1,rep,name=backup_vaults,json=backupVaults,proto3" json:"backup_vaults,omitempty"`
	// The token you can use to retrieve the next page of results. Not returned
	// if there are no more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListBackupVaultsResponse) Reset() {
	*x = ListBackupVaultsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBackupVaultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBackupVaultsResponse) ProtoMessage() {}

func (x *ListBackupVaultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBackupVaultsResponse.ProtoReflect.Descriptor instead.
func (*ListBackupVaultsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{3}
}

func (x *ListBackupVaultsResponse) GetBackupVaults() []*BackupVault {
	if x != nil {
		return x.BackupVaults
	}
	return nil
}

func (x *ListBackupVaultsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListBackupVaultsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// CreateBackupVaultRequest creates a backup vault.
type CreateBackupVaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The location to create the backup vaults, in the format
	// `projects/{project_id}/locations/{location}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The ID to use for the backupVault.
	// The ID must be unique within the specified location.
	// The max supported length is 63 characters.
	// This value must start with a lowercase letter followed by up to 62
	// lowercase letters, numbers, or hyphens, and cannot end with a hyphen.
	// Values that do not match this pattern will trigger an INVALID_ARGUMENT
	// error.
	BackupVaultId string `protobuf:"bytes,2,opt,name=backup_vault_id,json=backupVaultId,proto3" json:"backup_vault_id,omitempty"`
	// Required. A backupVault resource
	BackupVault *BackupVault `protobuf:"bytes,3,opt,name=backup_vault,json=backupVault,proto3" json:"backup_vault,omitempty"`
}

func (x *CreateBackupVaultRequest) Reset() {
	*x = CreateBackupVaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBackupVaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBackupVaultRequest) ProtoMessage() {}

func (x *CreateBackupVaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBackupVaultRequest.ProtoReflect.Descriptor instead.
func (*CreateBackupVaultRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBackupVaultRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBackupVaultRequest) GetBackupVaultId() string {
	if x != nil {
		return x.BackupVaultId
	}
	return ""
}

func (x *CreateBackupVaultRequest) GetBackupVault() *BackupVault {
	if x != nil {
		return x.BackupVault
	}
	return nil
}

// DeleteBackupVaultRequest deletes a backupVault.
type DeleteBackupVaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The backupVault resource name, in the format
	// `projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteBackupVaultRequest) Reset() {
	*x = DeleteBackupVaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBackupVaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBackupVaultRequest) ProtoMessage() {}

func (x *DeleteBackupVaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBackupVaultRequest.ProtoReflect.Descriptor instead.
func (*DeleteBackupVaultRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBackupVaultRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateBackupVaultRequest updates description and/or labels for a backupVault.
type UpdateBackupVaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Field mask is used to specify the fields to be overwritten in the
	// Backup resource to be updated.
	// The fields specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the mask. If the
	// user does not provide a mask then all fields will be overwritten.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. The backupVault being updated
	BackupVault *BackupVault `protobuf:"bytes,2,opt,name=backup_vault,json=backupVault,proto3" json:"backup_vault,omitempty"`
}

func (x *UpdateBackupVaultRequest) Reset() {
	*x = UpdateBackupVaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBackupVaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBackupVaultRequest) ProtoMessage() {}

func (x *UpdateBackupVaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBackupVaultRequest.ProtoReflect.Descriptor instead.
func (*UpdateBackupVaultRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBackupVaultRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateBackupVaultRequest) GetBackupVault() *BackupVault {
	if x != nil {
		return x.BackupVault
	}
	return nil
}

var File_google_cloud_netapp_v1_backup_vault_proto protoreflect.FileDescriptor

var file_google_cloud_netapp_v1_backup_vault_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbe, 0x04, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x3a, 0x87, 0x01, 0xea, 0x41, 0x83, 0x01,
	0x0a, 0x21, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x7d, 0x2a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x22, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x23, 0x0a, 0x21, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x12,
	0x21, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xae, 0x01, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x12,
	0x21, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x22, 0x59, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xa9, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x4b,
	0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x42, 0xb2, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70, 0x70, 0x62, 0x3b, 0x6e, 0x65, 0x74, 0x61, 0x70, 0x70,
	0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4e, 0x65, 0x74, 0x41, 0x70, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x41, 0x70, 0x70, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_netapp_v1_backup_vault_proto_rawDescOnce sync.Once
	file_google_cloud_netapp_v1_backup_vault_proto_rawDescData = file_google_cloud_netapp_v1_backup_vault_proto_rawDesc
)

func file_google_cloud_netapp_v1_backup_vault_proto_rawDescGZIP() []byte {
	file_google_cloud_netapp_v1_backup_vault_proto_rawDescOnce.Do(func() {
		file_google_cloud_netapp_v1_backup_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_netapp_v1_backup_vault_proto_rawDescData)
	})
	return file_google_cloud_netapp_v1_backup_vault_proto_rawDescData
}

var file_google_cloud_netapp_v1_backup_vault_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_netapp_v1_backup_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_netapp_v1_backup_vault_proto_goTypes = []interface{}{
	(BackupVault_State)(0),           // 0: google.cloud.netapp.v1.BackupVault.State
	(*BackupVault)(nil),              // 1: google.cloud.netapp.v1.BackupVault
	(*GetBackupVaultRequest)(nil),    // 2: google.cloud.netapp.v1.GetBackupVaultRequest
	(*ListBackupVaultsRequest)(nil),  // 3: google.cloud.netapp.v1.ListBackupVaultsRequest
	(*ListBackupVaultsResponse)(nil), // 4: google.cloud.netapp.v1.ListBackupVaultsResponse
	(*CreateBackupVaultRequest)(nil), // 5: google.cloud.netapp.v1.CreateBackupVaultRequest
	(*DeleteBackupVaultRequest)(nil), // 6: google.cloud.netapp.v1.DeleteBackupVaultRequest
	(*UpdateBackupVaultRequest)(nil), // 7: google.cloud.netapp.v1.UpdateBackupVaultRequest
	nil,                              // 8: google.cloud.netapp.v1.BackupVault.LabelsEntry
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),    // 10: google.protobuf.FieldMask
}
var file_google_cloud_netapp_v1_backup_vault_proto_depIdxs = []int32{
	0,  // 0: google.cloud.netapp.v1.BackupVault.state:type_name -> google.cloud.netapp.v1.BackupVault.State
	9,  // 1: google.cloud.netapp.v1.BackupVault.create_time:type_name -> google.protobuf.Timestamp
	8,  // 2: google.cloud.netapp.v1.BackupVault.labels:type_name -> google.cloud.netapp.v1.BackupVault.LabelsEntry
	1,  // 3: google.cloud.netapp.v1.ListBackupVaultsResponse.backup_vaults:type_name -> google.cloud.netapp.v1.BackupVault
	1,  // 4: google.cloud.netapp.v1.CreateBackupVaultRequest.backup_vault:type_name -> google.cloud.netapp.v1.BackupVault
	10, // 5: google.cloud.netapp.v1.UpdateBackupVaultRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 6: google.cloud.netapp.v1.UpdateBackupVaultRequest.backup_vault:type_name -> google.cloud.netapp.v1.BackupVault
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_netapp_v1_backup_vault_proto_init() }
func file_google_cloud_netapp_v1_backup_vault_proto_init() {
	if File_google_cloud_netapp_v1_backup_vault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupVault); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackupVaultRequest); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupVaultsRequest); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBackupVaultsResponse); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBackupVaultRequest); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBackupVaultRequest); i {
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
		file_google_cloud_netapp_v1_backup_vault_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBackupVaultRequest); i {
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
			RawDescriptor: file_google_cloud_netapp_v1_backup_vault_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_netapp_v1_backup_vault_proto_goTypes,
		DependencyIndexes: file_google_cloud_netapp_v1_backup_vault_proto_depIdxs,
		EnumInfos:         file_google_cloud_netapp_v1_backup_vault_proto_enumTypes,
		MessageInfos:      file_google_cloud_netapp_v1_backup_vault_proto_msgTypes,
	}.Build()
	File_google_cloud_netapp_v1_backup_vault_proto = out.File
	file_google_cloud_netapp_v1_backup_vault_proto_rawDesc = nil
	file_google_cloud_netapp_v1_backup_vault_proto_goTypes = nil
	file_google_cloud_netapp_v1_backup_vault_proto_depIdxs = nil
}
