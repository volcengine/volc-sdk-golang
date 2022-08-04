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
