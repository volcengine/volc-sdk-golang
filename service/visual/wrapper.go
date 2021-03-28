package visual

import (
	"encoding/json"
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

func (p *Visual) BankCard(form url.Values) (*model.BankCardResult, int, error) {
	resp := new(model.BankCardResult)
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

func (p *Visual) OCRNormal(form url.Values) (*model.OCRNormalResult, int, error) {
	resp := new(model.OCRNormalResult)
	statusCode, err := p.commonHandler("OCRNormal", form, resp)
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
