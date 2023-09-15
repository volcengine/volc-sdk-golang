package imagex

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

func (c *ImageX) GetImageEnhance(param *GetImageEnhanceParam) (*GetImageEnhanceResult, error) {
	query := url.Values{}
	query.Set("ServiceId", param.ServiceId)
	res := new(GetImageEnhanceResult)
	err := c.ImageXPost("GetImageEnhanceResult", query, param, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ImageX) GetImageEnhanceWithData(param *GetImageEnhanceParam, data io.Reader) ([]byte, error) {
	query := url.Values{}
	query.Set("ServiceId", param.ServiceId)
	paramStr, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	fieldItem := base.CreateMultiPartItemFormField("Input", string(paramStr))
	fileItem := base.CreateMultiPartItemFormFile("Data", "img", data)
	resp, _, err := c.CtxMultiPart(context.Background(), "GetImageEnhanceResultWithData", query, []*base.MultiPartItem{fieldItem, fileItem})
	if err != nil {
		return nil, err
	}
	result := ""
	if err := UnmarshalResultInto(resp, &result); err != nil {
		return nil, err
	}
	return base64.RawStdEncoding.DecodeString(result)
}

func (c *ImageX) GetImageEraseModel(eraseType int) ([]string, error) {
	res := new(struct {
		Models []string `json:"Models"`
	})
	query := make(url.Values)
	query.Add("Type", fmt.Sprintf("%d", eraseType))
	err := c.ImageXGet("GetImageEraseModels", query, res)
	if err != nil {
		return nil, err
	}
	return res.Models, nil
}

func (c *ImageX) GetImageErase(param *GetImageEraseParam) (*GetImageEraseResult, error) {
	res := new(GetImageEraseResult)
	err := c.ImageXPost("GetImageEraseResult", nil, param, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ImageX) GetImageSegment(param *GetImageSegmentParam) (*GetImageSegmentResult, error) {
	query := url.Values{}
	query.Set("ServiceId", param.ServiceId)
	res := new(GetImageSegmentResult)
	err := c.ImageXPost("GetSegmentImage", query, param, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ImageX) GetImageQuality(param *GetImageQualityParam) (*GetImageQualityResult, error) {
	u := url.Values{}
	u.Add("ServiceId", param.ServiceId)
	resp := new(GetImageQualityResult)
	err := c.ImageXPost("GetImageQuality", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageSuperResolution(param *GetImageSuperResolutionParam) (*GetImageSuperResolutionResp, error) {
	result := new(GetImageSuperResolutionResp)
	err := c.ImageXPost("GetImageSuperResolutionResult", nil, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageStyleResult(req *GetImageStyleResultReq) (*GetImageStyleResultResp, error) {
	query := url.Values{}
	query.Add("ServiceId", req.ServiceId)
	resp := new(GetImageStyleResultResp)
	err := c.ImageXPost("GetImageStyleResult", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageOCRV2(req *GetImageOCRParam) (*GetImageOCRResult, error) {
	query := url.Values{}
	query.Set("Scene", req.Scene)
	query.Set("ServiceId", req.ServiceId)
	query.Set("StoreUri", req.StoreUri)
	resp := new(GetImageOCRResult)
	err := c.ImageXPost("GetImageStyleResult", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageBgFill(param *GetImageBgFillParam) (*GetImageBgFillResult, error) {
	res := new(GetImageBgFillResult)
	err := c.ImageXPost("GetImageBgFillResult", nil, param, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *ImageX) GetImageDuplicateDetection(param *GetImageDuplicateDetectionParam) (*GetImageDuplicateDetectionResult, error) {
	arg := getImageDuplicateDetectionParam{
		Urls:     param.Urls,
		Async:    false,
		Callback: "",
	}
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetImageDuplicateDetectionResult)
	err := c.ImageXPost("GetImageDuplicateDetection", query, arg, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageDuplicateDetectionAsync(param *GetImageDuplicateDetectionAsyncParam) (*GetImageDuplicateDetectionAsyncResult, error) {
	arg := getImageDuplicateDetectionParam{
		Urls:     param.Urls,
		Async:    true,
		Callback: param.Callback,
	}
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetImageDuplicateDetectionAsyncResult)
	err := c.ImageXPost("GetImageDuplicateDetection", query, arg, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetDedupTaskStatus(taskId string) (*GetDedupTaskStatusResult, error) {
	query := make(url.Values)
	query.Add("TaskId", taskId)

	result := new(GetDedupTaskStatusResult)
	err := c.ImageXGet("GetDedupTaskStatus", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetDenoisingImage(param *GetDenoisingImageParam) (*GetDenoisingImageResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetDenoisingImageResult)
	err := c.ImageXPost("GetDenoisingImage", query, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageComic(param *GetImageComicParam) (*GetImageComicResult, error) {
	result := new(GetImageComicResult)
	err := c.ImageXPost("GetImageComicResult", nil, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageSmartCrop(param *GetImageSmartCropParam) (*GetImageSmartCropResult, error) {
	result := new(GetImageSmartCropResult)
	err := c.ImageXPost("GetImageSmartCropResult", nil, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetLicensePlateDetection(param *GetLicensePlateDetectionParam) (*GetLicensePlateDetectionResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetLicensePlateDetectionResult)
	err := c.ImageXPost("GetLicensePlateDetection", query, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImagePSDetection(param *GetImagePSDetectionParam) (*GetImagePSDetectionResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetImagePSDetectionResult)
	err := c.ImageXPost("GetImagePSDetection", query, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetPrivateImageType(param *GetPrivateImageTypeParam) (*GetPrivateImageTypeResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)

	result := new(GetPrivateImageTypeResult)
	err := c.ImageXPost("GetPrivateImageType", query, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) CreateImageHmEmbed(param *CreateImageHmEmbedParam) (*CreateImageHmEmbedResult, error) {
	result := new(CreateImageHmEmbedResult)
	err := c.ImageXPost("CreateImageHmEmbed", nil, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) CreateImageHmExtract(param *CreateImageHmExtractParam) (*CreateImageHmExtractResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)
	query.Add("StoreUri", param.StoreUri)
	query.Add("Algorithm", param.Algorithm)

	result := new(CreateImageHmExtractResult)
	err := c.ImageXGet("CreateImageHmExtract", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageAuditResult(req *GetImageAuditResultReq) (*GetImageAuditResultResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}
	resp := &GetImageAuditResultResp{}
	err = c.ImageXGet("GetImageAuditResult", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageAuditTasks(req *GetImageAuditTasksReq) (*GetImageAuditTasksResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}
	resp := &GetImageAuditTasksResp{}
	err = c.ImageXGet("GetImageAuditTasks", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetAuditEntrysCount(req *GetAuditEntrysCountReq) (*GetAuditEntrysCountResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}
	resp := &GetAuditEntrysCountResp{}
	err = c.ImageXGet("GetAuditEntrysCount", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) UpdateImageAuditTaskStatus(req *UpdateImageAuditTaskStatusReq) (*UpdateImageAuditTaskStatusResp, error) {
	result := new(UpdateImageAuditTaskStatusResp)
	err := c.ImageXPost("UpdateImageAuditTaskStatus", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) CreateImageRetryAuditTask(req *CreateImageRetryAuditTaskReq) (*CreateImageRetryAuditTaskResp, error) {
	result := new(CreateImageRetryAuditTaskResp)
	err := c.ImageXPost("CreateImageRetryAuditTask", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) DeleteImageAuditResult(req *DeleteImageAuditResultReq) (*DeleteImageAuditResultResp, error) {
	result := new(DeleteImageAuditResultResp)
	err := c.ImageXPost("DeleteImageAuditResult", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) UpdateAuditImageStatus(req *UpdateAuditImageStatusReq) (*UpdateAuditImageStatusResp, error) {
	result := new(UpdateAuditImageStatusResp)
	err := c.ImageXPost("UpdateAuditImageStatus", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) CreateImageAuditTask(req *CreateImageAuditTaskReq) (*CreateImageAuditTaskResp, error) {
	result := new(CreateImageAuditTaskResp)
	err := c.ImageXPost("CreateImageAuditTask", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) UpdateImageAuditTask(req *UpdateImageAuditTaskReq) (*UpdateImageAuditTaskResp, error) {
	result := new(UpdateImageAuditTaskResp)
	err := c.ImageXPost("UpdateImageAuditTask", nil, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
