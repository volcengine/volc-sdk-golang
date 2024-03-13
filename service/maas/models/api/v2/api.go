package api

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MessageContent struct {
	// ImageUrl corresponds to the JSON schema field "image_url".
	ImageUrl *MessageImageContent `json:"image_url,omitempty" yaml:"image_url,omitempty" mapstructure:"image_url,omitempty"`

	// Text corresponds to the JSON schema field "text".
	Text string `json:"text,omitempty" yaml:"text,omitempty" mapstructure:"text,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

type MessageImageContent struct {
	// Detail corresponds to the JSON schema field "detail".
	Detail string `json:"detail,omitempty" yaml:"detail,omitempty" mapstructure:"detail,omitempty"`

	// ImageBytes corresponds to the JSON schema field "image_bytes".
	ImageBytes string `json:"image_bytes,omitempty" yaml:"image_bytes,omitempty" mapstructure:"image_bytes,omitempty"`

	// Url corresponds to the JSON schema field "url".
	Url string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`
}

type ChatRole string

const ChatRoleAssistant ChatRole = "assistant"
const ChatRoleFunction ChatRole = "function"
const ChatRoleSystem ChatRole = "system"
const ChatRoleUser ChatRole = "user"

type Usage struct {
	// CompletionTokens corresponds to the JSON schema field "completion_tokens".
	CompletionTokens int `json:"completion_tokens,omitempty" yaml:"completion_tokens,omitempty" mapstructure:"completion_tokens,omitempty"`

	// PromptTokens corresponds to the JSON schema field "prompt_tokens".
	PromptTokens int `json:"prompt_tokens,omitempty" yaml:"prompt_tokens,omitempty" mapstructure:"prompt_tokens,omitempty"`

	// TotalTokens corresponds to the JSON schema field "total_tokens".
	TotalTokens int `json:"total_tokens,omitempty" yaml:"total_tokens,omitempty" mapstructure:"total_tokens,omitempty"`
}

type Error struct {
	// Code corresponds to the JSON schema field "code".
	Code string `json:"code" yaml:"code" mapstructure:"code"`

	// CodeN corresponds to the JSON schema field "code_n".
	CodeN int `json:"code_n" yaml:"code_n" mapstructure:"code_n"`

	// Message corresponds to the JSON schema field "message".
	Message string `json:"message" yaml:"message" mapstructure:"message"`

	ReqId string `json:"req_id" yaml:"req_id" mapstructure:"req_id"`
}

func (m *Error) Error() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s: %s, code=%d, req_id=%s", m.Code, m.Message, m.CodeN, m.ReqId)
}

func NewClientSDKRequestError(raw string, reqId string) *Error {
	return &Error{
		Code:    "ClientSDKRequestError",
		CodeN:   1709701,
		Message: fmt.Sprintf("MaaS SDK request error: %s", raw),
		ReqId:   reqId,
	}
}

