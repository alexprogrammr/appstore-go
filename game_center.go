package appstore

type GameCenterID string

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail
type GameCenter struct {
	ID   GameCenterID   `json:"id"`
	Attr GameCenterAttr `json:"attributes"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail/attributes
type GameCenterAttr struct {
	ArcadeEnabled    bool `json:"arcadeEnabled"`
	ChallengeEnabled bool `json:"challengeEnabled"`
}
