// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: test.proto

package testv1

import (
	context "context"
	subtest "github.com/sagikazarmark/protoc-gen-go-kit/test/subtest"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestService_TestProcedure_FullMethodName   = "/test.v1.TestService/TestProcedure"
	TestService_SubestProcedure_FullMethodName = "/test.v1.TestService/SubestProcedure"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	// TestProcedure is here for the sake of testing comments.
	TestProcedure(ctx context.Context, in *TestProcedureRequest, opts ...grpc.CallOption) (*TestProcedureResponse, error)
	SubestProcedure(ctx context.Context, in *subtest.SubtestProcedureRequest, opts ...grpc.CallOption) (*subtest.SubtestProcedureResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestProcedure(ctx context.Context, in *TestProcedureRequest, opts ...grpc.CallOption) (*TestProcedureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestProcedureResponse)
	err := c.cc.Invoke(ctx, TestService_TestProcedure_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SubestProcedure(ctx context.Context, in *subtest.SubtestProcedureRequest, opts ...grpc.CallOption) (*subtest.SubtestProcedureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(subtest.SubtestProcedureResponse)
	err := c.cc.Invoke(ctx, TestService_SubestProcedure_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility.
type TestServiceServer interface {
	// TestProcedure is here for the sake of testing comments.
	TestProcedure(context.Context, *TestProcedureRequest) (*TestProcedureResponse, error)
	SubestProcedure(context.Context, *subtest.SubtestProcedureRequest) (*subtest.SubtestProcedureResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) TestProcedure(context.Context, *TestProcedureRequest) (*TestProcedureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestProcedure not implemented")
}
func (UnimplementedTestServiceServer) SubestProcedure(context.Context, *subtest.SubtestProcedureRequest) (*subtest.SubtestProcedureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubestProcedure not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}
func (UnimplementedTestServiceServer) testEmbeddedByValue()                     {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestProcedure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestProcedureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestProcedure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_TestProcedure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestProcedure(ctx, req.(*TestProcedureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SubestProcedure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(subtest.SubtestProcedureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SubestProcedure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_SubestProcedure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SubestProcedure(ctx, req.(*subtest.SubtestProcedureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestProcedure",
			Handler:    _TestService_TestProcedure_Handler,
		},
		{
			MethodName: "SubestProcedure",
			Handler:    _TestService_SubestProcedure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
