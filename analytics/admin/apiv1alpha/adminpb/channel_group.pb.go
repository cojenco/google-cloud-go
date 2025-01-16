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
// source: google/analytics/admin/v1alpha/channel_group.proto

package adminpb

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

// How the filter will be used to determine a match.
type ChannelGroupFilter_StringFilter_MatchType int32

const (
	// Default match type.
	ChannelGroupFilter_StringFilter_MATCH_TYPE_UNSPECIFIED ChannelGroupFilter_StringFilter_MatchType = 0
	// Exact match of the string value.
	ChannelGroupFilter_StringFilter_EXACT ChannelGroupFilter_StringFilter_MatchType = 1
	// Begins with the string value.
	ChannelGroupFilter_StringFilter_BEGINS_WITH ChannelGroupFilter_StringFilter_MatchType = 2
	// Ends with the string value.
	ChannelGroupFilter_StringFilter_ENDS_WITH ChannelGroupFilter_StringFilter_MatchType = 3
	// Contains the string value.
	ChannelGroupFilter_StringFilter_CONTAINS ChannelGroupFilter_StringFilter_MatchType = 4
	// Full regular expression match with the string value.
	ChannelGroupFilter_StringFilter_FULL_REGEXP ChannelGroupFilter_StringFilter_MatchType = 5
	// Partial regular expression match with the string value.
	ChannelGroupFilter_StringFilter_PARTIAL_REGEXP ChannelGroupFilter_StringFilter_MatchType = 6
)

// Enum value maps for ChannelGroupFilter_StringFilter_MatchType.
var (
	ChannelGroupFilter_StringFilter_MatchType_name = map[int32]string{
		0: "MATCH_TYPE_UNSPECIFIED",
		1: "EXACT",
		2: "BEGINS_WITH",
		3: "ENDS_WITH",
		4: "CONTAINS",
		5: "FULL_REGEXP",
		6: "PARTIAL_REGEXP",
	}
	ChannelGroupFilter_StringFilter_MatchType_value = map[string]int32{
		"MATCH_TYPE_UNSPECIFIED": 0,
		"EXACT":                  1,
		"BEGINS_WITH":            2,
		"ENDS_WITH":              3,
		"CONTAINS":               4,
		"FULL_REGEXP":            5,
		"PARTIAL_REGEXP":         6,
	}
)

func (x ChannelGroupFilter_StringFilter_MatchType) Enum() *ChannelGroupFilter_StringFilter_MatchType {
	p := new(ChannelGroupFilter_StringFilter_MatchType)
	*p = x
	return p
}

func (x ChannelGroupFilter_StringFilter_MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelGroupFilter_StringFilter_MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_analytics_admin_v1alpha_channel_group_proto_enumTypes[0].Descriptor()
}

func (ChannelGroupFilter_StringFilter_MatchType) Type() protoreflect.EnumType {
	return &file_google_analytics_admin_v1alpha_channel_group_proto_enumTypes[0]
}

func (x ChannelGroupFilter_StringFilter_MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelGroupFilter_StringFilter_MatchType.Descriptor instead.
func (ChannelGroupFilter_StringFilter_MatchType) EnumDescriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A specific filter for a single dimension.
type ChannelGroupFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A StringFilter or InListFilter that defines this filters behavior.
	//
	// Types that are assignable to ValueFilter:
	//
	//	*ChannelGroupFilter_StringFilter_
	//	*ChannelGroupFilter_InListFilter_
	ValueFilter isChannelGroupFilter_ValueFilter `protobuf_oneof:"value_filter"`
	// Required. Immutable. The dimension name to filter.
	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
}

