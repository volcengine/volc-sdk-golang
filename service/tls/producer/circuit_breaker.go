package producer

import (
	"sync"
	"time"
)

type circuitState int

const (
	circuitStateClosed circuitState = iota
	circuitStateOpen
	circuitStateHalfOpen
)

type failureController struct {
	config CircuitBreakerConfig

	lock                sync.Mutex
	state               circuitState
	openedAt            time.Time
	requests            int
	failures            int
	consecutiveFailures int
	halfOpenInFlight    int
	halfOpenSuccesses   int
}

func normalizeCircuitBreakerConfig(config CircuitBreakerConfig) CircuitBreakerConfig {
	if config.MinRequests <= 0 {
		config.MinRequests = 100
	}
	if config.FailureRatio <= 0 || config.FailureRatio > 1 {
		config.FailureRatio = 0.8
	}
	if config.ConsecutiveFailures <= 0 {
		config.ConsecutiveFailures = 20
	}
	if config.OpenDuration <= 0 {
		config.OpenDuration = 30 * time.Second
	}
	if config.HalfOpenMaxRequests <= 0 {
		config.HalfOpenMaxRequests = 5
	}
	return config
}

func newFailureController(config *Config) *failureController {
	if config == nil || !config.CircuitBreaker.Enabled {
		return nil
	}
	return &failureController{config: config.CircuitBreaker}
}

func (controller *failureController) enabled() bool {
	return controller != nil && controller.config.Enabled
}

func (controller *failureController) allowSend() (bool, bool) {
	if !controller.enabled() {
		return true, false
	}
	controller.lock.Lock()
	defer controller.lock.Unlock()

	switch controller.state {
	case circuitStateClosed:
		return true, false
	case circuitStateOpen:
		if time.Since(controller.openedAt) < controller.config.OpenDuration {
			return false, false
		}
		controller.toHalfOpenLocked()
		fallthrough
	case circuitStateHalfOpen:
		if controller.halfOpenInFlight >= controller.config.HalfOpenMaxRequests {
			return false, false
		}
		controller.halfOpenInFlight++
		return true, true
	default:
		return true, false
	}
}

func (controller *failureController) rejectNewData() bool {
	if !controller.enabled() {
		return false
	}
	controller.lock.Lock()
	defer controller.lock.Unlock()

	switch controller.state {
	case circuitStateOpen:
		if time.Since(controller.openedAt) < controller.config.OpenDuration {
			return true
		}
		controller.toHalfOpenLocked()
		return false
	case circuitStateHalfOpen:
		return controller.halfOpenInFlight >= controller.config.HalfOpenMaxRequests
	default:
		return false
	}
}

func (controller *failureController) allowEnqueue() (bool, bool) {
	if !controller.enabled() {
		return true, false
	}
	controller.lock.Lock()
	defer controller.lock.Unlock()

	switch controller.state {
	case circuitStateClosed:
		return true, false
	case circuitStateOpen:
		if time.Since(controller.openedAt) < controller.config.OpenDuration {
			return false, false
		}
		controller.toHalfOpenLocked()
		fallthrough
	case circuitStateHalfOpen:
		if controller.halfOpenInFlight >= controller.config.HalfOpenMaxRequests {
			return false, false
		}
		controller.halfOpenInFlight++
		return true, true
	default:
		return true, false
	}
}

func (controller *failureController) afterSend(failure classifiedFailure, success bool, permitCount int) {
	if !controller.enabled() {
		return
	}
	controller.lock.Lock()
	defer controller.lock.Unlock()

	if controller.state == circuitStateHalfOpen {
		if permitCount <= 0 {
			return
		}
		if permitCount >= controller.halfOpenInFlight {
			controller.halfOpenInFlight = 0
		} else {
			controller.halfOpenInFlight -= permitCount
		}
	}
	if success {
		controller.onSuccessLocked()
		return
	}
	controller.onFailureLocked(failure)
}

func (controller *failureController) releasePermit() {
	controller.releasePermits(1)
}

func (controller *failureController) releasePermits(count int) {
	if !controller.enabled() {
		return
	}
	controller.lock.Lock()
	defer controller.lock.Unlock()

	if controller.state == circuitStateHalfOpen && count > 0 {
		if count >= controller.halfOpenInFlight {
			controller.halfOpenInFlight = 0
			return
		}
		controller.halfOpenInFlight -= count
	}
}

func (controller *failureController) onSuccessLocked() {
	switch controller.state {
	case circuitStateHalfOpen:
		controller.halfOpenSuccesses++
		if controller.halfOpenSuccesses >= controller.config.HalfOpenMaxRequests {
			controller.toClosedLocked()
		}
	case circuitStateClosed:
		controller.requests++
		controller.consecutiveFailures = 0
	}
}

func (controller *failureController) onFailureLocked(failure classifiedFailure) {
	if failure.Kind == FailureKindCircuitOpen || failure.Kind == FailureKindLocalResource || !failure.Retryable {
		return
	}
	switch controller.state {
	case circuitStateHalfOpen:
		controller.toOpenLocked()
	case circuitStateClosed:
		controller.requests++
		controller.failures++
		controller.consecutiveFailures++
		if controller.consecutiveFailures >= controller.config.ConsecutiveFailures ||
			(controller.requests >= controller.config.MinRequests &&
				float64(controller.failures)/float64(controller.requests) >= controller.config.FailureRatio) {
			controller.toOpenLocked()
		}
	}
}

func (controller *failureController) toOpenLocked() {
	controller.state = circuitStateOpen
	controller.openedAt = time.Now()
	controller.halfOpenInFlight = 0
	controller.halfOpenSuccesses = 0
}

func (controller *failureController) toHalfOpenLocked() {
	controller.state = circuitStateHalfOpen
	controller.halfOpenInFlight = 0
	controller.halfOpenSuccesses = 0
}

func (controller *failureController) toClosedLocked() {
	controller.state = circuitStateClosed
	controller.requests = 0
	controller.failures = 0
	controller.consecutiveFailures = 0
	controller.halfOpenInFlight = 0
	controller.halfOpenSuccesses = 0
}
