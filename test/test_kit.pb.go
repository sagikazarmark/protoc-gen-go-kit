// Code generated by protoc-gen-go-kit. DO NOT EDIT.
// versions:
// - protoc-gen-go-kit v0.6.0
// - protoc            v5.29.1

package testv1

import (
	context "context"
	subtest "github.com/sagikazarmark/protoc-gen-go-kit/test/subtest"
)

// TestServiceHandler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
//
// This interface is based on github.com/go-kit/kit/transport/grpc.Handler.
type TestServiceHandler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

// TestServiceKitServer is the Go kit server implementation for TestService service.
type TestServiceKitServer struct {
	UnimplementedTestServiceServer

	TestProcedureHandler   TestServiceHandler
	SubestProcedureHandler TestServiceHandler
}

// TestProcedure is here for the sake of testing comments.
func (s TestServiceKitServer) TestProcedure(ctx context.Context, req *TestProcedureRequest) (*TestProcedureResponse, error) {
	_, resp, err := s.TestProcedureHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*TestProcedureResponse), nil
}

func (s TestServiceKitServer) SubestProcedure(ctx context.Context, req *subtest.SubtestProcedureRequest) (*subtest.SubtestProcedureResponse, error) {
	_, resp, err := s.SubestProcedureHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*subtest.SubtestProcedureResponse), nil
}
