package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_apps
func (c *Client) ListApps(ctx context.Context) ([]Resource[App], error) {
	url := "https://api.appstoreconnect.apple.com/v1/apps"

	resp, err := doGet[listResponse[App]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list apps: %w", err)
	}

	return resp.Data, nil
}
