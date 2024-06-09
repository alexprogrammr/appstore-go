package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement
func (c *Client) CreateAchievement(ctx context.Context, gameCenter *Resource[GameCenter], ach *Achievement) (*Resource[Achievement], error) {
	url := "https://api.appstoreconnect.apple.com/v1/" + resourceTypeAchievements
	req := createResource{
		Type: resourceTypeAchievements,
		Attr: ach,
		Relations: relations{
			GameCenter: &relationd{
				Data: relation{
					ID:   gameCenter.ID,
					Type: gameCenter.Type,
				},
			},
		},
	}

	resp, err := doCreate[Achievement](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement: %w", err)
	}

	return resp, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_localization
func (c *Client) CreateAchievementLocalization(ctx context.Context, ach *Resource[Achievement], loc *AchievementLocalization) (*Resource[AchievementLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/" + resourceTypeAchievementLocalizations
	req := createResource{
		Type: resourceTypeAchievementLocalizations,
		Attr: loc,
		Relations: relations{
			Achievement: &relationd{
				Data: relation{
					ID:   ach.ID,
					Type: ach.Type,
				},
			},
		},
	}

	resp, err := doCreate[AchievementLocalization](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement localization: %w", err)
	}

	return resp, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_image
func (c *Client) CreateAchievementImage(ctx context.Context, loc *Resource[AchievementLocalization], name string, data []byte) (*Resource[Asset], error) {
	url := "https://api.appstoreconnect.apple.com/v1/" + resourceTypeAchievementImages
	req := createResource{
		Type: resourceTypeAchievementImages,
		Attr: createAsset{
			Name: name,
			Size: len(data),
		},
		Relations: relations{
			AchievementLocalization: &relationd{
				Data: relation{
					ID:   loc.ID,
					Type: loc.Type,
				},
			},
		},
	}

	asset, err := doCreate[Asset](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement image: %w", err)
	}

	asset, err = c.uploadAsset(ctx, asset, data)
	if err != nil {
		return nil, fmt.Errorf("failed to upload achievement image: %w", err)
	}

	return asset, nil
}
