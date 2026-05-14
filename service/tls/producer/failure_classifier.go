package producer

import (
	"context"
	"errors"
	"io"
	"net"
	"syscall"

	tls "github.com/volcengine/volc-sdk-golang/service/tls"
)

const CircuitOpenException = "CircuitOpenException"

var errCircuitOpen = errors.New(CircuitOpenException)

type FailureKind int

const (
	FailureKindUnknown FailureKind = iota
	FailureKindPermanent
	FailureKindThrottled
	FailureKindServer
	FailureKindNetwork
	FailureKindLocalResource
	FailureKindCircuitOpen
)

type classifiedFailure struct {
	Kind      FailureKind
	Retryable bool
	HTTPCode  int
	ErrorCode string
	Message   string
	RequestID string
}

func classifyFailure(err error, noRetryStatusCodeMap map[int]struct{}) classifiedFailure {
	if err == nil {
		return classifiedFailure{}
	}
	if isCircuitOpenError(err) {
		return classifiedFailure{Kind: FailureKindCircuitOpen, Message: err.Error()}
	}
	if isLocalResourceError(err) {
		return classifiedFailure{Kind: FailureKindLocalResource, Message: err.Error()}
	}

	var sdkError *tls.Error
	if errors.As(err, &sdkError) {
		httpCode := int(sdkError.HTTPCode)
		failure := classifiedFailure{
			Kind:      classifyHTTPCode(httpCode),
			Retryable: httpCode == 0 || httpCode == 429 || httpCode >= 500,
			HTTPCode:  httpCode,
			ErrorCode: sdkError.Code,
			Message:   sdkError.Message,
			RequestID: sdkError.RequestID,
		}
		if _, ok := noRetryStatusCodeMap[httpCode]; ok {
			failure.Kind = FailureKindPermanent
			failure.Retryable = false
		}
		return failure
	}

	var badResponse *tls.BadResponseError
	if errors.As(err, &badResponse) {
		httpCode := badResponse.HTTPCode
		failure := classifiedFailure{
			Kind:      classifyHTTPCode(httpCode),
			Retryable: httpCode == 429 || httpCode >= 500,
			HTTPCode:  httpCode,
			Message:   badResponse.RespBody,
		}
		if _, ok := noRetryStatusCodeMap[httpCode]; ok {
			failure.Kind = FailureKindPermanent
			failure.Retryable = false
		}
		return failure
	}

	if isRetryableNetworkError(err) {
		return classifiedFailure{Kind: FailureKindNetwork, Retryable: true, Message: err.Error()}
	}

	return classifiedFailure{Kind: FailureKindPermanent, Message: err.Error()}
}

func classifyHTTPCode(httpCode int) FailureKind {
	if httpCode == 429 {
		return FailureKindThrottled
	}
	if httpCode == 0 || httpCode >= 500 {
		return FailureKindServer
	}
	return FailureKindPermanent
}

func isCircuitOpenError(err error) bool {
	return errors.Is(err, errCircuitOpen)
}

func isLocalResourceError(err error) bool {
	return errors.Is(err, ErrMemoryTimeout)
}

func isRetryableNetworkError(err error) bool {
	if err == nil {
		return false
	}
	var nerr net.Error
	if errors.As(err, &nerr) && (nerr.Timeout() || nerr.Temporary()) {
		return true
	}
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, io.EOF) {
		return true
	}
	return errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.EPIPE) ||
		errors.Is(err, syscall.ETIMEDOUT) ||
		errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EHOSTUNREACH) ||
		errors.Is(err, syscall.ENETUNREACH)
}

func newCircuitOpenError() error {
	return errCircuitOpen
}
