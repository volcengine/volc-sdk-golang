package nlp

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Nlp) NlpPost(action string, query url.Values, req, result interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.Json(action, query, string(body))
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := UnmarshalResultInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

//中文文本纠错
func (c *Nlp) GetTextCorrectionZhCorrect(param *TextCorrectionZhCorrectParam) (*TextCorrectionZhCorrectResult, error) {
	u := url.Values{}
	u.Add("content", param.Content)
	resp := new(TextCorrectionZhCorrectResult)
	err := c.NlpPost("TextCorrectionZhCorrect", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//英文文本纠错
func (c *Nlp) GetTextCorrectionEnCorrect(param *TextCorrectionEnCorrectParam) (*TextCorrectionEnCorrectResult, error) {
	u := url.Values{}
	u.Add("content", param.Content)
	resp := new(TextCorrectionEnCorrectResult)
	err := c.NlpPost("TextCorrectionEnCorrect", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//关键词提取
func (c *Nlp) GetKeyPhraseExtractionExtract(param *KeyPhraseExtractionParam) (*KeyPhraseExtractionResult, error) {
	u := url.Values{}
	u.Add("request_type", param.RequestType)
	u.Add("title", param.Title)
	u.Add("content", param.Content)
	resp := new(KeyPhraseExtractionResult)
	err := c.NlpPost("KeyphraseExtractionExtract", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//情感分析
func (c *Nlp) GetSentimentAnalysis(param *SentimentAnalysisParam) (*SentimentAnalysisResult, error) {
	u := url.Values{}
	u.Add("text", param.Text)
	resp := new(SentimentAnalysisResult)
	err := c.NlpPost("SentimentAnalysis", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//文本摘要
func (c *Nlp) GetTextSummarization(param *TextSummarizationParam) (*TextSummarizationResult, error) {
	u := url.Values{}
	u.Add("text", param.Text)
	resp := new(TextSummarizationResult)
	err := c.NlpPost("TextSummarization", u, param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
