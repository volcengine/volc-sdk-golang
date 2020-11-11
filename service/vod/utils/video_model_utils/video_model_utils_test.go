package video_model_utils

import (
	"context"
	"reflect"
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestComposeVideoInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		infoStr     string
		params      *ComposePlayInfoWithFilter
		vodInstance *vod.Vod
	}
	type testStruct struct {
		name    string
		args    args
		want    *VideoInfo
		wantErr bool
	}
	vodIns := vod.NewInstance()
	info := `{
		"id": "v0c2c369007abu04ru8riko30uo9n73g",
		"media_type": "video",
		"poster_uri": "tos-boe-v-7a56cb/37d22ff1eaa142e3b1c868ab4db0cdc5",
		"original_stream": {
			"meta": {
				"file_id": "v0c2c369007abu04ru8riko30uo9n73g",
				"codec_type": "h264",
				"size": 465782,
				"duration": 2.2,
				"format_type": "mp4",
				"width": 720,
				"height": 1280,
				"bitrate": 1693752,
				"definition": "720p",
				"quality": "original",
				"md5": "f12d72ca14c7133f9afa270046a064a5"
			},
			"play_uri": "tos-boe-v-7a56cb/4cd3847372484ff59f051053415df485",
			"encrypt_info": {},
			"base_range_info": {},
			"check_info": {}
		},
		"static_streams": [
			{
				"meta": {
					"file_id": "25bae1cb73784839bb49fedd77da0ee5",
					"codec_type": "h265",
					"size": 69474,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 360,
					"height": 640,
					"bitrate": 246908,
					"definition": "360p",
					"quality": "normal",
					"fps": 30,
					"md5": "608e0fc708a9da18116283358a1f8807"
				},
				"play_uri": "tos-boe-v-7a56cb/f53e7583255c4c8bb58fdb0c86524f90",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-4032-392a|d:0-34736-4796,34737-69473-11e6"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "395aeef42ec841d19ab753bdfc7c8706",
					"codec_type": "h264",
					"size": 143782,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 360,
					"height": 640,
					"bitrate": 510997,
					"definition": "360p",
					"quality": "normal",
					"fps": 30,
					"md5": "e30d7382973656ffeb6aeaadf60b6c89"
				},
				"play_uri": "tos-boe-v-7a56cb/18bf671bf9ee44b8b884738a8878f251",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-3813-4bb1|d:0-71890-7729,71891-143781-c109"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "bf018b978ae24739ae08d38717d19a24",
					"codec_type": "h264",
					"size": 226001,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 480,
					"height": 854,
					"bitrate": 803202,
					"definition": "480p",
					"quality": "normal",
					"fps": 30,
					"md5": "3ea31feb61820cf83598b935571dc656"
				},
				"play_uri": "tos-boe-v-7a56cb/7e2830c16cad40eda3b5c8607e855b3e",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-3861-f9bc|d:0-112999-31a7,113000-226000-8cad"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "db7d9197ff6f455a8e70e399f2c7e95f",
					"codec_type": "h265",
					"size": 105355,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 480,
					"height": 854,
					"bitrate": 374429,
					"definition": "480p",
					"quality": "normal",
					"fps": 30,
					"md5": "ba62da95ff0b610a1f79697bb8ce8760"
				},
				"play_uri": "tos-boe-v-7a56cb/8db0a117492c45db8d543c646f522f0a",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-4033-2d2f|d:0-52676-9383,52677-105354-675e"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "b978675bb9be4fe3b6d74350390b3a8c",
					"codec_type": "h264",
					"size": 423048,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 720,
					"height": 1280,
					"bitrate": 1495221,
					"definition": "720p",
					"quality": "normal",
					"fps": 30,
					"md5": "a5b5e6bf0d57c6c71f5fe4cde0d6fc26"
				},
				"play_uri": "tos-boe-v-7a56cb/de29d2cc753d4c71b8e0afe870fc7bff",
				"encrypt_info": {
					"encrypt": true,
					"kid": "5f81824a31e944e34303f18900b6a8ca",
					"spade_a": "lbwb8GaIKc9gtyfOWbIV/l6FJtVqrRfka64m4lycE8xusCanpw=="
				},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-5885-549e|d:0-211523-a13d,211524-423047-a9af"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "76e96a14be2b4fcf83bbd97f55c43934",
					"codec_type": "h264",
					"size": 422301,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 720,
					"height": 1280,
					"bitrate": 1500847,
					"definition": "720p",
					"quality": "normal",
					"fps": 30,
					"md5": "2a2edbb4f2ed6c25e756fe593db685d3"
				},
				"play_uri": "tos-boe-v-7a56cb/a47f50af471d4b388817d30f2544c6d5",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-3715-a30b|d:0-211149-9e01,211150-422300-14e3"
				},
				"keyframe_alignment": "null"
			},
			{
				"meta": {
					"file_id": "4a1c36ec6fd244879210683c05f56984",
					"codec_type": "h265",
					"size": 160599,
					"duration": 2.251,
					"format_type": "mp4",
					"width": 720,
					"height": 1280,
					"bitrate": 570764,
					"definition": "720p",
					"quality": "normal",
					"fps": 30,
					"md5": "fcd2d8a4e8e210aaf3355df32ab71c91"
				},
				"play_uri": "tos-boe-v-7a56cb/8e5e74d9e57f4e389967fb171241b70a",
				"encrypt_info": {},
				"base_range_info": {},
				"check_info": {
					"check_info": "c:0-4033-5713|d:0-80298-294d,80299-160598-3d66"
				},
				"keyframe_alignment": "null"
			}
		],
		"dash_streams": [
			{
				"dynamic_video_list": [
					{
						"meta": {
							"file_id": "c7ef0308174047f6b99238aad6810016",
							"codec_type": "h264",
							"size": 269490,
							"duration": 2.251,
							"format_type": "dash",
							"width": 480,
							"height": 854,
							"bitrate": 742960,
							"definition": "480p",
							"quality": "normal",
							"fps": 30,
							"md5": "defee362aa532c7f60d8123e1ddcb287"
						},
						"play_uri": "tos-boe-v-7a56cb/98e4594b77d742f59061a0e2a042dd35/media-video-avc1",
						"encrypt_info": {},
						"base_range_info": {
							"init_range": "0-953",
							"index_range": "954-1021"
						},
						"check_info": {
							"check_info": "a:v0c2c369007abu04ru8riko30uo9n73g|b:0-1021-49580bff5d18e3001a1240b9ffa5bea3|c:0-1021-7719"
						},
						"keyframe_alignment": "null"
					}
				],
				"dynamic_audio_list": [
					{
						"meta": {
							"codec_type": "h264",
							"format_type": "dash",
							"bitrate": 258174,
							"quality": "normal",
							"md5": "671a405f0a0d102a3186a681fae550d5"
						},
						"play_uri": "tos-boe-v-7a56cb/98e4594b77d742f59061a0e2a042dd35/media-audio-und-mp4a",
						"encrypt_info": {},
						"base_range_info": {
							"init_range": "0-867",
							"index_range": "868-935"
						},
						"check_info": {
							"check_info": "a:v0c2c369007abu04ru8riko30uo9n73g|b:0-935-bc87283351d6addd564c59e5374f81c7|c:0-935-5e59"
						},
						"keyframe_alignment": "null"
					}
				]
			}
		]
	}`
	tests := []testStruct{
		{
			// TODO: Add test cases.
			name: "720MP4",
			args: args{
				ctx:     context.Background(),
				infoStr: info,
				params: &ComposePlayInfoWithFilter{
					FilterParams: FilterParams{
						Definition: V720P_DefinitionType,
						Format:     Normal_FormatType,
					},
					NeedOriginalVideoInfo: false,
				},
				vodInstance: vodIns,
			},
		},
		{
			// TODO: Add test cases.
			name: "720MP4DRM",
			args: args{
				ctx:     context.Background(),
				infoStr: info,
				params: &ComposePlayInfoWithFilter{
					FilterParams: FilterParams{
						Definition: V720P_DefinitionType,
						Format:     MP4_FormatType,
						FileType:     "encrypt",
					},
					NeedOriginalVideoInfo: false,
				},
				vodInstance: vodIns,
			},
		},
		{
			name: "test2",
			args: args{
				ctx:     context.Background(),
				infoStr: info,
				params: &ComposePlayInfoWithFilter{
					FilterParams: FilterParams{
						Definition: V480P_DefinitionType,
						Format:     DASH_FormatType,
					},
					NeedOriginalVideoInfo: false,
				},
				vodInstance: vodIns,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := ComposeVideoInfo(tt.args.ctx, tt.args.infoStr, tt.args.params, tt.args.vodInstance)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComposeVideoInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComposeVideoInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestMetaDataInfo_CheckAdaptive(t *testing.T) {
//	type fields struct {
//		state          protoimpl.MessageState
//		sizeCache      protoimpl.SizeCache
//		unknownFields  protoimpl.UnknownFields
//		Id             string
//		MediaType      string
//		Status         int64
//		PosterUri      string
//		Thumbs         []*BigThumb
//		OriginalStream *PlayInfo
//		StaticStreams  []*PlayInfo
//		HlsStreams     []*PlayInfo
//		DashStreams    []*DynamicPlayInfo
//		BarrageMaskUrl string
//	}
//	type args struct {
//		formatType FormatType
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			info := &MetaDataInfo{
//				state:          tt.fields.state,
//				sizeCache:      tt.fields.sizeCache,
//				unknownFields:  tt.fields.unknownFields,
//				Id:             tt.fields.Id,
//				MediaType:      tt.fields.MediaType,
//				Status:         tt.fields.Status,
//				PosterUri:      tt.fields.PosterUri,
//				Thumbs:         tt.fields.Thumbs,
//				OriginalStream: tt.fields.OriginalStream,
//				StaticStreams:  tt.fields.StaticStreams,
//				HlsStreams:     tt.fields.HlsStreams,
//				DashStreams:    tt.fields.DashStreams,
//				BarrageMaskUrl: tt.fields.BarrageMaskUrl,
//			}
//			if got := info.CheckAdaptive(tt.args.formatType); got != tt.want {
//				t.Errorf("CheckAdaptive() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMetaDataInfo_getDashVideoStreams(t *testing.T) {
//	type fields struct {
//		state          protoimpl.MessageState
//		sizeCache      protoimpl.SizeCache
//		unknownFields  protoimpl.UnknownFields
//		Id             string
//		MediaType      string
//		Status         int64
//		PosterUri      string
//		Thumbs         []*BigThumb
//		OriginalStream *PlayInfo
//		StaticStreams  []*PlayInfo
//		HlsStreams     []*PlayInfo
//		DashStreams    []*DynamicPlayInfo
//		BarrageMaskUrl string
//	}
//	type args struct {
//		params      *ComposePlayInfoWithFilter
//		vodInstance *vod.Vod
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *DynamicVideo
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			info := &MetaDataInfo{
//				state:          tt.fields.state,
//				sizeCache:      tt.fields.sizeCache,
//				unknownFields:  tt.fields.unknownFields,
//				Id:             tt.fields.Id,
//				MediaType:      tt.fields.MediaType,
//				Status:         tt.fields.Status,
//				PosterUri:      tt.fields.PosterUri,
//				Thumbs:         tt.fields.Thumbs,
//				OriginalStream: tt.fields.OriginalStream,
//				StaticStreams:  tt.fields.StaticStreams,
//				HlsStreams:     tt.fields.HlsStreams,
//				DashStreams:    tt.fields.DashStreams,
//				BarrageMaskUrl: tt.fields.BarrageMaskUrl,
//			}
//			got, err := info.getDashVideoStreams(tt.args.params, tt.args.vodInstance)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getDashVideoStreams() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getDashVideoStreams() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMetaDataInfo_getHlsVideoStreams(t *testing.T) {
//	type fields struct {
//		state          protoimpl.MessageState
//		sizeCache      protoimpl.SizeCache
//		unknownFields  protoimpl.UnknownFields
//		Id             string
//		MediaType      string
//		Status         int64
//		PosterUri      string
//		Thumbs         []*BigThumb
//		OriginalStream *PlayInfo
//		StaticStreams  []*PlayInfo
//		HlsStreams     []*PlayInfo
//		DashStreams    []*DynamicPlayInfo
//		BarrageMaskUrl string
//	}
//	type args struct {
//		params      *ComposePlayInfoWithFilter
//		vodInstance *vod.Vod
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *DynamicVideo
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			info := &MetaDataInfo{
//				state:          tt.fields.state,
//				sizeCache:      tt.fields.sizeCache,
//				unknownFields:  tt.fields.unknownFields,
//				Id:             tt.fields.Id,
//				MediaType:      tt.fields.MediaType,
//				Status:         tt.fields.Status,
//				PosterUri:      tt.fields.PosterUri,
//				Thumbs:         tt.fields.Thumbs,
//				OriginalStream: tt.fields.OriginalStream,
//				StaticStreams:  tt.fields.StaticStreams,
//				HlsStreams:     tt.fields.HlsStreams,
//				DashStreams:    tt.fields.DashStreams,
//				BarrageMaskUrl: tt.fields.BarrageMaskUrl,
//			}
//			got, err := info.getHlsVideoStreams(tt.args.params, tt.args.vodInstance)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getHlsVideoStreams() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getHlsVideoStreams() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestMetaDataInfo_getStaticVideoStreams(t *testing.T) {
//	type fields struct {
//		state          protoimpl.MessageState
//		sizeCache      protoimpl.SizeCache
//		unknownFields  protoimpl.UnknownFields
//		Id             string
//		MediaType      string
//		Status         int64
//		PosterUri      string
//		Thumbs         []*BigThumb
//		OriginalStream *PlayInfo
//		StaticStreams  []*PlayInfo
//		HlsStreams     []*PlayInfo
//		DashStreams    []*DynamicPlayInfo
//		BarrageMaskUrl string
//	}
//	type args struct {
//		params      *ComposePlayInfoWithFilter
//		vodInstance *vod.Vod
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []*Video
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			info := &MetaDataInfo{
//				state:          tt.fields.state,
//				sizeCache:      tt.fields.sizeCache,
//				unknownFields:  tt.fields.unknownFields,
//				Id:             tt.fields.Id,
//				MediaType:      tt.fields.MediaType,
//				Status:         tt.fields.Status,
//				PosterUri:      tt.fields.PosterUri,
//				Thumbs:         tt.fields.Thumbs,
//				OriginalStream: tt.fields.OriginalStream,
//				StaticStreams:  tt.fields.StaticStreams,
//				HlsStreams:     tt.fields.HlsStreams,
//				DashStreams:    tt.fields.DashStreams,
//				BarrageMaskUrl: tt.fields.BarrageMaskUrl,
//			}
//			got, err := info.getStaticVideoStreams(tt.args.params, tt.args.vodInstance)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getStaticVideoStreams() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getStaticVideoStreams() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_checkStaticStreamAdaptive(t *testing.T) {
//	type args struct {
//		infoList []*PlayInfo
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := checkStaticStreamAdaptive(tt.args.infoList); got != tt.want {
//				t.Errorf("checkStaticStreamAdaptive() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_matchCodecType(t *testing.T) {
//	type args struct {
//		reqCodecType  CodecType
//		dataCodecType string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := matchCodecType(tt.args.reqCodecType, tt.args.dataCodecType); got != tt.want {
//				t.Errorf("matchCodecType() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_matchDefinitionType(t *testing.T) {
//	type args struct {
//		reqDefinitionType  DefinitionType
//		dataDefinitionType string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := matchDefinitionType(tt.args.reqDefinitionType, tt.args.dataDefinitionType); got != tt.want {
//				t.Errorf("matchDefinitionType() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_matchQualityType(t *testing.T) {
//	type args struct {
//		reqQualityType  QualityType
//		dataQualityType string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := matchQualityType(tt.args.reqQualityType, tt.args.dataQualityType); got != tt.want {
//				t.Errorf("matchQualityType() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_parseAudioList(t *testing.T) {
//	type args struct {
//		req         []*PlayInfo
//		params      *ComposePlayInfoWithFilter
//		vodInstance *vod.Vod
//		builder     url_sign.URLBuilder
//	}
//	tests := []struct {
//		name string
//		args args
//		want []*Audio
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := parseAudioList(tt.args.req, tt.args.params, tt.args.vodInstance, tt.args.builder); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("parseAudioList() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_parseVideoList(t *testing.T) {
//	type args struct {
//		req         []*PlayInfo
//		params      *ComposePlayInfoWithFilter
//		vodInstance *vod.Vod
//		builder     url_sign.URLBuilder
//	}
//	tests := []struct {
//		name string
//		args args
//		want []*Video
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := parseVideoList(tt.args.req, tt.args.params, tt.args.vodInstance, tt.args.builder); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("parseVideoList() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
