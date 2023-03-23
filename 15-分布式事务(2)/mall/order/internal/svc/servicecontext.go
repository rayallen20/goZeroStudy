package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"order/internal/config"
	"user/userclient"
)

type ServiceContext struct {
	Config config.Config
	// UserRpc 实际上是user模块中的userclient/user.go中定义的接口User
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 其实userclient.NewUser()函数返回的接口User的实现 就是user/types/user_grpc.pb.go中定义的接口UserClient
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
