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

// Asynchronous video risk detection
// 内容安全视频异步接口
func (p *BusinessSecurity) AsyncVideoRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncVideoRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncVideoRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncVideoRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncVideoRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) VideoResult(req *VideoResultRequest) (*VideoResultResponse, error) {
	respBody, _, err := p.Client.Query("VideoResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("VideoResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("VideoResult: fail to do request, %v", err)
			}
			result := new(VideoResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("VideoResult: fail to do request, %v", err)
	}
	result := new(VideoResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous video risk detection
// 内容安全图片异步接口
func (p *BusinessSecurity) AsyncImageRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncImageRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncImageRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncImageRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetImageResult(req *VideoResultRequest) (*ImageResultResponse, error) {
	respBody, _, err := p.Client.Query("GetImageResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetImageResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// image risk deteciton
// 内容安全图片实时接口
func (p *BusinessSecurity) ImageContentRisk(req *RiskDetectionRequest) (*ImageResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ImageContentRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ImageContentRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ImageContentRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// text risk detection
// 内容安全文本实时接口
func (p *BusinessSecurity) TextRisk(req *RiskDetectionRequest) (*TextResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("TextRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("TextRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("TextRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("TextRisk: fail to do request, %v", err)
			}
			result := new(TextResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("TextRisk: fail to do request, %v", err)
	}
	result := new(TextResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous audio risk detection
// 内容安全音频异步接口
func (p *BusinessSecurity) AsyncAudioRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncAudioRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncAudioRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncAudioRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncAudioRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// audio Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetAudioResult(req *VideoResultRequest) (*AudioResultResponse, error) {
	respBody, _, err := p.Client.Query("GetAudioResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetAudioResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetAudioResult: fail to do request, %v", err)
			}
			result := new(AudioResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAudioResult: fail to do request, %v", err)
	}
	result := new(AudioResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous video risk detection
// 内容安全图片异步接口
func (p *BusinessSecurity) AsyncLiveVideoRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncLiveVideoRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncLiveVideoRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncLiveVideoRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetVideoLiveResult(req *VideoResultRequest) (*VideoResultResponse, error) {
	respBody, _, err := p.Client.Query("GetVideoLiveResul", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetVideoLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetVideoLiveResult: fail to do request, %v", err)
			}
			result := new(VideoResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
	}
	result := new(VideoResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous audio risk detection
// 内容安全音频异步接口
func (p *BusinessSecurity) AsyncLiveAudioRisk(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncLiveAudioRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncLiveAudioRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// audio Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) GetAudioLiveResult(req *VideoResultRequest) (*AudioResultResponse, error) {
	respBody, _, err := p.Client.Query("GetAudioLiveResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetAudioLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
			}
			result := new(AudioResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
	}
	result := new(AudioResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}