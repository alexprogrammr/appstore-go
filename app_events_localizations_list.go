package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-appevents-_id_-localizations
func (c *Client) ListAppEventLocalizations(ctx context.Context, event *Resource[AppEvent]) ([]Resource[AppEventLocalization], error) {
	url := event.Links.Self + "/localizations"

	resp, err := doList[AppEventLocalization](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app event localizations: %w", err)
	}

	return resp, nil
}
