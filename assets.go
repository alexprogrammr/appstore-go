package appstore

import "fmt"

type AssetState string

const (
	AssetStateReadyForUpload AssetState = "AWAITING_UPLOAD"
	AssetStateUploadComplete AssetState = "UPLOAD_COMPLETE"
	AssetStateComplete       AssetState = "COMPLETE"
	AssetStateFailed         AssetState = "FAILED"
)

type AssetStateError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AssetDeliveryState struct {
	State    AssetState        `json:"state"`
	Errors   []AssetStateError `json:"errors"`
	Warnings []AssetStateError `json:"warnings"`
}

func (s AssetDeliveryState) IsComplete() bool {
	return s.State == AssetStateComplete
}

func (s AssetDeliveryState) Error() error {
	if s.State == AssetStateFailed {
		return fmt.Errorf("asset failed: %v", s.Errors)
	}

	return nil
}

type ImageAsset struct {
	TemplateURL string `json:"templateUrl"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type HttpHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UploadOperation struct {
	Method  string       `json:"method"`
	URL     string       `json:"url"`
	Offset  int          `json:"offset"`
	Length  int          `json:"length"`
	Headers []HttpHeader `json:"requestHeaders"`
}

type Asset struct {
	Name       string             `json:"fileName"`
	Size       int                `json:"fileSize"`
	State      AssetDeliveryState `json:"assetDeliveryState"`
	Image      ImageAsset         `json:"imageAsset"`
	Operations []UploadOperation  `json:"uploadOperations"`
}

type CreateAsset struct {
	Name string `json:"fileName"`
	Size int    `json:"fileSize"`
}