type FunctionCall struct {
	// Arguments corresponds to the JSON schema field "arguments".
	Arguments string `json:"arguments,omitempty" yaml:"arguments,omitempty" mapstructure:"arguments,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}

type Logprobs struct {
	// TextOffset corresponds to the JSON schema field "text_offset".
	TextOffset []int `json:"text_offset,omitempty" yaml:"text_offset,omitempty" mapstructure:"text_offset,omitempty"`

	// TokenLogprobs corresponds to the JSON schema field "token_logprobs".
	TokenLogprobs []float64 `json:"token_logprobs,omitempty" yaml:"token_logprobs,omitempty" mapstructure:"token_logprobs,omitempty"`

	// Tokens corresponds to the JSON schema field "tokens".
	Tokens []string `json:"tokens,omitempty" yaml:"tokens,omitempty" mapstructure:"tokens,omitempty"`

	// TopLogprobs corresponds to the JSON schema field "top_logprobs".
	TopLogprobs []LogprobsTopLogprobsElem `json:"top_logprobs,omitempty" yaml:"top_logprobs,omitempty" mapstructure:"top_logprobs,omitempty"`
}

type LogprobsTopLogprobsElem map[string]interface{}

type Choice struct {
	// Action corresponds to the JSON schema field "action".
	Action *ChoiceLog `json:"action,omitempty" yaml:"action,omitempty" mapstructure:"action,omitempty"`

	// FinishReason corresponds to the JSON schema field "finish_reason".
	FinishReason string `json:"finish_reason,omitempty" yaml:"finish_reason,omitempty" mapstructure:"finish_reason,omitempty"`

	// Index corresponds to the JSON schema field "index".
	Index int `json:"index,omitempty" yaml:"index,omitempty" mapstructure:"index,omitempty"`

	// Logprobs corresponds to the JSON schema field "logprobs".
	Logprobs *Logprobs `json:"logprobs,omitempty" yaml:"logprobs,omitempty" mapstructure:"logprobs,omitempty"`

	// Message corresponds to the JSON schema field "message".
	Message *Message `json:"message,omitempty" yaml:"message,omitempty" mapstructure:"message,omitempty"`

	// Observation corresponds to the JSON schema field "observation".
	Observation *ChoiceLog `json:"observation,omitempty" yaml:"observation,omitempty" mapstructure:"observation,omitempty"`

	// Thought corresponds to the JSON schema field "thought".
	Thought *ChoiceLog `json:"thought,omitempty" yaml:"thought,omitempty" mapstructure:"thought,omitempty"`
}

type ChoiceLog struct {
	// Content corresponds to the JSON schema field "content".
	Content string `json:"content,omitempty" yaml:"content,omitempty" mapstructure:"content,omitempty"`

	// Input corresponds to the JSON schema field "input".
	Input string `json:"input,omitempty" yaml:"input,omitempty" mapstructure:"input,omitempty"`
}

type Message struct {
	// Content corresponds to the JSON schema field "content".
	Content interface{} `json:"content,omitempty" yaml:"content,omitempty" mapstructure:"content,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// References corresponds to the JSON schema field "references".
	References []*Reference `json:"references,omitempty" yaml:"references,omitempty" mapstructure:"references,omitempty"`

	// Role corresponds to the JSON schema field "role".
	Role ChatRole `json:"role,omitempty" yaml:"role,omitempty" mapstructure:"role,omitempty"`

	// ToolCallId corresponds to the JSON schema field "tool_call_id".
	ToolCallId string `json:"tool_call_id,omitempty" yaml:"tool_call_id,omitempty" mapstructure:"tool_call_id,omitempty"`

	// ToolCalls corresponds to the JSON schema field "tool_calls".
	ToolCalls []*ToolCall `json:"tool_calls,omitempty" yaml:"tool_calls,omitempty" mapstructure:"tool_calls,omitempty"`
}

type ChatReq struct {
	// CryptoToken corresponds to the JSON schema field "crypto_token".
	CryptoToken string `json:"crypto_token,omitempty" yaml:"crypto_token,omitempty" mapstructure:"crypto_token,omitempty"`

	// Extra corresponds to the JSON schema field "extra".
	Extra RequestExtra `json:"extra,omitempty" yaml:"extra,omitempty" mapstructure:"extra,omitempty"`

	// Messages corresponds to the JSON schema field "messages".
	Messages []*Message `json:"messages,omitempty" yaml:"messages,omitempty" mapstructure:"messages,omitempty"`

	// Parameters corresponds to the JSON schema field "parameters".
	Parameters *Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Stream corresponds to the JSON schema field "stream".
	Stream bool `json:"stream,omitempty" yaml:"stream,omitempty" mapstructure:"stream,omitempty"`

	// Tools corresponds to the JSON schema field "tools".
	Tools []*Tool `json:"tools,omitempty" yaml:"tools,omitempty" mapstructure:"tools,omitempty"`

	// User corresponds to the JSON schema field "user".
	User string `json:"user,omitempty" yaml:"user,omitempty" mapstructure:"user,omitempty"`
}

type RequestExtra map[string]string

type ChatResp struct {
	// Choices corresponds to the JSON schema field "choices".
	Choices []*Choice `json:"choices,omitempty" yaml:"choices,omitempty" mapstructure:"choices,omitempty"`

	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`

	// Extra corresponds to the JSON schema field "extra".
	Extra ChatRespExtra `json:"extra,omitempty" yaml:"extra,omitempty" mapstructure:"extra,omitempty"`

	// Usage corresponds to the JSON schema field "usage".
	Usage *Usage `json:"usage,omitempty" yaml:"usage,omitempty" mapstructure:"usage,omitempty"`

	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`
}

type ChatRespExtra map[string]string

