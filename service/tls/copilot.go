package tls

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// CreateAppInstance 创建应用实例，返回应用实例ID
func (c *LsClient) CreateAppInstance(request *CreateAppInstanceReq) (*CreateAppInstanceResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateAppInstance, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateAppInstanceResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeAppInstances(request *DescribeAppInstancesReq) (*DescribeAppInstancesResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}
	if request.PageNumber > 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}
	if request.PageSize > 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}
	if request.InstanceName != nil {
		params["InstanceName"] = *request.InstanceName
	}
	if request.InstanceType != nil {
		params["InstanceType"] = string(*request.InstanceType)
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeAppInstances, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeAppInstancesResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DeleteAppInstance(request *DeleteAppInstanceReq) (*DeleteAppInstanceResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteAppInstance, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DeleteAppInstanceResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) CreateAppSceneMeta(request *CreateAppSceneMetaReq) (*CreateAppSceneMetaResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathCreateAppSceneMeta, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &CreateAppSceneMetaResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeAppSceneMetas(request *DescribeAppSceneMetasReq) (*DescribeAppSceneMetasResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	if request.InstanceId != "" {
		params["InstanceId"] = request.InstanceId
	}
	if request.Id != nil {
		params["Id"] = *request.Id
	}
	if request.DescribeAPPMetaType != "" {
		params["DescribeAPPMetaType"] = string(request.DescribeAPPMetaType)
	}
	if request.PageNumber != 0 {
		params["PageNumber"] = strconv.Itoa(request.PageNumber)
	}
	if request.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(request.PageSize)
	}
	if request.PageContext != nil {
		params["PageContext"] = *request.PageContext
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeAppSceneMetas, params, c.assembleHeader(request.CommonRequest, reqHeaders), nil)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeAppSceneMetasResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) ModifyAppSceneMeta(request *ModifyAppSceneMetaReq) (*ModifyAppSceneMetaResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPut, PathModifyAppSceneMeta, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &ModifyAppSceneMetaResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DeleteAppSceneMeta(request *DeleteAppSceneMetaReq) (*DeleteAppSceneMetaResp, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	params := map[string]string{}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodDelete, PathDeleteAppSceneMeta, params, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DeleteAppSceneMetaResp{}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}
	response.FillRequestId(rawResponse)

	return response, nil
}

func (c *LsClient) DescribeSessionAnswer(request *DescribeSessionAnswerReq) (reader *CopilotSSEReader, err error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}

	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "text/event-stream",
	}

	bytesBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathDescribeSessionAnswer, nil, c.assembleHeader(request.CommonRequest, reqHeaders), bytesBody)
	if err != nil {
		return nil, err
	}
	if rawResponse.StatusCode != http.StatusOK {
		errorMessage, err := ioutil.ReadAll(rawResponse.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errorMessage))
	}

	return &CopilotSSEReader{
		reader:   bufio.NewReader(rawResponse.Body),
		response: rawResponse,
		finished: false,
	}, nil
}

type CopilotSSEReader struct {
	reader   *bufio.Reader
	response *http.Response
	finished bool

	once sync.Once
}

// ReadEvent 读取并解析 SSE 事件
func (r *CopilotSSEReader) ReadEvent() (_ *CopilotAnswer, err error) {
	defer func() {
		if err != nil {
			r.once.Do(func() {
				r.Close()
			})
		}
	}()

	if r.finished {
		return nil, io.EOF
	}

	var event string
	var dataBuffer bytes.Buffer

	for {
		line, err1 := r.reader.ReadBytes('\n')
		if err1 != nil {
			if errors.Is(err1, io.EOF) {
				// 处理流结束时的逻辑
				r.finished = true
				return nil, err1
			}
			return nil, err1
		}
		line = bytes.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if bytes.HasPrefix(line, []byte("event:")) {
			event = strings.TrimSpace(string(line[6:]))
		} else if bytes.HasPrefix(line, []byte("data:")) {
			if dataBuffer.Len() > 0 {
				dataBuffer.WriteByte('\n')
			}
			dataBuffer.Write(bytes.TrimSpace(line[5:]))
		}

		if event == "" || dataBuffer.Len() == 0 {
			continue
		}

		// 解析 JSON
		var describeRsp DescribeSessionAnswerResp
		if err := json.Unmarshal(dataBuffer.Bytes(), &describeRsp); err != nil {
			return nil, err
		}

		switch describeRsp.ConversationMessageType {
		case CopilotProgress:
			// 忽略进度消息
			dataBuffer.Reset()
			event = ""
			continue
		case CopilotMessage:
			copilotAnswer := &CopilotAnswer{
				QuestionId:       describeRsp.QuestionId,
				SessionId:        describeRsp.SessionId,
				MessageId:        describeRsp.MessageId,
				RspMsgType:       describeRsp.RspMsgType,
				IntentInfo:       describeRsp.IntentInfo,
				Tools:            describeRsp.Tools,
				RetrievalInfo:    describeRsp.RetrievalInfo,
				Suggestions:      describeRsp.Suggestions,
				ReasoningContent: describeRsp.ReasoningContent,
				StageInfo:        describeRsp.StageInfo,
			}
			if describeRsp.Message != nil {
				copilotAnswer.ModelAnswer = &ModelAnswer{
					Answer:     describeRsp.Message.Answer,
					PassDetect: describeRsp.Message.PassDetect,
				}
			}
			return copilotAnswer, nil
		case CopilotError:
			if describeRsp.Message != nil {
				return nil, fmt.Errorf("error occured, errorCode: %s, errorDetail: %s", describeRsp.Message.MessageId, describeRsp.Message.Answer)
			} else {
				return nil, fmt.Errorf("error occured, originRsp: %s", dataBuffer.Bytes())
			}
		}
	}
}

// Close 关闭 SSE 连接
func (r *CopilotSSEReader) Close() error {
	return r.response.Body.Close()
}
