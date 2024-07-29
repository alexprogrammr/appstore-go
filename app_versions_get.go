package appstore

import "context"

// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_information
func (c *Client) GetAppVersionByID(ctx context.Context, id string) (*Resource[AppVersion], error) {
	url := "https://api.appstoreconnect.apple.com/v1/appStoreVersions/" + id

	resp, err := doGet[AppVersion](c, ctx, url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
