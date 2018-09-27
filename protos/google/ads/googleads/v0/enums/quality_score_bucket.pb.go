// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/quality_score_bucket.proto

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

// Enum listing the possible quality score buckets.
type QualityScoreBucketEnum_QualityScoreBucket int32

const (
	// Not specified.
	QualityScoreBucketEnum_UNSPECIFIED QualityScoreBucketEnum_QualityScoreBucket = 0
	// Used for return value only. Represents value unknown in this version.
	QualityScoreBucketEnum_UNKNOWN QualityScoreBucketEnum_QualityScoreBucket = 1
	// Quality of the creative is below average.
	QualityScoreBucketEnum_BELOW_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 2
	// Quality of the creative is average.
	QualityScoreBucketEnum_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 3
	// Quality of the creative is above average.
	QualityScoreBucketEnum_ABOVE_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 4
)

var QualityScoreBucketEnum_QualityScoreBucket_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BELOW_AVERAGE",
	3: "AVERAGE",
	4: "ABOVE_AVERAGE",
}

var QualityScoreBucketEnum_QualityScoreBucket_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"BELOW_AVERAGE": 2,
	"AVERAGE":       3,
	"ABOVE_AVERAGE": 4,
}

func (x QualityScoreBucketEnum_QualityScoreBucket) String() string {
	return proto.EnumName(QualityScoreBucketEnum_QualityScoreBucket_name, int32(x))
}

func (QualityScoreBucketEnum_QualityScoreBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cc972712a91136d, []int{0, 0}
}

// The relative performance compared to other advertisers.
type QualityScoreBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QualityScoreBucketEnum) Reset()         { *m = QualityScoreBucketEnum{} }
func (m *QualityScoreBucketEnum) String() string { return proto.CompactTextString(m) }
func (*QualityScoreBucketEnum) ProtoMessage()    {}
func (*QualityScoreBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc972712a91136d, []int{0}
}
func (m *QualityScoreBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityScoreBucketEnum.Unmarshal(m, b)
}
func (m *QualityScoreBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityScoreBucketEnum.Marshal(b, m, deterministic)
}
func (m *QualityScoreBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityScoreBucketEnum.Merge(m, src)
}
func (m *QualityScoreBucketEnum) XXX_Size() int {
	return xxx_messageInfo_QualityScoreBucketEnum.Size(m)
}
func (m *QualityScoreBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityScoreBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QualityScoreBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QualityScoreBucketEnum)(nil), "google.ads.googleads.v0.enums.QualityScoreBucketEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.QualityScoreBucketEnum_QualityScoreBucket", QualityScoreBucketEnum_QualityScoreBucket_name, QualityScoreBucketEnum_QualityScoreBucket_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/enums/quality_score_bucket.proto", fileDescriptor_0cc972712a91136d)
}

var fileDescriptor_0cc972712a91136d = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xc2, 0xd2, 0xc4, 0x9c, 0xcc, 0x92, 0xca, 0xf8, 0xe2, 0xe4, 0xfc, 0xa2,
	0xd4, 0xf8, 0xa4, 0xd2, 0xe4, 0xec, 0xd4, 0x12, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59,
	0x88, 0x72, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x4e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x4e, 0xa5,
	0x7a, 0x2e, 0xb1, 0x40, 0x88, 0xe6, 0x60, 0x90, 0x5e, 0x27, 0xb0, 0x56, 0xd7, 0xbc, 0xd2, 0x5c,
	0xa5, 0x54, 0x2e, 0x21, 0x4c, 0x19, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67,
	0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff,
	0x70, 0x3f, 0x01, 0x46, 0x21, 0x41, 0x2e, 0x5e, 0x27, 0x57, 0x1f, 0xff, 0xf0, 0x78, 0xc7, 0x30,
	0xd7, 0x20, 0x47, 0x77, 0x57, 0x01, 0x26, 0x90, 0x3c, 0x8c, 0xc3, 0x0c, 0x92, 0x77, 0x74, 0xf2,
	0x0f, 0x73, 0x85, 0xcb, 0xb3, 0x38, 0x2d, 0x61, 0xe4, 0x52, 0x4c, 0xce, 0xcf, 0xd5, 0xc3, 0xeb,
	0x4c, 0x27, 0x71, 0x4c, 0xa7, 0x04, 0x80, 0xbc, 0x17, 0xc0, 0xb8, 0x88, 0x89, 0xd9, 0xdd, 0xd1,
	0x71, 0x15, 0x93, 0xac, 0x3b, 0xc4, 0x00, 0xc7, 0x94, 0x62, 0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33,
	0xd0, 0x03, 0x79, 0xa6, 0xf8, 0x14, 0x4c, 0x3e, 0xc6, 0x31, 0xa5, 0x38, 0x06, 0x2e, 0x1f, 0x13,
	0x66, 0x10, 0x03, 0x96, 0x7f, 0x44, 0x40, 0x3e, 0x89, 0x0d, 0x1c, 0x9a, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x67, 0x76, 0x39, 0x7c, 0x89, 0x01, 0x00, 0x00,
}
