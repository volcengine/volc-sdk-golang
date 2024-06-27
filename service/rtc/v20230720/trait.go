package rtc_v20230720

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

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

func (c *Rtc) SendBroadcast(ctx context.Context, arg *SendBroadcastBody) (*SendBroadcastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "SendBroadcast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(SendBroadcastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) SendRoomUnicast(ctx context.Context, arg *SendRoomUnicastBody) (*SendRoomUnicastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "SendRoomUnicast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(SendRoomUnicastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) SendUnicast(ctx context.Context, arg *SendUnicastBody) (*SendUnicastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "SendUnicast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(SendUnicastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}

func (c *Rtc) BatchSendRoomUnicast(ctx context.Context, arg *BatchSendRoomUnicastBody) (*BatchSendRoomUnicastRes, int, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, 0, err
	}

	data, statusCode, err := c.Client.CtxJson(ctx, "BatchSendRoomUnicast", url.Values{}, string(body))

	if len(data) > 0 {
		result := new(BatchSendRoomUnicastRes)
		if err2 := json.Unmarshal(data, result); err2 != nil {
			return nil, statusCode, err2
		}
		return result, statusCode, err
	}

	return nil, statusCode, err
}
