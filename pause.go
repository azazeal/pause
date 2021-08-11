// Package pause implements a method which suspends execution until its given
// Context is done or it's waited the given timeout.
package pause

import (
	"context"
	"time"
)

// For suspends execution until ctx is done or it's waited the given timeout.
func For(ctx context.Context, timeout time.Duration) {
	dl, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	<-dl.Done()
}
