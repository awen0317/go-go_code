// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/placeholder_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Possible placeholder types for a feed mapping.
type PlaceholderTypeEnum_PlaceholderType int32

const (
	// Not specified.
	PlaceholderTypeEnum_UNSPECIFIED PlaceholderTypeEnum_PlaceholderType = 0
	// Used for return value only. Represents value unknown in this version.
	PlaceholderTypeEnum_UNKNOWN PlaceholderTypeEnum_PlaceholderType = 1
	// Lets you show links in your ad to pages from your website, including the
	// main landing page.
	PlaceholderTypeEnum_SITELINK PlaceholderTypeEnum_PlaceholderType = 2
	// Lets you attach a phone number to an ad, allowing customers to call
	// directly from the ad.
	PlaceholderTypeEnum_CALL PlaceholderTypeEnum_PlaceholderType = 3
	// Lets you provide users with a link that points to a mobile app in
	// addition to a website.
	PlaceholderTypeEnum_APP PlaceholderTypeEnum_PlaceholderType = 4
	// Lets you show locations of businesses from your Google My Business
	// account in your ad. This helps people find your locations by showing your
	// ads with your address, a map to your location, or the distance to your
	// business. This extension type is useful to draw customers to your
	// brick-and-mortar location.
	PlaceholderTypeEnum_LOCATION PlaceholderTypeEnum_PlaceholderType = 5
	// If you sell your product through retail chains, affiliate location
	// extensions let you show nearby stores that carry your products.
	PlaceholderTypeEnum_AFFILIATE_LOCATION PlaceholderTypeEnum_PlaceholderType = 6
	// Lets you include additional text with your search ads that provide
	// detailed information about your business, including products and services
	// you offer. Callouts appear in ads at the top and bottom of Google search
	// results.
	PlaceholderTypeEnum_CALLOUT PlaceholderTypeEnum_PlaceholderType = 7
	// Lets you add more info to your ad, specific to some predefined categories
	// such as types, brands, styles, etc. A minimum of 3 text (SNIPPETS) values
	// are required.
	PlaceholderTypeEnum_STRUCTURED_SNIPPET PlaceholderTypeEnum_PlaceholderType = 8
	// Allows users to see your ad, click an icon, and contact you directly by
	// text message. With one tap on your ad, people can contact you to book an
	// appointment, get a quote, ask for information, or request a service.
	PlaceholderTypeEnum_MESSAGE PlaceholderTypeEnum_PlaceholderType = 9
	// Lets you display prices for a list of items along with your ads. A price
	// feed is composed of three to eight price table rows.
	PlaceholderTypeEnum_PRICE PlaceholderTypeEnum_PlaceholderType = 10
	// Allows you to highlight sales and other promotions that let users see how
	// they can save by buying now.
	PlaceholderTypeEnum_PROMOTION PlaceholderTypeEnum_PlaceholderType = 11
	// Lets you dynamically inject custom data into the title and description
	// of your ads.
	PlaceholderTypeEnum_AD_CUSTOMIZER PlaceholderTypeEnum_PlaceholderType = 12
	// Indicates that this feed is for education dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_EDUCATION PlaceholderTypeEnum_PlaceholderType = 13
	// Indicates that this feed is for flight dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_FLIGHT PlaceholderTypeEnum_PlaceholderType = 14
	// Indicates that this feed is for a custom dynamic remarketing type. Use
	// this only if the other business types don't apply to your products or
	// services.
	PlaceholderTypeEnum_DYNAMIC_CUSTOM PlaceholderTypeEnum_PlaceholderType = 15
	// Indicates that this feed is for hotels and rentals dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_HOTEL PlaceholderTypeEnum_PlaceholderType = 16
	// Indicates that this feed is for real estate dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_REAL_ESTATE PlaceholderTypeEnum_PlaceholderType = 17
	// Indicates that this feed is for travel dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_TRAVEL PlaceholderTypeEnum_PlaceholderType = 18
	// Indicates that this feed is for local deals dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_LOCAL PlaceholderTypeEnum_PlaceholderType = 19
	// Indicates that this feed is for job dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_JOB PlaceholderTypeEnum_PlaceholderType = 20
)

var PlaceholderTypeEnum_PlaceholderType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SITELINK",
	3:  "CALL",
	4:  "APP",
	5:  "LOCATION",
	6:  "AFFILIATE_LOCATION",
	7:  "CALLOUT",
	8:  "STRUCTURED_SNIPPET",
	9:  "MESSAGE",
	10: "PRICE",
	11: "PROMOTION",
	12: "AD_CUSTOMIZER",
	13: "DYNAMIC_EDUCATION",
	14: "DYNAMIC_FLIGHT",
	15: "DYNAMIC_CUSTOM",
	16: "DYNAMIC_HOTEL",
	17: "DYNAMIC_REAL_ESTATE",
	18: "DYNAMIC_TRAVEL",
	19: "DYNAMIC_LOCAL",
	20: "DYNAMIC_JOB",
}

