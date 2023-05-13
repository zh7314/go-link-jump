package api

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
)

type TestController struct {
	//Dependent services
}

func NewTestController() *TestController {
	return &TestController{
		//Inject services
	}
}

func (r *TestController) Test(ctx http.Context) {

	//panic("test1111111")
	//panic("test22222")
	//ctx.Response().Success().Json()
	utils.Success(ctx, 111, "成功")
}
