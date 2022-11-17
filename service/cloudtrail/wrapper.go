package cloudtrail

import (
	"encoding/json"
	"net/url"
)

func (p *CloudTrail) LookupEvents(req *LookupEventsReq) (*LookupEventsResp, int, error) {
	query := url.Values{}
	resp := new(LookupEventsResp)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}
	statusCode, err := p.commonHandlerJson("LookupEvents", query, reqBody, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *CloudTrail) commonHandlerJson(api string, query url.Values, reqBody []byte, resp interface{}) (statusCode int, err error) {
	var respBody []byte

	respBody, statusCode, err = p.Client.Json(api, query, string(reqBody))
	if err != nil {
		return
	}

	if err = json.Unmarshal(respBody, resp); err != nil {
		return
	}
	return
}
