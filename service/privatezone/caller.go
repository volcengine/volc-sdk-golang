package privatezone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/volcengine/volc-sdk-golang/base"
)

func NewDefaultServiceInfo() *base.Client {
	return base.NewClient(ServiceInfo, nil)
}

type VolcCaller struct {
	Volc *base.Client
}

// GetServiceInfo interface
func (p *VolcCaller) GetServiceInfo() *base.ServiceInfo {
	return p.Volc.ServiceInfo
}

func (p *VolcCaller) SetService(service string) {
	p.Volc.ServiceInfo.Credentials.Service = service
}

func (p *VolcCaller) SetRegion(region string) {
	p.Volc.ServiceInfo.Credentials.Region = region
}

func (p *VolcCaller) SetHost(host string) {
	p.Volc.ServiceInfo.Host = host
}

func (p *VolcCaller) SetSchema(schema string) {
	p.Volc.ServiceInfo.Scheme = schema
}

func NewVolcCaller() *VolcCaller {
	return &VolcCaller{Volc: NewDefaultServiceInfo()}
}

func InitDNSVolcClient() *Client {
	vc := NewVolcCaller()
	vc.Volc.SetSessionToken("")
	vc.Volc.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	vc.Volc.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))
	return NewClient(vc)
}

func (c *VolcCaller) Do(r *http.Request) (*http.Response, error) {
	r.URL.Host = c.Volc.ServiceInfo.Host
	r.URL.Scheme = c.Volc.ServiceInfo.Scheme
	r.Host = c.Volc.ServiceInfo.Host

	r.Header.Add("Content-Type", "application/json")
	q := r.URL.Query()
	q.Add("Version", ServiceVersion)
	r.URL.RawQuery = q.Encode()

	r = c.Volc.ServiceInfo.Credentials.Sign(r)

	ctx, cancel := context.WithTimeout(r.Context(), c.Volc.ServiceInfo.Timeout)
	defer cancel()
	r = r.WithContext(ctx)

	resp, err := c.Volc.Client.Do(r)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	var payload TopResponse
	err = json.NewDecoder(resp.Body).Decode(&payload)
	resp.Body.Close()

	newResp := &http.Response{}

	if err != nil {
		return newResp, errors.WithStack(err)
	}

	if payload.ResponseMetadata.Error.Code != "" {
		return newResp, NewTOPError(&payload.ResponseMetadata)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return newResp, errors.Wrap(errors.New("http error"), strconv.Itoa(resp.StatusCode))
	}

	str, err := json.Marshal(payload.Result)
	if err != nil {
		return newResp, errors.WithStack(err)
	}

	newResp = &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(str))}

	return newResp, errors.WithStack(err)
}

type TOPError struct {
	Code      string `form:"Code" json:"Code"`
	CodeN     int64  `form:"CodeN" json:"CodeN"`
	Message   string `form:"Message" json:"Message"`
	RequestID string `form:"RequestId" json:"RequestId"`
}

func NewTOPError(respMeta *TopRespMeta) *TOPError {
	return &TOPError{
		Code:      respMeta.Error.Code,
		CodeN:     respMeta.Error.CodeN,
		Message:   respMeta.Error.Message,
		RequestID: respMeta.RequestID,
	}
}

func (err *TOPError) GetCode() int64 {
	return err.CodeN
}

func (err *TOPError) Error() string {
	return fmt.Sprintf(
		"TOP DNS response failed with code %d: %s, %s, requestID: %s",
		err.CodeN, err.Code, err.Message, err.RequestID,
	)
}

type TopResponse struct {
	ResponseMetadata TopRespMeta     `form:"ResponseMetadata" json:"ResponseMetadata"`
	Result           json.RawMessage `form:"Result" json:"Result"`
}

type TopRespMeta struct {
	Action    string       `form:"Action" json:"Action"`
	Error     TopRespError `form:"Error" json:"Error"`
	Region    string       `form:"Region" json:"Region"`
	RequestID string       `form:"RequestId" json:"RequestId"`
	Service   string       `form:"Service" json:"Service"`
	Version   string       `form:"Version" json:"Version"`
}

type TopRespError struct {
	Code    string `form:"Code" json:"Code"`
	CodeN   int64  `form:"CodeN" json:"CodeN"`
	Message string `form:"Message" json:"Message"`
}
