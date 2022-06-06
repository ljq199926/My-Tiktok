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
 @Date               :2022/5/15 18:48
 @Description        :
 @Function List      :
 @History            :
*/

func AddToken(token string, user *User) error {
	data := map[string]interface{}{
		"UserId":        user.UserId,
		"Username":      user.Username,
		"Password":      user.Password,
		"FollowCount":   user.FollowCount,
		"FollowerCount": user.FollowerCount,
		"CreateDate":    user.CreateDate,
	}
	log.Info("AddToken: ", redisDb, token, user, data)
	err := redisDb.HMSet(context.Background(), token, data).Err()
	log.Info(err)
	return err
}
func QueryUserIdByToken(c context.Context, token string) (int64, error) {
	log.Info("QueryUserIdByToken", redisDb, token)
	tokens := redisDb.HMGet(c, token, "UserId").Val()
	if len(tokens) == 0 || tokens == nil {
		return -1, errors.New("query redis failed")
	}
	parseInt, err := strconv.ParseInt(tokens[0].(string), 10, 64)
	if err != nil {
		log.Info(err)
		return -1, err
	}
	log.Info("QueryUserIdByToken-END: ", parseInt)
	return parseInt, nil
}
