package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"userapi/internal/bizError"

	"userapi/internal/config"
	"userapi/internal/handler"
	"userapi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 入口文件处设置错误处理函数 该函数作用于所有路由
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (httpStatus int, errResp any) {
		switch e := err.(type) {
		case *bizError.BizError:
			errResp = e.GenResp()
			return http.StatusOK, errResp
		default:
			return http.StatusInternalServerError, nil
		}
	})

	httpx.SetErrorHandler(func(err error) (httpStatus int, errResp any) {
		switch e := err.(type) {
		case *bizError.BizError:
			errResp = e.GenResp()
			return http.StatusOK, errResp
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
