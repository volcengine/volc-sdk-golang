package tls

import (
	"net/http"
	"strconv"
)

func (c *LsClient) WebTracks(request *WebTracksRequest) (r *WebTracksResponse, e error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	params := map[string]string{
		"TopicId":   request.TopicID,
		"ProjectId": request.ProjectID,
	}

	reqHeaders := map[string]string{
		"Content-Type":       "application/json",
		"x-tls-compresstype": request.CompressType,
	}

	bytesBody, rawBodyLength, err := GetWebTracksBody(request.CompressType, request)
	if err != nil {
		return nil, err
	}
	reqHeaders[rawBodySizeHeader] = strconv.Itoa(rawBodyLength)

	rawResponse, err := c.Request(http.MethodPost, PathWebTracks, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	var response = &WebTracksResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}
