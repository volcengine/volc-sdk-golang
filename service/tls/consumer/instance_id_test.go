package consumer

import (
	"strings"
	"testing"
)

func TestObservabilityInstanceIDIsComponentScopedAndUnique(t *testing.T) {
	first := newObservabilityInstanceID("consumer")
	second := newObservabilityInstanceID("consumer")
	if first == second || !strings.HasPrefix(first, "consumer-") || !strings.HasPrefix(second, "consumer-") {
		t.Fatalf("IDs must be component-scoped and unique: %q %q", first, second)
	}
}
