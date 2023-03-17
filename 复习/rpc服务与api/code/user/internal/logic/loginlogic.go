package logic

import (
	"context"
	"user/internal/svc"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	resp := &user.LoginResponse{
		Code:    200,
		Message: "请求成功",
		Data: &user.LoginData{
			Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.AHsIKlhDgwkzQOn_MDlnsA2KqKRHcpOea3pZpgLj4pw",
			Role: &user.Role{
				Id:   1,
				Name: "super_admin",
			},
		},
	}
	return resp, nil
}
