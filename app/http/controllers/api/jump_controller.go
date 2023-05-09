package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/services"
	"time"
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

	id := ctx.Request().Input("id")
	//id, err := strconv.Atoi(ctx.Request().Input("id"))
	//if err != nil {
	//	fmt.Println(id)
	//}
	key := services.NewJumpService().Key + id
	fmt.Println(key)

	err := facades.Cache.Put(key, "2222222222222222222", 500*time.Second)
	fmt.Println(err)

	isHas := facades.Cache.Has(key)
	fmt.Println(isHas)

	value := facades.Cache.Get(key)

	ctx.Response().Success().Json(value)
}

func (r *JumpController) GetData(ctx http.Context) {

	jumpService := services.NewJumpService()
	data := jumpService.GetData()

	ctx.Response().Success().Json(data)
}
