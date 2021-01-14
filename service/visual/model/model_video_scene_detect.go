package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type VideoSceneDetectData struct {
	Results []string `json:"results"`
}

type VideoSceneDetectResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *VideoSceneDetectData  `json:"data"`
}
