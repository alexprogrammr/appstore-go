package appstore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailresponse
type GetGameCenterResponse struct {
	GameCenter GameCenter `json:"data"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/read_the_state_of_game_center_for_an_app
func (c *Client) GetGameCenter(ctx context.Context, appID string) (*GetGameCenterResponse, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/apps/%s/gameCenterDetail", appID)
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %s: %w", url, err)
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	rq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %s: %d", url, resp.StatusCode)
	}

	rp := GetGameCenterResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&rp); err != nil {
		return nil, fmt.Errorf("failed to decode response %s: %w", url, err)
	}

	return &rp, nil
}
