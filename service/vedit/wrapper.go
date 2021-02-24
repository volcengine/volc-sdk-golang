package vedit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (e *VEdit) SubmitDirectEditTaskAsync(request *SubmitDirectEditTaskRequest) (*SubmitDirectEditTaskAsyncResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := e.Json("SubmitDirectEditTaskAsync", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &SubmitDirectEditTaskAsyncResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "edit"
		return resp, nil
	}
}

func (e *VEdit) SubmitTemplateTaskAsync(request *SubmitTemplateTaskRequest) (*SubmitTemplateTaskAsyncResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := e.Json("SubmitTemplateTaskAsync", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &SubmitTemplateTaskAsyncResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "edit"
		return resp, nil
	}
}

func (e *VEdit) GetDirectEditResult(request *GetDirectEditResultRequest) (*GetDirectEditResultResponse, error) {
	bts, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	respBody, status, err := e.Json("GetDirectEditResult", nil, string(bts))
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("http error"), string(status))
	}

	resp := &GetDirectEditResultResponse{}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	} else {
		resp.ResponseMetadata.Service = "edit"
		return resp, nil
	}
}
