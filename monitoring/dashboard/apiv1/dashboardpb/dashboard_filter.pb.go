// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/monitoring/dashboard/v1/dashboard_filter.proto

package dashboardpb

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

// The type for the dashboard filter
type DashboardFilter_FilterType int32

const (
	// Filter type is unspecified. This is not valid in a well-formed request.
	DashboardFilter_FILTER_TYPE_UNSPECIFIED DashboardFilter_FilterType = 0
	// Filter on a resource label value
	DashboardFilter_RESOURCE_LABEL DashboardFilter_FilterType = 1
	// Filter on a metrics label value
	DashboardFilter_METRIC_LABEL DashboardFilter_FilterType = 2
	// Filter on a user metadata label value
	DashboardFilter_USER_METADATA_LABEL DashboardFilter_FilterType = 3
	// Filter on a system metadata label value
	DashboardFilter_SYSTEM_METADATA_LABEL DashboardFilter_FilterType = 4
	// Filter on a group id
	DashboardFilter_GROUP DashboardFilter_FilterType = 5
)

// Enum value maps for DashboardFilter_FilterType.
var (
	DashboardFilter_FilterType_name = map[int32]string{
		0: "FILTER_TYPE_UNSPECIFIED",
		1: "RESOURCE_LABEL",
		2: "METRIC_LABEL",
		3: "USER_METADATA_LABEL",
		4: "SYSTEM_METADATA_LABEL",
		5: "GROUP",
	}
	DashboardFilter_FilterType_value = map[string]int32{
		"FILTER_TYPE_UNSPECIFIED": 0,
		"RESOURCE_LABEL":          1,
		"METRIC_LABEL":            2,
		"USER_METADATA_LABEL":     3,
		"SYSTEM_METADATA_LABEL":   4,
		"GROUP":                   5,
	}
)

func (x DashboardFilter_FilterType) Enum() *DashboardFilter_FilterType {
	p := new(DashboardFilter_FilterType)
	*p = x
	return p
}

func (x DashboardFilter_FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DashboardFilter_FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_monitoring_dashboard_v1_dashboard_filter_proto_enumTypes[0].Descriptor()
}

func (DashboardFilter_FilterType) Type() protoreflect.EnumType {
	return &file_google_monitoring_dashboard_v1_dashboard_filter_proto_enumTypes[0]
}

func (x DashboardFilter_FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DashboardFilter_FilterType.Descriptor instead.
func (DashboardFilter_FilterType) EnumDescriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescGZIP(), []int{0, 0}
}

// A filter to reduce the amount of data charted in relevant widgets.
type DashboardFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The key for the label
	LabelKey string `protobuf:"bytes,1,opt,name=label_key,json=labelKey,proto3" json:"label_key,omitempty"`
	// The placeholder text that can be referenced in a filter string or MQL
	// query. If omitted, the dashboard filter will be applied to all relevant
	// widgets in the dashboard.
	TemplateVariable string `protobuf:"bytes,3,opt,name=template_variable,json=templateVariable,proto3" json:"template_variable,omitempty"`
	// The default value used in the filter comparison
	//
	// Types that are assignable to DefaultValue:
	//
	//	*DashboardFilter_StringValue
	DefaultValue isDashboardFilter_DefaultValue `protobuf_oneof:"default_value"`
	// The specified filter type
	FilterType DashboardFilter_FilterType `protobuf:"varint,5,opt,name=filter_type,json=filterType,proto3,enum=google.monitoring.dashboard.v1.DashboardFilter_FilterType" json:"filter_type,omitempty"`
}

func (x *DashboardFilter) Reset() {
	*x = DashboardFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardFilter) ProtoMessage() {}

func (x *DashboardFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardFilter.ProtoReflect.Descriptor instead.
func (*DashboardFilter) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescGZIP(), []int{0}
}

func (x *DashboardFilter) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *DashboardFilter) GetTemplateVariable() string {
	if x != nil {
		return x.TemplateVariable
	}
	return ""
}

func (m *DashboardFilter) GetDefaultValue() isDashboardFilter_DefaultValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (x *DashboardFilter) GetStringValue() string {
	if x, ok := x.GetDefaultValue().(*DashboardFilter_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *DashboardFilter) GetFilterType() DashboardFilter_FilterType {
	if x != nil {
		return x.FilterType
	}
	return DashboardFilter_FILTER_TYPE_UNSPECIFIED
}

type isDashboardFilter_DefaultValue interface {
	isDashboardFilter_DefaultValue()
}

type DashboardFilter_StringValue struct {
	// A variable-length string value.
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*DashboardFilter_StringValue) isDashboardFilter_DefaultValue() {}

var File_google_monitoring_dashboard_v1_dashboard_filter_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x0f, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x2b,
	0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x5b, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8e, 0x01,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4c, 0x41, 0x42, 0x45,
	0x4c, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x05, 0x42, 0x0f,
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0xfe, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescData = file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_dashboard_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_monitoring_dashboard_v1_dashboard_filter_proto_goTypes = []interface{}{
	(DashboardFilter_FilterType)(0), // 0: google.monitoring.dashboard.v1.DashboardFilter.FilterType
	(*DashboardFilter)(nil),         // 1: google.monitoring.dashboard.v1.DashboardFilter
}
var file_google_monitoring_dashboard_v1_dashboard_filter_proto_depIdxs = []int32{
	0, // 0: google.monitoring.dashboard.v1.DashboardFilter.filter_type:type_name -> google.monitoring.dashboard.v1.DashboardFilter.FilterType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_dashboard_filter_proto_init() }
func file_google_monitoring_dashboard_v1_dashboard_filter_proto_init() {
	if File_google_monitoring_dashboard_v1_dashboard_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardFilter); i {
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
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DashboardFilter_StringValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_dashboard_filter_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_dashboard_filter_proto_depIdxs,
		EnumInfos:         file_google_monitoring_dashboard_v1_dashboard_filter_proto_enumTypes,
		MessageInfos:      file_google_monitoring_dashboard_v1_dashboard_filter_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_dashboard_filter_proto = out.File
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_dashboard_filter_proto_depIdxs = nil
}
