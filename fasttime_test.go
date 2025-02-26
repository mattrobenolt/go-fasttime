package fasttime_test

import (
	"testing"
	"time"

	"go.withmatt.com/fasttime"
)

func stubTime(t *testing.T) {
	before := fasttime.ReplaceMonotonicRoot(time.Now())
	t.Cleanup(func() { fasttime.ReplaceMonotonicRoot(before) })
}

func assertEqual(t *testing.T, got, expected, margin time.Duration) {
	if (got - expected).Abs() > margin {
		t.Errorf("durations not equal, got %q, expected %q", got, expected)
	}
}

func testMonotonic(t *testing.T, approx bool) {
	stubTime(t)

	var margin time.Duration
	if approx {
		margin = 3 * time.Millisecond
	}

	now := fasttime.Now()
	timeNow := time.Now()

	assertEqual(t, 0, timeNow.Sub(now.ToTime()), margin)

	time.Sleep(10 * time.Millisecond)

	assertEqual(t, fasttime.Since(now), 10*time.Millisecond, margin)
}

func testCached(t *testing.T, approx bool) {
	var margin time.Duration
	if approx {
		margin = 3 * time.Millisecond
	}

	c := fasttime.NewClock(time.Second)
	defer c.Stop()

	start := c.Now()

	time.Sleep(100 * time.Millisecond)
	assertEqual(t, c.Since(start), 0, margin)

	time.Sleep(901 * time.Millisecond)
	assertEqual(t, c.Since(start), time.Second, margin)

	time.Sleep(100 * time.Millisecond)
	assertEqual(t, c.Since(start), time.Second, margin)
}
