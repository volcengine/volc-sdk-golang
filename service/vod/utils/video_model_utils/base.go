package video_model_utils

const (
	VideoModelVersion int64 = 3
)

type FormatType int

const (
	Normal_FormatType    FormatType = 0 //默认值，音频返回默认m4a,视频返回默认mp4
	MP4_FormatType       FormatType = 1 //视频mp4封装格式
	Audio_M4A_FormatType FormatType = 2 //音频m4a封装格式
	M3U8_FormatType      FormatType = 4
	GIF_FormatType       FormatType = 5
	DASH_FormatType      FormatType = 6
	Audio_OGG_FormatType FormatType = 7
	FMP4_FormatType      FormatType = 8
	HLS_FormatType       FormatType = 9
)

type DefinitionType int

const (
	ALL_DefinitionType    DefinitionType = 0
	V360P_DefinitionType  DefinitionType = 1
	V480P_DefinitionType  DefinitionType = 2
	V720P_DefinitionType  DefinitionType = 3
	V1080P_DefinitionType DefinitionType = 4
	V240P_DefinitionType  DefinitionType = 5
	V540P_DefinitionType  DefinitionType = 6
	HDR_DefinitionType    DefinitionType = 7
	V420P_DefinitionType  DefinitionType = 8
	V2K_DefinitionType    DefinitionType = 9
	V4K_DefinitionType    DefinitionType = 10
)

var DefinitionTypeDescription = map[DefinitionType]string{
	ALL_DefinitionType:    "all",
	V240P_DefinitionType:  "240p",
	V360P_DefinitionType:  "360p",
	V420P_DefinitionType:  "420p",
	V480P_DefinitionType:  "480p",
	V540P_DefinitionType:  "540p",
	V720P_DefinitionType:  "720p",
	V1080P_DefinitionType: "1080p",
	HDR_DefinitionType:    "hdr",
	V2K_DefinitionType:    "2k",
	V4K_DefinitionType:    "4k",
}

type CodecType int

const (
	ALL_CodecType            CodecType = 3 // 视频h264和bytevc1等都返回，音频所有格式都返回，通过response中CodecType来区分具体的格式
	H264_CodecType           CodecType = 0 // 视频h264
	MByteVC1_CodecType       CodecType = 1 // 视频bytevc1老格式, (h264+1)
	OByteVC1_CodecType       CodecType = 2 // (h264+1)_hvc1
	Audio_OPUS_CodecType     CodecType = 4 // 音频opus格式
	Audio_AAC_CodecType      CodecType = 5 // 音频aac格式
	Audio_MP3_CodecType      CodecType = 6 // 音频mp3格式
	ByteVC1_CodecType        CodecType = 7 // 视频bytevc1
	AllWithByteVC1_CodecType CodecType = 8 // 视频h264和bytevc1
)

var CodecTypeDescription = map[CodecType]string{
	ALL_CodecType:        "all",
	H264_CodecType:       "h264",
	Audio_OPUS_CodecType: "opus",
	Audio_AAC_CodecType:  "aac",
	Audio_MP3_CodecType:  "mp3",
	ByteVC1_CodecType:    "bytevc1",

	// TODO: 下面这两个value是从tcc中获取的动态配置，映射可能有问题，待讨论
	MByteVC1_CodecType: "mbytevc1",
	OByteVC1_CodecType: "obytevc1",
}

type QualityType int

const (
	Normal_QualityType      QualityType = 0 //正常
	Low_QualityType         QualityType = 1 // 低码率
	High_QualityType        QualityType = 2 // 高码率
	Medium_QualityType      QualityType = 3
	Lower_QualityType       QualityType = 4
	Lowest_QualityType      QualityType = 5
	Higher_QualityType      QualityType = 6
	Highest_QualityType     QualityType = 7
	Veryhigh_QualityType    QualityType = 9
	Superhigh_QualityType   QualityType = 10
	All_QualityType         QualityType = 8
	Adapt_QualityType       QualityType = 11
	AdaptLow_QualityType    QualityType = 12
	AdaptLower_QualityType  QualityType = 13
	AdaptLowest_QualityType QualityType = 14
	VLadder_QualityType     QualityType = 15
	AdaptHigh_QualityType   QualityType = 16
	AdaptHigher_QualityType QualityType = 17
)

var QualityTypeDescription = map[QualityType]string{
	Normal_QualityType:    "normal",
	Low_QualityType:       "low",
	High_QualityType:      "high",
	Medium_QualityType:    "medium",
	Lower_QualityType:     "lower",
	Lowest_QualityType:    "lowest",
	Higher_QualityType:    "higher",
	Highest_QualityType:   "highest",
	All_QualityType:       "all",
	Veryhigh_QualityType:  "veryhigh",
	Superhigh_QualityType: "superhigh",
	Adapt_QualityType:     "adapt",
	AdaptLow_QualityType:  "adapt_low",
}

type UserPlatform int

const (
	UNKNOWN_UserPlatform UserPlatform = 0
	IOS_UserPlatform     UserPlatform = 1
	ANDROID_UserPlatform UserPlatform = 2
	WEB_UserPlatform     UserPlatform = 3
)

type DeviceNetwork int

const (
	UNKNOWN_DeviceNetwork DeviceNetwork = 0
	WIFI_DeviceNetwork    DeviceNetwork = 1
	MOBILE_DeviceNetwork  DeviceNetwork = 2
	NT3G_DeviceNetwork    DeviceNetwork = 3
	NT4G_DeviceNetwork    DeviceNetwork = 4
	NT5G_DeviceNetwork    DeviceNetwork = 5
	WEB_DeviceNetwork     DeviceNetwork = 6
)
