// Copyright 2020 Google LLC
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
// 	protoc        v3.21.5
// source: google/cloud/notebooks/v1beta1/environment.proto

package notebookspb

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

// Definition of a software environment that is used to start a notebook
// instance.
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of this environment.
	// Format:
	// `projects/{project_id}/locations/{location}/environments/{environment_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Display name of this environment for the UI.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A brief description of this environment.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Type of the environment; can be one of VM image, or container image.
	//
	// Types that are assignable to ImageType:
	//
	//	*Environment_VmImage
	//	*Environment_ContainerImage
	ImageType isEnvironment_ImageType `protobuf_oneof:"image_type"`
	// Path to a Bash script that automatically runs after a notebook instance
	// fully boots up. The path must be a URL or
	// Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	PostStartupScript string `protobuf:"bytes,8,opt,name=post_startup_script,json=postStartupScript,proto3" json:"post_startup_script,omitempty"`
	// Output only. The time at which this environment was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1beta1_environment_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Environment) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Environment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Environment) GetImageType() isEnvironment_ImageType {
	if m != nil {
		return m.ImageType
	}
	return nil
}

func (x *Environment) GetVmImage() *VmImage {
	if x, ok := x.GetImageType().(*Environment_VmImage); ok {
		return x.VmImage
	}
	return nil
}

func (x *Environment) GetContainerImage() *ContainerImage {
	if x, ok := x.GetImageType().(*Environment_ContainerImage); ok {
		return x.ContainerImage
	}
	return nil
}

func (x *Environment) GetPostStartupScript() string {
	if x != nil {
		return x.PostStartupScript
	}
	return ""
}

func (x *Environment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type isEnvironment_ImageType interface {
	isEnvironment_ImageType()
}

type Environment_VmImage struct {
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage *VmImage `protobuf:"bytes,6,opt,name=vm_image,json=vmImage,proto3,oneof"`
}

type Environment_ContainerImage struct {
	// Use a container image to start the notebook instance.
	ContainerImage *ContainerImage `protobuf:"bytes,7,opt,name=container_image,json=containerImage,proto3,oneof"`
}

func (*Environment_VmImage) isEnvironment_ImageType() {}

func (*Environment_ContainerImage) isEnvironment_ImageType() {}

// Definition of a custom Compute Engine virtual machine image for starting a
// notebook instance with the environment installed directly on the VM.
type VmImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Google Cloud project that this VM image belongs to.
	// Format: `projects/{project_id}`
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The reference to an external Compute Engine VM image.
	//
	// Types that are assignable to Image:
	//
	//	*VmImage_ImageName
	//	*VmImage_ImageFamily
	Image isVmImage_Image `protobuf_oneof:"image"`
}

func (x *VmImage) Reset() {
	*x = VmImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmImage) ProtoMessage() {}

func (x *VmImage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmImage.ProtoReflect.Descriptor instead.
func (*VmImage) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1beta1_environment_proto_rawDescGZIP(), []int{1}
}

func (x *VmImage) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (m *VmImage) GetImage() isVmImage_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (x *VmImage) GetImageName() string {
	if x, ok := x.GetImage().(*VmImage_ImageName); ok {
		return x.ImageName
	}
	return ""
}

func (x *VmImage) GetImageFamily() string {
	if x, ok := x.GetImage().(*VmImage_ImageFamily); ok {
		return x.ImageFamily
	}
	return ""
}

type isVmImage_Image interface {
	isVmImage_Image()
}

type VmImage_ImageName struct {
	// Use VM image name to find the image.
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName,proto3,oneof"`
}

type VmImage_ImageFamily struct {
	// Use this VM image family to find the image; the newest image in this
	// family will be used.
	ImageFamily string `protobuf:"bytes,3,opt,name=image_family,json=imageFamily,proto3,oneof"`
}

func (*VmImage_ImageName) isVmImage_Image() {}

func (*VmImage_ImageFamily) isVmImage_Image() {}

// Definition of a container image for starting a notebook instance with the
// environment installed in a container.
type ContainerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The path to the container image repository. For example:
	// `gcr.io/{project_id}/{image_name}`
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// The tag of the container image. If not specified, this defaults
	// to the latest tag.
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *ContainerImage) Reset() {
	*x = ContainerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage) ProtoMessage() {}

func (x *ContainerImage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage.ProtoReflect.Descriptor instead.
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return file_google_cloud_notebooks_v1beta1_environment_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerImage) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ContainerImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_google_cloud_notebooks_v1beta1_environment_proto protoreflect.FileDescriptor

var file_google_cloud_notebooks_v1beta1_environment_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe6, 0x03, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x08, 0x76, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x56, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x76, 0x6d, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6f, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a, 0x24, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x77, 0x0a, 0x07, 0x56, 0x6d, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x47, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x42, 0xe7, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x6f,
	0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_notebooks_v1beta1_environment_proto_rawDescOnce sync.Once
	file_google_cloud_notebooks_v1beta1_environment_proto_rawDescData = file_google_cloud_notebooks_v1beta1_environment_proto_rawDesc
)

func file_google_cloud_notebooks_v1beta1_environment_proto_rawDescGZIP() []byte {
	file_google_cloud_notebooks_v1beta1_environment_proto_rawDescOnce.Do(func() {
		file_google_cloud_notebooks_v1beta1_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_notebooks_v1beta1_environment_proto_rawDescData)
	})
	return file_google_cloud_notebooks_v1beta1_environment_proto_rawDescData
}

var file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_notebooks_v1beta1_environment_proto_goTypes = []interface{}{
	(*Environment)(nil),           // 0: google.cloud.notebooks.v1beta1.Environment
	(*VmImage)(nil),               // 1: google.cloud.notebooks.v1beta1.VmImage
	(*ContainerImage)(nil),        // 2: google.cloud.notebooks.v1beta1.ContainerImage
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_notebooks_v1beta1_environment_proto_depIdxs = []int32{
	1, // 0: google.cloud.notebooks.v1beta1.Environment.vm_image:type_name -> google.cloud.notebooks.v1beta1.VmImage
	2, // 1: google.cloud.notebooks.v1beta1.Environment.container_image:type_name -> google.cloud.notebooks.v1beta1.ContainerImage
	3, // 2: google.cloud.notebooks.v1beta1.Environment.create_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_notebooks_v1beta1_environment_proto_init() }
func file_google_cloud_notebooks_v1beta1_environment_proto_init() {
	if File_google_cloud_notebooks_v1beta1_environment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmImage); i {
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
		file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage); i {
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
	file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Environment_VmImage)(nil),
		(*Environment_ContainerImage)(nil),
	}
	file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*VmImage_ImageName)(nil),
		(*VmImage_ImageFamily)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_notebooks_v1beta1_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_notebooks_v1beta1_environment_proto_goTypes,
		DependencyIndexes: file_google_cloud_notebooks_v1beta1_environment_proto_depIdxs,
		MessageInfos:      file_google_cloud_notebooks_v1beta1_environment_proto_msgTypes,
	}.Build()
	File_google_cloud_notebooks_v1beta1_environment_proto = out.File
	file_google_cloud_notebooks_v1beta1_environment_proto_rawDesc = nil
	file_google_cloud_notebooks_v1beta1_environment_proto_goTypes = nil
	file_google_cloud_notebooks_v1beta1_environment_proto_depIdxs = nil
}
