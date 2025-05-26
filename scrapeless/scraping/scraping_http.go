package scraping

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/scraping"
	sh "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/scraping/http"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

type ScrapingHttp struct{}

func New() Scraping {
	log.Info("Internal Router init")
	if sh.Default() == nil {
		sh.Init(env.Env.ScrapelessBaseApiUrl)
	}
	return ScrapingHttp{}
}

func (s ScrapingHttp) CreateTask(ctx context.Context, req ScrapingTaskRequest) ([]byte, error) {
	response, err := sh.Default().CreateTask(ctx, scraping.ScrapingTaskRequest{
		Actor: string(req.Actor),
		Input: req.Input,
		Proxy: scraping.TaskProxy{Country: req.ProxyCountry},
	})
	if err != nil {
		log.Errorf("scraping create err:%v", err)
		return nil, code.Format(err)
	}
	return response, nil
}

func (s ScrapingHttp) Close() error {
	return sh.Default().Close()
}

func (s ScrapingHttp) GetTaskResult(ctx context.Context, taskId string) ([]byte, error) {
	result, err := sh.Default().GetTaskResult(ctx, taskId)
	if err != nil {
		log.Errorf("get task result err:%v", err)
		return nil, code.Format(err)
	}
	return result, nil
}
