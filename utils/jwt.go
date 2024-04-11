package utils

import (
	"errors"
	"ginl/app/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var privateSigningKey = []byte("mongielee")

type CustomClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	jwt.RegisteredClaims
}

// GenRegisteredClaims 生成JWT
func GenRegisteredClaims(user *model.User, expired time.Duration) (string, error) {
	claims := &CustomClaims{
		Username: user.UserName,
		UserId:   user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-pure",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateSigningKey)
}

func GenerateAccessToken(user *model.User) (string, error) {
	return GenRegisteredClaims(user, time.Hour*12)
}

func GenerateRefreshToken(user *model.User) (string, error) {
	return GenRegisteredClaims(user, time.Hour*24)
}

// ValidRegisteredClaims 校验JWT有效性，不解析
func ValidRegisteredClaims(tokenStr string) bool {
	parse, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return privateSigningKey, nil
	})
	if err != nil {
		return false
	}
	return parse.Valid
}

// ParseJWTToken 解析JWT
func ParseJWTToken(tokenStr string) (*CustomClaims, error) {
	parse, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateSigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(*CustomClaims); ok && parse.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid parse")
}