func (x *ChannelGroupFilter) Reset() {
	*x = ChannelGroupFilter{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroupFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroupFilter) ProtoMessage() {}

func (x *ChannelGroupFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroupFilter.ProtoReflect.Descriptor instead.
func (*ChannelGroupFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{0}
}

func (m *ChannelGroupFilter) GetValueFilter() isChannelGroupFilter_ValueFilter {
	if m != nil {
		return m.ValueFilter
	}
	return nil
}

func (x *ChannelGroupFilter) GetStringFilter() *ChannelGroupFilter_StringFilter {
	if x, ok := x.GetValueFilter().(*ChannelGroupFilter_StringFilter_); ok {
		return x.StringFilter
	}
	return nil
}

func (x *ChannelGroupFilter) GetInListFilter() *ChannelGroupFilter_InListFilter {
	if x, ok := x.GetValueFilter().(*ChannelGroupFilter_InListFilter_); ok {
		return x.InListFilter
	}
	return nil
}

func (x *ChannelGroupFilter) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

type isChannelGroupFilter_ValueFilter interface {
	isChannelGroupFilter_ValueFilter()
}

type ChannelGroupFilter_StringFilter_ struct {
	// A filter for a string-type dimension that matches a particular pattern.
	StringFilter *ChannelGroupFilter_StringFilter `protobuf:"bytes,2,opt,name=string_filter,json=stringFilter,proto3,oneof"`
}

type ChannelGroupFilter_InListFilter_ struct {
	// A filter for a string dimension that matches a particular list of
	// options.
	InListFilter *ChannelGroupFilter_InListFilter `protobuf:"bytes,3,opt,name=in_list_filter,json=inListFilter,proto3,oneof"`
}

func (*ChannelGroupFilter_StringFilter_) isChannelGroupFilter_ValueFilter() {}

func (*ChannelGroupFilter_InListFilter_) isChannelGroupFilter_ValueFilter() {}

// A logical expression of Channel Group dimension filters.
type ChannelGroupFilterExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The expression applied to a filter.
	//
	// Types that are assignable to Expr:
	//
	//	*ChannelGroupFilterExpression_AndGroup
	//	*ChannelGroupFilterExpression_OrGroup
	//	*ChannelGroupFilterExpression_NotExpression
	//	*ChannelGroupFilterExpression_Filter
	Expr isChannelGroupFilterExpression_Expr `protobuf_oneof:"expr"`
}

func (x *ChannelGroupFilterExpression) Reset() {
	*x = ChannelGroupFilterExpression{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroupFilterExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroupFilterExpression) ProtoMessage() {}

func (x *ChannelGroupFilterExpression) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroupFilterExpression.ProtoReflect.Descriptor instead.
func (*ChannelGroupFilterExpression) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{1}
}

func (m *ChannelGroupFilterExpression) GetExpr() isChannelGroupFilterExpression_Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (x *ChannelGroupFilterExpression) GetAndGroup() *ChannelGroupFilterExpressionList {
	if x, ok := x.GetExpr().(*ChannelGroupFilterExpression_AndGroup); ok {
		return x.AndGroup
	}
	return nil
}

func (x *ChannelGroupFilterExpression) GetOrGroup() *ChannelGroupFilterExpressionList {
	if x, ok := x.GetExpr().(*ChannelGroupFilterExpression_OrGroup); ok {
		return x.OrGroup
	}
	return nil
}

func (x *ChannelGroupFilterExpression) GetNotExpression() *ChannelGroupFilterExpression {
	if x, ok := x.GetExpr().(*ChannelGroupFilterExpression_NotExpression); ok {
		return x.NotExpression
	}
	return nil
}

func (x *ChannelGroupFilterExpression) GetFilter() *ChannelGroupFilter {
	if x, ok := x.GetExpr().(*ChannelGroupFilterExpression_Filter); ok {
		return x.Filter
	}
	return nil
}

type isChannelGroupFilterExpression_Expr interface {
	isChannelGroupFilterExpression_Expr()
}

type ChannelGroupFilterExpression_AndGroup struct {
	// A list of expressions to be AND’ed together. It can only contain
	// ChannelGroupFilterExpressions with or_group. This must be set for the
	// top level ChannelGroupFilterExpression.
	AndGroup *ChannelGroupFilterExpressionList `protobuf:"bytes,1,opt,name=and_group,json=andGroup,proto3,oneof"`
}

type ChannelGroupFilterExpression_OrGroup struct {
	// A list of expressions to OR’ed together. It cannot contain
	// ChannelGroupFilterExpressions with and_group or or_group.
	OrGroup *ChannelGroupFilterExpressionList `protobuf:"bytes,2,opt,name=or_group,json=orGroup,proto3,oneof"`
}

type ChannelGroupFilterExpression_NotExpression struct {
	// A filter expression to be NOT'ed (that is inverted, complemented). It
	// can only include a dimension_or_metric_filter. This cannot be set on the
	// top level ChannelGroupFilterExpression.
	NotExpression *ChannelGroupFilterExpression `protobuf:"bytes,3,opt,name=not_expression,json=notExpression,proto3,oneof"`
}

type ChannelGroupFilterExpression_Filter struct {
	// A filter on a single dimension. This cannot be set on the top
	// level ChannelGroupFilterExpression.
	Filter *ChannelGroupFilter `protobuf:"bytes,4,opt,name=filter,proto3,oneof"`
}

func (*ChannelGroupFilterExpression_AndGroup) isChannelGroupFilterExpression_Expr() {}

func (*ChannelGroupFilterExpression_OrGroup) isChannelGroupFilterExpression_Expr() {}

func (*ChannelGroupFilterExpression_NotExpression) isChannelGroupFilterExpression_Expr() {}

func (*ChannelGroupFilterExpression_Filter) isChannelGroupFilterExpression_Expr() {}

// A list of Channel Group filter expressions.
type ChannelGroupFilterExpressionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Channel Group filter expressions.
	FilterExpressions []*ChannelGroupFilterExpression `protobuf:"bytes,1,rep,name=filter_expressions,json=filterExpressions,proto3" json:"filter_expressions,omitempty"`
}

