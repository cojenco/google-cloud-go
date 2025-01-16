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
// source: google/cloud/securityposture/v1/org_policy_config.proto

package securityposturepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	expr "google.golang.org/genproto/googleapis/type/expr"
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

// The operation for which this constraint will be applied. To apply this
// constraint only when creating new VMs, the `method_types` should be
// `CREATE` only. To apply this constraint when creating or deleting
// VMs, the `method_types` should be `CREATE` and `DELETE`.
//
// `UPDATE` only custom constraints are not supported. Use `CREATE` or
// `CREATE, UPDATE`.
type CustomConstraint_MethodType int32

const (
	// Unspecified. Results in an error.
	CustomConstraint_METHOD_TYPE_UNSPECIFIED CustomConstraint_MethodType = 0
	// Constraint applied when creating the resource.
	CustomConstraint_CREATE CustomConstraint_MethodType = 1
	// Constraint applied when updating the resource.
	CustomConstraint_UPDATE CustomConstraint_MethodType = 2
	// Constraint applied when deleting the resource.
	// Not supported yet.
	CustomConstraint_DELETE CustomConstraint_MethodType = 3
)

// Enum value maps for CustomConstraint_MethodType.
var (
	CustomConstraint_MethodType_name = map[int32]string{
		0: "METHOD_TYPE_UNSPECIFIED",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	CustomConstraint_MethodType_value = map[string]int32{
		"METHOD_TYPE_UNSPECIFIED": 0,
		"CREATE":                  1,
		"UPDATE":                  2,
		"DELETE":                  3,
	}
)

func (x CustomConstraint_MethodType) Enum() *CustomConstraint_MethodType {
	p := new(CustomConstraint_MethodType)
	*p = x
	return p
}

func (x CustomConstraint_MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomConstraint_MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes[0].Descriptor()
}

func (CustomConstraint_MethodType) Type() protoreflect.EnumType {
	return &file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes[0]
}

func (x CustomConstraint_MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomConstraint_MethodType.Descriptor instead.
func (CustomConstraint_MethodType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP(), []int{1, 0}
}

// Allow or deny type.
type CustomConstraint_ActionType int32

const (
	// Unspecified. Results in an error.
	CustomConstraint_ACTION_TYPE_UNSPECIFIED CustomConstraint_ActionType = 0
	// Allowed action type.
	CustomConstraint_ALLOW CustomConstraint_ActionType = 1
	// Deny action type.
	CustomConstraint_DENY CustomConstraint_ActionType = 2
)

// Enum value maps for CustomConstraint_ActionType.
var (
	CustomConstraint_ActionType_name = map[int32]string{
		0: "ACTION_TYPE_UNSPECIFIED",
		1: "ALLOW",
		2: "DENY",
	}
	CustomConstraint_ActionType_value = map[string]int32{
		"ACTION_TYPE_UNSPECIFIED": 0,
		"ALLOW":                   1,
		"DENY":                    2,
	}
)

func (x CustomConstraint_ActionType) Enum() *CustomConstraint_ActionType {
	p := new(CustomConstraint_ActionType)
	*p = x
	return p
}

func (x CustomConstraint_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomConstraint_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes[1].Descriptor()
}

func (CustomConstraint_ActionType) Type() protoreflect.EnumType {
	return &file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes[1]
}

func (x CustomConstraint_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomConstraint_ActionType.Descriptor instead.
func (CustomConstraint_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP(), []int{1, 1}
}

// A rule used to express this policy.
type PolicyRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*PolicyRule_Values
	//	*PolicyRule_AllowAll
	//	*PolicyRule_DenyAll
	//	*PolicyRule_Enforce
	Kind isPolicyRule_Kind `protobuf_oneof:"kind"`
	// A condition which determines whether this rule is used
	// in the evaluation of the policy. When set, the `expression` field in
	// the `Expr' must include from 1 to 10 subexpressions, joined by the "||"
	// or "&&" operators. Each subexpression must be of the form
	// "resource.matchTag('<ORG_ID>/tag_key_short_name,
	// 'tag_value_short_name')" or "resource.matchTagId('tagKeys/key_id',
	// 'tagValues/value_id')" where key_name and value_name are the resource
	// names for Label Keys and Values. These names are available from the Tag
	// Manager Service. An example expression is:
	// "resource.matchTag('123456789/environment,
	// 'prod')" or "resource.matchTagId('tagKeys/123',
	// 'tagValues/456')".
	Condition *expr.Expr `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *PolicyRule) Reset() {
	*x = PolicyRule{}
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRule) ProtoMessage() {}

func (x *PolicyRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRule.ProtoReflect.Descriptor instead.
func (*PolicyRule) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP(), []int{0}
}

func (m *PolicyRule) GetKind() isPolicyRule_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *PolicyRule) GetValues() *PolicyRule_StringValues {
	if x, ok := x.GetKind().(*PolicyRule_Values); ok {
		return x.Values
	}
	return nil
}

func (x *PolicyRule) GetAllowAll() bool {
	if x, ok := x.GetKind().(*PolicyRule_AllowAll); ok {
		return x.AllowAll
	}
	return false
}

func (x *PolicyRule) GetDenyAll() bool {
	if x, ok := x.GetKind().(*PolicyRule_DenyAll); ok {
		return x.DenyAll
	}
	return false
}

func (x *PolicyRule) GetEnforce() bool {
	if x, ok := x.GetKind().(*PolicyRule_Enforce); ok {
		return x.Enforce
	}
	return false
}

func (x *PolicyRule) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

type isPolicyRule_Kind interface {
	isPolicyRule_Kind()
}

type PolicyRule_Values struct {
	// List of values to be used for this policy rule. This field can be set
	// only in policies for list constraints.
	Values *PolicyRule_StringValues `protobuf:"bytes,1,opt,name=values,proto3,oneof"`
}

type PolicyRule_AllowAll struct {
	// Setting this to true means that all values are allowed. This field can
	// be set only in policies for list constraints.
	AllowAll bool `protobuf:"varint,2,opt,name=allow_all,json=allowAll,proto3,oneof"`
}

type PolicyRule_DenyAll struct {
	// Setting this to true means that all values are denied. This field can
	// be set only in policies for list constraints.
	DenyAll bool `protobuf:"varint,3,opt,name=deny_all,json=denyAll,proto3,oneof"`
}

type PolicyRule_Enforce struct {
	// If `true`, then the policy is enforced. If `false`, then any
	// configuration is acceptable.
	// This field can be set only in policies for boolean constraints.
	Enforce bool `protobuf:"varint,4,opt,name=enforce,proto3,oneof"`
}

func (*PolicyRule_Values) isPolicyRule_Kind() {}

func (*PolicyRule_AllowAll) isPolicyRule_Kind() {}

func (*PolicyRule_DenyAll) isPolicyRule_Kind() {}

func (*PolicyRule_Enforce) isPolicyRule_Kind() {}

// A custom constraint defined by customers which can *only* be applied to the
// given resource types and organization.
//
// By creating a custom constraint, customers can apply policies of this
// custom constraint. *Creating a custom constraint itself does NOT apply any
// policy enforcement*.
type CustomConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Name of the constraint. This is unique within the organization.
	// Format of the name should be
	//
	// -
	// `organizations/{organization_id}/customConstraints/{custom_constraint_id}`
	//
	// Example: `organizations/123/customConstraints/custom.createOnlyE2TypeVms`
	//
	// The max length is 70 characters and the minimum length is 1. Note that the
	// prefix `organizations/{organization_id}/customConstraints/` is not counted.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Immutable. The resource instance type on which this policy applies. Format
	// will be of the form : `<canonical service name>/<type>` Example:
	//
	//   - `compute.googleapis.com/Instance`.
	ResourceTypes []string `protobuf:"bytes,2,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	// All the operations being applied for this constraint.
	MethodTypes []CustomConstraint_MethodType `protobuf:"varint,3,rep,packed,name=method_types,json=methodTypes,proto3,enum=google.cloud.securityposture.v1.CustomConstraint_MethodType" json:"method_types,omitempty"`
	// Org policy condition/expression. For example:
	// `resource.instanceName.matches("[production|test]_.*_(\d)+")` or,
	// `resource.management.auto_upgrade == true`
	//
	// The max length of the condition is 1000 characters.
	Condition string `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	// Allow or deny type.
	ActionType CustomConstraint_ActionType `protobuf:"varint,5,opt,name=action_type,json=actionType,proto3,enum=google.cloud.securityposture.v1.CustomConstraint_ActionType" json:"action_type,omitempty"`
	// One line display name for the UI.
	// The max length of the display_name is 200 characters.
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Detailed information about this custom policy constraint.
	// The max length of the description is 2000 characters.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The last time this custom constraint was updated. This
	// represents the last time that the `CreateCustomConstraint` or
	// `UpdateCustomConstraint` RPC was called
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CustomConstraint) Reset() {
	*x = CustomConstraint{}
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConstraint) ProtoMessage() {}

func (x *CustomConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConstraint.ProtoReflect.Descriptor instead.
func (*CustomConstraint) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP(), []int{1}
}

func (x *CustomConstraint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomConstraint) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

func (x *CustomConstraint) GetMethodTypes() []CustomConstraint_MethodType {
	if x != nil {
		return x.MethodTypes
	}
	return nil
}

func (x *CustomConstraint) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *CustomConstraint) GetActionType() CustomConstraint_ActionType {
	if x != nil {
		return x.ActionType
	}
	return CustomConstraint_ACTION_TYPE_UNSPECIFIED
}

func (x *CustomConstraint) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CustomConstraint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomConstraint) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// A message that holds specific allowed and denied values.
// This message can define specific values and subtrees of the Resource
// Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that
// are allowed or denied. This is achieved by using the `under:` and
// optional `is:` prefixes.
// The `under:` prefix is used to denote resource subtree values.
// The `is:` prefix is used to denote specific values, and is required only
// if the value contains a ":". Values prefixed with "is:" are treated the
// same as values with no prefix.
// Ancestry subtrees must be in one of the following formats:
//
// - `projects/<project-id>` (for example, `projects/tokyo-rain-123`)
// - `folders/<folder-id>` (for example, `folders/1234`)
// - `organizations/<organization-id>` (for example, `organizations/1234`)
//
// The `supports_under` field of the associated `Constraint`  defines
// whether ancestry prefixes can be used.
type PolicyRule_StringValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of values allowed at this resource.
	AllowedValues []string `protobuf:"bytes,1,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	// List of values denied at this resource.
	DeniedValues []string `protobuf:"bytes,2,rep,name=denied_values,json=deniedValues,proto3" json:"denied_values,omitempty"`
}

func (x *PolicyRule_StringValues) Reset() {
	*x = PolicyRule_StringValues{}
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PolicyRule_StringValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRule_StringValues) ProtoMessage() {}

func (x *PolicyRule_StringValues) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRule_StringValues.ProtoReflect.Descriptor instead.
func (*PolicyRule_StringValues) Descriptor() ([]byte, []int) {
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PolicyRule_StringValues) GetAllowedValues() []string {
	if x != nil {
		return x.AllowedValues
	}
	return nil
}

func (x *PolicyRule_StringValues) GetDeniedValues() []string {
	if x != nil {
		return x.DeniedValues
	}
	return nil
}

var File_google_cloud_securityposture_v1_org_policy_config_proto protoreflect.FileDescriptor

var file_google_cloud_securityposture_v1_org_policy_config_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x75, 0x6c, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x08, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6e, 0x79,
	0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12,
	0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x5a, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6e, 0x69, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x22, 0xcb, 0x04, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x5f,
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x03, 0x22, 0x3e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59,
	0x10, 0x02, 0x42, 0x8c, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x4f, 0x72, 0x67, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x70, 0x62, 0x3b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescOnce sync.Once
	file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescData = file_google_cloud_securityposture_v1_org_policy_config_proto_rawDesc
)

func file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescGZIP() []byte {
	file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescData)
	})
	return file_google_cloud_securityposture_v1_org_policy_config_proto_rawDescData
}

