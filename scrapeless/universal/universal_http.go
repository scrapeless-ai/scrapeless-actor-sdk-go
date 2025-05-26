package universal

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/universal"
	sh "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/universal/http"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

type UniversalHttp struct{}

func New() Universal {
	log.Info("Internal Router init")
	if sh.Default() == nil {
		sh.Init(env.Env.ScrapelessBaseApiUrl)
	}
	return UniversalHttp{}
}

func (us UniversalHttp) CreateTask(ctx context.Context, req ScrapingTaskRequest) ([]byte, error) {
	response, err := sh.Default().CreateTask(ctx, universal.UniversalTaskRequest{
		Actor: string(ScraperUniversal),
		Input: req.Input,
		Proxy: universal.TaskProxy{Country: req.ProxyCountry},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return nil, err
	}
	return response, nil
}

func (us UniversalHttp) Close() error {
	return sh.Default().Close()
}

func (us UniversalHttp) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := sh.Default().GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, err
	}
	return result, nil
}
