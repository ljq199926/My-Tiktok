package main

import (
	"VideoService/handler"
	"VideoService/model"
	VideoService "VideoService/proto/VideoService"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	model.InitDB()

	// New Service
	reg := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.service.VideoService"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Server().Init(grpc.MaxMsgSize(1024 * 1024 * 1024))

	// Register Handler
	VideoService.RegisterVideoServiceHandler(service.Server(), new(handler.VideoService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
