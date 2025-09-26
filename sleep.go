package stdx

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, duration time.Duration) {
	var timer = time.NewTimer(duration)
	select {
	case <-ctx.Done():
		timer.Stop()
		return
	case <-timer.C:
		return
	}
}
