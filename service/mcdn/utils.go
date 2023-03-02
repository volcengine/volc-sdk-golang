package mcdn

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type withRequestOption func(r *http.Request)

func WithHeader(header http.Header) func(r *http.Request) {
	return func(r *http.Request) {
		if header != nil && len(header) > 0 {
			return
		}
		for k, v := range header {
			r.Header.Set(k, strings.Join(v, ";"))
		}
	}
}

func WithQuery(query url.Values) func(r *http.Request) {
	return func(r *http.Request) {
		if query == nil {
			return
		}
		q := r.URL.Query()
		for k, vv := range query {
			for _, v := range vv {
				q.Set(k, v)
			}
		}
		r.URL.RawQuery = q.Encode()
	}
}

func IsResponseOk(data base.ResponseMetadata) bool {
	return data.Error == nil
}

func validateResponse(data base.ResponseMetadata) error {
	if !IsResponseOk(data) {
		e := data.Error
		return fmt.Errorf("status: %d, code: %s, request id: %s, message: %s", e.CodeN, e.Code, data.RequestId, e.Message)
	}
	return nil
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
	for k, vv := range query1 {
		for _, v := range vv {
			query.Add(k, v)
		}
	}

	for k, vv := range query2 {
		for _, v := range vv {
			query.Add(k, v)
		}
	}
	return
}

func mergeHeaders(headers ...http.Header) (res http.Header) {
	res = http.Header{}
	for _, header := range headers {
		for k, v := range header {
			res.Set(k, strings.Join(v, ";"))
		}
	}
	return
}

func (s *MCDN) makeRequest(api string, req *http.Request, timeout time.Duration) ([]byte, int, error) {
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
		return body, resp.StatusCode, fmt.Errorf("api %s http code %d body %s", api, resp.StatusCode, string(body))
	}

	return body, resp.StatusCode, nil
}

func (s *MCDN) doRequest(apiName string, requestBody interface{}, responseBody interface{}, reqOptions []withRequestOption) (err error) {
	var body []byte
	if requestBody == nil {
		requestBody = struct{}{}
	}
	client := s.Client
	apiInfo := client.ApiInfoList[apiName]
	body, err = json.Marshal(requestBody)
	if err != nil {
		err = fmt.Errorf("serialize request body failed, err: %w", err)
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
		return fmt.Errorf("create http request failed: err: %w", err)
	}
	req.Header = header
	for _, option := range reqOptions {
		option(req)
	}
	data, _, err := s.makeRequest(apiName, req, timeout)
	if err != nil {
		err = fmt.Errorf("request %s failed, err: %w", apiName, err)
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		err = fmt.Errorf("deserilize response body failed: err: %w", err)
		return
	}
	return
}
