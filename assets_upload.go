package appstore

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func (c *Client) uploadAsset(ctx context.Context, asset *Resource[Asset], data []byte) (*Resource[Asset], error) {
	wg := sync.WaitGroup{}
	errs := make(chan error, len(asset.Attr.Operations))

	for _, op := range asset.Attr.Operations {
		wg.Add(1)

		go func(op UploadOperation) {
			defer wg.Done()

			buff := bytes.NewReader(data[op.Offset : op.Offset+op.Length])
			err := c.uploadAssetChunk(ctx, op, buff)
			if err != nil {
				errs <- fmt.Errorf("failed to upload asset chunk: %w", err)
			}
		}(op)
	}

	wg.Wait()
	close(errs)

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to upload asset: %w", <-errs)
	}

	if err := c.commitAsset(ctx, asset, data); err != nil {
		return nil, fmt.Errorf("failed to commit asset: %w", err)
	}

	resp, err := doGet[Asset](c, ctx, asset.Links.Self)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}

	return resp, nil
}

func (c *Client) uploadAssetChunk(ctx context.Context, op UploadOperation, data io.Reader) error {
	req, err := http.NewRequestWithContext(ctx, op.Method, op.URL, data)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for _, h := range op.Headers {
		req.Header.Set(h.Name, h.Value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

type commitAssetRequest struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Attr struct {
			Uploaded bool   `json:"uploaded"`
			Checksum string `json:"sourceFileChecksum,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

func (c *Client) commitAsset(ctx context.Context, asset *Resource[Asset], data []byte) error {
	commit := commitAssetRequest{}
	commit.Data.ID = asset.ID
	commit.Data.Type = asset.Type
	commit.Data.Attr.Uploaded = true

	// achievements images do not require checksum
	if asset.Type != resourceTypeAchievementImages {
		checksum := md5.Sum(data)
		commit.Data.Attr.Checksum = hex.EncodeToString(checksum[:])
	}

	body, err := json.Marshal(commit)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, asset.Links.Self, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
