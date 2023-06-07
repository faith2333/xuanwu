package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/pkg/errors"
	"reflect"
	"time"

	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var (
	SigningMethod = jwtv4.SigningMethodHS256
)

type ctxUser struct{}

type CurrentUser struct {
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	PhoneNumber string                 `json:"phoneNumber"`
	ExtraInfo   map[string]interface{} `json:"extraInfo"`
}

type CustomClaims struct {
	User CurrentUser `json:"user"`
	jwtv4.RegisteredClaims
}

func NewJWTFunc(secretKey []byte) func(token *jwtv4.Token) (interface{}, error) {
	return func(token *jwtv4.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv4.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v1", token.Header["alg"])
		}

		return secretKey, nil
	}
}

func CreateToken(secretKey []byte, username, email, phoneNumber string, extraInfo map[string]interface{}) (string, error) {
	token := jwtv4.NewWithClaims(SigningMethod, CustomClaims{
		User: CurrentUser{
			Username:    username,
			Email:       email,
			PhoneNumber: phoneNumber,
			ExtraInfo:   extraInfo,
		},
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwtv4.NewNumericDate(time.Now()),
			NotBefore: jwtv4.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
	})

	return token.SignedString(secretKey)
}

// AuthPlugin get jwt claims from context and put CurrentUser into it.
func AuthPlugin(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		claims, ok := jwt.FromContext(ctx)
		if !ok {
			// token not exists, return directly.
			return handler(ctx, req)
		}

		mapClaims, ok := claims.(jwtv4.MapClaims)
		if !ok {
			return nil, errors.Errorf("jwt claims asset failed, expect jwtv4.MapCalaims but got %s", reflect.TypeOf(mapClaims))
		}

		uInterface, ok := mapClaims["user"]
		if !ok {
			return nil, errors.New("get user info from context failed")
		}

		uBytes, err := json.Marshal(uInterface)
		if err != nil {
			return nil, errors.Wrap(err, "marshal failed")
		}

		user := CurrentUser{}
		err = json.Unmarshal(uBytes, &user)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal failed")
		}

		ctx = context.WithValue(ctx, ctxUser{}, user)
		return handler(ctx, req)
	}
}

func FromContext(ctx context.Context) (CurrentUser, error) {
	curUser, ok := ctx.Value(ctxUser{}).(CurrentUser)
	if !ok {
		return CurrentUser{}, errors.New("user not found in ctx")
	}

	return curUser, nil
}
