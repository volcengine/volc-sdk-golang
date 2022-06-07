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
	Score  int      `json:"Score"`
	Tags   []string `json:"Tags"`
	Detail string   `json:"Detail"`
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
	RequestId string `json:"RequestId"`
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
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
	VideoResult VideoResult `json:"VideoResult"`
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
	RequestId string    `json:"RequestId"`
	Code      int       `json:"Code"`
	Message   string    `json:"Message"`
	VideoResp VideoResp `json:"Data"`
}

type AudioResult struct {
	Decision string         `json:"Decision"`
	Details  []*AudioDetail `json:"Details"`
}

type AudioDetail struct {
	StartTime    float64        `json:"StartTime"`
	EndTime      float64        `json:"EndTime"`
	FrameUrl     string         `json:"FrameUrl"`
	AudioText    string         `json:"AudioText"`
	FrameID      int            `json:"FrameID"`
	FrameResults []*FrameResult `json:"FrameResults"`
}

type FrameResult struct {
	Label        string   `json:"Label"`
	SubLabel     string   `json:"SubLabel"`
	Decision     string   `json:"Decision"`
	MatchedWords []string `json:"MatchedWords"`
	LibId        string   `json:"LibId"`
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

type Label struct {
	Label    string    `json:"Label"`
	SubLabel string    `json:"SubLabel"`
	Decision string    `json:"Decision"`
	Contexts []Context `json:"Contexts"`
}

type Context struct {
	MatchedWords []string `json:"MatchedWords"`
	LibName      string   `json:"LibName"`
	Positions    Position `json:"Positions"`
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
	AppId      int64  `json:"AppId"`
	Service    string `json:"Service"`
	Parameters string `json:"Parameters"`
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

type NewCustomContentsReq struct {
	AppID       int64  `json:"app_id" form:"app_id" query:"app_id"`
	Service     string `json:"service" form:"service" query:"service"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	Decision    string `json:"decision" form:"decision" query:"decision"`
	MatchType   int    `json:"match_type" form:"match_type" query:"match_type"`
}

type UpdateContentReq struct {
	Contents   []string `json:"contents" form:"contents" query:"contents"`
	ModifyType int      `json:"modify_type" form:"modify_type" query:"modify_type"`
	AppID      int64    `json:"app_id" form:"app_id" query:"app_id"`
	Name       string   `json:"name" form:"name" query:"name"`
	IsFile     bool     `json:"is_file" form:"is_file" query:"is_file"`
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
