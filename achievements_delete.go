package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement
func (c *Client) DeleteAchievement(ctx context.Context, ach *Resource[Achievement]) error {
	if err := doDelete(c, ctx, ach.Links.Self); err != nil {
		return fmt.Errorf("failed to delete achievement: %w", err)
	}

	return nil
}

func (c *Client) DeleteAchievementByID(ctx context.Context, id string) error {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/" + id

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement: %w", err)
	}

	return nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_localization
func (c *Client) DeleteAchievementLocalization(ctx context.Context, ach *Resource[AchievementLocalization]) error {
	if err := doDelete(c, ctx, ach.Links.Self); err != nil {
		return fmt.Errorf("failed to delete achievement localization: %w", err)
	}

	return nil
}

func (c *Client) DeleteAchievementLocalizationByID(ctx context.Context, id string) error {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations/" + id

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement localization: %w", err)
	}

	return nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_image
func (c *Client) DeleteAchievementImage(ctx context.Context, asset *Resource[Asset]) error {
	if err := doDelete(c, ctx, asset.Links.Self); err != nil {
		return fmt.Errorf("failed to delete achievement image: %w", err)
	}

	return nil
}

func (c *Client) DeleteAchievementImageByID(ctx context.Context, id string) error {
	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementImages/" + id

	if err := doDelete(c, ctx, url); err != nil {
		return fmt.Errorf("failed to delete achievement image: %w", err)
	}

	return nil
}
