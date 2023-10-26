package aiot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (p *AIoT) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	for _, it := range p.queryInterceptors {
		query = it(api, query)
	}

	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		if string(respBody) != "" {
			if err := json.Unmarshal(respBody, resp); err != nil {
				return statusCode, err
			} else {
				return statusCode, nil
			}
		}
		return statusCode, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *AIoT) commonHandlerJson(api string, query url.Values, resp interface{}, req interface{}) (int, error) {
	for _, it := range p.queryInterceptors {
		query = it(api, query)
	}

	payload := ""
	if req != nil {
		body, _ := json.Marshal(req)
		payload = string(body)
	}

	respBody, statusCode, err := p.Client.Json(api, query, payload)
	if err != nil {
		if string(respBody) != "" {
			if err := json.Unmarshal(respBody, resp); err != nil {
				return statusCode, err
			} else {
				return statusCode, nil
			}
		}
		return statusCode, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

//space相关接口
func (p *AIoT) CreateSpace(request *CreateSpaceRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateSpace", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteSpace(request *SpaceRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("DeleteSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartSpace(request *SpaceRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("StartSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopSpace(request *SpaceRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("StopSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateSpace(request *UpdateSpaceRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("UpdateSpace", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetSpace(request *SpaceRequest) (*GetSpaceResponse, int, error) {
	resp := new(GetSpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandler("GetSpace", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListSpaces(request *ListSpacesRequest) (*ListSpacesResponse, int, error) {
	resp := new(ListSpacesResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	query.Add("Order", request.Order)
	statusCode, err := p.commonHandler("ListSpaces", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetSpaceTemplate(request *SpaceRequest) (*GetSpaceTemplateResponse, int, error) {
	resp := new(GetSpaceTemplateResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandler("GetSpaceTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CheckBindTemplate(request *CheckBindTemplateRequest) (*ListSpacesResponse, int, error) {
	resp := new(ListSpacesResponse)
	query := url.Values{
		"TemplateType": []string{request.TemplateType},
		"TemplateID":   []string{request.TemplateID},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandler("CheckBindTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CancelBindTemplate(request *CancelBindTemplateRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID":      []string{request.SpaceID},
		"TemplateType": []string{request.TemplateType},
	}
	statusCode, err := p.commonHandler("CancelBindTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) SetSpaceTemplate(request *SetSpaceTemplateRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID":      []string{request.SpaceID},
		"TemplateType": []string{request.TemplateType},
	}
	statusCode, err := p.commonHandlerJson("SetSpaceTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) AddSpaceDomain(request *AddSpaceDomainRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("AddSpaceDomain", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListSpaceDomains(spaceID string) (*ListSpaceDomainsResponse, int, error) {
	resp := new(ListSpaceDomainsResponse)
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	statusCode, err := p.commonHandler("ListSpaceDomains", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteSpaceDomain(spaceID, domain string) (*DeleteSpaceDomainResponse, int, error) {
	resp := new(DeleteSpaceDomainResponse)
	query := url.Values{
		"SpaceID": []string{spaceID},
		"Domain":  []string{domain},
	}
	statusCode, err := p.commonHandler("DeleteSpaceDomain", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetSpaceDomain(spaceID, domain string) (*GetSpaceDomainResponse, int, error) {
	resp := new(GetSpaceDomainResponse)
	query := url.Values{
		"SpaceID": []string{spaceID},
		"Domain":  []string{domain},
	}
	statusCode, err := p.commonHandler("GetSpaceDomain", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateDomainHTTPS(request *UpdateDomainHTTPSRequest) (*SpaceResponse, int, error) {
	resp := new(SpaceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
		"Domain":  []string{request.Domain},
	}
	statusCode, err := p.commonHandlerJson("UpdateDomainHTTPS", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateSpaceDomain(request *UpdateSpaceDomainRequest) (*UpdateSpaceDomainResponse, int, error) {
	resp := new(UpdateSpaceDomainResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("UpdateSpaceDomain", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateAuthInSpace(request *UpdateAuthInSpaceRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("UpdateAuthInSpace", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DisableAuthInSpace(spaceID, domain string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID": []string{spaceID},
		"Domain":  []string{domain},
	}
	statusCode, err := p.commonHandlerJson("DisableAuthInSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetDataProjectWithBindWidthAndFlow(request *GetDataProjectWithBindWidthAndFlowRequest) (*GetDataProjectWithBindWidthAndFlowResponse, int, error) {
	resp := new(GetDataProjectWithBindWidthAndFlowResponse)
	query := url.Values{
		"SpaceID":    []string{request.SpaceID},
		"StreamName": []string{request.StreamName},
		"StartTime":  []string{request.StartTime},
		"EndTime":    []string{request.EndTime},
		"Data":       []string{request.Data},
	}
	statusCode, err := p.commonHandlerJson("GetDataProjectWithBindWidthAndFlow", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetTotalData(t string) (*GetTotalDataResponse, int, error) {
	resp := new(GetTotalDataResponse)
	query := url.Values{
		"Time": []string{t},
	}
	statusCode, err := p.commonHandlerJson("GetTotalData", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetPushStreamCnt(request *GetPushStreamCntRequest) (*GetPushStreamCntResponse, int, error) {
	resp := new(GetPushStreamCntResponse)
	query := url.Values{
		"SpaceID":   []string{request.SpaceID},
		"StartTime": []string{request.StartTime},
		"EndTime":   []string{request.EndTime},
	}
	statusCode, err := p.commonHandler("GetPushStreamCnt", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//device相关
func (p *AIoT) CreateDevice(request *CreateDeviceRequest) (*CreateDeviceResponse, int, error) {
	resp := new(CreateDeviceResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("CreateDevice", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteDevice(request *DeviceRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}
	statusCode, err := p.commonHandlerJson("DeleteDevice", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopDevice(request *DeviceRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}
	statusCode, err := p.commonHandlerJson("StopDevice", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartDevice(request *DeviceRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}
	statusCode, err := p.commonHandlerJson("StartDevice", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) FreshDevice(request *FreshDeviceRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
		"SpaceID":  []string{request.SpaceID},
	}
	statusCode, err := p.commonHandler("FreshDevice", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) LocalMediaDownload(request *LocalMediaDownloadRequest) (*LocalMediaDownloadResponse, int, error) {
	resp := new(LocalMediaDownloadResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("LocalMediaDownload", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetMediaDownload(request *GetLocalMediaDownloadRequest) (*ListGBMediaResponse, int, error) {
	resp := new(ListGBMediaResponse)
	statusCode, err := p.commonHandler("GetLocalDownload", url.Values{
		"SpaceID": []string{request.SpaceID},
		"ID":      []string{request.ID},
	}, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteLocalDownload(spaceID, id string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID": []string{spaceID},
		"ID":      []string{id},
	}
	statusCode, err := p.commonHandlerJson("DeleteLocalDownload", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GenSipID(request *GenSipIDRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceType":  []string{request.DeviceType},
		"SipServerID": []string{request.SipServerID},
	}
	statusCode, err := p.commonHandler("GenSipID", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetDevice(request *GetDeviceRequest) (*GetDeviceResponse, int, error) {
	resp := new(GetDeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}
	if request.SpaceID != "" {
		query.Add("SpaceID", request.SpaceID)
	}
	if request.SipServerID != "" {
		query.Add("SipServerID", request.SipServerID)
	}
	statusCode, err := p.commonHandler("GetDevice", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListDevices(request *ListDevicesRequest) (*ListDevicesResponse, int, error) {
	resp := new(ListDevicesResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	if request.DeviceID != "" {
		query.Add("DeviceID", request.DeviceID)
	}
	if request.DeviceNSID != "" {
		query.Add("DeviceNSID", request.DeviceNSID)
	}
	if request.DeviceName != "" {
		query.Add("DeviceName", request.DeviceName)
	}
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandler("ListDevices", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetDeviceChannels(request *DeviceRequest, options ...GetDeviceChannelsOption) (*GetDeviceChannelsResponse, int, error) {
	resp := new(GetDeviceChannelsResponse)
	opts := getDeviceChannelsOptions{}
	for _, option := range options {
		option(&opts)
	}

	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}

	if opts.Mode != "" {
		query.Add("QueryMode", string(opts.Mode))
	}

	statusCode, err := p.commonHandler("GetDeviceChannels", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateDevice(request *UpdateDeviceRequest) (*DeviceResponse, int, error) {
	resp := new(DeviceResponse)
	query := url.Values{
		"DeviceID": []string{request.DeviceID},
	}
	statusCode, err := p.commonHandlerJson("UpdateDevice", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CloudRecordPlay(request *CloudRecordPlayRequest) (*CloudRecordPlayResponse, int, error) {
	resp := new(CloudRecordPlayResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CloudRecordPlay", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListDeviceScreenshots(request *ListHistoryRequest) (*ListHistoryResponse, int, error) {
	resp := new(ListHistoryResponse)
	query := url.Values{
		"SpaceID":  []string{request.SpaceID},
		"StreamID": []string{request.StreamID},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandlerJson("ListDeviceScreenshots", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListDeviceRecords(request *ListHistoryRequest) (*ListHistoryResponse, int, error) {
	resp := new(ListHistoryResponse)
	query := url.Values{
		"SpaceID":  []string{request.SpaceID},
		"StreamID": []string{request.StreamID},
		"ReqType":  []string{request.ReqType},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandlerJson("ListDeviceRecords", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CloudControl(request *CloudControlRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SipID": []string{request.SipID},
	}
	statusCode, err := p.commonHandlerJson("CloudControl", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) QueryPresetInfo(request *QueryPresetInfoRequest) (*QueryPresetInfoResponse, int, error) {
	resp := new(QueryPresetInfoResponse)
	query := url.Values{
		"SipID": []string{request.SipID},
	}
	statusCode, err := p.commonHandlerJson("QueryPresetInfo", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetRecordList(request *GetRecordListRequest) (*GetRecordListResponse, int, error) {
	resp := new(GetRecordListResponse)
	statusCode, err := p.commonHandlerJson("GetRecordList", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) PlayBackStart(request *PlayBackStartRequest) (*PlayBackStartResponse, int, error) {
	resp := new(PlayBackStartResponse)
	statusCode, err := p.commonHandlerJson("PlayBackStart", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) PlayBackStop(request *StreamRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("PlayBackStop", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) PlayBackControl(request *PlayBackControlRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	statusCode, err := p.commonHandlerJson("PlayBackControl", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CruiseControl(request *CruiseControlRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CruiseControl", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) SetAlarmGuard(request *SetAlarmGuardRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SipID":      []string{request.SipID},
		"DeviceNSID": []string{request.DeviceNSID},
		"ChannelID":  []string{request.ChannelID},
		"Enable":     []string{request.Enable},
	}
	statusCode, err := p.commonHandlerJson("SetAlarmGuard", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ResetAlarm(request *ResetAlarmRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SipID":      []string{request.SipID},
		"DeviceNSID": []string{request.DeviceNSID},
		"ChannelID":  []string{request.ChannelID},
	}
	statusCode, err := p.commonHandlerJson("ResetAlarm", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListAlarmNotify(request *ListAlarmNotifyRequest) (*ListAlarmNotifyResponse, int, error) {
	resp := new(ListAlarmNotifyResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandlerJson("ListAlarmNotify", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteAlarmNotify(alarmNotifyID string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"AlarmNotifyID": []string{alarmNotifyID},
	}
	statusCode, err := p.commonHandlerJson("DeleteAlarmNotify", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteAlarmNotifyAll(spaceID, deviceID string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID":  []string{spaceID},
		"DeviceID": []string{deviceID},
	}
	statusCode, err := p.commonHandlerJson("DeleteAlarmNotifyAll", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) AiotPlayBackStart(request *PlayBackStartRequest) (*PlayBackStartResponse, int, error) {
	resp := new(PlayBackStartResponse)
	statusCode, err := p.commonHandlerJson("AiotPlayBackStart", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) AiotPlayBackStop(request *StreamRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"StreamName": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("AiotPlayBackStop", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) AiotPlayBackControl(request *AiotPlayBackControlRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	statusCode, err := p.commonHandlerJson("AiotPlayBackControl", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//stream相关

func (p *AIoT) CreateStream(request *CreateStreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("CreateStream", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateStream(request *UpdateStreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"SpaceID":  []string{request.SpaceID},
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("UpdateStream", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartStream(request *StreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("StartStream", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartStreamWithParameters(request *StartStreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("StartStream", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopStream(request *StreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("StopStream", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ForbidStream(request *StreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("ForbidStream", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UnforbidStream(request *StreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("UnforbidStream", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteStream(request *StreamRequest) (*StreamResponse, int, error) {
	resp := new(StreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandlerJson("DeleteStream", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStream(request *StreamRequest) (*GetStreamResponse, int, error) {
	resp := new(GetStreamResponse)
	query := url.Values{
		"StreamID": []string{request.StreamID},
	}
	statusCode, err := p.commonHandler("GetStream", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStreams(request *ListStreamsRequest) (*ListStreamsResponse, int, error) {
	resp := new(ListStreamsResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	if request.StreamName != "" {
		query.Add("StreamName", request.StreamName)
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListStreams", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStreamData(request *GetStreamDataRequest) (*GetStreamDataResponse, int, error) {
	resp := new(GetStreamDataResponse)
	query := url.Values{}
	query.Add("StreamID", request.StreamID)
	query.Add("StartTime", request.StartTime)
	query.Add("EndTime", request.EndTime)
	statusCode, err := p.commonHandlerJson("GetStreamData", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StreamLogs(request *StreamLogsRequest) (*StreamLogsResponse, int, error) {
	resp := new(StreamLogsResponse)
	query := url.Values{}
	query.Add("StreamID", request.StreamID)
	query.Add("StartTs", request.StartTs)
	query.Add("EndTs", request.EndTs)
	query.Add("order", request.Order)
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandler("StreamLogs", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//模版相关

func (p *AIoT) CreateScreenshotTemplate(request *CreateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateScreenshotTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateRecordTemplate(request *CreateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateRecordTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateAITemplate(request *CreateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateAITemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateTransTemplate(request *CreateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateTransTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteScreenshotTemplate(request *DeleteTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteScreenshotTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteRecordTemplate(request *DeleteTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteRecordTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteAITemplate(request *DeleteTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteAITemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteTransTemplate(request *DeleteTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteTransTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetScreenshotTemplate(request *TemplateRequest) (*GetTemplateResponse, int, error) {
	resp := new(GetTemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandler("GetScreenshotTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetRecordTemplate(request *TemplateRequest) (*GetTemplateResponse, int, error) {
	resp := new(GetTemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandler("GetRecordTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetAITemplate(request *TemplateRequest) (*GetTemplateResponse, int, error) {
	resp := new(GetTemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandler("GetAITemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetTransTemplate(request *TemplateRequest) (*GetTemplateResponse, int, error) {
	resp := new(GetTemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandler("GetTransTemplate", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListScreenshotTemplates(request *ListTemplateRequest) (*ListTemplateResponse, int, error) {
	resp := new(ListTemplateResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListScreenshotTemplates", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListRecordTemplates(request *ListTemplateRequest) (*ListTemplateResponse, int, error) {
	resp := new(ListTemplateResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListRecordTemplates", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListAITemplates(request *ListTemplateRequest) (*ListTemplateResponse, int, error) {
	resp := new(ListTemplateResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListAITemplates", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListTransTemplates(request *ListTemplateRequest) (*ListTemplateResponse, int, error) {
	resp := new(ListTemplateResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListTransTemplates", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateScreenshotTemplate(request *UpdateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandlerJson("UpdateScreenshotTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateRecordTemplate(request *UpdateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandlerJson("UpdateRecordTemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateAITemplate(request *UpdateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandlerJson("UpdateAITemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateTransTemplate(request *UpdateTemplateRequest) (*TemplateResponse, int, error) {
	resp := new(TemplateResponse)
	query := url.Values{
		"TemplateID": []string{request.TemplateID},
	}
	statusCode, err := p.commonHandlerJson("UpdateAITemplate", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//证书相关
func (p *AIoT) UploadCert(request *UploadCertRequest) (*UploadCertResponse, int, error) {
	resp := new(UploadCertResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("UploadCert", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateCert(request *UpdateCertRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("UpdateCert", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) BindCert(request *BindCertRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("BindCert", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListCertificates(request *ListCertificatesRequest) (*ListCertificatesResponse, int, error) {
	resp := new(ListCertificatesResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandler("ListCertificates", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CertDetail(chainID string) (*CertDetailResponse, int, error) {
	resp := new(CertDetailResponse)
	query := url.Values{}
	query.Add("ChainID", chainID)
	statusCode, err := p.commonHandler("CertDetail", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UnbindCert(request *UnbindCertRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("UnbindCert", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteCert(request *DeleteCertRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteCert", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// 国标级联接口
func (p *AIoT) CreateCascadePlatform(request *CreateCascadePlatformRequest) (*CreateCascadePlatformResponse, int, error) {
	resp := &CreateCascadePlatformResponse{}
	statusCode, err := p.commonHandlerJson("CreateCascadePlatform", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateCascadePlatform(request *UpdateCascadePlatformRequest) (*UpdateCascadePlatformResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{request.PlatformID},
	}

	resp := &UpdateCascadePlatformResponse{}
	statusCode, err := p.commonHandlerJson("UpdateCascadePlatform", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteCascadePlatform(platformID string) (*RawResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}

	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("DeleteCascadePlatform", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) BatchDeleteCascadePlatform(request *BatchDeleteCascadePlatformRequest) (*RawResponse, int, error) {
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("BatchDeleteCascadePlatform", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetCascadePlatform(platformID string) (*GetCascadePlatformResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}

	resp := &GetCascadePlatformResponse{}
	statusCode, err := p.commonHandler("GetCascadePlatform", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListCascadePlatforms(request *ListCascadePlatformRequest) (*ListCascadePlatformResponse, int, error) {
	resp := new(ListCascadePlatformResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))

	if request.PlatformName != "" {
		query.Add("PlatformName", request.Order)
	}
	if request.Order != "" {
		query.Add("Order", request.Order)
	}

	statusCode, err := p.commonHandler("ListCascadePlatform", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateCascadeTask(request *CreateCascadeTaskRequest) (*CreateCascadeTaskResponse, int, error) {
	resp := &CreateCascadeTaskResponse{}
	statusCode, err := p.commonHandlerJson("CreateCascadeTask", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateCascadeTask(taskID string, arg *UpdateCascadeTaskArg) (*RawResponse, int, error) {
	query := url.Values{
		"TaskID": []string{taskID},
	}

	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("UpdateCascadeTask", query, resp, arg)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteCascadeTask(taskID string) (*RawResponse, int, error) {
	query := url.Values{
		"TaskID": []string{taskID},
	}

	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("DeleteCascadeTask", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) BatchDeleteCascadeTask(request *BatchDeleteCascadeTaskRequest) (*RawResponse, int, error) {
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("BatchDeleteCascadeTask", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetCascadeTask(taskID string) (*GetCascadeTaskResponse, int, error) {
	query := url.Values{
		"TaskID": []string{taskID},
	}

	resp := &GetCascadeTaskResponse{}
	statusCode, err := p.commonHandler("GetCascadeTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListCascadeTask(request *ListCascadeTaskRequest) (*ListCascadeTaskResponse, int, error) {
	resp := new(ListCascadeTaskResponse)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))

	if request.PlatformName != "" {
		query.Add("PlatformName", request.Order)
	}
	if request.PlatformID != "" {
		query.Add("PlatformID", request.Order)
	}
	if request.GroupTreeName != "" {
		query.Add("GroupTreeName", request.Order)
	}
	if request.GroupTreeID != "" {
		query.Add("GroupTreeID", request.Order)
	}
	if request.Order != "" {
		query.Add("Order", request.Order)
	}
	statusCode, err := p.commonHandler("ListCascadeTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartCascadeTask(taskID string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"TaskID": []string{taskID},
	}
	statusCode, err := p.commonHandlerJson("StartCascadeTask", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopCascadeTask(taskID string) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{
		"TaskID": []string{taskID},
	}
	statusCode, err := p.commonHandlerJson("StopCascadeTask", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ShareGroupNodesToCascadeTask(arg *ShareGroupNodesRequest) (*RawResponse, int, error) {
	resp := new(RawResponse)
	statusCode, err := p.commonHandlerJson("ShareGroupNodesToCascadeTask", url.Values{}, resp, arg)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetGroupNodesByCascadeTask(taskID string) (*GetGroupTreeNodeResponse, int, error) {
	query := url.Values{
		"TaskID": []string{taskID},
	}

	resp := &GetGroupTreeNodeResponse{}
	statusCode, err := p.commonHandler("GetGroupNodesByCascadeTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// 虚拟组织树接口
func (p *AIoT) GetGroupTreeNode(groupNodeID string) (*GetGroupTreeNodeResponse, int, error) {
	query := url.Values{
		"ID": []string{groupNodeID},
	}
	resp := &GetGroupTreeNodeResponse{}
	statusCode, err := p.commonHandler("GetGroupTreeNode", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) AddGroupTreeNode(request *AddGroupTreeNodeRequest) (*AddGroupTreeNodeResponse, int, error) {
	resp := &AddGroupTreeNodeResponse{}
	statusCode, err := p.commonHandlerJson("AddGroupTreeNode", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetSpaceGroupTreeRoot(spaceID string) (*GroupTreeNode, error) {
	query := url.Values{}
	query.Add("SpaceID", spaceID)

	resp := &GetGroupTreeNodeResponse{}
	statusCode, err := p.commonHandler("ListGroupTreeNodes", query, resp)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", statusCode)
	}

	if resp.ResponseMetadata.Error != nil {
		return nil, fmt.Errorf("code %d, reason:%s",
			resp.ResponseMetadata.Error.CodeN, resp.ResponseMetadata.Error.Message)
	}

	if len(resp.Result.Nodes) == 0 {
		return nil, errors.New("space has no group tree")
	}

	return resp.Result.Nodes[0], err
}

func (p *AIoT) ListGroupTreeNodes(spaceID string) (*ListGroupTreeNodesResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	resp := &ListGroupTreeNodesResponse{}
	statusCode, err := p.commonHandler("ListGroupTreeNodes", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateGroupTreeNode(request *UpdateGroupTreeNodeRequest) (*RawResponse, int, error) {
	query := url.Values{}
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("UpdateGroupTreeNode", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) BindDeviceToGroupTreeNode(request *UpdateGroupTreeNodeRequest) (*RawResponse, int, error) {
	query := url.Values{}
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("BindDeviceToGroupTreeNode", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteGroupTreeNodes(request *DeleteGroupTreeNodesRequest) (*RawResponse, int, error) {
	query := url.Values{}
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("DeleteGroupTreeNodes", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetDevicesByGroupTreeNode(request *GetDevicesByGroupTreeNodeRequest) (*GetDevicesByGroupTreeNodeResponse, int, error) {
	query := url.Values{
		"SearchName": []string{request.SearchName},
		"NodeID":     []string{request.NodeID},
	}
	resp := &GetDevicesByGroupTreeNodeResponse{}
	statusCode, err := p.commonHandler("GetDevicesByGroupTreeNode", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartVoiceTalk(request *StartVoiceTalkRequest) (*StartVoiceTalkResponse, int, error) {
	transport := "tcp"
	if !request.UseTcp {
		transport = "udp"
	}
	query := url.Values{
		"SpaceID":    []string{request.SpaceID},
		"DeviceNSID": []string{request.DeviceNSID},
		"transport":  []string{transport},
	}
	resp := &StartVoiceTalkResponse{}
	statusCode, err := p.commonHandler("StartVoiceTalk", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopVoiceTalk(request *StopVoiceTalkRequest) (*StopVoiceTalkResponse, int, error) {
	query := url.Values{
		"SpaceID":    []string{request.SpaceID},
		"DeviceNSID": []string{request.DeviceNSID},
		"transport":  []string{},
	}
	resp := &StopVoiceTalkResponse{}
	statusCode, err := p.commonHandler("StopVoiceTalk", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//慢直播相关
func (p *AIoT) CreateSlowLive(request *CreateSlowLiveRequest) (*SlowLiveResponse, int, error) {
	resp := &SlowLiveResponse{}
	statusCode, err := p.commonHandlerJson("CreateSlowLive", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateSlowLive(request *UpdateSlowLiveRequest) (*SlowLiveResponse, int, error) {
	query := url.Values{
		"LiveStreamID": []string{request.LiveStreamID},
	}
	resp := &SlowLiveResponse{}
	statusCode, err := p.commonHandlerJson("UpdateSlowLive", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteSlowLive(slowLiveID string) (*SlowLiveResponse, int, error) {
	query := url.Values{
		"LiveStreamID": []string{slowLiveID},
	}
	resp := &SlowLiveResponse{}
	statusCode, err := p.commonHandlerJson("DeleteSlowLive", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) BatchDeleteSlowLive(request *BatchDeleteSlowLiveRequest) (*RawResponse, int, error) {
	resp := &RawResponse{}
	statusCode, err := p.commonHandlerJson("BatchDeleteSlowLive", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StartSlowLive(slowLiveID string) (*SlowLiveResponse, int, error) {
	query := url.Values{
		"LiveStreamID": []string{slowLiveID},
	}
	resp := &SlowLiveResponse{}
	statusCode, err := p.commonHandlerJson("StartSlowLive", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) StopSlowLive(slowLiveID string) (*SlowLiveResponse, int, error) {
	query := url.Values{
		"LiveStreamID": []string{slowLiveID},
	}
	resp := &SlowLiveResponse{}
	statusCode, err := p.commonHandlerJson("StopSlowLive", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetSlowLive(slowLiveID string) (*GetSlowLiveResponse, int, error) {
	query := url.Values{
		"LiveStreamID": []string{slowLiveID},
	}
	resp := &GetSlowLiveResponse{}
	statusCode, err := p.commonHandler("GetSlowLive", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListSlowLive(request *SlowLiveRequest) (*ListSlowLiveResponse, int, error) {
	query := url.Values{
		"PageSize":   []string{fmt.Sprintf("%d", request.PageSize)},
		"PageNumber": []string{fmt.Sprintf("%d", request.PageNumber)},
	}
	resp := &ListSlowLiveResponse{}
	statusCode, err := p.commonHandler("ListSlowLive", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) SignSlowliveWsToken(request *SignSlowliveWsTokenRequest) (*SignSlowliveWsTokenResponse, int, error) {
	query := url.Values{
		"ValidDuration": []string{fmt.Sprintf("%d", request.ValidDuration)},
		"SlowliveID":    []string{request.SlowliveID},
	}
	resp := &SignSlowliveWsTokenResponse{}
	statusCode, err := p.commonHandler("SignSlowliveWsToken", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

//视图
func (p *AIoT) CreateStructuredViewSpace(request *CreateStructuredViewSpaceRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("CreateStructuredViewSpace", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateStructuredViewSpace(request *UpdateStructuredViewSpaceRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("UpdateStructuredViewSpace", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteStructuredViewSpace(spaceID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DeleteStructuredViewSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStructuredViewSpace(spaceID string) (*GetStructuredViewSpaceResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	resp := &GetStructuredViewSpaceResponse{}
	statusCode, err := p.commonHandler("GetStructuredViewSpace", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStructuredViewSpaces(request ListStructuredViewSpacesRequest) (*ListStructuredViewSpacesResponse, int, error) {
	query := url.Values{
		"SpaceName": []string{request.SpaceName},
		"Order":     []string{request.Order},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	resp := &ListStructuredViewSpacesResponse{}
	statusCode, err := p.commonHandler("ListStructuredViewSpaces", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EnableStructuredViewSpace(spaceID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("EnableStructuredViewSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DisableStructuredViewSpace(spaceID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"SpaceID": []string{spaceID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("EnableStructuredViewSpace", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateStructuredViewCode(viewType string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"ViewType": []string{viewType},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandler("CreateStructuredViewCode", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateStructuredView(request *CreateStructuredViewRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("CreateStructuredView", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateStructuredView(request *UpdateStructuredViewRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"ViewID": []string{request.ViewID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("UpdateStructuredView", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteStructuredView(viewID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"ViewID": []string{viewID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DeleteStructuredView", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStructuredView(viewID string) (*GetStructuredViewResponse, int, error) {
	query := url.Values{
		"ViewID": []string{viewID},
	}
	resp := &GetStructuredViewResponse{}
	statusCode, err := p.commonHandler("GetStructuredView", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStructuredViews(request *ListStructuredViewsRequest) (*ListStructuredViewsResponse, int, error) {
	query := url.Values{
		"ViewSpaceName": []string{request.ViewSpaceName},
		"ViewName":      []string{request.ViewName},
		"Order":         []string{request.Order},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	resp := &ListStructuredViewsResponse{}
	statusCode, err := p.commonHandler("ListStructuredViews", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EnableStructuredView(viewID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"ViewID": []string{viewID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("EnableStructuredView", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DisableStructuredView(viewID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"ViewID": []string{viewID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DisableStructuredView", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStructuredViewData(request *ListStructuredViewDataRequest) (*ListStructuredViewDataResponse, int, error) {
	query := url.Values{
		"ViewID":       []string{request.ViewID},
		"StartTs":      []string{request.StartTs},
		"EndTs":        []string{request.EndTs},
		"ViewDataType": []string{request.ViewDataType},
		"Order":        []string{request.Order},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	resp := &ListStructuredViewDataResponse{}
	statusCode, err := p.commonHandler("ListStructuredViewData", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStructuredViewDataInfo(dataID, viewID string) (*GetStructuredViewDataInfoResponse, int, error) {
	query := url.Values{
		"ViewID": []string{viewID},
		"DataID": []string{dataID},
	}
	resp := &GetStructuredViewDataInfoResponse{}
	statusCode, err := p.commonHandler("GetStructuredViewDataInfo", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateStructuredViewCascadePlatform(request *CreateStructuredViewCascadePlatformRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("CreateStructuredViewCascadePlatform", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateStructuredViewCascadePlatform(request *UpdateStructuredViewCascadePlatformRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{request.PlatformID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("UpdateStructuredViewCascadePlatform", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteStructuredViewCascadePlatform(platformID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DeleteStructuredViewCascadePlatform", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStructuredViewCascadePlatform(platformID string) (*GetStructuredViewCascadePlatformResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}
	resp := &GetStructuredViewCascadePlatformResponse{}
	statusCode, err := p.commonHandler("GetStructuredViewCascadePlatform", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStructuredViewCascadePlatform(request *ListStructuredViewCascadePlatformRequest) (*ListStructuredViewCascadePlatformResponse, int, error) {
	query := url.Values{
		"PlatformName": []string{request.PlatformName},
		"Order":        []string{request.Order},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	resp := &ListStructuredViewCascadePlatformResponse{}
	statusCode, err := p.commonHandler("ListStructuredViewCascadePlatform", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EnableStructuredViewCascadePlatform(platformID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("EnableStructuredViewCascadePlatform", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DisableStructuredViewCascadePlatform(platformID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"PlatformID": []string{platformID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DisableStructuredViewCascadePlatform", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateStructuredViewCascadeJob(request *CreateStructuredViewCascadeJobRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("CreateStructuredViewCascadeJob", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateStructuredCascadeJob(request *UpdateStructuredCascadeJobRequest) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"JobID": []string{request.JobID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("UpdateStructuredCascadeJob", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteStructuredViewCascadeJob(jobID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"JobID": []string{jobID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DeleteStructuredViewCascadeJob", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetStructuredViewCascadeJob(jobID string) (*GetStructuredViewCascadeJobResponse, int, error) {
	query := url.Values{
		"JobID": []string{jobID},
	}
	resp := &GetStructuredViewCascadeJobResponse{}
	statusCode, err := p.commonHandler("GetStructuredViewCascadeJob", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStructuredViewCascadeJob(request *ListStructuredViewCascadeJobRequest) (*ListStructuredViewCascadeJobResponse, int, error) {
	query := url.Values{
		"JobName": []string{request.JobName},
		"Order":   []string{request.Order},
	}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	resp := &ListStructuredViewCascadeJobResponse{}
	statusCode, err := p.commonHandler("ListStructuredViewCascadeJob", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EnableStructuredViewCascadeJob(jobID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"JobID": []string{jobID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("EnableStructuredViewCascadeJob", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DisableStructuredViewCascadeJob(jobID string) (*StructuredViewResponse, int, error) {
	query := url.Values{
		"JobID": []string{jobID},
	}
	resp := &StructuredViewResponse{}
	statusCode, err := p.commonHandlerJson("DisableStructuredViewCascadeJob", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EdgeInvite(streamID string, req *EdgeInviteRequest) (*EdgeResponse, int, error) {
	query := url.Values{}
	query.Add("StreamID", streamID)
	resp := &EdgeResponse{}

	statusCode, err := p.commonHandlerJson("EdgeInvite", query, resp, req)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EdgeBye(streamID string) (*EdgeResponse, int, error) {
	query := url.Values{}
	query.Add("StreamID", streamID)
	resp := &EdgeResponse{}

	statusCode, err := p.commonHandler("EdgeBye", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) EdgeStatus(streamID string, req *EdgeStatusRequest) (*EdgeResponse, int, error) {
	query := url.Values{}
	query.Add("StreamID", streamID)
	resp := &EdgeResponse{}

	statusCode, err := p.commonHandlerJson("EdgeStatus", query, resp, req)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetAIAnalysisResult(req *GetAIAnalysisResultRequest) (*GetAIAnalysisResultResponse, int, error) {
	resp := new(GetAIAnalysisResultResponse)
	query := url.Values{
		"StreamID": []string{req.StreamID},
	}
	statusCode, err := p.commonHandlerJson("GetAIAnalysisResult", query, resp, req)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CreateRecordPlan(req *CreateRecordPlanRequest) (*RecordPlanResultResponse, int, error) {
	resp := new(RecordPlanResultResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CreateRecordPlan", query, resp, req)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) UpdateRecordPlan(req *UpdateRecordPlanRequest) (*RecordPlanResultResponse, int, error) {
	resp := new(RecordPlanResultResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("UpdateRecordPlan", query, resp, req)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListRecordPlans(pageNumber, pageSize int) (*ListRecordPlanResponse, int, error) {
	resp := new(ListRecordPlanResponse)
	query := url.Values{
		"PageNumber": []string{fmt.Sprintf("%d", pageNumber)},
		"PageSize":   []string{fmt.Sprintf("%d", pageSize)},
	}
	statusCode, err := p.commonHandlerJson("ListRecordPlans", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListRecordPlanChannels(planID string, pageNumber, pageSize int) (*ListRecordPlanChannelResponse, int, error) {
	resp := new(ListRecordPlanChannelResponse)
	query := url.Values{
		"PlanID":     []string{planID},
		"PageNumber": []string{fmt.Sprintf("%d", pageNumber)},
		"PageSize":   []string{fmt.Sprintf("%d", pageSize)},
	}
	statusCode, err := p.commonHandler("ListRecordPlanChannels", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetRecordPlan(planID string) (*RecordPlanResponse, int, error) {
	resp := new(RecordPlanResponse)
	query := url.Values{
		"PlanID": []string{planID},
	}
	statusCode, err := p.commonHandlerJson("GetRecordPlan", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) DeleteRecordPlan(planID string) (*RecordPlanResultResponse, int, error) {
	resp := new(RecordPlanResultResponse)
	query := url.Values{
		"PlanID": []string{planID},
	}
	statusCode, err := p.commonHandlerJson("DeleteRecordPlan", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// SetCruiseTrack 设置巡航轨迹
func (p *AIoT) SetCruiseTrack(args *SetCruiseTrackArgs) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("SetCruiseTrack", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// GetCruiseTrack 查询通道巡航轨迹
func (p *AIoT) GetCruiseTrack(deviceNSID, channelID string, trackID uint8) (*GetCruiseTrackResponse, int, error) {
	resp := new(GetCruiseTrackResponse)
	query := url.Values{
		"DeviceNSID": []string{deviceNSID},
		"ChannelID":  []string{channelID},
		"TrackID":    []string{fmt.Sprintf("%d", trackID)},
	}
	statusCode, err := p.commonHandlerJson("GetCruiseTrack", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// ListCruiseTracks 查询通道巡航轨迹列表
func (p *AIoT) ListCruiseTracks(deviceNSID, channelID string) (*ListCruiseTracksResponse, int, error) {
	resp := new(ListCruiseTracksResponse)
	query := url.Values{
		"DeviceNSID": []string{deviceNSID},
		"ChannelID":  []string{channelID},
	}
	statusCode, err := p.commonHandlerJson("ListCruiseTracks", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// DeleteCruiseTrack 删除巡航轨迹
func (p *AIoT) DeleteCruiseTrack(args *DeleteCruiseTrackArgs) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteCruiseTrack", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// StartCruiseTrack 开始巡航
func (p *AIoT) StartCruiseTrack(args *StartCruiseTrackArgs) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("StartCruiseTrack", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// StopCruiseTrack 停止巡航
func (p *AIoT) StopCruiseTrack(args *StopCruiseTrackArgs) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("StopCruiseTrack", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetNssInfoList() (*NssInfoListResp, int, error) {
	resp := new(NssInfoListResp)
	query := url.Values{}
	query.Add("type", "relay")
	statusCode, err := p.commonHandler("GetNssInfoList", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// ------- section for V3 version API -----------------------------------------------------
// ------- add v3 api at below place  -----------------------------------------------------

func (p *AIoT) DeleteSpaceV3(request *SpaceRequest) (*SpaceResponseV3, int, error) {
	resp := new(SpaceResponseV3)
	query := url.Values{
		"SpaceID": []string{request.SpaceID},
	}
	statusCode, err := p.commonHandlerJson("DeleteSpaceV3", query, resp, "")
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) LocalMediaDownloadV3(request *LocalMediaDownloadRequestV3) (*LocalMediaDownloadResponseV3, int, error) {
	resp := new(LocalMediaDownloadResponseV3)
	statusCode, err := p.commonHandlerJson("LocalMediaDownloadV3", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CloudControlV3(request *CloudControlRequestV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CloudControlV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) QueryPresetInfoV3(request *QueryPresetInfoRequestV3) (*QueryPresetInfoResponse, int, error) {
	resp := new(QueryPresetInfoResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("QueryPresetInfoV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) GetRecordListV3(request *GetRecordListRequestV3) (*GetRecordListResponse, int, error) {
	resp := new(GetRecordListResponse)
	statusCode, err := p.commonHandlerJson("GetRecordListV3", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CruiseControlV3(request *CruiseControlRequestV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CruiseControlV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// SetCruiseTrack 设置巡航轨迹
func (p *AIoT) SetCruiseTrackV3(args *SetCruiseTrackArgsV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("SetCruiseTrackV3", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// GetCruiseTrack 查询通道巡航轨迹
func (p *AIoT) GetCruiseTrackV3(streamID string, trackID uint8) (*GetCruiseTrackResponse, int, error) {
	resp := new(GetCruiseTrackResponse)
	query := url.Values{
		"StreamID": []string{streamID},
		"TrackID":  []string{fmt.Sprintf("%d", trackID)},
	}
	statusCode, err := p.commonHandlerJson("GetCruiseTrackV3", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// ListCruiseTracks 查询通道巡航轨迹列表
func (p *AIoT) ListCruiseTracksV3(streamID string) (*ListCruiseTracksResponse, int, error) {
	resp := new(ListCruiseTracksResponse)
	query := url.Values{
		"StreamID": []string{streamID},
	}
	statusCode, err := p.commonHandlerJson("ListCruiseTracksV3", query, resp, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// DeleteCruiseTrack 删除巡航轨迹
func (p *AIoT) DeleteCruiseTrackV3(args *DeleteCruiseTrackArgsV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("DeleteCruiseTrackV3", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// StartCruiseTrack 开始巡航
func (p *AIoT) StartCruiseTrackV3(args *StartCruiseTrackArgsV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("StartCruiseTrackV3", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// StopCruiseTrack 停止巡航
func (p *AIoT) StopCruiseTrackV3(args *StopCruiseTrackArgsV3) (*RawResponse, int, error) {
	resp := new(RawResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("StopCruiseTrackV3", query, resp, args)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) PlayBackStartV3(request *PlayBackStartRequestV3) (*PlayBackStartResponse, int, error) {
	resp := new(PlayBackStartResponse)
	statusCode, err := p.commonHandlerJson("PlayBackStartV3", url.Values{}, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) CloudRecordPlayV3(request *CloudRecordPlayRequestV3) (*CloudRecordPlayResponse, int, error) {
	resp := new(CloudRecordPlayResponse)
	query := url.Values{}
	statusCode, err := p.commonHandlerJson("CloudRecordPlayV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStreamScreenshotsV3(request *ListHistoryRequestV3) (*ListHistoryResponseV3, int, error) {
	resp := new(ListHistoryResponseV3)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandlerJson("ListStreamScreenshotsV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *AIoT) ListStreamRecordsV3(request *ListHistoryRequestV3) (*ListHistoryResponseV3, int, error) {
	resp := new(ListHistoryResponseV3)
	query := url.Values{}
	query.Add("PageSize", fmt.Sprintf("%d", request.PageSize))
	query.Add("PageNumber", fmt.Sprintf("%d", request.PageNumber))
	statusCode, err := p.commonHandlerJson("ListStreamRecordsV3", query, resp, request)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
