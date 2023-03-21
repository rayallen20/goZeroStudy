package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL      MySQLConfig
	CacheRedis cache.CacheConf
}

type MySQLConfig struct {
	DataSource string
}
