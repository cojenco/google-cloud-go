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
// source: google/shopping/merchant/accounts/v1beta/tax_rule.proto

package accountspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	interval "google.golang.org/genproto/googleapis/type/interval"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Primary type convension
//
// percent micro : 100% = 1 000 000 and 1% = 10 000
//
//	cannot be negative.
//
// Information about tax nexus and related parameters applicable to orders
// delivered to the area covered by a single tax admin. Nexus is created when a
// merchant is doing business in an area administered by tax admin (only US
// states are supported for nexus configuration). If merchant has nexus in a US
// state, merchant needs to pay tax to all tax authorities associated with
// the shipping destination.
// Next Id : 8
type TaxRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describe the location through either postal code range or a criteria id.
	//
	// Types that are assignable to Location:
	//
	//	*TaxRule_LocationId
	//	*TaxRule_PostCodeRange
	Location isTaxRule_Location `protobuf_oneof:"location"`
	// What is the way to calculate tax rate for deliveries to this admin's area.
	// Can only be set on US states.
	//
	// Types that are assignable to RateCalculation:
	//
	//	*TaxRule_UseGoogleRate
	//	*TaxRule_SelfSpecifiedRateMicros
	RateCalculation isTaxRule_RateCalculation `protobuf_oneof:"rate_calculation"`
	// Region code in which this rule is applicable
	RegionCode string `protobuf:"bytes,1,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// If set, shipping charge is taxed (at the same rate as product) when
	// delivering to this admin's area.
	// Can only be set on US states without category.
	ShippingTaxed bool `protobuf:"varint,6,opt,name=shipping_taxed,json=shippingTaxed,proto3" json:"shipping_taxed,omitempty"`
	// Required. Time period when this rule is effective. If the duration is
	// missing from effective_time listed, then it is open ended to the future.
	// The start of this time period is inclusive, and the end is exclusive.
	EffectiveTimePeriod *interval.Interval `protobuf:"bytes,7,opt,name=effective_time_period,json=effectiveTimePeriod,proto3" json:"effective_time_period,omitempty"`
}

func (x *TaxRule) Reset() {
	*x = TaxRule{}
	mi := &file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaxRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxRule) ProtoMessage() {}

func (x *TaxRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxRule.ProtoReflect.Descriptor instead.
func (*TaxRule) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescGZIP(), []int{0}
}

func (m *TaxRule) GetLocation() isTaxRule_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (x *TaxRule) GetLocationId() int64 {
	if x, ok := x.GetLocation().(*TaxRule_LocationId); ok {
		return x.LocationId
	}
	return 0
}

func (x *TaxRule) GetPostCodeRange() *TaxRule_TaxPostalCodeRange {
	if x, ok := x.GetLocation().(*TaxRule_PostCodeRange); ok {
		return x.PostCodeRange
	}
	return nil
}

func (m *TaxRule) GetRateCalculation() isTaxRule_RateCalculation {
	if m != nil {
		return m.RateCalculation
	}
	return nil
}

func (x *TaxRule) GetUseGoogleRate() bool {
	if x, ok := x.GetRateCalculation().(*TaxRule_UseGoogleRate); ok {
		return x.UseGoogleRate
	}
	return false
}

func (x *TaxRule) GetSelfSpecifiedRateMicros() int64 {
	if x, ok := x.GetRateCalculation().(*TaxRule_SelfSpecifiedRateMicros); ok {
		return x.SelfSpecifiedRateMicros
	}
	return 0
}

func (x *TaxRule) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *TaxRule) GetShippingTaxed() bool {
	if x != nil {
		return x.ShippingTaxed
	}
	return false
}

func (x *TaxRule) GetEffectiveTimePeriod() *interval.Interval {
	if x != nil {
		return x.EffectiveTimePeriod
	}
	return nil
}

type isTaxRule_Location interface {
	isTaxRule_Location()
}

type TaxRule_LocationId struct {
	// The admin_id or criteria_id of the region in which this rule is
	// applicable.
	LocationId int64 `protobuf:"varint,2,opt,name=location_id,json=locationId,proto3,oneof"`
}

type TaxRule_PostCodeRange struct {
	// The range of postal codes in which this rule is applicable.
	PostCodeRange *TaxRule_TaxPostalCodeRange `protobuf:"bytes,3,opt,name=post_code_range,json=postCodeRange,proto3,oneof"`
}

func (*TaxRule_LocationId) isTaxRule_Location() {}

func (*TaxRule_PostCodeRange) isTaxRule_Location() {}

type isTaxRule_RateCalculation interface {
	isTaxRule_RateCalculation()
}

type TaxRule_UseGoogleRate struct {
	// Rate that depends on delivery location: if merchant has a nexus in
	// corresponding US state, rates from authorities with jurisdiction over
	// delivery area are added up.
	UseGoogleRate bool `protobuf:"varint,4,opt,name=use_google_rate,json=useGoogleRate,proto3,oneof"`
}

type TaxRule_SelfSpecifiedRateMicros struct {
	// A fixed rate specified in micros, where 100% = 1_000_000.
	// Suitable for origin-based states.
	SelfSpecifiedRateMicros int64 `protobuf:"varint,5,opt,name=self_specified_rate_micros,json=selfSpecifiedRateMicros,proto3,oneof"`
}

func (*TaxRule_UseGoogleRate) isTaxRule_RateCalculation() {}

func (*TaxRule_SelfSpecifiedRateMicros) isTaxRule_RateCalculation() {}

// A range of postal codes that defines the area.
type TaxRule_TaxPostalCodeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The start of the postal code range, which is also the smallest
	// in the range.
	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end of the postal code range. Will be the same as start if not
	// specified.
	End string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TaxRule_TaxPostalCodeRange) Reset() {
	*x = TaxRule_TaxPostalCodeRange{}
	mi := &file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaxRule_TaxPostalCodeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxRule_TaxPostalCodeRange) ProtoMessage() {}

func (x *TaxRule_TaxPostalCodeRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxRule_TaxPostalCodeRange.ProtoReflect.Descriptor instead.
func (*TaxRule_TaxPostalCodeRange) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TaxRule_TaxPostalCodeRange) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *TaxRule_TaxPostalCodeRange) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

var File_google_shopping_merchant_accounts_v1beta_tax_rule_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x78, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x80, 0x04, 0x0a, 0x07, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x6e, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x78, 0x50,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x1a, 0x73, 0x65, 0x6c,
	0x66, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x17, 0x73, 0x65, 0x6c, 0x66, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x61,
	0x74, 0x65, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x78, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x78, 0x65, 0x64,
	0x12, 0x4e, 0x0a, 0x15, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x1a, 0x41, 0x0a, 0x12, 0x54, 0x61, 0x78, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x12, 0x0a, 0x10, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x8e, 0x01, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x0c, 0x54, 0x61, 0x78, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescData = file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDesc
)

func file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescData)
	})
	return file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDescData
}

var file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_goTypes = []any{
	(*TaxRule)(nil),                    // 0: google.shopping.merchant.accounts.v1beta.TaxRule
	(*TaxRule_TaxPostalCodeRange)(nil), // 1: google.shopping.merchant.accounts.v1beta.TaxRule.TaxPostalCodeRange
	(*interval.Interval)(nil),          // 2: google.type.Interval
}
var file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_depIdxs = []int32{
	1, // 0: google.shopping.merchant.accounts.v1beta.TaxRule.post_code_range:type_name -> google.shopping.merchant.accounts.v1beta.TaxRule.TaxPostalCodeRange
	2, // 1: google.shopping.merchant.accounts.v1beta.TaxRule.effective_time_period:type_name -> google.type.Interval
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_init() }
func file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_init() {
	if File_google_shopping_merchant_accounts_v1beta_tax_rule_proto != nil {
		return
	}
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes[0].OneofWrappers = []any{
		(*TaxRule_LocationId)(nil),
		(*TaxRule_PostCodeRange)(nil),
		(*TaxRule_UseGoogleRate)(nil),
		(*TaxRule_SelfSpecifiedRateMicros)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_depIdxs,
		MessageInfos:      file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_accounts_v1beta_tax_rule_proto = out.File
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_rawDesc = nil
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_goTypes = nil
	file_google_shopping_merchant_accounts_v1beta_tax_rule_proto_depIdxs = nil
}
