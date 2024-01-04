package tls

import (
	"context"
	"math"
	"math/rand"
	"net"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

type ConditionOperation func() error

func needRetry(err *Error) bool {
	if err == nil {
		return false
	}
	if err.HTTPCode == 429 || err.HTTPCode == 500 || err.HTTPCode == 502 || err.HTTPCode == 503 {
		return true
	}
	return false
}

func RetryWithCondition(ctx context.Context, o ConditionOperation) error {
	var tryCount = 0
	var err error
	var expectedQuitTime = time.Now().Add(defaultRequestTimeout)
	for {
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
			err = o()
			tryCount += 1
			Err, ok := err.(*Error)
			if !ok {
				if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
					// 超时错误, 需要重试
				} else {
					// 预期之外的错误, 直接返回
					return err
				}
			} else {
				if !needRetry(Err) {
					// 预期之内的请求成功, 不需要重试, 计数减后直接返回
					for {
						curr := atomic.LoadInt32(&defaultRetryCounter)
						if curr <= 0 {
							break
						}
						if atomic.CompareAndSwapInt32(&defaultRetryCounter, curr, curr-1) {
							break // 操作成功，退出循环
						}
					}
					return err
				}
			}

			// 进入重试逻辑
			for {
				// 计数器+1
				curr := atomic.LoadInt32(&defaultRetryCounter)
				if curr >= defaultRetryCounterMaximum {
					break
				}
				if atomic.CompareAndSwapInt32(&defaultRetryCounter, curr, curr+1) {
					break // 操作成功，退出循环
				}
			}

			var retrySleepInterval = time.Duration(math.Floor(rand.Float64() * float64(defaultRetryCounter) * float64(defaultRetryInterval)))
			var maxSleepInterval = time.Until(expectedQuitTime)
			if retrySleepInterval > maxSleepInterval {
				retrySleepInterval = maxSleepInterval
			}
			time.Sleep(retrySleepInterval)

			if tryCount >= 5 || defaultRetryCounterMaximum <= 0 || time.Now().After(expectedQuitTime) {
				// 重试超过5次或已经超出预期超时, 直接返回最近一次请求结果
				return err
			}
		}
	}
}
