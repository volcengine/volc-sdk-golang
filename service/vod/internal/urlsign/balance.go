package url_sign

import (
	"sync/atomic"
)

type Balancer interface {
	NewPicker([]*CDN) Picker
}

type Picker interface {
	Pick(...string) (*CDN, error)
}

type RRBalancer struct {
}

var _ Balancer = (*RRBalancer)(nil)

func NewRRBalancer() Balancer {
	return &RRBalancer{}
}

func (b *RRBalancer) NewPicker(cdns []*CDN) Picker {
	return &rrPicker{
		cdns: cdns,
		size: uint32(len(cdns)),
	}
}

type rrPicker struct {
	cdns []*CDN
	size uint32
	next uint32
}

var _ Picker = (*rrPicker)(nil)

func (p *rrPicker) Pick(...string) (*CDN, error) {
	if p.size <= 0 {
		return nil, ErrNoCDN
	}
	return p.cdns[(atomic.AddUint32(&p.next, 1)-1)%p.size], nil
}

