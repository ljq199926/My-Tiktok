package main

import (
	"UserService/handler"
	"UserService/model"
	UserService "UserService/proto/UserService"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//初始化db
	model.InitDB()
	model.InitRedis()
	// New Service
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.UserService"),
		micro.Version("latest"),
		micro.Registry(reg),
	)
	// Initialise service
	service.Init()
	// Register Handler
	UserService.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
