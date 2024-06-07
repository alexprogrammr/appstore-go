package appstore

type AppID string

// https://developer.apple.com/documentation/appstoreconnectapi/app
type App struct {
	ID   AppID   `json:"id"`
	Attr AppAttr `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/app/attributes
type AppAttr struct {
	Name     string `json:"name"`
	BundleID string `json:"bundleId"`
	SKU      string `json:"sku"`
	// TODO: add more fields
}
