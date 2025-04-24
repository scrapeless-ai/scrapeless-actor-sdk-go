package http

import (
	"fmt"
	"net/http"
)

var (
	defaultGatewayClient *Client
)

func Init() {
	fmt.Println("browser init")
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
