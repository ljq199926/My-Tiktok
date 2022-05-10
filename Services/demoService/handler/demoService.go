package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	demoService "demoService/proto/demoService"
)

type DemoService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *DemoService) Call(ctx context.Context, req *demoService.Request, rsp *demoService.Response) error {
	log.Info("Received DemoService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
