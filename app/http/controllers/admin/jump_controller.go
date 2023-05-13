package admin

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
	"goravel/app/utils"
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

	utils.Success(ctx, data, "获取成功")
}

func (r *JumpController) AddLink(ctx http.Context) {

	url := ctx.Request().Input("url")
	end_time := ctx.Request().Input("end_time")

	//jumpService := services.NewJumpService()
	data, ok := services.NewJumpService().AddLink(url, end_time)
	if ok != nil {
		utils.Fail(ctx, "", ok.Error())
	} else {
		utils.Success(ctx, data, "增加成功")
	}
}
