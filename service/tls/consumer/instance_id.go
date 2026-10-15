package consumer

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"sync/atomic"
	"time"
)

var instanceIDFallbackCounter uint64

// newObservabilityInstanceID returns a process-independent identifier for one
// consumer. The fallback avoids returning an empty or reusable identifier if
// the OS random source fails.
func newObservabilityInstanceID(component string) string {
	var entropy [16]byte
	if _, err := rand.Read(entropy[:]); err == nil {
		return component + "-" + hex.EncodeToString(entropy[:])
	}
	sequence := atomic.AddUint64(&instanceIDFallbackCounter, 1)
	return component + "-" + strconv.FormatInt(time.Now().UnixNano(), 16) + "-" + strconv.FormatUint(sequence, 16)
}
