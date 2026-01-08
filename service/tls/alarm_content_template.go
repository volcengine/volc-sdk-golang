package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// ModifyAlarmContentTemplate API
func (c *LsClient) ModifyAlarmContentTemplate(request *ModifyAlarmContentTemplateRequest) (r *ModifyAlarmContentTemplateResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyAlarmContentTemplate, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, NewClientError(errors.New("nil http response"))
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &ModifyAlarmContentTemplateResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
