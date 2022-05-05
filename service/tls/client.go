package tls

import (
	"bytes"
	"context"

	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
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

func (c *LsClient) Request(method, uri string, params map[string]string, headers map[string]string, body []byte) (*http.Response, error) {
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
		return iErr
	})

	if err != nil {
		return r, err
	}
	return r, iErr
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
		h := sha256.New()
		bodyMD5 := fmt.Sprintf("%X", h.Sum(body))
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

	credential := Credentials{
		AccessKeyID:     c.AccessKeyID,
		SecretAccessKey: c.AccessKeySecret,
		Region:          c.Region,
		Service:         ServiceName,
		SessionToken:    c.SecurityToken,
	}

	c.accessLock.RUnlock()

	req = Sign(req, credential)

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
