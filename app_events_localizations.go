package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/appeventlocalizationcreaterequest/data-data.dictionary/attributes-data.dictionary
type AppEventLocalization struct {
	Locale           Locale `json:"locale"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
}
