package admin

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

func (r *JumpController) GetData(ctx http.Context) {

	jumpService := services.NewJumpService()
	data := jumpService.GetData()

	ctx.Response().Success().Json(data)
}

func (r *JumpController) AddLink(ctx http.Context) {

	url := ctx.Request().Input("url")
	end_time := ctx.Request().Input("end_time")

	jumpService := services.NewJumpService()
	data, ok := jumpService.AddLink(url, end_time)
	fmt.Println(ok)
	fmt.Println(data)
}
