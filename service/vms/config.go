package vms

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultTimeout = 5 * time.Second
	DefaultVersion = "2022-01-01"
	ServiceName    = "vms"
)

type Vms struct {
	*base.Client
}

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

func NewInstance() *Vms {
	instance := &Vms{
		Client: base.NewClient(ServiceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: DefaultTimeout,
		Host:    "cloud-vms.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Region:  base.RegionCnNorth1,
			Service: ServiceName,
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"CreateNumberApplication": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateNumberApplication"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXB"},
				"Version": []string{DefaultVersion},
			},
		},
		"SelectNumberAndBindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SelectNumberAndBindAXB"},
				"Version": []string{DefaultVersion},
			},
		},
		"UnbindAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAXB"},
				"Version": []string{DefaultVersion},
			},
		},
		"QuerySubscription": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QuerySubscription"},
				"Version": []string{DefaultVersion},
			},
		},
		"QuerySubscriptionForList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QuerySubscriptionForList"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpgradeAXToAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpgradeAXToAXB"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateAXB": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAXB"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXN"},
				"Version": []string{DefaultVersion},
			},
		},
		"SelectNumberAndBindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SelectNumberAndBindAXN"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAXN"},
				"Version": []string{DefaultVersion},
			},
		},
		"UnbindAXN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAXN"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindAXNE": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXNE"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindAXBForAXNE": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAXBForAXNE"},
				"Version": []string{DefaultVersion},
			},
		},
		"UnbindAXNE": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAXNE"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateAXNE": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAXNE"},
				"Version": []string{DefaultVersion},
			},
		},
		"Click2Call": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Click2Call"},
				"Version": []string{DefaultVersion},
			},
		},
		"CancelClick2Call": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CancelClick2Call"},
				"Version": []string{DefaultVersion},
			},
		},
		"Click2CallLite": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Click2CallLite"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindAxyb": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindAxyb"},
				"Version": []string{DefaultVersion},
			},
		},
		"BindYbForAxyb": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindYbForAxyb"},
				"Version": []string{DefaultVersion},
			},
		},
		"UnbindAxyb": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindAxyb"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateAxyb": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAxyb"},
				"Version": []string{DefaultVersion},
			},
		},
		"CreateNumberPool": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateNumberPool"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateNumberPool": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateNumberPool"},
				"Version": []string{DefaultVersion},
			},
		},
		"NumberPoolList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"NumberPoolList"},
				"Version": []string{DefaultVersion},
			},
		},
		"NumberList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"NumberList"},
				"Version": []string{DefaultVersion},
			},
		},
		"EnableOrDisableNumber": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableOrDisableNumber"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryNumberApplyRecordList": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryNumberApplyRecordList"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryAudioRecordFileUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryAudioRecordFileUrl"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryAudioRecordToTextFileUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryAudioRecordToTextFileUrl"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryCallRecordMsg": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryCallRecordMsg"},
				"Version": []string{"2021-01-01"},
			},
		},
		"QueryCallRecordMsgNew": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryCallRecordMsg"},
				"Version": []string{DefaultVersion},
			},
		},
		"CreateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTask"},
				"Version": []string{DefaultVersion},
			},
		},
		"BatchAppend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchAppend"},
				"Version": []string{DefaultVersion},
			},
		},
		"PauseTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PauseTask"},
				"Version": []string{DefaultVersion},
			},
		},
		"ResumeTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeTask"},
				"Version": []string{DefaultVersion},
			},
		},

		// policy
		"StopTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopTask"},
				"Version": []string{DefaultVersion},
			},
		},
		"UpdateTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTask"},
				"Version": []string{DefaultVersion},
			},
		},
		"SingleBatchAppend": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleBatchAppend"},
				"Version": []string{DefaultVersion},
			},
		},
		"SingleInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleInfo"},
				"Version": []string{DefaultVersion},
			},
		},
		"SingleCancel": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SingleCancel"},
				"Version": []string{DefaultVersion},
			},
		},
		"FetchResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FetchResource"},
				"Version": []string{DefaultVersion},
			},
		},
		"OpenCreateTts": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenCreateTts"},
				"Version": []string{DefaultVersion},
			},
		},

		// role
		"OpenDeleteResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenDeleteResource"},
				"Version": []string{DefaultVersion},
			},
		},
		"GetResourceUploadUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourceUploadUrl"},
				"Version": []string{DefaultVersion},
			},
		},
		"CommitResourceUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitResourceUpload"},
				"Version": []string{DefaultVersion},
			},
		},
		"OpenUpdateResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OpenUpdateResource"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryUsableResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryUsableResource"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryOpenGetResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryOpenGetResource"},
				"Version": []string{DefaultVersion},
			},
		},
		"AddQualification": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddQualification"},
				"Version": []string{DefaultVersion},
			},
		},

		"UpdateQualification": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateQualification"},
				"Version": []string{DefaultVersion},
			},
		},

		"AddQualificationScene": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddQualificationScene"},
				"Version": []string{DefaultVersion},
			},
		},

		"UpdateQualificationScene": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateQualificationScene"},
				"Version": []string{DefaultVersion},
			},
		},

		"QueryQualification": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryQualification"},
				"Version": []string{DefaultVersion},
			},
		},

		"UploadQualificationFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadQualificationFile"},
				"Version": []string{DefaultVersion},
			},
		},
		"QueryCanCall": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"QueryCanCall"},
				"Version": []string{DefaultVersion},
			},
		},
	}
)
