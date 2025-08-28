package main

import (
	_ "goframe-project/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"

	"goframe-project/internal/controller"
)

func main() {
	ctx := gctx.GetInitCtx()
	
	s := g.Server()
	
	// 注册路由
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			controller.Hello,
		)
	})
	
	// 启动服务器
	s.SetPort(8080)
	s.Run()
	
	g.Log().Info(ctx, "Server started on port 8080")
}