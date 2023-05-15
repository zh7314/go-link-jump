package models

import (
	"github.com/goravel/framework/database/orm"
)

type ApiLog struct {
	orm.Model
	orm.SoftDeletes
	Name   string
	Avatar string
}

func (r *ApiLog) TableName() string {
	return "api_log"
}
