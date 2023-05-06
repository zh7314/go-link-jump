package admin

import "github.com/goravel/framework/contracts/http"

type JumpController struct {
	//Dependent services
}

func NewJumpController() *JumpController {
	return &JumpController{
		//Inject services
	}
}

func (r *JumpController) GetData(ctx http.Context) {

}

func (r *JumpController) AddLink(ctx http.Context) {

}
