package router

import (
	rh "github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/remote/router/http"
	"io"

	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
)

type Internal struct{}

func New(baseUrl string) Router {
	log.Info("Internal Router init")
	if rh.Default() == nil {
		rh.Init(baseUrl)
	}
	return Internal{}
}

func (r Internal) Request(keyword string, method string, path string, body io.Reader, headers map[string]string) (data []byte, err error) {
	return rh.Default().Request(keyword, method, path, body, headers)
}

func (r Internal) Close() error {
	return nil
}
