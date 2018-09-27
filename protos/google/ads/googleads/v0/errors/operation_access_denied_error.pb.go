// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/operation_access_denied_error.proto

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

// Enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum_OperationAccessDeniedError int32

const (
	// Enum unspecified.
	OperationAccessDeniedErrorEnum_UNSPECIFIED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 0
	// The received error code is not known in this version.
	OperationAccessDeniedErrorEnum_UNKNOWN OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 1
	// Unauthorized invocation of a service's method (get, mutate, etc.)
	OperationAccessDeniedErrorEnum_ACTION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 2
	// Unauthorized CREATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_CREATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 3
	// Unauthorized REMOVE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_REMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 4
	// Unauthorized UPDATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_UPDATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 5
	// A mutate action is not allowed on this campaign, from this client.
	OperationAccessDeniedErrorEnum_MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 6
	// This operation is not permitted on this campaign type
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 7
	// A CREATE operation may not set status to REMOVED.
	OperationAccessDeniedErrorEnum_CREATE_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 8
	// This operation is not allowed because the campaign or adgroup is removed.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 9
	// This operation is not permitted on this ad group type.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 10
)

var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ACTION_NOT_PERMITTED",
	3:  "CREATE_OPERATION_NOT_PERMITTED",
	4:  "REMOVE_OPERATION_NOT_PERMITTED",
	5:  "UPDATE_OPERATION_NOT_PERMITTED",
	6:  "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT",
	7:  "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE",
	8:  "CREATE_AS_REMOVED_NOT_PERMITTED",
	9:  "OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE",
	10: "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE",
}

var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":                                  0,
	"UNKNOWN":                                      1,
	"ACTION_NOT_PERMITTED":                         2,
	"CREATE_OPERATION_NOT_PERMITTED":               3,
	"REMOVE_OPERATION_NOT_PERMITTED":               4,
	"UPDATE_OPERATION_NOT_PERMITTED":               5,
	"MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT":       6,
	"OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE":    7,
	"CREATE_AS_REMOVED_NOT_PERMITTED":              8,
	"OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE": 9,
	"OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE":    10,
}

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) String() string {
	return proto.EnumName(OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, int32(x))
}

func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb3b0aac42536863, []int{0, 0}
}

// Container for enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationAccessDeniedErrorEnum) Reset()         { *m = OperationAccessDeniedErrorEnum{} }
func (m *OperationAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*OperationAccessDeniedErrorEnum) ProtoMessage()    {}
func (*OperationAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb3b0aac42536863, []int{0}
}
func (m *OperationAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.Merge(m, src)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Size(m)
}
func (m *OperationAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperationAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperationAccessDeniedErrorEnum)(nil), "google.ads.googleads.v0.errors.OperationAccessDeniedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.OperationAccessDeniedErrorEnum_OperationAccessDeniedError", OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/operation_access_denied_error.proto", fileDescriptor_fb3b0aac42536863)
}

