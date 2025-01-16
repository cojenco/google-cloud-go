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
// source: google/cloud/oracledatabase/v1/db_system_shape.proto

package oracledatabasepb

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

// Details of the Database System Shapes resource.
// https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbSystemShapeSummary/
type DbSystemShape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The name of the Database System Shape resource with the format:
	// projects/{project}/locations/{region}/dbSystemShapes/{db_system_shape}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. shape
	Shape string `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// Optional. Minimum number of database servers.
	MinNodeCount int32 `protobuf:"varint,3,opt,name=min_node_count,json=minNodeCount,proto3" json:"min_node_count,omitempty"`
	// Optional. Maximum number of database servers.
	MaxNodeCount int32 `protobuf:"varint,4,opt,name=max_node_count,json=maxNodeCount,proto3" json:"max_node_count,omitempty"`
	// Optional. Minimum number of storage servers.
	MinStorageCount int32 `protobuf:"varint,5,opt,name=min_storage_count,json=minStorageCount,proto3" json:"min_storage_count,omitempty"`
	// Optional. Maximum number of storage servers.
	MaxStorageCount int32 `protobuf:"varint,6,opt,name=max_storage_count,json=maxStorageCount,proto3" json:"max_storage_count,omitempty"`
	// Optional. Number of cores per node.
	AvailableCoreCountPerNode int32 `protobuf:"varint,7,opt,name=available_core_count_per_node,json=availableCoreCountPerNode,proto3" json:"available_core_count_per_node,omitempty"`
	// Optional. Memory per database server node in gigabytes.
	AvailableMemoryPerNodeGb int32 `protobuf:"varint,8,opt,name=available_memory_per_node_gb,json=availableMemoryPerNodeGb,proto3" json:"available_memory_per_node_gb,omitempty"`
	// Optional. Storage per storage server in terabytes.
	AvailableDataStorageTb int32 `protobuf:"varint,9,opt,name=available_data_storage_tb,json=availableDataStorageTb,proto3" json:"available_data_storage_tb,omitempty"`
	// Optional. Minimum core count per node.
	MinCoreCountPerNode int32 `protobuf:"varint,10,opt,name=min_core_count_per_node,json=minCoreCountPerNode,proto3" json:"min_core_count_per_node,omitempty"`
	// Optional. Minimum memory per node in gigabytes.
	MinMemoryPerNodeGb int32 `protobuf:"varint,11,opt,name=min_memory_per_node_gb,json=minMemoryPerNodeGb,proto3" json:"min_memory_per_node_gb,omitempty"`
	// Optional. Minimum node storage per database server in gigabytes.
	MinDbNodeStoragePerNodeGb int32 `protobuf:"varint,12,opt,name=min_db_node_storage_per_node_gb,json=minDbNodeStoragePerNodeGb,proto3" json:"min_db_node_storage_per_node_gb,omitempty"`
}

func (x *DbSystemShape) Reset() {
	*x = DbSystemShape{}
	mi := &file_google_cloud_oracledatabase_v1_db_system_shape_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DbSystemShape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbSystemShape) ProtoMessage() {}

func (x *DbSystemShape) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oracledatabase_v1_db_system_shape_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbSystemShape.ProtoReflect.Descriptor instead.
func (*DbSystemShape) Descriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescGZIP(), []int{0}
}

func (x *DbSystemShape) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DbSystemShape) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

func (x *DbSystemShape) GetMinNodeCount() int32 {
	if x != nil {
		return x.MinNodeCount
	}
	return 0
}

func (x *DbSystemShape) GetMaxNodeCount() int32 {
	if x != nil {
		return x.MaxNodeCount
	}
	return 0
}

func (x *DbSystemShape) GetMinStorageCount() int32 {
	if x != nil {
		return x.MinStorageCount
	}
	return 0
}

func (x *DbSystemShape) GetMaxStorageCount() int32 {
	if x != nil {
		return x.MaxStorageCount
	}
	return 0
}

func (x *DbSystemShape) GetAvailableCoreCountPerNode() int32 {
	if x != nil {
		return x.AvailableCoreCountPerNode
	}
	return 0
}

func (x *DbSystemShape) GetAvailableMemoryPerNodeGb() int32 {
	if x != nil {
		return x.AvailableMemoryPerNodeGb
	}
	return 0
}

func (x *DbSystemShape) GetAvailableDataStorageTb() int32 {
	if x != nil {
		return x.AvailableDataStorageTb
	}
	return 0
}

func (x *DbSystemShape) GetMinCoreCountPerNode() int32 {
	if x != nil {
		return x.MinCoreCountPerNode
	}
	return 0
}

func (x *DbSystemShape) GetMinMemoryPerNodeGb() int32 {
	if x != nil {
		return x.MinMemoryPerNodeGb
	}
	return 0
}

func (x *DbSystemShape) GetMinDbNodeStoragePerNodeGb() int32 {
	if x != nil {
		return x.MinDbNodeStoragePerNodeGb
	}
	return 0
}

var File_google_cloud_oracledatabase_v1_db_system_shape_proto protoreflect.FileDescriptor

var file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa1, 0x06, 0x0a, 0x0d, 0x44, 0x62, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f,
	0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f,
	0x6d, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2f, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0f, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x45, 0x0a, 0x1d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x19, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x1c, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x18, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x50, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x62, 0x12, 0x3e, 0x0a, 0x19,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x16, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x62, 0x12, 0x39, 0x0a, 0x17,
	0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x16, 0x6d, 0x69, 0x6e, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67,
	0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x6d, 0x69,
	0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x62,
	0x12, 0x47, 0x0a, 0x1f, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x67, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x19,
	0x6d, 0x69, 0x6e, 0x44, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x62, 0x3a, 0x9a, 0x01, 0xea, 0x41, 0x96, 0x01,
	0x0a, 0x2b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x62, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x48, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x62, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x7d, 0x2a, 0x0e, 0x64, 0x62, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73, 0x32, 0x0d, 0x64, 0x62, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x53, 0x68, 0x61, 0x70, 0x65, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x44,
	0x62, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0x3b, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescOnce sync.Once
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescData = file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDesc
)

func file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescGZIP() []byte {
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescOnce.Do(func() {
		file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescData)
	})
	return file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDescData
}

var file_google_cloud_oracledatabase_v1_db_system_shape_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_oracledatabase_v1_db_system_shape_proto_goTypes = []any{
	(*DbSystemShape)(nil), // 0: google.cloud.oracledatabase.v1.DbSystemShape
}
var file_google_cloud_oracledatabase_v1_db_system_shape_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_oracledatabase_v1_db_system_shape_proto_init() }
func file_google_cloud_oracledatabase_v1_db_system_shape_proto_init() {
	if File_google_cloud_oracledatabase_v1_db_system_shape_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_oracledatabase_v1_db_system_shape_proto_goTypes,
		DependencyIndexes: file_google_cloud_oracledatabase_v1_db_system_shape_proto_depIdxs,
		MessageInfos:      file_google_cloud_oracledatabase_v1_db_system_shape_proto_msgTypes,
	}.Build()
	File_google_cloud_oracledatabase_v1_db_system_shape_proto = out.File
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_rawDesc = nil
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_goTypes = nil
	file_google_cloud_oracledatabase_v1_db_system_shape_proto_depIdxs = nil
}
