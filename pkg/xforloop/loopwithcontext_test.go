package xforloop

import (
	"context"
	"fmt"
	"testing"
)

func TestLoopWithContextAndRetryBackoffLimit(t *testing.T) {
	exeCount := 0
	LoopWithContextAndRetryBackoffLimit(context.Background(), func(ctx context.Context) bool {
		exeCount += 1
		return false
	}, WithRetryBackoffLimit(4), WithBaseBackoffDurationSec(0))

	if exeCount != 5 {
		t.Log(fmt.Sprintf("expect 5 but got %d", exeCount))
		t.FailNow()
	}
}
