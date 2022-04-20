package mars

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// CreateService : 新建服务
func (d *Diff) CreateService(request *CreateServiceReq, appid int64) (response *CreateServiceResp, status int, err error) {
	response = &CreateServiceResp{}
	status, err = d.handleRequest(createService, request, appid, response)

	return response, status, err
}

// Validate : 验证传入的原始包和差分包是否能正确合成新包
func (d *Diff) Validate(request *ValidateReq) (response *AsyncResp, status int, err error) {
	response = &AsyncResp{}
	status, err = d.handleRequest(validate, request, 1, response)

	return response, status, err
}

// DeletePackages : 删除原始包
func (d *Diff) DeletePackages(request *DeletePackagesReq, appid int64) (response *DeletePackagesResp, status int, err error) {
	response = &DeletePackagesResp{}
	status, err = d.handleRequest(deletePackages, request, appid, response)

	return response, status, err
}

// DeleteService : 删除服务
func (d *Diff) DeleteService(request *DeleteServiceReq, appid int64) (response *DeleteServiceResp, status int, err error) {
	response = &DeleteServiceResp{}
	status, err = d.handleRequest(deleteService, request, appid, response)

	return response, status, err
}

// QueryPatchByService : 查看服务中的包信息
func (d *Diff) QueryPatchByService(request *QueryPatchByServiceReq, appid int64) (response *QueryPatchByServiceResp, status int, err error) {
	response = &QueryPatchByServiceResp{}
	status, err = d.handleRequest(queryPatchByService, request, appid, response)

	return response, status, err
}

// GenByCount : 根据传入的个数生成差分包
func (d *Diff) GenByCount(request *GenByCountReq, appid int64) (response *AsyncResp, status int, err error) {
	response = &AsyncResp{}
	status, err = d.handleRequest(genByCount, request, appid, response)

	return response, status, err
}

// GenByVersion : 根据传入的版本生成差分包
func (d *Diff) GenByVersion(request *GenByVersionReq, appid int64) (response *AsyncResp, status int, err error) {
	response = &AsyncResp{}
	status, err = d.handleRequest(genByVersion, request, appid, response)

	return response, status, err
}

// GenByPkg : 根据传入的原始包和新包生成差分包
func (d *Diff) GenByPkg(request *GenByPkgReq, appid int64) (response *AsyncResp, status int, err error) {
	response = &AsyncResp{}
	status, err = d.handleRequest(genByPkg, request, appid, response)

	return response, status, err
}

// CheckResponse : 查询某个任务是否完成，如果完成会返回任务的处理结果
// 当response.Code = 0时，任务处理完成，result中会有值
// 当任务为生成差分包时（response.Data.Api = 0），result会返回GenResult结构体
// 当任务为验证差分包时（response.Data.Api = 1），result会返回ValidateResult结构体
func (d *Diff) CheckResponse(request *CheckResponseReq) (response *CheckResponseResp, result interface{}, status int, err error) {
	response = &CheckResponseResp{}
	status, err = d.handleRequest(checkResponse, request, 1, response)
	// 当出现错误或response没有值或任务没有完成时，直接返回
	if err != nil || response == nil || response.Code != 0 {
		return response, nil, status, err
	}

	// 根据任务类型确认result的实际类型
	switch response.Data.Api {
	case 0:
		result = &GenResult{}
	case 1:
		result = &ValidateResult{}
	}

	// 反序列化实际处理结果
	err = json.Unmarshal(response.Data.Response, result)
	if err != nil {
		result = nil
	}

	return
}

// private helper func
func (d *Diff) handleRequest(api string, request interface{}, appid int64, resp interface{}) (int, error) {
	// json marshal
	bts, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	// do request
	respBody, status, err := d.Json(api, url.Values{"aid": []string{fmt.Sprintf("%v", appid)}}, string(bts))
	if err != nil {
		return status, err
	}

	// check status code
	if status != http.StatusOK {
		return status, errors.Wrap(fmt.Errorf("http error"), strconv.Itoa(status))
	}

	// unmarshal response
	if err := json.Unmarshal(respBody, resp); err != nil {
		return status, err
	} else {
		return status, nil
	}
}
