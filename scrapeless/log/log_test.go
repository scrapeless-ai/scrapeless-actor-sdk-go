package log

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	// export SCRAPELESS_RUN_ID=21da69d8-e75a-4360-afe8-84aaecd68ea7
	GetLogger().Trace().Msg("hello world")
	GetLogger().Debug().Msg("hello world")
	GetLogger().Info().Msg("hello world")
	GetLogger().Warn().Msg("hello world")
	GetLogger().Error().Msg("hello world")
	// GetLogger().Fatal().Msg("hello world")
	// GetLogger().Panic().Msg("hello world")

	GetLogger().Trace().Msgf("say: %s", "hello world")
	GetLogger().Debug().Msgf("say: %s", "hello world")
	GetLogger().Info().Msgf("say: %s", "hello world")
	GetLogger().Warn().Msgf("say: %s", "hello world")
	GetLogger().Error().Msgf("say: %s", "hello world")
	// GetLogger().Fatal().Msgf("say: %s", "hello world")
	// GetLogger().Panic().Msgf("say: %s", "hello world")

	for i := 0; i < 10000; i++ {
		GetLogger().Trace().Msgf("hello world: %d", i)
		GetLogger().Debug().Msgf("hello world: %d", i)
		GetLogger().Info().Msgf("hello world: %d", i)
		GetLogger().Warn().Msgf("hello world: %d", i)
		GetLogger().Error().Msgf("hello world: %d", i)
	}

	// _ = archiveCurrentLog()
}

func TestLogWithTraceID(t *testing.T) {
	// export SCRAPELESS_RUN_ID=21da69d8-e75a-4360-afe8-84aaecd68ea7
	ctx := context.WithValue(context.Background(), traceKey, "trace-val")
	GetLogger().Info().Ctx(ctx).Msg("log with trace")
}
