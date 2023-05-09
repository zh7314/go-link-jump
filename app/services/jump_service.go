package services

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

const jumpKey = "JumpLink:"

type JumpService struct {
	//Dependent services
	Key string
}

func NewJumpService() *JumpService {
	return &JumpService{
		//Inject model
		Key: jumpKey,
	}
}

func (r *JumpService) GetData() []models.Jump {
	var jump []models.Jump
	facades.Orm.Query().Get(&jump)

	return jump
}

func (r *JumpService) AddLink() {

}
