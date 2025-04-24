package browser

import (
	"context"
	"fmt"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/config"
	remote_brwoser "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/browser/http"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/url"
	"strings"
)

type BHttp struct {
}

func NewBHttp() Browser {
	log.Info("browser http init", http.Default())
	if http.Default() == nil {
		http.Init()
	}
	return &BHttp{}
}
func (bh *BHttp) Create(ctx context.Context, req Actor) (*CreateResp, error) {
	create, err := http.Default().ScrapingBrowserCreate(ctx, &remote_brwoser.CreateBrowserRequest{
		ApiKey: env.Env.ApiKey,
		Input: map[string]string{
			"session_ttl": req.Input.SessionTtl,
		},
		Proxy: &remote_brwoser.ProxyParams{
			Url:             req.ProxyUrl,
			ChannelId:       req.ChannelId,
			Country:         strings.ToUpper(req.ProxyCountry),
			SessionDuration: req.SessionDuration,
			SessionId:       req.SessionId,
			Gateway:         req.Gateway,
		},
	})
	if err != nil {
		log.Errorf("scraping browser create err:%v\n", err)
		return nil, code.Format(err)
	}
	if create != nil {
		return &CreateResp{
			DevtoolsUrl: create.DevtoolsUrl,
			TaskId:      create.TaskId,
		}, nil
	}
	return nil, nil
}

func (bh *BHttp) CreateOnce(ctx context.Context, req ActorOnce) (*CreateResp, error) {
	u, err := url.Parse(config.ScrapingBrowserApiHost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "parse url error: %s", err.Error())
	}
	devtoolsUrl := fmt.Sprintf("wss://%s/browser", u.Host)
	value := &url.Values{}
	value.Set("token", env.Env.ApiKey)
	value.Set("session_ttl", req.Input.SessionTtl)
	value.Set("proxy_country", strings.ToUpper(req.ProxyCountry))
	return &CreateResp{
		DevtoolsUrl: devtoolsUrl + "?" + value.Encode(),
	}, nil
}
func (bh *BHttp) Close() error {
	return http.Default().Close()
}
