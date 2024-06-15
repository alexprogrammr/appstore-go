package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_achievement
func (c *Client) UpdateAchievement(ctx context.Context, upd AchievementUpdate) (*Resource[Achievement], error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/" + upd.ID
	req := updateResource{
		ID:   upd.ID,
		Type: resourceTypeAchievements,
		Attr: upd,
	}

	resp, err := doUpdate[Achievement](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update achievement: %w", err)
	}

	return resp, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/edit_an_achievement_localization
func (c *Client) UpdateAchievementLocalization(ctx context.Context, upd AchievementLocalizationUpdate) (*Resource[AchievementLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations/" + upd.ID
	req := updateResource{
		ID:   upd.ID,
		Type: resourceTypeAchievementLocalizations,
		Attr: upd,
	}

	resp, err := doUpdate[AchievementLocalization](c, ctx, url, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update achievement localization: %w", err)
	}

	return resp, nil
}
