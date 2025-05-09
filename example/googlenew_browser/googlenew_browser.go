package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	proxy2 "github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxy"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	actor := scrapeless.New(scrapeless.WithProxy(), scrapeless.WithBrowser(), scrapeless.WithStorage())
	defer actor.Close()

	proxy, err := actor.Proxy.Proxy(context.TODO(), proxy2.ProxyActor{
		Country:         "us",
		SessionDuration: 10,
		Gateway:         "gw-us.scrapeless.io:8789",
	})
	if err != nil {
		panic(err)
	}
	browserInfo, err := actor.Browser.Create(context.Background(), browser.Actor{
		Input:        browser.Input{SessionTtl: "180"},
		ProxyCountry: "US",
		ProxyUrl:     proxy,
	})
	if err != nil {
		panic(err)
	}
	data := chromedpScrape("https://news.google.com/search?q=nba&hl=en&gl=US", browserInfo.DevtoolsUrl)
	ok, err := actor.Storage.GetDataset().AddItems(context.Background(), []map[string]any{
		{
			"url":  "https://news.google.com/search?q=nba&hl=en&gl=US",
			"data": data,
		}})
	if err != nil {
		log.Errorf("add dataset err:%v", err)
		return
	}
	log.Info("ok:", ok)

}
func chromedpScrape(url string, devtoolsWsURL string) string {
	var htmlContent string
	allocatorCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), devtoolsWsURL, chromedp.NoModifyURL)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitReady("body"),
			chromedp.OuterHTML("html", &htmlContent),
		},
	)
	if err != nil {
		log.Errorf("chromedp err:%v", err)
	}
	return htmlContent
}
