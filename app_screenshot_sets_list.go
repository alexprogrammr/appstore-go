package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_screenshot_sets_for_an_app_store_version_localization
func (c *Client) ListAppScreenshotSets(ctx context.Context, localization *Resource[AppVersionLocalization]) ([]Resource[AppScreenshotSet], error) {
	url := localization.Links.Self + "/appScreenshotSets"

	resp, err := doList[AppScreenshotSet](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app screenshot sets: %w", err)
	}

	return resp, nil
}
