// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleClient interface {
	Ascii(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResult, error)
	AsciiStream(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (Example_AsciiStreamClient, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Ascii(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (*ExampleResult, error) {
	out := new(ExampleResult)
	err := c.cc.Invoke(ctx, "/grpc.Example/ascii", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) AsciiStream(ctx context.Context, in *ExampleRequest, opts ...grpc.CallOption) (Example_AsciiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Example_ServiceDesc.Streams[0], "/grpc.Example/asciiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleAsciiStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Example_AsciiStreamClient interface {
	Recv() (*ExampleResult, error)
	grpc.ClientStream
}

type exampleAsciiStreamClient struct {
	grpc.ClientStream
}

func (x *exampleAsciiStreamClient) Recv() (*ExampleResult, error) {
	m := new(ExampleResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServer is the server API for Example service.
// All implementations must embed UnimplementedExampleServer
// for forward compatibility
type ExampleServer interface {
	Ascii(context.Context, *ExampleRequest) (*ExampleResult, error)
	AsciiStream(*ExampleRequest, Example_AsciiStreamServer) error
	mustEmbedUnimplementedExampleServer()
}

// UnimplementedExampleServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (UnimplementedExampleServer) Ascii(context.Context, *ExampleRequest) (*ExampleResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ascii not implemented")
}
func (UnimplementedExampleServer) AsciiStream(*ExampleRequest, Example_AsciiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AsciiStream not implemented")
}
func (UnimplementedExampleServer) mustEmbedUnimplementedExampleServer() {}

// UnsafeExampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServer will
// result in compilation errors.
type UnsafeExampleServer interface {
	mustEmbedUnimplementedExampleServer()
}

func RegisterExampleServer(s grpc.ServiceRegistrar, srv ExampleServer) {
	s.RegisterService(&Example_ServiceDesc, srv)
}

func _Example_Ascii_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Ascii(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Example/ascii",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Ascii(ctx, req.(*ExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_AsciiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExampleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServer).AsciiStream(m, &exampleAsciiStreamServer{stream})
}

type Example_AsciiStreamServer interface {
	Send(*ExampleResult) error
	grpc.ServerStream
}

type exampleAsciiStreamServer struct {
	grpc.ServerStream
}

func (x *exampleAsciiStreamServer) Send(m *ExampleResult) error {
	return x.ServerStream.SendMsg(m)
}

// Example_ServiceDesc is the grpc.ServiceDesc for Example service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Example_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ascii",
			Handler:    _Example_Ascii_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "asciiStream",
			Handler:       _Example_AsciiStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
