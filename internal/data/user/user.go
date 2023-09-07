package user

import (
	"context"
	"fmt"
	biz "github.com/faith2333/xuanwu/internal/biz/user"
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/pkg/errors"
	"sync"
)

var userRunOnce = sync.Once{}

type User struct {
	ID          int64        `json:"id" gorm:"primaryKey"`
	Username    string       `json:"username" gorm:"type:varchar(64);uniqueIndex:CODE_DELETED"`
	Password    string       `json:"password" gorm:"type:varchar(128)"`
	Email       string       `json:"email" gorm:"type:varchar(128)"`
	PhoneNumber string       `json:"phoneNumber" gorm:"type:varchar(128)"`
	Enabled     bool         `json:"enabled"`
	ExtraInfo   base.TypeMap `json:"extraInfo" gorm:"type:json"`
	base.Model
}

func (User) TableName() string {
	return "user_users"
}

type RepoUser struct {
	base.Base
	data *base.Data
}

func NewRepoUser(data *base.Data) biz.IRepoUser {
	userRunOnce.Do(func() {
		err := data.DB(context.Background()).AutoMigrate(&User{})
		if err != nil {
			panic(err)
		}
	})

	return &RepoUser{
		data: data,
	}
}

func (repo *RepoUser) Create(ctx context.Context, bizUser *biz.User) error {
	dbUser := &User{}
	err := repo.TransformJson(bizUser, &dbUser)
	if err != nil {
		return err
	}

	return repo.data.DB(ctx).Model(&User{}).Create(&dbUser).Error
}

func (repo *RepoUser) Update(ctx context.Context, id int64, updateFields map[string]interface{}) error {
	query := repo.data.DB(ctx).Where("id = ? and deleted = 0", id).Updates(updateFields)
	return query.Error
}

func (repo *RepoUser) Delete(ctx context.Context, bizUser *biz.User) error {
	dbUser, err := repo.getUserByUsername(ctx, bizUser.Username)
	if err != nil {
		return err
	}

	return repo.data.DB(ctx).Where("username = ? and deleted = 0", bizUser.Username).Update("deleted = ?", dbUser.ID).Error
}

func (repo *RepoUser) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	dbUser, err := repo.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	bizUser := &biz.User{}
	err = repo.TransformJson(dbUser, &bizUser)
	if err != nil {
		return nil, err
	}

	return bizUser, nil
}

func (repo *RepoUser) GetByID(ctx context.Context, id int64) (*biz.User, error) {
	dbUser, err := repo.getUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bizUser := &biz.User{}
	err = repo.TransformJson(dbUser, &bizUser)
	if err != nil {
		return nil, err
	}

	return bizUser, nil
}

func (repo *RepoUser) getUserByUsername(ctx context.Context, username string) (*User, error) {
	users := make([]*User, 0)

	err := repo.data.DB(ctx).Where("username = ? and deleted = 0", username).Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("query by username %s not found", username))
	}

	return users[0], nil
}

func (repo *RepoUser) getUserByID(ctx context.Context, id int64) (*User, error) {
	users := make([]*User, 0)

	err := repo.data.DB(ctx).Where("id = ? and deleted = 0", id).Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("query by id %s not found", id))
	}

	return users[0], nil
}
