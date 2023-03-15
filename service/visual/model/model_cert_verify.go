package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type CertTokenRequest struct {
	ReqKey   string `json:"req_key"`
	StsToken string `json:"sts_token"`
	TosInfo  *struct {
		StsAk    string `json:"sts_ak"`
		StsSk    string `json:"sts_sk"`
		StsToken string `json:"sts_token"`
		Bucket   string `json:"bucket"`
		Endpoint string `json:"endpoint"`
		Region   string `json:"region"`
	} `json:"tos_info"`
	RefSource        string   `json:"ref_source"`
	LivenessType     string   `json:"liveness_type"`
	IdCardName       string   `json:"idcard_name"`
	IdCardNo         string   `json:"idcard_no"`
	RefImage         string   `json:"ref_image"`
	LivenessTimeout  int      `json:"liveness_timeout"`
	MotionList       []string `json:"motion_list"`
	FixedMotionList  []string `json:"fixed_motion_list"`
	MotionCount      int      `json:"motion_count"`
	MaxLivenessTrial int      `json:"max_liveness_trial"`
	CallbackInfo     *struct {
		Switch     bool   `json:"switch"`
		Block      bool   `json:"block"`
		Url        string `json:"url"`
		ClientName string `json:"client_name"`
	} `json:"callback_info"`
	ConfigID string `json:"config_id"`
}

type CertVerifyQueryRequest struct {
	ReqKey     string `json:"req_key"`
	BytedToken string `json:"byted_token"`
	OmitData   bool   `json:"omit_data"`
}

type CertConfigInitRequest struct {
	ReqKey         string `json:"req_key"`
	ConfigName     string `json:"config_name"`
	ConfigDesc     string `json:"config_desc"`
	TokenApiConfig *struct {
		RefSource    string `json:"ref_source"`
		LivenessType string `json:"liveness_type"`
		TosInfo      *struct {
			Bucket   string `json:"bucket"`
			Endpoint string `json:"endpoint"`
			Region   string `json:"region"`
		} `json:"tos_info"`
		RefImage         string   `json:"ref_image"`
		LivenessTimeout  int      `json:"liveness_timeout"`
		MotionList       []string `json:"motion_list"`
		FixedMotionList  []string `json:"fixed_motion_list"`
		MotionCount      int      `json:"motion_count"`
		MaxLivenessTrial int      `json:"max_liveness_trial"`
		CallbackInfo     *struct {
			Switch     bool   `json:"switch"`
			Block      bool   `json:"block"`
			Url        string `json:"url"`
			ClientName string `json:"client_name"`
		} `json:"callback_info"`
	} `json:"token_api_config"`
	H5Config *struct {
		RedirectUrl      string `json:"redirectUrl"`
		Source           string `json:"source"`
		Bucket           string `json:"bucket"`
		Endpoint         string `json:"endpoint"`
		Region           string `json:"region"`
		Type             string `json:"type"`
		ShowResult       string `json:"showResult"`
		ProtocolName     string `json:"protocolName"`
		ProtocolLink     string `json:"protocolLink"`
		DemoteType       int64  `json:"demoteType"`
		StyleConf        string `json:"styleConf"`
		ProtocolNav      string `json:"protocolNav"`
		ProtocolNavTitle string `json:"protocolNavTitle"`
	} `json:"h5_config"`
}

type CertConfigGetRequest struct {
	ReqKey   string `json:"req_key"`
	ConfigId bool   `json:"config_id"`
}

// 响应参数

type CertTokenData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	BytedToken       string   `json:"byted_token"`
	ClientConfig     string   `json:"client_config"`
}

type CertTokenResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *CertTokenData         `json:"data"`
}

type CertVerifyQueryData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	Result           bool     `json:"result"`
	Images           *struct {
		ImageBest string `json:"image_best"`
		ImageEnv  string `json:"image_env"`
	} `json:"images"`
	Video             string `json:"video"`
	SourceCompDetails *struct {
		Score      float64 `json:"score"`
		Thresholds *struct {
			E3 float64 `json:"1e-3"`
			E4 float64 `json:"1e-4"`
			E5 float64 `json:"1e-5"`
			E6 float64 `json:"1e-6"`
		} `json:"thresholds"`
	} `json:"source_comp_details"`
	TosData *struct {
		Bucket       string `json:"bucket"`
		ImageEnvKey  string `json:"image_env_key"`
		ImageBestKey string `json:"image_best_key"`
		VideoKey     string `json:"video_key"`
		CertDataKey  string `json:"cert_data_key"`
	} `json:"tos_data"`
	VerifyAlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"verify_algorithm_base_resp"`
	VerifyReqMeasureInfo *struct {
		MeasureType string `json:"measure_type"`
		Value       int    `json:"value"`
	} `json:"verify_req_measure_info"`
}

type CertVerifyQueryResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *CertVerifyQueryData   `json:"data"`
}

type CertConfigInitData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	ConfigId         string   `json:"config_id"`
}

type CertConfigInitResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *CertConfigInitData    `json:"data"`
}

type CertConfigGetData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	ClientConfig     string   `json:"client_config"`
	H5Config         *struct {
		Bucket           string `json:"bucket"`
		DemoteType       int    `json:"demoteType"`
		Endpoint         string `json:"endpoint"`
		ProtocolLink     string `json:"protocolLink"`
		ProtocolName     string `json:"protocolName"`
		ProtocolNav      string `json:"protocolNav"`
		ProtocolNavTitle string `json:"protocolNavTitle"`
		RedirectUrl      string `json:"redirectUrl"`
		Region           string `json:"region"`
		ShowResult       string `json:"showResult"`
		Source           string `json:"source"`
		StyleConf        string `json:"styleConf"`
		Type             string `json:"type"`
	} `json:"h5_config"`
	ReqMeasureInfo *struct {
		MeasureType string `json:"measure_type"`
		Value       int    `json:"value"`
	} `json:"req_measure_info"`
	TokenApiConfig *struct {
		CallbackInfo *struct {
			Switch     bool   `json:"switch"`
			Block      bool   `json:"block"`
			Url        string `json:"url"`
			ClientName string `json:"client_name"`
		} `json:"callback_info"`
		FixedMotionList  []string `json:"fixed_motion_list"`
		LivenessTimeout  int      `json:"liveness_timeout"`
		LivenessType     string   `json:"liveness_type"`
		MaxLivenessTrial int      `json:"max_liveness_trial"`
		MotionCount      int      `json:"motion_count"`
		MotionList       []string `json:"motion_list"`
		RefSource        string   `json:"ref_source"`
		TosInfo          *struct {
			Bucket   string `json:"bucket"`
			Endpoint string `json:"endpoint"`
			Region   string `json:"region"`
		} `json:"tos_info"`
	} `json:"token_api_config"`
}

type CertConfigGetResult struct {
	ResponseMetadata *base.ResponseMetadata `json:",omitempty"`
	RequestId        string                 `json:"request_id"`
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	Status           int                    `json:"status"`
	TimeElapsed      string                 `json:"time_elapsed"`
	Data             *CertConfigGetData     `json:"data"`
}
