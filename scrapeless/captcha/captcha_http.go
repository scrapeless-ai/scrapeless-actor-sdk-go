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

// Solver solves the captcha task by submitting it to the captcha solving service
//
// Parameters:
//
//	ctx: Context for controlling the request lifecycle and deadlines
//	req: Captcha solving request parameters object containing input data and configuration
func (c *CaHttp) Solver(ctx context.Context, req *CaptchaSolverReq) (*CaptchaSolverResp, error) {
	var (
		inputMap map[string]any
	)
	// Convert the input object to JSON format and unmarshal into a generic map to meet API requirements
	input, err := json.Marshal(req.Input)
	_ = json.Unmarshal(input, &inputMap)

	// Submit the captcha solving task to the remote service with provided parameters
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
	// Marshal the API response into JSON format and extract the 'token' field from the result
	marshal, _ := json.Marshal(response)
	token := gjson.Parse(string(marshal)).Get("token").String()
	return &CaptchaSolverResp{Token: token}, nil
}

// Create creates and submits a captcha solving task
//
// Parameters:
//
//	ctx: Context for controlling the request lifecycle and deadlines
//	req: Captcha solving task request parameters
func (c *CaHttp) Create(ctx context.Context, req *CaptchaSolverReq) (string, error) {
	var (
		inputMap map[string]any
	)
	// Convert input object to JSON and unmarshal into generic map (required by API)
	input, err := json.Marshal(req.Input)
	_ = json.Unmarshal(input, &inputMap)

	// Submit captcha solving task to remote service with provided configuration
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
		// Log error and return formatted error response
		log.Errorf("captcha creat err:%v\n", err)
		return "", code.Format(err)
	}
	return taskId, nil
}

// ResultGet retrieves the captcha solving result using the task ID
//
// Parameters:
//
//	ctx: context object for controlling the request lifecycle and timeouts
//	req: captcha solving request parameters containing the task ID
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
