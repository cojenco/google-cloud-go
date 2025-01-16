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
// source: google/cloud/translate/v3/common.proto

package translatepb

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

// Possible states of long running operations.
type OperationState int32

const (
	// Invalid.
	OperationState_OPERATION_STATE_UNSPECIFIED OperationState = 0
	// Request is being processed.
	OperationState_OPERATION_STATE_RUNNING OperationState = 1
	// The operation was successful.
	OperationState_OPERATION_STATE_SUCCEEDED OperationState = 2
	// Failed to process operation.
	OperationState_OPERATION_STATE_FAILED OperationState = 3
	// Request is in the process of being canceled after caller invoked
	// longrunning.Operations.CancelOperation on the request id.
	OperationState_OPERATION_STATE_CANCELLING OperationState = 4
	// The operation request was successfully canceled.
	OperationState_OPERATION_STATE_CANCELLED OperationState = 5
)

// Enum value maps for OperationState.
var (
	OperationState_name = map[int32]string{
		0: "OPERATION_STATE_UNSPECIFIED",
		1: "OPERATION_STATE_RUNNING",
		2: "OPERATION_STATE_SUCCEEDED",
		3: "OPERATION_STATE_FAILED",
		4: "OPERATION_STATE_CANCELLING",
		5: "OPERATION_STATE_CANCELLED",
	}
	OperationState_value = map[string]int32{
		"OPERATION_STATE_UNSPECIFIED": 0,
		"OPERATION_STATE_RUNNING":     1,
		"OPERATION_STATE_SUCCEEDED":   2,
		"OPERATION_STATE_FAILED":      3,
		"OPERATION_STATE_CANCELLING":  4,
		"OPERATION_STATE_CANCELLED":   5,
	}
)

func (x OperationState) Enum() *OperationState {
	p := new(OperationState)
	*p = x
	return p
}

func (x OperationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_translate_v3_common_proto_enumTypes[0].Descriptor()
}

func (OperationState) Type() protoreflect.EnumType {
	return &file_google_cloud_translate_v3_common_proto_enumTypes[0]
}

func (x OperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationState.Descriptor instead.
func (OperationState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{0}
}

// The Google Cloud Storage location for the input content.
type GcsInputSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Source data URI. For example, `gs://my_bucket/my_object`.
	InputUri string `protobuf:"bytes,1,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
}

func (x *GcsInputSource) Reset() {
	*x = GcsInputSource{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcsInputSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsInputSource) ProtoMessage() {}

func (x *GcsInputSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsInputSource.ProtoReflect.Descriptor instead.
func (*GcsInputSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{0}
}

func (x *GcsInputSource) GetInputUri() string {
	if x != nil {
		return x.InputUri
	}
	return ""
}

// An inlined file.
type FileInputSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The file's mime type.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Required. The file's byte contents.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Required. The file's display name.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *FileInputSource) Reset() {
	*x = FileInputSource{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInputSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInputSource) ProtoMessage() {}

func (x *FileInputSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInputSource.ProtoReflect.Descriptor instead.
func (*FileInputSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{1}
}

func (x *FileInputSource) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *FileInputSource) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *FileInputSource) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// The Google Cloud Storage location for the output content.
type GcsOutputDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage URI to output directory. For example,
	// `gs://bucket/directory`. The requesting user must have write permission to
	// the bucket. The directory will be created if it doesn't exist.
	OutputUriPrefix string `protobuf:"bytes,1,opt,name=output_uri_prefix,json=outputUriPrefix,proto3" json:"output_uri_prefix,omitempty"`
}