var file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_securityposture_v1_org_policy_config_proto_goTypes = []any{
	(CustomConstraint_MethodType)(0), // 0: google.cloud.securityposture.v1.CustomConstraint.MethodType
	(CustomConstraint_ActionType)(0), // 1: google.cloud.securityposture.v1.CustomConstraint.ActionType
	(*PolicyRule)(nil),               // 2: google.cloud.securityposture.v1.PolicyRule
	(*CustomConstraint)(nil),         // 3: google.cloud.securityposture.v1.CustomConstraint
	(*PolicyRule_StringValues)(nil),  // 4: google.cloud.securityposture.v1.PolicyRule.StringValues
	(*expr.Expr)(nil),                // 5: google.type.Expr
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_google_cloud_securityposture_v1_org_policy_config_proto_depIdxs = []int32{
	4, // 0: google.cloud.securityposture.v1.PolicyRule.values:type_name -> google.cloud.securityposture.v1.PolicyRule.StringValues
	5, // 1: google.cloud.securityposture.v1.PolicyRule.condition:type_name -> google.type.Expr
	0, // 2: google.cloud.securityposture.v1.CustomConstraint.method_types:type_name -> google.cloud.securityposture.v1.CustomConstraint.MethodType
	1, // 3: google.cloud.securityposture.v1.CustomConstraint.action_type:type_name -> google.cloud.securityposture.v1.CustomConstraint.ActionType
	6, // 4: google.cloud.securityposture.v1.CustomConstraint.update_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_securityposture_v1_org_policy_config_proto_init() }
func file_google_cloud_securityposture_v1_org_policy_config_proto_init() {
	if File_google_cloud_securityposture_v1_org_policy_config_proto != nil {
		return
	}
	file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes[0].OneofWrappers = []any{
		(*PolicyRule_Values)(nil),
		(*PolicyRule_AllowAll)(nil),
		(*PolicyRule_DenyAll)(nil),
		(*PolicyRule_Enforce)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securityposture_v1_org_policy_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securityposture_v1_org_policy_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_securityposture_v1_org_policy_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_securityposture_v1_org_policy_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_securityposture_v1_org_policy_config_proto_msgTypes,
	}.Build()
	File_google_cloud_securityposture_v1_org_policy_config_proto = out.File
	file_google_cloud_securityposture_v1_org_policy_config_proto_rawDesc = nil
	file_google_cloud_securityposture_v1_org_policy_config_proto_goTypes = nil
	file_google_cloud_securityposture_v1_org_policy_config_proto_depIdxs = nil
}
