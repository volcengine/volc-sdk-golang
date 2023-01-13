package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type CornerInfo struct {
	LeftTop     []float64 `json:"left_top"`
	RightTop    []float64 `json:"right_top"`
	RightBottom []float64 `json:"right_bottom"`
	LeftBottom  []float64 `json:"left_bottom"`
}

type BankCardData struct {
	CardNumber  string      `json:"card_number"`
	CardCorners *CornerInfo `json:"card_corners"`
}

type BankCardResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *BankCardData          `json:"data"`
}

type BankCardDataV2 struct {
	ExpiredDate        string      `json:"expired_date"`
	ExpiredDateCorners *CornerInfo `json:"expired_date_corners"`
	Number             string      `json:"number"`
	NumberCorners      *CornerInfo `json:"number_corners"`
	BankId             string      `json:"bank_id"`
	BankName           string      `json:"bank_name"`
	CardName           string      `json:"card_name"`
	CardType           string      `json:"card_type"`
}

type BankCardResultV2 struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *BankCardDataV2        `json:"data"`
}
