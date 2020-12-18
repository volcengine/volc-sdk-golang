package vod

import (
	"errors"
	"net/url"
)

const (
	FORMAT_JPEG     = "jpeg"
	FORMAT_PNG      = "png"
	FORMAT_WEBP     = "webp"
	FORMAT_AWEBP    = "awebp"
	FORMAT_GIF      = "gif"
	FORMAT_HEIC     = "heic"
	FORMAT_ORIGINAL = "image"

	HTTP  = "http"
	HTTPS = "https"

	KEY_SIG = "sig"

	VOD_TPL_OBJ         = "tplv-vod-obj"
	VOD_TPL_NOOP        = "tplv-vod-noop"
	VOD_TPL_RESIZE      = "tplv-vod-rs"
	VOD_TPL_CENTER_CROP = "tplv-vod-cc"
	VOD_TPL_SMART_CROP  = "tplv-vod-cs"
	VOD_TPL_SIG         = "tplv-bd-sig"
)

var (
	ErrKvSig = errors.New("Input kv already has sig query")
)

type option struct {
	isHttps bool
	format  string
	sigKey  string
	tpl     string
	w, h    int
	kv      url.Values
}

type OptionFun func(*option)

func WithHttps() OptionFun {
	return func(opt *option) {
		opt.isHttps = true
	}
}

func WithFormat(format string) OptionFun {
	return func(opt *option) {
		opt.format = format
	}
}

func WithSig(key string) OptionFun {
	return func(opt *option) {
		opt.sigKey = key
	}
}

//kv used in querystring,the key Must not use "sig"
func WithKV(kv url.Values) OptionFun {
	return func(opt *option) {
		opt.kv = kv
	}
}

func WithVodTplObj() OptionFun {
	return func(opt *option) {
		opt.tpl = VOD_TPL_OBJ
	}
}

func WithVodNoop() OptionFun {
	return func(opt *option) {
		opt.tpl = VOD_TPL_NOOP
	}
}

func WithVodTplCenterCrop(width, height int) OptionFun {
	return func(opt *option) {
		opt.tpl = VOD_TPL_CENTER_CROP
		opt.w = width
		opt.h = height
	}
}

func WithVodTplSmartCrop(width, height int) OptionFun {
	return func(opt *option) {
		opt.tpl = VOD_TPL_SMART_CROP
		opt.w = width
		opt.h = height
	}
}

func WithVodTplResize(width, height int) OptionFun {
	return func(opt *option) {
		opt.tpl = VOD_TPL_RESIZE
		opt.w = width
		opt.h = height
	}
}
