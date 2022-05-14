package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userId int64) (string, error) {
	expireTime := time.Now().Add(24 * 7 * time.Hour) //过期时间为7天
	nowTime := time.Now()                            //当前时间
	claims := MyClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "MyTikTok",        //颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString([]byte(JwtKey))
}
