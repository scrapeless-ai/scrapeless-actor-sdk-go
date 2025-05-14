package env

import (
	"errors"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	HTTPHeader              string `mapstructure:"SCRAPELESS_HTTP_HEADER"`
	ScrapingBrowserApiHost  string `mapstructure:"SCRAPELESS_BROWSER_API_HOST"`
	Token                   string `mapstructure:"SCRAPELESS_TOKEN"`
	ProxyCountry            string `mapstructure:"SCRAPELESS_PROXY_COUNTRY"`
	ProxySessionDurationMax int64  `mapstructure:"SCRAPELESS_PROXY_SESSION_DURATION_MAX"`
	ProxyGatewayHost        string `mapstructure:"SCRAPELESS_PROXY_GATEWAY_HOST"`
	ScrapelessApiHost       string `mapstructure:"SCRAPELESS_API_HOST"`
	ScrapelessCaptchaHost   string `mapstructure:"SCRAPELESS_CAPTCHA_HOST"`

	Actor actorEnv `mapstructure:",squash"`
	Log   logEnv   `mapstructure:",squash"`
}

type actorEnv struct {
	TeamId  string `mapstructure:"SCRAPELESS_TEAM_ID"`
	ActorId string `mapstructure:"SCRAPELESS_ACTOR_ID"`
	RunId   string `mapstructure:"SCRAPELESS_RUN_ID"`
	ApiKey  string `mapstructure:"SCRAPELESS_API_KEY"`

	KvNamespaceId string `mapstructure:"SCRAPELESS_KV_NAMESPACE_ID"`
	DatasetId     string `mapstructure:"SCRAPELESS_DATASET_ID"`
	BucketId      string `mapstructure:"SCRAPELESS_BUCKET_ID"`
	QueueId       string `mapstructure:"SCRAPELESS_QUEUE_ID"`

	HttpPort string `mapstructure:"SCRAPELESS_HTTP_PORT"`
}

type logEnv struct {
	MaxSize    int    `mapstructure:"SCRAPELESS_LOG_MAX_SIZE"`
	MaxBackups int    `mapstructure:"SCRAPELESS_LOG_MAX_BACKUPS"`
	MaxAge     int    `mapstructure:"SCRAPELESS_LOG_MAX_AGE"`
	LogRootDir string `mapstructure:"SCRAPELESS_LOG_ROOT_DIR"`
}

var Env config

func (c *config) Validate() error {
	if c.Actor.TeamId == "" {
		return errors.New("invalid env param team_Id")
	}
	if c.Actor.ActorId == "" {
		return errors.New("invalid env param actor_Id")
	}
	if c.Actor.RunId == "" {
		return errors.New("invalid env param run_Id")
	}
	if c.Actor.ApiKey == "" {
		return errors.New("invalid env param apikey")
	}

	return nil
}

func init() {
	// default value
	viper.SetDefault("SCRAPELESS_PROXY_COUNTRY", "ANY")
	viper.SetDefault("SCRAPELESS_BROWSER_API_HOST", "https://api.scrapeless.com")
	viper.SetDefault("SCRAPELESS_PROXY_SESSION_DURATION_MAX", 120)
	viper.SetDefault("SCRAPELESS_PROXY_GATEWAY_HOST", "gw-us.scrapeless.io:8789")
	viper.SetDefault("SCRAPELESS_HTTP_HEADER", "x-api-token")

	// Retrieve the directory where the current file is located (env directory)
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	viper.SetConfigFile(filepath.Join(dir, ".env"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("scrapeless: warn reading config file: %v", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		panic(err)
	}

	if Env.ScrapelessCaptchaHost == "" {
		Env.ScrapelessCaptchaHost = Env.ScrapelessApiHost
	}

	err = Env.Validate()
	if err != nil {
		log.Errorf("scrapeless: validate config err: %v", err)
	}
	log.Infof("scrapeless: conf: %+v", Env)
}

func GetLogEnv() *logEnv {
	return &Env.Log
}

func GetActorEnv() *actorEnv {
	return &Env.Actor
}
