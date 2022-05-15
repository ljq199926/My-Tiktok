package model

import (
	"context"
	"errors"
	log "github.com/micro/go-micro/v2/logger"
	"strconv"
)

/*
 @File Name          :redis.go
 @Author             :cc
 @Version            :1.0.0
 @Date               :2022/5/15 17:03
 @Description        :
 @Function List      :
 @History            :
*/

func QueryUserIdByToken(c context.Context, token string) (int64, error) {
	log.Info("QueryUserIdByToken", redisDB, token)
	tokens := redisDB.HMGet(c, token, "UserId").Val()
	if len(tokens) == 0 {
		return -1, errors.New("query redis failed")
	}
	return strconv.ParseInt(tokens[0].(string), 10, 64)
}
