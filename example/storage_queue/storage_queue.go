package main

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage/queue"
	log "github.com/sirupsen/logrus"
)

func main() {
	sl := scrapeless.New(scrapeless.WithStorage())

	// push a message to queue
	msgId, err := sl.Storage.GetQueue().Push(context.Background(), queue.PushQueue{
		Name:     "test-cy",
		Payload:  []byte("aaaa"),
		Retry:    0,
		Timeout:  0,
		Deadline: 0,
	})
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(msgId)

	// pull a message from queue
	pullResp, err := sl.Storage.GetQueue().Pull(context.Background(), 100)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(pullResp)
	for _, v := range pullResp {
		// ack message
		err = sl.Storage.GetQueue().Ack(context.Background(), v.QueueID)
		if err != nil {
			log.Error(err)
		}
	}

}
