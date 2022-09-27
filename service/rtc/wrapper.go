package rtc

import (
	"encoding/json"
	"net/url"
)

// GET method
// GetRecordTask ...
func GetRecordTask(r *RTC, query url.Values) (*GetRecordTaskResponse, int, error) {
	respBody, status, err := r.Client.Query(ActionGetRecordTask, query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetRecordTaskResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

// POST method
// StartRecord ...
func StartRecord(r *RTC, req *StartRecordRequest) (*StartRecordResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}

	respBody, status, err := r.Client.Json(ActionStartRecord, nil, string(bts))
	if err != nil {
		return nil, status, err
	}

	output := new(StartRecordResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

func StartWebRecord(r *RTC, req *StartWebRecordRequest) (*StartWebRecordResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}

	respBody, status, err := r.Client.Json(ActionStartWebRecord, nil, string(bts))
	if err != nil {
		return nil, status, err
	}

	output := new(StartWebRecordResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

func StopWebRecord(r *RTC, req *StopWebRecordRequest) (*StopWebRecordResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}

	respBody, status, err := r.Client.Json(ActionStopWebRecord, nil, string(bts))
	if err != nil {
		return nil, status, err
	}

	output := new(StopWebRecordResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

func GetWebRecordTask(r *RTC, query url.Values) (*GetWebRecordTaskResponse, int, error) {
	respBody, status, err := r.Client.Query(ActionGetWebRecordTask, query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetWebRecordTaskResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

func GetWebRecordList(r *RTC, query url.Values) (*GetWebRecordListResponse, int, error) {
	respBody, status, err := r.Client.Query(ActionGetWebRecordList, query)
	if err != nil {
		return nil, status, err
	}

	output := new(GetWebRecordListResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

/*
// ActionExample, add new action
func ActionExample(r *RTC, req interface{}) (*CommonResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}

    // use r.Client.Query when query
	respBody, status, err := r.Client.Json("ActionExample", nil, string(bts))
	if err != nil {
		return nil, status, err
	}

	output := new(CommonResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}
*/
