package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pierrec/lz4"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

const (
	rawBodySizeHeader = "x-tls-bodyrawsize"
)

func (c *LsClient) PutLogs(request *PutLogsRequest) (r *CommonResponse, e error) {
	if len(request.LogBody.LogGroups) == 0 {
		return nil, nil
	}

	for _, logGroup := range request.LogBody.LogGroups {
		if logGroup.Logs != nil {
			for _, log := range logGroup.Logs {
				if log.Time == 0 {
					log.Time = time.Now().Unix()
				}
			}
		}
	}

	params := map[string]string{
		"TopicId": request.TopicID,
	}

	bodyBytes, originalLength, err := GetPutLogsBody(request.CompressType, request.LogBody)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		rawBodySizeHeader:    strconv.Itoa(originalLength),
		"Content-Type":       "application/x-protobuf",
		"x-tls-hashkey":      request.HashKey,
		"x-tls-compresstype": request.CompressType,
	}

	rawResponse, err := c.Request(http.MethodPost, PathPutLogs, params, c.assembleHeader(request.CommonRequest, headers), bodyBytes)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	if rawResponse.StatusCode != http.StatusOK {
		return nil, errors.New(string(responseBody))
	}

	response := &CommonResponse{}
	response.FillRequestId(rawResponse)

	return response, nil
}

func lz4Decompress(input []byte, rawLength int64) ([]byte, error) {
	decompressedBuffer := make([]byte, rawLength)
	if _, err := lz4.UncompressBlock(input, decompressedBuffer); err != nil {
		return nil, err
	}

	return decompressedBuffer, nil
}

func parseLogList(input []byte, compression string, rawSize int64) (*pb.LogGroupList, error) {
	var (
		decompressed []byte
		err          error
	)

	switch strings.ToLower(compression) {
	case "lz4":
		decompressed, err = lz4Decompress(input, rawSize)
		if err != nil {
			return nil, err
		}
	default:
		decompressed = input
	}

	list := &pb.LogGroupList{}
	if err := proto.Unmarshal(decompressed, list); err != nil {
		return nil, err
	}

	return list, nil
}

func (c *LsClient) ConsumeLogs(request *ConsumeLogsRequest) (r *ConsumeLogsResponse, e error) {
	params := map[string]string{
		"TopicId": request.TopicID,
		"ShardId": strconv.Itoa(request.ShardID),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"Cursor": request.Cursor,
	}

	if request.EndCursor != nil {
		reqBody["EndCursor"] = *request.EndCursor
	}

	if request.LogGroupCount != nil {
		reqBody["LogGroupCount"] = *request.LogGroupCount
	}

	if request.Compression != nil {
		reqBody["Compression"] = *request.Compression
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathConsumeLogs, params, c.assembleHeader(request.CommonRequest, headers), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	if rawResponse.StatusCode != http.StatusOK {
		return nil, errors.New(string(responseBody))
	}

	response := &ConsumeLogsResponse{}
	response.FillRequestId(rawResponse)

	var (
		rawSize     int64
		compression string
	)

	if request.Compression != nil {
		compression = *request.Compression
		rawSize, err = strconv.ParseInt(rawResponse.Header.Get(rawBodySizeHeader), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	list, err := parseLogList(responseBody, compression, rawSize)
	if err != nil {
		return nil, err
	}

	response.Logs = list

	parsedCount, err := strconv.ParseInt(rawResponse.Header.Get("x-tls-count"), 10, 64)
	if err != nil {
		return nil, err
	}

	response.Count = int(parsedCount)
	response.Cursor = rawResponse.Header.Get("x-tls-cursor")

	return response, nil
}

func (c *LsClient) DescribeCursor(request *DescribeCursorRequest) (r *DescribeCursorResponse, e error) {
	params := map[string]string{
		"TopicId": request.TopicID,
		"ShardId": strconv.Itoa(request.ShardID),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"From": request.From,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PathDescribeCursor, params, c.assembleHeader(request.CommonRequest, headers), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeCursorResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *LsClient) DescribeLogContext(request *DescribeLogContextRequest) (r *DescribeLogContextResponse, e error) {
	params := map[string]string{}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]interface{}{
		"TopicId":       request.TopicId,
		"ContextFlow":   request.ContextFlow,
		"PackageOffset": request.PackageOffset,
		"Source":        request.Source,
	}

	if request.PrevLogs != nil {
		reqBody["PrevLogs"] = *request.PrevLogs
	}

	if request.NextLogs != nil {
		reqBody["NextLogs"] = *request.NextLogs
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodPost, PathDescribeLogContext, params, c.assembleHeader(request.CommonRequest, headers), bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &DescribeLogContextResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
