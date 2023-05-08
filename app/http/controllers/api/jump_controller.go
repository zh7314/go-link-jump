package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
)

type JumpController struct {
	//Dependent services
}

func NewJumpController() *JumpController {
	return &JumpController{
		//Inject services
	}
}

func (r *JumpController) DoJump(ctx http.Context) {

	//ctx.Request().
	id := ctx.Request().Input("id")
	fmt.Println(id)

	jumpService := services.NewJumpService()

	data := jumpService.GetData()

	//db, err := facades.Orm.Connection("mysql").DB()

	ctx.Response().Success().Json(data)
}
