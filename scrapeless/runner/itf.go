package runner

import "context"

type Runner interface {
	Abort(ctx context.Context) (bool, error)
	Close() error
}
