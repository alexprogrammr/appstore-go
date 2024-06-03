package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/app
type App struct {
	ID         string        `json:"id"`
	Attributes AppAttributes `json:"attributes"`
	// TODO: add more fields
}

// https://developer.apple.com/documentation/appstoreconnectapi/app/attributes
type AppAttributes struct {
	Name     string `json:"name"`
	BundleID string `json:"bundleId"`
	SKU      string `json:"sku"`
	// TODO: add more fields
}