func (x *GcsOutputDestination) Reset() {
	*x = GcsOutputDestination{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcsOutputDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsOutputDestination) ProtoMessage() {}

func (x *GcsOutputDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsOutputDestination.ProtoReflect.Descriptor instead.
func (*GcsOutputDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{2}
}

func (x *GcsOutputDestination) GetOutputUriPrefix() string {
	if x != nil {
		return x.OutputUriPrefix
	}
	return ""
}

// Represents a single entry in a glossary.
type GlossaryEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the entry.
	// Format:
	//
	//	`projects/*/locations/*/glossaries/*/glossaryEntries/*`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The different data for the glossary types (Unidirectional, Equivalent term
	// sets).
	//
	// Types that are assignable to Data:
	//
	//	*GlossaryEntry_TermsPair
	//	*GlossaryEntry_TermsSet
	Data isGlossaryEntry_Data `protobuf_oneof:"data"`
	// Describes the glossary entry.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GlossaryEntry) Reset() {
	*x = GlossaryEntry{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlossaryEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlossaryEntry) ProtoMessage() {}

func (x *GlossaryEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlossaryEntry.ProtoReflect.Descriptor instead.
func (*GlossaryEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{3}
}

func (x *GlossaryEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *GlossaryEntry) GetData() isGlossaryEntry_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *GlossaryEntry) GetTermsPair() *GlossaryEntry_GlossaryTermsPair {
	if x, ok := x.GetData().(*GlossaryEntry_TermsPair); ok {
		return x.TermsPair
	}
	return nil
}

func (x *GlossaryEntry) GetTermsSet() *GlossaryEntry_GlossaryTermsSet {
	if x, ok := x.GetData().(*GlossaryEntry_TermsSet); ok {
		return x.TermsSet
	}
	return nil
}

func (x *GlossaryEntry) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type isGlossaryEntry_Data interface {
	isGlossaryEntry_Data()
}

type GlossaryEntry_TermsPair struct {
	// Used for an unidirectional glossary.
	TermsPair *GlossaryEntry_GlossaryTermsPair `protobuf:"bytes,2,opt,name=terms_pair,json=termsPair,proto3,oneof"`
}

type GlossaryEntry_TermsSet struct {
	// Used for an equivalent term sets glossary.
	TermsSet *GlossaryEntry_GlossaryTermsSet `protobuf:"bytes,3,opt,name=terms_set,json=termsSet,proto3,oneof"`
}

func (*GlossaryEntry_TermsPair) isGlossaryEntry_Data() {}

func (*GlossaryEntry_TermsSet) isGlossaryEntry_Data() {}

// Represents a single glossary term
type GlossaryTerm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The language for this glossary term.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// The text for the glossary term.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GlossaryTerm) Reset() {
	*x = GlossaryTerm{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlossaryTerm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlossaryTerm) ProtoMessage() {}

func (x *GlossaryTerm) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlossaryTerm.ProtoReflect.Descriptor instead.
func (*GlossaryTerm) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{4}
}

func (x *GlossaryTerm) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *GlossaryTerm) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Represents a single entry for an unidirectional glossary.
type GlossaryEntry_GlossaryTermsPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source term is the term that will get match in the text,
	SourceTerm *GlossaryTerm `protobuf:"bytes,1,opt,name=source_term,json=sourceTerm,proto3" json:"source_term,omitempty"`
	// The term that will replace the match source term.
	TargetTerm *GlossaryTerm `protobuf:"bytes,2,opt,name=target_term,json=targetTerm,proto3" json:"target_term,omitempty"`
}

func (x *GlossaryEntry_GlossaryTermsPair) Reset() {
	*x = GlossaryEntry_GlossaryTermsPair{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlossaryEntry_GlossaryTermsPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlossaryEntry_GlossaryTermsPair) ProtoMessage() {}

func (x *GlossaryEntry_GlossaryTermsPair) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlossaryEntry_GlossaryTermsPair.ProtoReflect.Descriptor instead.
func (*GlossaryEntry_GlossaryTermsPair) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GlossaryEntry_GlossaryTermsPair) GetSourceTerm() *GlossaryTerm {
	if x != nil {
		return x.SourceTerm
	}
	return nil
}

func (x *GlossaryEntry_GlossaryTermsPair) GetTargetTerm() *GlossaryTerm {
	if x != nil {
		return x.TargetTerm
	}
	return nil
}

// Represents a single entry for an equivalent term set glossary. This is used
// for equivalent term sets where each term can be replaced by the other terms
// in the set.
type GlossaryEntry_GlossaryTermsSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each term in the set represents a term that can be replaced by the other
	// terms.
	Terms []*GlossaryTerm `protobuf:"bytes,1,rep,name=terms,proto3" json:"terms,omitempty"`
}

func (x *GlossaryEntry_GlossaryTermsSet) Reset() {
	*x = GlossaryEntry_GlossaryTermsSet{}
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlossaryEntry_GlossaryTermsSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlossaryEntry_GlossaryTermsSet) ProtoMessage() {}

