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

func (p *SMS) ApplySmsSignatureV2(req *ApplySmsSignatureRequestV2) (*ApplySmsSignatureResponse, int, error) {
	resp := new(ApplySmsSignatureResponse)
	statusCode, err := p.smsHandler("ApplySmsSignatureV2", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) UpdateSmsSignature(req *UpdateSmsSignatureRequestV2) (*UpdateSmsSignatureResponse, int, error) {
	resp := new(UpdateSmsSignatureResponse)
	statusCode, err := p.smsHandler("UpdateSmsSignature", req, resp)
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

func (p *SMS) ApplySignatureIdent(req *ApplySignatureIdentRequest) (*ApplySignatureIdentResponse, int, error) {
	resp := new(ApplySignatureIdentResponse)
	statusCode, err := p.smsHandler("ApplySignatureIdent", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetSignatureIdentList(req *GetSignatureIdentListRequest) (*GetSignatureIdentListResponse, int, error) {
	resp := new(GetSignatureIdentListResponse)
	statusCode, err := p.smsHandler("GetSignatureIdentList", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) BatchBindSignatureIdent(req *BatchBindSignatureIdentRequest) (*BatchBindSignatureIdentResponse, int, error) {
	resp := new(BatchBindSignatureIdentResponse)
	statusCode, err := p.smsHandler("BatchBindSignatureIdent", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) UpdateSignatureIdent(req *UpdateSignatureIdentificationRequest) (*UpdateSignatureIdentificationResponse, int, error) {
	resp := new(UpdateSignatureIdentificationResponse)
	statusCode, err := p.smsHandler("UpdateSignatureIdentification", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
