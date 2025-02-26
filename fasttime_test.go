//go:build go1.24 && goexperiment.synctest

package fasttime_test

import (
	"testing"
	"testing/synctest"
	"time"

	"go.withmatt.com/fasttime"
)

func stubTime(t *testing.T) {
	before := fasttime.ReplaceMonotonicRoot(time.Now())
	t.Cleanup(func() { fasttime.ReplaceMonotonicRoot(before) })
}

func TestMonotonic(t *testing.T) {
	synctest.Run(func() {
		stubTime(t)

		now := fasttime.Now()

		if time.Now() != now.ToTime() {
			t.Errorf("ToTime() didn't match time.Now(), expected: %q, got %q", time.Now(), now.ToTime())
		}

		time.Sleep(10 * time.Millisecond)

		if d := fasttime.Since(now); d != 10*time.Millisecond {
			t.Errorf("delta should have been 10ms, got %s", d)
		}
	})
}

func TestCached(t *testing.T) {
	synctest.Run(func() {
		c := fasttime.NewClock(time.Second)
		defer c.Stop()

		start := c.Now()
		time.Sleep(100 * time.Millisecond)
		delta := c.Now().Sub(start)
		if delta != 0 {
			t.Errorf("time should not have advanced yet, expected 0 got: %q", delta)
		}
		time.Sleep(901 * time.Millisecond)
		delta = c.Now().Sub(start)
		if delta != time.Second {
			t.Errorf("time should have advanced 1 second, got: %q", delta)
		}

		time.Sleep(100 * time.Millisecond)
		delta = c.Now().Sub(start)
		if delta != time.Second {
			t.Errorf("time should have advanced 1 second, got: %q", delta)
		}
	})
}
