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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/shopping/merchant/products/v1beta/productinputs.proto

package productspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	typepb "cloud.google.com/go/shopping/type/typepb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This resource represents input data you submit for a product, not the
// processed product that you see in Merchant Center, in Shopping ads, or across
// Google surfaces. Product inputs, rules and supplemental data source data are
// combined to create the processed
// [product][google.shopping.content.bundles.Products.Product].
//
// Required product input attributes to pass data validation checks are
// primarily defined in the [Products Data
// Specification](https://support.google.com/merchants/answer/188494).
//
// The following attributes are required:
// [feedLabel][google.shopping.content.bundles.Products.feed_label],
// [contentLanguage][google.shopping.content.bundles.Products.content_language]
// and [offerId][google.shopping.content.bundles.Products.offer_id].
//
// After inserting, updating, or deleting a product input, it may take several
// minutes before the processed product can be retrieved.
//
// All fields in the product input and its sub-messages match the English name
// of their corresponding attribute in the vertical spec with [some
// exceptions](https://support.google.com/merchants/answer/7052112).
type ProductInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The name of the product input.
	// Format:
	// `"{productinput.name=accounts/{account}/productInputs/{productinput}}"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The name of the processed product.
	// Format:
	// `"{product.name=accounts/{account}/products/{product}}"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	// Required. Immutable. The
	// [channel](https://support.google.com/merchants/answer/7361332) of the
	// product.
	Channel typepb.Channel_ChannelEnum `protobuf:"varint,3,opt,name=channel,proto3,enum=google.shopping.type.Channel_ChannelEnum" json:"channel,omitempty"`
	// Required. Immutable. Your unique identifier for the product. This is the
	// same for the product input and processed product. Leading and trailing
	// whitespaces are stripped and multiple whitespaces are replaced by a single
	// whitespace upon submission. See the [products data
	// specification](https://support.google.com/merchants/answer/188494#id) for
	// details.
	OfferId string `protobuf:"bytes,4,opt,name=offer_id,json=offerId,proto3" json:"offer_id,omitempty"`
	// Required. Immutable. The two-letter [ISO
	// 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code for the
	// product.
	ContentLanguage string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
	// Required. Immutable. The [feed
	// label](https://developers.google.com/shopping-content/guides/products/feed-labels)
	// for the product.
	FeedLabel string `protobuf:"bytes,6,opt,name=feed_label,json=feedLabel,proto3" json:"feed_label,omitempty"`
	// Optional. Represents the existing version (freshness) of the product, which
	// can be used to preserve the right order when multiple updates are done at
	// the same time.
	//
	// If set, the insertion is prevented when version number is lower than
	// the current version number of the existing product. Re-insertion (for
	// example, product refresh after 30 days) can be performed with the current
	// `version_number`.
	//
	// Only supported for insertions into primary data sources.
	//
	// If the operation is prevented, the aborted exception will be
	// thrown.
	VersionNumber *int64 `protobuf:"varint,7,opt,name=version_number,json=versionNumber,proto3,oneof" json:"version_number,omitempty"`
	// Optional. A list of product attributes.
	Attributes *Attributes `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Optional. A list of custom (merchant-provided) attributes. It can also be
	// used for submitting any attribute of the data specification in its generic
	// form (for example,
	// `{ "name": "size type", "value": "regular" }`).
	// This is useful for submitting attributes not explicitly exposed by the
	// API, such as additional attributes used for Buy on Google.
	// Maximum allowed number of characters for each
	// custom attribute is 10240 (represents sum of characters for name and
	// value). Maximum 2500 custom attributes can be set per product, with total
	// size of 102.4kB. Underscores in custom attribute names are replaced by
	// spaces upon insertion.
	CustomAttributes []*typepb.CustomAttribute `protobuf:"bytes,9,rep,name=custom_attributes,json=customAttributes,proto3" json:"custom_attributes,omitempty"`
}

func (x *ProductInput) Reset() {
	*x = ProductInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInput) ProtoMessage() {}

func (x *ProductInput) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInput.ProtoReflect.Descriptor instead.
func (*ProductInput) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescGZIP(), []int{0}
}

func (x *ProductInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductInput) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *ProductInput) GetChannel() typepb.Channel_ChannelEnum {
	if x != nil {
		return x.Channel
	}
	return typepb.Channel_ChannelEnum(0)
}

func (x *ProductInput) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *ProductInput) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

func (x *ProductInput) GetFeedLabel() string {
	if x != nil {
		return x.FeedLabel
	}
	return ""
}

func (x *ProductInput) GetVersionNumber() int64 {
	if x != nil && x.VersionNumber != nil {
		return *x.VersionNumber
	}
	return 0
}

func (x *ProductInput) GetAttributes() *Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ProductInput) GetCustomAttributes() []*typepb.CustomAttribute {
	if x != nil {
		return x.CustomAttributes
	}
	return nil
}

// Request message for the InsertProductInput method.
type InsertProductInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The account where this product will be inserted.
	// Format: accounts/{account}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The product input to insert.
	ProductInput *ProductInput `protobuf:"bytes,2,opt,name=product_input,json=productInput,proto3" json:"product_input,omitempty"`
	// Required. The primary or supplemental product data source name. If the
	// product already exists and data source provided is different, then the
	// product will be moved to a new data source. Format:
	// `accounts/{account}/dataSources/{datasource}`.
	DataSource string `protobuf:"bytes,3,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
}

