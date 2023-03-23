package businessSecurity

import (
	"encoding/json"
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

type DecisionData struct {
	Score      int        `json:"Score"`
	Tags       []string   `json:"Tags"`
	Detail     string     `json:"Detail"`
	DeviceInfo DeviceInfo `json:"DeviceInfo"`
}

type DeviceInfo struct {
	DevSecID      string `json:"DevSecID"`
	TokenCreateTs int64  `json:"TokenCreateTs"`
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

type RequestData struct {
	PassThrough interface{} `json:"PassThrough"`
}

type VideoResultRequest struct {
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

type VideoResp struct {
	DataId       string        `json:"DataId"`
	VideoResult  VideoResult   `json:"VideoResults"`
	AudioResults AudioResultV2 `json:"AudioResults"`
	PassThrough  interface{}   `json:"PassThrough"`
}

type AudioResultV2 struct {
	Decision    string           `json:"Decision"`
	DataId      string           `json:"DataId"`
	Details     []*AudioDetailV2 `json:"Details"`
	PassThrough interface{}      `json:"PassThrough"`
}

type AudioDetailV2 struct {
	StartTime    int              `json:"StartTime"`
	EndTime      int              `json:"EndTime"`
	FrameUrl     string           `json:"FrameUrl"`
	AudioText    string           `json:"AudioText"`
	FrameResults []*FrameResultV2 `json:"FrameResults"`
}

type FrameResultV2 struct {
	Label    string     `json:"Label"`
	SubLabel string     `json:"SubLabel"`
	Decision string     `json:"Decision"`
	Contexts []*Context `json:"Contexts"`
}

type VideoResult struct {
	Decision string  `json:"Decision"`
	Frames   []Frame `json:"Frames"`
}

type Frame struct {
	Url     string   `json:"Url"`
	Offset  float64  `json:"Offset"`
	Results []Result `json:"Results"`
}

type Result struct {
	Decision string   `json:"Decision"`
	Label    string   `json:"Label"`
	SubLabel string   `json:"SubLabel"`
	Detail   []string `json:"Detail"`
}

type ImageResultResponse struct {
	RequestId string               `json:"RequestId"`
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	ImageResp ImageContentRiskResp `json:"Data"`
}

type ImageContentRiskResp struct {
	DataId   string                 `json:"DataId"`
	Decision string                 `json:"Decision"`
	Results  []*ImageContentRiskTag `json:"Results"`
	Scores   map[string]float64     `json:"Scores"`
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
	//Id float64 `json:"id"`
	//Url   string  `json:"url"`
	//Score float64 `json:"score"`
}

type AudioResultResponse struct {
	RequestId string        `json:"RequestId"`
	Code      int           `json:"Code"`
	Message   string        `json:"Message"`
	AudioResp AudioResultV2 `json:"Data"`
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

type TextResultResponse struct {
	RequestId string         `json:"RequestId"`
	Code      int            `json:"Code"`
	Message   string         `json:"Message"`
	VideoResp TextRiskRespV2 `json:"Data"`
}

type TextRiskRespV2 struct {
	Decision string   `json:"Decision"`
	Result   []*Label `json:"Result"`
}

type TextSliceResultResponse struct {
	RequestId string         `json:"RequestId"`
	Code      int            `json:"Code"`
	Message   string         `json:"Message"`
	VideoResp TextRiskRespV3 `json:"Data"`
}

type TextRiskRespV3 struct {
	Decision   string        `json:"Decision"`
	FinalLabel string        `json:"FinalLabel"`
	TextCount  int           `json:"TextCount"`
	Results    []*TextResult `json:"Results"`
}

type TextResult struct {
	RiskText   string   `json:"RiskText"`
	RTStartPos int      `json:"RTStartPos"`
	RTEndPos   int      `json:"RTEndPos"`
	Labels     []*Label `json:"Labels"`
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
