package live

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
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
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCallback"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCallback"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCallback"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomain"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomain"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListDomainDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomainDetail"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDomain"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"EnableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableDomain"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DisableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableDomain"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ManagerPullPushDomainBind": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ManagerPullPushDomainBind"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAuthKey"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAuth"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ForbidStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ForbidStream"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ResumeStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeStream"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"BindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UnbindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCert"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateReferer"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteReferer"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeReferer"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRecordPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecordPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRecordPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListVhostRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostRecordPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTranscodePreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTranscodePreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodePreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListVhostTransCodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostTransCodePreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListVhostSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotPreset"},
				"Version": []string{ServiceVersion20200801},
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
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListVhostSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotAuditPreset"},
				"Version": []string{ServiceVersion20200801},
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
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWatermarkPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWatermarkPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWatermarkPreset"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeCDNSnapshotHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCDNSnapshotHistory"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeRecordTaskFileHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRecordTaskFileHistory"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeLiveStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamInfoByPage"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"KillStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KillStream"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeClosedStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeClosedStreamInfoByPage"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeLiveStreamState": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamState"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeForbiddenStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeForbiddenStreamInfoByPage"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySourceV2"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySourceV2"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeRelaySourceV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRelaySourceV2"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreateVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVQScoreTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVQScoreTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVQScoreTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"GeneratePlayURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePlayURL"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"GeneratePushURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePushURL"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"CreatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"ListPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"StopPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"RestartPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RestartPullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeletePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePullToPushTask"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"UpdateDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenyConfig"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DescribeDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDenyConfig"},
				"Version": []string{ServiceVersion20200801},
			},
		},
		"DeleteDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDenyConfig"},
				"Version": []string{ServiceVersion20200801},
			},
		},
	}
)

// DefaultInstance
var DefaultInstance = NewInstance()

// Live .
type Live struct {
	*base.Client
}

// NewInstance create a instance
func NewInstance() *Live {
	instance := &Live{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Live) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

// GetAPIInfo interface
func (p *Live) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetHost .
func (p *Live) SetRegion(region string) {
	ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Live) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Live) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
