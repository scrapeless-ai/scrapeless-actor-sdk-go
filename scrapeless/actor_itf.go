package scrapeless

import (
	"encoding/json"
	"errors"
	_ "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/log"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/httpserver"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxy"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/runner"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage"
	"github.com/spf13/viper"
	"reflect"
)

type Actor struct {
	Browser  browser.Browser
	Proxy    proxy.Proxy
	Captcha  captcha.Captcha
	Storage  storage.Storage
	runner   runner.Runner
	Server   httpserver.Server
	CloseFun []func() error
}

// New creates a new Actor with some options.
func New(opts ...Option) *Actor {
	//_ = env.LoadEnv()

	var actor = new(Actor)
	for _, opt := range opts {
		opt.Apply(actor)
	}
	return actor
}

// Close closes the actor.
func (a *Actor) Close() {
	for _, f := range a.CloseFun {
		_ = f()
	}
}

// Input get input data from env.
func (a *Actor) Input(data any) error {
	input := viper.GetStringMapString("SCRAPELESS_INPUT")
	inputData, _ := json.Marshal(input)
	tf := reflect.TypeOf(data)
	if tf.Kind() != reflect.Ptr {
		return errors.New("data must be ptr")
	}
	return json.Unmarshal(inputData, data)
}

func (a *Actor) Start() error {
	return a.Server.Start(":8848")
}
