package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/attributes
type Achievement struct {
	ReferenceName    string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points           int    `json:"points"`
	Repeatable       bool   `json:"repeatable"`
	ShowBeforeEarned bool   `json:"showBeforeEarned"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/attributes
type AchievementLocalization struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
}
