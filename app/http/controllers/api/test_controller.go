package api

import (
	"github.com/goravel/framework/contracts/http"
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

	panic("test1111111")
	panic("test22222")
	//ctx.Response().Success().Json()
}
