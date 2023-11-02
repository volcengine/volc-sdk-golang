package rtc

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
	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

func (c *Rtc) StartWBRecord(ctx context.Context, arg *StartWBRecordBody) (*common.CommonResponse, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StartWBRecord", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(common.CommonResponse)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Rtc) WbTranscodeCreate(ctx context.Context, arg *WbTranscodeCreateBody) (*WbTranscodeCreateRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "WbTranscodeCreate", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(WbTranscodeCreateRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Rtc) WbTranscodeQuery(ctx context.Context, arg *WbTranscodeQueryQuery) (*WbTranscodeQueryRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "WbTranscodeQuery", query)
	if err != nil {
		return nil, err
	}

	result := new(WbTranscodeQueryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Rtc) StopWBRecord(ctx context.Context, arg *StopWBRecordBody) (*StopWBRecordRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopWBRecord", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopWBRecordRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Rtc) WbTranscodeGet(ctx context.Context, arg *WbTranscodeGetQuery) (*WbTranscodeGetRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "WbTranscodeGet", query)
	if err != nil {
		return nil, err
	}

	result := new(WbTranscodeGetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
