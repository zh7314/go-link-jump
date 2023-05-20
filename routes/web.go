package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	sys_middleware "github.com/goravel/framework/http/middleware"
	"goravel/app/http/controllers/admin"
	"goravel/app/http/controllers/api"
	"goravel/app/http/middleware"
)

func Web() {

	//api
	/*
		调用多个中间件 中间件是顺序执行
	*/
	facades.Route.Middleware(sys_middleware.Cors(), middleware.Recovery()).Group(func(route route.Route) {
		//测试
		route.Get("/test", api.NewTestController().Test)
		//注意 不要用 facades.Route.Post("/getData", admin.NewJumpController().GetData) 在下来，因为闭包导致中间件无法执行
		route.Get("/p/{id}", api.NewJumpController().DoJump) //链接跳转

		route.Post("/getToken", admin.NewApiController().GetToken) //获取token
	})

	//admin
	/*
		分组 加上调用多个中间件 中间件是顺序执行
	*/
	facades.Route.Prefix("admin").Middleware(sys_middleware.Cors(), middleware.AdminCheck(), middleware.Recovery()).Group(func(route route.Route) {
		//注意 不要用 facades.Route.Post("/getData", admin.NewJumpController().GetData) 在下来，因为闭包导致中间件无法执行
		route.Post("/getData", admin.NewJumpController().GetData) //获取跳转数据
		route.Post("/AddLink", admin.NewJumpController().AddLink) //增加链接
	})
}
