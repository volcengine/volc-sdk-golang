package cdn

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/url"
)

func wrapperError(errObj *base.ErrorObj) (err error) {
	if errObj == nil {
		return
	}
	err = CDNError{
		Message: errObj.Message,
		Status:  errObj.CodeN,
		Code:    errObj.Code,
	}
	return
}

func validateResponse(meta *base.ResponseMetadata) (err error) {
	if meta == nil {
		return errors.New("response meta is not found")
	}
	err = wrapperError(meta.Error)
	return
}

func (s *CDN) post(apiName string, requestBody interface{}, responseBody interface{}) (err error) {
	var body []byte
	if requestBody == nil {
		requestBody = struct{}{}
	}
	body, err = json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("marshal request body failed, %w", err)
		return
	}
	query := url.Values{}
	data, _, err := s.Client.Json(apiName, query, string(body))
	if err != nil {
		err = fmt.Errorf("request %s api failed: %w", apiName, err)
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		err = fmt.Errorf("unmarshal response body failed: %w", err)
		return
	}
	return
}
