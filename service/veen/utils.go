package veen

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func packErrorInfo(metadata VolcResponseMetadata) error {
	return fmt.Errorf("RequestId=%s Error=%v", metadata.RequestId, metadata.Error)
}

func (v *Veen) post(apiName string, reqBody interface{}, respBody interface{}) error {
	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	query := url.Values{}
	resp, _, err := v.Client.Json(apiName, query, string(body))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp, respBody); err != nil {
		return err
	}
	return nil
}

func (v *Veen) get(apiName string, query url.Values, respBody interface{}) error {
	resp, _, err := v.Client.Query(apiName, query)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp, respBody); err != nil {
		return err
	}
	return nil
}

func ToJson(v interface{}) string {
	bs, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bs)
}
