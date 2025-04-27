package scrapeless

import (
	"encoding/json"
	"errors"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/helper"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxy"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/runner"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage"
	"os"
	"reflect"
)

type Actor struct {
	Browser  browser.Browser
	Proxy    proxy.Proxy
	Captcha  captcha.Captcha
	Storage  storage.Storage
	runner   runner.Runner
	CloseFun []func() error
}

// New creates a new Actor with some options.
func New(opts ...Option) *Actor {
	var actor = new(Actor)
	for _, opt := range opts {
		opt.Apply(actor)
	}
	return actor
}

// Close closes the actor.
func (a *Actor) Close() {
	for _, f := range a.CloseFun {
		f()
	}
	os.Exit(0)
}

// Input get input data from env.
func (a *Actor) Input(data any) error {
	input := helper.GetString("SCRAPELESS_INPUT", ``)
	tf := reflect.TypeOf(data)
	if tf.Kind() != reflect.Ptr {
		return errors.New("data must be ptr")
	}
	return json.Unmarshal([]byte(input), data)
}
