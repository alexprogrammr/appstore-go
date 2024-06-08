package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_image_information
func (c *Client) GetAchievementImage(ctx context.Context, id AchievementLocalizationID) (*AchievementImage, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations/%s/gameCenterAchievementImage", id)

	resp, err := doGet[getResponse[AchievementImage]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get achievement image: %w", err)
	}

	return &resp.Data, nil
}
