package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/attributes
type Achievement struct {
	ReferenceName    string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points           int    `json:"points"`
	Repeatable       bool   `json:"repeatable"`
	ShowBeforeEarned bool   `json:"showBeforeEarned"`
}

type AchievementUpdate struct {
	ID               string `json:"-"`
	ReferenceName    string `json:"referenceName,omitempty"`
	Points           int    `json:"points,omitempty"`
	Repeatable       bool   `json:"repeatable,omitempty"`
	ShowBeforeEarned bool   `json:"showBeforeEarned,omitempty"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/attributes
type AchievementLocalization struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
}

type AchievementLocalizationUpdate struct {
	ID                      string `json:"-"`
	Name                    string `json:"name,omitempty"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription,omitempty"`
	AfterEarnedDescription  string `json:"afterEarnedDescription,omitempty"`
}
