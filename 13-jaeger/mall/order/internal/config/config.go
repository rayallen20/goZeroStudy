package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// UserRpc user rpc服务的配置
	// 即:order-api.yaml中的 UserRpc字段下的配置项
	UserRpc zrpc.RpcClientConf
}
