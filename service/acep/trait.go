package acep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

// func skipEmptyValue() queryMarshalFilter {
// 	return func(_ string, value []string) (accept bool) {
// 		if len(value) == 0 {
// 			return false
// 		}
//
// 		for _, item := range value {
// 			if item == "" {
// 				return false
// 			}
// 		}
//
// 		return true
// 	}
// }

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

func (c *ACEP) ListContainerImages(ctx context.Context, arg *ListContainerImagesQuery) (*ListContainerImagesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListContainerImages", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListContainerImagesRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListContainerImagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DeleteContainerImages(ctx context.Context, arg *DeleteContainerImagesBody) (*DeleteContainerImagesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteContainerImages", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DeleteContainerImagesRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DeleteContainerImagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ImportContainerImage(ctx context.Context, arg *ImportContainerImageBody) (*ImportContainerImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ImportContainerImage", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(ImportContainerImageRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ImportContainerImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PullFile(ctx context.Context, arg *PullFileBody) (*PullFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PullFile", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PullFileRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PullFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdatePodProperty(ctx context.Context, arg *UpdatePodPropertyBody) (*UpdatePodPropertyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePodProperty", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdatePodPropertyRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdatePodPropertyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CloseApp(ctx context.Context, arg *CloseAppBody) (*CloseAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CloseApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CloseAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CloseAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreatePod(ctx context.Context, arg *CreatePodBody) (*CreatePodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreatePodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreatePodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreatePodOneStep(ctx context.Context, arg *CreatePodOneStepBody) (*CreatePodOneStepRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePodOneStep", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreatePodOneStepRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreatePodOneStepRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) AddPropertyRule(ctx context.Context, arg *AddPropertyRuleBody) (*AddPropertyRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddPropertyRule", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(AddPropertyRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(AddPropertyRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DeletePod(ctx context.Context, arg *DeletePodBody) (*DeletePodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeletePod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DeletePodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DeletePodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RunSyncCommand(ctx context.Context, arg *RunSyncCommandBody) (*RunSyncCommandRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RunSyncCommand", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RunSyncCommandRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RunSyncCommandRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) LaunchApp(ctx context.Context, arg *LaunchAppBody) (*LaunchAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "LaunchApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(LaunchAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(LaunchAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) LaunchApps(ctx context.Context, arg *LaunchAppsBody) (*LaunchAppsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "LaunchApps", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(LaunchAppsRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(LaunchAppsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PowerOffPod(ctx context.Context, arg *PowerOffPodBody) (*PowerOffPodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PowerOffPod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PowerOffPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PowerOffPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PodStop(ctx context.Context, arg *PodStopBody) (*PodStopRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PodStop", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PodStopRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PodStopRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PowerOnPod(ctx context.Context, arg *PowerOnPodBody) (*PowerOnPodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PowerOnPod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PowerOnPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PowerOnPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CopyPod(ctx context.Context, arg *CopyPodBody) (*CopyPodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CopyPod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CopyPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CopyPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PodDataTransfer(ctx context.Context, arg *PodDataTransferBody) (*PodDataTransferRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PodDataTransfer", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PodDataTransferRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PodDataTransferRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RebootPod(ctx context.Context, arg *RebootPodBody) (*RebootPodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RebootPod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RebootPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RebootPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ResetPod(ctx context.Context, arg *ResetPodBody) (*ResetPodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResetPod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(ResetPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ResetPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) BanUser(ctx context.Context, arg *BanUserBody) (*BanUserRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BanUser", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(BanUserRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(BanUserRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PushFile(ctx context.Context, arg *PushFileBody) (*PushFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PushFile", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PushFileRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PushFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DistributeFile(ctx context.Context, arg *DistributeFileBody) (*DistributeFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DistributeFile", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DistributeFileRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DistributeFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) StartRecording(ctx context.Context, arg *StartRecordingBody) (*StartRecordingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StartRecording", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(StartRecordingRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(StartRecordingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ScreenShot(ctx context.Context, arg *ScreenShotBody) (*ScreenShotRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ScreenShot", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(ScreenShotRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ScreenShotRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PodAdb(ctx context.Context, arg *PodAdbBody) (*PodAdbRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PodAdb", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PodAdbRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PodAdbRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RunCommand(ctx context.Context, arg *RunCommandBody) (*RunCommandRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RunCommand", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RunCommandRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RunCommandRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) BatchScreenShot(ctx context.Context, arg *BatchScreenShotBody) (*BatchScreenShotRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchScreenShot", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(BatchScreenShotRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(BatchScreenShotRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PodMute(ctx context.Context, arg *PodMuteBody) (*PodMuteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PodMute", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PodMuteRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PodMuteRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdatePod(ctx context.Context, arg *UpdatePodBody) (*UpdatePodRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePod", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdatePodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdatePodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListDc(ctx context.Context, arg *ListDcQuery) (*ListDcRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListDc", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListDcRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListDcRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetPodMetric(ctx context.Context, arg *GetPodMetricQuery) (*GetPodMetricRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetPodMetric", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetPodMetricRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetPodMetricRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailPod(ctx context.Context, arg *DetailPodQuery) (*DetailPodRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailPod", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListPod(ctx context.Context, arg *ListPodQuery) (*ListPodRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPod", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListPodRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListPodRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetPodProperty(ctx context.Context, arg *GetPodPropertyQuery) (*GetPodPropertyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetPodProperty", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetPodPropertyRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetPodPropertyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListPropertyRule(ctx context.Context, arg *ListPropertyRuleQuery) (*ListPropertyRuleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPropertyRule", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListPropertyRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListPropertyRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) PodDataDelete(ctx context.Context, arg *PodDataDeleteBody) (*PodDataDeleteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PodDataDelete", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(PodDataDeleteRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(PodDataDeleteRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RemovePropertyRule(ctx context.Context, arg *RemovePropertyRuleBody) (*RemovePropertyRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RemovePropertyRule", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RemovePropertyRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RemovePropertyRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) StopRecording(ctx context.Context, arg *StopRecordingBody) (*StopRecordingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopRecording", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(StopRecordingRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(StopRecordingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetPodAppList(ctx context.Context, arg *GetPodAppListQuery) (*GetPodAppListRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetPodAppList", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetPodAppListRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetPodAppListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) SetProxy(ctx context.Context, arg *SetProxyBody) (*SetProxyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetProxy", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(SetProxyRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(SetProxyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListTask(ctx context.Context, arg *ListTaskQuery) (*ListTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListTask", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListTaskRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetTaskInfo(ctx context.Context, arg *GetTaskInfoQuery) (*GetTaskInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetTaskInfo", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetTaskInfoRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetTaskInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreatePortMappingRule(ctx context.Context, arg *CreatePortMappingRuleBody) (*CreatePortMappingRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePortMappingRule", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreatePortMappingRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreatePortMappingRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListPortMappingRule(ctx context.Context, arg *ListPortMappingRuleQuery) (*ListPortMappingRuleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPortMappingRule", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListPortMappingRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListPortMappingRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailPortMappingRule(ctx context.Context, arg *DetailPortMappingRuleQuery) (*DetailPortMappingRuleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailPortMappingRule", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailPortMappingRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailPortMappingRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) BindPortMappingRule(ctx context.Context, arg *BindPortMappingRuleBody) (*BindPortMappingRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindPortMappingRule", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(BindPortMappingRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(BindPortMappingRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UnbindPortMappingRule(ctx context.Context, arg *UnbindPortMappingRuleBody) (*UnbindPortMappingRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindPortMappingRule", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UnbindPortMappingRuleRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UnbindPortMappingRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) AttachTag(ctx context.Context, arg *AttachTagBody) (*AttachTagRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AttachTag", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(AttachTagRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(AttachTagRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListTag(ctx context.Context, arg *ListTagQuery) (*ListTagRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListTag", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListTagRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListTagRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreateTag(ctx context.Context, arg *CreateTagBody) (*CreateTagRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTag", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreateTagRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreateTagRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DeleteTag(ctx context.Context, arg *DeleteTagBody) (*DeleteTagRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTag", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DeleteTagRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DeleteTagRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdateTag(ctx context.Context, arg *UpdateTagBody) (*UpdateTagRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTag", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdateTagRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdateTagRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UploadApp(ctx context.Context, arg *UploadAppBody) (*UploadAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UploadApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UploadAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UploadAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailApp(ctx context.Context, arg *DetailAppQuery) (*DetailAppRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailApp", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdateApp(ctx context.Context, arg *UpdateAppBody) (*UpdateAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdateAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdateAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListApp(ctx context.Context, arg *ListAppQuery) (*ListAppRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListApp", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DeleteApp(ctx context.Context, arg *DeleteAppBody) (*DeleteAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DeleteAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DeleteAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UninstallApp(ctx context.Context, arg *UninstallAppBody) (*UninstallAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UninstallApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UninstallAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UninstallAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) InstallApp(ctx context.Context, arg *InstallAppBody) (*InstallAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "InstallApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(InstallAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(InstallAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) InstallApps(ctx context.Context, arg *InstallAppsBody) (*InstallAppsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "InstallApps", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(InstallAppsRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(InstallAppsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListAppVersionDeploy(ctx context.Context, arg *ListAppVersionDeployQuery) (*ListAppVersionDeployRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListAppVersionDeploy", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListAppVersionDeployRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListAppVersionDeployRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) AutoInstallApp(ctx context.Context, arg *AutoInstallAppBody) (*AutoInstallAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AutoInstallApp", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(AutoInstallAppRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(AutoInstallAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetAppCrashLog(ctx context.Context, arg *GetAppCrashLogQuery) (*GetAppCrashLogRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAppCrashLog", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetAppCrashLogRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetAppCrashLogRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreateDisplayLayoutMini(ctx context.Context, arg *CreateDisplayLayoutMiniBody) (*CreateDisplayLayoutMiniRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDisplayLayoutMini", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreateDisplayLayoutMiniRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreateDisplayLayoutMiniRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DeleteDisplayLayout(ctx context.Context, arg *DeleteDisplayLayoutBody) (*DeleteDisplayLayoutRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDisplayLayout", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(DeleteDisplayLayoutRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DeleteDisplayLayoutRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListDisplayLayoutMini(ctx context.Context, arg *ListDisplayLayoutMiniQuery) (*ListDisplayLayoutMiniRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListDisplayLayoutMini", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListDisplayLayoutMiniRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListDisplayLayoutMiniRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailDisplayLayoutMini(ctx context.Context, arg *DetailDisplayLayoutMiniQuery) (*DetailDisplayLayoutMiniRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailDisplayLayoutMini", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailDisplayLayoutMiniRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailDisplayLayoutMiniRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreateAppImage(ctx context.Context, arg *CreateAppImageBody) (*CreateAppImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateAppImage", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreateAppImageRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreateAppImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailAppVersionImage(ctx context.Context, arg *DetailAppVersionImageQuery) (*DetailAppVersionImageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailAppVersionImage", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailAppVersionImageRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailAppVersionImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) CreateImageOneStep(ctx context.Context, arg *CreateImageOneStepBody) (*CreateImageOneStepRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageOneStep", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(CreateImageOneStepRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(CreateImageOneStepRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListImageResource(ctx context.Context, arg *ListImageResourceQuery) (*ListImageResourceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListImageResource", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListImageResourceRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListImageResourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetImagePreheating(ctx context.Context, arg *GetImagePreheatingQuery) (*GetImagePreheatingRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImagePreheating", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetImagePreheatingRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetImagePreheatingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListConfiguration(ctx context.Context, arg *ListConfigurationQuery) (*ListConfigurationRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListConfiguration", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListConfigurationRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListConfigurationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListPodResourceSet(ctx context.Context, arg *ListPodResourceSetQuery) (*ListPodResourceSetRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPodResourceSet", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListPodResourceSetRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListPodResourceSetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListPodResource(ctx context.Context, arg *ListPodResourceQuery) (*ListPodResourceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPodResource", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListPodResourceRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListPodResourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListResourceQuota(ctx context.Context, arg *ListResourceQuotaQuery) (*ListResourceQuotaRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListResourceQuota", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListResourceQuotaRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListResourceQuotaRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdatePodResourceApplyNum(ctx context.Context, arg *UpdatePodResourceApplyNumBody) (*UpdatePodResourceApplyNumRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePodResourceApplyNum", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdatePodResourceApplyNumRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdatePodResourceApplyNumRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) SubscribeResourceAuto(ctx context.Context, arg *SubscribeResourceAutoBody) (*SubscribeResourceAutoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SubscribeResourceAuto", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(SubscribeResourceAutoRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(SubscribeResourceAutoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RenewResourceAuto(ctx context.Context, arg *RenewResourceAutoBody) (*RenewResourceAutoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RenewResourceAuto", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RenewResourceAutoRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RenewResourceAutoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdateProductResource(ctx context.Context, arg *UpdateProductResourceBody) (*UpdateProductResourceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateProductResource", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdateProductResourceRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdateProductResourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) GetProductResource(ctx context.Context, arg *GetProductResourceQuery) (*GetProductResourceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetProductResource", query)
	if err != nil {
		if len(data) != 0 {
			result := new(GetProductResourceRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(GetProductResourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UnsubscribeHostResource(ctx context.Context, arg *UnsubscribeHostResourceBody) (*UnsubscribeHostResourceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnsubscribeHostResource", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UnsubscribeHostResourceRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UnsubscribeHostResourceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) UpdateHost(ctx context.Context, arg *UpdateHostBody) (*UpdateHostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHost", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(UpdateHostRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(UpdateHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ListHost(ctx context.Context, arg *ListHostQuery) (*ListHostRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListHost", query)
	if err != nil {
		if len(data) != 0 {
			result := new(ListHostRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ListHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) DetailHost(ctx context.Context, arg *DetailHostQuery) (*DetailHostRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailHost", query)
	if err != nil {
		if len(data) != 0 {
			result := new(DetailHostRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(DetailHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) RebootHost(ctx context.Context, arg *RebootHostBody) (*RebootHostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RebootHost", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(RebootHostRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(RebootHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ACEP) ResetHost(ctx context.Context, arg *ResetHostBody) (*ResetHostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResetHost", url.Values{}, string(body))
	if err != nil {
		if len(data) != 0 {
			result := new(ResetHostRes)
			err = unmarshalResultInto(data, result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, err
	}

	result := new(ResetHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
