package sms

import (
	"encoding/json"
)

func (p *SMS) Send(req *SmsRequest) (*SmsResponse, int, error) {
	resp := new(SmsResponse)
	statusCode, err := p.smsHandler("SendSms", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) smsHandler(api string, req interface{}, resp interface{}) (int, error) {
	reqJ, err := json.Marshal(req)
	respBody, statusCode, err := p.Client.Json(api, nil, string(reqJ))
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
