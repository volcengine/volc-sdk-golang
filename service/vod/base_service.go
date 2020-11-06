package vod

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func (p *Vod) GetPlayAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}
	if getPlayInfoToken, err := p.GetSignUrl("GetPlayInfo", query); err == nil {
		ret["GetPlayInfoToken"] = getPlayInfoToken
	} else {
		return "", err
	}

	b, _ := json.Marshal(ret)
	return base64.StdEncoding.EncodeToString(b), nil
}
