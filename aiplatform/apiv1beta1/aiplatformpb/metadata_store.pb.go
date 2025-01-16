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
// source: google/cloud/aiplatform/v1beta1/metadata_store.proto

package aiplatformpb

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

// Instance of a metadata store. Contains a set of metadata that can be
// queried.
type MetadataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the MetadataStore instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when this MetadataStore was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this MetadataStore was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Customer-managed encryption key spec for a Metadata Store. If set, this
	// Metadata Store and all sub-resources of this Metadata Store are secured
	// using this key.
	EncryptionSpec *EncryptionSpec `protobuf:"bytes,5,opt,name=encryption_spec,json=encryptionSpec,proto3" json:"encryption_spec,omitempty"`
	// Description of the MetadataStore.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. State information of the MetadataStore.
	State *MetadataStore_MetadataStoreState `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	// Optional. Dataplex integration settings.
	DataplexConfig *MetadataStore_DataplexConfig `protobuf:"bytes,8,opt,name=dataplex_config,json=dataplexConfig,proto3" json:"dataplex_config,omitempty"`
}

func (x *MetadataStore) Reset() {
	*x = MetadataStore{}
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataStore) ProtoMessage() {}

func (x *MetadataStore) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataStore.ProtoReflect.Descriptor instead.
func (*MetadataStore) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetadataStore) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *MetadataStore) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *MetadataStore) GetEncryptionSpec() *EncryptionSpec {
	if x != nil {
		return x.EncryptionSpec
	}
	return nil
}

func (x *MetadataStore) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MetadataStore) GetState() *MetadataStore_MetadataStoreState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *MetadataStore) GetDataplexConfig() *MetadataStore_DataplexConfig {
	if x != nil {
		return x.DataplexConfig
	}
	return nil
}

// Represents state information for a MetadataStore.
type MetadataStore_MetadataStoreState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The disk utilization of the MetadataStore in bytes.
	DiskUtilizationBytes int64 `protobuf:"varint,1,opt,name=disk_utilization_bytes,json=diskUtilizationBytes,proto3" json:"disk_utilization_bytes,omitempty"`
}

func (x *MetadataStore_MetadataStoreState) Reset() {
	*x = MetadataStore_MetadataStoreState{}
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataStore_MetadataStoreState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataStore_MetadataStoreState) ProtoMessage() {}

func (x *MetadataStore_MetadataStoreState) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataStore_MetadataStoreState.ProtoReflect.Descriptor instead.
func (*MetadataStore_MetadataStoreState) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MetadataStore_MetadataStoreState) GetDiskUtilizationBytes() int64 {
	if x != nil {
		return x.DiskUtilizationBytes
	}
	return 0
}

// Represents Dataplex integration settings.
type MetadataStore_DataplexConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Whether or not Data Lineage synchronization is enabled for
	// Vertex Pipelines.
	EnabledPipelinesLineage bool `protobuf:"varint,1,opt,name=enabled_pipelines_lineage,json=enabledPipelinesLineage,proto3" json:"enabled_pipelines_lineage,omitempty"`
}

func (x *MetadataStore_DataplexConfig) Reset() {
	*x = MetadataStore_DataplexConfig{}
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataStore_DataplexConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataStore_DataplexConfig) ProtoMessage() {}

func (x *MetadataStore_DataplexConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataStore_DataplexConfig.ProtoReflect.Descriptor instead.
func (*MetadataStore_DataplexConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MetadataStore_DataplexConfig) GetEnabledPipelinesLineage() bool {
	if x != nil {
		return x.EnabledPipelinesLineage
	}
	return false
}

var File_google_cloud_aiplatform_v1beta1_metadata_store_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x06, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x4a, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x74, 0x69,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x51, 0x0a,
	0x0e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3f, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65,
	0x3a, 0x75, 0xea, 0x41, 0x72, 0x0a, 0x27, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x47,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x42, 0xe4, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1beta1_metadata_store_proto_goTypes = []any{
	(*MetadataStore)(nil),                    // 0: google.cloud.aiplatform.v1beta1.MetadataStore
	(*MetadataStore_MetadataStoreState)(nil), // 1: google.cloud.aiplatform.v1beta1.MetadataStore.MetadataStoreState
	(*MetadataStore_DataplexConfig)(nil),     // 2: google.cloud.aiplatform.v1beta1.MetadataStore.DataplexConfig
	(*timestamppb.Timestamp)(nil),            // 3: google.protobuf.Timestamp
	(*EncryptionSpec)(nil),                   // 4: google.cloud.aiplatform.v1beta1.EncryptionSpec
}
var file_google_cloud_aiplatform_v1beta1_metadata_store_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1beta1.MetadataStore.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.aiplatform.v1beta1.MetadataStore.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.cloud.aiplatform.v1beta1.MetadataStore.encryption_spec:type_name -> google.cloud.aiplatform.v1beta1.EncryptionSpec
	1, // 3: google.cloud.aiplatform.v1beta1.MetadataStore.state:type_name -> google.cloud.aiplatform.v1beta1.MetadataStore.MetadataStoreState
	2, // 4: google.cloud.aiplatform.v1beta1.MetadataStore.dataplex_config:type_name -> google.cloud.aiplatform.v1beta1.MetadataStore.DataplexConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_metadata_store_proto_init() }
func file_google_cloud_aiplatform_v1beta1_metadata_store_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_metadata_store_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_encryption_spec_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_metadata_store_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_metadata_store_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_metadata_store_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_metadata_store_proto = out.File
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_metadata_store_proto_depIdxs = nil
}
