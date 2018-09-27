// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/resource_access_denied_error.proto

package google_ads_googleads_v0_errors

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

// Enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError int32

const (
	// Enum unspecified.
	ResourceAccessDeniedErrorEnum_UNSPECIFIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 0
	// The received error code is not known in this version.
	ResourceAccessDeniedErrorEnum_UNKNOWN ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 1
	// User did not have write access.
	ResourceAccessDeniedErrorEnum_WRITE_ACCESS_DENIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 3
)

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "WRITE_ACCESS_DENIED",
}

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WRITE_ACCESS_DENIED": 3,
}

func (x ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) String() string {
	return proto.EnumName(ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, int32(x))
}

func (ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be307f994c248b0a, []int{0, 0}
}

// Container for enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAccessDeniedErrorEnum) Reset()         { *m = ResourceAccessDeniedErrorEnum{} }
func (m *ResourceAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ResourceAccessDeniedErrorEnum) ProtoMessage()    {}
func (*ResourceAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be307f994c248b0a, []int{0}
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.Merge(m, src)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Size(m)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ResourceAccessDeniedErrorEnum)(nil), "google.ads.googleads.v0.errors.ResourceAccessDeniedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError", ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/resource_access_denied_error.proto", fileDescriptor_be307f994c248b0a)
}

var fileDescriptor_be307f994c248b0a = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xa2, 0xd4, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0xf8, 0xc4, 0xe4,
	0xe4, 0xd4, 0xe2, 0xe2, 0xf8, 0x94, 0xd4, 0xbc, 0xcc, 0xd4, 0x94, 0x78, 0xb0, 0xac, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0x9f, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0xdc, 0x08, 0xbd,
	0x32, 0x03, 0x3d, 0x88, 0x11, 0x4a, 0xc5, 0x5c, 0xb2, 0x41, 0x50, 0x53, 0x1c, 0xc1, 0x86, 0xb8,
	0x80, 0xcd, 0x70, 0x05, 0xc9, 0xba, 0xe6, 0x95, 0xe6, 0x2a, 0x05, 0x71, 0x49, 0xe2, 0x54, 0x20,
	0xc4, 0xcf, 0xc5, 0x1d, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0,
	0x20, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x28, 0x24, 0xce,
	0x25, 0x1c, 0x1e, 0xe4, 0x19, 0xe2, 0x1a, 0xef, 0xe8, 0xec, 0xec, 0x1a, 0x1c, 0x1c, 0xef, 0xe2,
	0xea, 0x07, 0x52, 0xc5, 0xec, 0xb4, 0x9e, 0x91, 0x4b, 0x29, 0x39, 0x3f, 0x57, 0x0f, 0xbf, 0xdb,
	0x9c, 0xe4, 0x70, 0x5a, 0x1c, 0x00, 0xf2, 0x5b, 0x00, 0xe3, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7,
	0x55, 0x4c, 0x72, 0xee, 0x10, 0x83, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40,
	0x0f, 0xac, 0xb8, 0xf8, 0x14, 0x4c, 0x41, 0x8c, 0x63, 0x4a, 0x71, 0x0c, 0x5c, 0x41, 0x4c, 0x98,
	0x41, 0x0c, 0x44, 0xc1, 0x23, 0x42, 0x0a, 0x92, 0xd8, 0xc0, 0xa1, 0x69, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0x76, 0xbf, 0xc5, 0x92, 0x01, 0x00, 0x00,
}
