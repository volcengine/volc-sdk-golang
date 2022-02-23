package codePipeline

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	. "github.com/volcengine/volc-sdk-golang/service/codePipeline/models"
)

func (p *CodePipeline) CreatePipeline(req *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreatePipelineRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("CreatePipeline", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("CreatePipeline", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("CreatePipeline: fail to do request, %v", err)
			}
			result := new(CreatePipelineResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("CreatePipeline: fail to do request, %v", err)
	}
	result := new(CreatePipelineResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) ListPipelines(req *ListPipelinesRequest) (*ListPipelinesResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListPipelinesRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ListPipelines", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ListPipelines", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ListPipelines: fail to do request, %v", err)
			}
			result := new(ListPipelinesResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ListPipelines: fail to do request, %v", err)
	}
	result := new(ListPipelinesResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) DeletePipeline(req *DeletePipelineRequest) (*DeletePipelineResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeletePipelineRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeletePipeline", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("DeletePipeline", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeletePipeline: fail to do request, %v", err)
			}
			result := new(DeletePipelineResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeletePipeline: fail to do request, %v", err)
	}
	result := new(DeletePipelineResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) UpdatePipeline(req *UpdatePipelineRequest) (*UpdatePipelineResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdatePipelineRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdatePipeline", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("UpdatePipeline", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdatePipeline: fail to do request, %v", err)
			}
			result := new(UpdatePipelineResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdatePipeline: fail to do request, %v", err)
	}
	result := new(UpdatePipelineResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) GetPipeline(req *GetPipelineRequest) (*GetPipelineResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetPipelineRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetPipeline", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetPipeline", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetPipeline: fail to do request, %v", err)
			}
			result := new(GetPipelineResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetPipeline: fail to do request, %v", err)
	}
	result := new(GetPipelineResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) UpdatePipelineProperties(req *UpdatePipelinePropertiesRequest) (*UpdatePipelinePropertiesResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdatePipelinePropertiesRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("UpdatePipelineProperties", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("UpdatePipelineProperties", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("UpdatePipelineProperties: fail to do request, %v", err)
			}
			result := new(UpdatePipelinePropertiesResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("UpdatePipelineProperties: fail to do request, %v", err)
	}
	result := new(UpdatePipelinePropertiesResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) RunPipeline(req *RunPipelineRequest) (*RunPipelineResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RunPipelineRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RunPipeline", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("RunPipeline", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("RunPipeline: fail to do request, %v", err)
			}
			result := new(RunPipelineResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("RunPipeline: fail to do request, %v", err)
	}
	result := new(RunPipelineResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) ListPipelineRecords(req *ListPipelineRecordsRequest) (*ListPipelineRecordsResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListPipelineRecordsRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ListPipelineRecords", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("ListPipelineRecords", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ListPipelineRecords: fail to do request, %v", err)
			}
			result := new(ListPipelineRecordsResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ListPipelineRecords: fail to do request, %v", err)
	}
	result := new(ListPipelineRecordsResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) GetPipelineRecord(req *GetPipelineRecordRequest) (*GetPipelineRecordResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetPipelineRecordRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("GetPipelineRecord", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("GetPipelineRecord", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("GetPipelineRecord: fail to do request, %v", err)
			}
			result := new(GetPipelineRecordResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("GetPipelineRecord: fail to do request, %v", err)
	}
	result := new(GetPipelineRecordResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) StopPipelineRecord(req *StopPipelineRecordRequest) (*StopPipelineRecordResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("StopPipelineRecordRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("StopPipelineRecord", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("StopPipelineRecord", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("StopPipelineRecord: fail to do request, %v", err)
			}
			result := new(StopPipelineRecordResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("StopPipelineRecord: fail to do request, %v", err)
	}
	result := new(StopPipelineRecordResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) RetryPipelineRecord(req *RetryPipelineRecordRequest) (*RetryPipelineRecordResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RetryPipelineRecordRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RetryPipelineRecord", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("RetryPipelineRecord", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("RetryPipelineRecord: fail to do request, %v", err)
			}
			result := new(RetryPipelineRecordResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("RetryPipelineRecord: fail to do request, %v", err)
	}
	result := new(RetryPipelineRecordResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *CodePipeline) DeletePipelineRecord(req *DeletePipelineRecordRequest) (*DeletePipelineRecordResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("DeletePipelineRecordRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DeletePipelineRecord", nil, string(reqData))
	if err != nil {
		if p.Retry() {
			respBody, _, err = p.Client.Json("DeletePipelineRecord", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("DeletePipelineRecord: fail to do request, %v", err)
			}
			result := new(DeletePipelineRecordResponse)
			if err := base.UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DeletePipelineRecord: fail to do request, %v", err)
	}
	result := new(DeletePipelineRecordResponse)
	if err := base.UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
