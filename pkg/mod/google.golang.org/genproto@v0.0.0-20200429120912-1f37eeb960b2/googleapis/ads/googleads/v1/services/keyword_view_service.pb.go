// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [KeywordViewService.GetKeywordView][google.ads.googleads.v1.services.KeywordViewService.GetKeywordView].
type GetKeywordViewRequest struct {
	// Required. The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf4cb0c6c667eaf, []int{0}
}

func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v1.services.GetKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/keyword_view_service.proto", fileDescriptor_eaf4cb0c6c667eaf)
}

var fileDescriptor_eaf4cb0c6c667eaf = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0x26, 0x11, 0x04, 0x83, 0x3a, 0x04, 0xe4, 0x57, 0xa3, 0x60, 0x29, 0x45, 0xa4, 0xc8, 0x9d,
	0x51, 0x41, 0xb8, 0xe2, 0x70, 0x1d, 0xac, 0x20, 0x48, 0x55, 0xc8, 0x20, 0x81, 0x70, 0x4d, 0x5e,
	0xe3, 0x69, 0x92, 0xab, 0x77, 0x69, 0x8a, 0x88, 0x8b, 0x5f, 0xc1, 0x6f, 0xe0, 0xe8, 0xb7, 0x70,
	0xed, 0xea, 0xe6, 0xe4, 0xe0, 0xe4, 0xee, 0xe2, 0x20, 0x92, 0x5e, 0x2e, 0x4d, 0xf5, 0x17, 0xba,
	0x3d, 0xe4, 0xf9, 0x73, 0xef, 0xfb, 0xbc, 0x71, 0xa6, 0xa9, 0x10, 0x69, 0x06, 0x98, 0x25, 0x0a,
	0x6b, 0x58, 0xa3, 0xca, 0xc7, 0x0a, 0x64, 0xc5, 0x63, 0x50, 0xf8, 0x35, 0xbc, 0xdd, 0x08, 0x99,
	0x44, 0x15, 0x87, 0x4d, 0xd4, 0x7c, 0x45, 0x2b, 0x29, 0x4a, 0xe1, 0x0e, 0xb5, 0x03, 0xb1, 0x44,
	0xa1, 0xd6, 0x8c, 0x2a, 0x1f, 0x19, 0xb3, 0x77, 0xb7, 0x2f, 0x5e, 0x82, 0x12, 0x6b, 0xf9, 0x6f,
	0xbe, 0xce, 0xf5, 0xae, 0x1a, 0xd7, 0x8a, 0x63, 0x56, 0x14, 0xa2, 0x64, 0x25, 0x17, 0x85, 0x6a,
	0xd8, 0x93, 0x0e, 0x1b, 0x67, 0x1c, 0x8a, 0xb2, 0x21, 0xae, 0x75, 0x88, 0x17, 0x1c, 0xb2, 0x24,
	0x5a, 0xc2, 0x4b, 0x56, 0x71, 0x21, 0x1b, 0xc1, 0xe5, 0x8e, 0xc0, 0x0c, 0xa0, 0xa9, 0xd1, 0x2b,
	0xe7, 0xd2, 0x1c, 0xca, 0x47, 0x7a, 0x96, 0x80, 0xc3, 0xe6, 0x29, 0xbc, 0x59, 0x83, 0x2a, 0xdd,
	0x27, 0xce, 0x05, 0x23, 0x8d, 0x0a, 0x96, 0xc3, 0xc0, 0x1a, 0x5a, 0x37, 0xce, 0xcd, 0x6e, 0x7e,
	0xa7, 0xf6, 0x6f, 0x7a, 0xdd, 0x19, 0xef, 0xf7, 0x6e, 0xd0, 0x8a, 0x2b, 0x14, 0x8b, 0x1c, 0x77,
	0xb3, 0xce, 0x9b, 0x88, 0xc7, 0x2c, 0x87, 0xdb, 0xbf, 0x2c, 0xc7, 0xed, 0xb0, 0xcf, 0x74, 0x59,
	0xee, 0x17, 0xcb, 0xb9, 0x78, 0x38, 0x83, 0x7b, 0x0f, 0x1d, 0x6b, 0x18, 0x9d, 0x3a, 0xb5, 0x87,
	0x7a, 0x8d, 0x6d, 0xf1, 0xa8, 0x63, 0x1b, 0x3d, 0xf8, 0x46, 0x0f, 0xd7, 0xfc, 0xf0, 0xf5, 0xc7,
	0x47, 0xfb, 0x96, 0x8b, 0xea, 0x5b, 0xbd, 0x3b, 0x60, 0xee, 0xc7, 0x6b, 0x55, 0x8a, 0x1c, 0xa4,
	0xc2, 0x13, 0x73, 0xbc, 0x3a, 0x43, 0xe1, 0xc9, 0x7b, 0xef, 0xca, 0x96, 0x0e, 0xfa, 0x1a, 0x99,
	0xfd, 0xb1, 0x9c, 0x71, 0x2c, 0xf2, 0xa3, 0x3b, 0xcd, 0x4e, 0xfe, 0x6f, 0x67, 0x51, 0x5f, 0x69,
	0x61, 0x3d, 0x7f, 0xd8, 0x98, 0x53, 0x91, 0xb1, 0x22, 0x45, 0x42, 0xa6, 0x38, 0x85, 0x62, 0x77,
	0x43, 0xbc, 0x7f, 0xae, 0xff, 0x77, 0x9e, 0x1a, 0xf0, 0xc9, 0x3e, 0x33, 0xa7, 0xf4, 0xb3, 0x3d,
	0x9c, 0xeb, 0x40, 0x9a, 0x28, 0xa4, 0x61, 0x8d, 0x02, 0x1f, 0x35, 0x0f, 0xab, 0xad, 0x91, 0x84,
	0x34, 0x51, 0x61, 0x2b, 0x09, 0x03, 0x3f, 0x34, 0x92, 0x9f, 0xf6, 0x58, 0x7f, 0x27, 0x84, 0x26,
	0x8a, 0x90, 0x56, 0x44, 0x48, 0xe0, 0x13, 0x62, 0x64, 0xcb, 0xb3, 0xbb, 0x39, 0xef, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xae, 0xac, 0x5d, 0x35, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordViewServiceClient(cc grpc.ClientConnInterface) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

// UnimplementedKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordViewServiceServer struct {
}

func (*UnimplementedKeywordViewServiceServer) GetKeywordView(ctx context.Context, req *GetKeywordViewRequest) (*resources.KeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordView not implemented")
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/keyword_view_service.proto",
}