var fileDescriptor_fb3b0aac42536863 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x18, 0x86, 0xdd, 0xaa, 0x9b, 0x66, 0x07, 0x96, 0xe0, 0x81, 0x78, 0xd0, 0x41, 0x05, 0x41, 0xd1,
	0xb4, 0xe0, 0x15, 0x64, 0x6d, 0x56, 0x8a, 0x36, 0x09, 0x59, 0x3a, 0x11, 0x0a, 0xa1, 0xae, 0x65,
	0x0c, 0xdc, 0x32, 0x9a, 0xb9, 0xfb, 0xd1, 0x43, 0x8f, 0xbc, 0x08, 0x8f, 0xbc, 0x0c, 0xaf, 0xe4,
	0xa7, 0xcd, 0xbf, 0x1d, 0xfc, 0xac, 0xdb, 0xd9, 0x0b, 0x79, 0xbe, 0xa7, 0x6f, 0xf9, 0x3e, 0x30,
	0x5b, 0x6b, 0xbd, 0xfe, 0x5e, 0x07, 0x65, 0x65, 0x02, 0x1b, 0xdb, 0x74, 0x0c, 0x83, 0xba, 0x69,
	0x74, 0x63, 0x02, 0xbd, 0xaf, 0x9b, 0xf2, 0xb0, 0xd1, 0x3b, 0x55, 0xae, 0x56, 0xb5, 0x31, 0xaa,
	0xaa, 0x77, 0x9b, 0xba, 0x52, 0xdd, 0x33, 0xda, 0x37, 0xfa, 0xa0, 0xa1, 0x67, 0x07, 0x51, 0x59,
	0x19, 0x74, 0x76, 0xa0, 0x63, 0x88, 0xac, 0xc3, 0xff, 0xeb, 0x00, 0x8f, 0x9d, 0x3c, 0xb8, 0xd3,
	0xc4, 0x9d, 0x85, 0xb4, 0xef, 0x64, 0xf7, 0x63, 0xeb, 0xff, 0x74, 0xc0, 0xab, 0x7e, 0x04, 0x3e,
	0x07, 0x93, 0x9c, 0x2e, 0x38, 0x89, 0xd2, 0x79, 0x4a, 0x62, 0xf7, 0x11, 0x9c, 0x80, 0x71, 0x4e,
	0x3f, 0x51, 0xf6, 0x85, 0xba, 0x03, 0xf8, 0x12, 0xbc, 0xc0, 0x91, 0x4c, 0x19, 0x55, 0x94, 0x49,
	0xc5, 0x89, 0xc8, 0x52, 0x29, 0x49, 0xec, 0x0e, 0xa1, 0x0f, 0xbc, 0x48, 0x10, 0x2c, 0x89, 0x62,
	0x9c, 0x08, 0x7c, 0x81, 0x71, 0x5a, 0x46, 0x90, 0x8c, 0x2d, 0xfb, 0x99, 0xc7, 0x2d, 0x93, 0xf3,
	0xf8, 0x9a, 0xe7, 0x09, 0x7c, 0x07, 0xde, 0x64, 0xb9, 0x6c, 0x99, 0x4b, 0x65, 0xd4, 0x9c, 0x09,
	0x15, 0x7d, 0x4e, 0x09, 0x95, 0xee, 0x08, 0x7e, 0x00, 0x6f, 0x7b, 0x44, 0x96, 0xc3, 0x19, 0xc7,
	0x69, 0x42, 0x95, 0xfc, 0xca, 0x89, 0x3b, 0x86, 0xaf, 0xc1, 0xf4, 0xfe, 0x37, 0xf0, 0x42, 0xd9,
	0xb2, 0xf1, 0x83, 0xef, 0x3f, 0x85, 0x21, 0x78, 0x7f, 0xcd, 0x79, 0x1a, 0x13, 0x64, 0xc1, 0x72,
	0x11, 0x11, 0xf7, 0xd9, 0xad, 0x16, 0x38, 0x56, 0x89, 0x60, 0x39, 0xb7, 0x2d, 0xc0, 0xec, 0xcf,
	0x00, 0xf8, 0x2b, 0xbd, 0x45, 0xd7, 0xb7, 0x3d, 0x9b, 0xf6, 0xef, 0x91, 0xb7, 0xe7, 0xc2, 0x07,
	0xbf, 0x86, 0x4e, 0x82, 0xf1, 0xef, 0xa1, 0x97, 0x58, 0x13, 0xae, 0x0c, 0xb2, 0xb1, 0x4d, 0xcb,
	0x10, 0x75, 0xb0, 0xf9, 0x77, 0x02, 0x0a, 0x5c, 0x99, 0xe2, 0x0c, 0x14, 0xcb, 0xb0, 0xb0, 0xc0,
	0xff, 0x5b, 0xc0, 0xb7, 0x51, 0x77, 0xa0, 0x1f, 0xef, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xd2,
	0x35, 0x89, 0xe6, 0x02, 0x00, 0x00,
}
