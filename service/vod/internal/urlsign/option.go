package url_sign

import (
	"time"
)

type Option func(*DefaultURLBuilder)

func WithCDNs(cdns []*CDN) Option {
	return func(builder *DefaultURLBuilder) {
		builder.cdns = cdns
	}
}

func WithBalancer(b Balancer) Option {
	return func(builder *DefaultURLBuilder) {
		if b != nil {
			builder.balancer = b
		}
	}
}

func WithHTTPS() Option {
	return func(builder *DefaultURLBuilder) {
		builder.scheme = schemeHTTPS
	}
}

func WithTTL(ttl time.Duration) Option {
	return func(builder *DefaultURLBuilder) {
		builder.ttl = ttl
	}
}

