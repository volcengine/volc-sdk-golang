package vedit

import "github.com/volcengine/volc-sdk-golang/base"

type TemplateParamItem struct {
	Name     string `json:"Name,omitempty"`
	Type     string `json:"Type,omitempty"`
	Position string `json:"Position,omitempty"`
	Source   string `json:"Source,omitempty"`
	Text     string `json:"Text,omitempty"`
}
type SubmitDirectEditTaskRequest struct {
	Uploader     string      `json:"Uploader,omitempty"`
	Application  string      `json:"Application,omitempty"`
	VideoName    string      `json:"VideoName,omitempty"`
	Param        interface{} `json:"EditParam"`
	Priority     int32       `json:"Priority"`
	CallbackUri  string      `json:"CallbackUri,omitempty"`
	CallbackArgs string      `json:"CallbackArgs,omitempty"`
}

type SubmitTemplateTaskRequest struct {
	TemplateId   string                `json:"TemplateId,omitempty"`
	Space        string                `json:"Space,omitempty"`
	VideoName    []string              `json:"VideoName,omitempty"`
	Params       [][]TemplateParamItem `json:"Params,omitempty"`
	Priority     int32                 `json:"Priority"`
	CallbackUri  string                `json:"CallbackUri,omitempty"`
	CallbackArgs string                `json:"CallbackArgs,omitempty"`
}

type GetDirectEditResultRequest struct {
	ReqIds []string `json:"ReqIds"`
}

type SubmitDirectEditTaskAsyncResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           struct {
		ReqId string `json:"ReqId"`
	} `json:"Result"`
}

type SubmitTemplateTaskAsyncResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           []string `json:"Result"`
}

type GetDirectEditResultResponse struct {
	ResponseMetadata *base.ResponseMetadata
	Result           []struct {
		ReqId        string      `json:"ReqId"`
		Param        interface{} `json:"EditParam"`
		Priority     int32       `json:"Priority"`
		CallbackUri  string      `json:"CallbackUri"`
		CallbackArgs string      `json:"CallbackArgs"`
		Status       string      `json:"Status"`
		OutputVid    string      `json:"OutputVid"`
		TaskId       string      `json:"TaskId"`
	} `json:"Result"`
}
