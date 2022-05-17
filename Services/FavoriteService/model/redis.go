package model

import (
	"context"
	"errors"
	log "github.com/micro/go-micro/v2/logger"
	"strconv"
)

func QueryUserIdByToken(c context.Context, token string) (int64, error) {
	log.Info("QueryUserIdByToken", redisDB, token)
	tokens := redisDB.HMGet(c, token, "UserId").Val()
	if len(tokens) == 0 {
		return -1, errors.New("query redis failed")
	}
	parseInt, err := strconv.ParseInt(tokens[0].(string), 10, 64)
	if err != nil {
		return -1, err
	}
	log.Info("QueryUserIdByToken-END: ", parseInt)
	return parseInt, nil
}
