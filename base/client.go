package base

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/net/http/httpproxy"
)

const (
	accessKey = "VOLC_ACCESSKEY"
	secretKey = "VOLC_SECRETKEY"

	// volc proxy
	httpProxy     = "VOLC_HTTP_PROXY"
	httpsProxy    = "VOLC_HTTPS_PROXY"
	noProxy       = "VOLC_NO_PROXY"
	requestMethod = "REQUEST_METHOD"

	defaultScheme = "http"
)

var (
	_GlobalClient   *http.Client
	emptyBytes      []byte
	emptyReadSeeker = bytes.Buffer{}
)

func volcProxy() func(req *http.Request) (*url.URL, error) {
	c := &httpproxy.Config{
		HTTPProxy:  os.Getenv(httpProxy),
		HTTPSProxy: os.Getenv(httpsProxy),
		NoProxy:    os.Getenv(noProxy),
		CGI:        os.Getenv(requestMethod) != "",
	}
	p := c.ProxyFunc()
	return func(req *http.Request) (*url.URL, error) { return p(req.URL) }
}

func init() {
	_GlobalClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     10 * time.Second,
			Proxy:               volcProxy(),
		},
	}
}

// Client
type Client struct {
	Client        *http.Client
	ServiceInfo   *ServiceInfo
	ApiInfoList   map[string]*ApiInfo
	CustomTimeout time.Duration
}

// NewClient
func NewClient(info *ServiceInfo, apiInfoList map[string]*ApiInfo) *Client {
	client := &Client{Client: _GlobalClient, ServiceInfo: info.Clone(), ApiInfoList: apiInfoList}

	if client.ServiceInfo.Scheme == "" {
		client.ServiceInfo.Scheme = defaultScheme
	}

	if os.Getenv(accessKey) != "" && os.Getenv(secretKey) != "" {
		client.ServiceInfo.Credentials.AccessKeyID = os.Getenv(accessKey)
		client.ServiceInfo.Credentials.SecretAccessKey = os.Getenv(secretKey)
	} else if _, err := os.Stat(os.Getenv("HOME") + "/.volc/config"); err == nil {
		if content, err := ioutil.ReadFile(os.Getenv("HOME") + "/.volc/config"); err == nil {
			m := make(map[string]string)
			json.Unmarshal(content, &m)
			if accessKey, ok := m["ak"]; ok {
				client.ServiceInfo.Credentials.AccessKeyID = accessKey
			}
			if secretKey, ok := m["sk"]; ok {
				client.ServiceInfo.Credentials.SecretAccessKey = secretKey
			}
		}
	}
	return client
}

func (serviceInfo *ServiceInfo) Clone() *ServiceInfo {
	ret := new(ServiceInfo)
	//base info
	ret.Timeout = serviceInfo.Timeout
	ret.Host = serviceInfo.Host
	ret.Scheme = serviceInfo.Scheme

	//credential
	ret.Credentials = serviceInfo.Credentials.Clone()

	// header
	ret.Header = serviceInfo.Header.Clone()
	return ret
}

func (cred Credentials) Clone() Credentials {
	return Credentials{
		Service:         cred.Service,
		Region:          cred.Region,
		SecretAccessKey: cred.SecretAccessKey,
		AccessKeyID:     cred.AccessKeyID,
		SessionToken:    cred.SessionToken,
	}
}

// SetRetrySettings
func (client *Client) SetRetrySettings(retrySettings *RetrySettings) {
	if retrySettings != nil {
		client.ServiceInfo.Retry = *retrySettings
	}
}

// SetAccessKey
func (client *Client) SetAccessKey(ak string) {
	if ak != "" {
		client.ServiceInfo.Credentials.AccessKeyID = ak
	}
}

// SetSecretKey
func (client *Client) SetSecretKey(sk string) {
	if sk != "" {
		client.ServiceInfo.Credentials.SecretAccessKey = sk
	}
}

// SetSessionToken
func (client *Client) SetSessionToken(token string) {
	if token != "" {
		client.ServiceInfo.Credentials.SessionToken = token
	}
}

// SetHost
func (client *Client) SetHost(host string) {
	if host != "" {
		client.ServiceInfo.Host = host
	}
}

func (client *Client) SetScheme(scheme string) {
	if scheme != "" {
		client.ServiceInfo.Scheme = scheme
	}
}

// SetCredential
func (client *Client) SetCredential(c Credentials) {
	if c.AccessKeyID != "" {
		client.ServiceInfo.Credentials.AccessKeyID = c.AccessKeyID
	}

	if c.SecretAccessKey != "" {
		client.ServiceInfo.Credentials.SecretAccessKey = c.SecretAccessKey
	}

	if c.Region != "" {
		client.ServiceInfo.Credentials.Region = c.Region
	}

	if c.SessionToken != "" {
		client.ServiceInfo.Credentials.SessionToken = c.SessionToken
	}

	if c.Service != "" {
		client.ServiceInfo.Credentials.Service = c.Service
	}
}

func (client *Client) SetTimeout(timeout time.Duration) {
	if timeout > 0 {
		client.ServiceInfo.Timeout = timeout
	}
}

func (client *Client) SetCustomTimeout(timeout time.Duration) {
	if timeout > 0 {
		client.CustomTimeout = timeout
	}
}

// GetSignUrl
func (client *Client) GetSignUrl(api string, query url.Values) (string, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return "", errors.New("The related api does not exist")
	}

	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), nil)

	if err != nil {
		return "", errors.New("Failed to build request")
	}

	return client.ServiceInfo.Credentials.SignUrl(req), nil
}

