// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RunRequest struct {
	J                    string   `protobuf:"bytes,1,opt,name=j,proto3" json:"j,omitempty"`
	Q                    string   `protobuf:"bytes,2,opt,name=q,proto3" json:"q,omitempty"`
	O                    []string `protobuf:"bytes,3,rep,name=o,proto3" json:"o,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetJ() string {
	if m != nil {
		return m.J
	}
	return ""
}

func (m *RunRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *RunRequest) GetO() []string {
	if m != nil {
		return m.O
	}
	return nil
}

type RunResponse struct {
	Out                  string   `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func init() {
	proto.RegisterType((*RunRequest)(nil), "api.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "api.RunResponse")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x32, 0xe3, 0xe2, 0x0a, 0x2a, 0xcd, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0xcc, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcc,
	0x02, 0xf1, 0x0a, 0x25, 0x98, 0x20, 0xbc, 0x42, 0x10, 0x2f, 0x5f, 0x82, 0x59, 0x81, 0x19, 0xc4,
	0xcb, 0x57, 0x92, 0xe7, 0xe2, 0x06, 0xeb, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe0,
	0x62, 0xce, 0x2f, 0x2d, 0x81, 0x6a, 0x05, 0x31, 0x8d, 0xbc, 0xb8, 0x38, 0xbd, 0x02, 0x83, 0x53,
	0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x6c, 0xb9, 0xd8, 0x03, 0xf2, 0x8b, 0x4b, 0x82, 0x4a, 0xf3,
	0x84, 0xf8, 0xf5, 0x40, 0xee, 0x43, 0xd8, 0x29, 0x25, 0x80, 0x10, 0x80, 0x18, 0xa6, 0xc4, 0xdf,
	0x74, 0xf9, 0xc9, 0x64, 0x26, 0x4e, 0x25, 0x16, 0xfd, 0xa2, 0xd2, 0x3c, 0x2b, 0x46, 0xad, 0x24,
	0x36, 0xb0, 0x5b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x2d, 0x32, 0xc2, 0xdb, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JQServiceClient is the client API for JQService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JQServiceClient interface {
	PostRun(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type jQServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJQServiceClient(cc grpc.ClientConnInterface) JQServiceClient {
	return &jQServiceClient{cc}
}

func (c *jQServiceClient) PostRun(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/api.JQService/PostRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JQServiceServer is the server API for JQService service.
type JQServiceServer interface {
	PostRun(context.Context, *RunRequest) (*RunResponse, error)
}

// UnimplementedJQServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJQServiceServer struct {
}

func (*UnimplementedJQServiceServer) PostRun(ctx context.Context, req *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRun not implemented")
}

func RegisterJQServiceServer(s *grpc.Server, srv JQServiceServer) {
	s.RegisterService(&_JQService_serviceDesc, srv)
}

func _JQService_PostRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JQServiceServer).PostRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JQService/PostRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JQServiceServer).PostRun(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JQService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.JQService",
	HandlerType: (*JQServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostRun",
			Handler:    _JQService_PostRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}