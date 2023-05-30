package safego

import (
	"fmt"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic catched")
			t.FailNow()
		}
	}()

	Go(func() {
		panic("Panic Happened")
	})

	time.Sleep(time.Second * 1)
}
