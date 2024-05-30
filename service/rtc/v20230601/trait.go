package rtc_v20230601

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

func (c *Rtc) StartRecord(ctx context.Context, arg *StartRecordBody) (*StartRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartRecord", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(StartRecordRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) StartPushSingleStreamToCDN(ctx context.Context, arg *StartPushSingleStreamToCDNBody) (*StartPushSingleStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartPushSingleStreamToCDN", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(StartPushSingleStreamToCDNRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) StartPushMixedStreamToCDN(ctx context.Context, arg *StartPushMixedStreamToCDNBody) (*StartPushMixedStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StartPushMixedStreamToCDN", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(StartPushMixedStreamToCDNRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) UpdateRecord(ctx context.Context, arg *UpdateRecordBody) (*UpdateRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdateRecord", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(UpdateRecordRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) UpdatePushMixedStreamToCDN(ctx context.Context, arg *UpdatePushMixedStreamToCDNBody) (*UpdatePushMixedStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "UpdatePushMixedStreamToCDN", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(UpdatePushMixedStreamToCDNRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) GetPushSingleStreamToCDNTask(ctx context.Context, arg *GetPushSingleStreamToCDNTaskQuery) (*GetPushSingleStreamToCDNTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetPushSingleStreamToCDNTask", query)
	if err != nil {
		return nil, statusCode, err
	}

	result := new(GetPushSingleStreamToCDNTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) GetPushMixedStreamToCDNTask(ctx context.Context, arg *GetPushMixedStreamToCDNTaskQuery) (*GetPushMixedStreamToCDNTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetPushMixedStreamToCDNTask", query)
	if err != nil {
		return nil, statusCode, err
	}

	result := new(GetPushMixedStreamToCDNTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) GetRecordTask(ctx context.Context, arg *GetRecordTaskQuery) (*GetRecordTaskRes, int, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.CtxQuery(ctx, "GetRecordTask", query)
	if err != nil {
		return nil, statusCode, err
	}

	result := new(GetRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) StopRecord(ctx context.Context, arg *StopRecordBody) (*StopRecordRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopRecord", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(StopRecordRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}

func (c *Rtc) StopPushStreamToCDN(ctx context.Context, arg *StopPushStreamToCDNBody) (*StopPushStreamToCDNRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "StopPushStreamToCDN", url.Values{}, string(body))
	if err != nil {
		return nil, statusCode, err
	}

	result := new(StopPushStreamToCDNRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, statusCode, err
	}

	return result, statusCode, nil
}
