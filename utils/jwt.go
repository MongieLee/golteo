package utils

import (
	"errors"
	"ginl/entitys/persistent"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var privateSigningKey = []byte("mongielee")

type CustomClaims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

// 生成JWT
func GenRegisteredClaims(user *persistent.User) (string, error) {
	claims := &CustomClaims{
		Username: user.UserName,
		Id:       user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-pure",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateSigningKey)
}

// 校验JWT有效性，不解析
func ValidRegisteredClaims(tokenStr string) bool {
	parse, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return privateSigningKey, nil
	})
	if err != nil {
		return false
	}
	return parse.Valid
}

// 解析JWT
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
