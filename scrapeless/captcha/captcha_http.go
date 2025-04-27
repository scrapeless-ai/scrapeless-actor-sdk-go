package captcha

import (
	"context"
	"encoding/json"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	gateway_captcha "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/captcha/http"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type CaHttp struct {
}

func NewCaHttp() Captcha {
	log.Info("captcha http init")
	if http.Default() == nil {
		http.Init()
	}
	return &CaHttp{}
}

func (c *CaHttp) Solver(ctx context.Context, req *CaptchaSolverReq) (*CaptchaSolverResp, error) {
	var (
		inputMap map[string]any
	)

	input, err := json.Marshal(req.Input)
	_ = json.Unmarshal(input, &inputMap)
	response, err := http.Default().CaptchaSolverSolverTask(ctx, &gateway_captcha.CreateTaskRequest{
		ApiKey: env.Env.ApiKey,
		Actor:  req.Actor,
		Input:  inputMap,
		Proxy: &gateway_captcha.ProxyParams{
			Url:             req.Proxy.Url,
			ChannelId:       req.Proxy.ChannelId,
			Country:         req.Proxy.Country,
			SessionDuration: req.Proxy.SessionDuration,
			SessionId:       req.Proxy.SessionId,
			Gateway:         req.Proxy.Gateway,
		},
		Timeout: req.TimeOut,
	})
	if err != nil {
		log.Errorf("captcha solver err:%v\n", err)
		return nil, code.Format(err)
	}
	marshal, _ := json.Marshal(response)

	token := gjson.Parse(string(marshal)).Get("token").String()
	return &CaptchaSolverResp{Token: token}, nil
}

func (c *CaHttp) Create(ctx context.Context, req *CaptchaSolverReq) (string, error) {
	var (
		inputMap map[string]any
	)

	input, err := json.Marshal(req.Input)
	_ = json.Unmarshal(input, &inputMap)
	taskId, err := http.Default().CaptchaSolverCreateTask(ctx, &gateway_captcha.CreateTaskRequest{
		ApiKey: env.Env.ApiKey,
		Actor:  req.Actor,
		Input:  inputMap,
		Proxy: &gateway_captcha.ProxyParams{
			Url:             req.Proxy.Url,
			ChannelId:       req.Proxy.ChannelId,
			Country:         req.Proxy.Country,
			SessionDuration: req.Proxy.SessionDuration,
			SessionId:       req.Proxy.SessionId,
			Gateway:         req.Proxy.Gateway,
		},
		Timeout: req.TimeOut,
	})
	if err != nil {
		log.Errorf("captcha creat err:%v\n", err)
		return "", code.Format(err)
	}
	return taskId, nil
}

func (c *CaHttp) ResultGet(ctx context.Context, req *CaptchaSolverReq) (*CaptchaSolverResp, error) {
	response, err := http.Default().CaptchaSolverGetTaskResult(ctx, &gateway_captcha.GetTaskResultRequest{
		ApiKey: env.Env.ApiKey,
		TaskId: req.TaskId,
	})
	if err != nil {
		log.Errorf("captcha result get err:%v\n", err)
		return nil, err
	}
	marshal, _ := json.Marshal(response)

	token := gjson.Parse(string(marshal)).Get("token").String()

	return &CaptchaSolverResp{Token: token}, nil
}

func (c *CaHttp) Close() error {
	return nil
}
