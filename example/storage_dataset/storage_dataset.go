package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

func main() {
	sl := actor.New(actor.WithStorage())
	defer sl.Close()

	success, err := sl.Storage.GetDataset().AddItems(context.Background(), []map[string]any{
		{
			"name": "John",
			"age":  20,
		},
		{
			"name": "lucy",
			"age":  19,
		},
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	if success {
		items, err := sl.Storage.GetDataset().GetItems(context.Background(), 1, 10, false)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Infof("%v", items)
	}

}