type Parameters struct {
	// DoSample corresponds to the JSON schema field "do_sample".
	DoSample bool `json:"do_sample,omitempty" yaml:"do_sample,omitempty" mapstructure:"do_sample,omitempty"`

	// FrequencyPenalty corresponds to the JSON schema field "frequency_penalty".
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty" yaml:"frequency_penalty,omitempty" mapstructure:"frequency_penalty,omitempty"`

	// Logprobs corresponds to the JSON schema field "logprobs".
	Logprobs int `json:"logprobs,omitempty" yaml:"logprobs,omitempty" mapstructure:"logprobs,omitempty"`

	// MaxNewTokens corresponds to the JSON schema field "max_new_tokens".
	MaxNewTokens int `json:"max_new_tokens,omitempty" yaml:"max_new_tokens,omitempty" mapstructure:"max_new_tokens,omitempty"`

	// MaxPromptTokens corresponds to the JSON schema field "max_prompt_tokens".
	MaxPromptTokens int `json:"max_prompt_tokens,omitempty" yaml:"max_prompt_tokens,omitempty" mapstructure:"max_prompt_tokens,omitempty"`

	// MaxTokens corresponds to the JSON schema field "max_tokens".
	MaxTokens int `json:"max_tokens,omitempty" yaml:"max_tokens,omitempty" mapstructure:"max_tokens,omitempty"`

	// MinNewTokens corresponds to the JSON schema field "min_new_tokens".
	MinNewTokens int `json:"min_new_tokens,omitempty" yaml:"min_new_tokens,omitempty" mapstructure:"min_new_tokens,omitempty"`

	// PresencePenalty corresponds to the JSON schema field "presence_penalty".
	PresencePenalty float64 `json:"presence_penalty,omitempty" yaml:"presence_penalty,omitempty" mapstructure:"presence_penalty,omitempty"`

	// RepetitionPenalty corresponds to the JSON schema field "repetition_penalty".
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty" yaml:"repetition_penalty,omitempty" mapstructure:"repetition_penalty,omitempty"`

	// Stop corresponds to the JSON schema field "stop".
	Stop []string `json:"stop,omitempty" yaml:"stop,omitempty" mapstructure:"stop,omitempty"`

	// Temperature corresponds to the JSON schema field "temperature".
	Temperature float64 `json:"temperature,omitempty" yaml:"temperature,omitempty" mapstructure:"temperature,omitempty"`

	// TopK corresponds to the JSON schema field "top_k".
	TopK int `json:"top_k,omitempty" yaml:"top_k,omitempty" mapstructure:"top_k,omitempty"`

	// TopP corresponds to the JSON schema field "top_p".
	TopP float64 `json:"top_p,omitempty" yaml:"top_p,omitempty" mapstructure:"top_p,omitempty"`

	Seed int `json:"seed,omitempty" yaml:"seed,omitempty" mapstructure:"seed,omitempty"`

	Strength float32 `json:"strength,omitempty" yaml:"strength,omitempty" mapstructure:"strength,omitempty"`

	Height int `json:"height,omitempty" yaml:"height,omitempty" mapstructure:"height,omitempty"`

	Width int `json:"width,omitempty" yaml:"width,omitempty" mapstructure:"width,omitempty"`

	NumInferenceSteps int `json:"num_inference_steps,omitempty" yaml:"num_inference_steps,omitempty" mapstructure:"num_inference_steps,omitempty"`
}

type Reference struct {
	// Idx corresponds to the JSON schema field "idx".
	Idx int `json:"idx,omitempty" yaml:"idx,omitempty" mapstructure:"idx,omitempty"`

	// LogoUrl corresponds to the JSON schema field "logo_url".
	LogoUrl string `json:"logo_url,omitempty" yaml:"logo_url,omitempty" mapstructure:"logo_url,omitempty"`

	// PcUrl corresponds to the JSON schema field "pc_url".
	PcUrl string `json:"pc_url,omitempty" yaml:"pc_url,omitempty" mapstructure:"pc_url,omitempty"`

	// SiteName corresponds to the JSON schema field "site_name".
	SiteName string `json:"site_name,omitempty" yaml:"site_name,omitempty" mapstructure:"site_name,omitempty"`

	// Url corresponds to the JSON schema field "url".
	Url string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`
}

var enumValues_ChatRole = []interface{}{
	"system",
	"assistant",
	"user",
	"function",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChatRole) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChatRole {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChatRole, v)
	}
	*j = ChatRole(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Usage) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain Usage
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["completion_tokens"]; !ok || v == nil {
		plain.CompletionTokens = 0.0
	}
	if v, ok := raw["prompt_tokens"]; !ok || v == nil {
		plain.PromptTokens = 0.0
	}
	if v, ok := raw["total_tokens"]; !ok || v == nil {
		plain.TotalTokens = 0.0
	}
	*j = Usage(plain)
	return nil
}

type Function struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	Parameters map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	Examples []string `json:"examples,omitempty" yaml:"examples,omitempty" mapstructure:"examples,omitempty"`
}

type Tool struct {
	// Type corresponds to the JSON schema field "type".
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	Function *Function `json:"function,omitempty" yaml:"function,omitempty" mapstructure:"function,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChatReq) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain ChatReq
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["stream"]; !ok || v == nil {
		plain.Stream = false
	}
	*j = ChatReq(plain)
	return nil
}

