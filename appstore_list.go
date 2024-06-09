package appstore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type listResponse[T any] struct {
	Data  []Resource[T] `json:"data"`
	Links Links         `json:"links"`
	Meta  meta          `json:"meta"`
}

type meta struct {
	Paging paging `json:"paging"`
}

type paging struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

func doList[T any](c *Client, ctx context.Context, url string) ([]Resource[T], error) {
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
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	rp := new(listResponse[T])
	if err := json.NewDecoder(resp.Body).Decode(rp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return rp.Data, nil
}
