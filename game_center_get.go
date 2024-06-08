package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/read_the_state_of_game_center_for_an_app
func (c *Client) GetGameCenter(ctx context.Context, appID AppID) (*GameCenter, error) {
	url := fmt.Sprintf("https://api.appstoreconnect.apple.com/v1/apps/%s/gameCenterDetail", appID)

	resp, err := doGet[getResponse[GameCenter]](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get game center: %w", err)
	}

	return &resp.Data, nil
}
