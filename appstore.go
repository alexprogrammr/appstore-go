package appstore

import "net/http"

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
