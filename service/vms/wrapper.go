package vms

import (
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/base"
)

func (p *Risk) QueryCanCall(req *RiskControlReq) (*RiskControlResponse, int, error) {
	resp := new(RiskControlResponse)
	if statusCode, err := p.Handler("QueryCanCall", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Risk) Handler(api string, req interface{}, resp interface{}) (int, error) {
	form := base.ToUrlValues(req)
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Client.Post(api, nil, form)
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
