package jwt

import (
	"fmt"
	"time"

	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var (
	SigningMethod = jwtv4.SigningMethodHS256
)

type CustomClaims struct {
	Username string `json:"username"`
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

func CreateToken(secretKey []byte, username string) (string, error) {
	token := jwtv4.NewWithClaims(SigningMethod, CustomClaims{
		Username: username,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwtv4.NewNumericDate(time.Now()),
			NotBefore: jwtv4.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
	})

	return token.SignedString(secretKey)
}
