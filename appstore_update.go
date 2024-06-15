package appstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type updateResource struct {
	ID   string       `json:"id"`
	Type resourceType `json:"type"`
	Attr any          `json:"attributes"`
}

func doUpdate[T any](c *Client, ctx context.Context, url string, resource updateResource) (*Resource[T], error) {
	body, err := json.Marshal(struct {
		Data any `json:"data"`
	}{resource})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewReader(body))
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rp := new(response[Resource[T]])
	if err := json.NewDecoder(resp.Body).Decode(rp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &rp.Data, nil
}
