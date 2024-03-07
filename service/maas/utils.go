package maas

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName       = "ml_maas"
	APICert           = "cert"
	APIChat           = "chat"
	APIStreamChat     = "stream_chat"
	APITokenization   = "tokenization"
	APIClassification = "classification"
	APIEmbeddings     = "embeddings"
	APIImagesQuickGen = "images.quick_gen"
	APIAudioSpeech    = "audio.speech"
	TOP               = "top"

	ServiceTimeout = time.Minute

	MaxBufferSize  = 4194304
	RespBufferSize = 32

	Terminator = "[DONE]"

	ChatRoleOfUser      = "user"
	ChatRoleOfAssistant = "assistant"
	ChatRoleOfSystem    = "system"
	ChatRoleOfFunction  = "function"
)

var (
	defaultRetryTimes    uint64 = 2
	defaultRetryInterval        = 1 * time.Second
)

func GetTimeout(serviceTimeout, apiTimeout time.Duration) time.Duration {
	timeout := time.Second
	if serviceTimeout != time.Duration(0) {
		timeout = serviceTimeout
	}
	if apiTimeout != time.Duration(0) {
		timeout = apiTimeout
	}
	return timeout
}

func MergeQuery(query1, query2 url.Values) (query url.Values) {
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

func MergeHeader(header1, header2 http.Header) (header http.Header) {
	header = http.Header{}

	for k, v := range header1 {
		header.Set(k, strings.Join(v, ";"))
	}

	for k, v := range header2 {
		header.Set(k, strings.Join(v, ";"))
	}

	return
}

func MakeRequest(apiInfo *base.ApiInfo, endpointId string, serviceInfo *base.ServiceInfo, query url.Values, ct string) (*http.Request, error) {
	header := MergeHeader(serviceInfo.Header, apiInfo.Header)
	query = MergeQuery(query, apiInfo.Query)
	path := apiInfo.Path
	if endpointId != "" {
		path = fmt.Sprintf(apiInfo.Path, endpointId)
	}

	u := url.URL{
		Scheme:   serviceInfo.Scheme,
		Host:     serviceInfo.Host,
		Path:     path,
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

func GetRetrySetting(serviceRetrySettings, apiRetrySettings *base.RetrySettings) *base.RetrySettings {
	retrySettings := &base.RetrySettings{
		AutoRetry:     false,
		RetryTimes:    new(uint64),
		RetryInterval: new(time.Duration),
	}
	if !apiRetrySettings.AutoRetry || !serviceRetrySettings.AutoRetry {
		return retrySettings
	}
	retrySettings.AutoRetry = true
	if serviceRetrySettings.RetryTimes != nil {
		retrySettings.RetryTimes = serviceRetrySettings.RetryTimes
	} else if apiRetrySettings.RetryTimes != nil {
		retrySettings.RetryTimes = apiRetrySettings.RetryTimes
	} else {
		retrySettings.RetryTimes = &defaultRetryTimes
	}
	if serviceRetrySettings.RetryInterval != nil {
		retrySettings.RetryInterval = serviceRetrySettings.RetryInterval
	} else if apiRetrySettings.RetryInterval != nil {
		retrySettings.RetryInterval = apiRetrySettings.RetryInterval
	} else {
		retrySettings.RetryInterval = &defaultRetryInterval
	}
	return retrySettings
}
