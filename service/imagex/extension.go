package imagex

import (
	"fmt"
	"net/url"
)

func (c *ImageX) GetImageEnhance(param *GetImageEnhanceParam) (*GetImageEnhanceResult, error) {
	res := new(GetImageEnhanceResult)
	err := c.ImageXPost("GetImageEnhanceResult", nil, param, res)
	if err != nil {
		return nil, err
	}
	return res, nil
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

func (c *ImageX) GetImageOCRLicense(param *GetImageOCRParam) (*GetImageOCRLicenseResult, error) {
	u := url.Values{}
	u.Set("Scene", "license")
	u.Set("ServiceId", param.ServiceId)
	u.Set("StoreUri", param.StoreUri)
	data, _, err := c.Post("GetImageOCR", u, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("fail to request api GetImageOCR, %v", err)
	}
	result := new(GetImageOCRLicenseResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (c *ImageX) GetImageOCRGeneral(param *GetImageOCRParam) (*GetImageOCRGeneralResult, error) {
	u := url.Values{}
	u.Set("Scene", "general")
	u.Set("ServiceId", param.ServiceId)
	u.Set("StoreUri", param.StoreUri)
	data, _, err := c.Post("GetImageOCR", u, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("fail to request api GetImageOCR, %v", err)
	}
	result := new(GetImageOCRGeneralResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
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
