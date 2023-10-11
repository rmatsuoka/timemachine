package timemachine

import (
	"context"
	"testing"
	"time"
)

func abs(t time.Duration) time.Duration {
	return max(t, -t)
}

// check t is approximately currnet time.
func isNow(t time.Time) bool {
	return abs(time.Until(t)) < time.Microsecond
}

func TestTimeMachine(t *testing.T) {
	cb := context.Background()
	tests := []struct {
		ctx        context.Context
		isExpected func(t time.Time) bool
	}{
		{
			NewContext(cb, func() time.Time { return time.Unix(0, 0) }),
			func(t time.Time) bool { return time.Unix(0, 0).Equal(t) },
		},
		{cb, isNow},
	}
	for _, test := range tests {
		if !test.isExpected(Now(test.ctx)) {
			t.Errorf("%v is unexpected", Now(test.ctx))
		}
	}
}

func TestClock(t *testing.T) {
	tests := []struct {
		clock      Clock
		isExpected func(t time.Time) bool
	}{
		{
			Clock{func() time.Time { return time.Unix(0, 0) }},
			func(t time.Time) bool { return time.Unix(0, 0).Equal(t) },
		},
		{Clock{}, isNow},
	}
	for _, test := range tests {
		if !test.isExpected(test.clock.Now()) {
			t.Errorf("%v is unexpected", test.clock.Now())
		}
	}
}