// SignSts2
func (client *Client) SignSts2(inlinePolicy *Policy, expire time.Duration) (*SecurityToken2, error) {
	var err error
	sts := new(SecurityToken2)
	if sts.AccessKeyID, sts.SecretAccessKey, err = createTempAKSK(); err != nil {
		return nil, err
	}

	if expire < time.Minute {
		expire = time.Minute
	}

	now := time.Now()
	expireTime := now.Add(expire)
	sts.CurrentTime = now.Format(time.RFC3339)
	sts.ExpiredTime = expireTime.Format(time.RFC3339)

	innerToken, err := createInnerToken(client.ServiceInfo.Credentials, sts, inlinePolicy, expireTime.Unix())
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(innerToken)
	sts.SessionToken = "STS2" + base64.StdEncoding.EncodeToString(b)
	return sts, nil
}

// Query Initiate a Get query request
func (client *Client) Query(api string, query url.Values) ([]byte, int, error) {
	return client.CtxQuery(context.Background(), api, query)
}

func (client *Client) CtxQuery(ctx context.Context, api string, query url.Values) ([]byte, int, error) {
	return client.request(ctx, api, query, emptyBytes, "")
}

// Json Initiate a Json post request
func (client *Client) Json(api string, query url.Values, body string) ([]byte, int, error) {
	return client.CtxJson(context.Background(), api, query, body)
}

func (client *Client) CtxJson(ctx context.Context, api string, query url.Values, body string) ([]byte, int, error) {
	return client.request(ctx, api, query, []byte(body), "application/json")
}
func (client *Client) PostWithContentType(api string, query url.Values, body string, ct string) ([]byte, int, error) {
	return client.CtxPostWithContentType(context.Background(), api, query, body, ct)
}

// CtxPostWithContentType Initiate a post request with a custom Content-Type, Content-Type cannot be empty
func (client *Client) CtxPostWithContentType(ctx context.Context, api string, query url.Values, body string, ct string) ([]byte, int, error) {
	return client.request(ctx, api, query, []byte(body), ct)
}

func (client *Client) Post(api string, query url.Values, form url.Values) ([]byte, int, error) {
	return client.CtxPost(context.Background(), api, query, form)
}

// CtxPost Initiate a Post request
func (client *Client) CtxPost(ctx context.Context, api string, query url.Values, form url.Values) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]
	form = mergeQuery(form, apiInfo.Form)
	return client.request(ctx, api, query, []byte(form.Encode()), "application/x-www-form-urlencoded")
}

func (client *Client) CtxMultiPart(ctx context.Context, api string, query url.Values, form []*MultiPartItem) ([]byte, int, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, item := range form {
		part, err := writer.CreatePart(item.header)
		if err != nil {
			return nil, 400, err
		}
		_, err = io.Copy(part, item.data)
		if err != nil {
			return nil, 400, err
		}
	}
	writer.Close()
	return client.request(ctx, api, query, body.Bytes(), writer.FormDataContentType())
}

func (client *Client) makeRequest(inputContext context.Context, api string, req *http.Request, timeout time.Duration) ([]byte, int, error, bool) {
	req = client.ServiceInfo.Credentials.Sign(req)

	ctx := inputContext
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := client.Client.Do(req)
	if err != nil {
		// should retry when client sends request error.
		return []byte(""), 500, err, true
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, err, false
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		needRetry := false
		// should retry when server returns 5xx error.
		if resp.StatusCode >= http.StatusInternalServerError {
			needRetry = true
		}
		return body, resp.StatusCode, fmt.Errorf("api %s http code %d body %s", api, resp.StatusCode, string(body)), needRetry
	}

	return body, resp.StatusCode, nil, false
}

func (client *Client) request(ctx context.Context, api string, query url.Values, body []byte, ct string) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("The related api does not exist")
	}
	return client.requestThumb(ctx, api, apiInfo, query, body, ct)
}

func (client *Client) requestThumb(ctx context.Context, api string, apiInfo *ApiInfo, query url.Values, body []byte, ct string) ([]byte, int, error) {
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout, client.CustomTimeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)
	retrySettings := getRetrySetting(&client.ServiceInfo.Retry, &apiInfo.Retry)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	requestBody := bytes.NewReader(body)
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), nil)
	if err != nil {
		return []byte(""), 500, fmt.Errorf("Failed to build request, err %w", err)
	}
	req.Header = header
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}
	// Because service info could be changed by SetRegion, so set UA header for every request here.
	req.Header.Set("User-Agent", strings.Join([]string{SDKName, SDKVersion}, "/"))

	var resp []byte
	var code int

	err = backoff.Retry(func() error {
		_, err = requestBody.Seek(0, io.SeekStart)
		if err != nil {
			// if seek failed, stop retry.
			return backoff.Permanent(err)
		}
		req.Body = ioutil.NopCloser(requestBody)
		var needRetry bool
		resp, code, err, needRetry = client.makeRequest(ctx, api, req, timeout)
		if needRetry {
			return err
		} else {
			return backoff.Permanent(err)
		}
	}, backoff.WithMaxRetries(backoff.NewConstantBackOff(*retrySettings.RetryInterval), *retrySettings.RetryTimes))
	return resp, code, err
}

func (client *Client) CtxQueryThumb(ctx context.Context, api string, apiInfo *ApiInfo, query url.Values) ([]byte, int, error) {
	return client.requestThumb(ctx, api, apiInfo, query, emptyBytes, "")
}

func (client *Client) CtxJsonThumb(ctx context.Context, api string, apiInfo *ApiInfo, query url.Values, body []byte) ([]byte, int, error) {
	return client.requestThumb(ctx, api, apiInfo, query, body, "application/json")
}
