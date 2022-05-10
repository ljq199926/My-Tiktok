package main

import (
	"demoService/handler"
	demoService "demoService/proto/demoService"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.demoService"),
		micro.Version("latest"),
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
