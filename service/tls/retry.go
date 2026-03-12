package tls

import (
	"context"
	"io"
	"net"
	"syscall"

	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
)

type ConditionOperation func() error

// ConditionOperationBool 返回是否需要重试以及错误
type ConditionOperationBool func() (bool, error)

func needRetry(err *Error) bool {
	if err == nil {
		return false
	}
	if err.HTTPCode == 429 || err.HTTPCode == 500 || err.HTTPCode == 502 || err.HTTPCode == 503 {
		return true
	}
	return false
}

func needRetryNetworkError(err error) bool {
	if err == nil {
		return false
	}
	var nerr net.Error
	if errors.As(err, &nerr) {
		if nerr.Timeout() {
			return true
		}
	}
	if errors.Is(err, io.EOF) {
		return true
	}
	if errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.EPIPE) ||
		errors.Is(err, syscall.ETIMEDOUT) ||
		errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EHOSTUNREACH) ||
		errors.Is(err, syscall.ENETUNREACH) {
		return true
	}
	return false
}

func needRetryError(err error) bool {
	if err == nil {
		return false
	}
	if clientErr, ok := err.(*Error); ok {
		return needRetry(clientErr)
	}
	if needRetryNetworkError(err) {
		return true
	}
	if badRespErr, ok := err.(*BadResponseError); ok {
		return badRespErr.HTTPCode == 429 || badRespErr.HTTPCode == 500 || badRespErr.HTTPCode == 502 || badRespErr.HTTPCode == 503
	}
	return false
}

// RetryWithConditionTicker 使用 backoff.NewTicker 驱动的重试实现，
// 当 backoff 计算完成或 ctx 取消时结束，返回最近一次错误。
func RetryWithConditionTicker(ctx context.Context, b backoff.BackOff, o ConditionOperationBool) error {
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()

	var err error
	var need bool
	for {
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
			select {
			case _, ok := <-ticker.C:
				if !ok {
					// backoff 已达到上限
					return err
				}
				need, err = o()
				if !need {
					return err
				}
			case <-ctx.Done():
				return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
			}
		}
	}
}
