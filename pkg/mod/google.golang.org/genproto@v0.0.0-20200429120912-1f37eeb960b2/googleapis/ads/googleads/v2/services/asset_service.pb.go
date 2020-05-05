// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/asset_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [AssetService.GetAsset][google.ads.googleads.v2.services.AssetService.GetAsset]
type GetAssetRequest struct {
	// Required. The resource name of the asset to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssetRequest) Reset()         { *m = GetAssetRequest{} }
func (m *GetAssetRequest) String() string { return proto.CompactTextString(m) }
func (*GetAssetRequest) ProtoMessage()    {}
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5e2650216741c6, []int{0}
}

func (m *GetAssetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssetRequest.Unmarshal(m, b)
}
func (m *GetAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssetRequest.Marshal(b, m, deterministic)
}
func (m *GetAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssetRequest.Merge(m, src)
}
func (m *GetAssetRequest) XXX_Size() int {
	return xxx_messageInfo_GetAssetRequest.Size(m)
}
func (m *GetAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssetRequest proto.InternalMessageInfo

func (m *GetAssetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AssetService.MutateAssets][google.ads.googleads.v2.services.AssetService.MutateAssets]
type MutateAssetsRequest struct {
	// Required. The ID of the customer whose assets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual assets.
	Operations           []*AssetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MutateAssetsRequest) Reset()         { *m = MutateAssetsRequest{} }
func (m *MutateAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAssetsRequest) ProtoMessage()    {}
func (*MutateAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5e2650216741c6, []int{1}
}

func (m *MutateAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetsRequest.Unmarshal(m, b)
}
func (m *MutateAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetsRequest.Merge(m, src)
}
func (m *MutateAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAssetsRequest.Size(m)
}
func (m *MutateAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetsRequest proto.InternalMessageInfo

func (m *MutateAssetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAssetsRequest) GetOperations() []*AssetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation to create an asset. Supported asset types are
// YoutubeVideoAsset, MediaBundleAsset, ImageAsset, and LeadFormAsset. TextAsset
// should be created with Ad inline.
type AssetOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AssetOperation_Create
	Operation            isAssetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AssetOperation) Reset()         { *m = AssetOperation{} }
func (m *AssetOperation) String() string { return proto.CompactTextString(m) }
func (*AssetOperation) ProtoMessage()    {}
func (*AssetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5e2650216741c6, []int{2}
}

func (m *AssetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetOperation.Unmarshal(m, b)
}
func (m *AssetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetOperation.Marshal(b, m, deterministic)
}
func (m *AssetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetOperation.Merge(m, src)
}
func (m *AssetOperation) XXX_Size() int {
	return xxx_messageInfo_AssetOperation.Size(m)
}
func (m *AssetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AssetOperation proto.InternalMessageInfo

type isAssetOperation_Operation interface {
	isAssetOperation_Operation()
}

type AssetOperation_Create struct {
	Create *resources.Asset `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*AssetOperation_Create) isAssetOperation_Operation() {}

func (m *AssetOperation) GetOperation() isAssetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AssetOperation) GetCreate() *resources.Asset {
	if x, ok := m.GetOperation().(*AssetOperation_Create); ok {
		return x.Create
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssetOperation_Create)(nil),
	}
}

// Response message for an asset mutate.
type MutateAssetsResponse struct {
	// All results for the mutate.
	Results              []*MutateAssetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutateAssetsResponse) Reset()         { *m = MutateAssetsResponse{} }
func (m *MutateAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAssetsResponse) ProtoMessage()    {}
func (*MutateAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5e2650216741c6, []int{3}
}

func (m *MutateAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetsResponse.Unmarshal(m, b)
}
func (m *MutateAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetsResponse.Merge(m, src)
}
func (m *MutateAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAssetsResponse.Size(m)
}
func (m *MutateAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetsResponse proto.InternalMessageInfo

func (m *MutateAssetsResponse) GetResults() []*MutateAssetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the asset mutate.
type MutateAssetResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAssetResult) Reset()         { *m = MutateAssetResult{} }
func (m *MutateAssetResult) String() string { return proto.CompactTextString(m) }
func (*MutateAssetResult) ProtoMessage()    {}
func (*MutateAssetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5e2650216741c6, []int{4}
}

func (m *MutateAssetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetResult.Unmarshal(m, b)
}
func (m *MutateAssetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetResult.Marshal(b, m, deterministic)
}
func (m *MutateAssetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetResult.Merge(m, src)
}
func (m *MutateAssetResult) XXX_Size() int {
	return xxx_messageInfo_MutateAssetResult.Size(m)
}
func (m *MutateAssetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetResult proto.InternalMessageInfo

func (m *MutateAssetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAssetRequest)(nil), "google.ads.googleads.v2.services.GetAssetRequest")
	proto.RegisterType((*MutateAssetsRequest)(nil), "google.ads.googleads.v2.services.MutateAssetsRequest")
	proto.RegisterType((*AssetOperation)(nil), "google.ads.googleads.v2.services.AssetOperation")
	proto.RegisterType((*MutateAssetsResponse)(nil), "google.ads.googleads.v2.services.MutateAssetsResponse")
	proto.RegisterType((*MutateAssetResult)(nil), "google.ads.googleads.v2.services.MutateAssetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/asset_service.proto", fileDescriptor_be5e2650216741c6)
}

var fileDescriptor_be5e2650216741c6 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xff, 0xec, 0x48, 0xfd, 0xe8, 0xa6, 0x80, 0xba, 0x20, 0x28, 0x01, 0x41, 0x64, 0x22, 0x14,
	0x05, 0x58, 0x17, 0x17, 0x10, 0x32, 0xaa, 0xc4, 0xe6, 0x92, 0x22, 0x54, 0x5a, 0x05, 0x91, 0x03,
	0x8a, 0x88, 0xb6, 0xf6, 0x62, 0x2c, 0xc5, 0xde, 0xe0, 0xd9, 0xe4, 0x52, 0xf5, 0xc2, 0x23, 0xc0,
	0x1b, 0x70, 0x02, 0x1e, 0xa5, 0x07, 0x2e, 0xdc, 0x7a, 0xea, 0x81, 0x13, 0x4f, 0x80, 0x38, 0x21,
	0x7b, 0xbd, 0x8e, 0x03, 0x44, 0xa1, 0xb7, 0x89, 0xe7, 0xf7, 0x67, 0x66, 0x67, 0x26, 0xe8, 0x6e,
	0x20, 0x44, 0x30, 0xe4, 0x36, 0xf3, 0xc1, 0x56, 0x61, 0x1a, 0x4d, 0x1c, 0x1b, 0x78, 0x32, 0x09,
	0x3d, 0x0e, 0x36, 0x03, 0xe0, 0x72, 0x90, 0xff, 0x24, 0xa3, 0x44, 0x48, 0x81, 0xeb, 0x0a, 0x4a,
	0x98, 0x0f, 0xa4, 0x60, 0x91, 0x89, 0x43, 0x34, 0xab, 0x76, 0x7b, 0x9e, 0x6e, 0xc2, 0x41, 0x8c,
	0x93, 0x42, 0x58, 0x09, 0xd6, 0xae, 0x68, 0xf8, 0x28, 0xb4, 0x59, 0x1c, 0x0b, 0xc9, 0x64, 0x28,
	0x62, 0xc8, 0xb3, 0x17, 0x4b, 0x59, 0x6f, 0x18, 0xf2, 0x58, 0xd3, 0xae, 0x95, 0x12, 0xaf, 0x42,
	0x3e, 0xf4, 0x07, 0x7b, 0xfc, 0x35, 0x9b, 0x84, 0x22, 0xc9, 0x01, 0x97, 0x4a, 0x00, 0xed, 0xac,
	0x52, 0xd6, 0x4b, 0x74, 0xb6, 0xc3, 0x25, 0x4d, 0x8b, 0xe8, 0xf2, 0x37, 0x63, 0x0e, 0x12, 0x3f,
	0x41, 0xa7, 0x35, 0x68, 0x10, 0xb3, 0x88, 0xaf, 0x19, 0x75, 0xa3, 0xb9, 0xdc, 0xbe, 0x71, 0x4c,
	0xcd, 0x9f, 0xb4, 0x8e, 0xae, 0x4e, 0x5b, 0xcd, 0xa3, 0x51, 0x08, 0xc4, 0x13, 0x91, 0xad, 0x54,
	0x56, 0x34, 0xf9, 0x29, 0x8b, 0xb8, 0xf5, 0xce, 0x40, 0xe7, 0xb6, 0xc7, 0x92, 0x49, 0x9e, 0x65,
	0x41, 0x9b, 0x34, 0x50, 0xd5, 0x1b, 0x83, 0x14, 0x11, 0x4f, 0x06, 0xa1, 0x9f, 0x5b, 0x54, 0x8e,
	0xa9, 0xd9, 0x45, 0xfa, 0xfb, 0x63, 0x1f, 0x3f, 0x47, 0x48, 0x8c, 0x78, 0xa2, 0x9e, 0x61, 0xcd,
	0xac, 0x57, 0x9a, 0x55, 0x67, 0x9d, 0x2c, 0x7a, 0x76, 0x92, 0x59, 0xed, 0x68, 0x62, 0x2e, 0x3b,
	0x15, 0xb2, 0x18, 0x3a, 0x33, 0x0b, 0xc1, 0x6d, 0xb4, 0xe4, 0x25, 0x9c, 0x49, 0xd5, 0x6c, 0xd5,
	0x69, 0xce, 0x35, 0x29, 0x26, 0xa7, 0x5c, 0xb6, 0xfe, 0xeb, 0xe6, 0xcc, 0x76, 0x15, 0x2d, 0x17,
	0x1e, 0x16, 0x47, 0xe7, 0x67, 0xdb, 0x86, 0x91, 0x88, 0x81, 0xe3, 0x6d, 0xf4, 0x7f, 0xc2, 0x61,
	0x3c, 0x94, 0xba, 0x9d, 0x8d, 0xc5, 0xed, 0x94, 0x84, 0xba, 0x19, 0xb7, 0xab, 0x35, 0xac, 0x07,
	0x68, 0xf5, 0x8f, 0x2c, 0xbe, 0xfe, 0xd7, 0x01, 0xce, 0x0e, 0xc6, 0xf9, 0x58, 0x41, 0x2b, 0x19,
	0xe9, 0x99, 0xb2, 0xc1, 0x9f, 0x0c, 0x74, 0x4a, 0xaf, 0x02, 0xbe, 0xb3, 0xb8, 0xaa, 0xdf, 0xd6,
	0xa6, 0xf6, 0xcf, 0x4f, 0x66, 0x3d, 0x3a, 0xa2, 0xb3, 0x05, 0xbe, 0xfd, 0xfa, 0xed, 0xbd, 0xd9,
	0xc2, 0xcd, 0xf4, 0x32, 0xf6, 0x67, 0x32, 0x9b, 0x7a, 0x19, 0xc0, 0x6e, 0xa9, 0x53, 0x01, 0xbb,
	0x75, 0x80, 0xbf, 0x18, 0x68, 0xa5, 0xfc, 0xbc, 0xf8, 0xde, 0x89, 0x5e, 0x51, 0x6f, 0x61, 0xed,
	0xfe, 0x49, 0x69, 0x6a, 0x8a, 0xd6, 0xce, 0x11, 0xbd, 0x50, 0x5a, 0xdf, 0x5b, 0xd3, 0xdd, 0xca,
	0x5a, 0x59, 0xb7, 0x6e, 0xa6, 0xad, 0x4c, 0x6b, 0xdf, 0x2f, 0x81, 0x37, 0x5b, 0x07, 0x79, 0x27,
	0x6e, 0x94, 0x69, 0xbb, 0x46, 0xab, 0x76, 0xf9, 0x90, 0xae, 0xcd, 0xbb, 0xab, 0xf6, 0x0f, 0x03,
	0x35, 0x3c, 0x11, 0x2d, 0xac, 0xb5, 0xbd, 0x5a, 0x1e, 0xe8, 0x6e, 0x7a, 0xdf, 0xbb, 0xc6, 0x8b,
	0xad, 0x9c, 0x16, 0x88, 0x21, 0x8b, 0x03, 0x22, 0x92, 0xc0, 0x0e, 0x78, 0x9c, 0x5d, 0xbf, 0x3d,
	0x35, 0x9a, 0xff, 0xd7, 0xf7, 0x50, 0x07, 0x1f, 0xcc, 0x4a, 0x87, 0xd2, 0xcf, 0x66, 0xbd, 0xa3,
	0x04, 0xa9, 0x0f, 0x44, 0x85, 0x69, 0xd4, 0x73, 0x48, 0x6e, 0x0c, 0x87, 0x1a, 0xd2, 0xa7, 0x3e,
	0xf4, 0x0b, 0x48, 0xbf, 0xe7, 0xf4, 0x35, 0xe4, 0xbb, 0xd9, 0x50, 0xdf, 0x5d, 0x97, 0xfa, 0xe0,
	0xba, 0x05, 0xc8, 0x75, 0x7b, 0x8e, 0xeb, 0x6a, 0xd8, 0xde, 0x52, 0x56, 0xe7, 0xc6, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb9, 0xa2, 0xe7, 0xd4, 0xa1, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Returns the requested asset in full detail.
	GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*resources.Asset, error)
	// Creates assets. Operation statuses are returned.
	MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*resources.Asset, error) {
	out := new(resources.Asset)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AssetService/GetAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error) {
	out := new(MutateAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AssetService/MutateAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Returns the requested asset in full detail.
	GetAsset(context.Context, *GetAssetRequest) (*resources.Asset, error)
	// Creates assets. Operation statuses are returned.
	MutateAssets(context.Context, *MutateAssetsRequest) (*MutateAssetsResponse, error)
}

// UnimplementedAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (*UnimplementedAssetServiceServer) GetAsset(ctx context.Context, req *GetAssetRequest) (*resources.Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAsset not implemented")
}
func (*UnimplementedAssetServiceServer) MutateAssets(ctx context.Context, req *MutateAssetsRequest) (*MutateAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssets not implemented")
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_GetAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).GetAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AssetService/GetAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).GetAsset(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_MutateAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).MutateAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AssetService/MutateAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).MutateAssets(ctx, req.(*MutateAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAsset",
			Handler:    _AssetService_GetAsset_Handler,
		},
		{
			MethodName: "MutateAssets",
			Handler:    _AssetService_MutateAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/asset_service.proto",
}
