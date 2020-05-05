// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_schedule_view.proto

package resources

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

// An ad schedule view summarizes the performance of campaigns by
// AdSchedule criteria.
type AdScheduleView struct {
	// Output only. The resource name of the ad schedule view.
	// AdSchedule view resource names have the form:
	//
	// `customers/{customer_id}/adScheduleViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdScheduleView) Reset()         { *m = AdScheduleView{} }
func (m *AdScheduleView) String() string { return proto.CompactTextString(m) }
func (*AdScheduleView) ProtoMessage()    {}
func (*AdScheduleView) Descriptor() ([]byte, []int) {
	return fileDescriptor_99aaa568f8d5a785, []int{0}
}

func (m *AdScheduleView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdScheduleView.Unmarshal(m, b)
}
func (m *AdScheduleView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdScheduleView.Marshal(b, m, deterministic)
}
func (m *AdScheduleView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdScheduleView.Merge(m, src)
}
func (m *AdScheduleView) XXX_Size() int {
	return xxx_messageInfo_AdScheduleView.Size(m)
}
func (m *AdScheduleView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdScheduleView.DiscardUnknown(m)
}

var xxx_messageInfo_AdScheduleView proto.InternalMessageInfo

func (m *AdScheduleView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdScheduleView)(nil), "google.ads.googleads.v3.resources.AdScheduleView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_schedule_view.proto", fileDescriptor_99aaa568f8d5a785)
}

var fileDescriptor_99aaa568f8d5a785 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x49, 0x0a, 0x82, 0x41, 0x1d, 0xea, 0xa2, 0x45, 0xd0, 0x0a, 0x45, 0x5d, 0xee, 0x86,
	0x0c, 0xca, 0x39, 0x5d, 0x97, 0x82, 0x83, 0x94, 0x2a, 0x19, 0x24, 0x10, 0xae, 0xb9, 0xbf, 0xe9,
	0x41, 0x92, 0x2b, 0xb9, 0x34, 0x1d, 0x4a, 0xc1, 0xcf, 0xe2, 0xe8, 0x47, 0x11, 0xfc, 0x0e, 0xce,
	0xfd, 0x08, 0x4e, 0xd2, 0x5e, 0xee, 0xda, 0x3a, 0xa8, 0xdb, 0x83, 0xff, 0xef, 0xbd, 0x7b, 0xbc,
	0xf3, 0x6e, 0x12, 0x29, 0x93, 0x14, 0x30, 0xe3, 0x0a, 0x6b, 0xb9, 0x54, 0x95, 0x8f, 0x0b, 0x50,
	0x72, 0x52, 0xc4, 0xa0, 0x30, 0xe3, 0x91, 0x8a, 0x47, 0xc0, 0x27, 0x29, 0x44, 0x95, 0x80, 0x29,
	0x1a, 0x17, 0xb2, 0x94, 0xcd, 0xb6, 0xc6, 0x11, 0xe3, 0x0a, 0x59, 0x27, 0xaa, 0x7c, 0x64, 0x9d,
	0xad, 0x53, 0x13, 0x3e, 0x16, 0xf8, 0x59, 0x40, 0xca, 0xa3, 0x21, 0x8c, 0x58, 0x25, 0x64, 0xa1,
	0x33, 0x5a, 0xc7, 0x1b, 0x80, 0xb1, 0xd5, 0xa7, 0x93, 0x8d, 0x13, 0xcb, 0x73, 0x59, 0xb2, 0x52,
	0xc8, 0x5c, 0xe9, 0xeb, 0xf9, 0x87, 0xe3, 0x1d, 0x50, 0xfe, 0x50, 0xd7, 0x0a, 0x04, 0x4c, 0x9b,
	0x8f, 0xde, 0xbe, 0x89, 0x88, 0x72, 0x96, 0xc1, 0x91, 0x73, 0xe6, 0x5c, 0xee, 0x76, 0xf1, 0x27,
	0x6d, 0x7c, 0xd1, 0x2b, 0xef, 0x62, 0xdd, 0xb1, 0x56, 0x63, 0xa1, 0x50, 0x2c, 0x33, 0xbc, 0x9d,
	0x33, 0xd8, 0x33, 0x29, 0xf7, 0x2c, 0x03, 0x02, 0x0b, 0x3a, 0xfc, 0xb7, 0xb7, 0x79, 0x1d, 0x4f,
	0x54, 0x29, 0x33, 0x28, 0x14, 0x9e, 0x19, 0x39, 0xc7, 0x6c, 0x0b, 0x52, 0x78, 0xf6, 0x73, 0xd1,
	0x79, 0xf7, 0xc5, 0xf5, 0x3a, 0xb1, 0xcc, 0xd0, 0x9f, 0x9b, 0x76, 0x0f, 0xb7, 0x9f, 0xec, 0x2f,
	0xe7, 0xe8, 0x3b, 0x4f, 0x77, 0xb5, 0x33, 0x91, 0x29, 0xcb, 0x13, 0x24, 0x8b, 0x04, 0x27, 0x90,
	0xaf, 0xc6, 0xc2, 0xeb, 0xce, 0xbf, 0x7c, 0xf3, 0xad, 0x55, 0xaf, 0x6e, 0xa3, 0x47, 0xe9, 0x9b,
	0xdb, 0xee, 0xe9, 0x48, 0xca, 0x15, 0xd2, 0x72, 0xa9, 0x02, 0x1f, 0x0d, 0x0c, 0xf9, 0x6e, 0x98,
	0x90, 0x72, 0x15, 0x5a, 0x26, 0x0c, 0xfc, 0xd0, 0x32, 0x0b, 0xb7, 0xa3, 0x0f, 0x84, 0x50, 0xae,
	0x08, 0xb1, 0x14, 0x21, 0x81, 0x4f, 0x88, 0xe5, 0x86, 0x3b, 0xab, 0xb2, 0xfe, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x58, 0xb5, 0xa2, 0x92, 0x02, 0x00, 0x00,
}
