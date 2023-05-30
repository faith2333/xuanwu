package xforloop

import (
	"context"
	"time"
)

const (
	// DefaultBackoffRetryLimit default no backoff retry limit
	DefaultBackoffRetryLimit int = 0

	DefaultBaseBackoffDurationSec = 3
)

type config struct {
	retryBackoffLimit      int
	baseBackoffDurationSec int
}

type OperationFunc func(c *config)

func WithRetryBackoffLimit(limits int) OperationFunc {
	return func(c *config) {
		c.retryBackoffLimit = limits
	}
}

func WithBaseBackoffDurationSec(sec int) OperationFunc {
	return func(c *config) {
		c.baseBackoffDurationSec = sec
	}
}

// LoopWithContextAndRetryBackoffLimit fn returns is stop or not
func LoopWithContextAndRetryBackoffLimit(ctx context.Context, fn func(ctx context.Context) bool, ops ...OperationFunc) {
	c := &config{
		retryBackoffLimit:      DefaultBackoffRetryLimit,
		baseBackoffDurationSec: DefaultBaseBackoffDurationSec,
	}

	for _, op := range ops {
		op(c)
	}

	exeCount := 0

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		exeCount += 1
		if fn(ctx) {
			return
		}

		if c.retryBackoffLimit != 0 {
			if exeCount > c.retryBackoffLimit {
				return
			}
			time.Sleep(time.Duration(c.baseBackoffDurationSec*exeCount) * time.Second)
		} else {
			time.Sleep(time.Duration(c.baseBackoffDurationSec) * time.Second)
		}

	}
}