func (x *GlossaryEntry_GlossaryTermsSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_translate_v3_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlossaryEntry_GlossaryTermsSet.ProtoReflect.Descriptor instead.
func (*GlossaryEntry_GlossaryTermsSet) Descriptor() ([]byte, []int) {
	return file_google_cloud_translate_v3_common_proto_rawDescGZIP(), []int{3, 1}
}

func (x *GlossaryEntry_GlossaryTermsSet) GetTerms() []*GlossaryTerm {
	if x != nil {
		return x.Terms
	}
	return nil
}

var File_google_cloud_translate_v3_common_proto protoreflect.FileDescriptor

var file_google_cloud_translate_v3_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x63, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x55, 0x72, 0x69, 0x22, 0x7a, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x63, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x11, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x55, 0x72, 0x69, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xbf, 0x05, 0x0a, 0x0d, 0x47,
	0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x70,
	0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54, 0x65,
	0x72, 0x6d, 0x73, 0x50, 0x61, 0x69, 0x72, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x73,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x5a, 0x0a, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x53, 0x65, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0xab, 0x01, 0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54,
	0x65, 0x72, 0x6d, 0x73, 0x50, 0x61, 0x69, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6c, 0x6f, 0x73,
	0x73, 0x61, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x65, 0x72, 0x6d, 0x12, 0x4a, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79,
	0x54, 0x65, 0x72, 0x6d, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x1a, 0x53, 0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x53, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x05,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x3a, 0xac, 0x01, 0xea, 0x41, 0xa8, 0x01, 0x0a, 0x26, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x5e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x67, 0x6c,
	0x6f, 0x73, 0x73, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x6c, 0x6f, 0x73, 0x73, 0x61,
	0x72, 0x79, 0x7d, 0x2f, 0x67, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x7d, 0x2a, 0x0f, 0x67, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0x0d, 0x67, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x0c,
	0x47, 0x6c, 0x6f, 0x73, 0x73, 0x61, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x2a, 0xc8, 0x01, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05,
	0x42, 0xc5, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x70, 0x62, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x62, 0xf8, 0x01,
	0x01, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x19,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x33, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_translate_v3_common_proto_rawDescOnce sync.Once
	file_google_cloud_translate_v3_common_proto_rawDescData = file_google_cloud_translate_v3_common_proto_rawDesc
)

func file_google_cloud_translate_v3_common_proto_rawDescGZIP() []byte {
	file_google_cloud_translate_v3_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_translate_v3_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_translate_v3_common_proto_rawDescData)
	})
	return file_google_cloud_translate_v3_common_proto_rawDescData
}

var file_google_cloud_translate_v3_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_translate_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_translate_v3_common_proto_goTypes = []any{
	(OperationState)(0),                     // 0: google.cloud.translation.v3.OperationState
	(*GcsInputSource)(nil),                  // 1: google.cloud.translation.v3.GcsInputSource
	(*FileInputSource)(nil),                 // 2: google.cloud.translation.v3.FileInputSource
	(*GcsOutputDestination)(nil),            // 3: google.cloud.translation.v3.GcsOutputDestination
	(*GlossaryEntry)(nil),                   // 4: google.cloud.translation.v3.GlossaryEntry
	(*GlossaryTerm)(nil),                    // 5: google.cloud.translation.v3.GlossaryTerm
	(*GlossaryEntry_GlossaryTermsPair)(nil), // 6: google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair
	(*GlossaryEntry_GlossaryTermsSet)(nil),  // 7: google.cloud.translation.v3.GlossaryEntry.GlossaryTermsSet
}
var file_google_cloud_translate_v3_common_proto_depIdxs = []int32{
	6, // 0: google.cloud.translation.v3.GlossaryEntry.terms_pair:type_name -> google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair
	7, // 1: google.cloud.translation.v3.GlossaryEntry.terms_set:type_name -> google.cloud.translation.v3.GlossaryEntry.GlossaryTermsSet
	5, // 2: google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair.source_term:type_name -> google.cloud.translation.v3.GlossaryTerm
	5, // 3: google.cloud.translation.v3.GlossaryEntry.GlossaryTermsPair.target_term:type_name -> google.cloud.translation.v3.GlossaryTerm
	5, // 4: google.cloud.translation.v3.GlossaryEntry.GlossaryTermsSet.terms:type_name -> google.cloud.translation.v3.GlossaryTerm
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_translate_v3_common_proto_init() }
func file_google_cloud_translate_v3_common_proto_init() {
	if File_google_cloud_translate_v3_common_proto != nil {
		return
	}
	file_google_cloud_translate_v3_common_proto_msgTypes[3].OneofWrappers = []any{
		(*GlossaryEntry_TermsPair)(nil),
		(*GlossaryEntry_TermsSet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_translate_v3_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_translate_v3_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_translate_v3_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_translate_v3_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_translate_v3_common_proto_msgTypes,
	}.Build()
	File_google_cloud_translate_v3_common_proto = out.File
	file_google_cloud_translate_v3_common_proto_rawDesc = nil
	file_google_cloud_translate_v3_common_proto_goTypes = nil
	file_google_cloud_translate_v3_common_proto_depIdxs = nil
}
