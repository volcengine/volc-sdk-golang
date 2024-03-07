package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	"io"
)

type audio struct {
	Speech
}

type Speech struct {
	m *MaaS
}

func (s *Speech) Create(endpointId string, speechReq *api.SpeechReq) (maas.ResponseContent, error) {
	apiInfo := s.m.ApiInfoList[maas.APIAudioSpeech]
	if apiInfo == nil {
		return nil, api.NewClientSDKRequestError("the related api does not exist", "")
	}

	body, err := json.Marshal(speechReq)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}

	// build request
	req, err := maas.MakeRequest(apiInfo, endpointId, s.m.ServiceInfo, nil, "application/json")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err), "")
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	req = s.m.ServiceInfo.Credentials.Sign(req)

	// do request
	resp, err := s.m.Client.Client.Do(req)
	logId := resp.Header.Get("x-tt-logid")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("request error: %v", err), logId)
	}

	if resp.StatusCode != 200 { // fast fail
		res := &api.ChatResp{}
		if er := json.NewDecoder(resp.Body).Decode(res); er != nil {
			res.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to call service: http status_code=%d", resp.StatusCode), logId)
		}
		_ = resp.Body.Close()
		return nil, res.Error
	}

	return maas.NewBinaryResponseContent(resp.Body, logId), nil
}
