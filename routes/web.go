package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers/admin"
	"goravel/app/http/controllers/api"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})

	//facades.Route.Get("/p/{id}", controllers.NewJumpController().doJump)
	//jumpController := api.NewJumpController()
	//facades.Route.Get("/p/{id}", jumpController.DoJump)

	facades.Route.Get("/p/{id}", api.NewJumpController().DoJump)

	facades.Route.Get("/admin/getData", admin.NewJumpController().GetData)

	facades.Route.Get("/admin/AddLink", admin.NewJumpController().AddLink)

	//facades.Route.Prefix("api").Get("/", userController.Show)

	facades.Route.Group(func(route route.Route) {
		route.Get("group/{id}", api.NewJumpController().DoJump)
	})
}
