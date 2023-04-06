package imagex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	ServiceName    = "ImageX"
	ApiVersion     = "2018-08-01"
	DefaultTimeout = 10 * time.Second

	ResourceServiceIdTRN = "trn:ImageX:*:*:ServiceId/%s"
	ResourceStoreKeyTRN  = "trn:ImageX:*:*:StoreKeys/%s"

	MinChunkSize   = 1024 * 1024 * 20
	LargeFileSize  = 1024 * 1024 * 1024
	UploadRoutines = 4
)

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ServiceName,
			},
		},
		base.RegionUsEast1: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-us-east-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionUsEast1,
				Service: ServiceName,
			},
		},
		base.RegionApSingapore: {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "imagex-ap-singapore-1.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionApSingapore,
				Service: ServiceName,
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetImageServiceSubscription": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageServiceSubscription"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageService"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageService": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageService"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetAllImageServices": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageServices"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageService"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageAuthKey"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageAuthKey": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageAuthKey"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageObjectAccess": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageObjectAccess"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateImageMirrorConf": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateImageMirrorConf"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DelDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DelDomain"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetServiceDomains": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetServiceDomains"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetDomainConfig": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDomainConfig"},
				"Version": []string{"2018-08-01"},
			},
		},
		"SetDefaultDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetDefaultDomain"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateResponseHeader": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateResponseHeader"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateRefer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRefer"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateHttps": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHttps"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetResponseHeaderValidateKeys": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetResponseHeaderValidateKeys"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageTemplate"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageTemplate": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageTemplate"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageTemplate": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTemplate"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetAllImageTemplates": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAllImageTemplates"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetTemplatesFromBin": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTemplatesFromBin"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateTemplatesFromBin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTemplatesFromBin"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteTemplatesFromBin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTemplatesFromBin"},
				"Version": []string{"2018-08-01"},
			},
		},
		"ApplyImageUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyImageUpload"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CommitImageUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitImageUpload"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFile"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageUploadFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUploadFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DeleteImageUploadFiles": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageUploadFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"PreviewImageUploadFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PreviewImageUploadFile"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageContentTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageContentTask"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageContentTaskDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageContentTaskDetail"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageContentBlockList": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageContentBlockList"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageUpdateFiles": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageUpdateFiles"},
				"Version": []string{"2018-08-01"},
			},
		},
		"FetchImageUrl": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FetchImageUrl"},
				"Version": []string{"2018-08-01"},
			},
		},
		"UpdateServiceName": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateServiceName"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageOCR": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageOCR"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageQuality": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageQuality"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageEraseModels": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseModels"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageEnhanceResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEnhanceResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageBgFillResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageBgFillResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageDuplicateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageDuplicateDetection"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetDedupTaskStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDedupTaskStatus"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetDenoisingImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDenoisingImage"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetSegmentImage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSegmentImage"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageComicResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageComicResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageSuperResolutionResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSuperResolutionResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageSmartCropResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageSmartCropResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetLicensePlateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetLicensePlateDetection"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImagePSDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImagePSDetection"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetPrivateImageType": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPrivateImageType"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageHmEmbed": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageHmEmbed"},
				"Version": []string{"2018-08-01"},
			},
		},
		"CreateImageHmExtract": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateImageHmExtract"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageEraseResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageEraseResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"GetImageStyleResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageStyleResult"},
				"Version": []string{"2018-08-01"},
			},
		},
		"DescribeImageVolcCdnAccessLog": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeImageVolcCdnAccessLog"},
				"Version": []string{"2018-08-01"},
			},
		},
	}
)

// DefaultInstance 默认实例，Region: cn-north-1
var DefaultInstance = NewInstance()

type ImageX struct {
	*base.Client
}

func NewInstance() *ImageX {
	instance := &ImageX{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *ImageX {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("ImageX not support region %s", region))
	}
	instance := &ImageX{
		Client: base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

func init() {
	g := []string{"DescribeImageXDomainTrafficData",
		"DescribeImageXDomainBandwidthData",
		"DescribeImageXBucketUsage",
		"DescribeImageXRequestCntUsage",
		"DescribeImageXBaseOpUsage",
		"DescribeImageXCompressUsage",
		"DescribeImageXEdgeRequest",
		"DescribeImageXHitRateTrafficData",
		"DescribeImageXHitRateRequestData",
		"DescribeImageXCDNTopRequestData",
		"DescribeImageXSummary",
		"DescribeImageXEdgeRequestBandwidth",
		"DescribeImageXEdgeRequestTraffic",
		"DescribeImageXEdgeRequestRegions",
		"DescribeImageXServiceQuality",
		"GetImageXQueryApps",
		"GetImageXQueryRegions",
		"GetImageXQueryDims",
		"GetImageXQueryVals",
	}
	p := []string{"DescribeImageXMirrorRequestTraffic",
		"DescribeImageXMirrorRequestBandwidth",
		"DescribeImageXMirrorRequestHttpCodeByTime",
		"DescribeImageXMirrorRequestHttpCodeOverview",
		"DescribeImageXUploadSuccessRateByTime",
		"DescribeImageXUploadErrorCodeAll",
		"DescribeImageXUploadErrorCodeByTime",
		"DescribeImageXUploadCountByTime",
		"DescribeImageXUploadFileSize",
		"DescribeImageXUploadSpeed",
		"DescribeImageXUploadDuration",
		"DescribeImageXUploadSegmentSpeedByTime",
		"DescribeImageXCdnSuccessRateByTime",
		"DescribeImageXCdnSuccessRateAll",
		"DescribeImageXCdnErrorCodeByTime",
		"DescribeImageXCdnErrorCodeAll",
		"DescribeImageXCdnDurationDetailByTime",
		"DescribeImageXCdnDurationAll",
		"DescribeImageXCdnReuseRateByTime",
		"DescribeImageXCdnReuseRateAll",
		"DescribeImageXCdnProtocolRateByTime",
		"DescribeImageXClientErrorCodeAll",
		"DescribeImageXClientErrorCodeByTime",
		"DescribeImageXClientDecodeSuccessRateByTime",
		"DescribeImageXClientDecodeDurationByTime",
		"DescribeImageXClientQueueDurationByTime",
		"DescribeImageXClientLoadDurationAll",
		"DescribeImageXClientLoadDuration",
		"DescribeImageXClientFailureRate",
		"DescribeImageXClientSdkVerByTime",
		"DescribeImageXClientFileSize",
		"DescribeImageXClientTopFileSize",
		"DescribeImageXClientCountByTime",
		"DescribeImageXClientScoreByTime",
		"DescribeImageXClientDemotionRateByTime",
		"DescribeImageXClientTopDemotionURL",
		"DescribeImageXClientQualityRateByTime",
		"DescribeImageXClientTopQualityURL",
		"DescribeImageXSensibleCountByTime",
		"DescribeImageXSensibleCacheHitRateByTime",
		"DescribeImageXSensibleTopSizeURL",
		"DescribeImageXSensibleTopRamURL",
		"DescribeImageXSensibleTopResolutionURL",
		"DescribeImageXSensibleTopUnknownURL",
	}
	for _, i := range g {
		ApiInfoList[i] = &base.ApiInfo{
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{i},
				"Version": []string{"2018-08-01"},
			},
		}
	}
	for _, i := range p {
		ApiInfoList[i] = &base.ApiInfo{
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{i},
				"Version": []string{"2018-08-01"},
			},
		}
	}
}
