// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"
	"user/internal/logic"
	"user/internal/svc"
	user2 "user/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user2.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user2.LoginRequest) (*user2.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
