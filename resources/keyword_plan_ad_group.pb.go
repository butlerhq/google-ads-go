// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/keyword_plan_ad_group.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A Keyword Planner ad group.
// Max number of keyword plan ad groups per plan: 50.
type KeywordPlanAdGroup struct {
	// The resource name of the Keyword Planner ad group.
	// KeywordPlanAdGroup resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanAdGroups/{kp_ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan campaign to which this ad group belongs.
	KeywordPlanCampaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3" json:"keyword_plan_campaign,omitempty"`
	// The ID of the keyword plan ad group.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the keyword plan ad group.
	//
	// This field is required and should not be empty when creating keyword plan
	// ad group.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A default ad group max cpc bid in micros in account currency for all
	// biddable keywords under the keyword plan ad group.
	// If not set, will inherit from parent campaign.
	CpcBidMicros         *wrappers.Int64Value `protobuf:"bytes,5,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeywordPlanAdGroup) Reset()         { *m = KeywordPlanAdGroup{} }
func (m *KeywordPlanAdGroup) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanAdGroup) ProtoMessage()    {}
func (*KeywordPlanAdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_3824dce59517be04, []int{0}
}

func (m *KeywordPlanAdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanAdGroup.Unmarshal(m, b)
}
func (m *KeywordPlanAdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanAdGroup.Marshal(b, m, deterministic)
}
func (m *KeywordPlanAdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanAdGroup.Merge(m, src)
}
func (m *KeywordPlanAdGroup) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanAdGroup.Size(m)
}
func (m *KeywordPlanAdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanAdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanAdGroup proto.InternalMessageInfo

func (m *KeywordPlanAdGroup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanAdGroup) GetKeywordPlanCampaign() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlanCampaign
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlanAdGroup) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanAdGroup)(nil), "google.ads.googleads.v0.resources.KeywordPlanAdGroup")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/keyword_plan_ad_group.proto", fileDescriptor_3824dce59517be04)
}

var fileDescriptor_3824dce59517be04 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x1c, 0xc5, 0x69, 0x36, 0x05, 0xe3, 0xf4, 0xa2, 0x22, 0x16, 0x15, 0xd9, 0x94, 0xc1, 0x40, 0x48,
	0x8b, 0x8a, 0x17, 0x11, 0x2f, 0x3a, 0x2f, 0x86, 0x8a, 0x52, 0x26, 0xf4, 0x42, 0x0a, 0x25, 0x4b,
	0x62, 0x08, 0x6b, 0x93, 0x90, 0xac, 0x1b, 0xde, 0xfb, 0x24, 0x5e, 0x0a, 0xbe, 0x88, 0x8f, 0xe2,
	0x53, 0x48, 0x9b, 0xb6, 0x22, 0x03, 0xf7, 0xdd, 0x1d, 0xda, 0xf3, 0x3b, 0xe7, 0x10, 0xfe, 0xf0,
	0xb5, 0xd0, 0x5a, 0x54, 0x3c, 0x26, 0xcc, 0xc5, 0x5e, 0xb6, 0xea, 0x98, 0xc4, 0x96, 0x3b, 0xdd,
	0x58, 0xca, 0x5d, 0xbc, 0xe7, 0x5f, 0x4f, 0xda, 0xb2, 0xd2, 0x54, 0x44, 0x95, 0x84, 0x95, 0xc2,
	0xea, 0xc6, 0x20, 0x63, 0xf5, 0x41, 0x87, 0x0b, 0xcf, 0x20, 0xc2, 0x1c, 0x1a, 0x71, 0x74, 0x4c,
	0xd0, 0x88, 0xdf, 0x7f, 0xd4, 0x37, 0x74, 0xc0, 0xae, 0xf9, 0x12, 0x9f, 0x2c, 0x31, 0x86, 0x5b,
	0xe7, 0x23, 0x1e, 0xff, 0x04, 0x30, 0x7c, 0xef, 0x2b, 0xb2, 0x8a, 0xa8, 0x94, 0x6d, 0xda, 0xfc,
	0xf0, 0x09, 0xbc, 0x35, 0x64, 0x94, 0x8a, 0xd4, 0x3c, 0x0a, 0xe6, 0xc1, 0xea, 0xc6, 0x76, 0x36,
	0x7c, 0xfc, 0x48, 0x6a, 0x1e, 0x66, 0xf0, 0xee, 0x3f, 0xeb, 0x28, 0xa9, 0x0d, 0x91, 0x42, 0x45,
	0x60, 0x1e, 0xac, 0x6e, 0x3e, 0x7b, 0xd8, 0x6f, 0x42, 0x43, 0x37, 0xfa, 0x74, 0xb0, 0x52, 0x89,
	0x9c, 0x54, 0x0d, 0xdf, 0xde, 0xd9, 0xff, 0x6d, 0x7d, 0xd3, 0x83, 0xe1, 0x53, 0x08, 0x24, 0x8b,
	0x26, 0x1d, 0xfe, 0xe0, 0x0c, 0x7f, 0xab, 0x0e, 0x2f, 0x5f, 0x78, 0x1a, 0x48, 0x16, 0x26, 0x70,
	0xda, 0x4d, 0x9b, 0x5e, 0xa1, 0xad, 0x73, 0x86, 0x29, 0xbc, 0x4d, 0x0d, 0x2d, 0x77, 0x92, 0x95,
	0xb5, 0xa4, 0x56, 0xbb, 0xe8, 0xda, 0xe5, 0xaa, 0x19, 0x35, 0x74, 0x2d, 0xd9, 0x87, 0x0e, 0x58,
	0x7f, 0x03, 0x70, 0x49, 0x75, 0x8d, 0x2e, 0xbe, 0xfc, 0xfa, 0xde, 0xf9, 0xb3, 0x66, 0x6d, 0x7c,
	0x16, 0x7c, 0x7e, 0xd7, 0xd3, 0x42, 0x57, 0x44, 0x09, 0xa4, 0xad, 0x88, 0x05, 0x57, 0x5d, 0xf9,
	0x70, 0x06, 0x46, 0xba, 0xff, 0x5c, 0xc5, 0xab, 0x51, 0x7d, 0x07, 0x93, 0x4d, 0x9a, 0xfe, 0x00,
	0x8b, 0x8d, 0x8f, 0x4c, 0x99, 0x43, 0x5e, 0xb6, 0x2a, 0x4f, 0xd0, 0x76, 0x70, 0xfe, 0x1a, 0x3c,
	0x45, 0xca, 0x5c, 0x31, 0x7a, 0x8a, 0x3c, 0x29, 0x46, 0xcf, 0x6f, 0xb0, 0xf4, 0x3f, 0x30, 0x4e,
	0x99, 0xc3, 0x78, 0x74, 0x61, 0x9c, 0x27, 0x18, 0x8f, 0xbe, 0xdd, 0xf5, 0x6e, 0xec, 0xf3, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x19, 0xa3, 0x92, 0xc1, 0x02, 0x00, 0x00,
}