package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	selfJwt "github.com/faith2333/xuanwu/pkg/middleware/jwt"
	"github.com/pkg/errors"
	"regexp"
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

func (d *defaultUser) GetCurrentUser(ctx context.Context) (selfJwt.CurrentUser, error) {
	return selfJwt.FromContext(ctx)
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

	tokenString, err := selfJwt.CreateToken([]byte(d.c.JWTSecretKey), dbUser.Username, dbUser.Email, dbUser.PhoneNumber, dbUser.ExtraInfo)
	if err != nil {
		return "", errors.Wrap(err, "generate jwt token failed")
	}

	return tokenString, nil
}

func (d *defaultUser) ChangePassword(ctx context.Context, req *ChangePasswordReq) error {
	user, err := d.uRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return err
	}

	if err = d.checkPassword(user, req.CurrentPassword); err != nil {
		return err
	}

	if err = d.validatePassword(req.NewPassword); err != nil {
		return err
	}

	return d.uRepo.Update(ctx, user.ID, map[string]interface{}{"password": req.NewPassword})
}

func (d *defaultUser) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return d.uRepo.GetByUsername(ctx, username)
}

func (d *defaultUser) checkPassword(user *User, reqPassword string) error {
	if user.Password == d.saltPassword(reqPassword) {
		return nil
	}
	return errors.New("validate password or username failed")
}

func (d *defaultUser) validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("the length of password must be greater than 8")
	}

	numsPart := `[0-9]{1}`
	matched, err := regexp.Match(numsPart, []byte(password))
	if !matched || err != nil {
		return errors.Errorf("password must contains number")
	}

	lowerPart := `[a-z]{1}`
	matched, err = regexp.Match(lowerPart, []byte(password))
	if !matched || err != nil {
		return errors.Errorf("password must contains lower character")
	}

	upperPart := `[A-Z]{1}`
	matched, err = regexp.Match(upperPart, []byte(password))
	if !matched || err != nil {
		return errors.Errorf("password must contains upper character")
	}

	symbolPart := `[!@#$%^&*()]{1}`
	matched, err = regexp.Match(symbolPart, []byte(password))
	if !matched || err != nil {
		return errors.Errorf("password must contains symbol")
	}

	return nil
}

func (d *defaultUser) saltPassword(password string) string {
	hash := md5.Sum([]byte(password + d.c.JWTSecretKey))
	return hex.EncodeToString(hash[:])
}
