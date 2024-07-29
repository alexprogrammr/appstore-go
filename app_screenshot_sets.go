package appstore

type DisplayType string

const (
	DisplayTypeAppIPhone67               DisplayType = "APP_IPHONE_67"
	DisplayTypeAppIPhone65               DisplayType = "APP_IPHONE_65"
	DisplayTypeAppIPhone61               DisplayType = "APP_IPHONE_61"
	DisplayTypeAppIPhone58               DisplayType = "APP_IPHONE_58"
	DisplayTypeAppIPhone55               DisplayType = "APP_IPHONE_55"
	DisplayTypeAppIPhone47               DisplayType = "APP_IPHONE_47"
	DisplayTypeAppIPhone40               DisplayType = "APP_IPHONE_40"
	DisplayTypeAppIPhone35               DisplayType = "APP_IPHONE_35"
	DisplayTypeAppIPadPro3Gen129         DisplayType = "APP_IPAD_PRO_3GEN_129"
	DisplayTypeAppIPadPro3Gen11          DisplayType = "APP_IPAD_PRO_3GEN_11"
	DisplayTypeAppIPadPro129             DisplayType = "APP_IPAD_PRO_129"
	DisplayTypeAppIPad105                DisplayType = "APP_IPAD_105"
	DisplayTypeAppIPad97                 DisplayType = "APP_IPAD_97"
	DisplayTypeAppWatchUltra             DisplayType = "APP_WATCH_ULTRA"
	DisplayTypeAppWatchSeries7           DisplayType = "APP_WATCH_SERIES_7"
	DisplayTypeAppWatchSeries4           DisplayType = "APP_WATCH_SERIES_4"
	DisplayTypeAppWatchSeries3           DisplayType = "APP_WATCH_SERIES_3"
	DisplayTypeAppDesktop                DisplayType = "APP_DESKTOP"
	DisplayTypeAppAppleTV                DisplayType = "APP_APPLE_TV"
	DisplayTypeIMessageAppIPhone67       DisplayType = "IMESSAGE_APP_IPHONE_67"
	DisplayTypeIMessageAppIPhone65       DisplayType = "IMESSAGE_APP_IPHONE_65"
	DisplayTypeIMessageAppIPhone61       DisplayType = "IMESSAGE_APP_IPHONE_61"
	DisplayTypeIMessageAppIPhone58       DisplayType = "IMESSAGE_APP_IPHONE_58"
	DisplayTypeIMessageAppIPhone55       DisplayType = "IMESSAGE_APP_IPHONE_55"
	DisplayTypeIMessageAppIPhone47       DisplayType = "IMESSAGE_APP_IPHONE_47"
	DisplayTypeIMessageAppIPhone40       DisplayType = "IMESSAGE_APP_IPHONE_40"
	DisplayTypeIMessageAppIPadPro3Gen129 DisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_129"
	DisplayTypeIMessageAppIPadPro3Gen11  DisplayType = "IMESSAGE_APP_IPAD_PRO_3GEN_11"
	DisplayTypeIMessageAppIPadPro129     DisplayType = "IMESSAGE_APP_IPAD_PRO_129"
	DisplayTypeIMessageAppIPad105        DisplayType = "IMESSAGE_APP_IPAD_105"
	DisplayTypeIMessageAppIPad97         DisplayType = "IMESSAGE_APP_IPAD_97"
	DisplayTypeAppAppleVisionPro         DisplayType = "APP_APPLE_VISION_PRO"
)

type AppScreenshotSet struct {
	DisplayType DisplayType `json:"screenshotDisplayType"`
}
