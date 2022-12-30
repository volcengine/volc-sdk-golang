package sms

import (
	"fmt"
)

func (p *SMS) ApplyVmsTemplate(req *ApplyVmsTemplateRequest) (*ApplyVmsTemplateResponse, int, error) {
	if len(req.Contents) == 0 {
		return nil, 0, fmt.Errorf("should contain contents")
	}
	containText := false
	for _, element := range req.Contents {
		if element.SourceType == SourceTypeText {
			containText = true
		}
	}
	if !containText {
		return nil, 0, fmt.Errorf("should contain text in contents")
	}
	resp := new(ApplyVmsTemplateResponse)
	if req.ChannelType == "" {
		req.ChannelType = "CN_VMS"
	}
	req.Caller = "sdk"
	statusCode, err := p.smsHandler("ApplyVmsTemplate", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetVmsTemplateStatus(req *GetVmsTemplateStatusRequest) (*GetVmsTemplateStatusResponse, int, error) {
	resp := new(GetVmsTemplateStatusResponse)
	statusCode, err := p.smsHandler("GetVmsTemplateStatus", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) SendVms(req *SendVmsRequest) (*SendVmsResponse, int, error) {
	resp := new(SendVmsResponse)
	statusCode, err := p.smsHandler("SendSms", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
