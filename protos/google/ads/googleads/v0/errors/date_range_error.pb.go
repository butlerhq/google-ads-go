// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/date_range_error.proto

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

// Enum describing possible date range errors.
type DateRangeErrorEnum_DateRangeError int32

const (
	// Enum unspecified.
	DateRangeErrorEnum_UNSPECIFIED DateRangeErrorEnum_DateRangeError = 0
	// The received error code is not known in this version.
	DateRangeErrorEnum_UNKNOWN DateRangeErrorEnum_DateRangeError = 1
	// Invalid date.
	DateRangeErrorEnum_INVALID_DATE DateRangeErrorEnum_DateRangeError = 2
	// The start date was after the end date.
	DateRangeErrorEnum_START_DATE_AFTER_END_DATE DateRangeErrorEnum_DateRangeError = 3
	// Cannot set date to past time
	DateRangeErrorEnum_CANNOT_SET_DATE_TO_PAST DateRangeErrorEnum_DateRangeError = 4
	// A date was used that is past the system "last" date.
	DateRangeErrorEnum_AFTER_MAXIMUM_ALLOWABLE_DATE DateRangeErrorEnum_DateRangeError = 5
	// Trying to change start date on a campaign that has started.
	DateRangeErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED DateRangeErrorEnum_DateRangeError = 6
)

var DateRangeErrorEnum_DateRangeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_DATE",
	3: "START_DATE_AFTER_END_DATE",
	4: "CANNOT_SET_DATE_TO_PAST",
	5: "AFTER_MAXIMUM_ALLOWABLE_DATE",
	6: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
}

var DateRangeErrorEnum_DateRangeError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"INVALID_DATE":                 2,
	"START_DATE_AFTER_END_DATE":    3,
	"CANNOT_SET_DATE_TO_PAST":      4,
	"AFTER_MAXIMUM_ALLOWABLE_DATE": 5,
	"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED": 6,
}

func (x DateRangeErrorEnum_DateRangeError) String() string {
	return proto.EnumName(DateRangeErrorEnum_DateRangeError_name, int32(x))
}

func (DateRangeErrorEnum_DateRangeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2dc01d50b2259f80, []int{0, 0}
}

// Container for enum describing possible date range errors.
type DateRangeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateRangeErrorEnum) Reset()         { *m = DateRangeErrorEnum{} }
func (m *DateRangeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DateRangeErrorEnum) ProtoMessage()    {}
func (*DateRangeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc01d50b2259f80, []int{0}
}
func (m *DateRangeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRangeErrorEnum.Unmarshal(m, b)
}
func (m *DateRangeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRangeErrorEnum.Marshal(b, m, deterministic)
}
func (m *DateRangeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRangeErrorEnum.Merge(m, src)
}
func (m *DateRangeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DateRangeErrorEnum.Size(m)
}
func (m *DateRangeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRangeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DateRangeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DateRangeErrorEnum)(nil), "google.ads.googleads.v0.errors.DateRangeErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.DateRangeErrorEnum_DateRangeError", DateRangeErrorEnum_DateRangeError_name, DateRangeErrorEnum_DateRangeError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/date_range_error.proto", fileDescriptor_2dc01d50b2259f80)
}

var fileDescriptor_2dc01d50b2259f80 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0x86, 0xff, 0x6e, 0xbf, 0x13, 0x32, 0xd1, 0x12, 0x0f, 0x44, 0xd4, 0x21, 0x3b, 0x15, 0xd2,
	0x82, 0x78, 0x01, 0xdf, 0x96, 0x74, 0x04, 0xdb, 0xb4, 0xb4, 0xd9, 0xa6, 0x50, 0x08, 0xd5, 0x96,
	0x9d, 0xb8, 0x55, 0xda, 0xba, 0x0b, 0x12, 0x8f, 0xbc, 0x14, 0x4f, 0xbc, 0x07, 0x0f, 0xbc, 0x0e,
	0x69, 0x53, 0x87, 0x3b, 0xd1, 0xb3, 0x37, 0x6f, 0x9e, 0x87, 0x90, 0xef, 0x43, 0x57, 0x8b, 0x3c,
	0x5f, 0x3c, 0x64, 0x56, 0x92, 0x96, 0x96, 0x8e, 0x75, 0x5a, 0xdb, 0x56, 0x56, 0x14, 0x79, 0x51,
	0x5a, 0x69, 0x52, 0x65, 0xaa, 0x48, 0x56, 0x8b, 0x4c, 0x35, 0x0d, 0x79, 0x2c, 0xf2, 0x2a, 0xc7,
	0x03, 0xcd, 0x92, 0x24, 0x2d, 0xc9, 0x46, 0x23, 0x6b, 0x9b, 0x68, 0x6d, 0xf8, 0x69, 0x20, 0x4c,
	0x93, 0x2a, 0x0b, 0x6b, 0x93, 0xd5, 0x1d, 0x5b, 0x3d, 0x2d, 0x87, 0xef, 0x06, 0xda, 0xdf, 0xae,
	0xf1, 0x01, 0xea, 0x4f, 0x45, 0x14, 0xb0, 0x31, 0x77, 0x38, 0xa3, 0xe6, 0x3f, 0xdc, 0x47, 0xbb,
	0x53, 0x71, 0x2d, 0xfc, 0xb9, 0x30, 0x0d, 0x6c, 0xa2, 0x3d, 0x2e, 0x66, 0xe0, 0x72, 0xaa, 0x28,
	0x48, 0x66, 0x76, 0xf0, 0x19, 0x3a, 0x8e, 0x24, 0x84, 0xb2, 0x39, 0x2b, 0x70, 0x24, 0x0b, 0x15,
	0x13, 0xed, 0x75, 0x17, 0x9f, 0xa0, 0xa3, 0x31, 0x08, 0xe1, 0x4b, 0x15, 0xb1, 0x96, 0x91, 0xbe,
	0x0a, 0x20, 0x92, 0xe6, 0x7f, 0x7c, 0x8e, 0x4e, 0xb5, 0xe0, 0xc1, 0x0d, 0xf7, 0xa6, 0x9e, 0x02,
	0xd7, 0xf5, 0xe7, 0x30, 0x72, 0x99, 0xd6, 0x77, 0xb0, 0x85, 0x2e, 0x5a, 0xdd, 0xf3, 0x29, 0x77,
	0x6e, 0xd5, 0x8f, 0xb7, 0xb8, 0xa3, 0xc0, 0x0d, 0x19, 0xd0, 0xb6, 0x65, 0xd4, 0xec, 0x8d, 0x5e,
	0x0c, 0x34, 0xbc, 0xcf, 0x97, 0xe4, 0xf7, 0x79, 0x8c, 0x0e, 0xb7, 0x7f, 0x1d, 0xd4, 0x43, 0x0c,
	0x8c, 0xe7, 0x4e, 0x77, 0x02, 0xf0, 0xda, 0x19, 0x4c, 0xb4, 0x0d, 0x69, 0x49, 0x74, 0xac, 0xd3,
	0xcc, 0x26, 0x0d, 0x5c, 0xbe, 0x7d, 0x03, 0x31, 0xa4, 0x65, 0xbc, 0x01, 0xe2, 0x99, 0x1d, 0x6b,
	0xe0, 0xe3, 0x2f, 0xe0, 0xae, 0xd7, 0xac, 0xed, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xed, 0xd1,
	0xa6, 0x1d, 0xef, 0x01, 0x00, 0x00,
}
