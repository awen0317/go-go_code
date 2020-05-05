// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/group_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [GroupPlacementViewService.GetGroupPlacementView][google.ads.googleads.v3.services.GroupPlacementViewService.GetGroupPlacementView].
type GetGroupPlacementViewRequest struct {
	// Required. The resource name of the Group Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupPlacementViewRequest) Reset()         { *m = GetGroupPlacementViewRequest{} }
func (m *GetGroupPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupPlacementViewRequest) ProtoMessage()    {}
func (*GetGroupPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_366e0655413e5b2b, []int{0}
}

func (m *GetGroupPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetGroupPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupPlacementViewRequest.Merge(m, src)
}
func (m *GetGroupPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Size(m)
}
func (m *GetGroupPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupPlacementViewRequest proto.InternalMessageInfo

func (m *GetGroupPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGroupPlacementViewRequest)(nil), "google.ads.googleads.v3.services.GetGroupPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/group_placement_view_service.proto", fileDescriptor_366e0655413e5b2b)
}

var fileDescriptor_366e0655413e5b2b = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x3d, 0x8b, 0xd5, 0x40,
	0x14, 0xe5, 0x45, 0x10, 0x0c, 0xda, 0x04, 0xc4, 0xdd, 0xb8, 0xe8, 0x63, 0xd9, 0x42, 0x56, 0x9c,
	0x01, 0xc3, 0x22, 0x8c, 0x1f, 0x30, 0xcf, 0x22, 0x36, 0xca, 0x43, 0x21, 0x88, 0x04, 0xc2, 0x6c,
	0x72, 0x8d, 0x03, 0xc9, 0x4c, 0x9c, 0x99, 0x64, 0x05, 0xb1, 0x11, 0xfc, 0x05, 0x36, 0xd6, 0x96,
	0xfe, 0x94, 0x6d, 0xed, 0x04, 0xc1, 0xc2, 0xca, 0x9f, 0x60, 0x25, 0xd9, 0xc9, 0x64, 0xdf, 0x63,
	0x37, 0xbe, 0xee, 0x90, 0x73, 0xee, 0xb9, 0x77, 0xce, 0x21, 0xfe, 0xe3, 0x52, 0xca, 0xb2, 0x02,
	0xcc, 0x0a, 0x8d, 0x2d, 0xec, 0x51, 0x17, 0x61, 0x0d, 0xaa, 0xe3, 0x39, 0x68, 0x5c, 0x2a, 0xd9,
	0x36, 0x59, 0x53, 0xb1, 0x1c, 0x6a, 0x10, 0x26, 0xeb, 0x38, 0x1c, 0x65, 0x03, 0x8b, 0x1a, 0x25,
	0x8d, 0x0c, 0xe6, 0x76, 0x12, 0xb1, 0x42, 0xa3, 0xd1, 0x04, 0x75, 0x11, 0x72, 0x26, 0xe1, 0x83,
	0xa9, 0x35, 0x0a, 0xb4, 0x6c, 0xd5, 0xd4, 0x1e, 0xeb, 0x1f, 0xee, 0xb8, 0xe9, 0x86, 0x63, 0x26,
	0x84, 0x34, 0xcc, 0x70, 0x29, 0xf4, 0xc0, 0x5e, 0x5b, 0x61, 0xf3, 0x8a, 0x83, 0x30, 0x03, 0x71,
	0x73, 0x85, 0x78, 0xcd, 0xa1, 0x2a, 0xb2, 0x43, 0x78, 0xc3, 0x3a, 0x2e, 0xd5, 0x20, 0xd8, 0x5e,
	0x11, 0xb8, 0x43, 0x2c, 0xb5, 0xfb, 0xce, 0xdf, 0x89, 0xc1, 0xc4, 0xfd, 0x4d, 0x4b, 0x77, 0x52,
	0xc2, 0xe1, 0xe8, 0x39, 0xbc, 0x6d, 0x41, 0x9b, 0xe0, 0xa5, 0x7f, 0xc5, 0x4d, 0x64, 0x82, 0xd5,
	0xb0, 0x35, 0x9b, 0xcf, 0x6e, 0x5d, 0x5a, 0x44, 0xbf, 0xa8, 0xf7, 0x97, 0xde, 0xf1, 0x6f, 0x9f,
	0xc6, 0x30, 0xa0, 0x86, 0x6b, 0x94, 0xcb, 0x1a, 0x9f, 0x63, 0x79, 0xd9, 0x39, 0x3d, 0x63, 0x35,
	0xdc, 0xfd, 0xe2, 0xf9, 0xdb, 0x67, 0x45, 0x2f, 0x6c, 0x92, 0xc1, 0xcf, 0x99, 0x7f, 0xf5, 0xdc,
	0xc3, 0x82, 0x47, 0x68, 0x53, 0x0b, 0xe8, 0x7f, 0x2f, 0x0a, 0x0f, 0x26, 0xe7, 0xc7, 0x8e, 0xd0,
	0xd9, 0xe9, 0xdd, 0xa7, 0x3f, 0xe8, 0x7a, 0x12, 0x1f, 0xbf, 0xff, 0xfe, 0xec, 0xdd, 0x0b, 0x0e,
	0xfa, 0x76, 0xdf, 0xaf, 0x31, 0x0f, 0xf3, 0x56, 0x1b, 0x59, 0x83, 0xd2, 0x78, 0xdf, 0xd6, 0xbd,
	0x66, 0xa5, 0xf1, 0xfe, 0x87, 0xf0, 0xfa, 0x31, 0xdd, 0x9a, 0xca, 0x6e, 0xf1, 0xc9, 0xf3, 0xf7,
	0x72, 0x59, 0x6f, 0x7c, 0xe8, 0xe2, 0xc6, 0x64, 0x80, 0xcb, 0xbe, 0xdd, 0xe5, 0xec, 0xd5, 0x93,
	0xc1, 0xa3, 0x94, 0x15, 0x13, 0x25, 0x92, 0xaa, 0xc4, 0x25, 0x88, 0x93, 0xee, 0xf1, 0xe9, 0xd6,
	0xe9, 0xdf, 0xe2, 0xbe, 0x03, 0x5f, 0xbd, 0x0b, 0x31, 0xa5, 0xdf, 0xbc, 0x79, 0x6c, 0x0d, 0x69,
	0xa1, 0x91, 0x85, 0x3d, 0x4a, 0x22, 0x34, 0x2c, 0xd6, 0xc7, 0x4e, 0x92, 0xd2, 0x42, 0xa7, 0xa3,
	0x24, 0x4d, 0xa2, 0xd4, 0x49, 0xfe, 0x78, 0x7b, 0xf6, 0x3b, 0x21, 0xb4, 0xd0, 0x84, 0x8c, 0x22,
	0x42, 0x92, 0x88, 0x10, 0x27, 0x3b, 0xbc, 0x78, 0x72, 0x67, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0x71, 0x3a, 0x5b, 0xbd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupPlacementViewServiceClient is the client API for GroupPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupPlacementViewServiceClient interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error)
}

type groupPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupPlacementViewServiceClient(cc grpc.ClientConnInterface) GroupPlacementViewServiceClient {
	return &groupPlacementViewServiceClient{cc}
}

func (c *groupPlacementViewServiceClient) GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error) {
	out := new(resources.GroupPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.GroupPlacementViewService/GetGroupPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupPlacementViewServiceServer is the server API for GroupPlacementViewService service.
type GroupPlacementViewServiceServer interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error)
}

// UnimplementedGroupPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupPlacementViewServiceServer struct {
}

func (*UnimplementedGroupPlacementViewServiceServer) GetGroupPlacementView(ctx context.Context, req *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPlacementView not implemented")
}

func RegisterGroupPlacementViewServiceServer(s *grpc.Server, srv GroupPlacementViewServiceServer) {
	s.RegisterService(&_GroupPlacementViewService_serviceDesc, srv)
}

func _GroupPlacementViewService_GetGroupPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.GroupPlacementViewService/GetGroupPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, req.(*GetGroupPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.GroupPlacementViewService",
	HandlerType: (*GroupPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupPlacementView",
			Handler:    _GroupPlacementViewService_GetGroupPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/group_placement_view_service.proto",
}
