package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps-_id_-appevents
func (c *Client) ListAppEvents(ctx context.Context, app *Resource[App]) ([]Resource[AppEvent], error) {
	url := app.Links.Self + "/appEvents"

	resp, err := doList[AppEvent](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list app events: %w", err)
	}

	return resp, nil
}
