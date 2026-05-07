package tls

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type requestIDFiller interface {
	FillRequestId(*http.Response)
}

func (c *LsClient) doProcessorRequest(method, path string, common CommonRequest, params map[string]string, body interface{}, response interface{}) error {
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
	}

	if body == nil {
		body = map[string]string{}
	}
	bytesBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	rawResponse, err := c.Request(method, path, params, c.assembleHeader(common, reqHeaders), bytesBody)
	if err != nil {
		return err
	}
	if rawResponse == nil {
		return fmt.Errorf("raw response is nil")
	}
	defer rawResponse.Body.Close()

	if filler, ok := response.(requestIDFiller); ok {
		filler.FillRequestId(rawResponse)
	}

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}
	if rawResponse.StatusCode != http.StatusOK {
		return errors.New(string(responseBody))
	}
	if response != nil && len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, response); err != nil {
			return err
		}
	}
	return nil
}

func addProcessorStringParam(params map[string]string, key, value string) {
	if len(value) != 0 {
		params[key] = value
	}
}

func addProcessorIntParam(params map[string]string, key string, value int) {
	if value != 0 {
		params[key] = strconv.Itoa(value)
	}
}

func (c *LsClient) CreateProcessor(request *CreateProcessorRequest) (*CreateProcessorResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CreateProcessorResponse{}
	if err := c.doProcessorRequest(http.MethodPost, PathCreateProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DeleteProcessor(request *DeleteProcessorRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodDelete, PathDeleteProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) ModifyProcessor(request *ModifyProcessorRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodPut, PathModifyProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeProcessor(request *DescribeProcessorRequest) (*DescribeProcessorResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	params := map[string]string{"ProcessorId": request.ProcessorID}
	response := &DescribeProcessorResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeProcessor, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeProcessors(request *DescribeProcessorsRequest) (*DescribeProcessorsResponse, error) {
	params := map[string]string{}
	addProcessorStringParam(params, "IamProjectName", request.IamProjectName)
	addProcessorStringParam(params, "ProjectId", request.ProjectID)
	addProcessorStringParam(params, "ProjectName", request.ProjectName)
	addProcessorStringParam(params, "ProcessorId", request.ProcessorID)
	addProcessorStringParam(params, "ProcessorName", request.ProcessorName)
	if request.ProcessorType != nil {
		params["ProcessorType"] = string(*request.ProcessorType)
	}
	if request.ProcessorStatus != nil {
		params["ProcessorStatus"] = string(*request.ProcessorStatus)
	}
	if request.ProcessorDSLType != nil {
		params["ProcessorDSLType"] = string(*request.ProcessorDSLType)
	}
	if request.OrderByProject != nil {
		params["OrderByProject"] = strconv.FormatBool(*request.OrderByProject)
	}
	addProcessorIntParam(params, "PageNumber", request.PageNumber)
	addProcessorIntParam(params, "PageSize", request.PageSize)

	response := &DescribeProcessorsResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeProcessors, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) ExecProcessor(request *ExecProcessorRequest) (*ExecProcessorResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &ExecProcessorResponse{}
	if err := c.doProcessorRequest(http.MethodPost, PathExecProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) OperateProcessor(request *OperateProcessorRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodPut, PathOperateProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeTopicsByProcessor(request *DescribeTopicsByProcessorRequest) (*DescribeTopicsByProcessorResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	params := map[string]string{"ProcessorId": request.ProcessorID}
	addProcessorIntParam(params, "PageNumber", request.PageNumber)
	addProcessorIntParam(params, "PageSize", request.PageSize)

	response := &DescribeTopicsByProcessorResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeTopicsByProcessor, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) BindTopicProcessor(request *BindTopicProcessorRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodPut, PathBindTopicProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) BatchBindTopics(request *BatchBindTopicsRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodPut, PathBatchBindTopics, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) UnbindTopicProcessor(request *UnbindTopicProcessorRequest) (*CommonResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	response := &CommonResponse{}
	if err := c.doProcessorRequest(http.MethodDelete, PathUnbindTopicProcessor, request.CommonRequest, nil, request, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeProcessorByTopic(request *DescribeProcessorByTopicRequest) (*DescribeProcessorResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	params := map[string]string{"TopicId": request.TopicID}
	response := &DescribeProcessorResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeProcessorByTopic, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeProcessorBindings(request *DescribeProcessorBindingsRequest) (*DescribeProcessorBindingsResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	params := map[string]string{"ProjectId": request.ProjectID}
	addProcessorIntParam(params, "PageNumber", request.PageNumber)
	addProcessorIntParam(params, "PageSize", request.PageSize)

	response := &DescribeProcessorBindingsResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeProcessorBindings, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *LsClient) DescribeProcessorFunctions(request *DescribeProcessorFunctionsRequest) (*DescribeProcessorFunctionsResponse, error) {
	if err := request.CheckValidation(); err != nil {
		return nil, NewClientError(err)
	}
	params := map[string]string{"ProcessorType": string(request.ProcessorType)}
	if len(request.ProcessorDSLType) != 0 {
		params["ProcessorDSLType"] = string(request.ProcessorDSLType)
	}

	response := &DescribeProcessorFunctionsResponse{}
	if err := c.doProcessorRequest(http.MethodGet, PathDescribeProcessorFunctions, request.CommonRequest, params, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
