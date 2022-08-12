package pause

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	ensure(ctx, t, forever, time.Second<<1)
}

func TestContextCancelation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	_ = time.AfterFunc(time.Second, cancel)

	ensure(ctx, t, forever, time.Second<<1)
}

func ensure(ctx context.Context, t *testing.T, timeout, delta time.Duration) {
	t.Helper()

	at := time.Now()
	For(ctx, timeout)

	require.WithinDuration(t, at, time.Now(), delta)
}
