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
		fmt.Println(str_id)
		fmt.Println("查询条件不能为空")
	}
	key := services.NewJumpService().Key + str_id
	fmt.Println(key)

	isHas := facades.Cache.Has(key)
	if isHas {
		fmt.Println(key)
		fmt.Println("key存在")
	}

	//err := facades.Cache.Put(key, "2222222222222222222", -1*time.Second)
	//if err != nil {
	//	fmt.Println("保存失败")
	//}

	res := facades.Cache.GetString(key)

	//res := facades.Cache.Get(key).(string)
	if len(res) == 0 {
		fmt.Println("值为空")
		fmt.Println(res)
	}

	fmt.Println(res)

	//zx := reflect.TypeOf(res)
	//fmt.Println("类型：")
	//fmt.Println(zx.Name(), zx.Kind())

	//if len(strings.TrimSpace(res)) == 0 {
	//	fmt.Println("url为空")
	//}
	query, err := url.QueryUnescape(res)
	if err != nil {
		fmt.Println("解析失败")
	}
	//ctx.Response().Success().Json(query)

	ctx.Response().Redirect(http.StatusMovedPermanently, query)
}

func (r *JumpController) GetData(ctx http.Context) {

	jumpService := services.NewJumpService()
	data := jumpService.GetData()

	ctx.Response().Success().Json(data)
}
