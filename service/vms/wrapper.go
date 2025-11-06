package vms

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

func (p *Vms) CreateNumberApplication(req *CreateNumberApplicationRequest) (*CreateNumberApplicationResponse, int, error) {
	resp := new(CreateNumberApplicationResponse)
	if statusCode, err := HandlerJson("CreateNumberApplication", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAXB(req *BindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) SelectNumberAndBindAXB(req *SelectNumberAndBindAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("SelectNumberAndBindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UnbindAXB(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UnbindAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QuerySubscription(req *SpecificSubIdRequest) (*QuerySubscriptionResponse, int, error) {
	resp := new(QuerySubscriptionResponse)
	if statusCode, err := p.Handler("QuerySubscription", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QuerySubscriptionForList(req *QuerySubscriptionForListRequest) (*QuerySubscriptionForListResponse, int, error) {
	resp := new(QuerySubscriptionForListResponse)
	if statusCode, err := p.Handler("QuerySubscriptionForList", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpgradeAXToAXB(req *UpgradeAXToAXBRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("UpgradeAXToAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAXB(req *UpdateAXBRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UpdateAXB", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAXN(req *BindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) SelectNumberAndBindAXN(req *SelectNumberAndBindAXNRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("SelectNumberAndBindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAXN(req *UpdateAXNRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UpdateAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UnbindAXN(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UnbindAXN", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAXNE(req *BindAXNERequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAXBForAXNE(req *BindAXBForAXNERequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindAXBForAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UnbindAXNE(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UnbindAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAXNE(req *UpdateAXNERequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UpdateAXNE", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) RouteAAuth(req *RouteAAuthRequest) (*RouteAAuthResponse, int, error) {
	resp := new(RouteAAuthResponse)
	if statusCode, err := p.Handler("RouteAAuth", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryAuth(req *AuthQueryRequest) (*AuthQueryResponse, int, error) {
	resp := new(AuthQueryResponse)
	if statusCode, err := p.Handler("QueryAuth", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) Click2Call(req *Click2CallRequest) (*Click2CallResponse, int, error) {
	resp := new(Click2CallResponse)
	if statusCode, err := p.Handler("Click2Call", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) CancelClick2Call(req *CancelClick2CallRequest) (*CancelClick2CallResponse, int, error) {
	resp := new(CancelClick2CallResponse)
	if statusCode, err := p.Handler("CancelClick2Call", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) Click2CallLite(req *Click2CallLiteRequest) (*Click2CallLiteResponse, int, error) {
	resp := new(Click2CallLiteResponse)
	if statusCode, err := p.Handler("Click2CallLite", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAxyb(req *BindAxybRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindAxyb", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindYbForAxyb(req *BindYbForAxybRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.Handler("BindYbForAxyb", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UnbindAxyb(req *SpecificSubIdRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UnbindAxyb", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAxyb(req *UpdateAxybRequest) (*OperationResponse, int, error) {
	resp := new(OperationResponse)
	if statusCode, err := p.Handler("UpdateAxyb", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) CreateNumberPool(req *CreateNumberPoolRequest) (*CreateNumberPoolResponse, int, error) {
	resp := new(CreateNumberPoolResponse)
	if statusCode, err := Handler("CreateNumberPool", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateNumberPool(req *UpdateNumberPoolRequest) (*UpdateNumberPoolResponse, int, error) {
	resp := new(UpdateNumberPoolResponse)
	if statusCode, err := Handler("UpdateNumberPool", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) NumberPoolList(req *NumberPoolListRequest) (*NumberPoolListResponse, int, error) {
	resp := new(NumberPoolListResponse)
	if statusCode, err := Handler("NumberPoolList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) NumberList(req *NumberListRequest) (*NumberListResponse, int, error) {
	resp := new(NumberListResponse)
	if statusCode, err := Handler("NumberList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) EnableOrDisableNumber(req *EnableOrDisableNumberRequest) (*EnableOrDisableNumberResponse, int, error) {
	resp := new(EnableOrDisableNumberResponse)
	if statusCode, err := Handler("EnableOrDisableNumber", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryNumberApplyRecordList(req *QueryNumberApplyRecordListRequest) (*QueryNumberApplyRecordListResponse, int, error) {
	resp := new(QueryNumberApplyRecordListResponse)
	if statusCode, err := Handler("QueryNumberApplyRecordList", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryAudioRecordFileUrl(req *QueryAudioRecordFileUrlRequest) (*QueryAudioRecordFileUrlResponse, int, error) {
	resp := new(QueryAudioRecordFileUrlResponse)
	if statusCode, err := p.Handler("QueryAudioRecordFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryAudioRecordToTextFileUrl(req *QueryAudioRecordToTextFileRequest) (*QueryAudioRecordToTextFileResponse, int, error) {
	resp := new(QueryAudioRecordToTextFileResponse)
	if statusCode, err := p.Handler("QueryAudioRecordToTextFileUrl", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryCallRecordMsg(req *QueryCallRecordMsgRequest) (*QueryCallRecordMsgResponse, int, error) {
	resp := new(QueryCallRecordMsgResponse)
	if statusCode, err := p.Handler("QueryCallRecordMsg", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryCallRecordMsgNew(req *QueryCallRecordMsgRequest) (*QueryCallRecordMsgNewResponse, int, error) {
	resp := new(QueryCallRecordMsgNewResponse)
	if statusCode, err := p.Handler("QueryCallRecordMsgNew", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) CreateTask(req *CreateTaskRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("CreateTask", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) BatchAppend(req *BatchAppendRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("BatchAppend", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) UpdateTask(req *EditTaskRequest) (*TaskAppendResponse, int, error) {
	resp := new(TaskAppendResponse)
	statusCode, err := p.doJson("UpdateTask", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) PauseTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("PauseTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) ResumeTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("ResumeTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) StopTask(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("TaskOpenId", taskOpenId)

	statusCode, err := p.doQuery("StopTask", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) SingleInfo(taskOpenId string) (*SingleInfoResponse, int, error) {
	resp := new(SingleInfoResponse)
	query := url.Values{}
	query.Add("SingleOpenId", taskOpenId)

	statusCode, err := p.doQuery("SingleInfo", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) SingleBatchAppend(req *SingleAppendRequest) (*SingleAppendResponse, int, error) {
	resp := new(SingleAppendResponse)
	statusCode, err := p.doJson("SingleBatchAppend", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) SingleCancel(taskOpenId string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("SingleOpenId", taskOpenId)

	statusCode, err := p.doQuery("SingleCancel", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) FetchVoiceResourceByUrl(req *FetchVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	req.ReqSourceType=3
	statusCode, err := p.doJson("FetchResource", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) CreateTtsResource(req *CreateTtsResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	req.ReqSourceType=3
	statusCode, err := p.doJson("OpenCreateTts", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) DeleteResourceByResourceKey(resourceKey string) (*CommonResponse, int, error) {
	resp := new(CommonResponse)
	query := url.Values{}
	query.Add("ResourceKey", resourceKey)

	statusCode, err := p.doQuery("OpenDeleteResource", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) GenerateVoiceResourceUploadUrl(req *UploadVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	req.ReqSourceType=3
	statusCode, err := p.doJson("GetResourceUploadUrl", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) CommitVoiceResourceUpload(req *UploadVoiceResourceRequest) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	req.ReqSourceType=3
	statusCode, err := p.doJson("CommitResourceUpload", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) OpenUpdateResource(resourceKey string, name string) (*BasicResourceResponse, int, error) {
	resp := new(BasicResourceResponse)
	query := url.Values{}
	query.Add("ResourceKey", resourceKey)
	query.Add("Name", name)

	statusCode, err := p.doQuery("OpenUpdateResource", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) QueryUsableResource(resourceType string) (*QueryUsableResourceResponse, int, error) {
	resp := new(QueryUsableResourceResponse)
	query := url.Values{}
	query.Add("Type", resourceType)

	statusCode, err := p.doQuery("QueryUsableResource", query, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) QueryOpenGetResource(req *QueryOpenGetResourceRequest) (*QueryOpenGetResourceResponse, int, error) {
	resp := new(QueryOpenGetResourceResponse)
	statusCode, err := p.doJson("QueryOpenGetResource", req, resp)

	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Vms) AddQualification(req *AddQualificationRequest) (*AddQualificationResponse, int, error) {
	resp := new(AddQualificationResponse)
	req.QualificationMainInfoFormDO.ReqSourceType=3
	if statusCode, err := HandlerJson("AddQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateQualification(req *UpdateQualificationRequest) (*UpdateQualificationResponse, int, error) {
	resp := new(UpdateQualificationResponse)
	req.QualificationMainInfoFormDO.ReqSourceType=3
	if statusCode, err := HandlerJson("UpdateQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) AddQualificationScene(req *AddQualificationSceneRequest) (*AddQualificationSceneResponse, int, error) {
	resp := new(AddQualificationSceneResponse)
	if statusCode, err := HandlerJson("AddQualificationScene", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateQualificationScene(req *UpdateQualificationSceneRequest) (*UpdateQualificationSceneResponse, int, error) {
	resp := new(UpdateQualificationSceneResponse)
	if statusCode, err := HandlerJson("UpdateQualificationScene", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryQualification(req *QueryQualificationRequest) (*QueryQualificationResponse, int, error) {
	resp := new(QueryQualificationResponse)
	if statusCode, err := HandlerJson("QueryQualification", req, resp, *p.Client); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) QueryCanCall(req *RiskControlReq) (*RiskControlResponse, int, error) {
	resp := new(RiskControlResponse)
	if statusCode, err := p.Handler("QueryCanCall", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) AddBlackList(req *AddBlackListReq) (*AddBlackListResponse, int, error) {
	resp := new(AddBlackListResponse)
	if statusCode, err := p.Handler("AddBlackList", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) DeleteBlackList(req *DeleteBlackListReq) (*DeleteBlackListResponse, int, error) {
	resp := new(DeleteBlackListResponse)
	if statusCode, err := p.Handler("DeleteBlackList", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) BindAXG(req *BindAXGRequest) (*SecretBindResponse, int, error) {
	resp := new(SecretBindResponse)
	if statusCode, err := p.doJson("BindAXG", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAXG(req *UpdateAXGRequest) (*OperateResponse, int, error) {
	resp := new(OperateResponse)
	if statusCode, err := p.doJson("UpdateAXG", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UnbindAXG(req *UnbindAXGRequest) (*OperateResponse, int, error) {
	resp := new(OperateResponse)
	if statusCode, err := p.doJson("UnbindAXG", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) CreateAXGGroup(req *CreateAXGGroupRequest) (*CreateAXGGroupResponse, int, error) {
	resp := new(CreateAXGGroupResponse)
	if statusCode, err := p.doJson("CreateAXGGroup", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) UpdateAXGGroup(req *UpdateAXGGroupRequest) (*OperateResponse, int, error) {
	resp := new(OperateResponse)
	if statusCode, err := p.doJson("UpdateAXGGroup", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) DeleteAXGGroup(req *DeleteAXGGroupRequest) (*OperateResponse, int, error) {
	resp := new(OperateResponse)
	if statusCode, err := p.doJson("DeleteAXGGroup", req, resp); err != nil {
		return nil, statusCode, err
	} else {
		return resp, statusCode, nil
	}
}

func (p *Vms) Handler(api string, req interface{}, resp interface{}) (int, error) {
	form := base.ToUrlValues(req)
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Client.Post(api, nil, form)
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func Handler(api string, req interface{}, resp interface{}, p base.Client) (int, error) {
	form := base.ToUrlValues(req)
	apiInfo := p.ApiInfoList[api]
	var respBody []byte
	var statusCode int
	var err error
	if http.MethodGet == apiInfo.Method {
		respBody, statusCode, err = p.Query(api, form)
	} else {
		respBody, statusCode, err = p.Post(api, nil, form)
	}
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		if http.MethodGet == apiInfo.Method {
			respBody, statusCode, err = p.Query(api, form)
		} else {
			respBody, statusCode, err = p.Post(api, nil, form)
		}
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func HandlerJson(api string, req interface{}, resp interface{}, p base.Client) (int, error) {
	jsonBody, _ := json.Marshal(req)
	respBody, statusCode, err := p.Json(api, nil, string(jsonBody))
	if err != nil {
		return statusCode, err
	}
	if statusCode >= 500 {
		respBody, statusCode, err = p.Json(api, nil, string(jsonBody))
		if err != nil {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Vms) doJson(api string, req interface{}, resp interface{}) (int, error) {
	reqJ, _ := json.Marshal(req)
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

func (p *Vms) doQuery(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
