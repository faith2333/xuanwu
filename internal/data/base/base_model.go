package base

import (
	"encoding/json"
	"gorm.io/gorm"
)

type ModelBase struct {
	gorm.Model
	CreateUser string `json:"createUser" gorm:"type:varchar(16);comment:'create user'"`
	ModifyUser string `json:"modifyUser" gorm:"type:varchar(16);comment:'modify user'"`
}

type RepoBase struct{}

// Transform convert two object of struct with json tag
func (b RepoBase) Transform(source, target interface{}) error {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(sBytes, &target)
}
