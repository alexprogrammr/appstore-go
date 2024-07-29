package appstore

import "time"

type AppVersion struct {
	Platform        string    `json:"platform"`
	AppStoreState   string    `json:"appStoreState"`
	AppVersionState string    `json:"appVersionState"`
	Version         string    `json:"versionString"`
	Copyright       string    `json:"copyright"`
	Downloadable    bool      `json:"downloadable"`
	CreatedDate     time.Time `json:"createdDate"`
	ReleaseDate     time.Time `json:"earliestReleaseDate"`
	ReleaseType     string    `json:"releaseType"`
	ReviewType      string    `json:"reviewType"`
}
