package mathTest

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := SecondsInRadians(thirtySeconds)

	if want != got {
		t.Errorf("Wanted %v radians, but got %v", want, got)
	}
}
