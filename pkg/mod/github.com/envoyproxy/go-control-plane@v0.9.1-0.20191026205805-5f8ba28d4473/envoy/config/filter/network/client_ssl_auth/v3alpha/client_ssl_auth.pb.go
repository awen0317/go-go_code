// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/client_ssl_auth/v3alpha/client_ssl_auth.proto

package envoy_config_filter_network_client_ssl_auth_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ClientSSLAuth struct {
	AuthApiCluster       string             `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	StatPrefix           string             `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RefreshDelay         *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	IpWhiteList          []*core.CidrRange  `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList,proto3" json:"ip_white_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClientSSLAuth) Reset()         { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()    {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2986cf35e37866a, []int{0}
}

func (m *ClientSSLAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSSLAuth.Unmarshal(m, b)
}
func (m *ClientSSLAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSSLAuth.Marshal(b, m, deterministic)
}
func (m *ClientSSLAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSSLAuth.Merge(m, src)
}
func (m *ClientSSLAuth) XXX_Size() int {
	return xxx_messageInfo_ClientSSLAuth.Size(m)
}
func (m *ClientSSLAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSSLAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSSLAuth proto.InternalMessageInfo

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*core.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.config.filter.network.client_ssl_auth.v3alpha.ClientSSLAuth")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/client_ssl_auth/v3alpha/client_ssl_auth.proto", fileDescriptor_c2986cf35e37866a)
}

var fileDescriptor_c2986cf35e37866a = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x6b, 0xeb, 0x30,
	0x14, 0xc5, 0x71, 0x12, 0xde, 0xe3, 0xc9, 0x2f, 0x8f, 0x87, 0x97, 0xba, 0x19, 0x8a, 0x5b, 0x3a,
	0x78, 0x92, 0x68, 0xb2, 0x17, 0xe2, 0xa4, 0x43, 0x21, 0x43, 0x70, 0x86, 0xd2, 0x49, 0x28, 0xb1,
	0x6c, 0x5f, 0x2a, 0x2c, 0x21, 0xc9, 0xf9, 0xf3, 0xb9, 0xbb, 0x75, 0x2a, 0xb2, 0x1c, 0x4a, 0x33,
	0x76, 0xb3, 0x75, 0xce, 0xfd, 0xdd, 0xa3, 0x23, 0xf4, 0xcc, 0x9b, 0xbd, 0x3c, 0x91, 0x9d, 0x6c,
	0x4a, 0xa8, 0x48, 0x09, 0xc2, 0x72, 0x4d, 0x1a, 0x6e, 0x0f, 0x52, 0xbf, 0x91, 0x9d, 0x00, 0xde,
	0x58, 0x6a, 0x8c, 0xa0, 0xac, 0xb5, 0x35, 0xd9, 0xcf, 0x98, 0x50, 0x35, 0xbb, 0x3c, 0xc7, 0x4a,
	0x4b, 0x2b, 0xa3, 0x59, 0x87, 0xc2, 0x1e, 0x85, 0x3d, 0x0a, 0xf7, 0x28, 0x7c, 0x39, 0xd2, 0xa3,
	0x26, 0xf7, 0x7e, 0x3f, 0x53, 0xf0, 0x45, 0x97, 0x9a, 0x13, 0x56, 0x14, 0x9a, 0x1b, 0xe3, 0xd1,
	0x93, 0x9b, 0x4a, 0xca, 0x4a, 0x70, 0xd2, 0xfd, 0x6d, 0xdb, 0x92, 0x14, 0xad, 0x66, 0x16, 0x64,
	0xd3, 0xeb, 0x57, 0x7b, 0x26, 0xa0, 0x60, 0x96, 0x93, 0xf3, 0x87, 0x17, 0xee, 0xde, 0x03, 0x34,
	0x5e, 0x74, 0xab, 0x37, 0x9b, 0xd5, 0xbc, 0xb5, 0x75, 0xf4, 0x80, 0xfe, 0xbb, 0x00, 0x94, 0x29,
	0xa0, 0x3b, 0xd1, 0x1a, 0xcb, 0x75, 0x1c, 0x24, 0x41, 0xfa, 0x27, 0xfb, 0xfd, 0x91, 0x8d, 0xf4,
	0x20, 0x09, 0xf2, 0x7f, 0xce, 0x30, 0x57, 0xb0, 0xf0, 0x72, 0x94, 0xa2, 0xd0, 0x58, 0x66, 0xa9,
	0xd2, 0xbc, 0x84, 0x63, 0x3c, 0xf8, 0xee, 0x46, 0x4e, 0x5b, 0x77, 0x52, 0xf4, 0x88, 0xc6, 0x9a,
	0x97, 0x9a, 0x9b, 0x9a, 0x16, 0x5c, 0xb0, 0x53, 0x3c, 0x4c, 0x82, 0x34, 0x9c, 0x5e, 0x63, 0x9f,
	0x1f, 0x9f, 0xf3, 0xe3, 0x65, 0x9f, 0x3f, 0xff, 0xdb, 0xfb, 0x97, 0xce, 0x1e, 0x3d, 0xa1, 0x31,
	0x28, 0x7a, 0xa8, 0xc1, 0x72, 0x2a, 0xc0, 0xd8, 0x78, 0x94, 0x0c, 0xd3, 0x70, 0x7a, 0x8b, 0x7d,
	0xb5, 0x4c, 0xc1, 0xb9, 0x38, 0xec, 0x5a, 0xc2, 0x0b, 0x28, 0x74, 0xce, 0x9a, 0x8a, 0xe7, 0x21,
	0xa8, 0x17, 0x37, 0xb6, 0x02, 0x63, 0xb3, 0x57, 0x34, 0x07, 0xe9, 0x67, 0x94, 0x96, 0xc7, 0x13,
	0xfe, 0xc1, 0xcb, 0x64, 0x51, 0xdf, 0x9b, 0x11, 0xae, 0xb7, 0xb5, 0x4b, 0xbe, 0x0e, 0xb6, 0xbf,
	0xba, 0x2b, 0xcc, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x2c, 0x37, 0x41, 0x38, 0x02, 0x00,
	0x00,
}
