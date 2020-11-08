package video_model_utils

import (
	"context"
	"os"
	"strings"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	url_sign "github.com/volcengine/volc-sdk-golang/service/vod/internal/urlsign"
)

func ComposeVideoInfo(ctx context.Context, infoStr string, params *ComposePlayInfoWithFilter, vodInstance *vod.Vod) (*VideoInfo, error) {
	metaDataInfo, err := ParseMetaDataInfoFromString(infoStr)
	if err != nil {
		return nil, err
	}
	videoInfo := &VideoInfo{
		Status:         metaDataInfo.GetStatus(),
		Version:        VideoModelVersion,
		VideoId:        metaDataInfo.GetId(),
		EnableSsl:      params.URLParams.SSL,
		VideoDuration:  metaDataInfo.GetOriginalStream().GetMeta().GetDuration(),
		MediaType:      "",
		EnableAdaptive: metaDataInfo.CheckAdaptive(params.FilterParams.FormatType),
	}
	switch params.FilterParams.FormatType {
	case Normal_FormatType, MP4_FormatType, FMP4_FormatType:
		videoInfo.VideoList, err = metaDataInfo.getStaticVideoStreams(params, vodInstance)
		if err != nil {
			return nil, err
		}
	case M3U8_FormatType, HLS_FormatType:
		videoInfo.DynamicVideo, err = metaDataInfo.getHlsVideoStreams(params, vodInstance)
	case DASH_FormatType:
		videoInfo.DynamicVideo, err = metaDataInfo.getDashVideoStreams(params, vodInstance)
	}

	return videoInfo, nil
}

func matchCodecType(reqCodecType CodecType, dataCodecType string) bool {
	switch reqCodecType {
	case ALL_CodecType:
		if len(dataCodecType) == 0 {
			return false
		}
	case AllWithByteVC1_CodecType:
		d1, _ := CodecTypeDescription[H264_CodecType]
		d2, _ := CodecTypeDescription[ByteVC1_CodecType]
		if d1 != dataCodecType && d2 != dataCodecType {
			return false
		}

	default:
		d, ok := CodecTypeDescription[reqCodecType]
		if !ok || d != dataCodecType {
			return false
		}
	}
	return true
}

func matchDefinitionType(reqDefinitionType DefinitionType, dataDefinitionType string) bool {
	switch reqDefinitionType {
	case ALL_DefinitionType:
		if len(dataDefinitionType) == 0 {
			return false
		}
	default:
		def := strings.Split(dataDefinitionType, " ")[0]
		d, ok := DefinitionTypeDescription[reqDefinitionType]
		if !ok || d != def {
			return false
		}
	}
	return true
}

func matchQualityType(reqQualityType QualityType, dataQualityType string) bool {
	switch reqQualityType {
	case All_QualityType:
		if len(dataQualityType) == 0 {
			return false
		}
	default:
		d, ok := QualityTypeDescription[reqQualityType]
		if !ok || d != dataQualityType {
			return false
		}
	}
	return true
}

func matchFormatType(reqFormatType FormatType, dataFormatType string) bool {
	switch reqFormatType {
	case Normal_FormatType:
		if strings.ToUpper(dataFormatType) == "MP4" ||
			strings.ToUpper(dataFormatType) == "M4A" {
			return true
		}
	case MP4_FormatType:
		if strings.ToUpper(dataFormatType) == "MP4" {
			return true
		}
	case FMP4_FormatType:
		if strings.ToUpper(dataFormatType) == "FMP4" {
			return true
		}
	case M3U8_FormatType:
		if strings.ToUpper(dataFormatType) == "M3U8" {
			return true
		}
	case HLS_FormatType:
		if strings.ToUpper(dataFormatType) == "HLS" {
			return true
		}
	case DASH_FormatType:
		if strings.ToUpper(dataFormatType) == "DASH" {
			return true
		}
	case Audio_M4A_FormatType:
		if strings.ToUpper(dataFormatType) == "M4A" {
			return true
		}
	}
	return false
}

