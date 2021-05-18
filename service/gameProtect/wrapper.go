package gameProtect

import (
	"fmt"
)

func (p *GameProtector) RiskResult(req *RiskResultRequest) (*RiskResultResponse, error) {
	respBody, _, err := p.Client.Query("RiskResult", req.ToQuery())
	if err == nil {
		result := new(RiskResultResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}

	// 支持错误重试
	if p.Retry() {
		respBody, _, err = p.Client.Query("RiskResult", req.ToQuery())
		if err != nil {
			return nil, fmt.Errorf("RiskResult: fail to do request, %v", err)
		}
		result := new(RiskResultResponse)
		if err := UnmarshalResultInto(respBody, result); err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, fmt.Errorf("RiskResult: fail to do request, %v", err)
}
