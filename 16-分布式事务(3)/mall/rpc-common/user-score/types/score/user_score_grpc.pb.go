// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: user_score.proto

package score

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

const (
	UserScore_SaveUserScore_FullMethodName         = "/score.UserScore/saveUserScore"
	UserScore_SaveUserScoreCallback_FullMethodName = "/score.UserScore/saveUserScoreCallback"
)

// UserScoreClient is the client API for UserScore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserScoreClient interface {
	SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	// saveUserScoreCallback 分布式事务中提交失败之后 回滚时需要执行一个callback
	SaveUserScoreCallback(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
}

type userScoreClient struct {
	cc grpc.ClientConnInterface
}

func NewUserScoreClient(cc grpc.ClientConnInterface) UserScoreClient {
	return &userScoreClient{cc}
}

func (c *userScoreClient) SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, UserScore_SaveUserScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userScoreClient) SaveUserScoreCallback(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, UserScore_SaveUserScoreCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserScoreServer is the server API for UserScore service.
// All implementations must embed UnimplementedUserScoreServer
// for forward compatibility
type UserScoreServer interface {
	SaveUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	// saveUserScoreCallback 分布式事务中提交失败之后 回滚时需要执行一个callback
	SaveUserScoreCallback(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	mustEmbedUnimplementedUserScoreServer()
}

// UnimplementedUserScoreServer must be embedded to have forward compatible implementations.
type UnimplementedUserScoreServer struct {
}

func (UnimplementedUserScoreServer) SaveUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserScore not implemented")
}
func (UnimplementedUserScoreServer) SaveUserScoreCallback(context.Context, *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserScoreCallback not implemented")
}
func (UnimplementedUserScoreServer) mustEmbedUnimplementedUserScoreServer() {}

// UnsafeUserScoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserScoreServer will
// result in compilation errors.
type UnsafeUserScoreServer interface {
	mustEmbedUnimplementedUserScoreServer()
}

func RegisterUserScoreServer(s grpc.ServiceRegistrar, srv UserScoreServer) {
	s.RegisterService(&UserScore_ServiceDesc, srv)
}

func _UserScore_SaveUserScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserScoreServer).SaveUserScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserScore_SaveUserScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserScoreServer).SaveUserScore(ctx, req.(*UserScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserScore_SaveUserScoreCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserScoreServer).SaveUserScoreCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserScore_SaveUserScoreCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserScoreServer).SaveUserScoreCallback(ctx, req.(*UserScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserScore_ServiceDesc is the grpc.ServiceDesc for UserScore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserScore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "score.UserScore",
	HandlerType: (*UserScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveUserScore",
			Handler:    _UserScore_SaveUserScore_Handler,
		},
		{
			MethodName: "saveUserScoreCallback",
			Handler:    _UserScore_SaveUserScoreCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_score.proto",
}
