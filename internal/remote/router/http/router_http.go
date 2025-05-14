package http

import (
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
	"io"
	"net/http"
)

func (c *Client) Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error) {
	if path[0] == '/' {
		path = path[1:]
	}
	u := fmt.Sprintf("%s/api/v1/%s/%s", env.Env.ScrapelessApiHost, keyword, path)
	fmt.Println(u)
	request, err := http.NewRequest(method, u, body)
	if err != nil {
		log.Errorf("new request error :%v\n", err)
		return nil, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	request.Header.Set(env.Env.HTTPHeader, env.Env.Token)
	do, err := c.client.Do(request)
	if err != nil {
		log.Errorf("do request error :%v\n", err)
		return nil, err
	}
	b, err := io.ReadAll(do.Body)
	if err != nil {
		log.Errorf("read body error :%v\n", err)
		return nil, err
	}
	defer do.Body.Close()
	return b, nil
}
