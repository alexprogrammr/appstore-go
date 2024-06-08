package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement
func (c *Client) DeleteAchievement(ctx context.Context, id AchievementID) error {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/%s", id)

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement: %w", err)
	}

	return nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_localization
func (c *Client) DeleteAchievementLocalization(ctx context.Context, id AchievementLocalizationID) error {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations/%s", id)

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement localization: %w", err)
	}

	return nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_image
func (c *Client) DeleteAchievementImage(ctx context.Context, id AchievementImageID) error {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievementImages/%s", id)

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement image: %w", err)
	}

	return nil
}