var PlaceholderTypeEnum_PlaceholderType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"SITELINK":            2,
	"CALL":                3,
	"APP":                 4,
	"LOCATION":            5,
	"AFFILIATE_LOCATION":  6,
	"CALLOUT":             7,
	"STRUCTURED_SNIPPET":  8,
	"MESSAGE":             9,
	"PRICE":               10,
	"PROMOTION":           11,
	"AD_CUSTOMIZER":       12,
	"DYNAMIC_EDUCATION":   13,
	"DYNAMIC_FLIGHT":      14,
	"DYNAMIC_CUSTOM":      15,
	"DYNAMIC_HOTEL":       16,
	"DYNAMIC_REAL_ESTATE": 17,
	"DYNAMIC_TRAVEL":      18,
	"DYNAMIC_LOCAL":       19,
	"DYNAMIC_JOB":         20,
}

func (x PlaceholderTypeEnum_PlaceholderType) String() string {
	return proto.EnumName(PlaceholderTypeEnum_PlaceholderType_name, int32(x))
}

func (PlaceholderTypeEnum_PlaceholderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4007a2ae92623c7b, []int{0, 0}
}

// Container for enum describing possible placeholder types for a feed mapping.
type PlaceholderTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaceholderTypeEnum) Reset()         { *m = PlaceholderTypeEnum{} }
func (m *PlaceholderTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PlaceholderTypeEnum) ProtoMessage()    {}
func (*PlaceholderTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4007a2ae92623c7b, []int{0}
}

func (m *PlaceholderTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceholderTypeEnum.Unmarshal(m, b)
}
func (m *PlaceholderTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceholderTypeEnum.Marshal(b, m, deterministic)
}
func (m *PlaceholderTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceholderTypeEnum.Merge(m, src)
}
func (m *PlaceholderTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PlaceholderTypeEnum.Size(m)
}
func (m *PlaceholderTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceholderTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceholderTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType", PlaceholderTypeEnum_PlaceholderType_name, PlaceholderTypeEnum_PlaceholderType_value)
	proto.RegisterType((*PlaceholderTypeEnum)(nil), "google.ads.googleads.v1.enums.PlaceholderTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/placeholder_type.proto", fileDescriptor_4007a2ae92623c7b)
}

