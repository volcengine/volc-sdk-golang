package businessSecurity

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/url"
)

type RiskDetectionRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type RiskDetectionResponse struct {
	RequestId string       `json:"RequestId"`
	Code      int          `json:"Code"`
	Message   string       `json:"Message"`
	Data      DecisionData `json:"Data"`
}

type DecisionData struct {
	Score  int      `json:"Score"`
	Tags   []string `json:"Tags"`
	Detail string   `json:"Detail"`
}

type AsyncRiskDetectionRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type AsyncRiskDetectionResponse struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type RiskResultRequest struct {
	AppId     int64  `json:"AppId" form:"AppId"`
	Service   string `json:"Service" form:"Service"`
	StartTime int64  `json:"StartTime" form:"StartTime"`
	EndTime   int64  `json:"EndTime" form:"EndTime"`
	Page
}

type RiskResultResponse struct {
	RequestId string                   `json:"RequestId"`
	Code      int                      `json:"Code"`
	Message   string                   `json:"Message"`
	Data      []map[string]interface{} `json:"Data"`
	Page      Page                     `json:"page"`
}

type DataReportRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type DataReportResponse struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type Page struct {
	PageNum  int64 `json:"PageNum" form:"PageNum"`
	PageSize int64 `json:"PageSize" form:"PageSize"`
	Total    int64 `json:"Total" form:"Total"` // used when return page information in response, no need when request
}

type ElementVerifyRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type ElementVerifyResponse struct {
	RequestId string            `json:"RequestId"`
	Code      int               `json:"Code"`
	Message   string            `json:"Message"`
	Data      ElementVerifyData `json:"Data"`
}

type ElementVerifyData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

type ElementVerifyResponseV2 struct {
	RequestId string              `json:"RequestId"`
	Code      int                 `json:"Code"`
	Message   string              `json:"Message"`
	Data      ElementVerifyDataV2 `json:"Data"`
}

type ElementVerifyDataV2 struct {
	Status int                     `json:"Status"`
	Detail ElementVerifyDataDetail `json:"Detail"`
}

type ElementVerifyDataDetail struct {
}

type MobileStatusRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileStatusResponse struct {
	RequestId string           `json:"RequestId"`
	Code      int              `json:"Code"`
	Message   string           `json:"Message"`
	Data      MobileStatusData `json:"Data"`
}

type MobileStatusData struct {
	Status int    `json:"Status"`
	Mobile string `json:"Mobile"`
	Detail string `json:"Detail"`
}

type MobileStatusResponseV2 struct {
	RequestId string             `json:"RequestId"`
	Code      int                `json:"Code"`
	Message   string             `json:"Message"`
	Data      MobileStatusDataV2 `json:"Data"`
}

type MobileStatusDataV2 struct {
	Status int                `json:"Status"`
	Mobile string             `json:"Mobile"`
	Detail MobileStatusDetail `json:"Detail"`
}

type MobileStatusDetail struct {
	Province     string `json:"Province"`
	ProvinceCode string `json:"ProvinceCode"`
	City         string `json:"City"`
	CityCode     string `json:"CityCode"`
	ISP          string `json:"ISP"`
	Mobile       string `json:"Mobile"`
	PostCode     string `json:"PostCode"`
	OldISP       int    `json:"OldISP"`
	NewISP       int    `json:"NewISP"`
}

func (r *RiskResultRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}
