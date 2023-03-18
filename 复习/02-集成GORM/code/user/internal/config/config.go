package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySQL MySQL
}

type MySQL struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Charset  string
}
