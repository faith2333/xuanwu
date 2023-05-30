package xforloop

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLoopWithContextAndRetryBackoffLimit(t *testing.T) {
	LoopWithContextAndRetryBackoffLimit(context.Background(), func(ctx context.Context) bool {
		fmt.Println(1)
		time.Sleep(time.Second * 1)
		return true
	}, WithRetryBackoffLimit(3))

	fmt.Println("OK")
}
