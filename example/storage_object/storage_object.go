package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

func main() {
	sl := scrapeless.New(scrapeless.WithStorage())
	defer sl.Close()

	// Put object The supported types include JSON、html、png
	objectId, err := sl.Storage.GetObject().Put(context.Background(), "object.json", []byte("data"))
	if err != nil {
		log.GetLogger().Error().Msg(err.Error())
		return
	}
	if objectId != "" {
		// Get object
		resp, err := sl.Storage.GetObject().Get(context.Background(), objectId)
		if err != nil {
			panic(err)
		}
		log.GetLogger().Info().Msg(string(resp))
	}
}
