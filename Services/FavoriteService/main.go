package main

import (
	"FavoriteService/handler"
	"FavoriteService/model"
	FavoriteService "FavoriteService/proto/FavoriteService"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	model.InitDB()
	model.InitRedis()
	// New Service
	reg := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("go.micro.service.FavoriteService"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	FavoriteService.RegisterFavoriteServiceHandler(service.Server(), new(handler.FavoriteService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
