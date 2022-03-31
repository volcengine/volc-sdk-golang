package notify

import (
	"encoding/json"
	"net/url"
)

func (p *Notify) doJson(api string, req interface{}, resp interface{}) (int, error) {
	reqJ, err := json.Marshal(req)
	respBody, statusCode, err := p.Client.Json(api, nil, string(reqJ))
	if err != nil {
		return statusCode, err
	}

	//log.Print(string(respBody))
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Notify) doQuery(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	//log.Print(string(respBody))
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Notify) CreateTask(req *CreateTaskRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("CreateTask", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) BatchAppend(req *BatchAppendRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("BatchAppend", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) UpdateTask(req *EditTaskRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("UpdateTask", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) PauseTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("PauseTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) ResumeTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("ResumeTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) StopTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("StopTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) SingleInfo(taskOpenId string) (*SingleInfoResponse, int, error) {
	resp := new(SingleInfoResponse)
	query := url.Values{}
	query.Add("SingleOpenId", taskOpenId)

	statusCode, err := p.doQuery("SingleInfo", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) SingleBatchAppend(req *SingleAppendRequest) (*SingleAppendResponse, int, error) {
	resp := new(SingleAppendResponse)
	statusCode, err := p.doJson("SingleBatchAppend", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) SingleCancel(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("SingleOpenId", taskOpenId)

	statusCode, err := p.doQuery("SingleCancel", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) FetchVoiceResourceByUrl(req *FetchVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	statusCode, err := p.doJson("FetchResource", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) CreateTtsResource(req *FetchVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	statusCode, err := p.doJson("OpenCreateTts", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) DeleteResourceByResourceKey(resourceKey string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("ResourceKey", resourceKey)

	statusCode, err := p.doQuery("OpenDeleteResource", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) GenerateVoiceResourceUploadUrl(req *UploadVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	statusCode, err := p.doJson("GetResourceUploadUrl", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Notify) CommitVoiceResourceUpload(req *UploadVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	statusCode, err := p.doJson("CommitResourceUpload", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}