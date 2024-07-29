package appstore

import "context"

// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_localization_information
func (c *Client) GetAppVersionLocalizationByID(ctx context.Context, id string) (*Resource[AppVersionLocalization], error) {
	url := "https://api.appstoreconnect.apple.com/v1/appStoreVersionLocalizations/" + id

	resp, err := doGet[AppVersionLocalization](c, ctx, url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
