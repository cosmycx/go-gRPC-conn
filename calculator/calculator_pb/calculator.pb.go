// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculator_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Sum struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Sum) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FibonacciRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{3}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FibonacciResponse struct {
	ResultNum            int32    `protobuf:"varint,1,opt,name=resultNum,proto3" json:"resultNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{4}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetResultNum() int32 {
	if m != nil {
		return m.ResultNum
	}
	return 0
}

type MeanRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeanRequest) Reset()         { *m = MeanRequest{} }
func (m *MeanRequest) String() string { return proto.CompactTextString(m) }
func (*MeanRequest) ProtoMessage()    {}
func (*MeanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{5}
}

func (m *MeanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeanRequest.Unmarshal(m, b)
}
func (m *MeanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeanRequest.Marshal(b, m, deterministic)
}
func (m *MeanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeanRequest.Merge(m, src)
}
func (m *MeanRequest) XXX_Size() int {
	return xxx_messageInfo_MeanRequest.Size(m)
}
func (m *MeanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeanRequest proto.InternalMessageInfo

func (m *MeanRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type MeanResponse struct {
	Mean                 int64    `protobuf:"varint,1,opt,name=mean,proto3" json:"mean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeanResponse) Reset()         { *m = MeanResponse{} }
func (m *MeanResponse) String() string { return proto.CompactTextString(m) }
func (*MeanResponse) ProtoMessage()    {}
func (*MeanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{6}
}

