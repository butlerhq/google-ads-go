// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/campaign_budget_error.proto

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

// Enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum_CampaignBudgetError int32

const (
	// Enum unspecified.
	CampaignBudgetErrorEnum_UNSPECIFIED CampaignBudgetErrorEnum_CampaignBudgetError = 0
	// The received error code is not known in this version.
	CampaignBudgetErrorEnum_UNKNOWN CampaignBudgetErrorEnum_CampaignBudgetError = 1
	// The requested campaign budget no longer exists.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_REMOVED CampaignBudgetErrorEnum_CampaignBudgetError = 2
	// The campaign budget is associated with at least one campaign, and so the
	// campaign budget cannot be removed.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_IN_USE CampaignBudgetErrorEnum_CampaignBudgetError = 3
	// Customer is not whitelisted for this campaign budget period.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE CampaignBudgetErrorEnum_CampaignBudgetError = 4
	// This field is not mutable on implicitly shared campaign budgets
	CampaignBudgetErrorEnum_CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET CampaignBudgetErrorEnum_CampaignBudgetError = 6
	// Cannot change explicitly shared campaign budgets back to implicitly
	// shared ones.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 7
	// An implicit campaign budget without a name cannot be changed to
	// explicitly shared campaign budget.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 8
	// Cannot change an implicitly shared campaign budget to an explicitly
	// shared one.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 9
	// Only explicitly shared campaign budgets can be used with multiple
	// campaigns.
	CampaignBudgetErrorEnum_CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS CampaignBudgetErrorEnum_CampaignBudgetError = 10
	// A campaign budget with this name already exists.
	CampaignBudgetErrorEnum_DUPLICATE_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 11
	// A money amount was not in the expected currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_IN_WRONG_CURRENCY CampaignBudgetErrorEnum_CampaignBudgetError = 12
	// A money amount was less than the minimum CPC for currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC CampaignBudgetErrorEnum_CampaignBudgetError = 13
	// A money amount was greater than the maximum allowed.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_TOO_LARGE CampaignBudgetErrorEnum_CampaignBudgetError = 14
	// A money amount was negative.
	CampaignBudgetErrorEnum_NEGATIVE_MONEY_AMOUNT CampaignBudgetErrorEnum_CampaignBudgetError = 15
	// A money amount was not a multiple of a minimum unit.
	CampaignBudgetErrorEnum_NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT CampaignBudgetErrorEnum_CampaignBudgetError = 16
)

var CampaignBudgetErrorEnum_CampaignBudgetError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "CAMPAIGN_BUDGET_REMOVED",
	3:  "CAMPAIGN_BUDGET_IN_USE",
	4:  "CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE",
	6:  "CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET",
	7:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED",
	8:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME",
	9:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED",
	10: "CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS",
	11: "DUPLICATE_NAME",
	12: "MONEY_AMOUNT_IN_WRONG_CURRENCY",
	13: "MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC",
	14: "MONEY_AMOUNT_TOO_LARGE",
	15: "NEGATIVE_MONEY_AMOUNT",
	16: "NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT",
}

var CampaignBudgetErrorEnum_CampaignBudgetError_value = map[string]int32{
	"UNSPECIFIED":                          0,
	"UNKNOWN":                              1,
	"CAMPAIGN_BUDGET_REMOVED":              2,
	"CAMPAIGN_BUDGET_IN_USE":               3,
	"CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE": 4,
	"CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET":             6,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED":                   7,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME":      8,
	"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED":                   9,
	"CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS": 10,
	"DUPLICATE_NAME":                              11,
	"MONEY_AMOUNT_IN_WRONG_CURRENCY":              12,
	"MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC": 13,
	"MONEY_AMOUNT_TOO_LARGE":                      14,
	"NEGATIVE_MONEY_AMOUNT":                       15,
	"NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT":       16,
}

func (x CampaignBudgetErrorEnum_CampaignBudgetError) String() string {
	return proto.EnumName(CampaignBudgetErrorEnum_CampaignBudgetError_name, int32(x))
}

