package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement
func (c *Client) CreateAchievement(ctx context.Context, gameCenterID GameCenterID, attr *AchievementAttr) (*Achievement, error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievements"
	req := newCreateRequest(attr, "gameCenterAchievements", relation{ID: string(gameCenterID), Type: "gameCenterDetails"})

	resp, err := doCreate[Achievement](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement: %w", err)
	}

	return resp, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_localization
func (c *Client) CreateAchievementLocalization(ctx context.Context, achievementID AchievementID, attr *AchievementLocalizationAttr) (*AchievementLocalization, error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations"
	req := newCreateRequest(attr, "gameCenterAchievementLocalizations", relation{ID: string(achievementID), Type: "gameCenterAchievements"})

	resp, err := doCreate[AchievementLocalization](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement localization: %w", err)
	}

	return resp, nil
}

type CreateAchievementImageRequest struct {
	Name string
	Data []byte
}

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_image
func (c *Client) CreateAchievementImage(ctx context.Context, id AchievementLocalizationID, r *CreateAchievementImageRequest) (*Asset, error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementImages"
	attr := CreateAssetAttr{
		Name: r.Name,
		Size: len(r.Data),
	}
	req := newCreateRequest(attr, "gameCenterAchievementImages", relation{ID: string(id), Type: "gameCenterAchievementLocalizations"})

	asset, err := doCreate[Asset](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create achievement image: %w", err)
	}

	asset, err = c.uploadAsset(ctx, asset, r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to upload achievement image: %w", err)
	}

	return asset, nil
}
