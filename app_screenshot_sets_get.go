package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/read_app_screenshot_set_information
func (c *Client) AppScreenshotSetByID(ctx context.Context, id string) (*Resource[AppScreenshotSet], error) {
	url := "https://api.appstoreconnect.apple.com/v1/appScreenshotSets/" + id

	resp, err := doGet[AppScreenshotSet](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app screenshot sets: %w", err)
	}

	return resp, nil
}
