package businessSecurity

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/url"
)

type RiskDetectionRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type RiskDetectionResponse struct {
	RequestId string       `json:"RequestId"`
	Code      int          `json:"Code"`
	Message   string       `json:"Message"`
	Data      DecisionData `json:"Data"`
}

type DeviceIdResponse struct {
	RequestId string     `json:"RequestId"`
	Code      int        `json:"Code"`
	Message   string     `json:"Message"`
	Data      DeviceInfo `json:"Data"`
}

type DecisionData struct {
	Score      int        `json:"Score"`
	Tags       []string   `json:"Tags"`
	Detail     string     `json:"Detail"`
	DeviceInfo DeviceInfo `json:"DeviceInfo"`
}

type DeviceInfo struct {
	DevSecID      string `json:"DevSecID"`
	TokenCreateTs int64  `json:"TokenCreateTs"`
	AccDevSecID   string `json:"AccDevSecID"`
}

type AsyncRiskDetectionRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type AsyncRiskDetectionResponse struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type RiskResultRequest struct {
	AppId     int64  `json:"AppId" form:"AppId"`
	Service   string `json:"Service" form:"Service"`
	StartTime int64  `json:"StartTime" form:"StartTime"`
	EndTime   int64  `json:"EndTime" form:"EndTime"`
	Page
}

type RiskResultResponse struct {
	RequestId string                   `json:"RequestId"`
	Code      int                      `json:"Code"`
	Message   string                   `json:"Message"`
	Data      []map[string]interface{} `json:"Data"`
	Page      Page                     `json:"page"`
}

type DataReportRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type DataReportResponse struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type AsyncVideoRiskResponse struct {
	RequestId string      `json:"RequestId"`
	Code      int         `json:"Code"`
	Message   string      `json:"Message"`
	Data      RequestData `json:"Data"`
}

type AsyncTextRiskResponse struct {
	RequestId string      `json:"RequestId"`
	Code      int         `json:"Code"`
	Message   string      `json:"Message"`
	Data      RequestData `json:"Data"`
}

type RequestData struct {
	PassThrough string `json:"PassThrough"`
}

type VideoResultRequest struct {
	DataId  string `json:"DataId"`
	AppId   int64  `json:"AppId"`
	Service string `json:"Service"`
}

type TextResultRequest struct {
	DataId  string `json:"DataId"`
	AppId   int64  `json:"AppId"`
	Service string `json:"Service"`
}

type VideoResultResponse struct {
	RequestId string    `json:"RequestId"`
	Code      int       `json:"Code"`
	Message   string    `json:"Message"`
	VideoResp VideoResp `json:"Data"`
}
type VideoLiveResultResponse struct {
	RequestId     string        `json:"RequestId"`
	Code          int           `json:"Code"`
	Message       string        `json:"Message"`
	VideoLiveResp VideoLiveResp `json:"Data"`
}

type VideoResp struct {
	DataId        string             `json:"DataId"`
	VideoResult   VideoResult        `json:"VideoResults"`
	AudioResults  AudioResultV2Basic `json:"AudioResults"`
	PassThrough   string             `json:"PassThrough"`
	FinalLabel    string             `json:"FinalLabel"`
	DecisionLabel string             `json:"DecisionLabel"`
}
type VideoLiveResp struct {
	DataId        string               `json:"DataId"`
	VideoResult   VideoLiveResult      `json:"VideoResults"`
	AudioResults  AudioLiveResultBasic `json:"AudioResults"`
	PassThrough   string               `json:"PassThrough"`
	FinalLabel    string               `json:"FinalLabel"`
	DecisionLabel string               `json:"DecisionLabel"`
}

type AudioResultV2 struct {
	AudioResultV2Basic
	FinalLabel    string `json:"FinalLabel"`
	DecisionLabel string `json:"DecisionLabel"`
}

type AudioResultV2Basic struct {
	Decision       string           `json:"Decision"`
	DecisionDetail string           `json:"DecisionDetail"`
	DataId         string           `json:"DataId"`
	AudioText      string           `json:"AudioText"`
	Details        []*AudioDetailV2 `json:"Details"`
	PassThrough    string           `json:"PassThrough"`
}

type AudioLiveResult struct {
	AudioLiveResultBasic
	FinalLabel    string `json:"FinalLabel"`
	DecisionLabel string `json:"DecisionLabel"`
}

type AudioLiveResultBasic struct {
	Decision    string               `json:"Decision"`
	DataId      string               `json:"DataId"`
	AudioText   string               `json:"AudioText"`
	Details     []*AudioLiveDetailV2 `json:"Details"`
	PassThrough string               `json:"PassThrough"`
}

type AudioDetailV2 struct {
	StartTime        int              `json:"StartTime"`
	EndTime          int              `json:"EndTime"`
	FrameUrl         string           `json:"FrameUrl"`
	AudioText        string           `json:"AudioText"`
	SliceId          string           `json:"SliceId"`
	UserId           string           `json:"UserId"`
	Decision         string           `json:"Decision"`
	DecisionDetail   string           `json:"DecisionDetail"`
	FrameResults     []*FrameResultV2 `json:"FrameResults"`
	FinalSubLabel    string           `json:"FinalSubLabel"`
	DecisionSubLabel string           `json:"DecisionSubLabel"`
}
type AudioLiveDetailV2 struct {
	StartTime              int              `json:"StartTime"`
	EndTime                int              `json:"EndTime"`
	StartTimeStamp         int64            `json:"StartTimeStamp"`
	EndTimeStamp           int64            `json:"EndTimeStamp"`
	FrameUrl               string           `json:"FrameUrl"`
	AudioText              string           `json:"AudioText"`
	SliceId                string           `json:"SliceId"`
	UserId                 string           `json:"UserId"`
	Decision               string           `json:"Decision"`
	FrameResults           []*FrameResultV2 `json:"FrameResults"`
	LiveFirstGetStreamTime int64            `json:"LiveFirstGetStreamTime"`
	FinalSubLabel          string           `json:"FinalSubLabel"`
	DecisionSubLabel       string           `json:"DecisionSubLabel"`
}
type FrameResultV2 struct {
	Label    string     `json:"Label"`
	SubLabel string     `json:"SubLabel"`
	Decision string     `json:"Decision"`
	Contexts []*Context `json:"Contexts"`
}