type ToolCall struct {
	// Function corresponds to the JSON schema field "function".
	Function *FunctionCall `json:"function,omitempty" yaml:"function,omitempty" mapstructure:"function,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id string `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Error) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["code"]; !ok || v == nil {
		return fmt.Errorf("field code in Error: required")
	}
	if v, ok := raw["code_n"]; !ok || v == nil {
		return fmt.Errorf("field code_n in Error: required")
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message in Error: required")
	}
	type Plain Error
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Error(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChatResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain ChatResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["choices"]; !ok || v == nil {
		plain.Choices = []*Choice{}
	}
	*j = ChatResp(plain)
	return nil
}

type Embedding struct {
	// Embedding corresponds to the JSON schema field "embedding".
	Embedding []float64 `json:"embedding,omitempty" yaml:"embedding,omitempty" mapstructure:"embedding,omitempty"`

	// Index corresponds to the JSON schema field "index".
	Index int `json:"index,omitempty" yaml:"index,omitempty" mapstructure:"index,omitempty"`

	// Object corresponds to the JSON schema field "object".
	Object string `json:"object,omitempty" yaml:"object,omitempty" mapstructure:"object,omitempty"`
}

type LabelLogprobosValue struct {
	// ReqId corresponds to the JSON schema field "req_id".
	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	// TokenLogprobs corresponds to the JSON schema field "token_logprobs".
	TokenLogprobs []float64 `json:"token_logprobs,omitempty" yaml:"token_logprobs,omitempty" mapstructure:"token_logprobs,omitempty"`

	// Tokens corresponds to the JSON schema field "tokens".
	Tokens []string `json:"tokens,omitempty" yaml:"tokens,omitempty" mapstructure:"tokens,omitempty"`
}

type CertReq map[string]interface{}