func matchStreamType(reqStreamType string, data *PlayInfo) bool {
	if reqStreamType == "" {
		reqStreamType = "normal"
	}
	//#normal普通非加密转码视频流 //#audio普通非加密转码音频流 //#encrypt加密转码视频流 //#audio_encrypt加密转码音频流 //#normal_short截断非加密转码视频 // perview预览
	switch reqStreamType {
	case "normal", "audio":
		if data.GetEncryptInfo().GetEncrypt() == false &&
			data.GetEncryptInfo().GetKid() == "" {
			if reqStreamType == "normal" {
				switch strings.ToUpper(data.GetMeta().GetFormatType()) {
				case "MP4", "FMP4", "M3U8", "HLS", "DASH", "MPD":
					return true
				}
			} else {
				switch strings.ToUpper(data.GetMeta().GetFormatType()) {
				case "M4A", "OGG":
					return true
				}
			}
		}
	case "encrypt", "audio_encrypt":
		if data.GetEncryptInfo().GetEncrypt() == true &&
			data.GetEncryptInfo().GetKid() != "" {
			if reqStreamType == "encrypt" {
				switch strings.ToUpper(data.GetMeta().GetFormatType()) {
				case "MP4", "FMP4", "M3U8", "HLS", "DASH", "MPD":
					return true
				}
			} else {
				switch strings.ToUpper(data.GetMeta().GetFormatType()) {
				case "M4A", "OGG":
					return true
				}
			}
		}
	}
	return false
}

func (info *MetaDataInfo) getStaticVideoStreams(params *ComposePlayInfoWithFilter, vodInstance *vod.Vod) ([]*Video, error) {
	list := info.GetStaticStreams()
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	var respList []*Video
	cdn := url_sign.NewCDN(os.Getenv("SIGN_HOST"), os.Getenv("SIGN_KEY"))
	builder := url_sign.NewDefaultURLBuilder(url_sign.WithCDNs([]*url_sign.CDN{cdn}))
	for _, row := range list {
		// 1. 类型
		checkFormat := matchFormatType(params.FilterParams.FormatType, row.GetMeta().GetFormatType())
		if !checkFormat {
			continue
		}
		// 2. 格式
		checkCodec := matchCodecType(params.FilterParams.CodecType, row.GetMeta().GetCodecType())
		if !checkCodec {
			continue
		}
		// 3. 清晰度
		checkDefinition := matchDefinitionType(params.FilterParams.DefinitionType, row.GetMeta().GetDefinition())
		if !checkDefinition {
			continue
		}
		// 4. 质量
		checkQuality := matchQualityType(params.FilterParams.VQuality, row.GetMeta().GetQuality())
		if !checkQuality {
			continue
		}
		// 5. 加密
		checkStream := matchStreamType(params.FilterParams.StreamType, row)
		if !checkStream {
			continue
		}

		// 6. url
		url, err := builder.BuildURL(row.GetPlayUri())
		if err != nil {
			continue
		}
		respRow := &Video{
			MainUrl:   url,
			BackupUrl: url,
			VideoMeta: &VideoMeta{
				Definition:  row.GetMeta().GetDefinition(),
				Quality:     row.GetMeta().GetQuality(),
				QualityDesc: row.GetMeta().GetQualityDesc(),
				Vtype:       row.GetMeta().GetFormatType(),
				Vwidth:      row.GetMeta().GetWidth(),
				Vheight:     row.GetMeta().GetHeight(),
				Bitrate:     row.GetMeta().GetBitrate(),
				CodecType:   row.GetMeta().GetCodecType(),
				Size:        row.GetMeta().GetSize(),
				FileId:      row.GetMeta().GetFileId(),
				Fps:         row.GetMeta().GetFps(),
			},
			EncryptInfo:   row.GetEncryptInfo(),
			BaseRangeInfo: row.GetBaseRangeInfo(),
			CheckInfo:     row.GetCheckInfo(),
		}
		respList = append(respList, respRow)
	}
	return respList, nil
}

