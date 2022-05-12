package nlp

import (
	"encoding/json"
	"fmt"
)

type ErrorResult struct {
	Message string `json:"message"`
}

type CharResult struct {
	Correct    string  `json:"correct"`
	Original   string  `json:"original"`
	IndexOfDoc int     `json:"index_of_doc"`
	Confidence float64 `json:"confidence"`
}

type WordResult struct {
	Correct    string  `json:"correct"`
	Original   string  `json:"original"`
	IndexOfDoc int     `json:"index_of_doc"`
	Confidence float64 `json:"confidence"`
}

type TextCorrectionZhCorrectParam struct {
	Content string `json:"content"`
}

type TextCorrectionZhCorrectResult struct {
	CharResult []CharResult `json:"char_result"`
	WordResult []WordResult `json:"word_result"`
}

type KeywordItem struct {
	Text     string   `json:"text"`
	Weight   float64  `json:"weight"`
	Fullname string   `json:"fullname"`
	Category []string `json:"category"`
}

type KeyPhraseExtractionParam struct {
	RequestType string `json:"request_type"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

type KeyPhraseExtractionResult struct {
	Keywords []KeywordItem `json:"keywords"`
}

type TextCorrectionEnCorrectParam struct {
	Content string `json:"content"`
}

type EnCorrectResultItem struct {
	Correct    string `json:"correct"`
	Original   string `json:"original"`
	IndexOfDoc int    `json:"index_of_doc"`
}

type TextCorrectionEnCorrectResult struct {
	Result []EnCorrectResultItem `json:"result"`
}

type SentimentAnalysisParam struct {
	Text string `json:"text"`
}

type SentimentAnalysisResult struct {
	Result        string `json:"result"`
	PositiveScore string `json:"positive_score"`
	NegativeScore string `json:"negative_score"`
}

type TextSummarizationParam struct {
	Text string `json:"text"`
}

type TextSummarizationResult struct {
	Result string `json:"result"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(ErrorResult)
	if err := json.Unmarshal(data, resp); err == nil && resp.Message != "" {
		return fmt.Errorf("request error %s", resp.Message)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}
