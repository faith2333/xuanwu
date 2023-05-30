package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	selfJwt "github.com/faith2333/xuanwu/pkg/middleware/jwt"
	"github.com/pkg/errors"
	"regexp"
)

const (
	// PasswordValidatePattern Regex pattern for password validation
	PasswordValidatePattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$"
)

type defaultUser struct {
	c     *Config
	uRepo IRepoUser
}

func NewDefaultUser(uRepo IRepoUser, c *Config) Interface {
	return &defaultUser{
		uRepo: uRepo,
		c:     c,
	}
}

func (d *defaultUser) SignUp(ctx context.Context, user *User) error {
	// salt the password
	user.Password = d.saltPassword(user.Password)
	return d.uRepo.Create(ctx, user)
}

func (d *defaultUser) Login(ctx context.Context, username, password string) (string, error) {
	dbUser, err := d.uRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	tokenString, err := selfJwt.CreateToken([]byte(d.c.JWTSecretKey), dbUser.Username)
	if err != nil {
		return "", errors.Wrap(err, "generate jwt token failed")
	}

	return tokenString, nil
}

func (d *defaultUser) checkPassword(user *User, reqPassword string) error {
	if user.Password == d.saltPassword(reqPassword) {
		return nil
	}
	return errors.New("validate password or username failed")
}

func (d *defaultUser) validatePassword(password string) error {
	match, err := regexp.MatchString(PasswordValidatePattern, password)
	if err != nil {
		return errors.Wrap(err, "regexp match failed")
	}

	if !match {
		return errors.New("password must contain at least one lowercase letter, one uppercase letter, one digit, one special character, and be at least 8 characters long")
	}

	return nil
}

func (d *defaultUser) saltPassword(password string) string {
	hash := md5.Sum([]byte(password + d.c.JWTSecretKey))
	return hex.EncodeToString(hash[:])
}
