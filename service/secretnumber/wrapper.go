package secretnumber

import (
	"encoding/json"
	"github.com/volcengine/volc-sdk-golang/base"
)

func (p *SecretNumber) BindAXB(req *BindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("BindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) SelectNumberAndBindAXB(req *SelectNumberAndBindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("SelectNumberAndBindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UnbindAXB(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UnbindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) QuerySubscription(req *SpecificSubIdRequest) (*QuerySubscriptionResponse, int, error) {
	resp := new(QuerySubscriptionResponse)
	if statusCode, err := p.handler("QuerySubscription", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) QuerySubscriptionForList(req *QuerySubscriptionForListRequest) (*QuerySubscriptionForListResponse, int, error) {
	resp := new(QuerySubscriptionForListResponse)
	if statusCode, err := p.handler("QuerySubscriptionForList", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpgradeAXToAXB(req *UpgradeAXToAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("UpgradeAXToAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpdateAXB(req *UpdateAXBRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UpdateAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) BindAXN(req *BindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("BindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) SelectNumberAndBindAXN(req *SelectNumberAndBindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.handler("SelectNumberAndBindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UpdateAXN(req *UpdateAXNRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UpdateAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) UnbindAXN(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.handler("UnbindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) Click2Call(req *Click2CallRequest) (*Click2CallResponse, int, error) {
	resp := new(Click2CallResponse)
	if statusCode, err := p.handler("Click2Call", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) Click2CallLite(req *Click2CallLiteRequest) (*Click2CallLiteResponse, int, error) {
	resp := new(Click2CallLiteResponse)
	if statusCode, err := p.handler("Click2CallLite", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryAudioRecordFileUrl(req *QueryAudioRecordFileUrlRequest) (*QueryAudioRecordFileUrlResponse, int, error) {
	resp := new(QueryAudioRecordFileUrlResponse)
	if statusCode, err := p.handler("QueryAudioRecordFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryAudioRecordToTextFileUrl(req *QueryAudioRecordToTextFileRequest) (*QueryAudioRecordToTextFileResponse, int, error) {
	resp := new(QueryAudioRecordToTextFileResponse)
	if statusCode, err := p.handler("QueryAudioRecordToTextFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *DataCenter) QueryCallRecordMsg(req *QueryCallRecordMsgRequest) (*QueryCallRecordMsgResponse, int, error) {
	resp := new(QueryCallRecordMsgResponse)
	if statusCode, err := p.handler("QueryCallRecordMsg", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *SecretNumber) handler(api string, req interface{}, resp interface{}) (int, error) {
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

func (p *DataCenter) handler(api string, req interface{}, resp interface{}) (int, error) {
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
