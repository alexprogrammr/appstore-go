package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_versions_for_an_app
func (c *Client) ListAppVersions(ctx context.Context, app *Resource[App]) ([]Resource[AppVersion], error) {
	url := app.Links.Self + "/appStoreVersions"

	resp, err := doList[AppVersion](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app versions: %w", err)
	}

	return resp, nil
}