func (CampaignBudgetErrorEnum_CampaignBudgetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9300e3888a22b411, []int{0, 0}
}

// Container for enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignBudgetErrorEnum) Reset()         { *m = CampaignBudgetErrorEnum{} }
func (m *CampaignBudgetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignBudgetErrorEnum) ProtoMessage()    {}
func (*CampaignBudgetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9300e3888a22b411, []int{0}
}
func (m *CampaignBudgetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignBudgetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignBudgetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBudgetErrorEnum.Merge(m, src)
}
func (m *CampaignBudgetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignBudgetErrorEnum.Size(m)
}
func (m *CampaignBudgetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBudgetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBudgetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CampaignBudgetErrorEnum)(nil), "google.ads.googleads.v0.errors.CampaignBudgetErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CampaignBudgetErrorEnum_CampaignBudgetError", CampaignBudgetErrorEnum_CampaignBudgetError_name, CampaignBudgetErrorEnum_CampaignBudgetError_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/errors/campaign_budget_error.proto", fileDescriptor_9300e3888a22b411)
}

var fileDescriptor_9300e3888a22b411 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0x87, 0x4d, 0x6a, 0x5b, 0x9d, 0x68, 0x3b, 0x8c, 0x68, 0xfd, 0x03, 0xb9, 0x08, 0x0a, 0x8a,
	0xb0, 0x09, 0x0a, 0x22, 0x22, 0xc8, 0x64, 0xe7, 0x64, 0x33, 0xb8, 0x33, 0xb3, 0xec, 0xce, 0x24,
	0x06, 0x02, 0x87, 0xb4, 0x09, 0x41, 0x30, 0xdd, 0x92, 0x6d, 0xfb, 0x38, 0x5e, 0x78, 0xa7, 0x8f,
	0xe0, 0x23, 0xf8, 0x18, 0x3e, 0x89, 0xec, 0x6e, 0x93, 0xd8, 0xb4, 0x18, 0xef, 0x0e, 0x73, 0xbe,
	0xef, 0xfc, 0x66, 0x98, 0x19, 0xf2, 0x6e, 0x9a, 0xa6, 0xd3, 0x2f, 0x93, 0xe6, 0x68, 0x9c, 0x35,
	0xcb, 0x32, 0xaf, 0xce, 0x5b, 0xcd, 0xc9, 0x7c, 0x9e, 0xce, 0xb3, 0xe6, 0xd1, 0x68, 0x76, 0x32,
	0xfa, 0x3c, 0x3d, 0xc6, 0xc3, 0xb3, 0xf1, 0x74, 0x72, 0x8a, 0xc5, 0xb2, 0x77, 0x32, 0x4f, 0x4f,
	0x53, 0x56, 0x2f, 0x05, 0x6f, 0x34, 0xce, 0xbc, 0xa5, 0xeb, 0x9d, 0xb7, 0xbc, 0xd2, 0x6d, 0xfc,
	0xdc, 0x26, 0x07, 0xfe, 0x85, 0xdf, 0x2e, 0x74, 0xc8, 0x1b, 0x70, 0x7c, 0x36, 0x6b, 0x7c, 0xdd,
	0x26, 0xf7, 0xae, 0xe9, 0xb1, 0x7d, 0x52, 0x73, 0x3a, 0x89, 0xc0, 0x97, 0x1d, 0x09, 0x82, 0xde,
	0x60, 0x35, 0xb2, 0xeb, 0xf4, 0x47, 0x6d, 0xfa, 0x9a, 0x56, 0xd8, 0x13, 0x72, 0xe0, 0x73, 0x15,
	0x71, 0x19, 0x68, 0x6c, 0x3b, 0x11, 0x80, 0xc5, 0x18, 0x94, 0xe9, 0x81, 0xa0, 0x55, 0xf6, 0x98,
	0x3c, 0x58, 0x6f, 0x4a, 0x8d, 0x2e, 0x01, 0xba, 0xc5, 0x9e, 0x93, 0xa7, 0xeb, 0xbd, 0x08, 0x62,
	0x69, 0x04, 0x6a, 0x63, 0x91, 0xf7, 0xb8, 0x0c, 0x79, 0x3b, 0x04, 0x7a, 0x93, 0xbd, 0x27, 0x6f,
	0x7d, 0xae, 0xf3, 0x55, 0x65, 0x84, 0xec, 0x0c, 0xb0, 0x23, 0x21, 0x14, 0x68, 0x3a, 0x28, 0x55,
	0x14, 0x4a, 0x5f, 0xda, 0x70, 0x80, 0x49, 0x97, 0xc7, 0x20, 0x70, 0x6d, 0x24, 0xdd, 0x61, 0x6f,
	0xc8, 0xab, 0x0b, 0xdb, 0x45, 0x82, 0x5b, 0x58, 0x47, 0xd0, 0x9a, 0xab, 0x73, 0xe8, 0x2e, 0xf3,
	0xc9, 0x87, 0x8d, 0x1e, 0x7c, 0x5a, 0xcf, 0xef, 0x4b, 0xdb, 0x35, 0xce, 0xa2, 0xe6, 0x0a, 0xe8,
	0xad, 0xff, 0x0a, 0xbf, 0x32, 0x84, 0xde, 0x66, 0x5d, 0x22, 0x16, 0x5e, 0x02, 0x9b, 0x8f, 0x59,
	0xc4, 0xa2, 0x72, 0xa1, 0x95, 0x51, 0xb8, 0x0a, 0x49, 0x28, 0x61, 0x8c, 0xec, 0x09, 0x97, 0xeb,
	0x79, 0x7a, 0xb1, 0xab, 0x1a, 0x6b, 0x90, 0xba, 0x32, 0x1a, 0x06, 0xc8, 0x95, 0x71, 0xba, 0xb8,
	0x93, 0x7e, 0x6c, 0x74, 0x80, 0xbe, 0x8b, 0x63, 0xd0, 0xfe, 0x80, 0xde, 0x61, 0x4d, 0xf2, 0xf2,
	0x12, 0x13, 0x42, 0x92, 0xa0, 0xed, 0x72, 0xbd, 0x84, 0x50, 0x49, 0x2d, 0x95, 0x53, 0xe8, 0x47,
	0x3e, 0xbd, 0x9b, 0xdf, 0xf5, 0x25, 0xc1, 0x1a, 0x83, 0x21, 0x8f, 0x03, 0xa0, 0x7b, 0xec, 0x11,
	0xb9, 0xaf, 0x21, 0xe0, 0x56, 0xf6, 0x00, 0xff, 0x86, 0xe8, 0x3e, 0x7b, 0x41, 0x9e, 0x69, 0xa3,
	0x57, 0x7b, 0x37, 0x9d, 0xd5, 0xdc, 0x45, 0x90, 0xd3, 0xd2, 0x52, 0xda, 0xfe, 0x5e, 0x21, 0x8d,
	0xa3, 0x74, 0xe6, 0xfd, 0xfb, 0x8d, 0xb7, 0x1f, 0x5e, 0xf3, 0x88, 0xa3, 0xfc, 0x77, 0x44, 0x95,
	0x6f, 0xd5, 0xad, 0x80, 0xf3, 0x1f, 0xd5, 0x7a, 0x50, 0x8e, 0xe0, 0xe3, 0xcc, 0x2b, 0xcb, 0xbc,
	0xea, 0xb5, 0xbc, 0x02, 0xce, 0x7e, 0x2d, 0x80, 0x21, 0x1f, 0x67, 0xc3, 0x25, 0x30, 0xec, 0xb5,
	0x86, 0x25, 0xf0, 0x7b, 0x13, 0x70, 0xb8, 0x53, 0xfc, 0xc7, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xea, 0x11, 0xc1, 0xe5, 0xcd, 0x03, 0x00, 0x00,
}
