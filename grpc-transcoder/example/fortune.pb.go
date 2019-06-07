// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/fortune.proto

package build_stack_fortune

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PredictionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictionRequest) Reset()         { *m = PredictionRequest{} }
func (m *PredictionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionRequest) ProtoMessage()    {}
func (*PredictionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a322b6a93c45146, []int{0}
}

func (m *PredictionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionRequest.Unmarshal(m, b)
}
func (m *PredictionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionRequest.Marshal(b, m, deterministic)
}
func (m *PredictionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionRequest.Merge(m, src)
}
func (m *PredictionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionRequest.Size(m)
}
func (m *PredictionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionRequest proto.InternalMessageInfo

type PredictionResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictionResponse) Reset()         { *m = PredictionResponse{} }
func (m *PredictionResponse) String() string { return proto.CompactTextString(m) }
func (*PredictionResponse) ProtoMessage()    {}
func (*PredictionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a322b6a93c45146, []int{1}
}

func (m *PredictionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionResponse.Unmarshal(m, b)
}
func (m *PredictionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionResponse.Marshal(b, m, deterministic)
}
func (m *PredictionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionResponse.Merge(m, src)
}
func (m *PredictionResponse) XXX_Size() int {
	return xxx_messageInfo_PredictionResponse.Size(m)
}
func (m *PredictionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionResponse proto.InternalMessageInfo

func (m *PredictionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PredictionRequest)(nil), "build.stack.fortune.PredictionRequest")
	proto.RegisterType((*PredictionResponse)(nil), "build.stack.fortune.PredictionResponse")
}

func init() { proto.RegisterFile("example/fortune.proto", fileDescriptor_2a322b6a93c45146) }

var fileDescriptor_2a322b6a93c45146 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x4f, 0xcb, 0x2f, 0x2a, 0x29, 0xcd, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x4e, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x2b, 0x2e, 0x49, 0x4c, 0xce, 0xd6, 0x83,
	0x4a, 0x29, 0x09, 0x73, 0x09, 0x06, 0x14, 0xa5, 0xa6, 0x64, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x05,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xe9, 0x71, 0x09, 0x21, 0x0b, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xd9, 0x5c, 0xbc, 0x6e, 0x10, 0xf3, 0x42, 0x52, 0x73,
	0x72, 0x52, 0x8b, 0x84, 0xa2, 0xb8, 0xd8, 0xa1, 0x06, 0x08, 0xa9, 0xe9, 0x61, 0xb1, 0x56, 0x0f,
	0xc3, 0x4e, 0x29, 0x75, 0x82, 0xea, 0x20, 0xce, 0x48, 0x62, 0x03, 0xfb, 0xc6, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xae, 0x6b, 0xa0, 0xe6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FortuneTellerClient is the client API for FortuneTeller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FortuneTellerClient interface {
	Predict(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error)
}

type fortuneTellerClient struct {
	cc *grpc.ClientConn
}

func NewFortuneTellerClient(cc *grpc.ClientConn) FortuneTellerClient {
	return &fortuneTellerClient{cc}
}

func (c *fortuneTellerClient) Predict(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*PredictionResponse, error) {
	out := new(PredictionResponse)
	err := c.cc.Invoke(ctx, "/build.stack.fortune.FortuneTeller/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FortuneTellerServer is the server API for FortuneTeller service.
type FortuneTellerServer interface {
	Predict(context.Context, *PredictionRequest) (*PredictionResponse, error)
}

// UnimplementedFortuneTellerServer can be embedded to have forward compatible implementations.
type UnimplementedFortuneTellerServer struct {
}

func (*UnimplementedFortuneTellerServer) Predict(ctx context.Context, req *PredictionRequest) (*PredictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}

func RegisterFortuneTellerServer(s *grpc.Server, srv FortuneTellerServer) {
	s.RegisterService(&_FortuneTeller_serviceDesc, srv)
}

func _FortuneTeller_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FortuneTellerServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build.stack.fortune.FortuneTeller/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FortuneTellerServer).Predict(ctx, req.(*PredictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FortuneTeller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "build.stack.fortune.FortuneTeller",
	HandlerType: (*FortuneTellerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _FortuneTeller_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/fortune.proto",
}
