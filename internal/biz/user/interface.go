package user

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/faith2333/xuanwu/api/user/v1"
	"github.com/faith2333/xuanwu/internal/conf"
	selfJwt "github.com/faith2333/xuanwu/pkg/middleware/jwt"
)

type Interface interface {
	SignUp(ctx context.Context, user *User) error
	GetCurrentUser(ctx context.Context) (user selfJwt.CurrentUser, err error)
	Login(ctx context.Context, username, password string) (string, error)
	ChangePassword(ctx context.Context, req *ChangePasswordReq) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type IRepoUser interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int64, updateFields map[string]interface{}) error
	Delete(ctx context.Context, user *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
}

type User struct {
	ID          int64                  `json:"id"`
	Username    string                 `json:"username"`
	DisplayName string                 `json:"displayName"`
	Password    string                 `json:"password"`
	AvatarURL   string                 `json:"avatarURL"`
	Email       string                 `json:"email"`
	PhoneNumber string                 `json:"phoneNumber"`
	Enabled     string                 `json:"enabled"`
	ExtraInfo   map[string]interface{} `json:"extraInfo"`
}

type CurrentUser struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarURL"`
}

type Config struct {
	Type                 Type   `json:"type"`
	JWTSecretKey         string `json:"jwtSecretKey"`
	DefaultAdminPassword string `json:"defaultAdminPassword"`
}

type ChangePasswordReq pb.ChangePasswordRequest

func NewUserFactory(userRepo IRepoUser, c *conf.Server, modelConfig *conf.ModelConfig) (Interface, error) {
	config := &Config{
		Type:                 TypeDefault,
		JWTSecretKey:         c.Auth.JwtSecretKey,
		DefaultAdminPassword: modelConfig.User.DefaultAdminPassword,
	}

	switch config.Type {
	case TypeDefault:
		return NewDefaultUser(userRepo, config), nil
	}

	return nil, errors.New(fmt.Sprintf("%s user has not been supported", config.Type))
}
