package appstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement
func (c *Client) CreateAchievement(ctx context.Context, gameCenterID GameCenterID, attr *AchievementAttr) (*Achievement, error) {
	rq := newCreateRequest(attr, "gameCenterAchievements", relation{ID: string(gameCenterID), Type: "gameCenterDetails"})
	body, err := json.Marshal(rq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievements"
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	re := createResponse[Achievement]{}
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		return nil, fmt.Errorf("failed to decode response %s: %w", url, err)
	}

	return &re.Data, nil
}

// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_localization
func (c *Client) CreateAchievementLocalization(ctx context.Context, achievementID AchievementID, attr *AchievementLocalizationAttr) (*AchievementLocalization, error) {
	rq := newCreateRequest(attr, "gameCenterAchievementLocalizations", relation{ID: string(achievementID), Type: "gameCenterAchievements"})
	body, err := json.Marshal(rq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := "https://api.appstoreconnect.apple.com/v1/gameCenterAchievementLocalizations"
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	re := createResponse[AchievementLocalization]{}
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		return nil, fmt.Errorf("failed to decode response %s: %w", url, err)
	}

	return &re.Data, nil
}
