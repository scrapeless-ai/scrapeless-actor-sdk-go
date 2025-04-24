package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage/queue"
)

func main() {
	sl := scrapeless.New(scrapeless.WithStorage())
	defer sl.Close()

	sl.Storage.GetQueue().Push(context.TODO(), queue.PushQueue{
		Name:     "task",
		Payload:  []byte("https://www.scrapeless.com/sitemap.xml"),
		Retry:    0,
		Timeout:  0,
		Deadline: 0,
	})
}
