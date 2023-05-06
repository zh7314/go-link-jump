package services

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type JumpService struct {
	//Dependent services
}

func NewJumpService() *JumpService {
	return &JumpService{
		//Inject model
	}
}

func (r *JumpService) GetData() []models.Jump {
	var jump []models.Jump
	facades.Orm.Query().Get(&jump)

	return jump
}

func (r *JumpService) AddLink() {

}
