package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
	proxy2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxy"
	"io"
	"net/http"
	"net/url"
)

func main() {
	sl := scrapeless.New(scrapeless.WithProxy())
	defer sl.Close()
	proxy, err := sl.Proxy.Proxy(context.TODO(), proxy2.ProxyActor{
		Country:         "US",
		SessionDuration: 180,
		SessionId:       "YOUR SESSION ID",
		Gateway:         "YOU GATEWAY",
	})
	if err != nil {
		panic(err)
	}
	log.GetLogger().Info().Msg(proxy)
	parse, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("HEADER KEY", "HEADER VALUE")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(parse)}}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.GetLogger().Info().Msg(string(body))
}
