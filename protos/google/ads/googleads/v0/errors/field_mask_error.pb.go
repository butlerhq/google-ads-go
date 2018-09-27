// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/field_mask_error.proto

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

// Enum describing possible field mask errors.
type FieldMaskErrorEnum_FieldMaskError int32

const (
	// Enum unspecified.
	FieldMaskErrorEnum_UNSPECIFIED FieldMaskErrorEnum_FieldMaskError = 0
	// The received error code is not known in this version.
	FieldMaskErrorEnum_UNKNOWN FieldMaskErrorEnum_FieldMaskError = 1
	// The field mask must be provided for update operations.
	FieldMaskErrorEnum_FIELD_MASK_MISSING FieldMaskErrorEnum_FieldMaskError = 5
	// The field mask must be empty for create and remove operations.
	FieldMaskErrorEnum_FIELD_MASK_NOT_ALLOWED FieldMaskErrorEnum_FieldMaskError = 4
	// The field mask contained an invalid field.
	FieldMaskErrorEnum_FIELD_NOT_FOUND FieldMaskErrorEnum_FieldMaskError = 2
	// The field mask updated a field with subfields. Fields with subfields may
	// be cleared, but not updated. To fix this, the field mask should select
	// all the subfields of the invalid field.
	FieldMaskErrorEnum_FIELD_HAS_SUBFIELDS FieldMaskErrorEnum_FieldMaskError = 3
)

var FieldMaskErrorEnum_FieldMaskError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "FIELD_MASK_MISSING",
	4: "FIELD_MASK_NOT_ALLOWED",
	2: "FIELD_NOT_FOUND",
	3: "FIELD_HAS_SUBFIELDS",
}

var FieldMaskErrorEnum_FieldMaskError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"FIELD_MASK_MISSING":     5,
	"FIELD_MASK_NOT_ALLOWED": 4,
	"FIELD_NOT_FOUND":        2,
	"FIELD_HAS_SUBFIELDS":    3,
}

func (x FieldMaskErrorEnum_FieldMaskError) String() string {
	return proto.EnumName(FieldMaskErrorEnum_FieldMaskError_name, int32(x))
}

func (FieldMaskErrorEnum_FieldMaskError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72820e40b31f95b7, []int{0, 0}
}

// Container for enum describing possible field mask errors.
type FieldMaskErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldMaskErrorEnum) Reset()         { *m = FieldMaskErrorEnum{} }
func (m *FieldMaskErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FieldMaskErrorEnum) ProtoMessage()    {}
func (*FieldMaskErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_72820e40b31f95b7, []int{0}
}
func (m *FieldMaskErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMaskErrorEnum.Unmarshal(m, b)
}
func (m *FieldMaskErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMaskErrorEnum.Marshal(b, m, deterministic)
}
func (m *FieldMaskErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMaskErrorEnum.Merge(m, src)
}
func (m *FieldMaskErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FieldMaskErrorEnum.Size(m)
}
func (m *FieldMaskErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMaskErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMaskErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FieldMaskErrorEnum)(nil), "google.ads.googleads.v0.errors.FieldMaskErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.FieldMaskErrorEnum_FieldMaskError", FieldMaskErrorEnum_FieldMaskError_name, FieldMaskErrorEnum_FieldMaskError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/field_mask_error.proto", fileDescriptor_72820e40b31f95b7)
}

var fileDescriptor_72820e40b31f95b7 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x51, 0x4a, 0xfb, 0x30,
	0x00, 0xc6, 0xff, 0xed, 0xfe, 0x2a, 0x64, 0xe0, 0x42, 0x0a, 0x13, 0x7c, 0xd8, 0x43, 0x0f, 0x90,
	0x16, 0xc4, 0x03, 0xa4, 0x36, 0xad, 0x65, 0x6d, 0x5a, 0x88, 0xdd, 0x5e, 0x0a, 0xa1, 0xda, 0x3a,
	0x64, 0xab, 0x91, 0x46, 0x77, 0x0e, 0xcf, 0x20, 0x82, 0xe0, 0x51, 0x3c, 0x86, 0x27, 0x91, 0x34,
	0x38, 0xdc, 0x8b, 0xbe, 0x7d, 0xf9, 0xe5, 0xf7, 0x41, 0xf2, 0x81, 0xf3, 0x95, 0x94, 0xab, 0x4d,
	0xeb, 0xd5, 0x8d, 0xf2, 0x4c, 0xd4, 0x69, 0xeb, 0x7b, 0x6d, 0xdf, 0xcb, 0x5e, 0x79, 0xb7, 0x77,
	0xed, 0xa6, 0x11, 0x5d, 0xad, 0xd6, 0x62, 0x20, 0xf8, 0xa1, 0x97, 0x8f, 0x12, 0xcd, 0x8c, 0x8b,
	0xeb, 0x46, 0xe1, 0x5d, 0x0d, 0x6f, 0x7d, 0x6c, 0x6a, 0xee, 0x9b, 0x05, 0x50, 0xa4, 0xab, 0x59,
	0xad, 0xd6, 0x54, 0x33, 0x7a, 0xff, 0xd4, 0xb9, 0xcf, 0x16, 0x38, 0xde, 0xc7, 0x68, 0x02, 0xc6,
	0x25, 0xe3, 0x05, 0xbd, 0x48, 0xa2, 0x84, 0x86, 0xf0, 0x1f, 0x1a, 0x83, 0xa3, 0x92, 0xcd, 0x59,
	0xbe, 0x64, 0xd0, 0x42, 0x53, 0x80, 0xa2, 0x84, 0xa6, 0xa1, 0xc8, 0x08, 0x9f, 0x8b, 0x2c, 0xe1,
	0x3c, 0x61, 0x31, 0x3c, 0x40, 0xa7, 0x60, 0xfa, 0x83, 0xb3, 0xfc, 0x4a, 0x90, 0x34, 0xcd, 0x97,
	0x34, 0x84, 0xff, 0x91, 0x03, 0x26, 0xe6, 0x4e, 0xe3, 0x28, 0x2f, 0x59, 0x08, 0x6d, 0x74, 0x02,
	0x1c, 0x03, 0x2f, 0x09, 0x17, 0xbc, 0x0c, 0x86, 0x03, 0x87, 0xa3, 0xe0, 0xd5, 0x02, 0xee, 0x8d,
	0xec, 0xf0, 0xef, 0x1f, 0x0a, 0x9c, 0xfd, 0x67, 0x17, 0x7a, 0x85, 0xc2, 0x7a, 0xb1, 0x47, 0x31,
	0x21, 0xef, 0xf6, 0x2c, 0x36, 0x6d, 0xd2, 0x28, 0x6c, 0xa2, 0x4e, 0x0b, 0x1f, 0x0f, 0xb2, 0xfa,
	0xf8, 0x16, 0x2a, 0xd2, 0xa8, 0x6a, 0x27, 0x54, 0x0b, 0xbf, 0x32, 0xc2, 0xe7, 0x5f, 0xc2, 0xf5,
	0xe1, 0xb0, 0xfb, 0xd9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xd1, 0x24, 0x06, 0xb0, 0x01,
	0x00, 0x00,
}
