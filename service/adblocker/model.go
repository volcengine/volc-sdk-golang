package adblocker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
)

type AdBlockRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type AdBlockResponse struct {
	RequestId string       `json:"RequestId"`
	Code      int          `json:"Code"`
	Message   string       `json:"Message"`
	Data      DecisionData `json:"Data"`
}

type DecisionData struct {
	Decision string   `json:"Decision"`
	Detail   string   `json:"Detail"`
	Tags     []string `json:"Tags"`
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

// openapi
type SimpleRiskStatResponse struct {
	Result *SimpleRiskStatResult `json:"Result"`
}

type OpenResult struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type SimpleRiskStatResult struct {
	OpenResult
	Data *SimpleProductStatisticsResult `json:"Data"`
}

func (result *OpenResult) GetErr() error {
	return errors.New(fmt.Sprintf("err msg is %s, reqId is %s", result.Message, result.RequestId))
}

type CommonProductStatisticsReq struct {
	Product    string  `json:"Product" form:"Product" query:"Product"`
	UnitType   string  `json:"UnitType" form:"UnitType" query:"UnitType"`
	Parameters string  `json:"Parameters" form:"Parameters" query:"Parameters"`
	AppId      *int64  `json:"AppId" form:"AppId" query:"AppId"`
	Service    *string `json:"Service" form:"Service" query:"Service"`
}
type CommonProductStatisticsParams struct {
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	NeedAppDetail bool   `json:"need_app_detail"`
	OperateTime   int64  `json:"operate_time"`
}
type SimpleProductStatisticsParams struct {
	CommonProductStatisticsParams
	NeedServiceDetail bool `json:"need_service_detail"`
}

type SimpleProductStatisticsResult struct {
	Total  *SimpleProductStatisticsResultTotal    `json:"Total"`
	Detail []*SimpleProductStatisticsResultDetail `json:"Detail"`
}

type CommonProductStatisticsResultTotal struct {
	AccountId        int64  `json:"AccountId"`
	RequestCnt       int64  `json:"RequestCnt"`
	ChargeRequestCnt *int64 `json:"ChargeRequestCnt,omitempty"`
}

type CommonProductStatisticsResultDetail struct {
	AccountId  int64  `json:"AccountId"`
	RequestCnt int64  `json:"RequestCnt"`
	DateTime   string `json:"DateTime"`

	ChargeRequestCnt *int64  `json:"ChargeRequestCnt,omitempty"`
	Service          *string `json:"Service,omitempty"`
	AppId            *int64  `json:"AppId,omitempty"`
	AppName          *string `json:"AppName,omitempty"`
}

type SimpleProductStatisticsResultTotal CommonProductStatisticsResultTotal
type SimpleProductStatisticsResultDetail CommonProductStatisticsResultDetail
