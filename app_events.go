package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/appevent/attributes-data.dictionary
type AppEvent struct {
	ReferenceName string `json:"referenceName"`
	PrimaryLocale Locale `json:"primaryLocale"`
	// TODO: add more fields
}
