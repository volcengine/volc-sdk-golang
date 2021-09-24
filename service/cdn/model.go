package cdn

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
)

type QueryOption struct {
	Key   string
	Value string
}

type CDNError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e CDNError) Error() string {
	return fmt.Sprintf("status: %d, code: %s, message: %s", e.Status, e.Code, e.Message)
}

type SubmitRefreshTaskParam struct {
	Type string
	Urls []string
}

type SubmitRefreshTaskRequest struct {
	Type string `json:"Type,omitempty"`
	Urls string `json:"Urls"`
}

type SubmitRefreshTaskResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		TaskID string `json:"TaskID"`
	} `json:",omitempty"`
}

type SubmitPreloadTaskParam struct {
	Urls []string
}

type SubmitPreloadTaskRequest struct {
	Urls string `json:"Urls"`
}

type SubmitPreloadTaskResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		TaskID string `json:"TaskID"`
	} `json:"Result,omitempty"`
}

type DescribeContentTasksParam struct {
	Url        string   `json:"Url,omitempty"`
	DomainName string   `json:"DomainName,omitempty"`
	TaskID     string   `json:"TaskID,omitempty"`
	TaskType   string   `json:"TaskType,omitempty"`
	Status     []string `json:"Status,omitempty"`
	StartTime  int64    `json:"StartTime,omitempty"`
	EndTime    int64    `json:"EndTime,omitempty"`
	PageNum    int64    `json:"PageNum,omitempty"`
	PageSize   int64    `json:"PageSize,omitempty"`
}

type ContentTask struct {
	Url       string `json:"Url"`
	Status    string `json:"Status"`
	TaskType  string `json:"TaskType"`
	CreatedAt int64  `json:"CreatedAt"`
	TaskID    string `json:"TaskID"`
}

type DescribeContentTasksResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Total    int64         `json:"Total"`
		PageNum  int64         `json:"PageNum"`
		PageSize int64         `json:"PageSize"`
		Data     []ContentTask `json:"Data"`
	} `json:"Result,omitempty"`
}

type DescribeContentQuotaResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		RefreshQuota     int64 `json:"RefreshQuota"`
		RefreshDirQuota  int64 `json:"RefreshDirQuota"`
		PreloadQuota     int64 `json:"PreloadQuota"`
		RefreshRemain    int64 `json:"RefreshRemain"`
		RefreshDirRemain int64 `json:"RefreshDirRemain"`
		PreloadRemain    int64 `json:"PreloadRemain"`
	} `json:"Result,omitempty"`
}

type TimeSeriesData struct {
	Timestamp int64   `json:"Timestamp"`
	Value     float64 `json:"Value"`
}

type MetricStatData struct {
	Metric string           `json:"Metric"`
	Values []TimeSeriesData `json:"Values"`
}

type ResourceStatData struct {
	Name    string           `json:"Name"`
	Metrics []MetricStatData `json:"Metrics"`
}

type DescribeCdnDataParam struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Metric    string `json:"Metric"`
	Domain    string `json:"Domain,omitempty"`
	Interval  string `json:"Interval,omitempty"`
	Isp       string `json:"Isp,omitempty"`
	Region    string `json:"Region,omitempty"`
	Protocol  string `json:"Protocol,omitempty"`
	IpVersion string `json:"IpVersion,omitempty"`
}

type DescribeCdnDataResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Resources []ResourceStatData `json:"Resources"`
	} `json:"Result,omitempty"`
}

type DescribeCdnOriginDataParam struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Metric    string `json:"Metric"`
	Domain    string `json:"Domain,omitempty"`
	Interval  string `json:"Interval,omitempty"`
}

type DescribeCdnOriginDataResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Resources []ResourceStatData `json:"Resources"`
	} `json:"Result,omitempty"`
}

type NamePair struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

type DescribeCdnRegionAndIspResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Isps    []NamePair `json:"Isps"`
		Regions []NamePair `json:"Regions"`
	} `json:"Result,omitempty"`
}

type DescribeCdnDomainTopDataParam struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Metric    string `json:"Metric"`
	Domain    string `json:"Domain"`
	Item      string `json:"Item"`
}

