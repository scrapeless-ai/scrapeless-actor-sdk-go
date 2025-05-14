package http

import (
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/log"
	"net/http"
)

var (
	defaultGatewayClient *Client
)

func Init() {
	log.Info("captcha init")
	var err error
	defaultGatewayClient, err = New()
	if err != nil {
		panic(err)
	}
}

type Client struct {
	client *http.Client
}

func Default() *Client {
	return defaultGatewayClient
}

func New() (*Client, error) {
	return &Client{
		client: &http.Client{},
	}, nil
}

func (c *Client) Close() error {
	c.client.CloseIdleConnections()
	return nil
}
