package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	sys_middleware "github.com/goravel/framework/http/middleware"
	"goravel/app/http/controllers/admin"
	"goravel/app/http/controllers/api"
	"goravel/app/http/middleware"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})
	//api
	facades.Route.Middleware(middleware.Recovery()).Middleware(sys_middleware.Cors()).Middleware(middleware.ApiLog()).Group(func(route route.Route) {

		//测试
		facades.Route.Get("/test", api.NewTestController().Test)
		facades.Route.Get("/p/{id}", api.NewJumpController().DoJump) //链接跳转
	})

	//admin
	facades.Route.Middleware(sys_middleware.Cors()).Middleware(middleware.ApiLog()).Group(func(route route.Route) {

		facades.Route.Middleware().Get("/admin/getData", admin.NewJumpController().GetData) //获取跳转数据
		facades.Route.Post("/admin/AddLink", admin.NewJumpController().AddLink)             //增加链接
	})

}
