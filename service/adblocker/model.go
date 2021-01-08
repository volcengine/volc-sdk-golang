package adblocker

import (
	"encoding/json"
	"fmt"
	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

type AdBlockRequest struct {
	AppID      int64  `json:"app_id"`
	Service    string `json:"service"`
	Parameters string `json:"parameters"`
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
