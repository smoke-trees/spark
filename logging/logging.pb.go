// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logging.proto

package main

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

func init() { proto.RegisterFile("logging.proto", fileDescriptor_9c8ad1e4de00dd2b) }

var fileDescriptor_9c8ad1e4de00dd2b = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xc9, 0x4f, 0x4f,
	0xcf, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x84, 0x8a,
	0x73, 0x13, 0x8b, 0x4a, 0x72, 0x53, 0x4b, 0x52, 0x8b, 0x52, 0x92, 0x20, 0x92, 0x46, 0x01, 0x5c,
	0x7c, 0x3e, 0x10, 0xe9, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x3b, 0x2e, 0x76, 0x97,
	0xc4, 0x92, 0x44, 0x9f, 0xfc, 0x74, 0x21, 0x09, 0xbd, 0x94, 0xc4, 0x92, 0xc4, 0xa4, 0xc4, 0xe2,
	0x54, 0x3d, 0xa8, 0x50, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x24, 0x16, 0x99, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x27, 0xb6, 0x28, 0x96, 0xdc, 0xc4, 0xcc, 0xbc, 0x24, 0x36, 0xb0,
	0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x61, 0x05, 0x6e, 0x8e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoggingServiceClient is the client API for LoggingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggingServiceClient interface {
	DataLog(ctx context.Context, in *DataLogRequest, opts ...grpc.CallOption) (*DataLogResponse, error)
}

type loggingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingServiceClient(cc grpc.ClientConnInterface) LoggingServiceClient {
	return &loggingServiceClient{cc}
}

func (c *loggingServiceClient) DataLog(ctx context.Context, in *DataLogRequest, opts ...grpc.CallOption) (*DataLogResponse, error) {
	out := new(DataLogResponse)
	err := c.cc.Invoke(ctx, "/logging.LoggingService/DataLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServiceServer is the server API for LoggingService service.
type LoggingServiceServer interface {
	DataLog(context.Context, *DataLogRequest) (*DataLogResponse, error)
}

// UnimplementedLoggingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceServer struct {
}

func (*UnimplementedLoggingServiceServer) DataLog(ctx context.Context, req *DataLogRequest) (*DataLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataLog not implemented")
}

func RegisterLoggingServiceServer(s *grpc.Server, srv LoggingServiceServer) {
	s.RegisterService(&_LoggingService_serviceDesc, srv)
}

func _LoggingService_DataLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).DataLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logging.LoggingService/DataLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).DataLog(ctx, req.(*DataLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logging.LoggingService",
	HandlerType: (*LoggingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataLog",
			Handler:    _LoggingService_DataLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logging.proto",
}
