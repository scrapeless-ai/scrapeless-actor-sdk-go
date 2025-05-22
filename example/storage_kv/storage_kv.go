package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/actor"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

func main() {
	// new scrapeless with storage.
	scrapeless := actor.New(actor.WithStorage())
	defer scrapeless.Close()

	// Set value use default namespace
	ok, err := scrapeless.Storage.GetKv().SetValue(context.Background(), "key", "nice boy", 20)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("ok:%v", ok)

	// Get value use default namespace
	value, err := scrapeless.Storage.GetKv().GetValue(context.Background(), "key")
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("value:%v", value)
}
