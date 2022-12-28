package tls

import (
	"bytes"
	"context"

	"github.com/volcengine/volc-sdk-golang/base"

	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/service/tls/innerlogger"
)

var (
	defaultHttpClient *http.Client
)

func init() {
	defaultHttpClient = &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     10 * time.Second,
			DisableCompression:  true,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
		},
	}
}

type LsClient struct {
	BaseClient      *base.Client
	Endpoint        string
	accessLock      *sync.RWMutex
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	UserAgent       string
	RequestTimeOut  time.Duration
	RetryTimeOut    time.Duration
	Region          string
}

// ResetAccessKeyToken reset client's access key token
func (c *LsClient) ResetAccessKeyToken(accessKeyID, accessKeySecret, securityToken string) {
	c.accessLock.Lock()
	c.AccessKeyID = accessKeyID
	c.AccessKeySecret = accessKeySecret
	c.SecurityToken = securityToken
	c.accessLock.Unlock()
}

func (c *LsClient) SetTimeout(timeout time.Duration) {
	defaultHttpClient.Timeout = timeout
}

func (c *LsClient) Request(method, uri string, params map[string]string, headers map[string]string, body []byte) (rsp *http.Response, e error) {
	defer func() {
		if e != nil {
			level.Error(innerlogger.DefaultLogger).Log(
				"sdk", GetCallerFuncName(3),
				"uri", fmt.Sprintf("%v", uri),
				"params", fmt.Sprintf("%+v", params),
				"headers", fmt.Sprintf("%+v", headers),
				"error", e,
			)
		}
		if innerlogger.DefaultLogger.GetLogLevel() <= innerlogger.LogLevelDebug {
			level.Debug(innerlogger.DefaultLogger).Log(
				"sdk", GetCallerFuncName(3),
				"uri", fmt.Sprintf("%v", uri),
				"params", fmt.Sprintf("%+v", params),
				"headers", fmt.Sprintf("%+v", headers),
				"body", string(body),
				"error", e,
			)
		}
	}()

	var r *http.Response
	var iErr error
	var err error
	var realUri string = uri

	if len(params) != 0 {
		realUri = appendParam(uri, params)
	} else {
		realUri = uri
	}

	ctx := context.Background()
	err = RetryWithCondition(ctx, backoff.NewExponentialBackOff(), func() error {
		r, iErr = c.realRequest(ctx, method, realUri, headers, body)
		if iErr != nil {
			level.Error(innerlogger.DefaultLogger).Log("msg", "Request failed", "reason", iErr)
			level.Debug(innerlogger.DefaultLogger).Log("method", method, "uri", realUri, "headers", headers, "body", string(body))
		}
		return iErr
	})

	if err != nil {
		return r, err
	}
	return r, iErr
}

func (c *LsClient) assembleHeader(request CommonRequest, headers map[string]string) map[string]string {
	newHeaders := map[string]string{}
	for key, value := range request.Headers {
		newHeaders[key] = value
	}
	for key, value := range headers {
		newHeaders[key] = value
	}
	return newHeaders
}

func appendParam(originalUrl string, params map[string]string) string {
	if len(params) == 0 {
		return originalUrl
	}

	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(originalUrl)
	urlBuilder.WriteString("?")

	var offset = 0

	for key, value := range params {
		urlBuilder.WriteString(key)
		urlBuilder.WriteString("=")
		urlBuilder.WriteString(value)

		if offset != len(params)-1 {
			urlBuilder.WriteString("&")
		}
		offset++
	}

	return urlBuilder.String()
}

func (c *LsClient) realRequest(ctx context.Context, method, uri string, headers map[string]string, body []byte) (*http.Response, error) {
	headers[AgentHeader] = defaultLogUserAgent

	if body != nil {
		bodyMD5 := fmt.Sprintf("%X", md5.Sum(body))
		headers[ContentMd5Header] = bodyMD5
	}

	reader := bytes.NewReader(body)
	urlStr := fmt.Sprintf("%s%s", c.Endpoint, uri)
	req, err := http.NewRequest(method, urlStr, reader)
	if err != nil {
		return nil, NewClientError(err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	c.accessLock.RLock()

	credential := base.Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.AccessKeySecret,
		Region:          c.Region,
		Service:         ServiceName,
		SessionToken:    c.SecurityToken,
	}

	c.accessLock.RUnlock()

	req = credential.Sign(req)

	// Get ready to do request
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := &Error{}
		err.HTTPCode = (int32)(resp.StatusCode)
		defer resp.Body.Close()
		buf, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return nil, NewBadResponseError(ioErr.Error(), resp.Header, resp.StatusCode)
		}
		if jErr := json.Unmarshal(buf, err); jErr != nil {
			return nil, NewBadResponseError(string(buf), resp.Header, resp.StatusCode)
		}
		err.RequestID = resp.Header.Get(RequestIDHeader)
		return nil, err
	}

	return resp, nil
}

func (c *LsClient) Close() error {
	return nil
}

func (c *LsClient) String() string {
	return c.Endpoint + " " + c.Region
}
