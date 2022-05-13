package router

import (
	"apiGateway/handler"
	"apiGateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.Cors())
	route.GET("/demo/:name", handler.Hello)

	userRouter := route.Group("/douyin")
	//这里使用中间件
	userRouter.POST("/user/register/", handler.Register)
	userRouter.POST("/user/login/", handler.Login)
	userRouter.GET("/user/", handler.Info)

	return route
}
