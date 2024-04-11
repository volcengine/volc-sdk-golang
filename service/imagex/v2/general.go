package imagex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	common "github.com/volcengine/volc-sdk-golang/base"
)

func (c *Imagex) ImageXGet(action string, query url.Values, result interface{}) error {
	respBody, _, err := c.Client.Query(action, query)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(respBody, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *Imagex) ImageXPost(action string, query url.Values, req, result interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.Json(action, query, string(body))
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *Imagex) GeneralGet(ctx context.Context, action, version string, query url.Values, result interface{}) error {
	apiInfo := &common.ApiInfo{
		Method: http.MethodGet,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{action},
			"Version": []string{version},
		},
	}
	respBody, _, err := c.Client.CtxQueryThumb(ctx, action, apiInfo, query)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(respBody, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *Imagex) GeneralPost(ctx context.Context, action, version string, query url.Values, req, result interface{}) error {
	apiInfo := &common.ApiInfo{
		Method: http.MethodPost,
		Path:   "/",
		Query: url.Values{
			"Action":  []string{action},
			"Version": []string{version},
		},
	}
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.CtxJsonThumb(ctx, action, apiInfo, query, body)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := unmarshalInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}
