package services

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type UserService struct {
	//Dependent services
}

func NewUserService() *UserService {
	return &UserService{
		//Inject model
	}
}

func (r *UserService) Show(id uint) {

	if id == 0 {
		errors.New("sadasda")
	}
}

func (r *UserService) Zx(name string) []models.User {

	fmt.Println(name)

	var user []models.User
	facades.Orm.Query().Get(&user)

	fmt.Println(user)

	return user
}

func Zxxx(name string) {

	fmt.Println(name)
}
