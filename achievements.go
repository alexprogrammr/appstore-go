package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement
type Achievement struct {
	ID         string                `json:"id"`
	Attributes AchievementAttributes `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/attributes
type AchievementAttributes struct {
	ReferenceName    string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points           int    `json:"points"`
	Repeatable       bool   `json:"repeatable"`
	ShowBeforeEarned bool   `json:"showBeforeEarned"`
	Archived         bool   `json:"archived"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization
type AchievementLocalization struct {
	ID         string                            `json:"id"`
	Attributes AchievementLocalizationAttributes `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/attributes
type AchievementLocalizationAttributes struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
}
