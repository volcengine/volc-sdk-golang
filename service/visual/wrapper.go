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
