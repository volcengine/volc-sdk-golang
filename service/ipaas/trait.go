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

func (c *IPaaS) UpdateHostUniversalInfo(ctx context.Context, arg *UpdateHostUniversalInfoBody) (*UpdateHostUniversalInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHostUniversalInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHostUniversalInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpdateInstanceUniversalInfo(ctx context.Context, arg *UpdateInstanceUniversalInfoBody) (*UpdateInstanceUniversalInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateInstanceUniversalInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateInstanceUniversalInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListAsyncAction(ctx context.Context, arg *ListAsyncActionBody) (*ListAsyncActionRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListAsyncAction", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListAsyncActionRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DetailProduct(ctx context.Context, arg *DetailProductQuery) (*DetailProductRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DetailProduct", query)
	if err != nil {
		return nil, err
	}

	result := new(DetailProductRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListInstancePackageBrief(ctx context.Context, arg *ListInstancePackageBriefBody) (*ListInstancePackageBriefRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListInstancePackageBrief", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListInstancePackageBriefRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListProductIDCData(ctx context.Context, arg *ListProductIDCDataQuery) (*ListProductIDCDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListProductIDCData", query)
	if err != nil {
		return nil, err
	}

	result := new(ListProductIDCDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetSecurityGroupVisibleConfig(ctx context.Context, arg *GetSecurityGroupVisibleConfigReq) (*GetSecurityGroupVisibleConfigRes, error) {
	query, err := marshalToQuery(arg.GetSecurityGroupVisibleConfigQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetSecurityGroupVisibleConfigBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetSecurityGroupVisibleConfig", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetSecurityGroupVisibleConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) CreateSecurityGroup(ctx context.Context, arg *CreateSecurityGroupReq) (*CreateSecurityGroupRes, error) {
	query, err := marshalToQuery(arg.CreateSecurityGroupQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateSecurityGroupBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSecurityGroup", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) CreateSecurityRule(ctx context.Context, arg *CreateSecurityRuleReq) (*CreateSecurityRuleRes, error) {
	query, err := marshalToQuery(arg.CreateSecurityRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateSecurityRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSecurityRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSecurityRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpdateSecurityRule(ctx context.Context, arg *UpdateSecurityRuleReq) (*UpdateSecurityRuleRes, error) {
	query, err := marshalToQuery(arg.UpdateSecurityRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateSecurityRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSecurityRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSecurityRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) BindInstanceSecurityGroup(ctx context.Context, arg *BindInstanceSecurityGroupBody) (*BindInstanceSecurityGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindInstanceSecurityGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindInstanceSecurityGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UnbindInstanceSecurityGroup(ctx context.Context, arg *UnbindInstanceSecurityGroupBody) (*UnbindInstanceSecurityGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindInstanceSecurityGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindInstanceSecurityGroupRes)
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

func (c *IPaaS) ModifyInstanceFps(ctx context.Context, arg *ModifyInstanceFpsBody) (*ModifyInstanceFpsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ModifyInstanceFps", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ModifyInstanceFpsRes)
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

func (c *IPaaS) SetInstanceBandwidth(ctx context.Context, arg *SetInstanceBandwidthBody) (*SetInstanceBandwidthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetInstanceBandwidth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetInstanceBandwidthRes)
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

func (c *IPaaS) ListTaskInfo(ctx context.Context, arg *ListTaskInfoQuery) (*ListTaskInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListTaskInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(ListTaskInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) UpgradeImage(ctx context.Context, arg *UpgradeImageBody) (*UpgradeImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpgradeImage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpgradeImageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListAOSPRepoACLEntries(ctx context.Context, arg *ListAOSPRepoACLEntriesBody) (*ListAOSPRepoACLEntriesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListAOSPRepoAclEntries", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListAOSPRepoACLEntriesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) AddAOSPRepoACLEntries(ctx context.Context, arg *AddAOSPRepoACLEntriesBody) (*AddAOSPRepoACLEntriesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddAOSPRepoAclEntries", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddAOSPRepoACLEntriesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GrantGitRepoPermission(ctx context.Context, arg *GrantGitRepoPermissionBody) (*GrantGitRepoPermissionRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GrantGitRepoPermission", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GrantGitRepoPermissionRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) RemoveAOSPRepoACLEntries(ctx context.Context, arg *RemoveAOSPRepoACLEntriesBody) (*RemoveAOSPRepoACLEntriesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RemoveAOSPRepoAclEntries", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RemoveAOSPRepoACLEntriesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetGitRepoUserInfo(ctx context.Context) (*GetGitRepoUserInfoRes, error) {

	data, _, err := c.CtxQuery(ctx, "GetGitRepoUserInfo", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GetGitRepoUserInfoRes)
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

func (c *IPaaS) InitializeHost(ctx context.Context, arg *InitializeHostBody) (*InitializeHostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "InitializeHost", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(InitializeHostRes)
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

func (c *IPaaS) ReconfigureDevicesPackage(ctx context.Context, arg *ReconfigureDevicesPackageBody) (*ReconfigureDevicesPackageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ReconfigureDevicesPackage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ReconfigureDevicesPackageRes)
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

func (c *IPaaS) GetFileDistributionJobDetail(ctx context.Context, arg *GetFileDistributionJobDetailBody) (*GetFileDistributionJobDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetFileDistributionJobDetail", url.Values{}, string(body))
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

func (c *IPaaS) TosPreSignURL(ctx context.Context, arg *TosPreSignURLBody) (*TosPreSignURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "TosPreSignUrl", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(TosPreSignURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetResourceNetworkCurveConsole(ctx context.Context, arg *GetResourceNetworkCurveConsoleQuery) (*GetResourceNetworkCurveConsoleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetResourceNetworkCurveConsole", query)
	if err != nil {
		return nil, err
	}

	result := new(GetResourceNetworkCurveConsoleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetResourcePodCurrentConsole(ctx context.Context, arg *GetResourcePodCurrentConsoleQuery) (*GetResourcePodCurrentConsoleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetResourcePodCurrentConsole", query)
	if err != nil {
		return nil, err
	}

	result := new(GetResourcePodCurrentConsoleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GetResourcePodCurveConsole(ctx context.Context, arg *GetResourcePodCurveConsoleQuery) (*GetResourcePodCurveConsoleRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetResourcePodCurveConsole", query)
	if err != nil {
		return nil, err
	}

	result := new(GetResourcePodCurveConsoleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (c *IPaaS) ListContainerImagesInner(ctx context.Context, arg *ListContainerImagesInnerQuery) (*ListContainerImagesInnerRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListContainerImagesInner", query)
	if err != nil {
		return nil, err
	}

	result := new(ListContainerImagesInnerRes)
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

func (c *IPaaS) GenerateAdbKey(ctx context.Context) (*GenerateAdbKeyRes, error) {

	data, _, err := c.CtxQuery(ctx, "GenerateAdbKey", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GenerateAdbKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListGitRepoWhiteIP(ctx context.Context, arg *ListGitRepoWhiteIPBody) (*ListGitRepoWhiteIPRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListGitRepoWhiteIP", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListGitRepoWhiteIPRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DeleteGitRepoWhiteIP(ctx context.Context, arg *DeleteGitRepoWhiteIPBody) (*DeleteGitRepoWhiteIPRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteGitRepoWhiteIP", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteGitRepoWhiteIPRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) OperateGitRepoWhiteIP(ctx context.Context, arg *OperateGitRepoWhiteIPBody) (*OperateGitRepoWhiteIPRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "OperateGitRepoWhiteIP", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(OperateGitRepoWhiteIPRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) AddGitRepoWhiteIP(ctx context.Context, arg *AddGitRepoWhiteIPBody) (*AddGitRepoWhiteIPRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddGitRepoWhiteIP", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddGitRepoWhiteIPRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListGitRepoSSHKey(ctx context.Context, arg *ListGitRepoSSHKeyBody) (*ListGitRepoSSHKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListGitRepoSSHKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListGitRepoSSHKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DeleteGitRepoSSHKey(ctx context.Context, arg *DeleteGitRepoSSHKeyBody) (*DeleteGitRepoSSHKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteGitRepoSSHKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteGitRepoSSHKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) OperateGitRepoSSHKey(ctx context.Context, arg *OperateGitRepoSSHKeyBody) (*OperateGitRepoSSHKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "OperateGitRepoSSHKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(OperateGitRepoSSHKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) AddGitRepoSSHKey(ctx context.Context, arg *AddGitRepoSSHKeyBody) (*AddGitRepoSSHKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddGitRepoSSHKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddGitRepoSSHKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) GenerateGitRepoSSHKey(ctx context.Context, arg *GenerateGitRepoSSHKeyBody) (*GenerateGitRepoSSHKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GenerateGitRepoSSHKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GenerateGitRepoSSHKeyRes)
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

func (c *IPaaS) AddEventCallback(ctx context.Context, arg *AddEventCallbackBody) (*AddEventCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddEventCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddEventCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) DelEventCallback(ctx context.Context, arg *DelEventCallbackBody) (*DelEventCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DelEventCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DelEventCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) SetEventCallback(ctx context.Context, arg *SetEventCallbackBody) (*SetEventCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetEventCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetEventCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListEventCallback(ctx context.Context, arg *ListEventCallbackBody) (*ListEventCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListEventCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListEventCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ValidateEventCallback(ctx context.Context, arg *ValidateEventCallbackBody) (*ValidateEventCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ValidateEventCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ValidateEventCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListPackage(ctx context.Context, arg *ListPackageBody) (*ListPackageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPackage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPackageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *IPaaS) ListDcCapacity(ctx context.Context, arg *ListDcCapacityBody) (*ListDcCapacityRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListDcCapacity", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListDcCapacityRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
