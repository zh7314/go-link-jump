package models

import (
	"github.com/goravel/framework/database/orm"
)

type ApiUser struct {
	orm.Model
	orm.SoftDeletes
	Name   string
	Avatar string
}

func (r *ApiUser) TableName() string {
	return "api_user"
}
