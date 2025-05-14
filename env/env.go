package env

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

var (
	Env    = &actorEnv{}
	LogEnv = &logEnv{}
	once   = sync.Once{}
)

const (
	HTTPHeader = "x-api-token"

	envActorId       = "SCRAPELESS_ACTOR_ID"
	envRunId         = "SCRAPELESS_RUN_ID"
	envTeamId        = "SCRAPELESS_TEAM_ID"
	envApiKey        = "SCRAPELESS_API_KEY"
	envToken         = "SCRAPELESS_TOKEN"
	envKvNamespaceId = "SCRAPELESS_KV_NAMESPACE_ID"
	envDatasetId     = "SCRAPELESS_DATASET_ID"
	envBucketId      = "SCRAPELESS_BUCKET_ID"
	envQueueId       = "SCRAPELESS_QUEUE_ID"
	envHttpPort      = "SCRAPELESS_HTTP_PORT"
	envInput         = "SCRAPELESS_INPUT"
	envXApiKey       = "SCRAPELESS_X_API_KEY"

	envBrowserApiHost          = "SCRAPELESS_BROWSER_API_HOST"
	envProxyCountry            = "SCRAPELESS_PROXY_COUNTRY"
	envProxySessionDurationMax = "SCRAPELESS_PROXY_SESSION_DURATION_MAX"
	envProxyGatewayHost        = "SCRAPELESS_PROXY_GATEWAY_HOST"
	envApiHost                 = "SCRAPELESS_API_HOST"
	envCaptchaHost             = "SCRAPELESS_CAPTCHA_HOST"

	envLogMaxSize    = "SCRAPELESS_LOG_MAX_SIZE"
	envLogMaxBackups = "SCRAPELESS_LOG_MAX_BACKUPS"
	envLogMaxAge     = "SCRAPELESS_LOG_MAX_AGE"
	envLogRootDir    = "SCRAPELESS_LOG_ROOT_DIR"
)

func LoadEnv() {
	once.Do(func() {
		viper.SetConfigFile(`.env`)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Error("Error reading config file: %v", err)
		}
		Env = &actorEnv{
			TeamId:  viper.GetString(envTeamId),
			ActorId: viper.GetString(envActorId),
			RunId:   viper.GetString(envRunId),
			ApiKey:  viper.GetString(envApiKey),

			KvNamespaceId: viper.GetString(envKvNamespaceId),
			DatasetId:     viper.GetString(envDatasetId),
			BucketId:      viper.GetString(envBucketId),
			QueueId:       viper.GetString(envQueueId),
			HttpPort:      viper.GetString(envHttpPort),
		}
		ScrapingBrowserApiHost = viper.GetString(envBrowserApiHost)
		Token = viper.GetString(envToken)
		ProxyCountry = viper.GetString(envProxyCountry)
		ProxySessionDurationMax = viper.GetInt64(envProxySessionDurationMax)
		ProxyGatewayHost = viper.GetString(envProxyGatewayHost)
		ScrapelessApiHost = viper.GetString(envApiHost)
		ScrapelessCaptchaHost = viper.GetString(envCaptchaHost)
		if ScrapelessCaptchaHost == "" {
			ScrapelessCaptchaHost = ScrapelessApiHost
		}
		LogEnv = &logEnv{
			MaxSize:    viper.GetInt(envLogMaxSize),
			MaxBackups: viper.GetInt(envLogMaxBackups),
			MaxAge:     viper.GetInt(envLogMaxAge),
			LogRootDir: viper.GetString(envLogRootDir),
		}
		if err := Env.param(); err != nil {
			log.Errorf("LoadEnv param err: %v", err)
			return
		}
		log.Infof("actor env init %+v", Env)
	})
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

	HttpPort string
}

type logEnv struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogRootDir string
}
