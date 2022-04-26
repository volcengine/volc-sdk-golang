package tls

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type ConditionOperation func() error

const (
	FlowControlSleepTime = time.Second * 1
)

func needRetry(statusCode int32) bool {
	if statusCode == 200 {
		return false
	}

	if statusCode == 500 || statusCode == 502 || statusCode == 503 {
		return true
	}

	return false
}

func needFlowControl(statusCode int32) bool {
	return statusCode == 429
}

func RetryWithCondition(ctx context.Context, b backoff.BackOff, o ConditionOperation) error {
	ticker := backoff.NewTicker(b)
	defer ticker.Stop()
	var err error
	for {
		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
		default:
			select {
			case _, ok := <-ticker.C:
				if !ok {
					return err
				}

				err = o()

				Err, ok := err.(Error)
				if !ok {
					return err
				}

				if !needRetry(Err.HTTPCode) {
					return err
				}

				if needFlowControl(Err.HTTPCode) {
					time.Sleep(FlowControlSleepTime)
				}

			case <-ctx.Done():
				return errors.Wrapf(ctx.Err(), "stopped retrying err: %v", err)
			}
		}
	}
}
