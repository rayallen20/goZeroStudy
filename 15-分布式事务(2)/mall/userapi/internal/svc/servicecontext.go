package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"rpc-common/user-score/userscore"
	"rpc-common/user/userclient"
	"userapi/internal/config"
	"userapi/internal/middlewares"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userclient.User
	UserScoreRpc   userscore.UserScore
	UserMiddleware *middlewares.UserMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserScoreRpc:   userscore.NewUserScore(zrpc.MustNewClient(c.UserScoreRpc)),
		UserMiddleware: middlewares.NewUserMiddleware(),
	}
}
