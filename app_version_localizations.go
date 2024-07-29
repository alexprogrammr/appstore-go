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
