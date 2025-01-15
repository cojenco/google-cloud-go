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
// source: google/cloud/eventarc/v1/logging_config.proto

package eventarcpb

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

// The different severities for logging supported by Eventarc Advanced
// resources.
// This enum is an exhaustive list of log severities and is FROZEN. Do not
// expect new values to be added.
type LoggingConfig_LogSeverity int32

const (
	// Log severity is not specified. This value is treated the same as NONE,
	// but is used to distinguish between no update and update to NONE in
	// update_masks.
	LoggingConfig_LOG_SEVERITY_UNSPECIFIED LoggingConfig_LogSeverity = 0
	// Default value at resource creation, presence of this value must be
	// treated as no logging/disable logging.
	LoggingConfig_NONE LoggingConfig_LogSeverity = 1
	// Debug or trace level logging.
	LoggingConfig_DEBUG LoggingConfig_LogSeverity = 2
	// Routine information, such as ongoing status or performance.
	LoggingConfig_INFO LoggingConfig_LogSeverity = 3
	// Normal but significant events, such as start up, shut down, or a
	// configuration change.
	LoggingConfig_NOTICE LoggingConfig_LogSeverity = 4
	// Warning events might cause problems.
	LoggingConfig_WARNING LoggingConfig_LogSeverity = 5
	// Error events are likely to cause problems.
	LoggingConfig_ERROR LoggingConfig_LogSeverity = 6
	// Critical events cause more severe problems or outages.
	LoggingConfig_CRITICAL LoggingConfig_LogSeverity = 7
	// A person must take action immediately.
	LoggingConfig_ALERT LoggingConfig_LogSeverity = 8
	// One or more systems are unusable.
	LoggingConfig_EMERGENCY LoggingConfig_LogSeverity = 9
)

// Enum value maps for LoggingConfig_LogSeverity.
var (
	LoggingConfig_LogSeverity_name = map[int32]string{
		0: "LOG_SEVERITY_UNSPECIFIED",
		1: "NONE",
		2: "DEBUG",
		3: "INFO",
		4: "NOTICE",
		5: "WARNING",
		6: "ERROR",
		7: "CRITICAL",
		8: "ALERT",
		9: "EMERGENCY",
	}
	LoggingConfig_LogSeverity_value = map[string]int32{
		"LOG_SEVERITY_UNSPECIFIED": 0,
		"NONE":                     1,
		"DEBUG":                    2,
		"INFO":                     3,
		"NOTICE":                   4,
		"WARNING":                  5,
		"ERROR":                    6,
		"CRITICAL":                 7,
		"ALERT":                    8,
		"EMERGENCY":                9,
	}
)

func (x LoggingConfig_LogSeverity) Enum() *LoggingConfig_LogSeverity {
	p := new(LoggingConfig_LogSeverity)
	*p = x
	return p
}

func (x LoggingConfig_LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggingConfig_LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_eventarc_v1_logging_config_proto_enumTypes[0].Descriptor()
}

func (LoggingConfig_LogSeverity) Type() protoreflect.EnumType {
	return &file_google_cloud_eventarc_v1_logging_config_proto_enumTypes[0]
}

func (x LoggingConfig_LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggingConfig_LogSeverity.Descriptor instead.
func (LoggingConfig_LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_v1_logging_config_proto_rawDescGZIP(), []int{0, 0}
}

// The configuration for Platform Telemetry logging for Eventarc Advanced
// resources.
type LoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The minimum severity of logs that will be sent to
	// Stackdriver/Platform Telemetry. Logs at severitiy ≥ this value will be
	// sent, unless it is NONE.
	LogSeverity LoggingConfig_LogSeverity `protobuf:"varint,1,opt,name=log_severity,json=logSeverity,proto3,enum=google.cloud.eventarc.v1.LoggingConfig_LogSeverity" json:"log_severity,omitempty"`
}

func (x *LoggingConfig) Reset() {
	*x = LoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_eventarc_v1_logging_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingConfig) ProtoMessage() {}

func (x *LoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_eventarc_v1_logging_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingConfig.ProtoReflect.Descriptor instead.
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_v1_logging_config_proto_rawDescGZIP(), []int{0}
}

func (x *LoggingConfig) GetLogSeverity() LoggingConfig_LogSeverity {
	if x != nil {
		return x.LogSeverity
	}
	return LoggingConfig_LOG_SEVERITY_UNSPECIFIED
}

var File_google_cloud_eventarc_v1_logging_config_proto protoreflect.FileDescriptor

var file_google_cloud_eventarc_v1_logging_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5b, 0x0a, 0x0c,
	0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x6c, 0x6f,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f, 0x47,
	0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52,
	0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x45, 0x52,
	0x54, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59,
	0x10, 0x09, 0x42, 0xc2, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x70, 0x62, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x63, 0x70, 0x62, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x61, 0x72, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_eventarc_v1_logging_config_proto_rawDescOnce sync.Once
	file_google_cloud_eventarc_v1_logging_config_proto_rawDescData = file_google_cloud_eventarc_v1_logging_config_proto_rawDesc
)

func file_google_cloud_eventarc_v1_logging_config_proto_rawDescGZIP() []byte {
	file_google_cloud_eventarc_v1_logging_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_eventarc_v1_logging_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_eventarc_v1_logging_config_proto_rawDescData)
	})
	return file_google_cloud_eventarc_v1_logging_config_proto_rawDescData
}

var file_google_cloud_eventarc_v1_logging_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_eventarc_v1_logging_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_eventarc_v1_logging_config_proto_goTypes = []any{
	(LoggingConfig_LogSeverity)(0), // 0: google.cloud.eventarc.v1.LoggingConfig.LogSeverity
	(*LoggingConfig)(nil),          // 1: google.cloud.eventarc.v1.LoggingConfig
}
var file_google_cloud_eventarc_v1_logging_config_proto_depIdxs = []int32{
	0, // 0: google.cloud.eventarc.v1.LoggingConfig.log_severity:type_name -> google.cloud.eventarc.v1.LoggingConfig.LogSeverity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_eventarc_v1_logging_config_proto_init() }
func file_google_cloud_eventarc_v1_logging_config_proto_init() {
	if File_google_cloud_eventarc_v1_logging_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_eventarc_v1_logging_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoggingConfig); i {
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
			RawDescriptor: file_google_cloud_eventarc_v1_logging_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_eventarc_v1_logging_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_eventarc_v1_logging_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_eventarc_v1_logging_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_eventarc_v1_logging_config_proto_msgTypes,
	}.Build()
	File_google_cloud_eventarc_v1_logging_config_proto = out.File
	file_google_cloud_eventarc_v1_logging_config_proto_rawDesc = nil
	file_google_cloud_eventarc_v1_logging_config_proto_goTypes = nil
	file_google_cloud_eventarc_v1_logging_config_proto_depIdxs = nil
}