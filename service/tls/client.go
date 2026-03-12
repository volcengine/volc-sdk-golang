package tls

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-kit/kit/log/level"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/tls/innerlogger"
)

type gzipReadCloser struct {
	reader *gzip.Reader
	body   io.Closer
}

func (r *gzipReadCloser) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *gzipReadCloser) Close() error {
	err1 := r.reader.Close()
	err2 := r.body.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

var (
	defaultHttpClient *http.Client

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

	retryPolicy atomic.Value
}

const (
	defaultRetryTimeout = 90 * time.Second
)

func NewClient(endpoint, accessKeyID, accessKeySecret, securityToken, region string) Client {
	apiVersion := APIVersion3
	client := &LsClient{
		Client:          defaultHttpClient,
		Endpoint:        endpoint,
		accessLock:      &sync.RWMutex{},
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Region:          region,
		APIVersion:      apiVersion,
	}
	client.retryPolicy.Store(DefaultRetryPolicy())
	return client
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

func (c *LsClient) SetRetryPolicy(policy RetryPolicy) {
	normalized := policy.Normalize()
	c.retryPolicy.Store(normalized)
}

func (c *LsClient) GetRetryPolicy() RetryPolicy {
	if policy, ok := c.retryPolicy.Load().(RetryPolicy); ok {
		return policy.Normalize()
	}
	return DefaultRetryPolicy()
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

	policy := c.GetRetryPolicy()

	// 为当前请求创建独立的上下文，限定重试的最大耗时
	ctx, cancel := context.WithTimeout(context.Background(), policy.TotalTimeout)
	defer cancel()
	attempts := 0
	retryOp := func() (bool, error) {
		attempts++
		// 为单次 HTTP 尝试创建独立的请求上下文，避免重试上下文在成功返回时立即取消导致服务端感知 cancel
		// 每次尝试的超时时间取 HTTP 客户端的 Timeout（若未设置则由默认客户端提供）
		perAttemptTimeout := c.Client.Timeout
		if perAttemptTimeout <= 0 {
			perAttemptTimeout = time.Second * 60
		}
		reqCtx, reqCancel := context.WithTimeout(ctx, perAttemptTimeout)
		defer reqCancel()
		r, iErr = c.realRequest(reqCtx, method, realUri, headers, body)
		if iErr != nil {
			level.Error(innerlogger.DefaultLogger).Log("msg", "Request failed", "reason", iErr)
			level.Debug(innerlogger.DefaultLogger).Log("method", method, "uri", realUri, "headers", headers, "body", string(body))
		}
		if iErr == nil {
			return false, nil
		}
		need := needRetryError(iErr)
		if need {
			// 检查是否超过最大重试次数
			if policy.MaxAttempts <= 0 || attempts < policy.MaxAttempts {
				return true, iErr
			}
		}
		return false, iErr
	}
	// 使用 RetryWithConditionTicker + backoff.ExponentialBackOff
	exp := policy.ExponentialBackOff()

	retryErr := RetryWithConditionTicker(ctx, backoff.WithContext(exp, ctx), retryOp)

	if retryErr != nil {
		return r, retryErr
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
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
		return nil, err
	}
	if resp.Header.Get("Content-Encoding") == CompressGz {
		rawBody := resp.Body
		resp.Body, err = gzip.NewReader(rawBody)
		if err != nil {
			_ = rawBody.Close()
			return nil, err
		}
		resp.Body = &gzipReadCloser{
			reader: resp.Body.(*gzip.Reader),
			body:   rawBody,
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
