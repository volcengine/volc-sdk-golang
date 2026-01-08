package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *LsClient) DescribeAlarmContentTemplates(request *DescribeAlarmContentTemplatesRequest) (r *DescribeAlarmContentTemplatesResponse, e error) {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	if request.AlarmContentTemplateName != nil {
		params["AlarmContentTemplateName"] = *request.AlarmContentTemplateName
	}

	if request.AlarmContentTemplateID != nil {
		params["AlarmContentTemplateId"] = *request.AlarmContentTemplateID
	}

	if request.OrderField != nil {
		params["OrderField"] = *request.OrderField
	}

	if request.ASC != nil {
		params["ASC"] = strconv.FormatBool(*request.ASC)
	}

	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)

	rawResponse, err := c.Request(http.MethodGet, PathDescribeAlarmContentTemplates, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeAlarmContentTemplatesResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
