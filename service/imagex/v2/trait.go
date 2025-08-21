package imagex

import (
	"context"
	"net/url"
)

func (c *Imagex) UpdateImageDomainVolcOrigin(ctx context.Context, arg *UpdateImageDomainVolcOriginReq) (*UpdateImageDomainVolcOriginRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainVolcOriginQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainVolcOriginBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainVolcOrigin", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainVolcOriginRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DelDomain(ctx context.Context, arg *DelDomainReq) (*DelDomainRes, error) {
	query, err := marshalToQuery(arg.DelDomainQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DelDomainBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DelDomain", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DelDomainRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) AddDomainV1(ctx context.Context, arg *AddDomainV1Req) (*AddDomainV1Res, error) {
	query, err := marshalToQuery(arg.AddDomainV1Query)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.AddDomainV1Body)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddDomainV1", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddDomainV1Res)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainIPAuth(ctx context.Context, arg *UpdateImageDomainIPAuthReq) (*UpdateImageDomainIPAuthRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainIPAuthQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainIPAuthBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainIPAuth", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainIPAuthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateRefer(ctx context.Context, arg *UpdateReferReq) (*UpdateReferRes, error) {
	query, err := marshalToQuery(arg.UpdateReferQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateReferBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRefer", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateReferRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainUaAccess(ctx context.Context, arg *UpdateImageDomainUaAccessReq) (*UpdateImageDomainUaAccessRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainUaAccessQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainUaAccessBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainUaAccess", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainUaAccessRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateHTTPS(ctx context.Context, arg *UpdateHTTPSReq) (*UpdateHTTPSRes, error) {
	query, err := marshalToQuery(arg.UpdateHTTPSQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateHTTPSBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHttps", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHTTPSRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainDownloadSpeedLimit(ctx context.Context, arg *UpdateImageDomainDownloadSpeedLimitReq) (*UpdateImageDomainDownloadSpeedLimitRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainDownloadSpeedLimitQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainDownloadSpeedLimitBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainDownloadSpeedLimit", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainDownloadSpeedLimitRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateResponseHeader(ctx context.Context, arg *UpdateResponseHeaderReq) (*UpdateResponseHeaderRes, error) {
	query, err := marshalToQuery(arg.UpdateResponseHeaderQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateResponseHeaderBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateResponseHeader", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateResponseHeaderRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainAreaAccess(ctx context.Context, arg *UpdateImageDomainAreaAccessReq) (*UpdateImageDomainAreaAccessRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainAreaAccessQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainAreaAccessBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainAreaAccess", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainAreaAccessRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateDomainAdaptiveFmt(ctx context.Context, arg *UpdateDomainAdaptiveFmtReq) (*UpdateDomainAdaptiveFmtRes, error) {
	query, err := marshalToQuery(arg.UpdateDomainAdaptiveFmtQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateDomainAdaptiveFmtBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDomainAdaptiveFmt", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDomainAdaptiveFmtRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainConfig(ctx context.Context, arg *UpdateImageDomainConfigReq) (*UpdateImageDomainConfigRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainConfigQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainConfigBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainConfig", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainConfigRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateAdvance(ctx context.Context, arg *UpdateAdvanceReq) (*UpdateAdvanceRes, error) {
	query, err := marshalToQuery(arg.UpdateAdvanceQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateAdvanceBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAdvance", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAdvanceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageDomainBandwidthLimit(ctx context.Context, arg *UpdateImageDomainBandwidthLimitReq) (*UpdateImageDomainBandwidthLimitRes, error) {
	query, err := marshalToQuery(arg.UpdateImageDomainBandwidthLimitQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageDomainBandwidthLimitBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageDomainBandwidthLimit", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageDomainBandwidthLimitRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateSlimConfig(ctx context.Context, arg *UpdateSlimConfigReq) (*UpdateSlimConfigRes, error) {
	query, err := marshalToQuery(arg.UpdateSlimConfigQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateSlimConfigBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSlimConfig", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSlimConfigRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) SetDefaultDomain(ctx context.Context, arg *SetDefaultDomainBody) (*SetDefaultDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SetDefaultDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SetDefaultDomainRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageVolcCdnAccessLog(ctx context.Context, arg *DescribeImageVolcCdnAccessLogReq) (*DescribeImageVolcCdnAccessLogRes, error) {
	query, err := marshalToQuery(arg.DescribeImageVolcCdnAccessLogQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageVolcCdnAccessLogBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageVolcCdnAccessLog", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageVolcCdnAccessLogRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) VerifyDomainOwner(ctx context.Context, arg *VerifyDomainOwnerReq) (*VerifyDomainOwnerRes, error) {
	query, err := marshalToQuery(arg.VerifyDomainOwnerQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.VerifyDomainOwnerBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "VerifyDomainOwner", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(VerifyDomainOwnerRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetResponseHeaderValidateKeys(ctx context.Context) (*GetResponseHeaderValidateKeysRes, error) {

	data, _, err := c.CtxQuery(ctx, "GetResponseHeaderValidateKeys", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GetResponseHeaderValidateKeysRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDomainConfig(ctx context.Context, arg *GetDomainConfigQuery) (*GetDomainConfigRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetDomainConfig", query)
	if err != nil {
		return nil, err
	}

	result := new(GetDomainConfigRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDomainOwnerVerifyContent(ctx context.Context, arg *GetDomainOwnerVerifyContentQuery) (*GetDomainOwnerVerifyContentRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetDomainOwnerVerifyContent", query)
	if err != nil {
		return nil, err
	}

	result := new(GetDomainOwnerVerifyContentRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetServiceDomains(ctx context.Context, arg *GetServiceDomainsQuery) (*GetServiceDomainsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetServiceDomains", query)
	if err != nil {
		return nil, err
	}

	result := new(GetServiceDomainsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageMonitorRules(ctx context.Context, arg *DeleteImageMonitorRulesReq) (*DeleteImageMonitorRulesRes, error) {
	query, err := marshalToQuery(arg.DeleteImageMonitorRulesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageMonitorRulesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageMonitorRules", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageMonitorRulesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageMonitorRecords(ctx context.Context, arg *DeleteImageMonitorRecordsReq) (*DeleteImageMonitorRecordsRes, error) {
	query, err := marshalToQuery(arg.DeleteImageMonitorRecordsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageMonitorRecordsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageMonitorRecords", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageMonitorRecordsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageMonitorRule(ctx context.Context, arg *CreateImageMonitorRuleReq) (*CreateImageMonitorRuleRes, error) {
	query, err := marshalToQuery(arg.CreateImageMonitorRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageMonitorRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageMonitorRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageMonitorRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageMonitorRule(ctx context.Context, arg *UpdateImageMonitorRuleReq) (*UpdateImageMonitorRuleRes, error) {
	query, err := marshalToQuery(arg.UpdateImageMonitorRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageMonitorRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageMonitorRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageMonitorRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageMonitorRuleStatus(ctx context.Context, arg *UpdateImageMonitorRuleStatusReq) (*UpdateImageMonitorRuleStatusRes, error) {
	query, err := marshalToQuery(arg.UpdateImageMonitorRuleStatusQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageMonitorRuleStatusBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageMonitorRuleStatus", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageMonitorRuleStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAlertRecords(ctx context.Context, arg *GetImageAlertRecordsReq) (*GetImageAlertRecordsRes, error) {
	query, err := marshalToQuery(arg.GetImageAlertRecordsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageAlertRecordsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageAlertRecords", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageAlertRecordsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageMonitorRules(ctx context.Context, arg *GetImageMonitorRulesQuery) (*GetImageMonitorRulesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageMonitorRules", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageMonitorRulesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageSettingRule(ctx context.Context, arg *CreateImageSettingRuleReq) (*CreateImageSettingRuleRes, error) {
	query, err := marshalToQuery(arg.CreateImageSettingRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageSettingRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageSettingRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageSettingRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageSettingRule(ctx context.Context, arg *DeleteImageSettingRuleReq) (*DeleteImageSettingRuleRes, error) {
	query, err := marshalToQuery(arg.DeleteImageSettingRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageSettingRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageSettingRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageSettingRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageSettingRulePriority(ctx context.Context, arg *UpdateImageSettingRulePriorityReq) (*UpdateImageSettingRulePriorityRes, error) {
	query, err := marshalToQuery(arg.UpdateImageSettingRulePriorityQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageSettingRulePriorityBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageSettingRulePriority", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageSettingRulePriorityRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageSettingRule(ctx context.Context, arg *UpdateImageSettingRuleReq) (*UpdateImageSettingRuleRes, error) {
	query, err := marshalToQuery(arg.UpdateImageSettingRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageSettingRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageSettingRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageSettingRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSettings(ctx context.Context, arg *GetImageSettingsQuery) (*GetImageSettingsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageSettings", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageSettingsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSettingRuleHistory(ctx context.Context, arg *GetImageSettingRuleHistoryQuery) (*GetImageSettingRuleHistoryRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageSettingRuleHistory", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageSettingRuleHistoryRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSettingRules(ctx context.Context, arg *GetImageSettingRulesQuery) (*GetImageSettingRulesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageSettingRules", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageSettingRulesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageMigrateTask(ctx context.Context, arg *CreateImageMigrateTaskBody) (*CreateImageMigrateTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageMigrateTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageMigrateTask(ctx context.Context, arg *DeleteImageMigrateTaskQuery) (*DeleteImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ExportFailedMigrateTask(ctx context.Context, arg *ExportFailedMigrateTaskQuery) (*ExportFailedMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ExportFailedMigrateTask", query)
	if err != nil {
		return nil, err
	}

	result := new(ExportFailedMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTaskStrategy(ctx context.Context, arg *UpdateImageTaskStrategyBody) (*UpdateImageTaskStrategyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTaskStrategy", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTaskStrategyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) TerminateImageMigrateTask(ctx context.Context, arg *TerminateImageMigrateTaskQuery) (*TerminateImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "TerminateImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(TerminateImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetVendorBuckets(ctx context.Context, arg *GetVendorBucketsBody) (*GetVendorBucketsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetVendorBuckets", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetVendorBucketsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageMigrateTasks(ctx context.Context, arg *GetImageMigrateTasksQuery) (*GetImageMigrateTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageMigrateTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageMigrateTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) RerunImageMigrateTask(ctx context.Context, arg *RerunImageMigrateTaskQuery) (*RerunImageMigrateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RerunImageMigrateTask", query, "")
	if err != nil {
		return nil, err
	}

	result := new(RerunImageMigrateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAddOnTag(ctx context.Context, arg *GetImageAddOnTagQuery) (*GetImageAddOnTagRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAddOnTag", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAddOnTagRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCubeUsage(ctx context.Context, arg *DescribeImageXCubeUsageQuery) (*DescribeImageXCubeUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXCubeUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCubeUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequestBandwidth(ctx context.Context, arg *DescribeImageXSourceRequestBandwidthQuery) (*DescribeImageXSourceRequestBandwidthRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequestBandwidth", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequestTraffic(ctx context.Context, arg *DescribeImageXSourceRequestTrafficQuery) (*DescribeImageXSourceRequestTrafficRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequestTraffic", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSourceRequest(ctx context.Context, arg *DescribeImageXSourceRequestQuery) (*DescribeImageXSourceRequestRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSourceRequest", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSourceRequestRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXStorageUsage(ctx context.Context, arg *DescribeImageXStorageUsageQuery) (*DescribeImageXStorageUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXStorageUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXStorageUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBucketRetrievalUsage(ctx context.Context, arg *DescribeImageXBucketRetrievalUsageQuery) (*DescribeImageXBucketRetrievalUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBucketRetrievalUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBucketRetrievalUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXAddOnQPSUsage(ctx context.Context, arg *DescribeImageXAddOnQPSUsageQuery) (*DescribeImageXAddOnQPSUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXAddOnQPSUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXAddOnQPSUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXAIRequestCntUsage(ctx context.Context, arg *DescribeImageXAIRequestCntUsageQuery) (*DescribeImageXAIRequestCntUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXAIRequestCntUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXAIRequestCntUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSummary(ctx context.Context, arg *DescribeImageXSummaryQuery) (*DescribeImageXSummaryRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXSummary", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSummaryRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainTrafficData(ctx context.Context, arg *DescribeImageXDomainTrafficDataQuery) (*DescribeImageXDomainTrafficDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainTrafficData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainTrafficDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainBandwidthData(ctx context.Context, arg *DescribeImageXDomainBandwidthDataQuery) (*DescribeImageXDomainBandwidthDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainBandwidthData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainBandwidthDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXDomainBandwidthNinetyFiveData(ctx context.Context, arg *DescribeImageXDomainBandwidthNinetyFiveDataQuery) (*DescribeImageXDomainBandwidthNinetyFiveDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXDomainBandwidthNinetyFiveData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXDomainBandwidthNinetyFiveDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBucketUsage(ctx context.Context, arg *DescribeImageXBucketUsageQuery) (*DescribeImageXBucketUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBucketUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBucketUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBillingRequestCntUsage(ctx context.Context, arg *DescribeImageXBillingRequestCntUsageQuery) (*DescribeImageXBillingRequestCntUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBillingRequestCntUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBillingRequestCntUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXRequestCntUsage(ctx context.Context, arg *DescribeImageXRequestCntUsageQuery) (*DescribeImageXRequestCntUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXRequestCntUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXRequestCntUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXBaseOpUsage(ctx context.Context, arg *DescribeImageXBaseOpUsageQuery) (*DescribeImageXBaseOpUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXBaseOpUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXBaseOpUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCompressUsage(ctx context.Context, arg *DescribeImageXCompressUsageQuery) (*DescribeImageXCompressUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXCompressUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCompressUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXScreenshotUsage(ctx context.Context, arg *DescribeImageXScreenshotUsageQuery) (*DescribeImageXScreenshotUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXScreenshotUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXScreenshotUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXVideoClipDurationUsage(ctx context.Context, arg *DescribeImageXVideoClipDurationUsageQuery) (*DescribeImageXVideoClipDurationUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXVideoClipDurationUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXVideoClipDurationUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMultiCompressUsage(ctx context.Context, arg *DescribeImageXMultiCompressUsageQuery) (*DescribeImageXMultiCompressUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXMultiCompressUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMultiCompressUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequest(ctx context.Context, arg *DescribeImageXEdgeRequestQuery) (*DescribeImageXEdgeRequestRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequest", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestBandwidth(ctx context.Context, arg *DescribeImageXEdgeRequestBandwidthQuery) (*DescribeImageXEdgeRequestBandwidthRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestBandwidth", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestTraffic(ctx context.Context, arg *DescribeImageXEdgeRequestTrafficQuery) (*DescribeImageXEdgeRequestTrafficRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestTraffic", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXEdgeRequestRegions(ctx context.Context, arg *DescribeImageXEdgeRequestRegionsQuery) (*DescribeImageXEdgeRequestRegionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXEdgeRequestRegions", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXEdgeRequestRegionsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestHTTPCodeByTime(ctx context.Context, arg *DescribeImageXMirrorRequestHTTPCodeByTimeReq) (*DescribeImageXMirrorRequestHTTPCodeByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXMirrorRequestHTTPCodeByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXMirrorRequestHTTPCodeByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestHttpCodeByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestHTTPCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestHTTPCodeOverview(ctx context.Context, arg *DescribeImageXMirrorRequestHTTPCodeOverviewReq) (*DescribeImageXMirrorRequestHTTPCodeOverviewRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXMirrorRequestHTTPCodeOverviewQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXMirrorRequestHTTPCodeOverviewBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestHttpCodeOverview", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestHTTPCodeOverviewRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestTraffic(ctx context.Context, arg *DescribeImageXMirrorRequestTrafficBody) (*DescribeImageXMirrorRequestTrafficRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestTraffic", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestTrafficRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXMirrorRequestBandwidth(ctx context.Context, arg *DescribeImageXMirrorRequestBandwidthBody) (*DescribeImageXMirrorRequestBandwidthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXMirrorRequestBandwidth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXMirrorRequestBandwidthRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXServerQPSUsage(ctx context.Context, arg *DescribeImageXServerQPSUsageQuery) (*DescribeImageXServerQPSUsageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXServerQPSUsage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXServerQPSUsageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHitRateTrafficData(ctx context.Context, arg *DescribeImageXHitRateTrafficDataQuery) (*DescribeImageXHitRateTrafficDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXHitRateTrafficData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHitRateTrafficDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHitRateRequestData(ctx context.Context, arg *DescribeImageXHitRateRequestDataQuery) (*DescribeImageXHitRateRequestDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXHitRateRequestData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHitRateRequestDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCDNTopRequestData(ctx context.Context, arg *DescribeImageXCDNTopRequestDataQuery) (*DescribeImageXCDNTopRequestDataRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXCDNTopRequestData", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCDNTopRequestDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeFileInSizeByTime(ctx context.Context, arg *DescribeImageXHeifEncodeFileInSizeByTimeReq) (*DescribeImageXHeifEncodeFileInSizeByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeFileInSizeByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeFileInSizeByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeFileInSizeByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeFileInSizeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeFileOutSizeByTime(ctx context.Context, arg *DescribeImageXHeifEncodeFileOutSizeByTimeReq) (*DescribeImageXHeifEncodeFileOutSizeByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeFileOutSizeByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeFileOutSizeByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeFileOutSizeByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeFileOutSizeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeSuccessCountByTime(ctx context.Context, arg *DescribeImageXHeifEncodeSuccessCountByTimeReq) (*DescribeImageXHeifEncodeSuccessCountByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeSuccessCountByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeSuccessCountByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeSuccessCountByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeSuccessCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeSuccessRateByTime(ctx context.Context, arg *DescribeImageXHeifEncodeSuccessRateByTimeReq) (*DescribeImageXHeifEncodeSuccessRateByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeSuccessRateByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeSuccessRateByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeSuccessRateByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeDurationByTime(ctx context.Context, arg *DescribeImageXHeifEncodeDurationByTimeReq) (*DescribeImageXHeifEncodeDurationByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeDurationByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeDurationByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeDurationByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeDurationByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXHeifEncodeErrorCodeByTime(ctx context.Context, arg *DescribeImageXHeifEncodeErrorCodeByTimeReq) (*DescribeImageXHeifEncodeErrorCodeByTimeRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXHeifEncodeErrorCodeByTimeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXHeifEncodeErrorCodeByTimeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXHeifEncodeErrorCodeByTime", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXHeifEncodeErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedResolutionRatioAll(ctx context.Context, arg *DescribeImageXExceedResolutionRatioAllBody) (*DescribeImageXExceedResolutionRatioAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedResolutionRatioAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedResolutionRatioAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedFileSize(ctx context.Context, arg *DescribeImageXExceedFileSizeBody) (*DescribeImageXExceedFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXExceedCountByTime(ctx context.Context, arg *DescribeImageXExceedCountByTimeBody) (*DescribeImageXExceedCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXExceedCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXExceedCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXServiceQuality(ctx context.Context, arg *DescribeImageXServiceQualityQuery) (*DescribeImageXServiceQualityRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeImageXServiceQuality", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXServiceQualityRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryApps(ctx context.Context, arg *GetImageXQueryAppsQuery) (*GetImageXQueryAppsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryApps", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryAppsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryRegions(ctx context.Context, arg *GetImageXQueryRegionsQuery) (*GetImageXQueryRegionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryRegions", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryRegionsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryDims(ctx context.Context, arg *GetImageXQueryDimsQuery) (*GetImageXQueryDimsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryDims", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryDimsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageXQueryVals(ctx context.Context, arg *GetImageXQueryValsQuery) (*GetImageXQueryValsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageXQueryVals", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageXQueryValsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadCountByTime(ctx context.Context, arg *DescribeImageXUploadCountByTimeBody) (*DescribeImageXUploadCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadDuration(ctx context.Context, arg *DescribeImageXUploadDurationBody) (*DescribeImageXUploadDurationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadDuration", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadDurationRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSuccessRateByTime(ctx context.Context, arg *DescribeImageXUploadSuccessRateByTimeBody) (*DescribeImageXUploadSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadFileSize(ctx context.Context, arg *DescribeImageXUploadFileSizeBody) (*DescribeImageXUploadFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadErrorCodeByTime(ctx context.Context, arg *DescribeImageXUploadErrorCodeByTimeBody) (*DescribeImageXUploadErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadErrorCodeAll(ctx context.Context, arg *DescribeImageXUploadErrorCodeAllBody) (*DescribeImageXUploadErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSpeed(ctx context.Context, arg *DescribeImageXUploadSpeedBody) (*DescribeImageXUploadSpeedRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSpeed", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSpeedRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXUploadSegmentSpeedByTime(ctx context.Context, arg *DescribeImageXUploadSegmentSpeedByTimeBody) (*DescribeImageXUploadSegmentSpeedByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXUploadSegmentSpeedByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXUploadSegmentSpeedByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnSuccessRateByTime(ctx context.Context, arg *DescribeImageXCdnSuccessRateByTimeBody) (*DescribeImageXCdnSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnSuccessRateAll(ctx context.Context, arg *DescribeImageXCdnSuccessRateAllBody) (*DescribeImageXCdnSuccessRateAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnSuccessRateAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnSuccessRateAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnErrorCodeByTime(ctx context.Context, arg *DescribeImageXCdnErrorCodeByTimeBody) (*DescribeImageXCdnErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnErrorCodeAll(ctx context.Context, arg *DescribeImageXCdnErrorCodeAllBody) (*DescribeImageXCdnErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnDurationDetailByTime(ctx context.Context, arg *DescribeImageXCdnDurationDetailByTimeBody) (*DescribeImageXCdnDurationDetailByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnDurationDetailByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnDurationDetailByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnDurationAll(ctx context.Context, arg *DescribeImageXCdnDurationAllBody) (*DescribeImageXCdnDurationAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnDurationAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnDurationAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnReuseRateByTime(ctx context.Context, arg *DescribeImageXCdnReuseRateByTimeBody) (*DescribeImageXCdnReuseRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnReuseRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnReuseRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnReuseRateAll(ctx context.Context, arg *DescribeImageXCdnReuseRateAllBody) (*DescribeImageXCdnReuseRateAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnReuseRateAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnReuseRateAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXCdnProtocolRateByTime(ctx context.Context, arg *DescribeImageXCdnProtocolRateByTimeBody) (*DescribeImageXCdnProtocolRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXCdnProtocolRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXCdnProtocolRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientFailureRate(ctx context.Context, arg *DescribeImageXClientFailureRateBody) (*DescribeImageXClientFailureRateRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientFailureRate", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientFailureRateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDecodeSuccessRateByTime(ctx context.Context, arg *DescribeImageXClientDecodeSuccessRateByTimeBody) (*DescribeImageXClientDecodeSuccessRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDecodeSuccessRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDecodeSuccessRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDecodeDurationByTime(ctx context.Context, arg *DescribeImageXClientDecodeDurationByTimeBody) (*DescribeImageXClientDecodeDurationByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDecodeDurationByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDecodeDurationByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientQueueDurationByTime(ctx context.Context, arg *DescribeImageXClientQueueDurationByTimeBody) (*DescribeImageXClientQueueDurationByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientQueueDurationByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientQueueDurationByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientErrorCodeByTime(ctx context.Context, arg *DescribeImageXClientErrorCodeByTimeBody) (*DescribeImageXClientErrorCodeByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientErrorCodeByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientErrorCodeByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientErrorCodeAll(ctx context.Context, arg *DescribeImageXClientErrorCodeAllBody) (*DescribeImageXClientErrorCodeAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientErrorCodeAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientErrorCodeAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientLoadDuration(ctx context.Context, arg *DescribeImageXClientLoadDurationBody) (*DescribeImageXClientLoadDurationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientLoadDuration", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientLoadDurationRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientLoadDurationAll(ctx context.Context, arg *DescribeImageXClientLoadDurationAllBody) (*DescribeImageXClientLoadDurationAllRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientLoadDurationAll", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientLoadDurationAllRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientSdkVerByTime(ctx context.Context, arg *DescribeImageXClientSdkVerByTimeBody) (*DescribeImageXClientSdkVerByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientSdkVerByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientSdkVerByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientFileSize(ctx context.Context, arg *DescribeImageXClientFileSizeBody) (*DescribeImageXClientFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopFileSize(ctx context.Context, arg *DescribeImageXClientTopFileSizeBody) (*DescribeImageXClientTopFileSizeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopFileSize", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopFileSizeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientCountByTime(ctx context.Context, arg *DescribeImageXClientCountByTimeBody) (*DescribeImageXClientCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientQualityRateByTime(ctx context.Context, arg *DescribeImageXClientQualityRateByTimeBody) (*DescribeImageXClientQualityRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientQualityRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientQualityRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopQualityURL(ctx context.Context, arg *DescribeImageXClientTopQualityURLBody) (*DescribeImageXClientTopQualityURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopQualityURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopQualityURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientDemotionRateByTime(ctx context.Context, arg *DescribeImageXClientDemotionRateByTimeBody) (*DescribeImageXClientDemotionRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientDemotionRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientDemotionRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientTopDemotionURL(ctx context.Context, arg *DescribeImageXClientTopDemotionURLBody) (*DescribeImageXClientTopDemotionURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientTopDemotionURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientTopDemotionURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXClientScoreByTime(ctx context.Context, arg *DescribeImageXClientScoreByTimeBody) (*DescribeImageXClientScoreByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXClientScoreByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXClientScoreByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleCountByTime(ctx context.Context, arg *DescribeImageXSensibleCountByTimeBody) (*DescribeImageXSensibleCountByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleCountByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleCountByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleCacheHitRateByTime(ctx context.Context, arg *DescribeImageXSensibleCacheHitRateByTimeBody) (*DescribeImageXSensibleCacheHitRateByTimeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleCacheHitRateByTime", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleCacheHitRateByTimeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopSizeURL(ctx context.Context, arg *DescribeImageXSensibleTopSizeURLBody) (*DescribeImageXSensibleTopSizeURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopSizeURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopSizeURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopResolutionURL(ctx context.Context, arg *DescribeImageXSensibleTopResolutionURLBody) (*DescribeImageXSensibleTopResolutionURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopResolutionURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopResolutionURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopRAMURL(ctx context.Context, arg *DescribeImageXSensibleTopRAMURLReq) (*DescribeImageXSensibleTopRAMURLRes, error) {
	query, err := marshalToQuery(arg.DescribeImageXSensibleTopRAMURLQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DescribeImageXSensibleTopRAMURLBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopRamURL", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopRAMURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DescribeImageXSensibleTopUnknownURL(ctx context.Context, arg *DescribeImageXSensibleTopUnknownURLBody) (*DescribeImageXSensibleTopUnknownURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeImageXSensibleTopUnknownURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeImageXSensibleTopUnknownURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateBatchProcessTask(ctx context.Context, arg *CreateBatchProcessTaskReq) (*CreateBatchProcessTaskRes, error) {
	query, err := marshalToQuery(arg.CreateBatchProcessTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateBatchProcessTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateBatchProcessTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateBatchProcessTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetBatchProcessResult(ctx context.Context, arg *GetBatchProcessResultReq) (*GetBatchProcessResultRes, error) {
	query, err := marshalToQuery(arg.GetBatchProcessResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetBatchProcessResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetBatchProcessResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetBatchProcessResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetBatchTaskInfo(ctx context.Context, arg *GetBatchTaskInfoQuery) (*GetBatchTaskInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetBatchTaskInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(GetBatchTaskInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) AIProcess(ctx context.Context, arg *AIProcessReq) (*AIProcessRes, error) {
	query, err := marshalToQuery(arg.AIProcessQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.AIProcessBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AIProcess", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AIProcessRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAITask(ctx context.Context, arg *CreateImageAITaskReq) (*CreateImageAITaskRes, error) {
	query, err := marshalToQuery(arg.CreateImageAITaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageAITaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAITask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAITaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAIProcessQueue(ctx context.Context, arg *CreateImageAIProcessQueueReq) (*CreateImageAIProcessQueueRes, error) {
	query, err := marshalToQuery(arg.CreateImageAIProcessQueueQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageAIProcessQueueBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAIProcessQueue", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAIProcessQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAIProcessQueue(ctx context.Context, arg *DeleteImageAIProcessQueueReq) (*DeleteImageAIProcessQueueRes, error) {
	query, err := marshalToQuery(arg.DeleteImageAIProcessQueueQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageAIProcessQueueBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAIProcessQueue", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAIProcessQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAIProcessCallback(ctx context.Context, arg *CreateImageAIProcessCallbackReq) (*CreateImageAIProcessCallbackRes, error) {
	query, err := marshalToQuery(arg.CreateImageAIProcessCallbackQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageAIProcessCallbackBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAIProcessCallback", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAIProcessCallbackRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAIProcessQueue(ctx context.Context, arg *UpdateImageAIProcessQueueReq) (*UpdateImageAIProcessQueueRes, error) {
	query, err := marshalToQuery(arg.UpdateImageAIProcessQueueQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageAIProcessQueueBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAIProcessQueue", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAIProcessQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAIProcessQueueStatus(ctx context.Context, arg *UpdateImageAIProcessQueueStatusReq) (*UpdateImageAIProcessQueueStatusRes, error) {
	query, err := marshalToQuery(arg.UpdateImageAIProcessQueueStatusQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageAIProcessQueueStatusBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAIProcessQueueStatus", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAIProcessQueueStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAIProcessDetail(ctx context.Context, arg *DeleteImageAIProcessDetailReq) (*DeleteImageAIProcessDetailRes, error) {
	query, err := marshalToQuery(arg.DeleteImageAIProcessDetailQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageAIProcessDetailBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAIProcessDetail", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAIProcessDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAIProcessQueues(ctx context.Context, arg *GetImageAIProcessQueuesQuery) (*GetImageAIProcessQueuesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAIProcessQueues", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAIProcessQueuesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAITasks(ctx context.Context, arg *GetImageAITasksQuery) (*GetImageAITasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAITasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAITasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAIDetails(ctx context.Context, arg *GetImageAIDetailsQuery) (*GetImageAIDetailsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAIDetails", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAIDetailsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ReportEvent(ctx context.Context, arg *ReportEventReq) (*ReportEventRes, error) {
	query, err := marshalToQuery(arg.ReportEventQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.ReportEventBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ReportEvent", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ReportEventRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageResourceStatus(ctx context.Context, arg *UpdateImageResourceStatusReq) (*UpdateImageResourceStatusRes, error) {
	query, err := marshalToQuery(arg.UpdateImageResourceStatusQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageResourceStatusBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageResourceStatus", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageResourceStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateFileStorageClass(ctx context.Context, arg *UpdateFileStorageClassReq) (*UpdateFileStorageClassRes, error) {
	query, err := marshalToQuery(arg.UpdateFileStorageClassQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateFileStorageClassBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateFileStorageClass", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateFileStorageClassRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStorageFiles(ctx context.Context, arg *GetImageStorageFilesQuery) (*GetImageStorageFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageStorageFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageStorageFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageUploadFiles(ctx context.Context, arg *DeleteImageUploadFilesReq) (*DeleteImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg.DeleteImageUploadFilesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageUploadFilesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageUploadFiles", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateFileRestore(ctx context.Context, arg *CreateFileRestoreReq) (*CreateFileRestoreRes, error) {
	query, err := marshalToQuery(arg.CreateFileRestoreQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateFileRestoreBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateFileRestore", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateFileRestoreRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageUploadFiles(ctx context.Context, arg *UpdateImageUploadFilesReq) (*UpdateImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg.UpdateImageUploadFilesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageUploadFilesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageUploadFiles", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CommitImageUpload(ctx context.Context, arg *CommitImageUploadReq) (*CommitImageUploadRes, error) {
	query, err := marshalToQuery(arg.CommitImageUploadQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CommitImageUploadBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CommitImageUpload", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CommitImageUploadRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageFileCT(ctx context.Context, arg *UpdateImageFileCTReq) (*UpdateImageFileCTRes, error) {
	query, err := marshalToQuery(arg.UpdateImageFileCTQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageFileCTBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageFileCT", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageFileCTRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ApplyVpcUploadInfo(ctx context.Context, arg *ApplyVpcUploadInfoQuery) (*ApplyVpcUploadInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ApplyVpcUploadInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(ApplyVpcUploadInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) ApplyImageUpload(ctx context.Context, arg *ApplyImageUploadQuery) (*ApplyImageUploadRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ApplyImageUpload", query)
	if err != nil {
		return nil, err
	}

	result := new(ApplyImageUploadRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUploadFile(ctx context.Context, arg *GetImageUploadFileQuery) (*GetImageUploadFileRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUploadFile", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUploadFileRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUploadFiles(ctx context.Context, arg *GetImageUploadFilesQuery) (*GetImageUploadFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUploadFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUploadFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageUpdateFiles(ctx context.Context, arg *GetImageUpdateFilesQuery) (*GetImageUpdateFilesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageUpdateFiles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageUpdateFilesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) PreviewImageUploadFile(ctx context.Context, arg *PreviewImageUploadFileQuery) (*PreviewImageUploadFileRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "PreviewImageUploadFile", query)
	if err != nil {
		return nil, err
	}

	result := new(PreviewImageUploadFileRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEraseResult(ctx context.Context, arg *GetImageEraseResultBody) (*GetImageEraseResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageEraseResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageEraseResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageService(ctx context.Context, arg *GetImageServiceQuery) (*GetImageServiceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageService", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAllImageServices(ctx context.Context, arg *GetAllImageServicesQuery) (*GetAllImageServicesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAllImageServices", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAllImageServicesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageCompressTask(ctx context.Context, arg *CreateImageCompressTaskReq) (*CreateImageCompressTaskRes, error) {
	query, err := marshalToQuery(arg.CreateImageCompressTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageCompressTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageCompressTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageCompressTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) FetchImageURL(ctx context.Context, arg *FetchImageURLBody) (*FetchImageURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "FetchImageUrl", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(FetchImageURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageStorageTTL(ctx context.Context, arg *UpdateImageStorageTTLBody) (*UpdateImageStorageTTLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageStorageTTL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageStorageTTLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCompressTaskInfo(ctx context.Context, arg *GetCompressTaskInfoQuery) (*GetCompressTaskInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetCompressTaskInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(GetCompressTaskInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetURLFetchTask(ctx context.Context, arg *GetURLFetchTaskQuery) (*GetURLFetchTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetUrlFetchTask", query)
	if err != nil {
		return nil, err
	}

	result := new(GetURLFetchTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetResourceURL(ctx context.Context, arg *GetResourceURLQuery) (*GetResourceURLRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetResourceURL", query)
	if err != nil {
		return nil, err
	}

	result := new(GetResourceURLRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageFromURI(ctx context.Context, arg *CreateImageFromURIReq) (*CreateImageFromURIRes, error) {
	query, err := marshalToQuery(arg.CreateImageFromURIQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageFromURIBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageFromUri", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageFromURIRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageFileKey(ctx context.Context, arg *UpdateImageFileKeyReq) (*UpdateImageFileKeyRes, error) {
	query, err := marshalToQuery(arg.UpdateImageFileKeyQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageFileKeyBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageFileKey", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageFileKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageContentTask(ctx context.Context, arg *CreateImageContentTaskBody) (*CreateImageContentTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageContentTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageContentTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageContentTaskDetail(ctx context.Context, arg *GetImageContentTaskDetailBody) (*GetImageContentTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageContentTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageContentTaskDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageContentBlockList(ctx context.Context, arg *GetImageContentBlockListBody) (*GetImageContentBlockListRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageContentBlockList", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageContentBlockListRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeQueue(ctx context.Context, arg *CreateImageTranscodeQueueBody) (*CreateImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTranscodeQueue(ctx context.Context, arg *DeleteImageTranscodeQueueBody) (*DeleteImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTranscodeQueue(ctx context.Context, arg *UpdateImageTranscodeQueueBody) (*UpdateImageTranscodeQueueRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTranscodeQueue", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTranscodeQueueRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageTranscodeQueueStatus(ctx context.Context, arg *UpdateImageTranscodeQueueStatusBody) (*UpdateImageTranscodeQueueStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageTranscodeQueueStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageTranscodeQueueStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTranscodeQueues(ctx context.Context, arg *GetImageTranscodeQueuesQuery) (*GetImageTranscodeQueuesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTranscodeQueues", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTranscodeQueuesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeTask(ctx context.Context, arg *CreateImageTranscodeTaskBody) (*CreateImageTranscodeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTranscodeDetails(ctx context.Context, arg *GetImageTranscodeDetailsQuery) (*GetImageTranscodeDetailsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTranscodeDetails", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTranscodeDetailsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTranscodeCallback(ctx context.Context, arg *CreateImageTranscodeCallbackBody) (*CreateImageTranscodeCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTranscodeCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTranscodeCallbackRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTranscodeDetail(ctx context.Context, arg *DeleteImageTranscodeDetailBody) (*DeleteImageTranscodeDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTranscodeDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTranscodeDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImagePSDetection(ctx context.Context, arg *GetImagePSDetectionReq) (*GetImagePSDetectionRes, error) {
	query, err := marshalToQuery(arg.GetImagePSDetectionQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImagePSDetectionBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImagePSDetection", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImagePSDetectionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSuperResolutionResult(ctx context.Context, arg *GetImageSuperResolutionResultBody) (*GetImageSuperResolutionResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageSuperResolutionResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageSuperResolutionResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDenoisingImage(ctx context.Context, arg *GetDenoisingImageReq) (*GetDenoisingImageRes, error) {
	query, err := marshalToQuery(arg.GetDenoisingImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetDenoisingImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetDenoisingImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetDenoisingImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageDuplicateDetection(ctx context.Context, arg *GetImageDuplicateDetectionReq) (*GetImageDuplicateDetectionRes, error) {
	query, err := marshalToQuery(arg.GetImageDuplicateDetectionQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageDuplicateDetectionBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageDuplicateDetection", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageDuplicateDetectionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageOCRV2(ctx context.Context, arg *GetImageOCRV2Req) (*GetImageOCRV2Res, error) {
	query, err := marshalToQuery(arg.GetImageOCRV2Query)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageOCRV2Body)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageOCRV2", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageOCRV2Res)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageBgFillResult(ctx context.Context, arg *GetImageBgFillResultBody) (*GetImageBgFillResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageBgFillResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageBgFillResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetSegmentImage(ctx context.Context, arg *GetSegmentImageReq) (*GetSegmentImageRes, error) {
	query, err := marshalToQuery(arg.GetSegmentImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetSegmentImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetSegmentImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetSegmentImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageSmartCropResult(ctx context.Context, arg *GetImageSmartCropResultBody) (*GetImageSmartCropResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageSmartCropResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageSmartCropResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageComicResult(ctx context.Context, arg *GetImageComicResultBody) (*GetImageComicResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageComicResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageComicResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEnhanceResult(ctx context.Context, arg *GetImageEnhanceResultBody) (*GetImageEnhanceResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageEnhanceResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageEnhanceResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageQuality(ctx context.Context, arg *GetImageQualityReq) (*GetImageQualityRes, error) {
	query, err := marshalToQuery(arg.GetImageQualityQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageQualityBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageQuality", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageQualityRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetLicensePlateDetection(ctx context.Context, arg *GetLicensePlateDetectionReq) (*GetLicensePlateDetectionRes, error) {
	query, err := marshalToQuery(arg.GetLicensePlateDetectionQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetLicensePlateDetectionBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetLicensePlateDetection", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetLicensePlateDetectionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetPrivateImageType(ctx context.Context, arg *GetPrivateImageTypeReq) (*GetPrivateImageTypeRes, error) {
	query, err := marshalToQuery(arg.GetPrivateImageTypeQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetPrivateImageTypeBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetPrivateImageType", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetPrivateImageTypeRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateCVImageGenerateTask(ctx context.Context, arg *CreateCVImageGenerateTaskReq) (*CreateCVImageGenerateTaskRes, error) {
	query, err := marshalToQuery(arg.CreateCVImageGenerateTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateCVImageGenerateTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCVImageGenerateTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCVImageGenerateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateHiddenWatermarkImage(ctx context.Context, arg *CreateHiddenWatermarkImageReq) (*CreateHiddenWatermarkImageRes, error) {
	query, err := marshalToQuery(arg.CreateHiddenWatermarkImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateHiddenWatermarkImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateHiddenWatermarkImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateHiddenWatermarkImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateHmExtractTask(ctx context.Context, arg *CreateHmExtractTaskReq) (*CreateHmExtractTaskRes, error) {
	query, err := marshalToQuery(arg.CreateHmExtractTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateHmExtractTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateHmExtractTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateHmExtractTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageExifData(ctx context.Context, arg *UpdateImageExifDataReq) (*UpdateImageExifDataRes, error) {
	query, err := marshalToQuery(arg.UpdateImageExifDataQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageExifDataBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageExifData", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageExifDataRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageDetectResult(ctx context.Context, arg *GetImageDetectResultReq) (*GetImageDetectResultRes, error) {
	query, err := marshalToQuery(arg.GetImageDetectResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageDetectResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageDetectResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageDetectResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCVImageGenerateResult(ctx context.Context, arg *GetCVImageGenerateResultReq) (*GetCVImageGenerateResultRes, error) {
	query, err := marshalToQuery(arg.GetCVImageGenerateResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetCVImageGenerateResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCVImageGenerateResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCVImageGenerateResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageHmExtract(ctx context.Context, arg *CreateImageHmExtractQuery) (*CreateImageHmExtractRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageHmExtract", query, "")
	if err != nil {
		return nil, err
	}

	result := new(CreateImageHmExtractRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCVTextGenerateImage(ctx context.Context, arg *GetCVTextGenerateImageReq) (*GetCVTextGenerateImageRes, error) {
	query, err := marshalToQuery(arg.GetCVTextGenerateImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetCVTextGenerateImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCVTextGenerateImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCVTextGenerateImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCVImageGenerateTask(ctx context.Context, arg *GetCVImageGenerateTaskReq) (*GetCVImageGenerateTaskRes, error) {
	query, err := marshalToQuery(arg.GetCVImageGenerateTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetCVImageGenerateTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCVImageGenerateTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCVImageGenerateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageHmEmbed(ctx context.Context, arg *CreateImageHmEmbedBody) (*CreateImageHmEmbedRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageHmEmbed", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageHmEmbedRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCVAnimeGenerateImage(ctx context.Context, arg *GetCVAnimeGenerateImageReq) (*GetCVAnimeGenerateImageRes, error) {
	query, err := marshalToQuery(arg.GetCVAnimeGenerateImageQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetCVAnimeGenerateImageBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetCVAnimeGenerateImage", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCVAnimeGenerateImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetComprehensiveEnhanceImage(ctx context.Context, arg *GetComprehensiveEnhanceImageBody) (*GetComprehensiveEnhanceImageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetComprehensiveEnhanceImage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetComprehensiveEnhanceImageRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAiGenerateTask(ctx context.Context, arg *GetImageAiGenerateTaskQuery) (*GetImageAiGenerateTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAiGenerateTask", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAiGenerateTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetProductAIGCResult(ctx context.Context, arg *GetProductAIGCResultReq) (*GetProductAIGCResultRes, error) {
	query, err := marshalToQuery(arg.GetProductAIGCResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetProductAIGCResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetProductAIGCResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetProductAIGCResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageEraseModels(ctx context.Context, arg *GetImageEraseModelsQuery) (*GetImageEraseModelsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageEraseModels", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageEraseModelsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetDedupTaskStatus(ctx context.Context, arg *GetDedupTaskStatusQuery) (*GetDedupTaskStatusRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetDedupTaskStatus", query)
	if err != nil {
		return nil, err
	}

	result := new(GetDedupTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageHmExtractTaskInfo(ctx context.Context, arg *GetImageHmExtractTaskInfoReq) (*GetImageHmExtractTaskInfoRes, error) {
	query, err := marshalToQuery(arg.GetImageHmExtractTaskInfoQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageHmExtractTaskInfoBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageHmExtractTaskInfo", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageHmExtractTaskInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageService(ctx context.Context, arg *CreateImageServiceBody) (*CreateImageServiceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageService", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageService(ctx context.Context, arg *DeleteImageServiceQuery) (*DeleteImageServiceRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageService", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageServiceRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuthKey(ctx context.Context, arg *UpdateImageAuthKeyReq) (*UpdateImageAuthKeyRes, error) {
	query, err := marshalToQuery(arg.UpdateImageAuthKeyQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageAuthKeyBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuthKey", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuthKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateResEventRule(ctx context.Context, arg *UpdateResEventRuleReq) (*UpdateResEventRuleRes, error) {
	query, err := marshalToQuery(arg.UpdateResEventRuleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateResEventRuleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateResEventRule", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateResEventRuleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateServiceName(ctx context.Context, arg *UpdateServiceNameReq) (*UpdateServiceNameRes, error) {
	query, err := marshalToQuery(arg.UpdateServiceNameQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateServiceNameBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateServiceName", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateServiceNameRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateStorageRules(ctx context.Context, arg *UpdateStorageRulesReq) (*UpdateStorageRulesRes, error) {
	query, err := marshalToQuery(arg.UpdateStorageRulesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateStorageRulesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateStorageRules", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateStorageRulesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateStorageRulesV2(ctx context.Context, arg *UpdateStorageRulesV2Req) (*UpdateStorageRulesV2Res, error) {
	query, err := marshalToQuery(arg.UpdateStorageRulesV2Query)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateStorageRulesV2Body)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateStorageRulesV2", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateStorageRulesV2Res)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageObjectAccess(ctx context.Context, arg *UpdateImageObjectAccessReq) (*UpdateImageObjectAccessRes, error) {
	query, err := marshalToQuery(arg.UpdateImageObjectAccessQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageObjectAccessBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageObjectAccess", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageObjectAccessRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageUploadOverwrite(ctx context.Context, arg *UpdateImageUploadOverwriteReq) (*UpdateImageUploadOverwriteRes, error) {
	query, err := marshalToQuery(arg.UpdateImageUploadOverwriteQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageUploadOverwriteBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageUploadOverwrite", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageUploadOverwriteRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageMirrorConf(ctx context.Context, arg *UpdateImageMirrorConfReq) (*UpdateImageMirrorConfRes, error) {
	query, err := marshalToQuery(arg.UpdateImageMirrorConfQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageMirrorConfBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageMirrorConf", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageMirrorConfRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageServiceSubscription(ctx context.Context, arg *GetImageServiceSubscriptionQuery) (*GetImageServiceSubscriptionRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageServiceSubscription", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageServiceSubscriptionRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuthKey(ctx context.Context, arg *GetImageAuthKeyQuery) (*GetImageAuthKeyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuthKey", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuthKeyRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAnalyzeTask(ctx context.Context, arg *CreateImageAnalyzeTaskBody) (*CreateImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAnalyzeTaskRun(ctx context.Context, arg *DeleteImageAnalyzeTaskRunBody) (*DeleteImageAnalyzeTaskRunRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAnalyzeTaskRun", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAnalyzeTaskRunRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAnalyzeTask(ctx context.Context, arg *DeleteImageAnalyzeTaskBody) (*DeleteImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAnalyzeTaskStatus(ctx context.Context, arg *UpdateImageAnalyzeTaskStatusBody) (*UpdateImageAnalyzeTaskStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAnalyzeTaskStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAnalyzeTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAnalyzeTask(ctx context.Context, arg *UpdateImageAnalyzeTaskBody) (*UpdateImageAnalyzeTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAnalyzeTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAnalyzeTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAnalyzeTasks(ctx context.Context, arg *GetImageAnalyzeTasksQuery) (*GetImageAnalyzeTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAnalyzeTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAnalyzeTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAnalyzeResult(ctx context.Context, arg *GetImageAnalyzeResultQuery) (*GetImageAnalyzeResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAnalyzeResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAnalyzeResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageElements(ctx context.Context, arg *DeleteImageElementsReq) (*DeleteImageElementsRes, error) {
	query, err := marshalToQuery(arg.DeleteImageElementsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageElementsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageElements", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageElementsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageBackgroundColors(ctx context.Context, arg *DeleteImageBackgroundColorsReq) (*DeleteImageBackgroundColorsRes, error) {
	query, err := marshalToQuery(arg.DeleteImageBackgroundColorsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageBackgroundColorsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageBackgroundColors", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageBackgroundColorsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageStyle(ctx context.Context, arg *DeleteImageStyleReq) (*DeleteImageStyleRes, error) {
	query, err := marshalToQuery(arg.DeleteImageStyleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageStyleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageStyle", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageStyleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageStyle(ctx context.Context, arg *CreateImageStyleReq) (*CreateImageStyleRes, error) {
	query, err := marshalToQuery(arg.CreateImageStyleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageStyleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageStyle", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageStyleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageStyleMeta(ctx context.Context, arg *UpdateImageStyleMetaReq) (*UpdateImageStyleMetaRes, error) {
	query, err := marshalToQuery(arg.UpdateImageStyleMetaQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageStyleMetaBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageStyleMeta", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageStyleMetaRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) AddImageElements(ctx context.Context, arg *AddImageElementsReq) (*AddImageElementsRes, error) {
	query, err := marshalToQuery(arg.AddImageElementsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.AddImageElementsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddImageElements", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddImageElementsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) AddImageBackgroundColors(ctx context.Context, arg *AddImageBackgroundColorsReq) (*AddImageBackgroundColorsRes, error) {
	query, err := marshalToQuery(arg.AddImageBackgroundColorsQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.AddImageBackgroundColorsBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddImageBackgroundColors", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddImageBackgroundColorsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageStyle(ctx context.Context, arg *UpdateImageStyleReq) (*UpdateImageStyleRes, error) {
	query, err := marshalToQuery(arg.UpdateImageStyleQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateImageStyleBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageStyle", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageStyleRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageFonts(ctx context.Context, arg *GetImageFontsQuery) (*GetImageFontsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageFonts", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageFontsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageElements(ctx context.Context, arg *GetImageElementsQuery) (*GetImageElementsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageElements", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageElementsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageBackgroundColors(ctx context.Context, arg *GetImageBackgroundColorsQuery) (*GetImageBackgroundColorsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageBackgroundColors", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageBackgroundColorsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStyles(ctx context.Context, arg *GetImageStylesQuery) (*GetImageStylesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageStyles", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageStylesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStyleDetail(ctx context.Context, arg *GetImageStyleDetailQuery) (*GetImageStyleDetailRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageStyleDetail", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageStyleDetailRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageStyleResult(ctx context.Context, arg *GetImageStyleResultReq) (*GetImageStyleResultRes, error) {
	query, err := marshalToQuery(arg.GetImageStyleResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetImageStyleResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetImageStyleResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetImageStyleResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DownloadCert(ctx context.Context, arg *DownloadCertQuery) (*DownloadCertRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DownloadCert", query)
	if err != nil {
		return nil, err
	}

	result := new(DownloadCertRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAllDomainCert(ctx context.Context, arg *GetImageAllDomainCertQuery) (*GetImageAllDomainCertRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAllDomainCert", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAllDomainCertRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetCertInfo(ctx context.Context, arg *GetCertInfoQuery) (*GetCertInfoRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetCertInfo", query)
	if err != nil {
		return nil, err
	}

	result := new(GetCertInfoRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAllCerts(ctx context.Context, arg *GetAllCertsQuery) (*GetAllCertsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAllCerts", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAllCertsRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTemplate(ctx context.Context, arg *CreateImageTemplateReq) (*CreateImageTemplateRes, error) {
	query, err := marshalToQuery(arg.CreateImageTemplateQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageTemplateBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTemplate", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteTemplatesFromBin(ctx context.Context, arg *DeleteTemplatesFromBinReq) (*DeleteTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg.DeleteTemplatesFromBinQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteTemplatesFromBinBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTemplatesFromBin", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageTemplate(ctx context.Context, arg *DeleteImageTemplateReq) (*DeleteImageTemplateRes, error) {
	query, err := marshalToQuery(arg.DeleteImageTemplateQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.DeleteImageTemplateBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageTemplate", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageTemplatesByImport(ctx context.Context, arg *CreateImageTemplatesByImportReq) (*CreateImageTemplatesByImportRes, error) {
	query, err := marshalToQuery(arg.CreateImageTemplatesByImportQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateImageTemplatesByImportBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageTemplatesByImport", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageTemplatesByImportRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateTemplatesFromBin(ctx context.Context, arg *CreateTemplatesFromBinReq) (*CreateTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg.CreateTemplatesFromBinQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateTemplatesFromBinBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTemplatesFromBin", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageTemplate(ctx context.Context, arg *GetImageTemplateQuery) (*GetImageTemplateRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageTemplate", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageTemplateRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetTemplatesFromBin(ctx context.Context, arg *GetTemplatesFromBinQuery) (*GetTemplatesFromBinRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetTemplatesFromBin", query)
	if err != nil {
		return nil, err
	}

	result := new(GetTemplatesFromBinRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAllImageTemplates(ctx context.Context, arg *GetAllImageTemplatesQuery) (*GetAllImageTemplatesRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAllImageTemplates", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAllImageTemplatesRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageAuditTask(ctx context.Context, arg *CreateImageAuditTaskBody) (*CreateImageAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateVideoAuditTask(ctx context.Context, arg *CreateVideoAuditTaskReq) (*CreateVideoAuditTaskRes, error) {
	query, err := marshalToQuery(arg.CreateVideoAuditTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateVideoAuditTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateVideoAuditTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateVideoAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateAudioAuditTask(ctx context.Context, arg *CreateAudioAuditTaskReq) (*CreateAudioAuditTaskRes, error) {
	query, err := marshalToQuery(arg.CreateAudioAuditTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.CreateAudioAuditTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateAudioAuditTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateAudioAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) DeleteImageAuditResult(ctx context.Context, arg *DeleteImageAuditResultBody) (*DeleteImageAuditResultRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteImageAuditResult", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteImageAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetSyncAuditResult(ctx context.Context, arg *GetSyncAuditResultReq) (*GetSyncAuditResultRes, error) {
	query, err := marshalToQuery(arg.GetSyncAuditResultQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetSyncAuditResultBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetSyncAuditResult", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetSyncAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) SingleImageAudit(ctx context.Context, arg *SingleImageAuditReq) (*SingleImageAuditRes, error) {
	query, err := marshalToQuery(arg.SingleImageAuditQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.SingleImageAuditBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "SingleImageAudit", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(SingleImageAuditRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) BatchImageAudit(ctx context.Context, arg *BatchImageAuditReq) (*BatchImageAuditRes, error) {
	query, err := marshalToQuery(arg.BatchImageAuditQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.BatchImageAuditBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BatchImageAudit", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BatchImageAuditRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuditTaskStatus(ctx context.Context, arg *UpdateImageAuditTaskStatusBody) (*UpdateImageAuditTaskStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuditTaskStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuditTaskStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateImageAuditTask(ctx context.Context, arg *UpdateImageAuditTaskBody) (*UpdateImageAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateImageAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateImageAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateAuditImageStatus(ctx context.Context, arg *UpdateAuditImageStatusBody) (*UpdateAuditImageStatusRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAuditImageStatus", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAuditImageStatusRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateVideoAuditTask(ctx context.Context, arg *UpdateVideoAuditTaskReq) (*UpdateVideoAuditTaskRes, error) {
	query, err := marshalToQuery(arg.UpdateVideoAuditTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateVideoAuditTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateVideoAuditTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateVideoAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) UpdateAudioAuditTask(ctx context.Context, arg *UpdateAudioAuditTaskReq) (*UpdateAudioAuditTaskRes, error) {
	query, err := marshalToQuery(arg.UpdateAudioAuditTaskQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.UpdateAudioAuditTaskBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAudioAuditTask", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAudioAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuditTasks(ctx context.Context, arg *GetImageAuditTasksQuery) (*GetImageAuditTasksRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuditTasks", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuditTasksRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuditTaskResult(ctx context.Context, arg *GetImageAuditTaskResultQuery) (*GetImageAuditTaskResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuditTaskResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuditTaskResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetImageAuditResult(ctx context.Context, arg *GetImageAuditResultQuery) (*GetImageAuditResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetImageAuditResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetImageAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAuditEntrysCount(ctx context.Context, arg *GetAuditEntrysCountQuery) (*GetAuditEntrysCountRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAuditEntrysCount", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAuditEntrysCountRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetVideoAuditResult(ctx context.Context, arg *GetVideoAuditResultQuery) (*GetVideoAuditResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetVideoAuditResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetVideoAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) GetAudioAuditResult(ctx context.Context, arg *GetAudioAuditResultQuery) (*GetAudioAuditResultRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "GetAudioAuditResult", query)
	if err != nil {
		return nil, err
	}

	result := new(GetAudioAuditResultRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Imagex) CreateImageRetryAuditTask(ctx context.Context, arg *CreateImageRetryAuditTaskBody) (*CreateImageRetryAuditTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateImageRetryAuditTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateImageRetryAuditTaskRes)
	err = unmarshalInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
