package iam

import (
	"encoding/json"
	"net/url"
)

// helper func
func (p *IAM) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *IAM) ListUsers(query url.Values) (*UserListResp, int, error) {
	resp := new(UserListResp)
	statusCode, err := p.commonHandler("ListUsers", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
