package sms

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

func (p *SMS) GetSmsTemplateAndOrderList(req *GetSmsTemplateAndOrderListRequest) (*GetSmsTemplateAndOrderListResponse, int, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	resp := new(GetSmsTemplateAndOrderListResponse)
	statusCode, err := p.smsHandlerWithQuery("GetSmsTemplateAndOrderList", nil, v, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ApplySmsTemplate(req *ApplySmsTemplateRequest) (*ApplySmsTemplateResponse, int, error) {
	resp := new(ApplySmsTemplateResponse)
	statusCode, err := p.smsHandler("ApplySmsTemplate", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) DeleteSmsTemplate(req *DeleteSmsTemplateRequest) (*DeleteSmsTemplateResponse, int, error) {
	resp := new(DeleteSmsTemplateResponse)
	statusCode, err := p.smsHandler("DeleteSmsTemplate", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
