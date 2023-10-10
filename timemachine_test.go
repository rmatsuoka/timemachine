package timemachine

import (
	"context"
	"testing"
	"time"
)

func abs(t time.Duration) time.Duration {
	return max(t, -t)
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
		{
			cb,
			func(t time.Time) bool {
				// We expect t is current time, but we do not know exact time.
				// Thus, we check it approximately.
				return abs(time.Until(t)) < time.Microsecond
			},
		},
	}
	for _, test := range tests {
		if !test.isExpected(Now(test.ctx)) {
			t.Errorf("%v is unexpected", Now(test.ctx))
		}
	}
}
