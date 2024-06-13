package appstore

import "context"

// https://developer.apple.com/documentation/appstoreconnectapi/read_app_information
func (c *Client) GetApp(ctx context.Context, id string) (*Resource[App], error) {
	url := "https://api.appstoreconnect.apple.com/v1/apps/" + id

	resp, err := doGet[App](c, ctx, url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