func (x *InsertProductInputRequest) Reset() {
	*x = InsertProductInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertProductInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertProductInputRequest) ProtoMessage() {}

func (x *InsertProductInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertProductInputRequest.ProtoReflect.Descriptor instead.
func (*InsertProductInputRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescGZIP(), []int{1}
}

func (x *InsertProductInputRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *InsertProductInputRequest) GetProductInput() *ProductInput {
	if x != nil {
		return x.ProductInput
	}
	return nil
}

func (x *InsertProductInputRequest) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

// Request message for the DeleteProductInput method.
type DeleteProductInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the product input resource to delete.
	// Format: accounts/{account}/productInputs/{product}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The primary or supplemental data source from which the product
	// input should be deleted. Format:
	// `accounts/{account}/dataSources/{datasource}`.
	DataSource string `protobuf:"bytes,2,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
}

func (x *DeleteProductInputRequest) Reset() {
	*x = DeleteProductInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductInputRequest) ProtoMessage() {}

func (x *DeleteProductInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductInputRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductInputRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteProductInputRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteProductInputRequest) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

var File_google_shopping_merchant_products_v1beta_productinputs_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x05, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x42,
	0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x21, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0,
	0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0,
	0x41, 0x05, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2f, 0x0a,
	0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x59,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x3a, 0x7a, 0xea, 0x41, 0x77, 0x0a, 0x27, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x7d, 0x2a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x32, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0xe7, 0x01, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x12, 0x22, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x60, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a,
	0x27, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x32, 0xfc, 0x03, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe3, 0x01,
	0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x3a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x39, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x3a, 0x69, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x12, 0xb4, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x41, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x2a, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x47, 0xca, 0x41, 0x1a, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x94, 0x01, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x70, 0x62, 0x3b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescData = file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDesc
)

func file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescData)
	})
	return file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDescData
}

var file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_shopping_merchant_products_v1beta_productinputs_proto_goTypes = []interface{}{
	(*ProductInput)(nil),              // 0: google.shopping.merchant.products.v1beta.ProductInput
	(*InsertProductInputRequest)(nil), // 1: google.shopping.merchant.products.v1beta.InsertProductInputRequest
	(*DeleteProductInputRequest)(nil), // 2: google.shopping.merchant.products.v1beta.DeleteProductInputRequest
	(typepb.Channel_ChannelEnum)(0),   // 3: google.shopping.type.Channel.ChannelEnum
	(*Attributes)(nil),                // 4: google.shopping.merchant.products.v1beta.Attributes
	(*typepb.CustomAttribute)(nil),    // 5: google.shopping.type.CustomAttribute
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_google_shopping_merchant_products_v1beta_productinputs_proto_depIdxs = []int32{
	3, // 0: google.shopping.merchant.products.v1beta.ProductInput.channel:type_name -> google.shopping.type.Channel.ChannelEnum
	4, // 1: google.shopping.merchant.products.v1beta.ProductInput.attributes:type_name -> google.shopping.merchant.products.v1beta.Attributes
	5, // 2: google.shopping.merchant.products.v1beta.ProductInput.custom_attributes:type_name -> google.shopping.type.CustomAttribute
	0, // 3: google.shopping.merchant.products.v1beta.InsertProductInputRequest.product_input:type_name -> google.shopping.merchant.products.v1beta.ProductInput
	1, // 4: google.shopping.merchant.products.v1beta.ProductInputsService.InsertProductInput:input_type -> google.shopping.merchant.products.v1beta.InsertProductInputRequest
	2, // 5: google.shopping.merchant.products.v1beta.ProductInputsService.DeleteProductInput:input_type -> google.shopping.merchant.products.v1beta.DeleteProductInputRequest
	0, // 6: google.shopping.merchant.products.v1beta.ProductInputsService.InsertProductInput:output_type -> google.shopping.merchant.products.v1beta.ProductInput
	6, // 7: google.shopping.merchant.products.v1beta.ProductInputsService.DeleteProductInput:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_products_v1beta_productinputs_proto_init() }
func file_google_shopping_merchant_products_v1beta_productinputs_proto_init() {
	if File_google_shopping_merchant_products_v1beta_productinputs_proto != nil {
		return
	}
	file_google_shopping_merchant_products_v1beta_products_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInput); i {
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
		file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertProductInputRequest); i {
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
		file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductInputRequest); i {
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
	file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_merchant_products_v1beta_productinputs_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_products_v1beta_productinputs_proto_depIdxs,
		MessageInfos:      file_google_shopping_merchant_products_v1beta_productinputs_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_products_v1beta_productinputs_proto = out.File
	file_google_shopping_merchant_products_v1beta_productinputs_proto_rawDesc = nil
	file_google_shopping_merchant_products_v1beta_productinputs_proto_goTypes = nil
	file_google_shopping_merchant_products_v1beta_productinputs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductInputsServiceClient is the client API for ProductInputsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductInputsServiceClient interface {
	// Uploads a product input to your Merchant Center account. If an input
	// with the same contentLanguage, offerId, and dataSource already exists,
	// this method replaces that entry.
	//
	// After inserting, updating, or deleting a product input, it may take several
	// minutes before the processed product can be retrieved.
	InsertProductInput(ctx context.Context, in *InsertProductInputRequest, opts ...grpc.CallOption) (*ProductInput, error)
	// Deletes a product input from your Merchant Center account.
	//
	// After inserting, updating, or deleting a product input, it may take several
	// minutes before the processed product can be retrieved.
	DeleteProductInput(ctx context.Context, in *DeleteProductInputRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type productInputsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInputsServiceClient(cc grpc.ClientConnInterface) ProductInputsServiceClient {
	return &productInputsServiceClient{cc}
}

func (c *productInputsServiceClient) InsertProductInput(ctx context.Context, in *InsertProductInputRequest, opts ...grpc.CallOption) (*ProductInput, error) {
	out := new(ProductInput)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.products.v1beta.ProductInputsService/InsertProductInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInputsServiceClient) DeleteProductInput(ctx context.Context, in *DeleteProductInputRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.products.v1beta.ProductInputsService/DeleteProductInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInputsServiceServer is the server API for ProductInputsService service.
type ProductInputsServiceServer interface {
	// Uploads a product input to your Merchant Center account. If an input
	// with the same contentLanguage, offerId, and dataSource already exists,
	// this method replaces that entry.
	//
	// After inserting, updating, or deleting a product input, it may take several
	// minutes before the processed product can be retrieved.
	InsertProductInput(context.Context, *InsertProductInputRequest) (*ProductInput, error)
	// Deletes a product input from your Merchant Center account.
	//
	// After inserting, updating, or deleting a product input, it may take several
	// minutes before the processed product can be retrieved.
	DeleteProductInput(context.Context, *DeleteProductInputRequest) (*emptypb.Empty, error)
}

// UnimplementedProductInputsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductInputsServiceServer struct {
}

func (*UnimplementedProductInputsServiceServer) InsertProductInput(context.Context, *InsertProductInputRequest) (*ProductInput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProductInput not implemented")
}
func (*UnimplementedProductInputsServiceServer) DeleteProductInput(context.Context, *DeleteProductInputRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductInput not implemented")
}

func RegisterProductInputsServiceServer(s *grpc.Server, srv ProductInputsServiceServer) {
	s.RegisterService(&_ProductInputsService_serviceDesc, srv)
}

func _ProductInputsService_InsertProductInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertProductInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInputsServiceServer).InsertProductInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.products.v1beta.ProductInputsService/InsertProductInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInputsServiceServer).InsertProductInput(ctx, req.(*InsertProductInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInputsService_DeleteProductInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInputsServiceServer).DeleteProductInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.products.v1beta.ProductInputsService/DeleteProductInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInputsServiceServer).DeleteProductInput(ctx, req.(*DeleteProductInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductInputsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.products.v1beta.ProductInputsService",
	HandlerType: (*ProductInputsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertProductInput",
			Handler:    _ProductInputsService_InsertProductInput_Handler,
		},
		{
			MethodName: "DeleteProductInput",
			Handler:    _ProductInputsService_DeleteProductInput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/products/v1beta/productinputs.proto",
}