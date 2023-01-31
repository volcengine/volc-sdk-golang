package visual

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func (p *Visual) commonHandler(api string, form url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonJsonHandler(api string, jsonStr string, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, nil, jsonStr)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonQueryHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) DistortionFree(form url.Values) (*model.DistortionFreeResult, int, error) {
	resp := new(model.DistortionFreeResult)
	statusCode, err := p.commonHandler("DistortionFree", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) StretchRecovery(form url.Values) (*model.StretchRecoveryResult, int, error) {
	resp := new(model.StretchRecoveryResult)
	statusCode, err := p.commonHandler("StretchRecovery", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EmoticonEdit(form url.Values) (*model.EmoticonEditResult, int, error) {
	resp := new(model.EmoticonEditResult)
	statusCode, err := p.commonHandler("EmoticonEdit", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EyeClose2Open(form url.Values) (*model.EyeClose2OpenResult, int, error) {
	resp := new(model.EyeClose2OpenResult)
	statusCode, err := p.commonHandler("EyeClose2Open", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageScore(form url.Values) (*model.ImageScoreResult, int, error) {
	resp := new(model.ImageScoreResult)
	statusCode, err := p.commonHandler("ImageScore", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageFlow(form url.Values) (*model.ImageFlowResult, int, error) {
	resp := new(model.ImageFlowResult)
	statusCode, err := p.commonHandler("ImageFlow", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) PoemMaterial(form url.Values) (*model.PoemMaterialResult, int, error) {
	resp := new(model.PoemMaterialResult)
	statusCode, err := p.commonHandler("PoemMaterial", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarPlateDetection(form url.Values) (*model.CarPlateDetectionResult, int, error) {
	resp := new(model.CarPlateDetectionResult)
	statusCode, err := p.commonHandler("CarPlateDetection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarSegment(form url.Values) (*model.CarSegmentResult, int, error) {
	resp := new(model.CarSegmentResult)
	statusCode, err := p.commonHandler("CarSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarDetection(form url.Values) (*model.CarDetectionResult, int, error) {
	resp := new(model.CarDetectionResult)
	statusCode, err := p.commonHandler("CarDetection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) SkySegment(form url.Values) (*model.SkySegmentResult, int, error) {
	resp := new(model.SkySegmentResult)
	statusCode, err := p.commonHandler("SkySegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) BankCard(form url.Values) (*model.BankCardResult, int, error) {
	resp := new(model.BankCardResult)
	statusCode, err := p.commonHandler("BankCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) BankCardV2(form url.Values) (*model.BankCardResultV2, int, error) {
	resp := new(model.BankCardResultV2)
	statusCode, err := p.commonHandler("BankCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) IDCard(form url.Values) (*model.IDCardResult, int, error) {
	resp := new(model.IDCardResult)
	statusCode, err := p.commonHandler("IDCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) IDCardV2(form url.Values) (*model.IDCardResultV2, int, error) {
	resp := new(model.IDCardResultV2)
	statusCode, err := p.commonHandler("IDCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OCRNormal(form url.Values) (*model.OCRNormalResult, int, error) {
	resp := new(model.OCRNormalResult)
	statusCode, err := p.commonHandler("OCRNormal", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OCRApi(form url.Values, action string) (*model.OcrApiResult, int, error) {
	resp := new(model.OcrApiResult)
	statusCode, err := p.commonHandler(action, form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) DrivingLicense(form url.Values) (*model.DrivingLicenseResult, int, error) {
	resp := new(model.DrivingLicenseResult)
	statusCode, err := p.commonHandler("DrivingLicense", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VehicleLicense(form url.Values) (*model.VehicleLicenseResult, int, error) {
	resp := new(model.VehicleLicenseResult)
	statusCode, err := p.commonHandler("VehicleLicense", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) Taibao(form url.Values) (*model.TaibaoResult, int, error) {
	resp := new(model.TaibaoResult)
	statusCode, err := p.commonHandler("OcrTaibao", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceSwap(form url.Values) (*model.FaceSwapResult, int, error) {
	resp := new(model.FaceSwapResult)
	statusCode, err := p.commonHandler("FaceSwap", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoon(form url.Values) (*model.JPCartoonResult, int, error) {
	resp := new(model.JPCartoonResult)
	statusCode, err := p.commonHandler("JPCartoon", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoonCut(form url.Values) (*model.JPCartoonCutResult, int, error) {
	resp := new(model.JPCartoonCutResult)
	statusCode, err := p.commonHandler("JPCartoonCut", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSceneDetect(form url.Values) (*model.VideoSceneDetectResult, int, error) {
	resp := new(model.VideoSceneDetectResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OverResolution(form url.Values) (*model.OverResolutionResult, int, error) {
	resp := new(model.OverResolutionResult)
	statusCode, err := p.commonHandler("OverResolution", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OverResolutionV2(imageBase64 []string) (*model.OverResolutionV2Result, int, error) {
	//处理入参
	jsonBody := &struct {
		ReqKey           string   `json:"req_key"`
		BinaryDataBase64 []string `json:"binary_data_base64"`
	}{
		ReqKey:           "lens_vida_nnsr", //算法名称，取固定值为lens_vida_nnsr
		BinaryDataBase64: imageBase64,      //图片文件，base64编码。此算法可选输入1张图片或多张图片
	}

	jsonStr, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, 500, errors.New("request json marshal error")
	}
	resp := new(model.OverResolutionV2Result)
	statusCode, err := p.commonJsonHandler("OverResolutionV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsSegment(form url.Values) (*model.GoodsSegmentResult, int, error) {
	resp := new(model.GoodsSegmentResult)
	statusCode, err := p.commonHandler("GoodsSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageOutpaint(form url.Values) (*model.ImageOutpaintResult, int, error) {
	resp := new(model.ImageOutpaintResult)
	statusCode, err := p.commonHandler("ImageOutpaint", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageInpaint(form url.Values) (*model.ImageInpaintResult, int, error) {
	resp := new(model.ImageInpaintResult)
	statusCode, err := p.commonHandler("ImageInpaint", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageCut(form url.Values) (*model.ImageCutResult, int, error) {
	resp := new(model.ImageCutResult)
	statusCode, err := p.commonHandler("ImageCut", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EntityDetect(form url.Values) (*model.EntityDetectResult, int, error) {
	resp := new(model.EntityDetectResult)
	statusCode, err := p.commonHandler("EntityDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsDetect(form url.Values) (*model.GoodsDetectResult, int, error) {
	resp := new(model.GoodsDetectResult)
	statusCode, err := p.commonHandler("GoodsDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ConvertPhoto(form url.Values) (*model.ConvertPhotoResult, int, error) {
	resp := new(model.ConvertPhotoResult)
	statusCode, err := p.commonHandler("ConvertPhoto", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EnhancePhoto(form url.Values) (*model.EnhancePhotoResult, int, error) {
	resp := new(model.EnhancePhotoResult)
	statusCode, err := p.commonHandler("EnhancePhoto", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GeneralSegment(form url.Values) (*model.GeneralSegmentResult, int, error) {
	resp := new(model.GeneralSegmentResult)
	statusCode, err := p.commonHandler("GeneralSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) HumanSegment(form url.Values) (*model.HumanSegmentResult, int, error) {
	resp := new(model.HumanSegmentResult)
	statusCode, err := p.commonHandler("HumanSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoInpaintSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoInpaintSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoInpaintQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoInpaintQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoRetargetingSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoRetargetingSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoRetargetingQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoRetargetingQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoCoverSelection(form url.Values) (*model.VideoCoverSelectResult, int, error) {
	resp := new(model.VideoCoverSelectResult)
	statusCode, err := p.commonHandler("VideoCoverSelection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSummarizationSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoSummarizationSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSummarizationQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoSummarizationQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoOverResolutionSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoOverResolutionQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
