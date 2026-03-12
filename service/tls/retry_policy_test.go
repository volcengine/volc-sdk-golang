package tls

import (
	"errors"
	"io"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestRetryPolicyNormalize(t *testing.T) {
	p := RetryPolicy{
		TotalTimeout:        -1,
		InitialInterval:     0,
		MaxInterval:         0,
		Multiplier:          0,
		RandomizationFactor: -1,
		MaxAttempts:         -10,
	}
	got := p.Normalize()

	if got.TotalTimeout != defaultRetryTimeout {
		t.Fatalf("TotalTimeout mismatch: got=%v want=%v", got.TotalTimeout, defaultRetryTimeout)
	}
	if got.InitialInterval != defaultRetryInterval {
		t.Fatalf("InitialInterval mismatch: got=%v want=%v", got.InitialInterval, defaultRetryInterval)
	}
	if got.MaxInterval != defaultRetryMaxInterval {
		t.Fatalf("MaxInterval mismatch: got=%v want=%v", got.MaxInterval, defaultRetryMaxInterval)
	}
	if got.Multiplier != defaultRetryMultiplier {
		t.Fatalf("Multiplier mismatch: got=%v want=%v", got.Multiplier, defaultRetryMultiplier)
	}
	if got.RandomizationFactor != defaultRetryRandomizationFactor {
		t.Fatalf("RandomizationFactor mismatch: got=%v want=%v", got.RandomizationFactor, defaultRetryRandomizationFactor)
	}
	if got.MaxAttempts != 0 {
		t.Fatalf("MaxAttempts mismatch: got=%v want=%v", got.MaxAttempts, 0)
	}
}

func TestRetryPolicyNormalizeClamp(t *testing.T) {
	p := RetryPolicy{
		TotalTimeout:        100 * time.Hour,
		InitialInterval:     1 * time.Millisecond,
		MaxInterval:         1 * time.Millisecond,
		Multiplier:          100,
		RandomizationFactor: 100,
		MaxAttempts:         1000,
	}
	got := p.Normalize()

	if got.TotalTimeout != RetryPolicyTotalTimeoutMax {
		t.Fatalf("TotalTimeout clamp mismatch: got=%v want=%v", got.TotalTimeout, RetryPolicyTotalTimeoutMax)
	}
	if got.InitialInterval != RetryPolicyInitialIntervalMin {
		t.Fatalf("InitialInterval clamp mismatch: got=%v want=%v", got.InitialInterval, RetryPolicyInitialIntervalMin)
	}
	if got.MaxInterval != RetryPolicyMaxIntervalMin {
		t.Fatalf("MaxInterval clamp mismatch: got=%v want=%v", got.MaxInterval, RetryPolicyMaxIntervalMin)
	}
	if got.MaxInterval < got.InitialInterval {
		t.Fatalf("MaxInterval should be >= InitialInterval: InitialInterval=%v MaxInterval=%v", got.InitialInterval, got.MaxInterval)
	}
	if got.Multiplier != RetryPolicyMultiplierMax {
		t.Fatalf("Multiplier clamp mismatch: got=%v want=%v", got.Multiplier, RetryPolicyMultiplierMax)
	}
	if got.RandomizationFactor != RetryPolicyRandomizationMax {
		t.Fatalf("RandomizationFactor clamp mismatch: got=%v want=%v", got.RandomizationFactor, RetryPolicyRandomizationMax)
	}
	if got.MaxAttempts != RetryPolicyMaxAttemptsMax {
		t.Fatalf("MaxAttempts clamp mismatch: got=%v want=%v", got.MaxAttempts, RetryPolicyMaxAttemptsMax)
	}
}

func TestRetryPolicyExponentialBackOff(t *testing.T) {
	p := RetryPolicy{
		TotalTimeout:        2 * time.Second,
		InitialInterval:     200 * time.Millisecond,
		MaxInterval:         3 * time.Second,
		Multiplier:          3,
		RandomizationFactor: 0,
		MaxAttempts:         10,
	}
	normalized := p.Normalize()
	exp := p.ExponentialBackOff()
	if exp.InitialInterval != normalized.InitialInterval {
		t.Fatalf("InitialInterval mismatch: got=%v want=%v", exp.InitialInterval, normalized.InitialInterval)
	}
	if exp.MaxInterval != normalized.MaxInterval {
		t.Fatalf("MaxInterval mismatch: got=%v want=%v", exp.MaxInterval, normalized.MaxInterval)
	}
	if exp.Multiplier != normalized.Multiplier {
		t.Fatalf("Multiplier mismatch: got=%v want=%v", exp.Multiplier, normalized.Multiplier)
	}
	if exp.RandomizationFactor != normalized.RandomizationFactor {
		t.Fatalf("RandomizationFactor mismatch: got=%v want=%v", exp.RandomizationFactor, normalized.RandomizationFactor)
	}
	if exp.MaxElapsedTime != normalized.TotalTimeout {
		t.Fatalf("MaxElapsedTime mismatch: got=%v want=%v", exp.MaxElapsedTime, normalized.TotalTimeout)
	}
}

func TestNeedRetryNetworkError(t *testing.T) {
	if !needRetryNetworkError(io.EOF) {
		t.Fatalf("io.EOF should be retryable")
	}

	if !needRetryNetworkError(syscall.ECONNRESET) {
		t.Fatalf("ECONNRESET should be retryable")
	}

	dnsTemp := &net.DNSError{Err: "tmp", IsTemporary: true}
	if !needRetryNetworkError(dnsTemp) {
		t.Fatalf("temporary DNS error should be retryable")
	}

	if needRetryNetworkError(errors.New("boom")) {
		t.Fatalf("generic error should not be retryable")
	}
}
