package live_v20230101

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	common "github.com/volcengine/volc-sdk-golang/base"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

func skipEmptyValue() queryMarshalFilter {
	return func(_ string, value []string) (accept bool) {
		if len(value) == 0 {
			return false
		}

		for _, item := range value {
			if item == "" {
				return false
			}
		}

		return true
	}
}

func marshalToQuery(model interface{}, filters ...queryMarshalFilter) (url.Values, error) {
	ret := url.Values{}
	if model == nil {
		return ret, nil
	}

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	if modelValue.IsNil() {
		return ret, nil
	}

	if modelType.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
		modelType = modelValue.Type()
	} else {
		return ret, fmt.Errorf("not struct")
	}
	fieldCount := modelType.NumField()

	for i := 0; i < fieldCount; i++ {
		field := modelType.Field(i)
		queryKey := field.Tag.Get("query")
		if queryKey == "" {
			continue
		}

		queryRawValue := modelValue.FieldByName(field.Name)
		queryValues := make([]string, 0)
		if field.Type.Kind() != reflect.Array && field.Type.Kind() != reflect.Slice {
			value := resolveQueryValue(queryRawValue)
			if value == nil {
				continue
			}
			queryValues = append(queryValues, fmt.Sprintf("%v", value))
		} else {
			length := queryRawValue.Len()
			for idx := 0; idx < length; idx++ {
				value := resolveQueryValue(queryRawValue.Index(idx))
				if value == nil {
					continue
				}
				queryValues = append(queryValues, fmt.Sprintf("%v", value))
			}
		}

		for _, fun := range filters {
			if !fun(queryKey, queryValues) {
				goto afterAddQuery
			}
		}

		for _, t := range queryValues {
			ret.Add(queryKey, t)
		}
	afterAddQuery:
	}

	return ret, nil
}

func resolveQueryValue(queryRawValue reflect.Value) interface{} {
	if queryRawValue.Kind() == reflect.Ptr {
		if queryRawValue.IsNil() {
			return nil
		}
		decayedQueryRawValue := queryRawValue.Elem()
		decayedReflectValue := decayedQueryRawValue.Interface()
		return fmt.Sprintf("%v", decayedReflectValue)
	} else {
		queryReflectValue := queryRawValue.Interface()
		return fmt.Sprintf("%v", queryReflectValue)
	}
}

func marshalToJson(model interface{}) ([]byte, error) {
	if model == nil {
		return make([]byte, 0), nil
	}

	result, err := json.Marshal(model)
	if err != nil {
		return []byte{}, fmt.Errorf("can not marshal model to json, %v", err)
	}
	return result, nil
}