func (info *MetaDataInfo) getHlsVideoStreams(params *ComposePlayInfoWithFilter, vodInstance *vod.Vod) (*DynamicVideo, error) {
	list := info.GetHlsStreams()
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	resp := &DynamicVideo{
		DynamicType:      "hls",
		DynamicVideoList: make([]*Video, 0),
		DynamicAudioList: make([]*Audio, 0),
	}
	cdn := url_sign.NewCDN(os.Getenv("SIGN_HOST"), os.Getenv("SIGN_KEY"))
	builder := url_sign.NewDefaultURLBuilder(url_sign.WithCDNs([]*url_sign.CDN{cdn}))
	for _, row := range list {
		// 1. 格式
		checkCodec := matchCodecType(params.FilterParams.CodecType, row.GetMeta().GetCodecType())
		if !checkCodec {
			continue
		}
		// 2. 清晰度
		checkDefinition := matchDefinitionType(params.FilterParams.DefinitionType, row.GetMeta().GetDefinition())
		if !checkDefinition {
			continue
		}
		// 3. 质量
		checkQuality := matchQualityType(params.FilterParams.VQuality, row.GetMeta().GetQuality())
		if !checkQuality {
			continue
		}
		// 4. url
		url, err := builder.BuildURL(row.GetPlayUri())
		if err != nil {
			continue
		}
		respRow := &Video{
			MainUrl:   url,
			BackupUrl: url,
			VideoMeta: &VideoMeta{
				Definition:  row.GetMeta().GetDefinition(),
				Quality:     row.GetMeta().GetQuality(),
				QualityDesc: row.GetMeta().GetQualityDesc(),
				Vtype:       row.GetMeta().GetFormatType(),
				Vwidth:      row.GetMeta().GetWidth(),
				Vheight:     row.GetMeta().GetHeight(),
				Bitrate:     row.GetMeta().GetBitrate(),
				CodecType:   row.GetMeta().GetCodecType(),
				Size:        row.GetMeta().GetSize(),
				FileId:      row.GetMeta().GetFileId(),
				Fps:         row.GetMeta().GetFps(),
			},
			EncryptInfo:   row.GetEncryptInfo(),
			BaseRangeInfo: row.GetBaseRangeInfo(),
			CheckInfo:     row.GetCheckInfo(),
		}
		resp.DynamicVideoList = append(resp.DynamicVideoList, respRow)
	}
	return resp, nil
}

func (info *MetaDataInfo) getDashVideoStreams(params *ComposePlayInfoWithFilter, vodInstance *vod.Vod) (*DynamicVideo, error) {
	list := info.GetDashStreams()
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	var respList []*DynamicVideo
	cdn := url_sign.NewCDN(os.Getenv("SIGN_HOST"), os.Getenv("SIGN_KEY"))
	builder := url_sign.NewDefaultURLBuilder(url_sign.WithCDNs([]*url_sign.CDN{cdn}))
	for _, dashlist := range list {
		resp := &DynamicVideo{
			DynamicType:      "segment_base",
			DynamicVideoList: parseVideoList(dashlist.GetDynamicVideoList(), params, vodInstance, builder),
			DynamicAudioList: parseAudioList(dashlist.GetDynamicVideoList(), params, vodInstance, builder),
		}
		respList = append(respList, resp)
	}
	if len(respList) > 0 {
		return respList[0], nil
	}
	return nil, nil
}

func parseVideoList(req []*PlayInfo, params *ComposePlayInfoWithFilter, vodInstance *vod.Vod, builder url_sign.URLBuilder) []*Video {
	var resp []*Video
	for _, row := range req {
		// 1. 格式
		checkCodec := matchCodecType(params.FilterParams.CodecType, row.GetMeta().GetCodecType())
		if !checkCodec {
			continue
		}
		// 2. 清晰度
		checkDefinition := matchDefinitionType(params.FilterParams.DefinitionType, row.GetMeta().GetDefinition())
		if !checkDefinition {
			continue
		}
		// 3. 质量
		checkQuality := matchQualityType(params.FilterParams.VQuality, row.GetMeta().GetQuality())
		if !checkQuality {
			continue
		}
		// 4. url
		url, err := builder.BuildURL(row.GetPlayUri())
		if err != nil {
			continue
		}
		respRow := &Video{
			MainUrl:   url,
			BackupUrl: url,
			VideoMeta: &VideoMeta{
				Definition:  row.GetMeta().GetDefinition(),
				Quality:     row.GetMeta().GetQuality(),
				QualityDesc: row.GetMeta().GetQualityDesc(),
				Vtype:       row.GetMeta().GetFormatType(),
				Vwidth:      row.GetMeta().GetWidth(),
				Vheight:     row.GetMeta().GetHeight(),
				Bitrate:     row.GetMeta().GetBitrate(),
				CodecType:   row.GetMeta().GetCodecType(),
				Size:        row.GetMeta().GetSize(),
				FileId:      row.GetMeta().GetFileId(),
				Fps:         row.GetMeta().GetFps(),
			},
			EncryptInfo:   row.GetEncryptInfo(),
			BaseRangeInfo: row.GetBaseRangeInfo(),
			CheckInfo:     row.GetCheckInfo(),
		}
		resp = append(resp, respRow)
	}
	return resp
}

