package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
)

type images struct {
	m *MaaS
}

// POST method
// Tokenization
func (i *images) ImagesQuickGen(endpointId string, req *api.ImagesQuickGenReq) (*api.ImagesQuickGenResp, int, error) {
	return i.ImagesQuickGenWithCtx(context.Background(), endpointId, req)
}

func (i *images) ImagesQuickGenWithCtx(ctx context.Context, endpointId string, req *api.ImagesQuickGenReq) (*api.ImagesQuickGenResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return i.imagesQuickGenImpl(ctx, endpointId, bts)
}

func (i *images) imagesQuickGenImpl(ctx context.Context, endpointId string, body []byte) (*api.ImagesQuickGenResp, int, error) {
	respBody, status, err, logId := i.m.request(ctx, maas.APIImagesQuickGen, nil, endpointId, body)
	if err != nil {
		errVal := &api.ImagesQuickGenResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), logId)
		}
		errVal.Error.ReqId = errVal.ReqId
		return nil, status, errVal.Error
	}

	output := new(api.ImagesQuickGenResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), logId)
	}
	return output, status, nil
}
