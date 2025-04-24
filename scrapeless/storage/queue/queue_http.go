package queue

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/storage/storage_http"
	log "github.com/sirupsen/logrus"
	"time"
)

type QueueHttp struct {
	queueId string
}

// NewQueueHttp new queue http
func NewQueueHttp(queueId ...string) Queue {
	log.Info("queue http init", storage_http.Default())
	if storage_http.Default() == nil {
		storage_http.Init()
	}
	q := &QueueHttp{queueId: env.Env.QueueId}
	if len(queueId) > 0 {
		q.queueId = queueId[0]
	}
	return q
}

// List retrieves a list of queues with pagination and sorting options.
// Parameters:
//
//	ctx: The context for the request.
//	page: int64 - The page number (minimum 1, defaults to 1 if invalid).
//	pageSize: int64 - Number of items per page (minimum 10, defaults to 10 if invalid).
//	desc: bool - Whether to sort results in descending order.
func (q *QueueHttp) List(ctx context.Context, page int64, pageSize int64, desc bool) (*ListQueuesResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	queues, err := storage_http.Default().GetQueues(ctx, &storage_http.GetQueuesRequest{
		Page:     page,
		PageSize: pageSize,
		Desc:     desc,
	})
	if err != nil {
		log.Errorf("failed to list queues: %v\n", code.Format(err))
		return nil, code.Format(err)
	}
	var items []Item
	for _, item := range queues.Items {
		items = append(items, Item{
			Id:          item.Id,
			Name:        item.Name,
			UserId:      item.UserId,
			TeamId:      item.TeamId,
			ActorId:     item.ActorId,
			RunId:       item.RunId,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return &ListQueuesResponse{
		Items:  items,
		Total:  queues.Total,
		Limit:  queues.Limit,
		Offset: queues.Offset,
	}, nil
}

// Create creates a new HTTP queue with the provided request parameters.
// Parameters:
//
//	ctx: The context for the request.
//	req: The request object containing queue configuration details.
func (q *QueueHttp) Create(ctx context.Context, req *CreateQueueReq) (string, error) {
	queue, err := storage_http.Default().CreateQueue(ctx, &storage_http.CreateQueueRequest{
		ActorId:     env.Env.ActorId,
		RunId:       env.Env.RunId,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		log.Errorf("failed to create queue: %v\n", code.Format(err))
		return "", code.Format(err)
	}

	return queue.Id, nil
}

// Get retrieves a queue item by name.
// Parameters:
//
//	ctx: The context for the request.
//	name: The name of the queue to retrieve.
func (q *QueueHttp) Get(ctx context.Context, name string) (*Item, error) {
	queue, err := storage_http.Default().GetQueue(ctx, &storage_http.GetQueueRequest{
		Id:   q.queueId,
		Name: name,
	})
	if err != nil {
		log.Errorf("failed to get queue: %v\n", code.Format(err))
		return nil, code.Format(err)
	}
	return &Item{
		Id:          queue.Id,
		Name:        queue.Name,
		TeamId:      queue.TeamId,
		ActorId:     queue.ActorId,
		RunId:       queue.RunId,
		Description: queue.Description,
		CreatedAt:   queue.CreatedAt,
		UpdatedAt:   queue.UpdatedAt,
	}, nil
}

// Update updates the queue information with the provided name and description.
// Parameters:
//
//	ctx: The context for the request.
//	name: The new name of the queue.
//	description: The new description of the queue.
func (q *QueueHttp) Update(ctx context.Context, name string, description string) error {
	err := storage_http.Default().UpdateQueue(ctx, &storage_http.UpdateQueueRequest{
		QueueId:     q.queueId,
		Name:        name,
		Description: description,
	})
	return err
}

// Delete deletes the queue using the storage HTTP service.
// Parameters:
//
//	ctx: The context for the request.
func (q *QueueHttp) Delete(ctx context.Context) error {
	err := storage_http.Default().DelQueue(ctx, &storage_http.DelQueueRequest{QueueId: q.queueId})
	if err != nil {
		log.Errorf("failed to delete queue: %v\n", code.Format(err))
		return code.Format(err)
	}
	return nil
}

// QueuePush  timeout-->[60,300]   deadline--> [300,86400]
func (q *QueueHttp) pushWithId(ctx context.Context, req PushQueue) (string, error) {
	// [60,300]
	if req.Timeout < 60 {
		req.Timeout = 60
	}
	if req.Timeout > 300 {
		req.Timeout = 300
	}

	// [300,86400]
	if req.Deadline < 300 {
		req.Deadline = 400
	}
	if req.Deadline > 86400 {
		req.Deadline = 86400
	}

	unix := time.Now().UTC().Add(time.Duration(req.Deadline) * time.Second).Unix()
	queue, err := storage_http.Default().CreateMsg(ctx, &storage_http.CreateMsgRequest{
		QueueId:  q.queueId,
		Name:     req.Name,
		PayLoad:  string(req.Payload),
		Retry:    req.Retry,
		Timeout:  req.Timeout,
		Deadline: unix,
	})
	if err != nil {
		log.Errorf("failed to push to queue: %v\n", code.Format(err))
		return "", code.Format(err)
	}
	return queue.MsgId, nil
}

// Push adds a request to the HTTP queue and returns the task ID.
//
// Parameters:
//
//	ctx context.Context: The context for the request, used for cancellation and timeouts.
//	req PushQueue: The request to be pushed into the queue.
func (q *QueueHttp) Push(ctx context.Context, req PushQueue) (string, error) {
	return q.pushWithId(ctx, req)
}
func (q *QueueHttp) pullWithId(ctx context.Context, size int32) (GetMsgResponse, error) {
	if size < 1 {
		size = 1
	}
	if size > 100 {
		size = 100
	}
	msgs, err := storage_http.Default().GetMsg(ctx, &storage_http.GetMsgRequest{
		QueueId: q.queueId,
		Limit:   size,
	})
	if err != nil {
		log.Errorf("failed to pull from queue: %v\n", code.Format(err))
		return nil, code.Format(err)
	}
	if msgs == nil {
		return nil, nil
	}
	var items []*Msg
	for _, msg := range *msgs {
		items = append(items, &Msg{
			ID:        msg.ID,
			QueueID:   msg.QueueID,
			Name:      msg.Name,
			Payload:   msg.Payload,
			Timeout:   msg.Timeout,
			Deadline:  msg.Deadline,
			Retry:     msg.Retry,
			Retried:   msg.Retried,
			SuccessAt: msg.SuccessAt,
			FailedAt:  msg.FailedAt,
			Desc:      msg.Desc,
		})
	}
	return items, nil
}

// Pull retrieves messages from the HTTP queue.
// Parameters:
//
//	ctx: The context used to control the request lifecycle (e.g., cancellation, deadlines).
//	size: The maximum number of messages to retrieve in this operation.
func (q *QueueHttp) Pull(ctx context.Context, size int32) (GetMsgResponse, error) {
	return q.pullWithId(ctx, size)
}

func (q *QueueHttp) ackWithId(ctx context.Context, msgId string) error {
	err := storage_http.Default().AckMsg(ctx, &storage_http.AckMsgRequest{
		QueueId: q.queueId,
		MsgId:   msgId,
	})
	if err != nil {
		log.Errorf("failed to ack msg: %v\n", code.Format(err))
		return code.Format(err)
	}
	return nil
}

// Ack confirms that a message has been processed successfully.
//
// Parameters:
//
//	ctx: The context used for request cancellation or timeout.
//	msgId: The unique identifier of the message to acknowledge.
func (q *QueueHttp) Ack(ctx context.Context, msgId string) error {
	return q.ackWithId(ctx, msgId)
}

func (k *QueueHttp) Close() error {
	return nil
}
