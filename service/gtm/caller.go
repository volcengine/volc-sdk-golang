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

type VolcCaller struct {
	Volc *base.Client
}

func NewVolcCaller() *VolcCaller {
	instance := &VolcCaller{}
	instance.Volc = base.NewClient(common.ServiceInfo, nil)
	instance.Volc.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	instance.Volc.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))
	instance.Volc.SetHost(common.ServiceInfo.Host)
	instance.Volc.SetScheme(common.ServiceInfo.Scheme)
	instance.Volc.SetTimeout(common.ServiceInfo.Timeout)

	return instance
}

func (c *VolcCaller) Do(r *http.Request) (*http.Response, error) {
	r.URL.Host = common.ServiceInfo.Host
	r.URL.Scheme = common.ServiceInfo.Scheme
	r.Host = common.ServiceInfo.Host

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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
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