func (m *MeanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeanResponse.Unmarshal(m, b)
}
func (m *MeanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeanResponse.Marshal(b, m, deterministic)
}
func (m *MeanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeanResponse.Merge(m, src)
}
func (m *MeanResponse) XXX_Size() int {
	return xxx_messageInfo_MeanResponse.Size(m)
}
func (m *MeanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeanResponse proto.InternalMessageInfo

func (m *MeanResponse) GetMean() int64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*FibonacciRequest)(nil), "calculator.FibonacciRequest")
	proto.RegisterType((*FibonacciResponse)(nil), "calculator.FibonacciResponse")
	proto.RegisterType((*MeanRequest)(nil), "calculator.MeanRequest")
	proto.RegisterType((*MeanResponse)(nil), "calculator.MeanResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x1d, 0x63, 0x03, 0xde, 0x56, 0x9a, 0xce, 0x22, 0x86, 0x50, 0xc1, 0x0e, 0x0a, 0x5d,
	0x55, 0xad, 0x0f, 0x20, 0xa8, 0xb8, 0xd3, 0x45, 0xb2, 0x73, 0x23, 0x33, 0x61, 0x16, 0x81, 0x64,
	0x12, 0x93, 0x8c, 0xaf, 0xed, 0x2b, 0xc8, 0xfc, 0x35, 0x63, 0x4b, 0x77, 0x73, 0xee, 0x3d, 0xf9,
	0x0e, 0xf7, 0x10, 0x88, 0x0a, 0x5a, 0x15, 0xb2, 0xa2, 0x43, 0xd3, 0x6d, 0xda, 0xae, 0x19, 0x1a,
	0x0c, 0xe3, 0x84, 0xac, 0x20, 0xc8, 0x65, 0x8d, 0x67, 0x80, 0x68, 0x82, 0xae, 0xd1, 0x7a, 0x92,
	0x21, 0xaa, 0x14, 0x4b, 0x4e, 0x8d, 0x62, 0xe4, 0x0e, 0x20, 0x97, 0x75, 0xc6, 0xbf, 0x25, 0xef,
	0x07, 0xbc, 0x82, 0xa0, 0x97, 0xb5, 0xf6, 0x4e, 0xb7, 0xf3, 0x8d, 0x07, 0x57, 0x26, 0xb5, 0x23,
	0xb7, 0x30, 0xd5, 0x1f, 0xf4, 0x6d, 0x23, 0x7a, 0x8e, 0x63, 0x08, 0x3b, 0xde, 0xcb, 0x6a, 0xb0,
	0x01, 0x56, 0x91, 0x1b, 0x88, 0xde, 0x4a, 0xd6, 0x08, 0x5a, 0x14, 0xa5, 0xa3, 0x47, 0x10, 0x08,
	0x4b, 0x9f, 0x64, 0xea, 0x49, 0x1e, 0x60, 0xe1, 0xb9, 0x2c, 0x72, 0x09, 0xe7, 0x06, 0xf2, 0xb1,
	0x33, 0x8f, 0x03, 0x95, 0xff, 0xce, 0xa9, 0x70, 0xcc, 0x18, 0x42, 0x21, 0x6b, 0xc6, 0x3b, 0xed,
	0x0c, 0x32, 0xab, 0x08, 0x81, 0x99, 0xb1, 0x59, 0x28, 0x86, 0xb3, 0x9a, 0x53, 0x61, 0x5d, 0xfa,
	0xbd, 0xfd, 0x45, 0xb0, 0x78, 0xd9, 0x9d, 0x98, 0xf3, 0xee, 0xa7, 0x2c, 0x38, 0x7e, 0xd2, 0x8d,
	0x38, 0x15, 0xef, 0x97, 0x60, 0x72, 0xd3, 0xcb, 0x83, 0xb9, 0x09, 0x22, 0x27, 0x38, 0xf7, 0x4e,
	0x77, 0x98, 0xa5, 0x6f, 0xdf, 0x2f, 0x26, 0xbd, 0x3a, 0xb2, 0x75, 0xc8, 0x7b, 0x84, 0x5f, 0xcd,
	0xd9, 0x8e, 0xf7, 0x2f, 0xde, 0xeb, 0x23, 0x4d, 0x0e, 0x17, 0x8e, 0xb2, 0x46, 0xcf, 0xf3, 0xcf,
	0x8b, 0x71, 0xfd, 0xd5, 0x32, 0x16, 0xea, 0x9f, 0xe6, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0x65, 0x81, 0x55, 0x48, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	SumService(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	FibonacciService(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciServiceClient, error)
	MeanService(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MeanServiceClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) SumService(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/SumService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) FibonacciService(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (CalculatorService_FibonacciServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/FibonacciService", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFibonacciServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_FibonacciServiceClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type calculatorServiceFibonacciServiceClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFibonacciServiceClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) MeanService(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MeanServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/MeanService", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMeanServiceClient{stream}
	return x, nil
}

type CalculatorService_MeanServiceClient interface {
	Send(*MeanRequest) error
	CloseAndRecv() (*MeanResponse, error)
	grpc.ClientStream
}

type calculatorServiceMeanServiceClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMeanServiceClient) Send(m *MeanRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMeanServiceClient) CloseAndRecv() (*MeanResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MeanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	SumService(context.Context, *SumRequest) (*SumResponse, error)
	FibonacciService(*FibonacciRequest, CalculatorService_FibonacciServiceServer) error
	MeanService(CalculatorService_MeanServiceServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_SumService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SumService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/SumService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SumService(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_FibonacciService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).FibonacciService(m, &calculatorServiceFibonacciServiceServer{stream})
}

type CalculatorService_FibonacciServiceServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type calculatorServiceFibonacciServiceServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFibonacciServiceServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_MeanService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).MeanService(&calculatorServiceMeanServiceServer{stream})
}

type CalculatorService_MeanServiceServer interface {
	SendAndClose(*MeanResponse) error
	Recv() (*MeanRequest, error)
	grpc.ServerStream
}

type calculatorServiceMeanServiceServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMeanServiceServer) SendAndClose(m *MeanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMeanServiceServer) Recv() (*MeanRequest, error) {
	m := new(MeanRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumService",
			Handler:    _CalculatorService_SumService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FibonacciService",
			Handler:       _CalculatorService_FibonacciService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MeanService",
			Handler:       _CalculatorService_MeanService_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
