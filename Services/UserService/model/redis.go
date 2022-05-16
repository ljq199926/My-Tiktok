package model

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
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
