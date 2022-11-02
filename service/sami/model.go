package sami

type GetTokenReq struct {
	Version    string `json:"token_version"`
	Appkey     string `json:"appkey"`
	Expiration int64  `json:"expiration"`
}

type GetTokenResponse struct {
	StatusCode int32  `form:"status_code,required" json:"status_code,required" query:"status_code,required"`
	StatusText string `form:"status_text,required" json:"status_text,required" query:"status_text,required"`
	TaskId     string `form:"task_id,required" json:"task_id,required" query:"task_id,required"`
	Token      string `form:"token,required" json:"token,required" query:"token,required"`
	ExpiresAt  int64  `form:"expires_at,required" json:"expires_at,required" query:"expires_at,required"`
}

type InvokeRequest struct {
	Data    []byte  `json:"data,omitempty"`
	Payload *string `json:"payload,omitempty"`
}

type InvokeResponse struct {
	StatusCode int32   `form:"status_code,required" json:"status_code,required" query:"status_code,required"`
	StatusText string  `form:"status_text,required" json:"status_text,required" query:"status_text,required"`
	TaskId     string  `form:"task_id,required" json:"task_id,required" query:"task_id,required"`
	Namespace  string  `form:"namespace,required" json:"namespace,required" query:"namespace,required"`
	Payload    *string `form:"payload,omitempty" json:"payload,omitempty" query:"payload,omitempty"`
	Data       []byte  `form:"data,omitempty" json:"data,omitempty" query:"data,omitempty"`
	State      *string `form:"state,omitempty" json:"state,omitempty" query:"state,omitempty"`
}

type SpeechSynthesisReq struct {
	Text    string      `json:"text"`
	SSML    string      `json:"ssml"`
	Speaker string      `json:"speaker,omitempty"`
	Config  AudioConfig `json:"audio_config"`
}

type AudioConfig struct {
	Format          string `json:"format"`
	SampleRate      int    `json:"sample_rate"`
	SpeechRate      int    `json:"speech_rate"`
	PitchRate       int    `json:"pitch_rate"`
	EnableTimestamp bool   `json:"enable_timestamp"`
}

type SpeechSynthesisResponse struct {
	AudioDuration float32     `json:"duration,omitempty"`
	Words         []WordInfo  `json:"words,omitempty"`
	Phonemes      []PhoneInfo `json:"phonemes,omitempty"`
}

// WordInfo includes ASR word information
type WordInfo struct {
	Word      string  `json:"word"`
	StartTime float32 `json:"start_time"`
	EndTime   float32 `json:"end_time"`
}

// PhoneInfo includes phone information
type PhoneInfo struct {
	Phone     string  `json:"phone"`
	StartTime float32 `json:"start_time"`
	EndTime   float32 `json:"end_time"`
}
