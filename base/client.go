package base

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	accessKey = "VOLCACCESSKEY"
	secretKey = "VOLCSECRETKEY"

	defaultScheme = "http"
)

// Client 基础客户端
type Client struct {
	Client      http.Client
	ServiceInfo *ServiceInfo
	ApiInfoList map[string]*ApiInfo
}

// NewClient 生成一个客户端
func NewClient(info *ServiceInfo, apiInfoList map[string]*ApiInfo) *Client {
	transport := &http.Transport{
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     10 * time.Second,
	}

	c := http.Client{Transport: transport}
	client := &Client{Client: c, ServiceInfo: info, ApiInfoList: apiInfoList}

	if client.ServiceInfo.Scheme == "" {
		client.ServiceInfo.Scheme = defaultScheme
	}

	if os.Getenv(accessKey) != "" && os.Getenv(secretKey) != "" {
		client.ServiceInfo.Credentials.AccessKeyID = os.Getenv(accessKey)
		client.ServiceInfo.Credentials.SecretAccessKey = os.Getenv(secretKey)
	} else if _, err := os.Stat(os.Getenv("HOME") + "/.volcconfig/config"); err == nil {
		if content, err := ioutil.ReadFile(os.Getenv("HOME") + "/.volcconfig/config"); err == nil {
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

// SetAccessKey 设置AK
func (client *Client) SetAccessKey(ak string) {
	if ak != "" {
		client.ServiceInfo.Credentials.AccessKeyID = ak
	}
}

// SetSecretKey 设置SK
func (client *Client) SetSecretKey(sk string) {
	if sk != "" {
		client.ServiceInfo.Credentials.SecretAccessKey = sk
	}
}

// SetHost 设置Host
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

// SetCredential 设置Credentials
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
}

// GetSignUrl 获取签名字符串
func (client *Client) GetSignUrl(api string, query url.Values) (string, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return "", errors.New("相关api不存在")
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
		return "", errors.New("构建request失败")
	}

	return client.ServiceInfo.Credentials.SignUrl(req), nil
}

// SignSts2 生成sts信息
func (client *Client) SignSts2(inlinePolicy *Policy, expire time.Duration) (*SecurityToken2, error) {
	var err error
	sts := new(SecurityToken2)
	if sts.AccessKeyId, sts.SecretAccessKey, err = createTempAKSK(); err != nil {
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

// Query 发起Get的query请求
func (client *Client) Query(api string, query url.Values) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}

	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), nil)
	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}
	req.Header = header
	return client.makeRequest(api, req, timeout)
}

// Json 发起Json的post请求
func (client *Client) Json(api string, query url.Values, body string) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), strings.NewReader(body))
	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}
	req.Header = header
	req.Header.Set("Content-Type", "application/json")

	return client.makeRequest(api, req, timeout)
}

// Post 发起Post请求
func (client *Client) Post(api string, query url.Values, form url.Values) ([]byte, int, error) {
	apiInfo := client.ApiInfoList[api]

	if apiInfo == nil {
		return []byte(""), 500, errors.New("相关api不存在")
	}
	timeout := getTimeout(client.ServiceInfo.Timeout, apiInfo.Timeout)
	header := mergeHeader(client.ServiceInfo.Header, apiInfo.Header)
	query = mergeQuery(query, apiInfo.Query)
	form = mergeQuery(form, apiInfo.Form)

	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     apiInfo.Path,
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest(strings.ToUpper(apiInfo.Method), u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return []byte(""), 500, errors.New("构建request失败")
	}
	req.Header = header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return client.makeRequest(api, req, timeout)
}

func (client *Client) makeRequest(api string, req *http.Request, timeout time.Duration) ([]byte, int, error) {
	req = client.ServiceInfo.Credentials.Sign(req)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := client.Client.Do(req)
	if err != nil {
		return []byte(""), 500, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, resp.StatusCode, fmt.Errorf("api %s http code %d body %s", api, resp.StatusCode, string(body))
	}

	return body, resp.StatusCode, nil
}
