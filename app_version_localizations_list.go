package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_version_localizations_for_an_app_store_version
func (c *Client) ListAppVersionLocalizations(ctx context.Context, version *Resource[AppVersion]) ([]Resource[AppVersionLocalization], error) {
	url := version.Links.Self + "/appStoreVersionLocalizations"

	resp, err := doList[AppVersionLocalization](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app version localizations: %w", err)
	}

	return resp, nil
}