type VideoResult struct {
	Decision        string  `json:"Decision"`
	DecisionDetail  string  `json:"DecisionDetail"`
	ImageSliceCount int64   `json:"ImageSliceCount"`
	Frames          []Frame `json:"Frames"`
}
type VideoLiveResult struct {
	Decision        string      `json:"Decision"`
	ImageSliceCount int64       `json:"ImageSliceCount"`
	Frames          []LiveFrame `json:"Frames"`
}
type Frame struct {
	Url              string   `json:"Url"`
	Offset           float64  `json:"Offset"`
	SliceId          string   `json:"SliceId"`
	UserId           string   `json:"UserId"`
	Decision         string   `json:"Decision"`
	DecisionDetail   string   `json:"DecisionDetail"`
	Text             string   `json:"Text"`
	Results          []Result `json:"Results"`
	FinalSubLabel    string   `json:"FinalSubLabel"`
	DecisionSubLabel string   `json:"DecisionSubLabel"`
}
type LiveFrame struct {
	Url              string       `json:"Url"`
	Offset           float64      `json:"Offset"`
	SliceId          string       `json:"SliceId"`
	UserId           string       `json:"UserId"`
	Decision         string       `json:"Decision"`
	Text             string       `json:"Text"`
	Results          []LiveResult `json:"Results"`
	FinalSubLabel    string       `json:"FinalSubLabel"`
	DecisionSubLabel string       `json:"DecisionSubLabel"`
}

type Result struct {
	Decision string      `json:"Decision"`
	Label    string      `json:"Label"`
	SubLabel string      `json:"SubLabel"`
	Detail   interface{} `json:"Detail"`
}
type LiveResult struct {
	Decision string      `json:"Decision"`
	Label    string      `json:"Label"`
	SubLabel string      `json:"SubLabel"`
	Detail   interface{} `json:"Detail"`
}

type ImageResultResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	ImageResp ImageContentRiskResp `json:"Data"`
}

type ImageContentRiskResp struct {
	DataId         string                     `json:"DataId"`
	Decision       string                     `json:"Decision"`
	OcrText        string                     `json:"OcrText"`
	DecisionDetail string                     `json:"DecisionDetail"`
	Results        []*ImageContentRiskTag     `json:"Results"`
	FinalLabel     string                     `json:"FinalLabel"`
	DecisionLabel  string                     `json:"DecisionLabel"`
	Scores         map[string]float64         `json:"Scores"`
	OcrDetails     []*ImageContentFrameDetail `json:"OcrDetails"`
	QrcodeDetails  []*QrcodeData              `json:"QrcodeDetails"`
	PassThrough    string                     `json:"PassThrough"`
}

type ImageContentRiskTag struct {
	Label    string               `json:"Label"`
	SubLabel string               `json:"SubLabel"`
	Decision string               `json:"Decision"`
	Score    float64              `json:"score"`
	Detail   interface{}          `json:"Detail"`
	Frames   []*ImageContentFrame `json:"Frames"`
}
type ImageContentFrame struct {
	Id  float64 `json:"id"`
	Url string  `json:"url"`
	//Score float64 `json:"score"`
}
type ImageContentFrameDetail struct {
	Id           float64      `json:"Id"`
	OcrFrameText string       `json:"OcrFrameText"`
	Positions    []*Positions `json:"Positions"`
}
type QrcodeData struct {
	Id           float64  `json:"Id"`
	QrcodeResult []string `json:"QrcodeResult"`
}
type Positions struct {
	DetPointsRelative []*DetPointsRelative `json:"DetPointsRelative"`
	Text              string               `json:"Text"`
}
type DetPointsRelative struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

type AudioResultResponse struct {
	RequestId string        `json:"RequestId"`
	Code      int           `json:"Code"`
	Message   string        `json:"Message"`
	AudioResp AudioResultV2 `json:"Data"`
}
type AudioLiveResultResponse struct {
	RequestId     string          `json:"RequestId"`
	Code          int             `json:"Code"`
	Message       string          `json:"Message"`
	AudioLiveResp AudioLiveResult `json:"Data"`
}

func (resp *AudioResultResponse) String() string {
	respJSON, _ := json.Marshal(resp)
	return string(respJSON)
}

type AudioResult struct {
	Decision string         `json:"Decision"`
	Details  []*AudioDetail `json:"Details"`
	DataId   string         `json:"DataId"`
}

type AudioDetail struct {
	StartTime    float64        `json:"StartTime"`
	EndTime      float64        `json:"EndTime"`
	FrameUrl     string         `json:"FrameUrl"`
	AudioText    string         `json:"AudioText"`
	FrameResults []*FrameResult `json:"FrameResults"`
}

type FrameResult struct {
	Label        string   `json:"Label"`
	SubLabel     string   `json:"SubLabel"`
	Decision     string   `json:"Decision"`
	MatchedWords []string `json:"MatchedWords"`
	LibName      string   `json:"LibName"`
}

func (t *TextResultResponse) ToString() string {
	respJSON, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(respJSON)
}

