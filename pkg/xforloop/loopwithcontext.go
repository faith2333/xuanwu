package xforloop

import (
	"context"
	"time"
)

const (
	// DefaultBackoffRetryLimit default no backoff retry limit
	DefaultBackoffRetryLimit int = 0

	baseBackoffTimes = 3
)

type config struct {
	retryBackoffLimit int
}

type OperationFunc func(c *config)

func WithRetryBackoffLimit(limits int) OperationFunc {
	return func(c *config) {
		c.retryBackoffLimit = limits
	}
}

func LoopWithContextAndRetryBackoffLimit(ctx context.Context, fn func(ctx context.Context) bool, ops ...OperationFunc) {
	c := &config{
		DefaultBackoffRetryLimit,
	}

	for _, op := range ops {
		op(c)
	}

	retryCount := 0

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(time.Duration(baseBackoffTimes*retryCount) * time.Second)

		if !fn(ctx) {
			return
		}

		if c.retryBackoffLimit != 0 {
			retryCount++
			if retryCount >= c.retryBackoffLimit {
				return
			}
		}

	}
}
