// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/enums/ad_group_criterion_status.proto

package enums

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

// The possible statuses of an AdGroupCriterion.
type AdGroupCriterionStatusEnum_AdGroupCriterionStatus int32

const (
	// No value has been specified.
	AdGroupCriterionStatusEnum_UNSPECIFIED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupCriterionStatusEnum_UNKNOWN AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 1
	// The ad group criterion is enabled.
	AdGroupCriterionStatusEnum_ENABLED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 2
	// The ad group criterion is paused.
	AdGroupCriterionStatusEnum_PAUSED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 3
	// The ad group criterion is removed.
	AdGroupCriterionStatusEnum_REMOVED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 4
)

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupCriterionStatusEnum_AdGroupCriterionStatus) String() string {
	return proto.EnumName(AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, int32(x))
}

func (AdGroupCriterionStatusEnum_AdGroupCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2ee8cdd92377e78, []int{0, 0}
}

// Message describing AdGroupCriterion statuses.
type AdGroupCriterionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionStatusEnum) Reset()         { *m = AdGroupCriterionStatusEnum{} }
func (m *AdGroupCriterionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionStatusEnum) ProtoMessage()    {}
func (*AdGroupCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ee8cdd92377e78, []int{0}
}
func (m *AdGroupCriterionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionStatusEnum.Merge(m, src)
}
func (m *AdGroupCriterionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Size(m)
}
func (m *AdGroupCriterionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupCriterionStatusEnum)(nil), "google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus", AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/enums/ad_group_criterion_status.proto", fileDescriptor_c2ee8cdd92377e78)
}

var fileDescriptor_c2ee8cdd92377e78 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xc4, 0x94, 0xf8, 0xf4, 0xa2, 0xfc, 0xd2, 0x82, 0xf8, 0xe4, 0xa2, 0xcc,
	0x92, 0xd4, 0xa2, 0xcc, 0xfc, 0xbc, 0xf8, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x1e, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x76, 0xbd, 0x32,
	0x03, 0x3d, 0xb0, 0x76, 0xa5, 0x2a, 0x2e, 0x29, 0xc7, 0x14, 0x77, 0x90, 0x01, 0xce, 0x30, 0xfd,
	0xc1, 0x60, 0xed, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0x31, 0x5c, 0x62, 0xd8, 0x65, 0x85, 0xf8, 0xb9,
	0xb8, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0xb8,
	0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x41, 0x1c, 0x57, 0x3f, 0x47,
	0x27, 0x1f, 0x57, 0x17, 0x01, 0x26, 0x21, 0x2e, 0x2e, 0xb6, 0x00, 0xc7, 0xd0, 0x60, 0x57, 0x17,
	0x01, 0x66, 0x90, 0x44, 0x90, 0xab, 0xaf, 0x7f, 0x98, 0xab, 0x8b, 0x00, 0x8b, 0xd3, 0x0a, 0x46,
	0x2e, 0xc5, 0xe4, 0xfc, 0x5c, 0x3d, 0xbc, 0x2e, 0x74, 0x92, 0xc6, 0xee, 0x82, 0x00, 0x90, 0xef,
	0x02, 0x18, 0x17, 0x31, 0x31, 0xbb, 0x3b, 0x3a, 0xae, 0x62, 0x92, 0x75, 0x87, 0x18, 0xe2, 0x98,
	0x52, 0xac, 0x07, 0x61, 0x82, 0x58, 0x61, 0x06, 0x7a, 0x20, 0x7f, 0x14, 0x9f, 0x82, 0xc9, 0xc7,
	0x38, 0xa6, 0x14, 0xc7, 0xc0, 0xe5, 0x63, 0xc2, 0x0c, 0x62, 0xc0, 0xf2, 0x8f, 0x08, 0xc8, 0x27,
	0xb1, 0x81, 0x03, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x20, 0x2e, 0xf0, 0x8d, 0x01,
	0x00, 0x00,
}