package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/app/attributes
type App struct {
	Name          string `json:"name"`
	BundleID      string `json:"bundleId"`
	SKU           string `json:"sku"`
	PrimaryLocale Locale `json:"primaryLocale"`
	// TODO: add more fields
}
