package nlp

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "nlp_gateway"
	ApiVersion1    = "2020-09-01"
	ApiVersion2    = "2020-12-01"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 10 * time.Second,
			Scheme:  "https",
			Host:    "open.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ServiceName,
			},
		},
	}
	ApiInfoList = map[string]*base.ApiInfo{
		// 中文文本纠错
		"TextCorrectionZhCorrect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextCorrectionZhCorrect"},
				"Version": []string{ApiVersion1},
			},
		},
		//	英文文本纠错
		"TextCorrectionEnCorrect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextCorrectionEnCorrect"},
				"Version": []string{ApiVersion1},
			},
		},
		//	情感分析
		"SentimentAnalysis": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SentimentAnalysis"},
				"Version": []string{ApiVersion2},
			},
		},
		//  文本摘要
		"TextSummarization": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextSummarization"},
				"Version": []string{ApiVersion2},
			},
		},
		//	 关键词提取
		"KeyphraseExtractionExtract": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KeyphraseExtractionExtract"},
				"Version": []string{ApiVersion1},
			},
		},
	}
)

// DefaultInstance 默认实例，Region: cn-north-1
var DefaultInstance = NewInstance()

type Nlp struct {
	*base.Client
}

func NewInstance() *Nlp {
	instance := &Nlp{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	return instance
}
