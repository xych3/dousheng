package util

import (
	"gitee.com/derrickball/douyin/pkg/constants"
	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(token *jwt.Token) int64 {
	claims := token.Claims.(jwt.MapClaims)
	return int64(claims[constants.IdentityKey].(float64))
}
