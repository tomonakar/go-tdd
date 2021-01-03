package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(tt *testing.T) {
		counter := Counter{0}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			assertCounter(t, counter, 3)
		}
	})
}

func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
