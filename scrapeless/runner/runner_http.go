package runner

import (
	"context"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/env"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/code"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/runner/http"
	log "github.com/sirupsen/logrus"
)

type RunHttp struct {
}

func NewRunHttp() Runner {
	log.Info("runner http init", http.Default())
	if http.Default() == nil {
		http.Init()
	}
	return &RunHttp{}
}

func (rh *RunHttp) Abort(ctx context.Context) (bool, error) {
	ok, err := http.Default().Abort(ctx, env.Env.RunId, env.Env.ActorId)
	if err != nil {
		return false, code.Format(err)
	}
	return ok, nil
}

func (rh *RunHttp) Close() error {
	return nil
}
