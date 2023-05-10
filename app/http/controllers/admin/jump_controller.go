package admin

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/models"
	"goravel/app/services"
	"net/url"
	"strings"
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

func (r *JumpController) GetData(ctx http.Context) {

	jumpService := services.NewJumpService()
	data := jumpService.GetData()

	ctx.Response().Success().Json(data)
}

func (r *JumpController) AddLink(ctx http.Context) {

	//err := facades.Cache.Put(key, "2222222222222222222", -1*time.Second)
	//if err != nil {
	//	fmt.Println("保存失败")
	//}

	str_url := ctx.Request().Input("url")
	str_end_time := ctx.Request().Input("end_time")

	if len(strings.TrimSpace(str_url)) == 0 {
		fmt.Println(str_url)
		fmt.Println("查询条件不能为空")
	}
	if len(strings.TrimSpace(str_end_time)) == 0 {
		fmt.Println(str_end_time)
		fmt.Println("查询条件不能为空")
	}

	var jump models.Jump
	jump.JumpUrl = url.QueryEscape(str_url)

	now := time.Now()
	start := now.Format("2006-01-02 15:04:05")
	fmt.Println(start)
}
