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
// source: google/cloud/securitycenter/v2/attack_path.proto

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

// The type of the incoming attack step node.
type AttackPath_AttackPathNode_NodeType int32

const (
	// Type not specified
	AttackPath_AttackPathNode_NODE_TYPE_UNSPECIFIED AttackPath_AttackPathNode_NodeType = 0
	// Incoming edge joined with AND
	AttackPath_AttackPathNode_NODE_TYPE_AND AttackPath_AttackPathNode_NodeType = 1
	// Incoming edge joined with OR
	AttackPath_AttackPathNode_NODE_TYPE_OR AttackPath_AttackPathNode_NodeType = 2
	// Incoming edge is defense
	AttackPath_AttackPathNode_NODE_TYPE_DEFENSE AttackPath_AttackPathNode_NodeType = 3
	// Incoming edge is attacker
	AttackPath_AttackPathNode_NODE_TYPE_ATTACKER AttackPath_AttackPathNode_NodeType = 4
)

// Enum value maps for AttackPath_AttackPathNode_NodeType.
var (
	AttackPath_AttackPathNode_NodeType_name = map[int32]string{
		0: "NODE_TYPE_UNSPECIFIED",
		1: "NODE_TYPE_AND",
		2: "NODE_TYPE_OR",
		3: "NODE_TYPE_DEFENSE",
		4: "NODE_TYPE_ATTACKER",
	}
	AttackPath_AttackPathNode_NodeType_value = map[string]int32{
		"NODE_TYPE_UNSPECIFIED": 0,
		"NODE_TYPE_AND":         1,
		"NODE_TYPE_OR":          2,
		"NODE_TYPE_DEFENSE":     3,
		"NODE_TYPE_ATTACKER":    4,
	}
)

func (x AttackPath_AttackPathNode_NodeType) Enum() *AttackPath_AttackPathNode_NodeType {
	p := new(AttackPath_AttackPathNode_NodeType)
	*p = x
	return p
}

func (x AttackPath_AttackPathNode_NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttackPath_AttackPathNode_NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v2_attack_path_proto_enumTypes[0].Descriptor()
}

func (AttackPath_AttackPathNode_NodeType) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v2_attack_path_proto_enumTypes[0]
}

func (x AttackPath_AttackPathNode_NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttackPath_AttackPathNode_NodeType.Descriptor instead.
func (AttackPath_AttackPathNode_NodeType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A path that an attacker could take to reach an exposed resource.
type AttackPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The attack path name, for example,
	//
	//	`organizations/12/simulations/34/valuedResources/56/attackPaths/78`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of nodes that exist in this attack path.
	PathNodes []*AttackPath_AttackPathNode `protobuf:"bytes,2,rep,name=path_nodes,json=pathNodes,proto3" json:"path_nodes,omitempty"`
	// A list of the edges between nodes in this attack path.
	Edges []*AttackPath_AttackPathEdge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *AttackPath) Reset() {
	*x = AttackPath{}
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttackPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackPath) ProtoMessage() {}

func (x *AttackPath) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackPath.ProtoReflect.Descriptor instead.
func (*AttackPath) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0}
}

func (x *AttackPath) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttackPath) GetPathNodes() []*AttackPath_AttackPathNode {
	if x != nil {
		return x.PathNodes
	}
	return nil
}

func (x *AttackPath) GetEdges() []*AttackPath_AttackPathEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

// Represents one point that an attacker passes through in this attack path.
type AttackPath_AttackPathNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the resource at this point in the attack path.
	// The format of the name follows the Cloud Asset Inventory [resource
	// name
	// format](https://cloud.google.com/asset-inventory/docs/resource-name-format)
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The [supported resource
	// type](https://cloud.google.com/asset-inventory/docs/supported-asset-types)
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Human-readable name of this resource.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The findings associated with this node in the attack path.
	AssociatedFindings []*AttackPath_AttackPathNode_PathNodeAssociatedFinding `protobuf:"bytes,4,rep,name=associated_findings,json=associatedFindings,proto3" json:"associated_findings,omitempty"`
	// Unique id of the attack path node.
	Uuid string `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// A list of attack step nodes that exist in this attack path node.
	AttackSteps []*AttackPath_AttackPathNode_AttackStepNode `protobuf:"bytes,6,rep,name=attack_steps,json=attackSteps,proto3" json:"attack_steps,omitempty"`
}

func (x *AttackPath_AttackPathNode) Reset() {
	*x = AttackPath_AttackPathNode{}
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttackPath_AttackPathNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackPath_AttackPathNode) ProtoMessage() {}

func (x *AttackPath_AttackPathNode) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackPath_AttackPathNode.ProtoReflect.Descriptor instead.
func (*AttackPath_AttackPathNode) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AttackPath_AttackPathNode) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AttackPath_AttackPathNode) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *AttackPath_AttackPathNode) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AttackPath_AttackPathNode) GetAssociatedFindings() []*AttackPath_AttackPathNode_PathNodeAssociatedFinding {
	if x != nil {
		return x.AssociatedFindings
	}
	return nil
}

func (x *AttackPath_AttackPathNode) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AttackPath_AttackPathNode) GetAttackSteps() []*AttackPath_AttackPathNode_AttackStepNode {
	if x != nil {
		return x.AttackSteps
	}
	return nil
}

// Represents a connection between a source node and a destination node in
// this attack path.
type AttackPath_AttackPathEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The attack node uuid of the source node.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The attack node uuid of the destination node.
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *AttackPath_AttackPathEdge) Reset() {
	*x = AttackPath_AttackPathEdge{}
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttackPath_AttackPathEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackPath_AttackPathEdge) ProtoMessage() {}

func (x *AttackPath_AttackPathEdge) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackPath_AttackPathEdge.ProtoReflect.Descriptor instead.
func (*AttackPath_AttackPathEdge) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AttackPath_AttackPathEdge) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *AttackPath_AttackPathEdge) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

// A finding that is associated with this node in the attack path.
type AttackPath_AttackPathNode_PathNodeAssociatedFinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical name of the associated findings. Example:
	// `organizations/123/sources/456/findings/789`
	CanonicalFinding string `protobuf:"bytes,1,opt,name=canonical_finding,json=canonicalFinding,proto3" json:"canonical_finding,omitempty"`
	// The additional taxonomy group within findings from a given source.
	FindingCategory string `protobuf:"bytes,2,opt,name=finding_category,json=findingCategory,proto3" json:"finding_category,omitempty"`
	// Full resource name of the finding.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) Reset() {
	*x = AttackPath_AttackPathNode_PathNodeAssociatedFinding{}
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackPath_AttackPathNode_PathNodeAssociatedFinding) ProtoMessage() {}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackPath_AttackPathNode_PathNodeAssociatedFinding.ProtoReflect.Descriptor instead.
func (*AttackPath_AttackPathNode_PathNodeAssociatedFinding) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) GetCanonicalFinding() string {
	if x != nil {
		return x.CanonicalFinding
	}
	return ""
}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) GetFindingCategory() string {
	if x != nil {
		return x.FindingCategory
	}
	return ""
}

func (x *AttackPath_AttackPathNode_PathNodeAssociatedFinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Detailed steps the attack can take between path nodes.
type AttackPath_AttackPathNode_AttackStepNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID for one Node
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Attack step type. Can be either AND, OR or DEFENSE
	Type AttackPath_AttackPathNode_NodeType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.securitycenter.v2.AttackPath_AttackPathNode_NodeType" json:"type,omitempty"`
	// User friendly name of the attack step
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Attack step labels for metadata
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Attack step description
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AttackPath_AttackPathNode_AttackStepNode) Reset() {
	*x = AttackPath_AttackPathNode_AttackStepNode{}
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttackPath_AttackPathNode_AttackStepNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackPath_AttackPathNode_AttackStepNode) ProtoMessage() {}

