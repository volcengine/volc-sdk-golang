package rtc_v20231101

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
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"LimitTokenPrivilege": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LimitTokenPrivilege"},
				"Version": []string{"2023-11-01"},
			},
		},
		"BanRoomUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanRoomUser"},
				"Version": []string{"2023-11-01"},
			},
		},
		"BanUserStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanUserStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdateBanRoomUserRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateBanRoomUserRule"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetRoomOnlineUsers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRoomOnlineUsers"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UnbanUserStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbanUserStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopPushPublicStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPushPublicStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRelayStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSnapshot"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartPushPublicStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushPublicStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartWebcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWebcast"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSnapshot"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartSegment"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRecord"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartPushSingleStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushSingleStreamToCDN"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartPushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartPushMixedStreamToCDN"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartRelayStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdateRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecord"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdatePublicStreamParam": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePublicStreamParam"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdatePushMixedStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePushMixedStreamToCDN"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdateRelayStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelayStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdateSnapshot": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshot"},
				"Version": []string{"2023-11-01"},
			},
		},
		"UpdateSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSegment"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetWebCastTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetWebCastTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetPushSingleStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushSingleStreamToCDNTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetPushMixedStreamToCDNTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPushMixedStreamToCDNTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetRecordTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecordTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetSnapshotTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSnapshotTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListRelayStream": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRelayStream"},
				"Version": []string{"2023-11-01"},
			},
		},
		"GetSegmentTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSegmentTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopWebcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWebcast"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopRecord"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopPushStreamToCDN": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPushStreamToCDN"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopSegment"},
				"Version": []string{"2023-11-01"},
			},
		},
		"CreateApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateApp"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ModifyAppStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyAppStatus"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListApps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListApps"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartWBRecord"},
				"Version": []string{"2023-11-01"},
			},
		},
		"WbTranscodeCreate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeCreate"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopWBRecord": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopWBRecord"},
				"Version": []string{"2023-11-01"},
			},
		},
		"WbTranscodeQuery": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeQuery"},
				"Version": []string{"2023-11-01"},
			},
		},
		"WbTranscodeGet": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"WbTranscodeGet"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListRealTimeQualityDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeQualityDistribution"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListRealTimeOperationData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeOperationData"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListQualityDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListQualityDistribution"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListQuality"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListOperationDistribution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListOperationDistribution"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListOperationData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListOperationData"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListRealTimeQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRealTimeQuality"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListCallDetail": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCallDetail"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListRoomInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRoomInfo"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListUserInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUserInfo"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StopDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopDetection"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListDetectionTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDetectionTask"},
				"Version": []string{"2023-11-01"},
			},
		},
		"StartDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartDetection"},
				"Version": []string{"2023-11-01"},
			},
		},
		"ListUsages": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListUsages"},
				"Version": []string{"2023-11-01"},
			},
		},
	}
)
