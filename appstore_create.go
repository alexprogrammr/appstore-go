package appstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type createRequest[T any] struct {
	Data struct {
		Type string `json:"type"`
		Attr T      `json:"attributes"`
		Rel  any    `json:"relationships"`
	} `json:"data"`
}

type relation struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newCreateRequest[T any](attr T, typ string, rel relation) createRequest[T] {
	r := createRequest[T]{}
	r.Data.Type = typ
	r.Data.Attr = attr
	r.Data.Rel = map[string]any{
		rel.Type[:len(rel.Type)-1]: map[string]any{
			"data": rel,
		},
	}

	return r
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

	rp := new(response[Resource[T]])
	if err := json.NewDecoder(resp.Body).Decode(rp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &rp.Data, nil
}