func unmarshalResultInto(data []byte, result interface{}) error {
	resp := new(common.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

func (c *Live) DeleteTranscodePreset(ctx context.Context, arg *DeleteTranscodePresetBody) (*DeleteTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateTranscodePreset(ctx context.Context, arg *UpdateTranscodePresetBody) (*UpdateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCommonTransPresetDetail(ctx context.Context, arg *ListCommonTransPresetDetailBody) (*ListCommonTransPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCommonTransPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCommonTransPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) TranscodingJobStatus(ctx context.Context, arg *TranscodingJobStatusQuery) (*TranscodingJobStatusRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "TranscodingJobStatus", query)
	if err != nil {
		return nil, err
	}

	result := new(TranscodingJobStatusRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostTransCodePreset(ctx context.Context, arg *ListVhostTransCodePresetBody) (*ListVhostTransCodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostTransCodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostTransCodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTranscodePreset(ctx context.Context, arg *CreateTranscodePresetBody) (*CreateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) RestartTranscodingJob(ctx context.Context, arg *RestartTranscodingJobQuery) (*RestartTranscodingJobRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "RestartTranscodingJob", query)
	if err != nil {
		return nil, err
	}

	result := new(RestartTranscodingJobRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteWatermarkPresetV2(ctx context.Context, arg *DeleteWatermarkPresetV2Body) (*DeleteWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateWatermarkPresetV2(ctx context.Context, arg *UpdateWatermarkPresetV2Body) (*UpdateWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListWatermarkPresetDetail(ctx context.Context, arg *ListWatermarkPresetDetailBody) (*ListWatermarkPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListWatermarkPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListWatermarkPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateWatermarkPresetV2(ctx context.Context, arg *CreateWatermarkPresetV2Body) (*CreateWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateWatermarkPreset(ctx context.Context, arg *CreateWatermarkPresetBody) (*CreateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateWatermarkPreset(ctx context.Context, arg *UpdateWatermarkPresetBody) (*UpdateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteWatermarkPreset(ctx context.Context, arg *DeleteWatermarkPresetBody) (*DeleteWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListWatermarkPreset(ctx context.Context, arg *ListWatermarkPresetBody) (*ListWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostWatermarkPreset(ctx context.Context, arg *ListVhostWatermarkPresetBody) (*ListVhostWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullRecordTask(ctx context.Context, arg *StopPullRecordTaskBody) (*StopPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateLiveStreamRecordIndexFiles(ctx context.Context, arg *CreateLiveStreamRecordIndexFilesBody) (*CreateLiveStreamRecordIndexFilesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateLiveStreamRecordIndexFiles", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateLiveStreamRecordIndexFilesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullRecordTask(ctx context.Context, arg *CreatePullRecordTaskBody) (*CreatePullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRecordPreset(ctx context.Context, arg *DeleteRecordPresetBody) (*DeleteRecordPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRecordPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRecordPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRecordPresetV2(ctx context.Context, arg *UpdateRecordPresetV2Body) (*UpdateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetPullRecordTask(ctx context.Context, arg *GetPullRecordTaskBody) (*GetPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRecordTaskFileHistory(ctx context.Context, arg *DescribeRecordTaskFileHistoryBody) (*DescribeRecordTaskFileHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRecordTaskFileHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRecordTaskFileHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostRecordPresetV2(ctx context.Context, arg *ListVhostRecordPresetV2Body) (*ListVhostRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullRecordTask(ctx context.Context, arg *ListPullRecordTaskBody) (*ListPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateRecordPresetV2(ctx context.Context, arg *CreateRecordPresetV2Body) (*CreateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSnapshotPreset(ctx context.Context, arg *DeleteSnapshotPresetBody) (*DeleteSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSnapshotPresetV2(ctx context.Context, arg *UpdateSnapshotPresetV2Body) (*UpdateSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCDNSnapshotHistory(ctx context.Context, arg *DescribeCDNSnapshotHistoryBody) (*DescribeCDNSnapshotHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCDNSnapshotHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCDNSnapshotHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSnapshotPresetV2(ctx context.Context, arg *ListVhostSnapshotPresetV2Body) (*ListVhostSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotPresetV2(ctx context.Context, arg *CreateSnapshotPresetV2Body) (*CreateSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTimeShiftPresetV3(ctx context.Context, arg *DeleteTimeShiftPresetV3Body) (*DeleteTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateTimeShiftPresetV3(ctx context.Context, arg *UpdateTimeShiftPresetV3Body) (*UpdateTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListTimeShiftPresetV2(ctx context.Context, arg *ListTimeShiftPresetV2Body) (*ListTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListTimeShiftPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListTimeShiftPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTimeShiftPresetV3(ctx context.Context, arg *CreateTimeShiftPresetV3Body) (*CreateTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCallback(ctx context.Context, arg *DeleteCallbackBody) (*DeleteCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCallback(ctx context.Context, arg *DescribeCallbackBody) (*DescribeCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCallback(ctx context.Context, arg *UpdateCallbackBody) (*UpdateCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCert(ctx context.Context, arg *DeleteCertBody) (*DeleteCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDetailSecretV2(ctx context.Context, arg *DescribeCertDetailSecretV2Body) (*DescribeCertDetailSecretV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCertDetailSecretV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDetailSecretV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCertV2(ctx context.Context, arg *ListCertV2Body) (*ListCertV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCertV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCertV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateCert(ctx context.Context, arg *CreateCertBody) (*CreateCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) BindCert(ctx context.Context, arg *BindCertBody) (*BindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UnbindCert(ctx context.Context, arg *UnbindCertBody) (*UnbindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteDomain(ctx context.Context, arg *DeleteDomainBody) (*DeleteDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) EnableDomain(ctx context.Context, arg *EnableDomainBody) (*EnableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "EnableDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(EnableDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateDomainV2(ctx context.Context, arg *CreateDomainV2Body) (*CreateDomainV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDomainV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDomainV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDomainVhost(ctx context.Context, arg *UpdateDomainVhostBody) (*UpdateDomainVhostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDomainVhost", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDomainVhostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeDomain(ctx context.Context, arg *DescribeDomainBody) (*DescribeDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListDomainDetail(ctx context.Context, arg *ListDomainDetailBody) (*ListDomainDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListDomainDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListDomainDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateDomain(ctx context.Context, arg *CreateDomainBody) (*CreateDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DisableDomain(ctx context.Context, arg *DisableDomainBody) (*DisableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DisableDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DisableDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateLiveVideoQualityAnalysisTask(ctx context.Context, arg *CreateLiveVideoQualityAnalysisTaskBody) (*CreateLiveVideoQualityAnalysisTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateLiveVideoQualityAnalysisTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateLiveVideoQualityAnalysisTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteLiveVideoQualityAnalysisTask(ctx context.Context, arg *DeleteLiveVideoQualityAnalysisTaskBody) (*DeleteLiveVideoQualityAnalysisTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteLiveVideoQualityAnalysisTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteLiveVideoQualityAnalysisTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetLiveVideoQualityAnalysisTaskDetail(ctx context.Context, arg *GetLiveVideoQualityAnalysisTaskDetailBody) (*GetLiveVideoQualityAnalysisTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetLiveVideoQualityAnalysisTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetLiveVideoQualityAnalysisTaskDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListLiveVideoQualityAnalysisTasks(ctx context.Context, arg *ListLiveVideoQualityAnalysisTasksBody) (*ListLiveVideoQualityAnalysisTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListLiveVideoQualityAnalysisTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListLiveVideoQualityAnalysisTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullToPushTask(ctx context.Context, arg *StopPullToPushTaskBody) (*StopPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopPullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullToPushTask(ctx context.Context, arg *CreatePullToPushTaskBody) (*CreatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullToPushGroup(ctx context.Context, arg *CreatePullToPushGroupBody) (*CreatePullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeletePullToPushTask(ctx context.Context, arg *DeletePullToPushTaskBody) (*DeletePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeletePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeletePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeletePullToPushGroup(ctx context.Context, arg *DeletePullToPushGroupBody) (*DeletePullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeletePullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeletePullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) RestartPullToPushTask(ctx context.Context, arg *RestartPullToPushTaskBody) (*RestartPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RestartPullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RestartPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdatePullToPushTask(ctx context.Context, arg *UpdatePullToPushTaskBody) (*UpdatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdatePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdatePullToPushGroup(ctx context.Context, arg *UpdatePullToPushGroupBody) (*UpdatePullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdatePullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullToPushGroup(ctx context.Context, arg *ListPullToPushGroupBody) (*ListPullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullToPushTask(ctx context.Context, arg *ListPullToPushTaskQuery) (*ListPullToPushTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPullToPushTask", query)
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullToPushTaskV2(ctx context.Context, arg *ListPullToPushTaskV2Body) (*ListPullToPushTaskV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPullToPushTaskV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushTaskV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceV4(ctx context.Context, arg *DeleteRelaySourceV4Body) (*DeleteRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceV3(ctx context.Context, arg *DeleteRelaySourceV3Body) (*DeleteRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySourceV4(ctx context.Context, arg *UpdateRelaySourceV4Body) (*UpdateRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListRelaySourceV4(ctx context.Context, arg *ListRelaySourceV4Body) (*ListRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRelaySourceV3(ctx context.Context, arg *DescribeRelaySourceV3Body) (*DescribeRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateRelaySourceV4(ctx context.Context, arg *CreateRelaySourceV4Body) (*CreateRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySourceV3(ctx context.Context, arg *UpdateRelaySourceV3Body) (*UpdateRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) KillStream(ctx context.Context, arg *KillStreamBody) (*KillStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "KillStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(KillStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeClosedStreamInfoByPage(ctx context.Context, arg *DescribeClosedStreamInfoByPageQuery) (*DescribeClosedStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeClosedStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeClosedStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamInfoByPage(ctx context.Context, arg *DescribeLiveStreamInfoByPageQuery) (*DescribeLiveStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeLiveStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamState(ctx context.Context, arg *DescribeLiveStreamStateQuery) (*DescribeLiveStreamStateRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeLiveStreamState", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamStateRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeForbiddenStreamInfoByPage(ctx context.Context, arg *DescribeForbiddenStreamInfoByPageQuery) (*DescribeForbiddenStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeForbiddenStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeForbiddenStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ForbidStream(ctx context.Context, arg *ForbidStreamBody) (*ForbidStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ForbidStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ForbidStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ResumeStream(ctx context.Context, arg *ResumeStreamBody) (*ResumeStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResumeStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResumeStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GeneratePlayURL(ctx context.Context, arg *GeneratePlayURLBody) (*GeneratePlayURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GeneratePlayURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GeneratePlayURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GeneratePushURL(ctx context.Context, arg *GeneratePushURLBody) (*GeneratePushURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GeneratePushURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GeneratePushURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteStreamQuotaConfig(ctx context.Context, arg *DeleteStreamQuotaConfigBody) (*DeleteStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeStreamQuotaConfig(ctx context.Context, arg *DescribeStreamQuotaConfigBody) (*DescribeStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateStreamQuotaConfig(ctx context.Context, arg *UpdateStreamQuotaConfigBody) (*UpdateStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSnapshotAuditPreset(ctx context.Context, arg *DeleteSnapshotAuditPresetBody) (*DeleteSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSnapshotAuditPreset(ctx context.Context, arg *UpdateSnapshotAuditPresetBody) (*UpdateSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSnapshotAuditPreset(ctx context.Context, arg *ListVhostSnapshotAuditPresetBody) (*ListVhostSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotAuditPreset(ctx context.Context, arg *CreateSnapshotAuditPresetBody) (*CreateSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeIPInfo(ctx context.Context, arg *DescribeIPInfoBody) (*DescribeIPInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeIpInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeIPInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTopPlayData(ctx context.Context, arg *DescribeLiveTopPlayDataBody) (*DescribeLiveTopPlayDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTopPlayData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTopPlayDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveRegionData(ctx context.Context) (*DescribeLiveRegionDataRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveRegionData", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveRegionDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceStreamMetrics(ctx context.Context, arg *DescribeLiveSourceStreamMetricsBody) (*DescribeLiveSourceStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePushStreamMetrics(ctx context.Context, arg *DescribeLivePushStreamMetricsBody) (*DescribeLivePushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePushStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchStreamSessionData(ctx context.Context, arg *DescribeLiveBatchStreamSessionDataBody) (*DescribeLiveBatchStreamSessionDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchStreamSessionData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchStreamSessionDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamSessionData(ctx context.Context, arg *DescribeLiveStreamSessionDataBody) (*DescribeLiveStreamSessionDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveStreamSessionData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamSessionDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePlayStatusCodeData(ctx context.Context, arg *DescribeLivePlayStatusCodeDataBody) (*DescribeLivePlayStatusCodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePlayStatusCodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePlayStatusCodeDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchSourceStreamMetrics(ctx context.Context, arg *DescribeLiveBatchSourceStreamMetricsBody) (*DescribeLiveBatchSourceStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchSourceStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchSourceStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchSourceStreamAvgMetrics(ctx context.Context, arg *DescribeLiveBatchSourceStreamAvgMetricsBody) (*DescribeLiveBatchSourceStreamAvgMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchSourceStreamAvgMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchSourceStreamAvgMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchPushStreamMetrics(ctx context.Context, arg *DescribeLiveBatchPushStreamMetricsBody) (*DescribeLiveBatchPushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchPushStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchPushStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchPushStreamAvgMetrics(ctx context.Context, arg *DescribeLiveBatchPushStreamAvgMetricsBody) (*DescribeLiveBatchPushStreamAvgMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchPushStreamAvgMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchPushStreamAvgMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamCountData(ctx context.Context, arg *DescribeLiveStreamCountDataBody) (*DescribeLiveStreamCountDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveStreamCountData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamCountDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePushStreamCountData(ctx context.Context, arg *DescribeLivePushStreamCountDataBody) (*DescribeLivePushStreamCountDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePushStreamCountData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamCountDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePushStreamInfoData(ctx context.Context, arg *DescribeLivePushStreamInfoDataBody) (*DescribeLivePushStreamInfoDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePushStreamInfoData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamInfoDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTranscodeInfoData(ctx context.Context, arg *DescribeLiveTranscodeInfoDataBody) (*DescribeLiveTranscodeInfoDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTranscodeInfoData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTranscodeInfoDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceBandwidthData(ctx context.Context, arg *DescribeLiveSourceBandwidthDataBody) (*DescribeLiveSourceBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceTrafficData(ctx context.Context, arg *DescribeLiveSourceTrafficDataBody) (*DescribeLiveSourceTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveMetricBandwidthData(ctx context.Context, arg *DescribeLiveMetricBandwidthDataBody) (*DescribeLiveMetricBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveMetricBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveMetricBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveMetricTrafficData(ctx context.Context, arg *DescribeLiveMetricTrafficDataBody) (*DescribeLiveMetricTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveMetricTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveMetricTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchStreamTrafficData(ctx context.Context, arg *DescribeLiveBatchStreamTrafficDataBody) (*DescribeLiveBatchStreamTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchStreamTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchStreamTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveEdgeStatData(ctx context.Context, arg *DescribeLiveEdgeStatDataBody) (*DescribeLiveEdgeStatDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveEdgeStatData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveEdgeStatDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveISPData(ctx context.Context) (*DescribeLiveISPDataRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveISPData", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveISPDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveP95PeakBandwidthData(ctx context.Context, arg *DescribeLiveP95PeakBandwidthDataBody) (*DescribeLiveP95PeakBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveP95PeakBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveP95PeakBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveAuditData(ctx context.Context, arg *DescribeLiveAuditDataBody) (*DescribeLiveAuditDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveAuditData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveAuditDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePullToPushBandwidthData(ctx context.Context, arg *DescribeLivePullToPushBandwidthDataBody) (*DescribeLivePullToPushBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePullToPushBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePullToPushBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePullToPushData(ctx context.Context, arg *DescribeLivePullToPushDataBody) (*DescribeLivePullToPushDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePullToPushData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePullToPushDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBandwidthData(ctx context.Context, arg *DescribeLiveBandwidthDataBody) (*DescribeLiveBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveRecordData(ctx context.Context, arg *DescribeLiveRecordDataBody) (*DescribeLiveRecordDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveRecordData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveRecordDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSnapshotData(ctx context.Context, arg *DescribeLiveSnapshotDataBody) (*DescribeLiveSnapshotDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSnapshotData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSnapshotDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTrafficData(ctx context.Context, arg *DescribeLiveTrafficDataBody) (*DescribeLiveTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTranscodeData(ctx context.Context, arg *DescribeLiveTranscodeDataBody) (*DescribeLiveTranscodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTranscodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTranscodeDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTimeShiftData(ctx context.Context, arg *DescribeLiveTimeShiftDataBody) (*DescribeLiveTimeShiftDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTimeShiftData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTimeShiftDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveLogData(ctx context.Context, arg *DescribeLiveLogDataBody) (*DescribeLiveLogDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveLogData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveLogDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteReferer(ctx context.Context, arg *DeleteRefererBody) (*DeleteRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeReferer(ctx context.Context, arg *DescribeRefererBody) (*DescribeRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeAuth(ctx context.Context, arg *DescribeAuthBody) (*DescribeAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeAuth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeAuthRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateReferer(ctx context.Context, arg *UpdateRefererBody) (*UpdateRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateAuthKey(ctx context.Context, arg *UpdateAuthKeyBody) (*UpdateAuthKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAuthKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAuthKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteHTTPHeaderConfig(ctx context.Context, arg *DeleteHTTPHeaderConfigBody) (*DeleteHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) EnableHTTPHeaderConfig(ctx context.Context, arg *EnableHTTPHeaderConfigBody) (*EnableHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "EnableHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(EnableHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeHTTPHeaderConfig(ctx context.Context, arg *DescribeHTTPHeaderConfigBody) (*DescribeHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateHTTPHeaderConfig(ctx context.Context, arg *UpdateHTTPHeaderConfigBody) (*UpdateHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateEncryptDRM(ctx context.Context, arg *UpdateEncryptDRMBody) (*UpdateEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateEncryptHLS(ctx context.Context, arg *UpdateEncryptHLSBody) (*UpdateEncryptHLSRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateEncryptHLS", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateEncryptHLSRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetHLSEncryptDataKey(ctx context.Context, arg *GetHLSEncryptDataKeyQuery) (*GetHLSEncryptDataKeyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetHLSEncryptDataKey", query)
	if err != nil {
		return nil, err
	}

	result := new(GetHLSEncryptDataKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeEncryptHLS(ctx context.Context) (*DescribeEncryptHLSRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeEncryptHLS", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeEncryptHLSRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLicenseDRM(ctx context.Context, arg *DescribeLicenseDRMQuery) (*DescribeLicenseDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLicenseDRM", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLicenseDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDRM(ctx context.Context, arg *DescribeCertDRMQuery) (*DescribeCertDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeCertDRM", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeEncryptDRM(ctx context.Context) (*DescribeEncryptDRMRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeEncryptDRM", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) BindEncryptDRM(ctx context.Context, arg *BindEncryptDRMBody) (*BindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UnBindEncryptDRM(ctx context.Context, arg *UnBindEncryptDRMBody) (*UnBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnBindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnBindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListBindEncryptDRM(ctx context.Context, arg *ListBindEncryptDRMBody) (*ListBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListBindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListBindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteIPAccessRule(ctx context.Context, arg *DeleteIPAccessRuleBody) (*DeleteIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateIPAccessRule(ctx context.Context, arg *UpdateIPAccessRuleBody) (*UpdateIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeIPAccessRule(ctx context.Context, arg *DescribeIPAccessRuleBody) (*DescribeIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateCloudMixTask(ctx context.Context, arg *CreateCloudMixTaskBody) (*CreateCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCloudMixTask(ctx context.Context, arg *UpdateCloudMixTaskBody) (*UpdateCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetCloudMixTaskDetail(ctx context.Context, arg *GetCloudMixTaskDetailBody) (*GetCloudMixTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCloudMixTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCloudMixTaskDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCloudMixTask(ctx context.Context, arg *ListCloudMixTaskBody) (*ListCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCloudMixTask(ctx context.Context, arg *DeleteCloudMixTaskBody) (*DeleteCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSubtitleTranscodePreset(ctx context.Context, arg *DeleteSubtitleTranscodePresetBody) (*DeleteSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSubtitleTranscodePreset(ctx context.Context, arg *UpdateSubtitleTranscodePresetBody) (*UpdateSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSubtitleTranscodePreset(ctx context.Context, arg *ListVhostSubtitleTranscodePresetBody) (*ListVhostSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSubtitleTranscodePreset(ctx context.Context, arg *CreateSubtitleTranscodePresetBody) (*CreateSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateLivePadPreset(ctx context.Context, arg *CreateLivePadPresetBody) (*CreateLivePadPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateLivePadPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateLivePadPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteLivePadPreset(ctx context.Context, arg *DeleteLivePadPresetBody) (*DeleteLivePadPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteLivePadPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteLivePadPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopLivePadStream(ctx context.Context, arg *StopLivePadStreamBody) (*StopLivePadStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopLivePadStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopLivePadStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateLivePadPreset(ctx context.Context, arg *UpdateLivePadPresetBody) (*UpdateLivePadPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateLivePadPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateLivePadPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePadStreamList(ctx context.Context, arg *DescribeLivePadStreamListBody) (*DescribeLivePadStreamListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePadStreamList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePadStreamListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePadPresetDetail(ctx context.Context, arg *DescribeLivePadPresetDetailBody) (*DescribeLivePadPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePadPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePadPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateCarouselTask(ctx context.Context, arg *CreateCarouselTaskBody) (*CreateCarouselTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCarouselTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCarouselTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCarouselTask(ctx context.Context, arg *DeleteCarouselTaskBody) (*DeleteCarouselTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCarouselTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCarouselTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCarouselTask(ctx context.Context, arg *UpdateCarouselTaskBody) (*UpdateCarouselTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCarouselTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCarouselTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetCarouselDetail(ctx context.Context, arg *GetCarouselDetailBody) (*GetCarouselDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCarouselDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCarouselDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCarouselTask(ctx context.Context, arg *ListCarouselTaskBody) (*ListCarouselTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCarouselTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCarouselTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateHighLightTask(ctx context.Context, arg *CreateHighLightTaskBody) (*CreateHighLightTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateHighLightTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateHighLightTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTaskByAccountID(ctx context.Context, arg *DeleteTaskByAccountIDBody) (*DeleteTaskByAccountIDRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTaskByAccountID", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTaskByAccountIDRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
