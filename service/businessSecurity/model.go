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

type MobileSecondSaleRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileSecondSaleResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	Data      MobileSecondSaleData `json:"Data"`
}

type MobileSecondSaleData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

type MobileEmptyCheckRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileEmptyCheckResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	Data      MobileEmptyCheckData `json:"Data"`
}

type MobileEmptyCheckData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

type MobileOnlineStatusRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileOnlineStatusResponse struct {
	RequestId string                 `json:"RequestId"`
	Code      int                    `json:"Code"`
	Message   string                 `json:"Message"`
	Data      MobileOnlineStatusData `json:"Data"`
}

type MobileOnlineStatusData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

type MobileOnlineTimeRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileOnlineTimeResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	Data      MobileOnlineTimeData `json:"Data"`
}

type MobileOnlineTimeData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
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
