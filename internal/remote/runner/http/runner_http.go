package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/config"
	request2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/request"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c *Client) Abort(ctx context.Context, runId, actorId string) (bool, error) {
	body, err := request2.Request(ctx, request2.ReqInfo{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("%s/api/v1/actors/%s/runs/%s", config.ScrapelessApiHost, actorId, runId),
		Headers: map[string]string{
			env.HTTPHeader: config.Token,
		},
	})
	if err != nil {
		return false, err
	}
	var resp request2.RespInfo
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}
	if resp.Err {
		log.Error(body)
		return false, fmt.Errorf("acto Abort err:%s", resp.Msg)
	}
	return true, nil
}
