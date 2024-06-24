package visual

import (
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion          = "cn-north-1"
	ServiceVersion20200826 = "2020-08-26"
	ServiceName            = "cv"
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: 10 * time.Second,
		Host:    "visual.volcengineapi.com",
		Header:  http.Header{},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"DistortionFree": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DistortionFree"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"StretchRecovery": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StretchRecovery"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EmoticonEdit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EmoticonEdit"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EyeClose2Open": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EyeClose2Open"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageScore": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageScore"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageFlow": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageFlow"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"PoemMaterial": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PoemMaterial"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarPlateDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarPlateDetection"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"CarDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CarDetection"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"SkySegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SkySegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"BankCard": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BankCard"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"IDCard": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"IDCard"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OCRNormal": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRNormal"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"MultiLanguageOCR": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MultiLanguageOCR"},
				"Version": []string{"2022-08-31"},
			},
		},
		"DrivingLicense": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DrivingLicense"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VehicleLicense": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VehicleLicense"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrTaibao": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrTaibao"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrVatInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrVatInvoice"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrTaxiInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrTaxiInvoice"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrQuotaInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrQuotaInvoice"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrTrainTicket": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrTrainTicket"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrFlightInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrFlightInvoice"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrFinance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrFinance"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrRollInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrRollInvoice"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrPassInvoice": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrPassInvoice"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrFoodProduction": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrFoodProduction"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrFoodBusiness": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrFoodBusiness"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OcrClueLicense": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrClueLicense"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OCRTrade": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRTrade"},
				"Version": []string{"2020-12-21"},
			},
		},
		"OCRRuanzhu": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRRuanzhu"},
				"Version": []string{"2020-12-21"},
			},
		},
		"OCRCosmeticProduct": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRCosmeticProduct"},
				"Version": []string{"2020-12-21"},
			},
		},
		"OcrSeal": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrSeal"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OcrTextAlignment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OcrTextAlignment"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OCRPdf": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRPdf"},
				"Version": []string{"2021-08-23"},
			},
		},
		"OCRTable": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OCRTable"},
				"Version": []string{"2021-08-23"},
			},
		},
		"FaceSwap": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceSwap"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"JPCartoon": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"JPCartoon"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"JPCartoonCut": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"JPCartoonCut"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoSceneDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSceneDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OverResolution": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OverResolution"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GoodsSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GoodsSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageOutpaint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageOutpaint"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageInpaint": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageInpaint"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageCut": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageCut"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EntityDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EntityDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GoodsDetect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GoodsDetect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ConvertPhoto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ConvertPhoto"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"EnhancePhoto": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnhancePhoto"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"GeneralSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneralSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"HumanSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"HumanSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoInpaintSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoInpaintSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoInpaintQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoInpaintQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoRetargetingSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoRetargetingSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoRetargetingQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoRetargetingQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoSummarizationSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSummarizationSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoSummarizationQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoSummarizationQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoOverResolutionSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionSubmitTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"VideoOverResolutionQueryTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionQueryTask"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"OverResolutionV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OverResolutionV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"ImageScoreV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageScoreV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"EnhancePhotoV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnhancePhotoV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"T2ILDM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"T2ILDM"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2ImgStyle": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2ImgStyle"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2ImgAnime": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2ImgAnime"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"BodyDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BodyDetection"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"AllAgeGeneration": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AllAgeGeneration"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"VideoOverResolutionSubmitTaskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionSubmitTaskV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"VideoOverResolutionQueryTaskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoOverResolutionQueryTaskV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"FaceFusionMovieSubmitTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceFusionMovieSubmitTask"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"FaceFusionMovieGetResult": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceFusionMovieGetResult"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"ImageCorrection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageCorrection"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CertToken": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CertToken"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CertConfigInit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CertConfigInit"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CertVerifyQuery": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CertVerifyQuery"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2Video3D": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2Video3D"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"VideoCoverSelection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoCoverSelection"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageStyleConversion": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageStyleConversion"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"PotraitEffect": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"PotraitEffect"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ImageAnimation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageAnimation"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"FacePretty": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FacePretty"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"HairStyle": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"HairStyle"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"DollyZoom": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DollyZoom"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ThreeDGameCartoon": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"3DGameCartoon"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"HairSegment": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"HairSegment"},
				"Version": []string{ServiceVersion20200826},
			},
		},
		"ConvertPhotoV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ConvertPhotoV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"LensVidaVideoSubmitTaskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LensVidaVideoSubmitTaskV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"LensVidaVideoGetResultV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"LensVidaVideoGetResultV2"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"FaceCompare": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceCompare"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"FaceSwapV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceSwap"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"FaceSwapAI": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"FaceSwapAI"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"HighAesSmartDrawing": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"HighAesSmartDrawing"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2ImgInpainting": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2ImgInpainting"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2ImgInpaintingEdit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2ImgInpaintingEdit"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"Img2ImgOutpainting": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"Img2ImgOutpainting"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CertH5Token": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CertH5Token"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"CertH5ConfigInit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CertH5ConfigInit"},
				"Version": []string{"2022-08-31"},
			},
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
	}
)

// DefaultInstance 默认的实例
var DefaultInstance = NewInstance()

// IAM .
type Visual struct {
	Client *base.Client
}

// NewInstance 创建一个实例
func NewInstance() *Visual {
	instance := &Visual{}
	instance.Client = base.NewClient(ServiceInfo, ApiInfoList)
	instance.Client.ServiceInfo.Credentials.Service = ServiceName
	instance.Client.ServiceInfo.Credentials.Region = DefaultRegion
	return instance
}

// GetServiceInfo interface
func (p *Visual) GetServiceInfo() *base.ServiceInfo {
	return p.Client.ServiceInfo
}

// GetAPIInfo interface
func (p *Visual) GetAPIInfo(api string) *base.ApiInfo {
	if apiInfo, ok := ApiInfoList[api]; ok {
		return apiInfo
	}
	return nil
}

// SetRegion
func (p *Visual) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

// SetHost .
func (p *Visual) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

// SetSchema .
func (p *Visual) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