func (x *ChannelGroupFilterExpressionList) Reset() {
	*x = ChannelGroupFilterExpressionList{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroupFilterExpressionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroupFilterExpressionList) ProtoMessage() {}

func (x *ChannelGroupFilterExpressionList) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroupFilterExpressionList.ProtoReflect.Descriptor instead.
func (*ChannelGroupFilterExpressionList) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelGroupFilterExpressionList) GetFilterExpressions() []*ChannelGroupFilterExpression {
	if x != nil {
		return x.FilterExpressions
	}
	return nil
}

// The rules that govern how traffic is grouped into one channel.
type GroupingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Customer defined display name for the channel.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The Filter Expression that defines the Grouping Rule.
	Expression *ChannelGroupFilterExpression `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *GroupingRule) Reset() {
	*x = GroupingRule{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupingRule) ProtoMessage() {}

func (x *GroupingRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupingRule.ProtoReflect.Descriptor instead.
func (*GroupingRule) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{3}
}

func (x *GroupingRule) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GroupingRule) GetExpression() *ChannelGroupFilterExpression {
	if x != nil {
		return x.Expression
	}
	return nil
}

// A resource message representing a Channel Group.
type ChannelGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name for this Channel Group resource.
	// Format: properties/{property}/channelGroups/{channel_group}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the Channel Group. Max length of 80
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the Channel Group. Max length of 256 characters.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. The grouping rules of channels. Maximum number of rules is 50.
	GroupingRule []*GroupingRule `protobuf:"bytes,4,rep,name=grouping_rule,json=groupingRule,proto3" json:"grouping_rule,omitempty"`
	// Output only. If true, then this channel group is the Default Channel Group
	// predefined by Google Analytics. Display name and grouping rules cannot be
	// updated for this channel group.
	SystemDefined bool `protobuf:"varint,5,opt,name=system_defined,json=systemDefined,proto3" json:"system_defined,omitempty"`
	// Optional. If true, this channel group will be used as the default channel
	// group for reports. Only one channel group can be set as `primary` at any
	// time. If the `primary` field gets set on a channel group, it will get unset
	// on the previous primary channel group.
	//
	// The Google Analytics predefined channel group is the primary by default.
	Primary bool `protobuf:"varint,6,opt,name=primary,proto3" json:"primary,omitempty"`
}

func (x *ChannelGroup) Reset() {
	*x = ChannelGroup{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroup) ProtoMessage() {}

func (x *ChannelGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroup.ProtoReflect.Descriptor instead.
func (*ChannelGroup) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelGroup) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ChannelGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChannelGroup) GetGroupingRule() []*GroupingRule {
	if x != nil {
		return x.GroupingRule
	}
	return nil
}

func (x *ChannelGroup) GetSystemDefined() bool {
	if x != nil {
		return x.SystemDefined
	}
	return false
}

func (x *ChannelGroup) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

// Filter where the field value is a String. The match is case insensitive.
type ChannelGroupFilter_StringFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The match type for the string filter.
	MatchType ChannelGroupFilter_StringFilter_MatchType `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=google.analytics.admin.v1alpha.ChannelGroupFilter_StringFilter_MatchType" json:"match_type,omitempty"`
	// Required. The string value to be matched against.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ChannelGroupFilter_StringFilter) Reset() {
	*x = ChannelGroupFilter_StringFilter{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroupFilter_StringFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroupFilter_StringFilter) ProtoMessage() {}

