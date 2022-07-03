package cmd

import (
	"context"
	"threebody/internal/controller/backend"
	"threebody/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"threebody/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().CORS)
				group.Bind(
					backend.User,
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().CORS, service.Middleware().Auth)
				group.Bind(
					controller.Hello,
					backend.Role,
				)
			})
			s.Run()
			return nil
		},
	}
)
