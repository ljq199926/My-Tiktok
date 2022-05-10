package handler

import (
	pb "apiGateway/proto/demoService"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func Hello(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	reg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(reg),
	)
	demoService := pb.NewDemoService("go.micro.service.demoService", microService.Client())
	resp, err := demoService.Call(context.Background(), &pb.Request{
		Name: name,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"msg": err})
		return
	}
	ctx.JSON(200, gin.H{"msg": resp.Msg})
}
