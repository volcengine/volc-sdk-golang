package tls

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/tls/innerlogger"
)

var (
	defaultHttpClient          *http.Client
	defaultRetryInterval       time.Duration
	defaultRetryCounter        int32
	defaultRetryCounterMaximum int32
	defaultRequestTimeout      time.Duration

	defaultUserAgent = fmt.Sprintf("%s/%s", base.SDKName, base.SDKVersion)
)

func init() {
	defaultHttpClient = &http.Client{
		Timeout: time.Second * 60,
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
	defaultRetryCounter = 0
	defaultRetryCounterMaximum = 50
	defaultRetryInterval = 100 * time.Millisecond
	defaultRequestTimeout = time.Second * 90
}

type LsClient struct {
	Client          *http.Client
	Endpoint        string
	accessLock      *sync.RWMutex
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	UserAgent       string
	RequestTimeOut  time.Duration
	Region          string
	APIVersion      string
	CustomUserAgent string
}

func (c *LsClient) SetAPIVersion(version string) {
	c.APIVersion = version
}

func (c *LsClient) SetCustomUserAgent(customUserAgent string) {
	c.CustomUserAgent = customUserAgent
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
	c.Client.Timeout = timeout
	defaultRequestTimeout = timeout
}

func (c *LsClient) GetHttpClient() *http.Client {
	return c.Client
}

func (c *LsClient) SetHttpClient(client *http.Client) error {
	if client == nil {
		return errors.New("client can not be nil")
	}
	c.Client = client
	return nil
}

func (c *LsClient) SetRetryIntervalMs(interval time.Duration) {
	defaultRetryInterval = interval
}

func (c *LsClient) SetRetryCounterMaximum(maximum int32) {
	defaultRetryCounterMaximum = maximum
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

	// 检查是否正确配置了 Region/AK/SK/Endpoint
	{
		if len(c.Region) <= 0 {
			return nil, NewClientError(errors.New("Empty Region; 请在初始化时填入 Region"))
		}
		if len(c.Endpoint) <= 0 {
			return nil, NewClientError(errors.New("Empty Endpoint; 请在初始化时填入 Endpoint"))
		}
	}

	if len(params) != 0 {
		realUri = appendParam(uri, params)
	} else {
		realUri = uri
	}

	ctx := context.Background()
	err = RetryWithCondition(ctx, func() error {
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
	if len(c.CustomUserAgent) == 0 {
		headers[AgentHeader] = defaultUserAgent
	} else {
		headers[AgentHeader] = c.CustomUserAgent
	}

	// 如果header没有配置api version，增加默认的api version 0.3.0
	if _, ok := headers[HeaderAPIVersion]; !ok {
		headers[HeaderAPIVersion] = c.APIVersion
	}

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
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Header.Get("Content-Encoding") == CompressGz {
		resp.Body, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
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

func (c *LsClient) DescribeScheduleSqlTasks(request *DescribeScheduleSqlTasksRequest) (*DescribeScheduleSqlTasksResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := make(map[string]string)
	if request.ProjectId != nil {
		params["ProjectId"] = *request.ProjectId
	}
	if request.ProjectName != nil {
		params["ProjectName"] = *request.ProjectName
	}
	if request.IamProjectName != nil {
		params["IamProjectName"] = *request.IamProjectName
	}
	if request.TopicId != nil {
		params["TopicId"] = *request.TopicId
	}
	if request.SourceTopicName != nil {
		params["SourceTopicName"] = *request.SourceTopicName
	}
	if request.TaskId != nil {
		params["TaskId"] = *request.TaskId
	}
	if request.TaskName != nil {
		params["TaskName"] = *request.TaskName
	}
	if request.Status != nil {
		params["Status"] = *request.Status
	}
	if request.PageNumber != nil {
		params["PageNumber"] = strconv.Itoa(*request.PageNumber)
	}
	if request.PageSize != nil {
		params["PageSize"] = strconv.Itoa(*request.PageSize)
	}

	body := map[string]string{}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeScheduleSqlTasks, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
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

	var response = &DescribeScheduleSqlTasksResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
