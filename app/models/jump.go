package models

import (
	"github.com/goravel/framework/database/orm"
	"time"
)

type Jump struct {
	orm.Model
	orm.SoftDeletes
	JumpUrl string    `gorm:"column:jump_url"`
	Md5Data string    `gorm:"column:md5_data"`
	EndTime time.Time `gorm:"column:end_time"`
}

func (r *Jump) TableName() string {
	return "jump_link"
}
