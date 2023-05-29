package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/faith2333/xuanwu/internal/conf"
)

type Interface interface {
	SignUp(ctx context.Context, user *User) error
	Login(ctx context.Context, username, password string) (string, error)
}

type IRepoUser interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, user *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
}

type User struct {
	ID          int64                  `json:"id"`
	Username    string                 `json:"username"`
	Password    string                 `json:"password"`
	Email       string                 `json:"email"`
	PhoneNumber string                 `json:"phoneNumber"`
	ExtraInfo   map[string]interface{} `json:"extraInfo"`
}

type Config struct {
	Type         Type   `json:"type"`
	JWTSecretKey string `json:"jwtSecretKey"`
}

func NewUserFactory(userRepo IRepoUser, c *conf.Server) (Interface, error) {
	config := &Config{
		Type:         TypeDefault,
		JWTSecretKey: c.Auth.JwtSecretKey,
	}

	switch config.Type {
	case TypeDefault:
		return NewDefaultUser(userRepo, config), nil
	}

	return nil, errors.New(fmt.Sprintf("%s user has not been supported", config.Type))
}
