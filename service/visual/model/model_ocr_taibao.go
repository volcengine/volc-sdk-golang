package model

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type TaibaoMainInfo struct {
	Title           string `json:"title"`
	ChnName         string `json:"chn_name"`
	EngName         string `json:"eng_name"`
	BirthDate       string `json:"birth_date"`
	Sex             string `json:"sex"`
	ValidDate       string `json:"valid_date"`
	RegisterOffice  string `json:"register_office"`
	RegisterAddress string `json:"register_address"`
	Num             string `json:"num"`
	RegisterCount   string `json:"register_count"`
}

type TaibaoData struct {
	MainInfo *TaibaoMainInfo `json:"license_main"`
}

type TaibaoResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Data             *TaibaoData            `json:"data"`
}
