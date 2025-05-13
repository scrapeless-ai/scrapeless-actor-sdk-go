package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

func main() {
	// new scrapeless with storage.
	scrapeless := scrapeless.New(scrapeless.WithStorage())
	defer scrapeless.Close()

	// Set value use default namespace
	ok, err := scrapeless.Storage.GetKv().SetValue(context.Background(), "key", "nice boy", 20)
	if err != nil {
		log.GetLogger().Error().Msg(err.Error())
	}
	log.GetLogger().Info().Msgf("ok:%v", ok)

	// Get value use default namespace
	value, err := scrapeless.Storage.GetKv().GetValue(context.Background(), "key")
	if err != nil {
		log.GetLogger().Error().Msg(err.Error())
	}
	log.GetLogger().Info().Msgf("value:%v", value)
}
