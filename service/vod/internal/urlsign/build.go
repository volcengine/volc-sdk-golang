package url_sign

import (
	"net/url"
	"time"
)

type URLBuilder interface {
	BuildURL(string) (string, error)
}

type DefaultURLBuilder struct {
	cdns     []*CDN
	balancer Balancer
	picker   Picker
	scheme   string
	ttl      time.Duration
}

var _ URLBuilder = (*DefaultURLBuilder)(nil)

func NewDefaultURLBuilder(opts ...Option) URLBuilder {
	builder := &DefaultURLBuilder{
		balancer: NewRRBalancer(),
		scheme:   schemeHTTP,
		ttl:      ttlOneHour,
	}

	for _, opt := range opts {
		opt(builder)
	}

	builder.SetCDNPicker()
	return builder
}

func (builder *DefaultURLBuilder) BuildURL(oid string) (string, error) {
	sr, err := getStorageAndRegion(oid)
	if err != nil {
		return "", err
	}

	cdn, err := builder.PickCDN()
	if err != nil {
		return "", err
	}

	authPath := buildAuthPath(sr, oid)
	expiration := buildExpiration(builder.ttl)
	token := buildMD5(cdn.Key, authPath, expiration)
	u := &url.URL{
		Scheme: builder.scheme,
		Host:   cdn.Host,
		Path:   buildURLPath(token, expiration, authPath),
	}
	return u.String(), nil
}

func (builder *DefaultURLBuilder) SetCDNPicker() {
	builder.picker = builder.balancer.NewPicker(builder.cdns)
}

func (builder *DefaultURLBuilder) PickCDN() (*CDN, error) {
	return builder.picker.Pick()
}

type CDN struct {
	Host string
	Key  string
}

func NewCDN(host, key string) *CDN {
	return &CDN{
		Host: host,
		Key:  key,
	}
}

