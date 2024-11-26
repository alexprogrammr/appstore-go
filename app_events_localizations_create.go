package appstore

import (
	"context"
	"fmt"
)

func (c *Client) CreateAppEventLocalization(ctx context.Context, event *Resource[AppEvent], loc *AppEventLocalization) (*Resource[AppEventLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/" + resourceTypeAppEventLocalizations
	req := createResource{
		Type: resourceTypeAppEventLocalizations,
		Attr: loc,
		Relations: relations{
			AppEvent: &relationd{
				Data: relation{
					ID:   event.ID,
					Type: event.Type,
				},
			},
		},
	}

	resp, err := doCreate[AppEventLocalization](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create app event localization: %w", err)
	}

	return resp, nil
}
