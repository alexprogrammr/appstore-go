package appstore

import "context"

// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-appevents-_id_
func (c *Client) GetAppEventByID(ctx context.Context, id string) (*Resource[AppEvent], error) {
	url := "https://api.appstoreconnect.apple.com/v1/appEvents/" + id

	resp, err := doGet[AppEvent](c, ctx, url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