func (x *AttackPath_AttackPathNode_AttackStepNode) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackPath_AttackPathNode_AttackStepNode.ProtoReflect.Descriptor instead.
func (*AttackPath_AttackPathNode_AttackStepNode) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *AttackPath_AttackPathNode_AttackStepNode) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AttackPath_AttackPathNode_AttackStepNode) GetType() AttackPath_AttackPathNode_NodeType {
	if x != nil {
		return x.Type
	}
	return AttackPath_AttackPathNode_NODE_TYPE_UNSPECIFIED
}

func (x *AttackPath_AttackPathNode_AttackStepNode) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AttackPath_AttackPathNode_AttackStepNode) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AttackPath_AttackPathNode_AttackStepNode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_securitycenter_v2_attack_path_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_attack_path_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x0c,
	0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x58, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x05, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68,
	0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x1a, 0xee, 0x07, 0x0a, 0x0e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x12, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x6b, 0x0a,
	0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x2e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x73, 0x1a, 0x87, 0x01, 0x0a, 0x19, 0x50,
	0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x6f,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x1a, 0xea, 0x02, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53,
	0x74, 0x65, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x56, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61,
	0x74, 0x68, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x79, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x44, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x45, 0x4e,
	0x53, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x04, 0x1a, 0x4a, 0x0a, 0x0e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x45, 0x64, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0xc3, 0x02, 0xea, 0x41, 0xbf, 0x02, 0x0a,
	0x28, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x71, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x73, 0x2f, 0x7b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x7d, 0x12, 0x86, 0x01, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x73, 0x2f, 0x7b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x7d, 0x2a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74,
	0x68, 0x73, 0x32, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x42, 0xe9,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x0f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_attack_path_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_attack_path_proto_rawDescData = file_google_cloud_securitycenter_v2_attack_path_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_attack_path_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_attack_path_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_attack_path_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_attack_path_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_attack_path_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_attack_path_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_securitycenter_v2_attack_path_proto_goTypes = []any{
	(AttackPath_AttackPathNode_NodeType)(0),                     // 0: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.NodeType
	(*AttackPath)(nil),                                          // 1: google.cloud.securitycenter.v2.AttackPath
	(*AttackPath_AttackPathNode)(nil),                           // 2: google.cloud.securitycenter.v2.AttackPath.AttackPathNode
	(*AttackPath_AttackPathEdge)(nil),                           // 3: google.cloud.securitycenter.v2.AttackPath.AttackPathEdge
	(*AttackPath_AttackPathNode_PathNodeAssociatedFinding)(nil), // 4: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding
	(*AttackPath_AttackPathNode_AttackStepNode)(nil),            // 5: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode
	nil, // 6: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.LabelsEntry
}
var file_google_cloud_securitycenter_v2_attack_path_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v2.AttackPath.path_nodes:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathNode
	3, // 1: google.cloud.securitycenter.v2.AttackPath.edges:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathEdge
	4, // 2: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.associated_findings:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathNode.PathNodeAssociatedFinding
	5, // 3: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.attack_steps:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode
	0, // 4: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.type:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathNode.NodeType
	6, // 5: google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.labels:type_name -> google.cloud.securitycenter.v2.AttackPath.AttackPathNode.AttackStepNode.LabelsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_attack_path_proto_init() }
func file_google_cloud_securitycenter_v2_attack_path_proto_init() {
	if File_google_cloud_securitycenter_v2_attack_path_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_attack_path_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_attack_path_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_attack_path_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v2_attack_path_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v2_attack_path_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_attack_path_proto = out.File
	file_google_cloud_securitycenter_v2_attack_path_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_attack_path_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_attack_path_proto_depIdxs = nil
}
