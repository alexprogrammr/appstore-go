package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail
type GameCenter struct {
	ID         string               `json:"id"`
	Attributes GameCenterAttributes `json:"attributes"`
	// TODO: add more fields
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail/attributes
type GameCenterAttributes struct {
	ArcadeEnabled    bool `json:"arcadeEnabled"`
	ChallengeEnabled bool `json:"challengeEnabled"`
}
