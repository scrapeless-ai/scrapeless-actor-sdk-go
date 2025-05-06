package env

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	viper.SetConfigFile(`C:\Users\Administrator\Desktop\dc\ai\scrapeless-actor-sdk-go\.env`)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file: %v", err)
		return err
	}
	Env = &actorEnv{
		TeamId:  viper.GetString(EnvTeamId),
		ActorId: viper.GetString(EnvActorId),
		RunId:   viper.GetString(EnvRunId),
		ApiKey:  viper.GetString(EnvApiKey),

		KvNamespaceId: viper.GetString(EnvKvNamespaceId),
		DatasetId:     viper.GetString(EnvDatasetId),
		BucketId:      viper.GetString(EnvBucketId),
		QueueId:       viper.GetString(EnvQueueId),
	}
	LogLevel = viper.GetString("SCRAPELESS_LOG_LEVEL")
	ScrapingBrowserApiHost = viper.GetString("SCRAPELESS_BROWSER_API_HOST")
	Token = viper.GetString("SCRAPELESS_TOKEN")
	ProxyCountry = viper.GetString("SCRAPELESS_PROXY_COUNTRY")
	ProxySessionDurationMax = viper.GetInt64("SCRAPELESS_PROXY_SESSION_DURATION_MAX")
	ProxyGatewayHost = viper.GetString("SCRAPELESS_PROXY_GATEWAY_HOST")
	ScrapelessApiHost = viper.GetString("SCRAPELESS_API_HOST")
	ScrapelessCaptchaHost = viper.GetString("SCRAPELESS_CAPTCHA_HOST")
	if ScrapelessCaptchaHost == "" {
		ScrapelessCaptchaHost = ScrapelessApiHost
	}
	if err := Env.param(); err != nil {
		log.Errorf("LoadEnv param err: %v", err)
		return err
	}
	log.Infof("actor env init %+v", Env)
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
