package sms

import "github.com/volcengine/volc-sdk-golang/base"

type GetSubAccountListRequest struct {
	SubAccount     string `url:"subAccount"`
	SubAccountName string `url:"subAccountName"`
	PageIndex      int    `url:"pageIndex"`
	PageSize       int    `url:"pageSize"`
}

type GetSubAccountListResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *struct {
		List  []*SmsSubAccountInfo `json:"list"`
		Total int                  `json:"total"`
	}
}

type SmsSubAccountInfo struct {
	SubAccountId   string `json:"subAccountId"`
	SubAccountName string `json:"subAccountName"`
	CreatedTime    int64  `json:"createdTime"`
	Status         int    `json:"status"`
	Desc           string `json:"desc"`
}

type GetSubAccountDetailRequest struct {
	SubAccount string `url:"subAccount"`
}

type GetSubAccountDetailResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *SmsSubAccountDetail
}

type SmsSubAccountDetail struct {
	SubAccountId       string `json:"subAccountId"`
	SubAccountName     string `json:"subAccountName"`
	EnabledChannelType []*struct {
		Name  string         `json:"name"`
		Value SmsChannelType `json:"value"`
		Area  Area
	} `json:"enabledChannelType"`
	Status      int    `json:"status"`
	Desc        string `json:"desc"`
	CreatedTime int64  `json:"createdTime"`
}
