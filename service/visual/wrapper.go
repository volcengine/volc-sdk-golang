package visual

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func (p *Visual) commonHandler(api string, form url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonJsonHandler(api string, jsonStr string, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, nil, jsonStr)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) commonQueryHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) DistortionFree(form url.Values) (*model.DistortionFreeResult, int, error) {
	resp := new(model.DistortionFreeResult)
	statusCode, err := p.commonHandler("DistortionFree", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) StretchRecovery(form url.Values) (*model.StretchRecoveryResult, int, error) {
	resp := new(model.StretchRecoveryResult)
	statusCode, err := p.commonHandler("StretchRecovery", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EmoticonEdit(form url.Values) (*model.EmoticonEditResult, int, error) {
	resp := new(model.EmoticonEditResult)
	statusCode, err := p.commonHandler("EmoticonEdit", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EyeClose2Open(form url.Values) (*model.EyeClose2OpenResult, int, error) {
	resp := new(model.EyeClose2OpenResult)
	statusCode, err := p.commonHandler("EyeClose2Open", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageScore(form url.Values) (*model.ImageScoreResult, int, error) {
	resp := new(model.ImageScoreResult)
	statusCode, err := p.commonHandler("ImageScore", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageFlow(form url.Values) (*model.ImageFlowResult, int, error) {
	resp := new(model.ImageFlowResult)
	statusCode, err := p.commonHandler("ImageFlow", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) PoemMaterial(form url.Values) (*model.PoemMaterialResult, int, error) {
	resp := new(model.PoemMaterialResult)
	statusCode, err := p.commonHandler("PoemMaterial", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarPlateDetection(form url.Values) (*model.CarPlateDetectionResult, int, error) {
	resp := new(model.CarPlateDetectionResult)
	statusCode, err := p.commonHandler("CarPlateDetection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarSegment(form url.Values) (*model.CarSegmentResult, int, error) {
	resp := new(model.CarSegmentResult)
	statusCode, err := p.commonHandler("CarSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CarDetection(form url.Values) (*model.CarDetectionResult, int, error) {
	resp := new(model.CarDetectionResult)
	statusCode, err := p.commonHandler("CarDetection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) SkySegment(form url.Values) (*model.SkySegmentResult, int, error) {
	resp := new(model.SkySegmentResult)
	statusCode, err := p.commonHandler("SkySegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) BankCard(form url.Values) (*model.BankCardResult, int, error) {
	resp := new(model.BankCardResult)
	statusCode, err := p.commonHandler("BankCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) BankCardV2(form url.Values) (*model.BankCardResultV2, int, error) {
	resp := new(model.BankCardResultV2)
	statusCode, err := p.commonHandler("BankCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) IDCard(form url.Values) (*model.IDCardResult, int, error) {
	resp := new(model.IDCardResult)
	statusCode, err := p.commonHandler("IDCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) IDCardV2(form url.Values) (*model.IDCardResultV2, int, error) {
	resp := new(model.IDCardResultV2)
	statusCode, err := p.commonHandler("IDCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OCRNormal(form url.Values) (*model.OCRNormalResult, int, error) {
	resp := new(model.OCRNormalResult)
	statusCode, err := p.commonHandler("OCRNormal", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OCRApi(form url.Values, action string) (*model.OcrApiResult, int, error) {
	resp := new(model.OcrApiResult)
	statusCode, err := p.commonHandler(action, form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) DrivingLicense(form url.Values) (*model.DrivingLicenseResult, int, error) {
	resp := new(model.DrivingLicenseResult)
	statusCode, err := p.commonHandler("DrivingLicense", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VehicleLicense(form url.Values) (*model.VehicleLicenseResult, int, error) {
	resp := new(model.VehicleLicenseResult)
	statusCode, err := p.commonHandler("VehicleLicense", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) Taibao(form url.Values) (*model.TaibaoResult, int, error) {
	resp := new(model.TaibaoResult)
	statusCode, err := p.commonHandler("OcrTaibao", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceSwap(form url.Values) (*model.FaceSwapResult, int, error) {
	resp := new(model.FaceSwapResult)
	statusCode, err := p.commonHandler("FaceSwap", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoon(form url.Values) (*model.JPCartoonResult, int, error) {
	resp := new(model.JPCartoonResult)
	statusCode, err := p.commonHandler("JPCartoon", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoonCut(form url.Values) (*model.JPCartoonCutResult, int, error) {
	resp := new(model.JPCartoonCutResult)
	statusCode, err := p.commonHandler("JPCartoonCut", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSceneDetect(form url.Values) (*model.VideoSceneDetectResult, int, error) {
	resp := new(model.VideoSceneDetectResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OverResolution(form url.Values) (*model.OverResolutionResult, int, error) {
	resp := new(model.OverResolutionResult)
	statusCode, err := p.commonHandler("OverResolution", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OverResolutionV2(imageBase64 []string) (*model.OverResolutionV2Result, int, error) {
	//处理入参
	jsonBody := &struct {
		ReqKey           string   `json:"req_key"`
		BinaryDataBase64 []string `json:"binary_data_base64"`
	}{
		ReqKey:           "lens_vida_nnsr", //算法名称，取固定值为lens_vida_nnsr
		BinaryDataBase64: imageBase64,      //图片文件，base64编码。此算法可选输入1张图片或多张图片
	}

	jsonStr, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.OverResolutionV2Result)
	statusCode, err := p.commonJsonHandler("OverResolutionV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageScoreV2(imageBase64 []string) (*model.ImageScoreV2Result, int, error) {
	//处理入参
	jsonBody := &struct {
		ReqKey           string   `json:"req_key"`
		BinaryDataBase64 []string `json:"binary_data_base64"`
	}{
		ReqKey:           "lens_vida_single_pic", //算法名称，取固定值为lens_vida_nnsr
		BinaryDataBase64: imageBase64,            //图片文件，base64编码。此算法可选输入1张图片或多张图片
	}

	jsonStr, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.ImageScoreV2Result)
	statusCode, err := p.commonJsonHandler("ImageScoreV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageCorrection(imageBase64 []string) (*model.ImageCorrectionResult, int, error) {
	//处理入参
	jsonBody := &struct {
		ReqKey           string   `json:"req_key"`
		BinaryDataBase64 []string `json:"binary_data_base64"`
	}{
		ReqKey:           "image_correction", //算法名称，取固定值为lens_vida_nnsr
		BinaryDataBase64: imageBase64,        //图片文件，base64编码。此算法可选输入1张图片或多张图片
	}

	jsonStr, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.ImageCorrectionResult)
	statusCode, err := p.commonJsonHandler("ImageCorrection", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EnhancePhotoV2(req *model.EnhancePhotoV2Request) (*model.EnhancePhotoV2Result, int, error) {
	if req.ResolutionBoundary == "" {
		req.ResolutionBoundary = "720p"
	}
	if req.JpgQuality == 0 {
		req.JpgQuality = 95
	}
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.EnhancePhotoV2Result)
	statusCode, err := p.commonJsonHandler("EnhancePhotoV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) T2ILDM(req *model.T2ILDMRequest) (*model.T2ILDMResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.T2ILDMResult)
	statusCode, err := p.commonJsonHandler("T2ILDM", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) Img2ImgStyle(req *model.Img2ImgStyleRequest) (*model.Img2ImgStyleResult, int, error) {
	if req.Strength == 0.0 {
		req.Strength = 0.5
	}
	if req.Seed == 0 {
		req.Seed = -1
	}
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.Img2ImgStyleResult)
	statusCode, err := p.commonJsonHandler("Img2ImgStyle", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) Img2ImgAnime(req *model.Img2ImgAnimeRequest) (*model.Img2ImgAnimeResult, int, error) {
	if req.Strength == 0.0 {
		req.Strength = 0.5
	}
	if req.Seed == 0 {
		req.Seed = -1
	}
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.Img2ImgAnimeResult)
	statusCode, err := p.commonJsonHandler("Img2ImgAnime", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) BodyDetection(req *model.BodyDetectionRequest) (*model.BodyDetectionResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.BodyDetectionResult)
	statusCode, err := p.commonJsonHandler("BodyDetection", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) AllAgeGeneration(req *model.AllAgeGenerationRequest) (*model.AllAgeGenerationResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.AllAgeGenerationResult)
	statusCode, err := p.commonJsonHandler("AllAgeGeneration", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionSubmitTaskV2(req *model.VideoOverResolutionSubmitTaskV2Request) (*model.VideoOverResolutionSubmitTaskV2Result, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VideoOverResolutionSubmitTaskV2Result)
	statusCode, err := p.commonJsonHandler("VideoOverResolutionSubmitTaskV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionQueryTaskV2(req *model.VideoOverResolutionQueryTaskV2Request) (*model.VideoOverResolutionQueryTaskV2Result, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VideoOverResolutionQueryTaskV2Result)
	statusCode, err := p.commonJsonHandler("VideoOverResolutionQueryTaskV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceFusionMovieSubmitTask(req *model.FaceFusionMovieSubmitTaskRequest) (*model.FaceFusionMovieSubmitTaskResult, int, error) {
	if req.SourceSimilarity == "" {
		req.SourceSimilarity = "1"
	}

	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceFusionMovieSubmitTaskResult)
	statusCode, err := p.commonJsonHandler("FaceFusionMovieSubmitTask", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceFusionMovieGetResult(req *model.FaceFusionMovieGetResultRequest) (*model.FaceFusionMovieGetResultResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceFusionMovieGetResultResult)
	statusCode, err := p.commonJsonHandler("FaceFusionMovieGetResult", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceFusionMovie(req *model.FaceFusionMovieRequest) (*model.FaceFusionMovieResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceFusionMovieResult)
	statusCode, err := p.commonJsonHandler("FaceFusionMovie", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CertToken(req *model.CertTokenRequest) (*model.CertTokenResult, int, error) {
	if req.LivenessTimeout == 0 {
		req.LivenessTimeout = 10
	}
	if req.MotionList == nil {
		req.MotionList = []string{"0", "1", "2", "3"}
	}
	if req.FixedMotionList == nil {
		req.FixedMotionList = []string{"0", "1"}
	}
	if req.MotionCount == 0 {
		req.MotionCount = 2
	}
	if req.MaxLivenessTrial == 0 {
		req.MaxLivenessTrial = 10
	}

	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.CertTokenResult)
	statusCode, err := p.commonJsonHandler("CertToken", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CertConfigInit(req *model.CertConfigInitRequest) (*model.CertConfigInitResult, int, error) {
	if req.TokenApiConfig.LivenessTimeout == 0 {
		req.TokenApiConfig.LivenessTimeout = 10
	}
	if req.TokenApiConfig.MotionList == nil {
		req.TokenApiConfig.MotionList = []string{"0", "1", "2", "3"}
	}
	if req.TokenApiConfig.FixedMotionList == nil {
		req.TokenApiConfig.FixedMotionList = []string{"0", "1"}
	}
	if req.TokenApiConfig.MotionCount == 0 {
		req.TokenApiConfig.MotionCount = 2
	}
	if req.TokenApiConfig.MaxLivenessTrial == 0 {
		req.TokenApiConfig.MaxLivenessTrial = 10
	}

	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.CertConfigInitResult)
	statusCode, err := p.commonJsonHandler("CertConfigInit", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) CertVerifyQuery(req *model.CertVerifyQueryRequest) (*model.CertVerifyQueryResult, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.CertVerifyQueryResult)
	statusCode, err := p.commonJsonHandler("CertVerifyQuery", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) Img2Video3D(req *model.Img2Video3DRequest) (*model.Img2Video3DResult, int, error) {
	if req.RenderSpec.LongSide == 0 {
		req.RenderSpec.LongSide = 960
	}
	if req.RenderSpec.FrameNum == 0 {
		req.RenderSpec.FrameNum = 90
	}
	if req.RenderSpec.Fps == 0 {
		req.RenderSpec.Fps = 30
	}
	if req.RenderSpec.SpeedShift == nil {
		req.RenderSpec.SpeedShift = []float64{}
	}

	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.Img2Video3DResult)
	statusCode, err := p.commonJsonHandler("Img2Video3D", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ConvertPhotoV2(req *model.ConvertPhotoV2Request) (*model.ConvertPhotoV2Result, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.ConvertPhotoV2Result)
	statusCode, err := p.commonJsonHandler("ConvertPhotoV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) LensVidaVideoSubmitTaskV2(req *model.LensVidaVideoSubmitTaskV2Request) (*model.LensVidaVideoSubmitTaskV2Result, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.LensVidaVideoSubmitTaskV2Result)
	statusCode, err := p.commonJsonHandler("LensVidaVideoSubmitTaskV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) LensVidaVideoGetResultV2(req *model.LensVidaVideoGetResultV2Request) (*model.LensVidaVideoGetResultV2Result, int, error) {
	jsonStr, err := json.Marshal(req)

	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.LensVidaVideoGetResultV2Result)
	statusCode, err := p.commonJsonHandler("LensVidaVideoGetResultV2", string(jsonStr), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageStyleConversion(form url.Values) (*model.ImageStyleConversionResult, int, error) {
	resp := new(model.ImageStyleConversionResult)
	statusCode, err := p.commonHandler("ImageStyleConversion", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) PotraitEffect(form url.Values) (*model.PotraitEffectResult, int, error) {
	resp := new(model.PotraitEffectResult)
	statusCode, err := p.commonHandler("PotraitEffect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageAnimation(form url.Values) (*model.ImageAnimationResult, int, error) {
	resp := new(model.ImageAnimationResult)
	statusCode, err := p.commonHandler("ImageAnimation", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FacePretty(form url.Values) (*model.FacePrettyResult, int, error) {
	resp := new(model.FacePrettyResult)
	statusCode, err := p.commonHandler("FacePretty", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) HairStyle(form url.Values) (*model.HairStyleResult, int, error) {
	resp := new(model.HairStyleResult)
	statusCode, err := p.commonHandler("HairStyle", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) DollyZoom(form url.Values) (*model.DollyZoomResult, int, error) {
	resp := new(model.DollyZoomResult)
	statusCode, err := p.commonHandler("DollyZoom", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ThreeDGameCartoon(form url.Values) (*model.ThreeDGameCartoonResult, int, error) {
	resp := new(model.ThreeDGameCartoonResult)
	statusCode, err := p.commonHandler("ThreeDGameCartoon", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) HairSegment(form url.Values) (*model.HairSegmentResult, int, error) {
	resp := new(model.HairSegmentResult)
	statusCode, err := p.commonHandler("HairSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsSegment(form url.Values) (*model.GoodsSegmentResult, int, error) {
	resp := new(model.GoodsSegmentResult)
	statusCode, err := p.commonHandler("GoodsSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageOutpaint(form url.Values) (*model.ImageOutpaintResult, int, error) {
	resp := new(model.ImageOutpaintResult)
	statusCode, err := p.commonHandler("ImageOutpaint", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageInpaint(form url.Values) (*model.ImageInpaintResult, int, error) {
	resp := new(model.ImageInpaintResult)
	statusCode, err := p.commonHandler("ImageInpaint", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageCut(form url.Values) (*model.ImageCutResult, int, error) {
	resp := new(model.ImageCutResult)
	statusCode, err := p.commonHandler("ImageCut", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EntityDetect(form url.Values) (*model.EntityDetectResult, int, error) {
	resp := new(model.EntityDetectResult)
	statusCode, err := p.commonHandler("EntityDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsDetect(form url.Values) (*model.GoodsDetectResult, int, error) {
	resp := new(model.GoodsDetectResult)
	statusCode, err := p.commonHandler("GoodsDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ConvertPhoto(form url.Values) (*model.ConvertPhotoResult, int, error) {
	resp := new(model.ConvertPhotoResult)
	statusCode, err := p.commonHandler("ConvertPhoto", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EnhancePhoto(form url.Values) (*model.EnhancePhotoResult, int, error) {
	resp := new(model.EnhancePhotoResult)
	statusCode, err := p.commonHandler("EnhancePhoto", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GeneralSegment(form url.Values) (*model.GeneralSegmentResult, int, error) {
	resp := new(model.GeneralSegmentResult)
	statusCode, err := p.commonHandler("GeneralSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) HumanSegment(form url.Values) (*model.HumanSegmentResult, int, error) {
	resp := new(model.HumanSegmentResult)
	statusCode, err := p.commonHandler("HumanSegment", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoInpaintSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoInpaintSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoInpaintQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoInpaintQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoRetargetingSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoRetargetingSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoRetargetingQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoRetargetingQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoCoverSelection(form url.Values) (*model.VideoCoverSelectResult, int, error) {
	resp := new(model.VideoCoverSelectResult)
	statusCode, err := p.commonHandler("VideoCoverSelection", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSummarizationSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoSummarizationSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSummarizationQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoSummarizationQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionSubmitTask(form url.Values) (*model.VideoEditSubmitTaskResult, int, error) {
	resp := new(model.VideoEditSubmitTaskResult)
	statusCode, err := p.commonHandler("VideoOverResolutionSubmitTask", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoOverResolutionQueryTask(query url.Values) (*model.VideoEditQueryTaskResult, int, error) {
	resp := new(model.VideoEditQueryTaskResult)
	statusCode, err := p.commonQueryHandler("VideoOverResolutionQueryTask", query, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceCompare(bodyMap interface{}) (*model.FaceCompareResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceCompareResult)
	statusCode, err := p.commonJsonHandler("FaceCompare", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// FaceSwapV3 入参说明，接口文档：https://www.volcengine.com/docs/6791/1130928
// bodyMap: 按照接口文档填写入参，其中 req_key：- 3.0版本取固定值faceswap ｜ 3.3版本取固定值face_swap3_3
func (p *Visual) FaceSwapV3(bodyMap interface{}) (*model.FaceSwapV3Result, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceSwapV3Result)
	statusCode, err := p.commonJsonHandler("FaceSwapV3", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// FaceSwapAI 入参说明，接口文档：https://www.volcengine.com/docs/6791/1167911
// bodyMap: 按照接口文档填写入参，其中 req_key：取固定值faceswap_ai
func (p *Visual) FaceSwapAI(bodyMap interface{}) (*model.FaceSwapAIResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.FaceSwapAIResult)
	statusCode, err := p.commonJsonHandler("FaceSwapAI", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// HighAesSmartDrawing 人像人体-智能绘图-高美感xx+单图写真,接口文档：
// https://www.volcengine.com/docs/6791/1213129
// https://www.volcengine.com/docs/6791/1214814
// https://www.volcengine.com/docs/6791/1213130
// https://www.volcengine.com/docs/6791/1213131
// https://www.volcengine.com/docs/6791/1213132
// https://www.volcengine.com/docs/6791/1213737
// bodyMap: 按照接口文档填写入参
func (p *Visual) HighAesSmartDrawing(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("HighAesSmartDrawing", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// Img2ImgInpainting 涂抹消除。接口文档：https://www.volcengine.com/docs/6791/1223709
// bodyMap: 按照接口文档填写入参
func (p *Visual) Img2ImgInpainting(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("Img2ImgInpainting", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// Img2ImgInpaintingEdit 涂抹编辑。接口文档：https://www.volcengine.com/docs/6791/1223721
// bodyMap: 按照接口文档填写入参
func (p *Visual) Img2ImgInpaintingEdit(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("Img2ImgInpaintingEdit", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// Img2ImgOutpainting 智能扩图。接口文档：https://www.volcengine.com/docs/6791/1223722
// bodyMap: 按照接口文档填写入参
func (p *Visual) Img2ImgOutpainting(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("Img2ImgOutpainting", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// FaceFusionMovieSubmitTaskNew 视频人脸融合-提交任务
// QPS-2.x接口文档: https://www.volcengine.com/docs/6792/145429
// 按量-2.x接口文档: https://www.volcengine.com/docs/6792/271359
// QPS-3.x接口文档: https://www.volcengine.com/docs/6792/1124893
// bodyMap: 按照接口文档填写入参
func (p *Visual) FaceFusionMovieSubmitTaskNew(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("FaceFusionMovieSubmitTask", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

// FaceFusionMovieGetResultNew 视频人脸融合-查询任务。
// QPS-2.x接口文档: https://www.volcengine.com/docs/6792/145429
// 按量-2.x接口文档: https://www.volcengine.com/docs/6792/271359
// QPS-3.x接口文档: https://www.volcengine.com/docs/6792/1124893
// bodyMap: 按照接口文档填写入参
func (p *Visual) FaceFusionMovieGetResultNew(bodyMap map[string]interface{}) (*model.VisualPubResult, int, error) {
	reqByte, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, 500, errors.New("request json marshal error" + err.Error())
	}
	resp := new(model.VisualPubResult)
	statusCode, err := p.commonJsonHandler("FaceFusionMovieGetResult", string(reqByte), resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
