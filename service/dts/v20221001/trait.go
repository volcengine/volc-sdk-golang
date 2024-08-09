package dts_v20221001

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

func (c *Dts) CreateDataValidationTask(ctx context.Context, arg *CreateDataValidationTaskBody) (*CreateDataValidationTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDataValidationTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDataValidationTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) CreateSubscriptionGroup(ctx context.Context, arg *CreateSubscriptionGroupBody) (*CreateSubscriptionGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSubscriptionGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSubscriptionGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) CreateTransmissionTask(ctx context.Context, arg *CreateTransmissionTaskBody) (*CreateTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DeleteDataValidationTask(ctx context.Context, arg *DeleteDataValidationTaskBody) (*DeleteDataValidationTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDataValidationTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDataValidationTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DeleteSubscriptionGroup(ctx context.Context, arg *DeleteSubscriptionGroupBody) (*DeleteSubscriptionGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSubscriptionGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSubscriptionGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DeleteTransmissionTask(ctx context.Context, arg *DeleteTransmissionTaskBody) (*DeleteTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DeleteTransmissionTasks(ctx context.Context, arg *DeleteTransmissionTasksBody) (*DeleteTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeDataValidationResult(ctx context.Context, arg *DescribeDataValidationResultBody) (*DescribeDataValidationResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDataValidationResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDataValidationResultRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeDataValidationTasks(ctx context.Context, arg *DescribeDataValidationTasksBody) (*DescribeDataValidationTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDataValidationTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDataValidationTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeSubscriptionGroup(ctx context.Context, arg *DescribeSubscriptionGroupBody) (*DescribeSubscriptionGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSubscriptionGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSubscriptionGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeSubscriptionGroupProgress(ctx context.Context, arg *DescribeSubscriptionGroupProgressBody) (*DescribeSubscriptionGroupProgressRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSubscriptionGroupProgress", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSubscriptionGroupProgressRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeSubscriptionGroups(ctx context.Context, arg *DescribeSubscriptionGroupsBody) (*DescribeSubscriptionGroupsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSubscriptionGroups", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSubscriptionGroupsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeTransmissionTaskInfo(ctx context.Context, arg *DescribeTransmissionTaskInfoBody) (*DescribeTransmissionTaskInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeTransmissionTaskInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeTransmissionTaskInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeTransmissionTaskProgress(ctx context.Context, arg *DescribeTransmissionTaskProgressBody) (*DescribeTransmissionTaskProgressRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeTransmissionTaskProgress", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeTransmissionTaskProgressRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) DescribeTransmissionTasks(ctx context.Context, arg *DescribeTransmissionTasksBody) (*DescribeTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) ModifyTransmissionTask(ctx context.Context, arg *ModifyTransmissionTaskBody) (*ModifyTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) ResumeDataValidationTask(ctx context.Context, arg *ResumeDataValidationTaskBody) (*ResumeDataValidationTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResumeDataValidationTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResumeDataValidationTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) ResumeTransmissionTask(ctx context.Context, arg *ResumeTransmissionTaskBody) (*ResumeTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResumeTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResumeTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) ResumeTransmissionTasks(ctx context.Context, arg *ResumeTransmissionTasksBody) (*ResumeTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResumeTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResumeTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) RetryTransmissionTask(ctx context.Context, arg *RetryTransmissionTaskBody) (*RetryTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RetryTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RetryTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) RetryTransmissionTasks(ctx context.Context, arg *RetryTransmissionTasksBody) (*RetryTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RetryTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RetryTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) SetBiSyncDDLSource(ctx context.Context, arg *SetBiSyncDDLSourceBody) (*SetBiSyncDDLSourceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetBiSyncDDLSource", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetBiSyncDDLSourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) StartTransmissionTask(ctx context.Context, arg *StartTransmissionTaskBody) (*StartTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StartTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StartTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) StartTransmissionTasks(ctx context.Context, arg *StartTransmissionTasksBody) (*StartTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StartTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StartTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) StopDataValidationTask(ctx context.Context, arg *StopDataValidationTaskBody) (*StopDataValidationTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopDataValidationTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopDataValidationTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) StopTransmissionTask(ctx context.Context, arg *StopTransmissionTaskBody) (*StopTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) StopTransmissionTasks(ctx context.Context, arg *StopTransmissionTasksBody) (*StopTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) SuspendDataValidationTask(ctx context.Context, arg *SuspendDataValidationTaskBody) (*SuspendDataValidationTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SuspendDataValidationTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SuspendDataValidationTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) SuspendTransmissionTask(ctx context.Context, arg *SuspendTransmissionTaskBody) (*SuspendTransmissionTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SuspendTransmissionTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SuspendTransmissionTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) SuspendTransmissionTasks(ctx context.Context, arg *SuspendTransmissionTasksBody) (*SuspendTransmissionTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SuspendTransmissionTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SuspendTransmissionTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Dts) UpdateSubscriptionGroup(ctx context.Context, arg *UpdateSubscriptionGroupBody) (*UpdateSubscriptionGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSubscriptionGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSubscriptionGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
