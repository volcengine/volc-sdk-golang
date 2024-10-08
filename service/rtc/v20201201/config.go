package rtc_v20201201

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "rtc"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
		"ap-singapore-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "open-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "ap-singapore-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"BanRoomUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanRoomUser"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BanUserStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanUserStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateBanRoomUserRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateBanRoomUserRule"},
				"Version": []string{"2020-12-01"},
			},
		},
		"KickUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KickUser"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UnbanUserStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbanUserStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DismissRoom": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DismissRoom"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopPushPublicStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPushPublicStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRelayStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSnapshot"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartPushPublicStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushPublicStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartWebcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWebcast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSnapshot"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSegment"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartPushSingleStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushSingleStreamToCDN"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartPushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushMixedStreamToCDN"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRelayStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdatePublicStreamParam": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePublicStreamParam"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdatePushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePushMixedStreamToCDN"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelayStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshot"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSegment"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetWebCastTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWebCastTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetPushSingleStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushSingleStreamToCDNTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetPushMixedStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushMixedStreamToCDNTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetRecordTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecordTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListRelayStream": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRelayStream"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopWebcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWebcast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopPushStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPushStreamToCDN"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSegment"},
				"Version": []string{"2020-12-01"},
			},
		},
		"SendBroadcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendBroadcast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"SendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendRoomUnicast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"SendUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendUnicast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"BatchSendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchSendRoomUnicast"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateApp"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyAppStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyAppStatus"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListApps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListApps"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ModifyBusinessRemarks": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyBusinessRemarks"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteBusinessID": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteBusinessID"},
				"Version": []string{"2020-12-01"},
			},
		},
		"AddBusinessID": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddBusinessID"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetAllBusinessID": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllBusinessID"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWBRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeCreate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeCreate"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeQuery": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeQuery"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWBRecord"},
				"Version": []string{"2020-12-01"},
			},
		},
		"WbTranscodeGet": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeGet"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListRealTimeQualityDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeQualityDistribution"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListRealTimeQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeQuality"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListRealTimeOperationData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeOperationData"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListQualityDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListQualityDistribution"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListUserInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUserInfo"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListQuality"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListOperationDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListOperationDistribution"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListOperationData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListOperationData"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListRoomInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRoomInfo"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListCallDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCallDetail"},
				"Version": []string{"2020-12-01"},
			},
		},
		"SearchMusics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SearchMusics"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListMusics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListMusics"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListHotMusic": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHotMusic"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListDetectionTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDetectionTask"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StopDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopDetection"},
				"Version": []string{"2020-12-01"},
			},
		},
		"StartDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartDetection"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateFailRecoveryPolicy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateFailRecoveryPolicy"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateVendorPolicy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVendorPolicy"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteFailRecoveryPolicy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteFailRecoveryPolicy"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetFailRecoveryPolicies": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetFailRecoveryPolicies"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateFailRecoveryPolicy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFailRecoveryPolicy"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateVendorPolicy": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateVendorPolicy"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListUsages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsages"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateByteplusApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateByteplusApp"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetResourcePackNum": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResourcePackNum"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListResourcePackages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourcePackages"},
				"Version": []string{"2020-12-01"},
			},
		},
		"ListResourcePackagesV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourcePackagesV2"},
				"Version": []string{"2020-12-01"},
			},
		},
		"CreateCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCallback"},
				"Version": []string{"2020-12-01"},
			},
		},
		"DeleteCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCallback"},
				"Version": []string{"2020-12-01"},
			},
		},
		"GetCallback": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCallback"},
				"Version": []string{"2020-12-01"},
			},
		},
		"UpdateCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCallback"},
				"Version": []string{"2020-12-01"},
			},
		},
	}
)
