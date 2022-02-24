package cdn

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestOption func(r *http.Request)

func WithHeader(header http.Header) func(r *http.Request) {
	return func(r *http.Request) {
		if header != nil {
			for k, v := range header {
				r.Header.Set(k, strings.Join(v, ";"))
			}
		}
	}
}

func WithQuery(query url.Values) func(r *http.Request) {
	return func(r *http.Request) {
		if query != nil {
			for k, vv := range query {
				for _, v := range vv {
					r.URL.Query().Add(k, v)
				}
			}
		}
	}
}

func validateResponse(meta *ResponseMetadata) error {
	if meta == nil {
		return errors.New("response meta is not found")
	}
	if meta.Error == nil {
		return nil
	}
	err := *meta.Error
	return fmt.Errorf("code: %v, text code: %v, message: %v", err.CodeN, err.Code, err.Message)
}

func getTimeout(serviceTimeout, apiTimeout time.Duration) time.Duration {
	timeout := time.Second * 60
	if serviceTimeout != time.Duration(0) {
		timeout = serviceTimeout
	}
	if apiTimeout != time.Duration(0) {
		timeout = apiTimeout
	}
	return timeout
}

func mergeQuery(query1, query2 url.Values) (query url.Values) {
	query = url.Values{}
	if query1 != nil {
		for k, vv := range query1 {
			for _, v := range vv {
				query.Add(k, v)
			}
		}
	}

	if query2 != nil {
		for k, vv := range query2 {
			for _, v := range vv {
				query.Add(k, v)
			}
		}
	}
	return
}

func mergeHeaders(headers ...http.Header) (res http.Header) {
	res = http.Header{}
	for _, header := range headers {
		if header != nil {
			for k, v := range header {
				res.Set(k, strings.Join(v, ";"))
			}
		}
	}
	return
}

func (s *CDN) makeRequest(api string, req *http.Request, timeout time.Duration) ([]byte, int, error) {
	req = s.Client.ServiceInfo.Credentials.Sign(req)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := s.Client.Client.Do(req)
	if err != nil {
		// should retry when client sends request error.
		return []byte(""), 500, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, err
	}

	if resp.StatusCode >= 400 {
		return body, resp.StatusCode, fmt.Errorf("[CDN] api %s http code %d body %s", api, resp.StatusCode, string(body))
	}

	return body, resp.StatusCode, nil
}

func (s *CDN) doRequest(apiName string, requestBody interface{}, responseBody interface{}, reqOptions ...RequestOption) (err error) {
	var body []byte
	if requestBody == nil {
		requestBody = struct{}{}
	}
	client := s.Client
	apiInfo := client.ApiInfoList[apiName]
	body, err = json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("[CDN] serialize request body failed, %w", err)
		return
	}
	query := url.Values{}
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeaders(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("[CDN] create http request failed: %v", err)
	}
	req.Header = header
	for _, option := range reqOptions {
		option(req)
	}
	data, _, err := s.makeRequest(apiName, req, timeout)
	if err != nil {
		err = fmt.Errorf("[CDN] request %s failed: %w", apiName, err)
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		err = fmt.Errorf("[CDN] deserilize response body failed: %v", err)
		return
	}
	return
}
