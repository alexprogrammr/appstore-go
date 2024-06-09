package appstore

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail/attributes
type GameCenter struct {
	ArcadeEnabled    bool `json:"arcadeEnabled"`
	ChallengeEnabled bool `json:"challengeEnabled"`
}
