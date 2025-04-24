package scrapeless

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage/queue"
	"testing"
)

func Test_GetQueue(t *testing.T) {
	s := New(WithStorage())
	queue, err := s.Storage.GetQueue().Get(context.Background(), "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", queue)
}

func Test_GetQueues(t *testing.T) {
	s := New(WithStorage())
	queues, err := s.Storage.GetQueue().List(context.Background(), 1, 20, false)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", queues)
}

func Test_CreateQueue(t *testing.T) {
	s := New(WithStorage())
	queueId, err := s.Storage.GetQueue().Create(context.Background(), &queue.CreateQueueReq{
		Name:        "test_001",
		Description: "desc_001",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("queueId: %s", queueId)
}

func Test_QueueUpdate(t *testing.T) {
	s := New(WithStorage())
	err := s.Storage.GetQueue().Update(context.Background(), "test_001_update1", "desc_update")
	if err != nil {
		t.Error(err)
	}
	t.Log("QueueUpdate success")
}

func Test_QueueDelete(t *testing.T) {
	s := New(WithStorage())
	err := s.Storage.GetQueue().Delete(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("QueueDelete success")
}

func Test_QueueAck(t *testing.T) {
	s := New(WithStorage())
	err := s.Storage.GetQueue().Ack(context.Background(), "")
	if err != nil {
		t.Error(err)
	}
	t.Log("QueueAck success")
}