type TextResultResponse struct {
	RequestId string         `json:"RequestId"`
	Code      int            `json:"Code"`
	Message   string         `json:"Message"`
	TextResp  TextRiskRespV3 `json:"Data"`
}

type TextRiskRespV2 struct {
	Decision       string   `json:"Decision"`
	DecisionDetail string   `json:"DecisionDetail"`
	Result         []*Label `json:"Result"`
}

type TextSliceResultResponse struct {
	RequestId    string         `json:"RequestId"`
	Code         int            `json:"Code"`
	Message      string         `json:"Message"`
	TextRiskResp TextRiskRespV3 `json:"Data"`
}

type TextRiskRespV3 struct {
	DataId         string        `json:"DataId"`
	PassThrough    string        `json:"PassThrough"`
	Decision       string        `json:"Decision"`
	DecisionDetail string        `json:"DecisionDetail"`
	FinalLabel     string        `json:"FinalLabel"`
	DecisionLabel  string        `json:"DecisionLabel"`
	TextCount      int           `json:"TextCount"`
	Results        []*TextResult `json:"Results"`
}

type TextResult struct {
	RiskText         string   `json:"RiskText"`
	RTStartPos       int      `json:"RTStartPos"`
	RTEndPos         int      `json:"RTEndPos"`
	Labels           []*Label `json:"Labels"`
	FinalSubLabel    string   `json:"FinalSubLabel"`
	DecisionSubLabel string   `json:"DecisionSubLabel"`
}

func (resp *TextSliceResultResponse) String() string {
	respJSON, err := json.Marshal(resp)
	if err != nil {
		return err.Error()
	}
	return string(respJSON)
}

type Label struct {
	Label    string    `json:"Label"`
	SubLabel string    `json:"SubLabel"`
	Decision string    `json:"Decision"`
	Contexts []Context `json:"Contexts"`
}

type Context struct {
	MatchedWords []string   `json:"MatchedWords"`
	LibName      string     `json:"LibName"`
	Description  string     `json:"Description"`
	Positions    []Position `json:"Positions"`
}

type Position struct {
	StartPos int `json:"StartPos"`
	EndPos   int `json:"EndPos"`
}

type Page struct {
	PageNum  int64 `json:"PageNum" form:"PageNum"`
	PageSize int64 `json:"PageSize" form:"PageSize"`
	Total    int64 `json:"Total" form:"Total"` // used when return page information in response, no need when request
}

type ElementVerifyRequest struct {
	AppId         int64  `json:"AppId"`
	Service       string `json:"Service"`
	EncryptedType string `json:"EncryptedType"`
	Parameters    string `json:"Parameters"`
}

type ElementVerifyResponse struct {
	RequestId string            `json:"RequestId"`
	Code      int               `json:"Code"`
	Message   string            `json:"Message"`
	Data      ElementVerifyData `json:"Data"`
}

type ElementVerifyData struct {
	Status int    `json:"Status"`
	Detail string `json:"Detail"`
}

type ElementVerifyResponseV2 struct {
	RequestId string              `json:"RequestId"`
	Code      int                 `json:"Code"`
	Message   string              `json:"Message"`
	Data      ElementVerifyDataV2 `json:"Data"`
}

type ElementVerifyDataV2 struct {
	Status int                     `json:"Status"`
	Detail ElementVerifyDataDetail `json:"Detail"`
}

type ElementVerifyDataDetail struct {
	BankName       string `json:"BankName,omitempty"`
	CardType       string `json:"CardType,omitempty"`
	CardCategory   string `json:"CardCategory,omitempty"`
	EnterpriseCode string `json:"EnterpriseCode,omitempty"`
	EnterpriseName string `json:"EnterpriseName,omitempty"`
	IdCardNo       string `json:"IdCardNo,omitempty"`
	IdCardName     string `json:"IdCardName,omitempty"`
	Reason         int64  `json:"Reason,omitempty"`
}

type MobileStatusRequest struct {
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
}

type MobileStatusResponse struct {
	RequestId string           `json:"RequestId"`
	Code      int              `json:"Code"`
	Message   string           `json:"Message"`
	Data      MobileStatusData `json:"Data"`
}

type MobileStatusData struct {
	Status int    `json:"Status"`
	Mobile string `json:"Mobile"`
	Detail string `json:"Detail"`
}

type MobileStatusResponseV2 struct {
	RequestId string             `json:"RequestId"`
	Code      int                `json:"Code"`
	Message   string             `json:"Message"`
	Data      MobileStatusDataV2 `json:"Data"`
}

type MobileStatusDataV2 struct {
	Status int                `json:"Status"`
	Mobile string             `json:"Mobile"`
	Detail MobileStatusDetail `json:"Detail"`
}

type MobileStatusDetail struct {
	Province     string `json:"Province"`
	ProvinceCode string `json:"ProvinceCode"`
	City         string `json:"City"`
	CityCode     string `json:"CityCode"`
	ISP          string `json:"ISP"`
	Mobile       string `json:"Mobile"`
	PostCode     string `json:"PostCode"`
	OldISP       int    `json:"OldISP"`
	NewISP       int    `json:"NewISP"`
}

