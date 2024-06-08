package appstore

import (
	"context"
	"fmt"
	"net/http"
)

// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_image
func (c *Client) DeleteAchievementImage(ctx context.Context, id AchievementImageID) error {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievementImages/%s", id)
	rq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	rq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(rq)
	if err != nil {
		return fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	return nil
}
