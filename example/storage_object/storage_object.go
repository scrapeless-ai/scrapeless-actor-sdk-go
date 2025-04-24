package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	log "github.com/sirupsen/logrus"
)

func main() {
	sl := scrapeless.New(scrapeless.WithStorage())
	defer sl.Close()

	// Put object The supported types include JSON、html、png
	objectId, err := sl.Storage.GetObject().Put(context.Background(), "object.json", []byte("data"))
	if err != nil {
		log.Error(err)
		return
	}
	if objectId != "" {
		// Get object
		resp, err := sl.Storage.GetObject().Get(context.Background(), objectId)
		if err != nil {
			panic(err)
		}
		log.Info(string(resp))
	}
}
