package businessSecurity

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	hi_sse "github.com/volcengine/volc-sdk-golang/service/businessSecurity/sse"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api"
	"io"
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

// Deprecated: use ElementVerifyV2 instead
// 已废弃，请使用ElementVerifyV2
// Deprecated
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

func (p *BusinessSecurity) AsyncImageRiskV2(req *AsyncRiskDetectionRequest) (*AsyncVideoRiskResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncImageRiskV2: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncImageRiskV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncImageRiskV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncImageRiskV2: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncImageRiskV2: fail to do request, %v", err)
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

func (p *BusinessSecurity) GetImageResultV2(req *VideoResultRequest) (*ImageResultResponse, error) {
	respBody, _, err := p.Client.Query("GetImageResultV2", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetImageResultV2", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetImageResultV2: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResultV2: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// image risk deteciton
// 已废弃-使用ImageContentRiskV2
// Deprecated
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

// image risk deteciton
// 内容安全图片实时接口
func (p *BusinessSecurity) ImageContentRiskV2(req *RiskDetectionRequest) (*ImageResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ImageContentRiskV2: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ImageContentRiskV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ImageContentRiskV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ImageContentRiskV2: fail to do request, %v", err)
			}
			result := new(ImageResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ImageContentRiskV2: fail to do request, %v", err)
	}
	result := new(ImageResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// 要素验证
func (p *BusinessSecurity) ElementVerifyV2(req *ElementVerifyRequest) (*ElementVerifyResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerifyV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerifyV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ImageContentRisk: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ElementVerifyEncrypted 加密要素验证
// encryptedType - 加密类型，例如：AES
// secretKey - 秘钥(需要申请)
func (p *BusinessSecurity) ElementVerifyEncrypted(encryptedType string, secretKey string, req *ElementVerifyRequest) (*ElementVerifyResponseV2, error) {
	if encryptedType == "" || secretKey == "" {
		return nil, fmt.Errorf("ElementVerifyEncrypted: encryptedType or secretKey empty")
	}
	if req == nil {
		return nil, fmt.Errorf("ElementVerifyEncrypted: requset is nil")
	}
	req.EncryptedType = encryptedType
	parameters, err := AesCBCEncryptWithBase64(secretKey, req.Parameters)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyEncrypted: fail encrypt parameters")
	}
	req.Parameters = parameters

	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyEncryptedRequest: fail to marshal request, %v", err)
	}
	respBody, _, err := p.Client.Json("ElementVerifyEncrypted", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerifyEncrypted", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerifyEncrypted: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerifyEncrypted: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// text risk detection
// 已废弃-使用TextSliceRisk
// Deprecated
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

// text slice risk detection
// 内容安全文本切片检测实时接口
func (p *BusinessSecurity) TextSliceRisk(req *RiskDetectionRequest) (*TextSliceResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("TextSliceRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("TextSliceRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("TextSliceRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
			}
			result := new(TextSliceResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
	}
	result := new(TextSliceResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// text slice risk detection with context
// 内容安全文本切片检测实时接口
func (p *BusinessSecurity) TextSliceRiskWithCtx(ctx context.Context, req *RiskDetectionRequest) (*TextSliceResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("TextSliceRisk: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.CtxJson(ctx, "TextSliceRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.CtxJson(ctx, "TextSliceRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
			}
			result := new(TextSliceResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("TextSliceRisk: fail to do request, %v", err)
	}
	result := new(TextSliceResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Deprecated: use MobileStatusV2 instead
// 已废弃，请使用MobileStatusV2
// Deprecated
func (p *BusinessSecurity) MobileStatus(req *MobileStatusRequest) (*MobileStatusResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
			}
			result := new(MobileStatusResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	result := new(MobileStatusResponse)
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

// 号码状态
func (p *BusinessSecurity) MobileStatusV2(req *MobileStatusRequest) (*MobileStatusResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatusV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatusV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatusV2: fail to do request, %v", err)
			}
			result := new(MobileStatusResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
	}
	result := new(MobileStatusResponseV2)
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

// audio Risk
// 风险识别接口
func (p *BusinessSecurity) AudioRisk(req *RiskDetectionRequest) (*AudioResultResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AudioRisk: fail to marshal request, %v", err)
	}
	respBody, _, err := p.Client.Json("AudioRisk", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AudioRisk", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AudioRisk: fail to do request, %v", err)
			}
			result := new(AudioResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AudioRisk: fail to do request, %v", err)
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
func (p *BusinessSecurity) GetVideoLiveResult(req *VideoResultRequest) (*VideoLiveResultResponse, error) {
	respBody, _, err := p.Client.Query("GetVideoLiveResul", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetVideoLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetVideoLiveResult: fail to do request, %v", err)
			}
			result := new(VideoLiveResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageResult: fail to do request, %v", err)
	}
	result := new(VideoLiveResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// close audio Risk
// 风险识别异步关闭接口
func (p *BusinessSecurity) CloseVideoLiveRisk(req *VideoResultRequest) (*AsyncVideoRiskResponse, error) {
	respBody, _, err := p.Client.Query("CloseVideoLiveRisk", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("CloseVideoLiveRisk", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("CloseVideoLiveRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CloseVideoLiveRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
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
func (p *BusinessSecurity) GetAudioLiveResult(req *VideoResultRequest) (*AudioLiveResultResponse, error) {
	respBody, _, err := p.Client.Query("GetAudioLiveResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("GetAudioLiveResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
			}
			result := new(AudioLiveResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAudioLiveResult: fail to do request, %v", err)
	}
	result := new(AudioLiveResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// close audio Risk
// 风险识别异步关闭接口
func (p *BusinessSecurity) CloseAudioLiveRisk(req *VideoResultRequest) (*AsyncVideoRiskResponse, error) {
	respBody, _, err := p.Client.Query("CloseAudioLiveRisk", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("CloseAudioLiveRisk", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("CloseAudioLiveRisk: fail to do request, %v", err)
			}
			result := new(AsyncVideoRiskResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CloseAudioLiveRisk: fail to do request, %v", err)
	}
	result := new(AsyncVideoRiskResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApp
// 创建app接口
func (p *BusinessSecurity) CreateApp(req *CreateAppReq) (*CreateAppResponse, error) {

	respBody, _, err := p.Client.Json("CreateApp", req.ToQuery(), "")
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("CreateApp", req.ToQuery(), "")
			if err != nil {
				return nil, fmt.Errorf("CreateApp: fail to do request, %v", err)
			}
			result := new(CreateAppResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncLiveAudioRisk: fail to do request, %v", err)
	}
	result := new(CreateAppResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListApps
// 获取APP列表
func (p *BusinessSecurity) ListApps(req *ListAppsReq) (*ListAppsResponse, error) {
	respBody, _, err := p.Client.Query("ListApps", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("ListApps", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("ListApps: fail to do request, %v", err)
			}
			result := new(ListAppsResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
		}
		return nil, fmt.Errorf("ListApps: fail to do request, %v", err)
	}
	result := new(ListAppsResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCustomContents 创建自定义库
func (p *BusinessSecurity) CreateCustomContents(req *NewCustomContentsReq) (*CustomContentResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("CreateCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateCustomContents: fail to do request, %v", err)
			}
			result := new(CustomContentResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateCustomContents: fail to do request, %v", err)
	}
	result := new(CustomContentResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// EnableCustomContents 启用文本自定义库
func (p *BusinessSecurity) EnableCustomContents(req *UpdateContentReq) (*CustomContentResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("EnableCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("EnableCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("EnableCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("EnableCustomContents: fail to do request, %v", err)
			}
			result := new(CustomContentResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("EnableCustomContents: fail to do request, %v", err)
	}
	result := new(CustomContentResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// disable custom contents
// 禁用文本自定义库
func (p *BusinessSecurity) DisableCustomContents(req *UpdateContentReq) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DisableCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DisableCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DisableCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// delete custom contents
// 删除文本自定义库内容
func (p *BusinessSecurity) DeleteCustomContents(req *ModifyTextContent) (*CustomContentResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DeleteCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteCustomContents: fail to do request, %v", err)
			}
			result := new(CustomContentResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DisableCustomContents: fail to do request, %v", err)
	}
	result := new(CustomContentResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// upload custom contents
// 上传文本自定义库内容
func (p *BusinessSecurity) UploadCustomContents(req *ModifyTextContent) (*CustomContentResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UploadCustomContents: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UploadCustomContents", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UploadCustomContents", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UploadCustomContents: fail to do request, %v", err)
			}
			result := new(CustomContentResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UploadCustomContents: fail to do request, %v", err)
	}
	result := new(CustomContentResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func handleSimpleRiskStatResp(respBody []byte) (*SimpleRiskStatResponse, error) {
	resultTmp := new(SimpleRiskStatResponse)
	if err := json.Unmarshal(respBody, resultTmp); err != nil {
		return nil, err
	}
	return resultTmp, nil
}

func (p *BusinessSecurity) SimpleRiskStat(req *CommonProductStatisticsReq) (*SimpleRiskStatResponse, error) {
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

func handleContentRiskStatResp(respBody []byte) (*ContentRiskStatResponse, error) {
	resultTmp := new(ContentRiskStatResponse)
	if err := json.Unmarshal(respBody, resultTmp); err != nil {
		return nil, err
	}
	return resultTmp, nil
}

func (p *BusinessSecurity) ContentRiskStat(req *CommonProductStatisticsReq) (*ContentRiskStatResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ContentRiskStat: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ContentRiskStat", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("ContentRiskStat", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ContentRiskStat: fail to do request, %v", err)
			}
			return handleContentRiskStatResp(respBody)
		}
		return nil, fmt.Errorf("ContentRiskStat: fail to do request, %v", err)
	}
	return handleContentRiskStatResp(respBody)
}

// query system name list item
// 查询系统名单内容
func (p *BusinessSecurity) QuerySystemNameListItem(req *QuerySystemNameListItemReq) (*QuerySystemNameListItemResp, error) {
	respBody, _, err := p.Client.Query("QuerySystemNameListItem", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("QuerySystemNameListItem", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("QuerySystemNameListItem: fail to do request, %v", err)
			}
			result := new(QuerySystemNameListItemResp)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("QuerySystemNameListItem: fail to do request, %v", err)
	}
	result := new(QuerySystemNameListItemResp)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// delete system name list item
// 删除系统名单内容，进行解封
func (p *BusinessSecurity) DelSystemNameListItem(req *DelSystemNameListItemReq) (*DelSystemNameListItemResp, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DelSystemNameListItem: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DelSystemNameListItem", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DelSystemNameListItem", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DelSystemNameListItem: fail to do request, %v", err)
			}
			result := new(DelSystemNameListItemResp)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DelSystemNameListItem: fail to do request, %v", err)
	}
	result := new(DelSystemNameListItemResp)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *SecuritySecurityClient) SecuritySource(req *RiskDetectionRequest) (*SecuritySourceResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("SecuritySource: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("SecuritySource", nil, string(reqData))
	if err != nil {
		return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
	}
	result := new(SecuritySourceResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *SecuritySecurityClient) SecuritySourceStream(req *RiskDetectionRequest) (<-chan *SecuritySourceResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("SecuritySource: fail to marshal request, %v", err)
	}

	apiInfo := p.ApiInfoList["SecuritySourceStream"]
	if apiInfo == nil {
		return nil, api.NewClientSDKRequestError("the related api does not exist")
	}

	r, err := makeRequest(apiInfo, p.ServiceInfo, nil, "application/json")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err))
	}

	r.Body = io.NopCloser(bytes.NewReader(reqData))
	timeout := getTimeout(p.ServiceInfo.Timeout, apiInfo.Timeout)

	r = p.ServiceInfo.Credentials.Sign(r)

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	r = r.WithContext(ctx)

	// do request
	resp, err := p.Client.Client.Do(r)
	if err != nil {
		cancel()
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("request error: %v", err))
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("error:%+v", resp.Body)
		cancel()
		_ = resp.Body.Close()
		return nil, err
	}

	requestId := resp.Header.Get("X-Tt-Logid")
	ch := make(chan *SecuritySourceResponse, respBufferSize)

	go func() {
		defer func() {
			_ = recover()
			_ = resp.Body.Close()
			cancel()
			close(ch)
		}()
		stream := hi_sse.NewEventStreamFromReader(resp.Body, maxBufferSize)

		for {
			event, err := stream.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				//if errors.Is(err, context.DeadlineExceeded) {
				//}
				ch <- &SecuritySourceResponse{
					RequestId: requestId,
					Code:      2000,
					Message:   "internal error",
				}
				return
			}
			if event != nil {
				if bytes.Equal(event.Data, []byte(terminator)) {
					return
				}
			}

			item := &SecuritySourceResponse{}
			if err = UnmarshalResultInto(event.Data, item); err != nil {
				//if err = json.Unmarshal(event.Data, item); err != nil {
				ch <- &SecuritySourceResponse{
					RequestId: requestId,
					Code:      2000,
					Message:   "internal error",
				}
				return
			}
			ch <- item

		}
	}()

	return ch, nil
}

func (p *BusinessSecurity) CreateCustomLib(req *CreateCustomLibRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomLib: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateCustomLib", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("CreateCustomLib", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateCustomLib: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateCustomLib: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) UpdateCustomLib(req *UpdateCustomLibRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateCustomLib: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdateCustomLib", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UpdateCustomLib", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdateCustomLib: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdateCustomLib: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) ChangeCustomContentsStatus(req *ChangeCustomLibStatusRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ChangeCustomContentsStatus: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ChangeCustomContentsStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("ChangeCustomContentsStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ChangeCustomContentsStatus: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ChangeCustomContentsStatus: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DeleteCustomLib(req *DeleteCustomLibRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteCustomLib: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteCustomLib", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DeleteCustomLib", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteCustomLib: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteCustomLib: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) GetCustomLib(req *GetCustomLibRequest) (*CustomLibListResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetCustomLib: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetCustomLib", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("GetCustomLib", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetCustomLib: fail to do request, %v", err)
			}
			result := new(CustomLibListResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetCustomLib: fail to do request, %v", err)
	}
	result := new(CustomLibListResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) CreateAccessConfig(req *CreateAccessConfigRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateAccessConfig: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreateAccessConfig", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("CreateAccessConfig", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreateAccessConfig: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreateAccessConfig: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) UpdateAccessConfig(req *UpdateAccessConfigRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccessConfig: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdateAccessConfig", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UpdateAccessConfig", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdateAccessConfig: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdateAccessConfig: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) UpdateConfigStatus(req *UpdateAccessConfigStatusRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateConfigStatus: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdateConfigStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UpdateConfigStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdateConfigStatus: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdateConfigStatus: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) GetAccessConfig(req *GetAccessConfigStatusRequest) (*AccessConfigListResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetAccessConfig: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetAccessConfig", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("GetAccessConfig", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetAccessConfig: fail to do request, %v", err)
			}
			result := new(AccessConfigListResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetAccessConfig: fail to do request, %v", err)
	}
	result := new(AccessConfigListResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) GetTextLibContent(req *GetCustomTextLibRequest) (*CustomTextLibListResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTextLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetTextLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("GetTextLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetTextLibContent: fail to do request, %v", err)
			}
			result := new(CustomTextLibListResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetTextLibContent: fail to do request, %v", err)
	}
	result := new(CustomTextLibListResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DeleteTextLibContent(req *DeleteCustomTextRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteTextLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteTextLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DeleteTextLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteTextLibContent: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteTextLibContent: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) UploadTextLibContent(req *UploadCustomTextRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UploadTextLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UploadTextLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UploadTextLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UploadTextLibContent: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UploadTextLibContent: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) GetImageLibContent(req *GetCustomImgLibRequest) (*CustomImgLibListResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetImageLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetImageLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("GetImageLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetImageLibContent: fail to do request, %v", err)
			}
			result := new(CustomImgLibListResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetImageLibContent: fail to do request, %v", err)
	}
	result := new(CustomImgLibListResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DeleteImageLibContent(req *DeleteCustomImgRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeleteImageLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeleteImageLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("DeleteImageLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeleteImageLibContent: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeleteImageLibContent: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) UploadImageLibContent(req *UploadCustomImgRequest) (*CommonResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UploadImageLibContent: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UploadImageLibContent", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err := p.Client.Json("UploadImageLibContent", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UploadImageLibContent: fail to do request, %v", err)
			}
			result := new(CommonResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UploadImageLibContent: fail to do request, %v", err)
	}
	result := new(CommonResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
