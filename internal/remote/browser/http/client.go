package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	defaultGatewayClient *Client
)

func Init() {
	log.Info("browser init")
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
