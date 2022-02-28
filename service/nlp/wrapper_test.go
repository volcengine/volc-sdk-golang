package nlp

import (
	"fmt"
	"testing"
)

const (
	testAk = "***REMOVED***"
	testSk = "***REMOVED***"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestGetTextCorrectionZhCorrect(t *testing.T) {
	res, err := DefaultInstance.GetTextCorrectionZhCorrect(&TextCorrectionZhCorrectParam{
		Content: "你好",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestGetTextCorrectionEnCorrect(t *testing.T) {
	res, err := DefaultInstance.GetTextCorrectionEnCorrect(&TextCorrectionEnCorrectParam{
		Content: "hello",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestGetTextKeyPhraseExtractionExtract(t *testing.T) {
	res, err := DefaultInstance.GetKeyPhraseExtractionExtract(&KeyPhraseExtractionParam{
		RequestType: "text",
		Title:       "title1",
		Content:     "hello",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestGetSentimentAnalysis(t *testing.T) {
	res, err := DefaultInstance.GetSentimentAnalysis(&SentimentAnalysisParam{
		Text: "happy",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}

func TestGetTextSummarization(t *testing.T) {
	res, err := DefaultInstance.GetTextSummarization(&TextSummarizationParam{
		Text: "1231",
	})
	fmt.Println(err)
	if res != nil {
		fmt.Println(*res)
	}
}
