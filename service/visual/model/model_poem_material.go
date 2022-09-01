package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type PoemMaterialData struct {
	Poems []string `json:"poems"`
}

type PoemMaterialResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *PoemMaterialData      `json:"data"`
}
