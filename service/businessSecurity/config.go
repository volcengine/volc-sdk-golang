package businessSecurity

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type BusinessSecurity struct {
	*base.Client
	retry bool
}

var DefaultInstance = NewInstance()

func NewInstance() *BusinessSecurity {
	instance := &BusinessSecurity{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		retry:  true,
	}
	return instance
}

func (p *BusinessSecurity) Retry() bool {
	return p.retry
}

func (p *BusinessSecurity) CloseRetry() {
	p.retry = false
}

func (p *BusinessSecurity) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	p.SetScheme("http")
	return nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "riskcontrol.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "BusinessSecurity"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "open-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "BusinessSecurity"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "open-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "BusinessSecurity"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"RiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskDetection"},
				"Version": []string{"2021-02-02"},
			},
		},
		"AsyncRiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncRiskDetection"},
				"Version": []string{"2021-02-25"},
			},
		},
		"RiskResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskResult"},
				"Version": []string{"2021-03-10"},
			},
		},
		"DataReport": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DataReport"},
				"Version": []string{"2021-08-31"},
			},
		},
		"TextRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextRisk"},
				"Version": []string{"2022-01-26"},
			},
		},
		"AsyncVideoRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncVideoRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"VideoResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoResult"},
				"Version": []string{"2021-11-29"},
			},
		},
		"AsyncImageRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncImageRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"AsyncImageRiskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncImageRisk"},
				"Version": []string{"2022-08-26"},
			},
		},
		"ImageContentRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageContentRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"GetImageResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageResult"},
				"Version": []string{"2021-11-29"},
			},
		},
		"GetImageResultV2": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageResult"},
				"Version": []string{"2022-08-26"},
			},
		},
		"AsyncAudioRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncAudioRisk"},
				"Version": []string{"2022-04-01"},
			},
		},
		"GetAudioResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioResult"},
				"Version": []string{"2022-04-01"},
			},
		},
		"AsyncLiveVideoRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncLiveVideoRisk"},
				"Version": []string{"2022-04-25"},
			},
		},
		"GetVideoLiveResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVideoLiveResult"},
				"Version": []string{"2022-04-25"},
			},
		},
		"CloseVideoLiveRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CloseVideoLive"},
				"Version": []string{"2022-04-25"},
			},
		},
		"AsyncLiveAudioRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncLiveAudioRisk"},
				"Version": []string{"2022-04-25"},
			},
		},
		"GetAudioLiveResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioLiveResult"},
				"Version": []string{"2022-04-25"},
			},
		},
		"CloseAudioLiveRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CloseAudioLive"},
				"Version": []string{"2022-04-25"},
			},
		},
		"EnableCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"DisableCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"CreateCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCustomContents"},
				"Version": []string{"2022-01-22"},
			},
		},
		"UploadCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadCustomContents"},
				"Version": []string{"2022-02-07"},
			},
		},
		"DeleteCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"ElementVerify": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerify"},
				"Version": []string{"2021-11-23"},
			},
		},
		"MobileStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MobileStatus"},
				"Version": []string{"2020-12-25"},
			},
		},
		"ElementVerifyV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerify"},
				"Version": []string{"2022-04-13"},
			},
		},
		"ElementVerifyEncrypted": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerifyEncrypted"},
				"Version": []string{"2022-11-24"},
			},
		},
		"MobileStatusV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MobileStatus"},
				"Version": []string{"2022-04-13"},
			},
		},
		"TextSliceRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextSliceRisk"},
				"Version": []string{"2022-11-07"},
			},
		},
	}
)
