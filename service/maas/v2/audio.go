package v2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
)

type audio struct {
	Speech
}

type Speech struct {
	m *MaaS
}

func (s *Speech) Create(endpointId string, speechReq *api.SpeechReq) (maas.ResponseContent, int, error) {
	return s.CreateWithCtx(context.Background(), endpointId, speechReq)
}

func (s *Speech) CreateWithCtx(ctx context.Context, endpointId string, req *api.SpeechReq) (maas.ResponseContent, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return s.createImpl(ctx, endpointId, bts)
}

func (s *Speech) createImpl(ctx context.Context, endpointId string, body []byte) (maas.ResponseContent, int, error) {
	ctx = getContext(ctx)
	apikey := s.m.settedApikey
	respBody, status, err, cancel := s.m.streamRequest(ctx, maas.APIAudioSpeech, nil, endpointId, body, apikey)
	if err != nil {
		return nil, status, err
	}

	return maas.NewBinaryResponseContent(respBody, reqIdFromCtx(ctx), cancel), status, nil
}
