package main

import (
	"demoService/handler"
	demoService "demoService/proto/demoService"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	reg := consul.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Address("127.0.0.1:12341"),
		micro.Name("go.micro.service.demoService"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	demoService.RegisterDemoServiceHandler(service.Server(), new(handler.DemoService))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
