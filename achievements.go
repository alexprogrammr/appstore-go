package appstore

type AchievementID string

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement
type Achievement struct {
	ID   AchievementID   `json:"id"`
	Attr AchievementAttr `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/attributes
type AchievementAttr struct {
	ReferenceName    string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points           int    `json:"points"`
	Repeatable       bool   `json:"repeatable"`
	ShowBeforeEarned bool   `json:"showBeforeEarned"`
}

type AchievementLocalizationID string

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization
type AchievementLocalization struct {
	ID   AchievementLocalizationID   `json:"id"`
	Attr AchievementLocalizationAttr `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/attributes
type AchievementLocalizationAttr struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
}

type AchievementImageID string

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimage
type AchievementImage struct {
	ID   AchievementImageID `json:"id"`
	Attr Asset              `json:"attributes"`
}
