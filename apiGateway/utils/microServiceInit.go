package utils

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func InitService() micro.Service {
	reg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(reg),
		//micro.Client(grpc.NewClient(grpc.MaxSendMsgSize(1024*1024*1024))),
		//micro.Client(grpc.NewClient(grpc.MaxRecvMsgSize(1024*1024*1024))),
	)
	return microService
}
