package appstore

import (
	"net/http"
)

type Resource[T any] struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Attr  T      `json:"attributes"`
	Links Links  `json:"links"`
}

type Links struct {
	Self  string `json:"self"`
	Next  string `json:"next"`
	First string `json:"first"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient  HTTPClient
	tokenSource TokenSource
}

func NewClient(httpClient HTTPClient, tokenSource TokenSource) *Client {
	return &Client{
		httpClient:  httpClient,
		tokenSource: tokenSource,
	}
}
