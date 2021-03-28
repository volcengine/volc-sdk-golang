package businessSecurity

import (
	"encoding/json"
	"fmt"
)

func (p *BusinessSecurity) RiskDetection(req *RiskDetectionRequest) (*RiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RiskDetection", nil, string(reqData))
	if err == nil {
		result := new(RiskDetectionResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}

	// 支持错误重试
	if p.Retry() {
		respBody, _, err = p.Client.Json("RiskDetection", nil, string(reqData))
		if err != nil {
			return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
		}
		result := new(RiskDetectionResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
}
