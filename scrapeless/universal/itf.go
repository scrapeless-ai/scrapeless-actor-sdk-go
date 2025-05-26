package universal

import "context"

type Universal interface {
	CreateTask(ctx context.Context, req ScrapingTaskRequest) ([]byte, error)
	GetTaskResult(ctx context.Context, taskId string) ([]byte, error)

	Close() error
}
