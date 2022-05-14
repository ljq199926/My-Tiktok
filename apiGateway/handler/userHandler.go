package handler

import (
	pb "apiGateway/proto/UserService"
	"apiGateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	microService := utils.InitService()
	userService := pb.NewUserService("go.micro.service.UserService", microService.Client())
	resp, err := userService.Login(context.Background(), &pb.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})

}
func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	microService := utils.InitService()
	userService := pb.NewUserService("go.micro.service.UserService", microService.Client())
	resp, err := userService.Register(context.Background(), &pb.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, gin.H{
		"status_code": resp.StatusCode,
		"status_msg":  resp.StatusMsg,
		"user_id":     resp.UserId,
		"token":       resp.Token,
	})
}

func Info(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	token := ctx.Query("token")
	_userid, _ := strconv.Atoi(user_id)
	userid := int64(_userid)
	microService := utils.InitService()
	userService := pb.NewUserService("go.micro.service.UserService", microService.Client())
	resp, err := userService.Info(context.Background(), &pb.DouyinUserRequest{
		UserId: userid,
		Token:  token,
	})
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status_code": resp.StatusCode,
		"user": gin.H{
			"id":             resp.User.Id,
			"name":           resp.User.Name,
			"follow_count":   resp.User.FollowCount,
			"follower_count": resp.User.FollowerCount,
			"is_follow":      resp.User.IsFollow,
		},
	})
}
