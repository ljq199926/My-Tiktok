package handler

import (
	pb "apiGateway/proto/demoService"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	client "github.com/micro/go-micro/v2/client/grpc"
)

func Hello(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	demoService := pb.NewDemoService("go.micro.service.demoService", client.NewClient())
	resp, err := demoService.Call(context.Background(), &pb.Request{
		Name: name,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"msg": err})
		return
	}
	ctx.JSON(200, gin.H{"msg": resp.Msg})
}
