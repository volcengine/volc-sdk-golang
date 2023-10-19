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

func (c *Live) UpdateSnapshotPreset(ctx context.Context, arg *UpdateSnapshotPresetBody) (*UpdateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotPresetRes)
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

func (c *Live) ListVhostSnapshotPreset(ctx context.Context, arg *ListVhostSnapshotPresetBody) (*ListVhostSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotPreset(ctx context.Context, arg *CreateSnapshotPresetBody) (*CreateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotPresetRes)
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

func (c *Live) CreateVerifyContent(ctx context.Context, arg *CreateVerifyContentBody) (*CreateVerifyContentRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateVerifyContent", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateVerifyContentRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) VerifyDomainOwner(ctx context.Context, arg *VerifyDomainOwnerBody) (*VerifyDomainOwnerRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "VerifyDomainOwner", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(VerifyDomainOwnerRes)
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

func (c *Live) DescribeDenyConfig(ctx context.Context, arg *DescribeDenyConfigBody) (*DescribeDenyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDenyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDenyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDenyConfig(ctx context.Context, arg *UpdateDenyConfigBody) (*UpdateDenyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDenyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDenyConfigRes)
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

func (c *Live) ListVqosDimensionValues(ctx context.Context, arg *ListVqosDimensionValuesReq) (*ListVqosDimensionValuesRes, error) {
	query, err := marshalToQuery(arg.ListVqosDimensionValuesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.ListVqosDimensionValuesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVqosDimensionValues", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVqosDimensionValuesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVqosMetricsDimensions(ctx context.Context, arg *ListVqosMetricsDimensionsQuery) (*ListVqosMetricsDimensionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListVqosMetricsDimensions", query)
	if err != nil {
		return nil, err
	}

	result := new(ListVqosMetricsDimensionsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetVqosRawData(ctx context.Context, arg *GetVqosRawDataReq) (*GetVqosRawDataRes, error) {
	query, err := marshalToQuery(arg.GetVqosRawDataQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetVqosRawDataBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetVqosRawData", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetVqosRawDataRes)
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

func (c *Live) DescribeLiveBatchStreamTranscodeData(ctx context.Context, arg *DescribeLiveBatchStreamTranscodeDataBody) (*DescribeLiveBatchStreamTranscodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchStreamTranscodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchStreamTranscodeDataRes)
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
