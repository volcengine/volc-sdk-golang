package adblocker

import (
	"encoding/json"
	"fmt"
)

func (p *AdBlocker) AdBlock(req *AdBlockRequest) (*AdBlockResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AdBlockRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AdBlock", nil, string(reqData))
	if err == nil {
		result := new(AdBlockResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}

	// 支持错误重试
	if p.Retry() {
		respBody, _, err = p.Client.Json("AdBlock", nil, string(reqData))
		if err != nil {
			return nil, fmt.Errorf("AdBlock: fail to do request, %v", err)
		}
		result := new(AdBlockResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, fmt.Errorf("AdBlock: fail to do request, %v", err)
}

func handleSimpleRiskStatResp(respBody []byte) (*SimpleRiskStatResponse, error) {
	resultTmp := new(SimpleRiskStatResponse)
	if err := json.Unmarshal(respBody, resultTmp); err != nil {
		return nil, err
	}
	return resultTmp, nil
}

func (p *AdBlocker) SimpleRiskStat(req *CommonProductStatisticsReq) (*SimpleRiskStatResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("SimpleRiskStat: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("SimpleRiskStat", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("SimpleRiskStat", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("SimpleRiskStat: fail to do request, %v", err)
			}
			return handleSimpleRiskStatResp(respBody)
		}
		return nil, fmt.Errorf("SimpleRiskStat: fail to do request, %v", err)
	}
	return handleSimpleRiskStatResp(respBody)
}
