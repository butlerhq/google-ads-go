// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/listing_group_type.proto

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

// The type of the listing group.
type ListingGroupTypeEnum_ListingGroupType int32

const (
	// Not specified.
	ListingGroupTypeEnum_UNSPECIFIED ListingGroupTypeEnum_ListingGroupType = 0
	// Used for return value only. Represents value unknown in this version.
	ListingGroupTypeEnum_UNKNOWN ListingGroupTypeEnum_ListingGroupType = 1
	// Subdivision of products along some listing dimension. These nodes
	// are not used by serving to target listing entries, but is purely
	// to define the structure of the tree.
	ListingGroupTypeEnum_SUBDIVISION ListingGroupTypeEnum_ListingGroupType = 2
	// Listing group unit that defines a bid.
	ListingGroupTypeEnum_UNIT ListingGroupTypeEnum_ListingGroupType = 3
)

var ListingGroupTypeEnum_ListingGroupType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SUBDIVISION",
	3: "UNIT",
}

var ListingGroupTypeEnum_ListingGroupType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SUBDIVISION": 2,
	"UNIT":        3,
}

func (x ListingGroupTypeEnum_ListingGroupType) String() string {
	return proto.EnumName(ListingGroupTypeEnum_ListingGroupType_name, int32(x))
}

func (ListingGroupTypeEnum_ListingGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9396e74fe0a1e231, []int{0, 0}
}

// Container for enum describing the type of the listing group.
type ListingGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListingGroupTypeEnum) Reset()         { *m = ListingGroupTypeEnum{} }
func (m *ListingGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ListingGroupTypeEnum) ProtoMessage()    {}
func (*ListingGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9396e74fe0a1e231, []int{0}
}
func (m *ListingGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListingGroupTypeEnum.Unmarshal(m, b)
}
func (m *ListingGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListingGroupTypeEnum.Marshal(b, m, deterministic)
}
func (m *ListingGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListingGroupTypeEnum.Merge(m, src)
}
func (m *ListingGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ListingGroupTypeEnum.Size(m)
}
func (m *ListingGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ListingGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ListingGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListingGroupTypeEnum)(nil), "google.ads.googleads.v0.enums.ListingGroupTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ListingGroupTypeEnum_ListingGroupType", ListingGroupTypeEnum_ListingGroupType_name, ListingGroupTypeEnum_ListingGroupType_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/listing_group_type.proto", fileDescriptor_9396e74fe0a1e231)
}

var fileDescriptor_9396e74fe0a1e231 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x9c, 0xcc, 0xe2, 0x92, 0xcc, 0xbc, 0xf4, 0xf8, 0xf4, 0xa2, 0xfc, 0xd2,
	0x82, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x62,
	0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x3e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x3e, 0xa5, 0x64, 0x2e,
	0x11, 0x1f, 0x88, 0x56, 0x77, 0x90, 0xce, 0x90, 0xca, 0x82, 0x54, 0xd7, 0xbc, 0xd2, 0x5c, 0x25,
	0x6f, 0x2e, 0x01, 0x74, 0x71, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f,
	0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70,
	0x3f, 0x01, 0x46, 0x90, 0x6c, 0x70, 0xa8, 0x93, 0x8b, 0x67, 0x98, 0x67, 0xb0, 0xa7, 0xbf, 0x9f,
	0x00, 0x93, 0x10, 0x07, 0x17, 0x4b, 0xa8, 0x9f, 0x67, 0x88, 0x00, 0xb3, 0xd3, 0x22, 0x46, 0x2e,
	0xc5, 0xe4, 0xfc, 0x5c, 0x3d, 0xbc, 0x4e, 0x71, 0x12, 0x45, 0xb7, 0x30, 0x00, 0xe4, 0x81, 0x00,
	0xc6, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab, 0x98, 0x64, 0xdd, 0x21, 0xda, 0x1d, 0x53, 0x8a,
	0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0xe4, 0xe0, 0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7,
	0x94, 0xe2, 0x18, 0xb8, 0x7c, 0x4c, 0x98, 0x41, 0x0c, 0x58, 0xfe, 0x11, 0x01, 0xf9, 0x24, 0x36,
	0x70, 0x78, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x79, 0xdb, 0x28, 0x69, 0x01, 0x00,
	0x00,
}
