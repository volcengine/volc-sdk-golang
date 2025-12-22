package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (c *LsClient) CancelDownloadTask(request *CancelDownloadTaskRequest) (r *CancelDownloadTaskResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCancelDownloadTask, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, errors.New("nil http response")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CancelDownloadTaskResponse{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}
