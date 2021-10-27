package gtm

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
	"github.com/volcengine/volc-sdk-golang/service/gtm/common"
)

type Conf struct {
	Host    string
	Name    string
	Region  string
	Timeout int
	Version string
}

func InitCommonParameter() *Conf {
	data := Conf{
		Host:    common.ServiceInfo.Host,
		Name:    common.ServiceName,
		Region:  common.DefaultRegion,
		Timeout: common.Timeout,
		Version: common.ServiceVersion,
	}
	return &data
}

func NewDefaultServiceInfo() *base.Client {
	return base.NewClient(common.ServiceInfo, nil)
}

type VolcCaller struct {
	Volc *base.Client
}

func NewVolcCaller() *VolcCaller {
	instance := &VolcCaller{Volc: NewDefaultServiceInfo()}
	instance.Volc.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	instance.Volc.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))
	instance.Volc.SetHost(instance.Volc.ServiceInfo.Host)
	instance.Volc.SetScheme(instance.Volc.ServiceInfo.Scheme)
	instance.Volc.SetTimeout(instance.Volc.ServiceInfo.Timeout)

	return instance
}

func (c *VolcCaller) Do(r *http.Request) (*http.Response, error) {
	r.URL.Host = c.Volc.ServiceInfo.Host
	r.URL.Scheme = c.Volc.ServiceInfo.Scheme
	r.Host = c.Volc.ServiceInfo.Host

	r = c.Volc.ServiceInfo.Credentials.Sign(r)

	ctx, cancel := context.WithTimeout(r.Context(), c.Volc.ServiceInfo.Timeout)
	defer cancel()
	r = r.WithContext(ctx)

	resp, err := c.Volc.Client.Do(r)
	if err != nil {
		return resp, errors.WithStack(err)
	}

	var payload common.TopResponse
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

	return newResp, nil
}

type TOPError struct {
	Code      string `form:"Code" json:"Code"`
	CodeN     int64  `form:"CodeN" json:"CodeN"`
	Message   string `form:"Message" json:"Message"`
	RequestID string `form:"RequestId" json:"RequestId"`
}

func NewTOPError(respMeta *common.TopRespMeta) *TOPError {
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
		"TOP GTM response failed with code %d: %s, %s, requestID: %s",
		err.CodeN, err.Code, err.Message, err.RequestID,
	)
}
