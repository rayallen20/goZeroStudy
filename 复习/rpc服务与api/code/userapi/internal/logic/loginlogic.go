package logic

import (
	"context"
	"time"
	"user/types/user"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	loginReq := &user.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	}
	loginResp, err := l.svcCtx.UserRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResponse{
		Code:    loginResp.Code,
		Message: loginResp.Message,
		Data: &types.LoginData{
			Jwt: loginResp.Data.Jwt,
			Role: &types.Role{
				Id:   loginResp.Data.Role.Id,
				Name: loginResp.Data.Role.Name,
			},
		},
	}
	return resp, nil
}
