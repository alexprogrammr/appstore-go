package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_information
func (c *Client) GetAchievementByID(ctx context.Context, id string) (*Resource[Achievement], error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/" + id

	resp, err := doGet[getResponse[Achievement]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get achievement: %w", err)
	}

	return &resp.Data, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_localization_information
func (c *Client) GetAchievementLocalizationByID(ctx context.Context, id string) (*Resource[AchievementLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations/" + id

	resp, err := doGet[getResponse[AchievementLocalization]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get achievement localization: %w", err)
	}

	return &resp.Data, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_image_information
func (c *Client) GetAchievementImage(ctx context.Context, ach *Resource[AchievementLocalization]) (*Resource[Asset], error) {
	url := ach.Links.Self + "/gameCenterAchievementImage"

	resp, err := doGet[getResponse[Asset]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get achievement image: %w", err)
	}

	return &resp.Data, nil
}

func (c *Client) GetAchievementImageByID(ctx context.Context, id string) (*Resource[Asset], error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementImages/" + id

	resp, err := doGet[getResponse[Asset]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get achievement image: %w", err)
	}

	return &resp.Data, nil
}
