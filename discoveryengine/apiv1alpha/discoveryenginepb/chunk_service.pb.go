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
// source: google/cloud/discoveryengine/v1alpha/chunk_service.proto

package discoveryenginepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [ChunkService.GetChunk][google.cloud.discoveryengine.v1alpha.ChunkService.GetChunk]
// method.
type GetChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of
	// [Chunk][google.cloud.discoveryengine.v1alpha.Chunk], such as
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document}/chunks/{chunk}`.
	//
	// If the caller does not have permission to access the
	// [Chunk][google.cloud.discoveryengine.v1alpha.Chunk], regardless of whether
	// or not it exists, a `PERMISSION_DENIED` error is returned.
	//
	// If the requested [Chunk][google.cloud.discoveryengine.v1alpha.Chunk] does
	// not exist, a `NOT_FOUND` error is returned.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetChunkRequest) Reset() {
	*x = GetChunkRequest{}
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkRequest) ProtoMessage() {}

func (x *GetChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkRequest.ProtoReflect.Descriptor instead.
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetChunkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for
// [ChunkService.ListChunks][google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks]
// method.
type ListChunksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent document resource name, such as
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document}`.
	//
	// If the caller does not have permission to list
	// [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s under this document,
	// regardless of whether or not this document exists, a `PERMISSION_DENIED`
	// error is returned.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Maximum number of [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s to
	// return. If unspecified, defaults to 100. The maximum allowed value is 1000.
	// Values above 1000 will be coerced to 1000.
	//
	// If this field is negative, an `INVALID_ARGUMENT` error is returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token
	// [ListChunksResponse.next_page_token][google.cloud.discoveryengine.v1alpha.ListChunksResponse.next_page_token],
	// received from a previous
	// [ChunkService.ListChunks][google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks]
	// call. Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to
	// [ChunkService.ListChunks][google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks]
	// must match the call that provided the page token. Otherwise, an
	// `INVALID_ARGUMENT` error is returned.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListChunksRequest) Reset() {
	*x = ListChunksRequest{}
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChunksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChunksRequest) ProtoMessage() {}

func (x *ListChunksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChunksRequest.ProtoReflect.Descriptor instead.
func (*ListChunksRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListChunksRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListChunksRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListChunksRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for
// [ChunkService.ListChunks][google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks]
// method.
type ListChunksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s.
	Chunks []*Chunk `protobuf:"bytes,1,rep,name=chunks,proto3" json:"chunks,omitempty"`
	// A token that can be sent as
	// [ListChunksRequest.page_token][google.cloud.discoveryengine.v1alpha.ListChunksRequest.page_token]
	// to retrieve the next page. If this field is omitted, there are no
	// subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListChunksResponse) Reset() {
	*x = ListChunksResponse{}
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChunksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChunksResponse) ProtoMessage() {}

func (x *ListChunksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChunksResponse.ProtoReflect.Descriptor instead.
func (*ListChunksResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListChunksResponse) GetChunks() []*Chunk {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *ListChunksResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_discoveryengine_v1alpha_chunk_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xed, 0x05, 0x0a, 0x0c, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb9, 0x02, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0xc8, 0x01, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xba, 0x01, 0x5a, 0x63, 0x12, 0x61,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x2f, 0x2a,
	0x7d, 0x12, 0x53, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a,
	0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xcc, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xca, 0x01, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xba, 0x01, 0x5a, 0x63, 0x12, 0x61, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x12, 0x53, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9d, 0x02, 0x0a, 0x28, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x11, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2,
	0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e,
	0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea,
	0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_goTypes = []any{
	(*GetChunkRequest)(nil),    // 0: google.cloud.discoveryengine.v1alpha.GetChunkRequest
	(*ListChunksRequest)(nil),  // 1: google.cloud.discoveryengine.v1alpha.ListChunksRequest
	(*ListChunksResponse)(nil), // 2: google.cloud.discoveryengine.v1alpha.ListChunksResponse
	(*Chunk)(nil),              // 3: google.cloud.discoveryengine.v1alpha.Chunk
}
var file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_depIdxs = []int32{
	3, // 0: google.cloud.discoveryengine.v1alpha.ListChunksResponse.chunks:type_name -> google.cloud.discoveryengine.v1alpha.Chunk
	0, // 1: google.cloud.discoveryengine.v1alpha.ChunkService.GetChunk:input_type -> google.cloud.discoveryengine.v1alpha.GetChunkRequest
	1, // 2: google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks:input_type -> google.cloud.discoveryengine.v1alpha.ListChunksRequest
	3, // 3: google.cloud.discoveryengine.v1alpha.ChunkService.GetChunk:output_type -> google.cloud.discoveryengine.v1alpha.Chunk
	2, // 4: google.cloud.discoveryengine.v1alpha.ChunkService.ListChunks:output_type -> google.cloud.discoveryengine.v1alpha.ListChunksResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_chunk_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_chunk_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_chunk_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChunkServiceClient is the client API for ChunkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChunkServiceClient interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*Chunk, error)
	// Gets a list of [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s.
	ListChunks(ctx context.Context, in *ListChunksRequest, opts ...grpc.CallOption) (*ListChunksResponse, error)
}

type chunkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServiceClient(cc grpc.ClientConnInterface) ChunkServiceClient {
	return &chunkServiceClient{cc}
}

func (c *chunkServiceClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.ChunkService/GetChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServiceClient) ListChunks(ctx context.Context, in *ListChunksRequest, opts ...grpc.CallOption) (*ListChunksResponse, error) {
	out := new(ListChunksResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.ChunkService/ListChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServiceServer is the server API for ChunkService service.
type ChunkServiceServer interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetChunk(context.Context, *GetChunkRequest) (*Chunk, error)
	// Gets a list of [Chunk][google.cloud.discoveryengine.v1alpha.Chunk]s.
	ListChunks(context.Context, *ListChunksRequest) (*ListChunksResponse, error)
}

// UnimplementedChunkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChunkServiceServer struct {
}

func (*UnimplementedChunkServiceServer) GetChunk(context.Context, *GetChunkRequest) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (*UnimplementedChunkServiceServer) ListChunks(context.Context, *ListChunksRequest) (*ListChunksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChunks not implemented")
}

func RegisterChunkServiceServer(s *grpc.Server, srv ChunkServiceServer) {
	s.RegisterService(&_ChunkService_serviceDesc, srv)
}

func _ChunkService_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.ChunkService/GetChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).GetChunk(ctx, req.(*GetChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkService_ListChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChunksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).ListChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.ChunkService/ListChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).ListChunks(ctx, req.(*ListChunksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChunkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.ChunkService",
	HandlerType: (*ChunkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChunk",
			Handler:    _ChunkService_GetChunk_Handler,
		},
		{
			MethodName: "ListChunks",
			Handler:    _ChunkService_ListChunks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/chunk_service.proto",
}
