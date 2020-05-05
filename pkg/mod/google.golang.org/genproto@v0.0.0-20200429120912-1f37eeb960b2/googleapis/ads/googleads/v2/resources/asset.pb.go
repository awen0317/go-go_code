// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/asset.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
type Asset struct {
	// Immutable. The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the asset.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Optional name of the asset.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// The specific type of the asset.
	//
	// Types that are valid to be assigned to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	//	*Asset_TextAsset
	AssetData            isAsset_AssetData `protobuf_oneof:"asset_data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_615ce8003e62b64c, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Asset) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if m != nil {
		return m.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

type Asset_TextAsset struct {
	TextAsset *common.TextAsset `protobuf:"bytes,8,opt,name=text_asset,json=textAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (*Asset_TextAsset) isAsset_AssetData() {}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := m.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (m *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := m.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (m *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := m.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

func (m *Asset) GetTextAsset() *common.TextAsset {
	if x, ok := m.GetAssetData().(*Asset_TextAsset); ok {
		return x.TextAsset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Asset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
		(*Asset_TextAsset)(nil),
	}
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.ads.googleads.v2.resources.Asset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/asset.proto", fileDescriptor_615ce8003e62b64c)
}

var fileDescriptor_615ce8003e62b64c = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0x97, 0x74, 0xdd, 0xff, 0x3f, 0x6f, 0x20, 0x08, 0x97, 0x32, 0xa6, 0xd1, 0x81, 0x86,
	0x06, 0x12, 0x4e, 0x09, 0x68, 0x87, 0x70, 0x4a, 0xa4, 0x69, 0x1d, 0x08, 0x36, 0x85, 0xa9, 0x12,
	0xa8, 0x52, 0xe4, 0xd6, 0x5e, 0x66, 0xa9, 0xb6, 0xa3, 0xd8, 0x29, 0xab, 0xa6, 0x7d, 0x19, 0x8e,
	0x7c, 0x14, 0xbe, 0x00, 0xd7, 0x9d, 0xf7, 0x05, 0x90, 0x38, 0xa1, 0xd8, 0x49, 0xda, 0x0d, 0x95,
	0x9e, 0xfa, 0x26, 0xef, 0xf3, 0xfc, 0xde, 0x47, 0xce, 0xeb, 0x82, 0x97, 0x89, 0x10, 0xc9, 0x88,
	0xb8, 0x08, 0x4b, 0xd7, 0x94, 0x45, 0x35, 0xf6, 0xdc, 0x8c, 0x48, 0x91, 0x67, 0x43, 0x22, 0x5d,
	0x24, 0x25, 0x51, 0x30, 0xcd, 0x84, 0x12, 0xce, 0xb6, 0xd1, 0x40, 0x84, 0x25, 0xac, 0xe5, 0x70,
	0xec, 0xc1, 0x5a, 0xbe, 0xd1, 0x99, 0x47, 0x1c, 0x0a, 0xc6, 0x04, 0x37, 0xb8, 0x58, 0x4d, 0x52,
	0x22, 0x0d, 0x74, 0x03, 0xce, 0x73, 0x10, 0x9e, 0x33, 0x39, 0x63, 0x28, 0xf5, 0x8f, 0x2b, 0x7d,
	0x4a, 0xdd, 0x53, 0x4a, 0x46, 0x38, 0x1e, 0x90, 0x33, 0x34, 0xa6, 0x22, 0x2b, 0x05, 0x0f, 0x67,
	0x04, 0x55, 0xb0, 0xb2, 0xb5, 0x55, 0xb6, 0xf4, 0xd3, 0x20, 0x3f, 0x75, 0xbf, 0x66, 0x28, 0x4d,
	0x49, 0x56, 0x65, 0xd9, 0x9c, 0xb1, 0x22, 0xce, 0x85, 0x42, 0x8a, 0x0a, 0x5e, 0x76, 0x9f, 0xfc,
	0x6c, 0x82, 0x66, 0x50, 0xc4, 0x71, 0xde, 0x83, 0x3b, 0x15, 0x39, 0xe6, 0x88, 0x91, 0x96, 0xd5,
	0xb6, 0x76, 0x57, 0xc3, 0x67, 0x57, 0x41, 0xf3, 0x77, 0xd0, 0x06, 0x5b, 0xd3, 0xc3, 0x29, 0xab,
	0x94, 0x4a, 0x38, 0x14, 0xcc, 0xd5, 0xf6, 0x68, 0xbd, 0x32, 0x7f, 0x44, 0x8c, 0x38, 0x1d, 0x60,
	0x53, 0xdc, 0xb2, 0xdb, 0xd6, 0xee, 0x9a, 0xf7, 0xa8, 0x34, 0xc0, 0x2a, 0x21, 0x3c, 0xe4, 0x6a,
	0xef, 0x4d, 0x0f, 0x8d, 0x72, 0x12, 0x36, 0xae, 0x82, 0x46, 0x64, 0x53, 0xec, 0x74, 0xc0, 0xb2,
	0x9e, 0xda, 0xd0, 0x9e, 0xcd, 0xbf, 0x3c, 0x9f, 0x54, 0x46, 0x79, 0xa2, 0x4d, 0x91, 0x56, 0x3a,
	0x47, 0x60, 0xb9, 0x38, 0xc2, 0xd6, 0x72, 0xdb, 0xda, 0xbd, 0xeb, 0xed, 0xc1, 0x79, 0x1f, 0x52,
	0x9f, 0x39, 0xd4, 0x29, 0x4f, 0x26, 0x29, 0xd9, 0xe7, 0x39, 0x9b, 0x3e, 0x99, 0x00, 0x1a, 0xe4,
	0x9c, 0x81, 0x07, 0x13, 0x91, 0xab, 0x7c, 0x40, 0xe2, 0x31, 0xc5, 0x44, 0xc4, 0xfa, 0x3b, 0xb5,
	0x9a, 0x3a, 0xd1, 0xab, 0xb9, 0x7c, 0xb3, 0x05, 0xf0, 0xb3, 0xb1, 0xf6, 0x0a, 0xa7, 0xc6, 0x17,
	0xe8, 0x66, 0x77, 0x29, 0xba, 0x3f, 0xb9, 0xdd, 0x71, 0x08, 0x70, 0x18, 0xc1, 0x14, 0xc5, 0x83,
	0x9c, 0xe3, 0x11, 0x29, 0x07, 0xad, 0xe8, 0x41, 0x9d, 0x45, 0x83, 0x3e, 0x14, 0xce, 0x50, 0x1b,
	0x6f, 0xcc, 0xb9, 0xc7, 0x6e, 0x35, 0x9c, 0x08, 0xac, 0x51, 0x86, 0x92, 0x8a, 0xff, 0x9f, 0xe6,
	0xbf, 0x58, 0xc4, 0x3f, 0x2c, 0x2c, 0x35, 0xb9, 0xd1, 0x5d, 0x8a, 0x00, 0xad, 0x5f, 0x39, 0x47,
	0x00, 0x28, 0x72, 0xae, 0x4a, 0xe4, 0xff, 0x1a, 0xf9, 0x7c, 0x11, 0xf2, 0x84, 0x9c, 0xab, 0x1b,
	0xc4, 0x55, 0x55, 0xbd, 0xf1, 0xbb, 0xd7, 0xc1, 0xfe, 0xa2, 0xed, 0x72, 0x9e, 0x0e, 0x73, 0xa9,
	0x04, 0x23, 0x99, 0x74, 0x2f, 0xaa, 0xf2, 0xd2, 0xdc, 0x23, 0xe9, 0x5e, 0xe8, 0xdf, 0xcb, 0x70,
	0x1d, 0x00, 0x73, 0xb3, 0x30, 0x52, 0x28, 0xfc, 0x65, 0x81, 0x9d, 0xa1, 0x60, 0x70, 0xe1, 0xfd,
	0x0e, 0x81, 0x9e, 0x71, 0x5c, 0x6c, 0xda, 0xb1, 0xf5, 0xe5, 0x5d, 0x69, 0x48, 0xc4, 0x08, 0xf1,
	0x04, 0x8a, 0x2c, 0x71, 0x13, 0xc2, 0xf5, 0x1e, 0xba, 0xd3, 0x6c, 0xff, 0xf8, 0x7b, 0x79, 0x5b,
	0x57, 0xdf, 0xec, 0xc6, 0x41, 0x10, 0x7c, 0xb7, 0xb7, 0x0f, 0x0c, 0x32, 0xc0, 0x12, 0x9a, 0xb2,
	0xa8, 0x7a, 0x1e, 0x8c, 0x2a, 0xe5, 0x8f, 0x4a, 0xd3, 0x0f, 0xb0, 0xec, 0xd7, 0x9a, 0x7e, 0xcf,
	0xeb, 0xd7, 0x9a, 0x6b, 0x7b, 0xc7, 0x34, 0x7c, 0x3f, 0xc0, 0xd2, 0xf7, 0x6b, 0x95, 0xef, 0xf7,
	0x3c, 0xdf, 0xaf, 0x75, 0x83, 0x15, 0x1d, 0xf6, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55,
	0x70, 0x6a, 0x51, 0x0a, 0x05, 0x00, 0x00,
}