type TopDataDetail struct {
	Item      string  `json:"Item"`
	PV        int64   `json:"PV"`
	PVRatio   float64 `json:"PVRatio"`
	Flux      int64   `json:"Flux"`
	FluxRatio float64 `json:"FluxRatio"`
	Bandwidth float64 `json:"Bandwidth"`
}

type DescribeCdnDomainTopDataResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Domain         string          `json:"Domain"`
		TopDataDetails []TopDataDetail `json:"TopDataDetails"`
	} `json:"Result,omitempty"`
}

type DescribeCdnAccountingDataParam struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Domain    string `json:"Domain"`
}

type DomainAccountingResult struct {
	Domain                  string `json:"Domain"`
	DomainAccountingDetails struct {
		DomainAccountingDetail []DomainAccountingDetail `json:"DomainAccountingDetail"`
	} `json:"DomainAccountingDetails"`
}

type DomainAccountingDetail struct {
	TimeStamp int64   `json:"TimeStamp"`
	Flux      float64 `json:"Flux"`
	Bandwidth float64 `json:"Bandwidth"`
}

type DescribeCdnAccountingDataResponse struct {
	ResponseMetadata *base.ResponseMetadata   `json:"ResponseMetadata"`
	Result           []DomainAccountingResult `json:"Result,omitempty"`
}

type DescribeCdnAccessLogParam struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Domain    string `json:"Domain"`
	PageNum   int64  `json:"PageNum,omitempty"`
	PageSize  int64  `json:"PageSize,omitempty"`
}

type DomainLogDetail struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	LogName   string `json:"LogName"`
	LogPath   string `json:"LogPath"`
	LogSize   int64  `json:"LogSize"`
}

type DescribeCdnAccessLogResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		Domain           string            `json:"Domain"`
		PageSize         int64             `json:"PageSize"`
		PageNum          int64             `json:"PageNum"`
		TotalCount       int64             `json:"TotalCount"`
		DomainLogDetails []DomainLogDetail `json:"DomainLogDetails"`
	} `json:"Result"`
}

type StartCdnDomainParam struct {
	DomainName string `json:"DomainName"`
}

type DeleteCdnDomainParam struct {
	DomainName string `json:"DomainName"`
}

type StopCdnDomainParam struct {
	DomainName string `json:"DomainName"`
}

type ListCdnDomainsParam struct {
	Domain      string `json:"Domain,omitempty"`
	ServiceType string `json:"ServiceType,omitempty"`
	ResourceTag string `json:"ResourceTag,omitempty"`
	PageNum     int64  `json:"PageNum,omitempty"`
	PageSize    int64  `json:"PageSize,omitempty"`
}

type DomainSummary struct {
	Domain        string             `json:"Domain"`
	ServiceType   string             `json:"ServiceType"`
	Status        string             `json:"Status"`
	Cname         string             `json:"Cname"`
	ResourceTags  []ResourceTagEntry `json:"ResourceTags"`
	ServiceRegion string             `json:"ServiceRegion"`
	CreateTime    int64              `json:"CreateTime"`
	UpdateTime    int64              `json:"UpdateTime"`
}

type ResourceTagEntry struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type ListCdnDomainsResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           struct {
		PageNum  int64           `json:"PageNum"`
		PageSize int64           `json:"PageSize"`
		Total    int64           `json:"Total"`
		Data     []DomainSummary `json:"Data"`
	} `json:"Result"`
}

type EmptyResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
}

type DescribeCdnUpperIpParam struct {
	Domain    string `json:"Domain"`
	IpVersion string `json:"IpVersion,omitempty"`
}

type DescribeCdnUpperIpResponse struct {
	ResponseMetadata *base.ResponseMetadata `json:"ResponseMetadata"`
	Result           DescribeCdnIpResult    `json:"Result"`
}

type DescribeCdnIpResult struct {
	CdnIpv4 []CdnIp `json:"CdnIpv4"`
	CdnIpv6 []CdnIp `json:"CdnIpv6"`
}

type CdnIp struct {
	Ip   string `json:"Ip"`
	Cidr string `json:"Cidr,omitempty"`
}
