// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/customer_client_service.proto

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

// Request message for [CustomerClientService.GetCustomerClient][google.ads.googleads.v2.services.CustomerClientService.GetCustomerClient].
type GetCustomerClientRequest struct {
	// The resource name of the client to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientRequest) Reset()         { *m = GetCustomerClientRequest{} }
func (m *GetCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientRequest) ProtoMessage()    {}
func (*GetCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfd3ae0efb37b5a5, []int{0}
}

func (m *GetCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientRequest.Merge(m, src)
}
func (m *GetCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientRequest.Size(m)
}
func (m *GetCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientRequest proto.InternalMessageInfo

func (m *GetCustomerClientRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientRequest)(nil), "google.ads.googleads.v2.services.GetCustomerClientRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/customer_client_service.proto", fileDescriptor_dfd3ae0efb37b5a5)
}

var fileDescriptor_dfd3ae0efb37b5a5 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc6, 0x49, 0x04, 0xc1, 0xa0, 0x83, 0x01, 0xb1, 0x44, 0x87, 0x52, 0x3b, 0x48, 0x87, 0x3b,
	0x9a, 0x0e, 0xe2, 0x89, 0x4a, 0xda, 0xa1, 0x4e, 0x52, 0x2a, 0x74, 0x90, 0x40, 0x39, 0x93, 0x23,
	0x04, 0x9a, 0xbb, 0x9a, 0xff, 0xb5, 0x8b, 0x38, 0xe8, 0x2b, 0xf8, 0x06, 0x8e, 0xbe, 0x83, 0x2f,
	0xd0, 0xd5, 0x57, 0x70, 0x72, 0xf6, 0x01, 0x24, 0xbd, 0x5c, 0xda, 0x62, 0x43, 0xb7, 0x8f, 0xfb,
	0x7f, 0xbf, 0xff, 0x77, 0xf7, 0x25, 0xd6, 0x55, 0x24, 0x44, 0x34, 0x62, 0x98, 0x86, 0x80, 0x95,
	0xcc, 0xd4, 0xd4, 0xc5, 0xc0, 0xd2, 0x69, 0x1c, 0x30, 0xc0, 0xc1, 0x04, 0xa4, 0x48, 0x58, 0x3a,
	0x0c, 0x46, 0x31, 0xe3, 0x72, 0x98, 0x0f, 0xd0, 0x38, 0x15, 0x52, 0xd8, 0x55, 0x05, 0x21, 0x1a,
	0x02, 0x2a, 0x78, 0x34, 0x75, 0x91, 0xe6, 0x9d, 0xb3, 0xb2, 0x84, 0x94, 0x81, 0x98, 0xa4, 0x6b,
	0x22, 0xd4, 0x6a, 0xe7, 0x58, 0x83, 0xe3, 0x18, 0x53, 0xce, 0x85, 0xa4, 0x32, 0x16, 0x1c, 0xf2,
	0xe9, 0xe1, 0xd2, 0x74, 0x19, 0xab, 0x5d, 0x5b, 0x95, 0x2e, 0x93, 0x9d, 0x7c, 0x65, 0x67, 0x3e,
	0xea, 0xb3, 0xc7, 0x09, 0x03, 0x69, 0x9f, 0x58, 0x7b, 0x3a, 0x75, 0xc8, 0x69, 0xc2, 0x2a, 0x46,
	0xd5, 0x38, 0xdd, 0xe9, 0xef, 0xea, 0xc3, 0x5b, 0x9a, 0x30, 0xf7, 0xd7, 0xb0, 0x0e, 0x56, 0xf1,
	0x3b, 0xf5, 0x16, 0xfb, 0xd3, 0xb0, 0xf6, 0xff, 0xed, 0xb6, 0x09, 0xda, 0xd4, 0x01, 0x2a, 0xbb,
	0x90, 0xd3, 0x2c, 0x65, 0x8b, 0x76, 0xd0, 0x2a, 0x59, 0x3b, 0x7f, 0xfd, 0xfa, 0x7e, 0x33, 0x5b,
	0x76, 0x33, 0xeb, 0xf0, 0x69, 0xe5, 0x39, 0x97, 0xba, 0x48, 0xc0, 0x8d, 0xa2, 0x54, 0x85, 0x01,
	0x6e, 0x3c, 0x3b, 0x47, 0x33, 0xaf, 0xb2, 0x08, 0xc9, 0xd5, 0x38, 0x06, 0x14, 0x88, 0xa4, 0xfd,
	0x62, 0x5a, 0xf5, 0x40, 0x24, 0x1b, 0x1f, 0xd3, 0x76, 0xd6, 0x96, 0xd3, 0xcb, 0xca, 0xef, 0x19,
	0xf7, 0x37, 0x39, 0x1f, 0x89, 0x11, 0xe5, 0x11, 0x12, 0x69, 0x84, 0x23, 0xc6, 0xe7, 0x9f, 0x06,
	0x2f, 0x12, 0xcb, 0xff, 0xb7, 0x0b, 0x2d, 0xde, 0xcd, 0xad, 0xae, 0xe7, 0x7d, 0x98, 0xd5, 0xae,
	0x5a, 0xe8, 0x85, 0x80, 0x94, 0xcc, 0xd4, 0xc0, 0x45, 0x79, 0x30, 0xcc, 0xb4, 0xc5, 0xf7, 0x42,
	0xf0, 0x0b, 0x8b, 0x3f, 0x70, 0x7d, 0x6d, 0xf9, 0x31, 0xeb, 0xea, 0x9c, 0x10, 0x2f, 0x04, 0x42,
	0x0a, 0x13, 0x21, 0x03, 0x97, 0x10, 0x6d, 0x7b, 0xd8, 0x9e, 0xdf, 0xb3, 0xf5, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x6a, 0xf6, 0x6d, 0xee, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClientServiceClient is the client API for CustomerClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientServiceClient interface {
	// Returns the requested client in full detail.
	GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error)
}

type customerClientServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClientServiceClient(cc *grpc.ClientConn) CustomerClientServiceClient {
	return &customerClientServiceClient{cc}
}

func (c *customerClientServiceClient) GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error) {
	out := new(resources.CustomerClient)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CustomerClientService/GetCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientServiceServer is the server API for CustomerClientService service.
type CustomerClientServiceServer interface {
	// Returns the requested client in full detail.
	GetCustomerClient(context.Context, *GetCustomerClientRequest) (*resources.CustomerClient, error)
}

// UnimplementedCustomerClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientServiceServer struct {
}

func (*UnimplementedCustomerClientServiceServer) GetCustomerClient(ctx context.Context, req *GetCustomerClientRequest) (*resources.CustomerClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClient not implemented")
}

func RegisterCustomerClientServiceServer(s *grpc.Server, srv CustomerClientServiceServer) {
	s.RegisterService(&_CustomerClientService_serviceDesc, srv)
}

func _CustomerClientService_GetCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CustomerClientService/GetCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, req.(*GetCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CustomerClientService",
	HandlerType: (*CustomerClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClient",
			Handler:    _CustomerClientService_GetCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/customer_client_service.proto",
}
