package sami

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (p *Sami) commonHandlerJson(api string, query url.Values, resp interface{}, body string) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, query, body)
	if err != nil {
		return statusCode, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Sami) commonSamiHandler(urlPath string, resp interface{}, body string) (int, error) {
	req, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return 500, fmt.Errorf("fail to new request %s, %v", urlPath, err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return rsp.StatusCode, fmt.Errorf("fail to do request to %s, %v", urlPath, err)
	}

	ret, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 500, fmt.Errorf("fail to read response body %s, %v", urlPath, err)
	}
	if rsp.StatusCode != http.StatusOK {
		return 500, fmt.Errorf("http status=%v, body=%s, url=%s", rsp.StatusCode, string(ret), urlPath)
	}
	defer rsp.Body.Close()

	if err = json.Unmarshal(ret, &resp); err != nil {
		return 500, fmt.Errorf("parse response failed ,ret=%s,err=%v", urlPath, err)
	}
	return http.StatusOK, nil
}

func (p *Sami) GetToken(req GetTokenReq) (*GetTokenResponse, int, error) {
	resp := new(GetTokenResponse)
	body, _ := json.Marshal(&req)
	statusCode, err := p.commonHandlerJson("GetToken", nil, &resp, string(body))
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Sami) TTS(req SpeechSynthesisReq, appkey string, token string) (*InvokeResponse, int, error) {
	resp := new(InvokeResponse)
	payload, _ := json.Marshal(&req)
	payloads := string(payload)
	iReq := InvokeRequest{
		Payload: &payloads,
	}
	iReqBytes, _ := json.Marshal(&iReq)

	urlPath := fmt.Sprintf(
		"%v/api/v1/invoke?version=%v&token=%v&appkey=%v&namespace=%v",
		SamiDomain, "v4", token, appkey, "TTS",
	)

	statusCode, err := p.commonSamiHandler(urlPath, resp, string(iReqBytes))
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
