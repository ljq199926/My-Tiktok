package utils

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func InitService() micro.Service {
	reg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(reg),
	)
	return microService
}
