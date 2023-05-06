package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type UserController struct {
	//Dependent services
	userService services.UserService
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

	//userService := services.NewUserService()

	//var userService services.UserService
	data := r.userService.Zx("ssss")

	//r := reflect.TypeOf(data)
	//
	//fmt.Println(r.Name(), r.Kind())

	ctx.Response().Success().Json(data)
}
