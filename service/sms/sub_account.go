package sms

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

func (p *SMS) GetSubAccountList(req *GetSubAccountListRequest) (*GetSubAccountListResponse, int, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	resp := new(GetSubAccountListResponse)
	statusCode, err := p.smsHandlerWithQuery("GetSubAccountList", req, v, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) GetSubAccountDetail(req *GetSubAccountDetailRequest) (*GetSubAccountDetailResponse, int, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	resp := new(GetSubAccountDetailResponse)
	statusCode, err := p.smsHandlerWithQuery("GetSubAccountDetail", req, v, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *SMS) InsertSmsSubAccount(req *InsertSmsSubAccountReq) (*InsertSmsSubAccountResponse, int, error) {
	resp := new(InsertSmsSubAccountResponse)
	statusCode, err := p.smsHandler("InsertSubAccount", req, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