func (r *RiskResultRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

func (r *VideoResultRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

func (r *TextResultRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

type CustomContentResponse struct {
	LogId     string      `json:"log_id"`
	ErrCode   int         `json:"err_code"`
	ErrMsg    string      `json:"err_msg"`
	Timestamp int         `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type NewCustomContentsReq struct {
	AppID       int64  `json:"app_id" form:"app_id" query:"app_id"`
	Service     string `json:"service" form:"service" query:"service"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	Decision    string `json:"decision" form:"decision" query:"decision"`
	MatchType   string `json:"match_type" form:"match_type" query:"match_type"`
	LibType     string `json:"lib_type" form:"lib_type" query:"lib_type"`
	Biztypes    string `json:"biztypes" form:"biztypes" query:"biztypes"`
}

type UpdateContentReq struct {
	AppID   int64  `json:"app_id" form:"app_id" query:"app_id"`
	Service string `json:"service" form:"service" query:"service"`
	LibType string `json:"lib_type" form:"lib_type" query:"lib_type"`
	Name    string `json:"name" form:"name" query:"name"`
	Status  int    `json:"status" form:"status" query:"status"`
}

type ModifyTextContent struct {
	Data    []string `json:"data" form:"data" query:"data"`
	AppID   int64    `json:"app_id" form:"app_id" query:"app_id"`
	Service string   `json:"service" form:"service" query:"service"`
	Name    string   `json:"name" form:"name" query:"name"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

// openapi
type SimpleRiskStatResponse struct {
	Result *SimpleRiskStatResult `json:"Result"`
}

type ContentRiskStatResponse struct {
	Result *ContentRiskStatResult `json:"Result"`
}

type OpenResult struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type SimpleRiskStatResult struct {
	OpenResult
	Data *SimpleProductStatisticsResult `json:"Data"`
}

type ContentRiskStatResult struct {
	OpenResult
	Data *ContentProductStatisticsResult `json:"Data"`
}

func (result *OpenResult) GetErr() error {
	return errors.New(fmt.Sprintf("err msg is %s, reqId is %s", result.Message, result.RequestId))
}
func (r *CreateAppReq) ToQuery() url.Values {
	return ToUrlValues(r)
}
func (r *ListAppsReq) ToQuery() url.Values {
	return ToUrlValues(r)
}

type CreateAppReq struct {
	AppName     string `json:"AppName" form:"AppName" query:"AppName"`
	AppNameZh   string `json:"AppNameZh" form:"AppNameZh" query:"AppNameZh"`
	AppType     string `json:"AppType" form:"AppType" query:"AppType"`
	Description string `json:"Description" form:"Description" query:"Description"`
}
type ListAppsReq struct {
	Limit  int `json:"Limit" form:"Limit" query:"Limit"`
	Offset int `json:"Offset" form:"Offset" query:"Offset"`
}
type OpenProductReq struct {
	Product string `json:"Product" form:"Product" query:"Product"`
}
type OpenProductRes struct {
	OpenResult
}
type CheckProductStatusReq struct {
	Product string `json:"Product" form:"Product" query:"Product"`
}
type CheckProductStatusRes struct {
	OpenResult
	Data CheckProductStatusResult `json:"Data"`
}
type CheckProductStatusResult struct {
	Status int `json:"Status"`
}

type EnableServiceReq struct {
	Service string `json:"Service" form:"Service" query:"Service"`
	AppID   int64  `json:"AppId" form:"AppId" query:"AppId"`
	Product string `json:"Product" form:"Product" query:"Product"`
}
type EnableServiceRes struct {
	OpenResult
}
type CheckServiceStatusReq struct {
	Service string `json:"Service" form:"Service" query:"Service"`
	AppID   int64  `json:"AppId" form:"AppId" query:"AppId"`
	Product string `json:"Product" form:"Product" query:"Product"`
}
type CheckServiceStatusRes struct {
	OpenResult
	Data CheckServiceStatusResult `json:"Data"`
}
type CheckServiceStatusResult struct {
	Status int `json:"Status"`
}
type CreateAppResponse struct {
	OpenResult
	Data CreateAppResult `json:"Data"`
}
type CreateAppResult struct {
	AppId       uint   `json:"AppId"`
	AppName     string `json:"AppName"`
	AppNameZh   string `json:"AppNameZh"`
	AppType     string `json:"AppType"`
	Description string `json:"Description"`
}
type ListAppsResponse struct {
	OpenResult
	Data ListAppsResult `json:"Data"`
}
type ListAppsResult struct {
	Total  int                 `json:"Total"`
	Limit  int                 `json:"Limit"`
	Offset int                 `json:"Offset"`
	Apps   []*ListAppsAppModel `json:"Apps"`
}
type ListAppsAppModel struct {
	AppId       uint   `json:"AppId"`
	AppName     string `json:"AppName"`
	AppNameZh   string `json:"AppNameZh"`
	AppType     string `json:"AppType"`
	Description string `json:"Description"`
	CreatedTime string `json:"CreatedTime"`
	UpdatedTime string `json:"UpdatedTime"`
	IsFavor     bool   `json:"IsFavor"`
}
type CommonProductStatisticsReq struct {
	Product    string  `json:"Product" form:"Product" query:"Product"`
	UnitType   string  `json:"UnitType" form:"UnitType" query:"UnitType"`
	Parameters string  `json:"Parameters" form:"Parameters" query:"Parameters"`
	AppId      *int64  `json:"AppId" form:"AppId" query:"AppId"`
	Service    *string `json:"Service" form:"Service" query:"Service"`
}
type CommonProductStatisticsParams struct {
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	NeedAppDetail bool   `json:"need_app_detail"`
	OperateTime   int64  `json:"operate_time"`
}
type SimpleProductStatisticsParams struct {
	CommonProductStatisticsParams
	NeedServiceDetail bool `json:"need_service_detail"`
}
type ContentProductStatisticsParams struct {
	CommonProductStatisticsParams
	Biztype            string `json:"biztype"`
	NeedBizTypeDetail  bool   `json:"need_biz_type_detail"`
	RiskType           string `json:"risk_type"`
	NeedRiskTypeDetail bool   `json:"need_risk_type_detail"`
}

type SimpleProductStatisticsResult struct {
	Total  *SimpleProductStatisticsResultTotal    `json:"Total"`
	Detail []*SimpleProductStatisticsResultDetail `json:"Detail"`
}

type ContentProductStatisticsResult struct {
	Total  *ContentProductStatisticsResultTotal    `json:"Total"`
	Detail []*ContentProductStatisticsResultDetail `json:"Detail"`
}

type CommonProductStatisticsResultTotal struct {
	AccountId        int64  `json:"AccountId"`
	RequestCnt       int64  `json:"RequestCnt"`
	ChargeRequestCnt *int64 `json:"ChargeRequestCnt,omitempty"`
}

type CommonProductStatisticsResultDetail struct {
	AccountId  int64  `json:"AccountId"`
	RequestCnt int64  `json:"RequestCnt"`
	DateTime   string `json:"DateTime"`

	ChargeRequestCnt *int64  `json:"ChargeRequestCnt,omitempty"`
	Service          *string `json:"Service,omitempty"`
	AppId            *int64  `json:"AppId,omitempty"`
	AppName          *string `json:"AppName,omitempty"`
}

type SimpleProductStatisticsResultTotal CommonProductStatisticsResultTotal
type SimpleProductStatisticsResultDetail CommonProductStatisticsResultDetail

type ContentProductStatisticsResultTotal struct {
	CommonProductStatisticsResultTotal

	RequestWordCnt  *int64   `json:"RequestWordCnt,omitempty"`
	RequestFrameCnt *int64   `json:"RequestFrameCnt,omitempty"`
	RequestDuration *float64 `json:"RequestAudioDuration,omitempty"`
}

type ContentProductStatisticsResultDetail struct {
	CommonProductStatisticsResultDetail
	Biztype         *string  `json:"Biztype,omitempty"`
	RequestWordCnt  *int64   `json:"RequestWordCnt,omitempty"`
	RiskType        *string  `json:"RiskType,omitempty"`
	RequestFrameCnt *int64   `json:"RequestFrameCnt,omitempty"`
	RequestDuration *float64 `json:"RequestAudioDuration,omitempty"`
}

type DelSystemNameListItemReq struct {
	Product string `json:"Product" form:"Product" query:"Product"`
	AppId   int64  `json:"AppId" form:"AppId" query:"AppId"`
	Service string `json:"Service" form:"Service" query:"Service"`
	Object  string `json:"Object" form:"Object" query:"Object"`
}

type QuerySystemNameListItemReq struct {
	Product string `json:"Product" form:"Product" query:"Product"`
	AppId   int64  `json:"AppId" form:"AppId" query:"AppId"`
	Service string `json:"Service" form:"Service" query:"Service"`
	Object  string `json:"Object" form:"Object" query:"Object"`
}

func (r *QuerySystemNameListItemReq) ToQuery() url.Values {
	return ToUrlValues(r)
}

type DelSystemNameListItemResp struct {
	OpenResult
	Data *DelSystemNameListItemResult `json:"Data"`
}

type QuerySystemNameListItemResp struct {
	OpenResult
	Data *QuerySystemNameListItemResult `json:"Data"`
}

type DelSystemNameListItemResult struct {
	Status int `json:"Status"`
}

type QuerySystemNameListItemResult struct {
	IsHit bool `json:"IsHit"`
}

type SecuritySourceResponse struct {
	RequestId string             `json:"RequestId"`
	Code      int                `json:"Code"`
	Message   string             `json:"Message"`
	Data      SecuritySourceData `json:"Data"`
}

type SecuritySourceData struct {
	Content      string                      `json:"Content"`
	References   []*SecuritySourceReferences `json:"References"`
	FinishReason string                      `json:"FinishReason"`
}

type SecuritySourceReferences struct {
	Url      string `json:"Url"`
	Idx      int64  `json:"Idx"`
	LogoUrl  string `json:"LogoUrl"`
	SiteName string `json:"SiteName"`
}

type TextModerationAsyncResponse struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      TextModerationAsyncData `json:"Data"`
}
type TextModerationSyncResponse struct {
	RequestId string                 `json:"RequestId"`
	Code      int                    `json:"Code"`
	Message   string                 `json:"Message"`
	Data      TextModerationSyncData `json:"Data"`
}

type TextModerationAsyncData struct {
	Decision      string `json:"Decision"`
	DecisionLabel string `json:"DecisionLabel"`
	PassThrough   string `json:"PassThrough"`
	DataId        string `json:"DataId"`
}
type TextModerationSyncData struct {
	Decision      string `json:"Decision"`
	DecisionLabel string `json:"DecisionLabel"`
}

type MultiModerationSyncResponse struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      MultiModerationSyncData `json:"Data"`
}
type MultiModerationAsyncResponse struct {
	RequestId string                   `json:"RequestId"`
	Code      int                      `json:"Code"`
	Message   string                   `json:"Message"`
	Data      MultiModerationAsyncData `json:"Data"`
}

type MultiModerationAsyncData struct {
	Decision      string   `json:"Decision"`
	DecisionLabel string   `json:"DecisionLabel"`
	HitLabels     []string `json:"HitLabels"`
	DataId        string   `json:"DataId"`
	PassThrough   string   `json:"PassThrough"`
}
type MultiModerationSyncData struct {
	Decision      string   `json:"Decision"`
	DecisionLabel string   `json:"DecisionLabel"`
	HitLabels     []string `json:"HitLabels"`
}
type CustomRiskAsyncResponse struct {
	RequestId string              `json:"RequestId"`
	Code      int                 `json:"Code"`
	Message   string              `json:"Message"`
	Data      CustomRiskAsyncData `json:"Data"`
}
type CustomRiskSyncResponse struct {
	RequestId string             `json:"RequestId"`
	Code      int                `json:"Code"`
	Message   string             `json:"Message"`
	Data      CustomRiskSyncData `json:"Data"`
}

type CustomRiskAsyncData struct {
	Decision       string  `json:"Decision"`
	DecisionLabel  string  `json:"DecisionLabel"`
	DataId         string  `json:"DataId"`
	DecisionRate   float64 `json:"DecisionRate"`
	DecisionReason string  `json:"DecisionReason"`
	PassThrough    string  `json:"PassThrough"`
}
type CustomRiskSyncData struct {
	Decision       string  `json:"Decision"`
	DecisionLabel  string  `json:"DecisionLabel"`
	DecisionRate   float64 `json:"DecisionRate"`
	DecisionReason string  `json:"DecisionReason"`
}

type AsyncRcLlmResponse struct {
	RequestId string      `json:"RequestId"`
	Code      int         `json:"Code"`
	Message   string      `json:"Message"`
	Data      RequestData `json:"Data"`
}

type RcLlmResultRequest struct {
	DataId  string `json:"DataId"`
	AppId   int64  `json:"AppId"`
	Service string `json:"Service"`
}

type CommonResponse struct {
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
}

type CreateCustomLibRequest struct {
	AppId       int64  `json:"app_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MatchType   string `json:"match_type"`
	Decision    string `json:"decision"`
	BizTypes    string `json:"biztypes"`
	Service     string `json:"service"`
	LibType     string `json:"lib_type"`
}

type CreateCustomLibResponse struct {
	RequestId string              `json:"RequestId"`
	Code      int                 `json:"Code"`
	Message   string              `json:"Message"`
	Data      CustomLibCreateInfo `json:"Data"`
}

type CustomLibCreateInfo struct {
	Id int `json:"id"`
}

type UpdateCustomLibRequest struct {
	AppId       int64  `json:"app_id"`
	Description string `json:"description"`
	BizTypes    string `json:"biztypes"`
	Service     string `json:"service"`
	LibType     string `json:"lib_type"`
	Name        string `json:"name"`
}

type ChangeCustomLibStatusRequest struct {
	AppId   int64  `json:"app_id"`
	Service string `json:"service"`
	LibType string `json:"lib_type"`
	Name    string `json:"name"`
	Status  int64  `json:"status"`
}

type DeleteCustomLibRequest struct {
	AppId   int64  `json:"app_id"`
	Name    string `json:"name"`
	Service string `json:"service"`
	LibType string `json:"lib_type"`
}

type GetCustomLibRequest struct {
	AppId   int64  `json:"app_id"`
	Service string `json:"service"`
	LibType string `json:"lib_type"`
}

type CustomLibListResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	Data      []*CustomLibListData `json:"Data"`
}

type CustomLibListData struct {
	Status      int64  `json:"status"`
	Decision    string `json:"decision"`
	Name        string `json:"name"`
	MatchType   string `json:"matchType"`
	BizTypes    string `json:"biztypes"`
	Description string `json:"description"`
	Count       int64  `json:"count"`
}

type CreateAccessConfigRequest struct {
	AppId     int64  `json:"app_id"`
	Service   string `json:"service_item"`
	Name      string `json:"name"`
	BizType   string `json:"biztype"`
	Scene     string `json:"scene"`
	Labels    string `json:"labels"`
	OcrLabels string `json:"ocr_labels"`
	AsrLabels string `json:"asr_labels"`
	TextLibs  string `json:"text_libs"`
	ImageLibs string `json:"image_libs"`
}

type CreateAccessConfigResponse struct {
	RequestId string                 `json:"RequestId"`
	Code      int                    `json:"Code"`
	Message   string                 `json:"Message"`
	Data      AccessConfigCreateInfo `json:"Data"`
}

type AccessConfigCreateInfo struct {
	Id int `json:"id"`
}

type UpdateAccessConfigRequest struct {
	Id        int64  `json:"id"`
	AppId     int64  `json:"app_id"`
	Service   string `json:"service_item"`
	Name      string `json:"name"`
	BizType   string `json:"biztype"`
	Scene     string `json:"scene"`
	Labels    string `json:"labels"`
	OcrLabels string `json:"ocr_labels"`
	AsrLabels string `json:"asr_labels"`
	TextLibs  string `json:"text_libs"`
	ImageLibs string `json:"image_libs"`
}

type UpdateAccessConfigStatusRequest struct {
	Id      int64  `json:"id"`
	AppId   int64  `json:"app_id"`
	Service string `json:"service_item"`
	BizType string `json:"biztype"`
	Status  int64  `json:"status"`
}

type GetAccessConfigStatusRequest struct {
	AppId   int64  `json:"app_id"`
	Service string `json:"service_item"`
}

type AccessConfigListResponse struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      []*AccessConfigListData `json:"Data"`
}

type AccessConfigListData struct {
	Id        int64  `json:"id"`
	BizType   string `json:"biztype"`
	Labels    string `json:"labels"`
	Name      string `json:"name"`
	Scene     string `json:"scene"`
	Status    int64  `json:"status"`
	OcrLabels string `json:"ocrLabels"`
	AsrLabels string `json:"asrLabels"`
	TextLibs  string `json:"textLibs"`
	ImageLibs string `json:"imageLibs"`
}

type GetCustomTextLibRequest struct {
	AppId    int64  `json:"app_id"`
	Service  string `json:"service"`
	Name     string `json:"name"`
	KeyWord  string `json:"keyword"`
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
}

type CustomTextLibListResponse struct {
	RequestId string                   `json:"RequestId"`
	Code      int                      `json:"Code"`
	Message   string                   `json:"Message"`
	Data      []*CustomTextLibListData `json:"Data"`
	Total     int                      `json:"TotalCnt"`
}

type CustomTextLibListData struct {
	Content         string `json:"content"`
	OriginContent   string `json:"originContent"`
	FirstUploadTime int64  `json:"firstUploadTime"`
	UpdateTime      int64  `json:"updateTime"`
}

type DeleteCustomTextRequest struct {
	AppId   int64    `json:"app_id"`
	Service string   `json:"service"`
	Name    string   `json:"name"`
	Data    []string `json:"data"`
}

type UploadCustomTextRequest struct {
	AppId   int64    `json:"app_id"`
	Service string   `json:"service"`
	Name    string   `json:"name"`
	Data    []string `json:"data"`
}

type GetCustomImgLibRequest struct {
	AppId    int64  `json:"app_id"`
	Service  string `json:"service"`
	Name     string `json:"name"`
	KeyWord  string `json:"keyword"`
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
}

type CustomImgLibListResponse struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      []*CustomImgLibListData `json:"Data"`
	Total     int                     `json:"TotalCnt"`
}

