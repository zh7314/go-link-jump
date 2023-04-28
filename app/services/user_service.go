package services

import (
	"errors"
	"fmt"
)

type UserService struct {
	//Dependent services
}

func NewUserService() *UserService {
	return &UserService{
		//Inject services
	}
}

func (r *UserService) Show(id uint) {

	if id == 0 {
		errors.New("sadasda")
	}
}

func (r *UserService) Zx(name string) {

	fmt.Println(name)
}

func Zxxx(name string) {

	fmt.Println(name)
}
