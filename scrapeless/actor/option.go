package actor

import (
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/browser"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/captcha"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/deepserp"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/httpserver"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/proxies"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/scraping"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/storage"
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/universal"
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
		a.Proxy = proxies.NewPHttp()
		a.CloseFun = append(a.CloseFun, a.Proxy.Close)
	}
}

// WithProxy choose proxies type.
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
		a.CloseFun = append(a.CloseFun, a.Storage.Close)
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

type DeepSerpOption struct {
	tp string
}

func (d *DeepSerpOption) Apply(a *Actor) {
	if d.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.DeepSerp.Close)
	} else {
		a.DeepSerp = deepserp.New()
		a.CloseFun = append(a.CloseFun, a.DeepSerp.Close)
	}
}

// WithDeepSerp choose DeepSerp type.
func WithDeepSerp(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &DeepSerpOption{
		tp: tp[0],
	}
}

type ScrapingOption struct {
	tp string
}

func (s *ScrapingOption) Apply(a *Actor) {
	if s.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Scraping.Close)
	} else {
		a.Scraping = scraping.New()
		a.CloseFun = append(a.CloseFun, a.Scraping.Close)
	}
}

// WithScraping choose scraping type.
func WithScraping(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &ScrapingOption{
		tp: tp[0],
	}
}

type UniversalOption struct {
	tp string
}

func (s *UniversalOption) Apply(a *Actor) {
	if s.tp == typeGrpc {
		a.CloseFun = append(a.CloseFun, a.Universal.Close)
	} else {
		a.Universal = universal.New()
		a.CloseFun = append(a.CloseFun, a.Universal.Close)
	}
}

// WithUniversal choose universal type.
func WithUniversal(tp ...string) Option {
	if len(tp) == 0 {
		tp = append(tp, typeHttp)
	}
	return &UniversalOption{
		tp: tp[0],
	}
}
