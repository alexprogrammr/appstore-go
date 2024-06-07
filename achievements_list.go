package appstore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementsresponse
type ListAchievementsResponse struct {
	Data  []Achievement `json:"data"`
	Links PagingLinks   `json:"links"`
	Meta  struct {
		Paging Paging `json:"paging"`
	} `json:"meta"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements
func (c *Client) ListAchievements(ctx context.Context, gameCenterID string) (*ListAchievementsResponse, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterDetails/%s/gameCenterAchievements", gameCenterID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	rp := ListAchievementsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&rp); err != nil {
		return nil, fmt.Errorf("failed to decode response %s: %w", url, err)
	}

	return &rp, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationsresponse
type ListAchievementLocalizationsResponse struct {
	Data  []AchievementLocalization `json:"data"`
	Links PagingLinks               `json:"links"`
	Meta  struct {
		Paging Paging `json:"paging"`
	} `json:"meta"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/list_all_localizations_for_an_achievement
func (c *Client) ListAchievementLocalizations(ctx context.Context, achievementID string) (*ListAchievementLocalizationsResponse, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/gameCenterAchievements/%s/localizations", achievementID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	rp := ListAchievementLocalizationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&rp); err != nil {
		return nil, fmt.Errorf("failed to decode response %s: %w", url, err)
	}

	return &rp, nil
}
