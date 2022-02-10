package businessSecurity

import (
	"encoding/json"
	"fmt"
)

// Synchronous risk detection
// 风险识别实时接口
func (p *BusinessSecurity) RiskDetection(req *RiskDetectionRequest) (*RiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
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
	result := new(RiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous risk detection
// 风险识别异步接口
func (p *BusinessSecurity) AsyncRiskDetection(req *AsyncRiskDetectionRequest) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncRiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncRiskDetection", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) RiskResult(req *RiskResultRequest) (*RiskResultResponse, error) {
	respBody, _, err := p.Client.Query("RiskResult", req.ToQuery())
	if err != nil {
		// Retry on error
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
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(RiskResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DataReport(req *DataReportRequest) (*DataReportResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DataReport", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("DataReport", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(DataReportResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DataReport: fail to do request, %v", err)
	}
	result := new(DataReportResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) ElementVerify(req *ElementVerifyRequest) (*ElementVerifyResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerify", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerify", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileSecondSale(req *MobileSecondSaleRequest) (*MobileSecondSaleResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileSecondSale", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileSecondSale", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileSecondSale: fail to do request, %v", err)
			}
			result := new(MobileSecondSaleResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileSecondSale: fail to do request, %v", err)
	}
	result := new(MobileSecondSaleResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileEmptyCheck(req *MobileEmptyCheckRequest) (*MobileEmptyCheckResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileEmptyCheckRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileEmptyCheck", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileEmptyCheck", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileEmptyCheck: fail to do request, %v", err)
			}
			result := new(MobileEmptyCheckResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileEmptyCheck: fail to do request, %v", err)
	}
	result := new(MobileEmptyCheckResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileOnlineStatus(req *MobileOnlineStatusRequest) (*MobileOnlineStatusResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileOnlineStatusRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileOnlineStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileOnlineStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileOnlineStatus: fail to do request, %v", err)
			}
			result := new(MobileOnlineStatusResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileOnlineStatus: fail to do request, %v", err)
	}
	result := new(MobileOnlineStatusResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileOnlineTime(req *MobileOnlineTimeRequest) (*MobileOnlineTimeResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileOnlineTimeRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileOnlineTime", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileOnlineTime", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileOnlineTime: fail to do request, %v", err)
			}
			result := new(MobileOnlineTimeResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileOnlineTime: fail to do request, %v", err)
	}
	result := new(MobileOnlineTimeResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
