package deepserp

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/deepserp"
	dh "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/deepserp/http"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

type DeepserpHttp struct{}

func New() Deepserp {
	log.Info("Internal Router init")
	if dh.Default() == nil {
		dh.Init(env.Env.ScrapelessBaseApiUrl)
	}
	return DeepserpHttp{}
}

func (s DeepserpHttp) CreateTask(ctx context.Context, req DeepserpTaskRequest) ([]byte, error) {
	response, err := dh.Default().CreateTask(ctx, deepserp.DeepserpTaskRequest{
		Actor: string(req.Actor),
		Input: req.Input,
		Proxy: deepserp.TaskProxy{Country: req.ProxyCountry},
	})
	if err != nil {
		log.Errorf("deepserp create err:%v", err)
		return nil, code.Format(err)
	}
	return response, nil
}

func (s DeepserpHttp) Close() error {
	return dh.Default().Close()
}

func (s DeepserpHttp) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := dh.Default().GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, code.Format(err)
	}
	return result, nil
}
