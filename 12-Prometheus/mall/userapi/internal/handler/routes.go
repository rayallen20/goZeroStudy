// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"userapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	handler := NewUserHandler(serverCtx)
	// 全局中间件
	server.Use(serverCtx.UserMiddleware.Global)
	// 不需要JWT认证的路由
	server.AddRoutes(
		rest.WithMiddlewares(
			// 路由中间件
			[]rest.Middleware{serverCtx.UserMiddleware.LoginAndRegister},
			[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: handler.Register,
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: handler.Login,
			},
		}...,),
	)

	// 需要JWT认证的路由
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/get/:id",
				Handler: handler.GetUser,
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.Secret),
	)
}