type CertResp struct {
	// Cert corresponds to the JSON schema field "cert".
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty" mapstructure:"cert,omitempty"`

	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`
}

type ClassificationReq struct {
	// Labels corresponds to the JSON schema field "labels".
	Labels []string `json:"labels,omitempty" yaml:"labels,omitempty" mapstructure:"labels,omitempty"`

	// Query corresponds to the JSON schema field "query".
	Query string `json:"query,omitempty" yaml:"query,omitempty" mapstructure:"query,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LabelLogprobosValue) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain LabelLogprobosValue
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["token_logprobs"]; !ok || v == nil {
		plain.TokenLogprobs = []float64{}
	}
	if v, ok := raw["tokens"]; !ok || v == nil {
		plain.Tokens = []string{}
	}
	*j = LabelLogprobosValue(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Embedding) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain Embedding
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["embedding"]; !ok || v == nil {
		plain.Embedding = []float64{}
	}
	*j = Embedding(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertResp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain CertResp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["cert"]; !ok || v == nil {
		plain.Cert = ""
	}
	*j = CertResp(plain)
	return nil
}

type ClassificationResp struct {
	// Label corresponds to the JSON schema field "label".
	Label string `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`

	// LabelLogprobos corresponds to the JSON schema field "label_logprobos".
	LabelLogprobos ClassificationRespLabelLogprobos `json:"label_logprobos,omitempty" yaml:"label_logprobos,omitempty" mapstructure:"label_logprobos,omitempty"`

	// Usage corresponds to the JSON schema field "usage".
	Usage *Usage `json:"usage,omitempty" yaml:"usage,omitempty" mapstructure:"usage,omitempty"`

	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`
}

type ClassificationRespLabelLogprobos map[string]*LabelLogprobosValue

type EmbeddingsReq struct {
	// EncodingFormat corresponds to the JSON schema field "encoding_format".
	EncodingFormat string `json:"encoding_format,omitempty" yaml:"encoding_format,omitempty" mapstructure:"encoding_format,omitempty"`

	// Input corresponds to the JSON schema field "input".
	Input []string `json:"input,omitempty" yaml:"input,omitempty" mapstructure:"input,omitempty"`

	// User corresponds to the JSON schema field "user".
	User string `json:"user,omitempty" yaml:"user,omitempty" mapstructure:"user,omitempty"`
}

type EmbeddingsResp struct {
	// Data corresponds to the JSON schema field "data".
	Data []Embedding `json:"data,omitempty" yaml:"data,omitempty" mapstructure:"data,omitempty"`

	// Object corresponds to the JSON schema field "object".
	Object string `json:"object,omitempty" yaml:"object,omitempty" mapstructure:"object,omitempty"`

	// Usage corresponds to the JSON schema field "usage".
	Usage *Usage `json:"usage,omitempty" yaml:"usage,omitempty" mapstructure:"usage,omitempty"`

	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`
}

type TokenizeReq struct {
	// Text corresponds to the JSON schema field "text".
	Text string `json:"text,omitempty" yaml:"text,omitempty" mapstructure:"text,omitempty"`
}

type TokenizeResp struct {
	// Tokens corresponds to the JSON schema field "tokens".
	Tokens []string `json:"tokens,omitempty" yaml:"tokens,omitempty" mapstructure:"tokens,omitempty"`

	// TotalTokens corresponds to the JSON schema field "total_tokens".
	TotalTokens int `json:"total_tokens,omitempty" yaml:"total_tokens,omitempty" mapstructure:"total_tokens,omitempty"`

	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`
}

type ImagesParameters struct {
	Seed int `json:"seed,omitempty" yaml:"seed,omitempty" mapstructure:"seed,omitempty"`

	Strength float32 `json:"strength,omitempty" yaml:"strength,omitempty" mapstructure:"strength,omitempty"`

	Height int `json:"height,omitempty" yaml:"height,omitempty" mapstructure:"height,omitempty"`

	Width int `json:"width,omitempty" yaml:"width,omitempty" mapstructure:"width,omitempty"`

	NumInferenceSteps int `json:"num_inference_steps,omitempty" yaml:"num_inference_steps,omitempty" mapstructure:"num_inference_steps,omitempty"`
}

type ImagesQuickGenReq struct {
	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	Prompt string `json:"prompt,omitempty" yaml:"prompt,omitempty" mapstructure:"prompt,omitempty"`

	NegativePrompt string `json:"negative_prompt,omitempty" yaml:"negative_prompt,omitempty" mapstructure:"negative_prompt,omitempty"`

	InitImage []byte `json:"init_image,omitempty" yaml:"init_image,omitempty" mapstructure:"init_image,omitempty"`

	ControlImage []byte `json:"control_image,omitempty" yaml:"control_image,omitempty" mapstructure:"control_image,omitempty"`

	// Parameters corresponds to the JSON schema field "parameters".
	Parameters *ImagesParameters `json:"parameters,omitempty" yaml:"parameters,omitempty" mapstructure:"parameters,omitempty"`
}

type ImageUrl struct {
	// Detail corresponds to the JSON schema field "detail".
	Detail string `json:"detail,omitempty" yaml:"detail,omitempty" mapstructure:"detail,omitempty"`

	// ImageBytes corresponds to the JSON schema field "image_bytes".
	ImageBytes []byte `json:"image_bytes,omitempty" yaml:"image_bytes,omitempty" mapstructure:"image_bytes,omitempty"`

	// Url corresponds to the JSON schema field "url".
	Url string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`
}

type ImagesQuickGenResp struct {
	ReqId string `json:"req_id,omitempty" yaml:"req_id,omitempty" mapstructure:"req_id,omitempty"`

	Data []*ImageUrl `json:"data,omitempty" yaml:"data,omitempty" mapstructure:"data,omitempty"`
}

type ErrorResp struct {
	// Error corresponds to the JSON schema field "error".
	Error *Error `json:"error,omitempty" yaml:"error,omitempty" mapstructure:"error,omitempty"`
}

type SpeechReq struct {
	Input string `json:"input,omitempty" yaml:"input,omitempty" mapstructure:"input,omitempty"`

	Voice string `json:"voice,omitempty" yaml:"voice,omitempty" mapstructure:"voice,omitempty"`

	ResponseFormat string `json:"response_format,omitempty" yaml:"response_format,omitempty" mapstructure:"response_format,omitempty"`

	Speed float64 `json:"speed,omitempty" yaml:"speed,omitempty" mapstructure:"speed,omitempty"`

	Extra RequestExtra `json:"extra,omitempty" yaml:"extra,omitempty" mapstructure:"extra,omitempty"`
}
