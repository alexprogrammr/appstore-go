package appstore

import (
	"context"
	"fmt"
)

// https://developer.apple.com/documentation/appstoreconnectapi/read_the_state_of_game_center_for_an_app
func (c *Client) GetGameCenter(ctx context.Context, app *Resource[App]) (*Resource[GameCenter], error) {
	url := app.Links.Self + "/gameCenterDetail"

	resp, err := doGet[GameCenter](c, ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get game center: %w", err)
	}

	return resp, nil
}
