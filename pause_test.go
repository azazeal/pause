package pause

import (
	"context"
	"math"
	"testing"
	"time"
)

const forever = time.Duration(math.MaxInt64)

func TestCanceledContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	ensure(ctx, t, time.Hour, time.Millisecond)
}

func TestTimedOutContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ensure(ctx, t, forever, time.Second<<1)
}

func TestContextCancelation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = time.AfterFunc(time.Second, cancel)

	ensure(ctx, t, forever, time.Second<<1)
}

// ensure ensures that the amount of time For(ctx, timeout) blocks for is within [-delta, delta].
func ensure(ctx context.Context, t *testing.T, timeout, delta time.Duration) {
	t.Helper()

	startedAt := time.Now()
	For(ctx, timeout)
	stoppedAt := time.Now()

	if dt := startedAt.Sub(stoppedAt); dt < -delta || dt > delta {
		t.Fatalf("max delta between %v and %v allowed is %v, but was %v", startedAt, stoppedAt, delta, dt)
	}
}
