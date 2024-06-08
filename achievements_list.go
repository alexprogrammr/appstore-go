package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements
func (c *Client) ListAchievements(ctx context.Context, gameCenterID GameCenterID) ([]Achievement, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterDetails/%s/gameCenterAchievements", gameCenterID)

	resp, err := doGet[listResponse[Achievement]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list achievements: %w", err)
	}

	return resp.Data, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_localizations_for_an_achievement
func (c *Client) ListAchievementLocalizations(ctx context.Context, achievementID AchievementID) ([]AchievementLocalization, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/%s/localizations", achievementID)

	resp, err := doGet[listResponse[AchievementLocalization]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list achievement localizations: %w", err)
	}

	return resp.Data, nil
}