type CustomImgLibListData struct {
	Content    string `json:"content"`
	ImageName  string `json:"imageName"`
	ImageId    string `json:"imageId"`
	UploadTime int64  `json:"uploadTime"`
}

type DeleteCustomImgRequest struct {
	AppId   int64    `json:"app_id"`
	Service string   `json:"service"`
	Name    string   `json:"name"`
	Data    []string `json:"data"`
}

type UploadCustomImgRequest struct {
	AppId   int64            `json:"app_id"`
	Service string           `json:"service"`
	Name    string           `json:"name"`
	Data    []*CustomImgInfo `json:"data"`
}

type CustomImgInfo struct {
	ImageName string `json:"image_name"`
	Url       string `json:"url"`
	ImageData string `json:"image_data"`
}

type UploadCustomImgResponse struct {
	RequestId string              `json:"RequestId"`
	Code      int                 `json:"Code"`
	Message   string              `json:"Message"`
	Data      UploadCustomImgInfo `json:"Data"`
}

type UploadCustomImgInfo struct {
	Ids []int `json:"ids"`
}
type ActivateRiskBasePackageReq struct {
	PackageId       string   `json:"PackageId"`
	TotalPackageNum int      `json:"TotalPackageNum"`
	PackageSeq      int      `json:"PackageSeq"`
	DataType        string   `json:"DataType"`
	Data            []string `json:"Data"`
}

