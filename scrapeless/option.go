package scrapeless

import (
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/httpserver"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxy"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/runner"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage"
)

const (
	typeHttp = "http"
	typeGrpc = "grpc"
)

type Option interface {
	Apply(*Actor)
}

type BrowserOption struct {
	tp string
}

func (o *BrowserOption) Apply(a *Actor) {
	if o.tp == typeGrpc {
		a.Browser = browser.NewBGrpc()
		a.CloseFun = append(a.CloseFun, a.Browser.Close)
	} else {
		a.Browser = browser.NewBHttp()
		a.CloseFun = append(a.CloseFun, a.Browser.Close)
	}
}

// WithBrowser choose browser type.
func WithBrowser(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &BrowserOption{
		tp: tp[0],
	}
}

type ProxyOption struct {
	tp string
}

func (o *ProxyOption) Apply(a *Actor) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Proxy.Close)
	} else {
		a.Proxy = proxy.NewPHttp()
		a.CloseFun = append(a.CloseFun, a.Proxy.Close)
	}
}

// WithProxy choose proxy type.
func WithProxy(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &ProxyOption{
		tp: tp[0],
	}
}

type CaptchaOption struct {
	tp string
}

func (o *CaptchaOption) Apply(a *Actor) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	} else {
		a.Captcha = captcha.NewCaHttp()
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	}
}

// WithCaptcha choose captcha type.
func WithCaptcha(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &CaptchaOption{
		tp: tp[0],
	}
}

type StorageOption struct {
	tp string
}

func (o *StorageOption) Apply(a *Actor) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	} else {
		a.Storage = storage.NewStorageHttp()
		a.CloseFun = append(a.CloseFun, a.Storage.Close)
	}
}

// WithStorage choose storage type.
func WithStorage(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &StorageOption{
		tp: tp[0],
	}
}

type RunnerOption struct {
	tp string
}

func (o *RunnerOption) Apply(a *Actor) {
	if o.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Captcha.Close)
	} else {
		a.runner = runner.NewRunHttp()
		a.CloseFun = append(a.CloseFun, a.runner.Close)
	}
}

// WithRunner choose runner type.
func withRunner(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &RunnerOption{
		tp: tp[0],
	}
}

type ServerOption struct {
	mode httpserver.ServerMode
}

func (s *ServerOption) Apply(a *Actor) {
	a.Server = httpserver.New(s.mode)
}

// WithServer choose server mode.
func WithServer(mode ...httpserver.ServerMode) Option {
	if len(mode) == 0 {
		mode = append(mode, httpserver.ReleaseMode)
	}
	return &ServerOption{mode: mode[0]}
}
