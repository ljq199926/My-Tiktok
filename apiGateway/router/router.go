package router

import (
	"apiGateway/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()
	route.GET("/demo/:name", handler.Hello)
	return route
}