type ActivateRiskSampleDataReq struct {
	PackageId       string       `json:"PackageId"`
	TotalPackageNum int          `json:"TotalPackageNum"`
	PackageSeq      int          `json:"PackageSeq"`
	BusinessType    string       `json:"BusinessType"`
	DataType        string       `json:"DataType"`
	Data            []SampleData `json:"Data"`
}

type SampleData struct {
	Id           string `json:"Id"`
	ReachType    string `json:"ReachType"`
	LaunchStatus string `json:"LaunchStatus"`
	Extra        string `json:"Extra"`
}

type ActivateRiskResultReq struct {
	PlanId       int64  `json:"PlanId"`
	ActivateCode string `json:"ActivateCode"`
	AppId        int64  `json:"AppId"`
}
type ActivateRiskBasePackageResp struct {
	RequestId string                      `json:"RequestId"`
	Code      int                         `json:"Code"`
	Message   string                      `json:"Message"`
	Data      ActivateRiskBasePackageInfo `json:"Data"`
}

type ActivateRiskBasePackageInfo struct {
	ActivateCode string `json:"ActivateCode"`
	Status       int    `json:"Status"`
}
type ActivateRiskSampleDataResp struct {
	RequestId string                     `json:"RequestId"`
	Code      int                        `json:"Code"`
	Message   string                     `json:"Message"`
	Data      ActivateRiskSampleDataInfo `json:"Data"`
}
type ActivateRiskSampleDataInfo struct {
	Status int `json:"Status"`
}
type ActivateRiskResultResp struct {
	RequestId string                 `json:"RequestId"`
	Code      int                    `json:"Code"`
	Message   string                 `json:"Message"`
	Data      ActivateRiskResultInfo `json:"Data"`
}
type ActivateRiskResultInfo struct {
	Success  bool `json:"Success"`
	ResultId int  `json:"ResultId"`
}

