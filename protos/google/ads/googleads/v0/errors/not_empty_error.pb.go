// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/not_empty_error.proto

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

// Enum describing possible not empty errors.
type NotEmptyErrorEnum_NotEmptyError int32

const (
	// Enum unspecified.
	NotEmptyErrorEnum_UNSPECIFIED NotEmptyErrorEnum_NotEmptyError = 0
	// The received error code is not known in this version.
	NotEmptyErrorEnum_UNKNOWN NotEmptyErrorEnum_NotEmptyError = 1
	// Empty list.
	NotEmptyErrorEnum_EMPTY_LIST NotEmptyErrorEnum_NotEmptyError = 2
)

var NotEmptyErrorEnum_NotEmptyError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EMPTY_LIST",
}

var NotEmptyErrorEnum_NotEmptyError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EMPTY_LIST":  2,
}

func (x NotEmptyErrorEnum_NotEmptyError) String() string {
	return proto.EnumName(NotEmptyErrorEnum_NotEmptyError_name, int32(x))
}

func (NotEmptyErrorEnum_NotEmptyError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_221b4006f4ee82bb, []int{0, 0}
}

// Container for enum describing possible not empty errors.
type NotEmptyErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotEmptyErrorEnum) Reset()         { *m = NotEmptyErrorEnum{} }
func (m *NotEmptyErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NotEmptyErrorEnum) ProtoMessage()    {}
func (*NotEmptyErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_221b4006f4ee82bb, []int{0}
}
func (m *NotEmptyErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotEmptyErrorEnum.Unmarshal(m, b)
}
func (m *NotEmptyErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotEmptyErrorEnum.Marshal(b, m, deterministic)
}
func (m *NotEmptyErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotEmptyErrorEnum.Merge(m, src)
}
func (m *NotEmptyErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NotEmptyErrorEnum.Size(m)
}
func (m *NotEmptyErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NotEmptyErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NotEmptyErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NotEmptyErrorEnum)(nil), "google.ads.googleads.v0.errors.NotEmptyErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.NotEmptyErrorEnum_NotEmptyError", NotEmptyErrorEnum_NotEmptyError_name, NotEmptyErrorEnum_NotEmptyError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/not_empty_error.proto", fileDescriptor_221b4006f4ee82bb)
}

var fileDescriptor_221b4006f4ee82bb = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xbc, 0xfc, 0x92, 0xf8, 0xd4, 0xdc, 0x82, 0x92, 0xca, 0x78, 0xb0,
	0x80, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0xa9, 0x5e, 0x62, 0x4a, 0xb1, 0x1e,
	0x5c, 0x97, 0x5e, 0x99, 0x81, 0x1e, 0x44, 0x97, 0x52, 0x10, 0x97, 0xa0, 0x5f, 0x7e, 0x89, 0x2b,
	0x48, 0x9f, 0x2b, 0x48, 0xc4, 0x35, 0xaf, 0x34, 0x57, 0xc9, 0x96, 0x8b, 0x17, 0x45, 0x50, 0x88,
	0x9f, 0x8b, 0x3b, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41,
	0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80, 0x51, 0x88, 0x8f, 0x8b,
	0xcb, 0xd5, 0x37, 0x20, 0x24, 0x32, 0xde, 0xc7, 0x33, 0x38, 0x44, 0x80, 0xc9, 0x69, 0x31, 0x23,
	0x97, 0x52, 0x72, 0x7e, 0xae, 0x1e, 0x7e, 0xab, 0x9d, 0x84, 0x50, 0xec, 0x08, 0x00, 0x39, 0x37,
	0x80, 0x71, 0x11, 0x13, 0xb3, 0xbb, 0xa3, 0xe3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0x66, 0xc7, 0x94,
	0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03, 0x2b, 0x2e, 0x3e, 0x05, 0x53, 0x10, 0xe3,
	0x98, 0x52, 0x1c, 0x03, 0x57, 0x10, 0x13, 0x66, 0x10, 0x03, 0x51, 0xf0, 0x88, 0x90, 0x82, 0x24,
	0x36, 0x70, 0x00, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xf5, 0x30, 0x78, 0x58, 0x01,
	0x00, 0x00,
}
