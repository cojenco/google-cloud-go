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
// source: google/cloud/oracledatabase/v1/db_node.proto

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

// The various lifecycle states of the database node.
type DbNodeProperties_State int32

const (
	// Default unspecified value.
	DbNodeProperties_STATE_UNSPECIFIED DbNodeProperties_State = 0
	// Indicates that the resource is in provisioning state.
	DbNodeProperties_PROVISIONING DbNodeProperties_State = 1
	// Indicates that the resource is in available state.
	DbNodeProperties_AVAILABLE DbNodeProperties_State = 2
	// Indicates that the resource is in updating state.
	DbNodeProperties_UPDATING DbNodeProperties_State = 3
	// Indicates that the resource is in stopping state.
	DbNodeProperties_STOPPING DbNodeProperties_State = 4
	// Indicates that the resource is in stopped state.
	DbNodeProperties_STOPPED DbNodeProperties_State = 5
	// Indicates that the resource is in starting state.
	DbNodeProperties_STARTING DbNodeProperties_State = 6
	// Indicates that the resource is in terminating state.
	DbNodeProperties_TERMINATING DbNodeProperties_State = 7
	// Indicates that the resource is in terminated state.
	DbNodeProperties_TERMINATED DbNodeProperties_State = 8
	// Indicates that the resource is in failed state.
	DbNodeProperties_FAILED DbNodeProperties_State = 9
)

// Enum value maps for DbNodeProperties_State.
var (
	DbNodeProperties_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PROVISIONING",
		2: "AVAILABLE",
		3: "UPDATING",
		4: "STOPPING",
		5: "STOPPED",
		6: "STARTING",
		7: "TERMINATING",
		8: "TERMINATED",
		9: "FAILED",
	}
	DbNodeProperties_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PROVISIONING":      1,
		"AVAILABLE":         2,
		"UPDATING":          3,
		"STOPPING":          4,
		"STOPPED":           5,
		"STARTING":          6,
		"TERMINATING":       7,
		"TERMINATED":        8,
		"FAILED":            9,
	}
)

func (x DbNodeProperties_State) Enum() *DbNodeProperties_State {
	p := new(DbNodeProperties_State)
	*p = x
	return p
}

func (x DbNodeProperties_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DbNodeProperties_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_oracledatabase_v1_db_node_proto_enumTypes[0].Descriptor()
}

func (DbNodeProperties_State) Type() protoreflect.EnumType {
	return &file_google_cloud_oracledatabase_v1_db_node_proto_enumTypes[0]
}

func (x DbNodeProperties_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DbNodeProperties_State.Descriptor instead.
func (DbNodeProperties_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_db_node_proto_rawDescGZIP(), []int{1, 0}
}

// Details of the database node resource.
// https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbNode/
type DbNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The name of the database node resource in the following format:
	// projects/{project}/locations/{location}/cloudVmClusters/{cloud_vm_cluster}/dbNodes/{db_node}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Various properties of the database node.
	Properties *DbNodeProperties `protobuf:"bytes,3,opt,name=properties,proto3" json:"properties,omitempty"`
}

func (x *DbNode) Reset() {
	*x = DbNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbNode) ProtoMessage() {}

func (x *DbNode) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbNode.ProtoReflect.Descriptor instead.
func (*DbNode) Descriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_db_node_proto_rawDescGZIP(), []int{0}
}

func (x *DbNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DbNode) GetProperties() *DbNodeProperties {
	if x != nil {
		return x.Properties
	}
	return nil
}