func parseAudioList(req []*PlayInfo, params *ComposePlayInfoWithFilter, vodInstance *vod.Vod, builder url_sign.URLBuilder) []*Audio {
	var resp []*Audio
	for _, row := range req {
		// 1. 格式
		checkCodec := matchCodecType(params.FilterParams.CodecType, row.GetMeta().GetCodecType())
		if !checkCodec {
			continue
		}
		// 2. 清晰度
		checkDefinition := matchDefinitionType(params.FilterParams.DefinitionType, row.GetMeta().GetDefinition())
		if !checkDefinition {
			continue
		}
		// 3. 质量
		checkQuality := matchQualityType(params.FilterParams.VQuality, row.GetMeta().GetQuality())
		if !checkQuality {
			continue
		}
		// 4. url
		url, err := builder.BuildURL(row.GetPlayUri())
		if err != nil {
			continue
		}
		respRow := &Audio{
			MainUrl:   url,
			BackupUrl: url,
			AudioMeta: &AudioMeta{
				Definition:  row.GetMeta().GetDefinition(),
				Quality:     row.GetMeta().GetQuality(),
				QualityDesc: row.GetMeta().GetQualityDesc(),
				Atype:       row.GetMeta().GetFormatType(),
				Bitrate:     row.GetMeta().GetBitrate(),
				CodecType:   row.GetMeta().GetCodecType(),
				Size:        row.GetMeta().GetSize(),
				FileId:      row.GetMeta().GetFileId(),
			},
			EncryptInfo:   row.GetEncryptInfo(),
			BaseRangeInfo: row.GetBaseRangeInfo(),
			CheckInfo:     row.GetCheckInfo(),
		}
		resp = append(resp, respRow)
	}
	return resp
}

func (info *MetaDataInfo) CheckAdaptive(formatType FormatType) bool {
	if info == nil || len(info.GetId()) == 0 {
		return false
	}
	switch formatType {
	case Normal_FormatType, MP4_FormatType, Audio_M4A_FormatType, FMP4_FormatType:
		return checkStaticStreamAdaptive(info.GetStaticStreams())
	case DASH_FormatType:
		return true
	default:
		return false
	}
}

func checkStaticStreamAdaptive(infoList []*PlayInfo) bool {
	version := "unset"
	for _, video := range infoList {
		fileVersion := "null"
		if video != nil && video.GetKeyframeAlignment() != "" {
			fileVersion = video.GetKeyframeAlignment()
		}
		if version == "unset" {
			version = fileVersion
			continue
		}
		if version != fileVersion {
			return false
		}
	}
	if version == "unset" || version == "null" {
		return false
	}
	if len(infoList) != 0 {
		return true
	}
	return false
}

type ThumbType int

const (
	SingleThumb ThumbType = 0 //单张雪碧图
	MultiThumb  ThumbType = 1 //多张雪碧图
)

type ComposePlayInfoWithFilter struct {
	FilterParams          FilterParams
	User                  User      //能传必须传过来,我们这边需要这个信息,进行cdn的调度,ab测试等
	URLParams             URLParams // url控制参数，如https、地址类型、cdn类型、过期时间等
	NeedPreloadInfo       bool      //是否需要预加载信息
	NeedOriginalVideoInfo bool      //是否需要片源信息
}

type FilterParams struct {
	DefinitionType DefinitionType
	CodecType      CodecType
	Watermark      string // 水印名称,不明确指定,按照我们服务播放策略返回(若无策略则返回无水印视频) //# 如果明确想获取某种水印，需要明确指定 //# 如果明确想要获取无水印视频，明确传递(unwatermarked)
	StreamType     string // 视频转码类型有一下几种，按需传参 //#normal普通非加密转码视频流 //#audio普通非加密转码音频流 //#encrypt加密转码视频流 //#audio_encrypt加密转码音频流 //#normal_short截断非加密转码视频 // perview预览
	FormatType     FormatType
	VQuality       QualityType
	EncodeUserTag  string
	VLadder        string
}

type URLParams struct {
	SSL    bool
	Indate int64 // 播放地址的有效期,单位s
}

type User struct {
	DeviceID          int64
	UserID            int64
	WebID             int64
	Platform          UserPlatform
	Version           int64
	AppID             int64
	UserIP            string
	OSVersion         string
	Network           DeviceNetwork
	DeviceType        string
	WifiIdentify      string
	PlayerVersion     string
	UpdateVersionCode int64
	OpenapiAction     string
	APIUserName       string
	TopAccountID      string
	VersionCode       string
	CountryCode       string
	ChannelID         int64
	Channel           string
	IsOrderFlow       int64
	Province          string
}
