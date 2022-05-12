package handler

import (
	pb "apiGateway/proto/UserService"
	"apiGateway/utils"
	"context"
	"github.com/gin-gonic/gin"
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
	ctx.JSON(200, resp)

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
	ctx.JSON(200, resp)
}
func Info(ctx *gin.Context) {

}
