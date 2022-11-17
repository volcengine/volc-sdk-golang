package live

import (
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

func (p *LIVE) CreateCaster(query url.Values, body string) (*CreateCasterResp, int, error) {
	resp := new(CreateCasterResp)
	statusCode, err := p.commonHandlerJson("CreateCaster", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) ListCasters(query url.Values, body string) (*ListCastersResp, int, error) {
	resp := new(ListCastersResp)
	statusCode, err := p.commonHandlerJson("ListCasters", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) GetCasterResource(query url.Values, body string) (*GetCasterResourceResp, int, error) {
	resp := new(GetCasterResourceResp)
	statusCode, err := p.commonHandlerJson("GetCasterInfo", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	resp.ResponseMetadata.Action = "GetCasterResource"
	return resp, statusCode, nil
}

func (p *LIVE) GetCasterLayout(query url.Values, body string) (*GetCasterLayoutResp, int, error) {
	resp := new(GetCasterLayoutResp)
	statusCode, err := p.commonHandlerJson("GetCasterInfo", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	resp.ResponseMetadata.Action = "GetCasterLayout"
	return resp, statusCode, nil
}

func (p *LIVE) GetCasterConfig(query url.Values, body string) (*GetCasterConfigResp, int, error) {
	resp := new(GetCasterConfigResp)
	statusCode, err := p.commonHandlerJson("GetCasterInfo", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	resp.ResponseMetadata.Action = "GetCasterConfig"
	return resp, statusCode, nil
}

func (p *LIVE) StartCaster(query url.Values, body string) (*StartCasterResp, int, error) {
	resp := new(StartCasterResp)
	statusCode, err := p.commonHandlerJson("StartCaster", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) StopCaster(query url.Values, body string) (*StopCasterResp, int, error) {
	resp := new(StopCasterResp)
	statusCode, err := p.commonHandlerJson("StopCaster", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) CreateCasterResource(query url.Values, body string) (*CreateCasterResourceResp, int, error) {
	resp := new(CreateCasterResourceResp)
	statusCode, err := p.commonHandlerJson("CreateCasterResource", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) RemoveCasterResource(query url.Values, body string) (*RemoveCasterResourceResp, int, error) {
	resp := new(RemoveCasterResourceResp)
	statusCode, err := p.commonHandlerJson("RemoveCasterResource", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) UpdateCasterLayout(query url.Values, body string) (*UpdateCasterLayoutResp, int, error) {
	resp := new(UpdateCasterLayoutResp)
	statusCode, err := p.commonHandlerJson("UpdateCasterLayout", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) DeleteCasterLayout(query url.Values, body string) (*DeleteCasterLayoutResp, int, error) {
	resp := new(DeleteCasterLayoutResp)
	statusCode, err := p.commonHandlerJson("DeleteCasterLayout", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) CreateCasterResourceImage(query url.Values, body string) (*CreateCasterResourceImageResp, int, error) {
	resp := new(CreateCasterResourceImageResp)
	statusCode, err := p.commonHandlerJson("CreateCasterResourceImage", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) DeleteCasterResourceImage(query url.Values, body string) (*DeleteCasterResourceImageResp, int, error) {
	resp := new(DeleteCasterResourceImageResp)
	statusCode, err := p.commonHandlerJson("DeleteCasterResourceImage", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) CreateCasterResourceImages(query url.Values, body string) (*CreateCasterResourceImagesResp, int, error) {
	resp := new(CreateCasterResourceImagesResp)
	statusCode, err := p.commonHandlerJson("CreateCasterResourceImages", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) UpdateCasterConfig(query url.Values, body string) (*UpdateCasterConfigResp, int, error) {
	resp := new(UpdateCasterConfigResp)
	statusCode, err := p.commonHandlerJson("UpdateCasterConfig", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) CreateCasterResourceOPED(query url.Values, body string) (*CreateCasterResourceOPEDResp, int, error) {
	resp := new(CreateCasterResourceOPEDResp)
	statusCode, err := p.commonHandlerJson("CreateCasterResourceOPED", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) DeleteCasterResourceOPED(query url.Values, body string) (*DeleteCasterResourceOPEDResp, int, error) {
	resp := new(DeleteCasterResourceOPEDResp)
	statusCode, err := p.commonHandlerJson("DeleteCasterResourceOPED", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) SwitchCasterResource(query url.Values, body string) (*SwitchCasterResourceResp, int, error) {
	resp := new(SwitchCasterResourceResp)
	statusCode, err := p.commonHandlerJson("SwitchCasterResource", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) SwitchCasterResourceImage(query url.Values, body string) (*SwitchCasterResourceImageResp, int, error) {
	resp := new(SwitchCasterResourceImageResp)
	statusCode, err := p.commonHandlerJson("SwitchCasterResourceImage", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) SwitchCasterResourceOPED(query url.Values, body string) (*SwitchCasterResourceOPEDResp, int, error) {
	resp := new(SwitchCasterResourceOPEDResp)
	statusCode, err := p.commonHandlerJson("SwitchCasterResourceOPED", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) StartCasterOPEDArrange(query url.Values, body string) (*StartCasterOPEDArrangeResp, int, error) {
	resp := new(StartCasterOPEDArrangeResp)
	statusCode, err := p.commonHandlerJson("StartCasterOPEDArrange", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) SwitchCasterLayout(query url.Values, body string) (*SwitchCasterLayoutResp, int, error) {
	resp := new(SwitchCasterLayoutResp)
	statusCode, err := p.commonHandlerJson("SwitchCasterLayout", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) CopyCasterPVWToPGM(query url.Values, body string) (*CopyCasterPVWToPGMResp, int, error) {
	resp := new(CopyCasterPVWToPGMResp)
	statusCode, err := p.commonHandlerJson("CopyCasterPVWToPGM", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) GetCasterResourceVodInfo(query url.Values, body string) (*GetCasterResourceVodInfoResp, int, error) {
	resp := new(GetCasterResourceVodInfoResp)
	statusCode, err := p.commonHandlerJson("GetCasterResourceVodInfo", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) GetCasterArrange(query url.Values, body string) (*GetCasterArrangeResp, int, error) {
	resp := new(GetCasterArrangeResp)
	statusCode, err := p.commonHandlerJson("GetCasterInfo", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	resp.ResponseMetadata.Action = "GetCasterArrange"
	return resp, statusCode, nil
}

func (p *LIVE) CreateCasterArrange(query url.Values, body string) (*CreateCasterArrangeResp, int, error) {
	resp := new(CreateCasterArrangeResp)
	statusCode, err := p.commonHandlerJson("CreateCasterArrange", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) UpdateCasterArrange(query url.Values, body string) (*UpdateCasterArrangeResp, int, error) {
	resp := new(UpdateCasterArrangeResp)
	statusCode, err := p.commonHandlerJson("UpdateCasterArrange", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) DeleteCasterArrange(query url.Values, body string) (*DeleteCasterArrangeResp, int, error) {
	resp := new(DeleteCasterArrangeResp)
	statusCode, err := p.commonHandlerJson("DeleteCasterArrange", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) UpdateCasterImageCollection(query url.Values, body string) (*UpdateCasterImageCollectionResp, int, error) {
	resp := new(UpdateCasterImageCollectionResp)
	statusCode, err := p.commonHandlerJson("UpdateCasterImageCollection", query, resp, body)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *LIVE) GetPlayerAuthWithExpiredTime() (*SecurityToken2, error) {
	inlinePolicy := new(base.Policy)
	actions := []string{"live:GetCasterPlayerInfo"}
	resources := make([]string, 0)
	statement := base.NewAllowStatement(actions, resources)
	inlinePolicy.Statement = append(inlinePolicy.Statement, statement)
	if resp, err := p.Client.SignSts2(inlinePolicy, time.Hour*24); err != nil {
		return nil, err
	} else {
		return &SecurityToken2{
			AccessKeyID:     resp.AccessKeyID,
			SecretAccessKey: resp.SecretAccessKey,
			SessionToken:    resp.SessionToken,
			ExpiredTime:     resp.ExpiredTime,
			CurrentTime:     resp.CurrentTime,
		}, nil
	}
}

type SecurityToken2 struct {
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string
	SessionToken    string
	ExpiredTime     string
	CurrentTime     string
}
