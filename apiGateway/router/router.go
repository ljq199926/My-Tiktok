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

	authRouter := route.Group("/douyin/")

	authRouter.Use(middleware.JwtMiddleware())
	{
		authRouter.GET("user/", handler.Info)
		authRouter.POST("/publish/action/", handler.UploadVideo)
		authRouter.POST("/favorite/action/", handler.FavorAction)
		authRouter.GET("/favorite/list/", handler.FavorList)
		authRouter.POST("/comment/action/", handler.CommentAction)
		authRouter.GET("/comment/list/", handler.CommentList)
		authRouter.GET("/publish/list/", handler.PublishList)
	}
	//userRouter.POST("/publish/action/", handler.UploadVideo)
	userRouter.GET("/feed/", handler.Feed)
	//userRouter.GET("/publish/list/", handler.PublishList)
	//userRouter.POST("/favorite/action/", handler.FavorAction)
	//userRouter.GET("/favorite/list/", handler.FavorList)

	return route
}
