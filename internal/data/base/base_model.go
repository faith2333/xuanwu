package base

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	CreateUser string `json:"createUser" gorm:"type:varchar(16);comment:'create user'"`
	ModifyUser string `json:"modifyUser" gorm:"type:varchar(16);comment:'modify user'"`
}
