package proxy

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	proxy2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/proxy"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/proxy/http"
	log "github.com/sirupsen/logrus"
)

type PHttp struct {
}

func NewPHttp() Proxy {
	log.Info("proxy http init")
	if http.Default() == nil {
		http.Init()
	}
	return &PHttp{}
}

// Proxy retrieves proxy information.
//
// Parameters:
//
//	ctx: context.Context - Context for the request.
//	proxy: ProxyActor - Struct containing proxy request parameters like country, session duration, etc.
func (ph *PHttp) Proxy(ctx context.Context, proxy ProxyActor) (string, error) {
	proxyUrl, err := http.Default().ProxyGetProxy(ctx, &proxy2.GetProxyRequest{
		ApiKey:          env.Env.ApiKey,
		Country:         proxy.Country,
		SessionDuration: proxy.SessionDuration,
		SessionId:       proxy.SessionId,
		Gateway:         proxy.Gateway,
		TaskId:          env.Env.RunId,
	})
	if err != nil {
		log.Errorf("get proxy err:%v\n", err)
		return "", code.Format(err)
	}
	return proxyUrl, nil
}

func (ph *PHttp) Close() error {
	return http.Default().Close()
}