var fileDescriptor_4007a2ae92623c7b = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0x9b, 0x3e,
	0x14, 0xc6, 0xff, 0x21, 0x6d, 0x93, 0x38, 0x4d, 0xe3, 0x38, 0xfd, 0x6f, 0xd2, 0xb4, 0x5e, 0xb4,
	0x0f, 0x00, 0x42, 0xdb, 0x15, 0xbb, 0x72, 0xc0, 0x49, 0xbd, 0x3a, 0x60, 0x81, 0xc9, 0xb4, 0x2a,
	0x12, 0x62, 0x05, 0xb1, 0x48, 0x09, 0x46, 0x21, 0xad, 0xd4, 0xb7, 0xd8, 0x33, 0xec, 0x72, 0x8f,
	0xb2, 0x37, 0xd9, 0x9e, 0x60, 0x97, 0x93, 0xa1, 0x64, 0x51, 0xa5, 0xed, 0x06, 0x1d, 0xfd, 0xce,
	0xf7, 0x7d, 0xd8, 0x3e, 0x07, 0xbc, 0xcd, 0xa4, 0xcc, 0xd6, 0xa9, 0x11, 0x27, 0xa5, 0x51, 0x97,
	0xaa, 0x7a, 0x30, 0x8d, 0x34, 0xbf, 0xdf, 0x94, 0x46, 0xb1, 0x8e, 0xef, 0xd2, 0xcf, 0x72, 0x9d,
	0xa4, 0xdb, 0x68, 0xf7, 0x58, 0xa4, 0x7a, 0xb1, 0x95, 0x3b, 0x89, 0x2e, 0x6a, 0xa9, 0x1e, 0x27,
	0xa5, 0xbe, 0x77, 0xe9, 0x0f, 0xa6, 0x5e, 0xb9, 0x5e, 0xbd, 0x6e, 0x42, 0x8b, 0x95, 0x11, 0xe7,
	0xb9, 0xdc, 0xc5, 0xbb, 0x95, 0xcc, 0xcb, 0xda, 0x7c, 0xf5, 0xa5, 0x0d, 0xc6, 0xfc, 0x4f, 0xae,
	0x78, 0x2c, 0x52, 0x92, 0xdf, 0x6f, 0xae, 0x7e, 0x69, 0x60, 0xf8, 0x8c, 0xa3, 0x21, 0xe8, 0x87,
	0x6e, 0xc0, 0x89, 0x4d, 0xa7, 0x94, 0x38, 0xf0, 0x3f, 0xd4, 0x07, 0x9d, 0xd0, 0xbd, 0x71, 0xbd,
	0x0f, 0x2e, 0x6c, 0xa1, 0x53, 0xd0, 0x0d, 0xa8, 0x20, 0x8c, 0xba, 0x37, 0x50, 0x43, 0x5d, 0x70,
	0x64, 0x63, 0xc6, 0x60, 0x1b, 0x75, 0x40, 0x1b, 0x73, 0x0e, 0x8f, 0x94, 0x80, 0x79, 0x36, 0x16,
	0xd4, 0x73, 0xe1, 0x31, 0x7a, 0x01, 0x10, 0x9e, 0x4e, 0x29, 0xa3, 0x58, 0x90, 0x68, 0xcf, 0x4f,
	0x54, 0xa6, 0x32, 0x7a, 0xa1, 0x80, 0x1d, 0x25, 0x0a, 0x84, 0x1f, 0xda, 0x22, 0xf4, 0x89, 0x13,
	0x05, 0x2e, 0xe5, 0x9c, 0x08, 0xd8, 0x55, 0xa2, 0x39, 0x09, 0x02, 0x3c, 0x23, 0xb0, 0x87, 0x7a,
	0xe0, 0x98, 0xfb, 0xd4, 0x26, 0x10, 0xa0, 0x01, 0xe8, 0x71, 0xdf, 0x9b, 0x7b, 0x55, 0x56, 0x1f,
	0x8d, 0xc0, 0x00, 0x3b, 0x91, 0x1d, 0x06, 0xc2, 0x9b, 0xd3, 0x5b, 0xe2, 0xc3, 0x53, 0xf4, 0x3f,
	0x18, 0x39, 0x1f, 0x5d, 0x3c, 0xa7, 0x76, 0x44, 0x9c, 0xf0, 0xe9, 0xaf, 0x03, 0x84, 0xc0, 0x59,
	0x83, 0xa7, 0x8c, 0xce, 0xae, 0x05, 0x3c, 0x3b, 0x64, 0x75, 0x04, 0x1c, 0xaa, 0xc4, 0x86, 0x5d,
	0x7b, 0x82, 0x30, 0x08, 0xd1, 0x4b, 0x30, 0x6e, 0x90, 0x4f, 0x30, 0x8b, 0x48, 0x20, 0xb0, 0x20,
	0x70, 0x74, 0xe8, 0x17, 0x3e, 0x5e, 0x10, 0x06, 0xd1, 0xa1, 0x5f, 0xdd, 0x99, 0xc1, 0xb1, 0x7a,
	0xd5, 0x06, 0xbd, 0xf7, 0x26, 0xf0, 0x7c, 0xf2, 0xa3, 0x05, 0x2e, 0xef, 0xe4, 0x46, 0xff, 0xe7,
	0x58, 0x27, 0xe7, 0xcf, 0xa6, 0xc3, 0xd5, 0x38, 0x79, 0xeb, 0x76, 0xf2, 0x64, 0xcb, 0xe4, 0x3a,
	0xce, 0x33, 0x5d, 0x6e, 0x33, 0x23, 0x4b, 0xf3, 0x6a, 0xd8, 0xcd, 0x4e, 0x15, 0xab, 0xf2, 0x2f,
	0x2b, 0xf6, 0xae, 0xfa, 0x7e, 0xd5, 0xda, 0x33, 0x8c, 0xbf, 0x69, 0x17, 0xb3, 0x3a, 0x0a, 0x27,
	0xa5, 0x5e, 0x97, 0xaa, 0x5a, 0x98, 0xba, 0xda, 0x90, 0xf2, 0x7b, 0xd3, 0x5f, 0xe2, 0xa4, 0x5c,
	0xee, 0xfb, 0xcb, 0x85, 0xb9, 0xac, 0xfa, 0x3f, 0xb5, 0xcb, 0x1a, 0x5a, 0x16, 0x4e, 0x4a, 0xcb,
	0xda, 0x2b, 0x2c, 0x6b, 0x61, 0x5a, 0x56, 0xa5, 0xf9, 0x74, 0x52, 0x1d, 0xec, 0xcd, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0x55, 0x4f, 0x8d, 0xfa, 0x02, 0x00, 0x00,
}
