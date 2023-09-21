package ipaas

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

func (c *IPaaS) ListInstance(ctx context.Context, arg *ListInstanceQuery) (*ListInstanceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListInstance", query)
	if err != nil {
		return nil, err
	}

	result := new(ListInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ExportInstance(ctx context.Context, arg *ExportInstanceQuery) (*ExportInstanceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ExportInstance", query)
	if err != nil {
		return nil, err
	}

	result := new(ExportInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListInstanceMetricData(ctx context.Context, arg *ListInstanceMetricDataBody) (*ListInstanceMetricDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListInstanceMetricData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListInstanceMetricDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListPortMapping(ctx context.Context, arg *ListPortMappingQuery) (*ListPortMappingRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPortMapping", query)
	if err != nil {
		return nil, err
	}

	result := new(ListPortMappingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DetailInstance(ctx context.Context, arg *DetailInstanceQuery) (*DetailInstanceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailInstance", query)
	if err != nil {
		return nil, err
	}

	result := new(DetailInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) LatestMetricInstance(ctx context.Context, arg *LatestMetricInstanceQuery) (*LatestMetricInstanceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "LatestMetricInstance", query)
	if err != nil {
		return nil, err
	}

	result := new(LatestMetricInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ModifyInstanceWindowDisplaySpec(ctx context.Context, arg *ModifyInstanceWindowDisplaySpecBody) (*ModifyInstanceWindowDisplaySpecRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyInstanceWindowDisplaySpec", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyInstanceWindowDisplaySpecRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) PowerDownInstance(ctx context.Context, arg *PowerDownInstanceBody) (*PowerDownInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PowerDownInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(PowerDownInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ColdRebootInstance(ctx context.Context, arg *ColdRebootInstanceBody) (*ColdRebootInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ColdRebootInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ColdRebootInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpgradeInstances(ctx context.Context, arg *UpgradeInstancesBody) (*UpgradeInstancesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpgradeInstances", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpgradeInstancesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) PowerUpInstance(ctx context.Context, arg *PowerUpInstanceBody) (*PowerUpInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PowerUpInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(PowerUpInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) AdbCommand(ctx context.Context, arg *AdbCommandBody) (*AdbCommandRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AdbCommand", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AdbCommandRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) InstallApplication(ctx context.Context, arg *InstallApplicationBody) (*InstallApplicationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "InstallApplication", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(InstallApplicationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ControlApplication(ctx context.Context, arg *ControlApplicationBody) (*ControlApplicationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ControlApplication", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ControlApplicationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetInstanceProperty(ctx context.Context, arg *GetInstancePropertyBody) (*GetInstancePropertyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetInstanceProperty", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetInstancePropertyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) PullFile(ctx context.Context, arg *PullFileBody) (*PullFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PullFile", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(PullFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) PushFile(ctx context.Context, arg *PushFileBody) (*PushFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "PushFile", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(PushFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) RecordScreen(ctx context.Context, arg *RecordScreenBody) (*RecordScreenRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RecordScreen", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RecordScreenRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ExecCmdSync(ctx context.Context, arg *ExecCmdSyncBody) (*ExecCmdSyncRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ExecCmdSync", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ExecCmdSyncRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ResetInstanceToFactory(ctx context.Context, arg *ResetInstanceToFactoryBody) (*ResetInstanceToFactoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResetInstanceToFactory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResetInstanceToFactoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpdateInstanceProperty(ctx context.Context, arg *UpdateInstancePropertyBody) (*UpdateInstancePropertyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateInstanceProperty", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateInstancePropertyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) WarmRebootInstance(ctx context.Context, arg *WarmRebootInstanceBody) (*WarmRebootInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "WarmRebootInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(WarmRebootInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ResetInstances(ctx context.Context, arg *ResetInstancesBody) (*ResetInstancesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResetInstances", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResetInstancesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetInstanceProperties(ctx context.Context, arg *GetInstancePropertiesBody) (*GetInstancePropertiesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetInstanceProperties", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetInstancePropertiesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) SetInstanceProperties(ctx context.Context, arg *SetInstancePropertiesBody) (*SetInstancePropertiesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetInstanceProperties", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetInstancePropertiesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ImportContainerImage(ctx context.Context, arg *ImportContainerImageBody) (*ImportContainerImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ImportContainerImage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ImportContainerImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListContainerImages(ctx context.Context, arg *ListContainerImagesQuery) (*ListContainerImagesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListContainerImages", query)
	if err != nil {
		return nil, err
	}

	result := new(ListContainerImagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DeleteContainerImages(ctx context.Context, arg *DeleteContainerImagesBody) (*DeleteContainerImagesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteContainerImages", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteContainerImagesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpdateContainerImage(ctx context.Context, arg *UpdateContainerImageBody) (*UpdateContainerImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateContainerImage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateContainerImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetJobDetails(ctx context.Context, arg *GetJobDetailsQuery) (*GetJobDetailsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetJobDetails", query)
	if err != nil {
		return nil, err
	}

	result := new(GetJobDetailsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListAdbKey(ctx context.Context, arg *ListAdbKeyQuery) (*ListAdbKeyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListAdbKey", query)
	if err != nil {
		return nil, err
	}

	result := new(ListAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) BindInstancesAdbKey(ctx context.Context, arg *BindInstancesAdbKeyBody) (*BindInstancesAdbKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindInstancesAdbKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindInstancesAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UnbindInstancesAdbKey(ctx context.Context, arg *UnbindInstancesAdbKeyBody) (*UnbindInstancesAdbKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindInstancesAdbKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindInstancesAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) BindInstanceAdbKey(ctx context.Context, arg *BindInstanceAdbKeyBody) (*BindInstanceAdbKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindInstanceAdbKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindInstanceAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UnbindInstanceAdbKey(ctx context.Context, arg *UnbindInstanceAdbKeyBody) (*UnbindInstanceAdbKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindInstanceAdbKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindInstanceAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DistributeFile(ctx context.Context, arg *DistributeFileBody) (*DistributeFileRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DistributeFile", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DistributeFileRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DistributeFileToInstances(ctx context.Context, arg *DistributeFileToInstancesBody) (*DistributeFileToInstancesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DistributeFileToInstances", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DistributeFileToInstancesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetFileDistributionJobDetail(ctx context.Context, arg *GetFileDistributionJobDetailQuery) (*GetFileDistributionJobDetailRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetFileDistributionJobDetail", query, "")
	if err != nil {
		return nil, err
	}

	result := new(GetFileDistributionJobDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetFileDistributionResult(ctx context.Context, arg *GetFileDistributionResultQuery) (*GetFileDistributionResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetFileDistributionResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetFileDistributionResultRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListHost(ctx context.Context, arg *ListHostQuery) (*ListHostRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListHost", query)
	if err != nil {
		return nil, err
	}

	result := new(ListHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetInfoAfterOrder(ctx context.Context, arg *GetInfoAfterOrderBody) (*GetInfoAfterOrderRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetInfoAfterOrder", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetInfoAfterOrderRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListHostMetricData(ctx context.Context, arg *ListHostMetricDataBody) (*ListHostMetricDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListHostMetricData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListHostMetricDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) RebootHost(ctx context.Context, arg *RebootHostBody) (*RebootHostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RebootHost", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RebootHostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) FixInstancesSGBound(ctx context.Context, arg *FixInstancesSGBoundBody) (*FixInstancesSGBoundRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "FixInstancesSGBound", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(FixInstancesSGBoundRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListSecurityGroup(ctx context.Context, arg *ListSecurityGroupQuery) (*ListSecurityGroupRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListSecurityGroup", query)
	if err != nil {
		return nil, err
	}

	result := new(ListSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DetailSecurityGroup(ctx context.Context, arg *DetailSecurityGroupQuery) (*DetailSecurityGroupRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailSecurityGroup", query)
	if err != nil {
		return nil, err
	}

	result := new(DetailSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) BindInstancesSecurityGroup(ctx context.Context, arg *BindInstancesSecurityGroupBody) (*BindInstancesSecurityGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindInstancesSecurityGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindInstancesSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UnbindInstancesSecurityGroup(ctx context.Context, arg *UnbindInstancesSecurityGroupBody) (*UnbindInstancesSecurityGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindInstancesSecurityGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindInstancesSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) CreateDevices(ctx context.Context, arg *CreateDevicesReq) (*CreateDevicesRes, error) {
	query, err := marshalToQuery(arg.CreateDevicesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateDevicesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDevices", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDevicesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DeleteDevices(ctx context.Context, arg *DeleteDevicesReq) (*DeleteDevicesRes, error) {
	query, err := marshalToQuery(arg.DeleteDevicesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteDevicesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDevices", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDevicesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) AcquireIdempotentToken(ctx context.Context, arg *AcquireIdempotentTokenBody) (*AcquireIdempotentTokenRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AcquireIdempotentToken", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AcquireIdempotentTokenRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListProduct(ctx context.Context, arg *ListProductQuery) (*ListProductRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListProduct", query)
	if err != nil {
		return nil, err
	}

	result := new(ListProductRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
