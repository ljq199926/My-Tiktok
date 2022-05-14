package middleware

import (
	"apiGateway/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Jwtkey = []byte(utils.JwtKey)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func CheckToken(token string) (*MyClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	fmt.Println(tokenObj)
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		tokenStr := c.Query("token")
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//token格式错误
		//tokenSlice := strings.SplitN(tokenStr, " ", 2)
		//if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
		//	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token格式错误"})
		//	c.Abort() //阻止执行
		//	return
		//}
		//验证token
		tokenStruck, ok := CheckToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token不正确"})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token过期"})
			c.Abort() //阻止执行
			return
		}
		fmt.Println("jwt校验正确，允许通过----------")
		c.Next()
	}
}