type CancelActivateRiskResultReq struct {
	ResultId int `json:"ResultId"`
}

type CancelActivateRiskResultResp struct {
	RequestId string                       `json:"RequestId"`
	Code      int                          `json:"Code"`
	Message   string                       `json:"Message"`
	Data      CancelActivateRiskResultInfo `json:"Data"`
}

type CancelActivateRiskResultInfo struct {
	Success bool `json:"Success"`
}

type ContentStatisticsReq struct {
	AppId      int64  `json:"AppId"`
	Scene      string `json:"Scene"`
	LabelsType string `json:"LabelsType"`
	StartTime  int64  `json:"StartTime"`
	EndTime    int64  `json:"EndTime"`
}

type TextRiskStatisticsRsp struct {
	RequestId string                 `json:"RequestId"`
	Code      int                    `json:"Code"`
	Message   string                 `json:"Message"`
	Data      TextRiskStatisticsInfo `json:"Data"`
}

type TextRiskStatisticsInfo struct {
	TextTotal           TextBaseStatisticsRsp     `json:"TextTotal"`
	TextTrendRecord     []*TextBaseTrendRecordRsp `json:"TextTrendRecord"`
	TextRiskTotal       ContentRiskBaseCountRsp   `json:"TextRiskTotal"`
	TextRiskTrendRecord ContentRiskBaseCountRsp   `json:"TextRiskTrendRecord"`
}

