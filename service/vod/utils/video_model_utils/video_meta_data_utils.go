package video_model_utils

import (
	json "github.com/json-iterator/go"
)



func ParseMetaDataInfoFromString(jsonStr string) (*MetaDataInfo, error) {
	return ParseMetaDataInfoFromByte([]byte(jsonStr))
}

func ParseMetaDataInfoFromByte(jsonByte []byte) (*MetaDataInfo, error) {
	var respData *MetaDataInfo
	err := json.Unmarshal(jsonByte, &respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func FormatMetaDataInfoToString(metaDataInfo *MetaDataInfo) (string, error) {
	resp, err := json.MarshalToString(metaDataInfo)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func GetVidFromString(jsonStr string) (string, error) {
	return GetVidFromByte([]byte(jsonStr))
}

func GetVidFromByte(jsonByte []byte) (string, error) {
	metaData, err :=  ParseMetaDataInfoFromByte(jsonByte)
	if err != nil {
		return "", err
	}
	return metaData.GetId(), nil
}

func GetVidFromMetaData(metaData *MetaDataInfo) string {
	return metaData.GetId()
}

func GetFileTypeFromString(jsonStr string) (string, error) {
	return GetFileTypeFromByte([]byte(jsonStr))
}

func GetFileTypeFromByte(jsonByte []byte) (string, error) {
	metaData, err :=  ParseMetaDataInfoFromByte(jsonByte)
	if err != nil {
		return "", err
	}
	return metaData.GetMediaType(), nil
}

func GetFileTypeFromMetaData(metaData *MetaDataInfo) string {
	return metaData.GetMediaType()
}

func GetPosterUriFromString(jsonStr string) (string, error) {
	return GetPosterUriFromByte([]byte(jsonStr))
}

func GetPosterUriFromByte(jsonByte []byte) (string, error) {
	metaData, err :=  ParseMetaDataInfoFromByte(jsonByte)
	if err != nil {
		return "", err
	}
	return metaData.GetPosterUri(), nil
}

func GetPosterUriFromMetaData(metaData *MetaDataInfo) string {
	return metaData.GetPosterUri()
}