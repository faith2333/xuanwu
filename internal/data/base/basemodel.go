package base

import (
	"context"
	"encoding/json"
	"github.com/faith2333/xuanwu/pkg/xtime"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreateUser string `json:"createUser" gorm:"type:varchar(16)"`
	ModifyUser string `json:"modifyUser" gorm:"type:varchar(16)"`
	GmtCreate  string `json:"gmtCreate" gorm:"type:varchar(64)"`
	GmtModify  string `json:"gmtModify" gorm:"type:varchar(64)"`
	Deleted    int64  `json:"deleted" gorm:"uniqueIndex:CODE_DELETED"`
}

type RepoBase struct{}

// Transform convert two object of struct with json tag
func (b RepoBase) Transform(source, target interface{}) error {
	sBytes, err := json.Marshal(source)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}
	return json.Unmarshal(sBytes, &target)
}

func (b RepoBase) StringToPBStruct(source string) (*structpb.Struct, error) {
	vMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(source), &vMap)
	if err != nil {
		return nil, err
	}

	return structpb.NewStruct(vMap)
}

func (m *Model) getUserFromCtx(ctx context.Context) string {
	// todo get user form context, which set up in middleware
	return "sys"
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	now := xtime.StringTime(time.Now())
	m.GmtCreate = now
	m.GmtModify = now
	user := m.getUserFromCtx(tx.Statement.Context)
	m.CreateUser = user
	m.ModifyUser = user
	m.Deleted = 0

	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	m.GmtModify = xtime.StringTime(time.Now())
	m.ModifyUser = m.getUserFromCtx(tx.Statement.Context)
	return
}
