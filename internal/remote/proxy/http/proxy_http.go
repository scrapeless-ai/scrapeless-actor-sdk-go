package http

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/proxy"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) ProxyGetProxy(ctx context.Context, req *proxy.GetProxyRequest) (string, error) {
	if req.ApiKey == "" {
		return "", status.Errorf(codes.InvalidArgument, "api key is required")
	}
	if req.Country == "" {
		req.Country = env.ProxyCountry
	}
	if int64(req.SessionDuration) > env.ProxySessionDurationMax {
		req.SessionDuration = uint64(env.ProxySessionDurationMax)
	}
	if req.SessionId == "" {

		req.SessionId = funk.RandomString(10)
	}
	if req.Gateway == "" {
		req.Gateway = env.ProxyGatewayHost
	}

	proxyURL := fmt.Sprintf(
		"http://CHANNEL-proxy.residential-country_%s-r_%dm-s_%s--scrapelesstaskid_%s:%s@%s",
		req.Country,
		req.SessionDuration,
		req.SessionId,
		req.TaskId,
		req.ApiKey,
		req.Gateway,
	)
	return proxyURL, nil
}
