package xtime

import (
	"testing"
	"time"
)

func TestStringTime(t *testing.T) {
	formatStr := StringTime(time.Date(2007, 1, 23, 1, 12, 32, 0, time.Local))
	if formatStr != "2007-01-23 01:12:32" {
		t.Logf("expect: 2007-01-23 01:01:01, but got: %s", formatStr)
		t.FailNow()
	}
}
