package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/volcengine/volc-sdk-golang/service/sami"
)

const (
	// dump output
	dataOutputFile    = "output.wav"
	payloadOutputFile = "output.json"
	isDump            = true
)

func main() {
	appKey := "appKey"
	token := "token"

	req := sami.SpeechSynthesisReq{
		Text:    "欢迎使用文本转语音服务。",
		Speaker: "speaker",
		Config: sami.AudioConfig{
			Format:          "wav",
			SampleRate:      16000,
			SpeechRate:      0,
			PitchRate:       0,
			EnableTimestamp: true,
		},
	}
	resp, status, err := sami.DefaultInstance.TTS(req, appKey, token)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
	// parse payload to SpeechSynthesisResponse
	payloadStr := *resp.Payload
	speechSynthesisResponse := new(sami.SpeechSynthesisResponse)
	json.Unmarshal([]byte(payloadStr), &speechSynthesisResponse)
	if isDump && resp.Payload != nil {
		_ = ioutil.WriteFile(payloadOutputFile, []byte(*resp.Payload), 0644)
	}
	if isDump && len(resp.Data) > 0 {
		_ = ioutil.WriteFile(dataOutputFile, resp.Data, 0644)
	}
}
