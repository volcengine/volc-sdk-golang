package sms

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

func (p *SMS) GetSignatureAndOrderList(req *GetSignatureAndOrderListRequest) (*GetSignatureAndOrderListResponse, int, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	resp := new(GetSignatureAndOrderListResponse)
	statusCode, err := p.smsHandlerWithQuery("GetSignatureAndOrderList", nil, v, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) ApplySmsSignature(req *ApplySmsSignatureRequest) (*ApplySmsSignatureResponse, int, error) {
	resp := new(ApplySmsSignatureResponse)
	statusCode, err := p.smsHandler("ApplySmsSignature", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) DeleteSignature(req *DeleteSignatureRequest) (*DeleteSignatureResponse, int, error) {
	resp := new(DeleteSignatureResponse)
	statusCode, err := p.smsHandler("DeleteSignature", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
