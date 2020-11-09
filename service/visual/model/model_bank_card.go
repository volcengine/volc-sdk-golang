package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type CornerInfo struct {
	LeftTop     []int `json:"left_top"`
	RightTop    []int `json:"right_top"`
	RightBottom []int `json:"right_bottom"`
	LeftBottom  []int `json:"left_bottom"`
}

type BankCardData struct {
	CardNumber  string      `json:"card_number"`
	CardCorners *CornerInfo `json:"card_corners"`
}

type BankCardResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *BankCardData          `json:"data"`
}
