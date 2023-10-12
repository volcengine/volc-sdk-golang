package live

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20230101 = "2023-01-01"
	ServiceVersion20200801 = "2020-08-01"
	ServiceVersion20180101 = "2018-01-01"
	ServiceName            = "live"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "live.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"ListCommonTransPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCommonTransPresetDetail"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCallback"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCallback"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCallback"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomain"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomain"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListDomainDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomainDetail"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDomain"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"EnableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableDomain"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DisableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableDomain"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ManagerPullPushDomainBind": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ManagerPullPushDomainBind"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAuthKey"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAuth"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ForbidStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ForbidStream"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ResumeStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeStream"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"BindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UnbindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCert"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateReferer"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteReferer"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeReferer"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRecordPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecordPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRecordPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListVhostRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostRecordPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTranscodePreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTranscodePreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodePreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListVhostTransCodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostTransCodePreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListVhostSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribePullToPushBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribePullToPushBandwidthData"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateCaster": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCaster"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListCasters": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCasters"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetCasterInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCasterInfo"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"StartCaster": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartCaster"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"StopCaster": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopCaster"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateCasterResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCasterResource"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"RemoveCasterResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RemoveCasterResource"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateCasterLayout": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCasterLayout"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteCasterLayout": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCasterLayout"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateCasterResourceImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCasterResourceImage"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteCasterResourceImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCasterResourceImage"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateCasterResourceImages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCasterResourceImages"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateCasterImageCollection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCasterImageCollection"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateCasterConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCasterConfig"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateCasterResourceOPED": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCasterResourceOPED"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteCasterResourceOPED": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCasterResourceOPED"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListVhostSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeLiveAuditData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveAuditData"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"SwitchCasterResource": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SwitchCasterResource"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"SwitchCasterResourceImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SwitchCasterResourceImage"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"SwitchCasterResourceOPED": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SwitchCasterResourceOPED"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"StartCasterOPEDArrange": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartCasterOPEDArrange"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"SwitchCasterLayout": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SwitchCasterLayout"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CopyCasterPVWToPGM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CopyCasterPVWToPGM"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"GetCasterResourceVodInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCasterResourceVodInfo"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"CreateCasterArrange": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCasterArrange"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"UpdateCasterArrange": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCasterArrange"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"DeleteCasterArrange": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCasterArrange"},
				"Version": []string{ServiceVersion20180101},
			},
		},
		"ListVhostWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostWatermarkPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWatermarkPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWatermarkPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWatermarkPreset"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeCDNSnapshotHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCDNSnapshotHistory"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeRecordTaskFileHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRecordTaskFileHistory"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeLiveStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamInfoByPage"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"KillStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KillStream"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeClosedStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeClosedStreamInfoByPage"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeLiveStreamState": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamState"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeForbiddenStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeForbiddenStreamInfoByPage"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySourceV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySourceV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRelaySourceV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreateVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVQScoreTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVQScoreTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVQScoreTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"GeneratePlayURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePlayURL"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"GeneratePushURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePushURL"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"CreatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"ListPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"StopPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"RestartPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RestartPullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeletePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePullToPushTask"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenyConfig"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDenyConfig"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDenyConfig"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"UpdateDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenyConfigV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDenyConfigV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DeleteDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDenyConfigV2"},
				"Version": []string{ServiceVersion20230101},
			},
		},
		"DescribeLiveMetricTrafficData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveMetricTrafficData"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeLiveMetricBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveMetricBandwidthData"},
				"Version": []string{ServiceVersion20200801},
			},
		},
	}
)

// DefaultInstance
var DefaultInstance = NewInstance()

// LIVE .
type LIVE struct {
	*base.Client
}

// NewInstance create a instance
func NewInstance() *LIVE {
	instance := &LIVE{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *LIVE) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *LIVE) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *LIVE) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *LIVE) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *LIVE) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
