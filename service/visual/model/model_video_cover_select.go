package model

import "github.com/volcengine/volc-sdk-golang/base"

type ImageInfoData struct {
	Score float64 `json:"score"`
	Data  string  `json:"data"`
}

type ImageInfoResult struct {
	Results []ImageInfoData `json:"results"`
}

type VideoCoverSelectResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *ImageInfoResult       `json:"data"`
}
