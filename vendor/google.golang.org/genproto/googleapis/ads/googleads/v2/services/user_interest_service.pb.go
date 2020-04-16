// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/user_interest_service.proto

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

// Request message for [UserInterestService.GetUserInterest][google.ads.googleads.v2.services.UserInterestService.GetUserInterest].
type GetUserInterestRequest struct {
	// Resource name of the UserInterest to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInterestRequest) Reset()         { *m = GetUserInterestRequest{} }
func (m *GetUserInterestRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInterestRequest) ProtoMessage()    {}
func (*GetUserInterestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d082b840ef8dbf, []int{0}
}

func (m *GetUserInterestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInterestRequest.Unmarshal(m, b)
}
func (m *GetUserInterestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInterestRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInterestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInterestRequest.Merge(m, src)
}
func (m *GetUserInterestRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInterestRequest.Size(m)
}
func (m *GetUserInterestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInterestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInterestRequest proto.InternalMessageInfo

func (m *GetUserInterestRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInterestRequest)(nil), "google.ads.googleads.v2.services.GetUserInterestRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/user_interest_service.proto", fileDescriptor_b9d082b840ef8dbf)
}

var fileDescriptor_b9d082b840ef8dbf = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x2e, 0x5c, 0xb8, 0xe1, 0x5e, 0x2e, 0xe4, 0xc2, 0xb5, 0x44, 0x17, 0xa5, 0x76,
	0x21, 0x5d, 0xcc, 0x60, 0x44, 0x94, 0xd1, 0x2e, 0xd2, 0x4d, 0x75, 0x23, 0xa5, 0x62, 0x17, 0x12,
	0x28, 0x31, 0x39, 0x84, 0x40, 0x33, 0x53, 0xe7, 0x4c, 0xba, 0x11, 0x37, 0xbe, 0x82, 0x6f, 0xe0,
	0xd2, 0xbd, 0x2f, 0xd1, 0xad, 0xaf, 0xe0, 0xaa, 0x2f, 0xa1, 0xa4, 0xc9, 0xa4, 0x55, 0x5b, 0xba,
	0xfb, 0x99, 0x39, 0xdf, 0xff, 0xcf, 0xf9, 0x13, 0xeb, 0x34, 0x16, 0x22, 0x1e, 0x01, 0x0d, 0x22,
	0xa4, 0x85, 0xcc, 0xd5, 0xc4, 0xa5, 0x08, 0x72, 0x92, 0x84, 0x80, 0x34, 0x43, 0x90, 0xc3, 0x84,
	0x2b, 0x90, 0x80, 0x6a, 0x58, 0x1e, 0x93, 0xb1, 0x14, 0x4a, 0xd8, 0xf5, 0x02, 0x21, 0x41, 0x84,
	0xa4, 0xa2, 0xc9, 0xc4, 0x25, 0x9a, 0x76, 0x0e, 0xd7, 0xf9, 0x4b, 0x40, 0x91, 0xc9, 0x6f, 0x01,
	0x85, 0xb1, 0xb3, 0xa3, 0xb1, 0x71, 0x42, 0x03, 0xce, 0x85, 0x0a, 0x54, 0x22, 0x38, 0x96, 0xb7,
	0x5b, 0x4b, 0xb7, 0xe1, 0x28, 0x01, 0x5e, 0x62, 0x8d, 0xb6, 0xf5, 0xbf, 0x0b, 0xea, 0x0a, 0x41,
	0x9e, 0x97, 0x7e, 0x7d, 0xb8, 0xcd, 0x00, 0x95, 0xbd, 0x6b, 0xfd, 0xd1, 0x89, 0x43, 0x1e, 0xa4,
	0x50, 0x33, 0xea, 0xc6, 0xde, 0xaf, 0xfe, 0x6f, 0x7d, 0x78, 0x11, 0xa4, 0xe0, 0xce, 0x0c, 0xeb,
	0xdf, 0x32, 0x7c, 0x59, 0x6c, 0x61, 0xbf, 0x18, 0xd6, 0xdf, 0x2f, 0xbe, 0xf6, 0x31, 0xd9, 0xb4,
	0x3b, 0x59, 0xfd, 0x14, 0x87, 0xae, 0x25, 0xab, 0x4e, 0xc8, 0x32, 0xd7, 0x38, 0x7a, 0x78, 0x7d,
	0x7b, 0x34, 0xf7, 0x6d, 0x9a, 0xf7, 0x76, 0xf7, 0x69, 0x8d, 0x76, 0x98, 0xa1, 0x12, 0x29, 0x48,
	0xa4, 0xad, 0x79, 0x91, 0x1a, 0x42, 0xda, 0xba, 0x77, 0xb6, 0xa7, 0x5e, 0x6d, 0x11, 0x50, 0xaa,
	0x71, 0x82, 0x24, 0x14, 0x69, 0xe7, 0xdd, 0xb0, 0x9a, 0xa1, 0x48, 0x37, 0xae, 0xd1, 0xa9, 0xad,
	0xa8, 0xa4, 0x97, 0xd7, 0xdd, 0x33, 0xae, 0xcf, 0x4a, 0x3a, 0x16, 0xa3, 0x80, 0xc7, 0x44, 0xc8,
	0x98, 0xc6, 0xc0, 0xe7, 0x1f, 0x83, 0x2e, 0xf2, 0xd6, 0xff, 0x5d, 0x27, 0x5a, 0x3c, 0x99, 0x3f,
	0xba, 0x9e, 0xf7, 0x6c, 0xd6, 0xbb, 0x85, 0xa1, 0x17, 0x21, 0x29, 0x64, 0xae, 0x06, 0x2e, 0x29,
	0x83, 0x71, 0xaa, 0x47, 0x7c, 0x2f, 0x42, 0xbf, 0x1a, 0xf1, 0x07, 0xae, 0xaf, 0x47, 0x66, 0x66,
	0xb3, 0x38, 0x67, 0xcc, 0x8b, 0x90, 0xb1, 0x6a, 0x88, 0xb1, 0x81, 0xcb, 0x98, 0x1e, 0xbb, 0xf9,
	0x39, 0x7f, 0xe7, 0xc1, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x42, 0x63, 0xbd, 0x04, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInterestServiceClient(cc grpc.ClientConnInterface) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
}

// UnimplementedUserInterestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserInterestServiceServer struct {
}

func (*UnimplementedUserInterestServiceServer) GetUserInterest(ctx context.Context, req *GetUserInterestRequest) (*resources.UserInterest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInterest not implemented")
}

func RegisterUserInterestServiceServer(s *grpc.Server, srv UserInterestServiceServer) {
	s.RegisterService(&_UserInterestService_serviceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInterestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/user_interest_service.proto",
}
