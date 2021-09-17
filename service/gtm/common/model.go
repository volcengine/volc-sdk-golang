package common

import "encoding/json"

type TopResponse struct {
	ResponseMetadata TopRespMeta     `form:"ResponseMetadata" json:"ResponseMetadata"`
	Result           json.RawMessage `form:"Result" json:"Result"`
}

type TopRespMeta struct {
	Action    string       `form:"Action" json:"Action"`
	Error     TopRespError `form:"Error" json:"Error"`
	Region    string       `form:"Region" json:"Region"`
	RequestID string       `form:"RequestId" json:"RequestId"`
	Service   string       `form:"Service" json:"Service"`
	Version   string       `form:"Version" json:"Version"`
}

type TopRespError struct {
	Code    string `form:"Code" json:"Code"`
	CodeN   int64  `form:"CodeN" json:"CodeN"`
	Message string `form:"Message" json:"Message"`
}
