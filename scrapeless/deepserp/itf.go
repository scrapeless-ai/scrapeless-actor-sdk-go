package deepserp

import "context"

type Deepserp interface {
	CreateTask(ctx context.Context, req DeepserpTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskId string) ([]byte, error)

	Close() error
}
