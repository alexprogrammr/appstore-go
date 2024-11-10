package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/patch-v1-appstoreversionlocalizations-_id_
func (c *Client) UpdateAppVersionLocalization(ctx context.Context, upd AppVersionLocalizationUpdate) (*Resource[AppVersionLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/appStoreVersionLocalizations/" + upd.ID
	req := updateResource{
		ID:   upd.ID,
		Type: resourceTypeAppVersionLocalizations,
		Attr: upd,
	}

	resp, err := doUpdate[AppVersionLocalization](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update app version localization: %w", err)
	}

	return resp, nil
}
