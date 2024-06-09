package appstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func doGet[T any](c *Client, ctx context.Context, url string) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rp := new(T)
	if err := json.NewDecoder(resp.Body).Decode(rp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return rp, nil
}

func doCreate[T any](c *Client, ctx context.Context, url string, data any) (*Resource[T], error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rp := new(createResponse[Resource[T]])
	if err := json.NewDecoder(resp.Body).Decode(rp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &rp.Data, nil
}

func doDelete(c *Client, ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
