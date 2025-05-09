package storage_http

import (
	"google.golang.org/grpc"
	"net/http"
)

var (
	defaultStorageClient *Client
)

type Client struct {
	client *http.Client
}

func Init() {
	var (
		err error
	)
	defaultStorageClient, err = New()
	regisHttpHandleFunc()
	if err != nil {
		panic(err)
	}
}

func Default() *Client {
	return defaultStorageClient
}

func New(opts ...grpc.DialOption) (*Client, error) {
	return &Client{
		client: &http.Client{},
	}, nil
}

func (c *Client) Close() error {
	return nil
}
