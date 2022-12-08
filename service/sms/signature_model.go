package sms

import "github.com/volcengine/volc-sdk-golang/base"

type GetSignatureAndOrderListRequest struct {
	SubAccount string `url:"subAccount"`
	Signature  string `url:"signature,omitempty"`
	PageIndex  int    `url:"pageIndex"`
	PageSize   int    `url:"pageSize"`
}

type GetSignatureAndOrderListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List  []*SmsSignatureInfo `json:"list"`
		Total int                 `json:"total"`
	}
}

type SmsSignatureInfo struct {
	Id          string         `json:"id"`
	ApplyId     string         `json:"applyId"`
	Content     string         `json:"content"`
	Source      string         `json:"source"`
	Application string         `json:"application"`
	CreatedTime int64          `json:"createdTime"`
	IsOrder     bool           `json:"isOrder"`
	Status      SmsOrderStatus `json:"status"`
	Reason      string         `json:"reason"`
}

type ApplySmsSignatureRequest struct {
	SubAccount    string `json:"subAccount"`
	Content       string `json:"content"`
	Source        string `json:"source"`
	Domain        string `json:"domain,omitempty"`
	Desc          string `json:"desc,omitempty"`
	UploadFileKey string `json:"uploadFileKey,omitempty"`
}

type ApplySmsSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsSignatureInfo
}

type DeleteSignatureRequest struct {
	SubAccount string `json:"subAccount"`
	Id         string `json:"id"`
	IsOrder    bool   `json:"isOrder"`
}

type DeleteSignatureResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           string
}
