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

func (p *SMS) ApplySmsTemplateV2(req *ApplyTemplateV2Request) (*ApplyTemplateV2Response, int, error) {
	resp := new(ApplyTemplateV2Response)
	statusCode, err := p.smsHandler("ApplySmsTemplateV2", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ApplySmsSubContentTemplateV2(req *CreateSubContentTemplate) (*ApplySmsSubContentTemplateV2Response, int, error) {
	resp := new(ApplySmsSubContentTemplateV2Response)
	statusCode, err := p.smsHandler("ApplySmsSubContentTemplateV2", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ListSmsTemplateV2(req *ListSmsTemplateV2Request) (*ListSmsTemplateV2Response, int, error) {
	resp := new(ListSmsTemplateV2Response)
	statusCode, err := p.smsHandler("ListSmsTemplateV2", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) BindSignatures(req *BindSignatureRequest) (*BindSignatureResponse, int, error) {
	resp := new(BindSignatureResponse)
	statusCode, err := p.smsHandler("BindSignatures", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ListSecondTemplate(req *ListSecondTemplateRequest) (*ListSecondTemplateResponse, int, error) {
	resp := new(ListSecondTemplateResponse)
	statusCode, err := p.smsHandler("ListSecondTemplate", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ListSubContent(req *ListSubContentRequest) (*ListSubContentResponse, int, error) {
	resp := new(ListSubContentResponse)
	statusCode, err := p.smsHandler("ListSubContent", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
