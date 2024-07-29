package appstore

import (
	"context"
	"fmt"
)

func (c *Client) CreateAppScreenshot(ctx context.Context, set *Resource[AppScreenshotSet], name string, data []byte) (*Resource[Asset], error) {
	url := "https://api.appstoreconnect.apple.com/v1/" + resourceTypeAppScreenshots
	req := createResource{
		Type: resourceTypeAppScreenshots,
		Attr: createAsset{
			Name: name,
			Size: len(data),
		},
		Relations: relations{
			AppScreenshotSet: &relationd{
				Data: relation{
					ID:   set.ID,
					Type: set.Type,
				},
			},
		},
	}

	asset, err := doCreate[Asset](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create app screenshot: %w", err)
	}

	asset, err = c.uploadAsset(ctx, asset, data)
	if err != nil {
		return nil, fmt.Errorf("failed to upload app screenshot: %w", err)
	}

	return asset, nil
}
