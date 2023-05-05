package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Zx(ctx http.Context) {

	var userService services.UserService
	data := userService.Zx("ssss")

	ctx.Response().Success().Json(data)
}