func (x *ChannelGroupFilter_StringFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroupFilter_StringFilter.ProtoReflect.Descriptor instead.
func (*ChannelGroupFilter_StringFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChannelGroupFilter_StringFilter) GetMatchType() ChannelGroupFilter_StringFilter_MatchType {
	if x != nil {
		return x.MatchType
	}
	return ChannelGroupFilter_StringFilter_MATCH_TYPE_UNSPECIFIED
}

func (x *ChannelGroupFilter_StringFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A filter for a string dimension that matches a particular list of options.
// The match is case insensitive.
type ChannelGroupFilter_InListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The list of possible string values to match against. Must be
	// non-empty.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ChannelGroupFilter_InListFilter) Reset() {
	*x = ChannelGroupFilter_InListFilter{}
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChannelGroupFilter_InListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGroupFilter_InListFilter) ProtoMessage() {}

func (x *ChannelGroupFilter_InListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGroupFilter_InListFilter.ProtoReflect.Descriptor instead.
func (*ChannelGroupFilter_InListFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ChannelGroupFilter_InListFilter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_google_analytics_admin_v1alpha_channel_group_proto protoreflect.FileDescriptor

var file_google_analytics_admin_v1alpha_channel_group_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xec, 0x04, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x67, 0x0a, 0x0e, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41,
	0x02, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0xa0, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x6d, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x44, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x04, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x50, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x50,
	0x10, 0x06, 0x1a, 0x2b, 0x0a, 0x0c, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42,
	0x0e, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x9b, 0x03, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x5f, 0x0a, 0x09, 0x61, 0x6e, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x61, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x5d, 0x0a, 0x08, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x65, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x22, 0x8f, 0x01,
	0x0a, 0x20, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x6b, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x99, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xfa, 0x02, 0x0a, 0x0c,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x56, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67,
	0x52, 0x75, 0x6c, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x3a, 0x64, 0xea, 0x41, 0x61, 0x0a, 0x2a, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x33, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x42, 0x79, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_analytics_admin_v1alpha_channel_group_proto_rawDescOnce sync.Once
	file_google_analytics_admin_v1alpha_channel_group_proto_rawDescData = file_google_analytics_admin_v1alpha_channel_group_proto_rawDesc
)

func file_google_analytics_admin_v1alpha_channel_group_proto_rawDescGZIP() []byte {
	file_google_analytics_admin_v1alpha_channel_group_proto_rawDescOnce.Do(func() {
		file_google_analytics_admin_v1alpha_channel_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_analytics_admin_v1alpha_channel_group_proto_rawDescData)
	})
	return file_google_analytics_admin_v1alpha_channel_group_proto_rawDescData
}

var file_google_analytics_admin_v1alpha_channel_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_analytics_admin_v1alpha_channel_group_proto_goTypes = []any{
	(ChannelGroupFilter_StringFilter_MatchType)(0), // 0: google.analytics.admin.v1alpha.ChannelGroupFilter.StringFilter.MatchType
	(*ChannelGroupFilter)(nil),                     // 1: google.analytics.admin.v1alpha.ChannelGroupFilter
	(*ChannelGroupFilterExpression)(nil),           // 2: google.analytics.admin.v1alpha.ChannelGroupFilterExpression
	(*ChannelGroupFilterExpressionList)(nil),       // 3: google.analytics.admin.v1alpha.ChannelGroupFilterExpressionList
	(*GroupingRule)(nil),                           // 4: google.analytics.admin.v1alpha.GroupingRule
	(*ChannelGroup)(nil),                           // 5: google.analytics.admin.v1alpha.ChannelGroup
	(*ChannelGroupFilter_StringFilter)(nil),        // 6: google.analytics.admin.v1alpha.ChannelGroupFilter.StringFilter
	(*ChannelGroupFilter_InListFilter)(nil),        // 7: google.analytics.admin.v1alpha.ChannelGroupFilter.InListFilter
}
var file_google_analytics_admin_v1alpha_channel_group_proto_depIdxs = []int32{
	6,  // 0: google.analytics.admin.v1alpha.ChannelGroupFilter.string_filter:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilter.StringFilter
	7,  // 1: google.analytics.admin.v1alpha.ChannelGroupFilter.in_list_filter:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilter.InListFilter
	3,  // 2: google.analytics.admin.v1alpha.ChannelGroupFilterExpression.and_group:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilterExpressionList
	3,  // 3: google.analytics.admin.v1alpha.ChannelGroupFilterExpression.or_group:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilterExpressionList
	2,  // 4: google.analytics.admin.v1alpha.ChannelGroupFilterExpression.not_expression:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilterExpression
	1,  // 5: google.analytics.admin.v1alpha.ChannelGroupFilterExpression.filter:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilter
	2,  // 6: google.analytics.admin.v1alpha.ChannelGroupFilterExpressionList.filter_expressions:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilterExpression
	2,  // 7: google.analytics.admin.v1alpha.GroupingRule.expression:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilterExpression
	4,  // 8: google.analytics.admin.v1alpha.ChannelGroup.grouping_rule:type_name -> google.analytics.admin.v1alpha.GroupingRule
	0,  // 9: google.analytics.admin.v1alpha.ChannelGroupFilter.StringFilter.match_type:type_name -> google.analytics.admin.v1alpha.ChannelGroupFilter.StringFilter.MatchType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_analytics_admin_v1alpha_channel_group_proto_init() }
func file_google_analytics_admin_v1alpha_channel_group_proto_init() {
	if File_google_analytics_admin_v1alpha_channel_group_proto != nil {
		return
	}
	file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[0].OneofWrappers = []any{
		(*ChannelGroupFilter_StringFilter_)(nil),
		(*ChannelGroupFilter_InListFilter_)(nil),
	}
	file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes[1].OneofWrappers = []any{
		(*ChannelGroupFilterExpression_AndGroup)(nil),
		(*ChannelGroupFilterExpression_OrGroup)(nil),
		(*ChannelGroupFilterExpression_NotExpression)(nil),
		(*ChannelGroupFilterExpression_Filter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_analytics_admin_v1alpha_channel_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_analytics_admin_v1alpha_channel_group_proto_goTypes,
		DependencyIndexes: file_google_analytics_admin_v1alpha_channel_group_proto_depIdxs,
		EnumInfos:         file_google_analytics_admin_v1alpha_channel_group_proto_enumTypes,
		MessageInfos:      file_google_analytics_admin_v1alpha_channel_group_proto_msgTypes,
	}.Build()
	File_google_analytics_admin_v1alpha_channel_group_proto = out.File
	file_google_analytics_admin_v1alpha_channel_group_proto_rawDesc = nil
	file_google_analytics_admin_v1alpha_channel_group_proto_goTypes = nil
	file_google_analytics_admin_v1alpha_channel_group_proto_depIdxs = nil
}
