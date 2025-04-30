package env

import (
	"errors"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/helper"
	log "github.com/sirupsen/logrus"
)

var Env = &actorEnv{}

const (
	EnvActorId       = "SCRAPELESS_ACTOR_ID"
	EnvRunId         = "SCRAPELESS_RUN_ID"
	EnvTeamId        = "SCRAPELESS_TEAM_ID"
	EnvInput         = "SCRAPELESS_INPUT"
	EnvApiKey        = "SCRAPELESS_API_KEY"
	EnvToken         = "SCRAPELESS_TOKEN"
	EnvXApiKey       = "SCRAPELESS_X_API_KEY"
	EnvKvNamespaceId = "SCRAPELESS_KV_NAMESPACE_ID"
	EnvDatasetId     = "SCRAPELESS_DATASET_ID"
	EnvBucketId      = "SCRAPELESS_BUCKET_ID"
	EnvQueueId       = "SCRAPELESS_QUEUE_ID"
	HTTPHeader       = "x-api-token"
)

func LoadEnv() error {
	Env = &actorEnv{
		TeamId:  helper.GetString(EnvTeamId, ""),
		ActorId: helper.GetString(EnvActorId, ""),
		RunId:   helper.GetString(EnvRunId, ""),
		ApiKey:  helper.GetString(EnvApiKey, ""),

		KvNamespaceId: helper.GetString(EnvKvNamespaceId, ""),
		DatasetId:     helper.GetString(EnvDatasetId, ""),
		BucketId:      helper.GetString(EnvBucketId, ""),
		QueueId:       helper.GetString(EnvQueueId, ""),
	}
	if err := Env.param(); err != nil {
		log.Errorf("LoadEnv param err: %v", err)
		return err
	}
	log.Info("actor env init", Env)
	return nil
}

func (ae *actorEnv) param() error {
	if ae.TeamId == "" {
		return errors.New("invalid env param team_Id")
	}
	if ae.ActorId == "" {
		return errors.New("invalid env param actor_Id")
	}
	if ae.RunId == "" {
		return errors.New("invalid env param run_Id")
	}
	if ae.ApiKey == "" {
		return errors.New("invalid env param apikey")
	}
	return nil
}

type actorEnv struct {
	TeamId  string
	ActorId string
	RunId   string
	ApiKey  string

	KvNamespaceId string
	DatasetId     string
	BucketId      string
	QueueId       string
}
