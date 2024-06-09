package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements
func (c *Client) ListAchievements(ctx context.Context, gameCenter *Resource[GameCenter]) ([]Resource[Achievement], error) {
	url := gameCenter.Links.Self + "/gameCenterAchievements"

	resp, err := doGet[listResponse[Achievement]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list achievements: %w", err)
	}

	return resp.Data, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_localizations_for_an_achievement
func (c *Client) ListAchievementLocalizations(ctx context.Context, ach *Resource[Achievement]) ([]Resource[AchievementLocalization], error) {
	url := ach.Links.Self + "/localizations"

	resp, err := doGet[listResponse[AchievementLocalization]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to list achievement localizations: %w", err)
	}

	return resp.Data, nil
}