type CommonStatisticsInfo struct {
	RequestCount int64 `json:"RequestCnt"`
	PassCount    int64 `json:"PassCnt"`
	BlockCount   int64 `json:"BlockCnt"`
	ReviewCount  int64 `json:"ReviewCnt"`
}

type FrameStatisticsInfo struct {
	PassFrameCnt   int64 `json:"PassFrameCnt"`
	BlockFrameCnt  int64 `json:"BlockFrameCnt"`
	ReviewFrameCnt int64 `json:"ReviewFrameCnt"`
}

type SplitStatisticsInfo struct {
	SplitCnt        int64   `json:"SplitCnt"`
	PassSplitCnt    int64   `json:"PassSplitCnt"`
	BlockSplitCnt   int64   `json:"BlockSplitCnt"`
	ReviewSplitCnt  int64   `json:"ReviewSplitCnt"`
	RequestDuration float64 `json:"RequestDuration"`
	PassDuration    float64 `json:"PassDuration"`
	BlockDuration   float64 `json:"BlockDuration"`
	ReviewDuration  float64 `json:"ReviewDuration"`
}

type TextBaseStatisticsRsp struct {
	CommonStatisticsInfo
	RequestWordCnt int64 `json:"RequestWordCnt"`
	PassWordCnt    int64 `json:"PassWordCnt"`
	BlockWordCnt   int64 `json:"BlockWordCnt"`
	ReviewWordCnt  int64 `json:"ReviewWordCnt"`
}

type TextBaseTrendRecordRsp struct {
	TextBaseStatisticsRsp
	StartTime int64 `json:"StartTime"`
}

type ContentRiskBaseCountRsp struct {
	All    []map[string]int64 `json:"All"`
	Block  []map[string]int64 `json:"Block"`
	Review []map[string]int64 `json:"Review"`
}

type ImageRiskStatisticsRsp struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      ImageRiskStatisticsInfo `json:"Data"`
}

type ImageRiskStatisticsInfo struct {
	ImageTotal           ImageBaseStatisticsRsp     `json:"ImageTotal"`
	ImageTrendRecord     []*ImageBaseTrendRecordRsp `json:"ImageTrendRecord"`
	ImageRiskTotal       ContentRiskBaseCountRsp    `json:"ImageRiskTotal"`
	ImageRiskTrendRecord ContentRiskBaseCountRsp    `json:"ImageRiskTrendRecord"`
}

type ImageBaseStatisticsRsp struct {
	CommonStatisticsInfo
	FrameStatisticsInfo
	RequestFrameCnt int64 `json:"RequestFrameCnt"`
}

type ImageBaseTrendRecordRsp struct {
	StartTime       int64 `json:"StartTime"`
	RequestFrameCnt int64 `json:"RequestFrameCnt"`
	CommonStatisticsInfo
	FrameStatisticsInfo
}

type VideoRiskStatisticsRsp struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      VideoRiskStatisticsInfo `json:"Data"`
}

type VideoRiskStatisticsInfo struct {
	VideoTotal           VideoBaseStatisticsRsp     `json:"VideoTotal"`
	VideoTrendRecord     []*VideoBaseTrendRecordRsp `json:"VideoTrendRecord"`
	VideoRiskTotal       VideoContentRiskCountRsp   `json:"VideoRiskTotal"`
	VideoRiskTrendRecord VideoContentRiskCountRsp   `json:"VideoRiskTrendRecord"`
}

type VideoContentRiskCountRsp struct {
	RequestStatistics ContentRiskBaseCountRsp `json:"Request"`
	FrameStatistics   ContentRiskBaseCountRsp `json:"Frame"`
	SplitStatistics   ContentRiskBaseCountRsp `json:"Split"`
}

type AudioContentRiskCountRsp struct {
	RequestStatistics ContentRiskBaseCountRsp `json:"Request"`
	SplitStatistics   ContentRiskBaseCountRsp `json:"Split"`
}

type VideoBaseStatisticsRsp struct {
	CommonStatisticsInfo
	SplitStatisticsInfo
	FrameStatisticsInfo
	FrameCnt int64 `json:"FrameCnt"`
}

type VideoBaseTrendRecordRsp struct {
	StartTime int64 `json:"StartTime"`
	CommonStatisticsInfo
	SplitStatisticsInfo
	FrameStatisticsInfo
	FrameCnt int64 `json:"FrameCnt"`
}

type AudioRiskStatisticsRsp struct {
	RequestId string                  `json:"RequestId"`
	Code      int                     `json:"Code"`
	Message   string                  `json:"Message"`
	Data      AudioRiskStatisticsInfo `json:"Data"`
}

type AudioRiskStatisticsInfo struct {
	AudioTotal           AudioBaseStatisticsRsp     `json:"AudioTotal"`
	AudioTrendRecord     []*AudioBaseTrendRecordRsp `json:"AudioTrendRecord"`
	AudioRiskTotal       AudioContentRiskCountRsp   `json:"AudioRiskTotal"`
	AudioRiskTrendRecord AudioContentRiskCountRsp   `json:"AudioRiskTrendRecord"`
}

type AudioBaseStatisticsRsp struct {
	CommonStatisticsInfo
	SplitStatisticsInfo
}

type AudioBaseTrendRecordRsp struct {
	StartTime int64 `json:"StartTime"`
	CommonStatisticsInfo
	SplitStatisticsInfo
}
