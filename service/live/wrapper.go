package live

import (
	"encoding/json"
	"net/url"
)

func (p *LIVE) commonHandler(api string, query url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Query(api, query)
	if err != nil {
		return statusCode, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *LIVE) commonHandlerJson(api string, query url.Values, resp interface{}, body string) (int, error) {
	respBody, statusCode, err := p.Client.Json(api, query, body)
	if err != nil {
		return statusCode, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *LIVE) ListCommonTransPresetDetail(query url.Values, body string) (*ListCommonTransPresetDetailResp, int, error) {
	resp := new(ListCommonTransPresetDetailResp)
	statusCode, err := p.commonHandlerJson("ListCommonTransPresetDetail", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
通过入参messageType区分不同服务：
	设置直播推流回调 - push
	设置直播断流回调 - push_end
	设置录制回调  - record
	设置截图回调 - snapshot
*/
func (p *LIVE) UpdateCallback(query url.Values, body string) (*UpdateCallbackResp, int, error) {
	resp := new(UpdateCallbackResp)
	statusCode, err := p.commonHandlerJson("UpdateCallback", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
通过入参messageType区分不同服务：
	查询直播推流回调
	查询直播断流回调
	查询录制回调
	查询截图回调
*/
func (p *LIVE) DescribeCallback(query url.Values, body string) (*DescribeCallbackResp, int, error) {
	resp := new(DescribeCallbackResp)
	statusCode, err := p.commonHandlerJson("DescribeCallback", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
通过入参messageType区分不同服务：
	删除直播推流回调
	删除直播断流回调
	删除录制回调
	删除截图回调
*/
func (p *LIVE) DeleteCallback(query url.Values, body string) (*DeleteCallbackResp, int, error) {
	resp := new(DeleteCallbackResp)
	statusCode, err := p.commonHandlerJson("DeleteCallback", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
域名相关的
*/
/**
添加域名
*/
func (p *LIVE) CreateDomain(query url.Values, body string) (*CreateDomainResp, int, error) {
	resp := new(CreateDomainResp)
	statusCode, err := p.commonHandlerJson("CreateDomain", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
删除域名
*/
func (p *LIVE) DeleteDomain(query url.Values, body string) (*DeleteDomainResp, int, error) {
	resp := new(DeleteDomainResp)
	statusCode, err := p.commonHandlerJson("DeleteDomain", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

/**
查询域名列表
*/
func (p *LIVE) ListDomainDetail(query url.Values, body string) (*ListDomainDetailResp, int, error) {
	resp := new(ListDomainDetailResp)
	statesCode, err := p.commonHandlerJson("ListDomainDetail", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
查询域名信息
*/
func (p *LIVE) DescribeDomain(query url.Values, body string) (*DescribeDomainResp, int, error) {
	resp := new(DescribeDomainResp)
	statesCode, err := p.commonHandlerJson("DescribeDomain", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
启用域名
*/
func (p *LIVE) EnableDomain(query url.Values, body string) (*EnableDomainResp, int, error) {
	resp := new(EnableDomainResp)
	statesCode, err := p.commonHandlerJson("EnableDomain", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
停用域名
*/
func (p *LIVE) DisableDomain(query url.Values, body string) (*DisableDomainResp, int, error) {
	resp := new(DisableDomainResp)
	statesCode, err := p.commonHandlerJson("DisableDomain", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
通过请求参数区分这两个服务
	拉流域名绑定推流域名
	删除拉流域名绑定推流域名
*/
func (p *LIVE) ManagerPullPushDomainBind(query url.Values, body string) (*ManagerPullPushDomainBindResp, int, error) {
	resp := new(ManagerPullPushDomainBindResp)
	statesCode, err := p.commonHandlerJson("ManagerPullPushDomainBind", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
鉴权相关的
*/

func (p *LIVE) UpdateAuthKey(query url.Values, body string) (*UpdateAuthKeyResp, int, error) {
	resp := new(UpdateAuthKeyResp)
	statesCode, err := p.commonHandlerJson("UpdateAuthKey", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVE) EnableAuth(query url.Values, body string) (*EnableAuthResp, int, error) {
	resp := new(EnableAuthResp)
	statesCode, err := p.commonHandlerJson("EnableAuth", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVE) DisableAuth(query url.Values, body string) (*DisableAuthResp, int, error) {
	resp := new(DisableAuthResp)
	statesCode, err := p.commonHandlerJson("DisableAuth", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVE) DescribeAuth(query url.Values, body string) (*DescribeAuthResp, int, error) {
	resp := new(DescribeAuthResp)
	statesCode, err := p.commonHandlerJson("DescribeAuth", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}


/**
流信息相关
*/
/**
禁播直播流
*/
func (p *LIVE) ForbidStream(query url.Values, body string) (*ForbidStreamResp, int, error) {
	resp := new(ForbidStreamResp)
	statesCode, err := p.commonHandlerJson("ForbidStream", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
恢复禁播直播流
*/
func (p *LIVE) ResumeStream(query url.Values, body string) (*ResumeStreamResp, int, error) {
	resp := new(ResumeStreamResp)
	statesCode, err := p.commonHandlerJson("ResumeStream", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
证书管理相关
*/
/**
1. 某个AccountID下的证书查询功能
*/
func (p *LIVE) ListCert(query url.Values, body string) (*ListCertConsoleResp, int, error) {
	resp := new(ListCertConsoleResp)
	statesCode, err := p.commonHandlerJson("ListCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
2. 上传第三方证书
*/
func (p *LIVE) CreateCert(query url.Values, body string) (*CreateCertConsoleResp, int, error) {
	resp := new(CreateCertConsoleResp)
	statesCode, err := p.commonHandlerJson("CreateCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
3. 查询证书详情 --》由DescribeCertDetail变更为DescribeCertDetailSecret接口
*/

func (p *LIVE) DescribeCertDetailSecret(query url.Values, body string) (*DescribeCertDetailConsoleResp, int, error) {
	resp := new(DescribeCertDetailConsoleResp)
	statesCode, err := p.commonHandlerJson("DescribeCertDetailSecret", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
4. 更新某个证书
*/
func (p *LIVE) UpdateCert(query url.Values, body string) (*UpdateCertConsoleResp, int, error) {
	resp := new(UpdateCertConsoleResp)
	statesCode, err := p.commonHandlerJson("UpdateCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
5. 为某个domain绑定证书
*/
func (p *LIVE) BindCert(query url.Values, body string) (*BindCertConsoleResp, int, error) {
	resp := new(BindCertConsoleResp)
	statesCode, err := p.commonHandlerJson("BindCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
6. 为这个domain解绑证书
*/
func (p *LIVE) UnbindCert(query url.Values, body string) (*UnbindCertConsoleResp, int, error) {
	resp := new(UnbindCertConsoleResp)
	statesCode, err := p.commonHandlerJson("UnbindCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
7. 删除一个证书
*/
func (p *LIVE) DeleteCert(query url.Values, body string) (*DeleteCertConsoleResp, int, error) {
	resp := new(DeleteCertConsoleResp)
	statesCode, err := p.commonHandlerJson("DeleteCert", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
防盗链相关
*/
/**
1. 创建/更新防盗链配置
*/
func (p *LIVE) UpdateReferer(query url.Values, body string) (*UpdateRefererResp, int, error) {
	resp := new(UpdateRefererResp)
	statesCode, err := p.commonHandlerJson("UpdateReferer", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
2. 关闭防盗链
*/
func (p *LIVE) DeleteReferer(query url.Values, body string) (*DeleteRefererResp, int, error) {
	resp := new(DeleteRefererResp)
	statesCode, err := p.commonHandlerJson("DeleteReferer", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
3. 查询防盗链内容
*/
func (p *LIVE) DescribeReferer(query url.Values, body string) (*DescribeRefererResp, int, error) {
	resp := new(DescribeRefererResp)
	statesCode, err := p.commonHandlerJson("DescribeReferer", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
录制相关
*/
/**
1,
*/
func (p *LIVE) CreateRecordPreset(query url.Values, body string) (*CreateRecordPresetResp, int, error) {
	resp := new(CreateRecordPresetResp)
	statesCode, err := p.commonHandlerJson("CreateRecordPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
2。
*/
func (p *LIVE) UpdateRecordPreset(query url.Values, body string) (*UpdateRecordPresetResp, int, error) {
	resp := new(UpdateRecordPresetResp)
	statesCode, err := p.commonHandlerJson("UpdateRecordPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
3.
*/

func (p *LIVE) DeleteRecordPreset(query url.Values, body string) (*DeleteRecordPresetResp, int, error) {
	resp := new(DeleteRecordPresetResp)
	statesCode, err := p.commonHandlerJson("DeleteRecordPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
6。
*/
func (p *LIVE) ListVhostRecordPreset(query url.Values, body string) (*ListVhostRecordPresetResp, int, error) {
	resp := new(ListVhostRecordPresetResp)
	statesCode, err := p.commonHandlerJson("ListVhostRecordPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
转码相关的
*/
/**
1。
*/
func (p *LIVE) CreateTranscodePreset(query url.Values, body string) (*CreateTranscodePresetResp, int, error) {
	resp := new(CreateTranscodePresetResp)
	statesCode, err := p.commonHandlerJson("CreateTranscodePreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

/**
2。
*/
func (p *LIVE) UpdateTranscodePreset(query url.Values, body string) (*UpdateTranscodePresetResp, int, error) {
	resp := new(UpdateTranscodePresetResp)
	statesCode, err := p.commonHandlerJson("UpdateTranscodePreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
/**
3。
*/
func (p *LIVE) DeleteTranscodePreset(query url.Values, body string) (*DeleteTranscodePresetResp, int, error) {
	resp := new(DeleteTranscodePresetResp)
	statesCode, err := p.commonHandlerJson("DeleteTranscodePreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}

func (p *LIVE) ListVhostTransCodePreset(query url.Values, body string) (*ListVhostTransCodePresetResp, int, error) {
	resp := new(ListVhostTransCodePresetResp)
	statesCode, err := p.commonHandlerJson("ListVhostTransCodePreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
/**
 截图相关的
*/
func (p *LIVE) CreateSnapshotPreset(query url.Values, body string) (*CreateSnapshotPresetResp, int, error) {
	resp := new(CreateSnapshotPresetResp)
	statesCode, err := p.commonHandlerJson("CreateSnapshotPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
func (p *LIVE) UpdateSnapshotPreset(query url.Values, body string) (*UpdateSnapshotPresetResp, int, error) {
	resp := new(UpdateSnapshotPresetResp)
	statesCode, err := p.commonHandlerJson("UpdateSnapshotPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
func (p *LIVE) DeleteSnapshotPreset(query url.Values, body string) (*DeleteSnapshotPresetResp, int, error) {
	resp := new(DeleteSnapshotPresetResp)
	statesCode, err := p.commonHandlerJson("DeleteSnapshotPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
func (p *LIVE) ListVhostSnapshotPreset(query url.Values, body string) (*ListVhostSnapshotPresetResp, int, error) {
	resp := new(ListVhostSnapshotPresetResp)
	statesCode, err := p.commonHandlerJson("ListVhostSnapshotPreset", query, resp, body)
	if err != nil {
		return nil, statesCode, err
	}
	return resp, statesCode, nil
}
