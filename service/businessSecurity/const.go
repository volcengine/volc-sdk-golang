package businessSecurity

// Product
const (
	PrdBusinessSecurity = "business_security"
	PrdAdblocker        = "adblocker"
	PrdElemVerify       = "element_verify"
	PRDMobileRisk       = "mobile_security"
	PRDContentRisk      = "content_security"
)

// service for content_security  product
const (
	ServiceContentVideoRisk        = "video_risk"
	ServiceContentImageContentRisk = "image_content_risk"
	ServiceContentTextRisk         = "text_risk"
	ServiceContentAudioRisk        = "audio_risk"
	ServiceContentVideoLiveRisk    = "video_live_risk"
	ServiceContentAudioLiveRisk    = "audio_live_risk"
)

// unit type
const (
	MONTHLY = "month"
	DAILY   = "day"
)

// RiskType
const (
	RiskTypeText  = "text"
	RiskTypeImage = "image"
	RiskTypeAudio = "audio"
)
