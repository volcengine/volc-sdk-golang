package maas

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

func getTimeout(serviceTimeout, apiTimeout time.Duration) time.Duration {
	timeout := time.Second
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

func mergeHeader(header1, header2 http.Header) (header http.Header) {
	header = http.Header{}

	for k, v := range header1 {
		header.Set(k, strings.Join(v, ";"))
	}

	for k, v := range header2 {
		header.Set(k, strings.Join(v, ";"))
	}

	return
}

func makeRequest(apiInfo *base.ApiInfo, serviceInfo *base.ServiceInfo, query url.Values, ct string) (*http.Request, error) {
	header := mergeHeader(serviceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   serviceInfo.Scheme,
		Host:     serviceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request")
	}
	req.Header = header
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	// Because service info could be changed by SetRegion, so set UA header for every request here.
	req.Header.Set("User-Agent", strings.Join([]string{base.SDKName, base.SDKVersion}, "/"))

	return req, nil
}