// Various properties and settings associated with Db node.
type DbNodeProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. OCID of database node.
	Ocid string `protobuf:"bytes,1,opt,name=ocid,proto3" json:"ocid,omitempty"`
	// Optional. OCPU count per database node.
	OcpuCount int32 `protobuf:"varint,2,opt,name=ocpu_count,json=ocpuCount,proto3" json:"ocpu_count,omitempty"`
	// Memory allocated in GBs.
	MemorySizeGb int32 `protobuf:"varint,3,opt,name=memory_size_gb,json=memorySizeGb,proto3" json:"memory_size_gb,omitempty"`
	// Optional. Local storage per database node.
	DbNodeStorageSizeGb int32 `protobuf:"varint,4,opt,name=db_node_storage_size_gb,json=dbNodeStorageSizeGb,proto3" json:"db_node_storage_size_gb,omitempty"`
	// Optional. Database server OCID.
	DbServerOcid string `protobuf:"bytes,5,opt,name=db_server_ocid,json=dbServerOcid,proto3" json:"db_server_ocid,omitempty"`
	// Optional. DNS
	Hostname string `protobuf:"bytes,8,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Output only. State of the database node.
	State DbNodeProperties_State `protobuf:"varint,9,opt,name=state,proto3,enum=google.cloud.oracledatabase.v1.DbNodeProperties_State" json:"state,omitempty"`
	// Total CPU core count of the database node.
	TotalCpuCoreCount int32 `protobuf:"varint,10,opt,name=total_cpu_core_count,json=totalCpuCoreCount,proto3" json:"total_cpu_core_count,omitempty"`
}

func (x *DbNodeProperties) Reset() {
	*x = DbNodeProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbNodeProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbNodeProperties) ProtoMessage() {}

func (x *DbNodeProperties) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbNodeProperties.ProtoReflect.Descriptor instead.
func (*DbNodeProperties) Descriptor() ([]byte, []int) {
	return file_google_cloud_oracledatabase_v1_db_node_proto_rawDescGZIP(), []int{1}
}

func (x *DbNodeProperties) GetOcid() string {
	if x != nil {
		return x.Ocid
	}
	return ""
}

func (x *DbNodeProperties) GetOcpuCount() int32 {
	if x != nil {
		return x.OcpuCount
	}
	return 0
}

func (x *DbNodeProperties) GetMemorySizeGb() int32 {
	if x != nil {
		return x.MemorySizeGb
	}
	return 0
}

func (x *DbNodeProperties) GetDbNodeStorageSizeGb() int32 {
	if x != nil {
		return x.DbNodeStorageSizeGb
	}
	return 0
}

func (x *DbNodeProperties) GetDbServerOcid() string {
	if x != nil {
		return x.DbServerOcid
	}
	return ""
}

func (x *DbNodeProperties) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DbNodeProperties) GetState() DbNodeProperties_State {
	if x != nil {
		return x.State
	}
	return DbNodeProperties_STATE_UNSPECIFIED
}

func (x *DbNodeProperties) GetTotalCpuCoreCount() int32 {
	if x != nil {
		return x.TotalCpuCoreCount
	}
	return 0
}

var File_google_cloud_oracledatabase_v1_db_node_proto protoreflect.FileDescriptor

var file_google_cloud_oracledatabase_v1_db_node_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x62, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x06, 0x44,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x55,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x99, 0x01, 0xea, 0x41, 0x95, 0x01, 0x0a, 0x24, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x62, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x5c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x56, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x76, 0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x2f, 0x64,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x62, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x7d,
	0x2a, 0x07, 0x64, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x06, 0x64, 0x62, 0x4e, 0x6f, 0x64,
	0x65, 0x22, 0xa6, 0x04, 0x0a, 0x10, 0x44, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x6f, 0x63, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6f, 0x63, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0a, 0x6f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x6f, 0x63, 0x70, 0x75, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x67, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x62, 0x12, 0x39, 0x0a, 0x17, 0x64, 0x62, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x67, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x13, 0x64, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x47, 0x62, 0x12, 0x29, 0x0a, 0x0e, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6f, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0c, 0x64, 0x62, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x63, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x51, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49,
	0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x42, 0xe5, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x44, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0x3b, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_oracledatabase_v1_db_node_proto_rawDescOnce sync.Once
	file_google_cloud_oracledatabase_v1_db_node_proto_rawDescData = file_google_cloud_oracledatabase_v1_db_node_proto_rawDesc
)

func file_google_cloud_oracledatabase_v1_db_node_proto_rawDescGZIP() []byte {
	file_google_cloud_oracledatabase_v1_db_node_proto_rawDescOnce.Do(func() {
		file_google_cloud_oracledatabase_v1_db_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_oracledatabase_v1_db_node_proto_rawDescData)
	})
	return file_google_cloud_oracledatabase_v1_db_node_proto_rawDescData
}

var file_google_cloud_oracledatabase_v1_db_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_oracledatabase_v1_db_node_proto_goTypes = []any{
	(DbNodeProperties_State)(0), // 0: google.cloud.oracledatabase.v1.DbNodeProperties.State
	(*DbNode)(nil),              // 1: google.cloud.oracledatabase.v1.DbNode
	(*DbNodeProperties)(nil),    // 2: google.cloud.oracledatabase.v1.DbNodeProperties
}
var file_google_cloud_oracledatabase_v1_db_node_proto_depIdxs = []int32{
	2, // 0: google.cloud.oracledatabase.v1.DbNode.properties:type_name -> google.cloud.oracledatabase.v1.DbNodeProperties
	0, // 1: google.cloud.oracledatabase.v1.DbNodeProperties.state:type_name -> google.cloud.oracledatabase.v1.DbNodeProperties.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_oracledatabase_v1_db_node_proto_init() }
func file_google_cloud_oracledatabase_v1_db_node_proto_init() {
	if File_google_cloud_oracledatabase_v1_db_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DbNode); i {
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
		file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DbNodeProperties); i {
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
			RawDescriptor: file_google_cloud_oracledatabase_v1_db_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_oracledatabase_v1_db_node_proto_goTypes,
		DependencyIndexes: file_google_cloud_oracledatabase_v1_db_node_proto_depIdxs,
		EnumInfos:         file_google_cloud_oracledatabase_v1_db_node_proto_enumTypes,
		MessageInfos:      file_google_cloud_oracledatabase_v1_db_node_proto_msgTypes,
	}.Build()
	File_google_cloud_oracledatabase_v1_db_node_proto = out.File
	file_google_cloud_oracledatabase_v1_db_node_proto_rawDesc = nil
	file_google_cloud_oracledatabase_v1_db_node_proto_goTypes = nil
	file_google_cloud_oracledatabase_v1_db_node_proto_depIdxs = nil
}
