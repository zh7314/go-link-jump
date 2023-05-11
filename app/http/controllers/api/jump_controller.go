package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/services"
	"net/url"
	"strings"
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

	str_id := ctx.Request().Input("id")
	//id, err := strconv.Atoi(ctx.Request().Input("id"))
	if len(strings.TrimSpace(str_id)) == 0 {
		fmt.Println("查询条件不能为空")
	}
	key := services.NewJumpService().Key + str_id

	isHas := facades.Cache.Has(key)
	if isHas {
		fmt.Println("key存在")
	}
	res := facades.Cache.GetString(key)

	if len(res) == 0 {
		fmt.Println("值为空")
	}
	jump, err := url.QueryUnescape(res)
	if err != nil {
		fmt.Println("解析失败")
	}
	//ctx.Response().Success().Json(query)
	ctx.Response().Redirect(http.StatusMovedPermanently, jump)
}
