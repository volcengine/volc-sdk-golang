package tls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/pierrec/lz4"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

const (
	rawBodySizeHeader = "x-tls-bodyrawsize"
)

func (c *LsClient) PutLogs(request *PutLogsRequest) (*CommonResponse, error) {
	if len(request.LogBody.LogGroups) == 0 {
		return nil, nil
	}

	params := map[string]string{
		"topic_id": request.TopicID,
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

	rawResponse, err := c.Request(http.MethodPost, PutLogsUrl, params, headers, bodyBytes)
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

func (c *LsClient) PullLogs(request *PullLogsRequest) (*PullLogsResponse, error) {
	params := map[string]string{
		"topic_id": request.TopicID,
		"shard_id": strconv.Itoa(request.ShardID),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"cursor": request.Cursor,
	}

	if request.EndCursor != nil {
		reqBody["end_cursor"] = *request.EndCursor
	}

	if request.Count != nil {
		reqBody["count"] = strconv.Itoa(*request.Count)
	}

	if request.Compression != nil {
		reqBody["compression"] = *request.Compression
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, PullLogsUrl, params, headers, bytesBody)
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

	response := &PullLogsResponse{}
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

func (c *LsClient) GetCursor(request *GetCursorRequest) (*GetCursorResponse, error) {
	params := map[string]string{
		"topic_id": request.TopicID,
		"shard_id": strconv.Itoa(request.ShardID),
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	reqBody := map[string]string{
		"from": request.From,
	}

	bytesBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	rawResponse, err := c.Request(http.MethodGet, GetCursorUrl, params, headers, bytesBody)
	if err != nil {
		return nil, err
	}
	defer rawResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	var response = &GetCursorResponse{}
	response.FillRequestId(rawResponse)

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}
