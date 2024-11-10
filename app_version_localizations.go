package appstore

type AppVersionLocalization struct {
	Locale       Locale `json:"locale"`
	WhatsNew     string `json:"whatsNew"`
	Description  string `json:"description"`
	Promotional  string `json:"promotionalText"`
	Keywords     string `json:"keywords"`
	MarketingURL string `json:"marketingUrl"`
	SupportURL   string `json:"supportUrl"`
}

type AppVersionLocalizationUpdate struct {
	ID           string `json:"-"`
	WhatsNew     string `json:"whatsNew,omitempty"`
	Description  string `json:"description,omitempty"`
	Promotional  string `json:"promotionalText,omitempty"`
	Keywords     string `json:"keywords,omitempty"`
	MarketingURL string `json:"marketingUrl,omitempty"`
	SupportURL   string `json:"supportUrl,omitempty"`
}
