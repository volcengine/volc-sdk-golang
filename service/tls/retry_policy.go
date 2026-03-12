package tls

import (
	"math"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const (
	RetryPolicyTotalTimeoutMin      = 30 * time.Second
	RetryPolicyTotalTimeoutMax      = 15 * time.Minute
	RetryPolicyInitialIntervalMin   = 100 * time.Millisecond
	RetryPolicyInitialIntervalMax   = 30 * time.Second
	RetryPolicyMaxIntervalMin       = 1 * time.Second
	RetryPolicyMaxIntervalMax       = 1 * time.Minute
	RetryPolicyMultiplierMin        = 1.0
	RetryPolicyMultiplierMax        = 3.0
	RetryPolicyRandomizationMin     = 0.1
	RetryPolicyRandomizationMax     = 1.0
	RetryPolicyMaxAttemptsMin       = 0
	RetryPolicyMaxAttemptsMax       = 50
	defaultRetryRandomizationFactor = 0.25
	defaultRetryMultiplier          = 2.0
	defaultRetryMaxInterval         = 10 * time.Second
	defaultRetryInterval            = 500 * time.Millisecond
)

type RetryPolicy struct {
	TotalTimeout        time.Duration
	InitialInterval     time.Duration
	MaxInterval         time.Duration
	Multiplier          float64
	RandomizationFactor float64
	MaxAttempts         int
}

func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		TotalTimeout:        defaultRetryTimeout,
		InitialInterval:     defaultRetryInterval,
		MaxInterval:         defaultRetryMaxInterval,
		Multiplier:          defaultRetryMultiplier,
		RandomizationFactor: defaultRetryRandomizationFactor,
		MaxAttempts:         0,
	}
}

func (p RetryPolicy) Normalize() RetryPolicy {
	out := p

	if out.TotalTimeout <= 0 {
		out.TotalTimeout = defaultRetryTimeout
	}
	if out.TotalTimeout < RetryPolicyTotalTimeoutMin {
		out.TotalTimeout = RetryPolicyTotalTimeoutMin
	}
	if out.TotalTimeout > RetryPolicyTotalTimeoutMax {
		out.TotalTimeout = RetryPolicyTotalTimeoutMax
	}

	if out.InitialInterval <= 0 {
		out.InitialInterval = defaultRetryInterval
	}
	if out.InitialInterval < RetryPolicyInitialIntervalMin {
		out.InitialInterval = RetryPolicyInitialIntervalMin
	}
	if out.InitialInterval > RetryPolicyInitialIntervalMax {
		out.InitialInterval = RetryPolicyInitialIntervalMax
	}

	if out.MaxInterval <= 0 {
		out.MaxInterval = defaultRetryMaxInterval
	}
	if out.MaxInterval < RetryPolicyMaxIntervalMin {
		out.MaxInterval = RetryPolicyMaxIntervalMin
	}
	if out.MaxInterval > RetryPolicyMaxIntervalMax {
		out.MaxInterval = RetryPolicyMaxIntervalMax
	}
	if out.MaxInterval < out.InitialInterval {
		out.MaxInterval = out.InitialInterval
	}

	if math.IsNaN(out.Multiplier) || out.Multiplier <= 0 {
		out.Multiplier = defaultRetryMultiplier
	}
	if out.Multiplier < RetryPolicyMultiplierMin {
		out.Multiplier = RetryPolicyMultiplierMin
	}
	if out.Multiplier > RetryPolicyMultiplierMax {
		out.Multiplier = RetryPolicyMultiplierMax
	}

	if math.IsNaN(out.RandomizationFactor) || out.RandomizationFactor < 0 {
		out.RandomizationFactor = defaultRetryRandomizationFactor
	}
	if out.RandomizationFactor < RetryPolicyRandomizationMin {
		out.RandomizationFactor = RetryPolicyRandomizationMin
	}
	if out.RandomizationFactor > RetryPolicyRandomizationMax {
		out.RandomizationFactor = RetryPolicyRandomizationMax
	}

	if out.MaxAttempts < RetryPolicyMaxAttemptsMin {
		out.MaxAttempts = RetryPolicyMaxAttemptsMin
	}
	if out.MaxAttempts > RetryPolicyMaxAttemptsMax {
		out.MaxAttempts = RetryPolicyMaxAttemptsMax
	}

	return out
}

func (p RetryPolicy) ExponentialBackOff() *backoff.ExponentialBackOff {
	normalized := p.Normalize()
	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = normalized.InitialInterval
	exp.RandomizationFactor = normalized.RandomizationFactor
	exp.Multiplier = normalized.Multiplier
	exp.MaxInterval = normalized.MaxInterval
	exp.MaxElapsedTime = normalized.TotalTimeout
	return exp
}
