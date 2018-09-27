// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/ad_group_status.proto

package google_ads_googleads_v0_enums

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The possible statuses of an ad group.
type AdGroupStatusEnum_AdGroupStatus int32

const (
	// The status has not been specified.
	AdGroupStatusEnum_UNSPECIFIED AdGroupStatusEnum_AdGroupStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupStatusEnum_UNKNOWN AdGroupStatusEnum_AdGroupStatus = 1
	// The ad group is enabled.
	AdGroupStatusEnum_ENABLED AdGroupStatusEnum_AdGroupStatus = 2
	// The ad group is paused.
	AdGroupStatusEnum_PAUSED AdGroupStatusEnum_AdGroupStatus = 3
	// The ad group is removed.
	AdGroupStatusEnum_REMOVED AdGroupStatusEnum_AdGroupStatus = 4
)

var AdGroupStatusEnum_AdGroupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupStatusEnum_AdGroupStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupStatusEnum_AdGroupStatus) String() string {
	return proto.EnumName(AdGroupStatusEnum_AdGroupStatus_name, int32(x))
}

func (AdGroupStatusEnum_AdGroupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22dc60055ea96e60, []int{0, 0}
}

// Container for enum describing possible statuses of an ad group.
type AdGroupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupStatusEnum) Reset()         { *m = AdGroupStatusEnum{} }
func (m *AdGroupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupStatusEnum) ProtoMessage()    {}
func (*AdGroupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_22dc60055ea96e60, []int{0}
}
func (m *AdGroupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupStatusEnum.Merge(m, src)
}
func (m *AdGroupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupStatusEnum.Size(m)
}
func (m *AdGroupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupStatusEnum)(nil), "google.ads.googleads.v0.enums.AdGroupStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupStatusEnum_AdGroupStatus", AdGroupStatusEnum_AdGroupStatus_name, AdGroupStatusEnum_AdGroupStatus_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/ad_group_status.proto", fileDescriptor_22dc60055ea96e60)
}

var fileDescriptor_22dc60055ea96e60 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xc4, 0x94, 0xf8, 0xf4, 0xa2, 0xfc, 0xd2, 0x82, 0xf8, 0xe2, 0x92, 0xc4,
	0x92, 0xd2, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x4a, 0xbd, 0xc4, 0x94,
	0x62, 0x3d, 0xb8, 0x26, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x26, 0xa5, 0x0c, 0x2e, 0x41, 0xc7, 0x14,
	0x77, 0x90, 0xb6, 0x60, 0xb0, 0x2e, 0xd7, 0xbc, 0xd2, 0x5c, 0xa5, 0x60, 0x2e, 0x5e, 0x14, 0x41,
	0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01,
	0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x10, 0xc7,
	0xd5, 0xcf, 0xd1, 0xc9, 0xc7, 0xd5, 0x45, 0x80, 0x49, 0x88, 0x8b, 0x8b, 0x2d, 0xc0, 0x31, 0x34,
	0xd8, 0xd5, 0x45, 0x80, 0x19, 0x24, 0x11, 0xe4, 0xea, 0xeb, 0x1f, 0xe6, 0xea, 0x22, 0xc0, 0xe2,
	0x34, 0x9f, 0x91, 0x4b, 0x31, 0x39, 0x3f, 0x57, 0x0f, 0xaf, 0x7b, 0x9c, 0x84, 0x50, 0x2c, 0x0e,
	0x00, 0x79, 0x21, 0x80, 0x71, 0x11, 0x13, 0xb3, 0xbb, 0xa3, 0xe3, 0x2a, 0x26, 0x59, 0x77, 0x88,
	0x5e, 0xc7, 0x94, 0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03, 0xb9, 0xba, 0xf8, 0x14,
	0x4c, 0x3e, 0xc6, 0x31, 0xa5, 0x38, 0x06, 0x2e, 0x1f, 0x13, 0x66, 0x10, 0x03, 0x96, 0x7f, 0x44,
	0x40, 0x3e, 0x89, 0x0d, 0x1c, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xa0, 0xc9,
	0xbe, 0x68, 0x01, 0x00, 0x00,
}
