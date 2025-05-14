package request

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
	"io"
	"net/http"
	"strings"
)

var (
	c *http.Client
)

func init() {
	c = &http.Client{}
}

type ReqInfo struct {
	Method  string `json:"method"`
	Url     string `json:"url"`
	Body    string `json:"body"`
	Headers map[string]string
}

type RespInfo struct {
	Data any    `json:"data"`
	Err  bool   `json:"err"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (resp RespInfo) IsErr() bool {
	return resp.Err
}

func (resp RespInfo) Error() string {
	return resp.Msg
}

func (resp RespInfo) GetData() any {
	return resp.Data
}

func Request(ctx context.Context, reqInfo ReqInfo) (string, error) {
	request, err := http.NewRequestWithContext(ctx, reqInfo.Method, reqInfo.Url, strings.NewReader(reqInfo.Body))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	for k, v := range reqInfo.Headers {
		request.Header.Set(k, v)
	}
	if reqInfo.Body != "" {
		if reqInfo.Body[0] == '[' || reqInfo.Body[0] == '{' {
			request.Header.Set("Content-Type", "application/json")
		} else {
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	do, err := c.Do(request)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	all, err := io.ReadAll(do.Body)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	defer do.Body.Close()
	return string(all), nil
}
